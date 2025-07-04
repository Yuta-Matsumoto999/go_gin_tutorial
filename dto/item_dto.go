package dto

type CreateItemInput struct {
	Name        string `json:"name" binding:"required,min=2,max=1000"`
	Price       uint   `json:"price" binding:"required,min=1,max=9999999999"`
	Description string `json:"description" binding:"min=1,max=1000"`
}

type UpdateItemInput struct {
	Name        string `json:"name,omitempty" binding:"omitempty,min=2,max=1000"`
	Price       uint   `json:"price,omitempty" binding:"omitempty,min=1,max=9999999999"`
	Description string `json:"description,omitempty" binding:"omitempty,min=1,max=1000"`
	SoldOut     bool   `json:"sold_out,omitempty"`
}
