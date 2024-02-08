package sqllites

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	"steve.care/network/domain/dashboards/widgets/viewports"
	"steve.care/network/domain/orms"
)

type instanceExec struct {
	name     string
	insatnce orms.Instance
}

func TestOrm_Success(t *testing.T) {
	dbDir := "./test_files"
	dbName := "testdb"
	basePath := filepath.Join(dbDir, dbName)
	defer func() {
		os.Remove(basePath)
	}()

	pDB, err := sql.Open("sqlite3", basePath)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	skeleton, err := NewSkeletonFactory().Create()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	//repository := NewOrmRepository(skeleton, pDB)

	pTx, err := pDB.Begin()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	service := NewOrmService(skeleton, pDB, pTx)

	// build viewport resource:
	//	pRandomHash, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes"))
	instances := []orms.Instance{
		viewports.NewViewportForTests(uint(33), uint(45)),
		/*tokens.NewTokenWithDashboardForTests(
			token_dashboards.NewDashboardWithViewportForTests(
				viewports.NewViewportForTests(uint(33), uint(45)),
			),
		),
		tokens.NewTokenWithDashboardForTests(
			token_dashboards.NewDashboardWithWidgetForTests(
				widgets.NewWidgetForTests(
					"this is a title",
					*pRandomHash,
					[]byte("this is an input"),
				),
			),
		),
		tokens.NewTokenWithDashboardForTests(
			token_dashboards.NewDashboardWithWidgetForTests(
				widgets.NewWidgetWithViewportForTests(
					"this is a title",
					*pRandomHash,
					[]byte("this is an input"),
					viewports.NewViewportForTests(uint(33), uint(45)),
				),
			),
		),*/
	}

	for idx, oneInstance := range instances {
		// init our service:
		err = service.Init()
		if err != nil {
			t.Errorf("index: %d, the error was expected to be nil, error returned: %s", idx, err.Error())
			return
		}

		// insert instance:
		err = service.Insert(oneInstance, []string{
			"dashboard",
			"widget",
			"viewport",
		})

		if err != nil {
			t.Errorf("index: %d, the error was expected to be nil, error returned: %s", idx, err.Error())
			return
		}

		/*insHash := oneInstance.Hash()
		retInstance, err := repository.RetrieveByHash(insHash)
		if err != nil {
			t.Errorf("index: %d, the error was expected to be nil, error returned: %s", idx, err.Error())
			return
		}

		if !bytes.Equal(retInstance.Hash().Bytes(), insHash.Bytes()) {
			t.Errorf("index: %d, the returned insatnce is invalid", idx)
			return
		}*/
	}
}
