package types

import (
	"fmt"
)

// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM and WAV formats with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
type ConvertWithTimestampsTextToSpeechRequestOutputFormat string

const (
	ConvertWithTimestampsTextToSpeechRequestOutputFormatAlaw8000     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "alaw_8000"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatMp32205032   ConvertWithTimestampsTextToSpeechRequestOutputFormat = "mp3_22050_32"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatMp32400048   ConvertWithTimestampsTextToSpeechRequestOutputFormat = "mp3_24000_48"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatMp344100128  ConvertWithTimestampsTextToSpeechRequestOutputFormat = "mp3_44100_128"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatMp344100192  ConvertWithTimestampsTextToSpeechRequestOutputFormat = "mp3_44100_192"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatMp34410032   ConvertWithTimestampsTextToSpeechRequestOutputFormat = "mp3_44100_32"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatMp34410064   ConvertWithTimestampsTextToSpeechRequestOutputFormat = "mp3_44100_64"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatMp34410096   ConvertWithTimestampsTextToSpeechRequestOutputFormat = "mp3_44100_96"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus48000128 ConvertWithTimestampsTextToSpeechRequestOutputFormat = "opus_48000_128"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus48000192 ConvertWithTimestampsTextToSpeechRequestOutputFormat = "opus_48000_192"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus4800032  ConvertWithTimestampsTextToSpeechRequestOutputFormat = "opus_48000_32"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus4800064  ConvertWithTimestampsTextToSpeechRequestOutputFormat = "opus_48000_64"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus4800096  ConvertWithTimestampsTextToSpeechRequestOutputFormat = "opus_48000_96"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm16000     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "pcm_16000"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm22050     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "pcm_22050"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm24000     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "pcm_24000"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm32000     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "pcm_32000"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm44100     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "pcm_44100"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm48000     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "pcm_48000"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm8000      ConvertWithTimestampsTextToSpeechRequestOutputFormat = "pcm_8000"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatUlaw8000     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "ulaw_8000"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatWav16000     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "wav_16000"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatWav22050     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "wav_22050"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatWav24000     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "wav_24000"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatWav32000     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "wav_32000"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatWav44100     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "wav_44100"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatWav48000     ConvertWithTimestampsTextToSpeechRequestOutputFormat = "wav_48000"
	ConvertWithTimestampsTextToSpeechRequestOutputFormatWav8000      ConvertWithTimestampsTextToSpeechRequestOutputFormat = "wav_8000"
)

func NewConvertWithTimestampsTextToSpeechRequestOutputFormatFromString(s string) (ConvertWithTimestampsTextToSpeechRequestOutputFormat, error) {
	switch s {
	case "alaw_8000":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatAlaw8000, nil
	case "mp3_22050_32":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatMp32205032, nil
	case "mp3_24000_48":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatMp32400048, nil
	case "mp3_44100_128":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatMp344100128, nil
	case "mp3_44100_192":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatMp344100192, nil
	case "mp3_44100_32":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatMp34410032, nil
	case "mp3_44100_64":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatMp34410064, nil
	case "mp3_44100_96":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatMp34410096, nil
	case "opus_48000_128":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus48000128, nil
	case "opus_48000_192":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus48000192, nil
	case "opus_48000_32":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus4800032, nil
	case "opus_48000_64":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus4800064, nil
	case "opus_48000_96":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus4800096, nil
	case "pcm_16000":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm16000, nil
	case "pcm_22050":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm22050, nil
	case "pcm_24000":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm24000, nil
	case "pcm_32000":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm32000, nil
	case "pcm_44100":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm44100, nil
	case "pcm_48000":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm48000, nil
	case "pcm_8000":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm8000, nil
	case "ulaw_8000":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatUlaw8000, nil
	case "wav_16000":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatWav16000, nil
	case "wav_22050":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatWav22050, nil
	case "wav_24000":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatWav24000, nil
	case "wav_32000":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatWav32000, nil
	case "wav_44100":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatWav44100, nil
	case "wav_48000":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatWav48000, nil
	case "wav_8000":
		return ConvertWithTimestampsTextToSpeechRequestOutputFormatWav8000, nil
	}
	var t ConvertWithTimestampsTextToSpeechRequestOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConvertWithTimestampsTextToSpeechRequestOutputFormat) Ptr() *ConvertWithTimestampsTextToSpeechRequestOutputFormat {
	return &e
}
