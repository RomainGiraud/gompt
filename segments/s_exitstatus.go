package segments

import(
)


type ExitStatus struct {
    Status int
    Value string
}

func (e ExitStatus) String() string {
    if e.Status != 0 {
        return "X"
    }
    return e.Value
}
