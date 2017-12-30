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

	prompt := segments.LoaderList{
		segments.TextLoader{exitStatusStyle, "\uf444"},
		segments.UsernameLoader{format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.NewColor("8")})},
		segments.TextLoader{format.StyleChameleon{}, "\ue0b0"},
		segments.HostnameLoader{format.NewStyleStandard(format.UniBrush{format.Black}, format.UniBrush{format.Blue})},
		segments.TextLoader{format.StyleChameleon{}, "\ue0b0"},
		segments.CurrentDirBlockLoader{
			format.NewStyleStandard(
				format.UniBrush{format.NewColor("#333")},
				format.GradientBrush{format.NewColor("#aaa"), format.NewColor("#eee")}),
			format.NewStyleStandard(
				format.UniBrush{format.NewColor("#333")},
				format.UniBrush{format.NewColor("#eee")}),
			"\ue0b4", 3, "\u2026"},
		segments.BindingLoader{
			segments.TextLoader{format.StyleChameleon{}, "\ue0b0"},
			segments.GitLoader{format.NewStyleStandard(format.UniBrush{format.Black}, format.UniBrush{format.Cyan})}},
		segments.TextLoader{format.StyleChameleon{}, "\ue0b0"},
		segments.TextLoader{format.NewStyleStandard(format.UniBrush{}, format.UniBrush{}), " "},
	}

	var buffer bytes.Buffer
	prompt.Load().Render(&buffer)
	fmt.Println(buffer.String())
}
