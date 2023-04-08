package main

import (
	"fmt"
	"io"
	"net/http"
)

func requestGoogleDomain(username string, password string, host string, ip string) string {
	url := fmt.Sprintf("https://%s:%s@domains.google.com/nic/update?hostname=%s&myip=%s", username, password, host, ip)
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	bodyBytes, _ := io.ReadAll(resp.Body)
	return string(bodyBytes)
}

func updateDns(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	fmt.Fprintf(w, requestGoogleDomain(query["username"][0], query["password"][0], query["hostname"][0], query["ip"][0]))
}
func main() {
	http.HandleFunc("/dns/update", updateDns)
	println("start server")
	http.ListenAndServe(":8080", nil)
}
