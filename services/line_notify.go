package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func LineNotify(message string) {

	url := "https://notify-api.line.me/api/notify"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("message=%s", message))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("LINE_NOTIFY_TOKEN")))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	if _, err := ioutil.ReadAll(res.Body); err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))
}
