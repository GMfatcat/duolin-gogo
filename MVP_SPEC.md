# Local Knowledge Duolingo MVP Spec

## 1. Product Summary

This product is a local-first micro-learning app that turns personal Markdown knowledge notes into timed learning cards, review quizzes, and adaptive repetition.

Primary goal:

- Help users continuously learn technical knowledge through short, interrupt-driven study sessions based on bilingual Markdown notes.

Product positioning:

- "A local Duolingo for technical knowledge, powered by your own Markdown notes."

## 2. Target Users

- Developers who keep personal notes in Markdown
- Learners who want repeated exposure instead of long study sessions
- Users who prefer local storage over cloud-first note systems

## 3. MVP Goals

- Import and parse local Markdown knowledge files
- Trigger a learning notification every fixed interval, default every 10 minutes
- Open a lightweight learning card when the notification is clicked
- Show one concept explanation and one quick question
- Track what the user studied today
- Run scheduled review sessions daily or weekly
- Increase repeat frequency for concepts with lower accuracy
- Support offline notification hook generation from card metadata and local templates

## 4. Non-Goals For MVP

- Multi-device sync
- Account system
- Online collaboration
- AI-generated content as a required dependency
- Complex spaced repetition algorithms like full Anki/SM-2 parity
- Rich gamification beyond basic streak and progress counters

## 5. Core User Stories

1. As a user, I want to point the app to a local Markdown folder so the app can build my study deck.
2. As a user, I want to receive short notifications during the day so I keep studying in small bursts.
3. As a user, I want the notification title to be curiosity-driven so I feel like clicking it.
4. As a user, I want to read a short explanation before answering a question.
5. As a user, I want each card to contain both Chinese and English learning content so I can switch depending on what I want to practice.
6. As a user, I want the app to remember what I already learned today.
7. As a user, I want a daily or weekly review mode to test older material.
8. As a user, I want weak concepts to show up more often so I can improve where I struggle.
9. As a user, I want all data stored locally.

## 6. MVP Experience Flow

### 6.1 First-Time Setup

1. User chooses a local knowledge folder.
2. App scans Markdown files and parses valid learning cards.
3. App stores parsed cards in the local database.
4. User configures:
   - notification interval
   - daily review time or weekly review schedule
   - optional active study hours

### 6.2 Learning Notification Flow

1. App picks the next card using weighted selection.
2. App shows an OS notification with a clickbait-style question.
3. User clicks notification.
4. App opens a learning window/card.
5. App enters a staged study flow:
   - `Learn`: show title and short concept explanation in the selected UI language
   - `Answer`: hide the explanation and show the question with answer options
   - `Feedback`: show correctness, hint, and a clear next-step action
6. User answers.
7. App records:
   - shown timestamp
   - answer correctness
   - response time if available
   - mastery update

### 6.3 Review Flow

1. At a configured time, app enters review mode.
2. Review mode selects cards from recently studied material.
3. User answers a small batch, for example 5 to 10 questions.
4. App updates mastery and future appearance probability.

## 7. Content Model

Each Markdown file contains one card. Each card includes both Chinese and English learning content inside the same Markdown file.

Current content focus:

- Git cards first

Planned near-term Git card coverage:

- `git add`
- `git commit`
- `git merge`
- `git fetch`
- `git pull`
- `git checkout`
- `git status`
- `git switch`
- `git restore`
- `git stash`

Recommended Markdown format:

```md
---
id: git-rebase-vs-merge
title: Rebase vs Merge
tags: [git, branching]
difficulty: 2
type: single-choice
question: "What does git rebase mainly do?"
choices:
  - "Creates a merge commit between two branches"
  - "Replays commits onto a new base"
  - "Deletes conflicting commits automatically"
answer: 1
clickbait: "Most Git beginners misunderstand rebase. Do you?"
body_format: bilingual-section
---

## zh-TW

`git rebase` 會把目前分支的提交重新套用到另一個 base 之上。
它常用來讓提交歷史更線性、更乾淨。

## en

`git rebase` replays commits from the current branch onto another base commit.
It is commonly used to keep history linear and cleaner than frequent merge commits.
```

Required fields:

- `id`
- `title`
- `question`
- `answer`
- bilingual explanation body content

Optional fields:

- `tags`
- `difficulty`
- `choices`
- `clickbait`
- `type`
- `body_format`

Supported question types for MVP:

- `single-choice`
- `true-false`

Question types deferred until later:

- fill-in-the-blank
- free text
- code ordering

## 8. Functional Requirements

### 8.1 Markdown Import

- User can choose one or more local folders to scan
- App recursively finds `.md` files
- App parses frontmatter and body
- Invalid files are skipped and shown in an import error list
- App supports re-scan to detect added, edited, or removed notes

