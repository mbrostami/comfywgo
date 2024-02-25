package comfywgo

func (w *Workflow) CheckpointLoaderSimple(modelName string) *Node {
	return w.Node("CheckpointLoaderSimple").
		Title("Load Checkpoint").
		WithInput("ckpt_name", modelName).
		WithOutput(Model, 0).
		WithOutput(Clip, 1).
		WithOutput(VAE, 2)
}

func (w *Workflow) VAELoader(modelName string) *Node {
	return w.Node("VAELoader").
		Title("Load VAE").
		WithInput("vae_name", modelName).
		WithOutput(VAE, 0)
}
