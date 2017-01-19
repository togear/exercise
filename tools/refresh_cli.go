/**
usage:
 refresh_cli -f http://host/uri

 purge both the whole file and the slices

output example:

http://mvvideo1.meitudata.com:8090/test/579b6b6b7c4419861.mp4
----------------------------------------
HTTP/1.0 200 OK
Content-Length: 134

<?xml version="1.0" encoding="UTF-8"?>
<url_purge_response sessionid="1484265789">
<url_ret id="0">200</url_ret>
</url_purge_response>

<?xml version="1.0" encoding="UTF-8"?>
<url_purge_response sessionid="1484043157">
<url_ret id="1">200</url_ret>
</url_purge_response>
**/

package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type ResponsFmt struct {
	XMLName xml.Name `xml:"url_purge_response"`
	Id      string   `xml:"sessionid,attr"`
	URLS    []URLSet `xml:"url_ret"`
}

type URLSet struct {
	Key   string `xml:"id,attr"`
	Value string `xml:",chardata"`
}

type purgeRequest struct {
	method    string
	purge_url string
	purge_dir string
	content   string
	response  string
}

var timeoutFlag time.Duration
var purgeUrlFlag string
var versionFlag string

func init() {
	flag.StringVar(&purgeUrlFlag, "f", "", "purge one url")
	flag.StringVar(&versionFlag, "v", "0.0.1", "refresh_cli version")
	flag.DurationVar(&timeoutFlag, "timeout", 3*time.Second, "timeout to refresh")
	flag.DurationVar(&timeoutFlag, "t", 3*time.Second, "timeout to refresh")
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	flag.Parse()

	origin_url := purgeUrlFlag
	if !strings.HasPrefix(origin_url, "http://") {
		panic("lack of purgeurl or url format is not illegal,see help!")
	}

	purge_url, origin_port := addPurgeKeyWordToOriginURL(origin_url)

	resp_flag := sendRequestForPurge(purge_url, origin_port)

	cl := getContentLengthFromHeadRangeRequest(origin_url, origin_port)
	if cl > 0 {
		resp_flag = sendRequestForPurgePieces(purge_url, origin_port, cl)
	}
	outputResultAsXmlFormat(origin_url, resp_flag)
}

/*
In order to Use Nginx purge module to purge, uri must begin with "purge"
example: http://www.test.com/index.html -> http://www.test.com/purge/index.html
*/

func addPurgeKeyWordToOriginURL(origin_url string) (string, int) {
	u, err := url.Parse(origin_url)
	if err != nil {
		panic(err)
	}
	//	fmt.Println("u.Host", u.Host)

	replace_str := u.Host + "/"
	subs_str := u.Host + "/purge/"

	_, port, _ := net.SplitHostPort(u.Host) //split host and port

	port_num := 80
	if port != "" {
		port_num, _ = strconv.Atoi(port)
	}

	//	return "http://www.test.com/purge/index.html"
	return strings.Replace(origin_url, replace_str, subs_str, 1), port_num
}

func sendRequestToCache(pr purgeRequest, port int) int64 {
	client := &http.Client{
		Transport: &http.Transport{
			//set compression off
			DisableCompression: true,
			//set keepalive off
			DisableKeepAlives: false,
			Dial: func(netw, addr string) (net.Conn, error) {
				localaddr := "127.0.0.1" + ":" + strconv.Itoa(port) //Fixme
				rAddr, err := net.ResolveTCPAddr(netw, localaddr)
				if err != nil {
					return nil, err
				}

				conn, err := net.DialTCP(netw, nil, rAddr)
				if err != nil {
					return nil, err
				}

				deadline := time.Now().Add(5 * time.Second)
				conn.SetDeadline(deadline)
				return conn, nil
			},
		},
	}

	var request = new(http.Request)
	var errRequest = errors.New("error Request")

	if pr.method == "GET" {
		request, errRequest = http.NewRequest(pr.method, pr.purge_url, nil)
	} else if pr.method == "POST" {
		var buffer bytes.Buffer
		buffer.WriteString(pr.content)
		request, errRequest = http.NewRequest(pr.method, pr.purge_url, &buffer)
		request.Header.Set("Accept", "*/*")
	} else {
		log.Println("unsupported method ")
		return int64(0)
	}

	if errRequest != nil {
		log.Println("errRequest")
		return int64(0)
	}

	response, _ := client.Do(request)
	if response == nil {
		log.Println("Cache response nil")
		return int64(0)
	}
	defer response.Body.Close() //must close resp.Body

	if pr.method == "POST" {
		data, _ := ioutil.ReadAll(response.Body)
		pr.response = string(data)
		//fmt.Println("response,", pr.response)
	}
	return int64(response.StatusCode)
}

