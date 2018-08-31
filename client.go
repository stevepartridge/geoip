package geoip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

type Client struct {
	Debug         bool
	IPStackApiKey string
	LastLookup    *IP
}

func New() *Client {
	return &Client{
		IPStackApiKey: os.Getenv("IPSTACK_API_KEY"),
	}
}

// Lookup takes a string based IPv4 address and queries for location and ISP info.
func (self *Client) Lookup(ipQuery string) (IP, error) {

	ip := IP{}
	var err error

	funcs := []func(ip *IP){

		// freegeoip.net
		func(ip *IP) {
			ipStack := IPStackResult{}
			var debug DebugResult
			debug, err = self.call(
				fmt.Sprintf("http://api.ipstack.com/%s?access_key=%s", ipQuery, self.IPStackApiKey),
				&ipStack,
			)

			if err != nil {
				fmt.Println("error", err)
				return
			}

			ip.IP = ipStack.IP
			ip.CountryName = ipStack.CountryName
			ip.CountryCode = ipStack.CountryCode
			ip.RegionName = ipStack.RegionName
			ip.RegionCode = ipStack.RegionCode
			ip.City = ipStack.City
			ip.ZipCode = ipStack.ZipCode
			ip.Timezone = ipStack.Timezone
			ip.Latitude = ipStack.Latitude
			ip.Longitude = ipStack.Longitude
			ip.Timezone = ipStack.Timezone
			ip.Location = ipStack.Location
			ip.Currency = ipStack.Currency
			ip.MetroCode = ipStack.MetroCode

			if self.Debug {
				fmt.Println("debug freegeoip", debug, ipStack)
			}
		},

		// ip-api.com
		func(ip *IP) {
			ipApi := IPApiResult{}
			var debug DebugResult
			debug, err = self.call(
				"http://ip-api.com/json/"+ipQuery,
				&ipApi,
			)

			if err != nil {
				fmt.Println("error", err)
				return
			}

			ip.ISP = ipApi.ISP
			ip.As = ipApi.As
			ip.Org = ipApi.Org

			if self.Debug {
				fmt.Println("debug ip-api", debug, ipApi)
			}
		},
	}

	wg := new(sync.WaitGroup)

	wg.Add(len(funcs))

	for _, fn := range funcs {
		go func(f func(ip *IP)) {
			f(&ip)
			wg.Done()
		}(fn)
	}

	wg.Wait()

	if self.Debug {
		fmt.Println("IP", ip, err)
	}

	return ip, err

}

// Call uses basic (GET) method to make a request to the API
func (self *Client) call(apiUrl string, result interface{}) (DebugResult, error) {
	fmt.Println("call", apiUrl)
	timeout := time.Duration(10 * time.Second)

	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(apiUrl)

	if err != nil {
		return DebugResult{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return DebugResult{}, err
	}

	debug := DebugResult{
		RequestedURL: apiUrl,
		Status:       resp.Status,
		StatusCode:   resp.StatusCode,
		RawResponse:  body,
	}

	fmt.Println("body", string(body))

	err = json.Unmarshal(body, &result)
	if err != nil {
		return debug, err
	}

	return debug, err
}
