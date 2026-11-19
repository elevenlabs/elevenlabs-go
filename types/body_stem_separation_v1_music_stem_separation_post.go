package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyStemSeparationV1MusicStemSeparationPost struct {
	// When enable_logging is set to false zero retention mode will be used for the request. Zero retention mode may only be used by enterprise customers.
	EnableLogging *bool `json:"-" url:"enable_logging,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_logging"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
	OutputFormat *AllowedOutputFormats `json:"-" url:"output_format,omitempty" explode:"true" serializationStyle:"form" queryParam:"output_format"`
	// The audio file to separate into stems.
	File io.Reader `json:"file" xml:"file" required:"true"`
	// The id of the stem variation to use.
	StemVariationId *SeparateStemsMusicRequestStemVariationId `json:"stem_variation_id,omitempty" xml:"stem_variation_id,omitempty"`
	// Whether to sign the generated song with C2PA. Applicable only for mp3 files.
	SignWithC2Pa *bool `json:"sign_with_c2pa,omitempty" xml:"sign_with_c2pa,omitempty"`
}

func (b *BodyStemSeparationV1MusicStemSeparationPost) GetEnableLogging() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableLogging
}

func (b *BodyStemSeparationV1MusicStemSeparationPost) GetOutputFormat() (value *AllowedOutputFormats) {
	if b == nil {
		return
	}
	return b.OutputFormat
}

func (b *BodyStemSeparationV1MusicStemSeparationPost) GetFile() (value io.Reader) {
	if b == nil {
		return
	}
	return b.File
}

func (b *BodyStemSeparationV1MusicStemSeparationPost) GetStemVariationId() (value *SeparateStemsMusicRequestStemVariationId) {
	if b == nil {
		return
	}
	return b.StemVariationId
}

func (b *BodyStemSeparationV1MusicStemSeparationPost) GetSignWithC2Pa() (value *bool) {
	if b == nil {
		return
	}
	return b.SignWithC2Pa
}

func (b *BodyStemSeparationV1MusicStemSeparationPost) SetEnableLogging(enableLogging *bool) {
	b.EnableLogging = enableLogging
}

func (b *BodyStemSeparationV1MusicStemSeparationPost) SetOutputFormat(outputFormat *AllowedOutputFormats) {
	b.OutputFormat = outputFormat
}

func (b *BodyStemSeparationV1MusicStemSeparationPost) SetFile(file io.Reader) {
	b.File = file
}

func (b *BodyStemSeparationV1MusicStemSeparationPost) SetStemVariationId(stemVariationId *SeparateStemsMusicRequestStemVariationId) {
	b.StemVariationId = stemVariationId
}

func (b *BodyStemSeparationV1MusicStemSeparationPost) SetSignWithC2Pa(signWithC2Pa *bool) {
	b.SignWithC2Pa = signWithC2Pa
}

func (b *BodyStemSeparationV1MusicStemSeparationPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyStemSeparationV1MusicStemSeparationPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyStemSeparationV1MusicStemSeparationPost to string"
	}
	return string(jsonData)
}

func (b *BodyStemSeparationV1MusicStemSeparationPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyStemSeparationV1MusicStemSeparationPost) MarshalJSON() ([]byte, error) {
	type embed BodyStemSeparationV1MusicStemSeparationPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
