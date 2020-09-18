package factory

import (
	"testing"
)

func TestDatabaseFactory(t *testing.T) {
	env1 := "production"
	env2 := "development"

	db1 := factory.DatabaseFactory(env1)
	db2 := factory.DatabaseFactory(env2)
}
