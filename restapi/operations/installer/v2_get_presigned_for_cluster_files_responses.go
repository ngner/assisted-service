// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2GetPresignedForClusterFilesOKCode is the HTTP code returned for type V2GetPresignedForClusterFilesOK
const V2GetPresignedForClusterFilesOKCode int = 200

/*V2GetPresignedForClusterFilesOK Success.

swagger:response v2GetPresignedForClusterFilesOK
*/
type V2GetPresignedForClusterFilesOK struct {

	/*
	  In: Body
	*/
	Payload *models.Presigned `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterFilesOK creates V2GetPresignedForClusterFilesOK with default headers values
func NewV2GetPresignedForClusterFilesOK() *V2GetPresignedForClusterFilesOK {

	return &V2GetPresignedForClusterFilesOK{}
}

// WithPayload adds the payload to the v2 get presigned for cluster files o k response
func (o *V2GetPresignedForClusterFilesOK) WithPayload(payload *models.Presigned) *V2GetPresignedForClusterFilesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster files o k response
func (o *V2GetPresignedForClusterFilesOK) SetPayload(payload *models.Presigned) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterFilesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterFilesBadRequestCode is the HTTP code returned for type V2GetPresignedForClusterFilesBadRequest
const V2GetPresignedForClusterFilesBadRequestCode int = 400

/*V2GetPresignedForClusterFilesBadRequest Error.

swagger:response v2GetPresignedForClusterFilesBadRequest
*/
type V2GetPresignedForClusterFilesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterFilesBadRequest creates V2GetPresignedForClusterFilesBadRequest with default headers values
func NewV2GetPresignedForClusterFilesBadRequest() *V2GetPresignedForClusterFilesBadRequest {

	return &V2GetPresignedForClusterFilesBadRequest{}
}

// WithPayload adds the payload to the v2 get presigned for cluster files bad request response
func (o *V2GetPresignedForClusterFilesBadRequest) WithPayload(payload *models.Error) *V2GetPresignedForClusterFilesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster files bad request response
func (o *V2GetPresignedForClusterFilesBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterFilesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterFilesUnauthorizedCode is the HTTP code returned for type V2GetPresignedForClusterFilesUnauthorized
const V2GetPresignedForClusterFilesUnauthorizedCode int = 401

/*V2GetPresignedForClusterFilesUnauthorized Unauthorized.

swagger:response v2GetPresignedForClusterFilesUnauthorized
*/
type V2GetPresignedForClusterFilesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterFilesUnauthorized creates V2GetPresignedForClusterFilesUnauthorized with default headers values
func NewV2GetPresignedForClusterFilesUnauthorized() *V2GetPresignedForClusterFilesUnauthorized {

	return &V2GetPresignedForClusterFilesUnauthorized{}
}

// WithPayload adds the payload to the v2 get presigned for cluster files unauthorized response
func (o *V2GetPresignedForClusterFilesUnauthorized) WithPayload(payload *models.InfraError) *V2GetPresignedForClusterFilesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster files unauthorized response
func (o *V2GetPresignedForClusterFilesUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterFilesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterFilesForbiddenCode is the HTTP code returned for type V2GetPresignedForClusterFilesForbidden
const V2GetPresignedForClusterFilesForbiddenCode int = 403

/*V2GetPresignedForClusterFilesForbidden Forbidden.

swagger:response v2GetPresignedForClusterFilesForbidden
*/
type V2GetPresignedForClusterFilesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterFilesForbidden creates V2GetPresignedForClusterFilesForbidden with default headers values
func NewV2GetPresignedForClusterFilesForbidden() *V2GetPresignedForClusterFilesForbidden {

	return &V2GetPresignedForClusterFilesForbidden{}
}

// WithPayload adds the payload to the v2 get presigned for cluster files forbidden response
func (o *V2GetPresignedForClusterFilesForbidden) WithPayload(payload *models.InfraError) *V2GetPresignedForClusterFilesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster files forbidden response
func (o *V2GetPresignedForClusterFilesForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterFilesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterFilesNotFoundCode is the HTTP code returned for type V2GetPresignedForClusterFilesNotFound
const V2GetPresignedForClusterFilesNotFoundCode int = 404

/*V2GetPresignedForClusterFilesNotFound Error.

swagger:response v2GetPresignedForClusterFilesNotFound
*/
type V2GetPresignedForClusterFilesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterFilesNotFound creates V2GetPresignedForClusterFilesNotFound with default headers values
func NewV2GetPresignedForClusterFilesNotFound() *V2GetPresignedForClusterFilesNotFound {

	return &V2GetPresignedForClusterFilesNotFound{}
}

// WithPayload adds the payload to the v2 get presigned for cluster files not found response
func (o *V2GetPresignedForClusterFilesNotFound) WithPayload(payload *models.Error) *V2GetPresignedForClusterFilesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster files not found response
func (o *V2GetPresignedForClusterFilesNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterFilesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterFilesMethodNotAllowedCode is the HTTP code returned for type V2GetPresignedForClusterFilesMethodNotAllowed
const V2GetPresignedForClusterFilesMethodNotAllowedCode int = 405

/*V2GetPresignedForClusterFilesMethodNotAllowed Method Not Allowed.

swagger:response v2GetPresignedForClusterFilesMethodNotAllowed
*/
type V2GetPresignedForClusterFilesMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterFilesMethodNotAllowed creates V2GetPresignedForClusterFilesMethodNotAllowed with default headers values
func NewV2GetPresignedForClusterFilesMethodNotAllowed() *V2GetPresignedForClusterFilesMethodNotAllowed {

	return &V2GetPresignedForClusterFilesMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 get presigned for cluster files method not allowed response
func (o *V2GetPresignedForClusterFilesMethodNotAllowed) WithPayload(payload *models.Error) *V2GetPresignedForClusterFilesMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster files method not allowed response
func (o *V2GetPresignedForClusterFilesMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterFilesMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterFilesConflictCode is the HTTP code returned for type V2GetPresignedForClusterFilesConflict
const V2GetPresignedForClusterFilesConflictCode int = 409

/*V2GetPresignedForClusterFilesConflict Error.

swagger:response v2GetPresignedForClusterFilesConflict
*/
type V2GetPresignedForClusterFilesConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterFilesConflict creates V2GetPresignedForClusterFilesConflict with default headers values
func NewV2GetPresignedForClusterFilesConflict() *V2GetPresignedForClusterFilesConflict {

	return &V2GetPresignedForClusterFilesConflict{}
}

// WithPayload adds the payload to the v2 get presigned for cluster files conflict response
func (o *V2GetPresignedForClusterFilesConflict) WithPayload(payload *models.Error) *V2GetPresignedForClusterFilesConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster files conflict response
func (o *V2GetPresignedForClusterFilesConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterFilesConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetPresignedForClusterFilesInternalServerErrorCode is the HTTP code returned for type V2GetPresignedForClusterFilesInternalServerError
const V2GetPresignedForClusterFilesInternalServerErrorCode int = 500

/*V2GetPresignedForClusterFilesInternalServerError Error.

swagger:response v2GetPresignedForClusterFilesInternalServerError
*/
type V2GetPresignedForClusterFilesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetPresignedForClusterFilesInternalServerError creates V2GetPresignedForClusterFilesInternalServerError with default headers values
func NewV2GetPresignedForClusterFilesInternalServerError() *V2GetPresignedForClusterFilesInternalServerError {

	return &V2GetPresignedForClusterFilesInternalServerError{}
}

// WithPayload adds the payload to the v2 get presigned for cluster files internal server error response
func (o *V2GetPresignedForClusterFilesInternalServerError) WithPayload(payload *models.Error) *V2GetPresignedForClusterFilesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get presigned for cluster files internal server error response
func (o *V2GetPresignedForClusterFilesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetPresignedForClusterFilesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}