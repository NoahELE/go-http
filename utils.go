package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/logrusorgru/aurora/v3"
)

func processUrl(urlStr string) (*url.URL, error) {
	// adds 'https://' before url if no protocol is specified
	b := strings.Builder{}
	if len(urlStr) < 8 || urlStr[:7] != "http://" || urlStr[:8] != "https://" {
		b.WriteString("https://")
	}
	b.WriteString(urlStr)
	return url.ParseRequestURI(b.String())
}

func printResp(resp *http.Response, bodySyntax string) error {
	// print http version and status code
	fmt.Println(aurora.BgBlue(aurora.Black(resp.Proto)),
		aurora.BgGreen(aurora.Black(resp.Status)))

	// print the headers
	for k, v := range resp.Header {
		fmt.Print(aurora.Blue(k + ": "))
		if len(v) == 1 {
			fmt.Println(v[0])
		} else {
			fmt.Println()
			for _, i := range v {
				fmt.Println("\t-", i)
			}
		}
	}
	fmt.Println()

	// print body
	lexer := lexers.Get(bodySyntax)
	style := styles.Get("monokai")
	if style == nil {
		style = styles.Fallback
	}
	formatter := formatters.Get("terminal")
	if formatter == nil {
		formatter = formatters.Fallback
	}
	b := make([]byte, 1024)
	for {
		n, e := resp.Body.Read(b)
		if n != 0 {
			iterator, e := lexer.Tokenise(nil, string(b[:n]))
			if e != nil {
				return e
			}
			e = formatter.Format(os.Stdout, style, iterator)
			if e != nil {
				return e
			}
		} else if e == io.EOF {
			break
		} else {
			return e
		}
	}

	return nil
}
