package segments

import(
    "fmt"
    "os"
)


type Hostname struct {
    style Style
}

func (h Hostname) Print(context Context, index int) {
    n, _ := os.Hostname()
    fmt.Print(h.style.Format(n, context, index))
}

func (h Hostname) GetStyle(context Context, index int) Style {
    return h.style
}

func NewHostname(style Style) Segment {
    return &Hostname{ style }
}

func LoadHostname(config map[string]interface{}) Segment {
    var style, _ = LoadStyle(config["style"])
    return &Hostname{ style }
}

func init() {
    RegisterSegmentLoader("hostname", LoadHostname)
}