package charts

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/stretchr/testify/assert"
)

func TestGraphicJson(t *testing.T) {
	g := &Graphic{
		ID:       "xxx",
		Elements: make([]interface{}, 0),
	}
	e := &opts.GraphicElementComAttr{
		Type: "adaddd",
	}
	g.AddElements(e)
	b, err := json.MarshalIndent(g, "", "\r\t")
	assert.NoError(t, err)
	fmt.Println(string(b))
}
