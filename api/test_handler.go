package api

import (
	"fmt"
	"net/http"

	"github.com/Hunter-Hancock/dbproject/db"
	"github.com/Hunter-Hancock/dbproject/ui"
)

type TestHandler struct {
	TestStore db.TestStore
}

func NewTestHandler(db db.TestStore) *TestHandler {
	return &TestHandler{TestStore: db}
}

func (t TestHandler) Test(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userID := ctx.Value("userID").(int)

	msg := fmt.Sprintf("User ID: %d", userID)

	ui.Test(msg).Render(r.Context(), w)
}
