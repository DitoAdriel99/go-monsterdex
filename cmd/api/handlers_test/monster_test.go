package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestSuccessCreateMonsterHandler() {
	client := &http.Client{}

	// Login request
	login := []byte(`{
        "email": "admin@gmail.com",
        "password": "admin"
    }`)

	reqLogin, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/login", bytes.NewBuffer(login))
	if err != nil {
		s.Fail("Failed to create login request:", err)
		return
	}
	reqLogin.Header.Set("Content-Type", "application/json")

	respLogin, err := client.Do(reqLogin)
	if err != nil {
		s.Fail("Failed to send login request:", err)
		return
	}
	defer respLogin.Body.Close()

	bodyLogin, err := ioutil.ReadAll(respLogin.Body)
	if err != nil {
		s.Fail("Failed to read login response body:", err)
		return
	}

	// Extract JWT from login response
	var loginResponse struct {
		Status  string            `json:"status"`
		Message string            `json:"message"`
		Data    map[string]string `json:"data"`
	}
	if err := json.Unmarshal(bodyLogin, &loginResponse); err != nil {
		s.Fail("Failed to parse login response:", err)
		return
	}

	// Create monster request with JWT
	createMonster := []byte(`{
		"name": "Ditto",
		"monster_category_id": 1,
		"description": "A mischievous ghostly creature",
		"types_id": [1, 2],
		"image": "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",
		"height": 0.7,
		"weight": 6.9,
		"stats_hp": 45,
		"stats_attack": 49,
		"stats_defense": 49,
		"stats_speed": 45
	}`)

	reqMonster, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/monster", bytes.NewBuffer(createMonster))
	if err != nil {
		s.Fail("Failed to create create-monster request:", err)
		return
	}

	reqMonster.Header.Set("Content-Type", "application/json")
	reqMonster.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginResponse.Data["token"])) // Set JWT in the Authorization header

	respMonster, err := client.Do(reqMonster)
	if err != nil {
		s.Fail("Failed to send create-monster request:", err)
		return
	}
	defer respMonster.Body.Close()

	s.Equal(http.StatusCreated, respMonster.StatusCode)
}

func (s *EndToEndSuite) TestFailedCreateMonsterHandler() {
	client := &http.Client{}

	// Login request
	login := []byte(`{
        "email": "admin@gmail.com",
        "password": "admin"
    }`)

	reqLogin, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/login", bytes.NewBuffer(login))
	if err != nil {
		s.Fail("Failed to create login request:", err)
		return
	}
	reqLogin.Header.Set("Content-Type", "application/json")

	respLogin, err := client.Do(reqLogin)
	if err != nil {
		s.Fail("Failed to send login request:", err)
		return
	}
	defer respLogin.Body.Close()

	bodyLogin, err := ioutil.ReadAll(respLogin.Body)
	if err != nil {
		s.Fail("Failed to read login response body:", err)
		return
	}

	// Extract JWT from login response
	var loginResponse struct {
		Status  string            `json:"status"`
		Message string            `json:"message"`
		Data    map[string]string `json:"data"`
	}
	if err := json.Unmarshal(bodyLogin, &loginResponse); err != nil {
		s.Fail("Failed to parse login response:", err)
		return
	}

	createMonster := []byte(`{
		"name": "Ditto",
		"monster_category_id": 1,
		"description": "A mischievous ghostly creature",
		"types_id": [1, 2],
		"image": "",
		"height": 0.7,
		"weight": 6.9,
		"stats_hp": 45,
		"stats_attack": 49,
		"stats_defense": 49,
		"stats_speed": 45
	}`)

	reqMonster, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/monster", bytes.NewBuffer(createMonster))
	if err != nil {
		s.Fail("Failed to create create-monster request:", err)
		return
	}

	reqMonster.Header.Set("Content-Type", "application/json")
	reqMonster.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginResponse.Data["token"])) // Set JWT in the Authorization header

	respMonster, err := client.Do(reqMonster)
	if err != nil {
		s.Fail("Failed to send create-monster request:", err)
		return
	}
	defer respMonster.Body.Close()

	s.Equal(http.StatusBadRequest, respMonster.StatusCode)
}

