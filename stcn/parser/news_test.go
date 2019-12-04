package parser

import (
	"NewsSpider/model"
	"io/ioutil"
	"testing"
)

func TestNewsParser(t *testing.T) {
	contents, err := ioutil.ReadFile("news_content_data.html")
	if err != nil {
		panic(err)
	}

	result := NewsParser(contents)

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1" + "element; but was %v", result.Items)
	}
	profile := result.Items[0].(model.NewsFields)
	expected := model.NewsFields{
		Title:       "海南省明年10月1日起实施生活垃圾分类_证券时报网",
		Content:     "据海南日报12月3日消息，12月2日，海南省新闻办在海口召开《海南省生活垃圾管理条例》（以下简称《条例》）新闻发布会，海南日报记者从会上获悉，海南省将从2020年10月1日起实施生活垃圾分类，对不按规定分类投放生活垃圾的单位将处以5000元以上5万元以下罚款，个人将处以200元以下罚款。",
		PublishTime: "2019-12-03 09:48",
	}
	if profile != expected {
		t.Errorf("expected %v\n; but was %v", expected, profile)
	}
}
