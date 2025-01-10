package createAuction

import (
	resp "auctions/internal/server/lib/response"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	ProductName string `json:"product_name" validate:"required"`
	Account     string `json:"account" validate:"required"`
}

type Response struct {
	resp.Response
	TXHash string `json:"tx_hash"`
}

// template. may be will add dataService
func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.create.New"

		log := log.With("op", op)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", slog.String("error", err.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error(err.Error()))
			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", slog.String("error", err.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error(err.Error()))
			return
		}

		// ToDo: add business logic

		txHash := "hash"

		log.Info("created auction", slog.String("hash", txHash))
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, Response{
			Response: resp.OK(),
			TXHash:   txHash,
		})
	}
}
