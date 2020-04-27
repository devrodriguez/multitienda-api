package utilities

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

func ToFile(data, filename string) (string, error) {
	idx := strings.Index(data, ";base64,")
	var filExt string
	var storageUrl = "storage/"
	var localPath = "stores/" + filename
	var filePath = storageUrl + localPath

	if idx < 0 {
		log.Println("Base 64 not valid")
		return "", errors.New("Base 64 not valid")
	}

	imageType := data[11:idx]
	unbased, err := base64.StdEncoding.DecodeString(data[idx+8:])

	if err != nil {
		log.Println("Error decoding string 64")
		return "", err
	}

	reader := bytes.NewReader(unbased)

	switch imageType {
	case "png":
		filExt = ".png"
		filePath = filePath + filExt
		img, err := png.Decode(reader)

		if err != nil {
			log.Println("Error decoding png reader")
			return "", err
		}

		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)

		if err != nil {
			log.Println("Error creating file")
			return "", err
		}

		defer file.Close()

		if err := png.Encode(file, img); err != nil {
			log.Println("Error writing png file")
			return "", err
		}
	case "jpeg":
		filExt = ".jpg"
		filePath = filePath + filExt
		img, err := jpeg.Decode(reader)

		if err != nil {
			log.Println("Error decoding jpeg reader")
			return "", err
		}

		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)

		if err != nil {
			log.Println("Error creating file")
			return "", err
		}

		defer file.Close()

		if err := jpeg.Encode(file, img, nil); err != nil {
			log.Println("Error writing jpeg file")
			return "", err
		}
	case "gif":
		filExt = ".gif"
		filePath = filePath + filExt
		img, err := gif.Decode(reader)

		if err != nil {
			log.Println("Error decoding gif reader")
			return "", err
		}

		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)

		if err != nil {
			log.Println("Error creating file")
			return "", err
		}

		defer file.Close()

		if err := gif.Encode(file, img, nil); err != nil {
			log.Println("Error writing gif file")
			return "", err
		}
	}

	return localPath + filExt, nil
}
