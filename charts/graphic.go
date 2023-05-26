package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
)

// https://echarts.apache.org/en/option.html#graphic
type Graphic struct {
	ID       string        `json:"id,omitempty"`
	Elements []interface{} `json:"elements,omitempty"`
}

func (g *GraphicChart) Type() string {
	return ""
}

type GraphicChart struct {
	BaseConfiguration
	BaseActions
}

func (c *GraphicChart) Validate() {
	c.Assets.Validate(c.AssetsHost)
}

func NewGraphicChart() *GraphicChart {
	c := &GraphicChart{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

func (b *BaseConfiguration) AddGraphic(e ...interface{}) {
	if b.Graphics == nil {
		b.Graphics = &Graphic{Elements: make([]interface{}, 0)}
		b.Graphics.Elements = append(b.Graphics.Elements, e...)
		return
	}
	b.Graphics.AddElements(e...)
}

func (g *Graphic) AddElements(e ...interface{}) {
	g.Elements = append(g.Elements, e...)
}


type GraphicGroup struct {
	*opts.GraphicElementComAttr
	Children []any `json:"children,omitempty"`
}


func NewGraphicGroup() *GraphicGroup {
	return &GraphicGroup{
		GraphicElementComAttr: &opts.GraphicElementComAttr{
			Type: "group",
		},
		Children: make([]any, 0),
	}
}
func (g *GraphicGroup) SetCommonAttr(attr *opts.GraphicElementComAttr) *GraphicGroup {
	if attr == nil {
		panic("shape is nil")
	}
	g.GraphicElementComAttr = attr
	g.GraphicElementComAttr.Type = "group"
	return g
}


func (g *GraphicGroup) AddChildren(child ...any) {
	g.Children = append(g.Children, child...)
}

type GraphicRect struct {
	*opts.GraphicElementComAttr
	Shape *opts.GraphicRectShape `json:"shape,omitempty"`
	Style *opts.GraphicStyle     `json:"style,omitempty"`
}
type GraphicRectShape struct {
	X          float64   `json:"x,omitempty"`
	Y          float64   `json:"y,omitempty"`
	Width      float64   `json:"width,omitempty"`
	Height     float64   `json:"height,omitempty"`
	R          []float64 `json:"r,omitempty"`
	Transition any       `json:"transition,omitempty"`
}

func NewGraphicRect() *GraphicRect {
	return &GraphicRect{
		GraphicElementComAttr: &opts.GraphicElementComAttr{
			Type: "rect",
		},
	}
}

func (g *GraphicRect) SetCommonAttr(attr *opts.GraphicElementComAttr) *GraphicRect {
	if attr == nil {
		panic("shape is nil")
	}
	g.GraphicElementComAttr = attr
	g.GraphicElementComAttr.Type = "rect"
	return g
}
func (g *GraphicRect) SetShape(shape *opts.GraphicRectShape) *GraphicRect {
	if shape == nil {
		panic("shape is nil")
	}
	g.Shape = shape
	return g
}

func (g *GraphicRect) SetStyle(style *opts.GraphicStyle) *GraphicRect {
	g.Style = style
	return g
}

type GraphicLine struct {
	*opts.GraphicElementComAttr
	Shape *opts.GraphicLineShape `json:"shape,omitempty"`
	Style *opts.GraphicStyle     `json:"style,omitempty"`
}


func NewGraphicLine() *GraphicLine {
	return &GraphicLine{
		GraphicElementComAttr: &opts.GraphicElementComAttr{
			Type: "line",
		},
	}
}
func (l *GraphicLine) SetZ(z int) *GraphicLine {
	l.Z = z
	return l
}

func (l *GraphicLine) SetStyle(style *opts.GraphicStyle) *GraphicLine {
	l.Style = style
	return l
}

func (l *GraphicLine) SetShape(shape *opts.GraphicLineShape) *GraphicLine {
	if shape == nil {
		panic("shape is nil")
	}
	l.Shape = shape
	return l
}

