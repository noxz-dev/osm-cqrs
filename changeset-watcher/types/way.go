package types

import "errors"

func (way Way) HasTags(tags ...string) bool {
	for _, tag := range tags {
		_, err := way.GetTag(tag)
		if err != nil {
			return false
		}
	}
	return true

}

func (way Way) GetTag(tagString string) (value string, err error) {
	for _, tag := range way.Tags {
		if tagString == tag.K {
			return tag.V, nil
		}
	}
	return "", errors.New("Tag " + tagString + " not found")

}
