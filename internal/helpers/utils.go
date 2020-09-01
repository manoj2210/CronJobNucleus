package helpers

import (
	"bytes"
	"github.com/google/go-github/github"
	"net/http"
	"reflect"
	"time"
)

func Filter(arr interface{}, cond func(interface{}) bool) interface{} {
	contentType := reflect.TypeOf(arr)
	contentValue := reflect.ValueOf(arr)

	newContent := reflect.MakeSlice(contentType, 0, 0)
	for i := 0; i < contentValue.Len(); i++ {
		if content := contentValue.Index(i); cond(content.Interface()) {
			newContent = reflect.Append(newContent, content)
		}
	}
	return newContent.Interface()
}

func FindIfLabelExists(issue *github.Issue,label string) bool{
	for _,x := range issue.Labels{
		if *x.Name == label{
			return true
		}
	}
	return false
}

func GetTimeDifference(startTime time.Time, endTime time.Time)  time.Duration{
	return endTime.Sub(startTime)
}

func NewPostRequest(url string,body []byte){
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200{
		panic("Error While sending a hook")
	}
}