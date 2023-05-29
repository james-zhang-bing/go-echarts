package opts

type GraphicElementComAttr struct {
	Type     string  `json:"type,omitempty"`
	ID       string  `json:"id,omitempty"`
	X        float64 `json:"x,omitempty"`
	Y        float64 `json:"y,omitempty"`
	Rotation float64 `json:"rotation,omitempty"`
	ScaleX   float64 `json:"scalex,omitempty"`
	ScaleY   float64 `json:"scaley,omitempty"`
	OriginX  float64 `json:"originx,omitempty"`
	OriginY  float64 `json:"originy,omitempty"`
	/*You can specify that all properties have transition animations turned on with `'all'', or you can specify a single property or an array of properties.
	The properties can be:
	Transform related properties:'x', 'y', 'scaleX', 'scaleY', 'rotation', 'originX', 'originY'*/
	Transition        any `json:"transition,omitempty"`
	EnterFrom         any `json:"enterFrom,omitempty"`
	LeaveTo           any `json:"leaveTo,omitempty"`
	EnterAnimation    any `json:"enterAnimation,omitempty"`
	UpdateAnimation   any `json:"updateAnimation,omitempty"`
	LeaveAnimation    any `json:"leaveAnimation,omitempty"`
	KeyFrameAnimation *KeyframeAnimation `json:"keyframeAnimation,omitempty"`
	/*
			Specify how to be positioned in its parent.

		When the element is at the top level, the parent is the container of the chart instance. Otherwise, the parent is a group element.

		Optional values:

		Pixel value: For example, can be a number 30, means 30px.
		Percent value: For example, can be a string '33%', means the final result should be calculated by this value and the height of its parent.
		'center': means position the element in the middle of according to its parent.
		Only one between left and right can work.

		If left or right is specified, positioning attributes in shape (like x, cx) will not work.
	*/
	Left        any                `json:"left,omitempty"`
	Right       any                `json:"right,omitempty"`
	Top         any                `json:"top,omitempty"`
	Bottom      any                `json:"bottom,omitempty"`
	Bounding    string             `json:"bounding,omitempty"`
	Z           int                `json:"z,omitempty"`
	Zlevel      int                `json:"zlevel,omitempty"`
	TextContent *TextContent       `json:"textContent,omitempty"`
	TextConfig  *GraphicTextConfig `json:"textConfig,omitempty"`
	Draggable   bool               `json:"draggable,omitempty"`
	Width       int                `json:"width,omitempty"`
}
type TextContent struct {
	*GraphicTextStyle `json:"style,omitempty"`
	*GraphicElementComAttr
}
type GraphicGroup struct {
	*GraphicElementComAttr
	Children []any
}

func (attr *GraphicElementComAttr) SetText(text string, z int, textStyle ...*GraphicTextStyle) *GraphicElementComAttr {
	var style *GraphicTextStyle
	if len(textStyle) != 0 {
		style = textStyle[0]
		style.Text = text
		attr.TextContent = &TextContent{
			GraphicTextStyle:      style,
			GraphicElementComAttr: &GraphicElementComAttr{Z: z},
		}
		return attr
	}
	style = &GraphicTextStyle{
		Text: text,
	}
	attr.TextContent = &TextContent{
		GraphicTextStyle:      style,
		GraphicElementComAttr: &GraphicElementComAttr{Z: z},
	}
	return attr
}

func (attr *GraphicElementComAttr) SetTextConfig(c *GraphicTextConfig) *GraphicElementComAttr {
	attr.TextConfig = c
	return attr
}
func (attr *GraphicElementComAttr)SetAnimation(anime *KeyframeAnimation)*GraphicElementComAttr{
	attr.KeyFrameAnimation=anime
	return attr
}
type GraphicRectShape struct {
	X          float64   `json:"x,omitempty"`
	Y          float64   `json:"y,omitempty"`
	Width      float64   `json:"width,omitempty"`
	Height     float64   `json:"height,omitempty"`
	R          []float64 `json:"r,omitempty"`
	Transition any       `json:"transition,omitempty"`
}

type GraphicLineShape struct {
	X1         float64 `json:"x1,omitempty"`
	Y1         float64 `json:"y1,omitempty"`
	X2         float64 `json:"x2,omitempty"`
	Y2         float64 `json:"y2,omitempty"`
	Percent    float64 `json:"percent,omitempty"`
	Transition any     `json:"transition,omitempty"`
}

type GraphicStyle struct {
	Fill          string  `json:"fill,omitempty"`
	Stroke        string  `json:"stroke,omitempty"`
	LineWidth     float64 `json:"lineWidth,omitempty"`
	ShadowBlur    float64 `json:"shadowBlur,omitempty"`
	ShadowOffsetX float64 `json:"shadowOffsetX,omitempty"`
	ShadowOffsetY float64 `json:"shadowOffsetY,omitempty"`
	ShadowColor   float64 `json:"shadowColor,omitempty"`
	Opacity       float64 `json:"opacity,omitempty"`
	Transition    any     `json:"transition,omitempty"`
}

type GraphicTextStyle struct {
	GraphicStyle
	Text              string  `json:"text,omitempty"`
	X                 float64 `json:"x,omitempty"`
	Y                 float64 `json:"y,omitempty"`
	Font              string  `json:"font,omitempty"`
	TextAlign         string  `json:"textAlign,omitempty"`
	Width             float64 `json:"width,omitempty"`
	OverFlow          string  `json:"overflow,omitempty"`
	Ellipsis          string  `json:"ellipsis,omitempty"`
	TextVerticalAlign string  `json:"textVerticalAlign,omitempty"`
}

type GraphicTextConfig struct {
	/*
		Position of textContent.
		'left'
		'right'
		'top'
		'bottom'
		'inside'
		'insideLeft'
		'insideRight'
		'insideTop'
		'insideBottom'
		'insideTopLeft'
		'insideTopRight'
		'insideBottomLeft'
		'insideBottomRight'
		or like [12, 33]
		or like ['50%', '50%']
	*/
	Position      string  `json:"position,omitempty"`
	Rotation      float64 `json:"rotation,omitempty"`
	InsideFill    string  `json:"insideFill,omitempty"`
	InsideStroke  string  `json:"insideStroke,omitempty"`
	OutsideFill   string  `json:"outsideFill,omitempty"`
	OutsideStroke string  `json:"outsideStroke,omitempty"`
}

type KeyframeAnimation struct {
	Duration  int         `json:"duration,omitempty"`
	Easing    string      `json:"easing,omitempty"`
	Delay     int         `json:"delay,omitempty"`
	Loop      bool        `json:"loop,omitempty"`
	KeyFrames []*Keyframe `json:"keyframes,omitempty"`
}

type Keyframe struct {
	// Keyframe position. 0 is the first frame, 1 is the last frame
	// The time of keyframe is percent * duration + delay
	Percent float64 `json:"percent"`
	// Easing function from the last keyframe to this keyframe. Optional
	Easing string `json:"easing,omitempty"`
	// Other properties are for configuring the state of target at this keyframe, such as x, y, style, shape, etc.
	Style *GraphicStyle     `json:"style,omitempty"`
	Shape any `json:"shape,omitempty"`
	
}

