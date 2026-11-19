package types

import (
	"encoding/json"
	"fmt"
)

type SendMessageMulti struct {
	InitializeConnectionMulti *InitializeConnectionMulti
	InitialiseContext         *InitialiseContext
	SendTextMulti             *SendTextMulti
	FlushContext              *FlushContext
	CloseContext              *CloseContext
	CloseSocket               *CloseSocket
	KeepContextAlive          *KeepContextAlive
	typ                       string
}

func (s *SendMessageMulti) GetInitializeConnectionMulti() *InitializeConnectionMulti {
	if s == nil {
		return nil
	}
	if s.InitializeConnectionMulti == nil {
		return nil
	}
	return s.InitializeConnectionMulti
}

func (s *SendMessageMulti) GetInitialiseContext() *InitialiseContext {
	if s == nil {
		return nil
	}
	if s.InitialiseContext == nil {
		return nil
	}
	return s.InitialiseContext
}

func (s *SendMessageMulti) GetSendTextMulti() *SendTextMulti {
	if s == nil {
		return nil
	}
	if s.SendTextMulti == nil {
		return nil
	}
	return s.SendTextMulti
}

func (s *SendMessageMulti) GetFlushContext() *FlushContext {
	if s == nil {
		return nil
	}
	if s.FlushContext == nil {
		return nil
	}
	return s.FlushContext
}

func (s *SendMessageMulti) GetCloseContext() *CloseContext {
	if s == nil {
		return nil
	}
	if s.CloseContext == nil {
		return nil
	}
	return s.CloseContext
}

func (s *SendMessageMulti) GetCloseSocket() *CloseSocket {
	if s == nil {
		return nil
	}
	if s.CloseSocket == nil {
		return nil
	}
	return s.CloseSocket
}

func (s *SendMessageMulti) GetKeepContextAlive() *KeepContextAlive {
	if s == nil {
		return nil
	}
	if s.KeepContextAlive == nil {
		return nil
	}
	return s.KeepContextAlive
}

func (s *SendMessageMulti) UnmarshalJSON(data []byte) error {
	var unionKeys map[string]json.RawMessage
	_ = json.Unmarshal(data, &unionKeys)
	hasAnyKey := func(keys ...string) bool {
		for _, key := range keys {
			if _, ok := unionKeys[key]; ok {
				return true
			}
		}
		return false
	}
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
		var valueInitializeConnectionMultiLiteral *InitializeConnectionMulti
		if err := json.Unmarshal(data, &valueInitializeConnectionMultiLiteral); err == nil {
			s.typ = "InitializeConnectionMulti"
			s.InitializeConnectionMulti = valueInitializeConnectionMultiLiteral
			return nil
		}
	}
	if hasLiteral("text", "") {
		matchedLiteral = true
		var valueKeepContextAliveLiteral *KeepContextAlive
		if err := json.Unmarshal(data, &valueKeepContextAliveLiteral); err == nil {
			s.typ = "KeepContextAlive"
			s.KeepContextAlive = valueKeepContextAliveLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, s)
	}
	var valueInitializeConnectionMulti *InitializeConnectionMulti
	if err := json.Unmarshal(data, &valueInitializeConnectionMulti); err == nil {
		s.typ = "InitializeConnectionMulti"
		s.InitializeConnectionMulti = valueInitializeConnectionMulti
		return nil
	}
	var valueInitialiseContext *InitialiseContext
	if err := json.Unmarshal(data, &valueInitialiseContext); err == nil {
		s.typ = "InitialiseContext"
		s.InitialiseContext = valueInitialiseContext
		return nil
	}
	var valueSendTextMulti *SendTextMulti
	if err := json.Unmarshal(data, &valueSendTextMulti); err == nil {
		s.typ = "SendTextMulti"
		s.SendTextMulti = valueSendTextMulti
		return nil
	}
	var valueFlushContext *FlushContext
	if err := json.Unmarshal(data, &valueFlushContext); err == nil {
		s.typ = "FlushContext"
		s.FlushContext = valueFlushContext
		return nil
	}
	var valueCloseContext *CloseContext
	if err := json.Unmarshal(data, &valueCloseContext); err == nil {
		s.typ = "CloseContext"
		s.CloseContext = valueCloseContext
		return nil
	}
	if hasAnyKey("close_socket") {
		var valueCloseSocket *CloseSocket
		if err := json.Unmarshal(data, &valueCloseSocket); err == nil {
			s.typ = "CloseSocket"
			s.CloseSocket = valueCloseSocket
			return nil
		}
	}
	var valueKeepContextAlive *KeepContextAlive
	if err := json.Unmarshal(data, &valueKeepContextAlive); err == nil {
		s.typ = "KeepContextAlive"
		s.KeepContextAlive = valueKeepContextAlive
		return nil
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueCloseSocketFallback *CloseSocket
	if err := json.Unmarshal(data, &valueCloseSocketFallback); err == nil {
		s.typ = "CloseSocket"
		s.CloseSocket = valueCloseSocketFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, s)
}

func (s SendMessageMulti) MarshalJSON() ([]byte, error) {
	if s.typ == "InitializeConnectionMulti" || s.InitializeConnectionMulti != nil {
		return json.Marshal(s.InitializeConnectionMulti)
	}
	if s.typ == "InitialiseContext" || s.InitialiseContext != nil {
		return json.Marshal(s.InitialiseContext)
	}
	if s.typ == "SendTextMulti" || s.SendTextMulti != nil {
		return json.Marshal(s.SendTextMulti)
	}
	if s.typ == "FlushContext" || s.FlushContext != nil {
		return json.Marshal(s.FlushContext)
	}
	if s.typ == "CloseContext" || s.CloseContext != nil {
		return json.Marshal(s.CloseContext)
	}
	if s.typ == "CloseSocket" || s.CloseSocket != nil {
		return json.Marshal(s.CloseSocket)
	}
	if s.typ == "KeepContextAlive" || s.KeepContextAlive != nil {
		return json.Marshal(s.KeepContextAlive)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", s)
}

type SendMessageMultiVisitor interface {
	VisitInitializeConnectionMulti(*InitializeConnectionMulti) error
	VisitInitialiseContext(*InitialiseContext) error
	VisitSendTextMulti(*SendTextMulti) error
	VisitFlushContext(*FlushContext) error
	VisitCloseContext(*CloseContext) error
	VisitCloseSocket(*CloseSocket) error
	VisitKeepContextAlive(*KeepContextAlive) error
}

func (s *SendMessageMulti) Accept(visitor SendMessageMultiVisitor) error {
	if s.typ == "InitializeConnectionMulti" || s.InitializeConnectionMulti != nil {
		return visitor.VisitInitializeConnectionMulti(s.InitializeConnectionMulti)
	}
	if s.typ == "InitialiseContext" || s.InitialiseContext != nil {
		return visitor.VisitInitialiseContext(s.InitialiseContext)
	}
	if s.typ == "SendTextMulti" || s.SendTextMulti != nil {
		return visitor.VisitSendTextMulti(s.SendTextMulti)
	}
	if s.typ == "FlushContext" || s.FlushContext != nil {
		return visitor.VisitFlushContext(s.FlushContext)
	}
	if s.typ == "CloseContext" || s.CloseContext != nil {
		return visitor.VisitCloseContext(s.CloseContext)
	}
	if s.typ == "CloseSocket" || s.CloseSocket != nil {
		return visitor.VisitCloseSocket(s.CloseSocket)
	}
	if s.typ == "KeepContextAlive" || s.KeepContextAlive != nil {
		return visitor.VisitKeepContextAlive(s.KeepContextAlive)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", s)
}
