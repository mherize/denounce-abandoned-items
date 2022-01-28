package clients

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"script-baja-items/domain"
)

const (
	emailsUrl = "https://internal-api.mercadolibre.com/internal/email"
)

func SendMail(payload domain.Email) (int, error) {
	client := &http.Client{}
	body, err := json.Marshal(payload)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(emailsUrl), bytes.NewBuffer(body))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return resp.StatusCode, errors.New(string(bodyBytes))
	}

	return resp.StatusCode, nil
}
