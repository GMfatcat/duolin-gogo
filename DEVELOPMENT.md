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

Next target:

- simplify the main sidebar so it only shows answer-related stats and weak topics
- move utility controls, hook settings, and import diagnostics into a top-right settings popout
- add review-time, notification-interval, and active-hours controls to the settings popout
- widen the settings popout and lay out controls horizontally enough to avoid routine vertical scrolling on desktop
- summarize import health inline in the settings header, with expandable details only when there are actual issues
- replace raw ISO timestamps in the visible UI with local `YYYY-MM-DD HH:mm` formatting
- keep active-hours scheduling editable from the same settings popout so notification timing can be debugged without hand-editing JSON
- tighten large-screen density so the app feels more like a desktop tool than a landing page
- reduce layout shift when switching hero summary copy between Chinese and English
