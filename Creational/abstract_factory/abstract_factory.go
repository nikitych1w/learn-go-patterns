package abstract_factory

import "fmt"

type (
	mongoDB struct {
		database map[string]string
	}

	sqlite struct {
		database map[string]string
	}

	file struct {
		name    string
		content string
	}

	ntfs struct {
		files map[string]file
	}

	ext4 struct {
		files map[string]file
	}

	Database interface {
		GetData(string) string
		PutData(string, string)
	}

	FileSystem interface {
		CreateFile(string)
		FindFile(string) file
	}

	Factory func(string) interface{}
)

func (mdb mongoDB) GetData(query string) string {
	if _, ok := mdb.database[query]; !ok {
		return ""
	}
	fmt.Println("MongoDB")
	return mdb.database[query]
}

func (sql sqlite) GetData(query string) string {
	if _, ok := sql.database[query]; !ok {
		return ""
	}

	fmt.Println("Sqlite")
	return sql.database[query]
}

func (mdb mongoDB) PutData(query string, data string) {
	mdb.database[query] = data
}

func (sql sqlite) PutData(query string, data string) {
	sql.database[query] = data
}

func (ntfs ntfs) CreateFile(path string) {
	ntfs.files[path] = file{
		name:    path,
		content: "NTFS file",
	}
	fmt.Println("NTFS")
}

func (ext ext4) CreateFile(path string) {
	ext.files[path] = file{
		name:    path,
		content: "EXT4 file",
	}
	fmt.Println("EXT4")
}

func (ntfs ntfs) FindFile(query string) file {
	if _, ok := ntfs.files[query]; !ok {
		return file{}
	}
	return ntfs.files[query]
}

func (ext ext4) FindFile(query string) file {
	if _, ok := ext.files[query]; !ok {
		return file{}
	}
	return ext.files[query]
}

func AbstractFactory(fact string) Factory {
	switch fact {
	case "database":
		return DatabaseFactory
	case "filesystem":
		return FileSystemFactory
	default:
		return nil
	}
}

func DatabaseFactory(env string) interface{} {
	switch env {
	case "production":
		return mongoDB{database: map[string]string{}}
	case "development":
		return sqlite{database: map[string]string{}}
	default:
		return nil
	}
}

func FileSystemFactory(env string) interface{} {
	switch env {
	case "production":
		return ntfs{files: map[string]file{}}
	case "development":
		return ext4{files: map[string]file{}}
	default:
		return nil
	}
}
