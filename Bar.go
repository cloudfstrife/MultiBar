package bar

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

/*Bar 处理进度*/
type Bar struct {
	Title            string
	TitleColor       string
	Prefix           string
	PrefixColor      string
	Postfix          string
	PostfixColor     string
	ProcessedFlag    byte
	ProcessedColor   string
	ProcessingFlag   byte
	ProcessingColor  string
	UnprocessedFlag  byte
	UnprocessedColor string
	Percent          int
	PercentColor     string
}

/*Default 创建默认处理*/
func Default() *Bar {
	return &Bar{
		TitleColor:       "\u001b[1m\u001b[38;5;21m",
		Prefix:           " | ",
		PrefixColor:      "\u001b[1m\u001b[38;5;87m",
		Postfix:          " | ",
		PostfixColor:     "\u001b[1m\u001b[38;5;87m",
		ProcessedFlag:    45,
		ProcessedColor:   "\u001b[1m\u001b[38;5;245m",
		ProcessingFlag:   62,
		ProcessingColor:  "\u001b[1m\u001b[38;5;245m",
		UnprocessedFlag:  32,
		UnprocessedColor: "\u001b[1m\u001b[38;5;245m",
		PercentColor:     "\u001b[1m\u001b[38;5;245m",
	}
}

/*Show 生成输出模板*/
func (process Bar) Show(w io.Writer, maxTitle int) {
	if process.Percent > 100 {
		process.Percent = 100
		process.ProcessingFlag = 32
	}

	var formatBuf bytes.Buffer
	// 添加title
	formatBuf.WriteString(process.TitleColor)
	if maxTitle < len(process.Title) {
		maxTitle = len(process.Title)
	}

	formatBuf.WriteString("%-" + strconv.Itoa(maxTitle) + "s")

	// 重置格式
	formatBuf.WriteString("\u001b[0m")
	// 添加处理
	formatBuf.WriteString(process.PrefixColor)
	formatBuf.WriteString("%s")

	//添加已处理部分
	formatBuf.WriteString("\u001b[0m")
	formatBuf.WriteString(process.ProcessedColor)
	var processedBuf bytes.Buffer
	for i := 0; i < process.Percent; i++ {
		processedBuf.WriteByte(process.ProcessedFlag)
	}
	formatBuf.WriteString("%s")

	// 添加正在处理标识
	formatBuf.WriteString("\u001b[0m")
	formatBuf.WriteString(process.ProcessingColor)
	formatBuf.WriteString("%c")

	// 添加未处理部分
	formatBuf.WriteString("\u001b[0m")
	formatBuf.WriteString(process.UnprocessedColor)
	var unprocessedBuf bytes.Buffer
	for i := 0; i < 100-process.Percent; i++ {
		unprocessedBuf.WriteByte(process.UnprocessedFlag)
	}
	formatBuf.WriteString("%s")

	// 添加后置处理
	formatBuf.WriteString("\u001b[0m")
	formatBuf.WriteString(process.PostfixColor)
	formatBuf.WriteString("%s")

	// 生成百分比
	formatBuf.WriteString("\u001b[0m")
	formatBuf.WriteString(process.PercentColor)
	formatBuf.WriteString("[ %d%% ]\n")

	fmt.Fprintf(w, formatBuf.String(),
		process.Title,
		process.Prefix,
		processedBuf.String(),
		process.ProcessingFlag,
		unprocessedBuf.String(),
		process.Postfix,
		process.Percent,
	)
}
