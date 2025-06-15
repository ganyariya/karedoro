# karedoro 開発タスク

## プロジェクト概要
ポモドーロタイマーの徹底継続を目的としたWails（Go + React/TypeScript）製デスクトップアプリケーション

## 開発フェーズ

### 📋 Phase 1: バックエンド基盤構築（Go） ✅ **完了**

#### 1.1 ドメインモデル設計
- [x] セッション状態の定義（Work/Break/Idle）
- [x] セッション時間の設定構造体（作業25分、休憩5分）
- [x] ドメインロジックのインターフェース設計

#### 1.2 セッション管理コア機能
- [x] タイマー機能の実装（カウントダウン機能）
- [x] セッション状態遷移ロジック
- [x] 一時停止・再開機能
- [x] セッション終了時の処理

#### 1.3 システム連携機能
- [x] Wailsランタイム統合（ウィンドウ制御）
- [x] 全画面表示の強制実装
- [x] OSネイティブ通知システム
- [x] 警告システム（5分間隔）

#### 1.4 Wails Bindings
- [x] フロントエンド向けAPI関数の実装
  - [x] セッション開始関数
  - [x] 一時停止・再開関数
  - [x] セッション状態取得関数

#### 1.5 Wails Events
- [x] リアルタイム状態更新イベント
- [x] タイマー残り時間の配信
- [x] 警告トリガーイベント

### 🎨 Phase 2: フロントエンド基盤構築（React/TypeScript）

#### 2.1 UI設計・実装
- [ ] メインタイマー画面の実装
- [ ] 全画面オーバーレイ画面の実装
- [ ] セッション開始ボタンUIの実装
- [ ] 一時停止・再開ボタンの実装

#### 2.2 状態管理
- [ ] React状態管理の設計
- [ ] Wails Eventsの受信処理
- [ ] UIのリアルタイム更新機能

#### 2.3 音響・通知機能
- [ ] HTML5 Audio API統合
- [ ] セッション開始・終了効果音
- [ ] ポップアップ通知システム

#### 2.4 Wails連携
- [ ] Go関数呼び出しの実装
- [ ] エラーハンドリング
- [ ] TypeScript型定義の整備

### 🎯 Phase 3: コア機能統合

#### 3.1 セッションフロー実装
- [ ] 作業セッション開始→終了フロー
- [ ] 休憩セッション開始→終了フロー
- [ ] セッション間遷移の全画面制御

#### 3.2 強制継続機能
- [ ] セッション終了時の全画面表示強制
- [ ] 他アプリ操作の無効化
- [ ] 警告システムの統合テスト

#### 3.3 ウィンドウ制御
- [ ] フルスクリーンモードの制御
- [ ] ウィンドウリサイズ・最小化の無効化
- [ ] OSレベルイベントハンドリング

#### 3.4 バックグラウンド実行機能 🆕
- [ ] ウィンドウクローズ時の継続動作実装
- [ ] システムトレイ統合（macOS: メニューバー、Windows: タスクトレイ）
- [ ] バックグラウンド状態でのセッション管理
- [ ] セッション終了時の強制ウィンドウ復帰
- [ ] バックグラウンド中の通知システム

### 🧪 Phase 4: テスト・品質保証

#### 4.1 単体テスト
- [ ] ドメインロジックのテスト
- [ ] タイマー機能のテスト
- [ ] セッション状態遷移のテスト

#### 4.2 統合テスト
- [ ] フロントエンド・バックエンド連携テスト
- [ ] Wailsランタイム機能のテスト
- [ ] エンドツーエンドシナリオテスト

#### 4.3 品質確認
- [ ] パフォーマンス検証
- [ ] メモリリーク検証
- [ ] クロスプラットフォーム動作確認

### 🚀 Phase 5: 完成・リリース準備

#### 5.1 最終調整
- [ ] UI/UXの微調整
- [ ] エラーメッセージの日本語化
- [ ] 設定値の最終確認

#### 5.2 ビルド・パッケージング
- [ ] リリースビルドの作成
- [ ] 各OS向けパッケージング
- [ ] インストーラー作成（必要に応じて）

