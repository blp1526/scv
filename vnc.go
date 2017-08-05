package scv

type VNC struct {
	Host     string `json:"Host"`
	Password string `json:"Password"`
	Port     string `json:"Port"`
}

func (vnc *VNC) Path() (path string) {
	path = "vnc://:" + vnc.Password + "@" + vnc.Host + ":" + vnc.Port
	return path
}
