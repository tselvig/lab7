package main

import (
	"hscan/hscan"
	"fmt"
	"os"
)

func main() {

	//To test this with other password files youre going to have to hash
	var md5hash = "77f62e3524cd583d698d51fa24fdff4f"
	var sha256hash = "95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced"

	//TODO - Try to find these (you may or may not based on your password lists)
	var drmike1 = "90f2c9c53f66540e67349e0ab83d8cd0" // p@ssword
	var drmike2 = "1c8bfe8f801d79745c4631d09fff36c82aa37fc4cce4fc946683d7b336b63032" // letmein

	// NON CODE - TODO
	// Download and use bigger password file from: https://weakpass.com/wordlist/tiny  (want to push yourself try /small ; to easy? /big )

	//TODO Grab the file to use from the command line instead; look at previous lab (e.g., #3 ) for examples of grabbing info from command line
	if len(os.Args) != 2 {
		fmt.Printf("Error: Incorrect number of args. Usage: ./main [filepath]")
		os.Exit(1);
	}
	var file = os.Args[1]

	fmt.Printf("GuessSingle md5: %s\n", hscan.GuessSingle(md5hash, file))
	fmt.Printf("GuessSingle sha256: %s\n", hscan.GuessSingle(sha256hash, file))
	hscan.GenHashMaps(file)
	hashmapMd5, _ := hscan.GetMD5(md5hash)
	hashmapSha, _ := hscan.GetSHA(sha256hash)
	fmt.Printf("Hashmap sha256: %s\n", hashmapSha)
	fmt.Printf("Hashmap md5: %s\n", hashmapMd5)
	fmt.Printf("GuessSingle drmike1: %s\n", hscan.GuessSingle(drmike1, file))
	fmt.Printf("GuessSingle drmike2: %s\n", hscan.GuessSingle(drmike2, file))

}
