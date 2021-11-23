package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("docker", "ps -a").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	log.Println(out)
}
