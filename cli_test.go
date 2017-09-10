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
		err     bool
	}{
		{
			version: "unexpected value",
			want:    "",
			err:     true,
		},
		{
			version: "v0.0.5-36-g8c7d97d",
			want:    "scv version 0.0.5, build g8c7d97d",
			err:     false,
		},
	}

	cli := &CLI{}
	for _, test := range tests {
		got, err := cli.Versionf(test.version)
		if test.err && err == nil {
			t.Fatalf("test.err: %s, err: %s", test.err, err)
		}
		if !test.err && err != nil {
			t.Fatalf("test.err: %s, err: %s", test.err, err)
		}
		if got != test.want {
			t.Fatalf("want: %s, got: %s", test.want, got)
		}
	}
}
