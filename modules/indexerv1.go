package modules

import (
	"enron_mail-indexer/models"
	"fmt"
	"os"
)

func inspectDirectory(path string, isDir bool, url string, username string, password string) {
	if !isDir {
		var mail models.Mail
		ReadFile(path, &mail)
		success, err := PostMailZincSearch(url, mail, username, password)
		if err != nil {
			fmt.Printf("Error en %v %v", path, err.Error())
		}
		if !success {
			fmt.Println("No se inserto ", path)
		}
		return
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}
	for _, file := range files {
		inspectDirectory(path+"\\"+file.Name(), file.IsDir(), url, username, password)
	}
}

func IndexerV1(filePath string) {
	inspectDirectory(filePath, true, url_api_data, username, password)
}
