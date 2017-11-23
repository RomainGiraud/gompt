package main

import(
    "os"
    "log"
    "flag"
    "io/ioutil"
    "prompt/segments"
)


func main() {
    var context segments.Context
    flag.IntVar(&context.Args.Status, "s", 0, "exit status")
    flag.StringVar(&context.Args.ConfigPath, "c", "", "config file path")
    flag.Parse()

    if len(context.Args.ConfigPath) != 0 {
        configFile, err := os.Open(context.Args.ConfigPath)
        if err != nil {
            log.Panic("wrong config file specified")
        }
        byteValue, _ := ioutil.ReadAll(configFile)
        context.LoadConfig(byteValue)
    } else {
        context.Segments = []segments.Segment {
            segments.NewUsername( segments.NewStyleUni(segments.FgRed, segments.BgGreen) ),
            segments.NewSeparator( "\ue0b0 ", segments.NewStyleUni(segments.FgGreen, segments.BgRed) ),
            segments.NewSeparator( "\ue0b0 ", segments.NewStyleUni(segments.FgRed, segments.BgYellow) ),
            segments.NewSeparator( "\ue0b0 ", segments.NewStyleUni(segments.FgYellow, segments.BgBlack) ),
        }
    }

    context.Display()
}
