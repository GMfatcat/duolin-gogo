import {
  GetStudyCard,
  LoadDashboard,
  RescanKnowledge,
  SendTestNotification,
  SnoozeNotifications,
  SubmitAnswer,
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
  importErrors: [],
  currentCard: {
    id: 'git-cherry-pick-purpose',
    title: 'Cherry-pick Purpose',
    titleZh: 'Cherry-pick 的用途',
    titleEn: 'Cherry-pick Purpose',
    questionType: 'true-false',
    questionText: '`git cherry-pick` applies a chosen commit to the current branch.',
    questionTextZh: '`git cherry-pick` 會把指定的 commit 套用到目前分支上。',
    questionTextEn: '`git cherry-pick` applies a chosen commit to the current branch.',
    choices: [
      { value: 'true', labelZh: '是', labelEn: 'True' },
      { value: 'false', labelZh: '否', labelEn: 'False' },
    ],
    clickbait: 'One Git command can steal just one commit. Know which?',
    clickbaitZh: '哪個 Git 指令可以只拿走一個 commit？',
    clickbaitEn: 'One Git command can steal just one commit. Know which?',
    reviewHint: 'Cherry-pick copies selected commit changes onto your current branch.',
    reviewHintZh: 'Cherry-pick 會把選定 commit 的變更套到目前分支。',
    reviewHintEn: 'Cherry-pick copies selected commit changes onto your current branch.',
    explanationZh: '`git cherry-pick` 會把你指定的一個 commit 套用到目前分支上。',
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

  return {
    cardId,
    isCorrect: selectedAnswer === 'true',
    correctAnswer: 'true',
    feedback: selectedAnswer === 'true' ? 'Correct.' : 'Not quite.',
    reviewHint: fallbackDashboard.currentCard.reviewHint,
    preferredLanguage: fallbackDashboard.preferredLanguage,
    stats: {
      studiedToday: 2,
      correctRate: selectedAnswer === 'true' ? 1 : 0.5,
    },
  }
}

export async function getStudyCard(cardId) {
  if (hasBackend()) {
    return GetStudyCard(cardId)
  }

  return structuredClone(fallbackDashboard.currentCard)
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
