package udp

import (
	"net"
)

// Handler responds to an UDP request.
type Handler interface {
	ServeUDP(conn *net.UDPConn, sourceAddr *net.UDPAddr, data []byte)
}

// ListenAndServe listens on the UDP network address addr and then calls
// Serve to handle requests on incoming connections.
func ListenAndServe(addr string, handler Handler) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	for {
		buf := make([]byte, 4096)
		n, srcAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			// todo
			continue
		}
		go handler.ServeUDP(conn, srcAddr, buf[:n])
	}
}
