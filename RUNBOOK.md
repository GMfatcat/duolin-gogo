# Runbook

## 1. Purpose

This document explains how to:

- run the app locally
- build the desktop executable
- test the learning flow
- test Windows notifications
- understand current MVP limitations

## 2. Project Entry Points

Repository root:

- `D:\duolin-gogo`

Desktop app project:

- `D:\duolin-gogo\app`

Knowledge files:

- `D:\duolin-gogo\knowledge`

Runtime data:

- `D:\duolin-gogo\data`

## 3. Development Commands

Run backend tests:

```powershell
cd D:\duolin-gogo\app
go test ./...
```

Run frontend tests:

```powershell
cd D:\duolin-gogo\app\frontend
npm test
```

Run frontend production build check:

```powershell
cd D:\duolin-gogo\app\frontend
npm run build
```

## 4. Run The App In Development

If you want the Wails desktop app in development mode:

```powershell
cd D:\duolin-gogo\app
C:\Users\GMfatcat\go\bin\wails.exe dev
```

What this gives you:

- a live desktop window
- frontend rebuild during development
- easier iteration while editing UI or Go bindings

## 5. Build The Desktop App

Build production executable:

```powershell
cd D:\duolin-gogo\app
C:\Users\GMfatcat\go\bin\wails.exe build
```

Output executable:

- `D:\duolin-gogo\app\build\bin\app.exe`

## 6. Launch The Built App

Start the built executable:

```powershell
Start-Process -FilePath "D:\duolin-gogo\app\build\bin\app.exe"
```

After launch, you should see:

- the main desktop window
- a study card or review card if one is available
- dashboard summary cards
- weak-topic summary
- import diagnostics section

## 7. Test The Learning Flow

Expected current UI behavior:

- one study card is shown
- `zh-TW` / `en` buttons switch explanation language
- answer options can be selected
- `Submit answer` shows correctness feedback
- studied-today / correct-rate values update

Current MVP note:

- the UI is still intentionally lightweight
- but it now includes a few utility controls for manual testing

Current useful controls:

- `Send test notification`
- `Snooze 15 min`
- `Rescan knowledge`
- `Validate knowledge`
- `Authoring preview` inside settings for single-card inspection
- `AI draft review` inside settings for pasted Markdown inspection
- `Save draft` from reviewed AI Markdown into `knowledge/<topic>/`
- saving a draft now auto-refreshes authoring preview onto the new card
- settings diagnostics now include a deck-level batch report
- `zh-TW` / `en` language toggle
- answer submission

## 8. Test Windows Notifications

Notification prerequisites:

- the app should already be running
- Windows notifications must be enabled on the machine
- current time must be inside the configured active hours

Current default settings:

- notification interval: 10 minutes
- active hours: `09:00` to `22:00`
- daily review time: `21:00`

You can now change all of these from the in-app settings popout:

- notification interval
- review time
- active hours enabled / disabled
- active hours start / end

### Automatic notification path

If the app is running, the scheduler checks every minute.

It may send:

- a normal learning notification
- or a review notification if review time is due

### Manual notification test path

The simplest manual path is now:

1. run the app
2. click `Send test notification`
3. wait for the Windows toast
4. click the toast

### Automatic interval test path

If you want to verify interval notifications instead of only the manual test button:

1. open the settings popout
2. set a short interval such as `5`
3. make sure active hours still include the current local time
4. save the schedule
5. leave the app running for 5 to 6 minutes

Expected result:

- a new Windows toast should appear automatically

## 9. Known Current Limitation

You reported:

- the desktop window appears
- Windows toast appears
- clicking the toast does not visibly bring anything forward

This was a real limitation in the earlier MVP pass.

The current implementation now explicitly tries to:

- show the window
- unminimise the window
- briefly raise always-on-top to bring it forward
- emit the card-open event into the frontend

That should improve the notification click behavior significantly.

If clicking the toast still does not visibly foreground the app on this machine, the remaining issue is more likely Windows/WebView2 focus behavior rather than missing app-side wiring.

## 10. Known UI Limitation

You also reported:

- the UI feels like it has very few buttons

The current UI mainly supports:

- `Send test notification`
- `Snooze 15 min`
- `Rescan knowledge`
- language switching
- answer submission
- passive dashboard display

It still does not include controls such as:

- `Open next review`
- `Open diagnostics details`

## 11. Data Files To Inspect During Testing

If the app appears to work but you want to verify state changes, inspect:

- `D:\duolin-gogo\data\cards-cache.json`
- `D:\duolin-gogo\data\progress.json`
- `D:\duolin-gogo\data\attempts.jsonl`
- `D:\duolin-gogo\data\import-errors.json`

What to expect:

- `progress.json` updates after answers
- `attempts.jsonl` appends one JSON line per answer
- `import-errors.json` records malformed knowledge files

## 12. Current Validation Status

The following have been verified:

- `go test ./...` passes
- `npm test` passes
- `npm run build` passes
- `wails build` succeeds
- `app.exe` launches
- Windows toast can be sent successfully
- UI has a manual test notification button
- UI has a snooze button
- UI has a manual knowledge rescan button

The following still needs improvement:

- richer review navigation controls
- more explicit diagnostics actions

## 13. Recommended Next Fix

The most valuable next implementation is likely:

- richer review navigation and queue controls
- diagnostics actions and deeper review controls
