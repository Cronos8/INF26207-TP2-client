package packet

import (
	"encoding/binary"
	"net"
)

// HeaderPacket header of packet
type HeaderPacket struct {
	HeaderIp       net.IP // 16 byte -> 128 octets
	HeaderPort     int32  // 4 byte -> 32 octets
	HeaderNbPacket uint64 // 8 byte -> 64 octets
	// // 28 bytes au total
}

func IsDuplicatePacket(packetMap map[uint64]int, nbPacket uint64) bool {
	if packetMap[nbPacket] == 1 {
		return true
	}
	packetMap[nbPacket]++
	return false
}

func DecapPacket2(packet []byte) (uint64, []byte) {
	buffnbpacket := packet[:8]
	buffbody := packet[8:]
	nbPacket := binary.LittleEndian.Uint64(buffnbpacket)

	return nbPacket, buffbody
}

func DecapPacket(packet []byte) (HeaderPacket, []byte) {

	buffnbpacket := packet[:8]
	nbPacket := binary.LittleEndian.Uint64(buffnbpacket)

	buffipacket := packet[8:24]

	buffportpacket := packet[24:28]
	nbPort := binary.LittleEndian.Uint32(buffportpacket)

	hpacket := HeaderPacket{
		net.IP(buffipacket),
		int32(nbPort),
		nbPacket,
	}
	buffbody := packet[28:]

	return hpacket, buffbody
}
