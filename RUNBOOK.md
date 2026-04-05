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
- a Windows tray icon for `duolin-gogo`
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

- global `Topic` filter for `all`, `git`, `docker`, `linux`, `go`, and `python`
- `Send test notification`
- `Snooze 15 min`
- `Rescan knowledge`
- `Validate knowledge`
- `Reset study data` with a confirmation warning before local progress is cleared
- `Authoring preview` inside the library/book popout for single-card inspection
- `AI draft review` inside the library/book popout for pasted Markdown inspection
- diagnostics now live in their own top-right diagnostics surface so neither settings nor library needs to scroll for deck health details
- `Save draft` from reviewed AI Markdown into `knowledge/<topic>/`
- saving a draft now auto-refreshes authoring preview onto the new card
- settings diagnostics now include a deck-level batch report
- batch diagnostics can now be filtered by severity and topic
- batch diagnostics now include a recently changed cards summary to highlight freshly edited knowledge files
- resetting study data removes local `progress.json` and `attempts.jsonl`, then reloads the dashboard from a clean state
- finishing the last card in a review batch now shows a dedicated completion state before returning to the next learn card
- active review batches now show explicit progress cues for completed, total, and remaining cards
- the review completion state now includes a lightweight session summary with answered count, estimated accuracy, and the weakest current topic
- `zh-TW` / `en` language toggle
- separate top-right `Settings`, `Library`, and `Diagnostics` buttons
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

## 9. Test Background Running And Tray Behavior

Current first-phase tray behavior:

- clicking the window close button should hide the window
- the app should remain alive in the Windows tray
- the tray menu should provide:
  - `Open duolin-gogo`
  - `Exit`

Manual verification path:

1. launch `D:\duolin-gogo\app\build\bin\app.exe`
2. confirm the tray icon appears in the Windows notification area
3. click the window `X`
4. confirm the main window hides instead of quitting
5. right-click the tray icon and click `Open duolin-gogo`
6. confirm the same app window is restored and focused
7. right-click the tray icon and click `Exit`
8. confirm the process fully exits

## 10. Known Current Limitation

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

## 11. Known UI Limitation

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

## 12. Data Files To Inspect During Testing

If the app appears to work but you want to verify state changes, inspect:

- `D:\duolin-gogo\data\cards-cache.json`
- `D:\duolin-gogo\data\progress.json`
- `D:\duolin-gogo\data\attempts.jsonl`
- `D:\duolin-gogo\data\import-errors.json`

What to expect:

- `progress.json` updates after answers
- `attempts.jsonl` appends one JSON line per answer
- `import-errors.json` records malformed knowledge files

## 13. Current Validation Status

The following have been verified:

- `go test ./...` passes
- `npm test` passes
- `npm run build` passes
- `wails build` succeeds
- `app.exe` launches
- tray-aware build succeeds
- Windows toast can be sent successfully
- UI has a manual test notification button
- UI has a snooze button
- UI has a manual knowledge rescan button
- close interception now hides instead of immediately quitting

The following still needs improvement:

- richer review navigation controls
- more explicit diagnostics actions
- optional minimize-to-tray behavior

## 14. Recommended Next Fix

The most valuable next implementation is likely:

- richer review navigation and queue controls
- diagnostics actions and deeper review controls
