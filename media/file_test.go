package media_test

import (
	"encoding/base64"
	"encoding/json"
	"github.com/LGuilhermeMoreira/loja-de-cafe/media"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

const (
	path = "../images/"
)

func getData() string {
	resp, err := http.Get("https://api.thecatapi.com/v1/images/search?size=full")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var imageData []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&imageData); err != nil {
		panic(err)
	}
	imageURL := imageData[0]["url"].(string)
	imageResp, err := http.Get(imageURL)
	if err != nil {
		panic(err)
	}
	defer imageResp.Body.Close()
	imageBytes, err := io.ReadAll(imageResp.Body)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(imageBytes)
}

func TestGetBase64(t *testing.T) {
	_, err := media.SaveImage(getData(), path, "imagem_test")
	assert.Nil(t, err)
}

func TestSaveImage(t *testing.T) {
	data := getData()
	pathImage, err := media.SaveImage(data, path, "imagem_test")
	assert.Nil(t, err)
	base64 := media.GetBase64(pathImage)
	if base64 == "" {
		t.Fatal("base64 is empty")
	}
	if base64 != data {
		t.Fatal("base64 is wrong")
	}
}
