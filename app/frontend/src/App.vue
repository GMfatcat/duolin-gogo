<script setup>
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { EventsOn } from '../wailsjs/runtime/runtime'
import {
  getStudyCard,
  loadDashboard,
  rescanKnowledge,
  sendTestNotification,
  snoozeNotifications,
  submitAnswer,
  updateNotificationSettings,
  updatePreferredLanguage,
} from './api'

const translations = {
  'zh-TW': {
    summary: '把你的雙語 Markdown 筆記變成定時提醒、微學習卡片與複習節奏。',
    bilingualCards: '雙語知識卡',
    reviewCard: '複習卡',
    nextCard: '下一張卡',
    reviewSessionActive: '目前正在複習，佇列數量：',
    learnPhase: '閱讀知識',
    answerPhase: '開始作答',
    feedbackPhase: '作答回饋',
    quickQuestion: '快速問題',
    submitAnswer: '送出答案',
    checking: '判斷中...',
    nextStep: '下一題',
    correctPrefix: '正確答案：',
    noCardsLabel: '目前沒有可用卡片',
    noCardsTitle: '知識匯入需要處理',
    noCardsBody: '請新增有效的雙語知識卡，或透過右上角設定檢查匯入問題。',
    studiedToday: '今日作答',
    correctRate: '正答率',
    reviewQueue: '複習佇列',
    nextReview: '下次複習',
    notScheduled: '尚未排程',
    weakTopicsLabel: '弱勢主題',
    weakTopicsTitle: '概念複習重點',
    noWeakTopics: '目前還沒有弱勢主題，繼續學習後會逐漸出現。',
    notificationSettings: '通知設定',
    hookMode: 'Hook 模式',
    style: '風格',
    titleSource: '標題來源',
    diagnosticsLabel: '匯入診斷',
    diagnosticsTitle: '知識檔案健康度',
    noDiagnostics: '目前沒有匯入問題。',
    sendTestNotification: '送出測試通知',
    snoozeNotifications: '稍後 15 分鐘',
    rescanKnowledge: '重新掃描知識庫',
    loading: '載入中',
    preparingCard: '正在準備下一張卡片...',
    shellLanguageUpdated: 'Language updated.',
    correctFeedback: '答對了。',
    incorrectFeedback: '再想一下。',
    trueLabel: '是',
    falseLabel: '否',
    accuracySuffix: '正確率',
    settingsLabel: '設定',
    close: '關閉',
  },
  en: {
    summary: 'Turn your bilingual Markdown notes into timed nudges, micro-lessons, and review rhythm.',
    bilingualCards: 'Bilingual knowledge cards',
    reviewCard: 'Review card',
    nextCard: 'Next card',
    reviewSessionActive: 'Review session active. Queue size:',
    learnPhase: 'Learn',
    answerPhase: 'Start question',
    feedbackPhase: 'Feedback',
    quickQuestion: 'Quick question',
    submitAnswer: 'Submit answer',
    checking: 'Checking...',
    nextStep: 'Next card',
    correctPrefix: 'Correct answer:',
    noCardsLabel: 'No cards available',
    noCardsTitle: 'Knowledge import needs attention',
    noCardsBody: 'Add valid bilingual knowledge cards, or inspect import issues from the settings popout.',
    studiedToday: 'Studied today',
    correctRate: 'Correct rate',
    reviewQueue: 'Review queue',
    nextReview: 'Next review',
    notScheduled: 'Not scheduled',
    weakTopicsLabel: 'Weak topics',
    weakTopicsTitle: 'Concepts to revisit',
    noWeakTopics: 'No weak topics yet. Keep studying to generate insights.',
    notificationSettings: 'Notification settings',
    hookMode: 'Hook mode',
    style: 'Style',
    titleSource: 'Title source',
    diagnosticsLabel: 'Import diagnostics',
    diagnosticsTitle: 'Knowledge file health',
    noDiagnostics: 'No import issues detected.',
    sendTestNotification: 'Send test notification',
    snoozeNotifications: 'Snooze 15 min',
    rescanKnowledge: 'Rescan knowledge',
    loading: 'Loading',
    preparingCard: 'Preparing the next card...',
    shellLanguageUpdated: 'Language updated.',
    correctFeedback: 'Correct.',
    incorrectFeedback: 'Not quite.',
    trueLabel: 'True',
    falseLabel: 'False',
    accuracySuffix: 'accuracy',
    settingsLabel: 'Settings',
    close: 'Close',
  },
}

