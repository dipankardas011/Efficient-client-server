package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/dipankardas011/Efficient-client-server/client"
	"github.com/dipankardas011/Efficient-client-server/payload"
	"github.com/dipankardas011/Efficient-client-server/server"
)

var (
	c   net.Conn
	err error
)

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	fmt.Printf("{\"Source\": \"backend-http\", \"Status\": {\"Port\": \"%v\"}}\n", port)
	return ":" + port
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("[ %s ] Hello from Efficient client-server\n", time.Now())))
}

func getHomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func getMessage(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		panic(fmt.Errorf("invalid method %s", r.Method).Error())
	}
	msg := r.FormValue("message")

	fmt.Printf("msg received %v %T\n\n", msg, msg)

	// client.Main_client get the encoded payload to be sent
	encodedMessage, err := client.MainEncoder(msg)
	if err != nil {
		panic(err)
	}

	_, err = c.Write(encodedMessage)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 2048)
	mLen, err := c.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	str1 := server.MainDecoder(buffer[:mLen])
	fmt.Println("[[server]] response: ", str1)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(str1))
}

func setup() {
	go func() {
		server.RunServer()
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c, err = net.Dial("tcp", payload.SERVER_HOST+":"+payload.SERVER_PORT)
		if err != nil {
			panic(err)
		}
	}()
}

func main() {
	setup()
	http.HandleFunc("/", getHomePage)
	//http.HandleFunc("/status", getStatusPage)
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/message", getMessage)
	http.ListenAndServe(getPort(), nil)
}