#### 5.3 ドキュメンテーション
- [ ] ユーザーマニュアル作成
- [ ] 開発者向けドキュメント整備
- [ ] README.mdの更新

## 重要な技術的考慮事項

### 🔧 実装上の注意点
1. **ドメイン駆動設計**：コアロジックをドメイン層として分離
2. **Wails特有のAPI**：ランタイム機能を適切に活用
3. **強制力の実装**：OSレベルでのウィンドウ制御
4. **リアルタイム更新**：Wails Eventsを活用した状態同期
5. **バックグラウンド実行**：ウィンドウを閉じてもタイマーが継続動作
   - ウィンドウクローズ時にアプリケーションプロセスは終了せず、バックグラウンドで継続実行
   - システムトレイ（macOS: メニューバー、Windows: タスクトレイ）にアイコンを表示
   - セッション終了時は強制的にウィンドウを復帰させ、全画面表示
   - バックグラウンド中もタイマーと警告システムは正常動作

### 📦 依存関係
- Wails v2
- Go標準ライブラリ
- React 18+
- TypeScript 4.5+
- 設定管理：spf13/viper（将来的な拡張用）

### 🎨 デザインガイドライン
- ミニマリスト・シンプルなデザイン
- ゲーミフィケーション要素は排除
- 統計機能は実装しない
- ポモドーロのコア機能に専念

## 🔍 開発イテレーション品質保証要件

### 必須確認プロセス
**すべての開発タスクにおいて、以下の4つの確認を必須とする：**

#### 1. テスト実行確認 ✅
- **Go バックエンドテスト**
  ```bash
  go test ./...
  ```
- **フロントエンドテスト**（実装後）
  ```bash
  npm test
  ```
- **すべてのテストがPASSすること**
- **新機能に対する適切なテストが追加されていること**

#### 2. ビルド確認 🔨
- **開発モードビルド**
  ```bash
  wails dev
  ```
- **プロダクションビルド**
  ```bash
  wails build
  ```
- **ビルドエラーが発生しないこと**
- **警告メッセージの確認と対応**

#### 3. 動作確認 🎯
- **基本動作の確認**
  - アプリケーションが正常に起動すること
  - 実装した機能が仕様通りに動作すること
  - エラーハンドリングが適切に機能すること
- **UI/UX確認**
  - ユーザーインターフェースが正常に表示されること
  - ボタン・入力フィールドが正常に機能すること
  - 画面遷移が正常に動作すること

#### 4. コード品質確認 🔧
- **Go言語ベストプラクティス確認**
  - マジックナンバー・マジック文字列の排除
  - 適切な定数・変数名の使用
  - エラーハンドリングの改善
  - 構造体・インターフェース設計の見直し
  - **型安全性・Entity設計の確認** 🎯
- **TypeScriptベストプラクティス確認**
  - 型安全性の向上
  - マジック値の定数化
  - コンポーネント設計の改善
  - イベントハンドリングの最適化
  - **ブランド型・制約型の適切な実装** 🎯
- **AI理解性確認** 🤖
  - 型定義による意図の明確化
  - 自己文書化されたコード構造
  - ドメイン知識の型システムへの組み込み

### 📋 各フェーズ完了時の確認チェックリスト

#### Phase 1 完了確認
- [ ] ✅ Go テスト実行：`go test ./...` PASS
- [ ] 🔨 ビルド確認：`wails build` 成功
- [ ] 🎯 タイマー機能の基本動作確認
- [ ] 🎯 セッション状態遷移ロジック確認
- [ ] 🎯 ドメインロジックの単体動作確認
- [ ] 🔧 Phase 1 コード品質リファクタリング
- [ ] 🏗️ Phase 1 型定義強化（Go基本エンティティ・値オブジェクト）

#### Phase 2 完了確認
- [ ] ✅ フロントエンド・バックエンド両方のテスト PASS
- [ ] 🔨 開発モード・プロダクションモード両方でビルド成功
- [ ] 🎯 UIコンポーネントの表示確認
- [ ] 🎯 Wails Bindings の動作確認
- [ ] 🎯 音響・通知機能の動作確認
- [ ] 🔧 Phase 2 コード品質リファクタリング
- [ ] 🏗️ Phase 2 型定義強化（TypeScript型システム・UIコンポーネント型）

