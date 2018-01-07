package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/RomainGiraud/gompt/format"
	"github.com/RomainGiraud/gompt/segment"
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

	sep := segment.NewTextStylized("\ue0b0", format.NewStyleChameleon())

	var prompt segment.SegmentList
	prompt = append(prompt, segment.NewTextStylized("\uf444", exitStatusStyle))
	prompt = append(prompt, segment.NewFullUsernameStylized(format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.NewColor("8")})))
	prompt = append(prompt, sep)
	prompt = append(prompt, segment.NewFullHostnameStylized(format.NewStyleStandard(format.UniBrush{format.Black}, format.UniBrush{format.Blue})))
	prompt = append(prompt, sep)

	if seg, err := segment.NewVirtualEnvStylized(format.NewStyleStandard(format.UniBrush{format.Black}, format.UniBrush{format.Yellow})); err == nil {
		prompt = append(prompt, seg)
		prompt = append(prompt, sep)
	}

	if seg, err := segment.NewDirEnvStylized(format.NewStyleStandard(format.UniBrush{format.Black}, format.UniBrush{format.Magenta})); err == nil {
		prompt = append(prompt, seg)
		prompt = append(prompt, sep)
	}

	cwd := segment.NewCurrentDirBlock()
	cwd.Style = format.NewStyleStandard(
		format.UniBrush{format.NewColor("#333")},
		format.GradientBrush{format.NewColor("#aaa"), format.NewColor("#eee")})
	cwd.StyleUnit = format.NewStyleStandard(
		format.UniBrush{format.NewColor("#333")},
		format.UniBrush{format.NewColor("#eee")})
	cwd.Separator = "\ue0b4"
	cwd.MaxDepth = 3
	prompt = append(prompt, cwd)

	if seg, err := segment.NewGit(); err == nil {
		prompt = append(prompt, sep)

		seg.Style = format.NewStyleStandard(format.UniBrush{format.Black}, format.UniBrush{format.Cyan})
		prompt = append(prompt, seg)
	}

	prompt = append(prompt, sep)
	prompt = append(prompt, segment.NewText(" "))

	var buffer bytes.Buffer
	prompt.Render(&buffer)
	fmt.Println(buffer.String())
}
