package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

var arg string

func main() {

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		arg = sc.Text()
		fmt.Println(url.QueryEscape(arg))
	}

}
