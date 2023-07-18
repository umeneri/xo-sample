/*
 * XXX API
 *
 * XXXについてのAPI
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LoginBody struct {

	Id string `json:"id,omitempty"`
}

// AssertLoginBodyRequired checks if the required fields are not zero-ed
func AssertLoginBodyRequired(obj LoginBody) error {
	return nil
}

// AssertRecurseLoginBodyRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of LoginBody (e.g. [][]LoginBody), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseLoginBodyRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aLoginBody, ok := obj.(LoginBody)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertLoginBodyRequired(aLoginBody)
	})
}
