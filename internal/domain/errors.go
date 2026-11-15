package domain

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrRateLimited  = errors.New("rate limited")
	ErrNotFound     = errors.New("not found")
	ErrUpstream     = errors.New("upstream error")
)