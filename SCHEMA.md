# Content And Data Schema

## 1. Purpose

This document defines the concrete schema for:

- `knowledge/` Markdown card files
- `data/` runtime storage files

The goal is to make parsing, scheduling, review logic, and future UI implementation predictable from day one.

## 2. Directory Layout

```text
knowledge/
  git/
    rebase.md
    cherry-pick.md

data/
  settings.json
  cards-cache.json
  progress.json
  attempts.jsonl
  import-errors.json
```

## 3. `knowledge/` Card File Schema

## 3.1 Design Rules

Each card file should:

- be a single `.md` file
- contain exactly one study card in MVP
- use YAML frontmatter
- include both Chinese and English explanation sections after frontmatter

MVP parsing rule:

- one file = one card
- one file contains both `zh-TW` and `en` content sections

This keeps the importer simple and makes file-level editing easy.

## 3.2 Required Frontmatter Fields

Required fields:

- `id`
- `type`
- `answer`

Required body:

- Markdown body must contain both `zh-TW` and `en` explanation sections

Localization rule:

- visible learning content should provide both `zh-TW` and `en` values
- for MVP, use bilingual frontmatter fields such as `title_zh` / `title_en`, `question_zh` / `question_en`
- legacy single-language fields like `title`, `question`, `choices`, `clickbait`, and `review_hint` are still accepted as fallback values

## 3.3 Optional Frontmatter Fields

Optional fields:

- `tags`
- `difficulty`
- `clickbait`
- `clickbait_zh`
- `clickbait_en`
- `choices`
- `choices_zh`
- `choices_en`
- `aliases`
- `source`
- `review_hint`
- `review_hint_zh`
- `review_hint_en`
- `confusion_with`
- `metaphor_seed`
- `hook_style_tags`
- `enabled`
- `body_format`

## 3.4 Field Definitions

### `id`

Type:

- string

Rules:

- globally unique across all imported cards
- recommended lowercase kebab-case
- stable over time once created

Example:

```yaml
id: git-rebase-vs-merge
```

### `title`

Type:

- string

Purpose:

- short label shown in UI

Example:

```yaml
title: Rebase vs Merge
```

Recommended bilingual fields:

```yaml
title_zh: Rebase 跟 Merge 的差別
title_en: Rebase vs Merge
```

### `type`

Type:

- string enum

Allowed MVP values:

- `single-choice`
- `true-false`

Example:

```yaml
type: single-choice
```

### `body_format`

Type:

- string enum

Allowed MVP values:

- `bilingual-section`

Default:

- `bilingual-section`

Purpose:

- tells the parser how the body is organized

### `question`

Type:

- string

Purpose:

- the prompt shown after reading the explanation

Recommended bilingual fields:

```yaml
question_zh: "git rebase 主要是在做什麼？"
question_en: "What does git rebase mainly do?"
```

### `answer`

Type depends on question type:

- `single-choice`: integer index, zero-based
- `true-false`: boolean

Examples:

```yaml
answer: 1
```

```yaml
answer: true
```

### `choices`

Type:

- array of strings

Required when:

- `type: single-choice`

Rules:

- minimum 2 choices
- maximum 6 choices for MVP
- `answer` must point to a valid index

Example:

```yaml
choices:
  - Creates a merge commit
  - Replays commits onto a new base
  - Deletes conflict history
```

Recommended bilingual fields:

```yaml
choices_zh:
  - 建立一個 merge commit
  - 把 commits 重新接到新的 base 上
choices_en:
  - Creates a merge commit
  - Replays commits onto a new base
```

### `tags`

Type:

- array of strings

Purpose:

- topic grouping
- filtering
- weak-topic reporting

Example:

```yaml
tags: [git, branching]
```

### `difficulty`

Type:

- integer

Allowed range:

- `1` to `5`

Suggested meaning:

- `1`: very basic
- `3`: intermediate
- `5`: advanced

Default if omitted:

- `2`

### `clickbait`

Type:

- string

Purpose:

- preferred notification title

Example:

```yaml
clickbait: Most developers misuse rebase. Do you?
```

Recommended bilingual fields:

```yaml
clickbait_zh: 你真的懂 rebase 跟 merge 的差別嗎？
clickbait_en: Most developers misuse rebase. Do you?
```

### `aliases`

Type:

- array of strings

Purpose:

- alternate search terms
- future deduplication or clustering support

### `source`

Type:

- string

Purpose:

- optional source note, book, article, or internal origin

### `review_hint`

Type:

- string

Purpose:

- short memory aid shown after an incorrect answer

Recommended bilingual fields:

