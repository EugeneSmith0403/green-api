import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'
import MessageForm from '../MessageForm.vue'

describe('MessageForm validation', () => {
  it('shows errors when required fields are empty', async () => {
    const wrapper = mount(MessageForm, { props: { loading: false } })
    await wrapper.find('button').trigger('click')
    expect(wrapper.findAll('.error-msg')).toHaveLength(2)
  })

  it('shows error only for empty chatId', async () => {
    const wrapper = mount(MessageForm, { props: { loading: false } })
    await wrapper.find('textarea').setValue('Hello')
    await wrapper.find('button').trigger('click')
    const errors = wrapper.findAll('.error-msg')
    expect(errors).toHaveLength(1)
    expect(errors[0].text()).toBe('Required')
  })

  it('emits sendMessage with trimmed values when form is valid', async () => {
    const wrapper = mount(MessageForm, { props: { loading: false } })
    await wrapper.find('input').setValue('  71234567890@c.us  ')
    await wrapper.find('textarea').setValue('  Hello  ')
    await wrapper.find('button').trigger('click')
    const emitted = wrapper.emitted('sendMessage')
    expect(emitted).toHaveLength(1)
    expect(emitted![0][0]).toEqual({ chatId: '71234567890@c.us', message: 'Hello' })
  })

  it('does not emit when loading', async () => {
    const wrapper = mount(MessageForm, { props: { loading: true } })
    expect(wrapper.find('button').attributes('disabled')).toBeDefined()
  })
})
