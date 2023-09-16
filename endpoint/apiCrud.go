package endpoint

import (
	"github.com/biangacila/bialib/dbcassandra"
	"github.com/biangacila/bialib/domain"
	"net/http"
)

var sessionCass = dbcassandra.CreateConnectionWithAuth("voip.easipath.com")

func WsCrudAccounts(w http.ResponseWriter, r *http.Request) {
	WsCrudRequest(w, r, "Accounts", true, "ACC", 1002, "Code", []domain.Accounts{}, sessionCass)
}
