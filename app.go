package main

import (
	"context"
	"time"

	"karedoro/domain"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx             context.Context
	pomodoroManager *domain.PomodoroManager
}

// NewApp creates a new App application struct
func NewApp() *App {
	app := &App{}
	
	config := domain.NewDefaultConfig()
	app.pomodoroManager = domain.NewPomodoroManager(config, app)
	
	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// StartWorkSession starts a new work session
func (a *App) StartWorkSession() error {
	return a.pomodoroManager.StartWorkSession()
}

// StartBreakSession starts a new break session
func (a *App) StartBreakSession() error {
	return a.pomodoroManager.StartBreakSession()
}

// PauseSession pauses the current session
func (a *App) PauseSession() error {
	return a.pomodoroManager.PauseSession()
}

// ResumeSession resumes the current session
func (a *App) ResumeSession() error {
	return a.pomodoroManager.ResumeSession()
}

// GetCurrentState returns the current application state
func (a *App) GetCurrentState() string {
	return a.pomodoroManager.GetCurrentState().String()
}

// GetRemainingTime returns the remaining time in seconds
func (a *App) GetRemainingTime() int {
	duration := a.pomodoroManager.GetRemainingTime()
	return int(duration.Seconds())
}

// EventHandler implementation for domain.EventHandler interface

// OnSessionStart handles session start events
func (a *App) OnSessionStart(session *domain.Session) {
	runtime.EventsEmit(a.ctx, "session:start", map[string]interface{}{
		"state":    session.State.String(),
		"duration": int(session.Duration.Seconds()),
	})
}

// OnSessionEnd handles session end events
func (a *App) OnSessionEnd(session *domain.Session) {
	runtime.EventsEmit(a.ctx, "session:end", map[string]interface{}{
		"state": session.State.String(),
	})
	
	// Enable fullscreen mode when session ends
	runtime.WindowFullscreen(a.ctx)
}

// OnSessionPause handles session pause events
func (a *App) OnSessionPause(session *domain.Session) {
	runtime.EventsEmit(a.ctx, "session:pause", map[string]interface{}{
		"state":         session.State.String(),
		"remainingTime": int(session.GetRemainingTime().Seconds()),
	})
}

// OnSessionResume handles session resume events
func (a *App) OnSessionResume(session *domain.Session) {
	runtime.EventsEmit(a.ctx, "session:resume", map[string]interface{}{
		"state":         session.State.String(),
		"remainingTime": int(session.GetRemainingTime().Seconds()),
	})
}

// OnTimerTick handles timer tick events
func (a *App) OnTimerTick(session *domain.Session, remainingTime time.Duration) {
	runtime.EventsEmit(a.ctx, "timer:tick", map[string]interface{}{
		"state":         session.State.String(),
		"remainingTime": int(remainingTime.Seconds()),
	})
}

// OnWarning handles warning events
func (a *App) OnWarning(idleDuration time.Duration) {
	runtime.EventsEmit(a.ctx, "warning", map[string]interface{}{
		"idleDuration": int(idleDuration.Minutes()),
	})
	
	// Show notification
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "Karedoro",
		Message: "まだ次のセッションを開始していません！",
	})
}
