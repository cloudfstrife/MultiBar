package bar

import (
	"bytes"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/google/go-cmp/cmp/cmpopts"
)

var testCases = map[string]struct {
	bar    *Bar
	max    int
	result string
}{
	"Percent-negative": {
		bar: &Bar{
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
		max:    10,
		result: "\x1b[31mPercent-Negative\x1b[0m\x1b[33m | \x1b[0m\x1b[32m\x1b[0m\x1b[32m>\x1b[0m\x1b[33m                                                                                                    \x1b[0m\x1b[33m | \x1b[0m\x1b[34m[   0% ]\x1b[0m",
	},
	"Percent-Zero": {
		bar: &Bar{
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
		max:    10,
		result: "\x1b[31mPercent-Zero\x1b[0m\x1b[33m | \x1b[0m\x1b[32m\x1b[0m\x1b[32m>\x1b[0m\x1b[33m                                                                                                    \x1b[0m\x1b[33m | \x1b[0m\x1b[34m[   0% ]\x1b[0m",
	},
	"Normality": {
		bar: &Bar{
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
		max:    9,
		result: "\x1b[31mNormality\x1b[0m\x1b[33m | \x1b[0m\x1b[32m===================================\x1b[0m\x1b[32m>\x1b[0m\x1b[33m                                                                 \x1b[0m\x1b[33m | \x1b[0m\x1b[34m[  35% ]\x1b[0m",
	},
	"Percent-100": {
		bar: &Bar{
			Title:            "Percent-100",
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
			Percent:          110,
			PercentColor:     Blue,
			lock:             &sync.Mutex{},
		},
		max:    0,
		result: "\x1b[31mPercent-100\x1b[0m\x1b[33m | \x1b[0m\x1b[32m====================================================================================================\x1b[0m\x1b[32m>\x1b[0m\x1b[33m\x1b[0m\x1b[33m | \x1b[0m\x1b[34m[ 100% ]\x1b[0m",
	},
	"Percent-More-Then-100": {
		bar: &Bar{
			Title:            "Percent-More-Then-100",
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
			Percent:          110,
			PercentColor:     Blue,
			lock:             &sync.Mutex{},
		},
		max:    25,
		result: "\x1b[31mPercent-More-Then-100    \x1b[0m\x1b[33m | \x1b[0m\x1b[32m====================================================================================================\x1b[0m\x1b[32m>\x1b[0m\x1b[33m\x1b[0m\x1b[33m | \x1b[0m\x1b[34m[ 100% ]\x1b[0m",
	},
}

func TestSout(t *testing.T) {
	for name, testCase := range testCases {
		v := testCase.bar.Sout(testCase.max)
		if v != testCase.result {
			t.Errorf("%s Faild : want : %#v got : %#v", name, testCase.result, v)
		}
	}
}

func TestNewDefault(t *testing.T) {
	want := Bar{
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

	got := *(NewDefault())

	opts := []cmp.Option{
		cmpopts.IgnoreUnexported(want, got),
	}

	if !cmp.Equal(want, got, opts...) {
		t.Errorf("TestNewDefault Failed : want : %#v  got : %#v", want, got)
	}
}

func TestShow(t *testing.T) {
	buf := bytes.NewBufferString("")
	for name, testCase := range testCases {
		testCase.bar.Show(buf, testCase.max, false)
		v := buf.String()
		if v != testCase.result+"\n" {
			t.Errorf("%s Faild : want : %#v got : %#v", name, testCase.result+"\n", v)
		}
		buf.Reset()
	}
}

func TestShowWithClean(t *testing.T) {
	for name, testCase := range testCases {
		buf := bytes.NewBufferString("")
		testCase.bar.Show(buf, testCase.max, true)
		v := buf.String()
		if v != "\u001b[1A\u001b[2K\u001b[0m"+testCase.result+"\n" {
			t.Errorf("%s Faild : want : %#v got : %#v", name, testCase.result+"\n", v)
		}
	}
}
