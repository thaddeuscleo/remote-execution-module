package executors

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/slc-na/roomnetman-cli/database"
	"github.com/slc-na/roomnetman-cli/models"
)

type MagicPacket [102]byte

func wakeExec(command models.Command) {
	var waitGroup sync.WaitGroup
	for _, computer := range command.Computers {
		waitGroup.Add(1)
		ipaddress := fmt.Sprintf("10.22.%s.%s", computer.Room, computer.Number)
		fmt.Printf("[info] waking up: %s\n", ipaddress)
		mac := database.DB().GetComputerMacAddress(ipaddress)
		go func(mac string, computer models.Computer) {
			if packet, err := createMagicPacket(mac); err == nil {
				packet.send(fmt.Sprintf("10.22.%s.255", computer.Room))
			}
			waitGroup.Done()
		}(mac, computer)
	}
	waitGroup.Wait()
}

func createMagicPacket(macAddr string) (packet MagicPacket, err error) {
	mac, err := net.ParseMAC(macAddr)
	if err != nil {
		return packet, err
	}

	if len(mac) != 6 {
		return packet, errors.New("invalid EUI-48 MAC address")
	}

	copy(packet[0:], []byte{255, 255, 255, 255, 255, 255})
	offset := 6

	for i := 0; i < 16; i++ {
		copy(packet[offset:], mac)
		offset += 6
	}

	return packet, nil
}

func sendMagicPacket(magicPacket MagicPacket, address string) (err error) {
	connection, err := net.Dial("udp", address)
	if err != nil {
		return err
	}
	defer connection.Close()
	_, err = connection.Write(magicPacket[:])
	return err
}

func (magicPacket MagicPacket) send(address string) (err error) {
	return sendMagicPacket(magicPacket, address+":40000")
}