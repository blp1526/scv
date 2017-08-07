package scv

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func TestAPIGET(t *testing.T) {
	mockHandler := func() func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			username, password, ok := r.BasicAuth()
			if !ok || !(username == "xxxx" && password == "xxxx") {
				w.WriteHeader(401)
				w.Write([]byte("401 Unauthorized"))
			} else {
				w.WriteHeader(201)
				w.Write([]byte("foo"))
			}
		}
	}
	mockServer := httptest.NewServer(http.HandlerFunc(mockHandler()))
	defer mockServer.Close()

	tests := []struct {
		url               string
		accessToken       string
		accessTokenSecret string
		want              []byte
		err               bool
	}{
		{
			url:               mockServer.URL,
			accessToken:       "",
			accessTokenSecret: "",
			want:              []byte(""),
			err:               true,
		},
		{
			url:               mockServer.URL,
			accessToken:       "xxxx",
			accessTokenSecret: "yyyy",
			want:              []byte(""),
			err:               true,
		},
		{
			url:               mockServer.URL,
			accessToken:       "yyyy",
			accessTokenSecret: "xxxx",
			want:              []byte(""),
			err:               true,
		},
		{
			url:               mockServer.URL,
			accessToken:       "xxxx",
			accessTokenSecret: "xxxx",
			want:              []byte("foo"),
			err:               false,
		},
	}

	api := &API{}
	for _, test := range tests {
		got, err := api.GET(test.url, test.accessToken, test.accessTokenSecret)
		if test.err && err == nil {
			t.Fatalf("test.err: %s, err: %s", test.err, err)
		}
		if !test.err && err != nil {
			t.Fatalf("test.err: %s, err: %s", test.err, err)
		}
		if string(test.want) != string(got) {
			t.Fatalf("test.want: %s, got: %s", string(test.want), string(got))
		}
	}
}
