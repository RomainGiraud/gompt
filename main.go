package main

import(
    "os"
    "log"
    "flag"
    "io/ioutil"
    "encoding/json"
    "prompt/segments"
    "prompt/color"
)


type Config struct {
    Segments []Segment `json:"segments"`
    Order []string `json:"prompt"`
}

type Segment struct {
    Name string `json:"name"`
    Type string `json:"type"`
    Options json.RawMessage `json:"options,omitempty"`
    Style string `json:"style,omitempty"`
    StyleOptions json.RawMessage `json:"style-options,omitempty"`
}

type SegmentCreator func(json.RawMessage, color.Style) segments.Segment

var registeredSegmentCreators map[string]SegmentCreator

func RegisterSegmentCreator(name string, fn SegmentCreator) {
    registeredSegmentCreators[name] = fn
}

func CreateSegment(segment Segment) segments.Segment {
    return registeredSegmentCreators[segment.Type](segment.Options, color.NewStyle(segment.Style, segment.StyleOptions))
}

func main() {
    registeredSegmentCreators = make(map[string]SegmentCreator)
    RegisterSegmentCreator("path", segments.NewPath)
    RegisterSegmentCreator("text", segments.NewText)
    RegisterSegmentCreator("separator", segments.NewSeparator)

    var context segments.Context
    flag.IntVar(&context.Args.Status, "s", 0, "exit status")
    flag.StringVar(&context.Args.ConfigPath, "c", "", "config file path")
    flag.Parse()

    var config Config
    configFile, err := os.Open(context.Args.ConfigPath)
    if err != nil {
        log.Panic("wrong config file specified")
    }
    byteValue, _ := ioutil.ReadAll(configFile)
    err = json.Unmarshal(byteValue, &config)
    if err != nil {
        log.Panic("config file wrong format")
    }

    context.Segments = make(map[string]segments.Segment)
    for _, segment := range config.Segments {
        context.Segments[segment.Name] = CreateSegment(segment)
    }
    context.Order = config.Order
    context.Display()
}
