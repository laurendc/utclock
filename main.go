package main

import (
	"fmt"
	"os"
	"time"

	"github.com/caseymrm/menuet"
)

func menuItems() []menuet.MenuItem {
	now := time.Now()
	utc := now.UTC()

	return []menuet.MenuItem{
		{
			Text: fmt.Sprintf("Local:  %s", now.Format("Mon 03:04:05 PM MST")),
		},
		{
			Text: fmt.Sprintf("UTC:    %s", utc.Format("Mon 15:04:05 MST")),
		},
		{
			Type: menuet.Separator,
		},
		{
			Text: "Quit",
			Clicked: func() {
				os.Exit(0)
			},
		},
	}
}

func updateTitle() {
	for {
		menuet.App().SetMenuState(&menuet.MenuState{
			Title: time.Now().UTC().Format("15:04 UTC"),
		})
		time.Sleep(time.Second)
	}
}

func main() {
	go updateTitle()

	menuet.App().Label = "com.lauren.utclock"
	menuet.App().Children = menuItems
	menuet.App().RunApplication()
}
