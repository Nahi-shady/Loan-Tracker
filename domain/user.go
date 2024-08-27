package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Firstname          string             `json:"firstname" bson:"firstname"`
	Lastname           string             `json:"lastname" bson:"lastname"`
	Username           string             `json:"username" bson:"username"`
	Password           string             `json:"password" bson:"password"`
	Email              string             `json:"email" bson:"email"`
	Bio                string             `json:"bio" bson:"bio"`
	ProfilePicture     string             `json:"profile_picture" bson:"profile_picture"`
	ContactInformation string             `json:"contact_information" bson:"contact_information"`
	IsAdmin            bool               `json:"isAdmin" bson:"isAdmin"`
	Active             bool               `json:"active" bson:"active"`
	VerificationToken  string             `json:"verification_token" bson:"verification_token"`
	TokenExpiry        time.Time          `json:"token_expiry" bson:"token_expiry"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at"`
}

type UpdateRequest struct {
	Firstname          string    `json:"firstname" bson:"firstname"`
	Lastname           string    `json:"lastname" bson:"lastname"`
	Username           string    `json:"username" bson:"username"`
	Bio                string    `json:"bio" bson:"bio"`
	ProfilePicture     string    `json:"profile_picture" bson:"profile_picture"`
	ContactInformation string    `json:"contact_information" bson:"contact_information"`
	VerificationToken  string    `json:"verification_token" bson:"verification_token"`
	TokenExpiry        time.Time `json:"token_expiry" bson:"token_expiry"`
}

type UserUsecase interface {
	SignUp(ctx context.Context, req SignupRequest) (SignupResponse, error)
	GetByUsernameOrEmail(ctx context.Context, identifier string) (User, error)
	UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser UpdateRequest) error
	Login(ctx context.Context, req LoginRequest) (LoginResponse, error)
	Logout(ctx context.Context, userID string) error
	PromoteDemote(ctx context.Context, identifier string, action string) error
	VerifyEmail(ctx context.Context, email, token string) error
	// RequestPasswordReset(ctx context.Context, email, frontendBaseURL string) error
	// ResetPassword(ctx context.Context, req ResetPasswordRequest) error
}

type UserRepository interface {
	Signup(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetByID(ctx context.Context, id string) (User, error)
	UpdatePassword(ctx context.Context, email, newPassword string) error
	UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser UpdateRequest) error
	PromoteDemote(ctx context.Context, userID primitive.ObjectID, action string) error
	SetVerificationToken(ctx context.Context, email, token string, expiry time.Time) error
	VerifyEmailToken(ctx context.Context, email, token string) (bool, error)
}
