package rendr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

type Cluster struct {
	Ip, Port string
}

func (c Cluster) url() string {
	return fmt.Sprintf("http://%s:%s", c.Ip, c.Port)
}

func (c Cluster) Fetch(url string) (body string, err error) {

	var urlStr = []byte(url)
	req, err := http.NewRequest("POST", c.url(), bytes.NewBuffer(urlStr))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	body = string(bodyBytes)
	return
}
