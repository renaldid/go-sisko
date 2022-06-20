package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/renaldid/go-sisko/controller"
	"github.com/renaldid/go-sisko/exception"
)

func NewRouter(siswaController controller.SiswaController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/siswas", siswaController.FindByAll)
	router.GET("/api/siswas/:siswaId", siswaController.FindById)
	router.POST("/api/siswas/", siswaController.Create)
	router.PUT("/api/siswas/:siswaId", siswaController.Update)
	router.DELETE("/api/siswas/:siswaId", siswaController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
