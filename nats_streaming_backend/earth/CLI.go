package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func closeCLReader(done chan bool) {
	done <- true
}
func CLReader(done chan bool) {
	defer closeCLReader(done)
	defer closeCLReader(done)
	// defer closeCLReader(done)
	var read string
	time.Sleep(time.Duration(100 * time.Millisecond))
	// fmt.Print("Input \"exit\" to exit CLI \nAdd order ID:")
	// fmt.Fscan(os.Stdin, &read)
	for read != "exit" {

		fmt.Print("Commands:\nexit - shut down\nlist - list of IDs\ninput ID if you want to see order info\nInput:")
		fmt.Fscan(os.Stdin, &read)
		id, _ := strconv.Atoi(read)
		if read == "exit" {

			continue
		}
		if _, ok := cache.Reports[id]; ok {
			fmt.Printf("%s\n", cache.to_json(id))

			// if val == nil

			// fmt.Print(val)
		} else if read == "list" {
			for key := range cache.Reports {
				fmt.Printf("ID: %s\n", key)
			}
		} else {
			err := fmt.Errorf("ERROR: try to get report by ID = %s", read)
			fmt.Errorf("\n%e\n", err)
			fmt.Printf("ERROR: try to get report by ID = %s(there is no such report)\n", read)
		}

	}

}
