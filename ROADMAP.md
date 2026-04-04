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
- Milestone 9 is `planned`
