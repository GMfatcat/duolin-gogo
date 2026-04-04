import { flushPromises, mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'
import App from './App.vue'

describe('App', () => {
  it('renders the study card shell and supports language toggle plus answer feedback', async () => {
    const wrapper = mount(App)

    await flushPromises()

    expect(wrapper.text()).toContain('duolin-gogo')
    expect(wrapper.text()).toContain('Bilingual knowledge cards')
    expect(wrapper.text()).toContain('TDD mode')
    expect(wrapper.text()).toContain('Learn session')
    expect(wrapper.text()).toContain('Cherry-pick 的用途')
    expect(wrapper.text()).toContain('哪個 Git 指令可以只拿走一個 commit')
    expect(wrapper.text()).toContain('`git cherry-pick` 會把你指定的一個 commit 套用到目前分支上。')
    expect(wrapper.text()).toContain('Concepts to revisit')
    expect(wrapper.text()).toContain('branching')
    expect(wrapper.text()).toContain('Next review')
    expect(wrapper.text()).toContain('Knowledge file health')
    expect(wrapper.text()).toContain('No import issues detected.')
    expect(wrapper.text()).toContain('Send test notification')
    expect(wrapper.text()).toContain('Snooze 15 min')
    expect(wrapper.text()).toContain('Rescan knowledge')

    const languageButtons = wrapper.findAll('.language-toggle button')
    await languageButtons[1].trigger('click')
    expect(wrapper.text()).toContain('Cherry-pick Purpose')
    expect(wrapper.text()).toContain('One Git command can steal just one commit')
    expect(wrapper.text()).toContain('lets you apply a chosen commit')

    await wrapper.find('input[value="true"]').setValue()
    await wrapper.find('.submit-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Correct.')
    expect(wrapper.text()).toContain('Correct answer: true')

    const toolbarButtons = wrapper.findAll('.toolbar-button')
    await toolbarButtons[0].trigger('click')
    await flushPromises()
    expect(wrapper.text()).toContain('Test notification sent.')

    await toolbarButtons[1].trigger('click')
    await flushPromises()
    expect(wrapper.text()).toContain('Notifications snoozed for 15 minutes.')

    await toolbarButtons[2].trigger('click')
    await flushPromises()
    expect(wrapper.text()).toContain('Knowledge refreshed: 2 cards, 0 errors.')
  })
})
