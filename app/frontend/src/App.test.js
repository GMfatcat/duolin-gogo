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
  })

  it('shows settings popout with schedule controls and hides duplicated sidebar cards', async () => {
    const wrapper = mount(App)

    await flushPromises()

    expect(wrapper.text()).not.toContain('通知設定')
    expect(wrapper.text()).not.toContain('送出測試通知')
    expect(wrapper.text()).toContain('下次複習2026-04-05 21:00')
    expect(wrapper.text()).not.toContain('雙語知識卡')

    await wrapper.find('.settings-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('通知設定')
    expect(wrapper.text()).toContain('送出測試通知')
    expect(wrapper.text()).toContain('複習時間')
    expect(wrapper.text()).toContain('通知間隔（分鐘）')

    const numberInput = wrapper.find('input[type="number"]')
    const timeInput = wrapper.find('input[type="time"]')
    await numberInput.setValue('30')
    await timeInput.setValue('20:30')
    await wrapper.find('.save-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Schedule settings updated.')
  })

  it('switches global shell copy and keeps the english summary visually stable', async () => {
    const wrapper = mount(App)

    await flushPromises()

    const summary = wrapper.find('.summary')
    expect(summary.text()).toBe('把你的雙語 Markdown 筆記變成定時提醒、微學習卡片與複習節奏。')

    const languageButtons = wrapper.findAll('.language-toggle button')
    await languageButtons[1].trigger('click')
    await flushPromises()

    expect(summary.text()).toBe('Turn Markdown notes into timely study nudges and review loops.')
    expect(wrapper.text()).toContain('Next review2026-04-05 21:00')
    expect(wrapper.text()).toContain('Concepts to revisit')
    expect(wrapper.text()).toContain('Cherry-pick Purpose')
    expect(wrapper.text()).toContain('Start question')
  })
})
