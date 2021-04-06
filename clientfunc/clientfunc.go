package clientfunc

import (
	"fmt"
	"log"
	"math/rand"
	"net"
)

func SendPaquetWithFiability(fiability float32) bool {
	if rand.Float32() <= fiability {
		return true
	}
	return false
}

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
