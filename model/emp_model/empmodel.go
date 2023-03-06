package empmodel

type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Userid int    `json:"user_id"`
}
