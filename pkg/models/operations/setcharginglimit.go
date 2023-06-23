// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"smartcar/pkg/models/shared"
)

type SetChargingLimitRequest struct {
	ChargeLimit *shared.ChargeLimit `request:"mediaType=application/json"`
	VehicleID   string              `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type SetChargingLimitResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// return Success Response
	SuccessResponse *shared.SuccessResponse
}