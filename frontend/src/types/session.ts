// Session state constants - must match Go backend domain/session.go
export const SESSION_STATE = {
  IDLE: 'Idle',
  WORK_SESSION: 'WorkSession',
  BREAK_SESSION: 'BreakSession',
} as const;

export type SessionStateType = typeof SESSION_STATE[keyof typeof SESSION_STATE];

// Session data interfaces
export interface SessionEventData {
  state: SessionStateType;
  duration?: number;
  remainingTime?: number;
}

export interface TimerTickEventData {
  state: SessionStateType;
  remainingTime: number;
}

export interface WarningEventData {
  idleDuration: number;
}

// Session configuration
export interface SessionConfig {
  workDuration: number;    // in minutes
  breakDuration: number;   // in minutes
}

// Utility functions for session state
export const isWorkSession = (state: SessionStateType): boolean => 
  state === SESSION_STATE.WORK_SESSION;

export const isBreakSession = (state: SessionStateType): boolean => 
  state === SESSION_STATE.BREAK_SESSION;

export const isIdleState = (state: SessionStateType): boolean => 
  state === SESSION_STATE.IDLE;