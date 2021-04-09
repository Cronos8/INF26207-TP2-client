package filebyte

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
)

// GetFileByteSignature affiche la signature d'un fichier
func GetFileByteSignature(fileByte []byte) {
	fmt.Printf("File signature : %x\n\n", sha1.Sum(fileByte))
}

// GetByteSignature retourne la signature d'un packet
func GetByteSignature(packet []byte) [20]byte {
	return sha1.Sum(packet)
}

// ConvertBytesToFile convertit une série d'octets en fichier
func ConvertBytesToFile(name string, bytesArr []byte, perm int) {
	err := ioutil.WriteFile(name, bytesArr, os.FileMode(perm))
	if err != nil {
		fmt.Println(err)
		return
	}
}
