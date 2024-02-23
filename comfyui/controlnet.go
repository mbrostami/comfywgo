package comfyui

func (w *Workflow) ControlNetLoader(modelName string) *Node {
	return w.Node("ControlNetLoader").
		Title("Load ControlNet Model").
		WithInput("control_net_name", modelName).
		WithOutput(ControlNet, 0)
}

func (w *Workflow) ControlNetApply(controlNet, conditioning, image *Node, strength float32) *Node {
	return w.Node("ControlNetApply").
		Title("Apply ControlNet").
		WithInput("strength", strength).
		WithInput("control_net", controlNet.Output(ControlNet)).
		WithInput("conditioning", conditioning.Output(Conditioning)).
		WithInput("image", image.Output(Image)).
		WithOutput(Conditioning, 0)
}
