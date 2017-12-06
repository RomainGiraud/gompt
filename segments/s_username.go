package segments

import(
    "os/user"
)


type Username struct {
    style Style
}

func (u Username) Print(context Context, index int) {
    uc, _ := user.Current()
    FormatString(uc.Username, u.style, context, index)
}

func (u Username) GetStyle(context Context, index int) Style {
    return u.style
}

func NewUsername(style Style) Segment {
    return &Username{ style }
}

func LoadUsername(config map[string]interface{}) Segment {
    var style, _ = LoadStyle(config["style"])
    return &Username{ style }
}

func init() {
    RegisterSegmentLoader("username", LoadUsername)
}
