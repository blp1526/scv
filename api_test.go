package scv

import "testing"

func TestAPIURL(t *testing.T) {
	tests := []struct {
		want string
	}{
		{
			want: "https://secure.sakura.ad.jp/cloud/zone/is1a/api/cloud/1.1/server/1129XXXXXXX1/vnc/proxy",
		},
	}

	for _, test := range tests {
		api := &API{ZoneName: "is1a", ServerID: "1129XXXXXXX1"}
		got := api.URL()
		if got != test.want {
			t.Fatalf("want: %s, got: %s", test.want, got)
		}
	}
}
