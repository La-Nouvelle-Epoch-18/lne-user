package types

type User struct {
	ID          string `xorm:"pk 'id'"`
	Name        string `xorm:"varchar(64) not null"`
	Email       string `xorm:"varchar(64) not null unique"`
	Username    string `xorm:"varchar(32) not null unique"`
	Password    string `xorm:"varchar(64) not null"`
	Type        string `xorm:"varchar(32) not null default 'student'"`
	Description string `xorm:"varchar(128)"`
}
