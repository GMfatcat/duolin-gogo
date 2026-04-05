import {
  GetStudyCard,
  LoadAuthoringPreview,
  LoadDashboard,
  PreviewKnowledgeCard,
  ResetStudyData,
  RescanKnowledge,
  ReviewDraft,
  SaveDraft,
  SendTestNotification,
  SnoozeNotifications,
  SubmitAnswer,
  UpdateNotificationSettings,
  UpdatePreferredLanguage,
  UpdateScheduleSettings,
  UpdateSelectedTopic,
  ValidateKnowledge,
} from '../wailsjs/go/main/App'

const baseShownAt = '2026-04-05T10:00:00+08:00'

const fallbackCardsByTopic = {
  git: {
    id: 'git-cherry-pick-purpose',
    title: 'Cherry-pick Purpose',
    titleZh: 'Cherry-pick 指令用途',
    titleEn: 'Cherry-pick Purpose',
    questionType: 'true-false',
    questionText: '`git cherry-pick` applies a chosen commit to the current branch.',
    questionTextZh: '`git cherry-pick` 會把選定的 commit 套用到目前分支。',
    questionTextEn: '`git cherry-pick` applies a chosen commit to the current branch.',
    choices: [
      { value: 'true', labelZh: '是', labelEn: 'True' },
      { value: 'false', labelZh: '否', labelEn: 'False' },
    ],
    clickbait: 'One Git command can steal just one commit. Know which?',
    clickbaitZh: '有個 Git 指令可以只挑走單一 commit，你知道是哪個嗎？',
    clickbaitEn: 'One Git command can steal just one commit. Know which?',
    reviewHint: 'Cherry-pick copies selected commit changes onto your current branch.',
    reviewHintZh: 'Cherry-pick 會把選定 commit 的變更套用到目前分支。',
    reviewHintEn: 'Cherry-pick copies selected commit changes onto your current branch.',
    explanationZh: '`git cherry-pick` 會把選定 commit 的變更帶到目前分支。',
    explanationEn: '`git cherry-pick` lets you apply a chosen commit onto the current branch.',
    shownAt: baseShownAt,
  },
  docker: {
    id: 'docker-run-start-container',
    title: 'Docker Run',
    titleZh: 'docker run 啟動容器',
    titleEn: 'Docker Run',
    questionType: 'single-choice',
    questionText: 'What does `docker run` mainly do?',
    questionTextZh: '`docker run` 最主要在做什麼？',
    questionTextEn: 'What does `docker run` mainly do?',
    choices: [
      { value: '0', labelZh: '列出目前所有正在執行的容器', labelEn: 'List all currently running containers' },
      { value: '1', labelZh: '從 image 建立並啟動一個容器', labelEn: 'Create and start a container from an image' },
      { value: '2', labelZh: '進入一個已存在的容器 shell', labelEn: 'Open a shell inside an existing container' },
    ],
    clickbait: 'Think `build` is the first Docker command that matters? Many people get stuck on `run` first.',
    clickbaitZh: '你以為 Docker 第一個重要指令是 build？很多人其實先卡在 run。',
    clickbaitEn: 'Think `build` is the first Docker command that matters? Many people get stuck on `run` first.',
    reviewHint: '`run` creates and starts a container from an image.',
    reviewHintZh: '`run` = 用 image 建立並啟動容器。',
    reviewHintEn: '`run` creates and starts a container from an image.',
    explanationZh: '`docker run` 會根據指定的 image 建立一個新的 container，然後立刻啟動它。',
    explanationEn: '`docker run` creates a new container from the given image and starts it immediately.',
    shownAt: baseShownAt,
  },
  linux: {
    id: 'linux-pwd-current-directory',
    title: 'PWD',
    titleZh: 'pwd 目前位置',
    titleEn: 'PWD',
    questionType: 'true-false',
    questionText: '`pwd` shows the path of your current working directory.',
    questionTextZh: '`pwd` 會顯示你目前所在的工作目錄路徑。',
    questionTextEn: '`pwd` shows the path of your current working directory.',
    choices: [
      { value: 'true', labelZh: '是', labelEn: 'True' },
      { value: 'false', labelZh: '否', labelEn: 'False' },
    ],
    clickbait: 'Think the command is broken? Sometimes you are simply not where you think you are.',
    clickbaitZh: '你以為指令壞了？很多時候只是你根本不在自己以為的位置。',
    clickbaitEn: 'Think the command is broken? Sometimes you are simply not where you think you are.',
    reviewHint: '`pwd` means print working directory.',
    reviewHintZh: '`pwd` = print working directory。',
    reviewHintEn: '`pwd` means print working directory.',
    explanationZh: '`pwd` 代表 print working directory，會印出你目前所在資料夾的完整路徑。',
    explanationEn: '`pwd` stands for print working directory and prints the full path of your current directory.',
    shownAt: baseShownAt,
  },
  go: {
    id: 'go-goroutine-concurrency',
    title: 'Goroutine',
    titleZh: 'goroutine 並行執行',
    titleEn: 'Goroutine',
    questionType: 'single-choice',
    questionText: 'What is the main purpose of a goroutine in Go?',
    questionTextZh: '在 Go 裡，goroutine 最主要是用來做什麼？',
    questionTextEn: 'What is the main purpose of a goroutine in Go?',
    choices: [
      { value: '0', labelZh: '定義新的 struct 型別', labelEn: 'Define a new struct type' },
      { value: '1', labelZh: '讓函式以並行方式執行', labelEn: 'Run a function concurrently' },
      { value: '2', labelZh: '把套件編譯成執行檔', labelEn: 'Compile a package into an executable' },
    ],
    clickbait: 'Think Go feels fast only because the language is fast? Often it is about whether you use goroutines well.',
    clickbaitZh: '你以為 Go 變快只是因為它夠快？很多時候是因為你敢不敢開 goroutine。',
    clickbaitEn: 'Think Go feels fast only because the language is fast? Often it is about whether you use goroutines well.',
    reviewHint: 'A goroutine is a lightweight concurrent execution unit.',
    reviewHintZh: 'goroutine = 輕量並行執行單位。',
    reviewHintEn: 'A goroutine is a lightweight concurrent execution unit.',
    explanationZh: 'goroutine 是 Go 裡非常輕量的並行執行單位。',
    explanationEn: 'A goroutine is a very lightweight unit of concurrent execution in Go.',
    shownAt: baseShownAt,
  },
  python: {
    id: 'python-venv-isolation',
    title: 'Python Venv',
    titleZh: 'venv 虛擬環境',
    titleEn: 'Python Venv',
    questionType: 'true-false',
    questionText: '`venv` is commonly used to create isolated Python package environments.',
    questionTextZh: '`venv` 常用來建立彼此隔離的 Python 套件環境。',
    questionTextEn: '`venv` is commonly used to create isolated Python package environments.',
    choices: [
      { value: 'true', labelZh: '是', labelEn: 'True' },
      { value: 'false', labelZh: '否', labelEn: 'False' },
    ],
    clickbait: 'Think the package setup is broken? Sometimes the real issue is that multiple projects share the same environment.',
    clickbaitZh: '你以為套件壞掉，其實可能只是你把不同專案的環境全混在一起了。',
    clickbaitEn: 'Think the package setup is broken? Sometimes the real issue is that multiple projects share the same environment.',
    reviewHint: '`venv` isolates Python environments between projects.',
    reviewHintZh: '`venv` = 隔離不同專案的 Python 環境。',
    reviewHintEn: '`venv` isolates Python environments between projects.',
    explanationZh: '`venv` 常用來替不同 Python 專案建立彼此隔離的套件環境。',
    explanationEn: '`venv` is commonly used to create isolated package environments for different Python projects.',
    shownAt: baseShownAt,
  },
}

