package comfyui

func (w *Workflow) EmptyLatentImage(width, height, batchSize int) *Node {
	return w.Node("EmptyLatentImage").
		Title("Empty Latent Image").
		WithInput("width", width).
		WithInput("height", height).
		WithInput("batch_size", batchSize).
		WithOutput(Latent, 0)
}
