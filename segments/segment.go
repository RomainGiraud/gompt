package segments

import(
    "fmt"
)


type Segment interface {
    Print(Context, string)
    GetStyle() Style
}

type Arguments struct {
    Status int
    ConfigPath string
}

type Context struct {
    Args Arguments
    Order []string
    Segments map[string]Segment
}

func (c Context) Display() {
    if len(c.Segments) == 0 || len(c.Order) == 0 {
        panic("Empty prompt")
    }

    for i, j := 0, 1; i < len(c.Order); i, j = i+1, j+1 {
        curr := c.Order[i]
        seg := c.Segments[curr]
        seg.Print(c, curr)

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

