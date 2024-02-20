package model

type Oauth2Client struct {
	AdminBaseModel
	ClientId    string `gorm:"type:varchar(255);not null"`
	Name        string `gorm:"type:varchar(255);not null"`
	Secret      string `gorm:"type:varchar(255);not null"`
	Domain      string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
}
