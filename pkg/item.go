package pmcollection

type Item struct {
	Name string `json:"name,omitempty"`

	Item []*Item `json:"item,omitempty"`

	ProtocolProfileBehavior ProtocolProfileBehavior `json:"protocolProfileBehavior,omitempty"`

	Request *Request `json:"request,omitempty"`

	Response []*Response `json:"response,omitempty"`
}

type Response struct {
}

type Request struct {
	Method string `json:"method,omitempty"`

	Body *Body `json:"body,omitempty"`

	Url *Url `json:"url,omitempty"`
}

type Url struct {
	Raw string `json:"raw,omitempty"`

	Protocol string `json:"protocol,omitempty"`

	Port string `json:"port,omitempty"`

	Host []string `json:"host,omitempty"`

	Path []string `json:"path,omitempty"`
}

type Body struct {
	Mode string `json:"mode,omitempty"`

	Raw string `json:"raw,omitempty"`

	Options *Options `json:"options,omitempty"`
}

type Options struct {
	Raw *RawOption `json:"raw,omitempty"`
}

type RawOption struct {
	Language string `json:"language,omitempty"`
}

type ProtocolProfileBehavior struct {
	DisableBodyPruning bool `json:"disableBodyPruning,omitempty"`
}