### 8.2 Notification System

- Default interval: every 20 minutes
- User can configure interval, for example 10, 15, 30, or 60 minutes
- Notifications only fire during configured active hours
- Notification title should prefer `clickbait` field if available
- Fallback notification title should be generated from question/title templates
- Clicking notification opens the learning card directly
- Closing the main window should hide the app to the background instead of stopping notifications
- The app should keep running while hidden so interval notifications and review scheduling continue
- The first background-running implementation should expose a tray icon with explicit `Open` and `Exit` actions
- Fully quitting the app should require an explicit tray `Exit` action rather than the window close button

### 8.3 Learning Card UI

- Shows one concept at a time
- Explanation must be short and readable within about 30 to 90 seconds
- Card explanation must support both Chinese and English in the same source file
- App shell should support a global `zh-TW` / `en` language toggle for all UI copy except the product name `duolin-gogo`
- Card content language should follow the current global language by default
- Layout should prefer a two-column structure:
  - left side for the active study flow
  - right side for answer-related stats and weak-topic context only
- Utility controls and diagnostics should move out of the main sidebar and into a settings popout triggered from the top-right app chrome
- Authoring workflows should not live in the same popout as end-user settings
- The top-right chrome should provide separate utility entry points:
  - a settings trigger for runtime settings and operational actions
  - a library/authoring trigger for card preview and AI draft review
  - a diagnostics trigger for detailed deck health, batch validation, and recently changed cards
- Settings popout should also contain scheduling controls such as review time, notification interval, and active hours
- Settings popout should prefer a wider horizontal layout so the main controls fit without vertical scrolling in normal desktop use
- Settings popout should prefer concise health summaries instead of full authoring diagnostics detail
- Detailed diagnostics and batch validation should live in a dedicated diagnostics popout instead of stretching the authoring surface
- Diagnostics should default to a collapsed disclosure state, with health summarized inline beside the diagnostics title
- During the `Learn` phase, the explanation is visible and the question is hidden
- During the `Answer` phase, the question and answer choices are visible and the explanation is hidden
- User can submit an answer
- During the `Feedback` phase, the app immediately shows whether the answer is correct
- The feedback state should restore a short explanation or hint after submission
- The feedback state must provide a clear next-step action such as `Next card`
- In normal learn mode, the app may treat every 3 answered cards as a mini batch and then pause new cards until the next notification rhythm
- That mini-batch rest state should also show a lightweight summary of answered count, estimated accuracy, and the concept or topic most worth revisiting
- When the rest window ends, the first returning card should visibly signal that a fresh mini batch has started
- The first session/progress baseline should feel complete enough for daily use:
  - reduce short-term card repetition
  - give review batches explicit progress and completion states
  - give normal learn mode a bounded 3-card rhythm with a rest window
  - show a simple local streak for daily momentum
- Displayed times should use a human-readable local format like `2026-04-05 21:00`
- Avoid raw ISO timestamps in the visible UI
- Grouped topic guidance should be promoted closer to the top of the shell as a compact assistant-style hint instead of living only lower in the sidebar
- The assistant-style hint should evolve into a lightweight `DG` mascot surface that can speak short contextual guidance such as the weakest deck, review completion, or gentle encouragement
- Mascot motion should stay subtle and low-noise, favoring small entrance, idle, and collapse animations over marquee-like movement
- `DG` should be able to evolve into a hidden-growth companion:
  - users should not see explicit `level`, `xp`, or stat bars
  - growth should instead show up as richer reactions, more varied moods, and slightly more familiar companion behavior over time
  - growth should come from real study activity such as answered cards, completed mini batches, completed reviews, streak continuation, and occasional direct clicks on `DG`
  - click interaction should feel playful but should not become a spammy mini-game
  - the first implementation can stay lightweight: store hidden local growth, let card answers feed it, and let direct clicks on `DG` surface stage-based reactions
- Language and study-mode controls should converge toward compact dropdown selectors instead of parallel button groups
- User can dismiss and return later

### 8.4 Progress Tracking

- Track all answer attempts locally
- Track cards studied today
- Track per-card accuracy
- Track per-card last seen time
- Track simple streak count for daily activity

### 8.5 Review Sessions

- User can configure review cadence:
  - daily at a specific time
  - weekly on a specific weekday and time
- Review session contains cards from previously studied material
- Review session prioritizes:
  - recently wrong cards
  - low-accuracy cards
  - not-reviewed-recently cards

### 8.6 Adaptive Repetition

- Each card has a dynamic appearance score
- Lower accuracy increases score
- Long time since last seen increases score
- Recently answered correctly lowers score
- New cards have a moderate boost so onboarding continues

