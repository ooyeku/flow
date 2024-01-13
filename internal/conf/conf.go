package conf

// The `once` variable is a sync.Once type that guarantees an initialization function is executed only once.
// The `dbPath` variable is a string that stores the file path of the database.
// The `taskRouter` variable is a pointer to a handle.TaskControl struct that handles task-related operations.
var (
	dbPath string = "internal/inmemory/goworkflow.db"
)

// GetDBPath returns the path of the database used by the application.
func GetDBPath() string {
	return dbPath
}
