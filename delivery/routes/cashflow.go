package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterCashflowRoutes(v1 fiber.Router, handler *http.CashflowHttp) {
	adminOnly := middleware.JWTMiddleware([]int{1})
	rtLeaderOnly := middleware.JWTMiddleware([]int{3})

	v1.Get("/finances/cashflow", adminOnly, handler.GetData)
	v1.Get("/finances/cashflow/:rt_profile_id/rt", adminOnly, rtLeaderOnly, handler.GetDataByRTProfileId)
	v1.Get("/finances/cashflow/:id", adminOnly, rtLeaderOnly, handler.GetDataById)
	v1.Put("/finances/cashflow/:id", adminOnly, rtLeaderOnly, handler.UpdateData)
	v1.Post("/finances/cashflow", handler.StoreData)
	v1.Delete("/finances/cashflow/:id", adminOnly, rtLeaderOnly, handler.DeleteData)
}
