package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	daoApi "api/dao_service"
	_ "app-service/algorithm-service/routers"
	"github.com/astaxie/beego"
	"model"
)

const (
	base_url = "http://localhost:8080/v1/algorithm"
)

func Test_Create(t *testing.T) {
	// create admin
	var cfg beego.AppConfig
	daoApi.UserDaoApi.Init(cfg.String("UserDaoService"))
	var admin model.User
	var resource model.Resource
	admin.Name = "admin"
	admin.Resource = &resource
	daoApi.UserDaoApi.Create(user)

	// create image
	var algorithm model.Algorithm
	algorithm.Id = 0
	algorithm.Name = "algorithm"
	algorithm.Version = "v1"
	algorithm.Description = "this is algorithm test"
	algorithm.Author = "qsg"
	algorithm.Parameters = "param1=1"
	algorithm.Image = ""
	algorithm.Downloads = 0
	algorithm.Stars = 0

	// post create action
	requestData, err := json.Marshal(&algorithm)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	res, err := http.Post(base_url+"/", "application/x-www-form-urlencoded", bytes.NewBuffer(requestData))
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}