## 9. Weighted Selection Logic

For MVP, use a simple deterministic scoring model rather than a complex spaced repetition algorithm.

Suggested card priority score:

```text
priority =
  new_card_bonus +
  weak_concept_bonus +
  time_since_last_seen_bonus +
  review_due_bonus -
  mastered_penalty
```

Definitions:

- `new_card_bonus`: added when card has never been seen
- `weak_concept_bonus`: based on wrong rate or low accuracy
- `time_since_last_seen_bonus`: increases as elapsed time grows
- `review_due_bonus`: added when current time is past next review time
- `mastered_penalty`: reduces frequency for cards with consistently high accuracy

Simple mastery model:

- `mastery_score` starts at `0`
- correct answer: `+1`
- wrong answer: `-2`
- score floor: `-5`
- score cap: `10`

Suggested review scheduling:

- first correct: review in 1 day
- second correct: review in 3 days
- third correct: review in 7 days
- wrong answer: review again soon, for example within same day or next day

## 10. Data Model

### 10.1 Card

- `id`
- `source_path`
- `title`
- `body_markdown_zh`
- `body_markdown_en`
- `body_plaintext_zh`
- `body_plaintext_en`
- `question_type`
- `question_text`
- `choices_json`
- `answer_value`
- `tags_json`
- `difficulty`
- `clickbait_text`
- `body_format`
- `created_at`
- `updated_at`

### 10.2 Study Progress

- `card_id`
- `seen_count`
- `correct_count`
- `wrong_count`
- `mastery_score`
- `last_seen_at`
- `last_correct_at`
- `last_wrong_at`
- `next_review_at`

### 10.3 Attempt Log

- `id`
- `card_id`
- `session_type` (`learn` or `review`)
- `shown_at`
- `answered_at`
- `selected_answer`
- `is_correct`
- `response_time_ms`

### 10.4 App Settings

- `knowledge_directories`
- `notification_interval_minutes`
- `active_hours_start`
- `active_hours_end`
- `review_mode`
- `review_weekday`
- `review_time`
- `review_batch_size`
- `preferred_language`

## 11. Suggested Screens

### 11.1 Setup Screen

- select knowledge folder
- import status
- settings for notification interval and review schedule

### 11.2 Home Dashboard

- cards imported
- studied today
- correct rate today
- next review time
- weak topics summary

### 11.3 Learning Card Modal / Window

- card title
- global language toggle
- short explanation in `Learn` phase
- question and answers in `Answer` phase
- correctness + hint in `Feedback` phase
- explicit next-step action after submission

### 11.6 Proposed UI Refinement

- left panel: study flow only
- right panel: answer-related stats and weak topics only
- top-right chrome: language toggle plus a settings icon button
- settings popout: test notification, snooze, rescan, hook settings, review time, notification interval, active hours, import diagnostics
- separate library/authoring popout: authoring preview, AI draft review, detailed diagnostics, batch validation, recently changed cards
- import health should be summarized inline near the settings title, with detailed diagnostics only expanded when issues exist
- active-hours values should be editable from the same settings popout so users can debug notification timing without hand-editing JSON
- keep `duolin-gogo` as the primary fixed product label across languages
- a playful eyebrow or slogan line above the main title is allowed to vary independently from the product name
- all other shell copy should follow the selected global UI language
- visible time values should use `YYYY-MM-DD HH:mm` style formatting
- avoid duplicated `next review` summaries across multiple cards in the sidebar
- keep the highlighted `next review` value in one place only

### 11.8 DG Mascot Interaction

- `DG` should live near the top of the shell as a compact assistant bubble rather than inside the lower sidebar
- the first implementation should use a small avatar or badge plus one short context line
- the hint should be collapsible so users can reduce it to a compact badge
- the first animation set should be minimal:
  - subtle entrance fade/slide
  - lightweight collapse/expand motion
  - small feedback states for review completion or weak-deck reminders
- avoid continuous marquee or aggressive movement that competes with study content
- the next mascot layer can introduce hidden pet-style growth:
  - no visible pet stats
  - click-to-react interaction
  - unlockable reaction pools and mood variations
  - stronger sense of companionship after repeated real study activity

### 11.4 Review Session Screen

- current question index
- correctness feedback
- session summary

### 11.5 Import Diagnostics

- invalid files
- missing frontmatter fields
- duplicate IDs

### 11.7 Authoring Support

