package executors

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/thaddeuscleo/remote-execution-module/database"
	"github.com/thaddeuscleo/remote-execution-module/utils"
)

type MagicPacket [102]byte

func wakeExec(command goExecution) {
	var waitGroup sync.WaitGroup
	for _, computer := range command.Computers {
		waitGroup.Add(1)
		content := fmt.Sprintf("waking up: %s\n", computer)
		utils.LogInfo(content)
		mac := database.DB().GetComputerMacAddress(computer)
		go func(mac string, ipAddress string) {
			if packet, err := createMagicPacket(mac); err == nil {
				broadcastAddress := strings.Split(ipAddress, ".")
				broadcastAddress[3] = "255"
				packet.send(strings.Join(broadcastAddress[:], "."))
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