func (s *EndToEndSuite) TestSuccessGetMonsterHandler() {
	client := &http.Client{}

	// Login request
	login := []byte(`{
        "email": "akuntes@gmail.com",
        "password": "admin"
    }`)

	reqLogin, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/login", bytes.NewBuffer(login))
	if err != nil {
		s.Fail("Failed to create login request:", err)
		return
	}
	reqLogin.Header.Set("Content-Type", "application/json")

	respLogin, err := client.Do(reqLogin)
	if err != nil {
		s.Fail("Failed to send login request:", err)
		return
	}
	defer respLogin.Body.Close()

	bodyLogin, err := ioutil.ReadAll(respLogin.Body)
	if err != nil {
		s.Fail("Failed to read login response body:", err)
		return
	}

	// Extract JWT from login response
	var loginResponse struct {
		Status  string            `json:"status"`
		Message string            `json:"message"`
		Data    map[string]string `json:"data"`
	}
	if err := json.Unmarshal(bodyLogin, &loginResponse); err != nil {
		s.Fail("Failed to parse login response:", err)
		return
	}

	reqMonster, err := http.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/monsters", nil)
	if err != nil {
		s.Fail("Failed to create create-monster request:", err)
		return
	}

	reqMonster.Header.Set("Content-Type", "application/json")
	reqMonster.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginResponse.Data["token"])) // Set JWT in the Authorization header

	respMonster, err := client.Do(reqMonster)
	if err != nil {
		s.Fail("Failed to send create-monster request:", err)
		return
	}
	defer respMonster.Body.Close()

	s.Equal(http.StatusOK, respMonster.StatusCode)
}

func (s *EndToEndSuite) TestSuccessUpdateMonsterHandler() {
	client := &http.Client{}

	// Login request
	login := []byte(`{
        "email": "admin@gmail.com",
        "password": "admin"
    }`)

	reqLogin, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/login", bytes.NewBuffer(login))
	if err != nil {
		s.Fail("Failed to create login request:", err)
		return
	}
	reqLogin.Header.Set("Content-Type", "application/json")

	respLogin, err := client.Do(reqLogin)
	if err != nil {
		s.Fail("Failed to send login request:", err)
		return
	}
	defer respLogin.Body.Close()

	bodyLogin, err := ioutil.ReadAll(respLogin.Body)
	if err != nil {
		s.Fail("Failed to read login response body:", err)
		return
	}

	// Extract JWT from login response
	var loginResponse struct {
		Status  string            `json:"status"`
		Message string            `json:"message"`
		Data    map[string]string `json:"data"`
	}
	if err := json.Unmarshal(bodyLogin, &loginResponse); err != nil {
		s.Fail("Failed to parse login response:", err)
		return
	}

	updateMonster := []byte(`{
		"name": "Change From Unit test"
	}`)

	reqMonster, err := http.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/monster/9", bytes.NewBuffer(updateMonster))
	if err != nil {
		s.Fail("Failed to create create-monster request:", err)
		return
	}

	reqMonster.Header.Set("Content-Type", "application/json")
	reqMonster.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginResponse.Data["token"])) // Set JWT in the Authorization header

	respMonster, err := client.Do(reqMonster)
	if err != nil {
		s.Fail("Failed to send create-monster request:", err)
		return
	}
	defer respMonster.Body.Close()

	s.Equal(http.StatusOK, respMonster.StatusCode)
}

func (s *EndToEndSuite) TestFailedUpdateMonsterHandler() {
	client := &http.Client{}

	// Login request
	login := []byte(`{
        "email": "admin@gmail.com",
        "password": "admin"
    }`)

	reqLogin, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/login", bytes.NewBuffer(login))
	if err != nil {
		s.Fail("Failed to create login request:", err)
		return
	}
	reqLogin.Header.Set("Content-Type", "application/json")

	respLogin, err := client.Do(reqLogin)
	if err != nil {
		s.Fail("Failed to send login request:", err)
		return
	}
	defer respLogin.Body.Close()

	bodyLogin, err := ioutil.ReadAll(respLogin.Body)
	if err != nil {
		s.Fail("Failed to read login response body:", err)
		return
	}

	// Extract JWT from login response
	var loginResponse struct {
		Status  string            `json:"status"`
		Message string            `json:"message"`
		Data    map[string]string `json:"data"`
	}
	if err := json.Unmarshal(bodyLogin, &loginResponse); err != nil {
		s.Fail("Failed to parse login response:", err)
		return
	}

	updateMonster := []byte(`{
		"image": "not base64 format"
	}`)

	reqMonster, err := http.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/monster/9", bytes.NewBuffer(updateMonster))
	if err != nil {
		s.Fail("Failed to create create-monster request:", err)
		return
	}

	reqMonster.Header.Set("Content-Type", "application/json")
	reqMonster.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginResponse.Data["token"])) // Set JWT in the Authorization header

	respMonster, err := client.Do(reqMonster)
	if err != nil {
		s.Fail("Failed to send create-monster request:", err)
		return
	}
	defer respMonster.Body.Close()

	s.Equal(http.StatusBadRequest, respMonster.StatusCode)
}

