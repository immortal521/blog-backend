package handler

import (
	"blog-server/api/response"
	"blog-server/config"
	"blog-server/errs"
	"blog-server/logger"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(log logger.Logger, cfg *config.Config) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		appErr := errs.ToAppError(err)
		httpCode := errs.MapToHTTPStatus(appErr.Code)

		msg := appErr.Msg
		if httpCode == 500 {
			msg = "Internal Server Error"
		}

		reqID := c.Get("X-Request-ID")

		errLogger := log.WithFields(
			logger.String("request_id", reqID),
			logger.String("method", c.Method()),
			logger.String("remote_ip", c.IP()),
			logger.String("path", c.Path()),
			logger.String("original_url", c.OriginalURL()),
		)

		if cfg.App.Environment == config.EnvDev {
			errLogger.Error(appErr.FormatStack())
		} else {
			errLogger.Error(msg, logger.Error(err))
		}

		writeErr := c.Status(httpCode).JSON(response.Error(appErr.Code, msg))
		if writeErr != nil {
			return c.Status(httpCode).SendString(msg)
		}

		return nil
	}
}
