package segments

import(
)


type Text struct {
    Value string
}

func (t Text) String() string {
    return t.Value
}
