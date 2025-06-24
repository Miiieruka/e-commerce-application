package handlers

type CreateRoomRequest struct {
	UserID      int    `json:"user_id" binding:"required"`
	RoomName    string `json:"room_name" binding:"required"`
	Description string `json:"description" binding:"required"`
	BuyerID     int    `json:"buyer_id" binding:"required"`
	SellerID    int    `json:"seller_id" binding:"required"`
}

type SendMessageRequest struct {
	RoomID  int    `json:"room_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}
