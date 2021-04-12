package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gitlab.com/cfs-service/store"
	"gitlab.com/cfs-service/utils"
)

const MaxPagingLimit = uint64(200) // TODO: Make this confirable

func GetPagingParam(c *fiber.Ctx) (*store.PagingOptions, error) {
	opts := &store.PagingOptions{}

	var err error
	offsetText := c.Query("offset", "0") // Default offset is 0
	opts.Offset, err = strconv.ParseUint(offsetText, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "Invalid offset")
	}

	limitText := c.Query("limit", "50") // Default is 50
	opts.ItemsPerPage, err = strconv.ParseUint(limitText, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "Invalid limit")
	}

	if opts.ItemsPerPage > MaxPagingLimit {
		return nil, errors.Wrap(err, fmt.Sprintf("Limit exceeded the maximum of %d", MaxPagingLimit))
	}

	return opts, nil
}

func GetQueryParamDate(c *fiber.Ctx, paramName string) (time.Time, error) {
	fromText := c.Query(paramName)
	return utils.StringToTime(fromText)
}

func GetSortParam(c *fiber.Ctx) string {
	order := "ASC" // Default

	if param := c.Query("sort"); len(param) > 0 && strings.ToLower(param) == "desc" {
		order = "DESC"
	}

	return order
}

func GetAgencyIDFromCtx(c *fiber.Ctx) string {
	return c.Context().UserValue("agency_id").(string)
}

func MarshalThenSend(c *fiber.Ctx, data interface{}) error {
	blob, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return c.SendString(string(blob))
}
