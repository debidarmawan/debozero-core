package repository

import (
	"context"

	"debozero-core/global"
	"debozero-core/model"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"gorm.io/gorm"
)

type Oauth2ClientRepository interface {
	GetById(id string) (*model.Oauth2Client, global.ErrorResponse)
	GetByClientId(clientId string) (*model.Oauth2Client, global.ErrorResponse)
	Create(clientDetail *model.Oauth2Client) global.ErrorResponse
	GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error)
}

type oauth2ClientRepository struct {
	db *gorm.DB
}

func NewOauth2ClientRepository(db *gorm.DB) Oauth2ClientRepository {
	return &oauth2ClientRepository{
		db: db,
	}
}

func (or *oauth2ClientRepository) GetById(id string) (*model.Oauth2Client, global.ErrorResponse) {
	var oauth2Client model.Oauth2Client

	err := or.db.Where("id = ?", id).First(&oauth2Client).Error
	if err != nil {
		return nil, global.NotFoundError()
	}

	return &oauth2Client, nil
}

func (or *oauth2ClientRepository) GetByClientId(clientId string) (*model.Oauth2Client, global.ErrorResponse) {
	var oauth2Client model.Oauth2Client

	err := or.db.Where("client_id = ?", clientId).First(&oauth2Client).Error
	if err != nil {
		return nil, global.NotFoundError()
	}

	return &oauth2Client, nil
}

func (or *oauth2ClientRepository) Create(clientDetail *model.Oauth2Client) global.ErrorResponse {
	err := or.db.Create(clientDetail).Error
	if err != nil {
		return global.InternalServerError(err)
	}

	return nil
}

func (or *oauth2ClientRepository) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	var oauth2Client model.Oauth2Client

	err := or.db.Where("client_id = ?", id).First(&oauth2Client).Error
	if err != nil {
		return nil, global.NotFoundError()
	}

	clientInfo := &models.Client{
		ID:     oauth2Client.ClientId,
		Secret: oauth2Client.Secret,
		Domain: oauth2Client.Domain,
	}

	return clientInfo, nil
}
