package elevenlabs

import (
	"errors"
	"testing"

	core "github.com/elevenlabs/elevenlabs-go/core"
)

func TestBadRequestErrorErrorCode(t *testing.T) {
	t.Parallel()
	build, ok := ErrorCodes[400]
	if !ok {
		t.Fatalf("no error is registered for status 400")
	}
	err := build(&core.APIError{StatusCode: 400})
	var typed *BadRequestError
	if !errors.As(err, &typed) {
		t.Fatalf("status 400 built %T, want *BadRequestError", err)
	}
	if typed.StatusCode != 400 {
		t.Fatalf("status code: got %d want 400", typed.StatusCode)
	}
}

func TestUnauthorizedErrorErrorCode(t *testing.T) {
	t.Parallel()
	build, ok := ErrorCodes[401]
	if !ok {
		t.Fatalf("no error is registered for status 401")
	}
	err := build(&core.APIError{StatusCode: 401})
	var typed *UnauthorizedError
	if !errors.As(err, &typed) {
		t.Fatalf("status 401 built %T, want *UnauthorizedError", err)
	}
	if typed.StatusCode != 401 {
		t.Fatalf("status code: got %d want 401", typed.StatusCode)
	}
}

func TestForbiddenErrorErrorCode(t *testing.T) {
	t.Parallel()
	build, ok := ErrorCodes[403]
	if !ok {
		t.Fatalf("no error is registered for status 403")
	}
	err := build(&core.APIError{StatusCode: 403})
	var typed *ForbiddenError
	if !errors.As(err, &typed) {
		t.Fatalf("status 403 built %T, want *ForbiddenError", err)
	}
	if typed.StatusCode != 403 {
		t.Fatalf("status code: got %d want 403", typed.StatusCode)
	}
}

func TestNotFoundErrorErrorCode(t *testing.T) {
	t.Parallel()
	build, ok := ErrorCodes[404]
	if !ok {
		t.Fatalf("no error is registered for status 404")
	}
	err := build(&core.APIError{StatusCode: 404})
	var typed *NotFoundError
	if !errors.As(err, &typed) {
		t.Fatalf("status 404 built %T, want *NotFoundError", err)
	}
	if typed.StatusCode != 404 {
		t.Fatalf("status code: got %d want 404", typed.StatusCode)
	}
}

func TestConflictErrorErrorCode(t *testing.T) {
	t.Parallel()
	build, ok := ErrorCodes[409]
	if !ok {
		t.Fatalf("no error is registered for status 409")
	}
	err := build(&core.APIError{StatusCode: 409})
	var typed *ConflictError
	if !errors.As(err, &typed) {
		t.Fatalf("status 409 built %T, want *ConflictError", err)
	}
	if typed.StatusCode != 409 {
		t.Fatalf("status code: got %d want 409", typed.StatusCode)
	}
}

func TestUnprocessableEntityErrorErrorCode(t *testing.T) {
	t.Parallel()
	build, ok := ErrorCodes[422]
	if !ok {
		t.Fatalf("no error is registered for status 422")
	}
	err := build(&core.APIError{StatusCode: 422})
	var typed *UnprocessableEntityError
	if !errors.As(err, &typed) {
		t.Fatalf("status 422 built %T, want *UnprocessableEntityError", err)
	}
	if typed.StatusCode != 422 {
		t.Fatalf("status code: got %d want 422", typed.StatusCode)
	}
}

func TestTooEarlyErrorErrorCode(t *testing.T) {
	t.Parallel()
	build, ok := ErrorCodes[425]
	if !ok {
		t.Fatalf("no error is registered for status 425")
	}
	err := build(&core.APIError{StatusCode: 425})
	var typed *TooEarlyError
	if !errors.As(err, &typed) {
		t.Fatalf("status 425 built %T, want *TooEarlyError", err)
	}
	if typed.StatusCode != 425 {
		t.Fatalf("status code: got %d want 425", typed.StatusCode)
	}
}
