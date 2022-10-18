package server

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/dipankardas011/Efficient-client-server/client"
	"github.com/dipankardas011/Efficient-client-server/payload"
)

func RunServer() {
	fmt.Println("Hello from [[server]]")

	l, err := net.Listen("tcp", payload.SERVER_HOST+":"+payload.SERVER_PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server Listening on ", payload.SERVER_HOST+":"+payload.SERVER_PORT)
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	// server is always open
	for true {
		buffer := make([]byte, 2048)
		mLen, err := c.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		choice := MainDecoder(buffer[:mLen])
		fmt.Println("Message from [[client]]: ", choice)

		switch choice {
		case "GET":
			byteFile, err := os.ReadFile("server/resources/index.html")
			if err != nil {
				panic(err.Error())
			}

			encodedMessage, err := client.MainEncoder(string(byteFile))
			if err != nil {
				panic(err)
			}

			if _, err = c.Write(encodedMessage); err != nil {
				panic(err)
			}

		default:
			var resp string
			resp = fmt.Sprintf(`<!DOCTYPE html>
<html>
<body>
<h1>Response from Server!!</h1>
<table>
<th>Status</th>
<th>Message</th>
<tr>
<td style="color: white; background-color: red">FAILED</td>
<td>Try running GET</td>
</tr>
</table>
<p>Message recieved %v</p>
</body>
</html>`, choice)

			encodedMessage, err := client.MainEncoder(resp)
			if err != nil {
				panic(err)
			}

			if _, err = c.Write(encodedMessage); err != nil {
				panic(err)
			}

		}

		buffer = nil
	}
}
func MainDecoder(jsonRecv []byte) string {
	// wait for the data
	recvData := payload.PayloadDS{}
	err := json.Unmarshal(jsonRecv, &recvData)
	if err != nil {
		panic(err)
	}
	ret := payload.DecodeMessage(recvData)
	return ret
}
