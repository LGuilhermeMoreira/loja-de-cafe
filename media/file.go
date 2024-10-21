package media

import (
	"encoding/base64"
	"log"
	"os"
)

func SaveImage(data, filePath string, label string) (string, error) {
	imageData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	name := filePath + label + ".jpeg"
	err = os.WriteFile(name, imageData, os.ModePerm)
	if err != nil {
		return "", err
	}
	return name, nil
}

func GetBase64(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(file)
}

func UpdateImage(data, path string) error {
	imageData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, imageData, os.ModePerm)
}

func DeleteImage(path string) error {
	return os.Remove(path)
}
