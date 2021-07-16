package TestTool

import (
	"log"
	"net"
)

func GetIp() {

	udpconn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Printf("net.Dial.err: %v", err)
	}
	localip := udpconn.LocalAddr()

	host, _, err := net.SplitHostPort(localip.String())
	if err != nil {
		log.Println("net.SplitHostPort.err：", err)
	}
	log.Println(host)
}
