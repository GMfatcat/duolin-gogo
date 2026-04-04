import { flushPromises, mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'
import App from './App.vue'

describe('App', () => {
  it('renders a focused workspace with staged learn, answer, and feedback flow', async () => {
    const wrapper = mount(App)

    await flushPromises()

    expect(wrapper.find('.workspace').exists()).toBe(true)
    expect(wrapper.find('.study-column').exists()).toBe(true)
    expect(wrapper.find('.sidebar-column').exists()).toBe(true)

    expect(wrapper.text()).toContain('duolin-gogo')
    expect(wrapper.text()).toContain('Cherry-pick 的用途')
    expect(wrapper.text()).toContain('雙語知識卡')
    expect(wrapper.text()).toContain('開始作答')
    expect(wrapper.text()).not.toContain('快速問題')
    expect(wrapper.findAll('input[type="radio"]').length).toBe(0)

    await wrapper.find('.phase-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('快速問題')
    expect(wrapper.text()).toContain('`git cherry-pick` 會把指定的 commit 套用到目前分支。')
    expect(wrapper.findAll('input[type="radio"]').length).toBe(2)

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

  it('shows a settings popout instead of putting utilities in the main sidebar', async () => {
    const wrapper = mount(App)

    await flushPromises()

    expect(wrapper.text()).not.toContain('通知設定')
    expect(wrapper.text()).not.toContain('知識檔案健康度')
    expect(wrapper.text()).not.toContain('送出測試通知')
    expect(wrapper.text()).not.toContain('內容模式')
    expect(wrapper.text()).not.toContain('語言模式')

    await wrapper.find('.settings-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('通知設定')
    expect(wrapper.text()).toContain('知識檔案健康度')
    expect(wrapper.text()).toContain('送出測試通知')

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

  it('switches global shell copy and formats next review time for humans', async () => {
    const wrapper = mount(App)

    await flushPromises()

    expect(wrapper.text()).toContain('下次複習2026-04-05 21:00')
    expect(wrapper.text()).toContain('概念複習重點')

    const languageButtons = wrapper.findAll('.language-toggle button')
    await languageButtons[1].trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Next review2026-04-05 21:00')
    expect(wrapper.text()).toContain('Concepts to revisit')
    expect(wrapper.text()).toContain('Cherry-pick Purpose')
    expect(wrapper.text()).toContain('Start question')
  })
})
