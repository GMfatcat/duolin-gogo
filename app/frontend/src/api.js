import {
  GetStudyCard,
  LoadDashboard,
  LoadAuthoringPreview,
  PreviewKnowledgeCard,
  ReviewDraft,
  RescanKnowledge,
  SendTestNotification,
  SnoozeNotifications,
  SubmitAnswer,
  UpdateNotificationSettings,
  UpdatePreferredLanguage,
  UpdateScheduleSettings,
  ValidateKnowledge,
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
    activeHoursEnabled: true,
    activeHoursStart: '09:00',
    activeHoursEnd: '22:00',
  },
  importErrors: [],
  currentCard: {
    id: 'git-cherry-pick-purpose',
    title: 'Cherry-pick Purpose',
    titleZh: 'Cherry-pick 的用途',
    titleEn: 'Cherry-pick Purpose',
    questionType: 'true-false',
    questionText: '`git cherry-pick` applies a chosen commit to the current branch.',
    questionTextZh: '`git cherry-pick` 會把指定的 commit 套到目前分支。',
    questionTextEn: '`git cherry-pick` applies a chosen commit to the current branch.',
    choices: [
      { value: 'true', labelZh: '是', labelEn: 'True' },
      { value: 'false', labelZh: '否', labelEn: 'False' },
    ],
    clickbait: 'One Git command can steal just one commit. Know which?',
    clickbaitZh: '有一個 Git 指令，可以只搬走一個 commit。你知道是哪個嗎？',
    clickbaitEn: 'One Git command can steal just one commit. Know which?',
    reviewHint: 'Cherry-pick copies selected commit changes onto your current branch.',
    reviewHintZh: 'Cherry-pick 會把指定 commit 的變更套用到目前分支。',
    reviewHintEn: 'Cherry-pick copies selected commit changes onto your current branch.',
    explanationZh: '`git cherry-pick` 讓你把特定 commit 的內容套用到目前分支。',
    explanationEn: '`git cherry-pick` lets you apply a chosen commit onto the current branch.',
    shownAt: '2026-04-05T10:00:00+08:00',
  },
  reviewQueue: [],
  reviewMode: false,
}

const fallbackAuthoringPreview = {
  files: [
    { path: 'D:/duolin-gogo/knowledge/git/cherry-pick.md', name: 'cherry-pick.md' },
    { path: 'D:/duolin-gogo/knowledge/git/rebase.md', name: 'rebase.md' },
  ],
  selectedPath: 'D:/duolin-gogo/knowledge/git/cherry-pick.md',
  currentCard: structuredClone(fallbackDashboard.currentCard),
  importErrors: [],
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
          : '這題還差一點。'
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

  return { message: 'Notifications snoozed until 10:15.' }
}

export async function rescanKnowledge() {
  if (hasBackend()) {
    return RescanKnowledge()
  }

  return { message: 'Knowledge refreshed: 12 cards, 0 errors.' }
}

export async function validateKnowledge() {
  if (hasBackend()) {
    return ValidateKnowledge()
  }

  return { message: 'Knowledge validated: 12 cards, 0 diagnostics.', importErrors: [] }
}

export async function loadAuthoringPreview() {
  if (hasBackend()) {
    return LoadAuthoringPreview()
  }

  return structuredClone(fallbackAuthoringPreview)
}

