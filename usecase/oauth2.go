package usecase

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/debidarmawan/debozero-core/config"
	"github.com/debidarmawan/debozero-core/constants"
	"github.com/debidarmawan/debozero-core/dto"
	"github.com/debidarmawan/debozero-core/global"
	"github.com/debidarmawan/debozero-core/helper"
	"github.com/debidarmawan/debozero-core/model"
	"github.com/debidarmawan/debozero-core/repository"
	"github.com/go-oauth2/mysql/v4"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v2"
	"gorm.io/gorm"
)

type Oauth2UseCase interface {
	GenerateToken(userId string, scope string) (*dto.TokenInfo, error)
	AddClient(request dto.Oauth2Client) (*dto.Oauth2ClientResponse, global.ErrorResponse)
}

type oauth2UseCase struct {
	server           *server.Server
	manager          *manage.Manager
	oauth2ClientRepo repository.Oauth2ClientRepository
}

func NewOauth2UseCase(db *gorm.DB, oauth2ClientRepo repository.Oauth2ClientRepository) Oauth2UseCase {
	m := manage.NewDefaultManager()

	accessTokenExp, err := strconv.Atoi(config.GetEnv(constants.AccessTokenExpiresMinute))
	if err != nil {
		log.Fatal("Invalid " + constants.AccessTokenExpiresMinute)
	}

	refreshTokenExp, err := strconv.Atoi(config.GetEnv(constants.RefreshTokenExpiresDay))
	if err != nil {
		log.Fatal("Invalid " + constants.RefreshTokenExpiresDay)
	}

	testToken := config.GetEnv(constants.TestAccessAndRefreshToken)

	var refreshTokenConfig manage.RefreshingConfig

	if testToken == "TRUE" && config.GetEnv(constants.Environment) == "development" {
		m.SetImplicitTokenCfg(&manage.Config{
			AccessTokenExp:    time.Minute * 1,
			RefreshTokenExp:   time.Minute * 5,
			IsGenerateRefresh: true,
		})
		refreshTokenConfig = manage.RefreshingConfig{
			AccessTokenExp:     time.Minute * 1,
			RefreshTokenExp:    time.Minute * 5,
			IsGenerateRefresh:  true,
			IsRemoveAccess:     true,
			IsRemoveRefreshing: true,
			IsResetRefreshTime: true,
		}
	} else {
		m.SetImplicitTokenCfg(&manage.Config{
			AccessTokenExp:    time.Minute * time.Duration(accessTokenExp),
			RefreshTokenExp:   24 * time.Hour * time.Duration(refreshTokenExp),
			IsGenerateRefresh: true,
		})
		refreshTokenConfig = manage.RefreshingConfig{
			AccessTokenExp:     time.Minute * time.Duration(accessTokenExp),
			RefreshTokenExp:    24 * time.Hour * time.Duration(refreshTokenExp),
			IsGenerateRefresh:  true,
			IsRemoveAccess:     true,
			IsRemoveRefreshing: true,
			IsResetRefreshTime: true,
		}
	}

	m.SetRefreshTokenCfg(&refreshTokenConfig)
	m.MapClientStorage(oauth2ClientRepo)

	sqlDb, _ := db.DB()
	dialect := gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}
	tokenStore := mysql.NewStoreWithOpts(sqlDb, mysql.WithSQLDialect(dialect))
	m.MapTokenStorage(tokenStore)

	s := server.NewDefaultServer(m)
	s.SetAllowGetAccessRequest(true)
	s.SetClientInfoHandler(server.ClientFormHandler)

	return &oauth2UseCase{server: s, manager: m, oauth2ClientRepo: oauth2ClientRepo}
}

func (ou *oauth2UseCase) GenerateToken(userId string, scope string) (*dto.TokenInfo, error) {
	ctx := context.Background()

	tokenGenerateRequest := oauth2.TokenGenerateRequest{
		ClientID: config.GetEnv(constants.ClientId),
		UserID:   userId,
		Scope:    scope,
	}

	ti, err := ou.manager.GenerateAuthToken(ctx, oauth2.Token, &tokenGenerateRequest)
	if err != nil {
		return nil, err
	}

	tokenInfo := dto.TokenInfo{
		AccessToken:  ti.GetAccess(),
		ExpiresAt:    time.Now().Add(ti.GetAccessExpiresIn()),
		RefreshToken: ti.GetRefresh(),
	}

	return &tokenInfo, nil
}

func (ou *oauth2UseCase) AddClient(request dto.Oauth2Client) (*dto.Oauth2ClientResponse, global.ErrorResponse) {
	secretKey := helper.GenerateRandomString(41)

	clientData := model.Oauth2Client{
		ClientId:    request.ClientId,
		Name:        request.Name,
		Secret:      secretKey,
		Domain:      request.Domain,
		Description: request.Description,
	}

	err := ou.oauth2ClientRepo.Create(&clientData)
	if err != nil {
		return nil, global.InternalServerError(err)
	}

	reuslt := dto.Oauth2ClientResponse{
		ClientId:  request.ClientId,
		Name:      request.Name,
		SecretKey: secretKey,
	}

	return &reuslt, nil
}
