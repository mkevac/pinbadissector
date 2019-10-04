package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/gogo/protobuf/proto"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	pinbaProto "github.com/mkevac/gopinba/Pinba"
)

var config struct {
	filename string
	device   string
	filter   string
}

func main() {
	flag.StringVar(&config.filename, "pcapfile", "", "pcap file")
	flag.StringVar(&config.device, "dev", "", "network device")
	flag.StringVar(&config.filename, "filter", "udp port 30006", "bpf filter to use")
	flag.Parse()

	var (
		handle *pcap.Handle
		err    error
	)

	if config.filename != "" {
		handle, err = pcap.OpenOffline(config.filename)
		if err != nil {
			log.Fatalf("error while opening file '%s': %s", config.filename, err)
		}
	} else if config.device != "" {
		handle, err = pcap.OpenLive(config.device, 1600, true, pcap.BlockForever)
		if err != nil {
			log.Fatalf("error while opening device '%s': %s", config.device, err)
		}
	} else {
		log.Fatalf("you should select either file or device")
	}

	if err := handle.SetBPFFilter(config.filter); err != nil {
		log.Fatalf("error while setting bpf filter '%s': %s", config.filter, err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {

		payload := packet.ApplicationLayer().Payload()

		pinbaMsg := pinbaProto.Request{}

		if err := proto.Unmarshal(payload, &pinbaMsg); err != nil {
			log.Printf("error while unmarshalling pinba packet, skipping...: %s", err)
			continue
		}

		jsonBuf, err := json.Marshal(pinbaMsg)
		if err != nil {
			log.Printf("error while converting pinba packet to json, skipping...: %s", err)
			continue
		}

		fmt.Println(string(jsonBuf))
	}
}
