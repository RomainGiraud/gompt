package segments

import(
    "os/user"
)


type Username struct {
}

func (u Username) String() string {
    uc, _ := user.Current()
    return uc.Username
}
