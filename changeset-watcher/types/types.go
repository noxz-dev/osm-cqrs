package types

import "encoding/xml"

type Osm struct {
	Version   string      `xml:"version"`
	ChageSets []ChangeSet `xml:"changeset"`
}

type ChangeSet struct {
	Id         int     `xml:"id,attr"`
	CreatedAt  string  `xml:"created_at,attr"`
	NumChanges int     `xml:"num_changes,attr"`
	MinLat     float32 `xml:"min_lat,attr"`
	MaxLat     float32 `xml:"max_lat,attr"`
	MinLong    float32 `xml:"min_lon,attr"`
	MaxLong    float32 `xml:"max_lon,attr"`

	Tags []Tag `xml:"tag"`
}

type Tag struct {
	K string `xml:"k,attr"`
	V string `xml:"v,attr"`
}

type OsmChange struct {
	Modify []Action `xml:"modify"`
	Create []Action `xml:"create"`
	Delete []Action `xml:"delete"`
}

type OsmChangeNormalizedXML struct {
	XMLName xml.Name `xml:"osmChange"`
	Modify  Action   `json:"modify"`
	Create  Action   `json:"create"`
	Delete  Action   `json:"delete"`
}
type OsmChangeNormalized struct {
	Modify   Action `json:"modify"`
	Create   Action `json:"create"`
	Delete   Action `json:"delete"`
	Reloaded Action `json:"reloaded"`
}

type Action struct {
	Nodes     []Node     `xml:"node"`
	Ways      []Way      `xml:"way"`
	Relations []Relation `xml:"relation"`
}

type Node struct {
	Id        int     `xml:"id,attr"`
	Version   int     `xml:"version,attr"`
	Timestamp string  `xml:"timestamp,attr"`
	Lat       float32 `xml:"lat,attr"`
	Lon       float32 `xml:"lon,attr"`
	Tags      []Tag   `xml:"tag"`
}

type Way struct {
	Id        int       `xml:"id,attr"`
	Version   int       `xml:"version,attr"`
	Timestamp string    `xml:"timestamp,attr"`
	NodeRefs  []NodeRef `xml:"nd"`
	Tags      []Tag     `xml:"tag"`
}

type NodeRef struct {
	Ref int `xml:"ref,attr"`
}

type Relation struct {
	Id        int      `xml:"id,attr"`
	Version   int      `xml:"version,attr"`
	Timestamp string   `xml:"timestamp,attr"`
	Member    []Member `xml:"member"`
	Tags      []Tag    `xml:"tag"`
}

type Member struct {
	Type string `xml:"type,attr"`
	Ref  int    `xml:"ref,attr"`
	Role string `xml:"role,attr"`
}

type OverPassAnswer struct {
	Nodes []Node `json:"elements" xml:"node"`
}

type SearchPoint struct {
	Name     string
	Location Location
	Id       string
	Tags     []Tag
}

type Location struct {
	Lat float32
	Lng float32
}

type SearchPayload struct {
	Modify []SearchPoint
	Create []SearchPoint
	Delete []SearchPoint
}

const (
	MODIFY_EVENT = "MODIFY"
	DELETE_EVENT = "DELETE"
	CREATE_EVENT = "CREATE"
)

func (payload *SearchPayload) Size() int {
	return len(payload.Delete) + len(payload.Modify) + len(payload.Create)
}
