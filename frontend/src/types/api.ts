export interface Victim {
  id: number;
  full_name: string;
  cause?: string;
  details?: string;
  created_at: string;
  death_time?: string;
  image_url: string;
  cause_added: boolean;
  details_added: boolean;
}

export interface ApiResponse<T = any> {
  data: T;
  status: number;
  statusText: string;
  headers: any;
  config: any;
}

export interface ErrorResponse {
  message: string;
  status?: number;
  errors?: Record<string, string[]>;
}

export type CauseRequest = {
  cause: string;
};

export type DetailsRequest = {
  details: string;
};