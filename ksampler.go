package comfyui

import "fmt"

type KSamplerParams struct {
	*Params
}

func NewKSamplerParams() KSamplerParams {
	return KSamplerParams{NewParams("model", "positive", "negative", "latent_image")}
}

func (k KSamplerParams) Steps(steps int) KSamplerParams {
	k.Inputs["steps"] = steps
	return k
}

func (k KSamplerParams) Seed(seed int) KSamplerParams {
	k.Inputs["seed"] = seed
	return k
}

func (k KSamplerParams) CFG(cfg float32) KSamplerParams {
	k.Inputs["cfg"] = cfg
	return k
}

func (k KSamplerParams) Sampler(samplerName string) KSamplerParams {
	k.Inputs["sampler_name"] = samplerName
	return k
}

func (k KSamplerParams) Scheduler(scheduler string) KSamplerParams {
	k.Inputs["scheduler"] = scheduler
	return k
}

func (k KSamplerParams) Denoise(denoise float32) KSamplerParams {
	k.Inputs["denoise"] = denoise
	return k
}

func (k KSamplerParams) Model(model *Node) KSamplerParams {
	k.Inputs["model"] = model.Output(Model)
	return k
}

func (k KSamplerParams) Positive(positive *Node) KSamplerParams {
	k.Inputs["positive"] = positive.Output(Conditioning)
	return k
}

func (k KSamplerParams) Negative(negative *Node) KSamplerParams {
	k.Inputs["negative"] = negative.Output(Conditioning)
	return k
}

func (k KSamplerParams) Latent(latent *Node) KSamplerParams {
	k.Inputs["latent_image"] = latent.Output(Latent)
	return k
}

func (w *Workflow) KSampler(config KSamplerParams) *Node {
	node := w.Node("KSampler").
		WithInput("steps", 30).
		WithInput("cfg", 8.0).
		WithInput("sampler_name", "euler").
		WithInput("scheduler", "kerras").
		WithInput("denoise", 1.0).
		InputMap(config.Inputs).
		WithOutput(Latent, 0)

	for _, k := range config.Required {
		if _, ok := config.Inputs[k]; !ok {
			node.error(fmt.Errorf("%s is Required for %s", k, node.ClassType))
		}
	}
	return node
}
