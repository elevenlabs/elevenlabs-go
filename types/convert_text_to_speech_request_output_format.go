package types

import (
	"fmt"
)

// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM and WAV formats with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
type ConvertTextToSpeechRequestOutputFormat string

const (
	ConvertTextToSpeechRequestOutputFormatAlaw8000     ConvertTextToSpeechRequestOutputFormat = "alaw_8000"
	ConvertTextToSpeechRequestOutputFormatMp32205032   ConvertTextToSpeechRequestOutputFormat = "mp3_22050_32"
	ConvertTextToSpeechRequestOutputFormatMp32400048   ConvertTextToSpeechRequestOutputFormat = "mp3_24000_48"
	ConvertTextToSpeechRequestOutputFormatMp344100128  ConvertTextToSpeechRequestOutputFormat = "mp3_44100_128"
	ConvertTextToSpeechRequestOutputFormatMp344100192  ConvertTextToSpeechRequestOutputFormat = "mp3_44100_192"
	ConvertTextToSpeechRequestOutputFormatMp34410032   ConvertTextToSpeechRequestOutputFormat = "mp3_44100_32"
	ConvertTextToSpeechRequestOutputFormatMp34410064   ConvertTextToSpeechRequestOutputFormat = "mp3_44100_64"
	ConvertTextToSpeechRequestOutputFormatMp34410096   ConvertTextToSpeechRequestOutputFormat = "mp3_44100_96"
	ConvertTextToSpeechRequestOutputFormatOpus48000128 ConvertTextToSpeechRequestOutputFormat = "opus_48000_128"
	ConvertTextToSpeechRequestOutputFormatOpus48000192 ConvertTextToSpeechRequestOutputFormat = "opus_48000_192"
	ConvertTextToSpeechRequestOutputFormatOpus4800032  ConvertTextToSpeechRequestOutputFormat = "opus_48000_32"
	ConvertTextToSpeechRequestOutputFormatOpus4800064  ConvertTextToSpeechRequestOutputFormat = "opus_48000_64"
	ConvertTextToSpeechRequestOutputFormatOpus4800096  ConvertTextToSpeechRequestOutputFormat = "opus_48000_96"
	ConvertTextToSpeechRequestOutputFormatPcm16000     ConvertTextToSpeechRequestOutputFormat = "pcm_16000"
	ConvertTextToSpeechRequestOutputFormatPcm22050     ConvertTextToSpeechRequestOutputFormat = "pcm_22050"
	ConvertTextToSpeechRequestOutputFormatPcm24000     ConvertTextToSpeechRequestOutputFormat = "pcm_24000"
	ConvertTextToSpeechRequestOutputFormatPcm32000     ConvertTextToSpeechRequestOutputFormat = "pcm_32000"
	ConvertTextToSpeechRequestOutputFormatPcm44100     ConvertTextToSpeechRequestOutputFormat = "pcm_44100"
	ConvertTextToSpeechRequestOutputFormatPcm48000     ConvertTextToSpeechRequestOutputFormat = "pcm_48000"
	ConvertTextToSpeechRequestOutputFormatPcm8000      ConvertTextToSpeechRequestOutputFormat = "pcm_8000"
	ConvertTextToSpeechRequestOutputFormatUlaw8000     ConvertTextToSpeechRequestOutputFormat = "ulaw_8000"
	ConvertTextToSpeechRequestOutputFormatWav16000     ConvertTextToSpeechRequestOutputFormat = "wav_16000"
	ConvertTextToSpeechRequestOutputFormatWav22050     ConvertTextToSpeechRequestOutputFormat = "wav_22050"
	ConvertTextToSpeechRequestOutputFormatWav24000     ConvertTextToSpeechRequestOutputFormat = "wav_24000"
	ConvertTextToSpeechRequestOutputFormatWav32000     ConvertTextToSpeechRequestOutputFormat = "wav_32000"
	ConvertTextToSpeechRequestOutputFormatWav44100     ConvertTextToSpeechRequestOutputFormat = "wav_44100"
	ConvertTextToSpeechRequestOutputFormatWav48000     ConvertTextToSpeechRequestOutputFormat = "wav_48000"
	ConvertTextToSpeechRequestOutputFormatWav8000      ConvertTextToSpeechRequestOutputFormat = "wav_8000"
)

func NewConvertTextToSpeechRequestOutputFormatFromString(s string) (ConvertTextToSpeechRequestOutputFormat, error) {
	switch s {
	case "alaw_8000":
		return ConvertTextToSpeechRequestOutputFormatAlaw8000, nil
	case "mp3_22050_32":
		return ConvertTextToSpeechRequestOutputFormatMp32205032, nil
	case "mp3_24000_48":
		return ConvertTextToSpeechRequestOutputFormatMp32400048, nil
	case "mp3_44100_128":
		return ConvertTextToSpeechRequestOutputFormatMp344100128, nil
	case "mp3_44100_192":
		return ConvertTextToSpeechRequestOutputFormatMp344100192, nil
	case "mp3_44100_32":
		return ConvertTextToSpeechRequestOutputFormatMp34410032, nil
	case "mp3_44100_64":
		return ConvertTextToSpeechRequestOutputFormatMp34410064, nil
	case "mp3_44100_96":
		return ConvertTextToSpeechRequestOutputFormatMp34410096, nil
	case "opus_48000_128":
		return ConvertTextToSpeechRequestOutputFormatOpus48000128, nil
	case "opus_48000_192":
		return ConvertTextToSpeechRequestOutputFormatOpus48000192, nil
	case "opus_48000_32":
		return ConvertTextToSpeechRequestOutputFormatOpus4800032, nil
	case "opus_48000_64":
		return ConvertTextToSpeechRequestOutputFormatOpus4800064, nil
	case "opus_48000_96":
		return ConvertTextToSpeechRequestOutputFormatOpus4800096, nil
	case "pcm_16000":
		return ConvertTextToSpeechRequestOutputFormatPcm16000, nil
	case "pcm_22050":
		return ConvertTextToSpeechRequestOutputFormatPcm22050, nil
	case "pcm_24000":
		return ConvertTextToSpeechRequestOutputFormatPcm24000, nil
	case "pcm_32000":
		return ConvertTextToSpeechRequestOutputFormatPcm32000, nil
	case "pcm_44100":
		return ConvertTextToSpeechRequestOutputFormatPcm44100, nil
	case "pcm_48000":
		return ConvertTextToSpeechRequestOutputFormatPcm48000, nil
	case "pcm_8000":
		return ConvertTextToSpeechRequestOutputFormatPcm8000, nil
	case "ulaw_8000":
		return ConvertTextToSpeechRequestOutputFormatUlaw8000, nil
	case "wav_16000":
		return ConvertTextToSpeechRequestOutputFormatWav16000, nil
	case "wav_22050":
		return ConvertTextToSpeechRequestOutputFormatWav22050, nil
	case "wav_24000":
		return ConvertTextToSpeechRequestOutputFormatWav24000, nil
	case "wav_32000":
		return ConvertTextToSpeechRequestOutputFormatWav32000, nil
	case "wav_44100":
		return ConvertTextToSpeechRequestOutputFormatWav44100, nil
	case "wav_48000":
		return ConvertTextToSpeechRequestOutputFormatWav48000, nil
	case "wav_8000":
		return ConvertTextToSpeechRequestOutputFormatWav8000, nil
	}
	var t ConvertTextToSpeechRequestOutputFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConvertTextToSpeechRequestOutputFormat) Ptr() *ConvertTextToSpeechRequestOutputFormat {
	return &e
}
