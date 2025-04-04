package resources

import (
	"encoding/xml"
	"github.com/HugeSpaceship/HugeSpaceship/internal/model/lbp_xml"
	"github.com/HugeSpaceship/HugeSpaceship/internal/resources"
	"github.com/HugeSpaceship/HugeSpaceship/internal/utils"
	"io"
	"log/slog"
	"net/http"
)

func ShowNotUploadedHandler(resMan *resources.ResourceManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := lbp_xml.Resources{}
		xmlBytes, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		err = xml.Unmarshal(xmlBytes, &res)
		if err != nil {
			utils.HttpLog(w, http.StatusBadRequest, "Failed to unmarshal XML")
			return
		}

		// This checks to see if the resources already exist in the DB
		resourcesToUpload, err := resMan.HasResources(r.Context(), res.Resources)
		if err != nil {
			utils.HttpLog(w, http.StatusInternalServerError, "Failed to check resources")
			slog.Error("Failed to check resources", slog.Any("error", err))
			return
		}

		err = utils.XMLMarshal(w, &resourcesToUpload)
		if err != nil {
			utils.HttpLog(w, http.StatusInternalServerError, "Failed to marshal XML")
			return
		}
	}
}
