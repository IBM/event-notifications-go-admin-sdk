/**
 * (C) Copyright IBM Corp. 2023.
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
				"EVENT_NOTIFICATIONS_URL": "https://eventnotificationsv1/api",
				"EVENT_NOTIFICATIONS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1UsingExternalConfig(&eventnotificationsv1.EventNotificationsV1Options{
				})
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
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1UsingExternalConfig(&eventnotificationsv1.EventNotificationsV1Options{
				})
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
				"EVENT_NOTIFICATIONS_URL": "https://eventnotificationsv1/api",
				"EVENT_NOTIFICATIONS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1UsingExternalConfig(&eventnotificationsv1.EventNotificationsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(eventNotificationsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"EVENT_NOTIFICATIONS_AUTH_TYPE":   "NOAuth",
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
	Describe(`SendNotifications(sendNotificationsOptions *SendNotificationsOptions) - Operation response error`, func() {
		sendNotificationsPath := "/v1/instances/testString/notifications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(sendNotificationsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SendNotifications with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the NotificationCreate model
				notificationCreateModel := new(eventnotificationsv1.NotificationCreate)
				notificationCreateModel.Specversion = core.StringPtr("1.0")
				notificationCreateModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				notificationCreateModel.ID = core.StringPtr("testString")
				notificationCreateModel.Source = core.StringPtr("testString")
				notificationCreateModel.Type = core.StringPtr("testString")
				notificationCreateModel.Ibmenseverity = core.StringPtr("testString")
				notificationCreateModel.Ibmensourceid = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultshort = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultlong = core.StringPtr("testString")
				notificationCreateModel.Subject = core.StringPtr("testString")
				notificationCreateModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				notificationCreateModel.Datacontenttype = core.StringPtr("application/json")
				notificationCreateModel.Ibmenpushto = core.StringPtr(`{"platforms":["push_android"]}`)
				notificationCreateModel.Ibmenfcmbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsheaders = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromebody = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromeheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenfirefoxbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenfirefoxheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenhuaweibody = core.StringPtr("testString")
				notificationCreateModel.Ibmensafaribody = core.StringPtr("testString")
				notificationCreateModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsOptions model
				sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
				sendNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendNotificationsOptionsModel.Body = notificationCreateModel
				sendNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.SendNotifications(sendNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.SendNotifications(sendNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SendNotifications(sendNotificationsOptions *SendNotificationsOptions)`, func() {
		sendNotificationsPath := "/v1/instances/testString/notifications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(sendNotificationsPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"notification_id": "NotificationID"}`)
				}))
			})
			It(`Invoke SendNotifications successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the NotificationCreate model
				notificationCreateModel := new(eventnotificationsv1.NotificationCreate)
				notificationCreateModel.Specversion = core.StringPtr("1.0")
				notificationCreateModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				notificationCreateModel.ID = core.StringPtr("testString")
				notificationCreateModel.Source = core.StringPtr("testString")
				notificationCreateModel.Type = core.StringPtr("testString")
				notificationCreateModel.Ibmenseverity = core.StringPtr("testString")
				notificationCreateModel.Ibmensourceid = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultshort = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultlong = core.StringPtr("testString")
				notificationCreateModel.Subject = core.StringPtr("testString")
				notificationCreateModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				notificationCreateModel.Datacontenttype = core.StringPtr("application/json")
				notificationCreateModel.Ibmenpushto = core.StringPtr(`{"platforms":["push_android"]}`)
				notificationCreateModel.Ibmenfcmbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsheaders = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromebody = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromeheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenfirefoxbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenfirefoxheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenhuaweibody = core.StringPtr("testString")
				notificationCreateModel.Ibmensafaribody = core.StringPtr("testString")
				notificationCreateModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsOptions model
				sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
				sendNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendNotificationsOptionsModel.Body = notificationCreateModel
				sendNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.SendNotificationsWithContext(ctx, sendNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.SendNotifications(sendNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.SendNotificationsWithContext(ctx, sendNotificationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(sendNotificationsPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"notification_id": "NotificationID"}`)
				}))
			})
			It(`Invoke SendNotifications successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.SendNotifications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NotificationCreate model
				notificationCreateModel := new(eventnotificationsv1.NotificationCreate)
				notificationCreateModel.Specversion = core.StringPtr("1.0")
				notificationCreateModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				notificationCreateModel.ID = core.StringPtr("testString")
				notificationCreateModel.Source = core.StringPtr("testString")
				notificationCreateModel.Type = core.StringPtr("testString")
				notificationCreateModel.Ibmenseverity = core.StringPtr("testString")
				notificationCreateModel.Ibmensourceid = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultshort = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultlong = core.StringPtr("testString")
				notificationCreateModel.Subject = core.StringPtr("testString")
				notificationCreateModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				notificationCreateModel.Datacontenttype = core.StringPtr("application/json")
				notificationCreateModel.Ibmenpushto = core.StringPtr(`{"platforms":["push_android"]}`)
				notificationCreateModel.Ibmenfcmbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsheaders = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromebody = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromeheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenfirefoxbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenfirefoxheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenhuaweibody = core.StringPtr("testString")
				notificationCreateModel.Ibmensafaribody = core.StringPtr("testString")
				notificationCreateModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsOptions model
				sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
				sendNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendNotificationsOptionsModel.Body = notificationCreateModel
				sendNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.SendNotifications(sendNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SendNotifications with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the NotificationCreate model
				notificationCreateModel := new(eventnotificationsv1.NotificationCreate)
				notificationCreateModel.Specversion = core.StringPtr("1.0")
				notificationCreateModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				notificationCreateModel.ID = core.StringPtr("testString")
				notificationCreateModel.Source = core.StringPtr("testString")
				notificationCreateModel.Type = core.StringPtr("testString")
				notificationCreateModel.Ibmenseverity = core.StringPtr("testString")
				notificationCreateModel.Ibmensourceid = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultshort = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultlong = core.StringPtr("testString")
				notificationCreateModel.Subject = core.StringPtr("testString")
				notificationCreateModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				notificationCreateModel.Datacontenttype = core.StringPtr("application/json")
				notificationCreateModel.Ibmenpushto = core.StringPtr(`{"platforms":["push_android"]}`)
				notificationCreateModel.Ibmenfcmbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsheaders = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromebody = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromeheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenfirefoxbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenfirefoxheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenhuaweibody = core.StringPtr("testString")
				notificationCreateModel.Ibmensafaribody = core.StringPtr("testString")
				notificationCreateModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsOptions model
				sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
				sendNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendNotificationsOptionsModel.Body = notificationCreateModel
				sendNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.SendNotifications(sendNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SendNotificationsOptions model with no property values
				sendNotificationsOptionsModelNew := new(eventnotificationsv1.SendNotificationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.SendNotifications(sendNotificationsOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke SendNotifications successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the NotificationCreate model
				notificationCreateModel := new(eventnotificationsv1.NotificationCreate)
				notificationCreateModel.Specversion = core.StringPtr("1.0")
				notificationCreateModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				notificationCreateModel.ID = core.StringPtr("testString")
				notificationCreateModel.Source = core.StringPtr("testString")
				notificationCreateModel.Type = core.StringPtr("testString")
				notificationCreateModel.Ibmenseverity = core.StringPtr("testString")
				notificationCreateModel.Ibmensourceid = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultshort = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultlong = core.StringPtr("testString")
				notificationCreateModel.Subject = core.StringPtr("testString")
				notificationCreateModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				notificationCreateModel.Datacontenttype = core.StringPtr("application/json")
				notificationCreateModel.Ibmenpushto = core.StringPtr(`{"platforms":["push_android"]}`)
				notificationCreateModel.Ibmenfcmbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsheaders = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromebody = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromeheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenfirefoxbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenfirefoxheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenhuaweibody = core.StringPtr("testString")
				notificationCreateModel.Ibmensafaribody = core.StringPtr("testString")
				notificationCreateModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsOptions model
				sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
				sendNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendNotificationsOptionsModel.Body = notificationCreateModel
				sendNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.SendNotifications(sendNotificationsOptionsModel)
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
	Describe(`SendBulkNotifications(sendBulkNotificationsOptions *SendBulkNotificationsOptions) - Operation response error`, func() {
		sendBulkNotificationsPath := "/v1/instances/testString/notifications/bulk"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(sendBulkNotificationsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SendBulkNotifications with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the NotificationCreate model
				notificationCreateModel := new(eventnotificationsv1.NotificationCreate)
				notificationCreateModel.Specversion = core.StringPtr("1.0")
				notificationCreateModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				notificationCreateModel.ID = core.StringPtr("testString")
				notificationCreateModel.Source = core.StringPtr("testString")
				notificationCreateModel.Type = core.StringPtr("testString")
				notificationCreateModel.Ibmenseverity = core.StringPtr("testString")
				notificationCreateModel.Ibmensourceid = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultshort = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultlong = core.StringPtr("testString")
				notificationCreateModel.Subject = core.StringPtr("testString")
				notificationCreateModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				notificationCreateModel.Datacontenttype = core.StringPtr("application/json")
				notificationCreateModel.Ibmenpushto = core.StringPtr(`{"platforms":["push_android"]}`)
				notificationCreateModel.Ibmenfcmbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsheaders = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromebody = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromeheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenfirefoxbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenfirefoxheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenhuaweibody = core.StringPtr("testString")
				notificationCreateModel.Ibmensafaribody = core.StringPtr("testString")
				notificationCreateModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendBulkNotificationsOptions model
				sendBulkNotificationsOptionsModel := new(eventnotificationsv1.SendBulkNotificationsOptions)
				sendBulkNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendBulkNotificationsOptionsModel.BulkMessages = []eventnotificationsv1.NotificationCreate{*notificationCreateModel}
				sendBulkNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.SendBulkNotifications(sendBulkNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.SendBulkNotifications(sendBulkNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SendBulkNotifications(sendBulkNotificationsOptions *SendBulkNotificationsOptions)`, func() {
		sendBulkNotificationsPath := "/v1/instances/testString/notifications/bulk"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(sendBulkNotificationsPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"bulk_notification_id": "BulkNotificationID", "bulk_messages": ["anyValue"]}`)
				}))
			})
			It(`Invoke SendBulkNotifications successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the NotificationCreate model
				notificationCreateModel := new(eventnotificationsv1.NotificationCreate)
				notificationCreateModel.Specversion = core.StringPtr("1.0")
				notificationCreateModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				notificationCreateModel.ID = core.StringPtr("testString")
				notificationCreateModel.Source = core.StringPtr("testString")
				notificationCreateModel.Type = core.StringPtr("testString")
				notificationCreateModel.Ibmenseverity = core.StringPtr("testString")
				notificationCreateModel.Ibmensourceid = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultshort = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultlong = core.StringPtr("testString")
				notificationCreateModel.Subject = core.StringPtr("testString")
				notificationCreateModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				notificationCreateModel.Datacontenttype = core.StringPtr("application/json")
				notificationCreateModel.Ibmenpushto = core.StringPtr(`{"platforms":["push_android"]}`)
				notificationCreateModel.Ibmenfcmbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsheaders = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromebody = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromeheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenfirefoxbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenfirefoxheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenhuaweibody = core.StringPtr("testString")
				notificationCreateModel.Ibmensafaribody = core.StringPtr("testString")
				notificationCreateModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendBulkNotificationsOptions model
				sendBulkNotificationsOptionsModel := new(eventnotificationsv1.SendBulkNotificationsOptions)
				sendBulkNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendBulkNotificationsOptionsModel.BulkMessages = []eventnotificationsv1.NotificationCreate{*notificationCreateModel}
				sendBulkNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.SendBulkNotificationsWithContext(ctx, sendBulkNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.SendBulkNotifications(sendBulkNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.SendBulkNotificationsWithContext(ctx, sendBulkNotificationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(sendBulkNotificationsPath))
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
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"bulk_notification_id": "BulkNotificationID", "bulk_messages": ["anyValue"]}`)
				}))
			})
			It(`Invoke SendBulkNotifications successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.SendBulkNotifications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the NotificationCreate model
				notificationCreateModel := new(eventnotificationsv1.NotificationCreate)
				notificationCreateModel.Specversion = core.StringPtr("1.0")
				notificationCreateModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				notificationCreateModel.ID = core.StringPtr("testString")
				notificationCreateModel.Source = core.StringPtr("testString")
				notificationCreateModel.Type = core.StringPtr("testString")
				notificationCreateModel.Ibmenseverity = core.StringPtr("testString")
				notificationCreateModel.Ibmensourceid = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultshort = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultlong = core.StringPtr("testString")
				notificationCreateModel.Subject = core.StringPtr("testString")
				notificationCreateModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				notificationCreateModel.Datacontenttype = core.StringPtr("application/json")
				notificationCreateModel.Ibmenpushto = core.StringPtr(`{"platforms":["push_android"]}`)
				notificationCreateModel.Ibmenfcmbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsheaders = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromebody = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromeheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenfirefoxbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenfirefoxheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenhuaweibody = core.StringPtr("testString")
				notificationCreateModel.Ibmensafaribody = core.StringPtr("testString")
				notificationCreateModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendBulkNotificationsOptions model
				sendBulkNotificationsOptionsModel := new(eventnotificationsv1.SendBulkNotificationsOptions)
				sendBulkNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendBulkNotificationsOptionsModel.BulkMessages = []eventnotificationsv1.NotificationCreate{*notificationCreateModel}
				sendBulkNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.SendBulkNotifications(sendBulkNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SendBulkNotifications with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the NotificationCreate model
				notificationCreateModel := new(eventnotificationsv1.NotificationCreate)
				notificationCreateModel.Specversion = core.StringPtr("1.0")
				notificationCreateModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				notificationCreateModel.ID = core.StringPtr("testString")
				notificationCreateModel.Source = core.StringPtr("testString")
				notificationCreateModel.Type = core.StringPtr("testString")
				notificationCreateModel.Ibmenseverity = core.StringPtr("testString")
				notificationCreateModel.Ibmensourceid = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultshort = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultlong = core.StringPtr("testString")
				notificationCreateModel.Subject = core.StringPtr("testString")
				notificationCreateModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				notificationCreateModel.Datacontenttype = core.StringPtr("application/json")
				notificationCreateModel.Ibmenpushto = core.StringPtr(`{"platforms":["push_android"]}`)
				notificationCreateModel.Ibmenfcmbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsheaders = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromebody = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromeheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenfirefoxbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenfirefoxheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenhuaweibody = core.StringPtr("testString")
				notificationCreateModel.Ibmensafaribody = core.StringPtr("testString")
				notificationCreateModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendBulkNotificationsOptions model
				sendBulkNotificationsOptionsModel := new(eventnotificationsv1.SendBulkNotificationsOptions)
				sendBulkNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendBulkNotificationsOptionsModel.BulkMessages = []eventnotificationsv1.NotificationCreate{*notificationCreateModel}
				sendBulkNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.SendBulkNotifications(sendBulkNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SendBulkNotificationsOptions model with no property values
				sendBulkNotificationsOptionsModelNew := new(eventnotificationsv1.SendBulkNotificationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.SendBulkNotifications(sendBulkNotificationsOptionsModelNew)
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
					res.WriteHeader(202)
				}))
			})
			It(`Invoke SendBulkNotifications successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the NotificationCreate model
				notificationCreateModel := new(eventnotificationsv1.NotificationCreate)
				notificationCreateModel.Specversion = core.StringPtr("1.0")
				notificationCreateModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				notificationCreateModel.ID = core.StringPtr("testString")
				notificationCreateModel.Source = core.StringPtr("testString")
				notificationCreateModel.Type = core.StringPtr("testString")
				notificationCreateModel.Ibmenseverity = core.StringPtr("testString")
				notificationCreateModel.Ibmensourceid = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultshort = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultlong = core.StringPtr("testString")
				notificationCreateModel.Subject = core.StringPtr("testString")
				notificationCreateModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				notificationCreateModel.Datacontenttype = core.StringPtr("application/json")
				notificationCreateModel.Ibmenpushto = core.StringPtr(`{"platforms":["push_android"]}`)
				notificationCreateModel.Ibmenfcmbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsheaders = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromebody = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromeheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenfirefoxbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenfirefoxheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenhuaweibody = core.StringPtr("testString")
				notificationCreateModel.Ibmensafaribody = core.StringPtr("testString")
				notificationCreateModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendBulkNotificationsOptions model
				sendBulkNotificationsOptionsModel := new(eventnotificationsv1.SendBulkNotificationsOptions)
				sendBulkNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendBulkNotificationsOptionsModel.BulkMessages = []eventnotificationsv1.NotificationCreate{*notificationCreateModel}
				sendBulkNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.SendBulkNotifications(sendBulkNotificationsOptionsModel)
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
	Describe(`CreateSources(createSourcesOptions *CreateSourcesOptions) - Operation response error`, func() {
		createSourcesPath := "/v1/instances/testString/sources"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSourcesPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSources with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the CreateSourcesOptions model
				createSourcesOptionsModel := new(eventnotificationsv1.CreateSourcesOptions)
				createSourcesOptionsModel.InstanceID = core.StringPtr("testString")
				createSourcesOptionsModel.Name = core.StringPtr("testString")
				createSourcesOptionsModel.Description = core.StringPtr("testString")
				createSourcesOptionsModel.Enabled = core.BoolPtr(true)
				createSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.CreateSources(createSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.CreateSources(createSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSources(createSourcesOptions *CreateSourcesOptions)`, func() {
		createSourcesPath := "/v1/instances/testString/sources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSourcesPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "enabled": false, "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateSources successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the CreateSourcesOptions model
				createSourcesOptionsModel := new(eventnotificationsv1.CreateSourcesOptions)
				createSourcesOptionsModel.InstanceID = core.StringPtr("testString")
				createSourcesOptionsModel.Name = core.StringPtr("testString")
				createSourcesOptionsModel.Description = core.StringPtr("testString")
				createSourcesOptionsModel.Enabled = core.BoolPtr(true)
				createSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.CreateSourcesWithContext(ctx, createSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.CreateSources(createSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.CreateSourcesWithContext(ctx, createSourcesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createSourcesPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "enabled": false, "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateSources successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.CreateSources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSourcesOptions model
				createSourcesOptionsModel := new(eventnotificationsv1.CreateSourcesOptions)
				createSourcesOptionsModel.InstanceID = core.StringPtr("testString")
				createSourcesOptionsModel.Name = core.StringPtr("testString")
				createSourcesOptionsModel.Description = core.StringPtr("testString")
				createSourcesOptionsModel.Enabled = core.BoolPtr(true)
				createSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.CreateSources(createSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSources with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the CreateSourcesOptions model
				createSourcesOptionsModel := new(eventnotificationsv1.CreateSourcesOptions)
				createSourcesOptionsModel.InstanceID = core.StringPtr("testString")
				createSourcesOptionsModel.Name = core.StringPtr("testString")
				createSourcesOptionsModel.Description = core.StringPtr("testString")
				createSourcesOptionsModel.Enabled = core.BoolPtr(true)
				createSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.CreateSources(createSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSourcesOptions model with no property values
				createSourcesOptionsModelNew := new(eventnotificationsv1.CreateSourcesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.CreateSources(createSourcesOptionsModelNew)
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
			It(`Invoke CreateSources successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the CreateSourcesOptions model
				createSourcesOptionsModel := new(eventnotificationsv1.CreateSourcesOptions)
				createSourcesOptionsModel.InstanceID = core.StringPtr("testString")
				createSourcesOptionsModel.Name = core.StringPtr("testString")
				createSourcesOptionsModel.Description = core.StringPtr("testString")
				createSourcesOptionsModel.Enabled = core.BoolPtr(true)
				createSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.CreateSources(createSourcesOptionsModel)
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
					fmt.Fprint(res, `} this is not valid json {`)
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
				listSourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
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
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "sources": [{"id": "ID", "name": "Name", "description": "Description", "type": "Type", "enabled": false, "updated_at": "2019-01-01T12:00:00.000Z", "topic_count": 0}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
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
				listSourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
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
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "sources": [{"id": "ID", "name": "Name", "description": "Description", "type": "Type", "enabled": false, "updated_at": "2019-01-01T12:00:00.000Z", "topic_count": 0}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
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
				listSourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
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
				listSourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
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
				listSourcesOptionsModel.Limit = core.Int64Ptr(int64(10))
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
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(eventnotificationsv1.SourceList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(eventnotificationsv1.SourceList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.SourceList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.SourceList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSourcesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"sources":[{"id":"ID","name":"Name","description":"Description","type":"Type","enabled":false,"updated_at":"2019-01-01T12:00:00.000Z","topic_count":0}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"sources":[{"id":"ID","name":"Name","description":"Description","type":"Type","enabled":false,"updated_at":"2019-01-01T12:00:00.000Z","topic_count":0}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use SourcesPager.GetNext successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listSourcesOptionsModel := &eventnotificationsv1.ListSourcesOptions{
					InstanceID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewSourcesPager(listSourcesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []eventnotificationsv1.SourceListItem
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use SourcesPager.GetAll successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listSourcesOptionsModel := &eventnotificationsv1.ListSourcesOptions{
					InstanceID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewSourcesPager(listSourcesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
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
					fmt.Fprint(res, `} this is not valid json {`)
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
	Describe(`DeleteSource(deleteSourceOptions *DeleteSourceOptions)`, func() {
		deleteSourcePath := "/v1/instances/testString/sources/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSourcePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSource successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := eventNotificationsService.DeleteSource(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSourceOptions model
				deleteSourceOptionsModel := new(eventnotificationsv1.DeleteSourceOptions)
				deleteSourceOptionsModel.InstanceID = core.StringPtr("testString")
				deleteSourceOptionsModel.ID = core.StringPtr("testString")
				deleteSourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = eventNotificationsService.DeleteSource(deleteSourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSource with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DeleteSourceOptions model
				deleteSourceOptionsModel := new(eventnotificationsv1.DeleteSourceOptions)
				deleteSourceOptionsModel.InstanceID = core.StringPtr("testString")
				deleteSourceOptionsModel.ID = core.StringPtr("testString")
				deleteSourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := eventNotificationsService.DeleteSource(deleteSourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSourceOptions model with no property values
				deleteSourceOptionsModelNew := new(eventnotificationsv1.DeleteSourceOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = eventNotificationsService.DeleteSource(deleteSourceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSource(updateSourceOptions *UpdateSourceOptions) - Operation response error`, func() {
		updateSourcePath := "/v1/instances/testString/sources/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSourcePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSource with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the UpdateSourceOptions model
				updateSourceOptionsModel := new(eventnotificationsv1.UpdateSourceOptions)
				updateSourceOptionsModel.InstanceID = core.StringPtr("testString")
				updateSourceOptionsModel.ID = core.StringPtr("testString")
				updateSourceOptionsModel.Name = core.StringPtr("testString")
				updateSourceOptionsModel.Description = core.StringPtr("testString")
				updateSourceOptionsModel.Enabled = core.BoolPtr(true)
				updateSourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.UpdateSource(updateSourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.UpdateSource(updateSourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSource(updateSourceOptions *UpdateSourceOptions)`, func() {
		updateSourcePath := "/v1/instances/testString/sources/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSourcePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "enabled": false, "type": "Type", "updated_at": "2019-01-01T12:00:00.000Z", "topic_count": 10, "topic_names": ["TopicNames"]}`)
				}))
			})
			It(`Invoke UpdateSource successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateSourceOptions model
				updateSourceOptionsModel := new(eventnotificationsv1.UpdateSourceOptions)
				updateSourceOptionsModel.InstanceID = core.StringPtr("testString")
				updateSourceOptionsModel.ID = core.StringPtr("testString")
				updateSourceOptionsModel.Name = core.StringPtr("testString")
				updateSourceOptionsModel.Description = core.StringPtr("testString")
				updateSourceOptionsModel.Enabled = core.BoolPtr(true)
				updateSourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.UpdateSourceWithContext(ctx, updateSourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.UpdateSource(updateSourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.UpdateSourceWithContext(ctx, updateSourceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateSourcePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "enabled": false, "type": "Type", "updated_at": "2019-01-01T12:00:00.000Z", "topic_count": 10, "topic_names": ["TopicNames"]}`)
				}))
			})
			It(`Invoke UpdateSource successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.UpdateSource(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateSourceOptions model
				updateSourceOptionsModel := new(eventnotificationsv1.UpdateSourceOptions)
				updateSourceOptionsModel.InstanceID = core.StringPtr("testString")
				updateSourceOptionsModel.ID = core.StringPtr("testString")
				updateSourceOptionsModel.Name = core.StringPtr("testString")
				updateSourceOptionsModel.Description = core.StringPtr("testString")
				updateSourceOptionsModel.Enabled = core.BoolPtr(true)
				updateSourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.UpdateSource(updateSourceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSource with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the UpdateSourceOptions model
				updateSourceOptionsModel := new(eventnotificationsv1.UpdateSourceOptions)
				updateSourceOptionsModel.InstanceID = core.StringPtr("testString")
				updateSourceOptionsModel.ID = core.StringPtr("testString")
				updateSourceOptionsModel.Name = core.StringPtr("testString")
				updateSourceOptionsModel.Description = core.StringPtr("testString")
				updateSourceOptionsModel.Enabled = core.BoolPtr(true)
				updateSourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.UpdateSource(updateSourceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSourceOptions model with no property values
				updateSourceOptionsModelNew := new(eventnotificationsv1.UpdateSourceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.UpdateSource(updateSourceOptionsModelNew)
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
			It(`Invoke UpdateSource successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the UpdateSourceOptions model
				updateSourceOptionsModel := new(eventnotificationsv1.UpdateSourceOptions)
				updateSourceOptionsModel.InstanceID = core.StringPtr("testString")
				updateSourceOptionsModel.ID = core.StringPtr("testString")
				updateSourceOptionsModel.Name = core.StringPtr("testString")
				updateSourceOptionsModel.Description = core.StringPtr("testString")
				updateSourceOptionsModel.Enabled = core.BoolPtr(true)
				updateSourceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.UpdateSource(updateSourceOptionsModel)
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
					fmt.Fprint(res, `} this is not valid json {`)
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

				// Construct an instance of the SourcesItems model
				sourcesItemsModel := new(eventnotificationsv1.SourcesItems)
				sourcesItemsModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				sourcesItemsModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the CreateTopicOptions model
				createTopicOptionsModel := new(eventnotificationsv1.CreateTopicOptions)
				createTopicOptionsModel.InstanceID = core.StringPtr("testString")
				createTopicOptionsModel.Name = core.StringPtr("testString")
				createTopicOptionsModel.Description = core.StringPtr("testString")
				createTopicOptionsModel.Sources = []eventnotificationsv1.SourcesItems{*sourcesItemsModel}
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

				// Construct an instance of the SourcesItems model
				sourcesItemsModel := new(eventnotificationsv1.SourcesItems)
				sourcesItemsModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				sourcesItemsModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the CreateTopicOptions model
				createTopicOptionsModel := new(eventnotificationsv1.CreateTopicOptions)
				createTopicOptionsModel.InstanceID = core.StringPtr("testString")
				createTopicOptionsModel.Name = core.StringPtr("testString")
				createTopicOptionsModel.Description = core.StringPtr("testString")
				createTopicOptionsModel.Sources = []eventnotificationsv1.SourcesItems{*sourcesItemsModel}
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

				// Construct an instance of the SourcesItems model
				sourcesItemsModel := new(eventnotificationsv1.SourcesItems)
				sourcesItemsModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				sourcesItemsModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the CreateTopicOptions model
				createTopicOptionsModel := new(eventnotificationsv1.CreateTopicOptions)
				createTopicOptionsModel.InstanceID = core.StringPtr("testString")
				createTopicOptionsModel.Name = core.StringPtr("testString")
				createTopicOptionsModel.Description = core.StringPtr("testString")
				createTopicOptionsModel.Sources = []eventnotificationsv1.SourcesItems{*sourcesItemsModel}
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

				// Construct an instance of the SourcesItems model
				sourcesItemsModel := new(eventnotificationsv1.SourcesItems)
				sourcesItemsModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				sourcesItemsModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the CreateTopicOptions model
				createTopicOptionsModel := new(eventnotificationsv1.CreateTopicOptions)
				createTopicOptionsModel.InstanceID = core.StringPtr("testString")
				createTopicOptionsModel.Name = core.StringPtr("testString")
				createTopicOptionsModel.Description = core.StringPtr("testString")
				createTopicOptionsModel.Sources = []eventnotificationsv1.SourcesItems{*sourcesItemsModel}
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

				// Construct an instance of the SourcesItems model
				sourcesItemsModel := new(eventnotificationsv1.SourcesItems)
				sourcesItemsModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				sourcesItemsModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the CreateTopicOptions model
				createTopicOptionsModel := new(eventnotificationsv1.CreateTopicOptions)
				createTopicOptionsModel.InstanceID = core.StringPtr("testString")
				createTopicOptionsModel.Name = core.StringPtr("testString")
				createTopicOptionsModel.Description = core.StringPtr("testString")
				createTopicOptionsModel.Sources = []eventnotificationsv1.SourcesItems{*sourcesItemsModel}
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
					fmt.Fprint(res, `} this is not valid json {`)
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
				listTopicsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "topics": [{"id": "ID", "name": "Name", "description": "Description", "source_count": 0, "sources_names": ["SourcesNames"], "subscription_count": 0}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
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
				listTopicsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "topics": [{"id": "ID", "name": "Name", "description": "Description", "source_count": 0, "sources_names": ["SourcesNames"], "subscription_count": 0}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
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
				listTopicsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
				listTopicsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
				listTopicsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(eventnotificationsv1.TopicList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(eventnotificationsv1.TopicList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.TopicList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.TopicList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTopicsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"topics":[{"id":"ID","name":"Name","description":"Description","source_count":0,"sources_names":["SourcesNames"],"subscription_count":0}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"topics":[{"id":"ID","name":"Name","description":"Description","source_count":0,"sources_names":["SourcesNames"],"subscription_count":0}],"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use TopicsPager.GetNext successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listTopicsOptionsModel := &eventnotificationsv1.ListTopicsOptions{
					InstanceID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewTopicsPager(listTopicsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []eventnotificationsv1.TopicsListItem
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use TopicsPager.GetAll successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listTopicsOptionsModel := &eventnotificationsv1.ListTopicsOptions{
					InstanceID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewTopicsPager(listTopicsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
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
					fmt.Fprint(res, `} this is not valid json {`)
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
					fmt.Fprint(res, `} this is not valid json {`)
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

				// Construct an instance of the SourcesItems model
				sourcesItemsModel := new(eventnotificationsv1.SourcesItems)
				sourcesItemsModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				sourcesItemsModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the ReplaceTopicOptions model
				replaceTopicOptionsModel := new(eventnotificationsv1.ReplaceTopicOptions)
				replaceTopicOptionsModel.InstanceID = core.StringPtr("testString")
				replaceTopicOptionsModel.ID = core.StringPtr("testString")
				replaceTopicOptionsModel.Name = core.StringPtr("testString")
				replaceTopicOptionsModel.Description = core.StringPtr("testString")
				replaceTopicOptionsModel.Sources = []eventnotificationsv1.SourcesItems{*sourcesItemsModel}
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

				// Construct an instance of the SourcesItems model
				sourcesItemsModel := new(eventnotificationsv1.SourcesItems)
				sourcesItemsModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				sourcesItemsModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the ReplaceTopicOptions model
				replaceTopicOptionsModel := new(eventnotificationsv1.ReplaceTopicOptions)
				replaceTopicOptionsModel.InstanceID = core.StringPtr("testString")
				replaceTopicOptionsModel.ID = core.StringPtr("testString")
				replaceTopicOptionsModel.Name = core.StringPtr("testString")
				replaceTopicOptionsModel.Description = core.StringPtr("testString")
				replaceTopicOptionsModel.Sources = []eventnotificationsv1.SourcesItems{*sourcesItemsModel}
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

				// Construct an instance of the SourcesItems model
				sourcesItemsModel := new(eventnotificationsv1.SourcesItems)
				sourcesItemsModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				sourcesItemsModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the ReplaceTopicOptions model
				replaceTopicOptionsModel := new(eventnotificationsv1.ReplaceTopicOptions)
				replaceTopicOptionsModel.InstanceID = core.StringPtr("testString")
				replaceTopicOptionsModel.ID = core.StringPtr("testString")
				replaceTopicOptionsModel.Name = core.StringPtr("testString")
				replaceTopicOptionsModel.Description = core.StringPtr("testString")
				replaceTopicOptionsModel.Sources = []eventnotificationsv1.SourcesItems{*sourcesItemsModel}
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

				// Construct an instance of the SourcesItems model
				sourcesItemsModel := new(eventnotificationsv1.SourcesItems)
				sourcesItemsModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				sourcesItemsModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the ReplaceTopicOptions model
				replaceTopicOptionsModel := new(eventnotificationsv1.ReplaceTopicOptions)
				replaceTopicOptionsModel.InstanceID = core.StringPtr("testString")
				replaceTopicOptionsModel.ID = core.StringPtr("testString")
				replaceTopicOptionsModel.Name = core.StringPtr("testString")
				replaceTopicOptionsModel.Description = core.StringPtr("testString")
				replaceTopicOptionsModel.Sources = []eventnotificationsv1.SourcesItems{*sourcesItemsModel}
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

				// Construct an instance of the SourcesItems model
				sourcesItemsModel := new(eventnotificationsv1.SourcesItems)
				sourcesItemsModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				sourcesItemsModel.Rules = []eventnotificationsv1.Rules{*rulesModel}

				// Construct an instance of the ReplaceTopicOptions model
				replaceTopicOptionsModel := new(eventnotificationsv1.ReplaceTopicOptions)
				replaceTopicOptionsModel.InstanceID = core.StringPtr("testString")
				replaceTopicOptionsModel.ID = core.StringPtr("testString")
				replaceTopicOptionsModel.Name = core.StringPtr("testString")
				replaceTopicOptionsModel.Description = core.StringPtr("testString")
				replaceTopicOptionsModel.Sources = []eventnotificationsv1.SourcesItems{*sourcesItemsModel}
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
	Describe(`CreateTemplate(createTemplateOptions *CreateTemplateOptions) - Operation response error`, func() {
		createTemplatePath := "/v1/instances/testString/templates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTemplatePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTemplate with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the TemplateConfig model
				templateConfigModel := new(eventnotificationsv1.TemplateConfig)
				templateConfigModel.Body = core.StringPtr("testString")
				templateConfigModel.Subject = core.StringPtr("testString")

				// Construct an instance of the CreateTemplateOptions model
				createTemplateOptionsModel := new(eventnotificationsv1.CreateTemplateOptions)
				createTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				createTemplateOptionsModel.Name = core.StringPtr("testString")
				createTemplateOptionsModel.Type = core.StringPtr("smtp_custom.notification")
				createTemplateOptionsModel.Params = templateConfigModel
				createTemplateOptionsModel.Description = core.StringPtr("testString")
				createTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.CreateTemplate(createTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.CreateTemplate(createTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTemplate(createTemplateOptions *CreateTemplateOptions)`, func() {
		createTemplatePath := "/v1/instances/testString/templates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTemplatePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "smtp_custom.notification", "params": {"body": "Body", "subject": "Subject"}, "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateTemplate successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the TemplateConfig model
				templateConfigModel := new(eventnotificationsv1.TemplateConfig)
				templateConfigModel.Body = core.StringPtr("testString")
				templateConfigModel.Subject = core.StringPtr("testString")

				// Construct an instance of the CreateTemplateOptions model
				createTemplateOptionsModel := new(eventnotificationsv1.CreateTemplateOptions)
				createTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				createTemplateOptionsModel.Name = core.StringPtr("testString")
				createTemplateOptionsModel.Type = core.StringPtr("smtp_custom.notification")
				createTemplateOptionsModel.Params = templateConfigModel
				createTemplateOptionsModel.Description = core.StringPtr("testString")
				createTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.CreateTemplateWithContext(ctx, createTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.CreateTemplate(createTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.CreateTemplateWithContext(ctx, createTemplateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTemplatePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "smtp_custom.notification", "params": {"body": "Body", "subject": "Subject"}, "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateTemplate successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.CreateTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TemplateConfig model
				templateConfigModel := new(eventnotificationsv1.TemplateConfig)
				templateConfigModel.Body = core.StringPtr("testString")
				templateConfigModel.Subject = core.StringPtr("testString")

				// Construct an instance of the CreateTemplateOptions model
				createTemplateOptionsModel := new(eventnotificationsv1.CreateTemplateOptions)
				createTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				createTemplateOptionsModel.Name = core.StringPtr("testString")
				createTemplateOptionsModel.Type = core.StringPtr("smtp_custom.notification")
				createTemplateOptionsModel.Params = templateConfigModel
				createTemplateOptionsModel.Description = core.StringPtr("testString")
				createTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.CreateTemplate(createTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTemplate with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the TemplateConfig model
				templateConfigModel := new(eventnotificationsv1.TemplateConfig)
				templateConfigModel.Body = core.StringPtr("testString")
				templateConfigModel.Subject = core.StringPtr("testString")

				// Construct an instance of the CreateTemplateOptions model
				createTemplateOptionsModel := new(eventnotificationsv1.CreateTemplateOptions)
				createTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				createTemplateOptionsModel.Name = core.StringPtr("testString")
				createTemplateOptionsModel.Type = core.StringPtr("smtp_custom.notification")
				createTemplateOptionsModel.Params = templateConfigModel
				createTemplateOptionsModel.Description = core.StringPtr("testString")
				createTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.CreateTemplate(createTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTemplateOptions model with no property values
				createTemplateOptionsModelNew := new(eventnotificationsv1.CreateTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.CreateTemplate(createTemplateOptionsModelNew)
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
			It(`Invoke CreateTemplate successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the TemplateConfig model
				templateConfigModel := new(eventnotificationsv1.TemplateConfig)
				templateConfigModel.Body = core.StringPtr("testString")
				templateConfigModel.Subject = core.StringPtr("testString")

				// Construct an instance of the CreateTemplateOptions model
				createTemplateOptionsModel := new(eventnotificationsv1.CreateTemplateOptions)
				createTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				createTemplateOptionsModel.Name = core.StringPtr("testString")
				createTemplateOptionsModel.Type = core.StringPtr("smtp_custom.notification")
				createTemplateOptionsModel.Params = templateConfigModel
				createTemplateOptionsModel.Description = core.StringPtr("testString")
				createTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.CreateTemplate(createTemplateOptionsModel)
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
	Describe(`ListTemplates(listTemplatesOptions *ListTemplatesOptions) - Operation response error`, func() {
		listTemplatesPath := "/v1/instances/testString/templates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTemplatesPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTemplates with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListTemplatesOptions model
				listTemplatesOptionsModel := new(eventnotificationsv1.ListTemplatesOptions)
				listTemplatesOptionsModel.InstanceID = core.StringPtr("testString")
				listTemplatesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplatesOptionsModel.Search = core.StringPtr("testString")
				listTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.ListTemplates(listTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.ListTemplates(listTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTemplates(listTemplatesOptions *ListTemplatesOptions)`, func() {
		listTemplatesPath := "/v1/instances/testString/templates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "templates": [{"id": "ID", "name": "Name", "description": "Description", "type": "smtp_custom.notification", "subscription_count": 17, "subscription_names": ["SubscriptionNames"], "updated_at": "2019-01-01T12:00:00.000Z"}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListTemplates successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the ListTemplatesOptions model
				listTemplatesOptionsModel := new(eventnotificationsv1.ListTemplatesOptions)
				listTemplatesOptionsModel.InstanceID = core.StringPtr("testString")
				listTemplatesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplatesOptionsModel.Search = core.StringPtr("testString")
				listTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.ListTemplatesWithContext(ctx, listTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.ListTemplates(listTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.ListTemplatesWithContext(ctx, listTemplatesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "templates": [{"id": "ID", "name": "Name", "description": "Description", "type": "smtp_custom.notification", "subscription_count": 17, "subscription_names": ["SubscriptionNames"], "updated_at": "2019-01-01T12:00:00.000Z"}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListTemplates successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.ListTemplates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTemplatesOptions model
				listTemplatesOptionsModel := new(eventnotificationsv1.ListTemplatesOptions)
				listTemplatesOptionsModel.InstanceID = core.StringPtr("testString")
				listTemplatesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplatesOptionsModel.Search = core.StringPtr("testString")
				listTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.ListTemplates(listTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTemplates with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListTemplatesOptions model
				listTemplatesOptionsModel := new(eventnotificationsv1.ListTemplatesOptions)
				listTemplatesOptionsModel.InstanceID = core.StringPtr("testString")
				listTemplatesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplatesOptionsModel.Search = core.StringPtr("testString")
				listTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.ListTemplates(listTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTemplatesOptions model with no property values
				listTemplatesOptionsModelNew := new(eventnotificationsv1.ListTemplatesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.ListTemplates(listTemplatesOptionsModelNew)
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
			It(`Invoke ListTemplates successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListTemplatesOptions model
				listTemplatesOptionsModel := new(eventnotificationsv1.ListTemplatesOptions)
				listTemplatesOptionsModel.InstanceID = core.StringPtr("testString")
				listTemplatesOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTemplatesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTemplatesOptionsModel.Search = core.StringPtr("testString")
				listTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.ListTemplates(listTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(eventnotificationsv1.TemplateList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(eventnotificationsv1.TemplateList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.TemplateList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.TemplateList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"templates":[{"id":"ID","name":"Name","description":"Description","type":"smtp_custom.notification","subscription_count":17,"subscription_names":["SubscriptionNames"],"updated_at":"2019-01-01T12:00:00.000Z"}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"templates":[{"id":"ID","name":"Name","description":"Description","type":"smtp_custom.notification","subscription_count":17,"subscription_names":["SubscriptionNames"],"updated_at":"2019-01-01T12:00:00.000Z"}],"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use TemplatesPager.GetNext successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listTemplatesOptionsModel := &eventnotificationsv1.ListTemplatesOptions{
					InstanceID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewTemplatesPager(listTemplatesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []eventnotificationsv1.Template
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use TemplatesPager.GetAll successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listTemplatesOptionsModel := &eventnotificationsv1.ListTemplatesOptions{
					InstanceID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewTemplatesPager(listTemplatesOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetTemplate(getTemplateOptions *GetTemplateOptions) - Operation response error`, func() {
		getTemplatePath := "/v1/instances/testString/templates/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTemplatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTemplate with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetTemplateOptions model
				getTemplateOptionsModel := new(eventnotificationsv1.GetTemplateOptions)
				getTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				getTemplateOptionsModel.ID = core.StringPtr("testString")
				getTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.GetTemplate(getTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.GetTemplate(getTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTemplate(getTemplateOptions *GetTemplateOptions)`, func() {
		getTemplatePath := "/v1/instances/testString/templates/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTemplatePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "smtp_custom.notification", "subscription_count": 17, "subscription_names": ["SubscriptionNames"], "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetTemplate successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the GetTemplateOptions model
				getTemplateOptionsModel := new(eventnotificationsv1.GetTemplateOptions)
				getTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				getTemplateOptionsModel.ID = core.StringPtr("testString")
				getTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.GetTemplateWithContext(ctx, getTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.GetTemplate(getTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.GetTemplateWithContext(ctx, getTemplateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTemplatePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "smtp_custom.notification", "subscription_count": 17, "subscription_names": ["SubscriptionNames"], "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetTemplate successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.GetTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTemplateOptions model
				getTemplateOptionsModel := new(eventnotificationsv1.GetTemplateOptions)
				getTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				getTemplateOptionsModel.ID = core.StringPtr("testString")
				getTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.GetTemplate(getTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTemplate with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetTemplateOptions model
				getTemplateOptionsModel := new(eventnotificationsv1.GetTemplateOptions)
				getTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				getTemplateOptionsModel.ID = core.StringPtr("testString")
				getTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.GetTemplate(getTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTemplateOptions model with no property values
				getTemplateOptionsModelNew := new(eventnotificationsv1.GetTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.GetTemplate(getTemplateOptionsModelNew)
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
			It(`Invoke GetTemplate successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetTemplateOptions model
				getTemplateOptionsModel := new(eventnotificationsv1.GetTemplateOptions)
				getTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				getTemplateOptionsModel.ID = core.StringPtr("testString")
				getTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.GetTemplate(getTemplateOptionsModel)
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
	Describe(`UpdateTemplate(updateTemplateOptions *UpdateTemplateOptions) - Operation response error`, func() {
		updateTemplatePath := "/v1/instances/testString/templates/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTemplatePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTemplate with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the TemplateConfig model
				templateConfigModel := new(eventnotificationsv1.TemplateConfig)
				templateConfigModel.Body = core.StringPtr("testString")
				templateConfigModel.Subject = core.StringPtr("testString")

				// Construct an instance of the UpdateTemplateOptions model
				updateTemplateOptionsModel := new(eventnotificationsv1.UpdateTemplateOptions)
				updateTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				updateTemplateOptionsModel.ID = core.StringPtr("testString")
				updateTemplateOptionsModel.Name = core.StringPtr("testString")
				updateTemplateOptionsModel.Description = core.StringPtr("testString")
				updateTemplateOptionsModel.Type = core.StringPtr("smtp_custom.notification")
				updateTemplateOptionsModel.Params = templateConfigModel
				updateTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.UpdateTemplate(updateTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.UpdateTemplate(updateTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTemplate(updateTemplateOptions *UpdateTemplateOptions)`, func() {
		updateTemplatePath := "/v1/instances/testString/templates/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTemplatePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "smtp_custom.notification", "subscription_count": 17, "subscription_names": ["SubscriptionNames"], "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateTemplate successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the TemplateConfig model
				templateConfigModel := new(eventnotificationsv1.TemplateConfig)
				templateConfigModel.Body = core.StringPtr("testString")
				templateConfigModel.Subject = core.StringPtr("testString")

				// Construct an instance of the UpdateTemplateOptions model
				updateTemplateOptionsModel := new(eventnotificationsv1.UpdateTemplateOptions)
				updateTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				updateTemplateOptionsModel.ID = core.StringPtr("testString")
				updateTemplateOptionsModel.Name = core.StringPtr("testString")
				updateTemplateOptionsModel.Description = core.StringPtr("testString")
				updateTemplateOptionsModel.Type = core.StringPtr("smtp_custom.notification")
				updateTemplateOptionsModel.Params = templateConfigModel
				updateTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.UpdateTemplateWithContext(ctx, updateTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.UpdateTemplate(updateTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.UpdateTemplateWithContext(ctx, updateTemplateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateTemplatePath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "smtp_custom.notification", "subscription_count": 17, "subscription_names": ["SubscriptionNames"], "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateTemplate successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.UpdateTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TemplateConfig model
				templateConfigModel := new(eventnotificationsv1.TemplateConfig)
				templateConfigModel.Body = core.StringPtr("testString")
				templateConfigModel.Subject = core.StringPtr("testString")

				// Construct an instance of the UpdateTemplateOptions model
				updateTemplateOptionsModel := new(eventnotificationsv1.UpdateTemplateOptions)
				updateTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				updateTemplateOptionsModel.ID = core.StringPtr("testString")
				updateTemplateOptionsModel.Name = core.StringPtr("testString")
				updateTemplateOptionsModel.Description = core.StringPtr("testString")
				updateTemplateOptionsModel.Type = core.StringPtr("smtp_custom.notification")
				updateTemplateOptionsModel.Params = templateConfigModel
				updateTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.UpdateTemplate(updateTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTemplate with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the TemplateConfig model
				templateConfigModel := new(eventnotificationsv1.TemplateConfig)
				templateConfigModel.Body = core.StringPtr("testString")
				templateConfigModel.Subject = core.StringPtr("testString")

				// Construct an instance of the UpdateTemplateOptions model
				updateTemplateOptionsModel := new(eventnotificationsv1.UpdateTemplateOptions)
				updateTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				updateTemplateOptionsModel.ID = core.StringPtr("testString")
				updateTemplateOptionsModel.Name = core.StringPtr("testString")
				updateTemplateOptionsModel.Description = core.StringPtr("testString")
				updateTemplateOptionsModel.Type = core.StringPtr("smtp_custom.notification")
				updateTemplateOptionsModel.Params = templateConfigModel
				updateTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.UpdateTemplate(updateTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTemplateOptions model with no property values
				updateTemplateOptionsModelNew := new(eventnotificationsv1.UpdateTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.UpdateTemplate(updateTemplateOptionsModelNew)
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
			It(`Invoke UpdateTemplate successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the TemplateConfig model
				templateConfigModel := new(eventnotificationsv1.TemplateConfig)
				templateConfigModel.Body = core.StringPtr("testString")
				templateConfigModel.Subject = core.StringPtr("testString")

				// Construct an instance of the UpdateTemplateOptions model
				updateTemplateOptionsModel := new(eventnotificationsv1.UpdateTemplateOptions)
				updateTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				updateTemplateOptionsModel.ID = core.StringPtr("testString")
				updateTemplateOptionsModel.Name = core.StringPtr("testString")
				updateTemplateOptionsModel.Description = core.StringPtr("testString")
				updateTemplateOptionsModel.Type = core.StringPtr("smtp_custom.notification")
				updateTemplateOptionsModel.Params = templateConfigModel
				updateTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.UpdateTemplate(updateTemplateOptionsModel)
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
	Describe(`DeleteTemplate(deleteTemplateOptions *DeleteTemplateOptions)`, func() {
		deleteTemplatePath := "/v1/instances/testString/templates/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTemplatePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTemplate successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := eventNotificationsService.DeleteTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTemplateOptions model
				deleteTemplateOptionsModel := new(eventnotificationsv1.DeleteTemplateOptions)
				deleteTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				deleteTemplateOptionsModel.ID = core.StringPtr("testString")
				deleteTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = eventNotificationsService.DeleteTemplate(deleteTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTemplate with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DeleteTemplateOptions model
				deleteTemplateOptionsModel := new(eventnotificationsv1.DeleteTemplateOptions)
				deleteTemplateOptionsModel.InstanceID = core.StringPtr("testString")
				deleteTemplateOptionsModel.ID = core.StringPtr("testString")
				deleteTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := eventNotificationsService.DeleteTemplate(deleteTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTemplateOptions model with no property values
				deleteTemplateOptionsModelNew := new(eventnotificationsv1.DeleteTemplateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = eventNotificationsService.DeleteTemplate(deleteTemplateOptionsModelNew)
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
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDestination with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DkimAttributes model
				dkimAttributesModel := new(eventnotificationsv1.DkimAttributes)
				dkimAttributesModel.PublicKey = core.StringPtr("testString")
				dkimAttributesModel.Selector = core.StringPtr("testString")
				dkimAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the SpfAttributes model
				spfAttributesModel := new(eventnotificationsv1.SpfAttributes)
				spfAttributesModel.TxtName = core.StringPtr("testString")
				spfAttributesModel.TxtValue = core.StringPtr("testString")
				spfAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the DestinationConfigOneOfCustomDomainEmailDestinationConfig model
				destinationConfigOneOfModel := new(eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig)
				destinationConfigOneOfModel.Domain = core.StringPtr("testString")
				destinationConfigOneOfModel.Dkim = dkimAttributesModel
				destinationConfigOneOfModel.Spf = spfAttributesModel

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigOneOfModel

				// Construct an instance of the CreateDestinationOptions model
				createDestinationOptionsModel := new(eventnotificationsv1.CreateDestinationOptions)
				createDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				createDestinationOptionsModel.Name = core.StringPtr("testString")
				createDestinationOptionsModel.Type = core.StringPtr("webhook")
				createDestinationOptionsModel.Description = core.StringPtr("testString")
				createDestinationOptionsModel.Config = destinationConfigModel
				createDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "config": {"params": {"domain": "Domain", "dkim": {"public_key": "PublicKey", "selector": "Selector", "verification": "Verification"}, "spf": {"txt_name": "TxtName", "txt_value": "TxtValue", "verification": "Verification"}}}, "created_at": "2019-01-01T12:00:00.000Z"}`)
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

				// Construct an instance of the DkimAttributes model
				dkimAttributesModel := new(eventnotificationsv1.DkimAttributes)
				dkimAttributesModel.PublicKey = core.StringPtr("testString")
				dkimAttributesModel.Selector = core.StringPtr("testString")
				dkimAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the SpfAttributes model
				spfAttributesModel := new(eventnotificationsv1.SpfAttributes)
				spfAttributesModel.TxtName = core.StringPtr("testString")
				spfAttributesModel.TxtValue = core.StringPtr("testString")
				spfAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the DestinationConfigOneOfCustomDomainEmailDestinationConfig model
				destinationConfigOneOfModel := new(eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig)
				destinationConfigOneOfModel.Domain = core.StringPtr("testString")
				destinationConfigOneOfModel.Dkim = dkimAttributesModel
				destinationConfigOneOfModel.Spf = spfAttributesModel

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigOneOfModel

				// Construct an instance of the CreateDestinationOptions model
				createDestinationOptionsModel := new(eventnotificationsv1.CreateDestinationOptions)
				createDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				createDestinationOptionsModel.Name = core.StringPtr("testString")
				createDestinationOptionsModel.Type = core.StringPtr("webhook")
				createDestinationOptionsModel.Description = core.StringPtr("testString")
				createDestinationOptionsModel.Config = destinationConfigModel
				createDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "config": {"params": {"domain": "Domain", "dkim": {"public_key": "PublicKey", "selector": "Selector", "verification": "Verification"}, "spf": {"txt_name": "TxtName", "txt_value": "TxtValue", "verification": "Verification"}}}, "created_at": "2019-01-01T12:00:00.000Z"}`)
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

				// Construct an instance of the DkimAttributes model
				dkimAttributesModel := new(eventnotificationsv1.DkimAttributes)
				dkimAttributesModel.PublicKey = core.StringPtr("testString")
				dkimAttributesModel.Selector = core.StringPtr("testString")
				dkimAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the SpfAttributes model
				spfAttributesModel := new(eventnotificationsv1.SpfAttributes)
				spfAttributesModel.TxtName = core.StringPtr("testString")
				spfAttributesModel.TxtValue = core.StringPtr("testString")
				spfAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the DestinationConfigOneOfCustomDomainEmailDestinationConfig model
				destinationConfigOneOfModel := new(eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig)
				destinationConfigOneOfModel.Domain = core.StringPtr("testString")
				destinationConfigOneOfModel.Dkim = dkimAttributesModel
				destinationConfigOneOfModel.Spf = spfAttributesModel

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigOneOfModel

				// Construct an instance of the CreateDestinationOptions model
				createDestinationOptionsModel := new(eventnotificationsv1.CreateDestinationOptions)
				createDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				createDestinationOptionsModel.Name = core.StringPtr("testString")
				createDestinationOptionsModel.Type = core.StringPtr("webhook")
				createDestinationOptionsModel.Description = core.StringPtr("testString")
				createDestinationOptionsModel.Config = destinationConfigModel
				createDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
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

				// Construct an instance of the DkimAttributes model
				dkimAttributesModel := new(eventnotificationsv1.DkimAttributes)
				dkimAttributesModel.PublicKey = core.StringPtr("testString")
				dkimAttributesModel.Selector = core.StringPtr("testString")
				dkimAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the SpfAttributes model
				spfAttributesModel := new(eventnotificationsv1.SpfAttributes)
				spfAttributesModel.TxtName = core.StringPtr("testString")
				spfAttributesModel.TxtValue = core.StringPtr("testString")
				spfAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the DestinationConfigOneOfCustomDomainEmailDestinationConfig model
				destinationConfigOneOfModel := new(eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig)
				destinationConfigOneOfModel.Domain = core.StringPtr("testString")
				destinationConfigOneOfModel.Dkim = dkimAttributesModel
				destinationConfigOneOfModel.Spf = spfAttributesModel

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigOneOfModel

				// Construct an instance of the CreateDestinationOptions model
				createDestinationOptionsModel := new(eventnotificationsv1.CreateDestinationOptions)
				createDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				createDestinationOptionsModel.Name = core.StringPtr("testString")
				createDestinationOptionsModel.Type = core.StringPtr("webhook")
				createDestinationOptionsModel.Description = core.StringPtr("testString")
				createDestinationOptionsModel.Config = destinationConfigModel
				createDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
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

				// Construct an instance of the DkimAttributes model
				dkimAttributesModel := new(eventnotificationsv1.DkimAttributes)
				dkimAttributesModel.PublicKey = core.StringPtr("testString")
				dkimAttributesModel.Selector = core.StringPtr("testString")
				dkimAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the SpfAttributes model
				spfAttributesModel := new(eventnotificationsv1.SpfAttributes)
				spfAttributesModel.TxtName = core.StringPtr("testString")
				spfAttributesModel.TxtValue = core.StringPtr("testString")
				spfAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the DestinationConfigOneOfCustomDomainEmailDestinationConfig model
				destinationConfigOneOfModel := new(eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig)
				destinationConfigOneOfModel.Domain = core.StringPtr("testString")
				destinationConfigOneOfModel.Dkim = dkimAttributesModel
				destinationConfigOneOfModel.Spf = spfAttributesModel

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigOneOfModel

				// Construct an instance of the CreateDestinationOptions model
				createDestinationOptionsModel := new(eventnotificationsv1.CreateDestinationOptions)
				createDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				createDestinationOptionsModel.Name = core.StringPtr("testString")
				createDestinationOptionsModel.Type = core.StringPtr("webhook")
				createDestinationOptionsModel.Description = core.StringPtr("testString")
				createDestinationOptionsModel.Config = destinationConfigModel
				createDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				createDestinationOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
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
					fmt.Fprint(res, `} this is not valid json {`)
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
				listDestinationsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "destinations": [{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "subscription_count": 17, "subscription_names": ["SubscriptionNames"], "updated_at": "2019-01-01T12:00:00.000Z"}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
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
				listDestinationsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "destinations": [{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "subscription_count": 17, "subscription_names": ["SubscriptionNames"], "updated_at": "2019-01-01T12:00:00.000Z"}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
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
				listDestinationsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
				listDestinationsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
				listDestinationsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(eventnotificationsv1.DestinationList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(eventnotificationsv1.DestinationList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.DestinationList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.DestinationList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDestinationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"destinations":[{"id":"ID","name":"Name","description":"Description","type":"webhook","subscription_count":17,"subscription_names":["SubscriptionNames"],"updated_at":"2019-01-01T12:00:00.000Z"}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"destinations":[{"id":"ID","name":"Name","description":"Description","type":"webhook","subscription_count":17,"subscription_names":["SubscriptionNames"],"updated_at":"2019-01-01T12:00:00.000Z"}],"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use DestinationsPager.GetNext successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listDestinationsOptionsModel := &eventnotificationsv1.ListDestinationsOptions{
					InstanceID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewDestinationsPager(listDestinationsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []eventnotificationsv1.DestinationListItem
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use DestinationsPager.GetAll successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listDestinationsOptionsModel := &eventnotificationsv1.ListDestinationsOptions{
					InstanceID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewDestinationsPager(listDestinationsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
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
					fmt.Fprint(res, `} this is not valid json {`)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "config": {"params": {"domain": "Domain", "dkim": {"public_key": "PublicKey", "selector": "Selector", "verification": "Verification"}, "spf": {"txt_name": "TxtName", "txt_value": "TxtValue", "verification": "Verification"}}}, "updated_at": "2019-01-01T12:00:00.000Z", "subscription_count": 0, "subscription_names": ["SubscriptionNames"]}`)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "config": {"params": {"domain": "Domain", "dkim": {"public_key": "PublicKey", "selector": "Selector", "verification": "Verification"}, "spf": {"txt_name": "TxtName", "txt_value": "TxtValue", "verification": "Verification"}}}, "updated_at": "2019-01-01T12:00:00.000Z", "subscription_count": 0, "subscription_names": ["SubscriptionNames"]}`)
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
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDestination with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DkimAttributes model
				dkimAttributesModel := new(eventnotificationsv1.DkimAttributes)
				dkimAttributesModel.PublicKey = core.StringPtr("testString")
				dkimAttributesModel.Selector = core.StringPtr("testString")
				dkimAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the SpfAttributes model
				spfAttributesModel := new(eventnotificationsv1.SpfAttributes)
				spfAttributesModel.TxtName = core.StringPtr("testString")
				spfAttributesModel.TxtValue = core.StringPtr("testString")
				spfAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the DestinationConfigOneOfCustomDomainEmailDestinationConfig model
				destinationConfigOneOfModel := new(eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig)
				destinationConfigOneOfModel.Domain = core.StringPtr("testString")
				destinationConfigOneOfModel.Dkim = dkimAttributesModel
				destinationConfigOneOfModel.Spf = spfAttributesModel

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigOneOfModel

				// Construct an instance of the UpdateDestinationOptions model
				updateDestinationOptionsModel := new(eventnotificationsv1.UpdateDestinationOptions)
				updateDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateDestinationOptionsModel.ID = core.StringPtr("testString")
				updateDestinationOptionsModel.Name = core.StringPtr("testString")
				updateDestinationOptionsModel.Description = core.StringPtr("testString")
				updateDestinationOptionsModel.Config = destinationConfigModel
				updateDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "config": {"params": {"domain": "Domain", "dkim": {"public_key": "PublicKey", "selector": "Selector", "verification": "Verification"}, "spf": {"txt_name": "TxtName", "txt_value": "TxtValue", "verification": "Verification"}}}, "updated_at": "2019-01-01T12:00:00.000Z", "subscription_count": 0, "subscription_names": ["SubscriptionNames"]}`)
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

				// Construct an instance of the DkimAttributes model
				dkimAttributesModel := new(eventnotificationsv1.DkimAttributes)
				dkimAttributesModel.PublicKey = core.StringPtr("testString")
				dkimAttributesModel.Selector = core.StringPtr("testString")
				dkimAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the SpfAttributes model
				spfAttributesModel := new(eventnotificationsv1.SpfAttributes)
				spfAttributesModel.TxtName = core.StringPtr("testString")
				spfAttributesModel.TxtValue = core.StringPtr("testString")
				spfAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the DestinationConfigOneOfCustomDomainEmailDestinationConfig model
				destinationConfigOneOfModel := new(eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig)
				destinationConfigOneOfModel.Domain = core.StringPtr("testString")
				destinationConfigOneOfModel.Dkim = dkimAttributesModel
				destinationConfigOneOfModel.Spf = spfAttributesModel

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigOneOfModel

				// Construct an instance of the UpdateDestinationOptions model
				updateDestinationOptionsModel := new(eventnotificationsv1.UpdateDestinationOptions)
				updateDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateDestinationOptionsModel.ID = core.StringPtr("testString")
				updateDestinationOptionsModel.Name = core.StringPtr("testString")
				updateDestinationOptionsModel.Description = core.StringPtr("testString")
				updateDestinationOptionsModel.Config = destinationConfigModel
				updateDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "type": "webhook", "config": {"params": {"domain": "Domain", "dkim": {"public_key": "PublicKey", "selector": "Selector", "verification": "Verification"}, "spf": {"txt_name": "TxtName", "txt_value": "TxtValue", "verification": "Verification"}}}, "updated_at": "2019-01-01T12:00:00.000Z", "subscription_count": 0, "subscription_names": ["SubscriptionNames"]}`)
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

				// Construct an instance of the DkimAttributes model
				dkimAttributesModel := new(eventnotificationsv1.DkimAttributes)
				dkimAttributesModel.PublicKey = core.StringPtr("testString")
				dkimAttributesModel.Selector = core.StringPtr("testString")
				dkimAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the SpfAttributes model
				spfAttributesModel := new(eventnotificationsv1.SpfAttributes)
				spfAttributesModel.TxtName = core.StringPtr("testString")
				spfAttributesModel.TxtValue = core.StringPtr("testString")
				spfAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the DestinationConfigOneOfCustomDomainEmailDestinationConfig model
				destinationConfigOneOfModel := new(eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig)
				destinationConfigOneOfModel.Domain = core.StringPtr("testString")
				destinationConfigOneOfModel.Dkim = dkimAttributesModel
				destinationConfigOneOfModel.Spf = spfAttributesModel

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigOneOfModel

				// Construct an instance of the UpdateDestinationOptions model
				updateDestinationOptionsModel := new(eventnotificationsv1.UpdateDestinationOptions)
				updateDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateDestinationOptionsModel.ID = core.StringPtr("testString")
				updateDestinationOptionsModel.Name = core.StringPtr("testString")
				updateDestinationOptionsModel.Description = core.StringPtr("testString")
				updateDestinationOptionsModel.Config = destinationConfigModel
				updateDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.UpdateDestination(updateDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDestination with error: Param validation error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:  testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the UpdateDestinationOptions model
				updateDestinationOptionsModel := new(eventnotificationsv1.UpdateDestinationOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := eventNotificationsService.UpdateDestination(updateDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke UpdateDestination with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DkimAttributes model
				dkimAttributesModel := new(eventnotificationsv1.DkimAttributes)
				dkimAttributesModel.PublicKey = core.StringPtr("testString")
				dkimAttributesModel.Selector = core.StringPtr("testString")
				dkimAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the SpfAttributes model
				spfAttributesModel := new(eventnotificationsv1.SpfAttributes)
				spfAttributesModel.TxtName = core.StringPtr("testString")
				spfAttributesModel.TxtValue = core.StringPtr("testString")
				spfAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the DestinationConfigOneOfCustomDomainEmailDestinationConfig model
				destinationConfigOneOfModel := new(eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig)
				destinationConfigOneOfModel.Domain = core.StringPtr("testString")
				destinationConfigOneOfModel.Dkim = dkimAttributesModel
				destinationConfigOneOfModel.Spf = spfAttributesModel

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigOneOfModel

				// Construct an instance of the UpdateDestinationOptions model
				updateDestinationOptionsModel := new(eventnotificationsv1.UpdateDestinationOptions)
				updateDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateDestinationOptionsModel.ID = core.StringPtr("testString")
				updateDestinationOptionsModel.Name = core.StringPtr("testString")
				updateDestinationOptionsModel.Description = core.StringPtr("testString")
				updateDestinationOptionsModel.Config = destinationConfigModel
				updateDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
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

				// Construct an instance of the DkimAttributes model
				dkimAttributesModel := new(eventnotificationsv1.DkimAttributes)
				dkimAttributesModel.PublicKey = core.StringPtr("testString")
				dkimAttributesModel.Selector = core.StringPtr("testString")
				dkimAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the SpfAttributes model
				spfAttributesModel := new(eventnotificationsv1.SpfAttributes)
				spfAttributesModel.TxtName = core.StringPtr("testString")
				spfAttributesModel.TxtValue = core.StringPtr("testString")
				spfAttributesModel.Verification = core.StringPtr("testString")

				// Construct an instance of the DestinationConfigOneOfCustomDomainEmailDestinationConfig model
				destinationConfigOneOfModel := new(eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig)
				destinationConfigOneOfModel.Domain = core.StringPtr("testString")
				destinationConfigOneOfModel.Dkim = dkimAttributesModel
				destinationConfigOneOfModel.Spf = spfAttributesModel

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				destinationConfigModel.Params = destinationConfigOneOfModel

				// Construct an instance of the UpdateDestinationOptions model
				updateDestinationOptionsModel := new(eventnotificationsv1.UpdateDestinationOptions)
				updateDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateDestinationOptionsModel.ID = core.StringPtr("testString")
				updateDestinationOptionsModel.Name = core.StringPtr("testString")
				updateDestinationOptionsModel.Description = core.StringPtr("testString")
				updateDestinationOptionsModel.Config = destinationConfigModel
				updateDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
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
	Describe(`UpdateVerifyDestination(updateVerifyDestinationOptions *UpdateVerifyDestinationOptions) - Operation response error`, func() {
		updateVerifyDestinationPath := "/v1/instances/testString/destinations/testString/verify"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateVerifyDestinationPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateVerifyDestination with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the UpdateVerifyDestinationOptions model
				updateVerifyDestinationOptionsModel := new(eventnotificationsv1.UpdateVerifyDestinationOptions)
				updateVerifyDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.ID = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.Type = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.UpdateVerifyDestination(updateVerifyDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.UpdateVerifyDestination(updateVerifyDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateVerifyDestination(updateVerifyDestinationOptions *UpdateVerifyDestinationOptions)`, func() {
		updateVerifyDestinationPath := "/v1/instances/testString/destinations/testString/verify"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateVerifyDestinationPath))
					Expect(req.Method).To(Equal("PATCH"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "Type", "verification": "Verification"}`)
				}))
			})
			It(`Invoke UpdateVerifyDestination successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateVerifyDestinationOptions model
				updateVerifyDestinationOptionsModel := new(eventnotificationsv1.UpdateVerifyDestinationOptions)
				updateVerifyDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.ID = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.Type = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.UpdateVerifyDestinationWithContext(ctx, updateVerifyDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.UpdateVerifyDestination(updateVerifyDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.UpdateVerifyDestinationWithContext(ctx, updateVerifyDestinationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateVerifyDestinationPath))
					Expect(req.Method).To(Equal("PATCH"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "Type", "verification": "Verification"}`)
				}))
			})
			It(`Invoke UpdateVerifyDestination successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.UpdateVerifyDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateVerifyDestinationOptions model
				updateVerifyDestinationOptionsModel := new(eventnotificationsv1.UpdateVerifyDestinationOptions)
				updateVerifyDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.ID = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.Type = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.UpdateVerifyDestination(updateVerifyDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateVerifyDestination with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the UpdateVerifyDestinationOptions model
				updateVerifyDestinationOptionsModel := new(eventnotificationsv1.UpdateVerifyDestinationOptions)
				updateVerifyDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.ID = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.Type = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.UpdateVerifyDestination(updateVerifyDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateVerifyDestinationOptions model with no property values
				updateVerifyDestinationOptionsModelNew := new(eventnotificationsv1.UpdateVerifyDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.UpdateVerifyDestination(updateVerifyDestinationOptionsModelNew)
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
			It(`Invoke UpdateVerifyDestination successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the UpdateVerifyDestinationOptions model
				updateVerifyDestinationOptionsModel := new(eventnotificationsv1.UpdateVerifyDestinationOptions)
				updateVerifyDestinationOptionsModel.InstanceID = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.ID = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.Type = core.StringPtr("testString")
				updateVerifyDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.UpdateVerifyDestination(updateVerifyDestinationOptionsModel)
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
	Describe(`CreateTagsSubscription(createTagsSubscriptionOptions *CreateTagsSubscriptionOptions) - Operation response error`, func() {
		createTagsSubscriptionPath := "/v1/instances/testString/destinations/testString/tag_subscriptions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTagsSubscriptionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTagsSubscription with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the CreateTagsSubscriptionOptions model
				createTagsSubscriptionOptionsModel := new(eventnotificationsv1.CreateTagsSubscriptionOptions)
				createTagsSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.ID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.DeviceID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.TagName = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.CreateTagsSubscription(createTagsSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.CreateTagsSubscription(createTagsSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTagsSubscription(createTagsSubscriptionOptions *CreateTagsSubscriptionOptions)`, func() {
		createTagsSubscriptionPath := "/v1/instances/testString/destinations/testString/tag_subscriptions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTagsSubscriptionPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "device_id": "DeviceID", "tag_name": "TagName", "user_id": "UserID", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateTagsSubscription successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the CreateTagsSubscriptionOptions model
				createTagsSubscriptionOptionsModel := new(eventnotificationsv1.CreateTagsSubscriptionOptions)
				createTagsSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.ID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.DeviceID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.TagName = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.CreateTagsSubscriptionWithContext(ctx, createTagsSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.CreateTagsSubscription(createTagsSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.CreateTagsSubscriptionWithContext(ctx, createTagsSubscriptionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTagsSubscriptionPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "device_id": "DeviceID", "tag_name": "TagName", "user_id": "UserID", "created_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateTagsSubscription successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.CreateTagsSubscription(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateTagsSubscriptionOptions model
				createTagsSubscriptionOptionsModel := new(eventnotificationsv1.CreateTagsSubscriptionOptions)
				createTagsSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.ID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.DeviceID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.TagName = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.CreateTagsSubscription(createTagsSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTagsSubscription with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the CreateTagsSubscriptionOptions model
				createTagsSubscriptionOptionsModel := new(eventnotificationsv1.CreateTagsSubscriptionOptions)
				createTagsSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.ID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.DeviceID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.TagName = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.CreateTagsSubscription(createTagsSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTagsSubscriptionOptions model with no property values
				createTagsSubscriptionOptionsModelNew := new(eventnotificationsv1.CreateTagsSubscriptionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.CreateTagsSubscription(createTagsSubscriptionOptionsModelNew)
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
			It(`Invoke CreateTagsSubscription successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the CreateTagsSubscriptionOptions model
				createTagsSubscriptionOptionsModel := new(eventnotificationsv1.CreateTagsSubscriptionOptions)
				createTagsSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.ID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.DeviceID = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.TagName = core.StringPtr("testString")
				createTagsSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.CreateTagsSubscription(createTagsSubscriptionOptionsModel)
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
	Describe(`ListTagsSubscription(listTagsSubscriptionOptions *ListTagsSubscriptionOptions) - Operation response error`, func() {
		listTagsSubscriptionPath := "/v1/instances/testString/destinations/testString/tag_subscriptions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTagsSubscriptionPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["device_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["user_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["tag_name"]).To(Equal([]string{"testString"}))
					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTagsSubscription with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListTagsSubscriptionOptions model
				listTagsSubscriptionOptionsModel := new(eventnotificationsv1.ListTagsSubscriptionOptions)
				listTagsSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.ID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.DeviceID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.UserID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.TagName = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTagsSubscriptionOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsSubscriptionOptionsModel.Search = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.ListTagsSubscription(listTagsSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.ListTagsSubscription(listTagsSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTagsSubscription(listTagsSubscriptionOptions *ListTagsSubscriptionOptions)`, func() {
		listTagsSubscriptionPath := "/v1/instances/testString/destinations/testString/tag_subscriptions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTagsSubscriptionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["device_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["user_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["tag_name"]).To(Equal([]string{"testString"}))
					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "tag_subscriptions": [{"id": "ID", "device_id": "DeviceID", "tag_name": "TagName", "user_id": "UserID", "updated_at": "2019-01-01T12:00:00.000Z"}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListTagsSubscription successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the ListTagsSubscriptionOptions model
				listTagsSubscriptionOptionsModel := new(eventnotificationsv1.ListTagsSubscriptionOptions)
				listTagsSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.ID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.DeviceID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.UserID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.TagName = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTagsSubscriptionOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsSubscriptionOptionsModel.Search = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.ListTagsSubscriptionWithContext(ctx, listTagsSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.ListTagsSubscription(listTagsSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.ListTagsSubscriptionWithContext(ctx, listTagsSubscriptionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTagsSubscriptionPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["device_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["user_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["tag_name"]).To(Equal([]string{"testString"}))
					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "tag_subscriptions": [{"id": "ID", "device_id": "DeviceID", "tag_name": "TagName", "user_id": "UserID", "updated_at": "2019-01-01T12:00:00.000Z"}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListTagsSubscription successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.ListTagsSubscription(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTagsSubscriptionOptions model
				listTagsSubscriptionOptionsModel := new(eventnotificationsv1.ListTagsSubscriptionOptions)
				listTagsSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.ID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.DeviceID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.UserID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.TagName = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTagsSubscriptionOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsSubscriptionOptionsModel.Search = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.ListTagsSubscription(listTagsSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTagsSubscription with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListTagsSubscriptionOptions model
				listTagsSubscriptionOptionsModel := new(eventnotificationsv1.ListTagsSubscriptionOptions)
				listTagsSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.ID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.DeviceID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.UserID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.TagName = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTagsSubscriptionOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsSubscriptionOptionsModel.Search = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.ListTagsSubscription(listTagsSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTagsSubscriptionOptions model with no property values
				listTagsSubscriptionOptionsModelNew := new(eventnotificationsv1.ListTagsSubscriptionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.ListTagsSubscription(listTagsSubscriptionOptionsModelNew)
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
			It(`Invoke ListTagsSubscription successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListTagsSubscriptionOptions model
				listTagsSubscriptionOptionsModel := new(eventnotificationsv1.ListTagsSubscriptionOptions)
				listTagsSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.ID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.DeviceID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.UserID = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.TagName = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.Limit = core.Int64Ptr(int64(10))
				listTagsSubscriptionOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsSubscriptionOptionsModel.Search = core.StringPtr("testString")
				listTagsSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.ListTagsSubscription(listTagsSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(eventnotificationsv1.TagsSubscriptionList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(eventnotificationsv1.TagsSubscriptionList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.TagsSubscriptionList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.TagsSubscriptionList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTagsSubscriptionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"limit":1,"tag_subscriptions":[{"id":"ID","device_id":"DeviceID","tag_name":"TagName","user_id":"UserID","updated_at":"2019-01-01T12:00:00.000Z"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"tag_subscriptions":[{"id":"ID","device_id":"DeviceID","tag_name":"TagName","user_id":"UserID","updated_at":"2019-01-01T12:00:00.000Z"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use TagsSubscriptionPager.GetNext successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listTagsSubscriptionOptionsModel := &eventnotificationsv1.ListTagsSubscriptionOptions{
					InstanceID: core.StringPtr("testString"),
					ID: core.StringPtr("testString"),
					DeviceID: core.StringPtr("testString"),
					UserID: core.StringPtr("testString"),
					TagName: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewTagsSubscriptionPager(listTagsSubscriptionOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []eventnotificationsv1.TagsSubscriptionListItem
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use TagsSubscriptionPager.GetAll successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listTagsSubscriptionOptionsModel := &eventnotificationsv1.ListTagsSubscriptionOptions{
					InstanceID: core.StringPtr("testString"),
					ID: core.StringPtr("testString"),
					DeviceID: core.StringPtr("testString"),
					UserID: core.StringPtr("testString"),
					TagName: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewTagsSubscriptionPager(listTagsSubscriptionOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`DeleteTagsSubscription(deleteTagsSubscriptionOptions *DeleteTagsSubscriptionOptions)`, func() {
		deleteTagsSubscriptionPath := "/v1/instances/testString/destinations/testString/tag_subscriptions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTagsSubscriptionPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["device_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["tag_name"]).To(Equal([]string{"testString"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTagsSubscription successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := eventNotificationsService.DeleteTagsSubscription(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTagsSubscriptionOptions model
				deleteTagsSubscriptionOptionsModel := new(eventnotificationsv1.DeleteTagsSubscriptionOptions)
				deleteTagsSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				deleteTagsSubscriptionOptionsModel.ID = core.StringPtr("testString")
				deleteTagsSubscriptionOptionsModel.DeviceID = core.StringPtr("testString")
				deleteTagsSubscriptionOptionsModel.TagName = core.StringPtr("testString")
				deleteTagsSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = eventNotificationsService.DeleteTagsSubscription(deleteTagsSubscriptionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTagsSubscription with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the DeleteTagsSubscriptionOptions model
				deleteTagsSubscriptionOptionsModel := new(eventnotificationsv1.DeleteTagsSubscriptionOptions)
				deleteTagsSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				deleteTagsSubscriptionOptionsModel.ID = core.StringPtr("testString")
				deleteTagsSubscriptionOptionsModel.DeviceID = core.StringPtr("testString")
				deleteTagsSubscriptionOptionsModel.TagName = core.StringPtr("testString")
				deleteTagsSubscriptionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := eventNotificationsService.DeleteTagsSubscription(deleteTagsSubscriptionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTagsSubscriptionOptions model with no property values
				deleteTagsSubscriptionOptionsModelNew := new(eventnotificationsv1.DeleteTagsSubscriptionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = eventNotificationsService.DeleteTagsSubscription(deleteTagsSubscriptionOptionsModelNew)
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
					fmt.Fprint(res, `} this is not valid json {`)
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
				subscriptionCreateAttributesModel.Invited = []string{"testString"}

				// Construct an instance of the CreateSubscriptionOptions model
				createSubscriptionOptionsModel := new(eventnotificationsv1.CreateSubscriptionOptions)
				createSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Name = core.StringPtr("testString")
				createSubscriptionOptionsModel.DestinationID = core.StringPtr("testString")
				createSubscriptionOptionsModel.TopicID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Description = core.StringPtr("testString")
				createSubscriptionOptionsModel.Attributes = subscriptionCreateAttributesModel
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {"subscribed": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z"}], "unsubscribed": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z"}], "invited": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z", "expires_at": "2019-01-01T12:00:00.000Z"}]}}`)
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
				subscriptionCreateAttributesModel.Invited = []string{"testString"}

				// Construct an instance of the CreateSubscriptionOptions model
				createSubscriptionOptionsModel := new(eventnotificationsv1.CreateSubscriptionOptions)
				createSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Name = core.StringPtr("testString")
				createSubscriptionOptionsModel.DestinationID = core.StringPtr("testString")
				createSubscriptionOptionsModel.TopicID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Description = core.StringPtr("testString")
				createSubscriptionOptionsModel.Attributes = subscriptionCreateAttributesModel
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {"subscribed": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z"}], "unsubscribed": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z"}], "invited": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z", "expires_at": "2019-01-01T12:00:00.000Z"}]}}`)
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
				subscriptionCreateAttributesModel.Invited = []string{"testString"}

				// Construct an instance of the CreateSubscriptionOptions model
				createSubscriptionOptionsModel := new(eventnotificationsv1.CreateSubscriptionOptions)
				createSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Name = core.StringPtr("testString")
				createSubscriptionOptionsModel.DestinationID = core.StringPtr("testString")
				createSubscriptionOptionsModel.TopicID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Description = core.StringPtr("testString")
				createSubscriptionOptionsModel.Attributes = subscriptionCreateAttributesModel
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
				subscriptionCreateAttributesModel.Invited = []string{"testString"}

				// Construct an instance of the CreateSubscriptionOptions model
				createSubscriptionOptionsModel := new(eventnotificationsv1.CreateSubscriptionOptions)
				createSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Name = core.StringPtr("testString")
				createSubscriptionOptionsModel.DestinationID = core.StringPtr("testString")
				createSubscriptionOptionsModel.TopicID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Description = core.StringPtr("testString")
				createSubscriptionOptionsModel.Attributes = subscriptionCreateAttributesModel
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
				subscriptionCreateAttributesModel.Invited = []string{"testString"}

				// Construct an instance of the CreateSubscriptionOptions model
				createSubscriptionOptionsModel := new(eventnotificationsv1.CreateSubscriptionOptions)
				createSubscriptionOptionsModel.InstanceID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Name = core.StringPtr("testString")
				createSubscriptionOptionsModel.DestinationID = core.StringPtr("testString")
				createSubscriptionOptionsModel.TopicID = core.StringPtr("testString")
				createSubscriptionOptionsModel.Description = core.StringPtr("testString")
				createSubscriptionOptionsModel.Attributes = subscriptionCreateAttributesModel
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
					fmt.Fprint(res, `} this is not valid json {`)
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
				listSubscriptionsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "subscriptions": [{"id": "ID", "name": "Name", "description": "Description", "destination_id": "DestinationID", "destination_name": "DestinationName", "destination_type": "sms_ibm", "topic_id": "TopicID", "topic_name": "TopicName", "updated_at": "2019-01-01T12:00:00.000Z"}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
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
				listSubscriptionsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "subscriptions": [{"id": "ID", "name": "Name", "description": "Description", "destination_id": "DestinationID", "destination_name": "DestinationName", "destination_type": "sms_ibm", "topic_id": "TopicID", "topic_name": "TopicName", "updated_at": "2019-01-01T12:00:00.000Z"}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
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
				listSubscriptionsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
				listSubscriptionsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
				listSubscriptionsOptionsModel.Limit = core.Int64Ptr(int64(10))
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
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(eventnotificationsv1.SubscriptionList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(eventnotificationsv1.SubscriptionList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.SubscriptionList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.SubscriptionList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSubscriptionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"subscriptions":[{"id":"ID","name":"Name","description":"Description","destination_id":"DestinationID","destination_name":"DestinationName","destination_type":"sms_ibm","topic_id":"TopicID","topic_name":"TopicName","updated_at":"2019-01-01T12:00:00.000Z"}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"subscriptions":[{"id":"ID","name":"Name","description":"Description","destination_id":"DestinationID","destination_name":"DestinationName","destination_type":"sms_ibm","topic_id":"TopicID","topic_name":"TopicName","updated_at":"2019-01-01T12:00:00.000Z"}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use SubscriptionsPager.GetNext successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listSubscriptionsOptionsModel := &eventnotificationsv1.ListSubscriptionsOptions{
					InstanceID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewSubscriptionsPager(listSubscriptionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []eventnotificationsv1.SubscriptionListItem
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use SubscriptionsPager.GetAll successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listSubscriptionsOptionsModel := &eventnotificationsv1.ListSubscriptionsOptions{
					InstanceID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewSubscriptionsPager(listSubscriptionsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
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
					fmt.Fprint(res, `} this is not valid json {`)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {"subscribed": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z"}], "unsubscribed": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z"}], "invited": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z", "expires_at": "2019-01-01T12:00:00.000Z"}]}}`)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {"subscribed": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z"}], "unsubscribed": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z"}], "invited": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z", "expires_at": "2019-01-01T12:00:00.000Z"}]}}`)
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
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSubscription with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the UpdateAttributesInvited model
				updateAttributesInvitedModel := new(eventnotificationsv1.UpdateAttributesInvited)
				updateAttributesInvitedModel.Add = []string{"testString"}
				updateAttributesInvitedModel.Remove = []string{"testString"}

				// Construct an instance of the UpdateAttributesSubscribed model
				updateAttributesSubscribedModel := new(eventnotificationsv1.UpdateAttributesSubscribed)
				updateAttributesSubscribedModel.Remove = []string{"testString"}

				// Construct an instance of the UpdateAttributesUnsubscribed model
				updateAttributesUnsubscribedModel := new(eventnotificationsv1.UpdateAttributesUnsubscribed)
				updateAttributesUnsubscribedModel.Remove = []string{"testString"}

				// Construct an instance of the SubscriptionUpdateAttributesSmsUpdateAttributes model
				subscriptionUpdateAttributesModel := new(eventnotificationsv1.SubscriptionUpdateAttributesSmsUpdateAttributes)
				subscriptionUpdateAttributesModel.Invited = updateAttributesInvitedModel
				subscriptionUpdateAttributesModel.Subscribed = updateAttributesSubscribedModel
				subscriptionUpdateAttributesModel.Unsubscribed = updateAttributesUnsubscribedModel

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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {"subscribed": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z"}], "unsubscribed": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z"}], "invited": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z", "expires_at": "2019-01-01T12:00:00.000Z"}]}}`)
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

				// Construct an instance of the UpdateAttributesInvited model
				updateAttributesInvitedModel := new(eventnotificationsv1.UpdateAttributesInvited)
				updateAttributesInvitedModel.Add = []string{"testString"}
				updateAttributesInvitedModel.Remove = []string{"testString"}

				// Construct an instance of the UpdateAttributesSubscribed model
				updateAttributesSubscribedModel := new(eventnotificationsv1.UpdateAttributesSubscribed)
				updateAttributesSubscribedModel.Remove = []string{"testString"}

				// Construct an instance of the UpdateAttributesUnsubscribed model
				updateAttributesUnsubscribedModel := new(eventnotificationsv1.UpdateAttributesUnsubscribed)
				updateAttributesUnsubscribedModel.Remove = []string{"testString"}

				// Construct an instance of the SubscriptionUpdateAttributesSmsUpdateAttributes model
				subscriptionUpdateAttributesModel := new(eventnotificationsv1.SubscriptionUpdateAttributesSmsUpdateAttributes)
				subscriptionUpdateAttributesModel.Invited = updateAttributesInvitedModel
				subscriptionUpdateAttributesModel.Subscribed = updateAttributesSubscribedModel
				subscriptionUpdateAttributesModel.Unsubscribed = updateAttributesUnsubscribedModel

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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {"subscribed": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z"}], "unsubscribed": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z"}], "invited": [{"phone_number": "PhoneNumber", "updated_at": "2019-01-01T12:00:00.000Z", "expires_at": "2019-01-01T12:00:00.000Z"}]}}`)
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

				// Construct an instance of the UpdateAttributesInvited model
				updateAttributesInvitedModel := new(eventnotificationsv1.UpdateAttributesInvited)
				updateAttributesInvitedModel.Add = []string{"testString"}
				updateAttributesInvitedModel.Remove = []string{"testString"}

				// Construct an instance of the UpdateAttributesSubscribed model
				updateAttributesSubscribedModel := new(eventnotificationsv1.UpdateAttributesSubscribed)
				updateAttributesSubscribedModel.Remove = []string{"testString"}

				// Construct an instance of the UpdateAttributesUnsubscribed model
				updateAttributesUnsubscribedModel := new(eventnotificationsv1.UpdateAttributesUnsubscribed)
				updateAttributesUnsubscribedModel.Remove = []string{"testString"}

				// Construct an instance of the SubscriptionUpdateAttributesSmsUpdateAttributes model
				subscriptionUpdateAttributesModel := new(eventnotificationsv1.SubscriptionUpdateAttributesSmsUpdateAttributes)
				subscriptionUpdateAttributesModel.Invited = updateAttributesInvitedModel
				subscriptionUpdateAttributesModel.Subscribed = updateAttributesSubscribedModel
				subscriptionUpdateAttributesModel.Unsubscribed = updateAttributesUnsubscribedModel

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

				// Construct an instance of the UpdateAttributesInvited model
				updateAttributesInvitedModel := new(eventnotificationsv1.UpdateAttributesInvited)
				updateAttributesInvitedModel.Add = []string{"testString"}
				updateAttributesInvitedModel.Remove = []string{"testString"}

				// Construct an instance of the UpdateAttributesSubscribed model
				updateAttributesSubscribedModel := new(eventnotificationsv1.UpdateAttributesSubscribed)
				updateAttributesSubscribedModel.Remove = []string{"testString"}

				// Construct an instance of the UpdateAttributesUnsubscribed model
				updateAttributesUnsubscribedModel := new(eventnotificationsv1.UpdateAttributesUnsubscribed)
				updateAttributesUnsubscribedModel.Remove = []string{"testString"}

				// Construct an instance of the SubscriptionUpdateAttributesSmsUpdateAttributes model
				subscriptionUpdateAttributesModel := new(eventnotificationsv1.SubscriptionUpdateAttributesSmsUpdateAttributes)
				subscriptionUpdateAttributesModel.Invited = updateAttributesInvitedModel
				subscriptionUpdateAttributesModel.Subscribed = updateAttributesSubscribedModel
				subscriptionUpdateAttributesModel.Unsubscribed = updateAttributesUnsubscribedModel

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

				// Construct an instance of the UpdateAttributesInvited model
				updateAttributesInvitedModel := new(eventnotificationsv1.UpdateAttributesInvited)
				updateAttributesInvitedModel.Add = []string{"testString"}
				updateAttributesInvitedModel.Remove = []string{"testString"}

				// Construct an instance of the UpdateAttributesSubscribed model
				updateAttributesSubscribedModel := new(eventnotificationsv1.UpdateAttributesSubscribed)
				updateAttributesSubscribedModel.Remove = []string{"testString"}

				// Construct an instance of the UpdateAttributesUnsubscribed model
				updateAttributesUnsubscribedModel := new(eventnotificationsv1.UpdateAttributesUnsubscribed)
				updateAttributesUnsubscribedModel.Remove = []string{"testString"}

				// Construct an instance of the SubscriptionUpdateAttributesSmsUpdateAttributes model
				subscriptionUpdateAttributesModel := new(eventnotificationsv1.SubscriptionUpdateAttributesSmsUpdateAttributes)
				subscriptionUpdateAttributesModel.Invited = updateAttributesInvitedModel
				subscriptionUpdateAttributesModel.Subscribed = updateAttributesSubscribedModel
				subscriptionUpdateAttributesModel.Unsubscribed = updateAttributesUnsubscribedModel

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
	Describe(`ListIntegrations(listIntegrationsOptions *ListIntegrationsOptions) - Operation response error`, func() {
		listIntegrationsPath := "/v1/instances/testString/integrations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listIntegrationsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for offset query parameter
					// TODO: Add check for limit query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListIntegrations with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListIntegrationsOptions model
				listIntegrationsOptionsModel := new(eventnotificationsv1.ListIntegrationsOptions)
				listIntegrationsOptionsModel.InstanceID = core.StringPtr("testString")
				listIntegrationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listIntegrationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listIntegrationsOptionsModel.Search = core.StringPtr("testString")
				listIntegrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.ListIntegrations(listIntegrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.ListIntegrations(listIntegrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListIntegrations(listIntegrationsOptions *ListIntegrationsOptions)`, func() {
		listIntegrationsPath := "/v1/instances/testString/integrations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listIntegrationsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for offset query parameter
					// TODO: Add check for limit query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "integrations": [{"id": "9fab83da-98cb-4f18-a7ba-b6f0435c9673", "type": "Type", "metadata": {"endpoint": "Endpoint", "crn": "CRN", "root_key_id": "RootKeyID"}, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListIntegrations successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the ListIntegrationsOptions model
				listIntegrationsOptionsModel := new(eventnotificationsv1.ListIntegrationsOptions)
				listIntegrationsOptionsModel.InstanceID = core.StringPtr("testString")
				listIntegrationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listIntegrationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listIntegrationsOptionsModel.Search = core.StringPtr("testString")
				listIntegrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.ListIntegrationsWithContext(ctx, listIntegrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.ListIntegrations(listIntegrationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.ListIntegrationsWithContext(ctx, listIntegrationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listIntegrationsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for offset query parameter
					// TODO: Add check for limit query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 0, "offset": 6, "limit": 5, "integrations": [{"id": "9fab83da-98cb-4f18-a7ba-b6f0435c9673", "type": "Type", "metadata": {"endpoint": "Endpoint", "crn": "CRN", "root_key_id": "RootKeyID"}, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}], "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}}`)
				}))
			})
			It(`Invoke ListIntegrations successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.ListIntegrations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListIntegrationsOptions model
				listIntegrationsOptionsModel := new(eventnotificationsv1.ListIntegrationsOptions)
				listIntegrationsOptionsModel.InstanceID = core.StringPtr("testString")
				listIntegrationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listIntegrationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listIntegrationsOptionsModel.Search = core.StringPtr("testString")
				listIntegrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.ListIntegrations(listIntegrationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListIntegrations with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListIntegrationsOptions model
				listIntegrationsOptionsModel := new(eventnotificationsv1.ListIntegrationsOptions)
				listIntegrationsOptionsModel.InstanceID = core.StringPtr("testString")
				listIntegrationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listIntegrationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listIntegrationsOptionsModel.Search = core.StringPtr("testString")
				listIntegrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.ListIntegrations(listIntegrationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListIntegrationsOptions model with no property values
				listIntegrationsOptionsModelNew := new(eventnotificationsv1.ListIntegrationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.ListIntegrations(listIntegrationsOptionsModelNew)
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
			It(`Invoke ListIntegrations successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListIntegrationsOptions model
				listIntegrationsOptionsModel := new(eventnotificationsv1.ListIntegrationsOptions)
				listIntegrationsOptionsModel.InstanceID = core.StringPtr("testString")
				listIntegrationsOptionsModel.Offset = core.Int64Ptr(int64(0))
				listIntegrationsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listIntegrationsOptionsModel.Search = core.StringPtr("testString")
				listIntegrationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.ListIntegrations(listIntegrationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(eventnotificationsv1.IntegrationList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(eventnotificationsv1.IntegrationList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.IntegrationList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(eventnotificationsv1.IntegrationList)
				nextObject := new(eventnotificationsv1.PageHrefResponse)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listIntegrationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"limit":1,"integrations":[{"id":"9fab83da-98cb-4f18-a7ba-b6f0435c9673","type":"Type","metadata":{"endpoint":"Endpoint","crn":"CRN","root_key_id":"RootKeyID"},"created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"integrations":[{"id":"9fab83da-98cb-4f18-a7ba-b6f0435c9673","type":"Type","metadata":{"endpoint":"Endpoint","crn":"CRN","root_key_id":"RootKeyID"},"created_at":"2019-01-01T12:00:00.000Z","updated_at":"2019-01-01T12:00:00.000Z"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use IntegrationsPager.GetNext successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listIntegrationsOptionsModel := &eventnotificationsv1.ListIntegrationsOptions{
					InstanceID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewIntegrationsPager(listIntegrationsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []eventnotificationsv1.IntegrationListItem
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use IntegrationsPager.GetAll successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				listIntegrationsOptionsModel := &eventnotificationsv1.ListIntegrationsOptions{
					InstanceID: core.StringPtr("testString"),
					Limit: core.Int64Ptr(int64(10)),
					Search: core.StringPtr("testString"),
				}

				pager, err := eventNotificationsService.NewIntegrationsPager(listIntegrationsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetIntegration(getIntegrationOptions *GetIntegrationOptions) - Operation response error`, func() {
		getIntegrationPath := "/v1/instances/testString/integrations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIntegrationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetIntegration with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetIntegrationOptions model
				getIntegrationOptionsModel := new(eventnotificationsv1.GetIntegrationOptions)
				getIntegrationOptionsModel.InstanceID = core.StringPtr("testString")
				getIntegrationOptionsModel.ID = core.StringPtr("testString")
				getIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.GetIntegration(getIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.GetIntegration(getIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetIntegration(getIntegrationOptions *GetIntegrationOptions)`, func() {
		getIntegrationPath := "/v1/instances/testString/integrations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getIntegrationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "9fab83da-98cb-4f18-a7ba-b6f0435c9673", "type": "Type", "metadata": {"endpoint": "Endpoint", "crn": "CRN", "root_key_id": "RootKeyID"}, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetIntegration successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the GetIntegrationOptions model
				getIntegrationOptionsModel := new(eventnotificationsv1.GetIntegrationOptions)
				getIntegrationOptionsModel.InstanceID = core.StringPtr("testString")
				getIntegrationOptionsModel.ID = core.StringPtr("testString")
				getIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.GetIntegrationWithContext(ctx, getIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.GetIntegration(getIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.GetIntegrationWithContext(ctx, getIntegrationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getIntegrationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "9fab83da-98cb-4f18-a7ba-b6f0435c9673", "type": "Type", "metadata": {"endpoint": "Endpoint", "crn": "CRN", "root_key_id": "RootKeyID"}, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetIntegration successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.GetIntegration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetIntegrationOptions model
				getIntegrationOptionsModel := new(eventnotificationsv1.GetIntegrationOptions)
				getIntegrationOptionsModel.InstanceID = core.StringPtr("testString")
				getIntegrationOptionsModel.ID = core.StringPtr("testString")
				getIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.GetIntegration(getIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetIntegration with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetIntegrationOptions model
				getIntegrationOptionsModel := new(eventnotificationsv1.GetIntegrationOptions)
				getIntegrationOptionsModel.InstanceID = core.StringPtr("testString")
				getIntegrationOptionsModel.ID = core.StringPtr("testString")
				getIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.GetIntegration(getIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetIntegrationOptions model with no property values
				getIntegrationOptionsModelNew := new(eventnotificationsv1.GetIntegrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.GetIntegration(getIntegrationOptionsModelNew)
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
			It(`Invoke GetIntegration successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetIntegrationOptions model
				getIntegrationOptionsModel := new(eventnotificationsv1.GetIntegrationOptions)
				getIntegrationOptionsModel.InstanceID = core.StringPtr("testString")
				getIntegrationOptionsModel.ID = core.StringPtr("testString")
				getIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.GetIntegration(getIntegrationOptionsModel)
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
	Describe(`ReplaceIntegration(replaceIntegrationOptions *ReplaceIntegrationOptions) - Operation response error`, func() {
		replaceIntegrationPath := "/v1/instances/testString/integrations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceIntegrationPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceIntegration with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the IntegrationMetadata model
				integrationMetadataModel := new(eventnotificationsv1.IntegrationMetadata)
				integrationMetadataModel.Endpoint = core.StringPtr("testString")
				integrationMetadataModel.CRN = core.StringPtr("testString")
				integrationMetadataModel.RootKeyID = core.StringPtr("testString")

				// Construct an instance of the ReplaceIntegrationOptions model
				replaceIntegrationOptionsModel := new(eventnotificationsv1.ReplaceIntegrationOptions)
				replaceIntegrationOptionsModel.InstanceID = core.StringPtr("testString")
				replaceIntegrationOptionsModel.ID = core.StringPtr("testString")
				replaceIntegrationOptionsModel.Type = core.StringPtr("testString")
				replaceIntegrationOptionsModel.Metadata = integrationMetadataModel
				replaceIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.ReplaceIntegration(replaceIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.ReplaceIntegration(replaceIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceIntegration(replaceIntegrationOptions *ReplaceIntegrationOptions)`, func() {
		replaceIntegrationPath := "/v1/instances/testString/integrations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceIntegrationPath))
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
					fmt.Fprintf(res, "%s", `{"id": "9fab83da-98cb-4f18-a7ba-b6f0435c9673", "type": "Type", "metadata": {"endpoint": "Endpoint", "crn": "CRN", "root_key_id": "RootKeyID"}, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke ReplaceIntegration successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the IntegrationMetadata model
				integrationMetadataModel := new(eventnotificationsv1.IntegrationMetadata)
				integrationMetadataModel.Endpoint = core.StringPtr("testString")
				integrationMetadataModel.CRN = core.StringPtr("testString")
				integrationMetadataModel.RootKeyID = core.StringPtr("testString")

				// Construct an instance of the ReplaceIntegrationOptions model
				replaceIntegrationOptionsModel := new(eventnotificationsv1.ReplaceIntegrationOptions)
				replaceIntegrationOptionsModel.InstanceID = core.StringPtr("testString")
				replaceIntegrationOptionsModel.ID = core.StringPtr("testString")
				replaceIntegrationOptionsModel.Type = core.StringPtr("testString")
				replaceIntegrationOptionsModel.Metadata = integrationMetadataModel
				replaceIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.ReplaceIntegrationWithContext(ctx, replaceIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.ReplaceIntegration(replaceIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.ReplaceIntegrationWithContext(ctx, replaceIntegrationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceIntegrationPath))
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
					fmt.Fprintf(res, "%s", `{"id": "9fab83da-98cb-4f18-a7ba-b6f0435c9673", "type": "Type", "metadata": {"endpoint": "Endpoint", "crn": "CRN", "root_key_id": "RootKeyID"}, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke ReplaceIntegration successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.ReplaceIntegration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the IntegrationMetadata model
				integrationMetadataModel := new(eventnotificationsv1.IntegrationMetadata)
				integrationMetadataModel.Endpoint = core.StringPtr("testString")
				integrationMetadataModel.CRN = core.StringPtr("testString")
				integrationMetadataModel.RootKeyID = core.StringPtr("testString")

				// Construct an instance of the ReplaceIntegrationOptions model
				replaceIntegrationOptionsModel := new(eventnotificationsv1.ReplaceIntegrationOptions)
				replaceIntegrationOptionsModel.InstanceID = core.StringPtr("testString")
				replaceIntegrationOptionsModel.ID = core.StringPtr("testString")
				replaceIntegrationOptionsModel.Type = core.StringPtr("testString")
				replaceIntegrationOptionsModel.Metadata = integrationMetadataModel
				replaceIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.ReplaceIntegration(replaceIntegrationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceIntegration with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the IntegrationMetadata model
				integrationMetadataModel := new(eventnotificationsv1.IntegrationMetadata)
				integrationMetadataModel.Endpoint = core.StringPtr("testString")
				integrationMetadataModel.CRN = core.StringPtr("testString")
				integrationMetadataModel.RootKeyID = core.StringPtr("testString")

				// Construct an instance of the ReplaceIntegrationOptions model
				replaceIntegrationOptionsModel := new(eventnotificationsv1.ReplaceIntegrationOptions)
				replaceIntegrationOptionsModel.InstanceID = core.StringPtr("testString")
				replaceIntegrationOptionsModel.ID = core.StringPtr("testString")
				replaceIntegrationOptionsModel.Type = core.StringPtr("testString")
				replaceIntegrationOptionsModel.Metadata = integrationMetadataModel
				replaceIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.ReplaceIntegration(replaceIntegrationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceIntegrationOptions model with no property values
				replaceIntegrationOptionsModelNew := new(eventnotificationsv1.ReplaceIntegrationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.ReplaceIntegration(replaceIntegrationOptionsModelNew)
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
			It(`Invoke ReplaceIntegration successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the IntegrationMetadata model
				integrationMetadataModel := new(eventnotificationsv1.IntegrationMetadata)
				integrationMetadataModel.Endpoint = core.StringPtr("testString")
				integrationMetadataModel.CRN = core.StringPtr("testString")
				integrationMetadataModel.RootKeyID = core.StringPtr("testString")

				// Construct an instance of the ReplaceIntegrationOptions model
				replaceIntegrationOptionsModel := new(eventnotificationsv1.ReplaceIntegrationOptions)
				replaceIntegrationOptionsModel.InstanceID = core.StringPtr("testString")
				replaceIntegrationOptionsModel.ID = core.StringPtr("testString")
				replaceIntegrationOptionsModel.Type = core.StringPtr("testString")
				replaceIntegrationOptionsModel.Metadata = integrationMetadataModel
				replaceIntegrationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.ReplaceIntegration(replaceIntegrationOptionsModel)
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
				// Construct an instance of the DkimAttributes model
				dkimAttributesModel := new(eventnotificationsv1.DkimAttributes)
				Expect(dkimAttributesModel).ToNot(BeNil())
				dkimAttributesModel.PublicKey = core.StringPtr("testString")
				dkimAttributesModel.Selector = core.StringPtr("testString")
				dkimAttributesModel.Verification = core.StringPtr("testString")
				Expect(dkimAttributesModel.PublicKey).To(Equal(core.StringPtr("testString")))
				Expect(dkimAttributesModel.Selector).To(Equal(core.StringPtr("testString")))
				Expect(dkimAttributesModel.Verification).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SpfAttributes model
				spfAttributesModel := new(eventnotificationsv1.SpfAttributes)
				Expect(spfAttributesModel).ToNot(BeNil())
				spfAttributesModel.TxtName = core.StringPtr("testString")
				spfAttributesModel.TxtValue = core.StringPtr("testString")
				spfAttributesModel.Verification = core.StringPtr("testString")
				Expect(spfAttributesModel.TxtName).To(Equal(core.StringPtr("testString")))
				Expect(spfAttributesModel.TxtValue).To(Equal(core.StringPtr("testString")))
				Expect(spfAttributesModel.Verification).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DestinationConfigOneOfCustomDomainEmailDestinationConfig model
				destinationConfigOneOfModel := new(eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig)
				Expect(destinationConfigOneOfModel).ToNot(BeNil())
				destinationConfigOneOfModel.Domain = core.StringPtr("testString")
				destinationConfigOneOfModel.Dkim = dkimAttributesModel
				destinationConfigOneOfModel.Spf = spfAttributesModel
				Expect(destinationConfigOneOfModel.Domain).To(Equal(core.StringPtr("testString")))
				Expect(destinationConfigOneOfModel.Dkim).To(Equal(dkimAttributesModel))
				Expect(destinationConfigOneOfModel.Spf).To(Equal(spfAttributesModel))

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				Expect(destinationConfigModel).ToNot(BeNil())
				destinationConfigModel.Params = destinationConfigOneOfModel
				Expect(destinationConfigModel.Params).To(Equal(destinationConfigOneOfModel))

				// Construct an instance of the CreateDestinationOptions model
				instanceID := "testString"
				name := "testString"
				typeVar := "webhook"
				createDestinationOptionsModel := eventNotificationsService.NewCreateDestinationOptions(instanceID, name, typeVar)
				createDestinationOptionsModel.SetInstanceID("testString")
				createDestinationOptionsModel.SetName("testString")
				createDestinationOptionsModel.SetType("webhook")
				createDestinationOptionsModel.SetDescription("testString")
				createDestinationOptionsModel.SetConfig(destinationConfigModel)
				createDestinationOptionsModel.SetCertificate(CreateMockReader("This is a mock file."))
				createDestinationOptionsModel.SetCertificateContentType("testString")
				createDestinationOptionsModel.SetIcon16x16(CreateMockReader("This is a mock file."))
				createDestinationOptionsModel.SetIcon16x16ContentType("testString")
				createDestinationOptionsModel.SetIcon16x162x(CreateMockReader("This is a mock file."))
				createDestinationOptionsModel.SetIcon16x162xContentType("testString")
				createDestinationOptionsModel.SetIcon32x32(CreateMockReader("This is a mock file."))
				createDestinationOptionsModel.SetIcon32x32ContentType("testString")
				createDestinationOptionsModel.SetIcon32x322x(CreateMockReader("This is a mock file."))
				createDestinationOptionsModel.SetIcon32x322xContentType("testString")
				createDestinationOptionsModel.SetIcon128x128(CreateMockReader("This is a mock file."))
				createDestinationOptionsModel.SetIcon128x128ContentType("testString")
				createDestinationOptionsModel.SetIcon128x1282x(CreateMockReader("This is a mock file."))
				createDestinationOptionsModel.SetIcon128x1282xContentType("testString")
				createDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDestinationOptionsModel).ToNot(BeNil())
				Expect(createDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Type).To(Equal(core.StringPtr("webhook")))
				Expect(createDestinationOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Config).To(Equal(destinationConfigModel))
				Expect(createDestinationOptionsModel.Certificate).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createDestinationOptionsModel.CertificateContentType).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Icon16x16).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createDestinationOptionsModel.Icon16x16ContentType).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Icon16x162x).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createDestinationOptionsModel.Icon16x162xContentType).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Icon32x32).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createDestinationOptionsModel.Icon32x32ContentType).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Icon32x322x).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createDestinationOptionsModel.Icon32x322xContentType).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Icon128x128).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createDestinationOptionsModel.Icon128x128ContentType).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Icon128x1282x).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createDestinationOptionsModel.Icon128x1282xContentType).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSourcesOptions successfully`, func() {
				// Construct an instance of the CreateSourcesOptions model
				instanceID := "testString"
				createSourcesOptionsName := "testString"
				createSourcesOptionsDescription := "testString"
				createSourcesOptionsModel := eventNotificationsService.NewCreateSourcesOptions(instanceID, createSourcesOptionsName, createSourcesOptionsDescription)
				createSourcesOptionsModel.SetInstanceID("testString")
				createSourcesOptionsModel.SetName("testString")
				createSourcesOptionsModel.SetDescription("testString")
				createSourcesOptionsModel.SetEnabled(true)
				createSourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSourcesOptionsModel).ToNot(BeNil())
				Expect(createSourcesOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createSourcesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createSourcesOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createSourcesOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createSourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSubscriptionOptions successfully`, func() {
				// Construct an instance of the SubscriptionCreateAttributesSmsAttributes model
				subscriptionCreateAttributesModel := new(eventnotificationsv1.SubscriptionCreateAttributesSmsAttributes)
				Expect(subscriptionCreateAttributesModel).ToNot(BeNil())
				subscriptionCreateAttributesModel.Invited = []string{"testString"}
				Expect(subscriptionCreateAttributesModel.Invited).To(Equal([]string{"testString"}))

				// Construct an instance of the CreateSubscriptionOptions model
				instanceID := "testString"
				createSubscriptionOptionsName := "testString"
				createSubscriptionOptionsDestinationID := "testString"
				createSubscriptionOptionsTopicID := "testString"
				createSubscriptionOptionsModel := eventNotificationsService.NewCreateSubscriptionOptions(instanceID, createSubscriptionOptionsName, createSubscriptionOptionsDestinationID, createSubscriptionOptionsTopicID)
				createSubscriptionOptionsModel.SetInstanceID("testString")
				createSubscriptionOptionsModel.SetName("testString")
				createSubscriptionOptionsModel.SetDestinationID("testString")
				createSubscriptionOptionsModel.SetTopicID("testString")
				createSubscriptionOptionsModel.SetDescription("testString")
				createSubscriptionOptionsModel.SetAttributes(subscriptionCreateAttributesModel)
				createSubscriptionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSubscriptionOptionsModel).ToNot(BeNil())
				Expect(createSubscriptionOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createSubscriptionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createSubscriptionOptionsModel.DestinationID).To(Equal(core.StringPtr("testString")))
				Expect(createSubscriptionOptionsModel.TopicID).To(Equal(core.StringPtr("testString")))
				Expect(createSubscriptionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createSubscriptionOptionsModel.Attributes).To(Equal(subscriptionCreateAttributesModel))
				Expect(createSubscriptionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTagsSubscriptionOptions successfully`, func() {
				// Construct an instance of the CreateTagsSubscriptionOptions model
				instanceID := "testString"
				id := "testString"
				createTagsSubscriptionOptionsDeviceID := "testString"
				createTagsSubscriptionOptionsTagName := "testString"
				createTagsSubscriptionOptionsModel := eventNotificationsService.NewCreateTagsSubscriptionOptions(instanceID, id, createTagsSubscriptionOptionsDeviceID, createTagsSubscriptionOptionsTagName)
				createTagsSubscriptionOptionsModel.SetInstanceID("testString")
				createTagsSubscriptionOptionsModel.SetID("testString")
				createTagsSubscriptionOptionsModel.SetDeviceID("testString")
				createTagsSubscriptionOptionsModel.SetTagName("testString")
				createTagsSubscriptionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTagsSubscriptionOptionsModel).ToNot(BeNil())
				Expect(createTagsSubscriptionOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createTagsSubscriptionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createTagsSubscriptionOptionsModel.DeviceID).To(Equal(core.StringPtr("testString")))
				Expect(createTagsSubscriptionOptionsModel.TagName).To(Equal(core.StringPtr("testString")))
				Expect(createTagsSubscriptionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTemplateOptions successfully`, func() {
				// Construct an instance of the TemplateConfig model
				templateConfigModel := new(eventnotificationsv1.TemplateConfig)
				Expect(templateConfigModel).ToNot(BeNil())
				templateConfigModel.Body = core.StringPtr("testString")
				templateConfigModel.Subject = core.StringPtr("testString")
				Expect(templateConfigModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(templateConfigModel.Subject).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateTemplateOptions model
				instanceID := "testString"
				createTemplateOptionsName := "testString"
				createTemplateOptionsType := "smtp_custom.notification"
				var createTemplateOptionsParams *eventnotificationsv1.TemplateConfig = nil
				createTemplateOptionsModel := eventNotificationsService.NewCreateTemplateOptions(instanceID, createTemplateOptionsName, createTemplateOptionsType, createTemplateOptionsParams)
				createTemplateOptionsModel.SetInstanceID("testString")
				createTemplateOptionsModel.SetName("testString")
				createTemplateOptionsModel.SetType("smtp_custom.notification")
				createTemplateOptionsModel.SetParams(templateConfigModel)
				createTemplateOptionsModel.SetDescription("testString")
				createTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTemplateOptionsModel).ToNot(BeNil())
				Expect(createTemplateOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createTemplateOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createTemplateOptionsModel.Type).To(Equal(core.StringPtr("smtp_custom.notification")))
				Expect(createTemplateOptionsModel.Params).To(Equal(templateConfigModel))
				Expect(createTemplateOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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

				// Construct an instance of the SourcesItems model
				sourcesItemsModel := new(eventnotificationsv1.SourcesItems)
				Expect(sourcesItemsModel).ToNot(BeNil())
				sourcesItemsModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				sourcesItemsModel.Rules = []eventnotificationsv1.Rules{*rulesModel}
				Expect(sourcesItemsModel.ID).To(Equal(core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")))
				Expect(sourcesItemsModel.Rules).To(Equal([]eventnotificationsv1.Rules{*rulesModel}))

				// Construct an instance of the CreateTopicOptions model
				instanceID := "testString"
				createTopicOptionsName := "testString"
				createTopicOptionsModel := eventNotificationsService.NewCreateTopicOptions(instanceID, createTopicOptionsName)
				createTopicOptionsModel.SetInstanceID("testString")
				createTopicOptionsModel.SetName("testString")
				createTopicOptionsModel.SetDescription("testString")
				createTopicOptionsModel.SetSources([]eventnotificationsv1.SourcesItems{*sourcesItemsModel})
				createTopicOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTopicOptionsModel).ToNot(BeNil())
				Expect(createTopicOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createTopicOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createTopicOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createTopicOptionsModel.Sources).To(Equal([]eventnotificationsv1.SourcesItems{*sourcesItemsModel}))
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
			It(`Invoke NewDeleteSourceOptions successfully`, func() {
				// Construct an instance of the DeleteSourceOptions model
				instanceID := "testString"
				id := "testString"
				deleteSourceOptionsModel := eventNotificationsService.NewDeleteSourceOptions(instanceID, id)
				deleteSourceOptionsModel.SetInstanceID("testString")
				deleteSourceOptionsModel.SetID("testString")
				deleteSourceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSourceOptionsModel).ToNot(BeNil())
				Expect(deleteSourceOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSourceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSourceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewDeleteTagsSubscriptionOptions successfully`, func() {
				// Construct an instance of the DeleteTagsSubscriptionOptions model
				instanceID := "testString"
				id := "testString"
				deleteTagsSubscriptionOptionsModel := eventNotificationsService.NewDeleteTagsSubscriptionOptions(instanceID, id)
				deleteTagsSubscriptionOptionsModel.SetInstanceID("testString")
				deleteTagsSubscriptionOptionsModel.SetID("testString")
				deleteTagsSubscriptionOptionsModel.SetDeviceID("testString")
				deleteTagsSubscriptionOptionsModel.SetTagName("testString")
				deleteTagsSubscriptionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTagsSubscriptionOptionsModel).ToNot(BeNil())
				Expect(deleteTagsSubscriptionOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTagsSubscriptionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTagsSubscriptionOptionsModel.DeviceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTagsSubscriptionOptionsModel.TagName).To(Equal(core.StringPtr("testString")))
				Expect(deleteTagsSubscriptionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTemplateOptions successfully`, func() {
				// Construct an instance of the DeleteTemplateOptions model
				instanceID := "testString"
				id := "testString"
				deleteTemplateOptionsModel := eventNotificationsService.NewDeleteTemplateOptions(instanceID, id)
				deleteTemplateOptionsModel.SetInstanceID("testString")
				deleteTemplateOptionsModel.SetID("testString")
				deleteTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTemplateOptionsModel).ToNot(BeNil())
				Expect(deleteTemplateOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				var params eventnotificationsv1.DestinationConfigOneOfIntf = nil
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
			It(`Invoke NewGetIntegrationOptions successfully`, func() {
				// Construct an instance of the GetIntegrationOptions model
				instanceID := "testString"
				id := "testString"
				getIntegrationOptionsModel := eventNotificationsService.NewGetIntegrationOptions(instanceID, id)
				getIntegrationOptionsModel.SetInstanceID("testString")
				getIntegrationOptionsModel.SetID("testString")
				getIntegrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getIntegrationOptionsModel).ToNot(BeNil())
				Expect(getIntegrationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getIntegrationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getIntegrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewGetTemplateOptions successfully`, func() {
				// Construct an instance of the GetTemplateOptions model
				instanceID := "testString"
				id := "testString"
				getTemplateOptionsModel := eventNotificationsService.NewGetTemplateOptions(instanceID, id)
				getTemplateOptionsModel.SetInstanceID("testString")
				getTemplateOptionsModel.SetID("testString")
				getTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTemplateOptionsModel).ToNot(BeNil())
				Expect(getTemplateOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewIntegrationMetadata successfully`, func() {
				endpoint := "testString"
				crn := "testString"
				rootKeyID := "testString"
				_model, err := eventNotificationsService.NewIntegrationMetadata(endpoint, crn, rootKeyID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListDestinationsOptions successfully`, func() {
				// Construct an instance of the ListDestinationsOptions model
				instanceID := "testString"
				listDestinationsOptionsModel := eventNotificationsService.NewListDestinationsOptions(instanceID)
				listDestinationsOptionsModel.SetInstanceID("testString")
				listDestinationsOptionsModel.SetLimit(int64(10))
				listDestinationsOptionsModel.SetOffset(int64(0))
				listDestinationsOptionsModel.SetSearch("testString")
				listDestinationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDestinationsOptionsModel).ToNot(BeNil())
				Expect(listDestinationsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listDestinationsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listDestinationsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listDestinationsOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listDestinationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListIntegrationsOptions successfully`, func() {
				// Construct an instance of the ListIntegrationsOptions model
				instanceID := "testString"
				listIntegrationsOptionsModel := eventNotificationsService.NewListIntegrationsOptions(instanceID)
				listIntegrationsOptionsModel.SetInstanceID("testString")
				listIntegrationsOptionsModel.SetOffset(int64(0))
				listIntegrationsOptionsModel.SetLimit(int64(10))
				listIntegrationsOptionsModel.SetSearch("testString")
				listIntegrationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listIntegrationsOptionsModel).ToNot(BeNil())
				Expect(listIntegrationsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listIntegrationsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listIntegrationsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listIntegrationsOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listIntegrationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSourcesOptions successfully`, func() {
				// Construct an instance of the ListSourcesOptions model
				instanceID := "testString"
				listSourcesOptionsModel := eventNotificationsService.NewListSourcesOptions(instanceID)
				listSourcesOptionsModel.SetInstanceID("testString")
				listSourcesOptionsModel.SetLimit(int64(10))
				listSourcesOptionsModel.SetOffset(int64(0))
				listSourcesOptionsModel.SetSearch("testString")
				listSourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSourcesOptionsModel).ToNot(BeNil())
				Expect(listSourcesOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listSourcesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
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
				listSubscriptionsOptionsModel.SetLimit(int64(10))
				listSubscriptionsOptionsModel.SetSearch("testString")
				listSubscriptionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSubscriptionsOptionsModel).ToNot(BeNil())
				Expect(listSubscriptionsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listSubscriptionsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listSubscriptionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listSubscriptionsOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listSubscriptionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTagsSubscriptionOptions successfully`, func() {
				// Construct an instance of the ListTagsSubscriptionOptions model
				instanceID := "testString"
				id := "testString"
				listTagsSubscriptionOptionsModel := eventNotificationsService.NewListTagsSubscriptionOptions(instanceID, id)
				listTagsSubscriptionOptionsModel.SetInstanceID("testString")
				listTagsSubscriptionOptionsModel.SetID("testString")
				listTagsSubscriptionOptionsModel.SetDeviceID("testString")
				listTagsSubscriptionOptionsModel.SetUserID("testString")
				listTagsSubscriptionOptionsModel.SetTagName("testString")
				listTagsSubscriptionOptionsModel.SetLimit(int64(10))
				listTagsSubscriptionOptionsModel.SetOffset(int64(0))
				listTagsSubscriptionOptionsModel.SetSearch("testString")
				listTagsSubscriptionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTagsSubscriptionOptionsModel).ToNot(BeNil())
				Expect(listTagsSubscriptionOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionOptionsModel.DeviceID).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionOptionsModel.TagName).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listTagsSubscriptionOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listTagsSubscriptionOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTemplatesOptions successfully`, func() {
				// Construct an instance of the ListTemplatesOptions model
				instanceID := "testString"
				listTemplatesOptionsModel := eventNotificationsService.NewListTemplatesOptions(instanceID)
				listTemplatesOptionsModel.SetInstanceID("testString")
				listTemplatesOptionsModel.SetLimit(int64(10))
				listTemplatesOptionsModel.SetOffset(int64(0))
				listTemplatesOptionsModel.SetSearch("testString")
				listTemplatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTemplatesOptionsModel).ToNot(BeNil())
				Expect(listTemplatesOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listTemplatesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listTemplatesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listTemplatesOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listTemplatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTopicsOptions successfully`, func() {
				// Construct an instance of the ListTopicsOptions model
				instanceID := "testString"
				listTopicsOptionsModel := eventNotificationsService.NewListTopicsOptions(instanceID)
				listTopicsOptionsModel.SetInstanceID("testString")
				listTopicsOptionsModel.SetLimit(int64(10))
				listTopicsOptionsModel.SetOffset(int64(0))
				listTopicsOptionsModel.SetSearch("testString")
				listTopicsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTopicsOptionsModel).ToNot(BeNil())
				Expect(listTopicsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listTopicsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listTopicsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listTopicsOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listTopicsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewNotificationCreate successfully`, func() {
				specversion := "1.0"
				id := "testString"
				source := "testString"
				typeVar := "testString"
				ibmensourceid := "testString"
				ibmendefaultshort := "testString"
				ibmendefaultlong := "testString"
				_model, err := eventNotificationsService.NewNotificationCreate(specversion, id, source, typeVar, ibmensourceid, ibmendefaultshort, ibmendefaultlong)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewReplaceIntegrationOptions successfully`, func() {
				// Construct an instance of the IntegrationMetadata model
				integrationMetadataModel := new(eventnotificationsv1.IntegrationMetadata)
				Expect(integrationMetadataModel).ToNot(BeNil())
				integrationMetadataModel.Endpoint = core.StringPtr("testString")
				integrationMetadataModel.CRN = core.StringPtr("testString")
				integrationMetadataModel.RootKeyID = core.StringPtr("testString")
				Expect(integrationMetadataModel.Endpoint).To(Equal(core.StringPtr("testString")))
				Expect(integrationMetadataModel.CRN).To(Equal(core.StringPtr("testString")))
				Expect(integrationMetadataModel.RootKeyID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplaceIntegrationOptions model
				instanceID := "testString"
				id := "testString"
				replaceIntegrationOptionsType := "testString"
				var replaceIntegrationOptionsMetadata *eventnotificationsv1.IntegrationMetadata = nil
				replaceIntegrationOptionsModel := eventNotificationsService.NewReplaceIntegrationOptions(instanceID, id, replaceIntegrationOptionsType, replaceIntegrationOptionsMetadata)
				replaceIntegrationOptionsModel.SetInstanceID("testString")
				replaceIntegrationOptionsModel.SetID("testString")
				replaceIntegrationOptionsModel.SetType("testString")
				replaceIntegrationOptionsModel.SetMetadata(integrationMetadataModel)
				replaceIntegrationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceIntegrationOptionsModel).ToNot(BeNil())
				Expect(replaceIntegrationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceIntegrationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceIntegrationOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(replaceIntegrationOptionsModel.Metadata).To(Equal(integrationMetadataModel))
				Expect(replaceIntegrationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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

				// Construct an instance of the SourcesItems model
				sourcesItemsModel := new(eventnotificationsv1.SourcesItems)
				Expect(sourcesItemsModel).ToNot(BeNil())
				sourcesItemsModel.ID = core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")
				sourcesItemsModel.Rules = []eventnotificationsv1.Rules{*rulesModel}
				Expect(sourcesItemsModel.ID).To(Equal(core.StringPtr("e7c3b3ee-78d9-4e02-95c3-c001a05e6ea5:api")))
				Expect(sourcesItemsModel.Rules).To(Equal([]eventnotificationsv1.Rules{*rulesModel}))

				// Construct an instance of the ReplaceTopicOptions model
				instanceID := "testString"
				id := "testString"
				replaceTopicOptionsModel := eventNotificationsService.NewReplaceTopicOptions(instanceID, id)
				replaceTopicOptionsModel.SetInstanceID("testString")
				replaceTopicOptionsModel.SetID("testString")
				replaceTopicOptionsModel.SetName("testString")
				replaceTopicOptionsModel.SetDescription("testString")
				replaceTopicOptionsModel.SetSources([]eventnotificationsv1.SourcesItems{*sourcesItemsModel})
				replaceTopicOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceTopicOptionsModel).ToNot(BeNil())
				Expect(replaceTopicOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceTopicOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceTopicOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(replaceTopicOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(replaceTopicOptionsModel.Sources).To(Equal([]eventnotificationsv1.SourcesItems{*sourcesItemsModel}))
				Expect(replaceTopicOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRules successfully`, func() {
				eventTypeFilter := "$.*"
				_model, err := eventNotificationsService.NewRules(eventTypeFilter)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSendBulkNotificationsOptions successfully`, func() {
				// Construct an instance of the NotificationCreate model
				notificationCreateModel := new(eventnotificationsv1.NotificationCreate)
				Expect(notificationCreateModel).ToNot(BeNil())
				notificationCreateModel.Specversion = core.StringPtr("1.0")
				notificationCreateModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				notificationCreateModel.ID = core.StringPtr("testString")
				notificationCreateModel.Source = core.StringPtr("testString")
				notificationCreateModel.Type = core.StringPtr("testString")
				notificationCreateModel.Ibmenseverity = core.StringPtr("testString")
				notificationCreateModel.Ibmensourceid = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultshort = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultlong = core.StringPtr("testString")
				notificationCreateModel.Subject = core.StringPtr("testString")
				notificationCreateModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				notificationCreateModel.Datacontenttype = core.StringPtr("application/json")
				notificationCreateModel.Ibmenpushto = core.StringPtr(`{"platforms":["push_android"]}`)
				notificationCreateModel.Ibmenfcmbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsheaders = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromebody = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromeheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenfirefoxbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenfirefoxheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenhuaweibody = core.StringPtr("testString")
				notificationCreateModel.Ibmensafaribody = core.StringPtr("testString")
				notificationCreateModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(notificationCreateModel.Specversion).To(Equal(core.StringPtr("1.0")))
				Expect(notificationCreateModel.Time).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(notificationCreateModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Source).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmenseverity).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmensourceid).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmendefaultshort).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmendefaultlong).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Subject).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Data).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(notificationCreateModel.Datacontenttype).To(Equal(core.StringPtr("application/json")))
				Expect(notificationCreateModel.Ibmenpushto).To(Equal(core.StringPtr(`{"platforms":["push_android"]}`)))
				Expect(notificationCreateModel.Ibmenfcmbody).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmenapnsbody).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmenapnsheaders).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmenchromebody).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmenchromeheaders).To(Equal(core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)))
				Expect(notificationCreateModel.Ibmenfirefoxbody).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmenfirefoxheaders).To(Equal(core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)))
				Expect(notificationCreateModel.Ibmenhuaweibody).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmensafaribody).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.GetProperties()).ToNot(BeEmpty())
				Expect(notificationCreateModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				notificationCreateModel.SetProperties(nil)
				Expect(notificationCreateModel.GetProperties()).To(BeEmpty())

				notificationCreateModelExpectedMap := make(map[string]interface{})
				notificationCreateModelExpectedMap["foo"] = core.StringPtr("testString")
				notificationCreateModel.SetProperties(notificationCreateModelExpectedMap)
				notificationCreateModelActualMap := notificationCreateModel.GetProperties()
				Expect(notificationCreateModelActualMap).To(Equal(notificationCreateModelExpectedMap))

				// Construct an instance of the SendBulkNotificationsOptions model
				instanceID := "testString"
				sendBulkNotificationsOptionsModel := eventNotificationsService.NewSendBulkNotificationsOptions(instanceID)
				sendBulkNotificationsOptionsModel.SetInstanceID("testString")
				sendBulkNotificationsOptionsModel.SetBulkMessages([]eventnotificationsv1.NotificationCreate{*notificationCreateModel})
				sendBulkNotificationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(sendBulkNotificationsOptionsModel).ToNot(BeNil())
				Expect(sendBulkNotificationsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(sendBulkNotificationsOptionsModel.BulkMessages).To(Equal([]eventnotificationsv1.NotificationCreate{*notificationCreateModel}))
				Expect(sendBulkNotificationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSendNotificationsOptions successfully`, func() {
				// Construct an instance of the NotificationCreate model
				notificationCreateModel := new(eventnotificationsv1.NotificationCreate)
				Expect(notificationCreateModel).ToNot(BeNil())
				notificationCreateModel.Specversion = core.StringPtr("1.0")
				notificationCreateModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				notificationCreateModel.ID = core.StringPtr("testString")
				notificationCreateModel.Source = core.StringPtr("testString")
				notificationCreateModel.Type = core.StringPtr("testString")
				notificationCreateModel.Ibmenseverity = core.StringPtr("testString")
				notificationCreateModel.Ibmensourceid = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultshort = core.StringPtr("testString")
				notificationCreateModel.Ibmendefaultlong = core.StringPtr("testString")
				notificationCreateModel.Subject = core.StringPtr("testString")
				notificationCreateModel.Data = map[string]interface{}{"anyKey": "anyValue"}
				notificationCreateModel.Datacontenttype = core.StringPtr("application/json")
				notificationCreateModel.Ibmenpushto = core.StringPtr(`{"platforms":["push_android"]}`)
				notificationCreateModel.Ibmenfcmbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenapnsheaders = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromebody = core.StringPtr("testString")
				notificationCreateModel.Ibmenchromeheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenfirefoxbody = core.StringPtr("testString")
				notificationCreateModel.Ibmenfirefoxheaders = core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)
				notificationCreateModel.Ibmenhuaweibody = core.StringPtr("testString")
				notificationCreateModel.Ibmensafaribody = core.StringPtr("testString")
				notificationCreateModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(notificationCreateModel.Specversion).To(Equal(core.StringPtr("1.0")))
				Expect(notificationCreateModel.Time).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(notificationCreateModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Source).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmenseverity).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmensourceid).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmendefaultshort).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmendefaultlong).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Subject).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Data).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(notificationCreateModel.Datacontenttype).To(Equal(core.StringPtr("application/json")))
				Expect(notificationCreateModel.Ibmenpushto).To(Equal(core.StringPtr(`{"platforms":["push_android"]}`)))
				Expect(notificationCreateModel.Ibmenfcmbody).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmenapnsbody).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmenapnsheaders).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmenchromebody).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmenchromeheaders).To(Equal(core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)))
				Expect(notificationCreateModel.Ibmenfirefoxbody).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmenfirefoxheaders).To(Equal(core.StringPtr(`{"TTL":3600,"Topic":"test","Urgency":"high"}`)))
				Expect(notificationCreateModel.Ibmenhuaweibody).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.Ibmensafaribody).To(Equal(core.StringPtr("testString")))
				Expect(notificationCreateModel.GetProperties()).ToNot(BeEmpty())
				Expect(notificationCreateModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				notificationCreateModel.SetProperties(nil)
				Expect(notificationCreateModel.GetProperties()).To(BeEmpty())

				notificationCreateModelExpectedMap := make(map[string]interface{})
				notificationCreateModelExpectedMap["foo"] = core.StringPtr("testString")
				notificationCreateModel.SetProperties(notificationCreateModelExpectedMap)
				notificationCreateModelActualMap := notificationCreateModel.GetProperties()
				Expect(notificationCreateModelActualMap).To(Equal(notificationCreateModelExpectedMap))

				// Construct an instance of the SendNotificationsOptions model
				instanceID := "testString"
				sendNotificationsOptionsModel := eventNotificationsService.NewSendNotificationsOptions(instanceID)
				sendNotificationsOptionsModel.SetInstanceID("testString")
				sendNotificationsOptionsModel.SetBody(notificationCreateModel)
				sendNotificationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(sendNotificationsOptionsModel).ToNot(BeNil())
				Expect(sendNotificationsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsOptionsModel.Body).To(Equal(notificationCreateModel))
				Expect(sendNotificationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSourcesItems successfully`, func() {
				id := "testString"
				rules := []eventnotificationsv1.Rules{}
				_model, err := eventNotificationsService.NewSourcesItems(id, rules)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTemplateConfig successfully`, func() {
				body := "testString"
				subject := "testString"
				_model, err := eventNotificationsService.NewTemplateConfig(body, subject)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateDestinationOptions successfully`, func() {
				// Construct an instance of the DkimAttributes model
				dkimAttributesModel := new(eventnotificationsv1.DkimAttributes)
				Expect(dkimAttributesModel).ToNot(BeNil())
				dkimAttributesModel.PublicKey = core.StringPtr("testString")
				dkimAttributesModel.Selector = core.StringPtr("testString")
				dkimAttributesModel.Verification = core.StringPtr("testString")
				Expect(dkimAttributesModel.PublicKey).To(Equal(core.StringPtr("testString")))
				Expect(dkimAttributesModel.Selector).To(Equal(core.StringPtr("testString")))
				Expect(dkimAttributesModel.Verification).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SpfAttributes model
				spfAttributesModel := new(eventnotificationsv1.SpfAttributes)
				Expect(spfAttributesModel).ToNot(BeNil())
				spfAttributesModel.TxtName = core.StringPtr("testString")
				spfAttributesModel.TxtValue = core.StringPtr("testString")
				spfAttributesModel.Verification = core.StringPtr("testString")
				Expect(spfAttributesModel.TxtName).To(Equal(core.StringPtr("testString")))
				Expect(spfAttributesModel.TxtValue).To(Equal(core.StringPtr("testString")))
				Expect(spfAttributesModel.Verification).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DestinationConfigOneOfCustomDomainEmailDestinationConfig model
				destinationConfigOneOfModel := new(eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig)
				Expect(destinationConfigOneOfModel).ToNot(BeNil())
				destinationConfigOneOfModel.Domain = core.StringPtr("testString")
				destinationConfigOneOfModel.Dkim = dkimAttributesModel
				destinationConfigOneOfModel.Spf = spfAttributesModel
				Expect(destinationConfigOneOfModel.Domain).To(Equal(core.StringPtr("testString")))
				Expect(destinationConfigOneOfModel.Dkim).To(Equal(dkimAttributesModel))
				Expect(destinationConfigOneOfModel.Spf).To(Equal(spfAttributesModel))

				// Construct an instance of the DestinationConfig model
				destinationConfigModel := new(eventnotificationsv1.DestinationConfig)
				Expect(destinationConfigModel).ToNot(BeNil())
				destinationConfigModel.Params = destinationConfigOneOfModel
				Expect(destinationConfigModel.Params).To(Equal(destinationConfigOneOfModel))

				// Construct an instance of the UpdateDestinationOptions model
				instanceID := "testString"
				id := "testString"
				updateDestinationOptionsModel := eventNotificationsService.NewUpdateDestinationOptions(instanceID, id)
				updateDestinationOptionsModel.SetInstanceID("testString")
				updateDestinationOptionsModel.SetID("testString")
				updateDestinationOptionsModel.SetName("testString")
				updateDestinationOptionsModel.SetDescription("testString")
				updateDestinationOptionsModel.SetConfig(destinationConfigModel)
				updateDestinationOptionsModel.SetCertificate(CreateMockReader("This is a mock file."))
				updateDestinationOptionsModel.SetCertificateContentType("testString")
				updateDestinationOptionsModel.SetIcon16x16(CreateMockReader("This is a mock file."))
				updateDestinationOptionsModel.SetIcon16x16ContentType("testString")
				updateDestinationOptionsModel.SetIcon16x162x(CreateMockReader("This is a mock file."))
				updateDestinationOptionsModel.SetIcon16x162xContentType("testString")
				updateDestinationOptionsModel.SetIcon32x32(CreateMockReader("This is a mock file."))
				updateDestinationOptionsModel.SetIcon32x32ContentType("testString")
				updateDestinationOptionsModel.SetIcon32x322x(CreateMockReader("This is a mock file."))
				updateDestinationOptionsModel.SetIcon32x322xContentType("testString")
				updateDestinationOptionsModel.SetIcon128x128(CreateMockReader("This is a mock file."))
				updateDestinationOptionsModel.SetIcon128x128ContentType("testString")
				updateDestinationOptionsModel.SetIcon128x1282x(CreateMockReader("This is a mock file."))
				updateDestinationOptionsModel.SetIcon128x1282xContentType("testString")
				updateDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDestinationOptionsModel).ToNot(BeNil())
				Expect(updateDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Config).To(Equal(destinationConfigModel))
				Expect(updateDestinationOptionsModel.Certificate).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateDestinationOptionsModel.CertificateContentType).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Icon16x16).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateDestinationOptionsModel.Icon16x16ContentType).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Icon16x162x).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateDestinationOptionsModel.Icon16x162xContentType).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Icon32x32).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateDestinationOptionsModel.Icon32x32ContentType).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Icon32x322x).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateDestinationOptionsModel.Icon32x322xContentType).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Icon128x128).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateDestinationOptionsModel.Icon128x128ContentType).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Icon128x1282x).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateDestinationOptionsModel.Icon128x1282xContentType).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSourceOptions successfully`, func() {
				// Construct an instance of the UpdateSourceOptions model
				instanceID := "testString"
				id := "testString"
				updateSourceOptionsModel := eventNotificationsService.NewUpdateSourceOptions(instanceID, id)
				updateSourceOptionsModel.SetInstanceID("testString")
				updateSourceOptionsModel.SetID("testString")
				updateSourceOptionsModel.SetName("testString")
				updateSourceOptionsModel.SetDescription("testString")
				updateSourceOptionsModel.SetEnabled(true)
				updateSourceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSourceOptionsModel).ToNot(BeNil())
				Expect(updateSourceOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateSourceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateSourceOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateSourceOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateSourceOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updateSourceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSubscriptionOptions successfully`, func() {
				// Construct an instance of the UpdateAttributesInvited model
				updateAttributesInvitedModel := new(eventnotificationsv1.UpdateAttributesInvited)
				Expect(updateAttributesInvitedModel).ToNot(BeNil())
				updateAttributesInvitedModel.Add = []string{"testString"}
				updateAttributesInvitedModel.Remove = []string{"testString"}
				Expect(updateAttributesInvitedModel.Add).To(Equal([]string{"testString"}))
				Expect(updateAttributesInvitedModel.Remove).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateAttributesSubscribed model
				updateAttributesSubscribedModel := new(eventnotificationsv1.UpdateAttributesSubscribed)
				Expect(updateAttributesSubscribedModel).ToNot(BeNil())
				updateAttributesSubscribedModel.Remove = []string{"testString"}
				Expect(updateAttributesSubscribedModel.Remove).To(Equal([]string{"testString"}))

				// Construct an instance of the UpdateAttributesUnsubscribed model
				updateAttributesUnsubscribedModel := new(eventnotificationsv1.UpdateAttributesUnsubscribed)
				Expect(updateAttributesUnsubscribedModel).ToNot(BeNil())
				updateAttributesUnsubscribedModel.Remove = []string{"testString"}
				Expect(updateAttributesUnsubscribedModel.Remove).To(Equal([]string{"testString"}))

				// Construct an instance of the SubscriptionUpdateAttributesSmsUpdateAttributes model
				subscriptionUpdateAttributesModel := new(eventnotificationsv1.SubscriptionUpdateAttributesSmsUpdateAttributes)
				Expect(subscriptionUpdateAttributesModel).ToNot(BeNil())
				subscriptionUpdateAttributesModel.Invited = updateAttributesInvitedModel
				subscriptionUpdateAttributesModel.Subscribed = updateAttributesSubscribedModel
				subscriptionUpdateAttributesModel.Unsubscribed = updateAttributesUnsubscribedModel
				Expect(subscriptionUpdateAttributesModel.Invited).To(Equal(updateAttributesInvitedModel))
				Expect(subscriptionUpdateAttributesModel.Subscribed).To(Equal(updateAttributesSubscribedModel))
				Expect(subscriptionUpdateAttributesModel.Unsubscribed).To(Equal(updateAttributesUnsubscribedModel))

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
			It(`Invoke NewUpdateTemplateOptions successfully`, func() {
				// Construct an instance of the TemplateConfig model
				templateConfigModel := new(eventnotificationsv1.TemplateConfig)
				Expect(templateConfigModel).ToNot(BeNil())
				templateConfigModel.Body = core.StringPtr("testString")
				templateConfigModel.Subject = core.StringPtr("testString")
				Expect(templateConfigModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(templateConfigModel.Subject).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateTemplateOptions model
				instanceID := "testString"
				id := "testString"
				updateTemplateOptionsModel := eventNotificationsService.NewUpdateTemplateOptions(instanceID, id)
				updateTemplateOptionsModel.SetInstanceID("testString")
				updateTemplateOptionsModel.SetID("testString")
				updateTemplateOptionsModel.SetName("testString")
				updateTemplateOptionsModel.SetDescription("testString")
				updateTemplateOptionsModel.SetType("smtp_custom.notification")
				updateTemplateOptionsModel.SetParams(templateConfigModel)
				updateTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTemplateOptionsModel).ToNot(BeNil())
				Expect(updateTemplateOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateTemplateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateTemplateOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateTemplateOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateTemplateOptionsModel.Type).To(Equal(core.StringPtr("smtp_custom.notification")))
				Expect(updateTemplateOptionsModel.Params).To(Equal(templateConfigModel))
				Expect(updateTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateVerifyDestinationOptions successfully`, func() {
				// Construct an instance of the UpdateVerifyDestinationOptions model
				instanceID := "testString"
				id := "testString"
				typeVar := "testString"
				updateVerifyDestinationOptionsModel := eventNotificationsService.NewUpdateVerifyDestinationOptions(instanceID, id, typeVar)
				updateVerifyDestinationOptionsModel.SetInstanceID("testString")
				updateVerifyDestinationOptionsModel.SetID("testString")
				updateVerifyDestinationOptionsModel.SetType("testString")
				updateVerifyDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateVerifyDestinationOptionsModel).ToNot(BeNil())
				Expect(updateVerifyDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateVerifyDestinationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateVerifyDestinationOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(updateVerifyDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDestinationConfigOneOfChromeDestinationConfig successfully`, func() {
				apiKey := "testString"
				websiteURL := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigOneOfChromeDestinationConfig(apiKey, websiteURL)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigOneOfCustomDomainEmailDestinationConfig successfully`, func() {
				domain := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigOneOfCustomDomainEmailDestinationConfig(domain)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigOneOfFirefoxDestinationConfig successfully`, func() {
				websiteURL := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigOneOfFirefoxDestinationConfig(websiteURL)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigOneOfHuaweiDestinationConfig successfully`, func() {
				clientID := "testString"
				clientSecret := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigOneOfHuaweiDestinationConfig(clientID, clientSecret)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigOneOfIBMCloudFunctionsDestinationConfig successfully`, func() {
				url := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigOneOfIBMCloudFunctionsDestinationConfig(url)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigOneOfIBMCloudObjectStorageDestinationConfig successfully`, func() {
				bucketName := "testString"
				instanceID := "testString"
				endpoint := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigOneOfIBMCloudObjectStorageDestinationConfig(bucketName, instanceID, endpoint)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigOneOfIosDestinationConfig successfully`, func() {
				certType := "p8"
				isSandbox := false
				_model, err := eventNotificationsService.NewDestinationConfigOneOfIosDestinationConfig(certType, isSandbox)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigOneOfMsTeamsDestinationConfig successfully`, func() {
				url := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigOneOfMsTeamsDestinationConfig(url)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigOneOfPagerDutyDestinationConfig successfully`, func() {
				apiKey := "testString"
				routingKey := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigOneOfPagerDutyDestinationConfig(apiKey, routingKey)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigOneOfSafariDestinationConfig successfully`, func() {
				certType := "p12"
				password := "testString"
				websiteURL := "testString"
				websiteName := "testString"
				urlFormatString := "testString"
				websitePushID := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigOneOfSafariDestinationConfig(certType, password, websiteURL, websiteName, urlFormatString, websitePushID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigOneOfServiceNowDestinationConfig successfully`, func() {
				clientID := "testString"
				clientSecret := "testString"
				username := "testString"
				password := "testString"
				instanceName := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigOneOfServiceNowDestinationConfig(clientID, clientSecret, username, password, instanceName)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigOneOfSlackDestinationConfig successfully`, func() {
				url := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigOneOfSlackDestinationConfig(url)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigOneOfWebhookDestinationConfig successfully`, func() {
				url := "testString"
				verb := "get"
				_model, err := eventNotificationsService.NewDestinationConfigOneOfWebhookDestinationConfig(url, verb)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubscriptionCreateAttributesCustomEmailAttributes successfully`, func() {
				invited := []string{"testString"}
				addNotificationPayload := false
				replyToMail := "testString"
				replyToName := "testString"
				fromName := "testString"
				fromEmail := "testString"
				_model, err := eventNotificationsService.NewSubscriptionCreateAttributesCustomEmailAttributes(invited, addNotificationPayload, replyToMail, replyToName, fromName, fromEmail)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubscriptionCreateAttributesEmailAttributes successfully`, func() {
				invited := []string{"testString"}
				addNotificationPayload := false
				replyToMail := "testString"
				replyToName := "testString"
				fromName := "testString"
				_model, err := eventNotificationsService.NewSubscriptionCreateAttributesEmailAttributes(invited, addNotificationPayload, replyToMail, replyToName, fromName)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubscriptionCreateAttributesSmsAttributes successfully`, func() {
				invited := []string{"testString"}
				_model, err := eventNotificationsService.NewSubscriptionCreateAttributesSmsAttributes(invited)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubscriptionCreateAttributesWebhookAttributes successfully`, func() {
				signingEnabled := true
				_model, err := eventNotificationsService.NewSubscriptionCreateAttributesWebhookAttributes(signingEnabled)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubscriptionUpdateAttributesCustomEmailUpdateAttributes successfully`, func() {
				addNotificationPayload := false
				replyToMail := "testString"
				replyToName := "testString"
				fromName := "testString"
				fromEmail := "testString"
				_model, err := eventNotificationsService.NewSubscriptionUpdateAttributesCustomEmailUpdateAttributes(addNotificationPayload, replyToMail, replyToName, fromName, fromEmail)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubscriptionUpdateAttributesEmailUpdateAttributes successfully`, func() {
				addNotificationPayload := false
				replyToMail := "testString"
				replyToName := "testString"
				fromName := "testString"
				_model, err := eventNotificationsService.NewSubscriptionUpdateAttributesEmailUpdateAttributes(addNotificationPayload, replyToMail, replyToName, fromName)
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
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
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
