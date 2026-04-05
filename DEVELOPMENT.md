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
  - next authoring slices:
    - A4: batch draft review for reviewing multiple pasted AI drafts in one pass
    - A5: fix suggestions layered on top of diagnostics
    - A6: batch import report for larger AI-assisted save operations
    - A7: Markdown-to-card assist for turning plain notes into draft scaffolds
  - thirteenth slice implemented: AI draft review now supports batch input with `===` separators, per-draft normalized previews, and per-draft diagnostics in one pass
  - fourteenth slice implemented: common authoring diagnostics now include direct fix suggestions, so preview, batch review, and diagnostics panels tell authors what to change instead of only reporting the problem
  - next authoring utility slice: surface the reusable AI card prompt inside the Library popout with a direct copy action
  - fifteenth slice implemented: the reusable AI authoring prompt now lives inside the Library surface with a one-click copy action, so users can generate new cards without leaving the app
  - sixteenth slice implemented: saving reviewed batch drafts now produces an import report with saved/skipped counts and per-draft outcomes, so larger AI-assisted authoring passes are auditable
  - seventeenth slice implemented: Markdown-to-card assist now turns plain notes into a card-shaped scaffold and feeds that scaffold into the existing draft review flow inside Library
  - eighteenth slice implemented: authoring UI is now split between `Library` and a dedicated `AI` popout, so authoring preview stays focused on card browsing while AI prompt, draft review, and note-to-card assist live in their own workspace
  - nineteenth slice implemented: authoring preview now includes a plain search box backed by tokenized regex matching across filename, path, card id, localized titles, and topic, so growing knowledge decks no longer depend on long dropdown scrolling
  - twentieth slice implemented: note-to-card assist now lives in `Library`, the `AI` surface now stays focused on prompt plus draft review, and draft review can target both existing topics and sanitized custom topic-folder names for newly created decks
  - twenty-first slice implemented: parsed knowledge cards now persist to `data/cards-cache.gob`, and app startup reuses that gob cache whenever the knowledge-tree fingerprint has not changed
- Phase B: session and progress UX
  - reduce short-term card repetition
  - add clearer session progress and completion cues
  - improve review-batch completion flow
  - make normal learn sessions feel bounded instead of endless
  - first slice implemented: the selection engine now penalizes cards seen in the last 10-30 minutes so alternate cards are preferred when available
  - second slice implemented: the final review card now transitions into a dedicated completion state before returning to the next learn card
  - third slice implemented: active review sessions now show explicit completed/total/remaining progress cues in the main study card
  - fourth slice implemented: review completion now shows a lightweight session summary with answered count, estimated accuracy, and the weakest current topic
  - fifth slice implemented: standard learn mode now enforces a 3-card mini batch, then shows a short rest state until the current notification interval has passed
  - sixth slice implemented: that learn-break state now also gives a small batch summary so normal learning has the same sense of closure as review mode
  - seventh slice implemented: dashboard stats and learn/review wrap-ups now surface a simple study streak for daily momentum
  - eighth slice implemented: when a learn break expires, the next card now re-enters with a visible `new batch ready` cue instead of silently resuming
  - current state: the first session/progress baseline is complete, so future work here can focus on stronger gamification or pet-layer progression instead of core flow fixes
- Phase C: multi-topic expansion
  - prepare the deck model for topics beyond Git
- support topic-aware summaries and selection filters
- first content expansion slice starts with `docker` and `linux`
- second content expansion slice adds foundational `go` and `python` decks
- third slice implemented: a global topic filter now constrains selected cards, review queues, and weak-topic summaries
- fourth slice implemented: the UI now reflects topic mode more clearly, including focused-topic copy and topic-aware weak-topic headings
- fifth slice implemented: the sidebar now includes per-topic progress cards for mixed-mode and focused deck overview
- sixth slice implemented: mixed mode now adds extra weight to weaker topics so broad study can revisit weaker decks more often
- seventh slice implemented: quick topic pin presets now jump between `all`, `backend-tools`, `languages`, and `git`
- eighth slice implemented: review notifications and test-notification feedback now use the current multi-topic scope instead of assuming Git
- ninth slice implemented: `docker`, `linux`, `go`, and `python` now each have a deeper second-wave deck for mixed-mode study
- tenth slice implemented: grouped topic modes now expose the weakest deck insight so focused presets can call out the weakest sub-deck directly
- eleventh slice implemented: Learn-phase knowledge copy now reveals line by line with a lightweight staggered animation and a replay control, so short card intros feel more productized without blocking the move into answering
- twelfth slice implemented: reveal timing now lives in Settings as a fast/normal/slow preference, so users can tune the guided intro cadence without exposing raw animation timings
- thirteenth slice implemented: multi-topic decks now have a third practical wave, adding Docker networking/storage cleanup, Linux file movement and log-tail basics, Go error/package/receiver ideas, and Python class/comprehension/context-manager/import concepts
- fourteenth slice implemented: multi-topic decks now have a fourth practical wave, adding Docker compose inspection and env/port workflows, Linux pipe/redirection/counting/xargs patterns, Go interface/context/error-wrapping edge cases, and Python iterator/init/module-boundary workflows
- fifteenth slice implemented: the content map now expands beyond tools and languages into three new high-value decks for SQL, HTTP, and backend engineering basics, covering query optimization pitfalls, protocol/transport concepts, and testing/async/cache strategy
- sixteenth slice implemented: those SQL, HTTP, and backend decks now have a second wave covering join/null/update pitfalls, headers/content-type/keep-alive/TLS handshake/REST, and practical backend patterns such as stubs, E2E tests, retries, TTL, hot keys, and queues
- next UI refinement queued: move grouped-topic guidance toward the top as an assistant hint, collapse diagnostics by default, and simplify language/topic controls into dropdowns
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
  - the shell eyebrow copy can be playful and separate from the main `duolin-gogo` product title
