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
                segments.UniBrush{ segments.NewColor("8") },
                segments.UniBrush{ segments.NewColor("8") } },
            "\uf444" )
    } else {
        exitStatus = segments.NewText(
            segments.StyleStandard{
                segments.UniBrush{ segments.NewColor("#f00") },
                segments.UniBrush{ segments.NewColor("8") } },
            "\uf444" )
    }

    segmentList := segments.SegmentList {
        exitStatus,
        segments.NewUsername( segments.StyleStandard{ segments.UniBrush{ segments.White }, segments.UniBrush{ segments.NewColor("8") } } ),
        segments.NewText( segments.StyleChameleon{ }, "\ue0b0" ),
        segments.NewHostname( segments.StyleStandard{ segments.UniBrush{ segments.Black }, segments.UniBrush{ segments.Blue } } ),
        segments.NewText( segments.StyleChameleon{ }, "\ue0b0" ),
        segments.NewComplexPathSplitted(
            segments.StyleStandard{
                segments.UniBrush{ segments.NewColor("#333") },
                segments.GradientBrush{ segments.NewColor("#aaa"), segments.NewColor("#eee") } },
            "\ue0b4", 3, "\u2026"),
        segments.NewBinding(
            segments.NewText( segments.StyleChameleon{ }, "\ue0b0" ),
            segments.NewGit( segments.StyleStandard{ segments.UniBrush{ segments.Black }, segments.UniBrush{ segments.Cyan } } ) ),
        segments.NewText( segments.StyleChameleon{ }, "\ue0b0" ),
        segments.NewText( segments.StyleStandard{ }, " " ),
    }

    var buffer bytes.Buffer
    segmentList.Render(&buffer)
    fmt.Println(buffer.String())
}
