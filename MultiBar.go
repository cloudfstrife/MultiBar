package bar

import (
	"fmt"
	"io"
)

//MultiBar 从进度条
type MultiBar struct {
	bars   []*Bar
	max    int
	showed bool
}

//Show 输出进度条
func (bars *MultiBar) Show(w io.Writer) {
	// 第一次输出不需要清除，之后的输出，先清除之前的输出，再进行输出
	// \u001b[1A 上移一行
	// \u001b[2K 删除整行
	if bars.showed {
		for i := 0; i < len(bars.bars); i++ {
			fmt.Fprintf(w, "\u001b[1A\u001b[2K")
		}
	} else {
		bars.showed = true
	}

	var i int
	for _, value := range bars.bars {
		i++
		fmt.Fprintf(w, "\u001b[0m[ %d/%d ]", i, len(bars.bars))
		value.Show(w, bars.max, false)
	}
}

//Append 添加一个进度条输出
// 参数说明
// bar 进度条指针
// 注意：title不可重名
func (bars *MultiBar) Append(bar *Bar) {
	if bars.bars == nil {
		bars.bars = []*Bar{}
	}
	bars.bars = append(bars.bars, bar)
	if len(bar.Title) > bars.max {
		bars.max = len(bar.Title)
	}
}
