package user

type User struct {
	Phone     string `gorm:"primaryKey"`
	SessionId string `gorm:"index"`
	Code      int
}
