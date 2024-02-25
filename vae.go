package main

func (w *Workflow) VAEDecode(latent, model *Node) *Node {
	return w.Node("VAEDecode").
		WithInput("samples", latent.Output(Latent)).
		WithInput("vae", model.Output(VAE)).
		WithOutput(Image, 0)
}

func (w *Workflow) VAEEncode(image, model *Node) *Node {
	return w.Node("VAEEncode").
		WithInput("pixels", image.Output(Image)).
		WithInput("vae", model.Output(VAE)).
		WithOutput(Latent, 0)
}
