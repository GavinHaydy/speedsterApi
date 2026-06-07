package errorx

import (
	"strconv"

	"google.golang.org/grpc/status"

	"speedsterApi/common/errno"
)

func Parse(err error) (int, string) {
	if err == nil {
		return errno.Ok, ""
	}

	st, ok := status.FromError(err)
	if !ok {
		return errno.ErrServer, err.Error()
	}

	code, e := strconv.Atoi(st.Message())
	if e != nil {
		return errno.ErrServer, err.Error()
	}

	msg := errno.CodeAlertMap[code]

	if msg == "" {
		msg = errno.CodeAlertMap[errno.ErrServer]
	}

	return code, msg
}
