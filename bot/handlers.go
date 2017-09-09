package blurblurbot

import (
	"log"
	"net/http"
	"strconv"
	// "net/url"
	"encoding/json"
	"fmt"
)

func BotUpdateHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(request.Body)
	encoder := json.NewEncoder(writer)
	var update Update
	err := decoder.Decode(&update)
	if err != nil {
		log.Print(err)
		NoResponse(writer)
	}

	// log.Printf("Message: %v\n", update.Message != nil)
	// log.Printf("CallbackQuery: %v\n", update.CallbackQuery != nil)

	if update.Message != nil {
		resp, err := processMessage(update)
		if err != nil {
			log.Print(err)
			NoResponse(writer)
		}
		encoder.Encode(&resp)

	} else {
		NoResponse(writer)
	}
}

func ResizeBlurHandler(writer http.ResponseWriter, request *http.Request) {
	var scale float64
	var blurFactor float64
	var sameSize bool
	var imageFileId string

	queries := request.URL.Query()

	if val, ok := queries["fileId"]; ok {
		imageFileId = val[0]
	} else {
		ErrorHandler(writer, request, "{\"ok\":false,\"code\":9000}")
	}
	if s, ok := queries["scale"]; ok {
		s, err := strconv.ParseFloat(s[0], 64)
		scale = s
		if err != nil {
			ErrorHandler(writer, request, "{\"ok\":false,\"code\":9001}")
		}
	} else {
		scale = 1.5
	}
	if bf, ok := queries["blurFactor"]; ok {
		bf, err := strconv.ParseFloat(bf[0], 64)
		blurFactor = bf
		if err != nil {
			ErrorHandler(writer, request, "{\"ok\":false,\"code\":9002}")
		}
	} else {
		blurFactor = 50
	}

	if retSameSize, ok := queries["same"]; ok {
		sameSize, _ = strconv.ParseBool(retSameSize[0])
	} else {
		blurFactor = 50
	}
	photoURL := getPhotoUrl(imageFileId)
	// log.Println(photoURL)
	writer.Header().Set("Content-Type", "image/jpeg")
	resizeBlurPasteImage(writer, photoURL, scale, blurFactor, sameSize)

	// img, err := os.Open("example_t500.jpg")
	// if err != nil {
	// log.Fatal(err) // perhaps handle this nicer
	// }
	// defer img.Close()
	// w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	// io.Copy(w, img)
	// w.Header().Set("Content-Type", "image/jpeg")
	// resizeBlurPasteImage(w, )
}

func NoResponse(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, "{}")
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, errorData string) {
	log.Println(errorData)
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, errorData)

}
