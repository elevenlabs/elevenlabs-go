package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ExportOptions struct {
	Format        string
	Docx          *DocxExportOptions
	Html          *HtmlExportOptions
	Pdf           *PdfExportOptions
	SegmentedJson *SegmentedJsonExportOptions
	Srt           *SrtExportOptions
	Txt           *TxtExportOptions
}

func (e *ExportOptions) GetDocx() *DocxExportOptions {
	if e == nil {
		return nil
	}
	if e.Docx == nil {
		return nil
	}
	return e.Docx
}

func (e *ExportOptions) GetHtml() *HtmlExportOptions {
	if e == nil {
		return nil
	}
	if e.Html == nil {
		return nil
	}
	return e.Html
}

func (e *ExportOptions) GetPdf() *PdfExportOptions {
	if e == nil {
		return nil
	}
	if e.Pdf == nil {
		return nil
	}
	return e.Pdf
}

func (e *ExportOptions) GetSegmentedJson() *SegmentedJsonExportOptions {
	if e == nil {
		return nil
	}
	if e.SegmentedJson == nil {
		return nil
	}
	return e.SegmentedJson
}

func (e *ExportOptions) GetSrt() *SrtExportOptions {
	if e == nil {
		return nil
	}
	if e.Srt == nil {
		return nil
	}
	return e.Srt
}

func (e *ExportOptions) GetTxt() *TxtExportOptions {
	if e == nil {
		return nil
	}
	if e.Txt == nil {
		return nil
	}
	return e.Txt
}

func (e *ExportOptions) GetFormat() string {
	if e == nil {
		return ""
	}
	return e.Format
}

func (e *ExportOptions) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Format string `json:"format"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Format = unmarshaler.Format
	if unmarshaler.Format == "" {
		return fmt.Errorf("%T did not include discriminant format", e)
	}
	switch unmarshaler.Format {
	case "docx":
		value := new(DocxExportOptions)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Docx = value
	case "html":
		value := new(HtmlExportOptions)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Html = value
	case "pdf":
		value := new(PdfExportOptions)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Pdf = value
	case "segmented_json":
		value := new(SegmentedJsonExportOptions)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.SegmentedJson = value
	case "srt":
		value := new(SrtExportOptions)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Srt = value
	case "txt":
		value := new(TxtExportOptions)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Txt = value
	}
	return nil
}

func (e ExportOptions) MarshalJSON() ([]byte, error) {
	if err := e.validate(); err != nil {
		return nil, err
	}
	if e.Docx != nil {
		return internal.MarshalJSONWithDiscriminant(e.Docx, "format", "docx")
	}
	if e.Html != nil {
		return internal.MarshalJSONWithDiscriminant(e.Html, "format", "html")
	}
	if e.Pdf != nil {
		return internal.MarshalJSONWithDiscriminant(e.Pdf, "format", "pdf")
	}
	if e.SegmentedJson != nil {
		return internal.MarshalJSONWithDiscriminant(e.SegmentedJson, "format", "segmented_json")
	}
	if e.Srt != nil {
		return internal.MarshalJSONWithDiscriminant(e.Srt, "format", "srt")
	}
	if e.Txt != nil {
		return internal.MarshalJSONWithDiscriminant(e.Txt, "format", "txt")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", e)
}

type ExportOptionsVisitor interface {
	VisitDocx(*DocxExportOptions) error
	VisitHtml(*HtmlExportOptions) error
	VisitPdf(*PdfExportOptions) error
	VisitSegmentedJson(*SegmentedJsonExportOptions) error
	VisitSrt(*SrtExportOptions) error
	VisitTxt(*TxtExportOptions) error
}

func (e *ExportOptions) Accept(visitor ExportOptionsVisitor) error {
	if e.Docx != nil {
		return visitor.VisitDocx(e.Docx)
	}
	if e.Html != nil {
		return visitor.VisitHtml(e.Html)
	}
	if e.Pdf != nil {
		return visitor.VisitPdf(e.Pdf)
	}
	if e.SegmentedJson != nil {
		return visitor.VisitSegmentedJson(e.SegmentedJson)
	}
	if e.Srt != nil {
		return visitor.VisitSrt(e.Srt)
	}
	if e.Txt != nil {
		return visitor.VisitTxt(e.Txt)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", e)
}

func (e *ExportOptions) validate() error {
	if e == nil {
		return fmt.Errorf("type %T is nil", e)
	}
	var fields []string
	if e.Docx != nil {
		fields = append(fields, "docx")
	}
	if e.Html != nil {
		fields = append(fields, "html")
	}
	if e.Pdf != nil {
		fields = append(fields, "pdf")
	}
	if e.SegmentedJson != nil {
		fields = append(fields, "segmented_json")
	}
	if e.Srt != nil {
		fields = append(fields, "srt")
	}
	if e.Txt != nil {
		fields = append(fields, "txt")
	}
	if len(fields) == 0 {
		if e.Format != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", e, e.Format)
		}
		return fmt.Errorf("type %T is empty", e)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", e, fields)
	}
	if e.Format != "" {
		field := fields[0]
		if e.Format != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				e,
				e.Format,
				e,
			)
		}
	}
	return nil
}
