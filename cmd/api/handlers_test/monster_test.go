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

func (s *EndToEndSuite) TestCreateMonsterHandler() {
	client := &http.Client{}

	// Login request
	login := []byte(`{
        "email": "admin@gmail.com",
        "password": "admin"
    }`)

	reqLogin, err := http.NewRequest(http.MethodPost, "http://localhost:2000/api/v1/login", bytes.NewBuffer(login))
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

	fmt.Println("Response login:", string(bodyLogin))

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

	fmt.Println("Token:", loginResponse.Data["token"])

	// Create monster request with JWT
	createMonster := []byte(`{
        "name": "MonsterName",
        "monster_type": "MonsterType",
        "description": "MonsterDescription",
        "type": ["Type1", "Type2"],
        "height": 1.5,
        "weight": 50.0,
        "stats_hp": 100,
        "stats_attack": 90,
        "stats_defense": 80,
        "stats_speed": 70
    }`)

	reqMonster, err := http.NewRequest(http.MethodPost, "http://localhost:2000/api/v1/monster", bytes.NewBuffer(createMonster))
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

	bodyMonster, err := ioutil.ReadAll(respMonster.Body)
	if err != nil {
		s.Fail("Failed to read create-monster response body:", err)
		return
	}

	s.Equal(http.StatusCreated, respMonster.StatusCode)
	fmt.Println("Response body:", string(bodyMonster))
}
