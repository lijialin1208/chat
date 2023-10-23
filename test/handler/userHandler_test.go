package handler

import (
	"bytes"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"io/ioutil"
	"net/http"
	"testing"
)

type UserBasic struct {
	Account  string `json:"username"`
	Password string `json:"password"`
	Mac      string `json:"mac"`
}

func TestLoginHandler(t *testing.T) {
	var userBasic UserBasic
	userBasic.Account = "2021024966"
	userBasic.Password = "101555"
	userBasic.Mac = "2C:0A:F9:2A:38:85"
	b, err := json.Marshal(userBasic)
	if err != nil {
		t.Fatal("json format error:", err)
	}
	body := bytes.NewBuffer(b)
	resp, err := http.Post("http://10.224.97.223:80/api/user/login", "application/json;charset=utf-8", body)
	if err != nil {
		t.Fatal("请求失败")
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	//var usb *UserBasic = &userBasic
	//var usb2 **UserBasic = &usb
	//fmt.Println(&usb)
	//fmt.Println(usb2)
	//fmt.Println(*usb2)
}

func TestRegisterHandler(t *testing.T) {
	var userBasic UserBasic
	userBasic.Account = "2021024966"
	userBasic.Password = "101555"
	userBasic.Mac = "2C:0A:F9:2A:38:73"
	b, err := json.Marshal(userBasic)
	if err != nil {
		t.Fatal("json format error:", err)
	}
	body := bytes.NewBuffer(b)
	resp, err := http.Post("http://10.224.97.223:80/api/user/register", "application/json;charset=utf-8", body)
	if err != nil {
		t.Fatal("请求失败")
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}
