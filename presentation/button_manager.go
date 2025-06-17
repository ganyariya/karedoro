package presentation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"karedoro/application"
)

type ButtonManager struct {
	buttons []Button
}

func NewButtonManager() *ButtonManager {
	return &ButtonManager{
		buttons: make([]Button, 0),
	}
}

func (bm *ButtonManager) SetupMainButtons(screenWidth, screenHeight int, sessionService *application.SessionService) {
	bm.buttons = []Button{
		{
			X: screenWidth/2 - ButtonWidth/2,
			Y: screenHeight/2 - ButtonHeight - ButtonPadding,
			W: ButtonWidth,
			H: ButtonHeight,
			Text: StartWorkButtonText,
			Action: func() {
				sessionService.StartWorkSession()
			},
		},
		{
			X: screenWidth/2 - ButtonWidth/2,
			Y: screenHeight/2 + ButtonPadding,
			W: ButtonWidth,
			H: ButtonHeight,
			Text: StartBreakButtonText,
			Action: func() {
				sessionService.StartBreakSession()
			},
		},
	}
}

func (bm *ButtonManager) SetupEndOfWorkButtons(screenWidth, screenHeight int, sessionService *application.SessionService) {
	bm.buttons = []Button{
		{
			X: screenWidth/2 - ButtonWidth/2,
			Y: screenHeight/2 - ButtonHeight - ButtonPadding,
			W: ButtonWidth,
			H: ButtonHeight,
			Text: StartBreakButtonText,
			Action: func() {
				sessionService.StartBreakSession()
			},
		},
		{
			X: screenWidth/2 - ButtonWidth/2,
			Y: screenHeight/2 + ButtonPadding,
			W: ButtonWidth,
			H: ButtonHeight,
			Text: SkipBreakButtonText,
			Action: func() {
				sessionService.StartWorkSession()
			},
		},
	}
}

func (bm *ButtonManager) SetupEndOfBreakButtons(screenWidth, screenHeight int, sessionService *application.SessionService) {
	bm.buttons = []Button{
		{
			X: screenWidth/2 - ButtonWidth/2,
			Y: screenHeight/2,
			W: ButtonWidth,
			H: ButtonHeight,
			Text: StartWorkButtonText,
			Action: func() {
				sessionService.StartWorkSession()
			},
		},
	}
}

func (bm *ButtonManager) UpdateButtonPositions(screenWidth, screenHeight int) {
	// Update button positions based on current screen size
	for i := range bm.buttons {
		switch len(bm.buttons) {
		case 1: // End of break (single button)
			bm.buttons[i].X = screenWidth/2 - ButtonWidth/2
			bm.buttons[i].Y = screenHeight/2
		case 2: // Main screen or end of work (two buttons)
			if i == 0 {
				bm.buttons[i].X = screenWidth/2 - ButtonWidth/2
				bm.buttons[i].Y = screenHeight/2 - ButtonHeight - ButtonPadding
			} else {
				bm.buttons[i].X = screenWidth/2 - ButtonWidth/2
				bm.buttons[i].Y = screenHeight/2 + ButtonPadding
			}
		}
	}
}

func (bm *ButtonManager) UpdateButtons() {
	mx, my := ebiten.CursorPosition()
	
	for i := range bm.buttons {
		button := &bm.buttons[i]
		button.Hovered = mx >= button.X && mx < button.X+button.W &&
			my >= button.Y && my < button.Y+button.H
		
		if button.Hovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			button.Action()
		}
	}
}

func (bm *ButtonManager) DrawButtons(screen *ebiten.Image) {
	for _, button := range bm.buttons {
		// Draw button shadow
		drawRect(screen, button.X+ButtonShadowOffset, button.Y+ButtonShadowOffset, button.W, button.H, ButtonShadow)
		
		// Draw button background
		buttonColor := ButtonColor
		if button.Hovered {
			buttonColor = ButtonHoverColor
		}
		drawRect(screen, button.X, button.Y, button.W, button.H, buttonColor)
		
		// Draw button border
		drawBorder(screen, button.X, button.Y, button.W, button.H, GrayBorder, ButtonBorderWidth)
		
		// Draw button text (centered)
		textX := button.X + button.W/2 - len(button.Text)*TextCharWidth
		textY := button.Y + button.H/2 - TextCharHeight
		ebitenutil.DebugPrintAt(screen, button.Text, textX, textY)
	}
}

func (bm *ButtonManager) GetButtons() []Button {
	return bm.buttons
}