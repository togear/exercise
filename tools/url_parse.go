package main

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

func main() {
	s := "http:/v.youku.com/test/dfsdfsdffffffffffffffffff61.mp4?auth=b31457284194303"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	fmt.Println("Host", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host:", host)
	fmt.Println("port:", port)
	fmt.Println("path:", u.Path)
	fmt.Println("Fragment", u.Fragment)
	fmt.Println("RawQuery", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["slice"])

	fmt.Println("HasPrefix: ", strings.HasPrefix(u.Path, "/purge"))

	if strings.HasPrefix(u.Path, "/purge") {

		fmt.Println("Url Parsing, begin with /purge")
	} else {
		fmt.Println("Error Url Parsing, no begin with /purge")
	}

	fmt.Println("path:", u.Path)

}
