package filebyte

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
)

func GetFileByteSignature(fileByte []byte) {
	fmt.Printf("File signature : %x\n\n", sha1.Sum(fileByte))
}

// GetByteSignature
func GetByteSignature(packet []byte) [20]byte {
	return sha1.Sum(packet)
}

func ConvertBytesToFile(name string, bytesArr []byte, perm int) {
	err := ioutil.WriteFile(name, bytesArr, os.FileMode(perm))
	if err != nil {
		fmt.Println(err)
		return
	}
}
