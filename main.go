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
            segments.NewUsername( segments.StyleStandard{ segments.UniBrush{ segments.Green }, segments.UniBrush{ segments.NewColor("#555") } } ),
            segments.NewSeparator( "\u2588\u2588\u2588\u2588\ue0b0        ", segments.StyleChameleon{ } ),
            segments.NewText( segments.StyleStandard{ segments.UniBrush{ segments.NewColor("#0000ff") }, segments.UniBrush{ segments.NewColor("#fff") } }, "${cmd> ls -la | wc -l}@" ),
            segments.NewHostname( segments.StyleStandard{ segments.UniBrush{ segments.NewColor("#0000f0") }, segments.GradientBrush{ segments.NewColor("#fff"), segments.NewColor("#aaa") } } ),
            segments.NewSeparator( "\ue0b0", segments.StyleChameleon{ } ),
            segments.NewComplexPath(
                segments.StyleStandard{
                    segments.UniBrush{ segments.NewColor("#555") },
                    segments.GradientBrush{ segments.NewColor("#0000ff"), segments.NewColor("#ff0000") } },
                "\ue0b1", segments.Red ),
            segments.NewSeparator( "\ue0b0", segments.StyleChameleon{ } ),
        }
    }

    context.Display()
}