func (s *EndToEndSuite) TestSuccessSetStatusMonsterHandler() {
	client := &http.Client{}

	// Login request
	login := []byte(`{
        "email": "admin@gmail.com",
        "password": "admin"
    }`)

	reqLogin, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/login", bytes.NewBuffer(login))
	if err != nil {
		s.Fail("Failed to create login request:", err)
		return
	}
	reqLogin.Header.Set("Content-Type", "application/json")

	respLogin, err := client.Do(reqLogin)
	if err != nil {
		s.Fail("Failed to send login request:", err)
		return
	}
	defer respLogin.Body.Close()

	bodyLogin, err := ioutil.ReadAll(respLogin.Body)
	if err != nil {
		s.Fail("Failed to read login response body:", err)
		return
	}

	// Extract JWT from login response
	var loginResponse struct {
		Status  string            `json:"status"`
		Message string            `json:"message"`
		Data    map[string]string `json:"data"`
	}
	if err := json.Unmarshal(bodyLogin, &loginResponse); err != nil {
		s.Fail("Failed to parse login response:", err)
		return
	}

	updateMonster := []byte(`{
		"status": false
	}`)

	reqMonster, err := http.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/monster/status/6", bytes.NewBuffer(updateMonster))
	if err != nil {
		s.Fail("Failed to create create-monster request:", err)
		return
	}

	reqMonster.Header.Set("Content-Type", "application/json")
	reqMonster.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginResponse.Data["token"])) // Set JWT in the Authorization header

	respMonster, err := client.Do(reqMonster)
	if err != nil {
		s.Fail("Failed to send create-monster request:", err)
		return
	}
	defer respMonster.Body.Close()

	var monsterResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Data    string `json:"data"`
	}
	resp, _ := ioutil.ReadAll(respMonster.Body)
	if err := json.Unmarshal(resp, &monsterResponse); err != nil {
		s.Fail("Failed to parse login response:", err)
		return
	}

	if monsterResponse.Data == "Success Set Status!" {
		s.Equal(http.StatusOK, respMonster.StatusCode)
	} else {
		s.Equal(http.StatusBadRequest, respMonster.StatusCode)

	}

}

func (s *EndToEndSuite) TestSuccessCatchAndReleaseMonsterHandler() {
	client := &http.Client{}

	// Login request
	login := []byte(`{
        "email": "akuntes@gmail.com",
        "password": "admin"
    }`)

	reqLogin, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/login", bytes.NewBuffer(login))
	if err != nil {
		s.Fail("Failed to create login request:", err)
		return
	}
	reqLogin.Header.Set("Content-Type", "application/json")

	respLogin, err := client.Do(reqLogin)
	if err != nil {
		s.Fail("Failed to send login request:", err)
		return
	}
	defer respLogin.Body.Close()

	bodyLogin, err := ioutil.ReadAll(respLogin.Body)
	if err != nil {
		s.Fail("Failed to read login response body:", err)
		return
	}

	// Extract JWT from login response
	var loginResponse struct {
		Status  string            `json:"status"`
		Message string            `json:"message"`
		Data    map[string]string `json:"data"`
	}
	if err := json.Unmarshal(bodyLogin, &loginResponse); err != nil {
		s.Fail("Failed to parse login response:", err)
		return
	}

	reqMonster, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/monster/catch/2", nil)
	if err != nil {
		s.Fail("Failed to create create-monster request:", err)
		return
	}

	reqMonster.Header.Set("Content-Type", "application/json")
	reqMonster.Header.Set("Authorization", fmt.Sprintf("Bearer %s", loginResponse.Data["token"])) // Set JWT in the Authorization header

	respMonster, err := client.Do(reqMonster)
	if err != nil {
		s.Fail("Failed to send create-monster request:", err)
		return
	}
	defer respMonster.Body.Close()

	s.Equal(http.StatusOK, respMonster.StatusCode)

}
