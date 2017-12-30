package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"os/user"
)

// UsernameLoader create a segment containing the current username.
type UsernameLoader struct {
	Style format.Style
}

func (s UsernameLoader) Load() []Segment {
	u, _ := user.Current()
	return []Segment{Text{s.Style, " " + u.Username + " "}}
}
