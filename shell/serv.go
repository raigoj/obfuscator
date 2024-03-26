package main

import (
	"io"
	"net"
	"os/exec"
	"syscall"

	"github.com/kr/pty"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func handleServerConn(chans <-chan ssh.NewChannel) {
	for newChan := range chans {
		if newChan.ChannelType() != "session" {
			_ = newChan.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}
		ch, reqs, err := newChan.Accept()
		Msg{Desc: "Accept err: ", Err: err, Fatal: false}.PrintManager()
		if err != nil {
			continue
		}
		go func(in <-chan *ssh.Request) {
			defer func() {
				_ = ch.Close()
			}()
			npty, ntty, err := pty.Open()
			Msg{Desc: "Pty err: ", Err: err, Fatal: false}.PrintManager()
			cmd := exec.Command("bash")
			cmd.Stdout = ntty
			cmd.Stdin = ntty
			cmd.Stderr = ntty
			cmd.SysProcAttr = &syscall.SysProcAttr{
				Setctty: true,
				Setsid:  true,
			}
			for req := range reqs {
				err = cmd.Start()
				Msg{Desc: "Bash err: ", Err: err, Fatal: false}.PrintManager()
				req.Reply(true, nil)
				go func() {
					_, _ = io.Copy(npty, ch)
				}()
				_, _ = io.Copy(ch, npty)
				ch.Close()
				Msg{Desc: "Bash err: ", Err: err, Fatal: false}.PrintManager()
				Msg{Desc: "Closed sess", Err: nil, Fatal: false}.PrintManager()
				return
			}
		}(reqs)
	}
}

func listener(config *ssh.ServerConfig) {
	listen, err := net.Listen("tcp", "0.0.0.0:4242")
	Msg{Desc: "Listen err: ", Err: err, Fatal: true}.PrintManager()
	for {
		con, err := listen.Accept()
		Msg{Desc: "Accept err: ", Err: err, Fatal: false}.PrintManager()
		if err == nil {
			go func() {
				sConn, chans, reqs, err := ssh.NewServerConn(con, config)
				if err != nil {
					if err == io.EOF || errors.Is(err, syscall.ECONNRESET) {
						Msg{Desc: "Handshake terminated: ", Err: err, Fatal: false}.PrintManager()
					} else {
						Msg{Desc: "Handshake err: ", Err: err, Fatal: false}.PrintManager()
					}
					return
				}
				Msg{Desc: "connecion: " + sConn.RemoteAddr().String(), Err: nil, Fatal: false}.PrintManager()
				go ssh.DiscardRequests(reqs)
				go handleServerConn(chans)
			}()
		}
	}
}

func setup() {
	conf := &ssh.ServerConfig{
		NoClientAuth: true,
	}
	key, err := ssh.ParsePrivateKey([]byte(KEY))
	Msg{Desc: "Parse err: ", Err: err, Fatal: true}.PrintManager()
	conf.AddHostKey(key)
	go listener(conf)
}
