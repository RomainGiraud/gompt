package main

import(
    "fmt"
    "flag"
    "prompt/segments"
    "prompt/separators"
)


type Elements []fmt.Stringer

func (segments Elements) Print() {
    if len(segments) == 0 {
        panic("Empty prompt")
    }

    for i, j := 0, 1; i < len(segments); i, j = i+1, j+1 {
        s := segments[i]
        fmt.Printf("%v", s)

        /*
        switch s.(type) {
        case segments.Segment:
            fmt.Printf("seg/")
        case separators.Separator:
            fmt.Printf("sep/")
        }
        */

        /*
        color.Set(s.GetFg(), s.GetBg())

        fmt.Printf("%v", s)

        color.Unset()
        if j < len(segments) {
            sn := segments[j]
            color.Set(convertColors[s.GetBg()], sn.GetBg())
        } else {
            color.Set(convertColors[s.GetBg()])
        }
        fmt.Printf("%v", sep)
        */
    }
    fmt.Printf("\n")
}


func main() {
    status := flag.Int("s", 0, "exit status")
    flag.Parse()

    //fmt.Println(Colorize("toto", Bg24(0, 155, 0), Fg(30)))
    //fmt.Println(fmt.Sprintf(Color.Bg(46).Fg(30)("toto")))

    sep := separators.Transition{" > "}
    //sep := separators.Transition{" \ue0b0 "}
    seg := Elements{
        Color{segments.Path{}, []Style{ Bg24(0, 155, 0), Fg(30), Underline }},
        sep,
        segments.ExitStatus{*status, "\u25CF"},
        sep,
        segments.Text{ "rom" },
        sep,
        segments.Username{},
        sep,
        segments.Hostname{},
        sep,
    }
    seg.Print()
}
