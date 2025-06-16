package presentation

import "image/color"

const (
	WindowWidth  = 800
	WindowHeight = 600
	WindowTitle  = "karedoro - ポモドーロタイマー"
	
	FontSize         = 24
	LargeFontSize    = 36
	ButtonWidth      = 200
	ButtonHeight     = 50
	ButtonPadding    = 10
	
	TimerFontSize    = 48
	MessageFontSize  = 32
)

var (
	BackgroundColor    = color.RGBA{R: 45, G: 45, B: 45, A: 255}
	TextColor          = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	ButtonColor        = color.RGBA{R: 70, G: 130, B: 180, A: 255}
	ButtonHoverColor   = color.RGBA{R: 100, G: 149, B: 237, A: 255}
	ButtonTextColor    = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	WorkSessionColor   = color.RGBA{R: 220, G: 20, B: 60, A: 255}
	BreakSessionColor  = color.RGBA{R: 34, G: 139, B: 34, A: 255}
	WarningColor       = color.RGBA{R: 255, G: 165, B: 0, A: 255}
)

const (
	WorkSessionStartMessage    = "Work Session Started!"
	BreakSessionStartMessage   = "Break Session Started!"
	WorkSessionEndMessage      = "Pomodoro Complete! Take a break!"
	BreakSessionEndMessage     = "Break Over! Back to work!"
	WarningMessage            = "Haven't started next session yet!"
	
	StartWorkButtonText       = "Start Work Session"
	StartBreakButtonText      = "Start Break Session"
	SkipBreakButtonText       = "Skip Break -> Work"
	PauseButtonText          = "Pause"
	ResumeButtonText         = "Resume"
	PausedText               = "PAUSED"
	
	IdleScreenMessage         = "Choose your next session:"
	WorkingText               = "WORKING"
	BreakText                 = "BREAK TIME"
	PauseInstructionText      = "Press SPACE to pause"
	ResumeInstructionText     = "Press SPACE to resume"
)