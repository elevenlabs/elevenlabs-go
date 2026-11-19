package procedures

import (
	elevenlabs "github.com/elevenlabs/elevenlabs-go"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
)

var ErrorCodes internal.ErrorCodes = internal.ErrorCodes{
	400: func(apiError *core.APIError) error {
		return &elevenlabs.BadRequestError{
			APIError: apiError,
		}
	},
	401: func(apiError *core.APIError) error {
		return &elevenlabs.UnauthorizedError{
			APIError: apiError,
		}
	},
	403: func(apiError *core.APIError) error {
		return &elevenlabs.ForbiddenError{
			APIError: apiError,
		}
	},
	404: func(apiError *core.APIError) error {
		return &elevenlabs.NotFoundError{
			APIError: apiError,
		}
	},
	409: func(apiError *core.APIError) error {
		return &elevenlabs.ConflictError{
			APIError: apiError,
		}
	},
	422: func(apiError *core.APIError) error {
		return &elevenlabs.UnprocessableEntityError{
			APIError: apiError,
		}
	},
	425: func(apiError *core.APIError) error {
		return &elevenlabs.TooEarlyError{
			APIError: apiError,
		}
	},
}
