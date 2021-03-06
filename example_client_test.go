package easyssh_test

import (
	"log"
	"net"

	"golang.org/x/crypto/ssh"

	"dev.justinjudd.org/justin/easyssh"
)

func ExampleDial() {
	config := &ssh.ClientConfig{
		User: "test",
		Auth: []ssh.AuthMethod{
			ssh.Password("test"),
		},
	}
	conn, err := easyssh.Dial("tcp", "localhost:2022", config)
	if err != nil {
		log.Fatalf("unable to connect: %s", err)
	}
	defer conn.Close()
}

func ExampleClient_LocalForward() {
	config := &ssh.ClientConfig{
		User: "test",
		Auth: []ssh.AuthMethod{
			ssh.Password("test"),
		},
	}
	conn, err := easyssh.Dial("tcp", "localhost:2022", config)
	if err != nil {
		log.Fatalf("unable to connect: %s", err)
	}
	defer conn.Close()

	laddr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")
	raddr, _ := net.ResolveTCPAddr("tcp", "localhost:6060")
	err = conn.LocalForward(laddr, raddr)
	if err != nil {
		log.Fatalf("unable to forward local port: %s", err)
	}

}

func ExampleClient_RemoteForward() {
	config := &ssh.ClientConfig{
		User: "test",
		Auth: []ssh.AuthMethod{
			ssh.Password("test"),
		},
	}
	conn, err := easyssh.Dial("tcp", "localhost:2022", config)
	if err != nil {
		log.Fatalf("unable to connect: %s", err)
	}
	defer conn.Close()

	err = conn.RemoteForward("localhost:8000", "localhost:6060")
	if err != nil {
		log.Fatalf("unable to forward local port: %s", err)
	}

}
