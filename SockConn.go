package conns

import "net"

type SocketConn struct {
	conn net.Conn
	bufLen int
}

// 发送消息
func (c *SocketConn) Send(msg string) error {
	if len(msg) < 1 {
		return nil
	}
	_, err := c.conn.Write([]byte(msg + "\n\r"))
	return err
}

func (c *SocketConn) Read() (string, error) {
	data := make([]byte, 0, 1024)
	bufData := make([]byte, c.bufLen)

	for {
		n, err := c.conn.Read(bufData)
		if err != nil {
			return "", err
		}
		data = append(data, bufData[:n]...)
		if n != c.bufLen {
			break
		}
	}
	return string(data), nil
}

func (c *SocketConn) Close() error {
	return c.conn.Close()
}

func (c *SocketConn) GetIPInfo() string {
	ipAddr := c.conn.RemoteAddr()
	return ipAddr.Network() + "|" + ipAddr.String()
}
