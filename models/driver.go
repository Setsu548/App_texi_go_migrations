package models

type ListDriver struct {
	UUIDModel   `gorm:"embedded"`
	FirstName   string `gorm:"column:first_name;size:100;not null" json:"first_name"`
	LastName    string `gorm:"column:last_name;size:100;not null" json:"last_name"`
	Address     string `gorm:"column:address;size:255;not null" json:"address"`
	Email       string `gorm:"column:email;size:100;not null;unique" json:"email"`
	PhoneNumber string `gorm:"column:phone_number;size:20;not null" json:"phone_number"`
	IsVerified  bool   `gorm:"column:is_verified;not null;default:false" json:"is_verified"`
	Status      string `gorm:"column:status;size:20;not null;default:'active'" json:"status"`
	BaseModel   `gorm:"embedded"`
}

type ListDrivers []ListDriver

// type Driver struct {
// 	UUIDModel   `gorm:"embedded"`
// 	FirstName   string `gorm:"column:first_name;size:100;not null" json:"first_name"`
// 	LastName    string `gorm:"column:last_name;size:100;not null" json:"last_name"`
// 	Address     string `gorm:"column:address;size:255;not null" json:"address"`
// 	Email       string `gorm:"column:email;size:100;not null;unique" json:"email"`
// 	PhoneNumber string `gorm:"column:phone_number;size:20;not null" json:"phone_number"`
// 	IsVerified  bool   `gorm:"column:is_verified;not null;default:false" json:"is_verified"`
// 	Status      string `gorm:"column:status;size:20;not null;default:'active'" json:"status"`
// 	BaseModel   `gorm:"embedded"`
// }
