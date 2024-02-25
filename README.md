# ComfyUI Workflow Builder
Go module to build ComfyUI workflow programmatically 

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
	model := w.CheckpointLoaderSimple("checkpoint_name")
	pos := w.CLIPTextEncode(model, "positive prompt")
	neg := w.CLIPTextEncode(model, "negative prompt")
	emptyLatent := w.EmptyLatentImage(1024, 1024, 1)
	sampler := w.KSampler(
		comfywgo.NewKSamplerParams().
			Model(model).
			Positive(pos).
			Negative(neg).
			Latent(emptyLatent),
	)
	w.VAEDecode(sampler, model)
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
Note: You need to specify inputs and outputs of a node during creation    
OutputName is for internal use only and it will not be used in ComfyUI   
