package filebyte

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GetFileByteSignature(fileByte []byte) {
	fmt.Printf("File signature : %x\n", sha1.Sum(fileByte))
}

func GetByteSignature(packet []byte) {
	log.Printf("Packet signature : %x\n", sha1.Sum(packet))
}

func ConvertBytesToFile(name string, bytesArr []byte, perm int) {
	err := ioutil.WriteFile(name, bytesArr, os.FileMode(perm))
	if err != nil {
		fmt.Println(err)
		return
	}
}
