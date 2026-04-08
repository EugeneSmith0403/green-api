export interface APIResponse {
  success: boolean
  action: string
  data?: Record<string, unknown>
  raw?: string
  error?: string
}
