package segments

import(
)


type ExitStatus struct {
    styleSuccess Style
    styleError Style
    textSuccess string
    textError string
}

func (e ExitStatus) Print(context Context, index int) {
    if context.Args.Status == 0 {
        FormatString(e.textSuccess, e.styleSuccess, context, index)
    } else {
        FormatString(e.textError, e.styleError, context, index)
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
