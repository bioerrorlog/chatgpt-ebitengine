package game

import "testing"

func TestInsertLineBreaks(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		charsPerLine   int
		expectedOutput string
	}{
		{"Regular case", "This is a test.", 5, "This \nis a \ntest."},
		{"No line breaks needed", "Short test", 50, "Short test"},
		{"Empty input string", "", 5, ""},
		{"Japanese input", "これは日本語のテストです", 5, "これは日本\n語のテスト\nです"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := insertLineBreaks(test.input, test.charsPerLine)
			if output != test.expectedOutput {
				t.Errorf("insertLineBreaks(%q, %d) = %q; want %q", test.input, test.charsPerLine, output, test.expectedOutput)
			}
		})
	}
}
