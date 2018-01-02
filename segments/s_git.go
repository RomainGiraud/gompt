package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
	"strconv"
	"strings"
)

// GitLoader displays information about git repository.
// If current directory is not a repository, do not display anything.
type GitLoader struct {
	Style format.Style
}

func (s GitLoader) Load() []Segment {
	statusOutput, err := ExecCommand("git", "status", "--porcelain", "--branch")
	if err != nil {
		return []Segment{}
	}

	status := strings.Split(statusOutput, "\n")

	branch, ahead, behind := parseBranch(status[0])
	clean := (len(status[1:]) == 0)
	stash := getStashCount()

	return []Segment{git{s.Style, branch, ahead, behind, stash, clean}}
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

type git struct {
	style  format.Style
	branch string
	ahead  int
	behind int
	stash  int
	clean  bool
}

func (s git) Print(writer io.Writer, segments []Segment, current int) {
	ff := []PartFormatter{
		PartFormatter{" ", nil, nil},
		PartFormatter{s.branch, nil, nil},
	}
	if s.ahead != 0 {
		ff = append(ff, PartFormatter{"\uf139" + strconv.Itoa(s.ahead), nil, nil})
	}
	if s.behind != 0 {
		ff = append(ff, PartFormatter{"\uf13a" + strconv.Itoa(s.behind), nil, nil})
	}
	if !s.clean || s.stash != 0 {
		//ff = append(ff, PartFormatter{ "|", nil, nil })
		if s.stash != 0 {
			ff = append(ff, PartFormatter{"\uf111", format.Blue, nil})
		}
		if !s.clean {
			ff = append(ff, PartFormatter{"\uf057", format.Red, nil})
		}
	}
	ff = append(ff, PartFormatter{" ", nil, nil})
	FormatParts(writer, s.style, segments, current, ff)
}

func (s git) GetStyle(segments []Segment, current int) format.Style {
	return s.style
}

