package presentation

import (
	"fmt"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"

	"karedoro/application"
	"karedoro/domain"
)

// EbitenUIApp はebitenuiを使用したアプリケーション
type EbitenUIApp struct {
	ui             *ebitenui.UI
	sessionService *application.SessionService
	audioService   domain.AudioPlayer
	buttonContainer *widget.Container
	progressBar    *widget.ProgressBar
	
	// ボタンラベル追跡用
	buttonLabels   map[*widget.Button]string
}

// NewEbitenUIApp は新しいebitenuiアプリケーションを作成
func NewEbitenUIApp(services *application.Services) *EbitenUIApp {
	app := &EbitenUIApp{
		sessionService: services.Session,
		audioService:   services.Audio,
		buttonLabels:   make(map[*widget.Button]string),
	}
	
	app.buildUI()
	return app
}

func (a *EbitenUIApp) buildUI() {
	// ルートコンテナ
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(40),
			widget.RowLayoutOpts.Padding(widget.Insets{Top: 120, Bottom: 120, Left: 100, Right: 100}),
		)),
	)

	// プログレスバー
	a.progressBar = widget.NewProgressBar(
		widget.ProgressBarOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				MaxWidth: 400,
			}),
			widget.WidgetOpts.MinSize(400, 30),
		),
		widget.ProgressBarOpts.Images(
			&widget.ProgressBarImage{
				Idle:  image.NewNineSliceColor(color.RGBA{60, 60, 60, 255}),
				Hover: image.NewNineSliceColor(color.RGBA{80, 80, 80, 255}),
			},
			&widget.ProgressBarImage{
				Idle:  image.NewNineSliceColor(color.RGBA{0, 180, 0, 255}),
				Hover: image.NewNineSliceColor(color.RGBA{0, 200, 0, 255}),
			},
		),
		widget.ProgressBarOpts.Values(0, 1500, 0), // 25分 = 1500秒
	)
	rootContainer.AddChild(a.progressBar)

	// ボタンコンテナ
	a.buttonContainer = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(20),
		)),
	)
	rootContainer.AddChild(a.buttonContainer)

	// 初期ボタンを作成
	a.setupInitialButtons()

	// UIを構築
	a.ui = &ebitenui.UI{
		Container: rootContainer,
	}
}

func (a *EbitenUIApp) setupInitialButtons() {
	startWorkBtn := a.createButton("Start Work Session", func() {
		log.Printf("Start Work Session clicked")
		err := a.sessionService.StartWorkSession()
		if err != nil {
			log.Printf("Failed to start work session: %v", err)
			return
		}
		a.audioService.PlayStartSound()
		a.updateButtons()
	})
	a.buttonContainer.AddChild(startWorkBtn)
}

func (a *EbitenUIApp) createButton(buttonText string, clickHandler func()) *widget.Button {
	// 明確な色でボタンを作成
	button := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
			widget.WidgetOpts.MinSize(250, 60),
		),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    image.NewNineSliceColor(color.RGBA{70, 130, 180, 255}),  // 鋼鉄青
			Hover:   image.NewNineSliceColor(color.RGBA{100, 149, 237, 255}), // コーンフラワー青
			Pressed: image.NewNineSliceColor(color.RGBA{65, 105, 225, 255}),  // 王室青
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			log.Printf("Button clicked: %s", buttonText)
			clickHandler()
		}),
	)
	
	// ボタンラベルを記録
	a.buttonLabels[button] = buttonText
	
	return button
}

func (a *EbitenUIApp) updateButtons() {
	log.Printf("Updating buttons...")
	
	// 既存のボタンを削除
	for button := range a.buttonLabels {
		delete(a.buttonLabels, button)
	}
	a.buttonContainer.RemoveChildren()

	session := a.sessionService.GetSession()
	sessionState := session.GetState()
	isPaused := session.IsSessionPaused()

	log.Printf("Session state: %v, isPaused: %v", sessionState, isPaused)

	switch sessionState {
	case domain.WorkSession:
		if isPaused {
			resumeBtn := a.createButton("Resume Work", func() {
				log.Printf("Resume Work clicked")
				err := a.sessionService.ResumeSession()
				if err != nil {
					log.Printf("Failed to resume session: %v", err)
					return
				}
				a.updateButtons()
			})
			a.buttonContainer.AddChild(resumeBtn)
		} else {
			pauseBtn := a.createButton("Pause Work", func() {
				log.Printf("Pause Work clicked")
				err := a.sessionService.PauseSession()
				if err != nil {
					log.Printf("Failed to pause session: %v", err)
					return
				}
				a.updateButtons()
			})
			a.buttonContainer.AddChild(pauseBtn)
		}
	case domain.BreakSession:
		if isPaused {
			resumeBtn := a.createButton("Resume Break", func() {
				log.Printf("Resume Break clicked")
				err := a.sessionService.ResumeSession()
				if err != nil {
					log.Printf("Failed to resume session: %v", err)
					return
				}
				a.updateButtons()
			})
			a.buttonContainer.AddChild(resumeBtn)
		} else {
			pauseBtn := a.createButton("Pause Break", func() {
				log.Printf("Pause Break clicked")
				err := a.sessionService.PauseSession()
				if err != nil {
					log.Printf("Failed to pause session: %v", err)
					return
				}
				a.updateButtons()
			})
			a.buttonContainer.AddChild(pauseBtn)
		}
	case domain.Idle:
		sessionType := session.GetSessionType()
		
		if sessionType == domain.Work {
			// 作業セッション終了後
			startBreakBtn := a.createButton("Start Break", func() {
				log.Printf("Start Break clicked")
				err := a.sessionService.StartBreakSession()
				if err != nil {
					log.Printf("Failed to start break session: %v", err)
					return
				}
				a.audioService.PlayStartSound()
				a.updateButtons()
			})
			a.buttonContainer.AddChild(startBreakBtn)
			
			skipBreakBtn := a.createButton("Skip Break", func() {
				log.Printf("Skip Break clicked")
				err := a.sessionService.StartWorkSession()
				if err != nil {
					log.Printf("Failed to start work session: %v", err)
					return
				}
				a.audioService.PlayStartSound()
				a.updateButtons()
			})
			a.buttonContainer.AddChild(skipBreakBtn)
		} else {
			// 休憩セッション終了後または初期状態
			startWorkBtn := a.createButton("Start Work", func() {
				log.Printf("Start Work clicked")
				err := a.sessionService.StartWorkSession()
				if err != nil {
					log.Printf("Failed to start work session: %v", err)
					return
				}
				a.audioService.PlayStartSound()
				a.updateButtons()
			})
			a.buttonContainer.AddChild(startWorkBtn)
		}
	}
	
	log.Printf("Buttons updated, container has %d children", len(a.buttonContainer.Children()))
}

