package handlers

import (
	"encoding/json"
	"errors"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
	"gitlab.com/cfs-service/server/views"
	"gitlab.com/cfs-service/store"
)

var dbStore store.IStore

func Initialize(s store.IStore) {
	dbStore = s
}

func PostEvent(c *fiber.Ctx) error {
	ev := &views.Event{}
	if err := c.BodyParser(ev); err != nil {
		pp.Println("err:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if errs := ev.Validate(); len(errs) > 0 {
		msg, _ := json.Marshal(errs)
		return errors.New("Invalid input:" + string(msg))
	}

	// pp.Println("ev:", ev)
	m, err := ev.ToStoreModel()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := dbStore.AddEvent(m); err != nil {
		logrus.Error("Save event to storage failed", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Save event to storage failed")
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetEvents(c *fiber.Ctx) error {
	// Get paging param (common)
	pagingOpts, err := GetPagingParam(c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid paging param")
	}

	//  Get range param
	from, err := GetQueryParamDate(c, "from")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid from")
	}

	to, err := GetQueryParamDate(c, "to")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid to")
	}

	searchResult, err := dbStore.SearchEventByTimeRange(&store.SearchEventOptions{
		AgentcyID:  GetAgencyIDFromCtx(c),
		From:       from,
		To:         to,
		PagingOpts: pagingOpts,
	})
	if err != nil {
		logrus.Error("Search events failed", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Search events failed")
	}

	pp.Println("events:", searchResult)
	reponseModel := views.GetEventResponseFromStoreResponse(searchResult)

	if err := MarshalThenSend(c, reponseModel); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error when sending response")
	}

	return nil
}
