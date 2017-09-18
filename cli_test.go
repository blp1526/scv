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

func TestCLIRun(t *testing.T) {
	tests := []struct {
		spec    string
		version string
		args    []string
		want    string
		err     bool
	}{
		{
			spec:    "short version",
			version: "v0.0.5-36-g8c7d97d",
			args:    []string{"-V"},
			want:    "scv version 0.0.5, build g8c7d97d",
			err:     false,
		},
		{
			spec:    "long version",
			version: "v0.0.5-36-g8c7d97d",
			args:    []string{"--version"},
			want:    "scv version 0.0.5, build g8c7d97d",
			err:     false,
		},
	}

	for _, tt := range tests {
		cli := &CLI{}
		Version = tt.version
		got, err := cli.Run(tt.args)

		if tt.err && err == nil {
			t.Errorf("spec: %s, tt.err: %v, err: %v", tt.spec, tt.err, err)
		}
		if !tt.err && err != nil {
			t.Errorf("spec: %s, tt.err: %v, err: %v", tt.spec, tt.err, err)
		}
		if got != tt.want {
			t.Errorf("spec: %s, got: %v, tt.want: %v", tt.spec, got, tt.want)
		}
	}
}
