package tui

import (
	"github.com/msugakov/release-stackrox/pkg/log"
	"github.com/pterm/pterm"
)

func ShowBanner(version string) {
	pterm.DefaultHeader.WithFullWidth().Println(pterm.LightBlue("StackRox") + "/" + pterm.Red("ACS") + " release tool")
	log.Infof("Version: %s", version)
}
