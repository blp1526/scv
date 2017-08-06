package scv

import "testing"

func TestCLIValidOptSize(t *testing.T) {
	tests := []struct {
		optSize int
		want    bool
	}{
		{
			optSize: 1,
			want:    false,
		},
		{
			optSize: 2,
			want:    true,
		},
	}

	cli := &CLI{}
	for _, test := range tests {
		got := cli.ValidateOptSize(test.optSize)
		if got != test.want {
			t.Fatalf("want: %s, got: %s", test.want, got)
		}
	}
}
