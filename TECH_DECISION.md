# Technical Decision Notes

## 1. Recommended Stack

Recommended MVP stack:

- Desktop shell: `Wails`
- Core runtime: `Go`
- Frontend: `Vue 3`
- Storage: `JSON + JSONL`
- Content source: local `Markdown`
- Notification style: Windows toast notification

This stack is chosen to optimize for:

- low memory usage
- local-first workflow
- simple packaging for Windows
- easy filesystem access
- small frontend surface area

## 2. Why This Stack Fits The Product

This product is not primarily a website. It is a local desktop utility that:

- watches or rescans local Markdown files
- keeps study history locally
- runs lightweight scheduling logic
- shows Windows-native notifications
- opens a focused study card window

Current content starts with Git, but the app shell and data model are intended to expand to broader technical topics later.

The first substantial content expansion after the MVP baseline is still Git-focused, but the notification layer now needs to support more expressive offline hooks rather than only literal study prompts.

That means we need:

- good local file access
- a lightweight always-available runtime
- clean communication between local logic and simple UI

`Go + Wails` fits this better than a browser-only app and with less overhead than Electron.

## 3. Why Not Electron

Electron is a strong option for many desktop apps, but it is not the best fit for this MVP because:

- memory footprint is larger
- startup cost is heavier
- the product does not need a full browser engine mindset
- your priority is utility and efficiency rather than rich cross-platform web app complexity

Electron would still work, but it goes against the stated goal of minimal resource consumption.

## 4. Why Not Docsify

`Docsify` is useful for rendering Markdown as a documentation site, but this product needs:

- notification scheduling
- persistent local progress tracking
- answer submission
- adaptive review logic
- app window behavior

Those needs push it beyond what a docs viewer is good at.

## 5. Why Not Browser-Only

A browser-only app sounds lightweight, but in practice it is weaker for this use case:

- background scheduling is less reliable
- OS notification integration is weaker or more awkward
- opening a study window from a notification is harder to control
- local file access is more limited

If the app must feel like a native Windows utility, browser-only is the wrong center of gravity.

## 6. Why Wails

`Wails` gives a lightweight desktop shell where:

- `Go` handles backend logic
- a thin web frontend handles UI
- the app can still be packaged like a desktop program

Practical benefits:

- smaller footprint than Electron
- good fit for a utility-style app
- clear separation between logic and UI
- easy to keep frontend minimal

## 7. Why Vue 3

You said the frontend should be as light as possible.

`Vue 3` is a good match because:

- simpler mental model for small UI
- easy to build focused screens without a large architecture
- lighter development overhead than many React-based setups
- good enough ecosystem without forcing complexity

For this app, the frontend only needs a few screens:

- setup/import
- dashboard
- study card
- review session
- settings

That is a strong use case for Vue.

If we wanted even thinner UI later, `Vanilla JS` or `Preact` could also work, but `Vue 3` is the most balanced recommendation.

## 8. Why JSON + JSONL First

This MVP does not need heavy relational querying on day one.

The data naturally splits into:

- content metadata
- current progress snapshot
- append-only study history
- settings

That maps well to `JSON` and `JSONL`.

Recommended local files:

- `data/settings.json`
- `data/cards-cache.json`
- `data/progress.json`
- `data/attempts.jsonl`

Benefits:

- minimal runtime complexity
- easy inspection and debugging
- append-friendly study logs
- no migration burden for the first prototype

## 9. When To Upgrade To SQLite

Move to SQLite later if one or more become true:

- review selection logic gets much more complex
- history grows large enough that loading JSON becomes annoying
- you want richer analytics or filtering
- you need stronger consistency guarantees

For MVP, `JSON + JSONL` is the right simplicity/performance tradeoff.

## 10. Proposed File Responsibilities

### Markdown Files

Source of truth for study content.

Suggested structure:

```text
knowledge/
  git/
    rebase.md
    cherry-pick.md
```

Each card file should contain both Chinese and English learning fields so the UI can switch languages without needing duplicate files.

### `settings.json`

Stores user preferences:

- knowledge directories
- notification interval
- active hours
- review schedule
- review batch size
- preferred language

### `cards-cache.json`

Stores parsed card data derived from Markdown:

