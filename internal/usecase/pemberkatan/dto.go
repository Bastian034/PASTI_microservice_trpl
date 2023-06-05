package pemberkatan

import (
	"gereja-services/internal/entities"
	"gereja-services/pkg/response"
)

//REQUEST

type CreateRequest struct {
	entities.PemberkatanEntity
}

type UpdateRequest struct {
	ID string `param:"id" validate:"required"`
}

type DeleteRequest struct {
	ID string `param:"id" validate:"required"`
}

type GetByIDRequest struct {
	ID string `param:"id" validate:"required"`
}

//RESPONSE

type CreateResponse struct {
	entities.PemberkatanEntityModel
}
type CreateResponseDoc struct {
	Body struct {
		Meta response.Meta  `json:"meta"`
		Data CreateResponse `json:"data"`
	} `json:"body"`
}

type UpdateResponse struct {
	entities.PemberkatanEntityModel
}
type UpdateResponseDoc struct {
	Body struct {
		Meta response.Meta  `json:"meta"`
		Data UpdateResponse `json:"data"`
	} `json:"body"`
}

type DeleteResponse struct {
	ID *string `json:"id"`
}
type DeleteResponseDoc struct {
	Body struct {
		Meta response.Meta  `json:"meta"`
		Data DeleteResponse `json:"data"`
	} `json:"body"`
}

type GetByIDResponse struct {
	entities.PemberkatanEntityModel
}

type GetByIDResponseDoc struct {
	Body struct {
		Meta response.Meta   `json:"meta"`
		Data GetByIDResponse `json:"data"`
	} `json:"body"`
}

type GetAllResponse struct {
	Datas []entities.PemberkatanEntityModel `json:"datas"`
}

type GetAllResponseDoc struct {
	Body struct {
		Meta response.Meta  `json:"meta"`
		Data GetAllResponse `json:"data"`
	}
}