- Cross-cut mascot and motion work
  - evolve the top assistant hint into a `DG` mascot surface
  - keep motion subtle, contextual, and secondary to the study flow
  - first slices should focus on weak-deck nudges, review-complete encouragement, and compact collapse/expand behavior
  - avoid marquee-style motion or persistent animation that competes with reading
  - first slice implemented: the existing top hint is now a `DG` bubble with stateful copy, a lightweight entrance animation, and a compact collapsed badge mode
  - second slice implemented: `DG` now reacts to learn, answer, correct-feedback, and wrong-feedback states with distinct copy and tone, so the mascot stays useful even outside review flows
- Cross-cut hidden pet growth work
  - evolve `DG` into a lightweight companion with hidden growth instead of visible game stats
  - do not expose `xp`, `level`, or progress bars to the user
  - let growth show up through richer click reactions, more moods, and more familiar tone over time
  - tie hidden progression to real study behavior such as answered cards, mini-batch completion, review completion, and streak continuation
  - keep reaction frequency controlled so the companion supports the study loop instead of becoming a toy layer
  - first slice implemented: `DG` now persists a hidden local bond state, grows from answered cards, and supports click-to-react behavior with a cooldown and stage-based reaction pool
  - next slice plan:
    - split reactions by trigger type instead of one generic stage pool
    - wire stage-aware pet reactions into correct, wrong, break, review-complete, and return states
    - preserve cooldown/probability guards so `DG` still feels compact and low-noise
  - second slice implemented: stage-aware reactions now differentiate direct clicks from correct/wrong answers, learn breaks, review completion, and returning after a pause
  - third slice implemented: those trigger-specific reactions now rotate through multi-line pools at each hidden stage, so DG feels more alive without adding visible progression UI
  - fourth slice implemented: ambient pet reactions now use low-noise cooldown/probability rules, and DG bubbles expose lightweight pose states like `nod`, `think`, `rest`, and `spark` so the UI can show subtle micro-expression changes
  - fifth slice implemented: the DG badge has now been replaced in-app by the first mascot asset pack, with SVG pose assets wired to the existing `pose-*` states and a dedicated collapsed-badge asset for the compact mode
  - sixth slice implemented: the mascot SVGs have already been compressed substantially after initial import, so the app keeps the new character visuals without carrying the original heavier asset footprint
  - seventh slice implemented: pose-specific motion polish now adds clearer one-shot transitions and idle float behavior, so users can actually feel `nod`, `think`, `rest`, `spark`, and `wave` without the mascot becoming visually noisy
  - eighth slice implemented: DG reaction copy is now topic-aware, so focused study scopes like `docker` and grouped scopes like `languages` can sound different from general mixed-mode nudges
  - next visual slice: tighten the shell palette around `DG teal + spark gold`, replacing leftover bright-blue focus accents so the mascot and the app chrome feel like the same system
  - latest visual slice: that palette now extends through the top controls and popout surfaces as well, so settings, library, diagnostics, and authoring panels read as part of the same mascot-led UI system
  - next easter-egg slice: introduce low-frequency hidden surprises without any visible unlock UI
  - latest easter-egg slice: `rapid_click` and `welcome_back` now live inside the hidden pet layer, so DG can occasionally react in playful or familiar ways when users boop too fast or return after a long gap
  - latest refinement: hidden easter eggs now also include topic inside jokes plus rarer celebration lines, and the reaction picker now advances through an internal step counter so repeated interactions feel less like a fixed script
