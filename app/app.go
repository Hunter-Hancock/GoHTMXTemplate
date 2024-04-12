package app

import (
	"github.com/Hunter-Hancock/dbproject/api"
	"github.com/Hunter-Hancock/dbproject/db"
	mw "github.com/Hunter-Hancock/dbproject/middleware"
)

type Application struct {
	TestStore   db.TestStore
	TestHandler *api.TestHandler
	Middleware  mw.MiddleWare
}

func NewApplication() (*Application, error) {
	sql, err := db.Open()
	if err != nil {
		return nil, err
	}

	testStore := db.NewSQLTestStore(sql)
	testHandler := api.NewTestHandler(testStore)

	app := &Application{
		TestStore:   testStore,
		TestHandler: testHandler,
		Middleware:  mw.MiddleWare{},
	}

	return app, nil
}