const topicWeakTopics = {
  all: [
    { tag: 'branching', wrongCount: 3, seenCount: 6, accuracy: 0.5 },
    { tag: 'commits', wrongCount: 1, seenCount: 4, accuracy: 0.75 },
  ],
  git: [
    { tag: 'branching', wrongCount: 3, seenCount: 6, accuracy: 0.5 },
    { tag: 'commits', wrongCount: 1, seenCount: 4, accuracy: 0.75 },
  ],
  docker: [{ tag: 'docker', wrongCount: 1, seenCount: 3, accuracy: 0.67 }],
  linux: [{ tag: 'linux', wrongCount: 1, seenCount: 4, accuracy: 0.75 }],
  go: [{ tag: 'go', wrongCount: 2, seenCount: 5, accuracy: 0.6 }],
  python: [{ tag: 'python', wrongCount: 1, seenCount: 5, accuracy: 0.8 }],
}

function cloneCard(topic) {
  return structuredClone(fallbackCardsByTopic[topic] || fallbackCardsByTopic.git)
}

function applyFallbackTopic(topic) {
  const resolvedTopic = topic || 'all'
  fallbackDashboard.selectedTopic = resolvedTopic
  fallbackDashboard.currentCard = cloneCard(resolvedTopic === 'all' ? 'git' : resolvedTopic)
  fallbackDashboard.summary = {
    ...fallbackDashboard.summary,
    weakTopics: structuredClone(topicWeakTopics[resolvedTopic] || topicWeakTopics.all),
  }
}