```yaml
review_hint_zh: Rebase = 把 commits 重放到新的 base 上。
review_hint_en: Rebase = replay commits on top of another base.
```

### `confusion_with`

Type:

- array of strings

Purpose:

- identify commonly confused related cards or commands
- improve review weighting and hook generation later

Example:

```yaml
confusion_with: [git-fetch-basic, git-pull-basic]
```

### `metaphor_seed`

Type:

- array of strings

Purpose:

- give the offline hook generator a small pool of analogy or metaphor hints

Example:

```yaml
metaphor_seed: [購物車, 暫存, 先觀察]
```

### `hook_style_tags`

Type:

- array of strings

Purpose:

- guide offline hook-generation templates
- examples: `misunderstood`, `fear_of_mistake`, `safer_first`, `comparison`

Example:

```yaml
hook_style_tags: [misunderstood, safer_first]
```

### `enabled`

Type:

- boolean

Default:

- `true`

Purpose:

- temporarily disable a card without deleting the file

## 3.5 Body Rules

The Markdown body is bilingual explanation content.

Rules:

- should be concise enough to read in 30 to 90 seconds
- recommended length: 1 to 5 short paragraphs
- may include inline code and bullet points
- should explain the concept needed to answer the question
- must contain these exact section headings:
  - `## zh-TW`
  - `## en`
- UI reads only one language section at a time, based on current user selection

Example body structure:

```md
## zh-TW

這裡放中文解釋。

## en

Put the English explanation here.
```

## 3.6 Valid Example: Single Choice

```md
---
id: git-rebase-vs-merge
title: Rebase vs Merge
type: single-choice
body_format: bilingual-section
tags: [git, branching]
difficulty: 2
question: "What does git rebase mainly do?"
choices:
  - "Creates a merge commit between branches"
  - "Replays commits onto a new base"
  - "Deletes conflicting commits automatically"
answer: 1
clickbait: "Most Git beginners misunderstand rebase. Do you?"
review_hint: "Rebase = replay commits on top of another base."
enabled: true
---

## zh-TW

`git rebase` 會把目前分支上的提交重新套用到另一個 base commit 上。
它常用來讓提交歷史更線性、更乾淨。
和 `merge` 不同，它通常不會產生額外的 merge commit。

## en

`git rebase` takes commits from your current branch and reapplies them onto another base commit.
It is often used to keep history linear and easier to read.
Unlike `merge`, it usually avoids creating an extra merge commit.
```

## 3.7 Valid Example: True/False

```md
---
id: git-fast-forward-merge
title: Fast-Forward Merge
type: true-false
body_format: bilingual-section
tags: [git, merge]
difficulty: 1
question: "A fast-forward merge creates a new merge commit every time."
answer: false
clickbait: "This Git merge behavior is simpler than most people think"
enabled: true
---

## zh-TW

當目標分支沒有額外分叉，只是單純往前推進時，Git 可以直接把指標往前移動。
這種情況叫做 fast-forward merge，通常不會建立新的 merge commit。

## en

When the target branch has not diverged, Git can simply move the branch pointer forward.
That is called a fast-forward merge, and it usually does not create a new merge commit.
```

## 3.8 Invalid Cases

Invalid examples:

- missing `id`
- duplicate `id`
- `single-choice` without `choices`
- `answer` index out of bounds
- empty body
- missing `## zh-TW` section
- missing `## en` section
- unsupported `type`

## 4. Import Validation Rules

Validation order:

1. file extension must be `.md`
2. frontmatter must parse successfully
3. required fields must exist
4. `type` must be supported
5. answer structure must match question type
6. body must contain both bilingual sections
7. `id` must be unique

On validation failure:

- skip card import
- add an entry to `data/import-errors.json`

## 5. `data/settings.json` Schema

Purpose:

- store app preferences and scheduling configuration

Suggested schema:

```json
{
  "version": 1,
  "knowledge_directories": [
    "D:\\duolin-gogo\\knowledge"
  ],
  "notification_interval_minutes": 10,
  "active_hours": {
    "enabled": true,
    "start": "09:00",
    "end": "22:00"
  },
  "review_schedule": {
    "mode": "daily",
    "weekday": null,
    "time": "21:00",
    "batch_size": 8
  },
  "language": {
    "default": "zh-TW",
    "allow_toggle": true
  },
  "notifications": {
    "style": "playful"
  },
  "study_rules": {
    "max_new_cards_per_day": 12,
    "snooze_minutes": 15,
    "cooldown_after_answer_minutes": 5
  },
  "ui": {
    "minimize_to_tray": true,
    "open_on_notification_click": true
  }
}
```

### Field Notes

