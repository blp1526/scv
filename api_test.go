package scv

import "testing"

func TestAPIURL(t *testing.T) {
	tests := []struct {
		zoneName string
		serverID string
		want     string
	}{
		{
			zoneName: "is1a",
			serverID: "1129XXXXXXX1",
			want:     "https://secure.sakura.ad.jp/cloud/zone/is1a/api/cloud/1.1/server/1129XXXXXXX1/vnc/proxy",
		},
	}

	for _, test := range tests {
		api := &API{}
		got := api.URL(test.zoneName, test.serverID)
		if got != test.want {
			t.Fatalf("want: %s, got: %s", test.want, got)
		}
	}
}
