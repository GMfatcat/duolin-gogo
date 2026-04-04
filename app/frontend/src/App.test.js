import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'
import App from './App.vue'

describe('App', () => {
  it('renders the project shell copy', () => {
    const wrapper = mount(App)

    expect(wrapper.text()).toContain('duolin-gogo')
    expect(wrapper.text()).toContain('Bilingual Git micro-learning')
    expect(wrapper.text()).toContain('TDD mode')
  })
})
