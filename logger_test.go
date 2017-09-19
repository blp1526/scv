package scv

import (
	"bytes"
	"testing"
)

func TestLoggerDebug(t *testing.T) {
	tests := []struct {
		a       string
		want    int
		err     bool
		out     string
		verbose bool
	}{
		{
			a:       "foo",
			want:    19,
			err:     false,
			out:     "\033[37mdebug: foo\033[m\n",
			verbose: true,
		},
		{
			a:       "foo",
			want:    0,
			err:     false,
			out:     "",
			verbose: false,
		},
	}

	for _, test := range tests {
		outStream := &bytes.Buffer{}
		errStream := &bytes.Buffer{}
		logger := &Logger{Verbose: test.verbose, OutStream: outStream, ErrStream: errStream}
		got, err := logger.Debug(test.a)
		if test.err && err == nil {
			t.Errorf("test.err: %v, err: %v", test.err, err)
		}
		if !test.err && err != nil {
			t.Errorf("test.err: %v, err: %v", test.err, err)
		}
		if got != test.want {
			t.Errorf("got: %d, test.want: %d", got, test.want)
		}
		if outStream.String() != test.out {
			t.Errorf("outStream.String(): %s, test.out: %s", outStream.String(), test.out)
		}
	}
}

func TestLoggerInfo(t *testing.T) {
	tests := []struct {
		a    string
		want int
		err  bool
		out  string
	}{
		{
			a:    "foo",
			want: 4,
			err:  false,
			out:  "foo\n",
		},
	}

	for _, test := range tests {
		outStream := &bytes.Buffer{}
		errStream := &bytes.Buffer{}
		logger := &Logger{OutStream: outStream, ErrStream: errStream}
		got, err := logger.Info(test.a)
		if test.err && err == nil {
			t.Errorf("test.err: %v, err: %v", test.err, err)
		}
		if !test.err && err != nil {
			t.Errorf("test.err: %v, err: %v", test.err, err)
		}
		if got != test.want {
			t.Errorf("got: %d, test.want: %d", got, test.want)
		}
		if outStream.String() != test.out {
			t.Errorf("outStream.String(): %s, test.out: %s", outStream.String(), test.out)
		}
	}
}

func TestLoggerFatal(t *testing.T) {
	tests := []struct {
		a    string
		want int
		err  bool
		out  string
	}{
		{
			a:    "foo",
			want: 19,
			err:  false,
			out:  "\033[31mfatal: foo\033[m\n",
		},
	}

	for _, test := range tests {
		outStream := &bytes.Buffer{}
		errStream := &bytes.Buffer{}
		logger := &Logger{OutStream: outStream, ErrStream: errStream}
		got, err := logger.Fatal(test.a)
		if test.err && err == nil {
			t.Errorf("test.err: %v, err: %v", test.err, err)
		}
		if !test.err && err != nil {
			t.Errorf("test.err: %v, err: %v", test.err, err)
		}
		if got != test.want {
			t.Errorf("got: %d, test.want: %d", got, test.want)
		}
		if errStream.String() != test.out {
			t.Errorf("\nerrStream.String(): %s, test.out: %s", errStream.String(), test.out)
		}
	}
}

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
		if test.err && err == nil {
			t.Errorf("test.err: %v, err %v", test.err, err)
		}
		if !test.err && err != nil {
			t.Errorf("test.err: %v, err %v", test.err, err)
		}
		if got != test.want {
			t.Errorf("colorName: %s, want: %s, got: %s",
				test.colorName, test.want, got)
		}
	}
}
