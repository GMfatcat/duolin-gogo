import { flushPromises, mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'
import App from './App.vue'

describe('App', () => {
  it('renders the study card shell and supports language toggle plus answer feedback', async () => {
    const wrapper = mount(App)

    await flushPromises()

    expect(wrapper.text()).toContain('duolin-gogo')
    expect(wrapper.text()).toContain('Bilingual Git micro-learning')
    expect(wrapper.text()).toContain('TDD mode')
    expect(wrapper.text()).toContain('Cherry-pick Purpose')
    expect(wrapper.text()).toContain('可以把某一個特定 commit')

    const languageButtons = wrapper.findAll('.language-toggle button')
    await languageButtons[1].trigger('click')
    expect(wrapper.text()).toContain('lets you apply a chosen commit')

    await wrapper.find('input[value="true"]').setValue()
    await wrapper.find('.submit-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Correct.')
    expect(wrapper.text()).toContain('Correct answer: true')
  })
})
