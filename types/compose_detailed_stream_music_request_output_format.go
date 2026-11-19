package types

import (
	"fmt"
)

type ComposeDetailedStreamMusicRequestOutputFormat string

const (
	ComposeDetailedStreamMusicRequestOutputFormatAuto         ComposeDetailedStreamMusicRequestOutputFormat = "auto"
	ComposeDetailedStreamMusicRequestOutputFormatMp348000128  ComposeDetailedStreamMusicRequestOutputFormat = "mp3_48000_128"
	ComposeDetailedStreamMusicRequestOutputFormatMp348000192  ComposeDetailedStreamMusicRequestOutputFormat = "mp3_48000_192"
	ComposeDetailedStreamMusicRequestOutputFormatMp348000240  ComposeDetailedStreamMusicRequestOutputFormat = "mp3_48000_240"
	ComposeDetailedStreamMusicRequestOutputFormatMp348000320  ComposeDetailedStreamMusicRequestOutputFormat = "mp3_48000_320"
	ComposeDetailedStreamMusicRequestOutputFormatMp32205032   ComposeDetailedStreamMusicRequestOutputFormat = "mp3_22050_32"
	ComposeDetailedStreamMusicRequestOutputFormatMp32400048   ComposeDetailedStreamMusicRequestOutputFormat = "mp3_24000_48"
	ComposeDetailedStreamMusicRequestOutputFormatMp34410032   ComposeDetailedStreamMusicRequestOutputFormat = "mp3_44100_32"
	ComposeDetailedStreamMusicRequestOutputFormatMp34410064   ComposeDetailedStreamMusicRequestOutputFormat = "mp3_44100_64"
	ComposeDetailedStreamMusicRequestOutputFormatMp34410096   ComposeDetailedStreamMusicRequestOutputFormat = "mp3_44100_96"
	ComposeDetailedStreamMusicRequestOutputFormatMp344100128  ComposeDetailedStreamMusicRequestOutputFormat = "mp3_44100_128"
	ComposeDetailedStreamMusicRequestOutputFormatMp344100192  ComposeDetailedStreamMusicRequestOutputFormat = "mp3_44100_192"
	ComposeDetailedStreamMusicRequestOutputFormatPcm8000      ComposeDetailedStreamMusicRequestOutputFormat = "pcm_8000"
	ComposeDetailedStreamMusicRequestOutputFormatPcm16000     ComposeDetailedStreamMusicRequestOutputFormat = "pcm_16000"
	ComposeDetailedStreamMusicRequestOutputFormatPcm22050     ComposeDetailedStreamMusicRequestOutputFormat = "pcm_22050"
	ComposeDetailedStreamMusicRequestOutputFormatPcm24000     ComposeDetailedStreamMusicRequestOutputFormat = "pcm_24000"
	ComposeDetailedStreamMusicRequestOutputFormatPcm32000     ComposeDetailedStreamMusicRequestOutputFormat = "pcm_32000"
	ComposeDetailedStreamMusicRequestOutputFormatPcm44100     ComposeDetailedStreamMusicRequestOutputFormat = "pcm_44100"
	ComposeDetailedStreamMusicRequestOutputFormatPcm48000     ComposeDetailedStreamMusicRequestOutputFormat = "pcm_48000"
	ComposeDetailedStreamMusicRequestOutputFormatUlaw8000     ComposeDetailedStreamMusicRequestOutputFormat = "ulaw_8000"
	ComposeDetailedStreamMusicRequestOutputFormatAlaw8000     ComposeDetailedStreamMusicRequestOutputFormat = "alaw_8000"
	ComposeDetailedStreamMusicRequestOutputFormatOpus4800032  ComposeDetailedStreamMusicRequestOutputFormat = "opus_48000_32"
	ComposeDetailedStreamMusicRequestOutputFormatOpus4800064  ComposeDetailedStreamMusicRequestOutputFormat = "opus_48000_64"
	ComposeDetailedStreamMusicRequestOutputFormatOpus4800096  ComposeDetailedStreamMusicRequestOutputFormat = "opus_48000_96"
	ComposeDetailedStreamMusicRequestOutputFormatOpus48000128 ComposeDetailedStreamMusicRequestOutputFormat = "opus_48000_128"
	ComposeDetailedStreamMusicRequestOutputFormatOpus48000192 ComposeDetailedStreamMusicRequestOutputFormat = "opus_48000_192"
)

