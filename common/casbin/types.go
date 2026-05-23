package casbin

type Policy struct {
	RoleCode string `db:"code"`
	Path     string `db:"path"`
	Method   string `db:"method"`
}
