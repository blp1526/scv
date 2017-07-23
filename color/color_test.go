package color

import "testing"

func TestRed(t *testing.T) {
	tests := []struct {
		text string
		want string
	}{
		{text: "foo", want: "\033[31mfoo\033[m"},
	}

	for _, test := range tests {
		got := Red(test.text)
		if got != test.want {
			t.Fatalf("Red(\"%s\") wants: %s, but got: %s", test.text, test.want, got)
		}
	}
}
