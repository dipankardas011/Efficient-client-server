package payload

type PayloadDS struct {
	Message  string          `json:"encoded"`
	TableEnc map[byte]string `json:"tableEnc"`
}

func (p PayloadDS) AddInfo(msg string, table map[byte]string) PayloadDS {
	p.Message = msg
	p.TableEnc = table
	return p
}

func (p PayloadDS) GetEncoded() string {
	return p.Message
}

func (p PayloadDS) GetTable() map[byte]string {
	return p.TableEnc
}

type Payload interface {
	AddInfo(string, map[byte]string) PayloadDS
	GetEncoded() string
	GetTable() map[byte]string
}
