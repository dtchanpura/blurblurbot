package bot

import (
	"github.com/disintegration/imaging"
	"github.com/esimov/stackblur-go"

	"fmt"
	"io"
	"time"

	"log"
	"net/http"
	"sort"
	"strings"

	"encoding/json"

	"strconv"

	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
)

func getMaxResolutionPhoto(photos []PhotoSize) PhotoSize {
	// maxResX, maxResY := 0, 0
	var photo PhotoSize
	// Sorting logic
	sort.Sort(ByResolution(photos))
	photo = photos[len(photos)-1]
	// fmt.Printf("Final - res: %dx%d", photo.Width, photo.Height)
	return photo
}

func getPhotoURL(fileID string) string {

	f := FileResult{Ok: false}
	response, err := http.Get(getBotAPIEndpoint(false) + "getFile?file_id=" + fileID)
	if err != nil {
		fmt.Println(err)
	}
	// data, err := ioutil.ReadAll(response.Body)
	// log.Println(string(data[:]))
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&f)
	if err != nil {
		log.Println(err)
	}
	// log.Println(f)
	defer response.Body.Close()
	return getBotAPIEndpoint(true) + f.Result.FilePath
}

// Function for adding blurred layer in background
func resizeBlurPasteImage(w io.Writer, imageURL string, scale float64, blurFactor float64, returnSameSize bool) {
	epch := time.Now().Unix()
	if scale == 0 {
		scale = 1.2
	}
	if blurFactor == 0 {
		blurFactor = 50
	}
	if blurFactor >= 254 {
		blurFactor = 254
	}
	radius := uint32(blurFactor)
	done := make(chan struct{}, radius)

	response, err := http.Get(imageURL)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	srcImage, err := imaging.Decode(response.Body)
	if err != nil {
		log.Println(err)
	}
	log.Printf("start resizeBlurPasteImage - %d\n", epch)
	bounds := srcImage.Bounds()
	width, height := float64(bounds.Dx()), float64(bounds.Dy())
	// log.Printf("debug: got width, height - %d\n", epch)
	pastePt := image.Pt(int(width*(scale-1)/2), int(height*(scale-1)/2))
	resizedImage := imaging.Resize(srcImage, int(width*scale), 0, imaging.Linear)
	// log.Printf("debug: got resizedImage - %d\n", epch)

	blurredImage := stackblur.Process(resizedImage, uint32(resizedImage.Bounds().Dx()), uint32(resizedImage.Bounds().Dy()), uint32(radius), done)
	<-done
	// blurredImage := imaging.Blur(resizedImage, blurFactor)
	// log.Printf("debug: got blurredImage - %d\n", epch)
	finalImage := imaging.Paste(blurredImage, srcImage, pastePt)
	// log.Printf("debug: pasted blurredImage - %d\n", epch)

	if returnSameSize {
		finalImage = imaging.Resize(finalImage, int(width), 0, imaging.Linear)
		// log.Printf("debug: resized finalImage - %d\n", epch)
	}

	// png.Encode(w, finalImage)
	jpeg.Encode(w, finalImage, nil)
	log.Printf("end resizeBlurPasteImage - %d\n", epch)

	// return finalImage
}

// Function for parsing caption string and getting scale and blur radius.
func captionScaleBlur(caption string) (float64, float64, bool, error) {
	values := strings.Split(caption, ",")
	returnSameSize := true
	if len(values) < 2 {
		values = strings.Split(caption, " ")
	}
	if len(values) >= 2 {
		scale, err := strconv.ParseFloat(strings.Trim(values[0], " ,"), -1)
		if err != nil {
			return 0, 0, true, err
		}
		blurFactor, err := strconv.ParseFloat(strings.Trim(values[1], " ,"), -1)
		if err != nil {
			return 0, 0, true, err
		}
		if len(values) >= 3 {
			returnSameSize, _ = strconv.ParseBool(strings.Trim(values[2], " ,"))
		}
		return scale, blurFactor, returnSameSize, nil
	}
	return 0, 0, true, nil

}
