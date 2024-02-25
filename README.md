# ComfyUI Workflow Builder
Go module to build ComfyUI workflow programmatically 

### Installation
```go
go get github.com/mbrostami/comfywgo
```
### API Workflow

Basic workflow looks like this:  
```go
package main

import (
	"fmt"
	"github.com/mbrostami/comfywgo"
)
func main() {
	w := comfywgo.New()
	
	// checkpoint loader (3 outputs, model/clip/vae)
	model := w.CheckpointLoaderSimple("checkpoint_name")
	
	// prompt (1 conditioning output, takes model's clip output as input)
	pos := w.CLIPTextEncode(model, "positive prompt")
	
	// prompt (1 conditioning output, takes model's clip output as input)
	neg := w.CLIPTextEncode(model, "negative prompt")
	
	emptyLatent := w.EmptyLatentImage(1024, 1024, 1)

	// KSampler (1 latent output)
	sampler := w.KSampler(
		comfywgo.NewKSamplerParams().
			Model(model).
			Positive(pos).
			Negative(neg).
			Latent(emptyLatent),
	)
	
	// VAEDecode (1 output, takes ksampler latent, and model vae)
	w.VAEDecode(sampler, model)
	
	// prints the whole api workflow as json
	fmt.Print(w.Json())
}

```

#### Create Custom Nodes
```go
package main

import (
	"fmt"

	"github.com/mbrostami/comfywgo"
)
func main() {
	w := comfywgo.New()
	model := w.CheckpointLoaderSimple("sd15.safetensors")
	node1 := w.Node("ExampleNode").
		WithInput("clip", model.Output(comfywgo.Clip)). // takes the clip output of the loader node
		WithOutput("OutputName", 0) // sets the first output as OutputName

	w.Node("SecondNode").
		WithInput("input", node1.Output("OutputName")) // takes the output (OutputName)

	fmt.Print(w.Json())
}
```
Note: Output names are for internal use only and will not be used in ComfyUI   