export async function previewKnowledgeCard(path) {
  if (hasBackend()) {
    return PreviewKnowledgeCard(path)
  }

  if (path.endsWith('rebase.md')) {
    return {
      ...structuredClone(fallbackAuthoringPreview),
      selectedPath: path,
      currentCard: {
        ...structuredClone(fallbackDashboard.currentCard),
        id: 'git-rebase-vs-merge',
        title: 'Rebase vs Merge',
        titleZh: 'Rebase 與 Merge 的差別',
        titleEn: 'Rebase vs Merge',
        questionText: 'What does git rebase mainly do?',
        questionTextZh: 'git rebase 最主要在做什麼？',
        questionTextEn: 'What does git rebase mainly do?',
        clickbait: 'Do you really know the difference between rebase and merge?',
        clickbaitZh: '你真的懂 rebase 跟 merge 差在哪嗎？',
        clickbaitEn: 'Do you really know the difference between rebase and merge?',
        reviewHint: 'Rebase replays commits on top of another base.',
        reviewHintZh: 'Rebase 會把 commits 重放到另一條 base 上。',
        reviewHintEn: 'Rebase replays commits on top of another base.',
        explanationZh: '`git rebase` 會把目前分支的 commits 重新接到另一條 base 上。',
        explanationEn: '`git rebase` replays your commits on top of another base.',
      },
    }
  }

  return {
    ...structuredClone(fallbackAuthoringPreview),
    selectedPath: path,
  }
}

export async function reviewDraft(raw) {
  if (hasBackend()) {
    return ReviewDraft(raw)
  }

  if (!raw.includes('## zh-TW') || !raw.includes('## en')) {
    return {
      currentCard: null,
      importErrors: [
        {
          source_path: 'draft://ai-card.md',
          severity: 'error',
          code: 'missing_language_section',
          field: 'body',
          message: 'Body must contain both ## zh-TW and ## en sections.',
        },
      ],
    }
  }

  const pickField = (name) => {
    const match = raw.match(new RegExp(`^${name}:\\s*"?(.+?)"?$`, 'm'))
    return match ? match[1] : ''
  }

  return {
    currentCard: {
      ...structuredClone(fallbackDashboard.currentCard),
      id: pickField('id') || fallbackDashboard.currentCard.id,
      title: pickField('title_en') || pickField('title') || fallbackDashboard.currentCard.title,
      titleZh: pickField('title_zh') || pickField('title') || fallbackDashboard.currentCard.titleZh,
      titleEn: pickField('title_en') || pickField('title') || fallbackDashboard.currentCard.titleEn,
      questionText: pickField('question_en') || pickField('question') || fallbackDashboard.currentCard.questionText,
      questionTextZh: pickField('question_zh') || pickField('question') || fallbackDashboard.currentCard.questionTextZh,
      questionTextEn: pickField('question_en') || pickField('question') || fallbackDashboard.currentCard.questionTextEn,
      clickbait: pickField('clickbait_en') || pickField('clickbait') || fallbackDashboard.currentCard.clickbait,
      clickbaitZh: pickField('clickbait_zh') || pickField('clickbait') || fallbackDashboard.currentCard.clickbaitZh,
      clickbaitEn: pickField('clickbait_en') || pickField('clickbait') || fallbackDashboard.currentCard.clickbaitEn,
      reviewHint: pickField('review_hint_en') || pickField('review_hint') || fallbackDashboard.currentCard.reviewHint,
      reviewHintZh: pickField('review_hint_zh') || pickField('review_hint') || fallbackDashboard.currentCard.reviewHintZh,
      reviewHintEn: pickField('review_hint_en') || pickField('review_hint') || fallbackDashboard.currentCard.reviewHintEn,
      explanationZh: raw.split('## zh-TW')[1]?.split('## en')[0]?.trim() || fallbackDashboard.currentCard.explanationZh,
      explanationEn: raw.split('## en')[1]?.trim() || fallbackDashboard.currentCard.explanationEn,
    },
    importErrors: [],
  }
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

export async function updateScheduleSettings({
  notificationIntervalMinutes,
  reviewTime,
  activeHoursEnabled,
  activeHoursStart,
  activeHoursEnd,
}) {
  if (hasBackend()) {
    return UpdateScheduleSettings(
      notificationIntervalMinutes,
      reviewTime,
      activeHoursEnabled,
      activeHoursStart,
      activeHoursEnd,
    )
  }

  fallbackDashboard.scheduleSettings = {
    notificationIntervalMinutes,
    reviewTime,
    activeHoursEnabled,
    activeHoursStart,
    activeHoursEnd,
  }
  return { message: 'Schedule settings updated.' }
}