- `version`: schema version for migration later
- `knowledge_directories`: list of absolute paths
- `notification_interval_minutes`: allowed MVP range `5` to `120`
- `active_hours.start/end`: 24-hour `HH:MM`
- `review_schedule.mode`: `daily`, `weekly`, or `off`
- `review_schedule.weekday`: `mon` to `sun`, only used in weekly mode
- `batch_size`: recommended `5` to `20`
- `language.default`: allowed MVP values `zh-TW` or `en`
- `notifications.style`: allowed initial values `safe`, `playful`, `aggressive`, `chaotic`

## 6. `data/cards-cache.json` Schema

Purpose:

- store parsed card snapshots so runtime does not need to fully reparse everything on every action

Suggested schema:

```json
{
  "version": 1,
  "generated_at": "2026-04-04T18:30:00+08:00",
  "cards": [
    {
      "id": "git-rebase-vs-merge",
      "source_path": "D:\\duolin-gogo\\knowledge\\git\\rebase.md",
      "source_modified_at": "2026-04-04T18:15:00+08:00",
      "source_hash": "sha256:example",
      "enabled": true,
      "title": "Rebase vs Merge",
      "type": "single-choice",
      "body_format": "bilingual-section",
      "tags": ["git", "branching"],
      "difficulty": 2,
      "question": "What does git rebase mainly do?",
      "choices": [
        "Creates a merge commit between branches",
        "Replays commits onto a new base",
        "Deletes conflicting commits automatically"
      ],
      "answer": 1,
      "clickbait": "Most Git beginners misunderstand rebase. Do you?",
      "review_hint": "Rebase = replay commits on top of another base.",
      "body_markdown_zh": "`git rebase` 會把目前分支上的提交重新套用到另一個 base commit 上。",
      "body_markdown_en": "`git rebase` takes commits from your current branch and reapplies them onto another base commit.",
      "body_plaintext_zh": "git rebase 會把目前分支上的提交重新套用到另一個 base commit 上。",
      "body_plaintext_en": "git rebase takes commits from your current branch and reapplies them onto another base commit."
    }
  ]
}
```

### Notes

- `source_hash` can be file-content hash or omitted in first implementation
- bilingual plaintext fields are optional but useful for previews, search, or notifications
- cache should only include valid cards

## 7. `data/progress.json` Schema

Purpose:

- store the latest state for each card

Suggested schema:

```json
{
  "version": 1,
  "updated_at": "2026-04-04T18:35:00+08:00",
  "cards": {
    "git-rebase-vs-merge": {
      "seen_count": 4,
      "correct_count": 3,
      "wrong_count": 1,
      "mastery_score": 2,
      "streak_correct": 2,
      "last_seen_at": "2026-04-04T17:10:00+08:00",
      "last_correct_at": "2026-04-04T17:10:12+08:00",
      "last_wrong_at": "2026-04-03T20:03:01+08:00",
      "last_session_type": "learn",
      "introduced_at": "2026-04-02T10:00:00+08:00",
      "next_review_at": "2026-04-05T21:00:00+08:00",
      "snoozed_until": null,
      "is_mastered": false
    }
  },
  "daily_summary": {
    "2026-04-04": {
      "cards_seen": 6,
      "answered": 6,
      "correct": 4,
      "wrong": 2,
      "review_answered": 2,
      "learn_answered": 4
    }
  }
}
```

### Per-card Field Notes

- `mastery_score`: integer range suggested `-5` to `10`
- `streak_correct`: number of consecutive correct answers
- `last_session_type`: `learn` or `review`
- `introduced_at`: first time card entered study rotation
- `snoozed_until`: temporary delay after user ignores or snoozes
- `is_mastered`: lightweight UI hint, not absolute truth

## 8. `data/attempts.jsonl` Schema

Purpose:

- append-only event history of question attempts

One JSON object per line.

Suggested fields:

```json
{"attempt_id":"att_20260404_001","card_id":"git-rebase-vs-merge","session_type":"learn","shown_at":"2026-04-04T17:10:00+08:00","answered_at":"2026-04-04T17:10:12+08:00","selected_answer":1,"is_correct":true,"response_time_ms":12000,"mastery_delta":1}
{"attempt_id":"att_20260404_002","card_id":"git-fast-forward-merge","session_type":"review","shown_at":"2026-04-04T21:00:00+08:00","answered_at":"2026-04-04T21:00:08+08:00","selected_answer":true,"is_correct":false,"response_time_ms":8000,"mastery_delta":-2}
```

### Field Definitions

