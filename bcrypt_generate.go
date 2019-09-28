package main
import (
	"flag"
	"fmt"
	"os"
	"code.google.com/p/go.crypto/bcrypt"
)
 
func main() {
	password := flag.String("password", "", "the password to hash or validate")
	hash := flag.String("hash", "", "the hash to validate the password against")
	cost := flag.Int("cost", 13, "the cost of the bcrypt function to execute")
 
	flag.Parse()
 
	if *hash != "" {
		// do validation
		err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*password))
		if err != nil {
			fmt.Print("Error: Password and hash do not match.")
			os.Exit(1)
		}
		fmt.Print("OK")
		os.Exit(0)
	} else {
		// do encryption
		generated_hash, err := bcrypt.GenerateFromPassword([]byte(*password), *cost)
		if err != nil {
			fmt.Print("Error: Couldn't generate hash.")
			os.Exit(1)
		}
 
		fmt.Printf("%s\n", generated_hash)
		os.Exit(0)
	}
}
