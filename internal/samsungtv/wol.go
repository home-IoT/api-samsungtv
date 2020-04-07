package samsungtv

import (
	"fmt"
	"net"

	"github.com/home-IoT/api-samsungtv/internal/log"
	wol "github.com/sabhiram/go-wol"
)

const defaultBroadcastIP = "255.255.255.255"
const defaultUDPPort = "9"

// wakeOnLan sends a magic packet to a mac address
func wakeOnLan(macAddr string) error {

	// can be configured later following the CLI model of https://github.com/sabhiram/go-wol
	bcastAddr := fmt.Sprintf("%s:%s", defaultBroadcastIP, defaultUDPPort)
	udpAddr, err := net.ResolveUDPAddr("udp", bcastAddr)
	if err != nil {
		return err
	}

	// Build the magic packet.
	mp, err := wol.New(macAddr)
	if err != nil {
		return err
	}

	// Grab a stream of bytes to send.
	bs, err := mp.Marshal()
	if err != nil {
		return err
	}

	var localAddr *net.UDPAddr

	// Grab a UDP connection to send our packet of bytes.
	conn, err := net.DialUDP("udp", localAddr, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	log.Debugf("Attempting to send a magic packet to MAC %s\n", macAddr)
	log.Debugf("... Broadcasting to: %s\n", bcastAddr)
	n, err := conn.Write(bs)
	if err == nil && n != 102 {
		err = fmt.Errorf("magic packet sent was %d bytes (expected 102 bytes sent)", n)
	}
	if err != nil {
		return err
	}

	log.Debugf("Magic packet sent successfully to %s\n", macAddr)
	return nil
}
