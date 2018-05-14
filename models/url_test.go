package models

import (
	"testing"
	"gopkg.in/mgo.v2"
	"github.com/stretchr/testify/assert"
)

func init() {
	mongoUrl := "127.0.0.1:3000"
	mongoDatabase := "short_url_test"
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		panic(err)
	}
	db = session.DB(mongoDatabase)
}
func TestMDB(t *testing.T) {
	var url = Url{
		Id:        1,
		SourceUrl: "http://www.qq.com",
	}

	// insert
	err := url.Insert()
	assert.Nil(t, err)

	// update
	url.SourceUrl = "http://www.weixin.com"
	err = url.Update()
	assert.Nil(t, err)

	// find
	err = url.FindById()
	assert.Nil(t, err)
	assert.Equal(t, url.SourceUrl, "http://www.weixin.com")

	//delete
	err = url.DeleteById()
	assert.Nil(t, err)
}
