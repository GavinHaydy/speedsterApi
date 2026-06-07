package utils

import "strings"

func EmptyToNil(v *string) *string {
	if v == nil {
		return nil
	}

	if strings.TrimSpace(*v) == "" {
		return nil
	}

	return v
}
