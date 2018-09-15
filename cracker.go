package main

import (
	"fmt"
	"strconv"
	"os/exec"
	"strings"
	"sync"
	)

func test_keys(keyset []string, worker_id int) {
	fmt.Println("Worker " + strconv.Itoa(worker_id) + " started.")
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
			fmt.Println("KEY HAS BEEN FOUND: " + key)
			return
		} else {
			if keys_tested % 100000 == 0 {
				fmt.Println(strconv.Itoa(keys_tested) + " keys tested by worker " + strconv.Itoa(worker_id))
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

	var partition_1 = 3960458
	var partition_2 = 7920917

	var wait_group sync.WaitGroup
	wait_group.Add(3)

	go test_keys(keys[0:partition_1], 0)
	go test_keys(keys[partition_1+1:partition_2], 1)
	go test_keys(keys[partition_2+1:len(keys)], 2)

	wait_group.Wait()

}
