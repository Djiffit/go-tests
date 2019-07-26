package roman

import "testing"

type TestCase struct {
	numeral int
	roman   string
}

func TestRomanNumerals(t *testing.T) {

	for _, test := range tests {
		want := test.roman
		got := ConvertToRoman(test.numeral)

		if got != want {
			t.Errorf("%d: got %q, want %q", test.numeral, got, want)
		}
	}

}

func TestRomanToNumber(t *testing.T) {

	for _, test := range tests {
		want := test.numeral
		got := ConvertToNumber(test.roman)

		if got != want {
			t.Errorf("%s: got %d, want %d", test.roman, got, want)
		}
	}

}

var tests = []TestCase{
	{
		1,
		"I",
	},
	{
		2,
		"II",
	},
	{
		3,
		"III",
	},
	{
		4,
		"IV",
	},
	{
		5,
		"V",
	},
	{
		6,
		"VI",
	},
	{
		7,
		"VII",
	},
	{
		8,
		"VIII",
	},
	{
		9,
		"IX",
	},
	{
		10,
		"X",
	},
	{
		14,
		"XIV",
	},
	{
		18,
		"XVIII",
	},
	{
		20,
		"XX",
	},
	{
		39,
		"XXXIX",
	},
	{
		40,
		"XL",
	},
	{
		47,
		"XLVII",
	},
	{
		49,
		"XLIX",
	},
	{
		50,
		"L",
	},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}
