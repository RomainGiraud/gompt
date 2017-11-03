package segments

import(
    "os"
)


type Hostname struct {
    Style
}

func (h Hostname) String() string {
    n, _ := os.Hostname()
    return n
}
