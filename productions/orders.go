package productions

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
	"time"
)

type CreateOrderRequest struct {
	// When true, creates a sandbox order that auto-progresses without producer intervention.
	Sandbox *bool `json:"sandbox,omitempty" xml:"sandbox,omitempty"`
}

func (c *CreateOrderRequest) GetSandbox() (value *bool) {
	if c == nil {
		return
	}
	return c.Sandbox
}

func (c *CreateOrderRequest) SetSandbox(sandbox *bool) {
	c.Sandbox = sandbox
}

func (c *CreateOrderRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreateOrderRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateOrderRequest to string"
	}
	return string(jsonData)
}

func (c *CreateOrderRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateOrderRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateOrderRequest(value)
	return nil
}

func (c *CreateOrderRequest) MarshalJSON() ([]byte, error) {
	type embed CreateOrderRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

type BodyUpdateOrderV1ProductionsOrdersOrderIdPatch struct {
	// The ID of the order.
	OrderId OrderId                   `json:"-" url:"-"`
	Request *types.UpdateOrderRequest `json:"request" xml:"request" required:"true"`
}

func (b *BodyUpdateOrderV1ProductionsOrdersOrderIdPatch) GetOrderId() (value OrderId) {
	if b == nil {
		return
	}
	return b.OrderId
}

func (b *BodyUpdateOrderV1ProductionsOrdersOrderIdPatch) GetRequest() (value *types.UpdateOrderRequest) {
	if b == nil {
		return
	}
	return b.Request
}

func (b *BodyUpdateOrderV1ProductionsOrdersOrderIdPatch) SetOrderId(orderId OrderId) {
	b.OrderId = orderId
}

func (b *BodyUpdateOrderV1ProductionsOrdersOrderIdPatch) SetRequest(request *types.UpdateOrderRequest) {
	b.Request = request
}

func (b *BodyUpdateOrderV1ProductionsOrdersOrderIdPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdateOrderV1ProductionsOrdersOrderIdPatch) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdateOrderV1ProductionsOrdersOrderIdPatch to string"
	}
	return string(jsonData)
}

func (b *BodyUpdateOrderV1ProductionsOrdersOrderIdPatch) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyUpdateOrderV1ProductionsOrdersOrderIdPatch
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyUpdateOrderV1ProductionsOrdersOrderIdPatch(tmp)
	return nil
}

func (b *BodyUpdateOrderV1ProductionsOrdersOrderIdPatch) MarshalJSON() ([]byte, error) {
	type embed BodyUpdateOrderV1ProductionsOrdersOrderIdPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	listOrdersRequestFieldPageSize  = big.NewInt(1 << 0)
	listOrdersRequestFieldOffset    = big.NewInt(1 << 1)
	listOrdersRequestFieldStatus    = big.NewInt(1 << 2)
	listOrdersRequestFieldStartDate = big.NewInt(1 << 3)
	listOrdersRequestFieldEndDate   = big.NewInt(1 << 4)
)

type ListOrdersRequest struct {
	PageSize  *int                `json:"-" url:"page_size,omitempty"`
	Offset    *int                `json:"-" url:"offset,omitempty"`
	Status    []*types.OrderState `json:"-" url:"status,omitempty"`
	StartDate *time.Time          `json:"-" url:"start_date,omitempty"`
	EndDate   *time.Time          `json:"-" url:"end_date,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListOrdersRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOrdersRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listOrdersRequestFieldPageSize)
}

// SetOffset sets the Offset field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOrdersRequest) SetOffset(offset *int) {
	l.Offset = offset
	l.require(listOrdersRequestFieldOffset)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOrdersRequest) SetStatus(status []*types.OrderState) {
	l.Status = status
	l.require(listOrdersRequestFieldStatus)
}

// SetStartDate sets the StartDate field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOrdersRequest) SetStartDate(startDate *time.Time) {
	l.StartDate = startDate
	l.require(listOrdersRequestFieldStartDate)
}

// SetEndDate sets the EndDate field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListOrdersRequest) SetEndDate(endDate *time.Time) {
	l.EndDate = endDate
	l.require(listOrdersRequestFieldEndDate)
}
