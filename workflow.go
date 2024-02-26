package comfywgo

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
)

type ID int

func (i ID) String() string {
	return strconv.Itoa(int(i))
}

// Workflow represent a workflow
type Workflow struct {
	mu    sync.Mutex
	id    ID
	nodes map[string]*Node
}

type Meta struct {
	Title string `json:"title"`
}

// Node represent a node in workflow,
// NOTE: modifying a node is not concurrent safe
type Node struct {
	errors    []error
	ID        ID             `json:"-"`
	Inputs    map[string]any `json:"inputs"`
	ClassType string         `json:"class_type"`
	Meta      Meta           `json:"_meta"`
	Outputs   map[string]int `json:"-"`
}

type Params struct {
	Required []string
	Inputs   map[string]any
}

func NewParams(required ...string) *Params {
	return &Params{
		Required: required,
		Inputs:   make(map[string]any),
	}
}

// New creates new workflow
func New() *Workflow {
	return &Workflow{
		nodes: make(map[string]*Node),
	}
}

// Title sets the title
func (n *Node) Title(title string) *Node {
	n.Meta.Title = title
	return n
}

// InputMap sets inputs to the node
func (n *Node) InputMap(inputs map[string]any) *Node {
	for k, v := range inputs {
		n.Inputs[k] = v
	}
	return n
}

// WithInput adds new input to the node
func (n *Node) WithInput(key string, value any) *Node {
	n.Inputs[key] = value
	return n
}

// OutputMap sets the outputs to the node
func (n *Node) OutputMap(outputs map[string]int) *Node {
	for k, v := range outputs {
		n.Outputs[k] = v
	}
	return n
}

// WithOutput adds new output to the node
// you can add multiple output for single index.
// e.g. String: 0, Conditioning: 0, Positive: 0, all outputs will be mapped to the given index
func (n *Node) WithOutput(key string, index int) *Node {
	n.Outputs[key] = index
	return n
}

// Output takes the output index from the node
func (n *Node) Output(name string) []any {
	if v, ok := n.Outputs[name]; ok {
		return []any{n.ID.String(), v}
	}
	n.Error(fmt.Errorf("%s output not found for %s", name, n.ClassType))
	return []any{n.ID.String(), 0} // take first output
}

func (n *Node) Error(err error) {
	n.errors = append(n.errors, err)
}

// Validate matches all the outputs
// returns false and first Error in case of missing output
// returns true and nil, if no Error found
func (w *Workflow) Validate() (bool, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	for _, node := range w.nodes {
		if node.errors != nil {
			return false, node.errors[0]
		}
	}
	return true, nil
}

// Json returns a valid json that can be used in ComfyUI API
func (w *Workflow) Json() ([]byte, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return json.Marshal(w.nodes)
}

// Nodes returns all the nodes
func (w *Workflow) Nodes() map[string]*Node {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.nodes
}
