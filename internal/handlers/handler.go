package handlers

import (
	"TestTask/internal/models"
	"TestTask/internal/repositories"
	"TestTask/pkg/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type handler struct {
	repo repositories.RateRepo
}

func New(repo repositories.RateRepo) handler {
	return handler{repo: repo}
}

type GetPairs struct {
	Pairs []string `json:"pairs"`
}

func (h *handler) GETPair(w http.ResponseWriter, r *http.Request) {
	pairsParam := r.URL.Query().Get("pairs")
	pairs := strings.Split(pairsParam, ",")
	if len(pairs) == 0 {
		utils.RespondWithJSON(w, http.StatusBadRequest, models.Response{Error: fmt.Sprint("Empty parameters")})
		return
	}
	message := make([]string, 0, len(pairs))
	for _, t := range pairs {
		if t == "" {
			continue
		}
		pair := strings.Split(t, "-")
		if rate, ok := h.repo.GetRate(fmt.Sprintf("%s%s", pair[0], pair[1])); !ok {
			if rate, ok = h.repo.GetRate(fmt.Sprintf("%s%s", pair[1], pair[0])); !ok {
				utils.RespondWithJSON(w, http.StatusNotFound, models.Response{Error: fmt.Sprintf("Can`t found pair %s", t)})
				return
			} else {
				message = append(message, fmt.Sprintf("%s:%s", t, rate.Price))
			}
		} else {
			message = append(message, fmt.Sprintf("%s:%s", t, rate.Price))
		}
	}
	utils.RespondWithJSON(w, http.StatusOK, models.Response{Message: strings.Join(message, ",")})
}

func (h *handler) POSTPair(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.Response{Error: err.Error()})
		return
	}

	var pairs GetPairs
	err = json.Unmarshal(body, &pairs)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.Response{Error: err.Error()})
		return
	}

	message := make([]string, 0, len(pairs.Pairs))
	for _, t := range pairs.Pairs {
		pair := strings.Split(t, "-")
		if rate, ok := h.repo.GetRate(fmt.Sprintf("%s%s", pair[0], pair[1])); !ok {
			if rate, ok = h.repo.GetRate(fmt.Sprintf("%s%s", pair[1], pair[0])); !ok {
				utils.RespondWithJSON(w, http.StatusNotFound, models.Response{Error: fmt.Sprintf("Can`t found pair %s", t)})
				return
			} else {
				message = append(message, fmt.Sprintf("%s:%s", t, rate.Price))
			}
		} else {
			message = append(message, fmt.Sprintf("%s:%s", t, rate.Price))
		}
	}
	utils.RespondWithJSON(w, http.StatusOK, models.Response{Message: strings.Join(message, ",")})
}
