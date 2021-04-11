package middleware

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/cfs-service/service"
)

func AuthByHTTPSecureCookie(c *fiber.Ctx) error {
	token := c.Cookies("auth-token")

	data, err := service.ExtractTokenMetadata(token)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	c.Context().SetUserValue("agency_id", data.AgencyID)

	return c.Next()
}
