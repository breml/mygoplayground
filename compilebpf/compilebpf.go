package main

import (
	"fmt"
	"github.com/breml/gopacket/pcap"
	"log"
	"time"
)

var (
	device string = "eth0"
	//snapshot_len int32  = 1024
	snapshot_len int32 = 1024
	promiscuous  bool  = false
	err          error
	timeout      time.Duration = 1 * time.Second
	handle       *pcap.Handle
	bpf          []pcap.BPFInstruction
)

func main() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	var filter string = "tcp[tcpflags] & (tcp-syn|tcp-ack) == (tcp-syn|tcp-ack)"
	bpf, err = handle.CompileBPFFilter(filter)

	for _, v := range bpf {
		fmt.Println(v)
	}

}
