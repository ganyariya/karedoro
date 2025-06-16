# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

**最終更新日**: 2025-01-16  
**プロジェクト状況**: ✅ **実装完了 - 動作可能状態**  
**日本語で回答してください。**

## Project Overview

karedoro is a **完全に実装済み**のPomodoro timer desktop applicationです。ebitengine v2で構築され、厳格なセッション遵守を強制します。フルスクリーンオーバーレイと持続的な通知により、ユーザーがポモドーロセッションを中断することを防ぎます。

## Current Implementation Status ✅

- **実装完了率**: 90% (イテレーション1-5完了)
- **動作状況**: 完全に動作する状態
- **テストカバレッジ**: 27テストケース通過
- **主要バグ**: 修正済み（タイマー一時停止、フォント表示）

## Architecture

- **Game Engine**: ebitengine v2 for cross-platform desktop application (macOS/Windows/Linux)
- **Language**: Go only
- **Domain Layer**: Session state management, timer functionality, state transition logic ✅
- **Application Layer**: Domain logic coordination, external system integration ✅
- **Presentation Layer**: ebitengine UI rendering, user input handling ✅
- **Design Pattern**: Domain-Driven Design (DDD) with core logic separated from UI layer ✅

## Implemented Features ✅

1. **Work Sessions (25min)** and **Break Sessions (5min)** with countdown timers ✅
2. **Fullscreen enforcement** when sessions end - prevents other app usage ✅
3. **Persistent warnings** every 5 minutes when idle between sessions ✅
4. **Pause/resume functionality** during active sessions (Space key) ✅
5. **Sound effects** and **popup notifications** for session transitions ✅

## Technology Stack (実装済み)

- Go language only ✅
- ebitengine v2 (v2.8.8) for UI rendering, event handling, and window management ✅
- ebitengine/oto v3 (v3.3.3) for programmatically generated sound effects ✅
- beeep library (v0.11.1) for OS-native system notifications ✅
- Cross-platform support (macOS/Windows/Linux) ✅

## Development Commands ✅

実装完了済みプロジェクトでの主要コマンド:
- `go run .` - アプリケーションを開発モードで実行 ✅
- `go build` - 実行可能ファイルをビルド ✅
- `go test ./...` - 全パッケージの単体テストを実行 ✅ (27テスト通過)
- Cross-platform builds using Go's build flags ✅

## Application States ✅

実装済みの3つの主要状態:
1. **Work Session**: 25分の集中作業期間 ✅
2. **Break Session**: 5分の休憩期間 ✅  
3. **Idle**: セッション間の待機状態 ✅

## Session Flow Details ✅

### Work Session Start ✅
- Triggered by "Start Work Session" button click or app launch
- Shows system notification popup using beeep ✅
- Plays programmatically generated sound effect using oto v3 ✅
- Starts 25-minute countdown ✅

### Break Session Start ✅
- Triggered by "Start Break Session" button after work session ends
- Exits fullscreen mode back to normal window size ✅
- Shows system notification popup using beeep ✅
- Plays programmatically generated sound effect using oto v3 ✅
- Starts 5-minute countdown ✅

### Session End Behavior ✅
- **Work session end**: Forces fullscreen mode using `ebiten.SetFullscreen(true)` with buttons:
  - "Start Break Session" ✅
  - "Skip Break -> Work" ✅
  - Message: "Pomodoro Complete! Take a break!" ✅
  - Plays sound effect and shows system notification ✅
- **Break session end**: Forces fullscreen mode using `ebiten.SetFullscreen(true)` with button:
  - "Start Work Session" ✅
  - Message: "Break Over! Back to work!" ✅
  - Plays sound effect and shows system notification ✅

### Warning System ✅
- Triggers every 5 minutes when idle in fullscreen overlay ✅
- Shows system notification using beeep with sound: "Haven't started next session yet!" ✅
- Continues until user starts next session ✅

### Pause/Resume System ✅
- **Space key** during sessions to pause/resume ✅
- Timer state properly maintained during pause ✅
- Visual indication when paused ✅

## Design Requirements ✅

- Minimalist and simple design ✅
- No gamification features ✅
- No detailed statistics ✅
- Focus purely on core Pomodoro functionality ✅

## Critical Implementation Notes ✅

- Use `ebiten.SetFullscreen(true)` for fullscreen window control ✅
- Use `ebiten.SetWindowClosingHandled(true)` to handle window close events and prevent app termination ✅
- Implement background operation - app continues running even when window is closed ✅
- Timer values (25min work, 5min break) implemented as constants with future configurability structure ✅
- Session state management is core domain logic - implement with proper state transitions ✅
- Use ebitengine/oto v3 for programmatically generated sound effects ✅
- Use beeep library for cross-platform system notifications ✅
- Settings persistence using Go standard library (`os`, `encoding/json`) ✅
- Comprehensive unit tests for core logic, timer and state management ✅ (27 tests)
- Pure Go implementation without external asset dependencies ✅

## Development Requirements ✅

- **Type Safety**: Proper types used throughout, no interface{} abuse ✅
- **No Magic Numbers**: All values properly defined as constants and types ✅
- **Commit Strategy**: Logical breakpoints with feature completion and passing tests ✅
- **Iterative Development**: ✅
  - Build and tests run after each iteration ✅
  - Errors fixed before proceeding ✅
  - Code refactored for quality maintenance ✅
- **Code Quality**: Clean, readable, and well-structured code maintained ✅

## Project Structure ✅

```
karedoro/
├── domain/           # Core business logic ✅
│   ├── types.go     # Constants and type definitions ✅
│   ├── timer.go     # Timer functionality ✅
│   ├── session.go   # Session state management ✅
│   └── *_test.go    # Domain tests (18 tests) ✅
├── application/      # Application services ✅
│   ├── session_service.go    # Session coordination ✅
│   ├── audio_service.go      # Sound effects ✅
│   ├── notification_service.go # System notifications ✅
│   ├── config_service.go     # Configuration management ✅
│   └── *_test.go            # Application tests (9 tests) ✅
├── presentation/     # UI layer ✅
│   ├── app.go       # Main application and game loop ✅
│   └── constants.go # UI constants ✅
└── main.go          # Entry point ✅
```

## Recent Fixes and Improvements ✅

1. **Timer Pause/Resume Bug**: Fixed timer state calculation during pause/resume ✅
2. **Font Display Issue**: Changed from Japanese to English text for compatibility ✅
3. **Audio System**: Migrated to oto v3 with programmatic sound generation ✅
4. **Build System**: Proper .gitignore and dependency management ✅
5. **Test Coverage**: Comprehensive test suite with 27 test cases ✅

## Usage Instructions ✅

1. **Run Application**: `go run .` または `./karedoro` ✅
2. **Controls**: 
   - マウス: ボタンクリックでセッション開始 ✅
   - スペースキー: セッション中の一時停止・再開 ✅
3. **Sessions**: 25分作業 → 5分休憩 のサイクル ✅
4. **Enforcement**: セッション終了時の自動フルスクリーン ✅