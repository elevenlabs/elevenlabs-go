package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Defines asset positioning and transformation on canvas.
type CanvasPlacement struct {
	XRelative       *float64 `json:"x_relative,omitempty" xml:"x_relative,omitempty"`
	YRelative       *float64 `json:"y_relative,omitempty" xml:"y_relative,omitempty"`
	ScaleX          *float64 `json:"scale_x,omitempty" xml:"scale_x,omitempty"`
	ScaleY          *float64 `json:"scale_y,omitempty" xml:"scale_y,omitempty"`
	PivotX          *float64 `json:"pivot_x,omitempty" xml:"pivot_x,omitempty"`
	PivotY          *float64 `json:"pivot_y,omitempty" xml:"pivot_y,omitempty"`
	SkewX           *float64 `json:"skew_x,omitempty" xml:"skew_x,omitempty"`
	SkewY           *float64 `json:"skew_y,omitempty" xml:"skew_y,omitempty"`
	CropTop         *float64 `json:"crop_top,omitempty" xml:"crop_top,omitempty"`
	CropRight       *float64 `json:"crop_right,omitempty" xml:"crop_right,omitempty"`
	CropBottom      *float64 `json:"crop_bottom,omitempty" xml:"crop_bottom,omitempty"`
	CropLeft        *float64 `json:"crop_left,omitempty" xml:"crop_left,omitempty"`
	FlipX           *bool    `json:"flip_x,omitempty" xml:"flip_x,omitempty"`
	FlipY           *bool    `json:"flip_y,omitempty" xml:"flip_y,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CanvasPlacement) GetXRelative() (value *float64) {
	if c == nil {
		return
	}
	return c.XRelative
}

func (c *CanvasPlacement) GetYRelative() (value *float64) {
	if c == nil {
		return
	}
	return c.YRelative
}

func (c *CanvasPlacement) GetScaleX() (value *float64) {
	if c == nil {
		return
	}
	return c.ScaleX
}

func (c *CanvasPlacement) GetScaleY() (value *float64) {
	if c == nil {
		return
	}
	return c.ScaleY
}

func (c *CanvasPlacement) GetPivotX() (value *float64) {
	if c == nil {
		return
	}
	return c.PivotX
}

func (c *CanvasPlacement) GetPivotY() (value *float64) {
	if c == nil {
		return
	}
	return c.PivotY
}

func (c *CanvasPlacement) GetSkewX() (value *float64) {
	if c == nil {
		return
	}
	return c.SkewX
}

func (c *CanvasPlacement) GetSkewY() (value *float64) {
	if c == nil {
		return
	}
	return c.SkewY
}

func (c *CanvasPlacement) GetCropTop() (value *float64) {
	if c == nil {
		return
	}
	return c.CropTop
}

func (c *CanvasPlacement) GetCropRight() (value *float64) {
	if c == nil {
		return
	}
	return c.CropRight
}

func (c *CanvasPlacement) GetCropBottom() (value *float64) {
	if c == nil {
		return
	}
	return c.CropBottom
}

func (c *CanvasPlacement) GetCropLeft() (value *float64) {
	if c == nil {
		return
	}
	return c.CropLeft
}

func (c *CanvasPlacement) GetFlipX() (value *bool) {
	if c == nil {
		return
	}
	return c.FlipX
}

func (c *CanvasPlacement) GetFlipY() (value *bool) {
	if c == nil {
		return
	}
	return c.FlipY
}

func (c *CanvasPlacement) SetXRelative(xRelative *float64) {
	c.XRelative = xRelative
}

func (c *CanvasPlacement) SetYRelative(yRelative *float64) {
	c.YRelative = yRelative
}

func (c *CanvasPlacement) SetScaleX(scaleX *float64) {
	c.ScaleX = scaleX
}

func (c *CanvasPlacement) SetScaleY(scaleY *float64) {
	c.ScaleY = scaleY
}

func (c *CanvasPlacement) SetPivotX(pivotX *float64) {
	c.PivotX = pivotX
}

func (c *CanvasPlacement) SetPivotY(pivotY *float64) {
	c.PivotY = pivotY
}

func (c *CanvasPlacement) SetSkewX(skewX *float64) {
	c.SkewX = skewX
}

func (c *CanvasPlacement) SetSkewY(skewY *float64) {
	c.SkewY = skewY
}

func (c *CanvasPlacement) SetCropTop(cropTop *float64) {
	c.CropTop = cropTop
}

func (c *CanvasPlacement) SetCropRight(cropRight *float64) {
	c.CropRight = cropRight
}

func (c *CanvasPlacement) SetCropBottom(cropBottom *float64) {
	c.CropBottom = cropBottom
}

func (c *CanvasPlacement) SetCropLeft(cropLeft *float64) {
	c.CropLeft = cropLeft
}

func (c *CanvasPlacement) SetFlipX(flipX *bool) {
	c.FlipX = flipX
}

func (c *CanvasPlacement) SetFlipY(flipY *bool) {
	c.FlipY = flipY
}

func (c *CanvasPlacement) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CanvasPlacement) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CanvasPlacement to string"
	}
	return string(jsonData)
}

func (c *CanvasPlacement) UnmarshalJSON(data []byte) error {
	type unmarshaler CanvasPlacement
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CanvasPlacement(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CanvasPlacement) MarshalJSON() ([]byte, error) {
	type embed CanvasPlacement
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
