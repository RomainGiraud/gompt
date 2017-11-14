package segments

import(
    "fmt"
)


type Segment interface {
    Print(Context, int)
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
        seg := c.Segments[c.Order[i]]
        seg.Print(c, i)
    }
    fmt.Printf("\n")
}