#### Phase 3 完了確認
- [ ] ✅ 統合テスト含む全テスト PASS
- [ ] 🔨 全機能統合後のビルド成功
- [ ] 🎯 完全なセッションフローの動作確認
- [ ] 🎯 全画面強制機能の動作確認
- [ ] 🎯 警告システムの動作確認
- [ ] 🎯 ウィンドウ制御機能の確認
- [ ] 🔧 Phase 3 コード品質リファクタリング
- [ ] 🏗️ Phase 3 型定義強化（API境界・ビジネスロジック型）

#### Phase 4 完了確認
- [ ] ✅ 全単体・統合・E2Eテスト PASS
- [ ] 🔨 リリースビルドの成功確認
- [ ] 🎯 パフォーマンス・メモリリーク検証
- [ ] 🎯 クロスプラットフォーム動作確認
- [ ] 🎯 長時間動作の安定性確認
- [ ] 🔧 Phase 4 コード品質リファクタリング
- [ ] 🏗️ Phase 4 型定義強化（テストデータ・エラー型・パフォーマンス型）

#### Phase 5 完了確認
- [ ] ✅ 最終テストスイート PASS
- [ ] 🔨 各OS向けリリースビルド成功
- [ ] 🎯 インストール・実行の動作確認
- [ ] 🎯 ユーザーシナリオ通りの動作確認
- [ ] 🎯 エラーケース対応の確認
- [ ] 🔧 最終コード品質リファクタリング
- [ ] 🏗️ 最終型定義確認（プロダクション型・保守性確保）

### ⚠️ 品質保証ルール

1. **ゼロテスト破綻原則**
   - 既存テストを破綻させる変更は許可しない
   - 新機能追加時は対応するテストも同時に実装

2. **ビルド成功原則**
   - 開発ブランチは常にビルド可能な状態を保つ
   - ビルドエラーが発生した場合は即座に修正

3. **動作確認原則**
   - 機能実装後は必ず手動での動作確認を実施
   - 想定外の動作がある場合は仕様確認・修正を行う

4. **段階的統合原則**
   - 各フェーズ完了時に必ず全体統合テストを実行
   - 問題発見時は該当フェーズに戻って修正

5. **進捗管理原則**
   - TASK.mdのチェックボックスを実際の進捗に合わせてリアルタイム更新
   - 作業開始時にタスクを開始状態に、完了時に完了状態に更新
   - 新たな課題や追加タスクが発見された場合はTASK.mdに追記

6. **定期コミット原則**
   - 意味のある機能単位（サブタスク完了時）で必ずコミット
   - フェーズ完了時には必ずコミット
   - ビルド・テスト通過を確認してからコミット実行

7. **イテレーション・リファクタリング原則**
   - 各フェーズ完了時に必ずコード品質の見直しを実施
   - Go/TypeScriptのベストプラクティスに準拠したリファクタリング
   - マジックナンバー・マジック文字列の完全排除
   - コードの可読性・保守性・拡張性の向上

8. **型定義徹底原則** 🎯 **AI開発最適化**
   - Entity・型・インターフェースの丁寧な実装により AI 理解性を最大化
   - 型安全性によるマジック定数の自然な抑制
   - ドメイン知識の型システムへの明示的な組み込み
   - AIが推論しやすい明確な型境界の定義

### 🚨 品質保証失敗時の対応

**テスト失敗時：**
- 失敗原因の特定と修正
- 関連テストの見直し
- 修正後の再テスト実行

**ビルド失敗時：**
- エラーログの詳細確認
- 依存関係・設定の見直し
- 段階的な問題切り分け

**動作不具合時：**
- 期待動作と実際の動作の比較
- 仕様書との照合
- 必要に応じて仕様変更の検討

## 🧹 コード品質・リファクタリング要件

### 💎 リファクタリング基本方針

