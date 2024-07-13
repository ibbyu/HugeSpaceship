package users

import (
	"github.com/HugeSpaceship/HugeSpaceship/internal/db"
	"github.com/HugeSpaceship/HugeSpaceship/internal/model/auth"
	"github.com/HugeSpaceship/HugeSpaceship/internal/model/lbp_xml"
	"github.com/HugeSpaceship/HugeSpaceship/pkg/utils"
	"github.com/rs/zerolog/log"
	"net/http"
)

func UpdateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := db.GetRequestConnection(r)
		if err != nil {
			panic(err)
		}

		session := utils.GetContextValue[auth.Session](r.Context(), "session")

		userUpdate, err := utils.XMLUnmarshal[lbp_xml.UpdateUser](r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		planetUpdate, er2 := utils.XMLUnmarshal[lbp_xml.PlanetUpdate](r)
		if er2 != nil {
			log.Debug().Err(er2).Msg("no bueno")
		}
		if planetUpdate.Planets != "" || planetUpdate.CCPlanet != "" {
			err := db.UpdatePlanet(conn, session.UserID, planetUpdate, session.Game)
			if err != nil {
				utils.HttpLog(w, http.StatusBadRequest, "failed to update user")
				return
			}
		}

		err = db.UpdateUser(conn, session.UserID, userUpdate)
		if err != nil {
			utils.HttpLog(w, http.StatusBadRequest, "failed to update user")
			return
		}
	}
}
