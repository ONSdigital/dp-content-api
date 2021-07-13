// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/ONSdigital/dp-content-api/api"
)

// Ensure, that ContentStoreMock does implement api.ContentStore.
// If this is not the case, regenerate this file with moq.
var _ api.ContentStore = &ContentStoreMock{}

// ContentStoreMock is a mock implementation of api.ContentStore.
//
//     func TestSomethingThatUsesContentStore(t *testing.T) {
//
//         // make and configure a mocked api.ContentStore
//         mockedContentStore := &ContentStoreMock{
//         }
//
//         // use mockedContentStore in code that requires api.ContentStore
//         // and then make assertions.
//
//     }
type ContentStoreMock struct {
	// calls tracks calls to the methods.
	calls struct {
	}
}
