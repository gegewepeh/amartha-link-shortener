package network

import (
	"encoding/json"

	utils "link-shortener/internal/pkg/utils"
)

// Meta is metadata of a JSON response
type Meta struct {
	Total      int    `json:"total,omitempty"`
	TotalPages int    `json:"totalPages,omitempty"`
	Page       int    `json:"page,omitempty"`
	PerPage    int    `json:"perPage,omitempty"`
	NextQuery  string `json:"nextQuery,omitempty"`
}

// Response format in JSON
type Response struct {
	Code int         `json:"-"`
	Data interface{} `json:"data"`
	Meta *Meta       `json:"meta,omitempty"`
}

// ToJSON format error as json response
func (resp *Response) ToJSON() ([]byte, *utils.AppError) {
	r, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		return nil, &utils.AppError{
			Code:      500,
			ErrorCode: "gfx5000000",
			Err:       jsonErr,
		}
	}

	return r, nil
}
