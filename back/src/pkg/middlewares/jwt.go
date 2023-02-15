package middlewares

import (
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func JwtMiddleware() (echo.MiddlewareFunc, error) {
	var err error
	issuerURL, err := url.Parse("www.hello.com")
	provider := jwks.NewCachingProvider(issuerURL, 15*time.Second)
	jwtValidator, err := validator.New(provider.KeyFunc, validator.RS256, issuerURL.String(), []string{"****"})
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Something went wrong while processing jwt")
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "No Authorization Header")
			}

			if !strings.HasPrefix(authorization, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization Header")
			}

			token := strings.TrimPrefix(authorization, "Bearer ")

			claims, err := jwtValidator.ValidateToken(c.Request().Context(), token)
			if err != nil {
				log.Println(err)
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token")
			}

			c.Set("claims", claims.(*validator.ValidatedClaims))

			return next(c)
		}
	}, nil

}
