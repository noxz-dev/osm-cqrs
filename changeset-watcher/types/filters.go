package types

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
