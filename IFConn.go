package conns

type IFConn interface {
	Send(msg string) error
	Read() (string, error)
	Close() error
	GetIPInfo() string
}
