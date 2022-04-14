package types

import (
	"encoding/xml"
	"time"
)

type Osm struct {
	Version   string      `xml:"version"`
	ChageSets []ChangeSet `xml:"changeset"`
}

type ChangeSet struct {
	Id         int     `xml:"id,attr"`
	CreatedAt  string  `xml:"created_at,attr"`
	NumChanges int     `xml:"num_changes,attr"`
	MinLat     float64 `xml:"min_lat,attr"`
	MaxLat     float64 `xml:"max_lat,attr"`
	MinLong    float64 `xml:"min_lon,attr"`
	MaxLong    float64 `xml:"max_lon,attr"`

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
	Version string   `xml:"version,attr"`
	Create  Action   `xml:"create"`
	Modify  Action   `xml:"modify"`
	Delete  Action   `xml:"delete"`
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
	Id        int       `xml:"id,attr"`
	Version   int       `xml:"version,attr"`
	Timestamp time.Time `xml:"timestamp,attr"`
	Lat       float64   `xml:"lat,attr"`
	Lon       float64   `xml:"lon,attr"`
	Tags      []Tag     `xml:"tag"`
}

type Way struct {
	Id        int       `xml:"id,attr"`
	Version   int       `xml:"version,attr"`
	Timestamp time.Time `xml:"timestamp,attr"`
	NodeRefs  []NodeRef `xml:"nd"`
	Tags      []Tag     `xml:"tag"`
}

type NodeRef struct {
	Ref int `xml:"ref,attr"`
}

type Relation struct {
	Id        int       `xml:"id,attr"`
	Version   int       `xml:"version,attr"`
	Timestamp time.Time `xml:"timestamp,attr"`
	Member    []Member  `xml:"member"`
	Tags      []Tag     `xml:"tag"`
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
	Lat float64
	Lng float64
}

type SearchPayload struct {
	Modify []SearchPoint
	Create []SearchPoint
	Delete []SearchPoint
}

type FilterConfig struct {
	Subjects []Subject
}

type Subject struct {
	Name           string
	NodeFilters    []NodeFilter
	WayFilters     []WayFilter
	Compress       bool
	ReduceToPoints bool
	Format         string
}

const (
	FormatJSON = "JSON"
	FormatXML  = "XML"
)

type NodeFilter struct {
	TagKeys []string
}

type WayFilter struct {
	TagKeys []string
}

const (
	MODIFY_EVENT = "MODIFY"
	DELETE_EVENT = "DELETE"
	CREATE_EVENT = "CREATE"
)

func (payload *SearchPayload) Size() int {
	return len(payload.Delete) + len(payload.Modify) + len(payload.Create)
}
