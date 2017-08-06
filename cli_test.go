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

func TestCLIVersionFormat(t *testing.T) {
	tests := []struct {
		version string
		want    string
	}{
		{
			version: "0.0.1",
			want:    "scv version 0.0.1",
		},
	}

	cli := &CLI{}
	for _, test := range tests {
		got := cli.VersionFormat(test.version)
		if got != test.want {
			t.Fatalf("want: %s, got: %s", test.want, got)
		}
	}

}
