package main

import (
	_ "embed"
	"sync"

	"fyne.io/systray"
)

//go:embed build/appicon.png
var trayIcon []byte

type backgroundTray struct {
	startOnce sync.Once
	stopOnce  sync.Once
	start     func()
	stop      func()
}

func newBackgroundTray(onOpen func(), onExit func()) *backgroundTray {
	start, stop := systray.RunWithExternalLoop(func() {
		if len(trayIcon) > 0 {
			systray.SetIcon(trayIcon)
		}
		systray.SetTooltip("duolin-gogo")

		openItem := systray.AddMenuItem("Open duolin-gogo", "Show the study window")
		systray.AddSeparator()
		exitItem := systray.AddMenuItem("Exit", "Quit duolin-gogo")

		go func() {
			for range openItem.ClickedCh {
				onOpen()
			}
		}()

		go func() {
			for range exitItem.ClickedCh {
				onExit()
			}
		}()
	}, func() {})

	return &backgroundTray{
		start: start,
		stop:  stop,
	}
}

func (t *backgroundTray) Start() {
	if t == nil || t.start == nil {
		return
	}
	t.startOnce.Do(t.start)
}

func (t *backgroundTray) Stop() {
	if t == nil || t.stop == nil {
		return
	}
	t.stopOnce.Do(t.stop)
}
