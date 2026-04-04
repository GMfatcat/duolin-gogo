import { flushPromises, mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'
import App from './App.vue'

describe('App', () => {
  it('renders a two-column study workspace with staged learn, answer, and feedback flow', async () => {
    const wrapper = mount(App)

    await flushPromises()

    expect(wrapper.find('.workspace').exists()).toBe(true)
    expect(wrapper.find('.study-column').exists()).toBe(true)
    expect(wrapper.find('.sidebar-column').exists()).toBe(true)

    expect(wrapper.text()).toContain('duolin-gogo')
    expect(wrapper.text()).toContain('雙語知識卡')
    expect(wrapper.text()).toContain('學習模式')
    expect(wrapper.text()).toContain('開始作答')

    expect(wrapper.text()).toContain('Cherry-pick 的用途')
    expect(wrapper.text()).toContain('`git cherry-pick` 可以把指定的 commit 套用到目前分支上。')
    expect(wrapper.text()).not.toContain('快速問題')
    expect(wrapper.findAll('input[type="radio"]').length).toBe(0)

    await wrapper.find('.phase-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('快速問題')
    expect(wrapper.text()).toContain('`git cherry-pick` 會把指定的 commit 套用到目前分支。')
    expect(wrapper.findAll('input[type="radio"]').length).toBe(2)
    expect(wrapper.find('.phase-button').exists()).toBe(false)

    await wrapper.find('input[value="true"]').setValue()
    await wrapper.find('.submit-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('答對了。')
    expect(wrapper.text()).toContain('正確答案：是')
    expect(wrapper.text()).toContain('下一題')
    expect(wrapper.text()).toContain('Cherry-pick 會把選定 commit 的變更複製到目前分支。')

    await wrapper.find('.next-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('開始作答')
    expect(wrapper.text()).not.toContain('答對了。')
  })

  it('switches shell copy and card language through the global language toggle', async () => {
    const wrapper = mount(App)

    await flushPromises()

    expect(wrapper.text()).toContain('通知設定')
    expect(wrapper.text()).toContain('概念複習重點')

    const languageButtons = wrapper.findAll('.language-toggle button')
    await languageButtons[1].trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Bilingual knowledge cards')
    expect(wrapper.text()).toContain('Notification settings')
    expect(wrapper.text()).toContain('Concepts to revisit')
    expect(wrapper.text()).toContain('Cherry-pick Purpose')
    expect(wrapper.text()).toContain('Start question')
  })

  it('supports notification controls and settings updates from the sidebar', async () => {
    const wrapper = mount(App)

    await flushPromises()

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

    const selects = wrapper.findAll('select')
    await selects[0].setValue('chaotic')
    await flushPromises()
    expect(wrapper.text()).toContain('Notification settings updated.')

    await selects[1].setValue('prefer_generated')
    await flushPromises()
    expect(wrapper.text()).toContain('Notification settings updated.')
  })
})
