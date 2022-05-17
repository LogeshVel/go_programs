package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	version := pcap.Version()
	fmt.Println(version)

	// handle, _ := pcap.OpenLive(
	// 	"Wireless LAN adapter WiFi",
	// 	int32(65535),
	// 	false,
	// 	time.Duration(10),
	// )
	fmt.Println("Rann")
	o_handle, _ := pcap.OpenOffline("DNS_pcap.pcapng")
	defer o_handle.Close()

	packetsource := gopacket.NewPacketSource(
		o_handle,
		o_handle.LinkType(),
	)
	pkt, _ := packetsource.NextPacket()
	// fmt.Println(pkt)
	fmt.Printf("pkt type %T", pkt)
	fmt.Println(pkt.Metadata().CaptureInfo)
	fmt.Printf("Type of pkt.Metadata().CaptureInfo is %T\n", pkt.Metadata().CaptureInfo)
	aa := pkt.Metadata().CaptureInfo.Timestamp
	fmt.Println(aa.Unix())
	fmt.Println(aa.UTC())
	fmt.Println(aa.Local())

}
