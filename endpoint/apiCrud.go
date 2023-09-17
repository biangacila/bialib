package endpoint

import (
	"biacibenga.com/lib/go/bialib/dbcassandra"
	"biacibenga.com/lib/go/bialib/domain"
	"net/http"
)

var sessionCass = dbcassandra.CreateConnectionWithAuth("voip.easipath.com")

func WsCrudAccounts(w http.ResponseWriter, r *http.Request) {
	WsCrudRequest(w, r, "Accounts", true, "ACC", 1002, "Code", []domain.Accounts{}, sessionCass)
}