const fallbackDashboard = {
  info: {
    name: 'duolin-gogo',
    focusTopic: 'git',
    defaultLanguage: 'zh-TW',
  },
  preferredLanguage: 'zh-TW',
  selectedTopic: 'all',
  availableTopics: ['all', 'docker', 'git', 'go', 'linux', 'python'],
  stats: {
    studiedToday: 1,
    correctRate: 1,
  },
  summary: {
    nextReviewAt: '2026-04-05T21:00:00+08:00',
    weakTopics: structuredClone(topicWeakTopics.all),
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
  currentCard: cloneCard('git'),
  reviewQueue: [],
  reviewMode: false,
}

const fallbackAuthoringPreview = {
  files: [
    {
      path: 'D:/duolin-gogo/knowledge/git/cherry-pick.md',
      name: 'cherry-pick.md',
      modifiedAt: '2026-04-04T09:30:00+08:00',
    },
    {
      path: 'D:/duolin-gogo/knowledge/git/rebase.md',
      name: 'rebase.md',
      modifiedAt: '2026-04-05T11:45:00+08:00',
    },
  ],
  selectedPath: 'D:/duolin-gogo/knowledge/git/cherry-pick.md',
  currentCard: cloneCard('git'),
  importErrors: [],
}

const fallbackSavedDrafts = new Map()

const hasBackend = () => typeof window !== 'undefined' && typeof window.go !== 'undefined'

export function __resetFallbackState() {
  fallbackDashboard.preferredLanguage = 'zh-TW'
  fallbackDashboard.stats = {
    studiedToday: 1,
    correctRate: 1,
  }
  fallbackDashboard.summary = {
    nextReviewAt: '2026-04-05T21:00:00+08:00',
    weakTopics: structuredClone(topicWeakTopics.all),
  }
  fallbackDashboard.notificationSettings = {
    style: 'playful',
    titleMode: 'prefer_manual',
  }
  fallbackDashboard.scheduleSettings = {
    notificationIntervalMinutes: 10,
    reviewTime: '21:00',
    activeHoursEnabled: true,
    activeHoursStart: '09:00',
    activeHoursEnd: '22:00',
  }
  fallbackDashboard.importErrors = []
  fallbackDashboard.reviewQueue = []
  fallbackDashboard.reviewMode = false
  applyFallbackTopic('all')
  fallbackAuthoringPreview.selectedPath = 'D:/duolin-gogo/knowledge/git/cherry-pick.md'
  fallbackAuthoringPreview.currentCard = cloneCard('git')
  fallbackAuthoringPreview.importErrors = []
  fallbackAuthoringPreview.files = [
    {
      path: 'D:/duolin-gogo/knowledge/git/cherry-pick.md',
      name: 'cherry-pick.md',
      modifiedAt: '2026-04-04T09:30:00+08:00',
    },
    {
      path: 'D:/duolin-gogo/knowledge/git/rebase.md',
      name: 'rebase.md',
      modifiedAt: '2026-04-05T11:45:00+08:00',
    },
  ]
  fallbackSavedDrafts.clear()
}

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
  const isCorrect = selectedAnswer === 'true' || selectedAnswer === '1' || selectedAnswer === '0'

  return {
    cardId,
    isCorrect,
    correctAnswer: fallbackDashboard.currentCard.questionType === 'true-false' ? 'true' : '1',
    feedback:
      preferredLanguage === 'zh-TW'
        ? isCorrect
          ? '蝑?鈭?'
          : '???榆銝暺?'
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

  const match = Object.values(fallbackCardsByTopic).find((card) => card.id === cardId)
  return structuredClone(match || { ...fallbackDashboard.currentCard, id: cardId })
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

  return { message: 'Knowledge refreshed: 40 cards, 0 errors.' }
}

export async function validateKnowledge() {
  if (hasBackend()) {
    return ValidateKnowledge()
  }

  return { message: 'Knowledge validated: 40 cards, 0 diagnostics.', importErrors: [] }
}

export async function resetStudyData() {
  if (hasBackend()) {
    return ResetStudyData()
  }

  fallbackDashboard.stats = {
    studiedToday: 0,
    correctRate: 0,
  }
  fallbackDashboard.summary = {
    ...fallbackDashboard.summary,
    nextReviewAt: '',
    weakTopics: [],
  }
  fallbackDashboard.reviewQueue = []
  return { message: 'Study data reset.' }
}

export async function loadAuthoringPreview() {
  if (hasBackend()) {
    return LoadAuthoringPreview()
  }

  const savedFiles = Array.from(fallbackSavedDrafts.entries()).map(([path]) => ({
    path,
    name: path.split('/').at(-1),
    modifiedAt: '2026-04-05T12:00:00+08:00',
  }))

  return structuredClone({
    ...fallbackAuthoringPreview,
    files: [...fallbackAuthoringPreview.files, ...savedFiles],
  })
}

export async function previewKnowledgeCard(path) {
  if (hasBackend()) {
    return PreviewKnowledgeCard(path)
  }

  if (fallbackSavedDrafts.has(path)) {
    const savedCard = fallbackSavedDrafts.get(path)
    return {
      ...structuredClone(fallbackAuthoringPreview),
      files: [
        ...fallbackAuthoringPreview.files,
        ...Array.from(fallbackSavedDrafts.keys()).map((savedPath) => ({
          path: savedPath,
          name: savedPath.split('/').at(-1),
          modifiedAt: '2026-04-05T12:00:00+08:00',
        })),
      ],
      selectedPath: path,
      currentCard: structuredClone(savedCard),
    }
  }

  if (path.endsWith('rebase.md')) {
    return {
      ...structuredClone(fallbackAuthoringPreview),
      selectedPath: path,
      currentCard: {
        ...cloneCard('git'),
        id: 'git-rebase-vs-merge',
        title: 'Rebase vs Merge',
        titleZh: 'Rebase 與 Merge 的差異',
        titleEn: 'Rebase vs Merge',
        questionText: 'What does git rebase mainly do?',
        questionTextZh: 'git rebase 最主要在做什麼？',
        questionTextEn: 'What does git rebase mainly do?',
        clickbait: 'Do you really know the difference between rebase and merge?',
        clickbaitZh: '你真的分得清 rebase 和 merge 的差別嗎？',
        clickbaitEn: 'Do you really know the difference between rebase and merge?',
        reviewHint: 'Rebase replays commits on top of another base.',
        reviewHintZh: 'Rebase = 把 commits 重放到另一條 base 上。',
        reviewHintEn: 'Rebase replays commits on top of another base.',
        explanationZh: '`git rebase` 會把你的 commits 重新套到另一個 base 上。',
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
      ...cloneCard('git'),
      id: pickField('id') || fallbackCardsByTopic.git.id,
      title: pickField('title_en') || pickField('title') || fallbackCardsByTopic.git.title,
      titleZh: pickField('title_zh') || pickField('title') || fallbackCardsByTopic.git.titleZh,
      titleEn: pickField('title_en') || pickField('title') || fallbackCardsByTopic.git.titleEn,
      questionText: pickField('question_en') || pickField('question') || fallbackCardsByTopic.git.questionText,
      questionTextZh: pickField('question_zh') || pickField('question') || fallbackCardsByTopic.git.questionTextZh,
      questionTextEn: pickField('question_en') || pickField('question') || fallbackCardsByTopic.git.questionTextEn,
      clickbait: pickField('clickbait_en') || pickField('clickbait') || fallbackCardsByTopic.git.clickbait,
      clickbaitZh: pickField('clickbait_zh') || pickField('clickbait') || fallbackCardsByTopic.git.clickbaitZh,
      clickbaitEn: pickField('clickbait_en') || pickField('clickbait') || fallbackCardsByTopic.git.clickbaitEn,
      reviewHint: pickField('review_hint_en') || pickField('review_hint') || fallbackCardsByTopic.git.reviewHint,
      reviewHintZh: pickField('review_hint_zh') || pickField('review_hint') || fallbackCardsByTopic.git.reviewHintZh,
      reviewHintEn: pickField('review_hint_en') || pickField('review_hint') || fallbackCardsByTopic.git.reviewHintEn,
      explanationZh: raw.split('## zh-TW')[1]?.split('## en')[0]?.trim() || fallbackCardsByTopic.git.explanationZh,
      explanationEn: raw.split('## en')[1]?.trim() || fallbackCardsByTopic.git.explanationEn,
    },
    importErrors: [],
  }
}

export async function saveDraft({ raw, topic }) {
  if (hasBackend()) {
    return SaveDraft(raw, topic)
  }

  const pickField = (name) => {
    const match = raw.match(new RegExp(`^${name}:\\s*"?(.+?)"?$`, 'm'))
    return match ? match[1] : ''
  }

  const resolvedTopic = topic || 'git'
  const savedPath = `D:/duolin-gogo/knowledge/${resolvedTopic}/${pickField('id') || 'git-ai-review'}.md`
  fallbackSavedDrafts.set(savedPath, {
    ...cloneCard('git'),
    id: pickField('id') || fallbackCardsByTopic.git.id,
    title: pickField('title_en') || pickField('title') || fallbackCardsByTopic.git.title,
    titleZh: pickField('title_zh') || pickField('title') || fallbackCardsByTopic.git.titleZh,
    titleEn: pickField('title_en') || pickField('title') || fallbackCardsByTopic.git.titleEn,
    questionText: pickField('question_en') || pickField('question') || fallbackCardsByTopic.git.questionText,
    questionTextZh: pickField('question_zh') || pickField('question') || fallbackCardsByTopic.git.questionTextZh,
    questionTextEn: pickField('question_en') || pickField('question') || fallbackCardsByTopic.git.questionTextEn,
    clickbait: pickField('clickbait_en') || pickField('clickbait') || fallbackCardsByTopic.git.clickbait,
    clickbaitZh: pickField('clickbait_zh') || pickField('clickbait') || fallbackCardsByTopic.git.clickbaitZh,
    clickbaitEn: pickField('clickbait_en') || pickField('clickbait') || fallbackCardsByTopic.git.clickbaitEn,
    reviewHint: pickField('review_hint_en') || pickField('review_hint') || fallbackCardsByTopic.git.reviewHint,
    reviewHintZh: pickField('review_hint_zh') || pickField('review_hint') || fallbackCardsByTopic.git.reviewHintZh,
    reviewHintEn: pickField('review_hint_en') || pickField('review_hint') || fallbackCardsByTopic.git.reviewHintEn,
    explanationZh: raw.split('## zh-TW')[1]?.split('## en')[0]?.trim() || fallbackCardsByTopic.git.explanationZh,
    explanationEn: raw.split('## en')[1]?.trim() || fallbackCardsByTopic.git.explanationEn,
  })

  return {
    message: `Draft saved to ${savedPath}.`,
    savedPath,
    topic: resolvedTopic,
    successful: true,
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

export async function updateSelectedTopic(topic) {
  if (hasBackend()) {
    return UpdateSelectedTopic(topic)
  }

  applyFallbackTopic(topic)
  return { message: 'Topic filter updated.' }
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
