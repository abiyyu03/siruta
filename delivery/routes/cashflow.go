package routes

import (
	"github.com/abiyyu03/siruta/delivery/http"
	"github.com/abiyyu03/siruta/delivery/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterCashflowRoutes(v1 fiber.Router, handler *http.CashflowHttp) {
	v1.Get("/finances/cashflow", middleware.JWTMiddleware([]int{1}), handler.GetData)
	v1.Get("/finances/cashflow/:rt_profile_id/rt", middleware.JWTMiddleware([]int{1, 3}), handler.GetDataByRTProfileId)
	v1.Get("/finances/cashflow/:id", middleware.JWTMiddleware([]int{1, 3}), handler.GetDataById)
	v1.Put("/finances/cashflow/:id", middleware.JWTMiddleware([]int{1, 3}), handler.UpdateData)
	v1.Post("/finances/cashflow", middleware.JWTMiddleware([]int{1, 3}), handler.StoreData)
	v1.Delete("/finances/cashflow/:id", middleware.JWTMiddleware([]int{1, 3}), handler.DeleteData)
}
