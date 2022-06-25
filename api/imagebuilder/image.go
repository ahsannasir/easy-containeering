package imagebuilder

type Image struct {
	Name       string   `json:"name,omitempty"`
	Tags       []string `json:"tags,omitempty"`
	Dockerfile string   `json:"dockerfile,omitempty"`
}
