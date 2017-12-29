package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/RomainGiraud/gompt/format"
	"github.com/RomainGiraud/gompt/segments"
)

func main() {
	var lastCommandStatus int
	flag.IntVar(&lastCommandStatus, "s", 0, "exit status")
	flag.Parse()

	var exitStatusStyle *format.StyleStandard
	if lastCommandStatus == 0 {
		exitStatusStyle = format.NewStyleStandard(
			format.UniBrush{format.NewColor("#0f0")},
			format.UniBrush{format.NewColor("8")})
	} else {
		exitStatusStyle = format.NewStyleStandard(
			format.UniBrush{format.NewColor("#f00")},
			format.UniBrush{format.NewColor("8")})
	}

	segmentList := segments.SegmentList{
		segments.NewText(exitStatusStyle, "\uf444"),
		segments.NewUsername(format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.NewColor("8")})),
		segments.NewText(format.StyleChameleon{}, "\ue0b0"),
		segments.NewHostname(format.NewStyleStandard(format.UniBrush{format.Black}, format.UniBrush{format.Blue})),
		segments.NewText(format.StyleChameleon{}, "\ue0b0"),
		segments.NewCurrentDirSplitted(
			format.NewStyleStandard(
				format.UniBrush{format.NewColor("#333")},
				format.GradientBrush{format.NewColor("#aaa"), format.NewColor("#eee")}),
			format.NewStyleStandard(
				format.UniBrush{format.NewColor("#333")},
				format.UniBrush{format.NewColor("#eee")}),
			"\ue0b4", 3, "\u2026"),
		segments.NewBinding(
			segments.NewText(format.StyleChameleon{}, "\ue0b0"),
			segments.NewGit(format.NewStyleStandard(format.UniBrush{format.Black}, format.UniBrush{format.Cyan}))),
		segments.NewText(format.StyleChameleon{}, "\ue0b0"),
		segments.NewText(format.NewStyleStandard(format.UniBrush{}, format.UniBrush{}), " "),
	}

	var buffer bytes.Buffer
	segmentList.Render(&buffer)
	fmt.Println(buffer.String())
}
