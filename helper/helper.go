package helper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Trim(value string) string {
	return strings.TrimSpace(value)
}

func ParseBody[T any](r *http.Request) T {
	var model T
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		json.Unmarshal(body, &model)
	}
	return model
}

func SendResponse(w http.ResponseWriter, statusCode int, payload map[string]any) {
	w.WriteHeader(statusCode)
	payload["status_code"] = statusCode
	b, _ := json.Marshal(payload)
	w.Write(b)
}

func SendData(w http.ResponseWriter, statusCode int, data any) {
	SendResponse(w, statusCode, map[string]any{
		"success": true,
		"data":    data,
	})
}

func SendDataMessage(w http.ResponseWriter, statusCode int, data any, message string) {
	SendResponse(w, statusCode, map[string]any{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func SendDataMessageFailed(w http.ResponseWriter, statusCode int, data any, message string) {
	SendResponse(w, statusCode, map[string]any{
		"success": false,
		"message": message,
		"data":    data,
	})
}

func SendMessage(w http.ResponseWriter, statusCode int, message string) {
	SendResponse(w, statusCode, map[string]any{
		"success": true,
		"message": message,
	})
}

func SendMessageFail(w http.ResponseWriter, statusCode int, message string) {
	SendResponse(w, statusCode, map[string]any{
		"success": false,
		"message": message,
	})
}

func GetLimitOffset(r *http.Request) (limit int, offset int) {
	limit = ParseToInt(r.URL.Query().Get("limit"))
	offset = ParseToInt(r.URL.Query().Get("offset"))
	if limit == 0 {
		limit = 10
	}
	return
}

func ParseToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
