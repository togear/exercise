package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	//	client := &http.Client{
	//		Transport: &http.Transport{
	//			Dial: func(netw, addr string) (net.Conn, error) {
	//				//lAddr, err := net.ResolveTCPAddr(netw, ipaddr+":0")
	//				//if err != nil {
	//				//return nil, err
	//				//}
	//				localaddr := "127.0.0.1:80"
	//				rAddr, err := net.ResolveTCPAddr(netw, localaddr)
	//				if err != nil {
	//					return nil, err
	//				}
	//
	//				//conn, err := net.DialTCP(netw, lAddr, rAddr)
	//				conn, err := net.DialTCP(netw, nil, rAddr)
	//				if err != nil {
	//					return nil, err
	//				}
	//
	//				deadline := time.Now().Add(5 * time.Second)
	//				conn.SetDeadline(deadline)
	//				return conn, nil
	//			},
	//		},
	//	}

	url := "http://www.baidu.com"

	request, err := http.NewRequest("GET", url, nil)
	//	reqest.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//	reqest.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	//	reqest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	//	reqest.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
	//	reqest.Header.Set("Cache-Control","max-age=0")
	//	reqest.Header.Set("Connection","keep-alive")

	if err != nil {
		panic(err)
	}

	response, _ := client.Do(request)

	//	stdout := os.Stdout
	//	_, err = io.Copy(stdout, response.Body)

	status := response.StatusCode

	//	if response.StatusCode == 200 {
	//		body, _ := ioutil.ReadAll(response.Body)
	//		bodystr := string(body);
	//		fmt.Println(bodystr)
	//	}

	fmt.Println(status)
}
