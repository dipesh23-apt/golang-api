package models

type User struct {
	Id         uint    `gorm:"primary key" json:"id"`
	CreatedAt  int64   `"autoCreateTime" json:"-"`
	UpdatedAt  int64   `"autoUpdateTime" json:"-"`
	First_name string  `json:"first_name" validate:"required"`
	Last_name  string  `json:"last_name" validate:"required"`
	City       string  `json:"city" validate:"required"`
	Phone      string  `json:"phone" validate:"required"`
	Height     float32 `json:"height" validate:"required"`
	Gender     string  `json:"gender" validate:"required"`
	Password   string  `json:"password" validate:"required"`
	Married    bool    `json:"married" validate:"required"`
}

type Body struct {
	Id []uint `json:"ids"`
}

func (User) TableName() string {
	return "user"
}
