package scv

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfigLoadFile(t *testing.T) {
	expectedPath, _ := filepath.Abs(".scv.sample.json")
	unexpectedPath, _ := filepath.Abs(".scv.not.sample.json")
	tests := []struct {
		dir  string
		want Config
		err  bool
	}{
		{
			dir: expectedPath,
			want: Config{
				AccessToken:       "xxxx",
				AccessTokenSecret: "xxxx",
				Servers: []Server{
					{Name: "ubuntu", ZoneName: "is1a", ID: "1129XXXXXXX1"},
				},
			},
			err: false,
		},
		{
			dir: unexpectedPath,
			want: Config{
				AccessToken:       "",
				AccessTokenSecret: "",
				Servers:           []Server{},
			},
			err: true,
		},
	}

	for _, test := range tests {
		got := Config{}
		err := got.LoadFile(test.dir)
		if !test.err && err != nil {
			t.Fatalf("dir: %s", test.dir)
		}
		if test.err && err == nil {
			t.Fatalf("dir: %s", test.dir)
		}
		if got.AccessToken != test.want.AccessToken {
			t.Fatalf("dir: %s, AccessToken want: %s, got: %s",
				test.dir, test.want.AccessToken, got.AccessToken)
		}
		if got.AccessTokenSecret != test.want.AccessTokenSecret {
			t.Fatalf("dir: %s, AccessTokenSecret want: %s, got: %s",
				test.dir, test.want.AccessTokenSecret, got.AccessTokenSecret)
		}
		if len(got.Servers) != len(test.want.Servers) {
			t.Fatalf("dir: %s, len(Servers) want: %v, got: %v",
				test.dir, len(test.want.Servers), len(got.Servers))
		}
		if len(got.Servers) == 0 {
			// do nothing
		} else if len(got.Servers) == 1 {
			if got.Servers[0].Name != test.want.Servers[0].Name {
				t.Fatalf("dir: %s, Servers[0].Name want: %s, got: %s",
					test.dir, test.want.Servers[0].Name, got.Servers[0].Name)
			}
			if got.Servers[0].ZoneName != test.want.Servers[0].ZoneName {
				t.Fatalf("dir: %s, Servers[0].ZoneName want: %s, got: %s",
					test.dir, test.want.Servers[0].ZoneName, got.Servers[0].ZoneName)
			}
			if got.Servers[0].ID != test.want.Servers[0].ID {
				t.Fatalf("dir: %s, Servers[0].ID want: %s, got: %s",
					test.dir, test.want.Servers[0].ID, got.Servers[0].ID)
			}
		} else {
			t.Fatalf("len(got.Servers) want 0 or 1, got: %v", len(got.Servers))
		}
	}
}

func TestConfigServerID(t *testing.T) {
	tests := []struct {
		zoneName   string
		serverName string
		want       string
		err        bool
	}{
		{
			zoneName:   "is1b",
			serverName: "centos",
			want:       "",
			err:        true,
		},
		{
			zoneName:   "is1b",
			serverName: "ubuntu",
			want:       "",
			err:        true,
		},
		{
			zoneName:   "is1a",
			serverName: "centos",
			want:       "",
			err:        true,
		},
		{
			zoneName:   "is1a",
			serverName: "ubuntu",
			want:       "1129XXXXXXX1",
			err:        false,
		},
	}

	absPath, _ := filepath.Abs(".scv.sample.json")
	config := Config{}
	_ = config.LoadFile(absPath)

	for _, test := range tests {
		got, err := config.ServerID(test.zoneName, test.serverName)
		if !test.err && err != nil {
			t.Fatalf("zoneName: %s, serverName: %s", test.zoneName, test.serverName)
		}
		if test.err && err == nil {
			t.Fatalf("zoneName: %s, serverName: %s", test.zoneName, test.serverName)
		}
		if got != test.want {
			t.Fatalf("zoneName: %s, serverName: %s, want: %s, got: %s",
				test.zoneName, test.serverName, test.want, got)
		}
	}
}

func TestConfigCreateFile(t *testing.T) {
	tests := []struct {
		dir  string
		want string
		err  bool
	}{
		{
			dir:  "tmp/.scv.test.json",
			want: "tmp/.scv.test.json created",
			err:  false,
		},
	}

	for _, tt := range tests {
		// setup
		err := os.MkdirAll("tmp", 0775)
		if err != nil {
			t.Fatalf("setup failed")
		}

		config := &Config{}
		got, err := config.CreateFile(tt.dir)

		if tt.err && err == nil {
			t.Fatalf("tt.err: %v, err: %v", tt.err, err)
		}
		if !tt.err && err != nil {
			t.Fatalf("tt.err: %v, err: %v", tt.err, err)
		}
		if got != tt.want {
			t.Fatalf("got: %s, tt.want: %s", got, tt.want)
		}

		// tearDown
		err = os.RemoveAll("tmp")
		if err != nil {
			t.Fatalf("tearDown failed")
		}
	}
}
