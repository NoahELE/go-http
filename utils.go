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
	"github.com/logrusorgru/aurora/v4"
)

func parseUrl(urlStr string) (*url.URL, error) {
	// adds 'https://' before url if no protocol is specified
	b := strings.Builder{}
	if len(urlStr) < 8 || urlStr[:7] != "http://" || urlStr[:8] != "https://" {
		b.WriteString("http://")
	}
	b.WriteString(urlStr)
	return url.ParseRequestURI(b.String())
}

func printResp(resp *http.Response, bodySyntax string) error {
	// print http version and status code
	fmt.Println(
		aurora.BgBlue(aurora.Black(resp.Proto)),
		aurora.BgGreen(aurora.Black(resp.Status)),
	)

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
		n, err := resp.Body.Read(b)
		if n != 0 {
			iterator, err := lexer.Tokenise(nil, string(b[:n]))
			if err != nil {
				return err
			}
			err = formatter.Format(os.Stdout, style, iterator)
			if err != nil {
				return err
			}
		} else if err == io.EOF {
			break
		} else {
			return err
		}
	}

	return nil
}
