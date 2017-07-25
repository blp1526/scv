package conf

import (
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	absPath, _ := filepath.Abs("../scv.sample.json")
	tests := []struct {
		dir  string
		want Config
	}{
		{
			dir: absPath,
			want: Config{
				AccessToken:       "xxxx",
				AccessTokenSecret: "xxxx",
				Servers: []Server{
					{Name: "ubuntu", ZoneName: "is1a", ID: "1129XXXXXXX1"},
				},
			},
		},
	}

	for _, test := range tests {
		config := Config{}
		_ = config.LoadFile(test.dir)
		if config.AccessToken != test.want.AccessToken {
			t.Fatalf("after LoadFile(%s), AccessToken wants: %s, but got: %s",
				test.dir, test.want.AccessToken, config.AccessToken)
		}
		if config.AccessTokenSecret != test.want.AccessTokenSecret {
			t.Fatalf("after LoadFile(%s), AccessTokenSecret wants: %s, but got: %s",
				test.dir, test.want.AccessTokenSecret, config.AccessTokenSecret)
		}
		if len(config.Servers) != len(test.want.Servers) {
			t.Fatalf("after LoadFile(%s), len(Servers) wants: %s, but got: %s",
				test.dir, len(test.want.Servers), len(config.Servers))
		}
		if config.Servers[0].Name != test.want.Servers[0].Name {
			t.Fatalf("after LoadFile(%s), Servers[0].Name wants: %s, but got: %s",
				test.dir, test.want.Servers[0].Name, config.Servers[0].Name)
		}
		if config.Servers[0].ZoneName != test.want.Servers[0].ZoneName {
			t.Fatalf("after LoadFile(%s), Servers[0].ZoneName wants: %s, but got: %s",
				test.dir, test.want.Servers[0].ZoneName, config.Servers[0].ZoneName)
		}
		if config.Servers[0].ID != test.want.Servers[0].ID {
			t.Fatalf("after LoadFile(%s), Servers[0].ID wants: %s, but got: %s",
				test.dir, test.want.Servers[0].ID, config.Servers[0].ID)
		}
	}
}
