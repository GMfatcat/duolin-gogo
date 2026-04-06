# Project Roadmap

## 1. Purpose

This document tracks the implementation milestones for the project.

It is intentionally separate from:

- [MVP_SPEC.md](D:\duolin-gogo\MVP_SPEC.md), which defines product requirements
- [TECH_DECISION.md](D:\duolin-gogo\TECH_DECISION.md), which records architecture choices
- [SCHEMA.md](D:\duolin-gogo\SCHEMA.md), which defines content and storage formats

This file is the execution-facing plan.

## 2. Development Mode

We will use `TDD` as the default development mode.

Working rule:

1. write or update tests first
2. implement the smallest code change that makes the test pass
3. refactor only after tests are passing

TDD expectations for this project:

- parser behavior should be defined by fixture-based tests
- scheduling logic should be covered by deterministic unit tests
- weighting logic should be tested with explicit input/output cases
- persistence logic should be tested against sample JSON and JSONL files
- UI logic can use component-level tests for key state transitions

## 3. Milestone Overview

### Milestone 0: Project Bootstrap

Goal:

- prepare the repo for implementation

Deliverables:

- initial docs in place
- sample `knowledge/` cards
- sample `data/` files
- git repo initialized
- roadmap and TDD workflow documented

Current status:

- completed

### Milestone 1: Test Harness And Project Skeleton

Goal:

- create the base app structure and test setup before feature work

Deliverables:

- Go module initialized
- Wails app scaffolded
- frontend scaffolded with Vue 3
- backend package layout created
- test runner commands documented
- first passing smoke tests

TDD focus:

- prove the test environment works before writing real feature code

Current status:

- completed

### Milestone 2: Markdown Card Parser

Goal:

- load bilingual Git knowledge cards from `knowledge/`

Deliverables:

- recursive Markdown scanner
- YAML frontmatter parser
- bilingual body section parser for `zh-TW` and `en`
- schema validation
- import error reporting to `data/import-errors.json`
- valid card cache written to `data/cards-cache.gob`

TDD focus:

- valid card fixture parses successfully
- invalid fixtures fail with expected error codes
- bilingual section extraction is verified by tests
- duplicate ID behavior is covered by tests

Current status:

- completed

### Milestone 3: Progress Store And Attempt Logging

Goal:

- persist study progress and answer history locally

Deliverables:

- `progress.json` load/save logic
- `attempts.jsonl` append logic
- daily summary updater
- stable progress updates keyed by card ID

TDD focus:

- progress update logic for correct and wrong answers
- JSON persistence round-trip tests
- JSONL append tests
- daily summary update tests

Current status:

- completed

### Milestone 4: Card Selection Engine

Goal:

- choose the next card based on simple adaptive weighting

Deliverables:

- new-card boost
- review-due boost
- wrong-answer boost
- mastery penalty
- deterministic selection behavior from weighted inputs

TDD focus:

- priority score tests for each factor
- regression tests for edge cases
- candidate ordering tests with fixed timestamps

Current status:

- completed

### Milestone 5: Minimal Study Card UI

Goal:

- let the user open one card, switch language, read, answer, and see feedback

Deliverables:

- dashboard shell
- study card view
- language toggle between `zh-TW` and `en`
- answer submission flow
- correctness feedback state

TDD focus:

- component tests for language toggle behavior
- answer submit and feedback tests
- UI state tests for loading one card

Current status:

- completed

### Milestone 6: Notification And App Scheduling

Goal:

- deliver timed learning prompts as Windows toast notifications

Deliverables:

- interval scheduler
- active-hours guard
- Windows toast notification integration
- notification click opens the study card
- cooldown and snooze support

TDD focus:

- scheduler decision tests with mocked time
- active-hours tests
- notification payload formatting tests
- click-routing behavior tests where practical

Current status:

- completed

### Milestone 7: Review Session Mode

Goal:

- run daily or weekly review batches from prior study data

Deliverables:

- review schedule trigger logic
- review queue builder
- review session UI flow
- review results persisted to progress and attempts

TDD focus:

- due-card selection tests
- review queue composition tests
- review result persistence tests

Current status:

- completed

### Milestone 8: Basic Dashboard And Weak Topic Summary

Goal:

- make learning progress visible without building a heavy analytics system

Deliverables:

- studied today count
- correct rate today
- next review time
- weak Git topic summary

TDD focus:

- dashboard aggregation tests from sample progress data
- weak-topic grouping tests by tag

Current status:

- completed

### Milestone 9: Polish And MVP Hardening

