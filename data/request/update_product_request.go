package request

type UpdateProductRequestBody struct {
	Name           string   `json:"name"  binding:"omitempty,min=3,max=50"`
	Brand          string   `json:"brand" binding:"omitempty,min=3,max=50"`
	Size           string   `json:"size"`
	Price          float64  `json:"price" binding:"omitempty,min=1.00,max=99999999.00"`
	MainImageUrl   string   `json:"mainImageUrl"`
	OtherImagesUrl []string `json:"otherImagesUrl"`
}
