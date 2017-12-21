package main

import(
    "fmt"
    "bytes"
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
            "\uf444" )
    } else {
        exitStatus = segments.NewText(
            segments.StyleStandard{
                segments.UniBrush{ segments.NewColor("#f00") },
                segments.UniBrush{ segments.NewColor("#555") } },
            "\uf444" )
    }

    segmentList := segments.SegmentList {
        exitStatus,
        segments.NewUsername( segments.StyleStandard{ segments.UniBrush{ segments.Green }, segments.UniBrush{ segments.NewColor("#555") } } ),
        segments.NewText( segments.StyleChameleon{ }, "\ue0b0" ),
        segments.NewHostname( segments.StyleStandard{ segments.UniBrush{ segments.Black }, segments.UniBrush{ segments.Blue } } ),
        segments.NewText( segments.StyleChameleon{ }, "\ue0b0" ),
        segments.NewComplexPathPlain(
            segments.StyleStandard{
                segments.UniBrush{ segments.NewColor("#555") },
                segments.GradientBrush{ segments.NewColor("#aaa"), segments.NewColor("#fff") } },
            "\ue0b1", segments.Red, 3, "\u2026"),
        segments.NewText( segments.StyleChameleon{ }, "\ue0b0" ),
        segments.NewBinding(
            segments.NewGit( segments.StyleStandard{ segments.UniBrush{ segments.White }, segments.UniBrush{ segments.NewColor("#444") } } ),
            segments.NewText( segments.StyleChameleon{ }, "\ue0b0" ) ),
        segments.NewText( segments.StyleStandard{ }, " " ),
    }

    var buffer bytes.Buffer
    segmentList.Render(&buffer)
    fmt.Println(buffer.String())
}
