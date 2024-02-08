package sqllites

import (
	"database/sql"

	"steve.care/network/domain/hash"
	"steve.care/network/domain/orms"
	"steve.care/network/domain/orms/skeletons"
)

type ormRepository struct {
	hashAdapter hash.Adapter
	builders    map[string]interface{}
	skeleton    skeletons.Skeleton
	dbPtr       *sql.DB
}

func createOrmRepository(
	hashAdapter hash.Adapter,
	builders map[string]interface{},
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
) orms.Repository {
	out := ormRepository{
		hashAdapter: hashAdapter,
		builders:    builders,
		skeleton:    skeleton,
		dbPtr:       dbPtr,
	}

	return &out
}

// AmountByQuery returns the amount of instance by criteria
func (app *ormRepository) AmountByQuery(query hash.Hash) (uint, error) {
	return 0, nil
}

// ListByQuery lists insatnce hashes by criteria
func (app *ormRepository) ListByQuery(query hash.Hash) ([]hash.Hash, error) {
	return nil, nil
}

// RetrieveByQuery retrieves an instance by criteria
func (app *ormRepository) RetrieveByQuery(query hash.Hash) (orms.Instance, error) {
	return nil, nil
}

// RetrieveByHash retrieves an instance by hash
func (app *ormRepository) RetrieveByHash(path []string, hash hash.Hash) (orms.Instance, error) {
	return nil, nil
}
