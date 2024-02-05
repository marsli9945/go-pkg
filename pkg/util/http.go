package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func Post(url string, data any) (string, error) {
	req, err := getRequest(url, data)
	if err != nil {
		return "", nil
	}
	return doClient(req)
}

func Get(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func CheckUrl(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		return false
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusForbidden && response.StatusCode != 567 {
		fmt.Println(response.StatusCode)
		return false
	}

	// 读取响应
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false
	}
	if len(body) < 1 {
		return false
	}
	return true
}

func Cmd(cmd string) (string, error) {
	// 创建Cmd对象
	command := exec.Command("bash", "-c", cmd)

	// 执行命令并获取输出
	output, err := command.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return string(output), nil
}

func PostAndHeader(url string, data any, headers map[string]string) (string, error) {
	req, err := getRequest(url, data)
	if err != nil {
		return "", nil
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return doClient(req)
}

func getRequest(url string, data any) (*http.Request, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func doClient(req *http.Request) (string, error) {
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
