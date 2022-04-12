/**
 * (C) Copyright IBM Corp. 2022.
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
	Describe(`SendNotifications(sendNotificationsOptions *SendNotificationsOptions) - Operation response error`, func() {
		sendNotificationsPath := "/v1/instances/testString/notifications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(sendNotificationsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Ce-Ibmenseverity"]).ToNot(BeNil())
					Expect(req.Header["Ce-Ibmenseverity"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Ibmendefaultshort"]).ToNot(BeNil())
					Expect(req.Header["Ce-Ibmendefaultshort"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Ibmendefaultlong"]).ToNot(BeNil())
					Expect(req.Header["Ce-Ibmendefaultlong"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Ibmenfcmbody"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenfcmbody"][0]).To(Equal(fmt.Sprintf("%v", notificationFcmBodyModel)))
					Expect(req.Header["Ce-Ibmenapnsbody"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenapnsbody"][0]).To(Equal(fmt.Sprintf("%v", notificationApnsBodyModel)))
					Expect(req.Header["Ce-Ibmenpushto"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenpushto"][0]).To(Equal(fmt.Sprintf("%v", notificationDevicesModel)))
					Expect(req.Header["Ce-Ibmenapnsheaders"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenapnsheaders"][0]).To(Equal(fmt.Sprintf("%v", make(map[string]interface{}))))
					Expect(req.Header["Ce-Ibmenchromebody"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenchromebody"][0]).To(Equal(fmt.Sprintf("%v", notificationChromeBodyModel)))
					Expect(req.Header["Ce-Ibmenfirefoxbody"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenfirefoxbody"][0]).To(Equal(fmt.Sprintf("%v", notificationFirefoxBodyModel)))
					Expect(req.Header["Ce-Ibmenchromeheaders"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenchromeheaders"][0]).To(Equal(fmt.Sprintf("%v", make(map[string]interface{}))))
					Expect(req.Header["Ce-Ibmenfirefoxheaders"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenfirefoxheaders"][0]).To(Equal(fmt.Sprintf("%v", make(map[string]interface{}))))
					Expect(req.Header["Ce-Ibmensourceid"]).ToNot(BeNil())
					Expect(req.Header["Ce-Ibmensourceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Id"]).ToNot(BeNil())
					Expect(req.Header["Ce-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Source"]).ToNot(BeNil())
					Expect(req.Header["Ce-Source"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Type"]).ToNot(BeNil())
					Expect(req.Header["Ce-Type"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Specversion"]).ToNot(BeNil())
					Expect(req.Header["Ce-Specversion"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.Header["Ce-Time"]).ToNot(BeNil())
					Expect(req.Header["Ce-Time"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime("2019-01-01T12:00:00.000Z"))))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
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

				// Construct an instance of the Lights model
				lightsModel := new(eventnotificationsv1.Lights)
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(0))
				lightsModel.LedOffMs = core.StringPtr("testString")

				// Construct an instance of the Style model
				styleModel := new(eventnotificationsv1.Style)
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}
				styleModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationFcmBodyMessageData model
				notificationFcmBodyMessageDataModel := new(eventnotificationsv1.NotificationFcmBodyMessageData)
				notificationFcmBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.CollapseKey = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.InteractiveCategory = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Icon = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.DelayWhileIdle = core.BoolPtr(true)
				notificationFcmBodyMessageDataModel.Sync = core.BoolPtr(true)
				notificationFcmBodyMessageDataModel.Visibility = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Redact = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Payload = make(map[string]interface{})
				notificationFcmBodyMessageDataModel.Priority = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Sound = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationFcmBodyMessageDataModel.Lights = lightsModel
				notificationFcmBodyMessageDataModel.AndroidTitle = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.GroupID = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Style = styleModel
				notificationFcmBodyMessageDataModel.Type = core.StringPtr("DEFAULT")

				// Construct an instance of the NotificationFcmBodyMessageEnData model
				notificationFcmBodyModel := new(eventnotificationsv1.NotificationFcmBodyMessageEnData)
				notificationFcmBodyModel.EnData = notificationFcmBodyMessageDataModel
				notificationFcmBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationApnsBodyMessageData model
				notificationApnsBodyMessageDataModel := new(eventnotificationsv1.NotificationApnsBodyMessageData)
				notificationApnsBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Badge = core.Int64Ptr(int64(38))
				notificationApnsBodyMessageDataModel.InteractiveCategory = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.IosActionKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				notificationApnsBodyMessageDataModel.Sound = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.TitleLocKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.LocKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.LaunchImage = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.TitleLocArgs = []string{"testString"}
				notificationApnsBodyMessageDataModel.LocArgs = []string{"testString"}
				notificationApnsBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Subtitle = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.AttachmentURL = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Type = core.StringPtr("DEFAULT")
				notificationApnsBodyMessageDataModel.ApnsCollapseID = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsThreadID = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))

				// Construct an instance of the NotificationApnsBodyMessageEnData model
				notificationApnsBodyModel := new(eventnotificationsv1.NotificationApnsBodyMessageEnData)
				notificationApnsBodyModel.EnData = notificationApnsBodyMessageDataModel
				notificationApnsBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationDevices model
				notificationDevicesModel := new(eventnotificationsv1.NotificationDevices)
				notificationDevicesModel.FcmDevices = []string{"testString"}
				notificationDevicesModel.ApnsDevices = []string{"testString"}
				notificationDevicesModel.UserIds = []string{"testString"}
				notificationDevicesModel.Tags = []string{"testString"}
				notificationDevicesModel.Platforms = []string{"testString"}

				// Construct an instance of the NotificationChromeBodyMessageData model
				notificationChromeBodyMessageDataModel := new(eventnotificationsv1.NotificationChromeBodyMessageData)
				notificationChromeBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.IconURL = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationChromeBodyMessageDataModel.Payload = make(map[string]interface{})

				// Construct an instance of the NotificationChromeBodyMessageEnData model
				notificationChromeBodyModel := new(eventnotificationsv1.NotificationChromeBodyMessageEnData)
				notificationChromeBodyModel.EnData = notificationChromeBodyMessageDataModel
				notificationChromeBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationFirefoxBodyMessageData model
				notificationFirefoxBodyMessageDataModel := new(eventnotificationsv1.NotificationFirefoxBodyMessageData)
				notificationFirefoxBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.IconURL = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationFirefoxBodyMessageDataModel.Payload = make(map[string]interface{})

				// Construct an instance of the NotificationFirefoxBodyMessageEnData model
				notificationFirefoxBodyModel := new(eventnotificationsv1.NotificationFirefoxBodyMessageEnData)
				notificationFirefoxBodyModel.EnData = notificationFirefoxBodyMessageDataModel
				notificationFirefoxBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsRequestNotificationCreate model
				sendNotificationsRequestModel := new(eventnotificationsv1.SendNotificationsRequestNotificationCreate)
				sendNotificationsRequestModel.Data = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmenseverity = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmenfcmbody = notificationFcmBodyModel
				sendNotificationsRequestModel.Ibmenapnsbody = notificationApnsBodyModel
				sendNotificationsRequestModel.Ibmenpushto = notificationDevicesModel
				sendNotificationsRequestModel.Ibmenapnsheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmendefaultshort = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmendefaultlong = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmenchromebody = notificationChromeBodyModel
				sendNotificationsRequestModel.Ibmenfirefoxbody = notificationFirefoxBodyModel
				sendNotificationsRequestModel.Ibmenchromeheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmenfirefoxheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmensourceid = core.StringPtr("testString")
				sendNotificationsRequestModel.Datacontenttype = core.StringPtr("application/json")
				sendNotificationsRequestModel.Subject = core.StringPtr("testString")
				sendNotificationsRequestModel.ID = core.StringPtr("testString")
				sendNotificationsRequestModel.Source = core.StringPtr("testString")
				sendNotificationsRequestModel.Type = core.StringPtr("testString")
				sendNotificationsRequestModel.Specversion = core.StringPtr("1.0")
				sendNotificationsRequestModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				sendNotificationsRequestModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsOptions model
				sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
				sendNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendNotificationsOptionsModel.Body = sendNotificationsRequestModel
				sendNotificationsOptionsModel.CeIbmenseverity = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmendefaultshort = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmendefaultlong = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmenfcmbody = notificationFcmBodyModel
				sendNotificationsOptionsModel.CeIbmenapnsbody = notificationApnsBodyModel
				sendNotificationsOptionsModel.CeIbmenpushto = notificationDevicesModel
				sendNotificationsOptionsModel.CeIbmenapnsheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmenchromebody = notificationChromeBodyModel
				sendNotificationsOptionsModel.CeIbmenfirefoxbody = notificationFirefoxBodyModel
				sendNotificationsOptionsModel.CeIbmenchromeheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmenfirefoxheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmensourceid = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeID = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeSource = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeType = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeSpecversion = core.StringPtr("1.0")
				sendNotificationsOptionsModel.CeTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
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

					Expect(req.Header["Ce-Ibmenseverity"]).ToNot(BeNil())
					Expect(req.Header["Ce-Ibmenseverity"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Ibmendefaultshort"]).ToNot(BeNil())
					Expect(req.Header["Ce-Ibmendefaultshort"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Ibmendefaultlong"]).ToNot(BeNil())
					Expect(req.Header["Ce-Ibmendefaultlong"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Ibmenfcmbody"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenfcmbody"][0]).To(Equal(fmt.Sprintf("%v", notificationFcmBodyModel)))
					Expect(req.Header["Ce-Ibmenapnsbody"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenapnsbody"][0]).To(Equal(fmt.Sprintf("%v", notificationApnsBodyModel)))
					Expect(req.Header["Ce-Ibmenpushto"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenpushto"][0]).To(Equal(fmt.Sprintf("%v", notificationDevicesModel)))
					Expect(req.Header["Ce-Ibmenapnsheaders"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenapnsheaders"][0]).To(Equal(fmt.Sprintf("%v", make(map[string]interface{}))))
					Expect(req.Header["Ce-Ibmenchromebody"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenchromebody"][0]).To(Equal(fmt.Sprintf("%v", notificationChromeBodyModel)))
					Expect(req.Header["Ce-Ibmenfirefoxbody"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenfirefoxbody"][0]).To(Equal(fmt.Sprintf("%v", notificationFirefoxBodyModel)))
					Expect(req.Header["Ce-Ibmenchromeheaders"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenchromeheaders"][0]).To(Equal(fmt.Sprintf("%v", make(map[string]interface{}))))
					Expect(req.Header["Ce-Ibmenfirefoxheaders"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenfirefoxheaders"][0]).To(Equal(fmt.Sprintf("%v", make(map[string]interface{}))))
					Expect(req.Header["Ce-Ibmensourceid"]).ToNot(BeNil())
					Expect(req.Header["Ce-Ibmensourceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Id"]).ToNot(BeNil())
					Expect(req.Header["Ce-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Source"]).ToNot(BeNil())
					Expect(req.Header["Ce-Source"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Type"]).ToNot(BeNil())
					Expect(req.Header["Ce-Type"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Specversion"]).ToNot(BeNil())
					Expect(req.Header["Ce-Specversion"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.Header["Ce-Time"]).ToNot(BeNil())
					Expect(req.Header["Ce-Time"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime("2019-01-01T12:00:00.000Z"))))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
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

				// Construct an instance of the Lights model
				lightsModel := new(eventnotificationsv1.Lights)
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(0))
				lightsModel.LedOffMs = core.StringPtr("testString")

				// Construct an instance of the Style model
				styleModel := new(eventnotificationsv1.Style)
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}
				styleModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationFcmBodyMessageData model
				notificationFcmBodyMessageDataModel := new(eventnotificationsv1.NotificationFcmBodyMessageData)
				notificationFcmBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.CollapseKey = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.InteractiveCategory = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Icon = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.DelayWhileIdle = core.BoolPtr(true)
				notificationFcmBodyMessageDataModel.Sync = core.BoolPtr(true)
				notificationFcmBodyMessageDataModel.Visibility = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Redact = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Payload = make(map[string]interface{})
				notificationFcmBodyMessageDataModel.Priority = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Sound = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationFcmBodyMessageDataModel.Lights = lightsModel
				notificationFcmBodyMessageDataModel.AndroidTitle = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.GroupID = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Style = styleModel
				notificationFcmBodyMessageDataModel.Type = core.StringPtr("DEFAULT")

				// Construct an instance of the NotificationFcmBodyMessageEnData model
				notificationFcmBodyModel := new(eventnotificationsv1.NotificationFcmBodyMessageEnData)
				notificationFcmBodyModel.EnData = notificationFcmBodyMessageDataModel
				notificationFcmBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationApnsBodyMessageData model
				notificationApnsBodyMessageDataModel := new(eventnotificationsv1.NotificationApnsBodyMessageData)
				notificationApnsBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Badge = core.Int64Ptr(int64(38))
				notificationApnsBodyMessageDataModel.InteractiveCategory = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.IosActionKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				notificationApnsBodyMessageDataModel.Sound = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.TitleLocKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.LocKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.LaunchImage = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.TitleLocArgs = []string{"testString"}
				notificationApnsBodyMessageDataModel.LocArgs = []string{"testString"}
				notificationApnsBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Subtitle = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.AttachmentURL = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Type = core.StringPtr("DEFAULT")
				notificationApnsBodyMessageDataModel.ApnsCollapseID = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsThreadID = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))

				// Construct an instance of the NotificationApnsBodyMessageEnData model
				notificationApnsBodyModel := new(eventnotificationsv1.NotificationApnsBodyMessageEnData)
				notificationApnsBodyModel.EnData = notificationApnsBodyMessageDataModel
				notificationApnsBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationDevices model
				notificationDevicesModel := new(eventnotificationsv1.NotificationDevices)
				notificationDevicesModel.FcmDevices = []string{"testString"}
				notificationDevicesModel.ApnsDevices = []string{"testString"}
				notificationDevicesModel.UserIds = []string{"testString"}
				notificationDevicesModel.Tags = []string{"testString"}
				notificationDevicesModel.Platforms = []string{"testString"}

				// Construct an instance of the NotificationChromeBodyMessageData model
				notificationChromeBodyMessageDataModel := new(eventnotificationsv1.NotificationChromeBodyMessageData)
				notificationChromeBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.IconURL = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationChromeBodyMessageDataModel.Payload = make(map[string]interface{})

				// Construct an instance of the NotificationChromeBodyMessageEnData model
				notificationChromeBodyModel := new(eventnotificationsv1.NotificationChromeBodyMessageEnData)
				notificationChromeBodyModel.EnData = notificationChromeBodyMessageDataModel
				notificationChromeBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationFirefoxBodyMessageData model
				notificationFirefoxBodyMessageDataModel := new(eventnotificationsv1.NotificationFirefoxBodyMessageData)
				notificationFirefoxBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.IconURL = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationFirefoxBodyMessageDataModel.Payload = make(map[string]interface{})

				// Construct an instance of the NotificationFirefoxBodyMessageEnData model
				notificationFirefoxBodyModel := new(eventnotificationsv1.NotificationFirefoxBodyMessageEnData)
				notificationFirefoxBodyModel.EnData = notificationFirefoxBodyMessageDataModel
				notificationFirefoxBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsRequestNotificationCreate model
				sendNotificationsRequestModel := new(eventnotificationsv1.SendNotificationsRequestNotificationCreate)
				sendNotificationsRequestModel.Data = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmenseverity = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmenfcmbody = notificationFcmBodyModel
				sendNotificationsRequestModel.Ibmenapnsbody = notificationApnsBodyModel
				sendNotificationsRequestModel.Ibmenpushto = notificationDevicesModel
				sendNotificationsRequestModel.Ibmenapnsheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmendefaultshort = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmendefaultlong = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmenchromebody = notificationChromeBodyModel
				sendNotificationsRequestModel.Ibmenfirefoxbody = notificationFirefoxBodyModel
				sendNotificationsRequestModel.Ibmenchromeheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmenfirefoxheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmensourceid = core.StringPtr("testString")
				sendNotificationsRequestModel.Datacontenttype = core.StringPtr("application/json")
				sendNotificationsRequestModel.Subject = core.StringPtr("testString")
				sendNotificationsRequestModel.ID = core.StringPtr("testString")
				sendNotificationsRequestModel.Source = core.StringPtr("testString")
				sendNotificationsRequestModel.Type = core.StringPtr("testString")
				sendNotificationsRequestModel.Specversion = core.StringPtr("1.0")
				sendNotificationsRequestModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				sendNotificationsRequestModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsOptions model
				sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
				sendNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendNotificationsOptionsModel.Body = sendNotificationsRequestModel
				sendNotificationsOptionsModel.CeIbmenseverity = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmendefaultshort = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmendefaultlong = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmenfcmbody = notificationFcmBodyModel
				sendNotificationsOptionsModel.CeIbmenapnsbody = notificationApnsBodyModel
				sendNotificationsOptionsModel.CeIbmenpushto = notificationDevicesModel
				sendNotificationsOptionsModel.CeIbmenapnsheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmenchromebody = notificationChromeBodyModel
				sendNotificationsOptionsModel.CeIbmenfirefoxbody = notificationFirefoxBodyModel
				sendNotificationsOptionsModel.CeIbmenchromeheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmenfirefoxheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmensourceid = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeID = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeSource = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeType = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeSpecversion = core.StringPtr("1.0")
				sendNotificationsOptionsModel.CeTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
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

					Expect(req.Header["Ce-Ibmenseverity"]).ToNot(BeNil())
					Expect(req.Header["Ce-Ibmenseverity"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Ibmendefaultshort"]).ToNot(BeNil())
					Expect(req.Header["Ce-Ibmendefaultshort"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Ibmendefaultlong"]).ToNot(BeNil())
					Expect(req.Header["Ce-Ibmendefaultlong"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Ibmenfcmbody"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenfcmbody"][0]).To(Equal(fmt.Sprintf("%v", notificationFcmBodyModel)))
					Expect(req.Header["Ce-Ibmenapnsbody"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenapnsbody"][0]).To(Equal(fmt.Sprintf("%v", notificationApnsBodyModel)))
					Expect(req.Header["Ce-Ibmenpushto"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenpushto"][0]).To(Equal(fmt.Sprintf("%v", notificationDevicesModel)))
					Expect(req.Header["Ce-Ibmenapnsheaders"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenapnsheaders"][0]).To(Equal(fmt.Sprintf("%v", make(map[string]interface{}))))
					Expect(req.Header["Ce-Ibmenchromebody"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenchromebody"][0]).To(Equal(fmt.Sprintf("%v", notificationChromeBodyModel)))
					Expect(req.Header["Ce-Ibmenfirefoxbody"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenfirefoxbody"][0]).To(Equal(fmt.Sprintf("%v", notificationFirefoxBodyModel)))
					Expect(req.Header["Ce-Ibmenchromeheaders"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenchromeheaders"][0]).To(Equal(fmt.Sprintf("%v", make(map[string]interface{}))))
					Expect(req.Header["Ce-Ibmenfirefoxheaders"]).ToNot(BeNil())
					//Expect(req.Header["Ce-Ibmenfirefoxheaders"][0]).To(Equal(fmt.Sprintf("%v", make(map[string]interface{}))))
					Expect(req.Header["Ce-Ibmensourceid"]).ToNot(BeNil())
					Expect(req.Header["Ce-Ibmensourceid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Id"]).ToNot(BeNil())
					Expect(req.Header["Ce-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Source"]).ToNot(BeNil())
					Expect(req.Header["Ce-Source"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Type"]).ToNot(BeNil())
					Expect(req.Header["Ce-Type"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Ce-Specversion"]).ToNot(BeNil())
					Expect(req.Header["Ce-Specversion"][0]).To(Equal(fmt.Sprintf("%v", "1.0")))
					Expect(req.Header["Ce-Time"]).ToNot(BeNil())
					Expect(req.Header["Ce-Time"][0]).To(Equal(fmt.Sprintf("%v", CreateMockDateTime("2019-01-01T12:00:00.000Z"))))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
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

				// Construct an instance of the Lights model
				lightsModel := new(eventnotificationsv1.Lights)
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(0))
				lightsModel.LedOffMs = core.StringPtr("testString")

				// Construct an instance of the Style model
				styleModel := new(eventnotificationsv1.Style)
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}
				styleModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationFcmBodyMessageData model
				notificationFcmBodyMessageDataModel := new(eventnotificationsv1.NotificationFcmBodyMessageData)
				notificationFcmBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.CollapseKey = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.InteractiveCategory = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Icon = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.DelayWhileIdle = core.BoolPtr(true)
				notificationFcmBodyMessageDataModel.Sync = core.BoolPtr(true)
				notificationFcmBodyMessageDataModel.Visibility = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Redact = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Payload = make(map[string]interface{})
				notificationFcmBodyMessageDataModel.Priority = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Sound = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationFcmBodyMessageDataModel.Lights = lightsModel
				notificationFcmBodyMessageDataModel.AndroidTitle = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.GroupID = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Style = styleModel
				notificationFcmBodyMessageDataModel.Type = core.StringPtr("DEFAULT")

				// Construct an instance of the NotificationFcmBodyMessageEnData model
				notificationFcmBodyModel := new(eventnotificationsv1.NotificationFcmBodyMessageEnData)
				notificationFcmBodyModel.EnData = notificationFcmBodyMessageDataModel
				notificationFcmBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationApnsBodyMessageData model
				notificationApnsBodyMessageDataModel := new(eventnotificationsv1.NotificationApnsBodyMessageData)
				notificationApnsBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Badge = core.Int64Ptr(int64(38))
				notificationApnsBodyMessageDataModel.InteractiveCategory = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.IosActionKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				notificationApnsBodyMessageDataModel.Sound = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.TitleLocKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.LocKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.LaunchImage = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.TitleLocArgs = []string{"testString"}
				notificationApnsBodyMessageDataModel.LocArgs = []string{"testString"}
				notificationApnsBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Subtitle = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.AttachmentURL = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Type = core.StringPtr("DEFAULT")
				notificationApnsBodyMessageDataModel.ApnsCollapseID = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsThreadID = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))

				// Construct an instance of the NotificationApnsBodyMessageEnData model
				notificationApnsBodyModel := new(eventnotificationsv1.NotificationApnsBodyMessageEnData)
				notificationApnsBodyModel.EnData = notificationApnsBodyMessageDataModel
				notificationApnsBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationDevices model
				notificationDevicesModel := new(eventnotificationsv1.NotificationDevices)
				notificationDevicesModel.FcmDevices = []string{"testString"}
				notificationDevicesModel.ApnsDevices = []string{"testString"}
				notificationDevicesModel.UserIds = []string{"testString"}
				notificationDevicesModel.Tags = []string{"testString"}
				notificationDevicesModel.Platforms = []string{"testString"}

				// Construct an instance of the NotificationChromeBodyMessageData model
				notificationChromeBodyMessageDataModel := new(eventnotificationsv1.NotificationChromeBodyMessageData)
				notificationChromeBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.IconURL = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationChromeBodyMessageDataModel.Payload = make(map[string]interface{})

				// Construct an instance of the NotificationChromeBodyMessageEnData model
				notificationChromeBodyModel := new(eventnotificationsv1.NotificationChromeBodyMessageEnData)
				notificationChromeBodyModel.EnData = notificationChromeBodyMessageDataModel
				notificationChromeBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationFirefoxBodyMessageData model
				notificationFirefoxBodyMessageDataModel := new(eventnotificationsv1.NotificationFirefoxBodyMessageData)
				notificationFirefoxBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.IconURL = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationFirefoxBodyMessageDataModel.Payload = make(map[string]interface{})

				// Construct an instance of the NotificationFirefoxBodyMessageEnData model
				notificationFirefoxBodyModel := new(eventnotificationsv1.NotificationFirefoxBodyMessageEnData)
				notificationFirefoxBodyModel.EnData = notificationFirefoxBodyMessageDataModel
				notificationFirefoxBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsRequestNotificationCreate model
				sendNotificationsRequestModel := new(eventnotificationsv1.SendNotificationsRequestNotificationCreate)
				sendNotificationsRequestModel.Data = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmenseverity = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmenfcmbody = notificationFcmBodyModel
				sendNotificationsRequestModel.Ibmenapnsbody = notificationApnsBodyModel
				sendNotificationsRequestModel.Ibmenpushto = notificationDevicesModel
				sendNotificationsRequestModel.Ibmenapnsheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmendefaultshort = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmendefaultlong = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmenchromebody = notificationChromeBodyModel
				sendNotificationsRequestModel.Ibmenfirefoxbody = notificationFirefoxBodyModel
				sendNotificationsRequestModel.Ibmenchromeheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmenfirefoxheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmensourceid = core.StringPtr("testString")
				sendNotificationsRequestModel.Datacontenttype = core.StringPtr("application/json")
				sendNotificationsRequestModel.Subject = core.StringPtr("testString")
				sendNotificationsRequestModel.ID = core.StringPtr("testString")
				sendNotificationsRequestModel.Source = core.StringPtr("testString")
				sendNotificationsRequestModel.Type = core.StringPtr("testString")
				sendNotificationsRequestModel.Specversion = core.StringPtr("1.0")
				sendNotificationsRequestModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				sendNotificationsRequestModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsOptions model
				sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
				sendNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendNotificationsOptionsModel.Body = sendNotificationsRequestModel
				sendNotificationsOptionsModel.CeIbmenseverity = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmendefaultshort = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmendefaultlong = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmenfcmbody = notificationFcmBodyModel
				sendNotificationsOptionsModel.CeIbmenapnsbody = notificationApnsBodyModel
				sendNotificationsOptionsModel.CeIbmenpushto = notificationDevicesModel
				sendNotificationsOptionsModel.CeIbmenapnsheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmenchromebody = notificationChromeBodyModel
				sendNotificationsOptionsModel.CeIbmenfirefoxbody = notificationFirefoxBodyModel
				sendNotificationsOptionsModel.CeIbmenchromeheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmenfirefoxheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmensourceid = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeID = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeSource = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeType = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeSpecversion = core.StringPtr("1.0")
				sendNotificationsOptionsModel.CeTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
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

				// Construct an instance of the Lights model
				lightsModel := new(eventnotificationsv1.Lights)
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(0))
				lightsModel.LedOffMs = core.StringPtr("testString")

				// Construct an instance of the Style model
				styleModel := new(eventnotificationsv1.Style)
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}
				styleModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationFcmBodyMessageData model
				notificationFcmBodyMessageDataModel := new(eventnotificationsv1.NotificationFcmBodyMessageData)
				notificationFcmBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.CollapseKey = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.InteractiveCategory = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Icon = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.DelayWhileIdle = core.BoolPtr(true)
				notificationFcmBodyMessageDataModel.Sync = core.BoolPtr(true)
				notificationFcmBodyMessageDataModel.Visibility = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Redact = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Payload = make(map[string]interface{})
				notificationFcmBodyMessageDataModel.Priority = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Sound = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationFcmBodyMessageDataModel.Lights = lightsModel
				notificationFcmBodyMessageDataModel.AndroidTitle = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.GroupID = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Style = styleModel
				notificationFcmBodyMessageDataModel.Type = core.StringPtr("DEFAULT")

				// Construct an instance of the NotificationFcmBodyMessageEnData model
				notificationFcmBodyModel := new(eventnotificationsv1.NotificationFcmBodyMessageEnData)
				notificationFcmBodyModel.EnData = notificationFcmBodyMessageDataModel
				notificationFcmBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationApnsBodyMessageData model
				notificationApnsBodyMessageDataModel := new(eventnotificationsv1.NotificationApnsBodyMessageData)
				notificationApnsBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Badge = core.Int64Ptr(int64(38))
				notificationApnsBodyMessageDataModel.InteractiveCategory = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.IosActionKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				notificationApnsBodyMessageDataModel.Sound = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.TitleLocKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.LocKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.LaunchImage = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.TitleLocArgs = []string{"testString"}
				notificationApnsBodyMessageDataModel.LocArgs = []string{"testString"}
				notificationApnsBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Subtitle = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.AttachmentURL = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Type = core.StringPtr("DEFAULT")
				notificationApnsBodyMessageDataModel.ApnsCollapseID = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsThreadID = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))

				// Construct an instance of the NotificationApnsBodyMessageEnData model
				notificationApnsBodyModel := new(eventnotificationsv1.NotificationApnsBodyMessageEnData)
				notificationApnsBodyModel.EnData = notificationApnsBodyMessageDataModel
				notificationApnsBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationDevices model
				notificationDevicesModel := new(eventnotificationsv1.NotificationDevices)
				notificationDevicesModel.FcmDevices = []string{"testString"}
				notificationDevicesModel.ApnsDevices = []string{"testString"}
				notificationDevicesModel.UserIds = []string{"testString"}
				notificationDevicesModel.Tags = []string{"testString"}
				notificationDevicesModel.Platforms = []string{"testString"}

				// Construct an instance of the NotificationChromeBodyMessageData model
				notificationChromeBodyMessageDataModel := new(eventnotificationsv1.NotificationChromeBodyMessageData)
				notificationChromeBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.IconURL = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationChromeBodyMessageDataModel.Payload = make(map[string]interface{})

				// Construct an instance of the NotificationChromeBodyMessageEnData model
				notificationChromeBodyModel := new(eventnotificationsv1.NotificationChromeBodyMessageEnData)
				notificationChromeBodyModel.EnData = notificationChromeBodyMessageDataModel
				notificationChromeBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationFirefoxBodyMessageData model
				notificationFirefoxBodyMessageDataModel := new(eventnotificationsv1.NotificationFirefoxBodyMessageData)
				notificationFirefoxBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.IconURL = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationFirefoxBodyMessageDataModel.Payload = make(map[string]interface{})

				// Construct an instance of the NotificationFirefoxBodyMessageEnData model
				notificationFirefoxBodyModel := new(eventnotificationsv1.NotificationFirefoxBodyMessageEnData)
				notificationFirefoxBodyModel.EnData = notificationFirefoxBodyMessageDataModel
				notificationFirefoxBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsRequestNotificationCreate model
				sendNotificationsRequestModel := new(eventnotificationsv1.SendNotificationsRequestNotificationCreate)
				sendNotificationsRequestModel.Data = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmenseverity = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmenfcmbody = notificationFcmBodyModel
				sendNotificationsRequestModel.Ibmenapnsbody = notificationApnsBodyModel
				sendNotificationsRequestModel.Ibmenpushto = notificationDevicesModel
				sendNotificationsRequestModel.Ibmenapnsheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmendefaultshort = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmendefaultlong = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmenchromebody = notificationChromeBodyModel
				sendNotificationsRequestModel.Ibmenfirefoxbody = notificationFirefoxBodyModel
				sendNotificationsRequestModel.Ibmenchromeheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmenfirefoxheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmensourceid = core.StringPtr("testString")
				sendNotificationsRequestModel.Datacontenttype = core.StringPtr("application/json")
				sendNotificationsRequestModel.Subject = core.StringPtr("testString")
				sendNotificationsRequestModel.ID = core.StringPtr("testString")
				sendNotificationsRequestModel.Source = core.StringPtr("testString")
				sendNotificationsRequestModel.Type = core.StringPtr("testString")
				sendNotificationsRequestModel.Specversion = core.StringPtr("1.0")
				sendNotificationsRequestModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				sendNotificationsRequestModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsOptions model
				sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
				sendNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendNotificationsOptionsModel.Body = sendNotificationsRequestModel
				sendNotificationsOptionsModel.CeIbmenseverity = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmendefaultshort = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmendefaultlong = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmenfcmbody = notificationFcmBodyModel
				sendNotificationsOptionsModel.CeIbmenapnsbody = notificationApnsBodyModel
				sendNotificationsOptionsModel.CeIbmenpushto = notificationDevicesModel
				sendNotificationsOptionsModel.CeIbmenapnsheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmenchromebody = notificationChromeBodyModel
				sendNotificationsOptionsModel.CeIbmenfirefoxbody = notificationFirefoxBodyModel
				sendNotificationsOptionsModel.CeIbmenchromeheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmenfirefoxheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmensourceid = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeID = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeSource = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeType = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeSpecversion = core.StringPtr("1.0")
				sendNotificationsOptionsModel.CeTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke SendNotifications successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the Lights model
				lightsModel := new(eventnotificationsv1.Lights)
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(0))
				lightsModel.LedOffMs = core.StringPtr("testString")

				// Construct an instance of the Style model
				styleModel := new(eventnotificationsv1.Style)
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}
				styleModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationFcmBodyMessageData model
				notificationFcmBodyMessageDataModel := new(eventnotificationsv1.NotificationFcmBodyMessageData)
				notificationFcmBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.CollapseKey = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.InteractiveCategory = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Icon = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.DelayWhileIdle = core.BoolPtr(true)
				notificationFcmBodyMessageDataModel.Sync = core.BoolPtr(true)
				notificationFcmBodyMessageDataModel.Visibility = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Redact = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Payload = make(map[string]interface{})
				notificationFcmBodyMessageDataModel.Priority = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Sound = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationFcmBodyMessageDataModel.Lights = lightsModel
				notificationFcmBodyMessageDataModel.AndroidTitle = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.GroupID = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Style = styleModel
				notificationFcmBodyMessageDataModel.Type = core.StringPtr("DEFAULT")

				// Construct an instance of the NotificationFcmBodyMessageEnData model
				notificationFcmBodyModel := new(eventnotificationsv1.NotificationFcmBodyMessageEnData)
				notificationFcmBodyModel.EnData = notificationFcmBodyMessageDataModel
				notificationFcmBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationApnsBodyMessageData model
				notificationApnsBodyMessageDataModel := new(eventnotificationsv1.NotificationApnsBodyMessageData)
				notificationApnsBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Badge = core.Int64Ptr(int64(38))
				notificationApnsBodyMessageDataModel.InteractiveCategory = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.IosActionKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				notificationApnsBodyMessageDataModel.Sound = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.TitleLocKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.LocKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.LaunchImage = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.TitleLocArgs = []string{"testString"}
				notificationApnsBodyMessageDataModel.LocArgs = []string{"testString"}
				notificationApnsBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Subtitle = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.AttachmentURL = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Type = core.StringPtr("DEFAULT")
				notificationApnsBodyMessageDataModel.ApnsCollapseID = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsThreadID = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))

				// Construct an instance of the NotificationApnsBodyMessageEnData model
				notificationApnsBodyModel := new(eventnotificationsv1.NotificationApnsBodyMessageEnData)
				notificationApnsBodyModel.EnData = notificationApnsBodyMessageDataModel
				notificationApnsBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationDevices model
				notificationDevicesModel := new(eventnotificationsv1.NotificationDevices)
				notificationDevicesModel.FcmDevices = []string{"testString"}
				notificationDevicesModel.ApnsDevices = []string{"testString"}
				notificationDevicesModel.UserIds = []string{"testString"}
				notificationDevicesModel.Tags = []string{"testString"}
				notificationDevicesModel.Platforms = []string{"testString"}

				// Construct an instance of the NotificationChromeBodyMessageData model
				notificationChromeBodyMessageDataModel := new(eventnotificationsv1.NotificationChromeBodyMessageData)
				notificationChromeBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.IconURL = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationChromeBodyMessageDataModel.Payload = make(map[string]interface{})

				// Construct an instance of the NotificationChromeBodyMessageEnData model
				notificationChromeBodyModel := new(eventnotificationsv1.NotificationChromeBodyMessageEnData)
				notificationChromeBodyModel.EnData = notificationChromeBodyMessageDataModel
				notificationChromeBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the NotificationFirefoxBodyMessageData model
				notificationFirefoxBodyMessageDataModel := new(eventnotificationsv1.NotificationFirefoxBodyMessageData)
				notificationFirefoxBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.IconURL = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationFirefoxBodyMessageDataModel.Payload = make(map[string]interface{})

				// Construct an instance of the NotificationFirefoxBodyMessageEnData model
				notificationFirefoxBodyModel := new(eventnotificationsv1.NotificationFirefoxBodyMessageEnData)
				notificationFirefoxBodyModel.EnData = notificationFirefoxBodyMessageDataModel
				notificationFirefoxBodyModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsRequestNotificationCreate model
				sendNotificationsRequestModel := new(eventnotificationsv1.SendNotificationsRequestNotificationCreate)
				sendNotificationsRequestModel.Data = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmenseverity = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmenfcmbody = notificationFcmBodyModel
				sendNotificationsRequestModel.Ibmenapnsbody = notificationApnsBodyModel
				sendNotificationsRequestModel.Ibmenpushto = notificationDevicesModel
				sendNotificationsRequestModel.Ibmenapnsheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmendefaultshort = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmendefaultlong = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmenchromebody = notificationChromeBodyModel
				sendNotificationsRequestModel.Ibmenfirefoxbody = notificationFirefoxBodyModel
				sendNotificationsRequestModel.Ibmenchromeheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmenfirefoxheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmensourceid = core.StringPtr("testString")
				sendNotificationsRequestModel.Datacontenttype = core.StringPtr("application/json")
				sendNotificationsRequestModel.Subject = core.StringPtr("testString")
				sendNotificationsRequestModel.ID = core.StringPtr("testString")
				sendNotificationsRequestModel.Source = core.StringPtr("testString")
				sendNotificationsRequestModel.Type = core.StringPtr("testString")
				sendNotificationsRequestModel.Specversion = core.StringPtr("1.0")
				sendNotificationsRequestModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				sendNotificationsRequestModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the SendNotificationsOptions model
				sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
				sendNotificationsOptionsModel.InstanceID = core.StringPtr("testString")
				sendNotificationsOptionsModel.Body = sendNotificationsRequestModel
				sendNotificationsOptionsModel.CeIbmenseverity = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmendefaultshort = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmendefaultlong = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeIbmenfcmbody = notificationFcmBodyModel
				sendNotificationsOptionsModel.CeIbmenapnsbody = notificationApnsBodyModel
				sendNotificationsOptionsModel.CeIbmenpushto = notificationDevicesModel
				sendNotificationsOptionsModel.CeIbmenapnsheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmenchromebody = notificationChromeBodyModel
				sendNotificationsOptionsModel.CeIbmenfirefoxbody = notificationFirefoxBodyModel
				sendNotificationsOptionsModel.CeIbmenchromeheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmenfirefoxheaders = make(map[string]interface{})
				sendNotificationsOptionsModel.CeIbmensourceid = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeID = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeSource = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeType = core.StringPtr("testString")
				sendNotificationsOptionsModel.CeSpecversion = core.StringPtr("1.0")
				sendNotificationsOptionsModel.CeTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
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
				createDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
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
				createDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
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
				createDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
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
				createDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
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
				createDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				createDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
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
				updateDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
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
				updateDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
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
				updateDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
				updateDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.UpdateDestination(updateDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDestination with error: Param validation error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
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
				updateDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
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
				updateDestinationOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				updateDestinationOptionsModel.CertificateContentType = core.StringPtr("testString")
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
	Describe(`ListDestinationDevices(listDestinationDevicesOptions *ListDestinationDevicesOptions) - Operation response error`, func() {
		listDestinationDevicesPath := "/v1/instances/testString/destinations/testString/devices"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDestinationDevicesPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDestinationDevices with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListDestinationDevicesOptions model
				listDestinationDevicesOptionsModel := new(eventnotificationsv1.ListDestinationDevicesOptions)
				listDestinationDevicesOptionsModel.InstanceID = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.ID = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listDestinationDevicesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listDestinationDevicesOptionsModel.Search = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.ListDestinationDevices(listDestinationDevicesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.ListDestinationDevices(listDestinationDevicesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDestinationDevices(listDestinationDevicesOptions *ListDestinationDevicesOptions)`, func() {
		listDestinationDevicesPath := "/v1/instances/testString/destinations/testString/devices"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDestinationDevicesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "devices": [{"id": "ID", "user_id": "UserID", "platform": "Platform", "token": "Token", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListDestinationDevices successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the ListDestinationDevicesOptions model
				listDestinationDevicesOptionsModel := new(eventnotificationsv1.ListDestinationDevicesOptions)
				listDestinationDevicesOptionsModel.InstanceID = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.ID = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listDestinationDevicesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listDestinationDevicesOptionsModel.Search = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.ListDestinationDevicesWithContext(ctx, listDestinationDevicesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.ListDestinationDevices(listDestinationDevicesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.ListDestinationDevicesWithContext(ctx, listDestinationDevicesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listDestinationDevicesPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					Expect(req.URL.Query()["search"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "devices": [{"id": "ID", "user_id": "UserID", "platform": "Platform", "token": "Token", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListDestinationDevices successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.ListDestinationDevices(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDestinationDevicesOptions model
				listDestinationDevicesOptionsModel := new(eventnotificationsv1.ListDestinationDevicesOptions)
				listDestinationDevicesOptionsModel.InstanceID = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.ID = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listDestinationDevicesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listDestinationDevicesOptionsModel.Search = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.ListDestinationDevices(listDestinationDevicesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDestinationDevices with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListDestinationDevicesOptions model
				listDestinationDevicesOptionsModel := new(eventnotificationsv1.ListDestinationDevicesOptions)
				listDestinationDevicesOptionsModel.InstanceID = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.ID = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listDestinationDevicesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listDestinationDevicesOptionsModel.Search = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.ListDestinationDevices(listDestinationDevicesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDestinationDevicesOptions model with no property values
				listDestinationDevicesOptionsModelNew := new(eventnotificationsv1.ListDestinationDevicesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.ListDestinationDevices(listDestinationDevicesOptionsModelNew)
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
			It(`Invoke ListDestinationDevices successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListDestinationDevicesOptions model
				listDestinationDevicesOptionsModel := new(eventnotificationsv1.ListDestinationDevicesOptions)
				listDestinationDevicesOptionsModel.InstanceID = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.ID = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.Limit = core.Int64Ptr(int64(1))
				listDestinationDevicesOptionsModel.Offset = core.Int64Ptr(int64(0))
				listDestinationDevicesOptionsModel.Search = core.StringPtr("testString")
				listDestinationDevicesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.ListDestinationDevices(listDestinationDevicesOptionsModel)
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
	Describe(`GetDestinationDevicesReport(getDestinationDevicesReportOptions *GetDestinationDevicesReportOptions) - Operation response error`, func() {
		getDestinationDevicesReportPath := "/v1/instances/testString/destinations/testString/devices/report"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDestinationDevicesReportPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for days query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDestinationDevicesReport with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetDestinationDevicesReportOptions model
				getDestinationDevicesReportOptionsModel := new(eventnotificationsv1.GetDestinationDevicesReportOptions)
				getDestinationDevicesReportOptionsModel.InstanceID = core.StringPtr("testString")
				getDestinationDevicesReportOptionsModel.ID = core.StringPtr("testString")
				getDestinationDevicesReportOptionsModel.Days = core.Int64Ptr(int64(1))
				getDestinationDevicesReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.GetDestinationDevicesReport(getDestinationDevicesReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.GetDestinationDevicesReport(getDestinationDevicesReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDestinationDevicesReport(getDestinationDevicesReportOptions *GetDestinationDevicesReportOptions)`, func() {
		getDestinationDevicesReportPath := "/v1/instances/testString/destinations/testString/devices/report"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDestinationDevicesReportPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for days query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"android": 7, "ios": 3, "chrome": 6, "firefox": 7, "safari": 6, "chromeAppExt": 12, "all": 3}`)
				}))
			})
			It(`Invoke GetDestinationDevicesReport successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the GetDestinationDevicesReportOptions model
				getDestinationDevicesReportOptionsModel := new(eventnotificationsv1.GetDestinationDevicesReportOptions)
				getDestinationDevicesReportOptionsModel.InstanceID = core.StringPtr("testString")
				getDestinationDevicesReportOptionsModel.ID = core.StringPtr("testString")
				getDestinationDevicesReportOptionsModel.Days = core.Int64Ptr(int64(1))
				getDestinationDevicesReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.GetDestinationDevicesReportWithContext(ctx, getDestinationDevicesReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.GetDestinationDevicesReport(getDestinationDevicesReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.GetDestinationDevicesReportWithContext(ctx, getDestinationDevicesReportOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDestinationDevicesReportPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for days query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"android": 7, "ios": 3, "chrome": 6, "firefox": 7, "safari": 6, "chromeAppExt": 12, "all": 3}`)
				}))
			})
			It(`Invoke GetDestinationDevicesReport successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.GetDestinationDevicesReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDestinationDevicesReportOptions model
				getDestinationDevicesReportOptionsModel := new(eventnotificationsv1.GetDestinationDevicesReportOptions)
				getDestinationDevicesReportOptionsModel.InstanceID = core.StringPtr("testString")
				getDestinationDevicesReportOptionsModel.ID = core.StringPtr("testString")
				getDestinationDevicesReportOptionsModel.Days = core.Int64Ptr(int64(1))
				getDestinationDevicesReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.GetDestinationDevicesReport(getDestinationDevicesReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDestinationDevicesReport with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetDestinationDevicesReportOptions model
				getDestinationDevicesReportOptionsModel := new(eventnotificationsv1.GetDestinationDevicesReportOptions)
				getDestinationDevicesReportOptionsModel.InstanceID = core.StringPtr("testString")
				getDestinationDevicesReportOptionsModel.ID = core.StringPtr("testString")
				getDestinationDevicesReportOptionsModel.Days = core.Int64Ptr(int64(1))
				getDestinationDevicesReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.GetDestinationDevicesReport(getDestinationDevicesReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDestinationDevicesReportOptions model with no property values
				getDestinationDevicesReportOptionsModelNew := new(eventnotificationsv1.GetDestinationDevicesReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.GetDestinationDevicesReport(getDestinationDevicesReportOptionsModelNew)
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
			It(`Invoke GetDestinationDevicesReport successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetDestinationDevicesReportOptions model
				getDestinationDevicesReportOptionsModel := new(eventnotificationsv1.GetDestinationDevicesReportOptions)
				getDestinationDevicesReportOptionsModel.InstanceID = core.StringPtr("testString")
				getDestinationDevicesReportOptionsModel.ID = core.StringPtr("testString")
				getDestinationDevicesReportOptionsModel.Days = core.Int64Ptr(int64(1))
				getDestinationDevicesReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.GetDestinationDevicesReport(getDestinationDevicesReportOptionsModel)
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
	Describe(`ListTagsSubscriptionsDevice(listTagsSubscriptionsDeviceOptions *ListTagsSubscriptionsDeviceOptions) - Operation response error`, func() {
		listTagsSubscriptionsDevicePath := "/v1/instances/testString/destinations/testString/tag_subscriptions/devices/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTagsSubscriptionsDevicePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["tag_name"]).To(Equal([]string{"testString"}))
					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTagsSubscriptionsDevice with error: Operation response processing error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListTagsSubscriptionsDeviceOptions model
				listTagsSubscriptionsDeviceOptionsModel := new(eventnotificationsv1.ListTagsSubscriptionsDeviceOptions)
				listTagsSubscriptionsDeviceOptionsModel.InstanceID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.ID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.DeviceID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.TagName = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTagsSubscriptionsDeviceOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsSubscriptionsDeviceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := eventNotificationsService.ListTagsSubscriptionsDevice(listTagsSubscriptionsDeviceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				eventNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = eventNotificationsService.ListTagsSubscriptionsDevice(listTagsSubscriptionsDeviceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTagsSubscriptionsDevice(listTagsSubscriptionsDeviceOptions *ListTagsSubscriptionsDeviceOptions)`, func() {
		listTagsSubscriptionsDevicePath := "/v1/instances/testString/destinations/testString/tag_subscriptions/devices/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTagsSubscriptionsDevicePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["tag_name"]).To(Equal([]string{"testString"}))
					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "tag_subscriptions": [{"id": "ID", "device_id": "DeviceID", "tag_name": "TagName", "user_id": "UserID", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListTagsSubscriptionsDevice successfully with retries`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())
				eventNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the ListTagsSubscriptionsDeviceOptions model
				listTagsSubscriptionsDeviceOptionsModel := new(eventnotificationsv1.ListTagsSubscriptionsDeviceOptions)
				listTagsSubscriptionsDeviceOptionsModel.InstanceID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.ID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.DeviceID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.TagName = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTagsSubscriptionsDeviceOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsSubscriptionsDeviceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := eventNotificationsService.ListTagsSubscriptionsDeviceWithContext(ctx, listTagsSubscriptionsDeviceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				eventNotificationsService.DisableRetries()
				result, response, operationErr := eventNotificationsService.ListTagsSubscriptionsDevice(listTagsSubscriptionsDeviceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = eventNotificationsService.ListTagsSubscriptionsDeviceWithContext(ctx, listTagsSubscriptionsDeviceOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTagsSubscriptionsDevicePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["tag_name"]).To(Equal([]string{"testString"}))
					// TODO: Add check for limit query parameter
					// TODO: Add check for offset query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "tag_subscriptions": [{"id": "ID", "device_id": "DeviceID", "tag_name": "TagName", "user_id": "UserID", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListTagsSubscriptionsDevice successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := eventNotificationsService.ListTagsSubscriptionsDevice(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTagsSubscriptionsDeviceOptions model
				listTagsSubscriptionsDeviceOptionsModel := new(eventnotificationsv1.ListTagsSubscriptionsDeviceOptions)
				listTagsSubscriptionsDeviceOptionsModel.InstanceID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.ID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.DeviceID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.TagName = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTagsSubscriptionsDeviceOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsSubscriptionsDeviceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = eventNotificationsService.ListTagsSubscriptionsDevice(listTagsSubscriptionsDeviceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTagsSubscriptionsDevice with error: Operation validation and request error`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListTagsSubscriptionsDeviceOptions model
				listTagsSubscriptionsDeviceOptionsModel := new(eventnotificationsv1.ListTagsSubscriptionsDeviceOptions)
				listTagsSubscriptionsDeviceOptionsModel.InstanceID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.ID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.DeviceID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.TagName = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTagsSubscriptionsDeviceOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsSubscriptionsDeviceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := eventNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := eventNotificationsService.ListTagsSubscriptionsDevice(listTagsSubscriptionsDeviceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTagsSubscriptionsDeviceOptions model with no property values
				listTagsSubscriptionsDeviceOptionsModelNew := new(eventnotificationsv1.ListTagsSubscriptionsDeviceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = eventNotificationsService.ListTagsSubscriptionsDevice(listTagsSubscriptionsDeviceOptionsModelNew)
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
			It(`Invoke ListTagsSubscriptionsDevice successfully`, func() {
				eventNotificationsService, serviceErr := eventnotificationsv1.NewEventNotificationsV1(&eventnotificationsv1.EventNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(eventNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListTagsSubscriptionsDeviceOptions model
				listTagsSubscriptionsDeviceOptionsModel := new(eventnotificationsv1.ListTagsSubscriptionsDeviceOptions)
				listTagsSubscriptionsDeviceOptionsModel.InstanceID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.ID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.DeviceID = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.TagName = core.StringPtr("testString")
				listTagsSubscriptionsDeviceOptionsModel.Limit = core.Int64Ptr(int64(1))
				listTagsSubscriptionsDeviceOptionsModel.Offset = core.Int64Ptr(int64(0))
				listTagsSubscriptionsDeviceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := eventNotificationsService.ListTagsSubscriptionsDevice(listTagsSubscriptionsDeviceOptionsModel)
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
				listTagsSubscriptionOptionsModel.Limit = core.Int64Ptr(int64(1))
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
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "tag_subscriptions": [{"id": "ID", "device_id": "DeviceID", "tag_name": "TagName", "user_id": "UserID", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
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
				listTagsSubscriptionOptionsModel.Limit = core.Int64Ptr(int64(1))
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
					fmt.Fprintf(res, "%s", `{"total_count": 10, "offset": 6, "limit": 5, "tag_subscriptions": [{"id": "ID", "device_id": "DeviceID", "tag_name": "TagName", "user_id": "UserID", "updated_at": "2019-01-01T12:00:00.000Z"}]}`)
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
				listTagsSubscriptionOptionsModel.Limit = core.Int64Ptr(int64(1))
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
				listTagsSubscriptionOptionsModel.Limit = core.Int64Ptr(int64(1))
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
				listTagsSubscriptionOptionsModel.Limit = core.Int64Ptr(int64(1))
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
				subscriptionCreateAttributesModel.To = []string{"testString"}

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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {}}`)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {}}`)
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
				subscriptionCreateAttributesModel.To = []string{"testString"}

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
				subscriptionCreateAttributesModel.To = []string{"testString"}

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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {}}`)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {}}`)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {}}`)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "updated_at": "UpdatedAt", "from": "From", "destination_type": "sms_ibm", "destination_id": "DestinationID", "destination_name": "DestinationName", "topic_id": "TopicID", "topic_name": "TopicName", "attributes": {}}`)
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
				createDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDestinationOptionsModel).ToNot(BeNil())
				Expect(createDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Type).To(Equal(core.StringPtr("webhook")))
				Expect(createDestinationOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createDestinationOptionsModel.Config).To(Equal(destinationConfigModel))
				Expect(createDestinationOptionsModel.Certificate).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createDestinationOptionsModel.CertificateContentType).To(Equal(core.StringPtr("testString")))
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
				subscriptionCreateAttributesModel.To = []string{"testString"}
				Expect(subscriptionCreateAttributesModel.To).To(Equal([]string{"testString"}))

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
			It(`Invoke NewGetDestinationDevicesReportOptions successfully`, func() {
				// Construct an instance of the GetDestinationDevicesReportOptions model
				instanceID := "testString"
				id := "testString"
				getDestinationDevicesReportOptionsModel := eventNotificationsService.NewGetDestinationDevicesReportOptions(instanceID, id)
				getDestinationDevicesReportOptionsModel.SetInstanceID("testString")
				getDestinationDevicesReportOptionsModel.SetID("testString")
				getDestinationDevicesReportOptionsModel.SetDays(int64(1))
				getDestinationDevicesReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDestinationDevicesReportOptionsModel).ToNot(BeNil())
				Expect(getDestinationDevicesReportOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getDestinationDevicesReportOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDestinationDevicesReportOptionsModel.Days).To(Equal(core.Int64Ptr(int64(1))))
				Expect(getDestinationDevicesReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewListDestinationDevicesOptions successfully`, func() {
				// Construct an instance of the ListDestinationDevicesOptions model
				instanceID := "testString"
				id := "testString"
				listDestinationDevicesOptionsModel := eventNotificationsService.NewListDestinationDevicesOptions(instanceID, id)
				listDestinationDevicesOptionsModel.SetInstanceID("testString")
				listDestinationDevicesOptionsModel.SetID("testString")
				listDestinationDevicesOptionsModel.SetLimit(int64(1))
				listDestinationDevicesOptionsModel.SetOffset(int64(0))
				listDestinationDevicesOptionsModel.SetSearch("testString")
				listDestinationDevicesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDestinationDevicesOptionsModel).ToNot(BeNil())
				Expect(listDestinationDevicesOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listDestinationDevicesOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listDestinationDevicesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listDestinationDevicesOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listDestinationDevicesOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listDestinationDevicesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				listTagsSubscriptionOptionsModel.SetLimit(int64(1))
				listTagsSubscriptionOptionsModel.SetOffset(int64(0))
				listTagsSubscriptionOptionsModel.SetSearch("testString")
				listTagsSubscriptionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTagsSubscriptionOptionsModel).ToNot(BeNil())
				Expect(listTagsSubscriptionOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionOptionsModel.DeviceID).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionOptionsModel.UserID).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionOptionsModel.TagName).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listTagsSubscriptionOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listTagsSubscriptionOptionsModel.Search).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTagsSubscriptionsDeviceOptions successfully`, func() {
				// Construct an instance of the ListTagsSubscriptionsDeviceOptions model
				instanceID := "testString"
				id := "testString"
				deviceID := "testString"
				listTagsSubscriptionsDeviceOptionsModel := eventNotificationsService.NewListTagsSubscriptionsDeviceOptions(instanceID, id, deviceID)
				listTagsSubscriptionsDeviceOptionsModel.SetInstanceID("testString")
				listTagsSubscriptionsDeviceOptionsModel.SetID("testString")
				listTagsSubscriptionsDeviceOptionsModel.SetDeviceID("testString")
				listTagsSubscriptionsDeviceOptionsModel.SetTagName("testString")
				listTagsSubscriptionsDeviceOptionsModel.SetLimit(int64(1))
				listTagsSubscriptionsDeviceOptionsModel.SetOffset(int64(0))
				listTagsSubscriptionsDeviceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTagsSubscriptionsDeviceOptionsModel).ToNot(BeNil())
				Expect(listTagsSubscriptionsDeviceOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionsDeviceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionsDeviceOptionsModel.DeviceID).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionsDeviceOptionsModel.TagName).To(Equal(core.StringPtr("testString")))
				Expect(listTagsSubscriptionsDeviceOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(listTagsSubscriptionsDeviceOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(0))))
				Expect(listTagsSubscriptionsDeviceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewNotificationApnsBodyMessageData successfully`, func() {
				alert := "testString"
				_model, err := eventNotificationsService.NewNotificationApnsBodyMessageData(alert)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNotificationChromeBodyMessageData successfully`, func() {
				alert := "testString"
				_model, err := eventNotificationsService.NewNotificationChromeBodyMessageData(alert)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNotificationFcmBodyMessageData successfully`, func() {
				alert := "testString"
				_model, err := eventNotificationsService.NewNotificationFcmBodyMessageData(alert)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNotificationFirefoxBodyMessageData successfully`, func() {
				alert := "testString"
				_model, err := eventNotificationsService.NewNotificationFirefoxBodyMessageData(alert)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
			It(`Invoke NewSendNotificationsOptions successfully`, func() {
				// Construct an instance of the Lights model
				lightsModel := new(eventnotificationsv1.Lights)
				Expect(lightsModel).ToNot(BeNil())
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(0))
				lightsModel.LedOffMs = core.StringPtr("testString")
				Expect(lightsModel.LedArgb).To(Equal(core.StringPtr("testString")))
				Expect(lightsModel.LedOnMs).To(Equal(core.Int64Ptr(int64(0))))
				Expect(lightsModel.LedOffMs).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Style model
				styleModel := new(eventnotificationsv1.Style)
				Expect(styleModel).ToNot(BeNil())
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}
				styleModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(styleModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(styleModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(styleModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(styleModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(styleModel.Lines).To(Equal([]string{"testString"}))
				Expect(styleModel.GetProperties()).ToNot(BeEmpty())
				Expect(styleModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				styleModel.SetProperties(nil)
				Expect(styleModel.GetProperties()).To(BeEmpty())

				styleModelExpectedMap := make(map[string]interface{})
				styleModelExpectedMap["foo"] = core.StringPtr("testString")
				styleModel.SetProperties(styleModelExpectedMap)
				styleModelActualMap := styleModel.GetProperties()
				Expect(styleModelActualMap).To(Equal(styleModelExpectedMap))

				// Construct an instance of the NotificationFcmBodyMessageData model
				notificationFcmBodyMessageDataModel := new(eventnotificationsv1.NotificationFcmBodyMessageData)
				Expect(notificationFcmBodyMessageDataModel).ToNot(BeNil())
				notificationFcmBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.CollapseKey = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.InteractiveCategory = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Icon = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.DelayWhileIdle = core.BoolPtr(true)
				notificationFcmBodyMessageDataModel.Sync = core.BoolPtr(true)
				notificationFcmBodyMessageDataModel.Visibility = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Redact = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Payload = make(map[string]interface{})
				notificationFcmBodyMessageDataModel.Priority = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Sound = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationFcmBodyMessageDataModel.Lights = lightsModel
				notificationFcmBodyMessageDataModel.AndroidTitle = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.GroupID = core.StringPtr("testString")
				notificationFcmBodyMessageDataModel.Style = styleModel
				notificationFcmBodyMessageDataModel.Type = core.StringPtr("DEFAULT")
				Expect(notificationFcmBodyMessageDataModel.Alert).To(Equal(core.StringPtr("testString")))
				Expect(notificationFcmBodyMessageDataModel.CollapseKey).To(Equal(core.StringPtr("testString")))
				Expect(notificationFcmBodyMessageDataModel.InteractiveCategory).To(Equal(core.StringPtr("testString")))
				Expect(notificationFcmBodyMessageDataModel.Icon).To(Equal(core.StringPtr("testString")))
				Expect(notificationFcmBodyMessageDataModel.DelayWhileIdle).To(Equal(core.BoolPtr(true)))
				Expect(notificationFcmBodyMessageDataModel.Sync).To(Equal(core.BoolPtr(true)))
				Expect(notificationFcmBodyMessageDataModel.Visibility).To(Equal(core.StringPtr("testString")))
				Expect(notificationFcmBodyMessageDataModel.Redact).To(Equal(core.StringPtr("testString")))
				Expect(notificationFcmBodyMessageDataModel.Payload).To(Equal(make(map[string]interface{})))
				Expect(notificationFcmBodyMessageDataModel.Priority).To(Equal(core.StringPtr("testString")))
				Expect(notificationFcmBodyMessageDataModel.Sound).To(Equal(core.StringPtr("testString")))
				Expect(notificationFcmBodyMessageDataModel.TimeToLive).To(Equal(core.Int64Ptr(int64(0))))
				Expect(notificationFcmBodyMessageDataModel.Lights).To(Equal(lightsModel))
				Expect(notificationFcmBodyMessageDataModel.AndroidTitle).To(Equal(core.StringPtr("testString")))
				Expect(notificationFcmBodyMessageDataModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(notificationFcmBodyMessageDataModel.Style).To(Equal(styleModel))
				Expect(notificationFcmBodyMessageDataModel.Type).To(Equal(core.StringPtr("DEFAULT")))

				// Construct an instance of the NotificationFcmBodyMessageEnData model
				notificationFcmBodyModel := new(eventnotificationsv1.NotificationFcmBodyMessageEnData)
				Expect(notificationFcmBodyModel).ToNot(BeNil())
				notificationFcmBodyModel.EnData = notificationFcmBodyMessageDataModel
				notificationFcmBodyModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(notificationFcmBodyModel.EnData).To(Equal(notificationFcmBodyMessageDataModel))
				Expect(notificationFcmBodyModel.GetProperties()).ToNot(BeEmpty())
				Expect(notificationFcmBodyModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				notificationFcmBodyModel.SetProperties(nil)
				Expect(notificationFcmBodyModel.GetProperties()).To(BeEmpty())

				notificationFcmBodyModelExpectedMap := make(map[string]interface{})
				notificationFcmBodyModelExpectedMap["foo"] = core.StringPtr("testString")
				notificationFcmBodyModel.SetProperties(notificationFcmBodyModelExpectedMap)
				notificationFcmBodyModelActualMap := notificationFcmBodyModel.GetProperties()
				Expect(notificationFcmBodyModelActualMap).To(Equal(notificationFcmBodyModelExpectedMap))

				// Construct an instance of the NotificationApnsBodyMessageData model
				notificationApnsBodyMessageDataModel := new(eventnotificationsv1.NotificationApnsBodyMessageData)
				Expect(notificationApnsBodyMessageDataModel).ToNot(BeNil())
				notificationApnsBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Badge = core.Int64Ptr(int64(38))
				notificationApnsBodyMessageDataModel.InteractiveCategory = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.IosActionKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				notificationApnsBodyMessageDataModel.Sound = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.TitleLocKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.LocKey = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.LaunchImage = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.TitleLocArgs = []string{"testString"}
				notificationApnsBodyMessageDataModel.LocArgs = []string{"testString"}
				notificationApnsBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Subtitle = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.AttachmentURL = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.Type = core.StringPtr("DEFAULT")
				notificationApnsBodyMessageDataModel.ApnsCollapseID = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsThreadID = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				notificationApnsBodyMessageDataModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))
				Expect(notificationApnsBodyMessageDataModel.Alert).To(Equal(core.StringPtr("testString")))
				Expect(notificationApnsBodyMessageDataModel.Badge).To(Equal(core.Int64Ptr(int64(38))))
				Expect(notificationApnsBodyMessageDataModel.InteractiveCategory).To(Equal(core.StringPtr("testString")))
				Expect(notificationApnsBodyMessageDataModel.IosActionKey).To(Equal(core.StringPtr("testString")))
				Expect(notificationApnsBodyMessageDataModel.Payload).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(notificationApnsBodyMessageDataModel.Sound).To(Equal(core.StringPtr("testString")))
				Expect(notificationApnsBodyMessageDataModel.TitleLocKey).To(Equal(core.StringPtr("testString")))
				Expect(notificationApnsBodyMessageDataModel.LocKey).To(Equal(core.StringPtr("testString")))
				Expect(notificationApnsBodyMessageDataModel.LaunchImage).To(Equal(core.StringPtr("testString")))
				Expect(notificationApnsBodyMessageDataModel.TitleLocArgs).To(Equal([]string{"testString"}))
				Expect(notificationApnsBodyMessageDataModel.LocArgs).To(Equal([]string{"testString"}))
				Expect(notificationApnsBodyMessageDataModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(notificationApnsBodyMessageDataModel.Subtitle).To(Equal(core.StringPtr("testString")))
				Expect(notificationApnsBodyMessageDataModel.AttachmentURL).To(Equal(core.StringPtr("testString")))
				Expect(notificationApnsBodyMessageDataModel.Type).To(Equal(core.StringPtr("DEFAULT")))
				Expect(notificationApnsBodyMessageDataModel.ApnsCollapseID).To(Equal(core.StringPtr("testString")))
				Expect(notificationApnsBodyMessageDataModel.ApnsThreadID).To(Equal(core.StringPtr("testString")))
				Expect(notificationApnsBodyMessageDataModel.ApnsGroupSummaryArg).To(Equal(core.StringPtr("testString")))
				Expect(notificationApnsBodyMessageDataModel.ApnsGroupSummaryArgCount).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the NotificationApnsBodyMessageEnData model
				notificationApnsBodyModel := new(eventnotificationsv1.NotificationApnsBodyMessageEnData)
				Expect(notificationApnsBodyModel).ToNot(BeNil())
				notificationApnsBodyModel.EnData = notificationApnsBodyMessageDataModel
				notificationApnsBodyModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(notificationApnsBodyModel.EnData).To(Equal(notificationApnsBodyMessageDataModel))
				Expect(notificationApnsBodyModel.GetProperties()).ToNot(BeEmpty())
				Expect(notificationApnsBodyModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				notificationApnsBodyModel.SetProperties(nil)
				Expect(notificationApnsBodyModel.GetProperties()).To(BeEmpty())

				notificationApnsBodyModelExpectedMap := make(map[string]interface{})
				notificationApnsBodyModelExpectedMap["foo"] = core.StringPtr("testString")
				notificationApnsBodyModel.SetProperties(notificationApnsBodyModelExpectedMap)
				notificationApnsBodyModelActualMap := notificationApnsBodyModel.GetProperties()
				Expect(notificationApnsBodyModelActualMap).To(Equal(notificationApnsBodyModelExpectedMap))

				// Construct an instance of the NotificationDevices model
				notificationDevicesModel := new(eventnotificationsv1.NotificationDevices)
				Expect(notificationDevicesModel).ToNot(BeNil())
				notificationDevicesModel.FcmDevices = []string{"testString"}
				notificationDevicesModel.ApnsDevices = []string{"testString"}
				notificationDevicesModel.UserIds = []string{"testString"}
				notificationDevicesModel.Tags = []string{"testString"}
				notificationDevicesModel.Platforms = []string{"testString"}
				Expect(notificationDevicesModel.FcmDevices).To(Equal([]string{"testString"}))
				Expect(notificationDevicesModel.ApnsDevices).To(Equal([]string{"testString"}))
				Expect(notificationDevicesModel.UserIds).To(Equal([]string{"testString"}))
				Expect(notificationDevicesModel.Tags).To(Equal([]string{"testString"}))
				Expect(notificationDevicesModel.Platforms).To(Equal([]string{"testString"}))

				// Construct an instance of the NotificationChromeBodyMessageData model
				notificationChromeBodyMessageDataModel := new(eventnotificationsv1.NotificationChromeBodyMessageData)
				Expect(notificationChromeBodyMessageDataModel).ToNot(BeNil())
				notificationChromeBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.IconURL = core.StringPtr("testString")
				notificationChromeBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationChromeBodyMessageDataModel.Payload = make(map[string]interface{})
				Expect(notificationChromeBodyMessageDataModel.Alert).To(Equal(core.StringPtr("testString")))
				Expect(notificationChromeBodyMessageDataModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(notificationChromeBodyMessageDataModel.IconURL).To(Equal(core.StringPtr("testString")))
				Expect(notificationChromeBodyMessageDataModel.TimeToLive).To(Equal(core.Int64Ptr(int64(0))))
				Expect(notificationChromeBodyMessageDataModel.Payload).To(Equal(make(map[string]interface{})))

				// Construct an instance of the NotificationChromeBodyMessageEnData model
				notificationChromeBodyModel := new(eventnotificationsv1.NotificationChromeBodyMessageEnData)
				Expect(notificationChromeBodyModel).ToNot(BeNil())
				notificationChromeBodyModel.EnData = notificationChromeBodyMessageDataModel
				notificationChromeBodyModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(notificationChromeBodyModel.EnData).To(Equal(notificationChromeBodyMessageDataModel))
				Expect(notificationChromeBodyModel.GetProperties()).ToNot(BeEmpty())
				Expect(notificationChromeBodyModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				notificationChromeBodyModel.SetProperties(nil)
				Expect(notificationChromeBodyModel.GetProperties()).To(BeEmpty())

				notificationChromeBodyModelExpectedMap := make(map[string]interface{})
				notificationChromeBodyModelExpectedMap["foo"] = core.StringPtr("testString")
				notificationChromeBodyModel.SetProperties(notificationChromeBodyModelExpectedMap)
				notificationChromeBodyModelActualMap := notificationChromeBodyModel.GetProperties()
				Expect(notificationChromeBodyModelActualMap).To(Equal(notificationChromeBodyModelExpectedMap))

				// Construct an instance of the NotificationFirefoxBodyMessageData model
				notificationFirefoxBodyMessageDataModel := new(eventnotificationsv1.NotificationFirefoxBodyMessageData)
				Expect(notificationFirefoxBodyMessageDataModel).ToNot(BeNil())
				notificationFirefoxBodyMessageDataModel.Alert = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.Title = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.IconURL = core.StringPtr("testString")
				notificationFirefoxBodyMessageDataModel.TimeToLive = core.Int64Ptr(int64(0))
				notificationFirefoxBodyMessageDataModel.Payload = make(map[string]interface{})
				Expect(notificationFirefoxBodyMessageDataModel.Alert).To(Equal(core.StringPtr("testString")))
				Expect(notificationFirefoxBodyMessageDataModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(notificationFirefoxBodyMessageDataModel.IconURL).To(Equal(core.StringPtr("testString")))
				Expect(notificationFirefoxBodyMessageDataModel.TimeToLive).To(Equal(core.Int64Ptr(int64(0))))
				Expect(notificationFirefoxBodyMessageDataModel.Payload).To(Equal(make(map[string]interface{})))

				// Construct an instance of the NotificationFirefoxBodyMessageEnData model
				notificationFirefoxBodyModel := new(eventnotificationsv1.NotificationFirefoxBodyMessageEnData)
				Expect(notificationFirefoxBodyModel).ToNot(BeNil())
				notificationFirefoxBodyModel.EnData = notificationFirefoxBodyMessageDataModel
				notificationFirefoxBodyModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(notificationFirefoxBodyModel.EnData).To(Equal(notificationFirefoxBodyMessageDataModel))
				Expect(notificationFirefoxBodyModel.GetProperties()).ToNot(BeEmpty())
				Expect(notificationFirefoxBodyModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				notificationFirefoxBodyModel.SetProperties(nil)
				Expect(notificationFirefoxBodyModel.GetProperties()).To(BeEmpty())

				notificationFirefoxBodyModelExpectedMap := make(map[string]interface{})
				notificationFirefoxBodyModelExpectedMap["foo"] = core.StringPtr("testString")
				notificationFirefoxBodyModel.SetProperties(notificationFirefoxBodyModelExpectedMap)
				notificationFirefoxBodyModelActualMap := notificationFirefoxBodyModel.GetProperties()
				Expect(notificationFirefoxBodyModelActualMap).To(Equal(notificationFirefoxBodyModelExpectedMap))

				// Construct an instance of the SendNotificationsRequestNotificationCreate model
				sendNotificationsRequestModel := new(eventnotificationsv1.SendNotificationsRequestNotificationCreate)
				Expect(sendNotificationsRequestModel).ToNot(BeNil())
				sendNotificationsRequestModel.Data = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmenseverity = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmenfcmbody = notificationFcmBodyModel
				sendNotificationsRequestModel.Ibmenapnsbody = notificationApnsBodyModel
				sendNotificationsRequestModel.Ibmenpushto = notificationDevicesModel
				sendNotificationsRequestModel.Ibmenapnsheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmendefaultshort = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmendefaultlong = core.StringPtr("testString")
				sendNotificationsRequestModel.Ibmenchromebody = notificationChromeBodyModel
				sendNotificationsRequestModel.Ibmenfirefoxbody = notificationFirefoxBodyModel
				sendNotificationsRequestModel.Ibmenchromeheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmenfirefoxheaders = make(map[string]interface{})
				sendNotificationsRequestModel.Ibmensourceid = core.StringPtr("testString")
				sendNotificationsRequestModel.Datacontenttype = core.StringPtr("application/json")
				sendNotificationsRequestModel.Subject = core.StringPtr("testString")
				sendNotificationsRequestModel.ID = core.StringPtr("testString")
				sendNotificationsRequestModel.Source = core.StringPtr("testString")
				sendNotificationsRequestModel.Type = core.StringPtr("testString")
				sendNotificationsRequestModel.Specversion = core.StringPtr("1.0")
				sendNotificationsRequestModel.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
				sendNotificationsRequestModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(sendNotificationsRequestModel.Data).To(Equal(make(map[string]interface{})))
				Expect(sendNotificationsRequestModel.Ibmenseverity).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsRequestModel.Ibmenfcmbody).To(Equal(notificationFcmBodyModel))
				Expect(sendNotificationsRequestModel.Ibmenapnsbody).To(Equal(notificationApnsBodyModel))
				Expect(sendNotificationsRequestModel.Ibmenpushto).To(Equal(notificationDevicesModel))
				Expect(sendNotificationsRequestModel.Ibmenapnsheaders).To(Equal(make(map[string]interface{})))
				Expect(sendNotificationsRequestModel.Ibmendefaultshort).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsRequestModel.Ibmendefaultlong).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsRequestModel.Ibmenchromebody).To(Equal(notificationChromeBodyModel))
				Expect(sendNotificationsRequestModel.Ibmenfirefoxbody).To(Equal(notificationFirefoxBodyModel))
				Expect(sendNotificationsRequestModel.Ibmenchromeheaders).To(Equal(make(map[string]interface{})))
				Expect(sendNotificationsRequestModel.Ibmenfirefoxheaders).To(Equal(make(map[string]interface{})))
				Expect(sendNotificationsRequestModel.Ibmensourceid).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsRequestModel.Datacontenttype).To(Equal(core.StringPtr("application/json")))
				Expect(sendNotificationsRequestModel.Subject).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsRequestModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsRequestModel.Source).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsRequestModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsRequestModel.Specversion).To(Equal(core.StringPtr("1.0")))
				Expect(sendNotificationsRequestModel.Time).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(sendNotificationsRequestModel.GetProperties()).ToNot(BeEmpty())
				Expect(sendNotificationsRequestModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				sendNotificationsRequestModel.SetProperties(nil)
				Expect(sendNotificationsRequestModel.GetProperties()).To(BeEmpty())

				sendNotificationsRequestModelExpectedMap := make(map[string]interface{})
				sendNotificationsRequestModelExpectedMap["foo"] = core.StringPtr("testString")
				sendNotificationsRequestModel.SetProperties(sendNotificationsRequestModelExpectedMap)
				sendNotificationsRequestModelActualMap := sendNotificationsRequestModel.GetProperties()
				Expect(sendNotificationsRequestModelActualMap).To(Equal(sendNotificationsRequestModelExpectedMap))

				// Construct an instance of the SendNotificationsOptions model
				instanceID := "testString"
				sendNotificationsOptionsModel := eventNotificationsService.NewSendNotificationsOptions(instanceID)
				sendNotificationsOptionsModel.SetInstanceID("testString")
				sendNotificationsOptionsModel.SetBody(sendNotificationsRequestModel)
				sendNotificationsOptionsModel.SetCeIbmenseverity("testString")
				sendNotificationsOptionsModel.SetCeIbmendefaultshort("testString")
				sendNotificationsOptionsModel.SetCeIbmendefaultlong("testString")
				sendNotificationsOptionsModel.SetCeIbmenfcmbody(notificationFcmBodyModel)
				sendNotificationsOptionsModel.SetCeIbmenapnsbody(notificationApnsBodyModel)
				sendNotificationsOptionsModel.SetCeIbmenpushto(notificationDevicesModel)
				sendNotificationsOptionsModel.SetCeIbmenapnsheaders(make(map[string]interface{}))
				sendNotificationsOptionsModel.SetCeIbmenchromebody(notificationChromeBodyModel)
				sendNotificationsOptionsModel.SetCeIbmenfirefoxbody(notificationFirefoxBodyModel)
				sendNotificationsOptionsModel.SetCeIbmenchromeheaders(make(map[string]interface{}))
				sendNotificationsOptionsModel.SetCeIbmenfirefoxheaders(make(map[string]interface{}))
				sendNotificationsOptionsModel.SetCeIbmensourceid("testString")
				sendNotificationsOptionsModel.SetCeID("testString")
				sendNotificationsOptionsModel.SetCeSource("testString")
				sendNotificationsOptionsModel.SetCeType("testString")
				sendNotificationsOptionsModel.SetCeSpecversion("1.0")
				sendNotificationsOptionsModel.SetCeTime(CreateMockDateTime("2019-01-01T12:00:00.000Z"))
				sendNotificationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(sendNotificationsOptionsModel).ToNot(BeNil())
				Expect(sendNotificationsOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsOptionsModel.Body).To(Equal(sendNotificationsRequestModel))
				Expect(sendNotificationsOptionsModel.CeIbmenseverity).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsOptionsModel.CeIbmendefaultshort).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsOptionsModel.CeIbmendefaultlong).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsOptionsModel.CeIbmenfcmbody).To(Equal(notificationFcmBodyModel))
				Expect(sendNotificationsOptionsModel.CeIbmenapnsbody).To(Equal(notificationApnsBodyModel))
				Expect(sendNotificationsOptionsModel.CeIbmenpushto).To(Equal(notificationDevicesModel))
				Expect(sendNotificationsOptionsModel.CeIbmenapnsheaders).To(Equal(make(map[string]interface{})))
				Expect(sendNotificationsOptionsModel.CeIbmenchromebody).To(Equal(notificationChromeBodyModel))
				Expect(sendNotificationsOptionsModel.CeIbmenfirefoxbody).To(Equal(notificationFirefoxBodyModel))
				Expect(sendNotificationsOptionsModel.CeIbmenchromeheaders).To(Equal(make(map[string]interface{})))
				Expect(sendNotificationsOptionsModel.CeIbmenfirefoxheaders).To(Equal(make(map[string]interface{})))
				Expect(sendNotificationsOptionsModel.CeIbmensourceid).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsOptionsModel.CeID).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsOptionsModel.CeSource).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsOptionsModel.CeType).To(Equal(core.StringPtr("testString")))
				Expect(sendNotificationsOptionsModel.CeSpecversion).To(Equal(core.StringPtr("1.0")))
				Expect(sendNotificationsOptionsModel.CeTime).To(Equal(CreateMockDateTime("2019-01-01T12:00:00.000Z")))
				Expect(sendNotificationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
				updateDestinationOptionsModel.SetCertificate(CreateMockReader("This is a mock file."))
				updateDestinationOptionsModel.SetCertificateContentType("testString")
				updateDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDestinationOptionsModel).ToNot(BeNil())
				Expect(updateDestinationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateDestinationOptionsModel.Config).To(Equal(destinationConfigModel))
				Expect(updateDestinationOptionsModel.Certificate).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateDestinationOptionsModel.CertificateContentType).To(Equal(core.StringPtr("testString")))
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
			It(`Invoke NewDestinationConfigParamsChromeDestinationConfig successfully`, func() {
				apiKey := "testString"
				websiteURL := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigParamsChromeDestinationConfig(apiKey, websiteURL)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigParamsFcmDestinationConfig successfully`, func() {
				serverKey := "testString"
				senderID := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigParamsFcmDestinationConfig(serverKey, senderID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigParamsFirefoxDestinationConfig successfully`, func() {
				websiteURL := "testString"
				_model, err := eventNotificationsService.NewDestinationConfigParamsFirefoxDestinationConfig(websiteURL)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigParamsIosDestinationConfig successfully`, func() {
				certType := "p8"
				isSandbox := false
				_model, err := eventNotificationsService.NewDestinationConfigParamsIosDestinationConfig(certType, isSandbox)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDestinationConfigParamsWebhookDestinationConfig successfully`, func() {
				url := "testString"
				verb := "get"
				_model, err := eventNotificationsService.NewDestinationConfigParamsWebhookDestinationConfig(url, verb)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSendNotificationsRequestNotificationCreate successfully`, func() {
				ibmenseverity := "testString"
				ibmensourceid := "testString"
				id := "testString"
				source := "testString"
				typeVar := "testString"
				time := CreateMockDateTime("2019-01-01T12:00:00.000Z")
				_model, err := eventNotificationsService.NewSendNotificationsRequestNotificationCreate(ibmenseverity, ibmensourceid, id, source, typeVar, time)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubscriptionCreateAttributesEmailAttributes successfully`, func() {
				to := []string{"testString"}
				addNotificationPayload := false
				replyToMail := "testString"
				replyToName := "testString"
				fromName := "testString"
				_model, err := eventNotificationsService.NewSubscriptionCreateAttributesEmailAttributes(to, addNotificationPayload, replyToMail, replyToName, fromName)
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
			It(`Invoke NewSubscriptionUpdateAttributesEmailUpdateAttributes successfully`, func() {
				var to *eventnotificationsv1.EmailUpdateAttributesTo = nil
				addNotificationPayload := false
				replyToMail := "testString"
				replyToName := "testString"
				fromName := "testString"
				_, err := eventNotificationsService.NewSubscriptionUpdateAttributesEmailUpdateAttributes(to, addNotificationPayload, replyToMail, replyToName, fromName)
				Expect(err).ToNot(BeNil())
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
