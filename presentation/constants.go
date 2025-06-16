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
	WorkSessionStartMessage    = "作業セッション開始！"
	BreakSessionStartMessage   = "休憩セッション開始！"
	WorkSessionEndMessage      = "ポモドーロ完了！休憩しましょう！"
	BreakSessionEndMessage     = "休憩終了！作業に戻りましょう！"
	WarningMessage            = "まだ次のセッションを開始していません！"
	
	StartWorkButtonText       = "作業セッションを開始"
	StartBreakButtonText      = "休憩セッションを開始"
	SkipBreakButtonText       = "休憩をスキップして作業セッションへ"
	PauseButtonText          = "一時停止"
	ResumeButtonText         = "再開"
	PausedText               = "一時停止中"
)