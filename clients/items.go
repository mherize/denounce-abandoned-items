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
	itemsNWUri = "https://internal-api.mercadolibre.com/items/%s?caller.scopes=admin&client.id=%v"
	itemsOWUri = "http://seguidor-test.portalinmobiliario.cl/Services/PropiedadService.svc/ajax/PauseItem"
	appID = "6614381939938131"
)

func PauseItemNW(itemID string) (int, error) {
	payload := domain.ItemsNWPayload{Status: "paused"}
	client := &http.Client{}
	body, err := json.Marshal(payload)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	//TODO: VALIDAR clientID
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(itemsNWUri, itemID, appID), bytes.NewBuffer(body))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return resp.StatusCode, errors.New(string(bodyBytes))
	}

	return resp.StatusCode, nil
}

func PauseItemOW(itemID string) (int, error) {
	payload := domain.ItemsOWPayload{
		ItemID:      itemID,
		Observation: "items-moderation",
		ActionType:  "2",
	}
	client := &http.Client{}
	body, err := json.Marshal(payload)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	req, err := http.NewRequest(http.MethodPost, itemsOWUri, bytes.NewBuffer(body))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("token_fury", "c7d6c054-bf1f-4486-9e36-7de92122e935")

	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return resp.StatusCode, errors.New(string(bodyBytes))
	}

	return resp.StatusCode, nil
}
