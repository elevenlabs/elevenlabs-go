package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WebhookSendResultResponseModel struct {
	// The delivery ID, also sent in the ElevenLabs-Webhook-Delivery-Id header.
	DeliveryId string         `json:"delivery_id" xml:"delivery_id" required:"true"`
	Outcome    WebhookOutcome `json:"outcome" xml:"outcome" required:"true"`
	// The HTTP status your endpoint returned, if it answered.
	ResponseCode *int `json:"response_code" xml:"response_code" required:"true"`
	// How long your endpoint took to answer, in milliseconds.
	LatencyMs *int `json:"latency_ms" xml:"latency_ms" required:"true"`
	// The Content-Type your endpoint answered with, if it answered.
	ResponseContentType *string `json:"response_content_type" xml:"response_content_type" required:"true"`
	// The size of your endpoint's response body in bytes, if it answered. The body itself is never stored or returned.
	ResponseBodyBytes *int `json:"response_body_bytes" xml:"response_body_bytes" required:"true"`
	// Where a failure came from: recipient, configuration, external_transport or timeout.
	FailureOrigin   *string `json:"failure_origin" xml:"failure_origin" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WebhookSendResultResponseModel) GetDeliveryId() (value string) {
	if w == nil {
		return
	}
	return w.DeliveryId
}

func (w *WebhookSendResultResponseModel) GetOutcome() (value WebhookOutcome) {
	if w == nil {
		return
	}
	return w.Outcome
}

func (w *WebhookSendResultResponseModel) GetResponseCode() (value *int) {
	if w == nil {
		return
	}
	return w.ResponseCode
}

func (w *WebhookSendResultResponseModel) GetLatencyMs() (value *int) {
	if w == nil {
		return
	}
	return w.LatencyMs
}

func (w *WebhookSendResultResponseModel) GetResponseContentType() (value *string) {
	if w == nil {
		return
	}
	return w.ResponseContentType
}

func (w *WebhookSendResultResponseModel) GetResponseBodyBytes() (value *int) {
	if w == nil {
		return
	}
	return w.ResponseBodyBytes
}

func (w *WebhookSendResultResponseModel) GetFailureOrigin() (value *string) {
	if w == nil {
		return
	}
	return w.FailureOrigin
}

func (w *WebhookSendResultResponseModel) SetDeliveryId(deliveryId string) {
	w.DeliveryId = deliveryId
}

func (w *WebhookSendResultResponseModel) SetOutcome(outcome WebhookOutcome) {
	w.Outcome = outcome
}

func (w *WebhookSendResultResponseModel) SetResponseCode(responseCode *int) {
	w.ResponseCode = responseCode
}

func (w *WebhookSendResultResponseModel) SetLatencyMs(latencyMs *int) {
	w.LatencyMs = latencyMs
}

func (w *WebhookSendResultResponseModel) SetResponseContentType(responseContentType *string) {
	w.ResponseContentType = responseContentType
}

func (w *WebhookSendResultResponseModel) SetResponseBodyBytes(responseBodyBytes *int) {
	w.ResponseBodyBytes = responseBodyBytes
}

func (w *WebhookSendResultResponseModel) SetFailureOrigin(failureOrigin *string) {
	w.FailureOrigin = failureOrigin
}

func (w *WebhookSendResultResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WebhookSendResultResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebhookSendResultResponseModel to string"
	}
	return string(jsonData)
}

func (w *WebhookSendResultResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WebhookSendResultResponseModel) MarshalJSON() ([]byte, error) {
	type embed WebhookSendResultResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
