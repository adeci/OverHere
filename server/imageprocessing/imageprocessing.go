package imageprocessing

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func DecodePNG(file string) string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toBase64(bytes)

	return base64Encoding
}

func EncodePNG(file bson.M) {
	fieldBase64Encoding := file["fieldbase64encoding"].(string)
	b64data := fieldBase64Encoding[strings.IndexByte(fieldBase64Encoding, ',')+1:]

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