**各フェーズ完了時に必須で実施するコード品質向上作業**

#### 1. マジック値の完全排除 🚨 **重要**
- **マジックナンバーの定数化**
  - `25 * time.Minute` → `const DefaultWorkDuration = 25 * time.Minute`
  - `5 * time.Minute` → `const DefaultBreakDuration = 5 * time.Minute`
  - `100 * time.Millisecond` → `const TimerTickInterval = 100 * time.Millisecond`
  - `5 * time.Minute` → `const WarningInterval = 5 * time.Minute`

- **マジック文字列の定数化**
  - `"session:start"` → `const EventSessionStart = "session:start"`
  - `"Idle"`, `"WorkSession"`, `"BreakSession"` → 適切な型定義
  - `"Karedoro"` → `const AppName = "Karedoro"`

#### 2. Go言語ベストプラクティス適用

##### 2.1 定数・設定管理
```go
// config/constants.go
package config

import "time"

const (
    // Session durations
    DefaultWorkDuration  = 25 * time.Minute
    DefaultBreakDuration = 5 * time.Minute
    
    // Timer intervals
    TimerTickInterval = 100 * time.Millisecond
    WarningInterval   = 5 * time.Minute
    
    // Application constants
    AppName = "Karedoro"
    AppVersion = "1.0.0"
)

// Event names
const (
    EventSessionStart  = "session:start"
    EventSessionEnd    = "session:end"
    EventSessionPause  = "session:pause"
    EventSessionResume = "session:resume"
    EventTimerTick     = "timer:tick"
    EventWarning       = "warning"
)
```

##### 2.2 エラーハンドリングの改善
- カスタムエラー型の定義
- センチネルエラーの使用
- エラーラッピングの適切な実装

##### 2.3 構造体・インターフェース設計
- 単一責任の原則の適用
- インターフェース分離の原則
- 依存性逆転の原則

#### 3. TypeScript/React ベストプラクティス適用

##### 3.1 型安全性の向上
```typescript
// types/session.ts
export const SessionState = {
  IDLE: 'Idle',
  WORK_SESSION: 'WorkSession',
  BREAK_SESSION: 'BreakSession'
} as const;

export type SessionStateType = typeof SessionState[keyof typeof SessionState];

// constants/app.ts
export const APP_CONSTANTS = {
  TIMER_UPDATE_INTERVAL: 1000,
  WARNING_IDLE_MINUTES: 5,
  APP_TITLE: 'Karedoro ポモドーロタイマー'
} as const;

// constants/events.ts
export const EVENTS = {
  SESSION_START: 'session:start',
  SESSION_END: 'session:end',
  SESSION_PAUSE: 'session:pause',
  SESSION_RESUME: 'session:resume',
  TIMER_TICK: 'timer:tick',
  WARNING: 'warning'
} as const;
```

##### 3.2 コンポーネント設計改善
- プロップスの型定義強化
- カスタムフックの抽出
- コンポーネントの責任分離

##### 3.3 状態管理の最適化
- useState の適切な使用
- useEffect の依存配列最適化
- メモ化の適切な適用

#### 4. ファイル構成・アーキテクチャ改善

##### 4.1 Go側のディレクトリ構造
```
/
├── cmd/                 # エントリーポイント
├── internal/           # プライベートパッケージ
│   ├── app/           # アプリケーション層
│   ├── domain/        # ドメイン層（既存）
│   ├── config/        # 設定・定数
│   └── errors/        # エラー定義
├── pkg/               # パブリックパッケージ
└── test/              # テストユーティリティ
```

##### 4.2 TypeScript側のディレクトリ構造
```
src/
├── components/        # UIコンポーネント
├── hooks/            # カスタムフック
├── types/            # 型定義
├── constants/        # 定数定義
├── utils/            # ユーティリティ関数
└── services/         # Wails API呼び出し
```

### 🔍 リファクタリング実施タイミング

#### フェーズ完了時のリファクタリング内容

**Phase 1 リファクタリング**
- [ ] ドメイン層の定数抽出
- [ ] 設定管理の分離
- [ ] エラー型の定義
- [ ] テストの改善

