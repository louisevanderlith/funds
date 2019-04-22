package requisitionstatus

import (
	"strings"
)

type Enum = int

const (
	Open Enum = iota
	Pending
	Delivered
	Paid
	Error
	Cancelled
)

var requisitionStatuses = [...]string{
	"Open",
	"Pending",
	"Delivered",
	"Paid",
	"Error",
	"Cancelled"}

func GetRequisitionStatus(name string) Enum {
	var result Enum

	for k, v := range requisitionStatuses {
		if strings.ToUpper(name) == strings.ToUpper(v) {
			result = Enum(k)
			break
		}
	}

	return result
}
