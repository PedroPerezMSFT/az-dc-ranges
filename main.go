package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// IPRanges contains a Subnet
type IPRanges struct {
	Subnet string `xml:"Subnet,attr" json:"Subnet"`
}

// Regions contains an array of IPRange(s)
type Regions struct {
	Name    string     `xml:"Name,attr" json:"Name"`
	IPRange []IPRanges `xml:"IpRange" json:"IPRange"`
}

// AzurePublicIPAddresses is the root field in the XML file
type AzurePublicIPAddresses struct {
	Region []Regions `xml:"Region" json:"Region"`
}

func findLink(resp *http.Response) string {

	link := ""

	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			resp.Body.Close()
			return "fail"
		case tt == html.StartTagToken:
			t := z.Token()

			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						if a.Val[len(a.Val)-3:] == "xml" {
							// We found the link, we're done
							link = a.Val
							resp.Body.Close()
							return link
						}
					}
				}
			}
		}
	}
}

func parseXML(link string) AzurePublicIPAddresses {
	var xmlfile []byte

	response, _ := http.Get(link)
	bytestream, _ := ioutil.ReadAll(response.Body)
	xmlfile = bytestream

	var azurepublicipaddresses AzurePublicIPAddresses
	xml.Unmarshal(xmlfile, &azurepublicipaddresses)
	return azurepublicipaddresses
}

func main() {
	url := "https://www.microsoft.com/en-us/download/confirmation.aspx?id=41653" // Global
	//url := "https://www.microsoft.com/en-us/download/confirmation.aspx?id=42064" // China

	resp, _ := http.Get(url)

	link := findLink(resp)

	if link == "fail" {
		fmt.Println("Fatal failure, link not found. Exiting...")
		os.Exit(1)
	}

	azurepublicipaddresses := parseXML(link)

	muhJSON, _ := json.MarshalIndent(azurepublicipaddresses, "", " ")
	fmt.Printf("%+v\n", string(muhJSON))
}
