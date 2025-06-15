// Event name constants - must match Go backend config/constants.go
export const EVENTS = {
  SESSION_START: 'session:start',
  SESSION_END: 'session:end',
  SESSION_PAUSE: 'session:pause',
  SESSION_RESUME: 'session:resume',
  TIMER_TICK: 'timer:tick',
  WARNING: 'warning',
} as const;

export type EventType = typeof EVENTS[keyof typeof EVENTS];