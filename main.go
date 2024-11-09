package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/suikurix/file-encrypter/filecrypt"
	"goland.org/x/term"
)

func main(){
	if len (os.Args) < 2{
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptionHandle()
	default:
		fmt.Println("Run encrypt to encrypt a file, and decrypt to decrypt a file")
		os.Exit(0)
	}
}

func printHelp(){
	fmt.Println("File Encryptor")
	fmt.Println("A file encrypter right in your terminal")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\\go run . encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypts a file given a password")
	fmt.Println("\t decrypt\tTries to decrypt a file using a password")
	fmt.Println("\t decrypt\tTries to decrypt a file using a password")
	fmt.Println("\t help\t\tDisplays help text")
	fmt.Println("")



}

func encryptHandle(){
	if len(os.Args) < 3(
		println("missing the path to the file. For more info, run . help")
		os.Exit(0)
	)
	file := os.Args[2]
	if !validateFile(file){
		panic("File not found")
	}
	password := getPassword()
	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n file sucessful protected")
}

func decryptionHandle(){
	if len(os.Args) < 3{
		println("missing the path to the file. For more info, run . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file){
		panic("File not found")
	}
	fmt.Print("Enter password:")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nDecrypting")
	filecrypt.Decrypt(file, password)
	fmt.Println("\n file sucessful decrypted")
}
func getPassword() []byte{

	fmt.Print("Enter password")
	password, _ := term.ReadPassword(0)
	fmt.Print("\nConfirm Password: ")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2){
		fmt.Print("\nPasswords do not mach. Please try again")
		return getPassword()
	}
	return password

}

func validatePassword(password1 []byte, password2 []byte) bool{
	if !bytes.Wqual(password1, password2){
		return fals
	}
	return tru

}

func validateFile(file string) bool{
	if _, err := os.Stat(file); os.IsNotExist(err){
		return fals
	}
	return tru
}