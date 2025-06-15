package config

import "time"

// Application constants
const (
	AppName    = "Karedoro"
	AppVersion = "1.0.0"
)

// Session duration constants
const (
	DefaultWorkDuration  = 25 * time.Minute
	DefaultBreakDuration = 5 * time.Minute
)

// Timer interval constants
const (
	TimerTickInterval = 100 * time.Millisecond
	WarningInterval   = 5 * time.Minute
)

// Event name constants
const (
	EventSessionStart  = "session:start"
	EventSessionEnd    = "session:end"
	EventSessionPause  = "session:pause"
	EventSessionResume = "session:resume"
	EventTimerTick     = "timer:tick"
	EventWarning       = "warning"
)

// UI message constants
const (
	MessageWarningIdle       = "まだ次のセッションを開始していません！"
	MessageWorkCompleted     = "ポモドーロ完了！休憩しましょう！"
	MessageBreakCompleted    = "休憩終了！作業に戻りましょう！"
	MessageSessionStartWork  = "作業セッション開始！"
	MessageSessionStartBreak = "休憩セッション開始！"
)