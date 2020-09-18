package abstract_factory_test

import (
	"fmt"
	"github.com/nikitych1w/learn-go-patterns/Creational/abstract_factory"
	"reflect"
	"testing"
)

func SetupCunstruction(env string) (abstract_factory.Database, abstract_factory.FileSystem) {
	fs := abstract_factory.AbstractFactory("filesystem")
	db := abstract_factory.AbstractFactory("database")

	return db(env).(abstract_factory.Database),
		fs(env).(abstract_factory.FileSystem)
}

func TestAbstractFactory(t *testing.T) {
	fsp, dbp := SetupCunstruction("production")
	fsd, dbd := SetupCunstruction("development")

	fmt.Println(reflect.TypeOf(fsp), reflect.TypeOf(dbp))
	fmt.Println(reflect.TypeOf(fsd), reflect.TypeOf(dbd))
}
