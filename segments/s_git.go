package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
	"strconv"
	"strings"
)

// Git segment prints git information.
type Git struct {
	Style        format.Style
	AheadSymbol  string
	AheadColor   format.Color
	BehindSymbol string
	BehindColor  format.Color
	StashSymbol  string
	StashColor   format.Color
	DirtySymbol  string
	DirtyColor   format.Color
	branch       string
	ahead        int
	behind       int
	stash        int
	clean        bool
}

// Create a Git segment.
func NewGit() *Git {
	// in a repository?
	_, err := ExecCommand("git", "status")
	if err != nil {
		return nil
	}

	return &Git{
		format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.Black}),
		"\uf139", format.Black,
		"\uf13a", format.Black,
		"\uf024", format.Black,
		"\uf057", format.Red,
		"", 0, 0, 0, true}
}

func (s *Git) Load() {
	statusOutput, err := ExecCommand("git", "status", "--porcelain", "--branch")
	if err != nil {
		return
	}

	status := strings.Split(statusOutput, "\n")

	s.branch, s.ahead, s.behind = parseBranch(status[0])
	s.clean = (len(status[1:]) == 0)
	s.stash = getStashCount()
}

func (s Git) Print(writer io.Writer, segments []Segment, current int) {
	ff := []PartFormatter{
		PartFormatter{" ", nil, nil},
		PartFormatter{s.branch, nil, nil},
	}
	if s.ahead != 0 {
		ff = append(ff, PartFormatter{s.AheadSymbol + strconv.Itoa(s.ahead), s.AheadColor, nil})
	}
	if s.behind != 0 {
		ff = append(ff, PartFormatter{s.BehindSymbol + strconv.Itoa(s.behind), s.BehindColor, nil})
	}
	if !s.clean || s.stash != 0 {
		//ff = append(ff, PartFormatter{ "|", nil, nil })
		if s.stash != 0 {
			ff = append(ff, PartFormatter{s.StashSymbol, s.StashColor, nil})
		}
		if !s.clean {
			ff = append(ff, PartFormatter{s.DirtySymbol, s.DirtyColor, nil})
		}
	}
	ff = append(ff, PartFormatter{" ", nil, nil})
	FormatParts(writer, s.Style, segments, current, ff)
}

func (s Git) GetStyle(segments []Segment, current int) format.Style {
	return s.Style
}

func parseBranch(str string) (string, int, int) {
	var branch string
	var ahead, behind int

	if strings.HasPrefix(str, "## No commits yet on ") {
		str = strings.TrimPrefix(str, "## No commits yet on ")
		branch = "(empty)"
	} else if strings.Contains(str, "(no branch)") {
		branch = "(detached)"
	} else {
		items := strings.Split(str, " ")
		// link to remote
		if strings.Contains(items[1], "...") {
			branches := strings.Split(items[1], "...")
			branch = branches[0]
			for i := 2; i < len(items); i += 1 {
				if items[i] == "[gone]" {
					branch += "(gone)"
				} else if strings.HasPrefix(items[i], "[ahead") {
					ahead, _ = strconv.Atoi(strings.TrimRight(items[i+1], "],"))
					i += 1
				} else if strings.HasPrefix(items[i], "[behind") || strings.HasPrefix(items[i], "behind") {
					behind, _ = strconv.Atoi(strings.TrimRight(items[i+1], "],"))
					i += 1
				}
			}
		} else {
			branch = items[1]
		}
	}

	return branch, ahead, behind
}

func getStashCount() int {
	stashOutput, err := ExecCommand("git", "stash", "list")
	if err != nil {
		return 0
	}

	stashOutput = strings.TrimSpace(stashOutput)
	if len(stashOutput) == 0 {
		return 0
	}
	return len(strings.Split(stashOutput, "\n"))
}
