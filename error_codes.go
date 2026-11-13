package elevenlabs

import (
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
)

var ErrorCodes internal.ErrorCodes = internal.ErrorCodes{
	400: func(apiError *core.APIError) error {
		return &BadRequestError{
			APIError: apiError,
		}
	},
	401: func(apiError *core.APIError) error {
		return &UnauthorizedError{
			APIError: apiError,
		}
	},
	403: func(apiError *core.APIError) error {
		return &ForbiddenError{
			APIError: apiError,
		}
	},
	404: func(apiError *core.APIError) error {
		return &NotFoundError{
			APIError: apiError,
		}
	},
	409: func(apiError *core.APIError) error {
		return &ConflictError{
			APIError: apiError,
		}
	},
	422: func(apiError *core.APIError) error {
		return &UnprocessableEntityError{
			APIError: apiError,
		}
	},
	425: func(apiError *core.APIError) error {
		return &TooEarlyError{
			APIError: apiError,
		}
	},
}
