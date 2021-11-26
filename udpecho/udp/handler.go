package udp

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
)

const (
	pheader = 0xAA
)

type ptype int32

const (
	// echo expecting a echo reply
	ptypeEcho ptype = iota
	// autoIncrement expecting a auto increment number reply
	ptypeAutoIncrement
)

// UDPHandler for handling udp packets
type UDPHandler struct {
}

// NewHandler create new udp handler
func NewHandler() *UDPHandler {
	return &UDPHandler{}
}

// ServeUDP serve udp packets
func (h *UDPHandler) ServeUDP(conn *net.UDPConn, sourceAddr *net.UDPAddr, packet []byte) {
	data, err := h.parse(packet)
	if err != nil {
		log.Printf("parse err %v from %v", err, sourceAddr)
		return
	}
	if _, err = conn.WriteToUDP(data, sourceAddr); err != nil {
		log.Printf("fail writing udp to %v. err %v", sourceAddr, err)
	}
}

func (h *UDPHandler) parse(packet []byte) ([]byte, error) {
	if len(packet) < 2 {
		return nil, errors.New("unvalid packets")
	}
	if packet[0] != pheader {
		return nil, errors.New("unknow packets")
	}
	switch ptype(packet[1]) {
	case ptypeEcho:
		return packet[2:], nil
	case ptypeAutoIncrement:
		num, err := strconv.Atoi(string(packet[2:]))
		if err == nil {
			return []byte(strconv.Itoa(num + 1)), nil
		}
		return nil, err
	default:
		return nil, fmt.Errorf("unsupport protocol type: %v", ptype(packet[1]))
	}
}
