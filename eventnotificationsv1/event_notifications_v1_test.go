/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package eventnotificationsv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`EventNotificationsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(eventNotificationsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(eventNotificationsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
				URL: "https://eventnotificationsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(eventNotificationsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EVENT_NOTIFICATIONS_URL":       "https://eventnotificationsv1/api",
				"EVENT_NOTIFICATIONS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1UsingExternalConfig(&eventnotificationsv1.EventNotificationsV1Options{})
				Expect(eventNotificationsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := eventNotificationsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != eventNotificationsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(eventNotificationsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(eventNotificationsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1UsingExternalConfig(&eventnotificationsv1.EventNotificationsV1Options{
					URL: "https://testService/api",
				})
				Expect(eventNotificationsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := eventNotificationsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != eventNotificationsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(eventNotificationsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(eventNotificationsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1UsingExternalConfig(&eventnotificationsv1.EventNotificationsV1Options{})
				err := eventNotificationsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := eventNotificationsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != eventNotificationsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(eventNotificationsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(eventNotificationsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EVENT_NOTIFICATIONS_URL":       "https://eventnotificationsv1/api",
				"EVENT_NOTIFICATIONS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1UsingExternalConfig(&eventnotificationsv1.EventNotificationsV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(eventNotificationsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EVENT_NOTIFICATIONS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1UsingExternalConfig(&eventnotificationsv1.EventNotificationsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(eventNotificationsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = eventnotificationsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListSources(listSourcesOptions *ListSourcesOptions) - Operation response error`, func() {
		listSourcesPath := "/v1/instances/testString/sources"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSourcesPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSources with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListSourcesOptions model
				listSourcesOptionsModel := new(eventnotificationsv1.ListSourcesOptions)
				listSourcesOptionsModel.InstanceID = core.StringPtr("testString")
				listSourcesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listSourcesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSourcesOptionsModel.Search = core.StringPtr("testString")
				listSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.ListSources(listSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.ListSources(listSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSources(listSourcesOptions *ListSourcesOptions)`, func() {
		listSourcesPath := "/v1/instances/testString/sources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSourcesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "sources": [{"id": "ID", "name": "Name", "description": "Description", "type": "Type", "enabled": false, "updated_at": "2019-01-01T12:00:00.000Z", "topic_count": 0}]}`)
				}))
			})
			It(`Invoke ListSources successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the ListSourcesOptions model
				listSourcesOptionsModel := new(eventnotificationsv1.ListSourcesOptions)
				listSourcesOptionsModel.InstanceID = core.StringPtr("testString")
				listSourcesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listSourcesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSourcesOptionsModel.Search = core.StringPtr("testString")
				listSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.ListSourcesWithContext(ctx, listSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.ListSources(listSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.ListSourcesWithContext(ctx, listSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSourcesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "sources": [{"id": "ID", "name": "Name", "description": "Description", "type": "Type", "enabled": false, "updated_at": "2019-01-01T12:00:00.000Z", "topic_count": 0}]}`)
				}))
			})
			It(`Invoke ListSources successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.ListSources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSourcesOptions model
				listSourcesOptionsModel := new(eventnotificationsv1.ListSourcesOptions)
				listSourcesOptionsModel.InstanceID = core.StringPtr("testString")
				listSourcesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listSourcesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSourcesOptionsModel.Search = core.StringPtr("testString")
				listSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.ListSources(listSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSources with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListSourcesOptions model
				listSourcesOptionsModel := new(eventnotificationsv1.ListSourcesOptions)
				listSourcesOptionsModel.InstanceID = core.StringPtr("testString")
				listSourcesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listSourcesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSourcesOptionsModel.Search = core.StringPtr("testString")
				listSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.ListSources(listSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListSourcesOptions model with no property values
				listSourcesOptionsModelNew := new(eventnotificationsv1.ListSourcesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.ListSources(listSourcesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListSources successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListSourcesOptions model
				listSourcesOptionsModel := new(eventnotificationsv1.ListSourcesOptions)
				listSourcesOptionsModel.InstanceID = core.StringPtr("testString")
				listSourcesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listSourcesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSourcesOptionsModel.Search = core.StringPtr("testString")
				listSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.ListSources(listSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSource(getSourceOptions *GetSourceOptions) - Operation response error`, func() {
		getSourcePath := "/v1/instances/testString/sources/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSourcePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSource with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetSourceOptions model
				getSourceOptionsModel := new(eventnotificationsv1.GetSourceOptions)
				getSourceOptionsModel.InstanceID = core.StringPtr("testString")
				getSourceOptionsModel.ID = core.StringPtr("testString")
				getSourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.GetSource(getSourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.GetSource(getSourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSource(getSourceOptions *GetSourceOptions)`, func() {
		getSourcePath := "/v1/instances/testString/sources/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSourcePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "enabled": false, "type": "Type", "updated_at": "2019-01-01T12:00:00.000Z", "topic_count": 10, "topic_names": ["TopicNames"]}`)
				}))
			})
			It(`Invoke GetSource successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the GetSourceOptions model
				getSourceOptionsModel := new(eventnotificationsv1.GetSourceOptions)
				getSourceOptionsModel.InstanceID = core.StringPtr("testString")
				getSourceOptionsModel.ID = core.StringPtr("testString")
				getSourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.GetSourceWithContext(ctx, getSourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.GetSource(getSourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.GetSourceWithContext(ctx, getSourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSourcePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "enabled": false, "type": "Type", "updated_at": "2019-01-01T12:00:00.000Z", "topic_count": 10, "topic_names": ["TopicNames"]}`)
				}))
			})
			It(`Invoke GetSource successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.GetSource(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSourceOptions model
				getSourceOptionsModel := new(eventnotificationsv1.GetSourceOptions)
				getSourceOptionsModel.InstanceID = core.StringPtr("testString")
				getSourceOptionsModel.ID = core.StringPtr("testString")
				getSourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.GetSource(getSourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSource with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetSourceOptions model
				getSourceOptionsModel := new(eventnotificationsv1.GetSourceOptions)
				getSourceOptionsModel.InstanceID = core.StringPtr("testString")
				getSourceOptionsModel.ID = core.StringPtr("testString")
				getSourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.GetSource(getSourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSourceOptions model with no property values
				getSourceOptionsModelNew := new(eventnotificationsv1.GetSourceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.GetSource(getSourceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSource successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetSourceOptions model
				getSourceOptionsModel := new(eventnotificationsv1.GetSourceOptions)
				getSourceOptionsModel.InstanceID = core.StringPtr("testString")
				getSourceOptionsModel.ID = core.StringPtr("testString")
				getSourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.GetSource(getSourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTopic(createTopicOptions *CreateTopicOptions) - Operation response error`, func() {
		createTopicPath := "/v1/instances/testString/topics"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTopicPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTopic with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the Rules model
				rulesModel := new(eventnotificationsv1.Rules)
				rulesModel.Enabled = core.BoolPtr(true)
				rulesModel.EventTypeFilter = core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")
				rulesModel.NotificationFilter = core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")

				// Construct an instance of the TopicUpdateSourcesItem model
				topicUpdateSourcesItemModel := new(eventnotificationsv1.TopicUpdateSourcesItem)
				topicUpdateSourcesItemModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				topicUpdateSourcesItemModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the CreateTopicOptions model
				createTopicOptionsModel := new(eventnotificationsv1.CreateTopicOptions)
				createTopicOptionsModel.InstanceID = core.StringPtr("testString")
				createTopicOptionsModel.Name = core.StringPtr("testString")
				createTopicOptionsModel.Description = core.StringPtr("testString")
				createTopicOptionsModel.Sources = []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel}
				createTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.CreateTopic(createTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.CreateTopic(createTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTopic(createTopicOptions *CreateTopicOptions)`, func() {
		createTopicPath := "/v1/instances/testString/topics"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTopicPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "created_at": "CreatedAt"}`)
				}))
			})
			It(`Invoke CreateTopic successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the Rules model
				rulesModel := new(eventnotificationsv1.Rules)
				rulesModel.Enabled = core.BoolPtr(true)
				rulesModel.EventTypeFilter = core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")
				rulesModel.NotificationFilter = core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")

				// Construct an instance of the TopicUpdateSourcesItem model
				topicUpdateSourcesItemModel := new(eventnotificationsv1.TopicUpdateSourcesItem)
				topicUpdateSourcesItemModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				topicUpdateSourcesItemModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the CreateTopicOptions model
				createTopicOptionsModel := new(eventnotificationsv1.CreateTopicOptions)
				createTopicOptionsModel.InstanceID = core.StringPtr("testString")
				createTopicOptionsModel.Name = core.StringPtr("testString")
				createTopicOptionsModel.Description = core.StringPtr("testString")
				createTopicOptionsModel.Sources = []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel}
				createTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.CreateTopicWithContext(ctx, createTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.CreateTopic(createTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.CreateTopicWithContext(ctx, createTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTopicPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "created_at": "CreatedAt"}`)
				}))
			})
			It(`Invoke CreateTopic successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.CreateTopic(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Rules model
				rulesModel := new(eventnotificationsv1.Rules)
				rulesModel.Enabled = core.BoolPtr(true)
				rulesModel.EventTypeFilter = core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")
				rulesModel.NotificationFilter = core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")

				// Construct an instance of the TopicUpdateSourcesItem model
				topicUpdateSourcesItemModel := new(eventnotificationsv1.TopicUpdateSourcesItem)
				topicUpdateSourcesItemModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				topicUpdateSourcesItemModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the CreateTopicOptions model
				createTopicOptionsModel := new(eventnotificationsv1.CreateTopicOptions)
				createTopicOptionsModel.InstanceID = core.StringPtr("testString")
				createTopicOptionsModel.Name = core.StringPtr("testString")
				createTopicOptionsModel.Description = core.StringPtr("testString")
				createTopicOptionsModel.Sources = []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel}
				createTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.CreateTopic(createTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTopic with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the Rules model
				rulesModel := new(eventnotificationsv1.Rules)
				rulesModel.Enabled = core.BoolPtr(true)
				rulesModel.EventTypeFilter = core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")
				rulesModel.NotificationFilter = core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")

				// Construct an instance of the TopicUpdateSourcesItem model
				topicUpdateSourcesItemModel := new(eventnotificationsv1.TopicUpdateSourcesItem)
				topicUpdateSourcesItemModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				topicUpdateSourcesItemModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the CreateTopicOptions model
				createTopicOptionsModel := new(eventnotificationsv1.CreateTopicOptions)
				createTopicOptionsModel.InstanceID = core.StringPtr("testString")
				createTopicOptionsModel.Name = core.StringPtr("testString")
				createTopicOptionsModel.Description = core.StringPtr("testString")
				createTopicOptionsModel.Sources = []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel}
				createTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.CreateTopic(createTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTopicOptions model with no property values
				createTopicOptionsModelNew := new(eventnotificationsv1.CreateTopicOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.CreateTopic(createTopicOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateTopic successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the Rules model
				rulesModel := new(eventnotificationsv1.Rules)
				rulesModel.Enabled = core.BoolPtr(true)
				rulesModel.EventTypeFilter = core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")
				rulesModel.NotificationFilter = core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")

				// Construct an instance of the TopicUpdateSourcesItem model
				topicUpdateSourcesItemModel := new(eventnotificationsv1.TopicUpdateSourcesItem)
				topicUpdateSourcesItemModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				topicUpdateSourcesItemModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the CreateTopicOptions model
				createTopicOptionsModel := new(eventnotificationsv1.CreateTopicOptions)
				createTopicOptionsModel.InstanceID = core.StringPtr("testString")
				createTopicOptionsModel.Name = core.StringPtr("testString")
				createTopicOptionsModel.Description = core.StringPtr("testString")
				createTopicOptionsModel.Sources = []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel}
				createTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.CreateTopic(createTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTopics(listTopicsOptions *ListTopicsOptions) - Operation response error`, func() {
		listTopicsPath := "/v1/instances/testString/topics"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTopicsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTopics with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListTopicsOptions model
				listTopicsOptionsModel := new(eventnotificationsv1.ListTopicsOptions)
				listTopicsOptionsModel.InstanceID = core.StringPtr("testString")
				listTopicsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTopicsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTopicsOptionsModel.Search = core.StringPtr("testString")
				listTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.ListTopics(listTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.ListTopics(listTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTopics(listTopicsOptions *ListTopicsOptions)`, func() {
		listTopicsPath := "/v1/instances/testString/topics"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTopicsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "topics": [{"id": "ID", "name": "Name", "description": "Description", "source_count": 0, "sources_names": ["SourcesNames"], "subscription_count": 0}]}`)
				}))
			})
			It(`Invoke ListTopics successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the ListTopicsOptions model
				listTopicsOptionsModel := new(eventnotificationsv1.ListTopicsOptions)
				listTopicsOptionsModel.InstanceID = core.StringPtr("testString")
				listTopicsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTopicsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTopicsOptionsModel.Search = core.StringPtr("testString")
				listTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.ListTopicsWithContext(ctx, listTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.ListTopics(listTopicsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.ListTopicsWithContext(ctx, listTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTopicsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "topics": [{"id": "ID", "name": "Name", "description": "Description", "source_count": 0, "sources_names": ["SourcesNames"], "subscription_count": 0}]}`)
				}))
			})
			It(`Invoke ListTopics successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.ListTopics(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTopicsOptions model
				listTopicsOptionsModel := new(eventnotificationsv1.ListTopicsOptions)
				listTopicsOptionsModel.InstanceID = core.StringPtr("testString")
				listTopicsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTopicsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTopicsOptionsModel.Search = core.StringPtr("testString")
				listTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.ListTopics(listTopicsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTopics with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListTopicsOptions model
				listTopicsOptionsModel := new(eventnotificationsv1.ListTopicsOptions)
				listTopicsOptionsModel.InstanceID = core.StringPtr("testString")
				listTopicsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTopicsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTopicsOptionsModel.Search = core.StringPtr("testString")
				listTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.ListTopics(listTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTopicsOptions model with no property values
				listTopicsOptionsModelNew := new(eventnotificationsv1.ListTopicsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.ListTopics(listTopicsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListTopics successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListTopicsOptions model
				listTopicsOptionsModel := new(eventnotificationsv1.ListTopicsOptions)
				listTopicsOptionsModel.InstanceID = core.StringPtr("testString")
				listTopicsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTopicsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTopicsOptionsModel.Search = core.StringPtr("testString")
				listTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.ListTopics(listTopicsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTopic(getTopicOptions *GetTopicOptions) - Operation response error`, func() {
		getTopicPath := "/v1/instances/testString/topics/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTopicPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTopic with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetTopicOptions model
				getTopicOptionsModel := new(eventnotificationsv1.GetTopicOptions)
				getTopicOptionsModel.InstanceID = core.StringPtr("testString")
				getTopicOptionsModel.ID = core.StringPtr("testString")
				getTopicOptionsModel.Include = core.StringPtr("testString")
				getTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.GetTopic(getTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.GetTopic(getTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTopic(getTopicOptions *GetTopicOptions)`, func() {
		getTopicPath := "/v1/instances/testString/topics/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTopicPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "description": "Description", "name": "Name", "updated_at": "UpdatedAt", "source_count": 11, "sources": [{"id": "ID", "name": "Name", "rules": [{"enabled": false, "event_type_filter": "$.*", "notification_filter": "NotificationFilter", "updated_at": "UpdatedAt", "id": "ID"}]}], "subscription_count": 17, "subscriptions": [{"id": "ID", "name": "Name", "description": "Description", "destination_id": "DestinationID", "destination_name": "DestinationName", "destination_type": "sms_ibm", "topic_id": "TopicID", "topic_name": "TopicName", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke GetTopic successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the GetTopicOptions model
				getTopicOptionsModel := new(eventnotificationsv1.GetTopicOptions)
				getTopicOptionsModel.InstanceID = core.StringPtr("testString")
				getTopicOptionsModel.ID = core.StringPtr("testString")
				getTopicOptionsModel.Include = core.StringPtr("testString")
				getTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.GetTopicWithContext(ctx, getTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.GetTopic(getTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.GetTopicWithContext(ctx, getTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTopicPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "description": "Description", "name": "Name", "updated_at": "UpdatedAt", "source_count": 11, "sources": [{"id": "ID", "name": "Name", "rules": [{"enabled": false, "event_type_filter": "$.*", "notification_filter": "NotificationFilter", "updated_at": "UpdatedAt", "id": "ID"}]}], "subscription_count": 17, "subscriptions": [{"id": "ID", "name": "Name", "description": "Description", "destination_id": "DestinationID", "destination_name": "DestinationName", "destination_type": "sms_ibm", "topic_id": "TopicID", "topic_name": "TopicName", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke GetTopic successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.GetTopic(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTopicOptions model
				getTopicOptionsModel := new(eventnotificationsv1.GetTopicOptions)
				getTopicOptionsModel.InstanceID = core.StringPtr("testString")
				getTopicOptionsModel.ID = core.StringPtr("testString")
				getTopicOptionsModel.Include = core.StringPtr("testString")
				getTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.GetTopic(getTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTopic with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetTopicOptions model
				getTopicOptionsModel := new(eventnotificationsv1.GetTopicOptions)
				getTopicOptionsModel.InstanceID = core.StringPtr("testString")
				getTopicOptionsModel.ID = core.StringPtr("testString")
				getTopicOptionsModel.Include = core.StringPtr("testString")
				getTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.GetTopic(getTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTopicOptions model with no property values
				getTopicOptionsModelNew := new(eventnotificationsv1.GetTopicOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.GetTopic(getTopicOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetTopic successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetTopicOptions model
				getTopicOptionsModel := new(eventnotificationsv1.GetTopicOptions)
				getTopicOptionsModel.InstanceID = core.StringPtr("testString")
				getTopicOptionsModel.ID = core.StringPtr("testString")
				getTopicOptionsModel.Include = core.StringPtr("testString")
				getTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.GetTopic(getTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceTopic(replaceTopicOptions *ReplaceTopicOptions) - Operation response error`, func() {
		replaceTopicPath := "/v1/instances/testString/topics/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTopicPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceTopic with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the Rules model
				rulesModel := new(eventnotificationsv1.Rules)
				rulesModel.Enabled = core.BoolPtr(true)
				rulesModel.EventTypeFilter = core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")
				rulesModel.NotificationFilter = core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")

				// Construct an instance of the TopicUpdateSourcesItem model
				topicUpdateSourcesItemModel := new(eventnotificationsv1.TopicUpdateSourcesItem)
				topicUpdateSourcesItemModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				topicUpdateSourcesItemModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the ReplaceTopicOptions model
				replaceTopicOptionsModel := new(eventnotificationsv1.ReplaceTopicOptions)
				replaceTopicOptionsModel.InstanceID = core.StringPtr("testString")
				replaceTopicOptionsModel.ID = core.StringPtr("testString")
				replaceTopicOptionsModel.Name = core.StringPtr("testString")
				replaceTopicOptionsModel.Description = core.StringPtr("testString")
				replaceTopicOptionsModel.Sources = []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel}
				replaceTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.ReplaceTopic(replaceTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.ReplaceTopic(replaceTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceTopic(replaceTopicOptions *ReplaceTopicOptions)`, func() {
		replaceTopicPath := "/v1/instances/testString/topics/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTopicPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "description": "Description", "name": "Name", "updated_at": "UpdatedAt", "source_count": 11, "sources": [{"id": "ID", "name": "Name", "rules": [{"enabled": false, "event_type_filter": "$.*", "notification_filter": "NotificationFilter", "updated_at": "UpdatedAt", "id": "ID"}]}], "subscription_count": 17, "subscriptions": [{"id": "ID", "name": "Name", "description": "Description", "destination_id": "DestinationID", "destination_name": "DestinationName", "destination_type": "sms_ibm", "topic_id": "TopicID", "topic_name": "TopicName", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ReplaceTopic successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the Rules model
				rulesModel := new(eventnotificationsv1.Rules)
				rulesModel.Enabled = core.BoolPtr(true)
				rulesModel.EventTypeFilter = core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")
				rulesModel.NotificationFilter = core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")

				// Construct an instance of the TopicUpdateSourcesItem model
				topicUpdateSourcesItemModel := new(eventnotificationsv1.TopicUpdateSourcesItem)
				topicUpdateSourcesItemModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				topicUpdateSourcesItemModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the ReplaceTopicOptions model
				replaceTopicOptionsModel := new(eventnotificationsv1.ReplaceTopicOptions)
				replaceTopicOptionsModel.InstanceID = core.StringPtr("testString")
				replaceTopicOptionsModel.ID = core.StringPtr("testString")
				replaceTopicOptionsModel.Name = core.StringPtr("testString")
				replaceTopicOptionsModel.Description = core.StringPtr("testString")
				replaceTopicOptionsModel.Sources = []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel}
				replaceTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.ReplaceTopicWithContext(ctx, replaceTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.ReplaceTopic(replaceTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.ReplaceTopicWithContext(ctx, replaceTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTopicPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "description": "Description", "name": "Name", "updated_at": "UpdatedAt", "source_count": 11, "sources": [{"id": "ID", "name": "Name", "rules": [{"enabled": false, "event_type_filter": "$.*", "notification_filter": "NotificationFilter", "updated_at": "UpdatedAt", "id": "ID"}]}], "subscription_count": 17, "subscriptions": [{"id": "ID", "name": "Name", "description": "Description", "destination_id": "DestinationID", "destination_name": "DestinationName", "destination_type": "sms_ibm", "topic_id": "TopicID", "topic_name": "TopicName", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ReplaceTopic successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.ReplaceTopic(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Rules model
				rulesModel := new(eventnotificationsv1.Rules)
				rulesModel.Enabled = core.BoolPtr(true)
				rulesModel.EventTypeFilter = core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")
				rulesModel.NotificationFilter = core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")

				// Construct an instance of the TopicUpdateSourcesItem model
				topicUpdateSourcesItemModel := new(eventnotificationsv1.TopicUpdateSourcesItem)
				topicUpdateSourcesItemModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				topicUpdateSourcesItemModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the ReplaceTopicOptions model
				replaceTopicOptionsModel := new(eventnotificationsv1.ReplaceTopicOptions)
				replaceTopicOptionsModel.InstanceID = core.StringPtr("testString")
				replaceTopicOptionsModel.ID = core.StringPtr("testString")
				replaceTopicOptionsModel.Name = core.StringPtr("testString")
				replaceTopicOptionsModel.Description = core.StringPtr("testString")
				replaceTopicOptionsModel.Sources = []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel}
				replaceTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.ReplaceTopic(replaceTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceTopic with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the Rules model
				rulesModel := new(eventnotificationsv1.Rules)
				rulesModel.Enabled = core.BoolPtr(true)
				rulesModel.EventTypeFilter = core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")
				rulesModel.NotificationFilter = core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")

				// Construct an instance of the TopicUpdateSourcesItem model
				topicUpdateSourcesItemModel := new(eventnotificationsv1.TopicUpdateSourcesItem)
				topicUpdateSourcesItemModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				topicUpdateSourcesItemModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the ReplaceTopicOptions model
				replaceTopicOptionsModel := new(eventnotificationsv1.ReplaceTopicOptions)
				replaceTopicOptionsModel.InstanceID = core.StringPtr("testString")
				replaceTopicOptionsModel.ID = core.StringPtr("testString")
				replaceTopicOptionsModel.Name = core.StringPtr("testString")
				replaceTopicOptionsModel.Description = core.StringPtr("testString")
				replaceTopicOptionsModel.Sources = []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel}
				replaceTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.ReplaceTopic(replaceTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceTopicOptions model with no property values
				replaceTopicOptionsModelNew := new(eventnotificationsv1.ReplaceTopicOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.ReplaceTopic(replaceTopicOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ReplaceTopic successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the Rules model
				rulesModel := new(eventnotificationsv1.Rules)
				rulesModel.Enabled = core.BoolPtr(true)
				rulesModel.EventTypeFilter = core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")
				rulesModel.NotificationFilter = core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")

				// Construct an instance of the TopicUpdateSourcesItem model
				topicUpdateSourcesItemModel := new(eventnotificationsv1.TopicUpdateSourcesItem)
				topicUpdateSourcesItemModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				topicUpdateSourcesItemModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the ReplaceTopicOptions model
				replaceTopicOptionsModel := new(eventnotificationsv1.ReplaceTopicOptions)
				replaceTopicOptionsModel.InstanceID = core.StringPtr("testString")
				replaceTopicOptionsModel.ID = core.StringPtr("testString")
				replaceTopicOptionsModel.Name = core.StringPtr("testString")
				replaceTopicOptionsModel.Description = core.StringPtr("testString")
				replaceTopicOptionsModel.Sources = []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel}
				replaceTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.ReplaceTopic(replaceTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteTopic(deleteTopicOptions *DeleteTopicOptions)`, func() {
		deleteTopicPath := "/v1/instances/testString/topics/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTopicPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTopic successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := eventNotificationsService.DeleteTopic(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTopicOptions model
				deleteTopicOptionsModel := new(eventnotificationsv1.DeleteTopicOptions)
				deleteTopicOptionsModel.InstanceID = core.StringPtr("testString")
				deleteTopicOptionsModel.ID = core.StringPtr("testString")
				deleteTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = eventNotificationsService.DeleteTopic(deleteTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTopic with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DeleteTopicOptions model
				deleteTopicOptionsModel := new(eventnotificationsv1.DeleteTopicOptions)
				deleteTopicOptionsModel.InstanceID = core.StringPtr("testString")
				deleteTopicOptionsModel.ID = core.StringPtr("testString")
				deleteTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := eventNotificationsService.DeleteTopic(deleteTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTopicOptions model with no property values
				deleteTopicOptionsModelNew := new(eventnotificationsv1.DeleteTopicOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = eventNotificationsService.DeleteTopic(deleteTopicOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDestination(createDestinationOptions *CreateDestinationOptions) - Operation response error`, func() {
		createDestinationPath := "/v1/instances/testString/destinations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDestinationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDestination with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DestinationConfigParamsWebhookDestinationConfig model
				destinationConfigParamsModel := new(eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig)
				destinationConfigParamsModel.URL = core.StringPtr("testString")
				destinationConfigParamsModel.Verb = core.StringPtr("get")
				destinationConfigParamsModel.CustomHeaders = make(map[string]string)
				destinationConfigParamsModel.SensitiveHeaders = []string{"testString"}

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigParamsModel

				// Construct an instance of the CreateDestinationOptions model
				createDestinationOptionsModel := new(eventnotificationsv1.CreateDestinationOptions)
				createDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				createDestinationOptionsModel.Name = core.StringPtr("testString")
				createDestinationOptionsModel.Type = core.StringPtr("webhook")
				createDestinationOptionsModel.Description = core.StringPtr("testString")
				createDestinationOptionsModel.Config = destinationConfigModel
				createDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.CreateDestination(createDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.CreateDestination(createDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDestination(createDestinationOptions *CreateDestinationOptions)`, func() {
		createDestinationPath := "/v1/instances/testString/destinations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDestinationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "config": {"params": {"url": "URL", "verb": "get", "custom_headers": {"mapKey": "Inner"}, "sensitive_headers": ["SensitiveHeaders"]}}, "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateDestination successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the DestinationConfigParamsWebhookDestinationConfig model
				destinationConfigParamsModel := new(eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig)
				destinationConfigParamsModel.URL = core.StringPtr("testString")
				destinationConfigParamsModel.Verb = core.StringPtr("get")
				destinationConfigParamsModel.CustomHeaders = make(map[string]string)
				destinationConfigParamsModel.SensitiveHeaders = []string{"testString"}

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigParamsModel

				// Construct an instance of the CreateDestinationOptions model
				createDestinationOptionsModel := new(eventnotificationsv1.CreateDestinationOptions)
				createDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				createDestinationOptionsModel.Name = core.StringPtr("testString")
				createDestinationOptionsModel.Type = core.StringPtr("webhook")
				createDestinationOptionsModel.Description = core.StringPtr("testString")
				createDestinationOptionsModel.Config = destinationConfigModel
				createDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.CreateDestinationWithContext(ctx, createDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.CreateDestination(createDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.CreateDestinationWithContext(ctx, createDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDestinationPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "config": {"params": {"url": "URL", "verb": "get", "custom_headers": {"mapKey": "Inner"}, "sensitive_headers": ["SensitiveHeaders"]}}, "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateDestination successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.CreateDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DestinationConfigParamsWebhookDestinationConfig model
				destinationConfigParamsModel := new(eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig)
				destinationConfigParamsModel.URL = core.StringPtr("testString")
				destinationConfigParamsModel.Verb = core.StringPtr("get")
				destinationConfigParamsModel.CustomHeaders = make(map[string]string)
				destinationConfigParamsModel.SensitiveHeaders = []string{"testString"}

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigParamsModel

				// Construct an instance of the CreateDestinationOptions model
				createDestinationOptionsModel := new(eventnotificationsv1.CreateDestinationOptions)
				createDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				createDestinationOptionsModel.Name = core.StringPtr("testString")
				createDestinationOptionsModel.Type = core.StringPtr("webhook")
				createDestinationOptionsModel.Description = core.StringPtr("testString")
				createDestinationOptionsModel.Config = destinationConfigModel
				createDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.CreateDestination(createDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDestination with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DestinationConfigParamsWebhookDestinationConfig model
				destinationConfigParamsModel := new(eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig)
				destinationConfigParamsModel.URL = core.StringPtr("testString")
				destinationConfigParamsModel.Verb = core.StringPtr("get")
				destinationConfigParamsModel.CustomHeaders = make(map[string]string)
				destinationConfigParamsModel.SensitiveHeaders = []string{"testString"}

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigParamsModel

				// Construct an instance of the CreateDestinationOptions model
				createDestinationOptionsModel := new(eventnotificationsv1.CreateDestinationOptions)
				createDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				createDestinationOptionsModel.Name = core.StringPtr("testString")
				createDestinationOptionsModel.Type = core.StringPtr("webhook")
				createDestinationOptionsModel.Description = core.StringPtr("testString")
				createDestinationOptionsModel.Config = destinationConfigModel
				createDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.CreateDestination(createDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDestinationOptions model with no property values
				createDestinationOptionsModelNew := new(eventnotificationsv1.CreateDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.CreateDestination(createDestinationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateDestination successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DestinationConfigParamsWebhookDestinationConfig model
				destinationConfigParamsModel := new(eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig)
				destinationConfigParamsModel.URL = core.StringPtr("testString")
				destinationConfigParamsModel.Verb = core.StringPtr("get")
				destinationConfigParamsModel.CustomHeaders = make(map[string]string)
				destinationConfigParamsModel.SensitiveHeaders = []string{"testString"}

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigParamsModel

				// Construct an instance of the CreateDestinationOptions model
				createDestinationOptionsModel := new(eventnotificationsv1.CreateDestinationOptions)
				createDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				createDestinationOptionsModel.Name = core.StringPtr("testString")
				createDestinationOptionsModel.Type = core.StringPtr("webhook")
				createDestinationOptionsModel.Description = core.StringPtr("testString")
				createDestinationOptionsModel.Config = destinationConfigModel
				createDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.CreateDestination(createDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDestinations(listDestinationsOptions *ListDestinationsOptions) - Operation response error`, func() {
		listDestinationsPath := "/v1/instances/testString/destinations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDestinationsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDestinations with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListDestinationsOptions model
				listDestinationsOptionsModel := new(eventnotificationsv1.ListDestinationsOptions)
				listDestinationsOptionsModel.InstanceID = core.StringPtr("testString")
				listDestinationsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listDestinationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listDestinationsOptionsModel.Search = core.StringPtr("testString")
				listDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.ListDestinations(listDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.ListDestinations(listDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDestinations(listDestinationsOptions *ListDestinationsOptions)`, func() {
		listDestinationsPath := "/v1/instances/testString/destinations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDestinationsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "destinations": [{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "subscription_count": 17, "subscription_names": ["SubscriptionNames"], "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListDestinations successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the ListDestinationsOptions model
				listDestinationsOptionsModel := new(eventnotificationsv1.ListDestinationsOptions)
				listDestinationsOptionsModel.InstanceID = core.StringPtr("testString")
				listDestinationsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listDestinationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listDestinationsOptionsModel.Search = core.StringPtr("testString")
				listDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.ListDestinationsWithContext(ctx, listDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.ListDestinations(listDestinationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.ListDestinationsWithContext(ctx, listDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDestinationsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "destinations": [{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "subscription_count": 17, "subscription_names": ["SubscriptionNames"], "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListDestinations successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.ListDestinations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDestinationsOptions model
				listDestinationsOptionsModel := new(eventnotificationsv1.ListDestinationsOptions)
				listDestinationsOptionsModel.InstanceID = core.StringPtr("testString")
				listDestinationsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listDestinationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listDestinationsOptionsModel.Search = core.StringPtr("testString")
				listDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.ListDestinations(listDestinationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDestinations with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListDestinationsOptions model
				listDestinationsOptionsModel := new(eventnotificationsv1.ListDestinationsOptions)
				listDestinationsOptionsModel.InstanceID = core.StringPtr("testString")
				listDestinationsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listDestinationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listDestinationsOptionsModel.Search = core.StringPtr("testString")
				listDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.ListDestinations(listDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDestinationsOptions model with no property values
				listDestinationsOptionsModelNew := new(eventnotificationsv1.ListDestinationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.ListDestinations(listDestinationsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDestinations successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListDestinationsOptions model
				listDestinationsOptionsModel := new(eventnotificationsv1.ListDestinationsOptions)
				listDestinationsOptionsModel.InstanceID = core.StringPtr("testString")
				listDestinationsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listDestinationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listDestinationsOptionsModel.Search = core.StringPtr("testString")
				listDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.ListDestinations(listDestinationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDestination(getDestinationOptions *GetDestinationOptions) - Operation response error`, func() {
		getDestinationPath := "/v1/instances/testString/destinations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDestinationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDestination with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetDestinationOptions model
				getDestinationOptionsModel := new(eventnotificationsv1.GetDestinationOptions)
				getDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				getDestinationOptionsModel.ID = core.StringPtr("testString")
				getDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.GetDestination(getDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.GetDestination(getDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDestination(getDestinationOptions *GetDestinationOptions)`, func() {
		getDestinationPath := "/v1/instances/testString/destinations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDestinationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "config": {"params": {"url": "URL", "verb": "get", "custom_headers": {"mapKey": "Inner"}, "sensitive_headers": ["SensitiveHeaders"]}}, "updated_at": "2019-01-01T12:00:00.000Z", "subscription_count": 0, "subscription_names": ["SubscriptionNames"]}`)
				}))
			})
			It(`Invoke GetDestination successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the GetDestinationOptions model
				getDestinationOptionsModel := new(eventnotificationsv1.GetDestinationOptions)
				getDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				getDestinationOptionsModel.ID = core.StringPtr("testString")
				getDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.GetDestinationWithContext(ctx, getDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.GetDestination(getDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.GetDestinationWithContext(ctx, getDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDestinationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "config": {"params": {"url": "URL", "verb": "get", "custom_headers": {"mapKey": "Inner"}, "sensitive_headers": ["SensitiveHeaders"]}}, "updated_at": "2019-01-01T12:00:00.000Z", "subscription_count": 0, "subscription_names": ["SubscriptionNames"]}`)
				}))
			})
			It(`Invoke GetDestination successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.GetDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDestinationOptions model
				getDestinationOptionsModel := new(eventnotificationsv1.GetDestinationOptions)
				getDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				getDestinationOptionsModel.ID = core.StringPtr("testString")
				getDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.GetDestination(getDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDestination with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetDestinationOptions model
				getDestinationOptionsModel := new(eventnotificationsv1.GetDestinationOptions)
				getDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				getDestinationOptionsModel.ID = core.StringPtr("testString")
				getDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.GetDestination(getDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDestinationOptions model with no property values
				getDestinationOptionsModelNew := new(eventnotificationsv1.GetDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.GetDestination(getDestinationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDestination successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetDestinationOptions model
				getDestinationOptionsModel := new(eventnotificationsv1.GetDestinationOptions)
				getDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				getDestinationOptionsModel.ID = core.StringPtr("testString")
				getDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.GetDestination(getDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDestination(updateDestinationOptions *UpdateDestinationOptions) - Operation response error`, func() {
		updateDestinationPath := "/v1/instances/testString/destinations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDestinationPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDestination with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DestinationConfigParamsWebhookDestinationConfig model
				destinationConfigParamsModel := new(eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig)
				destinationConfigParamsModel.URL = core.StringPtr("testString")
				destinationConfigParamsModel.Verb = core.StringPtr("get")
				destinationConfigParamsModel.CustomHeaders = make(map[string]string)
				destinationConfigParamsModel.SensitiveHeaders = []string{"testString"}

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigParamsModel

				// Construct an instance of the UpdateDestinationOptions model
				updateDestinationOptionsModel := new(eventnotificationsv1.UpdateDestinationOptions)
				updateDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateDestinationOptionsModel.ID = core.StringPtr("testString")
				updateDestinationOptionsModel.Name = core.StringPtr("testString")
				updateDestinationOptionsModel.Description = core.StringPtr("testString")
				updateDestinationOptionsModel.Config = destinationConfigModel
				updateDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.UpdateDestination(updateDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.UpdateDestination(updateDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDestination(updateDestinationOptions *UpdateDestinationOptions)`, func() {
		updateDestinationPath := "/v1/instances/testString/destinations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDestinationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "config": {"params": {"url": "URL", "verb": "get", "custom_headers": {"mapKey": "Inner"}, "sensitive_headers": ["SensitiveHeaders"]}}, "updated_at": "2019-01-01T12:00:00.000Z", "subscription_count": 0, "subscription_names": ["SubscriptionNames"]}`)
				}))
			})
			It(`Invoke UpdateDestination successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the DestinationConfigParamsWebhookDestinationConfig model
				destinationConfigParamsModel := new(eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig)
				destinationConfigParamsModel.URL = core.StringPtr("testString")
				destinationConfigParamsModel.Verb = core.StringPtr("get")
				destinationConfigParamsModel.CustomHeaders = make(map[string]string)
				destinationConfigParamsModel.SensitiveHeaders = []string{"testString"}

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigParamsModel

				// Construct an instance of the UpdateDestinationOptions model
				updateDestinationOptionsModel := new(eventnotificationsv1.UpdateDestinationOptions)
				updateDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateDestinationOptionsModel.ID = core.StringPtr("testString")
				updateDestinationOptionsModel.Name = core.StringPtr("testString")
				updateDestinationOptionsModel.Description = core.StringPtr("testString")
				updateDestinationOptionsModel.Config = destinationConfigModel
				updateDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.UpdateDestinationWithContext(ctx, updateDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.UpdateDestination(updateDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.UpdateDestinationWithContext(ctx, updateDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDestinationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "config": {"params": {"url": "URL", "verb": "get", "custom_headers": {"mapKey": "Inner"}, "sensitive_headers": ["SensitiveHeaders"]}}, "updated_at": "2019-01-01T12:00:00.000Z", "subscription_count": 0, "subscription_names": ["SubscriptionNames"]}`)
				}))
			})
			It(`Invoke UpdateDestination successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.UpdateDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DestinationConfigParamsWebhookDestinationConfig model
				destinationConfigParamsModel := new(eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig)
				destinationConfigParamsModel.URL = core.StringPtr("testString")
				destinationConfigParamsModel.Verb = core.StringPtr("get")
				destinationConfigParamsModel.CustomHeaders = make(map[string]string)
				destinationConfigParamsModel.SensitiveHeaders = []string{"testString"}

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigParamsModel

				// Construct an instance of the UpdateDestinationOptions model
				updateDestinationOptionsModel := new(eventnotificationsv1.UpdateDestinationOptions)
				updateDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateDestinationOptionsModel.ID = core.StringPtr("testString")
				updateDestinationOptionsModel.Name = core.StringPtr("testString")
				updateDestinationOptionsModel.Description = core.StringPtr("testString")
				updateDestinationOptionsModel.Config = destinationConfigModel
				updateDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.UpdateDestination(updateDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDestination with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DestinationConfigParamsWebhookDestinationConfig model
				destinationConfigParamsModel := new(eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig)
				destinationConfigParamsModel.URL = core.StringPtr("testString")
				destinationConfigParamsModel.Verb = core.StringPtr("get")
				destinationConfigParamsModel.CustomHeaders = make(map[string]string)
				destinationConfigParamsModel.SensitiveHeaders = []string{"testString"}

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigParamsModel

				// Construct an instance of the UpdateDestinationOptions model
				updateDestinationOptionsModel := new(eventnotificationsv1.UpdateDestinationOptions)
				updateDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateDestinationOptionsModel.ID = core.StringPtr("testString")
				updateDestinationOptionsModel.Name = core.StringPtr("testString")
				updateDestinationOptionsModel.Description = core.StringPtr("testString")
				updateDestinationOptionsModel.Config = destinationConfigModel
				updateDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.UpdateDestination(updateDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDestinationOptions model with no property values
				updateDestinationOptionsModelNew := new(eventnotificationsv1.UpdateDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.UpdateDestination(updateDestinationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateDestination successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DestinationConfigParamsWebhookDestinationConfig model
				destinationConfigParamsModel := new(eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig)
				destinationConfigParamsModel.URL = core.StringPtr("testString")
				destinationConfigParamsModel.Verb = core.StringPtr("get")
				destinationConfigParamsModel.CustomHeaders = make(map[string]string)
				destinationConfigParamsModel.SensitiveHeaders = []string{"testString"}

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigParamsModel

				// Construct an instance of the UpdateDestinationOptions model
				updateDestinationOptionsModel := new(eventnotificationsv1.UpdateDestinationOptions)
				updateDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateDestinationOptionsModel.ID = core.StringPtr("testString")
				updateDestinationOptionsModel.Name = core.StringPtr("testString")
				updateDestinationOptionsModel.Description = core.StringPtr("testString")
				updateDestinationOptionsModel.Config = destinationConfigModel
				updateDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.UpdateDestination(updateDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDestination(deleteDestinationOptions *DeleteDestinationOptions)`, func() {
		deleteDestinationPath := "/v1/instances/testString/destinations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDestinationPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDestination successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := eventNotificationsService.DeleteDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDestinationOptions model
				deleteDestinationOptionsModel := new(eventnotificationsv1.DeleteDestinationOptions)
				deleteDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				deleteDestinationOptionsModel.ID = core.StringPtr("testString")
				deleteDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = eventNotificationsService.DeleteDestination(deleteDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDestination with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DeleteDestinationOptions model
				deleteDestinationOptionsModel := new(eventnotificationsv1.DeleteDestinationOptions)
				deleteDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				deleteDestinationOptionsModel.ID = core.StringPtr("testString")
				deleteDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := eventNotificationsService.DeleteDestination(deleteDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDestinationOptions model with no property values
				deleteDestinationOptionsModelNew := new(eventnotificationsv1.DeleteDestinationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = eventNotificationsService.DeleteDestination(deleteDestinationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSubscription(createSubscriptionOptions *CreateSubscriptionOptions) - Operation response error`, func() {
		createSubscriptionPath := "/v1/instances/testString/subscriptions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSubscriptionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSubscription with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the SubscriptionCreateAttributesSmsAttributes model
				subscriptionCreateAttributesModel := new(eventnotificationsv1.SubscriptionCreateAttributesSmsAttributes)
				subscriptionCreateAttributesModel.To = []string{"testString"}

				// Construct an instance of the CreateSubscriptionOptions model
				createSubscriptionOptionsModel := new(eventnotificationsv1.CreateSubscriptionOptions)
				createSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Name = core.StringPtr("testString")
				createSubscriptionOptionsModel.DestinationID = core.StringPtr("testString")
				createSubscriptionOptionsModel.TopicID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Attributes = subscriptionCreateAttributesModel
				createSubscriptionOptionsModel.Description = core.StringPtr("testString")
				createSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.CreateSubscription(createSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.CreateSubscription(createSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSubscription(createSubscriptionOptions *CreateSubscriptionOptions)`, func() {
		createSubscriptionPath := "/v1/instances/testString/subscriptions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSubscriptionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {"to": ["To"], "recipient_selection": "only_destination"}}`)
				}))
			})
			It(`Invoke CreateSubscription successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the SubscriptionCreateAttributesSmsAttributes model
				subscriptionCreateAttributesModel := new(eventnotificationsv1.SubscriptionCreateAttributesSmsAttributes)
				subscriptionCreateAttributesModel.To = []string{"testString"}

				// Construct an instance of the CreateSubscriptionOptions model
				createSubscriptionOptionsModel := new(eventnotificationsv1.CreateSubscriptionOptions)
				createSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Name = core.StringPtr("testString")
				createSubscriptionOptionsModel.DestinationID = core.StringPtr("testString")
				createSubscriptionOptionsModel.TopicID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Attributes = subscriptionCreateAttributesModel
				createSubscriptionOptionsModel.Description = core.StringPtr("testString")
				createSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.CreateSubscriptionWithContext(ctx, createSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.CreateSubscription(createSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.CreateSubscriptionWithContext(ctx, createSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSubscriptionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {"to": ["To"], "recipient_selection": "only_destination"}}`)
				}))
			})
			It(`Invoke CreateSubscription successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.CreateSubscription(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SubscriptionCreateAttributesSmsAttributes model
				subscriptionCreateAttributesModel := new(eventnotificationsv1.SubscriptionCreateAttributesSmsAttributes)
				subscriptionCreateAttributesModel.To = []string{"testString"}

				// Construct an instance of the CreateSubscriptionOptions model
				createSubscriptionOptionsModel := new(eventnotificationsv1.CreateSubscriptionOptions)
				createSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Name = core.StringPtr("testString")
				createSubscriptionOptionsModel.DestinationID = core.StringPtr("testString")
				createSubscriptionOptionsModel.TopicID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Attributes = subscriptionCreateAttributesModel
				createSubscriptionOptionsModel.Description = core.StringPtr("testString")
				createSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.CreateSubscription(createSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSubscription with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the SubscriptionCreateAttributesSmsAttributes model
				subscriptionCreateAttributesModel := new(eventnotificationsv1.SubscriptionCreateAttributesSmsAttributes)
				subscriptionCreateAttributesModel.To = []string{"testString"}

				// Construct an instance of the CreateSubscriptionOptions model
				createSubscriptionOptionsModel := new(eventnotificationsv1.CreateSubscriptionOptions)
				createSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Name = core.StringPtr("testString")
				createSubscriptionOptionsModel.DestinationID = core.StringPtr("testString")
				createSubscriptionOptionsModel.TopicID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Attributes = subscriptionCreateAttributesModel
				createSubscriptionOptionsModel.Description = core.StringPtr("testString")
				createSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.CreateSubscription(createSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSubscriptionOptions model with no property values
				createSubscriptionOptionsModelNew := new(eventnotificationsv1.CreateSubscriptionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.CreateSubscription(createSubscriptionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateSubscription successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the SubscriptionCreateAttributesSmsAttributes model
				subscriptionCreateAttributesModel := new(eventnotificationsv1.SubscriptionCreateAttributesSmsAttributes)
				subscriptionCreateAttributesModel.To = []string{"testString"}

				// Construct an instance of the CreateSubscriptionOptions model
				createSubscriptionOptionsModel := new(eventnotificationsv1.CreateSubscriptionOptions)
				createSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Name = core.StringPtr("testString")
				createSubscriptionOptionsModel.DestinationID = core.StringPtr("testString")
				createSubscriptionOptionsModel.TopicID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Attributes = subscriptionCreateAttributesModel
				createSubscriptionOptionsModel.Description = core.StringPtr("testString")
				createSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.CreateSubscription(createSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSubscriptions(listSubscriptionsOptions *ListSubscriptionsOptions) - Operation response error`, func() {
		listSubscriptionsPath := "/v1/instances/testString/subscriptions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSubscriptionsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for offset query parameter
					// TODO: Add check for limit query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSubscriptions with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListSubscriptionsOptions model
				listSubscriptionsOptionsModel := new(eventnotificationsv1.ListSubscriptionsOptions)
				listSubscriptionsOptionsModel.InstanceID = core.StringPtr("testString")
				listSubscriptionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSubscriptionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listSubscriptionsOptionsModel.Search = core.StringPtr("testString")
				listSubscriptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.ListSubscriptions(listSubscriptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.ListSubscriptions(listSubscriptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSubscriptions(listSubscriptionsOptions *ListSubscriptionsOptions)`, func() {
		listSubscriptionsPath := "/v1/instances/testString/subscriptions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSubscriptionsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for offset query parameter
					// TODO: Add check for limit query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "subscriptions": [{"id": "ID", "name": "Name", "description": "Description", "destination_id": "DestinationID", "destination_name": "DestinationName", "destination_type": "sms_ibm", "topic_id": "TopicID", "topic_name": "TopicName", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListSubscriptions successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the ListSubscriptionsOptions model
				listSubscriptionsOptionsModel := new(eventnotificationsv1.ListSubscriptionsOptions)
				listSubscriptionsOptionsModel.InstanceID = core.StringPtr("testString")
				listSubscriptionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSubscriptionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listSubscriptionsOptionsModel.Search = core.StringPtr("testString")
				listSubscriptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.ListSubscriptionsWithContext(ctx, listSubscriptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.ListSubscriptions(listSubscriptionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.ListSubscriptionsWithContext(ctx, listSubscriptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSubscriptionsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for offset query parameter
					// TODO: Add check for limit query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "subscriptions": [{"id": "ID", "name": "Name", "description": "Description", "destination_id": "DestinationID", "destination_name": "DestinationName", "destination_type": "sms_ibm", "topic_id": "TopicID", "topic_name": "TopicName", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListSubscriptions successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.ListSubscriptions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSubscriptionsOptions model
				listSubscriptionsOptionsModel := new(eventnotificationsv1.ListSubscriptionsOptions)
				listSubscriptionsOptionsModel.InstanceID = core.StringPtr("testString")
				listSubscriptionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSubscriptionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listSubscriptionsOptionsModel.Search = core.StringPtr("testString")
				listSubscriptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.ListSubscriptions(listSubscriptionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSubscriptions with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListSubscriptionsOptions model
				listSubscriptionsOptionsModel := new(eventnotificationsv1.ListSubscriptionsOptions)
				listSubscriptionsOptionsModel.InstanceID = core.StringPtr("testString")
				listSubscriptionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSubscriptionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listSubscriptionsOptionsModel.Search = core.StringPtr("testString")
				listSubscriptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.ListSubscriptions(listSubscriptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListSubscriptionsOptions model with no property values
				listSubscriptionsOptionsModelNew := new(eventnotificationsv1.ListSubscriptionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.ListSubscriptions(listSubscriptionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListSubscriptions successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListSubscriptionsOptions model
				listSubscriptionsOptionsModel := new(eventnotificationsv1.ListSubscriptionsOptions)
				listSubscriptionsOptionsModel.InstanceID = core.StringPtr("testString")
				listSubscriptionsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listSubscriptionsOptionsModel.Limit = core.Int64Ptr(int64(1))
				listSubscriptionsOptionsModel.Search = core.StringPtr("testString")
				listSubscriptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.ListSubscriptions(listSubscriptionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSubscription(getSubscriptionOptions *GetSubscriptionOptions) - Operation response error`, func() {
		getSubscriptionPath := "/v1/instances/testString/subscriptions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSubscriptionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSubscription with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetSubscriptionOptions model
				getSubscriptionOptionsModel := new(eventnotificationsv1.GetSubscriptionOptions)
				getSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				getSubscriptionOptionsModel.ID = core.StringPtr("testString")
				getSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.GetSubscription(getSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.GetSubscription(getSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSubscription(getSubscriptionOptions *GetSubscriptionOptions)`, func() {
		getSubscriptionPath := "/v1/instances/testString/subscriptions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSubscriptionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {"to": ["To"], "recipient_selection": "only_destination"}}`)
				}))
			})
			It(`Invoke GetSubscription successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the GetSubscriptionOptions model
				getSubscriptionOptionsModel := new(eventnotificationsv1.GetSubscriptionOptions)
				getSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				getSubscriptionOptionsModel.ID = core.StringPtr("testString")
				getSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.GetSubscriptionWithContext(ctx, getSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.GetSubscription(getSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.GetSubscriptionWithContext(ctx, getSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSubscriptionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {"to": ["To"], "recipient_selection": "only_destination"}}`)
				}))
			})
			It(`Invoke GetSubscription successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.GetSubscription(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSubscriptionOptions model
				getSubscriptionOptionsModel := new(eventnotificationsv1.GetSubscriptionOptions)
				getSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				getSubscriptionOptionsModel.ID = core.StringPtr("testString")
				getSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.GetSubscription(getSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSubscription with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetSubscriptionOptions model
				getSubscriptionOptionsModel := new(eventnotificationsv1.GetSubscriptionOptions)
				getSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				getSubscriptionOptionsModel.ID = core.StringPtr("testString")
				getSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.GetSubscription(getSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSubscriptionOptions model with no property values
				getSubscriptionOptionsModelNew := new(eventnotificationsv1.GetSubscriptionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.GetSubscription(getSubscriptionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSubscription successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetSubscriptionOptions model
				getSubscriptionOptionsModel := new(eventnotificationsv1.GetSubscriptionOptions)
				getSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				getSubscriptionOptionsModel.ID = core.StringPtr("testString")
				getSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.GetSubscription(getSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSubscription(deleteSubscriptionOptions *DeleteSubscriptionOptions)`, func() {
		deleteSubscriptionPath := "/v1/instances/testString/subscriptions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSubscriptionPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSubscription successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := eventNotificationsService.DeleteSubscription(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSubscriptionOptions model
				deleteSubscriptionOptionsModel := new(eventnotificationsv1.DeleteSubscriptionOptions)
				deleteSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				deleteSubscriptionOptionsModel.ID = core.StringPtr("testString")
				deleteSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = eventNotificationsService.DeleteSubscription(deleteSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSubscription with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DeleteSubscriptionOptions model
				deleteSubscriptionOptionsModel := new(eventnotificationsv1.DeleteSubscriptionOptions)
				deleteSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				deleteSubscriptionOptionsModel.ID = core.StringPtr("testString")
				deleteSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := eventNotificationsService.DeleteSubscription(deleteSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSubscriptionOptions model with no property values
				deleteSubscriptionOptionsModelNew := new(eventnotificationsv1.DeleteSubscriptionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = eventNotificationsService.DeleteSubscription(deleteSubscriptionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSubscription(updateSubscriptionOptions *UpdateSubscriptionOptions) - Operation response error`, func() {
		updateSubscriptionPath := "/v1/instances/testString/subscriptions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSubscriptionPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSubscription with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the SubscriptionUpdateAttributesSmsAttributes model
				subscriptionUpdateAttributesModel := new(eventnotificationsv1.SubscriptionUpdateAttributesSmsAttributes)
				subscriptionUpdateAttributesModel.To = []string{"testString"}

				// Construct an instance of the UpdateSubscriptionOptions model
				updateSubscriptionOptionsModel := new(eventnotificationsv1.UpdateSubscriptionOptions)
				updateSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.ID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Name = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Description = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Attributes = subscriptionUpdateAttributesModel
				updateSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.UpdateSubscription(updateSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.UpdateSubscription(updateSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSubscription(updateSubscriptionOptions *UpdateSubscriptionOptions)`, func() {
		updateSubscriptionPath := "/v1/instances/testString/subscriptions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSubscriptionPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {"to": ["To"], "recipient_selection": "only_destination"}}`)
				}))
			})
			It(`Invoke UpdateSubscription successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the SubscriptionUpdateAttributesSmsAttributes model
				subscriptionUpdateAttributesModel := new(eventnotificationsv1.SubscriptionUpdateAttributesSmsAttributes)
				subscriptionUpdateAttributesModel.To = []string{"testString"}

				// Construct an instance of the UpdateSubscriptionOptions model
				updateSubscriptionOptionsModel := new(eventnotificationsv1.UpdateSubscriptionOptions)
				updateSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.ID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Name = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Description = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Attributes = subscriptionUpdateAttributesModel
				updateSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.UpdateSubscriptionWithContext(ctx, updateSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.UpdateSubscription(updateSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.UpdateSubscriptionWithContext(ctx, updateSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSubscriptionPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {"to": ["To"], "recipient_selection": "only_destination"}}`)
				}))
			})
			It(`Invoke UpdateSubscription successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.UpdateSubscription(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SubscriptionUpdateAttributesSmsAttributes model
				subscriptionUpdateAttributesModel := new(eventnotificationsv1.SubscriptionUpdateAttributesSmsAttributes)
				subscriptionUpdateAttributesModel.To = []string{"testString"}

				// Construct an instance of the UpdateSubscriptionOptions model
				updateSubscriptionOptionsModel := new(eventnotificationsv1.UpdateSubscriptionOptions)
				updateSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.ID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Name = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Description = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Attributes = subscriptionUpdateAttributesModel
				updateSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.UpdateSubscription(updateSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSubscription with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the SubscriptionUpdateAttributesSmsAttributes model
				subscriptionUpdateAttributesModel := new(eventnotificationsv1.SubscriptionUpdateAttributesSmsAttributes)
				subscriptionUpdateAttributesModel.To = []string{"testString"}

				// Construct an instance of the UpdateSubscriptionOptions model
				updateSubscriptionOptionsModel := new(eventnotificationsv1.UpdateSubscriptionOptions)
				updateSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.ID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Name = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Description = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Attributes = subscriptionUpdateAttributesModel
				updateSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.UpdateSubscription(updateSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSubscriptionOptions model with no property values
				updateSubscriptionOptionsModelNew := new(eventnotificationsv1.UpdateSubscriptionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.UpdateSubscription(updateSubscriptionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateSubscription successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the SubscriptionUpdateAttributesSmsAttributes model
				subscriptionUpdateAttributesModel := new(eventnotificationsv1.SubscriptionUpdateAttributesSmsAttributes)
				subscriptionUpdateAttributesModel.To = []string{"testString"}

				// Construct an instance of the UpdateSubscriptionOptions model
				updateSubscriptionOptionsModel := new(eventnotificationsv1.UpdateSubscriptionOptions)
				updateSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.ID = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Name = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Description = core.StringPtr("testString")
				updateSubscriptionOptionsModel.Attributes = subscriptionUpdateAttributesModel
				updateSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.UpdateSubscription(updateSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			eventNotificationsService, _ := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
				URL:           "http://eventnotificationsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateDestinationOptions successfully`, func() {
				// Construct an instance of the DestinationConfigParamsWebhookDestinationConfig model
				destinationConfigParamsModel := new(eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig)
				Expect(destinationConfigParamsModel).ToNot(BeNil())
				destinationConfigParamsModel.URL = core.StringPtr("testString")
				destinationConfigParamsModel.Verb = core.StringPtr("get")
				destinationConfigParamsModel.CustomHeaders = make(map[string]string)
				destinationConfigParamsModel.SensitiveHeaders = []string{"testString"}
				Expect(destinationConfigParamsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(destinationConfigParamsModel.Verb).To(Equal(core.StringPtr("get")))
				Expect(destinationConfigParamsModel.CustomHeaders).To(Equal(make(map[string]string)))
				Expect(destinationConfigParamsModel.SensitiveHeaders).To(Equal([]string{"testString"}))

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				Expect(destinationConfigModel).ToNot(BeNil())
				destinationConfigModel.Params = destinationConfigParamsModel
				Expect(destinationConfigModel.Params).To(Equal(destinationConfigParamsModel))

				// Construct an instance of the CreateDestinationOptions model
				instanceID := "testString"
				createDestinationOptionsName := "testString"
				createDestinationOptionsType := "webhook"
				createDestinationOptionsModel := eventNotificationsService.NewCreateDestinationOptions(instanceID, createDestinationOptionsName, createDestinationOptionsType)
				createDestinationOptionsModel.SetInstanceID("testString")
				createDestinationOptionsModel.SetName("testString")
				createDestinationOptionsModel.SetType("webhook")
				createDestinationOptionsModel.SetDescription("testString")
				createDestinationOptionsModel.SetConfig(destinationConfigModel)
				createDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDestinationOptionsModel).ToNot(BeNil())
				Expect(createDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Type).To(Equal(core.StringPtr("webhook")))
				Expect(createDestinationOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Config).To(Equal(destinationConfigModel))
				Expect(createDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSubscriptionOptions successfully`, func() {
				// Construct an instance of the SubscriptionCreateAttributesSmsAttributes model
				subscriptionCreateAttributesModel := new(eventnotificationsv1.SubscriptionCreateAttributesSmsAttributes)
				Expect(subscriptionCreateAttributesModel).ToNot(BeNil())
				subscriptionCreateAttributesModel.To = []string{"testString"}
				Expect(subscriptionCreateAttributesModel.To).To(Equal([]string{"testString"}))

				// Construct an instance of the CreateSubscriptionOptions model
				instanceID := "testString"
				createSubscriptionOptionsName := "testString"
				createSubscriptionOptionsDestinationID := "testString"
				createSubscriptionOptionsTopicID := "testString"
				var createSubscriptionOptionsAttributes eventnotificationsv1.SubscriptionCreateAttributesIntf = nil
				createSubscriptionOptionsModel := eventNotificationsService.NewCreateSubscriptionOptions(instanceID, createSubscriptionOptionsName, createSubscriptionOptionsDestinationID, createSubscriptionOptionsTopicID, createSubscriptionOptionsAttributes)
				createSubscriptionOptionsModel.SetInstanceID("testString")
				createSubscriptionOptionsModel.SetName("testString")
				createSubscriptionOptionsModel.SetDestinationID("testString")
				createSubscriptionOptionsModel.SetTopicID("testString")
				createSubscriptionOptionsModel.SetAttributes(subscriptionCreateAttributesModel)
				createSubscriptionOptionsModel.SetDescription("testString")
				createSubscriptionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSubscriptionOptionsModel).ToNot(BeNil())
				Expect(createSubscriptionOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createSubscriptionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createSubscriptionOptionsModel.DestinationID).To(Equal(core.StringPtr("testString")))
				Expect(createSubscriptionOptionsModel.TopicID).To(Equal(core.StringPtr("testString")))
				Expect(createSubscriptionOptionsModel.Attributes).To(Equal(subscriptionCreateAttributesModel))
				Expect(createSubscriptionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createSubscriptionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTopicOptions successfully`, func() {
				// Construct an instance of the Rules model
				rulesModel := new(eventnotificationsv1.Rules)
				Expect(rulesModel).ToNot(BeNil())
				rulesModel.Enabled = core.BoolPtr(true)
				rulesModel.EventTypeFilter = core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")
				rulesModel.NotificationFilter = core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")
				Expect(rulesModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(rulesModel.EventTypeFilter).To(Equal(core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")))
				Expect(rulesModel.NotificationFilter).To(Equal(core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")))

				// Construct an instance of the TopicUpdateSourcesItem model
				topicUpdateSourcesItemModel := new(eventnotificationsv1.TopicUpdateSourcesItem)
				Expect(topicUpdateSourcesItemModel).ToNot(BeNil())
				topicUpdateSourcesItemModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				topicUpdateSourcesItemModel.Rules = []eventnotificationsv1.Rules{*rulesModel}
				Expect(topicUpdateSourcesItemModel.ID).To(Equal(core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")))
				Expect(topicUpdateSourcesItemModel.Rules).To(Equal([]eventnotificationsv1.Rules{*rulesModel}))

				// Construct an instance of the CreateTopicOptions model
				instanceID := "testString"
				createTopicOptionsName := "testString"
				createTopicOptionsModel := eventNotificationsService.NewCreateTopicOptions(instanceID, createTopicOptionsName)
				createTopicOptionsModel.SetInstanceID("testString")
				createTopicOptionsModel.SetName("testString")
				createTopicOptionsModel.SetDescription("testString")
				createTopicOptionsModel.SetSources([]eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel})
				createTopicOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTopicOptionsModel).ToNot(BeNil())
				Expect(createTopicOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createTopicOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createTopicOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createTopicOptionsModel.Sources).To(Equal([]eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel}))
				Expect(createTopicOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDestinationOptions successfully`, func() {
				// Construct an instance of the DeleteDestinationOptions model
				instanceID := "testString"
				id := "testString"
				deleteDestinationOptionsModel := eventNotificationsService.NewDeleteDestinationOptions(instanceID, id)
				deleteDestinationOptionsModel.SetInstanceID("testString")
				deleteDestinationOptionsModel.SetID("testString")
				deleteDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDestinationOptionsModel).ToNot(BeNil())
				Expect(deleteDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDestinationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSubscriptionOptions successfully`, func() {
				// Construct an instance of the DeleteSubscriptionOptions model
				instanceID := "testString"
				id := "testString"
				deleteSubscriptionOptionsModel := eventNotificationsService.NewDeleteSubscriptionOptions(instanceID, id)
				deleteSubscriptionOptionsModel.SetInstanceID("testString")
				deleteSubscriptionOptionsModel.SetID("testString")
				deleteSubscriptionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSubscriptionOptionsModel).ToNot(BeNil())
				Expect(deleteSubscriptionOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSubscriptionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSubscriptionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTopicOptions successfully`, func() {
				// Construct an instance of the DeleteTopicOptions model
				instanceID := "testString"
				id := "testString"
				deleteTopicOptionsModel := eventNotificationsService.NewDeleteTopicOptions(instanceID, id)
				deleteTopicOptionsModel.SetInstanceID("testString")
				deleteTopicOptionsModel.SetID("testString")
				deleteTopicOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTopicOptionsModel).ToNot(BeNil())
				Expect(deleteTopicOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTopicOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTopicOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDestinationConfig successfully`, func() {
				var params eventnotificationsv1.DestinationConfigParamsIntf = nil
				_, err := eventNotificationsService.NewDestinationConfig(params)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewGetDestinationOptions successfully`, func() {
				// Construct an instance of the GetDestinationOptions model
				instanceID := "testString"
				id := "testString"
				getDestinationOptionsModel := eventNotificationsService.NewGetDestinationOptions(instanceID, id)
				getDestinationOptionsModel.SetInstanceID("testString")
				getDestinationOptionsModel.SetID("testString")
				getDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDestinationOptionsModel).ToNot(BeNil())
				Expect(getDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getDestinationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSourceOptions successfully`, func() {
				// Construct an instance of the GetSourceOptions model
				instanceID := "testString"
				id := "testString"
				getSourceOptionsModel := eventNotificationsService.NewGetSourceOptions(instanceID, id)
				getSourceOptionsModel.SetInstanceID("testString")
				getSourceOptionsModel.SetID("testString")
				getSourceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSourceOptionsModel).ToNot(BeNil())
				Expect(getSourceOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getSourceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getSourceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSubscriptionOptions successfully`, func() {
				// Construct an instance of the GetSubscriptionOptions model
				instanceID := "testString"
				id := "testString"
				getSubscriptionOptionsModel := eventNotificationsService.NewGetSubscriptionOptions(instanceID, id)
				getSubscriptionOptionsModel.SetInstanceID("testString")
				getSubscriptionOptionsModel.SetID("testString")
				getSubscriptionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSubscriptionOptionsModel).ToNot(BeNil())
				Expect(getSubscriptionOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getSubscriptionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getSubscriptionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTopicOptions successfully`, func() {
				// Construct an instance of the GetTopicOptions model
				instanceID := "testString"
				id := "testString"
				getTopicOptionsModel := eventNotificationsService.NewGetTopicOptions(instanceID, id)
				getTopicOptionsModel.SetInstanceID("testString")
				getTopicOptionsModel.SetID("testString")
				getTopicOptionsModel.SetInclude("testString")
				getTopicOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTopicOptionsModel).ToNot(BeNil())
				Expect(getTopicOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getTopicOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTopicOptionsModel.Include).To(Equal(core.StringPtr("testString")))
				Expect(getTopicOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDestinationsOptions successfully`, func() {
				// Construct an instance of the ListDestinationsOptions model
				instanceID := "testString"
				listDestinationsOptionsModel := eventNotificationsService.NewListDestinationsOptions(instanceID)
				listDestinationsOptionsModel.SetInstanceID("testString")
				listDestinationsOptionsModel.SetLimit(int64(1))
				listDestinationsOptionsModel.SetOffset(int64(0))
				listDestinationsOptionsModel.SetSearch("testString")
				listDestinationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDestinationsOptionsModel).ToNot(BeNil())
				Expect(listDestinationsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listDestinationsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listDestinationsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listDestinationsOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listDestinationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSourcesOptions successfully`, func() {
				// Construct an instance of the ListSourcesOptions model
				instanceID := "testString"
				listSourcesOptionsModel := eventNotificationsService.NewListSourcesOptions(instanceID)
				listSourcesOptionsModel.SetInstanceID("testString")
				listSourcesOptionsModel.SetLimit(int64(1))
				listSourcesOptionsModel.SetOffset(int64(0))
				listSourcesOptionsModel.SetSearch("testString")
				listSourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSourcesOptionsModel).ToNot(BeNil())
				Expect(listSourcesOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listSourcesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listSourcesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listSourcesOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listSourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSubscriptionsOptions successfully`, func() {
				// Construct an instance of the ListSubscriptionsOptions model
				instanceID := "testString"
				listSubscriptionsOptionsModel := eventNotificationsService.NewListSubscriptionsOptions(instanceID)
				listSubscriptionsOptionsModel.SetInstanceID("testString")
				listSubscriptionsOptionsModel.SetOffset(int64(0))
				listSubscriptionsOptionsModel.SetLimit(int64(1))
				listSubscriptionsOptionsModel.SetSearch("testString")
				listSubscriptionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSubscriptionsOptionsModel).ToNot(BeNil())
				Expect(listSubscriptionsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listSubscriptionsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listSubscriptionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listSubscriptionsOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listSubscriptionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTopicsOptions successfully`, func() {
				// Construct an instance of the ListTopicsOptions model
				instanceID := "testString"
				listTopicsOptionsModel := eventNotificationsService.NewListTopicsOptions(instanceID)
				listTopicsOptionsModel.SetInstanceID("testString")
				listTopicsOptionsModel.SetLimit(int64(1))
				listTopicsOptionsModel.SetOffset(int64(0))
				listTopicsOptionsModel.SetSearch("testString")
				listTopicsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTopicsOptionsModel).ToNot(BeNil())
				Expect(listTopicsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listTopicsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listTopicsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listTopicsOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listTopicsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceTopicOptions successfully`, func() {
				// Construct an instance of the Rules model
				rulesModel := new(eventnotificationsv1.Rules)
				Expect(rulesModel).ToNot(BeNil())
				rulesModel.Enabled = core.BoolPtr(true)
				rulesModel.EventTypeFilter = core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")
				rulesModel.NotificationFilter = core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")
				Expect(rulesModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(rulesModel.EventTypeFilter).To(Equal(core.StringPtr("$.notification_event_info.event_type == 'cert_manager'")))
				Expect(rulesModel.NotificationFilter).To(Equal(core.StringPtr("$.notification.findings[0].severity == 'MODERATE'")))

				// Construct an instance of the TopicUpdateSourcesItem model
				topicUpdateSourcesItemModel := new(eventnotificationsv1.TopicUpdateSourcesItem)
				Expect(topicUpdateSourcesItemModel).ToNot(BeNil())
				topicUpdateSourcesItemModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				topicUpdateSourcesItemModel.Rules = []eventnotificationsv1.Rules{*rulesModel}
				Expect(topicUpdateSourcesItemModel.ID).To(Equal(core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")))
				Expect(topicUpdateSourcesItemModel.Rules).To(Equal([]eventnotificationsv1.Rules{*rulesModel}))

				// Construct an instance of the ReplaceTopicOptions model
				instanceID := "testString"
				id := "testString"
				replaceTopicOptionsModel := eventNotificationsService.NewReplaceTopicOptions(instanceID, id)
				replaceTopicOptionsModel.SetInstanceID("testString")
				replaceTopicOptionsModel.SetID("testString")
				replaceTopicOptionsModel.SetName("testString")
				replaceTopicOptionsModel.SetDescription("testString")
				replaceTopicOptionsModel.SetSources([]eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel})
				replaceTopicOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceTopicOptionsModel).ToNot(BeNil())
				Expect(replaceTopicOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceTopicOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceTopicOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(replaceTopicOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(replaceTopicOptionsModel.Sources).To(Equal([]eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel}))
				Expect(replaceTopicOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRules successfully`, func() {
				eventTypeFilter := "$.*"
				_model, err := eventNotificationsService.NewRules(eventTypeFilter)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTopicUpdateSourcesItem successfully`, func() {
				id := "testString"
				rules := []eventnotificationsv1.Rules{}
				_model, err := eventNotificationsService.NewTopicUpdateSourcesItem(id, rules)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateDestinationOptions successfully`, func() {
				// Construct an instance of the DestinationConfigParamsWebhookDestinationConfig model
				destinationConfigParamsModel := new(eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig)
				Expect(destinationConfigParamsModel).ToNot(BeNil())
				destinationConfigParamsModel.URL = core.StringPtr("testString")
				destinationConfigParamsModel.Verb = core.StringPtr("get")
				destinationConfigParamsModel.CustomHeaders = make(map[string]string)
				destinationConfigParamsModel.SensitiveHeaders = []string{"testString"}
				Expect(destinationConfigParamsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(destinationConfigParamsModel.Verb).To(Equal(core.StringPtr("get")))
				Expect(destinationConfigParamsModel.CustomHeaders).To(Equal(make(map[string]string)))
				Expect(destinationConfigParamsModel.SensitiveHeaders).To(Equal([]string{"testString"}))

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				Expect(destinationConfigModel).ToNot(BeNil())
				destinationConfigModel.Params = destinationConfigParamsModel
				Expect(destinationConfigModel.Params).To(Equal(destinationConfigParamsModel))

				// Construct an instance of the UpdateDestinationOptions model
				instanceID := "testString"
				id := "testString"
				updateDestinationOptionsModel := eventNotificationsService.NewUpdateDestinationOptions(instanceID, id)
				updateDestinationOptionsModel.SetInstanceID("testString")
				updateDestinationOptionsModel.SetID("testString")
				updateDestinationOptionsModel.SetName("testString")
				updateDestinationOptionsModel.SetDescription("testString")
				updateDestinationOptionsModel.SetConfig(destinationConfigModel)
				updateDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDestinationOptionsModel).ToNot(BeNil())
				Expect(updateDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Config).To(Equal(destinationConfigModel))
				Expect(updateDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSubscriptionOptions successfully`, func() {
				// Construct an instance of the SubscriptionUpdateAttributesSmsAttributes model
				subscriptionUpdateAttributesModel := new(eventnotificationsv1.SubscriptionUpdateAttributesSmsAttributes)
				Expect(subscriptionUpdateAttributesModel).ToNot(BeNil())
				subscriptionUpdateAttributesModel.To = []string{"testString"}
				Expect(subscriptionUpdateAttributesModel.To).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateSubscriptionOptions model
				instanceID := "testString"
				id := "testString"
				updateSubscriptionOptionsModel := eventNotificationsService.NewUpdateSubscriptionOptions(instanceID, id)
				updateSubscriptionOptionsModel.SetInstanceID("testString")
				updateSubscriptionOptionsModel.SetID("testString")
				updateSubscriptionOptionsModel.SetName("testString")
				updateSubscriptionOptionsModel.SetDescription("testString")
				updateSubscriptionOptionsModel.SetAttributes(subscriptionUpdateAttributesModel)
				updateSubscriptionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSubscriptionOptionsModel).ToNot(BeNil())
				Expect(updateSubscriptionOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateSubscriptionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateSubscriptionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateSubscriptionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateSubscriptionOptionsModel.Attributes).To(Equal(subscriptionUpdateAttributesModel))
				Expect(updateSubscriptionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDestinationConfigParamsWebhookDestinationConfig successfully`, func() {
				url := "testString"
				verb := "get"
				_model, err := eventNotificationsService.NewDestinationConfigParamsWebhookDestinationConfig(url, verb)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubscriptionCreateAttributesEmailAttributes successfully`, func() {
				to := []string{"testString"}
				addNotificationPayload := false
				_model, err := eventNotificationsService.NewSubscriptionCreateAttributesEmailAttributes(to, addNotificationPayload)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubscriptionCreateAttributesSmsAttributes successfully`, func() {
				to := []string{"testString"}
				_model, err := eventNotificationsService.NewSubscriptionCreateAttributesSmsAttributes(to)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubscriptionCreateAttributesWebhookAttributes successfully`, func() {
				signingEnabled := true
				_model, err := eventNotificationsService.NewSubscriptionCreateAttributesWebhookAttributes(signingEnabled)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubscriptionUpdateAttributesEmailAttributes successfully`, func() {
				to := []string{"testString"}
				addNotificationPayload := false
				_model, err := eventNotificationsService.NewSubscriptionUpdateAttributesEmailAttributes(to, addNotificationPayload)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubscriptionUpdateAttributesSmsAttributes successfully`, func() {
				to := []string{"testString"}
				_model, err := eventNotificationsService.NewSubscriptionUpdateAttributesSmsAttributes(to)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubscriptionUpdateAttributesWebhookAttributes successfully`, func() {
				signingEnabled := true
				_model, err := eventNotificationsService.NewSubscriptionUpdateAttributesWebhookAttributes(signingEnabled)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
