package timeout

import (
	"errors"
	"net"
	"time"
)

var ErrNoDeadlines = errors.New("TimoutConn does not support deadlines")

var _ net.Conn = new(Conn)

type Conn struct {
	net.Conn
	Timeout time.Duration
}

func (t *Conn) update() {
	if t.Timeout == 0 {
		t.Conn.SetDeadline(time.Time{})
		return
	}
	t.Conn.SetDeadline(time.Now().Add(t.Timeout))
}

func (t *Conn) Read(b []byte) (int, error) {
	t.update()
	return t.Read(b)
}

func (t *Conn) Write(b []byte) (int, error) {
	t.update()
	return t.Write(b)
}

func (t *Conn) SetDeadline(_ time.Time) error      { return ErrNoDeadlines }
func (t *Conn) SetWriteDeadline(_ time.Time) error { return ErrNoDeadlines }
func (t *Conn) SetReadDeadline(_ time.Time) error  { return ErrNoDeadlines }
