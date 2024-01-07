package dashboards

import "steve.care/network/domain/hash"

type dashboards struct {
	hash hash.Hash
	list []Dashboard
}

func createDashboards(
	hash hash.Hash,
	list []Dashboard,
) Dashboards {
	out := dashboards{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *dashboards) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *dashboards) List() []Dashboard {
	return obj.list
}
