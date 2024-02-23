package comfyui

func (w *Workflow) CLIPTextEncode(clip *Node, text string) *Node {
	return w.Node("CLIPTextEncode").
		Title("CLIP Text Encode (Prompt)").
		WithInput("text", text).
		WithInput("clip", clip.Output(Clip)).
		WithOutput(Conditioning, 0).
		WithOutput(String, 0)
}
