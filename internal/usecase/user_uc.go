package usecase

import (
	"errors"
	"evermos/internal/domain"
	"fmt"
	"strings"

	//"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo  domain.UserRepository
	storeRepo domain.StoreRepository
}

func NewUserUsecase(u domain.UserRepository, s domain.StoreRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo:  u,
		storeRepo: s,
	}
}

func (uc *userUsecase) Register(user *domain.User) error {
	// Cek email
	if existing, _ := uc.userRepo.FindByEmail(user.Email); existing.ID != 0 {
		return errors.New("email already used")
	}
	// Cek no telepon
	if existing, _ := uc.userRepo.FindByPhone(user.Phone); existing.ID != 0 {
		return errors.New("phone already used")
	}

	// Hash password
	// hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return err
	// }
	// user.Password = strings.TrimSpace(string(hash))

	// Simpan user
	if err := uc.userRepo.Save(user); err != nil {
		return err
	}

	// Auto create store (pakai nama user + " Store")
	storeName := user.Name + "'s Store"
	if err := uc.storeRepo.CreateForUser(user.ID, storeName); err != nil {
		return err
	}

	return nil
}

func (uc *userUsecase) Login(email, password string) (*domain.User, error) {
	email=strings.TrimSpace(email)
	password=strings.TrimSpace(password)

	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	fmt.Print("Trying to login:",email)
	fmt.Println("|")
	fmt.Print("Password input:",password)
	fmt.Println("|")
	fmt.Print("Password in DB:",user.Password)
	fmt.Println("|")

	fmt.Print("Password input (raw bytes): ")
	fmt.Println([]byte(password))

	fmt.Print("Password DB (raw bytes): ")
	fmt.Println([]byte(user.Password))
	
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
	// 	fmt.Println("Bcrypt error:", err)
	// 	return nil, errors.New("not hashed")
	// }

	if user.Password != password {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func (uc *userUsecase) GetProfile(userID uint) (*domain.User, error) {
	return uc.userRepo.FindByID(userID)
}

func (uc *userUsecase) Update(userID uint, updated *domain.User) error {
	user, err := uc.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	// Perbarui field yang boleh diubah
	user.Name = updated.Name
	user.Email = updated.Email
	user.Phone = updated.Phone

	return uc.userRepo.Update(userID, user)
}