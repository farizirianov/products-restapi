package request

type CreateProductRequestBody struct {
	Sku            string   `json:"sku"   binfing:"required"`
	Name           string   `json:"name"  binding:"required,min=3,max=50"`
	Brand          string   `json:"brand" binding:"required,min=3,max=50"`
	Size           string   `json:"size"  binding:"required"`
	Price          float64  `json:"price" binding:"required,min=1.00,max=99999999.00"`
	MainImageUrl   string   `json:"mainImageUrl" binding:"required"`
	OtherImagesUrl []string `json:"otherImagesUrl"`
}
