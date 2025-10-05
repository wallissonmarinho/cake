package controllers

import (
	"github.com/wallissonmarinho/cake/models"
	"github.com/wallissonmarinho/cake/views"

	"github.com/gofiber/fiber/v2"
)

func selectLayout(isStaff, isUserProfile bool) string {
	if isStaff {
		return views.BaseLayout
	}
	if isUserProfile {
		return views.StoreLayout
	}
	return views.BaseLayout
}

func getUserID(ctx *fiber.Ctx) uint {
	profile, ok := ctx.Locals("Profile").(*models.Profile)
	if !ok {
		return 0
	}
	return profile.ID
}
