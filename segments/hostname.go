package segments

import(
    "os"
)


type Hostname struct {
}

func (h Hostname) String() string {
    n, _ := os.Hostname()
    return n
}