const dashboard = ref(null)
const selectedLanguage = ref('zh-TW')
const selectedAnswer = ref('')
const feedback = ref(null)
const actionMessage = ref('')
const loading = ref(true)
const submitting = ref(false)
const savingNotificationSettings = ref(false)
const changingLanguage = ref(false)
const phase = ref('learn')
const settingsOpen = ref(false)
let unsubscribe = null

const card = computed(() => dashboard.value?.currentCard ?? null)
const stats = computed(() => dashboard.value?.stats ?? { studiedToday: 0, correctRate: 0 })
const reviewMode = computed(() => dashboard.value?.reviewMode ?? false)
const reviewQueue = computed(() => dashboard.value?.reviewQueue ?? [])
const summary = computed(() => dashboard.value?.summary ?? { nextReviewAt: '', weakTopics: [] })
const importErrors = computed(() => dashboard.value?.importErrors ?? [])
const notificationSettings = computed(() =>
  dashboard.value?.notificationSettings ?? { style: 'playful', titleMode: 'prefer_manual' },
)
const t = computed(() => translations[selectedLanguage.value] ?? translations['zh-TW'])

const currentPhaseLabel = computed(() => {
  if (phase.value === 'answer') return t.value.answerPhase
  if (phase.value === 'feedback') return t.value.feedbackPhase
  return t.value.learnPhase
})

const titleText = computed(() => {
  if (!card.value) return ''
  return selectedLanguage.value === 'en' ? card.value.titleEn : card.value.titleZh
})

const questionText = computed(() => {
  if (!card.value) return ''
  return selectedLanguage.value === 'en' ? card.value.questionTextEn : card.value.questionTextZh
})

const clickbaitText = computed(() => {
  if (!card.value) return ''
  return selectedLanguage.value === 'en' ? card.value.clickbaitEn : card.value.clickbaitZh
})

const localizedChoices = computed(() =>
  (card.value?.choices ?? []).map((choice) => ({
    value: choice.value,
    label: selectedLanguage.value === 'en' ? choice.labelEn : choice.labelZh,
  })),
)

const explanation = computed(() => {
  if (!card.value) return ''
  return selectedLanguage.value === 'en' ? card.value.explanationEn : card.value.explanationZh
})

const reviewHintText = computed(() => {
  if (!card.value) return ''
  return selectedLanguage.value === 'en' ? card.value.reviewHintEn : card.value.reviewHintZh
})

const formattedCorrectRate = computed(() => `${Math.round((stats.value.correctRate ?? 0) * 100)}%`)
const nextReviewText = computed(() => formatDisplayTime(summary.value.nextReviewAt, t.value.notScheduled))

const correctAnswerLabel = computed(() => {
  if (!feedback.value) return ''
  const matchingChoice = localizedChoices.value.find((choice) => choice.value === feedback.value.correctAnswer)
  if (matchingChoice) return matchingChoice.label
  if (feedback.value.correctAnswer === 'true') return t.value.trueLabel
  if (feedback.value.correctAnswer === 'false') return t.value.falseLabel
  return feedback.value.correctAnswer
})

const feedbackMessage = computed(() => {
  if (!feedback.value) return ''
  return feedback.value.isCorrect ? t.value.correctFeedback : t.value.incorrectFeedback
})

