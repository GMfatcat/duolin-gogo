package main

import "testing"

func TestNewAppReturnsApp(t *testing.T) {
	app := NewApp()

	if app == nil {
		t.Fatal("expected app instance")
	}
}

func TestAppInfo(t *testing.T) {
	app := NewApp()

	info := app.AppInfo()

	if info.Name != "duolin-gogo" {
		t.Fatalf("expected app name duolin-gogo, got %q", info.Name)
	}

	if info.FocusTopic != "git" {
		t.Fatalf("expected focus topic git, got %q", info.FocusTopic)
	}

	if info.DefaultLanguage != "zh-TW" {
		t.Fatalf("expected default language zh-TW, got %q", info.DefaultLanguage)
	}
}
