package errorx

import (
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(code int) error {
	return status.Error(
		codes.Internal,
		strconv.Itoa(code),
	)
}