- card id
- source path
- parsed frontmatter
- bilingual explanation body
- content hash or modified time

This lets the app avoid reparsing every file on every action.

### `progress.json`

Stores current per-card state:

- seen count
- correct count
- wrong count
- mastery score
- last seen
- next review

### `attempts.jsonl`

Append-only event log, one answer attempt per line.

Example:

```json
{"card_id":"git-rebase-vs-merge","session_type":"learn","shown_at":"2026-04-04T10:00:00+08:00","answered_at":"2026-04-04T10:00:12+08:00","selected_answer":1,"is_correct":true,"response_time_ms":12000}
```

This is useful for:

- later analytics
- rebuilding progress if needed
- debugging scoring issues

## 11. Notification Strategy

Windows notification behavior should use native toast notifications.

Desired flow:

1. Scheduler selects next eligible card
2. App resolves notification style from settings
3. App sends a Windows toast
4. Toast title uses localized `clickbait` text or offline hook-generator output
4. User clicks notification
5. App opens the study card window for that selected card

Notification rules:

- do not notify outside active hours
- do not notify when a review session is already in progress
- optionally avoid duplicate notifications for the same card in a short period

Recommended offline hook generator inputs:

- preferred language
- card tags
- card title and question
- `confusion_with`
- `metaphor_seed`
- `hook_style_tags`
- chosen notification style

Recommended initial notification styles:

- `safe`
- `playful`
- `aggressive`
- `chaotic`

Current implementation note:

- `chaotic` is intentionally allowed to sound more like a headline, shopping hook, or light personality-quiz opener
- the generator remains offline and deterministic so the same card/style pair stays predictable

Recommended title-source modes:

- `prefer_manual`
- `prefer_generated`

## 12. Scheduling Strategy

For MVP, scheduling should be app-managed while the app is running.

Two scheduler categories:

- interval learning scheduler
- fixed-time review scheduler

### Interval Learning Scheduler

- default every 10 minutes
- checks active hours first
- fetches one weighted candidate card
- dispatches a toast

### Fixed-Time Review Scheduler

- daily at configured hour
- or weekly on configured weekday and hour
- creates a review queue of cards due for review

For MVP, it is acceptable if notifications only work while the app is open or minimized to tray.

Future improvement:

- add startup-on-login
- add stronger OS-level trigger support if needed

## 13. Suggested App Lifecycle

At app startup:

1. Load settings
2. Ensure data directory exists
3. Scan or refresh card cache
4. Load progress snapshot
5. Start scheduler
6. Restore user's last selected content language

During runtime:

1. Wait for timer tick
2. Determine whether a learn or review notification should fire
3. Send notification
4. Open study UI on click
5. Persist result immediately after answer

At content refresh:

1. Re-scan knowledge directories
2. Parse changed Markdown files
3. Update cache
4. Preserve progress for unchanged card IDs

## 14. Suggested Folder Structure

```text
duolin-gogo/
  app/
    backend/
      cards/
      progress/
      scheduler/
      notifications/
      settings/
    frontend/
      src/
        components/
        views/
        stores/
  knowledge/
  data/
    settings.json
    cards-cache.json
    progress.json
    attempts.jsonl
  MVP_SPEC.md
  TECH_DECISION.md
```

If using default Wails layout, adjust names to match Wails conventions, but keep the logical separation.

## 15. Suggested Backend Modules

### `cards`

Responsibilities:

- scan Markdown directories
- parse frontmatter
- validate required fields
- build in-memory card models
- update cache

### `progress`

Responsibilities:

- load/save progress snapshot
- load/save attempts log
- calculate mastery changes
- return daily stats

### `scheduler`

Responsibilities:

- maintain timers
- respect active hours
- choose next card
- trigger review sessions

### `notifications`

Responsibilities:

- send Windows toast notifications
- attach routing payload for click handling
- open correct card when toast is clicked

### `settings`

Responsibilities:

- load and validate settings
- expose update methods for UI

## 16. Suggested Frontend Screens

### Setup View

- choose knowledge folder
- run scan/import
- configure notification interval
- configure review schedule

### Dashboard View

- studied today
- correct rate today
- next review time
- weak concepts

### Study Card View

- title
- global language toggle
- staged `Learn -> Answer -> Feedback` flow
- left-heavy study area with explicit next-step controls
- human-readable local time formatting instead of raw ISO strings

