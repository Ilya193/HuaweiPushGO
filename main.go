package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	content, err := ioutil.ReadFile("json.txt")
	if err != nil {
		panic(err)
	}

	url := "https://push-api.cloud.huawei.com/v1/[APP_ID]/messages:send"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(content)))
	req.Header.Set("Authorization", "Bearer [ACCESS_TOKEN]")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