Goal:

- make the MVP stable enough for daily personal use

Deliverables:

- import diagnostics view
- better empty states
- better notification fallback copy
- recovery behavior for malformed files
- packaging checklist for Windows

TDD focus:

- regression tests for known parser and scheduler failures
- persistence safety tests

Current status:

- completed

## 4. Recommended Build Order

Recommended implementation order:

1. Milestone 1: Test Harness And Project Skeleton
2. Milestone 2: Markdown Card Parser
3. Milestone 3: Progress Store And Attempt Logging
4. Milestone 4: Card Selection Engine
5. Milestone 5: Minimal Study Card UI
6. Milestone 6: Notification And App Scheduling
7. Milestone 7: Review Session Mode
8. Milestone 8: Basic Dashboard And Weak Topic Summary
9. Milestone 9: Polish And MVP Hardening

## 5. File Placement Decision

Milestones should live in this dedicated file:

- [ROADMAP.md](D:\duolin-gogo\ROADMAP.md)

Why this file is the right place:

- milestone planning changes more often than product spec
- it keeps execution tracking separate from architecture notes
- it gives us one place to update status as implementation moves forward

## 6. Status Tracking Convention

Suggested status labels for future updates:

- `planned`
- `in_progress`
- `blocked`
- `completed`

For now:

- Milestone 0 is `completed`
- Milestone 1 is `completed`
- Milestone 2 is `completed`
- Milestone 3 is `completed`
- Milestone 4 is `completed`
- Milestone 5 is `completed`
- Milestone 6 is `completed`
- Milestone 7 is `completed`
- Milestone 8 is `completed`
- Milestone 9 is `completed`

## 7. Post-MVP Follow-Ups

### Utility Controls

Goal:

- make the desktop app easier to operate without restarting it during local study

Deliverables:

- manual `Rescan knowledge` control
- refreshed dashboard state after rescan
- runbook updates for the new control

TDD focus:

- backend test for manual rescan updating cached knowledge
- frontend test for rescan action feedback

Current status:

- completed

### Content Expansion And Hook Generation

Goal:

- grow the Git card set enough for real spaced repetition
- make notifications more clickable through offline hook templates

Deliverables:

- new Git cards for `add`, `commit`, `merge`, `fetch`, `pull`, `checkout`
- second-wave Git cards for `status`, `switch`, `restore`, `stash`
- schema support for hook metadata
- offline `hook generator` module with `safe`, `playful`, `aggressive`, and `chaotic` modes
- notification flow wired to localized generated hooks

TDD focus:

- parser tests for hook metadata fields
- deterministic hook-generator tests by style and language
- notification tests for localized hook selection

Current status:

- completed

### Study Flow And Global i18n Refresh

Goal:

- make the desktop learning flow feel more guided, more bilingual, and less cluttered

Deliverables:

- two-column layout with study content on the left and progress/context on the right
- global `zh-TW` / `en` shell-language toggle
- shell copy localized consistently except for `duolin-gogo`
- staged study flow with `Learn`, `Answer`, and `Feedback` states
- explicit next-step action after answer submission
- settings icon plus popout for utility actions, hook controls, and import diagnostics
- settings popout extended with review-time, notification-interval, and active-hours controls
- settings popout widened so the core controls fit horizontally on a normal desktop screen without needing routine vertical scrolling
- human-readable local date/time formatting in the UI
- tighter density on large screens
- deduplicated sidebar summaries with one emphasized `next review` display
- more stable hero copy layout across language switches

TDD focus:

- frontend tests for global language switching across shell copy
- component tests for staged visibility of explanation, question, and feedback
- interaction tests for answer submission followed by `Next card`
- component tests for settings popout visibility and sidebar simplification
- formatting tests for visible date/time strings
- interaction tests for updating review time, notification interval, and active hours from settings

Current status:

- completed

### Next Phase A: Content Authoring And Validation

Goal:

- make it much easier to grow the knowledge deck without breaking schema or restarting the app

Deliverables:

- more Git cards beyond the current baseline deck
- current expansion set now includes `clone`, `init`, `branch`, `log`, `reset`, and `diff`
- second-wave Git deck now also includes `revert`, `tag`, `remote`, `merge conflict`, `rebase --continue`, and `push -u`
- authoring-oriented validation for `knowledge/` files
- preview-friendly card diagnostics for missing bilingual fields
- clearer import feedback for malformed card content
- optional card-preview or validate action for local authoring workflows
- a reusable LLM prompt for generating schema-compliant bilingual cards

