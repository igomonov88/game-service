package tests

import (
	"fmt"
	"testing"

	"githib.com/igomonov88/game-service/internal/database/dbtest"
	"githib.com/igomonov88/game-service/internal/tests/docker"
)

var c *docker.Container

func TestMain(m *testing.M) {
	var err error
	c, err = dbtest.StartDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbtest.StopDB(c)

	m.Run()
}
