package pmcollection

type Collection struct {
	Info Info `json:"info"`

	Item []Item `json:"item"`
}

type Info struct {
	Name string `json:"name"`

	Postman_Id string `json:"_postman_id"`

	//"https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	Schema string `json:"schema"`
}

type Item struct {
	Name string `json:"name"`

	Item []Item `json:"item"`

	ProtocolProfileBehavior ProtocolProfileBehavior `json:"protocolProfileBehavior"`
}

type FolderItem struct {
	Item
}

type RequestItem struct {
	Item

	Request Request `json:"request"`

	Response Response `json:"response"`
}

type Response struct {
}

type Request struct {
	Method string `json:"method"`

	Body Body `json:"body"`

	Url Url `json:"url"`
}

type Url struct {
	Raw string `json:"raw"`

	Protocol string `json:"protocol"`

	Port string `json:"port"`

	Host []string `json:"host"`

	Path []string `json:"path"`
}

type Body struct {
	Mode string `json:"mode"`

	Raw string `json:"raw"`

	Options Options `json:"options"`
}

type Options struct {
	Raw RawOption `json:"raw"`
}

type RawOption struct {
	Language string `json:"language"`
}

type ProtocolProfileBehavior struct {
	DisableBodyPruning bool `json:"disableBodyPruning"`
}
