package segments

import(
    "os/user"
)


type Username struct {
    Style
}

func (u Username) String() string {
    uc, _ := user.Current()
    return uc.Username
}
