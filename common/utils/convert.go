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
