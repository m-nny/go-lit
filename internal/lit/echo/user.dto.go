package echo

import "github.com/m-nny/go-lit/internal/lit"

type GetUsersResult struct {
	Items []*lit.User `json:"users"`
	Count int         `json:"count"`
}
