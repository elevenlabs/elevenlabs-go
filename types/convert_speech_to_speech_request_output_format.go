package types

import (
	"fmt"
)

// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM and WAV formats with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
type ConvertSpeechToSpeechRequestOutputFormat string

const (
	ConvertSpeechToSpeechRequestOutputFormatAlaw8000     ConvertSpeechToSpeechRequestOutputFormat = "alaw_8000"
	ConvertSpeechToSpeechRequestOutputFormatMp32205032   ConvertSpeechToSpeechRequestOutputFormat = "mp3_22050_32"
	ConvertSpeechToSpeechRequestOutputFormatMp32400048   ConvertSpeechToSpeechRequestOutputFormat = "mp3_24000_48"
	ConvertSpeechToSpeechRequestOutputFormatMp344100128  ConvertSpeechToSpeechRequestOutputFormat = "mp3_44100_128"
	ConvertSpeechToSpeechRequestOutputFormatMp344100192  ConvertSpeechToSpeechRequestOutputFormat = "mp3_44100_192"
	ConvertSpeechToSpeechRequestOutputFormatMp34410032   ConvertSpeechToSpeechRequestOutputFormat = "mp3_44100_32"
	ConvertSpeechToSpeechRequestOutputFormatMp34410064   ConvertSpeechToSpeechRequestOutputFormat = "mp3_44100_64"
	ConvertSpeechToSpeechRequestOutputFormatMp34410096   ConvertSpeechToSpeechRequestOutputFormat = "mp3_44100_96"
	ConvertSpeechToSpeechRequestOutputFormatOpus48000128 ConvertSpeechToSpeechRequestOutputFormat = "opus_48000_128"
	ConvertSpeechToSpeechRequestOutputFormatOpus48000192 ConvertSpeechToSpeechRequestOutputFormat = "opus_48000_192"
	ConvertSpeechToSpeechRequestOutputFormatOpus4800032  ConvertSpeechToSpeechRequestOutputFormat = "opus_48000_32"
	ConvertSpeechToSpeechRequestOutputFormatOpus4800064  ConvertSpeechToSpeechRequestOutputFormat = "opus_48000_64"
	ConvertSpeechToSpeechRequestOutputFormatOpus4800096  ConvertSpeechToSpeechRequestOutputFormat = "opus_48000_96"
	ConvertSpeechToSpeechRequestOutputFormatPcm16000     ConvertSpeechToSpeechRequestOutputFormat = "pcm_16000"
	ConvertSpeechToSpeechRequestOutputFormatPcm22050     ConvertSpeechToSpeechRequestOutputFormat = "pcm_22050"
	ConvertSpeechToSpeechRequestOutputFormatPcm24000     ConvertSpeechToSpeechRequestOutputFormat = "pcm_24000"
	ConvertSpeechToSpeechRequestOutputFormatPcm32000     ConvertSpeechToSpeechRequestOutputFormat = "pcm_32000"
	ConvertSpeechToSpeechRequestOutputFormatPcm44100     ConvertSpeechToSpeechRequestOutputFormat = "pcm_44100"
	ConvertSpeechToSpeechRequestOutputFormatPcm48000     ConvertSpeechToSpeechRequestOutputFormat = "pcm_48000"
	ConvertSpeechToSpeechRequestOutputFormatPcm8000      ConvertSpeechToSpeechRequestOutputFormat = "pcm_8000"
	ConvertSpeechToSpeechRequestOutputFormatUlaw8000     ConvertSpeechToSpeechRequestOutputFormat = "ulaw_8000"
	ConvertSpeechToSpeechRequestOutputFormatWav16000     ConvertSpeechToSpeechRequestOutputFormat = "wav_16000"
	ConvertSpeechToSpeechRequestOutputFormatWav22050     ConvertSpeechToSpeechRequestOutputFormat = "wav_22050"
	ConvertSpeechToSpeechRequestOutputFormatWav24000     ConvertSpeechToSpeechRequestOutputFormat = "wav_24000"
	ConvertSpeechToSpeechRequestOutputFormatWav32000     ConvertSpeechToSpeechRequestOutputFormat = "wav_32000"
	ConvertSpeechToSpeechRequestOutputFormatWav44100     ConvertSpeechToSpeechRequestOutputFormat = "wav_44100"
	ConvertSpeechToSpeechRequestOutputFormatWav48000     ConvertSpeechToSpeechRequestOutputFormat = "wav_48000"
	ConvertSpeechToSpeechRequestOutputFormatWav8000      ConvertSpeechToSpeechRequestOutputFormat = "wav_8000"
)