func (a *EbitenUIApp) updateProgressBar() {
	session := a.sessionService.GetSession()
	sessionState := session.GetState()
	remaining := session.GetTimeRemaining()

	switch sessionState {
	case domain.WorkSession:
		total := float64(25 * 60) // 25分
		elapsed := total - remaining.Seconds()
		progress := int(elapsed)
		if progress > int(total) {
			progress = int(total)
		}
		if progress < 0 {
			progress = 0
		}
		a.progressBar.SetCurrent(progress)
	case domain.BreakSession:
		total := float64(5 * 60) // 5分
		elapsed := total - remaining.Seconds()
		progress := int(elapsed)
		if progress > int(total) {
			progress = int(total)
		}
		if progress < 0 {
			progress = 0
		}
		a.progressBar.SetCurrent(progress)
	default:
		a.progressBar.SetCurrent(0)
	}
}

func (a *EbitenUIApp) Update() error {
	a.ui.Update()
	
	// セッションサービスを更新
	a.sessionService.Update()
	
	// プログレスバーを更新
	a.updateProgressBar()
	
	return nil
}

func (a *EbitenUIApp) Draw(screen *ebiten.Image) {
	// 背景を暗い色で塗りつぶし
	screen.Fill(color.RGBA{25, 25, 25, 255})
	
	// ebitenuiを描画
	a.ui.Draw(screen)
	
	// タイマー情報とボタンラベルをオーバーレイで描画
	a.drawOverlayText(screen)
}

func (a *EbitenUIApp) drawOverlayText(screen *ebiten.Image) {
	session := a.sessionService.GetSession()
	remaining := session.GetTimeRemaining()
	minutes := int(remaining.Minutes())
	seconds := int(remaining.Seconds()) % 60
	timerText := fmt.Sprintf("%02d:%02d", minutes, seconds)
	
	sessionState := session.GetState()
	statusText := ""
	switch sessionState {
	case domain.WorkSession:
		if session.IsSessionPaused() {
			statusText = "Work Session (Paused)"
		} else {
			statusText = "Work Session"
		}
	case domain.BreakSession:
		if session.IsSessionPaused() {
			statusText = "Break Session (Paused)"
		} else {
			statusText = "Break Session"
		}
	default:
		statusText = "Ready to start"
	}
	
	// タイマーテキストを描画（上部中央）
	timerBounds := text.BoundString(basicfont.Face7x13, timerText)
	timerX := screen.Bounds().Dx()/2 - timerBounds.Dx()/2
	timerY := 60
	text.Draw(screen, timerText, basicfont.Face7x13, timerX, timerY, color.White)
	
	// ステータステキストを描画
	statusBounds := text.BoundString(basicfont.Face7x13, statusText)
	statusX := screen.Bounds().Dx()/2 - statusBounds.Dx()/2
	statusY := 90
	text.Draw(screen, statusText, basicfont.Face7x13, statusX, statusY, color.RGBA{200, 200, 200, 255})
	
	// ボタンラベルを各ボタンの上に描画
	buttonY := 220 // ボタンエリアの開始位置
	for button, label := range a.buttonLabels {
		if button.GetWidget().Disabled {
			continue
		}
		
		labelBounds := text.BoundString(basicfont.Face7x13, label)
		labelX := screen.Bounds().Dx()/2 - labelBounds.Dx()/2
		text.Draw(screen, label, basicfont.Face7x13, labelX, buttonY, color.White)
		buttonY += 80 // 次のボタンの位置
	}
}

func (a *EbitenUIApp) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}