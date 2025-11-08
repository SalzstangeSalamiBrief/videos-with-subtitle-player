import type { ApiError } from './ApiError';

export interface IApiResponse<T> {
  data?: T;
  error?: ApiError;
}
