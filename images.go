package comfywgo

func (w *Workflow) SaveImage(image *Node) *Node {
	return w.Node("SaveImage").
		WithInput("filename_prefix", "ComfyUI").
		WithInput("images", image.Output(Image))
}

func (w *Workflow) PreviewImage(image *Node) *Node {
	return w.Node("PreviewImage").WithInput("images", image.Output(Image))
}

func (w *Workflow) LoadImage(name string) *Node {
	return w.Node("LoadImage").
		Title("Load Image").
		WithInput("image", name).
		WithInput("upload", "image").
		WithOutput(Image, 0)
}
