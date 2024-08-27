package user_usecase

import (
	"context"
	"errors"
	"loan-tracker/domain"
	"loan-tracker/infrastructure/validation"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserUsecase) SignUp(ctx context.Context, req domain.SignupRequest) (domain.SignupResponse, error) {
	err := validation.ValidateEmail(req.Email)
	if err != nil {
		return domain.SignupResponse{}, err
	}

	if _, err := u.userRepo.GetByEmail(ctx, req.Email); err == nil {
		return domain.SignupResponse{}, errors.New("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.SignupResponse{}, err
	}

	token := generateVerificationToken()
	tokenExpiry := time.Now().Add(24 * time.Hour)

	user := &domain.User{
		ID:                 primitive.NewObjectID(),
		Firstname:          req.Firstname,
		Lastname:           req.Lastname,
		Username:           req.Username,
		Password:           string(hashedPassword),
		Email:              req.Email,
		Bio:                "",
		ProfilePicture:     "",
		ContactInformation: "",
		IsAdmin:            false,
		CreatedAt:          time.Now(),
		VerificationToken:  token,
		TokenExpiry:        tokenExpiry,
		Active:             false,
	}

	err = u.userRepo.Signup(ctx, user)
	if err != nil {
		return domain.SignupResponse{}, err
	}

	err = u.emailService.SendEmailVerification(user.Email, token)
	if err != nil {
		u.DeleteUser(ctx, user.ID)
		return domain.SignupResponse{}, err
	}

	return domain.SignupResponse{Success: true, Message: "Verification email sent."}, nil
}

func generateVerificationToken() string {
	return primitive.NewObjectID().Hex()
}
