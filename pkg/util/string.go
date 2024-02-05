package util

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"strings"
)

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func JsonMarshal(v any) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(marshal)
}

func PrettyString(str string) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", " "); err != nil {
		return ""
	}
	return prettyJSON.String()
}

func JsonUnMarshal(str string, v any) {
	json.Unmarshal([]byte(str), v)
}

func GetUUId() string {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return strings.ReplaceAll(newUUID.String(), "-", "")
}

type kvString struct {
	Key   string
	Value string
}

func ParseUrlTags(tagStr string) []kvString {
	var result []kvString
	var kvArr []string
	for _, kvStr := range strings.Split(tagStr, "&") {
		if !strings.Contains(kvStr, "=") {
			continue
		}
		kvArr = strings.Split(kvStr, "=")
		result = append(result, kvString{
			Key:   kvArr[0],
			Value: kvArr[1],
		})
	}
	return result
}

func PixelIdsSplit(contentIds string) []string {
	all := strings.ReplaceAll(contentIds, "[", "")
	all = strings.ReplaceAll(all, "]", "")
	all = strings.ReplaceAll(all, "\"", "")
	all = strings.ReplaceAll(all, "\\", "")
	return strings.Split(all, ",")
}
