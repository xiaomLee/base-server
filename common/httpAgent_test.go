package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func BenchmarkHttpAgent(t *testing.B) {
	for i := 0; i < 20; i++ {
		request := NewHttpAgent()
		request = request.Post("https://localhost:8080/user/loginByAccount")

		data := map[string]string{
			"orgId":    "99",
			"account":  "tom",
			"password": "123456",
		}

		_, body, err := request.ContentType(TypeFormUrlencoded).SendForm(data).End()
		if err != nil {
			println(err.Error())
			return
		}
		println(string(body))
	}
}

func TestHttpAgent_Timeout(t *testing.T) {
	request := NewHttpAgent()
	request = request.Timeout(time.Millisecond * 500)
	request = request.Post("https://localhost:8080/user/loginByAccount")

	data := map[string]string{
		"account":  "tom",
		"password": "123456",
	}

	_, body, err := request.ContentType(TypeFormUrlencoded).SendForm(data).End()
	if err != nil {
		println(err.Error())
		return
	}
	println(string(body))
}

func TestHttpAgent(t *testing.T) {

	// 1. POST
	request := NewHttpAgent()
	request = request.Post("https://onereg-email-suggest.mail.com/email-alias/availability")
	request.AddHeader("Authorization", "Bearer qXeyJhbGciOiJIUzI1NiJ9.eyJjdCI6Im9faU9JSVQwcVkzdkUwYmVCeUxVZ2RpbnlqWFlIdWpDb2hvNE5kZHdybFRKUkV5cFEwSjJfYWE0TjlZZnI5aVNfRG5sMWZFdTFBUGZxWkJ4SHRxaVRnd1dtQXdfWEVSdjdBQ0FSeTNYU1RRQ3R2VVFYNWNyZWtZREllWExSQjVLTHRNRE5ZWDFoSlFVa3FSMFpNeWl5ZWxFVjlraGxObWhtVWNmVkdYUUFTWSIsInNjb3BlIjoicmVnaXN0cmF0aW9uIiwia2lkIjoiNDk0YjBlMjAiLCJleHAiOjE1ODU1NjY2MjcxOTAsIml2Ijoic0dHQUx6UGRaankwVl9RZ3AxWmFadyIsImlhdCI6MTU4NTU1OTQyNzE5MCwidmVyc2lvbiI6Mn0.2nF1B1t5XGmF8D9U9S8NaJdU7HHeXb1n_avNBiOHhTk")
	request.AddHeader("User-Agent", "Chrome/80.0.3987.149")
	request.AddHeader("Origin", "https://signup.mail.com")
	request.AddHeader("X-REQUEST-ID", "613fef8e-1f94-4abe-a7e8-ed7c71fdcf20")

	data := []byte(`{"emailAddress":"litter_tom@mail.com","countryCode":"VG","suggestionProducts":["mailcomFree"],"maxResultCountPerProduct":"10","mdhMaxResultCount":"5","requestedEmailAddressProduct":"mailcomFree"}`)

	resp, body, err := request.ContentType(TypeJSON).SendData(data).End()
	fmt.Printf("resp:%+v body:%+v err: %v \n", resp, body, err)
	if err != nil {
		println(err.Error())
		return
	}
	println(string(body))
	return
	// GET查询用户信息
	_, body, err = request.Get("https://localhost:8080/user/getUserInfo").End()
	if err != nil {
		println(err.Error())
		return
	}
	println(string(body))

	// 上传文件
	file1, err := os.OpenFile("httpUtil.go", os.O_RDONLY, 0755)
	if err != nil {
		println(err.Error())
		return
	}
	defer file1.Close()
	file2, err := os.OpenFile("logger.go", os.O_RDONLY, 0755)
	if err != nil {
		println(err.Error())
		return
	}
	defer file2.Close()
	data1, _ := ioutil.ReadAll(file1)
	data2, _ := ioutil.ReadAll(file2)
	request = request.SendFile(File{FileName: "httpUtil.go", FieldName: "a", Data: data1})
	request = request.SendFile(File{FileName: "logger.go", FieldName: "b", Data: data2})
	_, body, err = request.Post("https://localhost:8080/octopus/uploadResource").ContentType(TypeMultipartFormData).End()
	if err != nil {
		println(err.Error())
		return
	}
	println(string(body))
}
