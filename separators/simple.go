package separators

import(
)


type Transition struct {
    Value string
}

func (t Transition) String() string {
    return t.Value
}
