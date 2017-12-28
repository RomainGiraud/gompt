package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/RomainGiraud/gompt/segments"
)

func main() {
	var lastCommandStatus int
	flag.IntVar(&lastCommandStatus, "s", 0, "exit status")
	flag.Parse()

	var exitStatusStyle *segments.StyleStandard
	if lastCommandStatus == 0 {
		exitStatusStyle = segments.NewStyleStandard(
			segments.UniBrush{segments.NewColor("#0f0")},
			segments.UniBrush{segments.NewColor("8")})
	} else {
		exitStatusStyle = segments.NewStyleStandard(
			segments.UniBrush{segments.NewColor("#f00")},
			segments.UniBrush{segments.NewColor("8")})
	}

	segmentList := segments.SegmentList{
		segments.NewText(exitStatusStyle, "\uf444"),
		segments.NewUsername(segments.NewStyleStandard(segments.UniBrush{segments.White}, segments.UniBrush{segments.NewColor("8")})),
		segments.NewText(segments.StyleChameleon{}, "\ue0b0"),
		segments.NewHostname(segments.NewStyleStandard(segments.UniBrush{segments.Black}, segments.UniBrush{segments.Blue})),
		segments.NewText(segments.StyleChameleon{}, "\ue0b0"),
		segments.NewComplexPathSplitted(
			segments.NewStyleStandard(
				segments.UniBrush{segments.NewColor("#333")},
				segments.GradientBrush{segments.NewColor("#aaa"), segments.NewColor("#eee")}),
			segments.NewStyleStandard(
				segments.UniBrush{segments.NewColor("#333")},
				segments.UniBrush{segments.NewColor("#eee")}),
			"\ue0b4", 3, "\u2026"),
		segments.NewBinding(
			segments.NewText(segments.StyleChameleon{}, "\ue0b0"),
			segments.NewGit(segments.NewStyleStandard(segments.UniBrush{segments.Black}, segments.UniBrush{segments.Cyan}))),
		segments.NewText(segments.StyleChameleon{}, "\ue0b0"),
		segments.NewText(segments.NewStyleStandard(segments.UniBrush{}, segments.UniBrush{}), " "),
	}

	var buffer bytes.Buffer
	segmentList.Render(&buffer)
	fmt.Println(buffer.String())
}
