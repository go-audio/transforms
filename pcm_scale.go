package transforms

import "github.com/go-audio/audio"

// PCMScale converts a buffer with audio content from -1 to 1 into
// the PCM scale based on the buffer's bitdepth.
// Note that while the PCM data is scaled, the PCM format is not changed.
func PCMScale(buf *audio.FloatBuffer) error {
	if buf == nil || buf.Format == nil {
		return audio.ErrInvalidBuffer
	}
	factor := float64(audio.IntMaxSignedValue(buf.Format.BitDepth))
	for i := 0; i < len(buf.Data); i++ {
		buf.Data[i] *= factor
	}

	return nil
}
