package models

type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// 检验用户是否存在
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{
		Username:username,
		Password: password,
	}).First(&auth)
	//log.Printf("get the auth: %v", auth)
	if auth.ID > 0 {
		return true
	}

	return false
}
