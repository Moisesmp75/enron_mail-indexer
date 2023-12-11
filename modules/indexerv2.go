package modules

import (
	"bytes"
	"encoding/json"
	"enron_mail-indexer/models"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func readFile2(path string, mail *models.Mail, channel chan models.Mail) {
	fmt.Println("Leyendo ", path)
	file, err1 := os.Open(path)
	if err1 != nil {
		panic(err1.Error())
	}
	file.Close()
	content, err2 := os.ReadFile(path)
	if err2 != nil {
		panic(err2.Error())
	}
	strContent := string(content)
	values := strings.Split(strContent, "\n")
	for _, value := range values {
		data := strings.SplitN(value, ": ", 2)
		if len(mail.To) > 0 && mail.Subject == "" && len(data) == 1 {
			mails := strings.Split(data[0], ",")
			mail.To = append(mail.To, mails...)
		}
		if len(data) == 2 && mail.X_filename == "" {
			switch data[0] {
			case "Message-ID":
				mail.Message_ID = data[1]
			case "Date":
				mail.Date = data[1]
			case "From":
				mail.From = data[1]
			case "To":
				mail.To = strings.Split(data[1], ",")
			case "Subject":
				mail.Subject = data[1]
			case "Cc":
				mail.Cc = strings.Split(data[1], ",")
			case "Mime-Version":
				mail.Mime_version = data[1]
			case "Content-Type":
				mail.Content_Type = data[1]
			case "Content-Transfer-Encoding":
				mail.Content_Transfer_Encoding = data[1]
			case "Bcc":
				mail.Bcc = strings.Split(data[1], ",")
			case "X-From":
				mail.X_from = data[1]
			case "X-To":
				mail.X_to = strings.Split(data[1], ",")
			case "X-Cc":
				mail.X_cc = strings.Split(data[1], ",")
			case "X-Bcc":
				mail.X_bcc = strings.Split(data[1], ",")
			case "X-Folder":
				mail.X_folder = data[1]
			case "X-Origin":
				mail.X_origin = data[1]
			case "X-FileName":
				mail.X_filename = data[1]
			}
		} else {
			if data[0] == "" {
				mail.Content += "\n"
			}
			if len(data) == 2 {
				mail.Content += data[0] + ": " + data[1] + "\n"
			} else {
				mail.Content += data[len(data)-1]
			}
		}
	}
	channel <- *mail
	defer wg.Done()
}

func postMailZincSearch2(url string, username string, password string, channel chan models.Mail) (bool, error) {
	json, err := json.Marshal(<-channel)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	req, err1 := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(json))
	if err1 != nil {
		fmt.Println(err1.Error())
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
	defer wg.Done()
	return true, nil
}

func inspectDirectory2(path string, url string, username string, password string, channel chan models.Mail) {
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}
	if !files[0].IsDir() {
		for _, file := range files {
			wg.Add(2)
			var mail models.Mail
			go readFile2(path+"\\"+file.Name(), &mail, channel)
			go postMailZincSearch2(url, username, password, channel)
		}
		return
	}
	for _, file := range files {
		inspectDirectory2(path+"\\"+file.Name(), url, username, password, channel)
	}
}

func IndexerV2(filePath string) {
	now := time.Now()
	channel := make(chan models.Mail, 10)
	inspectDirectory2(filePath, url_api_data, username, password, channel)
	wg.Wait()
	close(channel)
	fmt.Println("Tiempo de ejecucion: ", time.Since(now))
}
