package modules

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

var (
	url_index         = "ttp://localhost:4080/api/index"
	create_json_index = "create_index.json"
	username          = "admin"
	password          = "Complexpass#123"
	url_api_data      = "http://localhost:4080/api/enron_mail/_doc"
)

func readJSON(path string) ([]byte, error) {
	json, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer json.Close()
	content, err2 := os.ReadFile(path)
	if err2 != nil {
		return nil, err2
	}
	return content, nil
}

func postIndex(url string, jsonData []byte, username string, password string) (bool, error) {
	req, err1 := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err1 != nil {
		return false, err1
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err2 := client.Do(req)
	if err2 != nil {
		return false, err2
	}
	defer resp.Body.Close()

	return true, nil
}

func Create_index_zincsearch() {
	jsonData, err := readJSON(create_json_index)
	if err != nil {
		panic(err.Error())
	}

	success, err2 := postIndex(url_index, jsonData, username, password)
	if err2 != nil {
		panic(err2.Error())
	}
	fmt.Println("Create Index enron_mail is ", success)
}
