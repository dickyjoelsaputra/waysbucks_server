package authdto

type RegisterRequest struct {
	Name     string `json:"name" gorm:"type: varchar(255) required"`
	Email    string `json:"email" gorm:"type: varchar(255) required"`
	Password string `json:"password" gorm:"type: varchar(255) required"`
}

type LoginRequest struct {
	Email    string `json:"email" gorm:"type: varchar(255) required"`
	Password string `json:"password" gorm:"type: varchar(255) required"`
}

type LoginResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Token  string `json:"token"`
	Status string `json:"status"`
}

type CheckAuthResponse struct {
	Id     int    `gorm:"type: int" json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}
