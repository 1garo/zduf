package internal

// Teacher schema.
type Teacher struct {
	ID      uint `json:"-"`
	FirstName string    `json:"first_name"      binding:"required"`
	LastName string    `json:"last_name"      binding:"required"`
	Age     uint      `json:"age"  binding:"required"`
	//Teaches []Class   `json:"-"`
}
