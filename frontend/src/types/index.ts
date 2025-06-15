// Re-export all types for easy importing
export * from './session';

// Common utility types
export type StringLiteral<T> = T extends string ? string extends T ? never : T : never;