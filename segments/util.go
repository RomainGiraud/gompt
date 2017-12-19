package segments

import(
    "os/exec"
    "strings"
    "bytes"
)

func ExecCommand(name string, arg ...string) string {
    cmd := exec.Command(name, arg...)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        return ""
    }
    return strings.Trim(out.String(), "\n")
}