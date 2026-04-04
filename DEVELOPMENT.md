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

- content expansion for core Git commands
- offline `hook generator` with TDD-first tests for localized notification styles
