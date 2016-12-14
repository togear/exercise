package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type method struct {
	XMLName xml.Name `xml:method"`
	Name    string   `xml:"name,attr"`
	Id      string   `xml:"sessionid,attr"`
	UrlL    UrlList  `xml:"url_list"`
}

type UrlList struct {
	Value string `xml:",chardata"`
	Urls  []Url  `xml:"url"`
}

type Url struct {
	Key   string `xml:"id,attr"`
	Value string `xml:",chardata"`
}

func main() {

	var item method
	item.Name = "url_purge"
	item.Id = "1234567"

	url_1 := Url{Key: "1", Value: "http://www.test.com/1.index.html"}
	url_2 := Url{Key: "2", Value: "http://www.test.com/2.index.html"}
	url_3 := Url{Key: "3", Value: "http://www.test.com/3.index.html"}
	url_4 := Url{Key: "4", Value: "http://www.test.com/4.index.html"}
	item.UrlL.Urls = make([]Url, 4)
	item.UrlL.Urls[0] = url_1
	item.UrlL.Urls[1] = url_2
	item.UrlL.Urls[2] = url_3
	item.UrlL.Urls[3] = url_4

	xmlP, _ := xml.Marshal(item) // HL

	fmt.Printf("About me (xml):\n %s\n", xmlP)

	xml_header := `<?xml version="1.0" encoding="UTF-8"?>\n`
	var buffer bytes.Buffer
	buffer.WriteString(xml_header)
	buffer.WriteString(string(xmlP))

	fmt.Printf("About me (xml):\n %s\n", buffer.String())

}
