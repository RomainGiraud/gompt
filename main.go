package main

import(
    "flag"
    "github.com/RomainGiraud/gompt/segments"
)


func main() {
    var lastCommandStatus int
    flag.IntVar(&lastCommandStatus, "s", 0, "exit status")
    flag.Parse()

    var exitStatus segments.Segment
    if lastCommandStatus == 0 {
        exitStatus = segments.NewText(
            segments.StyleStandard{
                segments.UniBrush{ segments.NewColor("#555") },
                segments.UniBrush{ segments.NewColor("#555") } },
            "\ue0b0" )
    } else {
        exitStatus = segments.NewText(
            segments.StyleStandard{
                segments.UniBrush{ segments.NewColor("#f00") },
                segments.UniBrush{ segments.NewColor("#555") } },
            "\ue0b0" )
    }

    segmentList := segments.SegmentList {
        exitStatus,
        segments.NewUsername( segments.StyleStandard{ segments.UniBrush{ segments.Green }, segments.UniBrush{ segments.NewColor("#555") } } ),
        segments.NewSeparator( "\u2588\u2588\u2588\u2588\ue0b0        ", segments.StyleChameleon{ } ),
        segments.NewText( segments.StyleStandard{ segments.UniBrush{ segments.NewColor("#0000ff") }, segments.UniBrush{ segments.NewColor("#fff") } }, "${cmd> ls -la | wc -l}@" ),
        segments.NewHostname( segments.StyleStandard{ segments.UniBrush{ segments.NewColor("#0000f0") }, segments.GradientBrush{ segments.NewColor("#fff"), segments.NewColor("#aaa") } } ),
        segments.NewSeparator( "\ue0b0", segments.StyleChameleon{ } ),
        segments.NewComplexPath(
            segments.StyleStandard{
                segments.UniBrush{ segments.NewColor("#555") },
                segments.GradientBrush{ segments.NewColor("#aaa"), segments.NewColor("#fff") } },
            "\ue0b1",
            segments.UniBrush{ segments.Red } ),
        segments.NewSeparator( "\ue0b0", segments.StyleChameleon{ } ),
    }

    segmentList.Display()
}
