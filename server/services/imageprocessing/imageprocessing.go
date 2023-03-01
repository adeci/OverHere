package imageprocessing

import (
	"encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func ReadImage(file string) []byte {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func DecodePNG(file string) string {
	var base64Encoding string

	mimeType := http.DetectContentType(ReadImage(file))

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += ToBase64(ReadImage(file))

	return base64Encoding
}

func EncodePNG(file bson.M) {
	fieldBase64Encoding := file["encoding"].(string)
	b64data := fieldBase64Encoding[strings.IndexByte(fieldBase64Encoding, ',')+1:]
	fmt.Println(b64data)

	imgData, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("output.png")
	if err != nil {
		fmt.Println(err)
	}

	if _, err := f.Write(imgData); err != nil {
		fmt.Println(err)
	}

	if err := f.Close(); err != nil {
		fmt.Println(err)
	}
}
