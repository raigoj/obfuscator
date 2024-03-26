package main

import (
	"fmt"
	"io"
	"net"

	"golang.org/x/crypto/ssh"
)

func (endpoint *Endpoint) String() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}

func clients(client net.Conn, remote net.Conn) {
	defer client.Close()
	chDone := make(chan bool)
	go func() {
		_, err := io.Copy(client, remote)
		Msg{Desc: "Copy err: ", Err: err, Fatal: false}.PrintManager()
		chDone <- true
	}()
	go func() {
		_, err := io.Copy(remote, client)
		Msg{Desc: "Copy err: ", Err: err, Fatal: false}.PrintManager()
		chDone <- true
	}()
	<-chDone
}

func main() {
	go setup()
	sshConfig := &ssh.ClientConfig{
		User:            "shell",
		Auth:            []ssh.AuthMethod{ssh.Password("password")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	con, err := ssh.Dial("tcp", "192.168.1.100:22", sshConfig)
	Msg{Desc: "Dial err: ", Err: err, Fatal: true}.PrintManager()
	listen, err := con.Listen("tcp", "localhost:4242")
	Msg{Desc: "Listen err: ", Err: err, Fatal: true}.PrintManager()
	defer listen.Close()
	for {
		local, err := net.Dial("tcp", "localhost:4242")
		Msg{Desc: "Dial err: ", Err: err, Fatal: true}.PrintManager()
		client, err := listen.Accept()
		Msg{Desc: "Listen err: ", Err: err, Fatal: true}.PrintManager()
		clients(client, local)
	}
}
