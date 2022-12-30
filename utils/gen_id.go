package utils

import "github.com/lithammer/shortuuid/v4"

func GenRequestId() string {
	return "req_" + shortuuid.New()
}
