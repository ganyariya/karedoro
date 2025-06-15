// Application constants
export const APP_CONSTANTS = {
  // Application information
  NAME: 'Karedoro',
  VERSION: '1.0.0',
  TITLE: 'Karedoro ポモドーロタイマー',
  
  // Timer intervals (in milliseconds)
  TIMER_UPDATE_INTERVAL: 1000,
  
  // Warning intervals (in minutes)
  WARNING_IDLE_MINUTES: 5,
  
  // Session durations (in minutes)
  DEFAULT_WORK_DURATION: 25,
  DEFAULT_BREAK_DURATION: 5,
} as const;

// UI Messages
export const UI_MESSAGES = {
  WARNING_IDLE: 'まだ次のセッションを開始していません！',
  WORK_COMPLETED: 'ポモドーロ完了！休憩しましょう！',
  BREAK_COMPLETED: '休憩終了！作業に戻りましょう！',
  SESSION_START_WORK: '作業セッション開始！',
  SESSION_START_BREAK: '休憩セッション開始！',
} as const;

// Button labels
export const BUTTON_LABELS = {
  START_WORK: '作業セッション開始',
  START_BREAK: '休憩セッション開始',
  PAUSE: '一時停止',
  RESUME: '再開',
  SKIP_BREAK: '休憩をスキップして作業セッションへ',
} as const;