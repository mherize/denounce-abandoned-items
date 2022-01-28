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
	donounceItem = "https://internal-api.mercadolibre.com/moderations/v2/denounces/denounce/%s-ITM"
)

func denounceItem(itemID string) (int, error) {
	payload := domain.DenounceItem{ReportReasonID: "NODISPONIBLE", Comment: "Script pausado de items", CallerID: "612055121", ItemID: itemID, ElementID: "", Type: "ITM", Origin: "fe"}
	client := &http.Client{}
	body, err := json.Marshal(payload)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(donounceItem, itemID), bytes.NewBuffer(body))
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
