package segments

import(
    "fmt"
)


type ExitStatus struct {
    styleSuccess Style
    styleError Style
    textSuccess string
    textError string
}

func (e ExitStatus) Print(context Context, index int) {
    if context.Args.Status == 0 {
        fmt.Print(e.styleSuccess.Format(e.textSuccess, context, index, 0))
    } else {
        fmt.Print(e.styleError.Format(e.textError, context, index, 0))
    }
}

func (e ExitStatus) GetStyle(context Context, index int) Style {
    if context.Args.Status == 0 {
        return e.styleSuccess
    } else {
        return e.styleError
    }
}

func NewExitStatus(styleSuccess Style, styleError Style, textSuccess string, textError string) Segment {
    return &ExitStatus{ styleSuccess, styleError, textSuccess, textError }
}

func LoadExitStatus(config map[string]interface{}) Segment {
    var sSuccess, _ = LoadStyle(config["style-success"])
    var sError,   _ = LoadStyle(config["style-error"])
    var tSuccess, _ = config["text-success"].(string)
    var tError,   _ = config["text-error"].(string)
    return &ExitStatus{ sSuccess, sError, tSuccess, tError }
}

func init() {
    RegisterSegmentLoader("exit-status", LoadExitStatus)
}
