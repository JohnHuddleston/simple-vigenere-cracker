package main

import (
	"fmt"
	"strconv"
	"os/exec"
	"strings"
	"sync"
	)

func test_keys(keyset []string, worker_id int, completion_bus chan<- string) {
	str_id := strconv.Itoa(worker_id)
	fmt.Println("Worker " + str_id + " started.")
	var keys_tested = 0
	for _, key := range keyset {
		command := strings.Split("/path/to/executable -d -i /path/to/input/ciphertext.txt -c vigenere -k " + key, " ")
		cmd := exec.Command(command[0], command[1:]...)
		out, err := cmd.CombinedOutput()
		keys_tested += 1

		if err != nil {
			fmt.Println(err)
			fmt.Println(string(out))
		} else if strings.Contains(string(out), "successfully") {
			fmt.Println("KEY HAS BEEN FOUND BY WORKER " + str_id)
			completion_bus <- key
			return
		} else {
			if keys_tested % 1000 == 0 {
				fmt.Println("Worker " + str_id + " has tested " + strconv.Itoa(keys_tested) + " keys.")
			}
		}
	}
}

func main() {
	keys := make([]string, 0)

	var alphabet = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`

	fmt.Println("Generating keys...")

	for _, c := range alphabet {
		for _, d := range alphabet {
			for _, e := range alphabet {
				for _, f := range alphabet {
					for _, g := range alphabet {
						keys = append(keys, string(c) + string(d) + string(e) + string(f) + string(g))
					}
				}
			}
		}
	}

	fmt.Println("Key generation finished.  Starting worker threads to test...")

	batch_size := 300000
	var batches [][]string

	for batch_size < len(keys) {
		keys, batches = keys[batch_size:], append(batches, keys[0:batch_size:batch_size])
	}

	batches = append(batches, keys)

	for i, partition := range batches {
		go test_keys(partition, i, completion_bus)
	}

	the_key := <-completion_bus
	fmt.Println("Program ending, correct key found: " + the_key)

}
