import type { APIResponse } from '../types'

async function get(url: string): Promise<APIResponse> {
  const res = await fetch(url, { method: 'GET' })
  return res.json() as Promise<APIResponse>
}

async function post(url: string, body: object): Promise<APIResponse> {
  const res = await fetch(url, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  })
  return res.json() as Promise<APIResponse>
}

export function getSettings(idInstance: string, apiTokenInstance: string): Promise<APIResponse> {
  return get(`/api/green/instance/${idInstance}/settings/${apiTokenInstance}`)
}

export function getStateInstance(
  idInstance: string,
  apiTokenInstance: string,
): Promise<APIResponse> {
  return get(`/api/green/instance/${idInstance}/state/${apiTokenInstance}`)
}

export function sendMessage(
  idInstance: string,
  apiTokenInstance: string,
  chatId: string,
  message: string,
): Promise<APIResponse> {
  return post(`/api/green/instance/${idInstance}/send-message/${apiTokenInstance}`, {
    chatId,
    message,
  })
}

export function sendFileByUrl(
  idInstance: string,
  apiTokenInstance: string,
  chatId: string,
  urlFile: string,
  fileName?: string,
  caption?: string,
): Promise<APIResponse> {
  return post(`/api/green/instance/${idInstance}/send-file/${apiTokenInstance}`, {
    chatId,
    urlFile,
    fileName,
    caption,
  })
}
