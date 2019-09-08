package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Requisitions struct {
}

// @Title GetUserRequisitions
// @Description Gets the current user's requisitions.
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [get]
func (req *Requisitions) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusNotImplemented, nil
	/*filter := funds.Requisition{}
	var container []*funds.Requisition*/
	//err := funds.Get .Read(&filter, &container)

	//req.Serve(err, container)
}

func (req *Requisitions) Search(ctx context.Requester) (int, interface{}) {
	return http.StatusNotImplemented, nil
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
func (req *Requisitions) View(ctx context.Requester) (int, interface{}) {
	return http.StatusNotImplemented, nil
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
func (req *Requisitions) Create(ctx context.Requester) (int, interface{}) {
	return http.StatusNotImplemented, nil
	/*var requisition funds.Requisition
	json.Unmarshal(req.Ctx.Input.RequestBody, &requisition)

	_, err := funds.Ctx.Requisition.Create(&requisition)

	req.Serve(err, "Requisition has been created.")*/
}

// @Title UpdateRequisition
// @Description Update requisition to confirm delivery of goods or services
// @Param	body		body 	funds.Requisition	true		"requisition entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *Requisitions) Update(ctx context.Requester) (int, interface{}) {
	return http.StatusNotImplemented, nil
	/*var requisition funds.Requisition
	json.Unmarshal(req.Ctx.Input.RequestBody, &requisition)

	err := funds.Ctx.Requisition.Update(&requisition)

	req.Serve(err, "Requisition has been updated.")*/
}
