package model

type Machine struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(50);unique_index"`
	Host     string `json:"host" gorm:"type:varchar(50)"`
	Ip       string `json:"ip" gorm:"type:varchar(80)"`
	Port     int    `json:"port" gorm:"type:int(6)"`
	User     string `json:"user" gorm:"type:varchar(20)"`
	Password string `json:"password,omitempty" gorm:"type:varchar(20)"`
	Key      string `json:"key,omitempty" gorm:"type:varchar(20)"`
	Type     string `json:"type" gorm:"type:varchar(20)"`
}
