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
            segments.NewUsername( segments.NewStyleUni(segments.NewColor("22"), segments.Green) ),
            segments.NewSeparator( "\ue0b0", segments.NewStyleChameleon() ),
            segments.NewPath( segments.NewStyleUni(segments.Blue, segments.NewColor("25")) ),
            segments.NewSeparator( "\ue0b0", segments.NewStyleChameleon() ),
        }
    }

    context.Display()
}
