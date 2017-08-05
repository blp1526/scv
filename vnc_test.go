package scv

import "testing"

func TestVNCPath(t *testing.T) {
	tests := []struct {
		want string
	}{
		{
			want: "vnc://:XXXXXXXX@example.com:XXXXX",
		},
	}

	for _, test := range tests {
		vnc := &VNC{Host: "example.com", Password: "XXXXXXXX", Port: "XXXXX"}
		got := vnc.Path()
		if got != test.want {
			t.Fatalf("want: %s, got: %s", test.want, got)
		}
	}
}
