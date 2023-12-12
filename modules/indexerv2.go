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

//Esta funcion se ejecuta en una go rutina y se encarga de leer el archivo y
//agregar el Mail en en canal
func readFile2(path string, mail *models.Mail, channel chan models.Mail) {
	fmt.Println("Leyendo ", path)
	file, err1 := os.Open(path)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	file.Close()
	content, err2 := os.ReadFile(path)
	if err2 != nil {
		fmt.Println(err2.Error())
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

//Esta funcion se ejecuta en una go rutina y se encarga de obtener un Mail
//del canal e indexarlo en zincsearch
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

//Esta funcion recursiva se encarga de verificar si una ruta contiene archivos o es una carpeta
//en caso de que contenga archivos para leer, se recorre con un bucle for cada archivo y se
//ejecutan las go rutinas para leer el archivo mediante la funcion readFile2, mientras que
//por otro lado se ejecutan las go rutinas para indexar las datos a zincsearch mediante la funcion
//postMailZincSearch2, esta funcion escucha el canal de Mail y obtiene el dato para luego indexarlo.
func inspectDirectory2(path string, url string, username string, password string, channel chan models.Mail) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(path)
	if !files[0].IsDir() {
		for _, file := range files {
			var mail models.Mail
			wg.Add(2)
			go readFile2(path+"\\"+file.Name(), &mail, channel)
			go postMailZincSearch2(url, username, password, channel)
		}
		return
	}
	for _, file := range files {
		inspectDirectory2(path+"\\"+file.Name(), url, username, password, channel)
	}
}

//Esta es una version mejorada de la funcion IndexerV1, ya que se hace uso de concurrencia para procesar datos
//de manera paralela. Para realizarlo, se creó un canal de tipo Mail con un buffer de 10, este canal permitirá
//que nuestras go rutinas se comuniquen. Las Funciones ReadFile2 y postMailZincSearch2 seran las go rutinas.
//Mientras que inspectDirectory2 se encarga de ejecutar estas funciones. Para ello, la funcion inspectDirectory2
//recorre recursivamente las carpetas, una vez que se detecta que una carpeta contiene los archivos necesarios
//para la lectura, comienza a ejecutar la funcion readFile2 en una go rutina. mientras que por la otra parte, la
//funcion postMailZincSearch2 espera los Mail entrantes para indexarlo en zincsearch, todo esto de manera concurrente.
func IndexerV2(filePath string) {
	now := time.Now()
	channel := make(chan models.Mail, 10)
	inspectDirectory2(filePath, url_api_data, username, password, channel)
	wg.Wait()
	close(channel)
	fmt.Println("Tiempo de ejecucion: ", time.Since(now))
}
