package user

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/salmanfr/catatann-api/pkg/common"
	"github.com/salmanfr/catatann-api/pkg/entities"
	"github.com/salmanfr/catatann-api/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	GoogleSignin(code string) (*models.SigninResponse, *models.CustomHttpErrors)
	Signup(dto models.SignupDto) (*models.SigninResponse, error)
	Signin(dto models.SigninDto) (*models.SigninResponse, *models.CustomHttpErrors)
	GetSelf(access_token string) (*entities.User, *models.CustomHttpErrors)
	RefreshToken(refresh_token string) (*models.SigninResponse, *models.CustomHttpErrors)
	InvalidateRefreshToken(refresh_token string) error
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		db,
	}
}

func (s *service) GetSelf(access_token string) (*entities.User, *models.CustomHttpErrors) {
	claims, err := common.VerifyJwt(access_token, os.Getenv("USER_ACCESS_TOKEN_JWT_SECRET"))
	
	var user entities.User
	
	if err != nil {
		return nil, models.CreateCustomHttpError(http.StatusUnauthorized, err)
	}
	
	user_id := claims["sub"].(string)

	res := s.db.First(&user, "user_id = ?", user_id)

	if res.Error != nil {
		return nil, models.CreateCustomHttpError(http.StatusNotFound, "user not registered")
	}
	
	user.Password = ""
	
	return &user, nil
}

func (s *service) RefreshToken(refresh_token string) (*models.SigninResponse, *models.CustomHttpErrors) {
	var token entities.Token
	var user entities.User
	
	claims, err := common.VerifyJwt(refresh_token, os.Getenv("USER_REFRESH_TOKEN_JWT_SECRET"))

	if err != nil {
		return nil, models.CreateCustomHttpError(http.StatusUnauthorized, err)
	}
	
	token_id := claims["sub"].(string)
	
	res := s.db.First(&token, "token_id = ?", token_id)

	if res.Error != nil {
		return nil, models.CreateCustomHttpError(http.StatusUnauthorized, "session has expired")
	}

	res = s.db.First(&user, "user_id = ? AND disabled_at IS NULL", token.UserId)
	
	if res.Error != nil {
		return nil, models.CreateCustomHttpError(http.StatusNotFound, "user not registered")
	}
	
	if err := s.InvalidateRefreshToken(token.TokenId); err != nil {
		return nil, models.CreateCustomHttpError(http.StatusInternalServerError, err)
	}

	// ? Save Refresh Token to get the auto generated token_id
	newRefreshToken := entities.Token{
		UserId: user.UserId,
	}
	
	res = s.db.Create(&newRefreshToken) 
	
	if res.Error != nil {
		return nil, models.CreateCustomHttpError(http.StatusInternalServerError, "internal server error")
	}

	// ? New Refresh Token
	newRft, err := common.GenerateJwt(time.Hour * 1, newRefreshToken.TokenId, os.Getenv("USER_REFRESH_TOKEN_JWT_SECRET")) 
	// ? New Acces Token
	newAct, err := common.GenerateJwt(time.Hour * 1, user.UserId, os.Getenv("USER_ACCESS_TOKEN_JWT_SECRET")) 
	
	if err != nil {
		return nil, models.CreateCustomHttpError(http.StatusInternalServerError, "internal server error")
	}
	
	user.Password = ""
	
	return &models.SigninResponse{User: user, AccessToken: newAct, RefreshToken: newRft}, nil
}

