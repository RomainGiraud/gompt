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

	sep := segments.NewTextStylized("\ue0b0", format.NewStyleChameleon())

	var prompt segments.SegmentList
	prompt = append(prompt, segments.NewTextStylized("\uf444", exitStatusStyle))
	prompt = append(prompt, segments.NewFullUsernameStylized(format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.NewColor("8")})))
	prompt = append(prompt, sep)
	prompt = append(prompt, segments.NewFullHostnameStylized(format.NewStyleStandard(format.UniBrush{format.Black}, format.UniBrush{format.Blue})))
	prompt = append(prompt, sep)

	venv := segments.NewVirtualEnvStylized(format.NewStyleStandard(format.UniBrush{format.Black}, format.UniBrush{format.Yellow}))
	if venv != nil {
		prompt = append(prompt, venv)
		prompt = append(prompt, sep)
	}

	direnv := segments.NewDirEnvStylized(format.NewStyleStandard(format.UniBrush{format.Black}, format.UniBrush{format.Magenta}))
	if direnv != nil {
		prompt = append(prompt, direnv)
		prompt = append(prompt, sep)
	}

	cwd := segments.NewCurrentDirBlock()
	cwd.Style = format.NewStyleStandard(
		format.UniBrush{format.NewColor("#333")},
		format.GradientBrush{format.NewColor("#aaa"), format.NewColor("#eee")})
	cwd.StyleUnit = format.NewStyleStandard(
		format.UniBrush{format.NewColor("#333")},
		format.UniBrush{format.NewColor("#eee")})
	cwd.Separator = "\ue0b4"
	cwd.MaxDepth = 3
	prompt = append(prompt, cwd)

	if git := segments.NewGit(); git != nil {
		prompt = append(prompt, sep)

		git.Style = format.NewStyleStandard(format.UniBrush{format.Black}, format.UniBrush{format.Cyan})
		prompt = append(prompt, git)
	}

	prompt = append(prompt, sep)
	prompt = append(prompt, segments.NewText(" "))

	var buffer bytes.Buffer
	prompt.Render(&buffer)
	fmt.Println(buffer.String())
}
