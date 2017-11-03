package segments

import(
    "os"
    "log"
)


type Path struct {
    Style
}

func (p Path) String() string {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    return dir
}