**Phase 2 リファクタリング**
- [ ] フロントエンド定数の抽出
- [ ] 型定義の強化
- [ ] コンポーネント分離
- [ ] カスタムフックの抽出

**Phase 3 リファクタリング**
- [ ] アーキテクチャ全体の見直し
- [ ] パフォーマンス最適化
- [ ] メモリリーク対策
- [ ] セキュリティ向上

**Phase 4 リファクタリング**
- [ ] コードレビューによる品質向上
- [ ] ドキュメント整備
- [ ] テストカバレッジ向上
- [ ] 静的解析結果の対応

**Phase 5 リファクタリング**
- [ ] 最終コード監査
- [ ] パフォーマンス測定・改善
- [ ] プロダクション準備
- [ ] 保守性確保

### ⚡ リファクタリング実施ルール

1. **ゼロ・リグレッション原則**
   - リファクタリング後もすべてのテストが通ること
   - 機能的な変更は一切含めない

2. **段階的改善原則**
   - 一度に大きな変更をせず、小さな改善を積み重ねる
   - 各改善後にテスト・ビルド・動作確認を実施

3. **文書化原則**
   - 変更理由と改善内容を明確にコミットメッセージに記載
   - 破壊的変更がある場合は事前に影響範囲を調査

4. **パフォーマンス維持原則**
   - リファクタリングによってパフォーマンスが劣化しないこと
   - 必要に応じてベンチマークテストを実施

## 🏗️ 型定義・Entity設計要件

### 💎 AI開発最適化のための型設計原則

**型システムを通じたドメイン知識の明示化により、AIの理解性と開発効率を最大化する**

#### 1. 型安全性によるマジック値抑制 🚨 **最重要**

##### 1.1 Go側の型安全設計
```go
// ❌ 悪い例：マジック値が混入しやすい
func processSession(sessionType string, duration int) error {
    if sessionType == "work" && duration > 1500 { // マジック値
        // ...
    }
}

// ✅ 良い例：型で制約を表現
type SessionDuration time.Duration
type SessionType int

const (
    SessionTypeWork SessionType = iota
    SessionTypeBreak
)

func (st SessionType) DefaultDuration() SessionDuration {
    switch st {
    case SessionTypeWork:
        return SessionDuration(config.DefaultWorkDuration)
    case SessionTypeBreak:
        return SessionDuration(config.DefaultBreakDuration)
    }
}

func processSession(sessionType SessionType, duration SessionDuration) error {
    if sessionType == SessionTypeWork && duration > sessionType.DefaultDuration() {
        // 型安全でマジック値なし
    }
}
```

##### 1.2 TypeScript側の型安全設計
```typescript
// ❌ 悪い例：マジック値と型安全性の欠如
function updateTimer(state: string, time: number) {
    if (state === "work" && time < 25 * 60) { // マジック値
        // ...
    }
}

// ✅ 良い例：型とブランディングによる安全性
type Minutes = number & { readonly __brand: unique symbol };
type Seconds = number & { readonly __brand: unique symbol };

const createMinutes = (value: number): Minutes => value as Minutes;
const createSeconds = (value: number): Seconds => value as Seconds;

interface SessionTimerConfig {
    readonly workDuration: Minutes;
    readonly breakDuration: Minutes;
}

function updateTimer(state: SessionStateType, remainingTime: Seconds, config: SessionTimerConfig) {
    if (state === SESSION_STATE.WORK_SESSION && remainingTime < (config.workDuration * 60)) {
        // 型安全でマジック値なし
    }
}
```

#### 2. Entity・ドメインオブジェクト設計

