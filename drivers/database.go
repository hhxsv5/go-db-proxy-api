package drivers

type Database interface {
	Open() error
	Query(q string) ([]map[string]interface{})
	Close()
}
