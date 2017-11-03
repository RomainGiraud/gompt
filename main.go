package main

import(
    "fmt"
    "flag"
    "github.com/fatih/color"
    "prompt/segments"
)


func Print(segments []segments.Segment) {
    if len(segments) == 0 {
        panic("Empty prompt")
    }

    sep := "\ue0b0"
    for i, j := 0, 1; i < len(segments); i, j = i+1, j+1 {
        s := segments[i]
        color.Set(s.GetFg(), s.GetBg())

        if i != 0 {
            fmt.Printf(" ")
        }
        fmt.Printf("%v ", s)

        color.Unset()
        if j < len(segments) {
            sn := segments[j]
            color.Set(convertColors[s.GetBg()], sn.GetBg())
        } else {
            color.Set(convertColors[s.GetBg()])
        }
        fmt.Printf("%v", sep)
    }
    color.Unset()
    fmt.Printf(" \n");
}


func main() {
    status := flag.Int("s", 0, "exit status")
    flag.Parse()
    fmt.Println("=> ", *status)


    seg := make([]segments.Segment, 0, 8)
    seg = append(seg, segments.ExitStatus{
        segments.Style{color.FgWhite, color.BgWhite},
        segments.Style{color.FgWhite, color.BgRed},
        *status})
    seg = append(seg, segments.Path{ segments.Style{color.FgBlack, color.BgWhite}})
    seg = append(seg, segments.Text{ segments.Style{color.FgWhite, color.BgBlue}, "rom"})
    seg = append(seg, segments.Username{ segments.Style{color.FgWhite, color.BgGreen}})
    seg = append(seg, segments.Hostname{ segments.Style{color.FgWhite, color.BgMagenta}})
    Print(seg)
}
