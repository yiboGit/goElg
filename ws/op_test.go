package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextColor(t *testing.T) {
	c := "1,2,3,1"
	assert.Equal(t, color.RGBA{1, 2, 3, 1}, extractColor(c))
	c = "1,2,3"
	assert.Equal(t, color.RGBA{1, 2, 3, 255}, extractColor(c))
}

func TestImageOperation(t *testing.T) {
	imgOp := ImageOperation{
		Base: "http://epj-images.oss-cn-shanghai.aliyuncs.com/activities/card/5-98-fans-1527556989839-card.png",
		SubImages: []SubImage{
			SubImage{
				URL:     "http://epj-images.oss-cn-shanghai.aliyuncs.com/avatar/20180423105613-5-105.jpg",
				Left:    Pos{false, 28},
				Top:     Pos{false, 28},
				Width:   80,
				Height:  80,
				WithArc: true,
			},
			SubImage{
				URL:    "http://epj-images.oss-cn-shanghai.aliyuncs.com/fans/template-2-back.png",
				Left:   Pos{false, 0},
				Top:    Pos{true, 0},
				Width:  0,
				Height: 164,
				// WithArc: true,
			},
		},
		Text: TextOption{
			Left:    Pos{false, 300},
			Top:     Pos{false, 0},
			Content: "小天好啊好啊去去啊去啊啊啊",
			Middle:  true,
		},
		Extra: Extra{
			AddHeight: 164,
		},
		Appid: "wx7cc34ee2a43fe772",
	}
	// buf, _ := json.Marshal(imgOp)
	r, error := DoImageOperation(&imgOp)
	// reqs.JPost("http://localhost:3003", nil, )
	// b, _ := json.Marshal(&imgOp)
	// r, error := http.Post("http://localhost:3002/api/v2/image/composite", "application/json", bytes.NewReader(b))
	assert.Equal(t, nil, error)
	assert.Equal(t, "", r)
}

func TestResize(t *testing.T) {
	// b, err := ioutil.ReadFile("/Users/amyli/Documents/WechatIMG2796.jpeg")
	f, _ := os.Open("/Users/amyli/Documents/WechatIMG2796.jpeg")
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	fmt.Print(content)
	r := base64.StdEncoding.EncodeToString(content)
	assert.Equal(t, "111", r)
}
