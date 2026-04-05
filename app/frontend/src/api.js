import {
  GetStudyCard,
  GetDGReaction,
  InteractWithDG,
  LoadAuthoringPreview,
  LoadDashboard,
  PreviewKnowledgeCard,
  ResetStudyData,
  RescanKnowledge,
  ReviewDraft,
  SaveDraft,
  SendTestNotification,
  SnoozeNotifications,
  StartLearnBreak,
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
    titleZh: 'Cherry-pick 用途',
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
    clickbaitZh: '有個 Git 指令只會拿走一個 commit，你知道是哪個嗎？',
    clickbaitEn: 'One Git command can steal just one commit. Know which?',
    reviewHint: 'Cherry-pick copies selected commit changes onto your current branch.',
    reviewHintZh: 'Cherry-pick 會把選定 commit 的變更複製到目前分支。',
    reviewHintEn: 'Cherry-pick copies selected commit changes onto your current branch.',
    explanationZh: '`git cherry-pick` 可以把選定的 commit 套用到目前分支。',
    explanationEn: '`git cherry-pick` lets you apply a chosen commit onto the current branch.',
    shownAt: baseShownAt,
  },
  docker: {
    id: 'docker-run-start-container',
    title: 'Docker Run',
    titleZh: 'Docker Run',
    titleEn: 'Docker Run',
    questionType: 'single-choice',
    questionText: 'What does `docker run` mainly do?',
    questionTextZh: '`docker run` 主要在做什麼？',
    questionTextEn: 'What does `docker run` mainly do?',
    choices: [
      { value: '0', labelZh: '列出目前執行中的容器', labelEn: 'List all currently running containers' },
      { value: '1', labelZh: '從 image 建立並啟動容器', labelEn: 'Create and start a container from an image' },
      { value: '2', labelZh: '進入既有容器的 shell', labelEn: 'Open a shell inside an existing container' },
    ],
    clickbait: 'Think `build` is the first Docker command that matters? Many people get stuck on `run` first.',
    clickbaitZh: '你以為先學 Docker build 就夠了？很多人其實先卡在 run。',
    clickbaitEn: 'Think `build` is the first Docker command that matters? Many people get stuck on `run` first.',
    reviewHint: '`run` creates and starts a container from an image.',
    reviewHintZh: '`run` = 從 image 建立並啟動容器。',
    reviewHintEn: '`run` creates and starts a container from an image.',
    explanationZh: '`docker run` 會從指定 image 建立新的 container，並立即啟動它。',
    explanationEn: '`docker run` creates a new container from the given image and starts it immediately.',
    shownAt: baseShownAt,
  },
  linux: {
    id: 'linux-pwd-current-directory',
    title: 'PWD',
    titleZh: 'PWD',
    titleEn: 'PWD',
    questionType: 'true-false',
    questionText: '`pwd` shows the path of your current working directory.',
    questionTextZh: '`pwd` 會顯示你目前工作目錄的路徑。',
    questionTextEn: '`pwd` shows the path of your current working directory.',
    choices: [
      { value: 'true', labelZh: '是', labelEn: 'True' },
      { value: 'false', labelZh: '否', labelEn: 'False' },
    ],
    clickbait: 'Think the command is broken? Sometimes you are simply not where you think you are.',
    clickbaitZh: '你以為指令壞了？有時只是你根本不在自己以為的位置。',
    clickbaitEn: 'Think the command is broken? Sometimes you are simply not where you think you are.',
    reviewHint: '`pwd` means print working directory.',
    reviewHintZh: '`pwd` = print working directory。',
    reviewHintEn: '`pwd` means print working directory.',
    explanationZh: '`pwd` 是 print working directory 的縮寫，會印出目前目錄的完整路徑。',
    explanationEn: '`pwd` stands for print working directory and prints the full path of your current directory.',
    shownAt: baseShownAt,
  },
  go: {
    id: 'go-goroutine-concurrency',
    title: 'Goroutine',
    titleZh: 'Goroutine',
    titleEn: 'Goroutine',
    questionType: 'single-choice',
    questionText: 'What is the main purpose of a goroutine in Go?',
    questionTextZh: '在 Go 裡，goroutine 的主要用途是什麼？',
    questionTextEn: 'What is the main purpose of a goroutine in Go?',
    choices: [
      { value: '0', labelZh: '定義新的 struct 型別', labelEn: 'Define a new struct type' },
      { value: '1', labelZh: '讓函式並行執行', labelEn: 'Run a function concurrently' },
      { value: '2', labelZh: '把 package 編譯成可執行檔', labelEn: 'Compile a package into an executable' },
    ],
    clickbait: 'Think Go feels fast only because the language is fast? Often it is about whether you use goroutines well.',
    clickbaitZh: '你以為 Go 快只是因為語言快？很多時候差在你會不會用 goroutine。',
    clickbaitEn: 'Think Go feels fast only because the language is fast? Often it is about whether you use goroutines well.',
    reviewHint: 'A goroutine is a lightweight concurrent execution unit.',
    reviewHintZh: 'goroutine = 輕量的並行執行單位。',
    reviewHintEn: 'A goroutine is a lightweight concurrent execution unit.',
    explanationZh: 'goroutine 是 Go 裡非常輕量的並行執行單位。',
    explanationEn: 'A goroutine is a very lightweight unit of concurrent execution in Go.',
    shownAt: baseShownAt,
  },
  python: {
    id: 'python-venv-isolation',
    title: 'Python Venv',
    titleZh: 'Python Venv',
    titleEn: 'Python Venv',
    questionType: 'true-false',
    questionText: '`venv` is commonly used to create isolated Python package environments.',
    questionTextZh: '`venv` 常用來建立隔離的 Python 套件環境。',
    questionTextEn: '`venv` is commonly used to create isolated Python package environments.',
    choices: [
      { value: 'true', labelZh: '是', labelEn: 'True' },
      { value: 'false', labelZh: '否', labelEn: 'False' },
    ],
    clickbait: 'Think the package setup is broken? Sometimes the real issue is that multiple projects share the same environment.',
    clickbaitZh: '你以為套件壞了？真正的問題常常是多個專案共用同一個環境。',
    clickbaitEn: 'Think the package setup is broken? Sometimes the real issue is that multiple projects share the same environment.',
    reviewHint: '`venv` isolates Python environments between projects.',
    reviewHintZh: '`venv` 會隔離不同專案的 Python 環境。',
    reviewHintEn: '`venv` isolates Python environments between projects.',
    explanationZh: '`venv` 常用來為不同 Python 專案建立彼此隔離的套件環境。',
    explanationEn: '`venv` is commonly used to create isolated package environments for different Python projects.',
    shownAt: baseShownAt,
  },
}
const topicWeakTopics = {
  all: [
    { tag: 'branching', wrongCount: 3, seenCount: 6, accuracy: 0.5 },
    { tag: 'commits', wrongCount: 1, seenCount: 4, accuracy: 0.75 },
  ],
  'backend-tools': [
    { tag: 'docker', wrongCount: 1, seenCount: 3, accuracy: 0.67 },
    { tag: 'linux', wrongCount: 1, seenCount: 4, accuracy: 0.75 },
    { tag: 'commits', wrongCount: 1, seenCount: 4, accuracy: 0.75 },
  ],
  languages: [
    { tag: 'go', wrongCount: 2, seenCount: 5, accuracy: 0.6 },
    { tag: 'python', wrongCount: 1, seenCount: 5, accuracy: 0.8 },
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

const topicProgressByMode = {
  all: [
    { topic: 'docker', seenCount: 3, correctCount: 2, wrongCount: 1, accuracy: 0.67 },
    { topic: 'go', seenCount: 5, correctCount: 3, wrongCount: 2, accuracy: 0.6 },
    { topic: 'git', seenCount: 10, correctCount: 8, wrongCount: 2, accuracy: 0.8 },
    { topic: 'linux', seenCount: 4, correctCount: 3, wrongCount: 1, accuracy: 0.75 },
    { topic: 'python', seenCount: 5, correctCount: 4, wrongCount: 1, accuracy: 0.8 },
  ],
  'backend-tools': [
    { topic: 'docker', seenCount: 3, correctCount: 2, wrongCount: 1, accuracy: 0.67 },
    { topic: 'git', seenCount: 10, correctCount: 8, wrongCount: 2, accuracy: 0.8 },
    { topic: 'linux', seenCount: 4, correctCount: 3, wrongCount: 1, accuracy: 0.75 },
  ],
  languages: [
    { topic: 'go', seenCount: 5, correctCount: 3, wrongCount: 2, accuracy: 0.6 },
    { topic: 'python', seenCount: 5, correctCount: 4, wrongCount: 1, accuracy: 0.8 },
  ],
  git: [{ topic: 'git', seenCount: 10, correctCount: 8, wrongCount: 2, accuracy: 0.8 }],
  docker: [{ topic: 'docker', seenCount: 3, correctCount: 2, wrongCount: 1, accuracy: 0.67 }],
  linux: [{ topic: 'linux', seenCount: 4, correctCount: 3, wrongCount: 1, accuracy: 0.75 }],
  go: [{ topic: 'go', seenCount: 5, correctCount: 3, wrongCount: 2, accuracy: 0.6 }],
  python: [{ topic: 'python', seenCount: 5, correctCount: 4, wrongCount: 1, accuracy: 0.8 }],
}

const weakestDeckByMode = {
  all: { topic: 'docker', seenCount: 3, correctCount: 2, wrongCount: 1, accuracy: 0.67 },
  'backend-tools': { topic: 'docker', seenCount: 3, correctCount: 2, wrongCount: 1, accuracy: 0.67 },
  languages: { topic: 'go', seenCount: 5, correctCount: 3, wrongCount: 2, accuracy: 0.6 },
  git: null,
  docker: null,
  linux: null,
  go: null,
  python: null,
}

function cloneCard(topic) {
  return structuredClone(fallbackCardsByTopic[topic] || fallbackCardsByTopic.git)
}

function primaryTopicForMode(topic) {
  switch (topic) {
    case 'backend-tools':
      return 'docker'
    case 'languages':
      return 'go'
    case 'all':
      return 'git'
    default:
      return topic
  }
}

function applyFallbackTopic(topic) {
  const resolvedTopic = topic || 'all'
  fallbackDashboard.selectedTopic = resolvedTopic
  fallbackDashboard.currentCard = cloneCard(primaryTopicForMode(resolvedTopic))
  fallbackDashboard.summary = {
    ...fallbackDashboard.summary,
    weakTopics: structuredClone(topicWeakTopics[resolvedTopic] || topicWeakTopics.all),
    topicProgress: structuredClone(topicProgressByMode[resolvedTopic] || topicProgressByMode.all),
    weakestDeck: structuredClone(weakestDeckByMode[resolvedTopic] ?? weakestDeckByMode.all),
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
  availableTopics: ['all', 'backend-tools', 'languages', 'docker', 'git', 'go', 'linux', 'python'],
  stats: {
    studiedToday: 1,
    correctRate: 1,
  },
  summary: {
    nextReviewAt: '2026-04-05T21:00:00+08:00',
    studyStreak: 3,
    weakTopics: structuredClone(topicWeakTopics.all),
    topicProgress: structuredClone(topicProgressByMode.all),
    weakestDeck: structuredClone(weakestDeckByMode.all),
  },
  notificationSettings: {
    style: 'playful',
    titleMode: 'prefer_manual',
  },
  scheduleSettings: {
    notificationIntervalMinutes: 20,
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
const fallbackPetState = {
  bondXp: 0,
  stage: 0,
  lastInteractionAt: '',
  lastReactionAt: '',
}

function fallbackPetIndex(trigger, size) {
  if (!size) {
    return 0
  }
  const minuteSeed = Math.floor(new Date(baseShownAt).getTime() / 60000)
  return (minuteSeed + fallbackPetState.bondXp + trigger.length) % size
}

function pickFallbackPetReaction(trigger, pool) {
  return pool[fallbackPetIndex(trigger, pool.length)]
}

function fallbackClickPool(language, topic, stage) {
  const normalizedTopic = topic || 'all'

  if (language === 'zh-TW') {
    if (normalizedTopic === 'docker') {
      return [
        { title: 'DG', body: 'Docker 模式開著，我們把這堆容器整理乾淨。', variant: 'focus', pose: 'wave', stage },
        { title: 'DG', body: '回到 docker 區了，我陪你盯好這些移動零件。', variant: 'focus', pose: 'wave', stage },
      ]
    }
    if (normalizedTopic === 'languages') {
      return [
        { title: 'DG', body: '今天這輪偏語言感，我們把細節咬準。', variant: 'focus', pose: 'wave', stage },
        { title: 'DG', body: '程式語言模式開始了，節奏會比較講究。', variant: 'focus', pose: 'wave', stage },
      ]
    }
  } else {
    if (normalizedTopic === 'docker') {
      return [
        { title: 'DG', body: 'Docker mode is on. Let us keep this stack tidy.', variant: 'focus', pose: 'wave', stage },
        { title: 'DG', body: 'Back in docker land. I am watching the moving parts with you.', variant: 'focus', pose: 'wave', stage },
      ]
    }
    if (normalizedTopic === 'languages') {
      return [
        { title: 'DG', body: 'This round leans language-heavy. Let us tune the details.', variant: 'focus', pose: 'wave', stage },
        { title: 'DG', body: 'Language mode is on. The rhythm here is more about nuance.', variant: 'focus', pose: 'wave', stage },
      ]
    }
  }

  if (language === 'zh-TW') {
    if (stage >= 2) {
      return [
        { title: 'DG', body: '你回來了，我開始抓到你的節奏了。', variant: 'celebration', pose: 'spark', stage },
        { title: 'DG', body: '這輪我跟得上，你只管繼續。', variant: 'celebration', pose: 'spark', stage },
      ]
    }
    if (stage >= 1) {
      return [
        { title: 'DG', body: '好，這一輪我們一起走完。', variant: 'focus', pose: 'wave', stage },
        { title: 'DG', body: '我準備好了，你先開題。', variant: 'focus', pose: 'wave', stage },
      ]
    }
    return [
      { title: 'DG', body: '我在這裡，慢慢來就好。', variant: 'neutral', pose: 'idle', stage },
      { title: 'DG', body: '多點我幾次，我會更快進入狀態。', variant: 'neutral', pose: 'idle', stage },
    ]
  }

  if (stage >= 2) {
    return [
      { title: 'DG', body: 'You are back. I am starting to learn your rhythm.', variant: 'celebration', pose: 'spark', stage },
      { title: 'DG', body: 'I am in sync now. You can keep the pace up.', variant: 'celebration', pose: 'spark', stage },
    ]
  }
  if (stage >= 1) {
    return [
      { title: 'DG', body: 'Alright, let us work through this batch together.', variant: 'focus', pose: 'wave', stage },
      { title: 'DG', body: 'I am ready. You take the first step.', variant: 'focus', pose: 'wave', stage },
    ]
  }
  return [
    { title: 'DG', body: 'I am here. Keep tapping in and I will warm up.', variant: 'neutral', pose: 'idle', stage },
    { title: 'DG', body: 'Tap back in a little more and I will wake up faster.', variant: 'neutral', pose: 'idle', stage },
  ]
}

function fallbackTriggerPool(trigger, language, topic, stage) {
  const normalizedTopic = topic || 'all'
  const zh = language === 'zh-TW'
  if (trigger === 'correct' && normalizedTopic === 'languages') {
    return zh
      ? [
          { title: 'DG', body: '漂亮，你的語言直覺正在穩下來。', variant: 'celebration', pose: 'nod', stage },
          { title: 'DG', body: '這題很順，語言這側開始有感了。', variant: 'celebration', pose: 'nod', stage },
        ]
      : [
          { title: 'DG', body: 'Nice catch. Your language instincts are settling in.', variant: 'celebration', pose: 'nod', stage },
          { title: 'DG', body: 'That was clean. The language side is starting to click.', variant: 'celebration', pose: 'nod', stage },
        ]
  }
  if (trigger === 'return' && normalizedTopic === 'docker') {
    return zh
      ? [
          { title: 'DG', body: 'Docker 這輪回來了，我們把它穩穩轉起來。', variant: 'focus', pose: 'wave', stage },
          { title: 'DG', body: '回到 docker，這次把容器掌握乾淨。', variant: 'focus', pose: 'wave', stage },
        ]
      : [
          { title: 'DG', body: 'Docker is back on deck. We can spin this up cleanly.', variant: 'focus', pose: 'wave', stage },
          { title: 'DG', body: 'Back to docker. Let us keep the containers under control.', variant: 'focus', pose: 'wave', stage },
        ]
  }

  const pools = {
    correct: zh
      ? [
          { title: 'DG', body: stage >= 1 ? '這題很乾淨，感覺開始黏住了。' : '漂亮，這一題先收下來。', variant: 'celebration', pose: 'nod', stage },
          { title: 'DG', body: stage >= 1 ? '對，就是這種手感，先記住。' : '很好，把這個感覺帶到下一題。', variant: 'celebration', pose: 'nod', stage },
        ]
      : [
          { title: 'DG', body: stage >= 1 ? 'That was clean. I can tell this is starting to stick.' : 'Nice hit. Hold on to that feeling for the next one.', variant: 'celebration', pose: 'nod', stage },
          { title: 'DG', body: stage >= 1 ? 'Yes, that is the feeling. Keep it for the next card.' : 'Good catch. Bring that same energy into the next card.', variant: 'celebration', pose: 'nod', stage },
        ],
    wrong: zh
      ? [
          { title: 'DG', body: stage >= 1 ? '沒關係，這種差一點最值得記。' : '先抓住差異，下一輪就會好很多。', variant: 'warning', pose: 'think', stage },
          { title: 'DG', body: stage >= 1 ? '先記住差異，下次會更穩。' : '這題先別怕，等下再看一次。', variant: 'warning', pose: 'think', stage },
        ]
      : [
          { title: 'DG', body: stage >= 1 ? 'That is okay. These almost-right misses are worth keeping.' : 'Just hold on to the difference. The next pass will feel steadier.', variant: 'warning', pose: 'think', stage },
          { title: 'DG', body: stage >= 1 ? 'Keep the difference in view. The next pass will feel steadier.' : 'Do not worry about this one yet. We can loop back cleanly.', variant: 'warning', pose: 'think', stage },
        ],
    learn_break: zh
      ? [
          { title: 'DG', body: stage >= 1 ? '這輪收得不錯，先讓腦袋留點空間。' : '先休息一下，下一輪不急。', variant: 'focus', pose: 'rest', stage },
          { title: 'DG', body: stage >= 1 ? '停一下剛剛好，讓記憶沉一沉。' : '這裡先停一下，等等再回來。', variant: 'focus', pose: 'rest', stage },
        ]
      : [
          { title: 'DG', body: stage >= 1 ? 'That batch landed well. Give your brain a little room now.' : 'Take a short beat. The next batch can wait.', variant: 'focus', pose: 'rest', stage },
          { title: 'DG', body: stage >= 1 ? 'A short pause is right. Let the last few cards settle.' : 'Pause here for a moment. The next round is fine waiting.', variant: 'focus', pose: 'rest', stage },
        ],
    review_complete: zh
      ? [
          { title: 'DG', body: stage >= 1 ? '這輪複習收得很漂亮，節奏在成形了。' : '這輪複習完成了，先讓它停一下。', variant: 'celebration', pose: 'spark', stage },
          { title: 'DG', body: stage >= 1 ? '很好，讓這一輪在腦中沉下去。' : '複習做完了，現在先喘一口氣。', variant: 'celebration', pose: 'spark', stage },
        ]
      : [
          { title: 'DG', body: stage >= 1 ? 'That review batch closed out nicely. I can feel the loop settling in.' : 'That review batch is done. Take a moment and let it settle.', variant: 'celebration', pose: 'spark', stage },
          { title: 'DG', body: stage >= 1 ? 'Nice finish. Let that review loop settle in a bit.' : 'Review complete. Take a breath before you move on.', variant: 'celebration', pose: 'spark', stage },
        ],
    return: zh
      ? [
          { title: 'DG', body: stage >= 1 ? '你回來了，我們從這裡接著走。' : '下一輪已經準備好了。', variant: 'focus', pose: 'wave', stage },
          { title: 'DG', body: stage >= 1 ? '剛剛那條線還在，現在可以繼續。' : '好，現在可以重新開始。', variant: 'focus', pose: 'wave', stage },
        ]
      : [
          { title: 'DG', body: stage >= 1 ? 'You are back. We can pick up the thread from here.' : 'Alright, the next round is ready.', variant: 'focus', pose: 'wave', stage },
          { title: 'DG', body: stage >= 1 ? 'That thread is still here. We can keep going now.' : 'Okay, we can start fresh from here.', variant: 'focus', pose: 'wave', stage },
        ],
  }

  return pools[trigger] || fallbackClickPool(language, topic, stage)
}

function shouldEmitFallbackReaction(trigger, now) {
  if (['learn_break', 'review_complete', 'return'].includes(trigger)) {
    return true
  }

  if (fallbackPetState.lastReactionAt) {
    const last = new Date(fallbackPetState.lastReactionAt)
    if (now.getTime() - last.getTime() < 6000) {
      return false
    }
  }

  const seed = Math.floor(now.getTime() / 60000) + fallbackPetState.bondXp + trigger.length
  if (trigger === 'correct') {
    return seed % 3 !== 0
  }
  if (trigger === 'wrong') {
    return seed % 2 === 0
  }
  return true
}

const hasBackend = () => typeof window !== 'undefined' && typeof window.go !== 'undefined'

export function __resetFallbackState() {
  fallbackDashboard.preferredLanguage = 'zh-TW'
  fallbackDashboard.stats = {
    studiedToday: 1,
    correctRate: 1,
  }
  fallbackDashboard.summary = {
    nextReviewAt: '2026-04-05T21:00:00+08:00',
    studyStreak: 3,
    weakTopics: structuredClone(topicWeakTopics.all),
    topicProgress: structuredClone(topicProgressByMode.all),
    weakestDeck: structuredClone(weakestDeckByMode.all),
  }
  fallbackDashboard.notificationSettings = {
    style: 'playful',
    titleMode: 'prefer_manual',
  }
  fallbackDashboard.scheduleSettings = {
    notificationIntervalMinutes: 20,
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
  fallbackPetState.bondXp = 0
  fallbackPetState.stage = 0
  fallbackPetState.lastInteractionAt = ''
  fallbackPetState.lastReactionAt = ''
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
  fallbackPetState.bondXp += isCorrect ? 2 : 1
  fallbackPetState.stage = fallbackPetState.bondXp >= 16 ? 2 : fallbackPetState.bondXp >= 6 ? 1 : 0

  return {
    cardId,
    isCorrect,
    correctAnswer: fallbackDashboard.currentCard.questionType === 'true-false' ? 'true' : '1',
    feedback:
      preferredLanguage === 'zh-TW'
        ? isCorrect
          ? '????哨???'
          : '?謕??????綜窖??'
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

export async function startLearnBreak() {
  if (hasBackend()) {
    return StartLearnBreak()
  }

  const unlockAt = new Date(Date.now() + fallbackDashboard.scheduleSettings.notificationIntervalMinutes * 60 * 1000)

  return {
    message: `Learn break started until ${unlockAt.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}.`,
    unlockAt: unlockAt.toISOString(),
    durationMinutes: fallbackDashboard.scheduleSettings.notificationIntervalMinutes,
  }
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
    weakestDeck: null,
  }
  fallbackDashboard.reviewQueue = []
  fallbackPetState.bondXp = 0
  fallbackPetState.stage = 0
  fallbackPetState.lastInteractionAt = ''
  fallbackPetState.lastReactionAt = ''
  return { message: 'Study data reset.' }
}

export async function interactWithDG() {
  if (hasBackend()) {
    return InteractWithDG()
  }

  const now = new Date('2026-04-05T10:00:00+08:00')
  if (fallbackPetState.lastInteractionAt) {
    const last = new Date(fallbackPetState.lastInteractionAt)
    if (now.getTime() - last.getTime() < 15000) {
      return {
        title: 'DG',
        body: fallbackDashboard.preferredLanguage === 'zh-TW' ? '我有聽到，先讓我喘一口氣。' : 'I heard you. Give me a beat.',
        variant: 'focus',
        stage: fallbackPetState.stage,
      }
    }
  }

  fallbackPetState.bondXp += 1
  fallbackPetState.stage = fallbackPetState.bondXp >= 16 ? 2 : fallbackPetState.bondXp >= 6 ? 1 : 0
  fallbackPetState.lastInteractionAt = now.toISOString()
  fallbackPetState.lastReactionAt = now.toISOString()
  return pickFallbackPetReaction(
    'clicked',
    fallbackClickPool(fallbackDashboard.preferredLanguage, fallbackDashboard.selectedTopic, fallbackPetState.stage),
  )
}

export async function getDGReaction(trigger) {
  if (hasBackend()) {
    return GetDGReaction(trigger)
  }

  const now = new Date('2026-04-05T10:00:00+08:00')
  if (!shouldEmitFallbackReaction(trigger, now)) {
    return { title: '', body: '', variant: '', pose: '', stage: fallbackPetState.stage }
  }

  fallbackPetState.lastReactionAt = now.toISOString()
  return pickFallbackPetReaction(
    trigger,
    fallbackTriggerPool(trigger, fallbackDashboard.preferredLanguage, fallbackDashboard.selectedTopic, fallbackPetState.stage),
  )
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
        questionTextZh: 'git rebase 主要在做什麼？',
        questionTextEn: 'What does git rebase mainly do?',
        clickbait: 'Do you really know the difference between rebase and merge?',
        clickbaitZh: '你真的知道 rebase 和 merge 差在哪嗎？',
        clickbaitEn: 'Do you really know the difference between rebase and merge?',
        reviewHint: 'Rebase replays commits on top of another base.',
        reviewHintZh: 'Rebase = 把 commits 重新套到另一條 base 上。',
        reviewHintEn: 'Rebase replays commits on top of another base.',
        explanationZh: '`git rebase` 會把你的 commits 重新套用到另一條 base 之上。',
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
