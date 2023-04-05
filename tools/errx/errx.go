package errx

import (
	"errors"

	"google.golang.org/grpc/status"
)

func ConverRpcErr(err error) error {
	return errors.New(status.Convert(err).Message())
}
