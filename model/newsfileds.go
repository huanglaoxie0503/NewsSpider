package model

import "encoding/json"

type NewsFields struct {
	Title       string
	Content 	string
	PublishTime string
}

func FromJsonObj(o interface{}) (NewsFields, error)  {
	var newsFields NewsFields
	s, err := json.Marshal(o)
	if err != nil {
		return newsFields, err
	}
	err = json.Unmarshal(s, &newsFields)
	return newsFields, err
}