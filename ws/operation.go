package main

import "github.com/labstack/echo"
import "log"

// ImageOperation 图片操作
type ImageOperation struct {
	Base      string       `json:"base"` // 基准图片链接
	SubImages []SubImage   `json:"subImages"`
	Text      TextOption   `json:"text"`
	Texts     []TextOption `json:"texts"`
	Extra     Extra        `json:"extra"`
	Appid     string       `json:"appid"`
}

// Extra 额外操作
type Extra struct {
	Width     int `json:"width"`   // 指定高度
	Height    int `json:"height"`  // 指定高度
	AddWidth  int `json:addWidth`  // base 基础上增加宽度
	AddHeight int `json:addHeight` // base 基础上增加高度
}

// TextOption 文本操作
type TextOption struct {
	Content string `json:"content"`
	Size    int    `json:"size"`
	Left    Pos    `json:"left"` // 水平偏移 居中时，left无效
	Top     Pos    `json:"top"`  // 垂直偏移
	Color   string `json:"color"`
	Middle  bool   `json:"middle"` // 是否居中
}

// SubImage overlay图片定义
type SubImage struct {
	URL     string `json:"url"`     // 图片链接
	Left    Pos    `json:"left"`    // 水平偏移
	Top     Pos    `json:"top"`     // 垂直偏移
	Width   int    `json:"width"`   // 绘制的宽度， url的图片会resize到这个宽度
	Height  int    `json:"height"`  // 绘制的宽度， url的图片会resize到这个宽度
	WithArc bool   `json:"withArc"` // 是否有圆
}

// Pos 位置定义
type Pos struct {
	Relative bool `json:"relative"` // 是否相对于base
	Value    int  `json:"value"`    // 偏移，如果relative = true, 绝对值为 base[width|height] + value
}

// Composite 图片合成
func Composite(c echo.Context) error {
	var operation ImageOperation
	c.Bind(&operation)
	log.Print(operation)
	result, err := DoImageOperation(&operation)
	if err != nil {
		c.JSON(403, err.Error())
		return err
	}
	return c.JSON(200, result)
}

type ResizeOptions struct {
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	ImageBase64 string `json:"imageBase64"`
}

func ResizeImage(c echo.Context) error {
	var resizeOp ResizeOptions
	c.Bind(&resizeOp)
	log.Print(resizeOp)
	ossUrl, error := Resize(resizeOp.ImageBase64, resizeOp.Width, resizeOp.Height)
	if error != nil {
		return error
	}
	c.String(200, ossUrl)
	return nil
}
