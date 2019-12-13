package bar

import (
	"bytes"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var multiBarTestCase = map[string]struct {
	bars []*Bar
	want string
}{
	"Normality": {
		bars: []*Bar{
			&Bar{
				Title:            "Normality",
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
				Percent:          35,
				PercentColor:     Blue,
				lock:             &sync.Mutex{},
			},
			&Bar{
				Title:            "Percent-Zero",
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
				Percent:          0,
				PercentColor:     Blue,
				lock:             &sync.Mutex{},
			},
			&Bar{
				Title:            "Percent-Negative",
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
				Percent:          -10,
				PercentColor:     Blue,
				lock:             &sync.Mutex{},
			},
		},
		want: "[ 1/3 ] \x1b[31mNormality       \x1b[0m\x1b[33m | \x1b[0m\x1b[32m===================================\x1b[0m\x1b[32m>\x1b[0m\x1b[33m                                                                 \x1b[0m\x1b[33m | \x1b[0m\x1b[34m[  35% ]\x1b[0m\n[ 2/3 ] \x1b[31mPercent-Zero    \x1b[0m\x1b[33m | \x1b[0m\x1b[32m\x1b[0m\x1b[32m>\x1b[0m\x1b[33m                                                                                                    \x1b[0m\x1b[33m | \x1b[0m\x1b[34m[   0% ]\x1b[0m\n[ 3/3 ] \x1b[31mPercent-Negative\x1b[0m\x1b[33m | \x1b[0m\x1b[32m\x1b[0m\x1b[32m>\x1b[0m\x1b[33m                                                                                                    \x1b[0m\x1b[33m | \x1b[0m\x1b[34m[   0% ]\x1b[0m\n",
	},
}

func TestNewMultiBar(t *testing.T) {
	got := NewMultiBar()
	want := &MultiBar{
		bars:   []*Bar{},
		max:    0,
		showed: false,
		lock:   sync.Mutex{},
	}
	opts := []cmp.Option{
		cmp.Comparer(func(x, y *MultiBar) bool {
			return len(x.bars) == 0 && len(y.bars) == 0 && !x.showed && !y.showed && x.max == 0 && y.max == 0
		}),
	}

	if !cmp.Equal(want, got, opts...) {
		t.Errorf("TestNewDefault Failed : want : %#v  got : %#v", want, got)
	}
}

func TestAppend(t *testing.T) {
	for name, testCase := range multiBarTestCase {
		multiBar := NewMultiBar()
		for _, bar := range testCase.bars {
			multiBar.Append(bar)
		}
		if len(multiBar.bars) != len(testCase.bars) {
			t.Errorf("%s Failed : want : %#v  got : %#v", name, len(testCase.bars), len(multiBar.bars))
		}
	}

}

func TestMultiBarShow(t *testing.T) {
	buf := bytes.NewBufferString("")
	for name, testCase := range multiBarTestCase {
		multiBar := NewMultiBar()
		for _, bar := range testCase.bars {
			multiBar.Append(bar)
		}
		multiBar.Show(buf)
		v := buf.String()
		if v != testCase.want {
			t.Errorf("%s Faild : want : %#v got : %#v", name, testCase.want, v)
		}
		buf.Reset()
	}
}

func TestMultiBarShowAgain(t *testing.T) {
	buf := bytes.NewBufferString("")
	for name, testCase := range multiBarTestCase {
		multiBar := NewMultiBar()
		wantbuf := bytes.NewBufferString("")
		for _, bar := range testCase.bars {
			multiBar.Append(bar)
			wantbuf.WriteString("\u001b[1A\u001b[2K")
		}
		wantbuf.WriteString(testCase.want)
		multiBar.Show(buf)
		buf.Reset()
		multiBar.Show(buf)
		v := buf.String()
		if v != wantbuf.String() {
			t.Errorf("%s Faild : want : %#v got : %#v", name, wantbuf.String(), v)
		}
		buf.Reset()
	}
}
