package server

import (
	"os"
	"regexp"

	"github.com/debidarmawan/debozero-core/docs"
	"github.com/debidarmawan/debozero-core/handler"
	"github.com/debidarmawan/debozero-core/helper"
	"github.com/debidarmawan/debozero-core/repository"
	"github.com/debidarmawan/debozero-core/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"gorm.io/gorm"
)

func Routes(f *fiber.App, db *gorm.DB) {

	f.Use(cors.New())
	f.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
	}))

	baseRoute := "/api/v1"

	docs.SwaggerInfo.BasePath = baseRoute

	if os.Getenv("GO_ENV") == "development" {
		f.Get("/swagger/*", fiberSwagger.WrapHandler)
	}

	routerGroup := f.Group(baseRoute)

	routerGroup.Use(PanicHandler)

	redactedInfoPaths := []string{
		"/api/v1/auth/login",
		"/api/v1/auth/refresh",
		"/api/v1/users/reset-password",
		"/api/v1/users/change-password",
		"/api/v1/users/register",
	}

	f.Use(logger.New(logger.Config{
		Format: "[${time}] ${method} ${path} - ${status} | ${latency} | ReqId ${reqHeader:X-RequestId} \n[Request Headers] ${reqHeaders}\n[Request Query Params] ${queryParams}\n[Request Body] ${maskedRequestBody}\n[Response Body] ${maskedResponseBody}\n\n",
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/api/v1/auth/verify"
		},
		CustomTags: map[string]logger.LogFunc{
			"maskedRequestBody": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				if !helper.ContainsString(redactedInfoPaths, c.Path()) {
					return output.Write([]byte(c.Request().Body()))
				}

				if c.Body() == nil || len(c.Body()) == 0 {
					return output.Write([]byte("{}"))
				}
				body := string(c.Body())
				sensitiveFields := []string{"password", "old_password", "new_password", "confirm_password", "pin", "new_pin"}

				for _, field := range sensitiveFields {
					re := regexp.MustCompile(`"` + field + `"\s*:\s*"[^"]*"`)
					body = re.ReplaceAllString(body, `"`+field+`":"[REDACTED]"`)
				}
				return output.Write([]byte(body))
			},
			"maskedResponseBody": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				if !helper.ContainsString(redactedInfoPaths, c.Path()) {
					return output.Write([]byte(c.Response().Body()))
				}
				if c.Response().Body() == nil || len(c.Response().Body()) == 0 {
					return output.Write([]byte("{}"))
				}
				body := string(c.Response().Body())

				// TODO: Activate token redaction after expiration token issuse is resolved
				// REF: https://bitbucket.org/mp-digital-initiative/gpos-b2b-account-service/pull-requests/155#comment-thread-436801379
				// sensitiveFields := []string{"access_token", "refresh_token"}

				sensitiveFields := []string{}

				for _, field := range sensitiveFields {
					re := regexp.MustCompile(`"` + field + `"\s*:\s*"[^"]*"`)
					body = re.ReplaceAllString(body, `"`+field+`":"[REDACTED]"`)
				}
				return output.Write([]byte(body))
			},
		},
	}))

	txManager := helper.NewTxManager(db)

	// INIT REPOSITORY
	userRepo := repository.NewUserRepository(db)
	oauth2ClientRepo := repository.NewOauth2ClientRepository(db)
	roleRepo := repository.NewRoleRepository(db)

	// INIT USECASE
	userUseCase := usecase.NewUserUseCase(txManager, userRepo)
	oauth2UseCase := usecase.NewOauth2UseCase(db, oauth2ClientRepo)
	roleUseCase := usecase.NewRoleUseCase(roleRepo)
	authUseCase := usecase.NewAuthUseCase(txManager, userRepo, oauth2UseCase, roleUseCase)

	// INIT HANDLER
	userHandler := handler.NewUserHanler(userUseCase)
	authHandler := handler.NewAuthHandler(authUseCase)
	oauth2Handler := handler.NewOauth2Handler(oauth2UseCase)

	// ROUTING HANDLER
	userHandler.Routes(routerGroup)
	authHandler.Routes(routerGroup)
	oauth2Handler.Routes(routerGroup)
}