onMounted(async () => {
  await refreshDashboard()

  if (typeof window !== 'undefined' && typeof window.runtime !== 'undefined') {
    EventsOn('notification:open-card', async (cardId) => {
      const nextCard = await getStudyCard(cardId)
      dashboard.value = {
        ...dashboard.value,
        currentCard: nextCard,
      }
      resetStudyFlow()
      actionMessage.value = `Opened from notification: ${cardId}`
    })
    unsubscribe = true
  }
})

onUnmounted(() => {
  unsubscribe = null
})

async function refreshDashboard() {
  dashboard.value = await loadDashboard()
  selectedLanguage.value = dashboard.value.preferredLanguage || dashboard.value.info.defaultLanguage
  resetStudyFlow()
  loading.value = false
}

function resetStudyFlow() {
  phase.value = 'learn'
  feedback.value = null
  selectedAnswer.value = ''
}

function formatDisplayTime(value, fallback) {
  if (!value) return fallback

  const parsed = new Date(value)
  if (Number.isNaN(parsed.getTime())) return fallback

  const year = parsed.getFullYear()
  const month = `${parsed.getMonth() + 1}`.padStart(2, '0')
  const day = `${parsed.getDate()}`.padStart(2, '0')
  const hours = `${parsed.getHours()}`.padStart(2, '0')
  const minutes = `${parsed.getMinutes()}`.padStart(2, '0')
  return `${year}-${month}-${day} ${hours}:${minutes}`
}

async function handleSubmit() {
  if (!card.value || !selectedAnswer.value) return

  submitting.value = true
  const result = await submitAnswer({
    cardId: card.value.id,
    sessionType: reviewMode.value ? 'review' : 'learn',
    selectedAnswer: selectedAnswer.value,
    shownAt: card.value.shownAt,
  })
  feedback.value = result
  dashboard.value = {
    ...dashboard.value,
    stats: result.stats,
  }
  phase.value = 'feedback'
  submitting.value = false
}

async function handleNextCard() {
  await refreshDashboard()
}

async function handleSendTestNotification() {
  try {
    const result = await sendTestNotification()
    actionMessage.value = result.message
  } catch (error) {
    actionMessage.value = `Notification failed: ${error?.message ?? String(error)}`
  }
}

async function handleSnooze() {
  try {
    const result = await snoozeNotifications()
    actionMessage.value = result.message
  } catch (error) {
    actionMessage.value = `Snooze failed: ${error?.message ?? String(error)}`
  }
}

async function handleRescanKnowledge() {
  try {
    const result = await rescanKnowledge()
    await refreshDashboard()
    actionMessage.value = result.message
  } catch (error) {
    actionMessage.value = `Rescan failed: ${error?.message ?? String(error)}`
  }
}

async function handleNotificationSettingChange(field, value) {
  try {
    savingNotificationSettings.value = true
    const nextSettings = {
      style: notificationSettings.value.style,
      titleMode: notificationSettings.value.titleMode,
      [field]: value,
    }
    const result = await updateNotificationSettings(nextSettings)
    dashboard.value = {
      ...dashboard.value,
      notificationSettings: nextSettings,
    }
    actionMessage.value = result.message
  } catch (error) {
    actionMessage.value = `Settings update failed: ${error?.message ?? String(error)}`
  } finally {
    savingNotificationSettings.value = false
  }
}

async function handleLanguageChange(language) {
  if (selectedLanguage.value === language) return

  try {
    changingLanguage.value = true
    const result = await updatePreferredLanguage(language)
    selectedLanguage.value = language
    if (dashboard.value) {
      dashboard.value = {
        ...dashboard.value,
        preferredLanguage: language,
      }
    }
    actionMessage.value = result.message || t.value.shellLanguageUpdated
  } catch (error) {
    actionMessage.value = `Language update failed: ${error?.message ?? String(error)}`
  } finally {
    changingLanguage.value = false
  }
}

function toggleSettings() {
  settingsOpen.value = !settingsOpen.value
}
</script>

