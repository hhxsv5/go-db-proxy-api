package drivers

type MySQL struct {
	base
}

func NewMySQL(dsn string) (*MySQL) {
	b := newBase("mysql", dsn)
	return &MySQL{*b}
}
