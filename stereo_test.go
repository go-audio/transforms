package transforms

import (
	"testing"

	"reflect"

	"github.com/go-audio/audio"
)

func TestMonoToStereoF32Validation(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		buf     *audio.Float32Buffer
		wantErr bool
	}{
		{
			name:    "valid blank",
			buf:     &audio.Float32Buffer{Format: &audio.Format{NumChannels: 1}},
			wantErr: false,
		},
		{
			name:    "nil format",
			buf:     &audio.Float32Buffer{},
			wantErr: true,
		},
		{
			name:    "stereo",
			buf:     &audio.Float32Buffer{Format: &audio.Format{NumChannels: 2}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MonoToStereoF32(tt.buf); (err != nil) != tt.wantErr {
				t.Errorf("MonoToStereoF32() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMonoToStereoF32(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name       string
		buf        *audio.Float32Buffer
		stereoData []float32
		wantErr    bool
	}{
		{
			name: "valid blank",
			buf: &audio.Float32Buffer{
				Format: &audio.Format{NumChannels: 1},
				Data:   []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0},
			},
			stereoData: []float32{1.0, 1.0, 2.0, 2.0, 3.0, 3.0, 4.0, 4.0, 5.0, 5.0, 6.0, 6.0, 7.0, 7.0, 8.0, 8.0, 9.0, 9.0, 10.0, 10.0, 11.0, 11.0, 12.0, 12.0},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MonoToStereoF32(tt.buf); (err != nil) != tt.wantErr {
				t.Errorf("MonoToStereoF32() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.buf.Format.NumChannels != 2 {
				t.Errorf("Expected the buffer channel numbers to be 2 got %d", tt.buf.Format.NumChannels)
			}
			if !reflect.DeepEqual(tt.buf.Data, tt.stereoData) {
				t.Errorf("Expected the conversion to match, got %#v expected:\n%#v", tt.buf.Data, tt.stereoData)
			}
		})
	}
}
