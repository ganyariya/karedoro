# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

**日本語で回答してください。**

## Project Overview

karedoro is a Pomodoro timer desktop application built with Wails (Go backend + React/TypeScript frontend) that enforces strict session adherence. The app prevents users from interrupting Pomodoro sessions through fullscreen overlays and persistent notifications.

## Architecture

- **Backend (Go)**: Core logic, session state management, timer functionality, OS-level integrations
- **Frontend (React/TypeScript)**: UI rendering, user input handling, communication with Go backend
- **Framework**: Wails for cross-platform desktop app (macOS/Windows/Linux)
- **Design Pattern**: Domain-Driven Design (DDD) with core logic in Go domain layer

## Key Features

1. **Work Sessions (25min)** and **Break Sessions (5min)** with countdown timers
2. **Fullscreen enforcement** when sessions end - prevents other app usage
3. **Persistent warnings** every 5 minutes when idle between sessions
4. **Pause/resume functionality** during active sessions
5. **Sound effects** and **popup notifications** for session transitions

## Technology Stack

- Go backend with Wails runtime APIs for window control
- React/TypeScript frontend 
- HTML5 Audio API for sound effects
- OS-native notifications through Go libraries
- Wails Bindings for frontend-to-backend function calls
- Wails Events for backend-to-frontend state updates

## Development Commands

Since this is a new project, common commands will be:
- `wails build` - Build the application
- `wails dev` - Development mode with hot reload
- `go test ./...` - Run Go backend tests
- `npm test` - Run frontend tests (when implemented)

## Application States

The app has 3 main states:
1. **Work Session**: 25-minute focused work period
2. **Break Session**: 5-minute rest period  
3. **Idle**: Waiting state between sessions

## Session Flow Details

### Work Session Start
- Triggered by "作業開始" button click or app launch
- Shows popup notification with sound effect
- Starts 25-minute countdown

### Break Session Start  
- Triggered by "休憩開始" button after work session ends
- Exits fullscreen mode back to normal desktop
- Shows popup notification with sound effect
- Starts 5-minute countdown

### Session End Behavior
- **Work session end**: Forces fullscreen overlay with buttons:
  - "休憩セッションを開始" 
  - "休憩をスキップして作業セッションへ"
  - Message: "ポモドーロ完了！休憩しましょう！"
- **Break session end**: Forces fullscreen overlay with button:
  - "作業セッションを開始"
  - Message: "休憩終了！作業に戻りましょう！"

### Warning System
- Triggers every 5 minutes when idle in fullscreen overlay
- Shows popup with sound: "まだ次のセッションを開始していません！"
- Continues until user starts next session

## Design Requirements

- Minimalist and simple design
- No gamification features
- No detailed statistics
- Focus purely on core Pomodoro functionality

## Critical Implementation Notes

- Use `runtime.WindowSetFullscreen` and related Wails APIs for window control
- Use `runtime.WindowSetResizable`, `runtime.WindowSetMinSize` to prevent window manipulation
- Implement OS-level event handling to prevent forced app closure/minimization
- Timer values (25min work, 5min break) should be variables for future configurability
- Session state management is core domain logic - implement with proper state transitions
- All user interactions must go through Wails Bindings to Go functions
- Real-time UI updates use Wails Events from Go to React
- Settings persistence using Go standard library (`os`, `encoding/json`) or `spf13/viper`
- Consider meaningful unit tests for core logic, especially timer and state management