TDD focus:

- fixture tests for incomplete bilingual card files
- validation tests for missing hook metadata and choice mismatches
- refresh tests that preserve healthy cards while reporting bad ones

Current status:

- in_progress
- first slice implemented: the reusable prompt from `AI_CARD_PROMPT.md` now loads inside the Library surface and can be copied directly from the app
- latest slice implemented: the authoring workspace is now split into `Library` and `AI`, and authoring preview has a built-in search flow so larger decks can be browsed without relying on one long file dropdown
- latest refinement: note-to-card assist moved back into `Library`, the `AI` workspace now centers on prompt plus draft review, and draft saves can target sanitized custom topic folders for newly created decks
- latest infrastructure slice: parsed knowledge cards now persist to `data/cards-cache.gob`, and startup can reuse that gob cache whenever the knowledge fingerprint is unchanged

### Next Phase A1: Card Preview

Goal:

- let authors inspect one card before it enters the normal study flow

Deliverables:

- load a single local card file for author preview
- show parsed result inside the app settings flow before the card enters normal study
- allow card switching from a local Markdown file list
- support `zh-TW` / `en` preview switching
- show diagnostics for that card without relying on the full dashboard
- keep preview refreshable after `Validate knowledge` and `Rescan knowledge`

TDD focus:

- parser-to-preview data mapping tests
- component tests for language switching in preview
- diagnostics rendering tests for single-card preview

Current status:

- completed

### Next Phase A2: AI Draft Review Flow

Goal:

- make AI-generated cards reviewable before they are trusted as knowledge cards

Deliverables:

- draft input area for pasted AI Markdown
- schema validation result for the draft
- normalized preview of parsed card data
- first slice implemented inside settings popout for pasted Markdown review
- save flow for accepted drafts into a topic folder under `knowledge/`
- saving a reviewed draft now auto-rescans the deck and focuses authoring preview on the new file
- support for the reusable prompt and review prompt from `AI_CARD_PROMPT.md`

TDD focus:

- validation tests for pasted draft content
- preview tests for accepted vs rejected drafts
- save-flow tests for approved drafts

Current status:

- in_progress

### Next Phase A3: Batch Validation Report

Goal:

- make larger AI-assisted deck growth easy to audit

Deliverables:

- whole-deck validation summary
- first slice implemented inside settings with total, clean, warning, and error card counts
- grouped warning/error report
- second slice implemented with severity and topic filters
- third slice implemented with a recently changed cards summary based on authoring preview timestamps
- quick visibility into recently changed or problematic cards

TDD focus:

- aggregation tests for warning/error totals
- filtering tests by severity and topic
- UI tests for grouped diagnostics report
- preview metadata tests for recent-change timestamps

Current status:

- completed

### Next Phase A4: Batch Draft Review

Goal:

- let authors review multiple AI-generated Markdown drafts in one pass before saving anything into `knowledge/`

Deliverables:

- accept multiple pasted Markdown drafts in a single review surface
- split the raw input into draft-sized review units
- validate each draft independently
- show per-draft pass, warning, or error state
- show normalized preview per draft instead of only one global preview
- keep save actions scoped to reviewed items rather than forcing an all-or-nothing flow

TDD focus:

- backend tests for mixed valid/invalid draft batches
- frontend tests for rendering multiple review results from one paste
- per-item diagnostics tests for valid, warning, and error cases

Current status:

- in_progress
- first slice implemented: pasted AI Markdown can now be split into multiple drafts with `===`, reviewed per item, and rendered as separate normalized preview cards with their own diagnostics

### Next Phase A5: Fix Suggestions

Goal:

- turn diagnostics into direct authoring guidance instead of only error reporting

Deliverables:

- suggest likely fixes for common draft and card issues
- explain missing bilingual sections and field mismatches in plain language
- surface the most important fix first for each broken draft
- keep suggestions visible beside diagnostics in preview and batch review flows

TDD focus:

- tests for mapping parser codes to human-readable fix suggestions
- UI tests for suggestion rendering alongside diagnostics

Current status:

- in_progress
- first slice implemented: common diagnostics now carry direct fix suggestions, and authoring surfaces render those suggestions beside the original error or warning message

### Next Phase A6: Batch Import Report

Goal:

- make larger AI-assisted content imports easier to audit and trust

Deliverables:

- summarize imported, skipped, warned, and errored cards for a batch operation
- show per-topic import breakdown
- make recently added or changed cards easy to identify after a batch save

TDD focus:

