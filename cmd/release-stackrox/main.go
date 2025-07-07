package main

import (
	"github.com/msugakov/release-stackrox/pkg/tui"
	"github.com/msugakov/release-stackrox/pkg/version"
)

func main() {
	ver := version.GetVersion()

	tui.ShowBanner(ver)
}
