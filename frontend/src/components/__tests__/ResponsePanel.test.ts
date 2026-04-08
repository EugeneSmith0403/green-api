import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'
import ResponsePanel from '../ResponsePanel.vue'
import type { APIResponse } from '../../types'

describe('ResponsePanel', () => {
  it('shows placeholder when response is null', () => {
    const wrapper = mount(ResponsePanel, { props: { response: null, loading: false } })
    expect(wrapper.text()).toContain('Response will appear here')
    expect(wrapper.find('pre').exists()).toBe(false)
  })

  it('formats JSON data with indentation', () => {
    const response: APIResponse = {
      success: true,
      action: 'getSettings',
      data: { wid: '71234567890@c.us', webhookUrl: '' },
    }
    const wrapper = mount(ResponsePanel, { props: { response, loading: false } })
    expect(wrapper.find('pre').text()).toBe(JSON.stringify(response.data, null, 2))
  })

  it('displays raw text when data is absent', () => {
    const response: APIResponse = {
      success: true,
      action: 'getSettings',
      raw: 'plain text response',
    }
    const wrapper = mount(ResponsePanel, { props: { response, loading: false } })
    expect(wrapper.find('pre').text()).toBe('plain text response')
  })

  it('shows error text on failure', () => {
    const response: APIResponse = {
      success: false,
      action: 'getSettings',
      error: 'upstream returned 401',
    }
    const wrapper = mount(ResponsePanel, { props: { response, loading: false } })
    expect(wrapper.find('pre').text()).toBe('upstream returned 401')
  })

  it('shows loading spinner when loading is true', () => {
    const wrapper = mount(ResponsePanel, { props: { response: null, loading: true } })
    expect(wrapper.find('.spinner').exists()).toBe(true)
    expect(wrapper.find('pre').exists()).toBe(false)
  })

  it('shows success badge on successful response', () => {
    const response: APIResponse = { success: true, action: 'sendMessage', data: {} }
    const wrapper = mount(ResponsePanel, { props: { response, loading: false } })
    expect(wrapper.find('.badge.success').exists()).toBe(true)
  })

  it('shows failure badge on error response', () => {
    const response: APIResponse = { success: false, action: 'sendMessage', error: 'err' }
    const wrapper = mount(ResponsePanel, { props: { response, loading: false } })
    expect(wrapper.find('.badge.failure').exists()).toBe(true)
  })
})