### Review View

- review queue progress
- correctness feedback
- summary after batch completion

### Settings View

- active hours
- review cadence
- import/rescan

## 17. UI Direction For The Next Iteration

The next UI pass should move away from a stacked all-at-once layout and toward a clearer desktop learning workspace.

Recommended direction:

- two-column layout
- left column reserved for the active study flow
- right column reserved for answer-related stats and weak topics
- a compact settings trigger should live in the top-right chrome next to the global language toggle
- utility controls and diagnostics should live in a settings popout instead of the main sidebar
- scheduling controls such as review time, notification interval, and active hours should also live in that settings popout
- that settings popout should favor a wider horizontal layout so tool actions and scheduling controls can fit without forcing routine vertical scrolling on desktop
- current MVP verification shows interval notifications working after widening active hours from the settings popout, confirming that schedule changes can be tested from the UI instead of editing `data/settings.json`

Recommended study-state model:

- `Learn`: explanation visible, question hidden
- `Answer`: question visible, explanation hidden
- `Feedback`: correctness and hint visible, plus a clear `Next card` action

Recommended i18n rule:

- `duolin-gogo` remains fixed
- all other shell copy should follow a global UI language setting
- card content and notification language should default to that same global setting

Recommended density rule:

- large screens should feel denser and more tool-like, not like a stretched landing page
- reduce hero weight and oversized spacing
- keep the main study card visually dominant while making the sidebar lighter
- keep top summary copy visually stable when switching languages to avoid obvious layout jumps

Recommended time-format rule:

- use local human-readable timestamps such as `2026-04-05 21:00`
- do not expose raw ISO date strings in visible UI copy
- highlight `next review` in one location instead of repeating the same value across multiple sidebar cards
## 18. Card Selection Logic For MVP

Use a simple priority model.

Inputs:

- whether card is new
- current mastery score
- accuracy rate
- time since last seen
- whether review is due
- whether card was recently wrong

Example weighting idea:

```text
priority =
  30 if new
  + 25 if review overdue
  + 20 if recently wrong
  + min(days_since_seen * 5, 20)
  + weakness bonus
  - mastery penalty
```

Where:

- `weakness bonus` can be based on `wrong_count - correct_count/2`
- `mastery penalty` grows when the card is answered correctly repeatedly

This is intentionally simple and tunable.

## 19. MVP Performance Philosophy

Design principle:

- read files simply
- write state incrementally
- avoid large background work
- keep UI mostly idle unless opened

Practical implications:

- use cached parse results
- append attempts instead of rewriting history
- rewrite `progress.json` only when needed
- keep scheduler lightweight

## 20. Known Tradeoffs

### Tradeoff 1: JSON Simplicity vs Query Power

Using JSON is simpler now, but later analytics may be more awkward than SQL.

### Tradeoff 2: App-Running Requirement

If the scheduler is app-managed, notifications depend on the app being open or tray-resident.

### Tradeoff 3: Thin UI vs Fancy UX

Keeping frontend light means the app should favor clarity over heavy visual polish in early versions.

These are good tradeoffs for MVP.

## 21. Recommended Build Order

### Step 1

Create card parser and cache pipeline.

Deliverable:

- app can scan a folder and list valid cards

### Step 2

Create progress store and attempt logging.

Deliverable:

- app can answer a card and save results

### Step 3

Create minimal study card UI.

Deliverable:

- app can open one selected card and submit an answer

### Step 4

Add interval scheduler and Windows toast notification.

Deliverable:

- notification appears and opens the card view

### Step 5

Add daily/weekly review session generation.

Deliverable:

- app can produce a review batch from prior study data

### Step 6

Add simple dashboard and weak-topic summary.

Deliverable:

- user can see lightweight learning stats

## 22. Final Recommendation

The best-fit technical direction for this product is:

- `Wails` as desktop shell
- `Go` for all local logic
- `Vue 3` for a very small UI layer
- `Markdown` as the editable content source
- bilingual Markdown card content in the same file
- `JSON + JSONL` for MVP persistence
- Windows toast notifications for study prompts

This direction keeps the app:

- lightweight
- practical
- local-first
- easier to reason about than a heavy web-stack desktop app
