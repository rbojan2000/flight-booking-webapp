package model

type User struct {
	ID        uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string   `json:"name" gorm:"not null;type:string"`
	Surname   string   `json:"surname" gorm:"not null;type:string"`
	Email     string   `json:"email" gorm:"not null;type:string"`
	Password  string   `json:"password" gorm:"not null;type:string"`
	Type      UserType `json:"type" gorm:"not null;type:int"`
	Address   Location `json:"address" gorm:"foreignKey:AddressID"`
	AddressID uint     `json:"addressID" gorm:"index"`
}
