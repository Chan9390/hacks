package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/weppos/publicsuffix-go/publicsuffix"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain, err := publicsuffix.DomainFromListWithOptions(publicsuffix.DefaultList, sc.Text(), &publicsuffix.FindOptions{IgnorePrivate: true})
		if err == nil {
			fmt.Println(domain)
		}
	}
}