##### 2.1 Goドメインエンティティ強化
```go
// SessionID型の導入
type SessionID string

func NewSessionID() SessionID {
    return SessionID(uuid.New().String())
}

// 時刻情報の型安全化
type SessionTimestamp time.Time
type SessionDuration time.Duration

// セッション統計のエンティティ
type SessionStats struct {
    TotalSessions     int           `json:"totalSessions"`
    CompletedSessions int           `json:"completedSessions"`
    TotalWorkTime     SessionDuration `json:"totalWorkTime"`
    TotalBreakTime    SessionDuration `json:"totalBreakTime"`
}

// セッション進捗の値オブジェクト
type SessionProgress struct {
    ElapsedTime   SessionDuration `json:"elapsedTime"`
    RemainingTime SessionDuration `json:"remainingTime"`
    Progress      float64         `json:"progress"` // 0.0-1.0
}

func (sp SessionProgress) IsNearCompletion() bool {
    return sp.Progress >= 0.9
}
```

##### 2.2 TypeScript型システム強化
```typescript
// ブランド型による型安全性
type SessionID = string & { readonly __brand: 'SessionID' };
type Timestamp = number & { readonly __brand: 'Timestamp' };
type Duration = number & { readonly __brand: 'Duration' };

// セッション進捗の型定義
interface SessionProgress {
    readonly elapsedTime: Duration;
    readonly remainingTime: Duration;
    readonly progress: number; // 0.0-1.0の制約をコメントで明示
}

// セッション統計の型定義
interface SessionStats {
    readonly totalSessions: number;
    readonly completedSessions: number;
    readonly totalWorkTime: Duration;
    readonly totalBreakTime: Duration;
}

// イベントデータの型安全化
interface SessionStartEvent {
    readonly sessionId: SessionID;
    readonly sessionType: SessionStateType;
    readonly duration: Duration;
    readonly timestamp: Timestamp;
}

// ユーティリティ型関数
const isNearCompletion = (progress: SessionProgress): boolean => 
    progress.progress >= 0.9;
```

#### 3. インターフェース・契約設計

##### 3.1 サービス層インターフェース
```go
// ポモドーロサービスの完全なインターフェース定義
type PomodoroService interface {
    // セッション管理
    StartSession(sessionType SessionType) (*Session, error)
    PauseSession(sessionID SessionID) error
    ResumeSession(sessionID SessionID) error
    EndSession(sessionID SessionID) (*SessionStats, error)
    
    // 状態照会
    GetCurrentSession() (*Session, error)
    GetSessionProgress(sessionID SessionID) (*SessionProgress, error)
    GetSessionStats() (*SessionStats, error)
    
    // 設定管理
    UpdateConfig(config *SessionConfig) error
    GetConfig() *SessionConfig
}

// イベントハンドラーの型安全化
type SessionEventType int

const (
    SessionEventStart SessionEventType = iota
    SessionEventEnd
    SessionEventPause
    SessionEventResume
    SessionEventTick
    SessionEventWarning
)

type SessionEvent struct {
    Type      SessionEventType `json:"type"`
    SessionID SessionID       `json:"sessionId"`
    Timestamp SessionTimestamp `json:"timestamp"`
    Data      interface{}     `json:"data"`
}

type TypedEventHandler interface {
    HandleSessionEvent(event SessionEvent) error
}
```

##### 3.2 APIレスポンス型定義
```typescript
// APIレスポンスの型安全化
interface ApiResponse<T> {
    readonly success: boolean;
    readonly data?: T;
    readonly error?: {
        readonly code: string;
        readonly message: string;
    };
}

// セッション操作のレスポンス型
type StartSessionResponse = ApiResponse<{
    readonly session: {
        readonly id: SessionID;
        readonly type: SessionStateType;
        readonly startTime: Timestamp;
        readonly duration: Duration;
    };
}>;

type SessionProgressResponse = ApiResponse<SessionProgress>;
type SessionStatsResponse = ApiResponse<SessionStats>;

// Wails API呼び出しの型安全ラッパー
interface PomodoroAPI {
    startWorkSession(): Promise<StartSessionResponse>;
    startBreakSession(): Promise<StartSessionResponse>;
    pauseSession(): Promise<ApiResponse<void>>;
    resumeSession(): Promise<ApiResponse<void>>;
    getCurrentState(): Promise<ApiResponse<SessionStateType>>;
    getProgress(): Promise<SessionProgressResponse>;
    getStats(): Promise<SessionStatsResponse>;
}
```

#### 4. バリデーション・制約の型表現

