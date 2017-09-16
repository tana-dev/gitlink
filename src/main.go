package main

import (
	"fmt"
    "os/exec"
)

func main() {
    out, _ := exec.Command("ls", "-la").Output()
    fmt.Println(string(out))
}