- `attempt_id`: unique event id
- `card_id`: references imported card id
- `session_type`: `learn` or `review`
- `shown_at`: when UI displayed the card
- `answered_at`: when user submitted answer
- `selected_answer`: integer or boolean depending on type
- `is_correct`: boolean
- `response_time_ms`: integer
- `mastery_delta`: integer change applied to progress state

## 9. `data/import-errors.json` Schema

Purpose:

- keep importer diagnostics visible and debuggable

Suggested schema:

```json
{
  "version": 1,
  "generated_at": "2026-04-04T18:30:00+08:00",
  "errors": [
    {
      "source_path": "D:\\duolin-gogo\\knowledge\\git\\bad-card.md",
      "code": "missing_required_field",
      "field": "question",
      "message": "Required field 'question' is missing."
    },
    {
      "source_path": "D:\\duolin-gogo\\knowledge\\git\\weird.md",
      "code": "answer_out_of_range",
      "field": "answer",
      "message": "Answer index 4 exceeds available choices."
    }
  ]
}
```

### Suggested Error Codes

- `frontmatter_parse_failed`
- `missing_required_field`
- `unsupported_type`
- `missing_choices`
- `answer_out_of_range`
- `empty_body`
- `missing_language_section`
- `duplicate_id`

## 10. In-Memory App Model

During runtime, it will help to think in terms of these core models.

### Card

```json
{
  "id": "string",
  "title": "string",
  "title_zh": "string",
  "title_en": "string",
  "type": "single-choice | true-false",
  "question": "string",
  "question_zh": "string",
  "question_en": "string",
  "choices": ["string"],
  "choices_zh": ["string"],
  "choices_en": ["string"],
  "answer": "number | boolean",
  "tags": ["string"],
  "difficulty": 2,
  "clickbait": "string",
  "clickbait_zh": "string",
  "clickbait_en": "string",
  "review_hint": "string",
  "review_hint_zh": "string",
  "review_hint_en": "string",
  "confusion_with": ["string"],
  "metaphor_seed": ["string"],
  "hook_style_tags": ["string"],
  "body_format": "bilingual-section",
  "body_markdown_zh": "string",
  "body_markdown_en": "string",
  "source_path": "string",
  "enabled": true
}
```

### ProgressState

```json
{
  "card_id": "string",
  "seen_count": 0,
  "correct_count": 0,
  "wrong_count": 0,
  "mastery_score": 0,
  "streak_correct": 0,
  "last_seen_at": null,
  "last_correct_at": null,
  "last_wrong_at": null,
  "next_review_at": null,
  "snoozed_until": null
}
```

### AttemptEvent

```json
{
  "attempt_id": "string",
  "card_id": "string",
  "session_type": "learn",
  "shown_at": "timestamp",
  "answered_at": "timestamp",
  "selected_answer": 1,
  "is_correct": true,
  "response_time_ms": 12000,
  "mastery_delta": 1
}
```

## 11. Default Values

Recommended defaults when fields are omitted:

- `difficulty`: `2`
- `tags`: `[]`
- `clickbait`: generated fallback
- `review_hint`: `""`
- `enabled`: `true`
- `body_format`: `bilingual-section`
- `choices`: `[]` only for `true-false`

## 12. Naming And Stability Rules

Rules:

- `id` should never be auto-regenerated if the file is renamed
- progress is keyed by `id`, not file path
- file path changes should preserve learning history if the `id` is unchanged
- deleting a card file should not immediately erase its history from `progress.json`

This prevents accidental progress loss during note organization.

## 13. Recommended Parser Behavior

When scanning files:

1. read Markdown file
2. parse YAML frontmatter
3. validate schema
4. trim body
5. split body into bilingual sections
6. normalize defaults
7. add valid card to cache
8. write invalid result to import errors

Normalization suggestions:

- trim strings
- lowercase enum values
- normalize tags to lowercase
- preserve original markdown body
- preserve both language bodies separately after parsing

## 14. Recommended Future-Compatible Extensions

Fields we may add later without breaking the MVP:

- `explanation_short`
- `code_snippet`
- `related_card_ids`
- `estimated_read_seconds`
- `memory_level`
- `last_notification_at`
- `ignore_count`
- `hook_body_zh`
- `hook_body_en`

For now, keep parser permissive for unknown fields and ignore them.

## 15. Final Recommendation

For MVP, lock in these principles:

- one Markdown file = one card
- YAML frontmatter + bilingual Markdown body
- `progress.json` as current snapshot
- `attempts.jsonl` as append-only event history
- `settings.json` as user config
- `import-errors.json` as parser diagnostics

This is simple enough to implement quickly, while still stable enough to support scheduling, weak-topic weighting, and future migration to SQLite if needed.