##### 4.1 型レベルでの制約表現
```go
// 値オブジェクトによる制約の表現
type TimerInterval struct {
    value time.Duration
}

func NewTimerInterval(d time.Duration) (*TimerInterval, error) {
    if d < 100*time.Millisecond || d > 10*time.Second {
        return nil, errors.New("timer interval must be between 100ms and 10s")
    }
    return &TimerInterval{value: d}, nil
}

func (ti TimerInterval) Duration() time.Duration {
    return ti.value
}

// セッション設定の制約
type SessionConfig struct {
    workDuration  SessionDuration
    breakDuration SessionDuration
    timerInterval TimerInterval
}

func NewSessionConfig(workMin, breakMin int, intervalMs int) (*SessionConfig, error) {
    if workMin < 1 || workMin > 120 {
        return nil, errors.New("work duration must be 1-120 minutes")
    }
    if breakMin < 1 || breakMin > 60 {
        return nil, errors.New("break duration must be 1-60 minutes")
    }
    
    interval, err := NewTimerInterval(time.Duration(intervalMs) * time.Millisecond)
    if err != nil {
        return nil, err
    }
    
    return &SessionConfig{
        workDuration:  SessionDuration(time.Duration(workMin) * time.Minute),
        breakDuration: SessionDuration(time.Duration(breakMin) * time.Minute),
        timerInterval: *interval,
    }, nil
}
```

##### 4.2 TypeScript制約型
```typescript
// 数値制約の型表現
type PositiveNumber = number & { readonly __constraint: 'positive' };
type Range<T extends number, U extends number> = number & { 
    readonly __range: [T, U] 
};

type WorkDurationMinutes = Range<1, 120>;
type BreakDurationMinutes = Range<1, 60>;
type ProgressPercentage = Range<0, 100>;

// バリデーション関数
const createPositiveNumber = (value: number): PositiveNumber | null =>
    value > 0 ? value as PositiveNumber : null;

const createWorkDuration = (minutes: number): WorkDurationMinutes | null =>
    (minutes >= 1 && minutes <= 120) ? minutes as WorkDurationMinutes : null;

// セッション設定の型安全ファクトリー
interface SessionConfigBuilder {
    setWorkDuration(minutes: WorkDurationMinutes): SessionConfigBuilder;
    setBreakDuration(minutes: BreakDurationMinutes): SessionConfigBuilder;
    build(): SessionConfig | null;
}
```

#### 5. AI理解性向上のための型設計

##### 5.1 自己文書化型定義
```go
// 目的が明確な型名
type UserActionContext struct {
    ActionType     UserActionType    `json:"actionType"`
    SessionContext *SessionContext   `json:"sessionContext"`
    Timestamp      ActionTimestamp   `json:"timestamp"`
}

type UserActionType int

const (
    UserActionStartWork UserActionType = iota // ユーザーが作業セッションを開始
    UserActionStartBreak                      // ユーザーが休憩セッションを開始
    UserActionPause                           // ユーザーがセッションを一時停止
    UserActionResume                          // ユーザーがセッションを再開
    UserActionSkipBreak                       // ユーザーが休憩をスキップ
)

// AIが理解しやすいメソッド名と型
func (ctx UserActionContext) IsWorkRelatedAction() bool {
    return ctx.ActionType == UserActionStartWork || ctx.ActionType == UserActionResume
}

func (ctx UserActionContext) RequiresFullscreenMode() bool {
    return ctx.ActionType == UserActionStartWork || ctx.ActionType == UserActionStartBreak
}
```

