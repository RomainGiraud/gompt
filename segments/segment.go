package segments

import(
    "fmt"
    "log"
    "errors"
    "encoding/json"
)


type Segment interface {
    Print(Context, int)
    GetStyle(context Context, index int) Style
}

type Arguments struct {
    Status int
    ConfigPath string
}

type Context struct {
    Args Arguments
    Segments []Segment
}

func (c *Context) LoadConfig(conf []byte) {
    var config map[string]interface{}
    err := json.Unmarshal(conf, &config)
    if err != nil {
        log.Panic("config file wrong format")
    }

    if val, ok := config["segments"]; ok {
        c.Segments = make([]Segment, 0, 8)
        for _, segment := range val.([]interface{}) {
            segmentConfig := segment.(map[string]interface{})
            if seg, err := LoadSegment(segmentConfig); err == nil {
                c.Segments = append(c.Segments, seg)
            }
        }
    }
}

func (c *Context) Display() {
    if len(c.Segments) == 0 {
        panic("Empty prompt")
    }

    for i, j := 0, 1; i < len(c.Segments); i, j = i+1, j+1 {
        seg := c.Segments[i]
        seg.Print(*c, i)
    }
    fmt.Printf("\n")
}


type SegmentLoader func(map[string]interface{}) Segment

var segmentLoaders = map[string]SegmentLoader{}

func RegisterSegmentLoader(name string, fn SegmentLoader) {
    segmentLoaders[name] = fn
}

func LoadSegment(config map[string]interface{}) (Segment, error) {
    typeName, ok := config["type"].(string);
    if ! ok {
        return nil, errors.New("LoadSegment: key 'type' does not exists in configuration")
    }

    val, ok := segmentLoaders[typeName];
    if ! ok {
        return nil, errors.New("unknown segment type: " + typeName)
    }

    return val(config), nil
}
