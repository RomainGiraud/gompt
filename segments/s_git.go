package segments

import (
	"bufio"
	"errors"
	"github.com/RomainGiraud/gompt/format"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Git segment prints git information.
type Git struct {
	Style         format.Style
	DetachedColor format.Color
	BareColor     format.Color
	StateColor    format.Color

	StagedSymbol    string
	StagedColor     format.Color
	UnstagedSymbol  string
	UnstagedColor   format.Color
	UntrackedSymbol string
	UntrackedColor  format.Color
	StashSymbol     string
	StashColor      format.Color
	AheadSymbol     string
	AheadColor      format.Color
	BehindSymbol    string
	BehindColor     format.Color

	branch        string
	state         string
	isDetached    bool
	isBare        bool
	hasStaged     bool
	hasUnstaged   bool
	hasUntracked  bool
	hasStash      bool
	ahead, behind int
}

// Create a Git segment.
func NewGit() (*Git, error) {
	out, err := ExecCommand("git", "rev-parse", "--git-dir", "--is-inside-git-dir", "--is-bare-repository", "--is-inside-work-tree", "--short", "HEAD")
	if len(out) == 0 {
		return nil, errors.New("Not a git repository.")
	}

	info := strings.Split(out, "\n")
	git_dir := info[0]
	in_gitdir := (info[1] == "true")
	is_bare := (info[2] == "true")
	in_wd := (info[3] == "true")
	sha := ""
	if err == nil {
		sha = info[4]
	}

	branch := ""
	state := ""
	step := ""
	total := ""
	detached := false

	if stat, err := os.Stat(git_dir + "/rebase-merge"); err == nil && stat.IsDir() {
		branch = readFirstLine(git_dir + "/rebase-merge/head-name")
		step = readFirstLine(git_dir + "/rebase-merge/msgnum")
		total = readFirstLine(git_dir + "/rebase-merge/end")
		if stat, err = os.Stat(git_dir + "/rebase-merge/interactive"); err == nil && stat.Mode().IsRegular() {
			state = "|REBASE-i"
		} else {
			state = "|REBASE-m"
		}
	} else {
		if stat, err := os.Stat(git_dir + "/rebase-apply"); err == nil && stat.IsDir() {
			step = readFirstLine(git_dir + "/rebase-apply/next")
			total = readFirstLine(git_dir + "/rebase-apply/last")
			if stat, err = os.Stat(git_dir + "/rebase-apply/rebasing"); err == nil && stat.Mode().IsRegular() {
				branch = readFirstLine(git_dir + "/rebase-apply/head-name")
				state = "|REBASE"
			} else if stat, err = os.Stat(git_dir + "/rebase-apply/applying"); err == nil && stat.Mode().IsRegular() {
				state = "|AM"
			} else {
				state = "|AM/REBASE"
			}
		} else if stat, err = os.Stat(git_dir + "/MERGE_HEAD"); err == nil && stat.Mode().IsRegular() {
			state = "|MERGING"
		} else if stat, err = os.Stat(git_dir + "/CHERRY_PICK_HEAD"); err == nil && stat.Mode().IsRegular() {
			state = "|CHERRY-PICKING"
		} else if stat, err = os.Stat(git_dir + "/REVERT_HEAD"); err == nil && stat.Mode().IsRegular() {
			state = "|REVERTING"
		} else if stat, err = os.Stat(git_dir + "/BISECT_LOG"); err == nil && stat.Mode().IsRegular() {
			state = "|BISECTING"
		}
	}

	if len(branch) != 0 {
	} else if stat, err := os.Stat(git_dir + "/HEAD"); err == nil && (stat.Mode()&os.ModeSymlink) != 0 {
		branch, _ = ExecCommand("git", "symbolic-ref", "HEAD")
	} else {
		head := ""
		if head = readFirstLine(git_dir + "/HEAD"); len(head) == 0 {
			return nil, errors.New("Cannot read .git/HEAD")
		}
		branch = strings.TrimPrefix(head, "ref: ")
		if branch == head {
			detached = true
			branch, _ = ExecCommand("git", "describe", "--contains", "--all", "HEAD")
			if len(branch) == 0 {
				branch = ":" + sha
			}
		}
	}

	if len(step) != 0 && len(total) != 0 {
		state = state + "(" + step + "/" + total + ")"
	}

	bare := false
	unstaged := false
	staged := false
	stash := false
	untracked := false
	ahead, behind := 0, 0

	if in_gitdir {
		if is_bare {
			bare = true
		} else {
			branch = "GIT_DIR!"
		}
	} else if in_wd {
		out, err = ExecCommand("git", "diff", "--no-ext-diff", "--quiet")
		unstaged = (err != nil)

		out, err = ExecCommand("git", "diff", "--no-ext-diff", "--quiet", "--cached")
		staged = (err != nil)

		out, err = ExecCommand("git", "rev-parse", "--verify", "--quiet", "refs/stash")
		stash = (err == nil)

		out, err = ExecCommand("git", "ls-files", "--others", "--exclude-standard", "--directory", "--no-empty-directory", "--error-unmatch", "--", ":/*")
		untracked = (err == nil)

		out, err = ExecCommand("git", "rev-list", "--count", "--left-right", "@{upstream}...HEAD")
		if err == nil {
			split := strings.Split(out, "\t")
			behind, _ = strconv.Atoi(split[0])
			ahead, _ = strconv.Atoi(split[1])
		}
	}

	return &Git{
		format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.Black}), // Style
		format.Red,        // DetachedColor
		format.Magenta,    // BareColor
		format.Black,      // StateColor
		"+", format.Black, // Staged
		"*", format.Black, // Unstaged
		"\uf057", format.Black, // Untracked
		"\uf024", format.Black, // Stash
		"\uf139", format.Black, // Ahead
		"\uf13a", format.Black, // Behind

		strings.TrimPrefix(branch, "refs/heads/"), // branch
		state,     // state
		detached,  // isDetached
		bare,      // isBare
		staged,    // hasStaged
		unstaged,  // hasUnstaged
		untracked, // hasUntracked
		stash,     // hasStash
		ahead,     // ahead
		behind,    // behind
	}, nil
}

