# ComfyUI Workflow Builder
Go module to build ComfyUI workflow programmatically 

### API Workflow

Basic workflow looks like this:  
```go
import github.com/mbrostami/comfywgo

w := comfywgo.New()
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
fmt.Print(w.Json())
```

### Create New Node
```go
w := comfywgo.New()
node := w.Node("CLIPTextEncode").
    WithInput("text", text).
    WithInput("clip", clip.Output(Clip)).
    WithOutput("OutputName", 0)
```
Note: You need to specify inputs and outputs of a node during creation    
OutputName is for internal use only and it will not be used in ComfyUI   
