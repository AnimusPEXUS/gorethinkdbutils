package rethinkdbutils

type IndexStatus struct {
	Index    *string  `gorethink:"index,omitempty"`
	Ready    *bool    `gorethink:"ready,omitempty"`
	Progress *float64 `gorethink:"progress,omitempty"`
	Function []byte   `gorethink:"function,omitempty"`
	Multi    *bool    `gorethink:"multi,omitempty"`
	Geo      *bool    `gorethink:"geo,omitempty"`
	Outdated *bool    `gorethink:"outdated,omitempty"`
}
