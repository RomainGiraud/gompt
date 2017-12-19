package segments

import(
    "io"
    "os"
    "strings"
)


type Text struct {
    style Style
    value string
}

func (s Text) Load() []Segment {
    return []Segment{ s }
}

func (s Text) Print(writer io.Writer, segments []Segment, current int) {
    text := os.Expand(s.value, getenv)
    FormatString(writer, text, s.style, segments, current)
}

func (s Text) GetStyle(segments []Segment, current int) Style {
    return s.style
}

func NewText(style Style, text string) Segment {
    return &Text{ style, text }
}

func getenv(key string) string {
    // Escape $ by doubling it: $$
    if key == "$" {
        return "$"
    }

    // Execute "my_command" from ${cmd> my_command}
    if strings.HasPrefix(key, "cmd> ") {
        return ExecCommand("bash", "-c", key[5:])
    }

    return os.Getenv(key)
}