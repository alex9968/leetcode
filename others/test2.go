package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	url := "http://localhost:8099/api/v1/logFile/upload"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("project_id", "35")
	_ = writer.WriteField("times", "1")
	_ = writer.WriteField("type", "3")
	_ = writer.WriteField("category", "Agent")
	_ = writer.WriteField("ip", "127.0.0.122")
	_ = writer.WriteField("agent_type", "Linux")
	file, errFile7 := os.Open("C:/Users/AH/Desktop/log-demo/log (2) (2).zip")
	defer file.Close()
	part7,
	errFile7 := writer.CreateFormFile("file",filepath.Base("/C:/Users/AH/Desktop/log-demo/log (2) (2).zip"))
	_, errFile7 = io.Copy(part7, file)
	if errFile7 != nil {
		fmt.Println(errFile7)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}


	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Api-Key", "X427ro4hkEwBJjsQuGzk")

	req.Header.Set("Content-Type", writer.FormDataContentType())

	byt22, err := ioutil.ReadAll(req.Body)
	fmt.Println("ss",string(byt22))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}