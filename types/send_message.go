package types

import (
	"encoding/json"
	"fmt"
)

type SendMessage struct {
	InitializeConnection *InitializeConnection
	SendText             *SendText
	CloseConnection      *CloseConnection
	typ                  string
}

func (s *SendMessage) GetInitializeConnection() *InitializeConnection {
	if s == nil {
		return nil
	}
	if s.InitializeConnection == nil {
		return nil
	}
	return s.InitializeConnection
}

func (s *SendMessage) GetSendText() *SendText {
	if s == nil {
		return nil
	}
	if s.SendText == nil {
		return nil
	}
	return s.SendText
}

func (s *SendMessage) GetCloseConnection() *CloseConnection {
	if s == nil {
		return nil
	}
	if s.CloseConnection == nil {
		return nil
	}
	return s.CloseConnection
}

func (s *SendMessage) UnmarshalJSON(data []byte) error {
	var unionKeys map[string]json.RawMessage
	_ = json.Unmarshal(data, &unionKeys)
	hasLiteral := func(key string, want interface{}) bool {
		raw, ok := unionKeys[key]
		if !ok {
			return false
		}
		var got interface{}
		return json.Unmarshal(raw, &got) == nil && got == want
	}
	matchedLiteral := false
	if hasLiteral("text", " ") {
		matchedLiteral = true
		var valueInitializeConnectionLiteral *InitializeConnection
		if err := json.Unmarshal(data, &valueInitializeConnectionLiteral); err == nil {
			s.typ = "InitializeConnection"
			s.InitializeConnection = valueInitializeConnectionLiteral
			return nil
		}
	}
	if hasLiteral("text", "") {
		matchedLiteral = true
		var valueCloseConnectionLiteral *CloseConnection
		if err := json.Unmarshal(data, &valueCloseConnectionLiteral); err == nil {
			s.typ = "CloseConnection"
			s.CloseConnection = valueCloseConnectionLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, s)
	}
	var valueInitializeConnection *InitializeConnection
	if err := json.Unmarshal(data, &valueInitializeConnection); err == nil {
		s.typ = "InitializeConnection"
		s.InitializeConnection = valueInitializeConnection
		return nil
	}
	var valueSendText *SendText
	if err := json.Unmarshal(data, &valueSendText); err == nil {
		s.typ = "SendText"
		s.SendText = valueSendText
		return nil
	}
	var valueCloseConnection *CloseConnection
	if err := json.Unmarshal(data, &valueCloseConnection); err == nil {
		s.typ = "CloseConnection"
		s.CloseConnection = valueCloseConnection
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, s)
}

func (s SendMessage) MarshalJSON() ([]byte, error) {
	if s.typ == "InitializeConnection" || s.InitializeConnection != nil {
		return json.Marshal(s.InitializeConnection)
	}
	if s.typ == "SendText" || s.SendText != nil {
		return json.Marshal(s.SendText)
	}
	if s.typ == "CloseConnection" || s.CloseConnection != nil {
		return json.Marshal(s.CloseConnection)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", s)
}

type SendMessageVisitor interface {
	VisitInitializeConnection(*InitializeConnection) error
	VisitSendText(*SendText) error
	VisitCloseConnection(*CloseConnection) error
}

func (s *SendMessage) Accept(visitor SendMessageVisitor) error {
	if s.typ == "InitializeConnection" || s.InitializeConnection != nil {
		return visitor.VisitInitializeConnection(s.InitializeConnection)
	}
	if s.typ == "SendText" || s.SendText != nil {
		return visitor.VisitSendText(s.SendText)
	}
	if s.typ == "CloseConnection" || s.CloseConnection != nil {
		return visitor.VisitCloseConnection(s.CloseConnection)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", s)
}
