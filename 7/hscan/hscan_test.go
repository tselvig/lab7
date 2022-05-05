// Optional Todo

package hscan

import (
	"testing"
	"fmt"
	"time"
)

func TestGuessSingle(t *testing.T) {
	got := GuessSingle("77f62e3524cd583d698d51fa24fdff4f", "../main/wordlist.txt") // Currently function returns only number of open ports
	want := "Nickelback4life"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}

func TestGenHashMapsNoRoutines(t *testing.T) {
	// Without go routines
	start := time.Now()
	GenHashMaps("../main/small.txt", false)
	duration := time.Since(start)
	fmt.Printf("Time to generate hash maps without go routines: %f\nTime per password: %f\n", duration.Seconds(), duration.Seconds() / 1667462)

}

func TestGenHashMapsRoutines(t *testing.T) {
	// Without go routines
	start := time.Now()
	GenHashMaps("../main/small.txt", true)
	duration := time.Since(start)
	fmt.Printf("Time to generate hash maps with go routines: %f\nTime per password: %f\n", duration.Seconds(), duration.Seconds() / 1667462)

}
