package main

// outputs
const (
	Image        string = "image"
	Latent       string = "latent"
	String       string = "string"
	Positive     string = "positive"
	Negative     string = "negative"
	Conditioning string = "conditioning"
	Clip         string = "clip"
	Model        string = "model"
	VAE          string = "vae"
	ControlNet   string = "control_net"
)

// Node creates an new node
func (w *Workflow) Node(classType string) *Node {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.id++
	w.nodes[w.id.String()] = &Node{
		ID:        w.id,
		Inputs:    map[string]any{},
		ClassType: classType,
		Meta: Meta{
			Title: classType,
		},
		Outputs: map[string]int{},
	}
	return w.nodes[w.id.String()]
}
