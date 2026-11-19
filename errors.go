package elevenlabs

import (
	json "encoding/json"
	core "github.com/elevenlabs/elevenlabs-go/core"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

// Invalid request
type BadRequestError struct {
	*core.APIError
	Body any
}

func (b *BadRequestError) UnmarshalJSON(data []byte) error {
	var body any
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	b.StatusCode = 400
	b.Body = body
	return nil
}

func (b *BadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Body)
}

func (b *BadRequestError) Unwrap() error {
	return b.APIError
}

func (b *BadRequestError) GetBody() any {
	if b == nil {
		var zero any
		return zero
	}
	return b.Body
}

// Authentication required
type UnauthorizedError struct {
	*core.APIError
	Body any
}

func (u *UnauthorizedError) UnmarshalJSON(data []byte) error {
	var body any
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 401
	u.Body = body
	return nil
}

func (u *UnauthorizedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnauthorizedError) Unwrap() error {
	return u.APIError
}

func (u *UnauthorizedError) GetBody() any {
	if u == nil {
		var zero any
		return zero
	}
	return u.Body
}

// Permission denied
type ForbiddenError struct {
	*core.APIError
	Body any
}

func (f *ForbiddenError) UnmarshalJSON(data []byte) error {
	var body any
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	f.StatusCode = 403
	f.Body = body
	return nil
}

func (f *ForbiddenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Body)
}

func (f *ForbiddenError) Unwrap() error {
	return f.APIError
}

func (f *ForbiddenError) GetBody() any {
	if f == nil {
		var zero any
		return zero
	}
	return f.Body
}

// Dubbing not found
type NotFoundError struct {
	*core.APIError
	Body any
}

func (n *NotFoundError) UnmarshalJSON(data []byte) error {
	var body any
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	n.StatusCode = 404
	n.Body = body
	return nil
}

func (n *NotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Body)
}

func (n *NotFoundError) Unwrap() error {
	return n.APIError
}

func (n *NotFoundError) GetBody() any {
	if n == nil {
		var zero any
		return zero
	}
	return n.Body
}

// A generation for this preview is already in progress. Retry the request.
type ConflictError struct {
	*core.APIError
	Body any
}

func (c *ConflictError) UnmarshalJSON(data []byte) error {
	var body any
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	c.StatusCode = 409
	c.Body = body
	return nil
}

func (c *ConflictError) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Body)
}

func (c *ConflictError) Unwrap() error {
	return c.APIError
}

func (c *ConflictError) GetBody() any {
	if c == nil {
		var zero any
		return zero
	}
	return c.Body
}

// Validation Error
type UnprocessableEntityError struct {
	*core.APIError
	Body *types.HttpValidationError
}

func (u *UnprocessableEntityError) UnmarshalJSON(data []byte) error {
	var body *types.HttpValidationError
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 422
	u.Body = body
	return nil
}

func (u *UnprocessableEntityError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnprocessableEntityError) Unwrap() error {
	return u.APIError
}

func (u *UnprocessableEntityError) GetBody() *types.HttpValidationError {
	if u == nil {
		var zero *types.HttpValidationError
		return zero
	}
	return u.Body
}

// Dubbing not ready
type TooEarlyError struct {
	*core.APIError
	Body any
}

func (t *TooEarlyError) UnmarshalJSON(data []byte) error {
	var body any
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	t.StatusCode = 425
	t.Body = body
	return nil
}

func (t *TooEarlyError) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Body)
}

func (t *TooEarlyError) Unwrap() error {
	return t.APIError
}

func (t *TooEarlyError) GetBody() any {
	if t == nil {
		var zero any
		return zero
	}
	return t.Body
}
