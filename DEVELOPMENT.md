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
  - eleventh slice implemented: batch report now includes a recently changed cards summary using authoring preview file timestamps
  - twelfth slice implemented: settings tools now include a guarded `Reset study data` action that clears local progress and attempts only after an explicit confirmation step
- Phase B: session and progress UX
  - reduce short-term card repetition
  - add clearer session progress and completion cues
  - improve review-batch completion flow
  - first slice implemented: the selection engine now penalizes cards seen in the last 10-30 minutes so alternate cards are preferred when available
  - second slice implemented: the final review card now transitions into a dedicated completion state before returning to the next learn card
  - third slice implemented: active review sessions now show explicit completed/total/remaining progress cues in the main study card
  - fourth slice implemented: review completion now shows a lightweight session summary with answered count, estimated accuracy, and the weakest current topic
- Phase C: multi-topic expansion
  - prepare the deck model for topics beyond Git
  - support topic-aware summaries and selection filters
  - first content expansion slice starts with `docker` and `linux`
  - second content expansion slice adds foundational `go` and `python` decks
  - third slice implemented: a global topic filter now constrains selected cards, review queues, and weak-topic summaries
  - fourth slice implemented: the UI now reflects topic mode more clearly, including focused-topic copy and topic-aware weak-topic headings
- Cross-cut lifecycle work
  - add close-to-background behavior so `X` hides instead of quitting
  - keep notifications active while the app window is hidden
  - add a minimal Windows tray surface with `Open duolin-gogo` and `Exit`
  - treat minimize-to-tray as a follow-up after close-to-tray is stable
  - first slice implemented: explicit close now hides the window through `OnBeforeClose` instead of terminating the app
  - second slice implemented: a minimal tray icon now exposes `Open duolin-gogo` and `Exit`
- Cross-cut UI structure work
  - split runtime settings from authoring tools
  - keep `Settings` focused on schedule, notification, and utility controls
  - move `Authoring preview` and `AI draft review` into a separate library/authoring popout
  - move detailed diagnostics and batch validation into a dedicated diagnostics popout
  - implemented: the top-right chrome now uses separate gear and book buttons so authoring tools no longer crowd the main settings surface
  - implemented: diagnostics and batch validation now use their own top-right popout so the library surface stays shorter and cleaner
