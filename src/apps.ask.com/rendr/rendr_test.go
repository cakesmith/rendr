package rendr_test

import (
	"apps.ask.com/rendr"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCluster(t *testing.T) {
	c := rendr.Cluster{Ip: "54.152.154.104", Port: "4200"}
	url := "http://hp.myway.com/fromdoctopdf/TTAB02/index.html"
	fmt.Println(c.Fetch(url))
}

func BenchmarkCluster(b *testing.B) {
	cluster := rendr.Cluster{Ip: "54.152.154.104", Port: "4200"}
	url := "http://hp.myway.com/fromdoctopdf/TTAB02/index.html"
	b.ResetTimer()
	for i := 0; i < 25 * b.N; i++ {
		_, err := cluster.Fetch(url)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSingle(b *testing.B) {
	single := rendr.Cluster{Ip: "54.197.213.44", Port: "4200"}
	url := "http://hp.myway.com/fromdoctopdf/TTAB02/index.html"
	b.ResetTimer()
	for i := 0; i < 25* b.N; i++ {
		_, err := single.Fetch(url)
		if err != nil {
			panic(err)
		}
	}
}

func singleFetch(url string) (body string, err error) {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	body = string(bodyBytes)
	return
}
