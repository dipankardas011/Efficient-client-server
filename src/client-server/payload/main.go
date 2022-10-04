package payload

type PayloadDS struct {
	Message  string          `json:"encoded"`
	TableEnc map[byte]string `json:"tableEnc"`
}

const (
	SERVER_HOST = "127.0.0.1"
	SERVER_PORT = "9988"
)

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

func DecodeMessage(encoded PayloadDS) string {
	var reverseHashMap map[string]byte
	reverseHashMap = make(map[string]byte, len(encoded.GetTable()))
	for key, value := range encoded.GetTable() {
		reverseHashMap[value] = key
	}
	i := 0
	j := 1
	decodedMsg := ""
	for i < len(encoded.GetEncoded()) && j <= len(encoded.GetEncoded()) {
		if value, found := reverseHashMap[encoded.GetEncoded()[i:j]]; found {
			decodedMsg += string(value)
			i = j
		}
		j++
	}
	return decodedMsg
}
