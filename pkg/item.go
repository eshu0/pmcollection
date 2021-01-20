package pmcollection

type Item struct {
	Name string `json:"name,omitempty"`

	Item []*Item `json:"item,omitempty"`

	ProtocolProfileBehavior ProtocolProfileBehavior `json:"protocolProfileBehavior,omitempty"`

	Event []*Event `json:"event,omitempty"`

	Request *Request `json:"request,omitempty"`

	Response []*Response `json:"response,omitempty"`
}

type Response struct {
}

type Event struct {
	Listen string `json:"listen,omitempty"`

	Script *Script `json:"script,omitempty"`
}

type Script struct {
	Exec []string `json:"exec,omitempty"`

	Type string `json:"type,omitempty"`
}

type Request struct {
	Method string `json:"method,omitempty"`

	Body *Body `json:"body,omitempty"`

	Header []*Header `json:"header,omitempty"`

	Url *Url `json:"url,omitempty"`

	Description string `json:"description,omitempty"`
}

type Header struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	Type string `json:"type,omitempty"`

	Disabled bool `json:"disabled,omitempty"`
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
