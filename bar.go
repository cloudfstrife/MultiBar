package bar

import (
	"bytes"
	"fmt"
	"io"
	"sync"
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
	//lock 输出锁
	lock *sync.Mutex
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
	//FormatTemplate 格式模板
	FormatTemplate = "%s%%-%ds\u001b[0m%s%%s\u001b[0m%s%%s\u001b[0m%s%%c\u001b[0m%s%%s\u001b[0m%s%%s\u001b[0m%s[ %%3d%%%% ]\u001b[0m"
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
		lock:             &sync.Mutex{},
	}
}

//Sout 获取输出
func (bar *Bar) Sout(max int) string {
	if bar.Percent > 100 {
		bar.Percent = 100
	}
	var processed = bytes.NewBufferString("")
	for i := 0; i < bar.Percent; i++ {
		processed.WriteRune(bar.ProcessedFlag)
	}
	// 计算最长Title
	if max < len(bar.Title) {
		max = len(bar.Title)
	}

	//未处理部分
	var unprocessed = bytes.NewBufferString("")
	for i := 0; i < 100-bar.Percent; i++ {
		unprocessed.WriteRune(bar.UnprocessedFlag)
	}
	format := fmt.Sprintf(FormatTemplate, bar.TitleColor, max, bar.PrefixColor, bar.ProcessedColor, bar.ProcessingColor, bar.UnprocessedColor, bar.PostfixColor, bar.PercentColor)
	return fmt.Sprintf(format,
		bar.Title,
		bar.Prefix,
		processed.String(),
		bar.ProcessingFlag,
		unprocessed.String(),
		bar.Postfix,
		bar.Percent,
	)
}

// Show 输出
// 参数说明
// w io.Writer 	输出目标
// max			title长度
// clean		是否清除上次的输出
func (bar *Bar) Show(w io.Writer, max int, clean bool) {
	bar.lock.Lock()
	defer bar.lock.Unlock()
	if clean && bar.showed {
		fmt.Fprintf(w, "\u001b[1A\u001b[2K\u001b[0m")
	} else {
		bar.showed = true
	}
	fmt.Fprintln(w, bar.Sout(max))
}
