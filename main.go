package main

import(
    "os"
    "fmt"
    "log"
    "flag"
    //"reflect"
    "io/ioutil"
    "encoding/json"
    "prompt/segments"
    "prompt/color"
    //"prompt/separators"
)


func PrintPrompt(segments map[string]fmt.Stringer, order []string) {
    if len(segments) == 0 || len(order) == 0 {
        panic("Empty prompt")
    }

    for i, j := 0, 1; i < len(order); i, j = i+1, j+1 {
        curr := order[i]
        fmt.Printf("%v", segments[curr])

        /*
        switch s.(type) {
        case segments.Segment:
            fmt.Printf("seg/")
        case separators.Separator:
            fmt.Printf("sep/")
        }
        */

        /*
        color.Set(s.GetFg(), s.GetBg())

        fmt.Printf("%v", s)

        color.Unset()
        if j < len(segments) {
            sn := segments[j]
            color.Set(convertColors[s.GetBg()], sn.GetBg())
        } else {
            color.Set(convertColors[s.GetBg()])
        }
        fmt.Printf("%v", sep)
        */
    }
    fmt.Printf("\n")
}

type Segment struct {
    Name string `json:"name"`
    Type string `json:"type"`
    Options json.RawMessage `json:"options"`
    Style color.StyleConfig `json:"style,omitempty"`
}

type Config struct {
    Segments []Segment `json:"segments"`
    Order []string `json:"prompt"`
}

type SegmentCreator func(json.RawMessage, color.StyleConfig) fmt.Stringer

var registeredSegmentCreators map[string]SegmentCreator

func RegisterSegmentCreator(name string, fn SegmentCreator) {
    registeredSegmentCreators[name] = fn
}

func CreateSegment(segment Segment) fmt.Stringer {
    return registeredSegmentCreators[segment.Type](segment.Options, segment.Style)
}

func main() {
    registeredSegmentCreators = make(map[string]SegmentCreator)
    RegisterSegmentCreator("path", segments.NewPath)
    RegisterSegmentCreator("text", segments.NewText)
    RegisterSegmentCreator("separator", segments.NewSeparator)

    status      := flag.Int("s", 0, "exit status")
    configPath  := flag.String("c", "", "config file path")
    flag.Parse()
    _ = *status

    var config Config
    configFile, err := os.Open(*configPath)
    if err != nil {
        log.Panic("wrong config file specified")
    }
    byteValue, _ := ioutil.ReadAll(configFile)
    err = json.Unmarshal(byteValue, &config)
    if err != nil {
        log.Panic("config file wrong format")
    }

    segs := make(map[string]fmt.Stringer)
    for _, segment := range config.Segments {
        segs[segment.Name] = CreateSegment(segment)
    }
    PrintPrompt(segs, config.Order)

    //fmt.Println(Colorize("toto", Bg24(0, 155, 0), Fg(30)))
    //fmt.Println(fmt.Sprintf(Color.Bg(46).Fg(30)("toto")))

    /*
    sep := separators.Transition{" > "}
    //sep := separators.Transition{" \ue0b0 "}
    seg := Elements{
        //Color{segments.Path{}, []Style{ Bg24(0, 155, 0), Fg(30), Underline }},
        sep,
        segments.ExitStatus{*status, "\u25CF"},
        sep,
        segments.Text{ "rom" },
        sep,
        segments.Username{},
        sep,
        segments.Hostname{},
        sep,
    }
    seg.Print()
    */
}
