import {
  GetStudyCard,
  LoadDashboard,
  RescanKnowledge,
  SendTestNotification,
  SnoozeNotifications,
  SubmitAnswer,
  UpdateNotificationSettings,
  UpdatePreferredLanguage,
  UpdateScheduleSettings,
} from '../wailsjs/go/main/App'

const fallbackDashboard = {
  info: {
    name: 'duolin-gogo',
    focusTopic: 'git',
    defaultLanguage: 'zh-TW',
  },
  preferredLanguage: 'zh-TW',
  stats: {
    studiedToday: 1,
    correctRate: 1,
  },
  summary: {
    studiedToday: 1,
    correctRate: 1,
    nextReviewAt: '2026-04-05T21:00:00+08:00',
    weakTopics: [
      { tag: 'branching', wrongCount: 3, seenCount: 6, accuracy: 0.5 },
      { tag: 'commits', wrongCount: 1, seenCount: 4, accuracy: 0.75 },
    ],
  },
  notificationSettings: {
    style: 'playful',
    titleMode: 'prefer_manual',
  },
  scheduleSettings: {
    notificationIntervalMinutes: 10,
    reviewTime: '21:00',
  },
  importErrors: [],
  currentCard: {
    id: 'git-cherry-pick-purpose',
    title: 'Cherry-pick Purpose',
    titleZh: 'Cherry-pick 的用途',
    titleEn: 'Cherry-pick Purpose',
    questionType: 'true-false',
    questionText: '`git cherry-pick` applies a chosen commit to the current branch.',
    questionTextZh: '`git cherry-pick` 會把指定的 commit 套用到目前分支。',
    questionTextEn: '`git cherry-pick` applies a chosen commit to the current branch.',
    choices: [
      { value: 'true', labelZh: '是', labelEn: 'True' },
      { value: 'false', labelZh: '否', labelEn: 'False' },
    ],
    clickbait: 'One Git command can steal just one commit. Know which?',
    clickbaitZh: '有個 Git 指令能只偷一個 commit，你知道是哪個嗎？',
    clickbaitEn: 'One Git command can steal just one commit. Know which?',
    reviewHint: 'Cherry-pick copies selected commit changes onto your current branch.',
    reviewHintZh: 'Cherry-pick 會把選定 commit 的變更複製到目前分支。',
    reviewHintEn: 'Cherry-pick copies selected commit changes onto your current branch.',
    explanationZh: '`git cherry-pick` 可以把指定的 commit 套用到目前分支上。',
    explanationEn: '`git cherry-pick` lets you apply a chosen commit onto the current branch.',
    shownAt: '2026-04-05T10:00:00+08:00',
  },
  reviewQueue: [],
  reviewMode: false,
}

const hasBackend = () => typeof window !== 'undefined' && typeof window.go !== 'undefined'

export async function loadDashboard() {
  if (hasBackend()) {
    return LoadDashboard()
  }

  return structuredClone(fallbackDashboard)
}

export async function submitAnswer({ cardId, sessionType, selectedAnswer, shownAt }) {
  if (hasBackend()) {
    return SubmitAnswer(cardId, sessionType, selectedAnswer, shownAt)
  }

  const preferredLanguage = fallbackDashboard.preferredLanguage
  const isCorrect = selectedAnswer === 'true'

  return {
    cardId,
    isCorrect,
    correctAnswer: 'true',
    feedback:
      preferredLanguage === 'zh-TW'
        ? isCorrect
          ? '答對了。'
          : '再想一下。'
        : isCorrect
          ? 'Correct.'
          : 'Not quite.',
    reviewHint:
      preferredLanguage === 'zh-TW'
        ? fallbackDashboard.currentCard.reviewHintZh
        : fallbackDashboard.currentCard.reviewHintEn,
    preferredLanguage,
    stats: {
      studiedToday: 2,
      correctRate: isCorrect ? 1 : 0.5,
    },
  }
}

export async function getStudyCard(cardId) {
  if (hasBackend()) {
    return GetStudyCard(cardId)
  }

  return structuredClone({ ...fallbackDashboard.currentCard, id: cardId })
}

export async function sendTestNotification() {
  if (hasBackend()) {
    return SendTestNotification()
  }

  return { message: 'Test notification sent.' }
}

export async function snoozeNotifications() {
  if (hasBackend()) {
    return SnoozeNotifications()
  }

  return { message: 'Notifications snoozed for 15 minutes.' }
}

export async function rescanKnowledge() {
  if (hasBackend()) {
    return RescanKnowledge()
  }

  return { message: 'Knowledge refreshed: 2 cards, 0 errors.' }
}

export async function updateNotificationSettings({ style, titleMode }) {
  if (hasBackend()) {
    return UpdateNotificationSettings(style, titleMode)
  }

  fallbackDashboard.notificationSettings = { style, titleMode }
  return { message: 'Notification settings updated.' }
}

export async function updatePreferredLanguage(language) {
  if (hasBackend()) {
    return UpdatePreferredLanguage(language)
  }

  fallbackDashboard.preferredLanguage = language
  return { message: 'Language updated.' }
}

export async function updateScheduleSettings({ notificationIntervalMinutes, reviewTime }) {
  if (hasBackend()) {
    return UpdateScheduleSettings(notificationIntervalMinutes, reviewTime)
  }

  fallbackDashboard.scheduleSettings = {
    notificationIntervalMinutes,
    reviewTime,
  }
  return { message: 'Schedule settings updated.' }
}
