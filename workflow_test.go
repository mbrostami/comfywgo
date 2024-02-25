package main

import (
	"reflect"
	"testing"
)

func TestWorkflowJsonHappy(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "test simple workflow",
			want:    "{\"1\":{\"Inputs\":{\"ckpt_name\":\"checkpoint_name\"},\"class_type\":\"CheckpointLoaderSimple\",\"_meta\":{\"title\":\"Load Checkpoint\"}},\"2\":{\"Inputs\":{\"clip\":[\"1\",1],\"text\":\"positive prompt\"},\"class_type\":\"CLIPTextEncode\",\"_meta\":{\"title\":\"CLIP Text Encode (Prompt)\"}},\"3\":{\"Inputs\":{\"clip\":[\"1\",1],\"text\":\"negative prompt\"},\"class_type\":\"CLIPTextEncode\",\"_meta\":{\"title\":\"CLIP Text Encode (Prompt)\"}},\"4\":{\"Inputs\":{\"batch_size\":1,\"height\":1024,\"width\":1024},\"class_type\":\"EmptyLatentImage\",\"_meta\":{\"title\":\"Empty Latent Image\"}},\"5\":{\"Inputs\":{\"cfg\":8,\"denoise\":1,\"latent_image\":[\"4\",0],\"model\":[\"1\",0],\"negative\":[\"3\",0],\"positive\":[\"2\",0],\"sampler_name\":\"euler\",\"scheduler\":\"kerras\",\"steps\":30},\"class_type\":\"KSampler\",\"_meta\":{\"title\":\"KSampler\"}},\"6\":{\"Inputs\":{\"samples\":[\"5\",0],\"vae\":[\"1\",2]},\"class_type\":\"VAEDecode\",\"_meta\":{\"title\":\"VAEDecode\"}}}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New()
			model := w.CheckpointLoaderSimple("checkpoint_name")
			pos := w.CLIPTextEncode(model, "positive prompt")
			neg := w.CLIPTextEncode(model, "negative prompt")
			emptyLatent := w.EmptyLatentImage(1024, 1024, 1)
			sampler := w.KSampler(
				NewKSamplerParams().
					Model(model).
					Positive(pos).
					Negative(neg).
					Latent(emptyLatent),
			)
			w.VAEDecode(sampler, model)
			if _, err := w.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() got = %v, want %v", err, tt.wantErr)
			}

			got, _ := w.Json()
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("Json() got = %v, want %v", string(got), tt.want)
			}
		})
	}
}