func NewComposeDetailedStreamMusicRequestOutputFormatFromString(s string) (ComposeDetailedStreamMusicRequestOutputFormat, error) {
	switch s {
	case "auto":
		return ComposeDetailedStreamMusicRequestOutputFormatAuto, nil
	case "mp3_48000_128":
		return ComposeDetailedStreamMusicRequestOutputFormatMp348000128, nil
	case "mp3_48000_192":
		return ComposeDetailedStreamMusicRequestOutputFormatMp348000192, nil
	case "mp3_48000_240":
		return ComposeDetailedStreamMusicRequestOutputFormatMp348000240, nil
	case "mp3_48000_320":
		return ComposeDetailedStreamMusicRequestOutputFormatMp348000320, nil
	case "mp3_22050_32":
		return ComposeDetailedStreamMusicRequestOutputFormatMp32205032, nil
	case "mp3_24000_48":
		return ComposeDetailedStreamMusicRequestOutputFormatMp32400048, nil
	case "mp3_44100_32":
		return ComposeDetailedStreamMusicRequestOutputFormatMp34410032, nil
	case "mp3_44100_64":
		return ComposeDetailedStreamMusicRequestOutputFormatMp34410064, nil
	case "mp3_44100_96":
		return ComposeDetailedStreamMusicRequestOutputFormatMp34410096, nil
	case "mp3_44100_128":
		return ComposeDetailedStreamMusicRequestOutputFormatMp344100128, nil
	case "mp3_44100_192":
		return ComposeDetailedStreamMusicRequestOutputFormatMp344100192, nil
	case "pcm_8000":
		return ComposeDetailedStreamMusicRequestOutputFormatPcm8000, nil
	case "pcm_16000":
		return ComposeDetailedStreamMusicRequestOutputFormatPcm16000, nil
	case "pcm_22050":
		return ComposeDetailedStreamMusicRequestOutputFormatPcm22050, nil
	case "pcm_24000":
		return ComposeDetailedStreamMusicRequestOutputFormatPcm24000, nil
	case "pcm_32000":
		return ComposeDetailedStreamMusicRequestOutputFormatPcm32000, nil
	case "pcm_44100":
		return ComposeDetailedStreamMusicRequestOutputFormatPcm44100, nil
	case "pcm_48000":
		return ComposeDetailedStreamMusicRequestOutputFormatPcm48000, nil
	case "ulaw_8000":
		return ComposeDetailedStreamMusicRequestOutputFormatUlaw8000, nil
	case "alaw_8000":
		return ComposeDetailedStreamMusicRequestOutputFormatAlaw8000, nil
	case "opus_48000_32":
		return ComposeDetailedStreamMusicRequestOutputFormatOpus4800032, nil
	case "opus_48000_64":
		return ComposeDetailedStreamMusicRequestOutputFormatOpus4800064, nil
	case "opus_48000_96":
		return ComposeDetailedStreamMusicRequestOutputFormatOpus4800096, nil
	case "opus_48000_128":
		return ComposeDetailedStreamMusicRequestOutputFormatOpus48000128, nil
	case "opus_48000_192":
		return ComposeDetailedStreamMusicRequestOutputFormatOpus48000192, nil
	}
	var t ComposeDetailedStreamMusicRequestOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ComposeDetailedStreamMusicRequestOutputFormat) Ptr() *ComposeDetailedStreamMusicRequestOutputFormat {
	return &e
}
