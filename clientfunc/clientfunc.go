package clientfunc

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
)

// SendPaquetWithFiability simule le pourcentage de fiabilité du client
func SendPaquetWithFiability(fiability float32) bool {
	if rand.Float32() <= fiability {
		return true
	}
	return false
}

// SetFileName récupère le nom du fichier
func SetFileName(conn net.Conn, buff []byte) string {
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if strings.Contains(string(buff[:n]), "FILE ") {
		conn.Write([]byte("PACKET RECEIVE"))
		return strings.Fields(string(buff[:n]))[1]
	}
	return ""
}

// NewServerConnexion établit une connexion avec le serveur
func NewServerConnexion(conn net.Conn) int {

	buffCo := make([]byte, 1000)

	for {
		n, err := conn.Read(buffCo)
		fmt.Println(string(buffCo))
		if err != nil {
			fmt.Printf("ERROR : %s\n", err)
			continue
		}
		if string(buffCo[:n]) == "Serveur - CONNEXION OK" {
			conn.Write([]byte("READY TO RECEIVE"))
			log.Printf("Send: %s", "READY TO RECEIVE")
			break
		}
	}
	return 0
}
