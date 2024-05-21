package im

import (
	"github.com/wjp-letgo/letgo/lib"
	lazadaConfig "github.com/zhouzongyan/lazadago/config"
	messageEntity "github.com/zhouzongyan/lazadago/im/entity"
)

type Message struct {
	Config *lazadaConfig.Config
}

func (s *Message) GetSeesions(start, pageSize int, lastSeesionId string) messageEntity.GetSeesions {
	method := "/im/session/list"
	params := lib.InRow{
		"start_time": start,
		"page_size":  pageSize,
	}
	if lastSeesionId != "" {
		params = lib.InRow{
			"start_time":      start,
			"page_size":       pageSize,
			"last_session_id": lastSeesionId,
		}
	}
	result := messageEntity.GetSeesions{}
	err := s.Config.HttpGet(method, params, &result)
	if err != nil {
		result.Code = err.Error()
	}
	return result
}

func (s *Message) GetMessages(sessionId string, start, pageSize int, lastSeesionId string) messageEntity.GetMessages {
	method := "/im/message/list"
	params := lib.InRow{
		"session_id": sessionId,
		"start_time": start,
		"page_size":  pageSize,
	}
	if lastSeesionId != "" {
		params = lib.InRow{
			"session_id":      sessionId,
			"start_time":      start,
			"page_size":       pageSize,
			"last_message_id": lastSeesionId,
		}
	}
	result := messageEntity.GetMessages{}
	err := s.Config.HttpGet(method, params, &result)
	if err != nil {
		result.Code = err.Error()
	}
	return result
}
