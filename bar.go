package bar

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

// Bar 处理进度
type Bar struct {
	//Title 进度条title
	Title string
	//TitleColor title颜色
	TitleColor string
	//Prefix 前缀字符
	Prefix string
	//PrefixColor 前缀字符颜色
	PrefixColor string
	//Postfix 后缀字符
	Postfix string
	//PostfixColor 后缀字符颜色
	PostfixColor string
	//ProcessedFlag 已处理部分字符
	ProcessedFlag rune
	//ProcessedColor 已处理部分字符颜色
	ProcessedColor string
	//ProcessingFlag 处理字符
	ProcessingFlag rune
	//ProcessingColor 处理字符颜色
	ProcessingColor string
	//UnprocessedFlag 未处理部分字符
	UnprocessedFlag rune
	//UnprocessedColor 未处理部分颜色
	UnprocessedColor string
	//Percent 比例
	Percent int
	//PercentColor 比例颜色
	PercentColor string
	//showed 是否已经显示过，如果已经显示过，会做光标上移并清除行的操作
	showed bool
}

const (
	//Black 黑色
	Black = "\u001b[30m"
	//Red 红色
	Red = "\u001b[31m"
	//Green 绿色
	Green = "\u001b[32m"
	//Yellow 黄色
	Yellow = "\u001b[33m"
	//Blue 蓝色
	Blue = "\u001b[34m"
	//Carmine 洋红色
	Carmine = "\u001b[35m"
	//Cyan 青色
	Cyan = "\u001b[36m"
	//White 白色
	White = "\u001b[37m"
	//Reset 重置
	Reset = "\u001b[0m"
)

// NewDefault 创建默认处理
func NewDefault() *Bar {
	return &Bar{
		TitleColor:       Red,
		Prefix:           " | ",
		PrefixColor:      Yellow,
		Postfix:          " | ",
		PostfixColor:     Yellow,
		ProcessedFlag:    '=',
		ProcessedColor:   Green,
		ProcessingFlag:   '>',
		ProcessingColor:  Green,
		UnprocessedFlag:  ' ',
		UnprocessedColor: Yellow,
		PercentColor:     Blue,
	}
}

// Show 输出
// 参数说明
// w io.Writer 	输出目标
// max			title长度
// clean		是否清除上次的输出
func (bar *Bar) Show(w io.Writer, max int, clean bool) {
	if clean && bar.showed {
		fmt.Fprintf(w, "\u001b[1A\u001b[2K\u001b[0m")
	} else {
		bar.showed = true
	}

	if bar.Percent > 100 {
		bar.Percent = 100
	}
	var format = bytes.NewBufferString("")
	// 添加title
	format.WriteString(bar.TitleColor)
	if max < len(bar.Title) {
		max = len(bar.Title)
	}
	format.WriteString("%-" + strconv.Itoa(max) + "s")
	format.WriteString(Reset)

	// 添加处理
	format.WriteString(bar.PrefixColor)
	format.WriteString("%s")
	format.WriteString(Reset)

	//添加已处理部分
	format.WriteString(bar.ProcessedColor)
	var processed = bytes.NewBufferString("")
	for i := 0; i < bar.Percent; i++ {
		processed.WriteRune(bar.ProcessedFlag)
	}
	format.WriteString("%s")
	format.WriteString(Reset)
	// 添加正在处理标识
	format.WriteString(bar.ProcessingColor)
	format.WriteString("%c")
	format.WriteString(Reset)

	// 添加未处理部分
	format.WriteString(bar.UnprocessedColor)
	var unprocessed = bytes.NewBufferString("")
	for i := 0; i < 100-bar.Percent; i++ {
		unprocessed.WriteRune(bar.UnprocessedFlag)
	}
	format.WriteString("%s")
	format.WriteString(Reset)

	// 添加后置处理
	format.WriteString(bar.PostfixColor)
	format.WriteString("%s")
	format.WriteString(Reset)
	// 生成百分比
	format.WriteString(bar.PercentColor)
	format.WriteString("[ %3d%% ]\n")
	format.WriteString(Reset)

	fmt.Fprintf(w, format.String(),
		bar.Title,
		bar.Prefix,
		processed.String(),
		bar.ProcessingFlag,
		unprocessed.String(),
		bar.Postfix,
		bar.Percent,
	)
}