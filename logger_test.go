package scv

import (
	"testing"
)

func TestLoggerFormat(t *testing.T) {
	tests := []struct {
		colorName string
		level     string
		want      string
		err       bool
	}{
		{
			colorName: "purple",
			level:     "foo",
			want:      "",
			err:       true,
		},
		{
			colorName: "black",
			level:     "foo",
			want:      "\033[30mfoo: %s\033[m\n",
			err:       false,
		},
	}

	logger := &Logger{}
	for _, test := range tests {
		got, err := logger.Format(test.colorName, test.level)
		if !test.err && err != nil {
			t.Fatalf("colorName: %s", test.colorName)
		}
		if test.err && err == nil {
			t.Fatalf("colorName: %s", test.colorName)
		}
		if got != test.want {
			t.Fatalf("colorName: %s, want: %s, got: %s",
				test.colorName, test.want, got)
		}
	}
}
