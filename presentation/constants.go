package presentation

import "image/color"

const (
	WindowWidth  = 800
	WindowHeight = 600
	WindowTitle  = "karedoro - ポモドーロタイマー"
	
	MinWindowWidth  = 600
	MinWindowHeight = 400
	
	FontSize         = 24
	LargeFontSize    = 36
	ButtonWidth      = 200
	ButtonHeight     = 50
	ButtonPadding    = 10
	
	TimerFontSize    = 48
	MessageFontSize  = 32
	
	// Layout constants
	TimerBoxWidth         = 120
	TimerBoxHeight        = 30
	TimerOffsetX          = 60
	TimerOffsetY          = 110
	ProgressBarWidth      = 300
	ProgressBarHeight     = 10
	ProgressBarOffsetY    = 40
	MessageBoxPadding     = 40
	MessageBoxHeight      = 60
	MessageBoxBorderWidth = 3
	ButtonShadowOffset    = 2
	ButtonBorderWidth     = 2
	
	// Text positioning
	TextCharWidth     = 3
	TextCharWidthLg   = 4
	TextCharHeight    = 6
	TextLineHeight    = 50
	IdleMessageOffset = 150
)

var (
	BackgroundColor     = color.RGBA{R: 45, G: 45, B: 45, A: 255}
	TextColor           = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	ButtonColor         = color.RGBA{R: 70, G: 130, B: 180, A: 255}
	ButtonHoverColor    = color.RGBA{R: 100, G: 149, B: 237, A: 255}
	ButtonTextColor     = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	WorkSessionColor    = color.RGBA{R: 220, G: 20, B: 60, A: 255}
	BreakSessionColor   = color.RGBA{R: 34, G: 139, B: 34, A: 255}
	WarningColor        = color.RGBA{R: 255, G: 165, B: 0, A: 255}
	
	// Enhanced enforcement colors
	ForceRedBackground  = color.RGBA{R: 180, G: 0, B: 0, A: 255}
	ForceYellowBox      = color.RGBA{R: 255, G: 255, B: 0, A: 200}
	WhiteBorder         = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	BlackShadow         = color.RGBA{R: 0, G: 0, B: 0, A: 100}
	ButtonShadow        = color.RGBA{R: 0, G: 0, B: 0, A: 50}
	GrayBorder          = color.RGBA{R: 200, G: 200, B: 200, A: 255}
	ProgressBarBg       = color.RGBA{R: 60, G: 60, B: 60, A: 255}
	ProgressBarFill     = color.RGBA{R: 100, G: 200, B: 100, A: 255}
	ProgressBarBorder   = color.RGBA{R: 180, G: 180, B: 180, A: 255}
)

const (
	WorkSessionStartMessage    = "WORK SESSION STARTED!"
	BreakSessionStartMessage   = "BREAK SESSION STARTED!"
	WorkSessionEndMessage      = "POMODORO COMPLETE! You MUST take a break!"
	BreakSessionEndMessage     = "BREAK OVER! Get back to work NOW!"
	WarningMessage            = "WARNING! Start your next session!"
	
	StartWorkButtonText       = "START WORK SESSION"
	StartBreakButtonText      = "START BREAK SESSION"
	SkipBreakButtonText       = "SKIP BREAK -> WORK"
	PauseButtonText          = "PAUSE"
	ResumeButtonText         = "RESUME"
	PausedText               = "PAUSED"
	
	IdleScreenMessage         = "You MUST choose your next session:"
	WorkingText               = "WORKING - STAY FOCUSED!"
	BreakText                 = "BREAK TIME - RELAX!"
	PauseInstructionText      = "Press SPACE to pause"
	ResumeInstructionText     = "Press SPACE to resume"
)