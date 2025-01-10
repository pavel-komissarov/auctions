package makeBet

import (
	resp "auctions/internal/server/handlers"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Bet     int    `json:"bet" validate:"required"`
	Account string `json:"account" validate:"required"`
}

func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.make.New"

		log := log.With("op", op)

		auctionID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("failed to decode request id", slog.String("error", err.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resp.Error(err.Error()))
			return
		}

		log.Info("auction ID received", slog.Int("auctionID", auctionID))

		var req Request

		err = render.DecodeJSON(r.Body, &req)
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

		log.Info("bet placed", slog.Int("auctionID", auctionID), slog.Int("Bet", req.Bet))
		render.Status(r, http.StatusOK)
		render.JSON(w, r, resp.OK())
	}
}
