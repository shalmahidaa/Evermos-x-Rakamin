package domain

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Phone     string    `gorm:"unique" json:"phone"`
	Password  string    `json:"password"`
	Role      string    `gorm:"default:user" json:"role"` // "admin" atau "user"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Store     Store     `json:"store"` // One-to-one
}

type UserRepository interface {
	Save(user *User) error
	FindByEmail(email string) (*User, error)
	FindByPhone(phone string) (*User, error)
	FindByID(id uint) (*User, error)
	Update(userID uint, user *User) error
}

type UserUsecase interface {
	Register(user *User) error
	Login(email, password string) (*User, error)
	GetProfile(userID uint) (*User, error)
	Update(userID uint, updated *User) error
}