package main

//go:generate go run ./metadata -pretty -baseURL https://graph.microsoft.com/v1.0 -out metadata/v1.0.xml
//go:generate go run ./metadata -pretty -baseURL https://graph.microsoft.com/beta -out metadata/beta.xml
//go:generate go run . -baseURL https://graph.microsoft.com/v1.0 -in metadata/v1.0.xml -out ../v1.0
//go:generate go run . -baseURL https://graph.microsoft.com/beta -in metadata/beta.xml -out ../beta

import (
	"bufio"
	"flag"
	"log"
	"os"

	_ "github.com/rickb777/date/period"
)

func fix_ModelIP(outputFolder string) {
	file, err := os.OpenFile(outputFolder+"/ModelIP.go", os.O_RDWR, 0660)
	if err != nil {
		log.Fatal(err)
	}
	allText := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		allText = allText + scanner.Text() + "\n"
	}
	appendContent := `
	func filterIPRange(ss []IPRange, v4 bool) (ret []IPRange) {
		odata := "#microsoft.graph.iPv6CidrRange"
		if v4 {
			odata = "#microsoft.graph.iPv4CidrRange"
		}
		for _, s := range ss {
			if s.ODataType == odata {
				ret = append(ret, s)
			}
		}
		return
	}
	
	func NewIPv6Range1() (*IPRange, error) {
		newIPRange := &IPRange{
			ODataType: "#microsoft.graph.IPv6CidrRange",
		}
		return newIPRange, nil
	}
	
	func NewIPv4Range1() (*IPRange, error) {
		newIPRange := &IPRange{
			ODataType: "#microsoft.graph.IPv4CidrRange",
		}
		return newIPRange, nil
	}`
	allText = allText + appendContent
	file.Seek(0, 0)
	if _, err := file.Write([]byte(allText)); err != nil {
		file.Close()
		log.Fatal(err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	g := &Generator{Created: map[string]bool{}, SymTypeMap: map[string]string{}}

	flag.StringVar(&g.BaseURL, "baseURL", "https://graph.microsoft.com/v1.0", "Base URL")
	flag.StringVar(&g.In, "in", "metadata/v1.0.xml", "Input file name")
	flag.StringVar(&g.Out, "out", "out", "Output folder name")
	flag.StringVar(&g.Fmt, "fmt", "goimports", "Formatter")
	flag.Parse()

	err := g.Clean()
	if err != nil {
		log.Fatalf("Failed to clean: %s", err)
	}

	err = g.Generate()
	if err != nil {
		log.Fatalf("Failed to generate: %s", err)
	}

	err = g.Format()
	if err != nil {
		log.Fatalf("Failed to format: %s", err)
	}

	fix_ModelIP(g.Out)
}
