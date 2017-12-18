package segments

import(
    "io"
    "os"
    "os/exec"
    "strings"
    "bytes"
    "log"
)


type Text struct {
    style Style
    value string
}

func getenv(key string) string {
    // Escape $ by doubling it: $$
    if key == "$" {
        return "$"
    }

    // Execute "my_command" from ${cmd> my_command}
    if strings.HasPrefix(key, "cmd> ") {
        cmd := exec.Command("bash", "-c", key[5:])
        var out bytes.Buffer
        cmd.Stdout = &out
        err := cmd.Run()
        if err != nil {
            log.Fatal(err)
        }
        return strings.Trim(out.String(), "\n")
    }

    return os.Getenv(key)
}

func (s Text) Print(writer io.Writer, segments []Segment, current int) {
    str := os.Expand(s.value, getenv)
    FormatString(writer, str, s.style, segments, current)
}

func (s Text) GetStyle(segments []Segment, current int) Style {
    return s.style
}

func NewText(style Style, text string) Segment {
    return &Text{ style, text }
}