package upload

import (
	"bytes"
	"context"
	"encoding/json"
	"go-pdf/pkg/httpclient"
	"os"

	"log"
)

type ResponseFileService struct {
	OriginalFilename    string `json:"originalFilename"`
	NewFilename         string `json:"newFilename"`
	ObjectHash          string `json:"objectHash"`
	ObjectURL           string `json:"objectUrl"`
	ObjectSize          int    `json:"objectSize"`
	ObjectHashThumbnail string `json:"objectHashThumbnail"`
	ObjectURLThumbnail  string `json:"objectUrlThumbnail"`
	ObjectSizeThumbnail int    `json:"objectSizeThumbnail"`
	Enlarged            bool   `json:"enlarged"`
	ExpirationDate      string `json:"expirationDate"`
}

func UploadFile(ctx context.Context, filename string, fileBuffer []byte) (res *ResponseFileService, err error) {
	client := httpclient.FileService()

	resp, errResp := client.R().
		SetContext(ctx).
		SetResult(&ResponseFileService{}).
		SetFileReader("file", filename, bytes.NewReader(fileBuffer)).
		SetHeader("x-jubelio-fileservice-api-key", os.Getenv("fileservice_api_key")).
		Post("/api/v1/object/file")

	if errResp != nil {
		log.Println("error while upload file to file service", errResp)
		return res, errResp
	}

	if resp.StatusCode() > 399 {
		log.Println("error while upload file to file service", errResp)
		return res, errResp
	}

	var responseResp ResponseFileService
	byteRes, _ := json.Marshal(resp.Result())
	json.Unmarshal(byteRes, &responseResp)

	res = &responseResp

	return res, nil
}

func UploadImage(ctx context.Context, filename string, fileBuffer []byte) (res *ResponseFileService, err error) {
	client := httpclient.FileService()

	resp, errResp := client.R().
		SetContext(ctx).
		SetResult(&ResponseFileService{}).
		SetFileReader("file", filename, bytes.NewReader(fileBuffer)).
		SetHeader("x-jubelio-fileservice-api-key", os.Getenv("fileservice_api_key")).
		Post("/api/v1/object/image")

	if errResp != nil {
		log.Println("error while upload file to file service", errResp)
		return res, errResp
	}

	if resp.StatusCode() > 399 {
		log.Println("body error:", string(resp.Body()))
		log.Println("error while upload file to file service", errResp)
		return res, errResp
	}

	var responseResp ResponseFileService
	byteRes, _ := json.Marshal(resp.Result())
	json.Unmarshal(byteRes, &responseResp)

	res = &responseResp

	return res, nil
}
