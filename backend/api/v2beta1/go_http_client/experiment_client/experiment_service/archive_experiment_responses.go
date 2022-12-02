// Code generated by go-swagger; DO NOT EDIT.

package experiment_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ArchiveExperimentReader is a Reader for the ArchiveExperiment structure.
type ArchiveExperimentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveExperimentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewArchiveExperimentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewArchiveExperimentOK creates a ArchiveExperimentOK with default headers values
func NewArchiveExperimentOK() *ArchiveExperimentOK {
	return &ArchiveExperimentOK{}
}

/*ArchiveExperimentOK handles this case with default header values.

A successful response.
*/
type ArchiveExperimentOK struct {
	Payload interface{}
}

func (o *ArchiveExperimentOK) Error() string {
	return fmt.Sprintf("[POST /apis/v2beta1/experiments/{experiment_id}:archive][%d] archiveExperimentOK  %+v", 200, o.Payload)
}

func (o *ArchiveExperimentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}