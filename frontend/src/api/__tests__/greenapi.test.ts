import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import { getSettings, getStateInstance, sendFileByUrl, sendMessage } from '../greenapi'

const mockFetch = vi.fn()

beforeEach(() => {
  vi.stubGlobal('fetch', mockFetch)
})

afterEach(() => {
  vi.unstubAllGlobals()
  vi.clearAllMocks()
})

function mockResponse(body: unknown) {
  mockFetch.mockResolvedValue({ json: () => Promise.resolve(body) })
}

const id = '1101000001'
const token = 'test-token'

describe('getSettings', () => {
  it('GET /api/green/instance/:id/settings/:token', async () => {
    mockResponse({ success: true, action: 'getSettings', data: { wid: '71234567890@c.us' } })

    await getSettings(id, token)

    expect(mockFetch).toHaveBeenCalledWith(
      `/api/green/instance/${id}/settings/${token}`,
      expect.objectContaining({ method: 'GET' }),
    )
  })
})

describe('getStateInstance', () => {
  it('GET /api/green/instance/:id/state/:token', async () => {
    mockResponse({ success: true, action: 'getStateInstance', data: { stateInstance: 'authorized' } })

    await getStateInstance(id, token)

    expect(mockFetch).toHaveBeenCalledWith(
      `/api/green/instance/${id}/state/${token}`,
      expect.objectContaining({ method: 'GET' }),
    )
  })
})

describe('sendMessage', () => {
  it('POST /api/green/instance/:id/send-message/:token with chatId and message', async () => {
    mockResponse({ success: true, action: 'sendMessage', data: { idMessage: 'BAE5' } })

    await sendMessage(id, token, '71234567890@c.us', 'Hello')

    expect(mockFetch).toHaveBeenCalledWith(
      `/api/green/instance/${id}/send-message/${token}`,
      expect.objectContaining({ method: 'POST' }),
    )
    const body = JSON.parse(mockFetch.mock.calls[0][1].body)
    expect(body.chatId).toBe('71234567890@c.us')
    expect(body.message).toBe('Hello')
  })
})

describe('sendFileByUrl', () => {
  it('POST /api/green/instance/:id/send-file/:token with all fields', async () => {
    mockResponse({ success: true, action: 'sendFileByUrl', data: { idMessage: 'BAE6' } })

    await sendFileByUrl(id, token, '71234567890@c.us', 'https://example.com/file.pdf', 'file.pdf', 'Look at this')

    expect(mockFetch).toHaveBeenCalledWith(
      `/api/green/instance/${id}/send-file/${token}`,
      expect.objectContaining({ method: 'POST' }),
    )
    const body = JSON.parse(mockFetch.mock.calls[0][1].body)
    expect(body.urlFile).toBe('https://example.com/file.pdf')
    expect(body.fileName).toBe('file.pdf')
    expect(body.caption).toBe('Look at this')
  })
})
