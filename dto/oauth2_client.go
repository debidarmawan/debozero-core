package dto

type Oauth2Client struct {
	ClientId    string `json:"client_id" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type Oauth2ClientResponse struct {
	ClientId  string `json:"client_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	SecretKey string `json:"secret_key"`
}
