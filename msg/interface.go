package msg

type Msg interface {
	Bytes() []byte
	Marshal() ([]byte, error)
}
