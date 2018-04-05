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
		value.Show(w, multiBar.MaxTitle)

	}

}

/*Append 添加进度*/
func (multiBar *MultiBar) Append(index string, process *Bar) {
	if multiBar.Bars == nil {
		multiBar.Bars = make(map[string]*Bar)
	}
	multiBar.Bars[index] = process

	if len(process.Title) > multiBar.MaxTitle {
		multiBar.MaxTitle = len(process.Title)
	}
}
