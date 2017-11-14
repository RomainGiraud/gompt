package segments

import(
    "fmt"
    "log"
    "os"
    "encoding/json"
)


type Hostname struct {
    style Style
}

func (h Hostname) Print(context Context, index int) {
    n, _ := os.Hostname()
    fmt.Print(h.style.Format(n, context, index))
}

func (h Hostname) GetStyle() Style {
    return h.style
}

type hostnameConfig struct {
}

func NewHostname(bytes json.RawMessage, style Style) Segment {
    var config hostnameConfig
    err := json.Unmarshal(bytes, &config)
    if err != nil {
        log.Fatal(err)
    }
    return &Hostname{ style }
}
