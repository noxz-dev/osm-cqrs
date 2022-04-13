package types

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
