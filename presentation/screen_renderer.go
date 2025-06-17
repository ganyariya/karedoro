package presentation

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"karedoro/domain"
)

type ScreenRenderer struct{}

func NewScreenRenderer() *ScreenRenderer {
	return &ScreenRenderer{}
}

func (sr *ScreenRenderer) DrawMainScreen(screen *ebiten.Image, session *domain.Session, buttonManager *ButtonManager) {
	switch session.GetState() {
	case domain.WorkSession:
		sr.drawWorkSession(screen, session)
	case domain.BreakSession:
		sr.drawBreakSession(screen, session)
	case domain.Idle:
		sr.drawIdleScreen(screen, buttonManager)
	}
}

func (sr *ScreenRenderer) DrawFullscreenOverlay(screen *ebiten.Image, session *domain.Session, buttonManager *ButtonManager) {
	screenWidth, screenHeight := ebiten.WindowSize()
	
	// 強制的な赤い背景で注意を引く
	screen.Fill(ForceRedBackground)
	
	var message string
	if session.GetSessionType() == domain.Work {
		message = WorkSessionEndMessage
	} else {
		message = BreakSessionEndMessage
	}
	
	// メッセージを大きく強調表示
	messageX := screenWidth/2 - len(message)*TextCharWidthLg
	messageY := screenHeight/3
	
	// 背景の強調ボックス
	boxWidth := len(message)*8 + MessageBoxPadding
	boxHeight := MessageBoxHeight
	boxX := screenWidth/2 - boxWidth/2
	boxY := messageY - 20
	drawRect(screen, boxX, boxY, boxWidth, boxHeight, ForceYellowBox)
	drawBorder(screen, boxX, boxY, boxWidth, boxHeight, WhiteBorder, MessageBoxBorderWidth)
	
	ebitenutil.DebugPrintAt(screen, message, messageX, messageY)
	
	// 警告メッセージを追加
	warningMsg := "YOU CANNOT CONTINUE UNTIL YOU CHOOSE!"
	warningX := screenWidth/2 - len(warningMsg)*TextCharWidth
	warningY := screenHeight/2 - TextLineHeight
	ebitenutil.DebugPrintAt(screen, warningMsg, warningX, warningY)
	
	buttonManager.DrawButtons(screen)
}

func (sr *ScreenRenderer) drawWorkSession(screen *ebiten.Image, session *domain.Session) {
	sr.drawSessionState(screen, session, WorkSessionColor, WorkingText)
}

func (sr *ScreenRenderer) drawBreakSession(screen *ebiten.Image, session *domain.Session) {
	sr.drawSessionState(screen, session, BreakSessionColor, BreakText)
}

func (sr *ScreenRenderer) drawSessionState(screen *ebiten.Image, session *domain.Session, sessionColor color.Color, statusText string) {
	remaining := session.GetTimeRemaining()
	screenWidth, screenHeight := ebiten.WindowSize()
	
	screen.Fill(sessionColor)
	
	timerText := fmt.Sprintf("%02d:%02d", int(remaining.Minutes()), int(remaining.Seconds())%60)
	// Draw timer background for better visibility
	timerX := screenWidth/2 - TimerOffsetX
	timerY := screenHeight/2 - TimerOffsetY
	drawRect(screen, timerX, timerY, TimerBoxWidth, TimerBoxHeight, BlackShadow)
	ebitenutil.DebugPrintAt(screen, timerText, screenWidth/2-TimerOffsetX+10, screenHeight/2-TimerOffsetY+10)
	
	// Draw progress bar
	sr.drawProgressBar(screen, session.GetProgress(), screenWidth, screenHeight)
	
	if session.IsSessionPaused() {
		ebitenutil.DebugPrintAt(screen, PausedText, screenWidth/2-len(PausedText)*TextCharWidth, screenHeight/2-TextLineHeight)
		ebitenutil.DebugPrintAt(screen, ResumeInstructionText, screenWidth/2-len(ResumeInstructionText)*TextCharWidth, screenHeight/2-20)
	} else {
		ebitenutil.DebugPrintAt(screen, statusText, screenWidth/2-len(statusText)*TextCharWidth, screenHeight/2-TextLineHeight)
		ebitenutil.DebugPrintAt(screen, PauseInstructionText, screenWidth/2-len(PauseInstructionText)*TextCharWidth, screenHeight/2-20)
	}
}

func (sr *ScreenRenderer) drawIdleScreen(screen *ebiten.Image, buttonManager *ButtonManager) {
	screenWidth, screenHeight := ebiten.WindowSize()
	ebitenutil.DebugPrintAt(screen, IdleScreenMessage, screenWidth/2-len(IdleScreenMessage)*TextCharWidth, screenHeight/2-IdleMessageOffset)
	
	buttonManager.DrawButtons(screen)
}

func (sr *ScreenRenderer) drawProgressBar(screen *ebiten.Image, progress float64, screenWidth, screenHeight int) {
	barX := screenWidth/2 - ProgressBarWidth/2
	barY := screenHeight/2 + ProgressBarOffsetY
	
	// Draw progress bar background
	drawRect(screen, barX, barY, ProgressBarWidth, ProgressBarHeight, ProgressBarBg)
	
	// Draw progress bar fill
	progressWidth := int(float64(ProgressBarWidth) * progress)
	if progressWidth > 0 {
		drawRect(screen, barX, barY, progressWidth, ProgressBarHeight, ProgressBarFill)
	}
	
	// Draw progress bar border
	drawBorder(screen, barX, barY, ProgressBarWidth, ProgressBarHeight, ProgressBarBorder, 1)
}