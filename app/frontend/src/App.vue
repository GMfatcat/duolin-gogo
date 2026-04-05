<script setup>
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { EventsOn } from '../wailsjs/runtime/runtime'
import {
  getStudyCard,
  loadAuthoringPreview,
  loadDashboard,
  previewKnowledgeCard,
  reviewDraft,
  rescanKnowledge,
  resetStudyData,
  saveDraft,
  sendTestNotification,
  snoozeNotifications,
  submitAnswer,
  updateNotificationSettings,
  updatePreferredLanguage,
  updateScheduleSettings,
  validateKnowledge,
} from './api'

const translations = {
  'zh-TW': {
    summary: '把筆記變成定時提醒、微課與複習節奏。',
    reviewCard: '複習卡',
    nextCard: '下一張卡',
    reviewSessionActive: '複習模式進行中，待答題數：',
    reviewProgress: '複習進度',
    remainingCards: '剩餘',
    sessionSummaryTitle: '這輪小結',
    sessionAnswered: '本輪作答',
    sessionAccuracy: '本輪正答率',
    sessionWeakTopic: '建議回頭看',
    noSessionWeakTopic: '目前沒有特別弱勢的主題',
    learnPhase: '先看觀念',
    answerPhase: '開始作答',
    feedbackPhase: '作答回饋',
    reviewCompleteTitle: '複習完成',
    reviewCompleteBody: '這一輪複習已經做完了，先收一下成果，再決定要不要繼續新卡。',
    reviewCompleteAction: '回到下一張卡',
    quickQuestion: '快速問題',
    submitAnswer: '送出答案',
    checking: '檢查中...',
    nextStep: '下一張卡',
    correctPrefix: '正確答案：',
    noCardsLabel: '目前沒有卡片',
    noCardsTitle: '知識匯入需要留意',
    noCardsBody: '請加入有效的雙語知識卡，或從設定面板查看匯入問題。',
    studiedToday: '今日作答',
    correctRate: '正答率',
    reviewQueue: '複習佇列',
    nextReview: '下次複習',
    notScheduled: '尚未安排',
    weakTopicsLabel: '弱勢主題',
    weakTopicsTitle: '建議多看幾次',
    noWeakTopics: '目前沒有特別弱勢的主題，繼續作答就會慢慢產生洞察。',
    notificationSettings: '設定',
    hookMode: '通知 hook',
    style: '風格',
    titleSource: '標題來源',
    scheduleLabel: '排程',
    reviewTime: '複習時間',
    intervalMinutes: '通知間隔（分鐘）',
    activeHours: '推送時段',
    activeHoursEnabled: '限制在這段時間內推送',
    activeHoursStart: '開始時間',
    activeHoursEnd: '結束時間',
    diagnosticsTitle: '匯入診斷',
    noDiagnostics: '目前沒有匯入問題。',
    importHealthOk: '匯入正常',
    importHealthWarnings: '個警告',
    importHealthErrors: '個錯誤',
    sendTestNotification: '送出測試通知',
    snoozeNotifications: '延後 15 分鐘',
    rescanKnowledge: '重新掃描知識庫',
    validateKnowledge: '檢查題庫格式',
    resetStudyData: 'Reset study data',
    resetWarningTitle: '重置學習紀錄',
    resetWarningBody: '這會清空本地的作答紀錄、今日統計與複習進度。這個動作無法復原。',
    cancel: '取消',
    confirmReset: '確認重置',
    loading: '載入中',
    preparingCard: '正在準備下一張卡...',
    shellLanguageUpdated: '語言已更新。',
    correctFeedback: '答對了。',
    incorrectFeedback: '這題還差一點。',
    trueLabel: '是',
    falseLabel: '否',
    accuracySuffix: '正答率',
    settingsLabel: '設定',
    libraryLabel: '書庫',
    diagnosticsLabel: '診斷',
    close: '關閉',
    saveSchedule: '儲存排程',
    openSettings: '開啟設定',
    toolsLabel: '工具',
    warningLabel: 'warning',
    errorLabel: 'error',
    deckReport: '題庫報告',
    totalCards: '總卡片數',
    cleanCards: '乾淨卡片',
    warningCards: '有警告的卡片',
    errorCards: '有錯誤的卡片',
    recentChanges: '最近變動卡片',
    lastUpdated: '最後更新',
    noRecentChanges: '目前沒有最近變動資料。',
    severityFilter: '嚴重度',
    topicFilter: '主題',
    allFilter: '全部',
    authoringPreview: '作者預覽',
    previewFile: '卡片檔案',
    previewDiagnostics: '這張卡的診斷',
    noPreviewDiagnostics: '這張卡目前沒有額外診斷。',
    aiDraftReview: 'AI 草稿審查',
    draftInput: '貼上 AI 產生的 Markdown',
    reviewDraft: '審查草稿',
    draftPreview: '草稿預覽',
    noDraftYet: '貼上草稿後，就可以在這裡看到 normalized preview 與診斷。',
    saveDraft: '儲存草稿',
    draftTopic: '主題資料夾',
  },
  en: {
    summary: 'Turn notes into study nudges and review loops.',
    reviewCard: 'Review card',
    nextCard: 'Next card',
    reviewSessionActive: 'Review session active. Queue size:',
    reviewProgress: 'Review progress',
    remainingCards: 'remaining',
    sessionSummaryTitle: 'Session summary',
    sessionAnswered: 'Answered this batch',
    sessionAccuracy: 'Batch accuracy',
    sessionWeakTopic: 'Topic to revisit',
    noSessionWeakTopic: 'No standout weak topic right now.',
    learnPhase: 'Learn',
    answerPhase: 'Start question',
    feedbackPhase: 'Feedback',
    reviewCompleteTitle: 'Review session complete',
    reviewCompleteBody: 'You finished this review batch. Take a beat, then jump back into the next card when ready.',
    reviewCompleteAction: 'Back to the next card',
    quickQuestion: 'Quick question',
    submitAnswer: 'Submit answer',
    checking: 'Checking...',
    nextStep: 'Next card',
    correctPrefix: 'Correct answer: ',
    noCardsLabel: 'No cards available',
    noCardsTitle: 'Knowledge import needs attention',
    noCardsBody: 'Add valid bilingual knowledge cards, or inspect import issues from the settings panel.',
    studiedToday: 'Studied today',
    correctRate: 'Correct rate',
    reviewQueue: 'Review queue',
    nextReview: 'Next review',
    notScheduled: 'Not scheduled',
    weakTopicsLabel: 'Weak topics',
    weakTopicsTitle: 'Concepts to revisit',
    noWeakTopics: 'No weak topics yet. Keep studying to generate insights.',
    notificationSettings: 'Settings',
    hookMode: 'Notification hooks',
    style: 'Style',
    titleSource: 'Title source',
    scheduleLabel: 'Schedule',
    reviewTime: 'Review time',
    intervalMinutes: 'Notification interval (minutes)',
    activeHours: 'Active hours',
    activeHoursEnabled: 'Only notify during this time range',
    activeHoursStart: 'Start time',
    activeHoursEnd: 'End time',
    diagnosticsTitle: 'Import diagnostics',
    noDiagnostics: 'No import issues detected.',
    importHealthOk: 'import OK',
    importHealthWarnings: 'warnings',
    importHealthErrors: 'errors',
    sendTestNotification: 'Send test notification',
    snoozeNotifications: 'Snooze 15 min',
    rescanKnowledge: 'Rescan knowledge',
    validateKnowledge: 'Validate knowledge',
    resetStudyData: 'Reset study data',
    resetWarningTitle: 'Reset study data',
    resetWarningBody: 'This clears your local progress history, today stats, and review scheduling progress. This cannot be undone.',
    cancel: 'Cancel',
    confirmReset: 'Confirm reset',
    loading: 'Loading',
    preparingCard: 'Preparing the next card...',
    shellLanguageUpdated: 'Language updated.',
    correctFeedback: 'Correct.',
    incorrectFeedback: 'Not quite.',
    trueLabel: 'True',
    falseLabel: 'False',
    accuracySuffix: 'accuracy',
    settingsLabel: 'Settings',
    libraryLabel: 'Library',
    diagnosticsLabel: 'Diagnostics',
    close: 'Close',
    saveSchedule: 'Save schedule',
    openSettings: 'Open settings',
    toolsLabel: 'Tools',
    warningLabel: 'warning',
    errorLabel: 'error',
    deckReport: 'Deck report',
    totalCards: 'Total cards',
    cleanCards: 'Clean cards',
    warningCards: 'Cards with warnings',
    errorCards: 'Cards with errors',
    recentChanges: 'Recently changed cards',
    lastUpdated: 'Last updated',
    noRecentChanges: 'No recent changes yet.',
    severityFilter: 'Severity',
    topicFilter: 'Topic',
    allFilter: 'All',
    authoringPreview: 'Authoring preview',
    previewFile: 'Card file',
    previewDiagnostics: 'Diagnostics for this card',
    noPreviewDiagnostics: 'No diagnostics for this card.',
    aiDraftReview: 'AI draft review',
    draftInput: 'Paste AI-generated Markdown',
    reviewDraft: 'Review draft',
    draftPreview: 'Draft preview',
    noDraftYet: 'Paste a draft to inspect the normalized preview and diagnostics here.',
    saveDraft: 'Save draft',
    draftTopic: 'Topic folder',
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
const savingScheduleSettings = ref(false)
const changingLanguage = ref(false)
const phase = ref('learn')
const settingsOpen = ref(false)
const libraryOpen = ref(false)
const diagnosticsOpen = ref(false)
const resetWarningOpen = ref(false)
const reviewCompleted = ref(false)
const reviewSessionProgress = ref({
  active: false,
  total: 0,
  remaining: 0,
})
const reviewSessionSnapshot = ref({
  started: false,
  studiedToday: 0,
  correctAnswers: 0,
})
const sessionSummary = ref({
  visible: false,
  answered: 0,
  accuracy: 0,
  weakTopic: '',
})
const diagnosticsFilter = ref({
  severity: 'all',
  topic: 'all',
})
const authoringPreview = ref({
  files: [],
  selectedPath: '',
  currentCard: null,
  importErrors: [],
})
const draftReview = ref({
  raw: '',
  currentCard: null,
  importErrors: [],
  topic: 'git',
})
const scheduleForm = ref({
  notificationIntervalMinutes: 10,
  reviewTime: '21:00',
  activeHoursEnabled: true,
  activeHoursStart: '09:00',
  activeHoursEnd: '22:00',
})
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
const scheduleSettings = computed(() =>
  dashboard.value?.scheduleSettings ?? {
    notificationIntervalMinutes: 10,
    reviewTime: '21:00',
    activeHoursEnabled: true,
    activeHoursStart: '09:00',
    activeHoursEnd: '22:00',
  },
)
const previewCard = computed(() => authoringPreview.value?.currentCard ?? null)
const previewFiles = computed(() => authoringPreview.value?.files ?? [])
const previewErrors = computed(() => authoringPreview.value?.importErrors ?? [])
const draftReviewErrors = computed(() => draftReview.value?.importErrors ?? [])
const draftReviewCard = computed(() => draftReview.value?.currentCard ?? null)
const t = computed(() => translations[selectedLanguage.value] ?? translations['zh-TW'])
const reviewProgressText = computed(() => {
  if (!reviewSessionProgress.value.active || reviewSessionProgress.value.total <= 0) return ''
  const completed = reviewSessionProgress.value.total - reviewSessionProgress.value.remaining
  return `${completed} / ${reviewSessionProgress.value.total}`
})

const currentPhaseLabel = computed(() => {
  if (phase.value === 'answer') return t.value.answerPhase
  if (phase.value === 'feedback') return t.value.feedbackPhase
  return t.value.learnPhase
})

const warningCount = computed(
  () => importErrors.value.filter((item) => (item.severity || 'error') === 'warning').length,
)
const errorCount = computed(
  () => importErrors.value.filter((item) => (item.severity || 'error') === 'error').length,
)
const warningItems = computed(() =>
  importErrors.value.filter((item) => (item.severity || 'error') === 'warning'),
)
const errorItems = computed(() =>
  importErrors.value.filter((item) => (item.severity || 'error') === 'error'),
)

const titleText = computed(() =>
  card.value ? (selectedLanguage.value === 'en' ? card.value.titleEn : card.value.titleZh) : '',
)
const questionText = computed(() =>
  card.value ? (selectedLanguage.value === 'en' ? card.value.questionTextEn : card.value.questionTextZh) : '',
)
const clickbaitText = computed(() =>
  card.value ? (selectedLanguage.value === 'en' ? card.value.clickbaitEn : card.value.clickbaitZh) : '',
)
const explanation = computed(() =>
  card.value ? (selectedLanguage.value === 'en' ? card.value.explanationEn : card.value.explanationZh) : '',
)
const reviewHintText = computed(() =>
  card.value ? (selectedLanguage.value === 'en' ? card.value.reviewHintEn : card.value.reviewHintZh) : '',
)

const localizedChoices = computed(() =>
  (card.value?.choices ?? []).map((choice) => ({
    value: choice.value,
    label: selectedLanguage.value === 'en' ? choice.labelEn : choice.labelZh,
  })),
)

const formattedCorrectRate = computed(() => `${Math.round((stats.value.correctRate ?? 0) * 100)}%`)
const nextReviewText = computed(() => formatDisplayTime(summary.value.nextReviewAt, t.value.notScheduled))
const diagnosticsSummary = computed(() => {
  if (warningCount.value === 0 && errorCount.value === 0) {
    return `(${t.value.importHealthOk})`
  }

  const parts = []
  if (warningCount.value > 0) {
    parts.push(`${warningCount.value} ${t.value.importHealthWarnings}`)
  }
  if (errorCount.value > 0) {
    parts.push(`${errorCount.value} ${t.value.importHealthErrors}`)
  }
  return `(${parts.join(', ')})`
})
const activeHoursSummary = computed(() => {
  if (!scheduleForm.value.activeHoursEnabled) {
    return t.value.notScheduled
  }
  return `${scheduleForm.value.activeHoursStart} - ${scheduleForm.value.activeHoursEnd}`
})
const diagnosticsBySource = computed(() => {
  const grouped = new Map()
  for (const item of importErrors.value) {
    const key = item.source_path || item.sourcePath || ''
    const entry = grouped.get(key) ?? { hasWarning: false, hasError: false }
    if ((item.severity || 'error') === 'warning') {
      entry.hasWarning = true
    } else {
      entry.hasError = true
    }
    grouped.set(key, entry)
  }
  return grouped
})
const totalCardsCount = computed(() => previewFiles.value.length)
const warningCardCount = computed(() => {
  let count = 0
  diagnosticsBySource.value.forEach((entry) => {
    if (entry.hasWarning) count += 1
  })
  return count
})
const errorCardCount = computed(() => {
  let count = 0
  diagnosticsBySource.value.forEach((entry) => {
    if (entry.hasError) count += 1
  })
  return count
})
const cleanCardCount = computed(() => {
  const dirty = new Set(diagnosticsBySource.value.keys())
  return Math.max(0, totalCardsCount.value - dirty.size)
})
const availableTopics = computed(() => {
  const set = new Set()
  for (const item of importErrors.value) {
    const match = (item.source_path || '').match(/knowledge\/([^/]+)\//)
    if (match?.[1]) {
      set.add(match[1])
    }
  }
  return ['all', ...Array.from(set).sort()]
})
const filteredWarningItems = computed(() =>
  warningItems.value.filter((item) => matchesDiagnosticFilter(item)),
)
const filteredErrorItems = computed(() =>
  errorItems.value.filter((item) => matchesDiagnosticFilter(item)),
)
const recentlyChangedFiles = computed(() =>
  [...previewFiles.value]
    .filter((file) => file.modifiedAt)
    .sort((left, right) => new Date(right.modifiedAt).getTime() - new Date(left.modifiedAt).getTime())
    .slice(0, 5),
)

const correctAnswerLabel = computed(() => {
  if (!feedback.value) return ''
  const matchingChoice = localizedChoices.value.find((choice) => choice.value === feedback.value.correctAnswer)
  if (matchingChoice) return matchingChoice.label
  if (feedback.value.correctAnswer === 'true') return t.value.trueLabel
  if (feedback.value.correctAnswer === 'false') return t.value.falseLabel
  return feedback.value.correctAnswer
})

const feedbackMessage = computed(() =>
  feedback.value ? (feedback.value.isCorrect ? t.value.correctFeedback : t.value.incorrectFeedback) : '',
)

function severityLabel(item) {
  return (item.severity || 'error') === 'warning' ? t.value.warningLabel : t.value.errorLabel
}

function diagnosticTopic(item) {
  const match = (item.source_path || '').match(/knowledge\/([^/]+)\//)
  return match?.[1] || 'unknown'
}

function matchesDiagnosticFilter(item) {
  const severity = item.severity || 'error'
  const severityPass =
    diagnosticsFilter.value.severity === 'all' || diagnosticsFilter.value.severity === severity
  const topicPass =
    diagnosticsFilter.value.topic === 'all' || diagnosticsFilter.value.topic === diagnosticTopic(item)
  return severityPass && topicPass
}

onMounted(async () => {
  await refreshDashboard()
  await refreshAuthoringPreview()

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
  const nextDashboard = await loadDashboard()
  if (nextDashboard.reviewMode) {
    if (!reviewSessionProgress.value.active) {
      reviewSessionSnapshot.value = {
        started: true,
        studiedToday: nextDashboard.stats?.studiedToday ?? 0,
        correctAnswers: estimatedCorrectAnswers(nextDashboard.stats),
      }
      sessionSummary.value = {
        visible: false,
        answered: 0,
        accuracy: 0,
        weakTopic: '',
      }
    }
    const nextTotal =
      reviewSessionProgress.value.active && reviewSessionProgress.value.total >= nextDashboard.reviewQueue.length
        ? reviewSessionProgress.value.total
        : nextDashboard.reviewQueue.length
    reviewSessionProgress.value = {
      active: true,
      total: nextTotal,
      remaining: nextDashboard.reviewQueue.length,
    }
  } else if (!reviewCompleted.value) {
    reviewSessionProgress.value = {
      active: false,
      total: 0,
      remaining: 0,
    }
    reviewSessionSnapshot.value = {
      started: false,
      studiedToday: 0,
      correctAnswers: 0,
    }
  }
  dashboard.value = nextDashboard
  selectedLanguage.value = dashboard.value.preferredLanguage || dashboard.value.info.defaultLanguage
  scheduleForm.value = {
    notificationIntervalMinutes: scheduleSettings.value.notificationIntervalMinutes,
    reviewTime: scheduleSettings.value.reviewTime,
    activeHoursEnabled: scheduleSettings.value.activeHoursEnabled,
    activeHoursStart: scheduleSettings.value.activeHoursStart,
    activeHoursEnd: scheduleSettings.value.activeHoursEnd,
  }
  resetStudyFlow()
  loading.value = false
}

async function refreshAuthoringPreview(selectedPath = '') {
  authoringPreview.value = selectedPath
    ? await previewKnowledgeCard(selectedPath)
    : await loadAuthoringPreview()
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

function estimatedCorrectAnswers(dayStats) {
  return Math.round((dayStats?.studiedToday ?? 0) * (dayStats?.correctRate ?? 0))
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
  const wasReviewMode = reviewMode.value
  const previousQueueLength = reviewQueue.value.length
  const previousReviewTotal = reviewSessionProgress.value.total
  await refreshDashboard()
  if (wasReviewMode && previousQueueLength > 0 && !dashboard.value.reviewMode) {
    const answered = Math.max(0, (dashboard.value.stats?.studiedToday ?? 0) - reviewSessionSnapshot.value.studiedToday)
    const correctAnswers =
      Math.max(0, estimatedCorrectAnswers(dashboard.value.stats) - reviewSessionSnapshot.value.correctAnswers)
    reviewCompleted.value = true
    reviewSessionProgress.value = {
      active: true,
      total: previousReviewTotal || previousQueueLength,
      remaining: 0,
    }
    sessionSummary.value = {
      visible: true,
      answered: answered || previousReviewTotal || previousQueueLength,
      accuracy: answered > 0 ? correctAnswers / answered : dashboard.value.stats?.correctRate ?? 0,
      weakTopic: dashboard.value.summary?.weakTopics?.[0]?.tag ?? '',
    }
  }
}

function handleReturnToLearning() {
  reviewCompleted.value = false
  reviewSessionProgress.value = {
    active: false,
    total: 0,
    remaining: 0,
  }
  reviewSessionSnapshot.value = {
    started: false,
    studiedToday: 0,
    correctAnswers: 0,
  }
  sessionSummary.value = {
    visible: false,
    answered: 0,
    accuracy: 0,
    weakTopic: '',
  }
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
    await refreshAuthoringPreview(authoringPreview.value.selectedPath)
    actionMessage.value = result.message
  } catch (error) {
    actionMessage.value = `Rescan failed: ${error?.message ?? String(error)}`
  }
}

async function handleValidateKnowledge() {
  try {
    const result = await validateKnowledge()
    dashboard.value = {
      ...dashboard.value,
      importErrors: result.importErrors ?? [],
    }
    await refreshAuthoringPreview(authoringPreview.value.selectedPath)
    actionMessage.value = result.message
  } catch (error) {
    actionMessage.value = `Validate failed: ${error?.message ?? String(error)}`
  }
}

function openResetWarning() {
  resetWarningOpen.value = true
}

function closeResetWarning() {
  resetWarningOpen.value = false
}

async function handleResetStudyData() {
  try {
    const result = await resetStudyData()
    await refreshDashboard()
    closeResetWarning()
    actionMessage.value = result.message
  } catch (error) {
    actionMessage.value = `Reset failed: ${error?.message ?? String(error)}`
  }
}

async function handlePreviewSelection(path) {
  await refreshAuthoringPreview(path)
}

async function handleDraftReview() {
  const result = await reviewDraft(draftReview.value.raw)
  draftReview.value = {
    ...draftReview.value,
    currentCard: result.currentCard ?? null,
    importErrors: result.importErrors ?? [],
  }
}

async function handleSaveDraft() {
  try {
    const result = await saveDraft({
      raw: draftReview.value.raw,
      topic: draftReview.value.topic,
    })
    await rescanKnowledge()
    await refreshDashboard()
    await refreshAuthoringPreview(result.savedPath)
    actionMessage.value = result.message
  } catch (error) {
    actionMessage.value = `Save draft failed: ${error?.message ?? String(error)}`
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

async function handleScheduleSave() {
  try {
    savingScheduleSettings.value = true
    const nextSettings = {
      notificationIntervalMinutes: Number(scheduleForm.value.notificationIntervalMinutes),
      reviewTime: scheduleForm.value.reviewTime,
      activeHoursEnabled: scheduleForm.value.activeHoursEnabled,
      activeHoursStart: scheduleForm.value.activeHoursStart,
      activeHoursEnd: scheduleForm.value.activeHoursEnd,
    }
    const result = await updateScheduleSettings(nextSettings)
    dashboard.value = {
      ...dashboard.value,
      scheduleSettings: nextSettings,
    }
    actionMessage.value = result.message
  } catch (error) {
    actionMessage.value = `Schedule update failed: ${error?.message ?? String(error)}`
  } finally {
    savingScheduleSettings.value = false
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

function toggleLibrary() {
  libraryOpen.value = !libraryOpen.value
}

function toggleDiagnostics() {
  diagnosticsOpen.value = !diagnosticsOpen.value
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

        <button class="library-button" type="button" :aria-label="t.libraryLabel" @click="toggleLibrary">
          <svg class="settings-icon" viewBox="0 0 24 24" aria-hidden="true">
            <path
              d="M4 5.5A2.5 2.5 0 0 1 6.5 3H20v16.5a1.5 1.5 0 0 0-1.5-1.5H6.5A2.5 2.5 0 0 0 4 20.5V5.5Zm4 1.5h8"
              fill="none"
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="1.8"
            />
            <path
              d="M8 11h8M8 15h5"
              fill="none"
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="1.8"
            />
          </svg>
        </button>

        <button class="diagnostics-button" type="button" :aria-label="t.diagnosticsLabel" @click="toggleDiagnostics">
          <svg class="settings-icon" viewBox="0 0 24 24" aria-hidden="true">
            <path
              d="M5 19V9M12 19V5M19 19v-8"
              fill="none"
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="1.8"
            />
            <circle cx="5" cy="7" r="2" fill="currentColor" />
            <circle cx="12" cy="3" r="2" fill="currentColor" />
            <circle cx="19" cy="9" r="2" fill="currentColor" />
          </svg>
        </button>

        <button class="settings-button" type="button" :aria-label="t.openSettings" @click="toggleSettings">
          <svg class="settings-icon" viewBox="0 0 24 24" aria-hidden="true">
            <path
              d="M19.14 12.94c.04-.31.06-.63.06-.94s-.02-.63-.06-.94l2.03-1.58a.5.5 0 0 0 .12-.64l-1.92-3.32a.5.5 0 0 0-.6-.22l-2.39.96a7.2 7.2 0 0 0-1.63-.94l-.36-2.54a.5.5 0 0 0-.5-.42h-3.84a.5.5 0 0 0-.5.42l-.36 2.54c-.58.23-1.12.54-1.63.94l-2.39-.96a.5.5 0 0 0-.6.22L2.71 8.84a.5.5 0 0 0 .12.64l2.03 1.58c-.04.31-.06.63-.06.94s.02.63.06.94L2.83 14.52a.5.5 0 0 0-.12.64l1.92 3.32a.5.5 0 0 0 .6.22l2.39-.96c.5.39 1.05.71 1.63.94l.36 2.54a.5.5 0 0 0 .5.42h3.84a.5.5 0 0 0 .5-.42l.36-2.54c.58-.23 1.13-.55 1.63-.94l2.39.96a.5.5 0 0 0 .6-.22l1.92-3.32a.5.5 0 0 0-.12-.64l-2.03-1.58ZM12 15.5A3.5 3.5 0 1 1 12 8.5a3.5 3.5 0 0 1 0 7Z"
              fill="currentColor"
            />
          </svg>
        </button>
      </div>
    </section>

    <section class="workspace">
      <div class="study-column">
        <section v-if="loading" class="study-card emphasis">
          <p class="label">{{ t.loading }}</p>
          <strong>{{ t.preparingCard }}</strong>
        </section>

        <section v-else-if="reviewCompleted" class="study-card emphasis completion-card">
          <div class="study-header">
            <div>
              <p class="label">{{ t.reviewCard }}</p>
              <h2>{{ t.reviewCompleteTitle }}</h2>
            </div>
            <span class="phase-pill">{{ t.feedbackPhase }}</span>
          </div>
          <p v-if="reviewSessionProgress.total" class="review-progress-line">
            {{ t.reviewProgress }} {{ reviewProgressText }}
          </p>
          <p class="callout">{{ t.reviewCompleteBody }}</p>
          <div v-if="sessionSummary.visible" class="session-summary">
            <p class="label">{{ t.sessionSummaryTitle }}</p>
            <div class="session-summary-grid">
              <article class="status-card batch-stat">
                <span class="label">{{ t.sessionAnswered }}</span>
                <strong>{{ sessionSummary.answered }}</strong>
              </article>
              <article class="status-card batch-stat">
                <span class="label">{{ t.sessionAccuracy }}</span>
                <strong>{{ Math.round((sessionSummary.accuracy ?? 0) * 100) }}%</strong>
              </article>
            </div>
            <p class="explanation session-summary-note">
              {{ t.sessionWeakTopic }}:
              <strong>{{ sessionSummary.weakTopic || t.noSessionWeakTopic }}</strong>
            </p>
          </div>
          <button class="next-button complete-review-button" type="button" @click="handleReturnToLearning">
            {{ t.reviewCompleteAction }}
          </button>
        </section>

        <section v-else-if="card" class="study-card emphasis">
          <div class="study-header">
            <div>
              <p class="label">{{ reviewMode ? t.reviewCard : t.nextCard }}</p>
              <h2>{{ titleText }}</h2>
            </div>
            <span class="phase-pill">{{ currentPhaseLabel }}</span>
          </div>

          <p v-if="reviewMode" class="review-banner">{{ t.reviewSessionActive }} {{ reviewQueue.length }}</p>
          <p v-if="reviewMode && reviewSessionProgress.active" class="review-progress-line">
            {{ t.reviewProgress }} {{ reviewProgressText }} · {{ t.remainingCards }} {{ reviewSessionProgress.remaining }}
          </p>
          <p class="callout">{{ clickbaitText }}</p>

          <div v-if="phase === 'learn'" class="phase-panel">
            <p class="explanation">{{ explanation }}</p>
            <button class="phase-button" type="button" @click="phase = 'answer'">{{ t.answerPhase }}</button>
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
            <button class="next-button" type="button" @click="handleNextCard">{{ t.nextStep }}</button>
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
          <article class="status-card review-highlight">
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
            <h2>{{ t.settingsLabel }} <span class="settings-meta">{{ diagnosticsSummary }}</span></h2>
          </div>
          <button class="close-button" type="button" @click="toggleSettings">{{ t.close }}</button>
        </div>

        <div class="settings-layout">
          <section class="study-card inset-card tools-panel">
            <div class="study-header">
              <div>
                <h2>{{ t.toolsLabel }}</h2>
              </div>
            </div>

            <div class="toolbar">
              <button class="toolbar-button" type="button" @click="handleSendTestNotification">{{ t.sendTestNotification }}</button>
              <button class="toolbar-button secondary" type="button" @click="handleSnooze">{{ t.snoozeNotifications }}</button>
              <button class="toolbar-button secondary" type="button" @click="handleRescanKnowledge">{{ t.rescanKnowledge }}</button>
              <button class="toolbar-button secondary" type="button" @click="handleValidateKnowledge">{{ t.validateKnowledge }}</button>
              <button class="toolbar-button danger" type="button" @click="openResetWarning">{{ t.resetStudyData }}</button>
            </div>
            <span v-if="actionMessage" class="toolbar-message">{{ actionMessage }}</span>
          </section>

          <section class="study-card inset-card">
            <div class="study-header">
              <div>
                <h2>{{ t.hookMode }}</h2>
              </div>
            </div>

            <div class="settings-grid">
              <label class="settings-field">
                <span>{{ t.style }}</span>
                <select :value="notificationSettings.style" :disabled="savingNotificationSettings" @change="handleNotificationSettingChange('style', $event.target.value)">
                  <option value="safe">safe</option>
                  <option value="playful">playful</option>
                  <option value="aggressive">aggressive</option>
                  <option value="chaotic">chaotic</option>
                </select>
              </label>

              <label class="settings-field">
                <span>{{ t.titleSource }}</span>
                <select :value="notificationSettings.titleMode" :disabled="savingNotificationSettings" @change="handleNotificationSettingChange('titleMode', $event.target.value)">
                  <option value="prefer_manual">prefer_manual</option>
                  <option value="prefer_generated">prefer_generated</option>
                </select>
              </label>
            </div>
          </section>

          <section class="study-card inset-card">
            <div class="study-header">
              <div>
                <h2>{{ t.scheduleLabel }}</h2>
              </div>
              <span class="phase-pill schedule-pill">{{ activeHoursSummary }}</span>
            </div>

            <div class="settings-grid schedule-grid">
              <label class="settings-field">
                <span>{{ t.intervalMinutes }}</span>
                <input v-model="scheduleForm.notificationIntervalMinutes" type="number" min="5" max="120">
              </label>

              <label class="settings-field">
                <span>{{ t.reviewTime }}</span>
                <input v-model="scheduleForm.reviewTime" type="time">
              </label>

              <label class="settings-field settings-field-wide checkbox-field">
                <span>{{ t.activeHours }}</span>
                <label class="toggle-line">
                  <input v-model="scheduleForm.activeHoursEnabled" type="checkbox">
                  <span>{{ t.activeHoursEnabled }}</span>
                </label>
              </label>

              <label class="settings-field">
                <span>{{ t.activeHoursStart }}</span>
                <input v-model="scheduleForm.activeHoursStart" type="time" :disabled="!scheduleForm.activeHoursEnabled">
              </label>

              <label class="settings-field">
                <span>{{ t.activeHoursEnd }}</span>
                <input v-model="scheduleForm.activeHoursEnd" type="time" :disabled="!scheduleForm.activeHoursEnabled">
              </label>
            </div>

            <button class="phase-button save-button" type="button" :disabled="savingScheduleSettings" @click="handleScheduleSave">
              {{ t.saveSchedule }}
            </button>
          </section>
        </div>

      </section>
    </div>

    <div v-if="libraryOpen" class="settings-overlay" @click.self="toggleLibrary">
      <section class="settings-popout library-popout">
        <div class="study-header">
          <div>
            <h2>{{ t.libraryLabel }}</h2>
          </div>
          <button class="close-button" type="button" @click="toggleLibrary">{{ t.close }}</button>
        </div>

        <div class="settings-layout library-layout">
          <section class="study-card inset-card preview-panel">
            <div class="study-header">
              <div>
                <h2>{{ t.authoringPreview }}</h2>
              </div>
            </div>

            <label class="settings-field">
              <span>{{ t.previewFile }}</span>
              <select
                class="preview-select"
                :value="authoringPreview.selectedPath"
                @change="handlePreviewSelection($event.target.value)"
              >
                <option v-for="file in previewFiles" :key="file.path" :value="file.path">
                  {{ file.name }}
                </option>
              </select>
            </label>

            <div v-if="previewCard" class="preview-card">
              <p class="label">{{ previewCard.id }}</p>
              <strong>{{ selectedLanguage === 'en' ? previewCard.titleEn : previewCard.titleZh }}</strong>
              <p class="callout compact">{{ selectedLanguage === 'en' ? previewCard.clickbaitEn : previewCard.clickbaitZh }}</p>
              <p class="explanation compact">{{ selectedLanguage === 'en' ? previewCard.explanationEn : previewCard.explanationZh }}</p>
              <p class="label">{{ selectedLanguage === 'en' ? previewCard.questionTextEn : previewCard.questionTextZh }}</p>
            </div>

            <div class="preview-diagnostics">
              <span class="label">{{ t.previewDiagnostics }}</span>
              <div v-if="previewErrors.length" class="diagnostics-list compact">
                <article
                  v-for="item in previewErrors"
                  :key="`preview-${item.source_path}-${item.code}-${item.field || ''}`"
                  class="diagnostic-item"
                  :class="(item.severity || 'error') === 'warning' ? 'warning' : 'error'"
                >
                  <div class="diagnostic-head">
                    <span class="severity-pill" :class="(item.severity || 'error') === 'warning' ? 'warning' : 'error'">
                      {{ severityLabel(item) }}
                    </span>
                    <strong>{{ item.code }}</strong>
                  </div>
                  <p>{{ item.message }}</p>
                </article>
              </div>
              <p v-else class="explanation">{{ t.noPreviewDiagnostics }}</p>
            </div>
          </section>

          <section class="study-card inset-card preview-panel">
            <div class="study-header">
              <div>
                <h2>{{ t.aiDraftReview }}</h2>
              </div>
            </div>

            <label class="settings-field">
              <span>{{ t.draftInput }}</span>
              <textarea
                v-model="draftReview.raw"
                class="draft-input"
                rows="12"
                spellcheck="false"
              />
            </label>

            <label class="settings-field">
              <span>{{ t.draftTopic }}</span>
              <select v-model="draftReview.topic" class="draft-topic-select">
                <option value="git">git</option>
              </select>
            </label>

            <div class="draft-actions">
              <button class="phase-button" type="button" @click="handleDraftReview">{{ t.reviewDraft }}</button>
              <button
                class="toolbar-button secondary"
                type="button"
                :disabled="!draftReviewCard"
                @click="handleSaveDraft"
              >
                {{ t.saveDraft }}
              </button>
            </div>

            <div v-if="draftReviewCard" class="preview-card">
              <p class="label">{{ t.draftPreview }}</p>
              <strong>{{ selectedLanguage === 'en' ? draftReviewCard.titleEn : draftReviewCard.titleZh }}</strong>
              <p class="callout compact">{{ selectedLanguage === 'en' ? draftReviewCard.clickbaitEn : draftReviewCard.clickbaitZh }}</p>
              <p class="explanation compact">{{ selectedLanguage === 'en' ? draftReviewCard.explanationEn : draftReviewCard.explanationZh }}</p>
              <p class="label">{{ selectedLanguage === 'en' ? draftReviewCard.questionTextEn : draftReviewCard.questionTextZh }}</p>
            </div>
            <p v-else class="explanation">{{ t.noDraftYet }}</p>

            <div class="preview-diagnostics">
              <span class="label">{{ t.previewDiagnostics }}</span>
              <div v-if="draftReviewErrors.length" class="diagnostics-list compact">
                <article
                  v-for="item in draftReviewErrors"
                  :key="`draft-${item.source_path}-${item.code}-${item.field || ''}`"
                  class="diagnostic-item"
                  :class="(item.severity || 'error') === 'warning' ? 'warning' : 'error'"
                >
                  <div class="diagnostic-head">
                    <span class="severity-pill" :class="(item.severity || 'error') === 'warning' ? 'warning' : 'error'">
                      {{ severityLabel(item) }}
                    </span>
                    <strong>{{ item.code }}</strong>
                  </div>
                  <p>{{ item.message }}</p>
                </article>
              </div>
              <p v-else class="explanation">{{ t.noPreviewDiagnostics }}</p>
            </div>
          </section>
        </div>

        <span v-if="actionMessage" class="toolbar-message">{{ actionMessage }}</span>
      </section>
    </div>

    <div v-if="diagnosticsOpen" class="settings-overlay" @click.self="toggleDiagnostics">
      <section class="settings-popout diagnostics-popout">
        <div class="study-header">
          <div>
            <h2>{{ t.diagnosticsLabel }}</h2>
          </div>
          <button class="close-button" type="button" @click="toggleDiagnostics">{{ t.close }}</button>
        </div>

        <details class="diagnostics-disclosure" open>
          <summary>{{ t.diagnosticsTitle }}</summary>
          <p class="label batch-report-title">{{ t.deckReport }}</p>
          <section class="batch-report">
            <article class="status-card batch-stat">
              <span class="label">{{ t.totalCards }}</span>
              <strong>{{ totalCardsCount }}</strong>
            </article>
            <article class="status-card batch-stat">
              <span class="label">{{ t.cleanCards }}</span>
              <strong>{{ cleanCardCount }}</strong>
            </article>
            <article class="status-card batch-stat warning">
              <span class="label">{{ t.warningCards }}</span>
              <strong>{{ warningCardCount }}</strong>
            </article>
            <article class="status-card batch-stat error">
              <span class="label">{{ t.errorCards }}</span>
              <strong>{{ errorCardCount }}</strong>
            </article>
          </section>
          <section class="recent-changes">
            <div class="diagnostic-group-head">
              <span class="severity-pill neutral">{{ t.recentChanges }}</span>
              <strong>{{ recentlyChangedFiles.length }}</strong>
            </div>
            <div v-if="recentlyChangedFiles.length" class="recent-change-list">
              <article
                v-for="file in recentlyChangedFiles"
                :key="`recent-${file.path}`"
                class="recent-change-item"
              >
                <strong>{{ file.name }}</strong>
                <span>{{ t.lastUpdated }}: {{ formatDisplayTime(file.modifiedAt, t.notScheduled) }}</span>
              </article>
            </div>
            <p v-else class="explanation">{{ t.noRecentChanges }}</p>
          </section>
          <div class="diagnostic-filters">
            <label class="settings-field">
              <span>{{ t.severityFilter }}</span>
              <select v-model="diagnosticsFilter.severity" class="diagnostic-filter">
                <option value="all">{{ t.allFilter }}</option>
                <option value="warning">{{ t.warningLabel }}</option>
                <option value="error">{{ t.errorLabel }}</option>
              </select>
            </label>
            <label class="settings-field">
              <span>{{ t.topicFilter }}</span>
              <select v-model="diagnosticsFilter.topic" class="diagnostic-filter">
                <option v-for="topic in availableTopics" :key="topic" :value="topic">
                  {{ topic === 'all' ? t.allFilter : topic }}
                </option>
              </select>
            </label>
          </div>
          <p v-if="!importErrors.length" class="explanation">{{ t.noDiagnostics }}</p>
          <div class="diagnostics-groups">
            <section v-if="filteredWarningItems.length" class="diagnostic-group">
              <div class="diagnostic-group-head">
                <span class="severity-pill warning">{{ t.warningLabel }}</span>
                <strong>{{ filteredWarningItems.length }}</strong>
              </div>
              <div class="diagnostics-list">
                <article
                  v-for="item in filteredWarningItems"
                  :key="`${item.source_path}-${item.code}-${item.field || ''}`"
                  class="diagnostic-item warning"
                >
                  <div class="diagnostic-head">
                    <span class="severity-pill warning">{{ severityLabel(item) }}</span>
                    <strong>{{ item.code }}</strong>
                  </div>
                  <p>{{ item.message }}</p>
                  <span>{{ item.source_path }}</span>
                </article>
              </div>
            </section>

            <section v-if="filteredErrorItems.length" class="diagnostic-group">
              <div class="diagnostic-group-head">
                <span class="severity-pill error">{{ t.errorLabel }}</span>
                <strong>{{ filteredErrorItems.length }}</strong>
              </div>
              <div class="diagnostics-list">
                <article
                  v-for="item in filteredErrorItems"
                  :key="`${item.source_path}-${item.code}-${item.field || ''}`"
                  class="diagnostic-item error"
                >
                  <div class="diagnostic-head">
                    <span class="severity-pill error">{{ severityLabel(item) }}</span>
                    <strong>{{ item.code }}</strong>
                  </div>
                  <p>{{ item.message }}</p>
                  <span>{{ item.source_path }}</span>
                </article>
              </div>
            </section>
          </div>
        </details>
      </section>
    </div>

    <div v-if="resetWarningOpen" class="settings-overlay danger-overlay" @click.self="closeResetWarning">
      <section class="confirm-popout">
        <div class="study-header">
          <div>
            <p class="label">{{ t.resetStudyData }}</p>
            <h2>{{ t.resetWarningTitle }}</h2>
          </div>
        </div>
        <p class="explanation">{{ t.resetWarningBody }}</p>
        <div class="confirm-actions">
          <button class="close-button" type="button" @click="closeResetWarning">{{ t.cancel }}</button>
          <button class="toolbar-button danger confirm-reset-button" type="button" @click="handleResetStudyData">
            {{ t.confirmReset }}
          </button>
        </div>
      </section>
    </div>
  </main>
</template>
