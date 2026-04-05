import { flushPromises, mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import App from './App.vue'
import { __resetFallbackState } from './api'

async function switchToEnglish(wrapper) {
  await wrapper.find('.language-select select').setValue('en')
  await flushPromises()
}

describe('App', () => {
  beforeEach(() => {
    __resetFallbackState()
  })

  it('renders a staged learn, answer, and feedback flow', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    expect(wrapper.find('.workspace').exists()).toBe(true)
    expect(wrapper.find('.study-column').exists()).toBe(true)
    expect(wrapper.find('.sidebar-column').exists()).toBe(true)
    expect(wrapper.text()).toContain('duolin-gogo')
    expect(wrapper.text()).toContain('Cherry-pick Purpose')
    expect(wrapper.text()).toContain('Start question')
    expect(wrapper.text()).not.toContain('Quick question')
    expect(wrapper.findAll('input[type="radio"]').length).toBe(0)

    await wrapper.find('.phase-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Quick question')
    expect(wrapper.text()).toContain('applies a chosen commit to the current branch')
    expect(wrapper.findAll('input[type="radio"]').length).toBe(2)

    await wrapper.find('input[value="true"]').setValue()
    await wrapper.find('.submit-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Correct.')
    expect(wrapper.text()).toContain('Correct answer: True')
    expect(wrapper.text()).toContain('Next card')
    expect(wrapper.text()).toContain('copies selected commit changes onto your current branch')
  })

  it('shows settings popout with runtime controls only', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)
    await wrapper.find('.settings-button').trigger('click')
    await flushPromises()

    expect(wrapper.find('.settings-layout').exists()).toBe(true)
    expect(wrapper.text()).toContain('Send test notification')
    expect(wrapper.text()).toContain('Rescan knowledge')
    expect(wrapper.text()).toContain('Validate knowledge')
    expect(wrapper.text()).toContain('Reset study data')
    expect(wrapper.find('.preview-select').exists()).toBe(false)
    expect(wrapper.text()).not.toContain('AI draft review')

    const numberInput = wrapper.find('input[type="number"]')
    const timeInputs = wrapper.findAll('input[type="time"]')
    const checkbox = wrapper.find('input[type="checkbox"]')

    await numberInput.setValue('30')
    await timeInputs[0].setValue('20:30')
    await checkbox.setValue(false)
    await wrapper.find('.save-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Schedule settings updated.')
    expect(wrapper.find('.settings-meta').text()).toContain('import OK')
  })

  it('shows the current study streak in the sidebar stats', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    expect(wrapper.text()).toContain('Streak')
    expect(wrapper.text()).toContain('3 days')
  })

  it('uses 20 minutes as the default notification interval', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)
    await wrapper.find('.settings-button').trigger('click')
    await flushPromises()

    const numberInput = wrapper.find('input[type="number"]')
    expect(numberInput.element.value).toBe('20')
  })

  it('validates knowledge without reloading the current study card', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

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
    await switchToEnglish(wrapper)
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
    await switchToEnglish(wrapper)

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

    expect(wrapper.text()).toContain('Review session complete')
    expect(wrapper.text()).toContain('Session summary')
    expect(wrapper.text()).toContain('Answered this batch')
    expect(wrapper.find('.complete-review-button').exists()).toBe(true)

    await wrapper.find('.complete-review-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Cherry-pick Purpose')
  })

  it('forces a short learn break after three answered learn cards', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    wrapper.vm.learnSessionSnapshot = {
      started: true,
      studiedToday: 0,
      correctAnswers: 0,
    }
    wrapper.vm.dashboard = {
      ...wrapper.vm.dashboard,
      stats: {
        studiedToday: 3,
        correctRate: 2 / 3,
      },
      summary: {
        ...wrapper.vm.dashboard.summary,
        weakestDeck: { topic: 'docker', accuracy: 0.67 },
      },
    }
    wrapper.vm.learnSessionProgress = {
      answered: 2,
      total: 3,
      cooldownUntil: '',
    }
    wrapper.vm.phase = 'feedback'
    wrapper.vm.feedback = {
      isCorrect: true,
      correctAnswer: 'true',
    }
    await wrapper.vm.$nextTick()

    await wrapper.find('.next-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Take a short break')
    expect(wrapper.text()).toContain('Next card unlocks at')
    expect(wrapper.text()).toContain('3')
    expect(wrapper.text()).toContain('Session summary')
    expect(wrapper.text()).toContain('67%')
    expect(wrapper.text()).toContain('docker')
    expect(wrapper.text()).toContain('3 days')
  })

  it('shows review session progress cues while a batch is active', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

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
    expect(wrapper.text()).toContain('remaining 2')
  })

  it('shows a new-batch cue after a learn break ends', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    wrapper.vm.learnRestartCue = true
    await wrapper.vm.$nextTick()

    expect(wrapper.text()).toContain('New batch ready')
    expect(wrapper.text()).toContain('Cards are open again')

    await wrapper.find('.phase-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).not.toContain('New batch ready')
  })

  it('switches global shell copy to english without changing the product name', async () => {
    const wrapper = mount(App)

    await flushPromises()

    expect(wrapper.find('.language-select').exists()).toBe(true)

    await switchToEnglish(wrapper)

    expect(wrapper.find('.summary').text()).toBe('Turn notes into study nudges and review loops.')
    expect(wrapper.text()).toContain('Next review')
    expect(wrapper.text()).toContain('Concepts to revisit')
    expect(wrapper.text()).toContain('Cherry-pick Purpose')
    expect(wrapper.text()).toContain('Start question')
    expect(wrapper.text()).toContain('duolin-gogo')
  })

  it('updates the current card when the global mode filter changes', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    expect(wrapper.text()).toContain('Cherry-pick')

    await wrapper.find('.mode-select select').setValue('docker')
    await flushPromises()

    expect(wrapper.text()).toContain('Docker Run')
    expect(wrapper.find('.study-header h2').text()).toContain('Docker Run')
    expect(wrapper.text()).toContain('Focused on docker.')
    expect(wrapper.text()).toContain('docker concepts to revisit')
    expect(wrapper.text()).toContain('docker deck overview')
    expect(wrapper.text()).toContain('3 answers')
  })

  it('switches grouped study modes from the mode selector', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    await wrapper.find('.mode-select select').setValue('backend-tools')
    await flushPromises()

    expect(wrapper.text()).toContain('Focused on Git, Docker, and Linux workflows.')
    expect(wrapper.text()).toContain('docker')
    expect(wrapper.text()).toContain('git')

    await wrapper.find('.mode-select select').setValue('languages')
    await flushPromises()

    expect(wrapper.text()).toContain('Focused on Go and Python language concepts.')
    expect(wrapper.text()).toContain('go')
    expect(wrapper.text()).toContain('python')
  })

  it('shows per-topic progress in mixed mode', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    await wrapper.find('.insights-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Topic progress')
    expect(wrapper.text()).toContain('docker')
    expect(wrapper.text()).toContain('git')
    expect(wrapper.text()).toContain('10 answers')
  })

  it('surfaces weakest deck hints near the top of the shell', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    wrapper.vm.phase = 'feedback'
    wrapper.vm.feedback = {
      isCorrect: false,
      correctAnswer: 'false',
    }
    await wrapper.vm.$nextTick()

    expect(wrapper.find('.assistant-hint').text()).toContain('Lock in the key difference')

    await wrapper.find('.mode-select select').setValue('languages')
    await flushPromises()

    wrapper.vm.phase = 'learn'
    wrapper.vm.feedback = null
    await wrapper.vm.$nextTick()

    expect(wrapper.find('.assistant-hint').text()).toContain('Take the concept in first')

    await wrapper.find('.assistant-collapse').trigger('click')
    await flushPromises()

    expect(wrapper.find('.assistant-hint').classes()).toContain('collapsed')
  })

  it('lets users click DG for a pet-style reaction', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    await wrapper.find('.assistant-hint').trigger('click')
    await flushPromises()

    expect(wrapper.find('.assistant-hint').text()).toContain('I am here. Keep tapping in and I will warm up.')
  })

  it('switches the dg bubble into a review-complete encouragement state', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    wrapper.vm.reviewCompleted = true
    await wrapper.vm.$nextTick()

    const bubble = wrapper.find('.assistant-hint')
    expect(bubble.classes()).toContain('celebration')
    expect(bubble.text()).toContain('That review batch is done.')
  })

  it('uses a focused learn-stage dg prompt before answering', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    const bubble = wrapper.find('.assistant-hint')
    expect(bubble.text()).toContain('Take the concept in first')
  })

  it('switches the dg bubble after a correct answer', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    await wrapper.find('.phase-button').trigger('click')
    await flushPromises()
    await wrapper.find('input[value="true"]').setValue()
    await wrapper.find('.submit-button').trigger('click')
    await flushPromises()

    const bubble = wrapper.find('.assistant-hint')
    expect(bubble.classes()).toContain('celebration')
    expect(wrapper.find('.assistant-avatar').classes()).toContain('pose-nod')
    expect(
      bubble.text().includes('Nice hit') ||
        bubble.text().includes('Good catch') ||
        bubble.text().includes('That was clean') ||
        bubble.text().includes('Yes, that is the feeling.'),
    ).toBe(true)
  })

  it('switches the dg bubble after a wrong answer', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    await wrapper.find('.phase-button').trigger('click')
    await flushPromises()
    await wrapper.find('input[value="false"]').setValue()
    await wrapper.find('.submit-button').trigger('click')
    await flushPromises()

    const bubble = wrapper.find('.assistant-hint')
    expect(bubble.classes()).toContain('warning')
    expect(wrapper.find('.assistant-avatar').classes()).toContain('pose-think')
    expect(
      bubble.text().includes('difference') ||
        bubble.text().includes('almost-right') ||
        bubble.text().includes('Do not worry') ||
        bubble.text().includes('Keep the difference'),
    ).toBe(true)
  })

  it('uses a dedicated dg reaction during a learn break', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

    wrapper.vm.learnSessionSnapshot = {
      started: true,
      studiedToday: 0,
      correctAnswers: 0,
    }
    wrapper.vm.dashboard = {
      ...wrapper.vm.dashboard,
      stats: {
        studiedToday: 3,
        correctRate: 2 / 3,
      },
      summary: {
        ...wrapper.vm.dashboard.summary,
        weakestDeck: { topic: 'docker', accuracy: 0.67 },
      },
    }
    wrapper.vm.learnSessionProgress = {
      answered: 2,
      total: 3,
      cooldownUntil: '',
    }
    wrapper.vm.phase = 'feedback'
    wrapper.vm.feedback = {
      isCorrect: true,
      correctAnswer: 'true',
    }
    await wrapper.vm.$nextTick()

    await wrapper.find('.next-button').trigger('click')
    await flushPromises()

    expect(
      wrapper.find('.assistant-hint').text().includes('Take a short beat') ||
        wrapper.find('.assistant-hint').text().includes('Pause here') ||
        wrapper.find('.assistant-hint').text().includes('That batch landed well') ||
        wrapper.find('.assistant-hint').text().includes('A short pause is right'),
    ).toBe(true)
    expect(wrapper.find('.assistant-avatar').classes()).toContain('pose-rest')
  })

  it('keeps diagnostics collapsed by default and shows severity grouping', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

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

    const disclosure = wrapper.find('.diagnostics-disclosure')
    expect(disclosure.attributes('open')).toBeUndefined()
    expect(wrapper.find('.diagnostics-popout h2').text()).toContain('warnings')
    expect(wrapper.find('.diagnostics-popout h2').text()).toContain('errors')
    expect(wrapper.find('.diagnostics-disclosure summary').text()).not.toContain('import OK')
    expect(wrapper.find('.diagnostics-disclosure summary').text()).not.toContain('warnings')

    await disclosure.element.setAttribute('open', '')
    await wrapper.vm.$nextTick()

    expect(wrapper.text()).toContain('warning')
    expect(wrapper.text()).toContain('error')
    expect(wrapper.text()).toContain('Deck report')
    expect(wrapper.text()).toContain('Total cards')
    expect(wrapper.findAll('.diagnostic-group').length).toBe(2)
  })

  it('filters diagnostics by severity and topic inside batch report', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)

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
    const disclosure = wrapper.find('.diagnostics-disclosure')
    await disclosure.element.setAttribute('open', '')
    await wrapper.vm.$nextTick()

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
    await switchToEnglish(wrapper)
    await wrapper.find('.diagnostics-button').trigger('click')
    await flushPromises()
    const disclosure = wrapper.find('.diagnostics-disclosure')
    await disclosure.element.setAttribute('open', '')
    await wrapper.vm.$nextTick()

    expect(wrapper.text()).toContain('rebase.md')
    expect(wrapper.text()).toContain('2026-04-05 11:45')
    expect(wrapper.text()).toContain('cherry-pick.md')
  })

  it('shows authoring preview controls inside library and updates the preview selection', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)
    await wrapper.find('.library-button').trigger('click')
    await flushPromises()

    expect(wrapper.find('.preview-select').exists()).toBe(true)
    expect(wrapper.find('.diagnostics-disclosure').exists()).toBe(false)
    expect(wrapper.find('.preview-card').text()).toContain('Cherry-pick Purpose')

    await wrapper.find('.preview-select').setValue('D:/duolin-gogo/knowledge/git/rebase.md')
    await flushPromises()

    expect(wrapper.find('.preview-card').text()).toContain('Rebase vs Merge')
  })

  it('reviews pasted AI draft markdown and shows normalized preview', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)
    await wrapper.find('.library-button').trigger('click')
    await flushPromises()

    const draft = `---
id: git-ai-review
title_zh: Git Fetch 草稿
title_en: Git Fetch Draft
type: true-false
question_zh: "git fetch 會直接 merge 到目前分支。"
question_en: "git fetch merges into the current branch."
clickbait_zh: "這個指令看起來很安靜，但多數人都忽略它。"
clickbait_en: "This command looks quiet, but it matters"
review_hint_zh: "fetch 只更新追蹤資訊，不會直接 merge。"
review_hint_en: "Fetch updates tracking refs without merging."
answer: false
---

## zh-TW

git fetch 只會更新遠端追蹤參照，不會直接把變更合併進目前分支。

## en

git fetch only updates remote-tracking refs and does not merge into the current branch.`

    await wrapper.find('.draft-input').setValue(draft)
    const reviewButton = wrapper.findAll('.phase-button').find((button) => button.text().includes('Review draft'))
    await reviewButton.trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Git Fetch Draft')
    expect(wrapper.text()).toContain('git fetch only updates remote-tracking refs and does not merge into the current branch')
  })

  it('reviews multiple pasted drafts as separate batch items', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)
    await wrapper.find('.library-button').trigger('click')
    await flushPromises()

    const batchDraft = `---
id: git-batch-one
title_zh: 第一張
title_en: First Draft
type: true-false
question_zh: "git fetch 會 merge。"
question_en: "git fetch merges."
clickbait_zh: "第一張草稿"
clickbait_en: "First draft"
review_hint_zh: "fetch 不會 merge。"
review_hint_en: "fetch does not merge."
answer: false
---

## zh-TW

第一張草稿內容。

## en

First draft body.

===

---
id: git-batch-two
title: Broken
type: true-false
question: "broken?"
answer: true
---

## zh-TW

Only one language section.`

    await wrapper.find('.draft-input').setValue(batchDraft)
    const reviewButton = wrapper.findAll('.phase-button').find((button) => button.text().includes('Review draft'))
    await reviewButton.trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Draft 1')
    expect(wrapper.text()).toContain('First Draft')
    expect(wrapper.text()).toContain('Draft 2')
    expect(wrapper.text()).toContain('Needs fixes')
    expect(wrapper.text()).toContain('missing_language_section')
    expect(wrapper.text()).toContain('Suggested fix')
    expect(wrapper.text()).toContain('Add both `## zh-TW` and `## en` sections')
    expect(wrapper.findAll('.batch-review-card').length).toBe(2)
    const saveButton = wrapper.findAll('.toolbar-button.secondary').find((button) => button.text().includes('Save valid drafts'))
    expect(saveButton).toBeDefined()
    expect(saveButton.attributes('disabled')).toBeUndefined()
  })

  it('saves a reviewed draft and reports the saved path', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)
    await wrapper.find('.library-button').trigger('click')
    await flushPromises()

    const draft = `---
id: git-ai-review
title_zh: Git Fetch 草稿
title_en: Git Fetch Draft
type: true-false
question_zh: "git fetch 會直接 merge 到目前分支。"
question_en: "git fetch merges into the current branch."
clickbait_zh: "這個指令看起來很安靜，但多數人都忽略它。"
clickbait_en: "This command looks quiet, but it matters"
review_hint_zh: "fetch 只更新追蹤資訊，不會直接 merge。"
review_hint_en: "Fetch updates tracking refs without merging."
answer: false
---

## zh-TW

git fetch 只會更新遠端追蹤參照，不會直接把變更合併進目前分支。

## en

git fetch only updates remote-tracking refs and does not merge into the current branch.`

    await wrapper.find('.draft-input').setValue(draft)
    const reviewButton = wrapper.findAll('.phase-button').find((button) => button.text().includes('Review draft'))
    await reviewButton.trigger('click')
    await flushPromises()
    const secondaryButtons = wrapper.findAll('.toolbar-button.secondary')
    await secondaryButtons[secondaryButtons.length - 1].trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Draft saved to')
    expect(wrapper.text()).toContain('knowledge/git/git-ai-review.md')
    expect(wrapper.findAll('.preview-card')[0].text()).toContain('Git Fetch')
  })

  it('shows the in-app AI prompt panel and copies the prompt', async () => {
    const wrapper = mount(App)
    const writeText = vi.fn().mockResolvedValue(undefined)
    Object.defineProperty(global.navigator, 'clipboard', {
      value: { writeText },
      configurable: true,
    })

    await flushPromises()
    await switchToEnglish(wrapper)
    await wrapper.find('.library-button').trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('AI card prompt')
    expect(wrapper.text()).toContain('You are generating one Markdown study card for duolin-gogo.')

    const promptButton = wrapper.findAll('.phase-button').find((button) => button.text().includes('Copy prompt'))
    await promptButton.trigger('click')
    await flushPromises()

    expect(writeText).toHaveBeenCalled()
    expect(wrapper.text()).toContain('Copied.')
  })

  it('shows a batch import report after saving valid drafts from a mixed batch', async () => {
    const wrapper = mount(App)

    await flushPromises()
    await switchToEnglish(wrapper)
    await wrapper.find('.library-button').trigger('click')
    await flushPromises()

    const batchDraft = `---
id: git-import-one
title_zh: 第一張
title_en: First Import
type: true-false
question_zh: "git fetch 會 merge。"
question_en: "git fetch merges."
clickbait_zh: "第一張"
clickbait_en: "First"
review_hint_zh: "fetch 不會 merge。"
review_hint_en: "fetch does not merge."
answer: false
---

## zh-TW

第一張可存。

## en

First draft can be saved.

===

---
id: git-import-two
title: Broken
type: true-false
question: "broken?"
answer: true
---

## zh-TW

Only one language section.`

    await wrapper.find('.draft-input').setValue(batchDraft)
    const reviewButton = wrapper.findAll('.phase-button').find((button) => button.text().includes('Review draft'))
    await reviewButton.trigger('click')
    await flushPromises()

    const saveButton = wrapper.findAll('.toolbar-button.secondary').find((button) => button.text().includes('Save valid drafts'))
    await saveButton.trigger('click')
    await flushPromises()

    expect(wrapper.text()).toContain('Saved 1 drafts. Skipped 1 drafts.')
    expect(wrapper.text()).toContain('Import report')
    expect(wrapper.text()).toContain('Saved')
    expect(wrapper.text()).toContain('Skipped')
    expect(wrapper.text()).toContain('git-import-one')
  })
})
