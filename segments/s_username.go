package segments

import(
    "fmt"
    "log"
    "os/user"
    "encoding/json"
)


type Username struct {
    style Style
}

func (u Username) Print(context Context, index int) {
    uc, _ := user.Current()
    fmt.Print(u.style.Format(uc.Username, context, index))
}

func (u Username) GetStyle() Style {
    return u.style
}

type usernameConfig struct {
}

func NewUsername(bytes json.RawMessage, style Style) Segment {
    var config usernameConfig
    err := json.Unmarshal(bytes, &config)
    if err != nil {
        log.Fatal(err)
    }
    return &Username{ style }
}
