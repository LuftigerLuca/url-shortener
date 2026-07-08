package model

type CreateUrlRequest struct {
	Url      string `json:"url" binding:"required"`
	Lifespan *int   `json:"lifespan"`
}

type DeleteUrlRequest struct {
	Short string `json:"short" binding:"required"`
}