func (s *service ) GoogleSignin(code string) (*models.SigninResponse, *models.CustomHttpErrors) {
	token_res, err := common.GetGoogleOAuthToken(code)
	
	if err != nil {
		return nil, models.CreateCustomHttpError(http.StatusForbidden, err.Error())
	}

	google_user, err := common.GetGoogleUser(token_res.AccessToken, token_res.IDToken)

	if err != nil {
		return nil, models.CreateCustomHttpError(http.StatusUnauthorized, err.Error())
	}

	var users []entities.User
	
	res := s.db.Find(&users, "email = ?", google_user.Email).Limit(1)

	if res.Error != nil {
		return nil, models.CreateCustomHttpError(500, res.Error.Error())
	}
	
	if len(users) == 0 {
		var new_user entities.User

		new_user.Email = google_user.Email
		new_user.FullName = google_user.Name
		new_user.Provider = "google"
		new_user.PictureUrl = google_user.Picture

		res := s.db.Create(&new_user)

		if res.Error != nil {
			return nil, models.CreateCustomHttpError(http.StatusInternalServerError, res.Error.Error())
		}

		users = []entities.User{new_user}
	} else {
		users[0].FullName = google_user.Name
		users[0].PictureUrl = google_user.Picture

		res := s.db.Save(&users[0])

		if res.Error != nil {
			return nil, models.CreateCustomHttpError(http.StatusInternalServerError, res.Error.Error())
		}
	}

	user := users[0]
	
	access_token, err := common.GenerateJwt(time.Minute * 10, user.UserId, os.Getenv("USER_ACCESS_TOKEN_JWT_SECRET"))
	refresh_token, err := s.createRefreshToken(user.UserId)

	if err != nil {
		return nil, models.CreateCustomHttpError(http.StatusInternalServerError, err)
	}
	
	user.Password = ""
	
	response := models.SigninResponse{
		User: user,
		AccessToken: access_token,
		RefreshToken: refresh_token,
	}
	
	return &response, nil
}

func (s *service) Signup(dto models.SignupDto) (*models.SigninResponse, error) {
	var existing_user entities.User

	res := s.db.First(&existing_user, "email = ?", dto.Email)

	fmt.Println("user found ", existing_user)
	
	// ? Encountered error besides not found
	if res.Error != nil && res.Error.Error() != "record not found" {
		return nil, res.Error
	}

	// ? User exists
	if existing_user.UserId != "" {
		return nil, errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	
	if err != nil {
		return nil, err
	}

	user := entities.User{
		Email: dto.Email,
		FullName: dto.FullName,
		Password: string(hash),
	}
	
	res = s.db.Create(&user)
	
	if res.Error != nil {
		return nil, res.Error
	}
	
	user.Password = ""
	
	return &models.SigninResponse{User: user, AccessToken: ""}, nil
}

func (s *service) Signin(dto models.SigninDto) (*models.SigninResponse, *models.CustomHttpErrors) {
	var user entities.User

	res := s.db.First(&user, "email = ? AND disabled_at IS NULL", dto.Email)

	if res.Error != nil && res.Error.Error() != "record not found" {
		return nil, models.CreateCustomHttpError(http.StatusInternalServerError, "internal server error")
	}

	if user.UserId == "" {
		return nil,  models.CreateCustomHttpError(http.StatusNotFound, "user not registered")
	}
	
	if user.Provider != "local" {
		return nil,  models.CreateCustomHttpError(http.StatusBadRequest, "incorrect email/password")
	}
	
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))

	if err != nil {
		return nil,  models.CreateCustomHttpError(http.StatusBadRequest, "incorrect email/password")
	}

	user.Password = ""
	user.Provider = ""
	
	access_token, err := common.GenerateJwt(time.Minute * 10, user.UserId, os.Getenv("USER_ACCESS_TOKEN_JWT_SECRET"))
	refresh_token, err := s.createRefreshToken(user.UserId)

	if err != nil {
		return nil, models.CreateCustomHttpError(http.StatusInternalServerError, err)
	}

	if err != nil {
		return nil,  models.CreateCustomHttpError(http.StatusInternalServerError, "unable to generate tokens")
	}
	
	return &models.SigninResponse{User: user, AccessToken: access_token, RefreshToken: refresh_token}, nil
}

func (s *service) createRefreshToken(user_id string) (string, error) {
	var refresh_token entities.Token

	refresh_token.UserId = user_id
	
	res := s.db.Create(&refresh_token)

	if res.Error != nil {
		return "", res.Error
	}
	
	token, err := common.GenerateJwt(time.Hour * 24 * 365, refresh_token.TokenId, os.Getenv("USER_REFRESH_TOKEN_JWT_SECRET"))
	
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) InvalidateRefreshToken(token_id string) error {
	res := s.db.Model(&entities.Token{}).Where("token_id = ?", token_id).Delete(&entities.Token{})

	return res.Error
}