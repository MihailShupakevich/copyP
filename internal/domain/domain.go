package domain

type User struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"unique"`
	Age      int    `json:"age"`
	UserName string `json:"username" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
	Posts    []Post `json:"posts" gorm:"foreignKey:IdUser;references:Id"`
}

type Post struct {
	Id      int    `json:"id" gorm:"primary_key"`
	Title   string `json:"title" gorm:"unique"`
	Content string `json:"content" binding:"required"`
	IdUser  int    `json:"id_user"`
}
