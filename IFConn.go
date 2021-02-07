package conns

// 通讯的接口
type IFConn interface {
	// 发送消息
	Send(msg string) error
	// 读取消息
	Read() (string, error)
	// 关闭
	Close() error
	// 获取IP信息
	GetIPInfo() string
}
