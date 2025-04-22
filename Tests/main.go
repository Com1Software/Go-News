package main

import (
	"io"
	"net/http"
	"os"
	"runtime"
	"time"

	"fmt"
)

// ----------------------------------------------------------------
func main() {
	fmt.Println("News")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	url := "https://www.cbsnews.com/latest/rss/main"

	fmt.Printf("Current Time : %s\n", GetNews(url, 0))
	fmt.Printf("Channel : %s\n", GetNews(url, 1))
}

func GetNews(url string, opt int) string {
	xdata := ""
	chr := ""
	ton := false
	word := ""
	loc := ""

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching URL: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "HTTP error: %v\n", resp.Status)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response body: %v\n", err)
		os.Exit(1)
	}
	switch {
	case opt == 0:
		t := time.Now()
		formattedTime := t.Format(time.Kitchen)
		xdata = formattedTime
		//------------------------------------------------------------------------ Location
	case opt == 1:
		for x := 1; x < len(body); x++ {
			chr = string(body[x : x+1])
			if chr == "<" {
				ton = true
			}
			if chr == ">" {
				ton = false
				word = ""
			}
			if ton {
				word = word + chr
			}
			if word == "<channel" {
				tmp := ""
				tdata := string(body[x+50 : x+170])
				fmt.Println(tdata)
				for xx := 1; xx < len(tdata)-5; xx++ {
					if tdata[xx:xx+5] == "<title>" {
						xx = xx + 5
						for xx := xx; xx < len(tdata)-5; xx++ {
							chr = string(tdata[xx : xx+1])
							if chr == "<" {
								break
							}
							tmp = tmp + chr
						}

					}

				}
				loc = tmp
			}

		}
		xdata = loc
	}
	return xdata

}
