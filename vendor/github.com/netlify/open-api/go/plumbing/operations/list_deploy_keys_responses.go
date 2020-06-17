// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/netlify/open-api/go/models"
)

// ListDeployKeysReader is a Reader for the ListDeployKeys structure.
type ListDeployKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDeployKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDeployKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDeployKeysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDeployKeysOK creates a ListDeployKeysOK with default headers values
func NewListDeployKeysOK() *ListDeployKeysOK {
	return &ListDeployKeysOK{}
}

/*ListDeployKeysOK handles this case with default header values.

OK
*/
type ListDeployKeysOK struct {
	Payload []*models.DeployKey
}

func (o *ListDeployKeysOK) Error() string {
	return fmt.Sprintf("[GET /deploy_keys][%d] listDeployKeysOK  %+v", 200, o.Payload)
}

func (o *ListDeployKeysOK) GetPayload() []*models.DeployKey {
	return o.Payload
}

func (o *ListDeployKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDeployKeysDefault creates a ListDeployKeysDefault with default headers values
func NewListDeployKeysDefault(code int) *ListDeployKeysDefault {
	return &ListDeployKeysDefault{
		_statusCode: code,
	}
}

/*ListDeployKeysDefault handles this case with default header values.

error
*/
type ListDeployKeysDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list deploy keys default response
func (o *ListDeployKeysDefault) Code() int {
	return o._statusCode
}

func (o *ListDeployKeysDefault) Error() string {
	return fmt.Sprintf("[GET /deploy_keys][%d] listDeployKeys default  %+v", o._statusCode, o.Payload)
}

func (o *ListDeployKeysDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDeployKeysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
