package types

import "time"

func (relation *Relation) getCreationTime() (creationTime time.Time, err error) {
	creationTime, err = time.Parse(time.RFC3339, relation.Timestamp)
	return
}
