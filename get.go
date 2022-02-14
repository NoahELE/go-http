package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/logrusorgru/aurora/v3"
	"github.com/urfave/cli/v2"
)

func get(c *cli.Context) error {
	// throw error if no args provided
	if c.Args().Len() == 0 {
		return errors.New("no args provided")
	}

	// process the url
	url := c.Args().First()
	// add the defult 'https://' before url
	if url[:4] != "http" {
		url = "https://" + url
	}

	// get the response
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

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

	// print body
	fmt.Println()
	p := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(p)
		if n != 0 {
			fmt.Print(string(p[:n]))
		} else if err == io.EOF {
			break
		} else {
			return err
		}
	}
	return nil
}
