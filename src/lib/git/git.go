package git

import (
	"fmt"
	"os/exec"
)

func Clone(url string) {

	out, _ := exec.Command("git", "clone", url).Output()
	fmt.Println(string(out))
}

func Branch() {

	out, _ := exec.Command("git", "branch", "-r").Output()
	fmt.Println(string(out))
}

func Diff() {

	out, _ := exec.Command("git", "diff").Output()
	fmt.Println(string(out))
}
