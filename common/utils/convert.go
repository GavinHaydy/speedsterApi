package utils

import "database/sql"

func ToNullString(v *string) sql.NullString {
	if v == nil {
		return sql.NullString{}
	}

	return sql.NullString{
		String: *v,
		Valid:  true,
	}
}

func NullString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}
