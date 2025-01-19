package helper

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"
)

func Upload(wg *sync.WaitGroup, files []*multipart.FileHeader) ([]CdnResponse, error) {
	var results []CdnResponse
	var err error
	for _, file := range files {
		wg.Add(1)

		go func() {
			defer wg.Done()
			var f multipart.File
			f, err = file.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			var part io.Writer
			part, err = writer.CreateFormFile("image", filepath.Base(""))
			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(part, f)

			err = writer.Close()
			if err != nil {
				log.Fatal(err)
			}

			var request *http.Request
			request, err = http.NewRequest("POST", "https://cdn-lumoshive-academy.vercel.app/api/v1/upload", body)
			if err != nil {
				log.Fatal(err)
			}
			request.Header.Add("Content-Type", writer.FormDataContentType())
			client := &http.Client{}
			var response *http.Response
			response, err = client.Do(request)
			if err != nil {
				log.Fatal(err)
			}

			defer response.Body.Close()

			var res []byte
			res, err = io.ReadAll(response.Body)
			if err != nil {
				log.Fatal("Error reading response:", err)
			}

			var result CdnResponse
			json.Unmarshal(res, &result)
			results = append(results, result)
		}()
	}
	wg.Wait()
	return results, err
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func BadResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, Response{
		Status:  false,
		Message: message,
	})
}

func GoodResponseWithData(c *gin.Context, message string, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func GoodResponseWithPage(c *gin.Context, message string, statusCode, total, totalPages, page, Limit int, data interface{}) {
	c.JSON(statusCode, DataPage{
		Status:      true,
		Message:     message,
		Total:       int64(total),
		Pages:       totalPages,
		CurrentPage: uint(page),
		Limit:       uint(Limit),
		Data:        data,
	})
}

type CdnResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		FileId      string `json:"fileId"`
		Name        string `json:"name"`
		Size        int    `json:"size"`
		VersionInfo struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"versionInfo"`
		FilePath     string      `json:"filePath"`
		Url          string      `json:"url"`
		FileType     string      `json:"fileType"`
		Height       int         `json:"height"`
		Width        int         `json:"width"`
		ThumbnailUrl string      `json:"thumbnailUrl"`
		AITags       interface{} `json:"AITags"`
	} `json:"data"`
}

type DataPage struct {
	Status      bool        `json:"status"`
	Message     string      `json:"message"`
	Total       int64       `json:"total"`
	Pages       int         `json:"pages"`
	CurrentPage uint        `json:"current_page"`
	Limit       uint        `json:"per_page"`
	Data        interface{} `json:"data"`
}
