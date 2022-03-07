package types

import (
	"strings"
	"time"
)

func (way Way) HasTags(tags ...string) bool {
	for _, tag := range tags {
		_, exists := way.GetTag(tag)
		if !exists {
			return false
		}
	}
	return true

}

func (way Way) GetTag(tagString string) (value string, exists bool) {
	for _, tag := range way.Tags {
		if tagString == tag.K {
			return tag.V, true
		}
	}
	return "", false

}

func (way *Way) getCreationTime() (creationTime time.Time, err error) {
	creationTime, err = time.Parse(time.RFC3339, way.Timestamp)
	return
}

func (way *Way) GetAddressString() string {
	sb := strings.Builder{}
	street, useSeparator := way.GetTag("addr:street")
	sb.WriteString(street)

	if houseNumber, exists := way.GetTag("addr:housenumber"); exists {
		sb.WriteString(" " + houseNumber)
		useSeparator = true
	}

	if postCode, exists := way.GetTag("addr:postcode"); exists {
		if useSeparator {
			sb.WriteString(", ")
		}
		sb.WriteString(postCode)
		useSeparator = true
	}

	if city, exists := way.GetTag("addr:city"); exists {
		if useSeparator {
			sb.WriteString(", ")
		}
		sb.WriteString(city)
		useSeparator = true
	}

	if country, exists := way.GetTag("addr:country"); exists {
		if useSeparator {
			sb.WriteString(", ")
		}
		sb.WriteString(country)
		useSeparator = true
	}

	return sb.String()
}
