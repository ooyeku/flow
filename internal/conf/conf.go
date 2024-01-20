package conf

// The `dbPath` variable is a string that stores the file path of the database.
var (
	dbPath string = "internal/inmemory/goworkflow.db"
)

// GetDBPath returns the path to the database used for the application.
func GetDBPath() string {
	return dbPath
}
