/*
 * XXX API
 *
 * XXXについてのAPI
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DefaultErrorResponse struct {

	Errors []DefaultErrorResponseErrorsInner `json:"errors"`
}

// AssertDefaultErrorResponseRequired checks if the required fields are not zero-ed
func AssertDefaultErrorResponseRequired(obj DefaultErrorResponse) error {
	elements := map[string]interface{}{
		"errors": obj.Errors,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Errors {
		if err := AssertDefaultErrorResponseErrorsInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseDefaultErrorResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DefaultErrorResponse (e.g. [][]DefaultErrorResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDefaultErrorResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDefaultErrorResponse, ok := obj.(DefaultErrorResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDefaultErrorResponseRequired(aDefaultErrorResponse)
	})
}
