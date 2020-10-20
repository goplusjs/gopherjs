// +build teststds

package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	c := exec.Command("go", "build", "-tags", "gopherjsdev")
	c.Dir = filepath.Join(dir, "./..")
	err = c.Run()
	if err != nil {
		log.Fatalln(err)
	}

	gopherjs := filepath.Join(dir, "./..", "gopherjs")

	cmd := exec.Command("go", "list", "std")
	list, err := cmd.Output()
	if err != nil {
		log.Fatal("go list error", err)
	}
	data, err := ioutil.ReadFile("../.std_test_pkg_exclusions")
	if err != nil {
		log.Fatal("read std_test_pkg_exclusions error", err)
	}
	lines := strings.Split(string(data), "\n")
	check_skip := func(v string) bool {
		for _, line := range lines {
			if line == v {
				return true
			}
		}
		return false
	}
	var init bool
	for _, pkg := range strings.Split(string(list), "\n") {
		if pkg == "reflect" {
			init = true
		}
		if pkg == "" || check_skip(pkg) {
			continue
		}
		if !init {
			continue
		}
		testPkg(pkg, gopherjs)
	}
}

func testPkg(pkg string, bin string) {
	log.Println("test", pkg)
	cmd := exec.Command(bin, "test", "-v", "--minify", "--short", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Dir = os.TempDir()
	err := cmd.Run()
	if err != nil {
		log.Fatalf("test pkg %q error %v\n", pkg, err)
	}
}
