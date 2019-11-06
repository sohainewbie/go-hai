package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func StringToTime(layout, str string) (error, time.Time) {
	t, err := time.Parse(layout, str)
	return err, t
}

func StringToTimeStr(layout, str string) (error, string) {
	t, err := time.Parse(layout, str)

	return err, t.Format("2006-01-02 15:04:05")
}

func TimeToTimeStr(t time.Time) string {

	return t.Format("2006-01-02 15:04:05")
}

func StringMustMatchContains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func StringContains(arr []string, str string) bool {
	for _, a := range arr {
		if strings.Contains(str, a) {
			return true
		}
	}
	return false
}

func RandStringBytes(n int) (output string) {
	var (
		temp        = make([]byte, n)
		letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	)
	for i := range temp {
		temp[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	output = string(temp)
	return
}

// FormatArrayOfIntExpandToString - format array of integer
// and return string containing that integers separated with comma
func FormatArrayOfIntExpandToString(intArray []uint64) string {
	var buffer bytes.Buffer
	lenInt := len(intArray)
	if lenInt <= 0 {
		return ""
	}
	for i, value := range intArray {
		if i == lenInt-1 {
			buffer.WriteString(fmt.Sprintf(`%d`, value))
		} else {
			buffer.WriteString(fmt.Sprintf(`%d,`, value))
		}
	}
	return buffer.String()
}