<template>
  <main class="shell">
    <section class="hero">
      <div class="hero-copy">
        <p class="eyebrow">duolin-gogo</p>
        <h1>duolin-gogo</h1>
        <p class="summary">{{ t.summary }}</p>
      </div>

      <div class="hero-actions">
        <div class="language-toggle hero-toggle">
          <button
            :class="{ active: selectedLanguage === 'zh-TW' }"
            :disabled="changingLanguage"
            type="button"
            @click="handleLanguageChange('zh-TW')"
          >
            zh-TW
          </button>
          <button
            :class="{ active: selectedLanguage === 'en' }"
            :disabled="changingLanguage"
            type="button"
            @click="handleLanguageChange('en')"
          >
            en
          </button>
        </div>

        <button class="settings-button" type="button" @click="toggleSettings">
          ⚙
        </button>
      </div>
    </section>

    <section class="workspace">
      <div class="study-column">
        <section v-if="loading" class="study-card emphasis">
          <p class="label">{{ t.loading }}</p>
          <strong>{{ t.preparingCard }}</strong>
        </section>

        <section v-else-if="card" class="study-card emphasis">
          <div class="study-header">
            <div>
              <p class="label">{{ reviewMode ? t.reviewCard : t.nextCard }}</p>
              <h2>{{ titleText }}</h2>
            </div>
            <span class="phase-pill">{{ currentPhaseLabel }}</span>
          </div>

          <p v-if="reviewMode" class="review-banner">
            {{ t.reviewSessionActive }} {{ reviewQueue.length }}
          </p>
          <p class="callout">{{ clickbaitText }}</p>

          <div v-if="phase === 'learn'" class="phase-panel">
            <p class="explanation">{{ explanation }}</p>
            <button class="phase-button" type="button" @click="phase = 'answer'">
              {{ t.answerPhase }}
            </button>
          </div>

          <div v-else-if="phase === 'answer'" class="phase-panel">
            <div class="question-block">
              <p class="label">{{ t.quickQuestion }}</p>
              <h3>{{ questionText }}</h3>

              <div class="answers">
                <label v-for="choice in localizedChoices" :key="choice.value" class="answer-option">
                  <input v-model="selectedAnswer" type="radio" name="answer" :value="choice.value">
                  <span>{{ choice.label }}</span>
                </label>
              </div>

              <button class="submit-button" type="button" :disabled="!selectedAnswer || submitting" @click="handleSubmit">
                {{ submitting ? t.checking : t.submitAnswer }}
              </button>
            </div>
          </div>

          <div v-else class="phase-panel">
            <div class="feedback" :class="{ correct: feedback?.isCorrect, wrong: !feedback?.isCorrect }">
              <strong>{{ feedbackMessage }}</strong>
              <p>{{ t.correctPrefix }}{{ correctAnswerLabel }}</p>
              <p>{{ reviewHintText }}</p>
            </div>
            <p class="explanation feedback-explanation">{{ explanation }}</p>
            <button class="next-button" type="button" @click="handleNextCard">
              {{ t.nextStep }}
            </button>
          </div>
        </section>

        <section v-else class="study-card emphasis">
          <div class="study-header">
            <div>
              <p class="label">{{ t.noCardsLabel }}</p>
              <h2>{{ t.noCardsTitle }}</h2>
            </div>
          </div>
          <p class="explanation">{{ t.noCardsBody }}</p>
        </section>
      </div>

      <aside class="sidebar-column">
        <section class="status-card lead-card">
          <span class="label">{{ t.bilingualCards }}</span>
          <strong>{{ t.nextReview }}</strong>
          <p class="sidebar-time">{{ nextReviewText }}</p>
        </section>

        <section class="status-grid compact">
          <article class="status-card">
            <span class="label">{{ t.studiedToday }}</span>
            <strong>{{ stats.studiedToday }}</strong>
          </article>
          <article class="status-card">
            <span class="label">{{ t.correctRate }}</span>
            <strong>{{ formattedCorrectRate }}</strong>
          </article>
          <article class="status-card">
            <span class="label">{{ t.reviewQueue }}</span>
            <strong>{{ reviewQueue.length }}</strong>
          </article>
          <article class="status-card">
            <span class="label">{{ t.nextReview }}</span>
            <strong>{{ nextReviewText }}</strong>
          </article>
        </section>

        <section class="study-card sidebar-panel">
          <div class="study-header">
            <div>
              <p class="label">{{ t.weakTopicsLabel }}</p>
              <h2>{{ t.weakTopicsTitle }}</h2>
            </div>
          </div>

          <div v-if="summary.weakTopics.length" class="weak-topics">
            <article v-for="topic in summary.weakTopics" :key="topic.tag" class="topic-chip">
              <strong>{{ topic.tag }}</strong>
              <span>{{ Math.round(topic.accuracy * 100) }}% {{ t.accuracySuffix }}</span>
            </article>
          </div>
          <p v-else class="explanation">{{ t.noWeakTopics }}</p>
        </section>
      </aside>
    </section>

    <div v-if="settingsOpen" class="settings-overlay" @click.self="toggleSettings">
      <section class="settings-popout">
        <div class="study-header">
          <div>
            <p class="label">{{ t.notificationSettings }}</p>
            <h2>{{ t.settingsLabel }}</h2>
          </div>
          <button class="close-button" type="button" @click="toggleSettings">
            {{ t.close }}
          </button>
        </div>

        <section class="study-card inset-card">
          <div class="toolbar">
            <button class="toolbar-button" type="button" @click="handleSendTestNotification">
              {{ t.sendTestNotification }}
            </button>
            <button class="toolbar-button secondary" type="button" @click="handleSnooze">
              {{ t.snoozeNotifications }}
            </button>
            <button class="toolbar-button secondary" type="button" @click="handleRescanKnowledge">
              {{ t.rescanKnowledge }}
            </button>
          </div>
          <span v-if="actionMessage" class="toolbar-message">{{ actionMessage }}</span>
        </section>

        <section class="study-card inset-card">
          <div class="study-header">
            <div>
              <p class="label">{{ t.notificationSettings }}</p>
              <h2>{{ t.hookMode }}</h2>
            </div>
          </div>

          <div class="settings-grid">
            <label class="settings-field">
              <span>{{ t.style }}</span>
              <select
                :value="notificationSettings.style"
                :disabled="savingNotificationSettings"
                @change="handleNotificationSettingChange('style', $event.target.value)"
              >
                <option value="safe">safe</option>
                <option value="playful">playful</option>
                <option value="aggressive">aggressive</option>
                <option value="chaotic">chaotic</option>
              </select>
            </label>

            <label class="settings-field">
              <span>{{ t.titleSource }}</span>
              <select
                :value="notificationSettings.titleMode"
                :disabled="savingNotificationSettings"
                @change="handleNotificationSettingChange('titleMode', $event.target.value)"
              >
                <option value="prefer_manual">prefer_manual</option>
                <option value="prefer_generated">prefer_generated</option>
              </select>
            </label>
          </div>
        </section>

        <section class="study-card inset-card">
          <div class="study-header">
            <div>
              <p class="label">{{ t.diagnosticsLabel }}</p>
              <h2>{{ t.diagnosticsTitle }}</h2>
            </div>
          </div>

          <div v-if="importErrors.length" class="diagnostics-list">
            <article v-for="item in importErrors" :key="`${item.source_path}-${item.code}`" class="diagnostic-item">
              <strong>{{ item.code }}</strong>
              <p>{{ item.message }}</p>
              <span>{{ item.source_path }}</span>
            </article>
          </div>
          <p v-else class="explanation">{{ t.noDiagnostics }}</p>
        </section>
      </section>
    </div>
  </main>
</template>
