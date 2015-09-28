package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
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
)

func main() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	bpfInstructions := []pcap.BPFInstruction{
		{0x20, 0, 0, 0xfffff038}, // ld rand
		{0x54, 0, 0, 0x00000004},
		{0x15, 0, 1, 0x00000004},
		{0x06, 0, 0, 0x0000ffff},
		{0x06, 0, 0, 0000000000},
	}

	err = handle.SetBPFInstructionFilter(bpfInstructions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Capturing ~4th packet (randomly).")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Do something with a packet here.
		fmt.Println(packet)
	}

}
