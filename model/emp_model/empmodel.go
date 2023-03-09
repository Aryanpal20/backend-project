package empmodel

type Employee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `gorm:"Unique:email"`
	Phone_No string `gorm:"Unique:phone_no"`
	Userid   int    `json:"user_id"`
}
