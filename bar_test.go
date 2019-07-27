package bar

import (
	"sync"
	"testing"
)

func TestBar(t *testing.T) {
	testCases := map[string]struct {
		bar    *Bar
		max    int
		result string
	}{
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

	for name, testCase := range testCases {
		v := testCase.bar.Sout(testCase.max)
		if v != testCase.result {
			t.Errorf("%s Faild : want : %s got : %s", name, testCase.result, v)
		}
	}
}
