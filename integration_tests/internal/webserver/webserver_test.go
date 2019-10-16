package webserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"integration_tests/internal/config"
	"io/ioutil"
	"net/http"
	"testing"
)

type AddUser struct {
	Id        string  `json:"id"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	City      string `json:"city"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Status struct {
	Status bool `json:"status"`
	Detail string `json:"detail"`
	UserID string `json:"user_id"`
}

func TestHttp(t *testing.T) {
	conf := config.NewHttpConf()
	domain := fmt.Sprintf("http://%s:%d", conf.Host, conf.Port)
	//
	client := &http.Client{}
	addReq := &AddUser{
		Login:     "login_http",
		Password:  "123456",
		Email:     "site@mail.ru",
		City:      "kazan",
		Gender:    "Female",
		FirstName: "Ivanov",
		LastName:  "Popov",
	}
	t.Run("Adding user by http", func(t *testing.T) {
		b := new(bytes.Buffer)
		err := json.NewEncoder(b).Encode(addReq)
		if err != nil {
			t.Fatal(err)
		}
		url := fmt.Sprintf("%s/%s", domain, "v1/users")
		req, err := http.NewRequest("POST", url, b)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		status:=&Status{}
		err = json.Unmarshal(body, status)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(status)
	})

	t.Run("Get user by id", func(t *testing.T) {
		url := fmt.Sprintf("%s/%s", domain, "v1/user/2")
		req, err := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		var st Status
		body, err:=ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body,&st)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(st)
	})

}
