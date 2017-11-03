package segments

import(
)


type Text struct {
    Style
    Value string
}

func (t Text) String() string {
    return t.Value
}
