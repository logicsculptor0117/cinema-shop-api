// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	Film "cinema-shop/services/film/api/internal/handler/Film"
	"cinema-shop/services/film/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/list/:page/:limit",
				Handler: Film.FilmListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/detail/:film_id",
				Handler: Film.FilmDetailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/film/v1"),
	)
}
