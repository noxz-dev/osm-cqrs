package types

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
	Version string   `xml:"version"`
	Modify  []Action `xml:"modify"`
	Create  []Action `xml:"create"`
	Delete  []Action `xml:"delete"`
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

const (
	MODIFY_EVENT = "MODIFY"
	DELETE_EVENT = "DELETE"
	CREATE_EVENT = "CREATE"
)

func (action Action) ContainsNodeByRef(ref NodeRef) bool {
	for _, node := range action.Nodes {
		if node.Id == ref.Ref {
			return true
		}
	}
	return false
}