- a card preview view should help authors inspect one card before it enters the normal study flow
- preview should show parsed content, language switching, and validation results
- the first implementation can live inside the settings panel as an authoring preview block with file selection and per-card diagnostics
- AI-generated drafts should be reviewable against the same schema before saving into `knowledge/`
- the first AI-draft implementation can use a pasted Markdown textarea plus normalized preview and diagnostics before any save flow exists
- whole-deck validation should later support grouped warning/error review for batch AI generation

## 12. Local Storage Strategy

MVP recommendation:

- JSON/JSONL for structured app data
- Raw Markdown files remain the source of truth for content

Reasoning:

- Lower complexity and lighter runtime than adding SQLite immediately
- Easy review scheduling and attempt history for MVP scale
- Keeps user content editable in plain files
- Fits the resource-saving goal of the app

## 13. Notification Copy Strategy

Notification copy should feel curiosity-driven, not like a dry reminder.

Priority order:

1. Use localized `clickbait_zh` or `clickbait_en` from content
2. Generate a localized hook from offline templates based on card metadata
3. Fall back to title/question templates

Template examples:

- "Can you answer this before your coffee gets cold?"
- "Most developers get this Git concept wrong"
- "Quick check: do you really know this one?"
- "You have 30 seconds. Can you solve this?"

Recommended hook styles:

- `safe`
- `playful`
- `aggressive`
- `chaotic`

Recommended title-source modes:

- `prefer_manual`
- `prefer_generated`

Offline hook generator rules:

- must not require network access
- should use card metadata such as `tags`, `confusion_with`, `metaphor_seed`, and `hook_style_tags`
- should generate 1 to 3 candidate hook lines per card
- should prefer metaphor, analogy, contrast, or curiosity-gap framing over dry technical wording
- should still keep a recognizable relation to the underlying concept
- `chaotic` hooks may lean into headline, shopping, or light personality-quiz energy, but should still preserve an obvious conceptual link to the study card

## 14. Error Handling

- Missing required frontmatter fields: skip card and log error
- Duplicate `id`: import warning and skip newer or older based on clear policy
- Invalid answer index: mark file invalid
- Missing knowledge folder: app shows empty-state guidance
- Notification click without valid card: open dashboard instead

## 15. Metrics For MVP

Local metrics only, no analytics backend required.

Useful measurements:

- cards studied today
- review completion rate
- per-topic accuracy
- active days streak
- total attempts
- weakest 5 concepts

The first streak implementation can stay local and simple:

- count consecutive calendar days with at least one answered card
- show the streak in the sidebar and in session wrap-up states

## 16. Acceptance Criteria

The MVP is successful if:

- User can import at least 20 Markdown cards from a local folder
- App can send a timed notification on the local machine
- Clicking a notification opens one study card
- User can answer a question and get immediate feedback
- The app records study activity and accuracy locally
- The app can run at least one scheduled review session
- Cards with low accuracy appear more often than mastered cards

## 17. Suggested Tech Evaluation Criteria

When choosing the framework, compare against:

- local file system access quality
- system notification support
- scheduling reliability while app is running
- lightweight desktop footprint
- SQLite support
- simple packaging for Windows first
- future support for macOS if needed

## 18. Background Running Requirement

For Windows-first usage, the app should behave like a lightweight background utility:

- clicking the window close button should hide the window instead of exiting
- hidden app instances should continue running notifications and review scheduling
- the app should be restorable from the Windows notification area tray
- the tray menu should provide a minimal first-phase control surface:
  - `Open duolin-gogo`
  - `Exit`

The first implementation may defer minimize-to-tray behavior if close-to-tray is stable first.

## 19. Recommended MVP Delivery Order

### Phase 1: Core Content + Study Loop

- parse Markdown cards
- store cards in local DB
- basic dashboard
- open one study card manually
- record answers

### Phase 2: Notifications + Scheduling

- timed notifications
- click notification to open card
- active hours support

### Phase 3: Review + Adaptive Weighting

- review session generation
- weak-topic weighting
- simple streak and stats

### Phase 4: Polish

- import diagnostics
- cleaner card UI
- better notification copy
- staged study flow with clearer next actions
- global i18n across shell copy

## 20. Open Decisions For Framework Discussion

- Desktop shell: Tauri or Electron
- Frontend: React, Svelte, or another lightweight UI layer
- Database: SQLite from day one or JSON-first prototype
- Scheduler: in-app timer only or OS-level scheduled trigger fallback
- Markdown parser: strict frontmatter schema vs forgiving parser

## 21. Recommendation

For this product shape, the MVP should optimize for:

- local-first behavior
- fast interaction
- low idle resource usage
- easy parsing of Markdown and persistence of study history

That means the likely strongest MVP direction is:

- desktop app
- local Markdown as source content
- JSON/JSONL for progress
- simple weighted repetition instead of a full academic SRS model
