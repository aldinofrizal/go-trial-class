package helpers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
)

func SendMail(email, address, productName string) error {
	url := "http://localhost:8001/mail"
	mailData := map[string]string{
		"email":        email,
		"address":      address,
		"product_name": productName,
	}

	marshallMailData, err := json.Marshal(mailData)
	if err != nil {
		return err
	}

	jsonStr := []byte(marshallMailData)

	request, error := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if error != nil {
		return error
	}

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return error
	}

	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	return nil
}
