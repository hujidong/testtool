package testtool

import (
	"log"
	"net"
	"os"
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
func GetHostName() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Println("os.HostName.err: ", err)
	}
	log.Println(hostname)
}