- aggregation tests for batch import summaries
- UI tests for grouped report rendering after multi-card import operations

Current status:

- in_progress
- first slice implemented: batch draft saves now return an import report with saved, skipped, warning, and error counts, plus per-draft outcome details

### Next Phase A5.5: In-App AI Prompt Viewer

Goal:

- make the authoring prompt available inside the app so users can generate cards without leaving the authoring workflow

Deliverables:

- load the reusable prompt from `AI_CARD_PROMPT.md`
- show a prompt viewer inside the Library surface
- provide a one-click copy action for the prompt text

TDD focus:

- backend test for loading prompt content from disk
- frontend test for rendering the prompt panel and copy action feedback

Current status:

- in_progress

### Next Phase A7: Markdown-To-Card Assist

Goal:

- help authors convert plain technical notes into card-shaped Markdown drafts faster

Deliverables:

- accept general-purpose Markdown notes as source input
- scaffold a draft card structure from that input
- preserve bilingual and schema expectations in the generated scaffold
- keep the result reviewable through the same draft-review flow

TDD focus:

- conversion tests for scaffold generation
- validation tests for generated draft skeletons

Current status:

- in_progress
- first slice implemented: plain Markdown notes can now be scaffolded into a card-shaped draft, then passed directly into the existing draft-review flow inside Library

### Next Phase B: Session And Progress UX

Goal:

- make day-to-day study sessions feel more complete and less repetitive

Deliverables:

- clearer per-session progress cues
- better transition after finishing a review batch
- reduced short-term repetition of the same card
- optional streak or completion messaging
- lightweight end-of-session summary

TDD focus:

- selection tests for short-term repeat avoidance
- UI tests for review completion and next-step transitions
- aggregation tests for session summary data

Current status:

- completed
- first slice implemented: cards seen in the last 10-30 minutes now receive a repeat-avoidance penalty so alternative cards are preferred when available
- second slice implemented: finishing the last review card now lands in a dedicated review-complete state with a clear next-step action
- third slice implemented: review batches now show explicit progress cues for completed, total, and remaining cards
- fourth slice implemented: review completion now includes a lightweight end-of-session summary with answered count, accuracy, and the current weakest topic
- fifth slice implemented: normal learn mode now pauses after each 3-card mini batch and waits until the next notification interval before unlocking another card
- sixth slice implemented: the learn-break state now includes a lightweight session summary with answered count, accuracy, and the current topic to revisit
- seventh slice implemented: dashboard stats and session wrap-ups now expose a simple study streak so day-to-day momentum is visible
- eighth slice implemented: when a learn break ends, the next card now starts with a lightweight `new batch ready` cue so the restart boundary is visible

### Next Phase C: Multi-Topic Expansion

Goal:

- expand `duolin-gogo` beyond Git while keeping the same local-first flow

Deliverables:

- topic-aware card organization
- optional topic filters in study selection
- dashboard summaries grouped by topic
- support for decks such as `docker`, `linux`, `go`, and `python`
- first slice implemented: a global topic filter now drives next-card selection, review queue composition, and filtered dashboard summaries
- second slice implemented: the UI now makes mixed-mode versus focused-topic study clearer with topic-aware copy and weak-topic framing
- third slice implemented: the sidebar now shows per-topic progress so mixed-mode study has a visible deck-level overview
- fourth slice implemented: mixed mode now boosts weaker topics so `all` mode is less likely to over-serve already-strong decks
- fifth slice implemented: topic pin presets now provide one-click filters for `all`, `backend-tools`, `languages`, and `git`
- sixth slice implemented: review notifications and test-notification feedback now respect multi-topic scopes such as `backend-tools` and `languages`
- seventh slice implemented: each non-Git deck now has a second wave of cards, giving `docker`, `linux`, `go`, and `python` more study depth
- eighth slice implemented: grouped topic modes now surface the weakest deck inside the group, so presets like `languages` can directly point back to `go` or `python`
- ninth slice implemented: Learn-phase concept copy now reveals line by line with a lightweight staggered animation and replay control, so short knowledge notes feel more guided without slowing the whole card flow
- tenth slice implemented: reveal timing is now user-configurable through Settings with simple fast/normal/slow options, so guided intros can match different reading preferences
- eleventh slice implemented: a third wave of practical deck cards now covers Docker networking/storage cleanup, Linux file inspection and movement, Go error/package/receiver concepts, and Python classes/comprehensions/context managers/imports
- twelfth slice implemented: a fourth wave of practical deck cards now covers Docker compose inspection and env/port workflows, Linux pipes/redirection/counting/xargs, Go interface/context/error-wrapping edge cases, and Python iterator/init/module-boundary workflows
- thirteenth slice implemented: three new high-value decks now broaden the app into SQL, HTTP, and backend engineering basics, covering query logic and optimization traps, transport/protocol fundamentals, and testing/async/cache design concepts
- fourteenth slice implemented: the second wave for SQL, HTTP, and backend decks now adds optimization pitfalls, protocol metadata and connection behavior, plus practical retry/TTL/queue/test-double concepts
- fifteenth slice implemented: practical expansion now also covers Bash command flow, middleware pipeline concepts, backend CRUD semantics, and SQL transaction ACID basics
- sixteenth slice implemented: a concept-oriented wave now adds OOP, functional programming, and lightweight design-patterns decks that stay close to day-to-day engineering instead of drifting into overly abstract theory
- next UI refinement: promote grouped-topic guidance into a top assistant-style hint, collapse diagnostics by default, and replace parallel language/mode controls with dropdown selectors
- next interaction refinement: evolve the top assistant hint into a lightweight `DG` mascot surface with contextual copy and subtle animation states