func sendRequestForPurge(purge_url string, port int) bool {

	pr := purgeRequest{}
	pr.purge_url = purge_url
	pr.method = "GET"

	rc := sendRequestToCache(pr, port)

	return rc == 200
}

/**
breif: get content_length from HEAD or GET Request
	   Response Header Content_Range:xxx/content_length

return value content_length
**/

func getContentLengthFromHeadRangeRequest(origin_url string, port int) int64 {

	client := &http.Client{
		Transport: &http.Transport{
			//set compression off
			DisableCompression: true,
			//set keepalive on
			DisableKeepAlives: false,
			Dial: func(netw, addr string) (net.Conn, error) {
				//localaddr := "10.10.100.33" + ":" + strconv.Itoa(port) //Fixme
				localaddr := "127.0.0.1" + ":" + strconv.Itoa(port) //Fixme
				//localaddr := "127,0.0.1:80" //Fixme
				rAddr, err := net.ResolveTCPAddr(netw, localaddr)
				if err != nil {
					return nil, err
				}

				conn, err := net.DialTCP(netw, nil, rAddr)
				if err != nil {
					return nil, err
				}

				deadline := time.Now().Add(5 * time.Second)
				conn.SetDeadline(deadline)
				return conn, nil
			},
		},
	}

	request, err := http.NewRequest("HEAD", origin_url, nil)
	//Fixme
	request.Header.Set("Range", "bytes=0-1048575")

	if err != nil {
		return int64(0)
	}

	response, _ := client.Do(request)
	if response == nil {
		log.Println("HEAD response nil")
		return int64(0)
	}
	defer response.Body.Close() //must close resp.Body

	if response.StatusCode == 206 {
		cr := response.Header.Get("Content-Range")
		cl := cr[strings.Index(cr, "/")+1:]
		sourceSize, _ := strconv.Atoi(cl)
		return int64(sourceSize)
	}

	return int64(0)
}

func sendRequestForPurgePieces(purge_url string, port int, contentLength int64) bool {

	const bytesPrefix1 string = "?slice=bytes="
	const bytesPrefix2 string = "&slice=bytes="
	var flag_purge bool = false

	m_result := make(map[string]int64)

	var sliceSize int64 = 1024 * 1024
	var rangeurl string = ""

	numf := float64(contentLength) / float64(sliceSize)
	num := int64(math.Ceil(float64(numf)))

	pr := purgeRequest{}
	pr.purge_url = purge_url
	pr.method = "GET"

	for i := int64(0); i < num; i++ {
		start := i * sliceSize
		end := (i+1)*sliceSize - 1
		var rangeValue1 string = fmt.Sprintf("%s%d-%d", bytesPrefix1, start, end)
		var rangeValue2 string = fmt.Sprintf("%s%d-%d", bytesPrefix2, start, end)
		if strings.Contains(purge_url, "?") {
			rangeurl = purge_url + rangeValue2
		} else {
			rangeurl = purge_url + rangeValue1
		}
		//		log.Println("slice file size", contentLength, "for", sliceSize, "At Num [", num, "] ", i, rangeurl)
		pr.purge_url = rangeurl

		sc := sendRequestToCache(pr, port)
		m_result[rangeurl] = sc
		if sc == 200 {
			flag_purge = true
		}
		//		log.Println("purge range File return code", sc, "for", pr.purge_url, rangeurl, i, "-", num)
	}

	return flag_purge
}

func outputResultAsXmlFormat(origin_url string, f bool) {
	var item ResponsFmt
	var RC string
	//if any of the urls got 200,the response is 200,else 404
	if f {
		RC = "200"
	} else {
		RC = "404"
	}
	unixnano := strconv.FormatInt(time.Now().Unix(), 10)
	item.Id = unixnano
	item.URLS = make([]URLSet, 1)
	url_pair := URLSet{Key: "0", Value: RC}
	item.URLS[0] = url_pair

	xmlP, _ := xml.Marshal(item) // HL
	xml_header := `<?xml version="1.0" encoding="UTF-8"?>`
	var buffer bytes.Buffer
	buffer.WriteString(xml_header)
	buffer.WriteString(string(xmlP))

	formatBuf := buffer.String()
	formatBuf = strings.Replace(formatBuf, "><", ">\n<", -1)

	buffer_len := len(formatBuf)

	outputString := origin_url +
		"\n-----------------------------------\n" +
		"HTTP/1.0 200 OK\nContent-Length: " + strconv.Itoa(buffer_len) + "\n\n" +
		formatBuf

	fmt.Println(outputString)

}
