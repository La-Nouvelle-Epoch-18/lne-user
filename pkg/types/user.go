package types

type User struct {
	ID          string `json:"id" xorm:"pk 'id'"`
	Name        string `json:"name" xorm:"varchar(64) not null"`
	Email       string `json:"email" xorm:"varchar(64) not null unique"`
	Username    string `json:"username" xorm:"varchar(32) not null unique"`
	Password    string `json:"password" xorm:"varchar(64) not null"`
	Type        string `json:"type" xorm:"varchar(32) not null default 'student'"`
	Description string `json:"description" xorm:"varchar(128)"`
}
