package bot

import (
	"bytes"
	"encoding/json"
	"time"
	// "fmt"
	// "io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	// "os"
	"strconv"
)

func SendTgPhotoFormData(chatId int64, fileId string, scale, blurFactor float64, returnSameSize bool) {
	epch := time.Now().Unix()

	log.Printf("start - SendTgPhotoFormData - %d\n", epch)
	imgBuffer := new(bytes.Buffer)

	resizeBlurPasteImage(imgBuffer, getPhotoUrl(fileId), scale, blurFactor, returnSameSize)
	// log.Println(GetBotApiEndpoint(false) + "sendPhoto")
	// file, err := os.Open(path)

	// if err != nil {
	// 	return nil, err
	// }
	imgContents, err := ioutil.ReadAll(imgBuffer)
	// log.Println(len(imgContents))
	if err != nil {
		return // nil, err
	}
	// imgBuffer.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("photo", "photo.png")
	if err != nil {
		log.Println(err)
		return // nil, err
	}
	part.Write(imgContents)

	_ = writer.WriteField("chat_id", strconv.FormatInt(chatId, 10))
	err = writer.Close()
	if err != nil {
		log.Println(err)
		return //nil, err
	}

	// For reply markup

	// _ = writer.WriteField("reply_markup", "{\"inline_keyboard\":[[{\"text\": \"test\", \"url\":\"http://dcpri.me\"}]]}")
	// err = writer.Close()
	// if err != nil {
	// 	log.Println(err)
	// 	return //nil, err
	// }

	request, _ := http.NewRequest("POST", GetBotApiEndpoint(false)+"sendPhoto", body)

	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		// log.Fatal(err)
		defer resp.Body.Close()
		bodyContent, _ := ioutil.ReadAll(resp.Body)
		log.Println(resp.StatusCode)
		log.Println(resp.Header)
		// resp.Body.Read(bodyContent)
		// resp.Body.Close()
		log.Println(string(bodyContent))
	} else {
		resp.Body.Close()
		// var bodyContent []byte
		// defer resp.Body.Close()
		// bodyContent, _ := ioutil.ReadAll(resp.Body)
		// log.Println(resp.StatusCode)
		// log.Println(resp.Header)
		// resp.Body.Read(bodyContent)
		// log.Println(string(bodyContent))
	}
	log.Printf("end - SendTgPhotoFormData - %d\n", epch)
}

func SendTgApiRequest(method BaseMethod) {
	// var requestBytes []byte
	// requestBuffer := bytes.NewBuffer(requestBytes)
	var requestBuffer bytes.Buffer
	encoder := json.NewEncoder(&requestBuffer)
	encoder.Encode(method)
	resp, err := http.Post(GetBotApiEndpoint(false), "application/json", &requestBuffer)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		bodyContent, _ := ioutil.ReadAll(resp.Body)
		log.Println(resp.StatusCode)
		log.Println(resp.Header)
		resp.Body.Read(bodyContent)
		log.Println(string(bodyContent))
	}
}
