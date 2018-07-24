package go_get_ip

import (
	"net/http"
	"io/ioutil"
)

//GetIPV4 returns the IP version 4 address string format
func GetIPV4()(string, error) {
	return getIP("https://v4.ident.me/")
}

//GetIPV6 returns the IP version 4 address in a string format
func GetIPV6()(string, error) {
	return getIP("https://v6.ident.me/")
}

func getIP(url string)(string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}