import { flushPromises, mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'
import App from './App.vue'

describe('App', () => {
  it('renders a staged learn, answer, and feedback flow', async () => {
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
    expect(wrapper.text()).toContain('`git cherry-pick` 會把指定的 commit 套到目前分支。')
    expect(wrapper.findAll('input[type="radio"]').length).toBe(2)

    await wrapper.find('input[value="true"]').setValue()
    await wrapper.find('.submit-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('答對了。')
    expect(wrapper.text()).toContain('正確答案：是')
    expect(wrapper.text()).toContain('下一張卡')
    expect(wrapper.text()).toContain('Cherry-pick 會把指定 commit 的變更套用到目前分支。')
  })

  it('shows settings popout with horizontal controls and active hours fields', async () => {
    const wrapper = mount(App)

    await flushPromises()

    expect(wrapper.text()).not.toContain('送出測試通知')
    expect(wrapper.text()).toContain('下次複習2026-04-05 21:00')

    await wrapper.find('.settings-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('送出測試通知')
    expect(wrapper.text()).toContain('重新掃描知識庫')
    expect(wrapper.text()).toContain('通知間隔（分鐘）')
    expect(wrapper.text()).toContain('複習時間')
    expect(wrapper.text()).toContain('推送時段')
    expect(wrapper.find('.settings-layout').exists()).toBe(true)
    expect(wrapper.findAll('.toolbar-button').length).toBe(3)

    const numberInput = wrapper.find('input[type="number"]')
    const timeInputs = wrapper.findAll('input[type="time"]')
    const checkbox = wrapper.find('input[type="checkbox"]')

    await numberInput.setValue('30')
    await timeInputs[0].setValue('20:30')
    await checkbox.setValue(false)
    await wrapper.find('.save-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Schedule settings updated.')
    expect(wrapper.find('.settings-meta').text()).toContain('匯入正常')
  })

  it('switches global shell copy to english without changing the product name', async () => {
    const wrapper = mount(App)

    await flushPromises()

    const summary = wrapper.find('.summary')
    expect(summary.text()).toBe('把筆記變成定時提醒、微課與複習節奏。')

    const languageButtons = wrapper.findAll('.language-toggle button')
    await languageButtons[1].trigger('click')
    await flushPromises()

    expect(summary.text()).toBe('Turn notes into study nudges and review loops.')
    expect(wrapper.text()).toContain('Next review2026-04-05 21:00')
    expect(wrapper.text()).toContain('Concepts to revisit')
    expect(wrapper.text()).toContain('Cherry-pick Purpose')
    expect(wrapper.text()).toContain('Start question')
    expect(wrapper.text()).toContain('duolin-gogo')
  })

  it('distinguishes import warnings from import errors inside settings diagnostics', async () => {
    const wrapper = mount(App)

    await flushPromises()

    wrapper.vm.dashboard = {
      ...wrapper.vm.dashboard,
      importErrors: [
        {
          source_path: 'D:/duolin-gogo/knowledge/git/legacy.md',
          severity: 'warning',
          code: 'missing_localized_field',
          field: 'title_en',
          message: "Missing localized field 'title_en'; using fallback title value.",
        },
        {
          source_path: 'D:/duolin-gogo/knowledge/git/broken.md',
          severity: 'error',
          code: 'missing_language_section',
          field: 'body',
          message: 'Body must contain both ## zh-TW and ## en sections.',
        },
      ],
    }
    await wrapper.vm.$nextTick()

    await wrapper.find('.settings-button').trigger('click')
    await flushPromises()

    expect(wrapper.find('.settings-meta').text()).toContain('1')
    expect(wrapper.find('.settings-meta').text()).toContain('warning')
    expect(wrapper.find('.settings-meta').text()).toContain('error')
    expect(wrapper.text()).toContain('warning')
    expect(wrapper.text()).toContain('error')
    expect(wrapper.findAll('.severity-pill').length).toBe(2)
  })
})