TDD focus:

- parser tests for multi-topic knowledge trees
- selection tests for topic filtering
- dashboard tests for per-topic summaries

Current status:

- in_progress

### Cross-Cut: DG Mascot And Motion

Goal:

- turn the top assistant hint into a recognizable mascot surface without adding noisy motion

Deliverables:

- compact `DG` mascot bubble near the top of the shell
- context-aware hint copy for weak-deck nudges, review completion, and general encouragement
- collapse/expand interaction for the mascot text
- subtle animation states for entrance, idle emphasis, and completion feedback
- motion guidelines that keep the study card as the visual priority

TDD focus:

- component tests for mascot hint collapse and restore behavior
- UI tests for context-aware copy switching by topic mode and review state
- style-regression checks for collapsed versus expanded mascot layout

Current status:

- in_progress
- first slice implemented: `DG` now has hidden local growth state plus click-to-react behavior, without exposing visible pet stats, levels, or progress bars
- first slice implemented: the top-shell assistant hint now behaves as a compact `DG` mascot bubble with weak-deck and review-complete states, plus subtle entrance and collapse motion
- second slice implemented: `DG` now changes tone across learn, answer, correct-feedback, and wrong-feedback states so users can feel the study flow without waiting for review mode

### Cross-Cut: DG Pet Growth

Goal:

- evolve `DG` from a contextual hint bubble into a lightweight companion that feels more alive over time without turning the app into a visible stats game

Deliverables:

- hidden growth state for `DG`, such as internal bond or stage progression that is not exposed as `xp`, `level`, or public meters
- click interaction on `DG` so users can poke or greet the mascot directly
- unlock-based reaction growth, where later usage gradually reveals more lines, moods, and micro-reactions
- growth driven by real study behavior such as answered cards, completed mini batches, completed review batches, and streak continuation
- low-noise behavior rules so `DG` does not react on every single event
- next planned slices:
  - expand the reaction pool by trigger type: `clicked`, `correct`, `wrong`, `learn-break`, `review-complete`, `return`
  - make later hidden stages unlock richer variants inside each trigger type instead of only swapping one generic line
  - add cooldown and probability rules so contextual pet reactions stay occasional and pleasant
  - keep the companion layer invisible as a system, with no public level meters or growth charts

TDD focus:

- state-transition tests for hidden growth thresholds
- interaction tests for click-to-react behavior
- UI tests ensuring unlocked reactions change visible copy without adding public progress meters
- cooldown tests so repeated clicking does not spam reactions or growth

Current status:

- planned
- first slice implemented: `DG` now stores hidden local bond state and supports click-to-react with stage-based unlocks
- second slice implemented: `DG` now has trigger-specific reaction pools for `correct`, `wrong`, `learn-break`, `review-complete`, and `return`, so the companion can react across the study loop instead of only on direct clicks
- third slice implemented: each trigger-stage pair now rotates through multiple short reaction variants, so later hidden stages feel richer without exposing visible pet stats
- fourth slice implemented: low-noise rules now let some ambient `correct` and `wrong` cues stay quiet, while always-on moments like breaks, returns, and review completion still react; DG also now exposes lightweight pose states for micro-expression styling
- fifth slice implemented: Mascot V1 is now integrated into the top DG bubble using pose-mapped SVG assets for `idle`, `wave`, `nod`, `think`, `rest`, `spark`, plus a dedicated collapsed badge asset
- sixth slice implemented: the first mascot asset pack has already been SVG-compressed for lighter frontend delivery while keeping the pose mapping intact
- seventh slice implemented: pose-specific motion polish now makes `idle`, `wave`, `nod`, `think`, `rest`, and `spark` transitions easier to notice without turning DG into a noisy animated element
- eighth slice implemented: `DG` persona copy is now topic-aware, so focused scopes such as `docker` and grouped scopes such as `languages` can surface different companion tone instead of reusing the same generic lines
- ninth slice planned: pull the app chrome and mascot states into a shared `DG teal + spark gold` color system so the UI feels like one visual world instead of splitting between gold and unrelated bright blue accents
- ninth slice implemented: the shell, top controls, and popout surfaces now share the same `DG teal + spark gold` palette, so mascot chrome, settings, library, and diagnostics no longer feel like separate themes
- tenth slice implemented: Easter Egg V1 has begun with hidden `rapid_click` and `welcome_back` reactions, keeping rare companion surprises inside the existing DG bubble without adding visible unlock UI
- eleventh slice implemented: Easter Egg V1 now also includes topic inside jokes and rare celebration lines, and DG reaction rotation is less script-like because repeated triggers no longer lock into the same exact sentence order
- twelfth slice implemented: DG now has a thicker everyday reaction pool plus low-frequency time-of-day flavor, while topic-aware click tones stay primary so normal interactions feel less scripted without drowning in easter eggs
- thirteenth slice implemented: DG can now notice short topic streaks and occasionally surface "almost there" encouragement when a weaker deck starts improving, so mascot surprises feel tied to real study momentum instead of pure randomness
- DG Pet V2 planned slices:
  - `V2-A Stage Visuals`
    - make hidden stage 0 / 1 / 2 feel visible through subtle mascot differences instead of explicit levels
    - likely cues: brighter spark nub, richer gloss, stronger aura, or slightly more expressive pose treatment
    - first visual pass should focus on high-visibility poses such as `idle`, `wave`, and `spark`
  - `V2-B Context Memory`
    - let DG feel more aware of the user's recent study rhythm, not just the immediate trigger
    - reinforce welcome-back, topic streak, and weak-deck recovery with more familiar copy
  - `V2-C Interaction Growth`
    - deepen direct click interaction so later hidden stages feel more alive without becoming a toy layer
    - later stages should unlock thicker click pools and slightly more attitude-rich reactions
  - `V2-D Rare Surprises`
    - reserve low-frequency visual or behavioral surprises for later polish, such as rare pose variants or tiny badge-level changes
- V2-A wiring slice implemented: frontend mascot asset resolution now has a stable fallback chain for future `stage0/1/2` art, so `idle`, `wave`, and `spark` can gain staged assets incrementally without breaking current V1 poses

### Cross-Cut: Background Running And Tray Lifecycle

Goal:

- make `duolin-gogo` behave like a background helper instead of a close-to-exit window

Deliverables:

- clicking `X` hides the main window instead of quitting
- notifications and review scheduling continue while hidden
- Windows tray icon with `Open duolin-gogo` and `Exit`
- explicit full shutdown only through the tray `Exit` action
- first implementation may defer minimize-to-tray until close-to-tray is stable

TDD focus:

- app-level tests for close interception policy
- lifecycle tests for explicit quit vs hide behavior
- manual verification checklist for tray restore and explicit exit

Current status:

- in_progress
- first slice implemented: clicking `X` now routes through close interception and hides the window instead of quitting
- second slice implemented: a minimal Windows tray surface now provides `Open duolin-gogo` and `Exit`
- manual verification passed for tray restore and explicit exit on the built app

### Cross-Cut: Settings And Authoring Surface Split

Goal:

- keep runtime settings lightweight while moving card-authoring tools into a separate workspace

Deliverables:

- keep gear/settings popout focused on runtime settings and operational actions
- add a separate library or authoring button in the top-right chrome
- move authoring preview and AI draft review out of settings
- move detailed diagnostics and batch validation into the authoring/library surface
- keep only a concise knowledge-health summary inside settings

TDD focus:

- component tests for separate popout toggles
- UI tests ensuring authoring preview is no longer rendered inside settings
- UI tests for detailed diagnostics living in the authoring surface

Current status:

- in_progress
- first slice implemented: settings and authoring are already split into separate gear and book surfaces
- second slice implemented: diagnostics now live in a dedicated top-right popout so settings and library stay shorter