func NewConvertSpeechToSpeechRequestOutputFormatFromString(s string) (ConvertSpeechToSpeechRequestOutputFormat, error) {
	switch s {
	case "alaw_8000":
		return ConvertSpeechToSpeechRequestOutputFormatAlaw8000, nil
	case "mp3_22050_32":
		return ConvertSpeechToSpeechRequestOutputFormatMp32205032, nil
	case "mp3_24000_48":
		return ConvertSpeechToSpeechRequestOutputFormatMp32400048, nil
	case "mp3_44100_128":
		return ConvertSpeechToSpeechRequestOutputFormatMp344100128, nil
	case "mp3_44100_192":
		return ConvertSpeechToSpeechRequestOutputFormatMp344100192, nil
	case "mp3_44100_32":
		return ConvertSpeechToSpeechRequestOutputFormatMp34410032, nil
	case "mp3_44100_64":
		return ConvertSpeechToSpeechRequestOutputFormatMp34410064, nil
	case "mp3_44100_96":
		return ConvertSpeechToSpeechRequestOutputFormatMp34410096, nil
	case "opus_48000_128":
		return ConvertSpeechToSpeechRequestOutputFormatOpus48000128, nil
	case "opus_48000_192":
		return ConvertSpeechToSpeechRequestOutputFormatOpus48000192, nil
	case "opus_48000_32":
		return ConvertSpeechToSpeechRequestOutputFormatOpus4800032, nil
	case "opus_48000_64":
		return ConvertSpeechToSpeechRequestOutputFormatOpus4800064, nil
	case "opus_48000_96":
		return ConvertSpeechToSpeechRequestOutputFormatOpus4800096, nil
	case "pcm_16000":
		return ConvertSpeechToSpeechRequestOutputFormatPcm16000, nil
	case "pcm_22050":
		return ConvertSpeechToSpeechRequestOutputFormatPcm22050, nil
	case "pcm_24000":
		return ConvertSpeechToSpeechRequestOutputFormatPcm24000, nil
	case "pcm_32000":
		return ConvertSpeechToSpeechRequestOutputFormatPcm32000, nil
	case "pcm_44100":
		return ConvertSpeechToSpeechRequestOutputFormatPcm44100, nil
	case "pcm_48000":
		return ConvertSpeechToSpeechRequestOutputFormatPcm48000, nil
	case "pcm_8000":
		return ConvertSpeechToSpeechRequestOutputFormatPcm8000, nil
	case "ulaw_8000":
		return ConvertSpeechToSpeechRequestOutputFormatUlaw8000, nil
	case "wav_16000":
		return ConvertSpeechToSpeechRequestOutputFormatWav16000, nil
	case "wav_22050":
		return ConvertSpeechToSpeechRequestOutputFormatWav22050, nil
	case "wav_24000":
		return ConvertSpeechToSpeechRequestOutputFormatWav24000, nil
	case "wav_32000":
		return ConvertSpeechToSpeechRequestOutputFormatWav32000, nil
	case "wav_44100":
		return ConvertSpeechToSpeechRequestOutputFormatWav44100, nil
	case "wav_48000":
		return ConvertSpeechToSpeechRequestOutputFormatWav48000, nil
	case "wav_8000":
		return ConvertSpeechToSpeechRequestOutputFormatWav8000, nil
	}
	var t ConvertSpeechToSpeechRequestOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConvertSpeechToSpeechRequestOutputFormat) Ptr() *ConvertSpeechToSpeechRequestOutputFormat {
	return &e
}
