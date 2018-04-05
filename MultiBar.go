package bar

import (
	"fmt"
	"io"
)

/*MultiBar 从进度条*/
type MultiBar struct {
	Bars     map[string]*Bar
	MaxTitle int
	showed   bool
}

/*Show 输出进度条*/
func (multiBar *MultiBar) Show(w io.Writer) {
	// 第一次输出不需要清除，之后的输出，先清除之前的输出，再进行输出
	// \u001b[1A 上移一行
	// \u001b[2K 删除整行
	if multiBar.showed {
		for i := 0; i < len(multiBar.Bars); i++ {
			fmt.Fprintf(w, "\u001b[1A\u001b[2K")
		}
	} else {
		multiBar.showed = true
	}

	var i int
	for _, value := range multiBar.Bars {
		i++
		fmt.Fprintf(w, "\u001b[0m[ %d/%d ]", i, len(multiBar.Bars))
		value.Show(w, multiBar.MaxTitle, false)
	}
}

/*Append 添加一个进度条输出*/
func (multiBar *MultiBar) Append(index string, process *Bar) {
	if multiBar.Bars == nil {
		multiBar.Bars = make(map[string]*Bar)
	}
	multiBar.Bars[index] = process

	if len(process.Title) > multiBar.MaxTitle {
		multiBar.MaxTitle = len(process.Title)
	}
}
