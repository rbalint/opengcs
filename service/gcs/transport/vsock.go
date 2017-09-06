package transport

import (
	"fmt"
	"time"

	"github.com/linuxkit/virtsock/pkg/vsock"
	"github.com/sirupsen/logrus"
)

const (
	vmaddrCidHost = 2
	vmaddrCidAny  = 0xffffffff
)

// VsockTransport is an implementation of Transport which uses vsock
// sockets.
type VsockTransport struct{}

var _ Transport = &VsockTransport{}

// Dial accepts a vsock socket port number as configuration, and
// returns an unconnected VsockConnection struct.
func (t *VsockTransport) Dial(port uint32) (Connection, error) {
	logrus.Infof("vsock Dial port (%d)", port)

	for i := 0; i < 10; i++ {
		conn, err := vsock.Dial(vmaddrCidHost, port)
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		logrus.Infof("vsock Connect port (%d)", port)
		return conn, nil
	}
	return nil, fmt.Errorf("failed connecting the VsockConnection")
}
