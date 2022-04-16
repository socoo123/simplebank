package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
