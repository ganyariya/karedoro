# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

**日本語で回答してください。**

## Project Overview

karedoro is a Pomodoro timer desktop application built with ebitengine v2 that enforces strict session adherence. The app prevents users from interrupting Pomodoro sessions through fullscreen overlays and persistent notifications.

## Architecture

- **Game Engine**: ebitengine v2 for cross-platform desktop application (macOS/Windows/Linux)
- **Language**: Go only
- **Domain Layer**: Session state management, timer functionality, state transition logic
- **Application Layer**: Domain logic coordination, external system integration
- **Presentation Layer**: ebitengine UI rendering, user input handling
- **Design Pattern**: Domain-Driven Design (DDD) with core logic separated from UI layer

## Key Features

1. **Work Sessions (25min)** and **Break Sessions (5min)** with countdown timers
2. **Fullscreen enforcement** when sessions end - prevents other app usage
3. **Persistent warnings** every 5 minutes when idle between sessions
4. **Pause/resume functionality** during active sessions
5. **Sound effects** and **popup notifications** for session transitions

## Technology Stack

- Go language only
- ebitengine v2 for UI rendering, event handling, and window management
- ebiten/oto for sound effects and audio playback
- beeep library for OS-native system notifications
- Cross-platform support (macOS/Windows/Linux)

## Development Commands

Since this is a new project, common commands will be:
- `go run .` - Run the application in development mode
- `go build` - Build the application executable
- `go test ./...` - Run unit tests for core logic
- Cross-platform builds using Go's build flags

## Application States

The app has 3 main states:
1. **Work Session**: 25-minute focused work period
2. **Break Session**: 5-minute rest period  
3. **Idle**: Waiting state between sessions

## Session Flow Details

### Work Session Start
- Triggered by "作業開始" button click or app launch
- Shows system notification popup using beeep
- Plays sound effect using ebiten/oto
- Starts 25-minute countdown

### Break Session Start  
- Triggered by "休憩開始" button after work session ends
- Exits fullscreen mode back to normal window size
- Shows system notification popup using beeep
- Plays sound effect using ebiten/oto
- Starts 5-minute countdown

### Session End Behavior
- **Work session end**: Forces fullscreen mode using `ebiten.SetFullscreen(true)` with buttons:
  - "休憩セッションを開始" 
  - "休憩をスキップして作業セッションへ"
  - Message: "ポモドーロ完了！休憩しましょう！"
  - Plays sound effect and shows system notification
- **Break session end**: Forces fullscreen mode using `ebiten.SetFullscreen(true)` with button:
  - "作業セッションを開始"
  - Message: "休憩終了！作業に戻りましょう！"
  - Plays sound effect and shows system notification

### Warning System
- Triggers every 5 minutes when idle in fullscreen overlay
- Shows system notification using beeep with sound: "まだ次のセッションを開始していません！"
- Continues until user starts next session

## Design Requirements

- Minimalist and simple design
- No gamification features
- No detailed statistics
- Focus purely on core Pomodoro functionality

## Critical Implementation Notes

- Use `ebiten.SetFullscreen(true)` for fullscreen window control
- Use `ebiten.SetWindowClosingHandled(true)` to handle window close events and prevent app termination
- Implement background operation - app continues running even when window is closed
- Timer values (25min work, 5min break) should be variables for future configurability
- Session state management is core domain logic - implement with proper state transitions
- Use ebiten/oto for sound effects (WAV/OGG format assets)
- Use beeep library for cross-platform system notifications
- Settings persistence using Go standard library (`os`, `encoding/json`) or similar
- Consider meaningful unit tests for core logic, especially timer and state management
- Asset embedding for sound files and other resources in the executable

## Development Requirements

- **Type Safety**: Use proper types wherever possible, avoid interface{} and any unsafe type operations
- **No Magic Numbers**: Replace all magic values with properly defined constants and types
- **Commit Strategy**: Commit at logical breakpoints when features are complete and tests pass
- **Iterative Development**: 
  - Run build and tests after each iteration
  - Fix any errors before proceeding to next iteration
  - Refactor code after each iteration to maintain code quality
- **Code Quality**: Maintain clean, readable, and well-structured code throughout development