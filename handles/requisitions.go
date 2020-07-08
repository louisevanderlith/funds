package handles

import (
	"net/http"
)

// @Title GetUserRequisitions
// @Description Gets the current user's requisitions.
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [get]
func GetRequisitions (w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotImplemented)
	/*filter := funds.Requisition{}
	var container []*funds.Requisition*/
	//err := funds.Get .Read(&filter, &container)

	//req.Serve(err, container)
}

func SearchRequisitions(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotImplemented)
	/*filter := funds.Requisition{}
	var container []*funds.Requisition*/
	//err := funds.Get .Read(&filter, &container)

	//req.Serve(err, container)
}

// @Title GetRequisitionDetail
// @Description Gets all details (including transactions) related to a requisition.
// @Param	requisitionID			path	string 	true		"requisition's ID"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router /:requisitionID [get]
func ViewRequisition(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotImplemented)
	/*var result db.IRecord

	reqID, err := strconv.ParseInt(req.Ctx.Input.Param(":requisitionID"), 10, 64)

	if err == nil {
		filter := funds.Requisition{}
		filter.Id = reqID

		result, err = funds.Ctx.Requisition.ReadOne(&filter)
	}

	req.Serve(err, result)*/
}

// @Title CreateRequisition
// @Description Create requisition on good or services checkout
// @Param	body		body 	funds.Requisition	true		"requisition entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func CreateRequisition(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotImplemented)
	/*var requisition funds.Requisition
	json.Unmarshal(req.Ctx.Input.RequestBody, &requisition)

	_, err := funds.Ctx.Requisition.Create(&requisition)

	req.Serve(err, "Requisition has been created.")*/
}
