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

// FixedEbitenUIApp はebitenuiを使用した修正版アプリケーション
type FixedEbitenUIApp struct {
	ui             *ebitenui.UI
	sessionService *application.SessionService
	audioService   domain.AudioPlayer
	buttonContainer *widget.Container
	progressBar    *widget.ProgressBar
}

// NewFixedEbitenUIApp は新しいebitenuiアプリケーションを作成
func NewFixedEbitenUIApp(services *application.Services) *FixedEbitenUIApp {
	app := &FixedEbitenUIApp{
		sessionService: services.Session,
		audioService:   services.Audio,
	}
	
	app.buildUI()
	return app
}

func (a *FixedEbitenUIApp) buildUI() {
	// ルートコンテナをRowLayoutで作成
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(30),
			widget.RowLayoutOpts.Padding(widget.Insets{Top: 100, Bottom: 100, Left: 100, Right: 100}),
		)),
	)

	// プログレスバー
	a.progressBar = widget.NewProgressBar(
		widget.ProgressBarOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				MaxWidth: 400,
			}),
			widget.WidgetOpts.MinSize(400, 20),
		),
		widget.ProgressBarOpts.Images(
			&widget.ProgressBarImage{
				Idle:  image.NewNineSliceColor(color.RGBA{100, 100, 100, 255}),
				Hover: image.NewNineSliceColor(color.RGBA{120, 120, 120, 255}),
			},
			&widget.ProgressBarImage{
				Idle:  image.NewNineSliceColor(color.RGBA{0, 255, 0, 255}),
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
			widget.RowLayoutOpts.Spacing(15),
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

func (a *FixedEbitenUIApp) setupInitialButtons() {
	// "Start Work Session"ボタン
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

func (a *FixedEbitenUIApp) createButton(buttonText string, clickHandler func()) *widget.Button {
	// ボタンの最小サイズを設定
	button := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
			widget.WidgetOpts.MinSize(200, 50),
		),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    image.NewNineSliceColor(color.RGBA{70, 130, 180, 255}),  // SteelBlue
			Hover:   image.NewNineSliceColor(color.RGBA{100, 149, 237, 255}), // CornflowerBlue
			Pressed: image.NewNineSliceColor(color.RGBA{65, 105, 225, 255}),  // RoyalBlue
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			log.Printf("Button clicked: %s", buttonText)
			clickHandler()
		}),
	)
	
	return button
}

func (a *FixedEbitenUIApp) updateButtons() {
	log.Printf("Updating buttons...")
	// 既存のボタンを削除
	a.buttonContainer.RemoveChildren()

	session := a.sessionService.GetSession()
	sessionState := session.GetState()
	isPaused := session.IsSessionPaused()

	log.Printf("Session state: %v, isPaused: %v", sessionState, isPaused)

	switch sessionState {
	case domain.WorkSession:
		if isPaused {
			resumeBtn := a.createButton("Resume Work", func() {
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
		// セッション終了後のボタン配置
		startBreakBtn := a.createButton("Start Break", func() {
			err := a.sessionService.StartBreakSession()
			if err != nil {
				log.Printf("Failed to start break session: %v", err)
				return
			}
			a.audioService.PlayStartSound()
			a.updateButtons()
		})
		
		startWorkBtn := a.createButton("Start Work", func() {
			err := a.sessionService.StartWorkSession()
			if err != nil {
				log.Printf("Failed to start work session: %v", err)
				return
			}
			a.audioService.PlayStartSound()
			a.updateButtons()
		})
		
		// 前のセッションタイプに応じてボタンを表示
		sessionType := session.GetSessionType()
		if sessionType == domain.Work {
			// 作業セッション終了後
			a.buttonContainer.AddChild(startBreakBtn)
			a.buttonContainer.AddChild(startWorkBtn)
		} else {
			// 休憩セッション終了後または初期状態
			a.buttonContainer.AddChild(startWorkBtn)
		}
	}
	
	log.Printf("Buttons updated, container has %d children", len(a.buttonContainer.Children()))
}

func (a *FixedEbitenUIApp) updateProgressBar() {
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

func (a *FixedEbitenUIApp) Update() error {
	a.ui.Update()
	
	// セッションサービスを更新
	a.sessionService.Update()
	
	// プログレスバーを更新
	a.updateProgressBar()
	
	return nil
}

func (a *FixedEbitenUIApp) Draw(screen *ebiten.Image) {
	// 背景を黒で塗りつぶし
	screen.Fill(color.RGBA{30, 30, 30, 255})
	
	// ebitenuiを描画
	a.ui.Draw(screen)
	
	// タイマー情報をオーバーレイで手動描画
	session := a.sessionService.GetSession()
	remaining := session.GetTimeRemaining()
	minutes := int(remaining.Minutes())
	seconds := int(remaining.Seconds()) % 60
	timerText := fmt.Sprintf("%02d:%02d", minutes, seconds)
	
	sessionState := session.GetState()
	statusText := ""
	switch sessionState {
	case domain.WorkSession:
		statusText = "Work Session"
	case domain.BreakSession:
		statusText = "Break Session"
	default:
		statusText = "Ready to start"
	}
	
	// テキストを手動描画（上部中央）
	bounds := text.BoundString(basicfont.Face7x13, timerText)
	timerX := screen.Bounds().Dx()/2 - bounds.Dx()/2
	timerY := 50
	text.Draw(screen, timerText, basicfont.Face7x13, timerX, timerY, color.White)
	
	// ステータステキストを描画（タイマーの下）
	statusBounds := text.BoundString(basicfont.Face7x13, statusText)
	statusX := screen.Bounds().Dx()/2 - statusBounds.Dx()/2
	statusY := 80
	text.Draw(screen, statusText, basicfont.Face7x13, statusX, statusY, color.RGBA{200, 200, 200, 255})
}

func (a *FixedEbitenUIApp) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}