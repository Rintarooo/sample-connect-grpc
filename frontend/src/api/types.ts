// API Error types
export interface ApiError {
  message: string;
  code?: string;
  details?: unknown;
}

// API Response wrapper
export interface ApiResponse<T> {
  data: T;
  error?: ApiError;
}

// API Configuration
export interface ApiConfig {
  baseUrl: string;
  timeout?: number;
}
