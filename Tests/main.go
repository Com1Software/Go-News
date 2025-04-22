package main

import (
	"io"
	"math"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"fmt"

	asciistring "github.com/Com1Software/Go-ASCII-String-Package"
)

// ----------------------------------------------------------------
func main() {
	fmt.Println("Weather")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	//url := "https://forecast.weather.gov/MapClick.php?lat=41.5&lon=-81.7&unit=0&lg=english&FcstType=dwml"
	url := "https://forecast.weather.gov/MapClick.php?lat=41.25&lon=-81.44&unit=0&lg=english&FcstType=dwml"

	fmt.Printf("Current Time : %s\n", GetWeather(url, 0))
}


func GetWeather(url string, opt int) string {
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
  }
 return xdata

}
