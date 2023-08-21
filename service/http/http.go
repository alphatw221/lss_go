package http

import (
	"bytes"
	"fmt"
	"net/url"

	// "bufio"
	"encoding/json"
	"io"
	"net/http"
)

func Get(host string, path string, params map[string]string, headers map[string]string) []byte {

	query := ""
	if params != nil {
		values := url.Values{}
		for key, element := range params {
			values.Add(key, element)
		}
		query = values.Encode()

	}

	target_url := fmt.Sprintf("https://%s/%s?%s", host, path, query)

	fmt.Println(target_url)

	// return
	req, err := http.NewRequest("GET", target_url, nil)
	if err != nil {
		fmt.Println("創建請求錯誤:", err)

	}

	if headers != nil {
		for key, element := range headers {
			req.Header.Set(key, element)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("請求錯誤:", err)
	}
	defer resp.Body.Close()

	// 讀取響應內容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("讀取響應內容錯誤:", err)
	}

	return body

}

func GetStreamEvent(host string, path string, params map[string]string, headers map[string]string) *http.Response {

	query := ""
	if params != nil {
		values := url.Values{}
		for key, element := range params {
			values.Add(key, element)
		}
		query = values.Encode()

	}

	target_url := fmt.Sprintf("https://%s/%s?%s", host, path, query)

	fmt.Println(target_url)

	// return
	req, err := http.NewRequest("GET", target_url, nil)
	if err != nil {
		fmt.Println("創建請求錯誤:", err)

	}

	if headers != nil {
		for key, element := range headers {
			req.Header.Set(key, element)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("請求錯誤:", err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("請求返回狀態碼:", resp.Status)
	}

	return resp

}
func Post(host string, path string, params map[string]string, data map[string]string, headers map[string]string) []byte {

	query := ""
	if params != nil {
		values := url.Values{}
		for key, element := range params {
			values.Add(key, element)
		}
		query = values.Encode()

	}

	target_url := fmt.Sprintf("https://%s/%s&%s", host, path, query)

	jsonStr, err := json.Marshal(data)

	req, err := http.NewRequest("POST", target_url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("創建請求錯誤:", err)
	}

	if headers != nil {
		for key, element := range headers {
			req.Header.Set(key, element)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("請求錯誤:", err)
	}
	defer resp.Body.Close()

	// 讀取響應內容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("讀取響應內容錯誤:", err)
	}

	return body

}
