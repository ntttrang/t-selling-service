package messagehandler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/kataras/golog"
)

//Init init message error
func Init() {
	loadMessages()
}

type (
	messageStore = map[string]Mess
	//MessCode messCode
	MessCode string
	//Mess define message error struct
	Mess struct {
		Code       string `json:"code"`
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
	}
)

var messStoreLck = sync.RWMutex{}
var messStore = messageStore{}

func loadMessages() {
	var tempStruct = &struct {
		Messages []Mess `json:"messages"`
	}{}

	// Read json in messages.json.go
	err := json.Unmarshal([]byte(messagesJSON), &tempStruct)
	if nil != err {
		golog.Fatal(err)
	}

	messStoreLck.Lock()
	defer messStoreLck.Unlock()

	for _, message := range tempStruct.Messages {
		messStore[message.Code] = message
	}
}

// GetStatusCode Get status code
func (mc MessCode) GetStatusCode(params ...interface{}) (int, error) {
	messStoreLck.RLock()
	defer messStoreLck.Unlock()

	if m, ok := messStore[mc.ToString()]; ok {
		return m.StatusCode, nil
	}
	return 0, fmt.Errorf("Message code %s not found", mc)

}

//GetMessage get error message
func (mc MessCode) GetMessage(params ...interface{}) *Mess {
	messStoreLck.RLock()
	defer messStoreLck.RUnlock()

	if m, ok := messStore[mc.ToString()]; ok {
		msgStr, e := mc.GetMessageStr(params...)
		if nil != e {
			golog.Error(e)
		}
		m.Message = msgStr
		return &m
	}
	return nil
}

// GetMessageStr Get message str
func (mc MessCode) GetMessageStr(params ...interface{}) (string, error) {
	messStoreLck.RLock()
	defer messStoreLck.RUnlock()

	if m, ok := messStore[mc.ToString()]; ok {
		if nil != params && len(params) > 0 && fmt.Sprintf("%v", params...) != "[]" {
			return fmt.Sprintf(m.Message, params...), nil
		}
		return fmt.Sprint(m.Message), nil
	}

	return "", fmt.Errorf("Message code %s not found", mc)
}

//ToString convert MessCode to string
func (mc MessCode) ToString() string {
	return fmt.Sprint(mc)
}

const (
	E000001 MessCode = "E000001"
	W000001 MessCode = "W000001"
	W000002 MessCode = "W000002"
	W000003 MessCode = "W000003"
	W000004 MessCode = "W000004"
	W000005 MessCode = "W000005"
	W000006 MessCode = "W000006"
	W000007 MessCode = "W000007"
	W000008 MessCode = "W000008"
	W000009 MessCode = "W000009"

	I000008 MessCode = "I000008"
)
