// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"sync"

	"github.com/ONSdigital/dp-content-api/service"
)

// Ensure, that HTTPServerMock does implement service.HTTPServer.
// If this is not the case, regenerate this file with moq.
var _ service.HTTPServer = &HTTPServerMock{}

// HTTPServerMock is a mock implementation of service.HTTPServer.
//
//     func TestSomethingThatUsesHTTPServer(t *testing.T) {
//
//         // make and configure a mocked service.HTTPServer
//         mockedHTTPServer := &HTTPServerMock{
//             ListenAndServeFunc: func() error {
// 	               panic("mock out the ListenAndServe method")
//             },
//             ShutdownFunc: func(ctx context.Context) error {
// 	               panic("mock out the Shutdown method")
//             },
//         }
//
//         // use mockedHTTPServer in code that requires service.HTTPServer
//         // and then make assertions.
//
//     }
type HTTPServerMock struct {
	// ListenAndServeFunc mocks the ListenAndServe method.
	ListenAndServeFunc func() error

	// ShutdownFunc mocks the Shutdown method.
	ShutdownFunc func(ctx context.Context) error

	// calls tracks calls to the methods.
	calls struct {
		// ListenAndServe holds details about calls to the ListenAndServe method.
		ListenAndServe []struct {
		}
		// Shutdown holds details about calls to the Shutdown method.
		Shutdown []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockListenAndServe sync.RWMutex
	lockShutdown       sync.RWMutex
}

// ListenAndServe calls ListenAndServeFunc.
func (mock *HTTPServerMock) ListenAndServe() error {
	if mock.ListenAndServeFunc == nil {
		panic("HTTPServerMock.ListenAndServeFunc: method is nil but HTTPServer.ListenAndServe was just called")
	}
	callInfo := struct {
	}{}
	mock.lockListenAndServe.Lock()
	mock.calls.ListenAndServe = append(mock.calls.ListenAndServe, callInfo)
	mock.lockListenAndServe.Unlock()
	return mock.ListenAndServeFunc()
}

// ListenAndServeCalls gets all the calls that were made to ListenAndServe.
// Check the length with:
//     len(mockedHTTPServer.ListenAndServeCalls())
func (mock *HTTPServerMock) ListenAndServeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockListenAndServe.RLock()
	calls = mock.calls.ListenAndServe
	mock.lockListenAndServe.RUnlock()
	return calls
}

// Shutdown calls ShutdownFunc.
func (mock *HTTPServerMock) Shutdown(ctx context.Context) error {
	if mock.ShutdownFunc == nil {
		panic("HTTPServerMock.ShutdownFunc: method is nil but HTTPServer.Shutdown was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockShutdown.Lock()
	mock.calls.Shutdown = append(mock.calls.Shutdown, callInfo)
	mock.lockShutdown.Unlock()
	return mock.ShutdownFunc(ctx)
}

// ShutdownCalls gets all the calls that were made to Shutdown.
// Check the length with:
//     len(mockedHTTPServer.ShutdownCalls())
func (mock *HTTPServerMock) ShutdownCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockShutdown.RLock()
	calls = mock.calls.Shutdown
	mock.lockShutdown.RUnlock()
	return calls
}
