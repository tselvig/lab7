package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
)

//==========================================================================\\

var shalookup map[string]string
var md5lookup map[string]string

func GuessSingle(sourceHash string, filename string) string{

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()

		// TODO - From the length of the hash you should know which one of these to check ...
		// add a check and logicial structure
		if len(sourceHash) == 32{
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				return password
			}
		}
		if len(sourceHash) > 32{
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				return password
			}
		}
		
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return ""
}

func GenHashMaps(filename string, useGoRoutine bool) {
	shalookup = make(map[string]string)
	md5lookup = make(map[string]string)
	//TODO
	//itterate through a file (look in the guessSingle function above)
	//rather than check for equality add each hash:passwd entry to a map SHA and MD5 where the key = hash and the value = password
	//TODO at the very least use go subroutines to generate the sha and md5 hashes at the same time
	//OPTIONAL -- Can you use workers to make this even faster
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var md5Hash string
	var shaHash string
	for scanner.Scan() {
		password := scanner.Text()
		if(useGoRoutine) {
			c := make(chan string)
			c2 := make(chan string)
			
			go mapMd5(password, c)
			go mapSha(password, c2)
			md5Hash = <- c
			shaHash = <- c2
		} else {
			md5Hash = fmt.Sprintf("%x", md5.Sum([]byte(password)))
			shaHash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		}
		md5lookup[md5Hash] = password
		shalookup[shaHash] = password
	}

	//TODO create a test in hscan_test.go so that you can time the performance of your implementation
	//Test and record the time it takes to scan to generate these Maps
	// 1. With and without using go subroutines
	// 2. Compute the time per password (hint the number of passwords for each file is listed on the site...)
}

func mapMd5(password string, c chan string) {
	hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	c <- hash
}

func mapSha(password string, c chan string) {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	c <- hash
}

func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	if ok {
		return password, nil

	} else {

		return "", errors.New("password does not exist")

	}
}

//TODO
func GetMD5(hash string) (string, error) {
	password, ok := md5lookup[hash]
	if ok {
		return password, nil
	} else {
		return "", errors.New("password does not exist")
	}
}