func readFirstLine(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func (s *Git) Load() {
}

func (s Git) Print(writer io.Writer, segments []Segment, current int) {
	ff := []PartFormatter{}

	ff = append(ff, PartFormatter{" ", nil, nil})

	var branchColor format.Color
	if s.isBare {
		branchColor = s.BareColor
	} else if s.isDetached {
		branchColor = s.DetachedColor
	}
	ff = append(ff, PartFormatter{s.branch, branchColor, nil})

	if len(s.state) != 0 {
		ff = append(ff, PartFormatter{s.state, s.StateColor, nil})
	}

	if s.ahead != 0 {
		ff = append(ff, PartFormatter{s.AheadSymbol + strconv.Itoa(s.ahead), s.AheadColor, nil})
	}
	if s.behind != 0 {
		ff = append(ff, PartFormatter{s.BehindSymbol + strconv.Itoa(s.behind), s.BehindColor, nil})
	}

	if s.hasStaged {
		ff = append(ff, PartFormatter{s.StagedSymbol, s.StagedColor, nil})
	}
	if s.hasUnstaged {
		ff = append(ff, PartFormatter{s.UnstagedSymbol, s.UnstagedColor, nil})
	}
	if s.hasUntracked {
		ff = append(ff, PartFormatter{s.UntrackedSymbol, s.UntrackedColor, nil})
	}
	if s.hasStash {
		ff = append(ff, PartFormatter{s.StashSymbol, s.StashColor, nil})
	}

	ff = append(ff, PartFormatter{" ", nil, nil})

	FormatParts(writer, s.Style, segments, current, ff)
}

func (s Git) GetStyle(segments []Segment, current int) format.Style {
	return s.Style
}
