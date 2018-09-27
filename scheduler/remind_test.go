package main

import (
	"testing"

	"encoding/json"

	"eglass.com/utils"
	"github.com/stretchr/testify/assert"
)

func TestLoadTasks(t *testing.T) {
	sess, error := utils.NewMysql(false, true)
	if error != nil {
		panic("mysql con error")
	}
	defer sess.Close()
	er := ERemindRecord{
		Id: 128,
	}
	error = er.Run(sess)
	assert.Equal(t, nil, error)
}

func TestUnMarshal(t *testing.T) {
	type M struct {
		ID int `json:"id"`
	}
	s := `{"id":1}`
	var m M
	json.Unmarshal([]byte(s), &m)
	assert.Equal(t, m, M{1})
}
