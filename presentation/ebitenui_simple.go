package presentation

import (
	"fmt"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2"

	"karedoro/application"
	"karedoro/domain"
)

// SimpleEbitenUIApp はebitenuiを使用したシンプルなアプリケーション
type SimpleEbitenUIApp struct {
	ui             *ebitenui.UI
	sessionService *application.SessionService
	audioService   domain.AudioPlayer
	buttonContainer *widget.Container
}

// NewSimpleEbitenUIApp は新しいebitenuiアプリケーションを作成
func NewSimpleEbitenUIApp(services *application.Services) *SimpleEbitenUIApp {
	app := &SimpleEbitenUIApp{
		sessionService: services.Session,
		audioService:   services.Audio,
	}
	
	app.buildUI()
	return app
}

func (a *SimpleEbitenUIApp) buildUI() {
	// ルートコンテナをRowLayoutで作成（シンプルに）
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(20),
			widget.RowLayoutOpts.Padding(widget.Insets{Top: 50, Bottom: 50, Left: 50, Right: 50}),
		)),
	)

	// プログレスバー（シンプル版）
	progressBar := widget.NewProgressBar(
		widget.ProgressBarOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				MaxWidth: 400,
			}),
		),
		widget.ProgressBarOpts.Images(
			&widget.ProgressBarImage{
				Idle:  image.NewNineSliceColor(color.RGBA{50, 50, 50, 255}),
				Hover: image.NewNineSliceColor(color.RGBA{60, 60, 60, 255}),
			},
			&widget.ProgressBarImage{
				Idle:  image.NewNineSliceColor(color.RGBA{0, 255, 0, 255}),
				Hover: image.NewNineSliceColor(color.RGBA{0, 200, 0, 255}),
			},
		),
		widget.ProgressBarOpts.Values(0, 1500, 0), // 25分 = 1500秒
	)
	rootContainer.AddChild(progressBar)

	// ボタンコンテナ
	a.buttonContainer = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
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

func (a *SimpleEbitenUIApp) setupInitialButtons() {
	// "Start Work Session"ボタン
	startWorkBtn := a.createButton("Start Work Session", func() {
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

func (a *SimpleEbitenUIApp) createButton(text string, clickHandler func()) *widget.Button {
	return widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    image.NewNineSliceColor(color.RGBA{100, 100, 100, 255}),
			Hover:   image.NewNineSliceColor(color.RGBA{120, 120, 120, 255}),
			Pressed: image.NewNineSliceColor(color.RGBA{80, 80, 80, 255}),
		}),
		widget.ButtonOpts.TextPadding(widget.Insets{Left: 30, Right: 30, Top: 15, Bottom: 15}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			clickHandler()
		}),
	)
}

func (a *SimpleEbitenUIApp) updateButtons() {
	// 既存のボタンを削除
	a.buttonContainer.RemoveChildren()

	session := a.sessionService.GetSession()
	sessionState := session.GetState()
	isPaused := session.IsSessionPaused()

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
			a.buttonContainer.AddChild(startWorkBtn) // "Skip Break"的な意味
		} else {
			// 休憩セッション終了後または初期状態
			a.buttonContainer.AddChild(startWorkBtn)
		}
	}
}

func (a *SimpleEbitenUIApp) Update() error {
	a.ui.Update()
	
	// セッションサービスを更新
	a.sessionService.Update()
	
	// ボタン状態を必要に応じて更新
	// Note: 毎フレーム更新は重いので、状態変化時のみに最適化可能
	
	return nil
}

func (a *SimpleEbitenUIApp) Draw(screen *ebiten.Image) {
	// 背景を黒で塗りつぶし
	screen.Fill(color.RGBA{20, 20, 20, 255})
	
	// タイマー情報を手動描画（ebitenuiのText問題を回避）
	session := a.sessionService.GetSession()
	remaining := session.GetTimeRemaining()
	minutes := int(remaining.Minutes())
	seconds := int(remaining.Seconds()) % 60
	timerText := fmt.Sprintf("%02d:%02d", minutes, seconds)
	
	// 簡単なテキスト描画（ebitenの基本機能使用）
	// NOTE: 実際のプロダクトではより良いフォント描画が必要
	_ = timerText // 将来の実装で使用予定
	
	a.ui.Draw(screen)
}

func (a *SimpleEbitenUIApp) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}