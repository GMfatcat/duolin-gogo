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
- valid card cache written to `data/cards-cache.json`

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

- in_progress
- first slice implemented: cards seen in the last 10-30 minutes now receive a repeat-avoidance penalty so alternative cards are preferred when available

### Next Phase C: Multi-Topic Expansion

Goal:

- expand `duolin-gogo` beyond Git while keeping the same local-first flow

Deliverables:

- topic-aware card organization
- optional topic filters in study selection
- dashboard summaries grouped by topic
- support for decks such as `javascript`, `typescript`, or `react`

TDD focus:

- parser tests for multi-topic knowledge trees
- selection tests for topic filtering
- dashboard tests for per-topic summaries

Current status:

- planned