##### 5.2 意図明確なTypeScript型
```typescript
// AIが推論しやすい明確な型名
interface TimerDisplayState {
    readonly currentPhase: SessionPhase;
    readonly timeDisplay: FormattedTimeString;
    readonly progressIndicator: ProgressPercentage;
    readonly userInteractionEnabled: boolean;
}

type SessionPhase = 
    | 'preparation'      // セッション開始前の準備状態
    | 'active'          // セッション実行中
    | 'paused'          // ユーザーが一時停止中
    | 'transitioning'   // セッション間の遷移中
    | 'completed';      // セッション完了状態

type FormattedTimeString = string & { readonly __format: 'MM:SS' };

// AIが理解しやすいビジネスロジック関数
const determineNextUserAction = (
    currentPhase: SessionPhase,
    sessionType: SessionStateType
): UserActionType[] => {
    switch (currentPhase) {
        case 'preparation':
            return [UserActionType.START_WORK, UserActionType.START_BREAK];
        case 'active':
            return [UserActionType.PAUSE];
        case 'paused':
            return [UserActionType.RESUME];
        case 'completed':
            return sessionType === SESSION_STATE.WORK_SESSION
                ? [UserActionType.START_BREAK, UserActionType.SKIP_BREAK]
                : [UserActionType.START_WORK];
        default:
            return [];
    }
};
```

### 📋 型定義実装チェックリスト

#### 各フェーズでの型定義要件

**Phase 1 型強化**
- [ ] Go基本エンティティの型安全化
- [ ] SessionID・Timestamp等の値オブジェクト
- [ ] 基本的な制約型の実装

**Phase 2 型強化**
- [ ] TypeScript型システムの強化
- [ ] イベントデータの型安全化
- [ ] UIコンポーネントpropsの型定義

**Phase 3 型強化**
- [ ] API境界の型定義完全化
- [ ] 複雑なビジネスロジックの型表現
- [ ] バリデーション型の実装

**Phase 4 型強化**
- [ ] テストデータの型安全化
- [ ] エラー型の体系化
- [ ] パフォーマンス型の測定

**Phase 5 型強化**
- [ ] プロダクション型定義の最終確認
- [ ] 型ドキュメントの整備
- [ ] 型システムの保守性確保

## 📝 開発ワークフロー

### 作業進捗管理
**すべての開発作業において以下のワークフローを必須とする：**

#### 1. タスク開始時
- [ ] TASK.mdの該当チェックボックスを確認
- [ ] 作業内容を明確に把握
- [ ] 関連する既存コードの確認

#### 2. 作業中
- [ ] TASK.mdのチェックボックスをリアルタイム更新
- [ ] 新たな課題や追加作業が発見された場合は即座にTASK.mdに追記
- [ ] 定期的な進捗確認（1時間毎目安）

#### 3. サブタスク完了時
- [ ] ✅ テスト実行確認
- [ ] 🔨 ビルド確認  
- [ ] 🎯 動作確認
- [ ] TASK.mdのチェックボックス更新
- [ ] **意味のある機能単位でコミット実行**

#### 4. フェーズ完了時
- [ ] フェーズ完了確認チェックリストの実行
- [ ] 品質保証要件の完全確認
- [ ] TASK.mdの進捗状況更新
- [ ] **フェーズ完了コミット実行**

### コミット方針

#### コミットタイミング
1. **サブタスク完了時**：機能単位での小さなコミット
2. **フェーズ完了時**：大きな機能群の統合コミット
3. **バグ修正時**：問題解決のタイミングでコミット
4. **リファクタリング時**：コード品質改善のタイミングでコミット

#### コミット前必須確認
- [ ] `go test ./...` 全テスト通過
- [ ] `wails build` ビルド成功
- [ ] 基本機能の動作確認完了
- [ ] TASK.mdの進捗状況が最新

#### コミットメッセージ規則
- **形式**：`動詞 + 概要説明`（例：Add, Update, Fix, Refactor）
- **詳細**：何を実装/修正したかを具体的に記載
- **影響範囲**：変更が及ぼす影響を明記
- **テスト状況**：テスト実行結果を記載

## 開発進捗確認ポイント
- [x] Phase 1完了：基本的なタイマー機能が動作 + 品質保証確認完了
- [ ] Phase 2完了：UIが表示され、基本操作が可能 + 品質保証確認完了
- [ ] Phase 3完了：全画面強制とセッション遷移が動作 + 品質保証確認完了
- [ ] Phase 4完了：安定動作とテスト通過 + 品質保証確認完了
- [ ] Phase 5完了：リリース可能状態 + 最終品質保証確認完了