package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/logrusorgru/aurora/v3"
)

func processUrl(url string) string {
	// adds 'https://' before url if no protocol is specified
	if len(url) < 8 || url[:7] != "http://" || url[:8] != "https://" {
		url = "https://" + url
	}
	return url
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
	p := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(p)
		if n != 0 {
			iterator, err := lexer.Tokenise(nil, string(p[:n]))
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
