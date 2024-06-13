package request

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`                  // golang to json
	Currency string `json:"currency" binding:"required,oneof=USD EUR"` // golang to json
}

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type ListAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
