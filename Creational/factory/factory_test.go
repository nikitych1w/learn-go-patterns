package factory_test

import (
	"fmt"
	"github.com/nikitych1w/learn-go-patterns/Creational/factory"
	"reflect"
	"testing"
)

func TestDatabaseFactory(t *testing.T) {
	env1 := "production"
	env2 := "development"

	db1 := factory.DatabaseFactory(env1)
	db2 := factory.DatabaseFactory(env2)

	fmt.Println(reflect.TypeOf(db1), reflect.TypeOf(db2))
}
