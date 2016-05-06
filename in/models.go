package in

import "github.com/mrosecrance/tracker-resource"

type InRequest struct {
	Source  resource.Source  `json:"source"`
	Version resource.Version `json:"version"`
}

type InResponse struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
}
