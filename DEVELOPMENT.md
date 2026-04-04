# Development Notes

## 1. Workflow

This project uses `TDD` by default.

Working loop:

1. write or update tests first
2. run the failing tests
3. implement the smallest change that makes them pass
4. refactor with tests still passing

## 2. App Structure

- `app/` contains the Wails application
- `app/frontend/` contains the Vue frontend
- `knowledge/` contains bilingual Markdown cards, currently starting with Git
- `data/` contains local runtime JSON and JSONL files

## 3. Test Commands

Backend tests:

```bash
cd app
go test ./...
```

Frontend tests:

```bash
cd app/frontend
npm test
```

Frontend production build check:

```bash
cd app/frontend
npm run build
```

## 4. Current Milestone Focus

Current completed milestone:

- Milestone 9: Polish And MVP Hardening
- Post-MVP utility controls: manual `Rescan knowledge`

Next targets:

- Phase A: content authoring and validation
  - add more Git cards
  - latest deck expansion added `git clone`, `git init`, `git branch`, `git log`, `git reset`, and `git diff`
  - second-wave deck expansion added `git revert`, `git tag`, `git remote`, `git merge conflict`, `git rebase --continue`, and `git push -u`
  - add a reusable LLM authoring prompt so AI-generated cards follow our Markdown schema and bilingual rules
  - next authoring UX slices:
    - A1: card preview
    - A2: AI draft review flow
    - A3: batch validation report
  - improve authoring-time validation for bilingual card files
  - make import feedback clearer for broken Markdown cards
  - first slice implemented: missing localized fields and bilingual choice mismatches now surface as warnings without blocking import
  - second slice implemented: settings diagnostics visually distinguish warnings from errors
  - third slice implemented: `Validate knowledge` checks authoring quality without resetting the current study card
  - fourth slice implemented: diagnostics group warnings and errors separately for faster scan-and-fix authoring flow
  - fifth slice implemented: `Authoring preview` inside settings can load one local card, switch files, and show per-card diagnostics
  - sixth slice implemented: `AI draft review` can validate pasted Markdown and show a normalized preview before saving
  - seventh slice implemented: reviewed drafts can now be saved straight into `knowledge/<topic>/<id>.md`
  - eighth slice implemented: saving a reviewed draft now auto-rescans knowledge and focuses authoring preview on the new card
  - ninth slice implemented: settings diagnostics now include a batch report for total, clean, warning, and error card counts
  - tenth slice implemented: batch report diagnostics can now be filtered by severity and topic
- Phase B: session and progress UX
  - reduce short-term card repetition
  - add clearer session progress and completion cues
  - improve review-batch completion flow
- Phase C: multi-topic expansion
  - prepare the deck model for topics beyond Git
  - support topic-aware summaries and selection filters
