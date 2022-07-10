package sshtest

import (
	"io"

	"github.com/gliderlabs/ssh"
)

// HoneyPot encapsulates the initialized struct
type HoneyPot struct {
	server *ssh.Server
}

// NewHoneyPot takes in IP address to be used for honeypot
func NewHoneyPot(addr string) *HoneyPot {
	return &HoneyPot{
		server: &ssh.Server{
			Addr: addr,
			Handler: func(s ssh.Session) {
				io.WriteString(s, "Honey pot")
			},
			PasswordHandler: func(ctx ssh.Context, password string) bool {
				return true
			},
		},
	}
}

// ListenAndServe listens on the TCP network address srv.Addr
// and then calls Serve to handle incoming connections
func (h *HoneyPot) ListenAndServe() error {
	return h.server.ListenAndServe()
}

// Close returns any error returned from closing
// the Server's underlying Listener(s).
func (h *HoneyPot) Close() error {
	return h.server.Close()
}

// SetReturnString takes in a string and set it as
// the response from the server
func (h *HoneyPot) SetReturnString(str string) {
	h.server.Handler = func(s ssh.Session) {
		io.WriteString(s, str)
	}
}
