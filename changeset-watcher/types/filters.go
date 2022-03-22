package types

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

func NewNodeFilter(tagKeys ...string) NodeFilter {
	return NodeFilter{
		TagKeys: tagKeys,
	}
}

func NewWayFilter(tagKeys ...string) WayFilter {
	return WayFilter{
		TagKeys: tagKeys,
	}

}
