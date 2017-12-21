package segments

import(
    "io"
    _ "fmt"
    "strings"
    "strconv"
)


type Git struct {
    style Style
    branch string
    ahead int
    behind int
    stash int
    clean bool
}

func ParseBranch(str string) (string, int, int) {
    //fmt.Println("status: ", str)
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
                    ahead, _ = strconv.Atoi(strings.TrimRight(items[i + 1], "],"))
                    i += 1
                } else if strings.HasPrefix(items[i], "[behind") || strings.HasPrefix(items[i], "behind") {
                    behind, _ = strconv.Atoi(strings.TrimRight(items[i + 1], "],"))
                    i += 1
                }
            }
        } else {
            branch = items[1]
        }
    }

    return branch, ahead, behind
}

func GetStashCount() int {
    stashOutput := ExecCommand("git", "stash", "list")
    if len(stashOutput) == 0 {
        return 0
    }

    return len(strings.Split(stashOutput, "\n"))
}

func (s Git) Load() []Segment {
    statusOutput := ExecCommand("git", "status", "--porcelain", "--branch")
    if len(statusOutput) == 0 {
        return []Segment{}
    }

    status := strings.Split(statusOutput, "\n")

    s.branch, s.ahead, s.behind = ParseBranch(status[0])
    s.clean = (len(status[1:]) == 0)
    s.stash = GetStashCount()

    return []Segment{ s }
}

func (s Git) Print(writer io.Writer, segments []Segment, current int) {
    ff := []PartFormatter{
        PartFormatter{ " ", nil, nil },
        PartFormatter{ s.branch, nil, nil },
    }
    if s.ahead != 0 {
        ff = append(ff, PartFormatter{ "\uf139" + strconv.Itoa(s.ahead), nil, nil })
    }
    if s.behind != 0 {
        ff = append(ff, PartFormatter{ "\uf13a" + strconv.Itoa(s.behind), nil, nil })
    }
    if !s.clean || s.stash != 0 {
        //ff = append(ff, PartFormatter{ "|", nil, nil })
        if s.stash != 0 {
            ff = append(ff, PartFormatter{ "\uf111", Blue, nil })
        }
        if !s.clean {
            ff = append(ff, PartFormatter{ "\uf057", Red, nil })
        }
    }
    ff = append(ff, PartFormatter{ " ", nil, nil })
    FormatParts(writer, s.style, segments, current, ff)
}

func (s Git) GetStyle(segments []Segment, current int) Style {
    return s.style
}

func NewGit(style Style) Segment {
    return &Git{ style, "", 0, 0, 0, true }
}