import { flushPromises, mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it } from 'vitest'
import App from './App.vue'
import { __resetFallbackState } from './api'

describe('App', () => {
  beforeEach(() => {
    __resetFallbackState()
  })

  it('renders a staged learn, answer, and feedback flow', async () => {
    const wrapper = mount(App)

    await flushPromises()

    expect(wrapper.find('.workspace').exists()).toBe(true)
    expect(wrapper.find('.study-column').exists()).toBe(true)
    expect(wrapper.find('.sidebar-column').exists()).toBe(true)

    expect(wrapper.text()).toContain('duolin-gogo')
    expect(wrapper.text()).toContain('Cherry-pick 指令用途')
    expect(wrapper.text()).toContain('開始作答')
    expect(wrapper.text()).not.toContain('快速問題')
    expect(wrapper.findAll('input[type="radio"]').length).toBe(0)

    await wrapper.find('.phase-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('快速問題')
    expect(wrapper.text()).toContain('`git cherry-pick` 會把選定的 commit 套用到目前分支。')
    expect(wrapper.findAll('input[type="radio"]').length).toBe(2)

    await wrapper.find('input[value="true"]').setValue()
    await wrapper.find('.submit-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('答對了。')
    expect(wrapper.text()).toContain('正確答案：是')
    expect(wrapper.text()).toContain('下一張卡')
    expect(wrapper.text()).toContain('Cherry-pick 會把選定 commit 的變更套用到目前分支。')
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
    expect(wrapper.text()).toContain('檢查題庫格式')
    expect(wrapper.text()).toContain('通知間隔（分鐘）')
    expect(wrapper.text()).toContain('複習時間')
    expect(wrapper.text()).toContain('推送時段')
    expect(wrapper.find('.settings-layout').exists()).toBe(true)
    expect(wrapper.text()).toContain('Reset study data')
    expect(wrapper.findAll('.toolbar-button').length).toBe(5)
    expect(wrapper.find('.preview-select').exists()).toBe(false)
    expect(wrapper.text()).not.toContain('AI')

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

  it('validates knowledge without reloading the current study card', async () => {
    const wrapper = mount(App)

    await flushPromises()

    const currentTitle = wrapper.find('.study-header h2').text()

    await wrapper.find('.settings-button').trigger('click')
    await flushPromises()
    const buttons = wrapper.findAll('.toolbar-button')
    await buttons[3].trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Knowledge validated: 40 cards, 0 diagnostics.')
    expect(wrapper.find('.study-header h2').text()).toBe(currentTitle)
  })

  it('requires confirmation before resetting study data', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await wrapper.find('.settings-button').trigger('click')
    await flushPromises()

    const resetButton = wrapper.findAll('.toolbar-button').find((button) => button.text().includes('Reset study data'))
    await resetButton.trigger('click')
    await flushPromises()

    expect(wrapper.find('.confirm-reset-button').exists()).toBe(true)

    await wrapper.find('.confirm-reset-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Study data reset.')
    expect(wrapper.findAll('.status-card strong')[0].text()).toBe('0')
  })

  it('shows a completion state after the last review card', async () => {
    const wrapper = mount(App)

    await flushPromises()

    wrapper.vm.dashboard = {
      ...wrapper.vm.dashboard,
      reviewMode: true,
      reviewQueue: [wrapper.vm.dashboard.currentCard],
    }
    wrapper.vm.phase = 'feedback'
    wrapper.vm.feedback = {
      isCorrect: true,
      correctAnswer: 'true',
    }
    await wrapper.vm.$nextTick()

    await wrapper.find('.next-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('複習完成')
    expect(wrapper.text()).toContain('這輪小結')
    expect(wrapper.text()).toContain('本輪作答')
    expect(wrapper.find('.complete-review-button').exists()).toBe(true)

    await wrapper.find('.complete-review-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Cherry-pick 指令用途')
  })

  it('shows review session progress cues while a batch is active', async () => {
    const wrapper = mount(App)

    await flushPromises()

    wrapper.vm.dashboard = {
      ...wrapper.vm.dashboard,
      reviewMode: true,
      reviewQueue: [wrapper.vm.dashboard.currentCard, { ...wrapper.vm.dashboard.currentCard, id: 'git-rebase-vs-merge' }],
    }
    wrapper.vm.reviewSessionProgress = {
      active: true,
      total: 3,
      remaining: 2,
    }
    await wrapper.vm.$nextTick()

    expect(wrapper.text()).toContain('1 / 3')
    expect(wrapper.text()).toContain('剩餘 2')
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

  it('updates the current card when the global topic filter changes', async () => {
    const wrapper = mount(App)

    await flushPromises()

    expect(wrapper.text()).toContain('Cherry-pick')

    await wrapper.find('.topic-filter select').setValue('docker')
    await flushPromises()

    expect(wrapper.text()).toContain('docker run')
    expect(wrapper.find('.study-header h2').text()).toContain('docker run')
    expect(wrapper.text()).toContain('目前專注在 docker 主題。')
    expect(wrapper.text()).toContain('docker 建議多看幾次')
    expect(wrapper.text()).toContain('docker 主題概況')
    expect(wrapper.text()).toContain('3 次作答')
  })

  it('switches to preset topic groups from quick pins', async () => {
    const wrapper = mount(App)

    await flushPromises()

    const presetButtons = wrapper.findAll('.topic-preset')
    await presetButtons[1].trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('後端工具')
    expect(wrapper.text()).toContain('Git、Docker 與 Linux')
    expect(wrapper.text()).toContain('docker')
    expect(wrapper.text()).toContain('git')

    await presetButtons[2].trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('程式語言')
    expect(wrapper.text()).toContain('Go 與 Python')
    expect(wrapper.text()).toContain('go')
    expect(wrapper.text()).toContain('python')
  })

  it('shows per-topic progress in mixed mode', async () => {
    const wrapper = mount(App)

    await flushPromises()

    expect(wrapper.text()).toContain('各主題概況')
    expect(wrapper.text()).toContain('docker')
    expect(wrapper.text()).toContain('git')
    expect(wrapper.text()).toContain('10 次作答')
  })

  it('distinguishes import warnings from import errors inside diagnostics popout', async () => {
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

    await wrapper.find('.diagnostics-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('題庫報告')
    expect(wrapper.text()).toContain('warning')
    expect(wrapper.text()).toContain('error')
    expect(wrapper.text()).toContain('題庫報告')
    expect(wrapper.text()).toContain('總卡片數')
    expect(wrapper.text()).toContain('乾淨卡片')
    expect(wrapper.findAll('.severity-pill').length).toBe(5)
    expect(wrapper.findAll('.diagnostic-group').length).toBe(2)
  })

  it('filters diagnostics by severity and topic inside batch report', async () => {
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

    await wrapper.find('.diagnostics-button').trigger('click')
    await flushPromises()

    const selects = wrapper.findAll('.diagnostic-filter')
    await selects[0].setValue('warning')
    await flushPromises()
    expect(wrapper.text()).toContain('missing_localized_field')
    expect(wrapper.text()).not.toContain('missing_language_section')

    await selects[1].setValue('git')
    await flushPromises()
    expect(wrapper.text()).toContain('legacy.md')
  })

  it('shows recently changed cards inside batch report', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await wrapper.find('.diagnostics-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('rebase.md')
    expect(wrapper.text()).toContain('2026-04-05 11:45')
    expect(wrapper.text()).toContain('cherry-pick.md')
  })

  it('shows authoring preview controls inside library and updates the preview selection', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await wrapper.find('.library-button').trigger('click')
    await flushPromises()

    expect(wrapper.find('.preview-select').exists()).toBe(true)
    expect(wrapper.find('.diagnostics-disclosure').exists()).toBe(false)
    expect(wrapper.find('.preview-card').text()).toContain('Cherry-pick')

    await wrapper.find('.preview-select').setValue('D:/duolin-gogo/knowledge/git/rebase.md')
    await flushPromises()

    expect(wrapper.find('.preview-card').text()).toContain('Rebase')
  })

  it('reviews pasted AI draft markdown and shows normalized preview', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await wrapper.find('.library-button').trigger('click')
    await flushPromises()

    const draft = `---
id: git-ai-review
title_zh: Git Fetch 草稿
title_en: Git Fetch Draft
type: true-false
question_zh: "git fetch 會直接 merge 到目前分支。"
question_en: "git fetch merges into the current branch."
clickbait_zh: "這個指令看起來沒做事，但很多人第一步會按它"
clickbait_en: "This command looks quiet, but it matters"
review_hint_zh: "fetch 只更新追蹤資訊，不會直接 merge。"
review_hint_en: "Fetch updates tracking refs without merging."
answer: false
---

## zh-TW

git fetch 只會更新遠端追蹤資訊，不會直接合併到目前分支。

## en

git fetch only updates remote-tracking refs and does not merge into the current branch.`

    await wrapper.find('.draft-input').setValue(draft)
    await wrapper.findAll('.phase-button')[1].trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Git Fetch')
    expect(wrapper.text()).toContain('fetch')
  })

  it('saves a reviewed draft and reports the saved path', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await wrapper.find('.library-button').trigger('click')
    await flushPromises()

    const draft = `---
id: git-ai-review
title_zh: Git Fetch 草稿
title_en: Git Fetch Draft
type: true-false
question_zh: "git fetch 會直接 merge 到目前分支。"
question_en: "git fetch merges into the current branch."
clickbait_zh: "這個指令看起來沒做事，但很多人第一步會按它"
clickbait_en: "This command looks quiet, but it matters"
review_hint_zh: "fetch 只更新追蹤資訊，不會直接 merge。"
review_hint_en: "Fetch updates tracking refs without merging."
answer: false
---

## zh-TW

git fetch 只會更新遠端追蹤資訊，不會直接合併到目前分支。

## en

git fetch only updates remote-tracking refs and does not merge into the current branch.`

    await wrapper.find('.draft-input').setValue(draft)
    await wrapper.findAll('.phase-button')[1].trigger('click')
    await flushPromises()
    const secondaryButtons = wrapper.findAll('.toolbar-button.secondary')
    await secondaryButtons[secondaryButtons.length - 1].trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Draft saved to')
    expect(wrapper.text()).toContain('knowledge/git/git-ai-review.md')
    expect(wrapper.findAll('.preview-card')[0].text()).toContain('Git Fetch 草稿')
  })
})
