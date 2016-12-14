package zoogo

import (
	"io"
	"net"
	"sync"
	"sync/atomic"
)

type Conn struct {
	xid            uint32
	lastZxid       int64
	sessionID      int64
	sessionTimeout int32
}

func (c *Conn) Get() {

}

func (c *Conn) Set() {

}

func (c *Conn) Create() {

}

func (c *Conn) Delete(path string, version int32) error {

}

func (c *Conn) Exists(path string) bool {

}

func (c *Conn) Close() {

}

func (c *Conn) SessionID() int64 {

}

func Connect(servers []string, timeout time.Duration) {

}
