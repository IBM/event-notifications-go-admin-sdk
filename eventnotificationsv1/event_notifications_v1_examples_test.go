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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.46.1-a5569134-20220316-164819
 */

package eventnotificationsv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Event Notifications service.
//
// The following configuration properties are assumed to be defined:
// EVENT_NOTIFICATIONS_URL=<service base url>
// EVENT_NOTIFICATIONS_AUTH_TYPE=iam
// EVENT_NOTIFICATIONS_APIKEY=<IAM apikey>
// EVENT_NOTIFICATIONS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../event_notifications_v1.env"

var (
	eventNotificationsService *eventnotificationsv1.EventNotificationsV1
	config                    map[string]string
	configLoaded              bool = false
	instanceID                string
	safariCertificatePath     string
	topicName                 string = "Admin Topic Compliance"
	sourceID                  string = ""
	topicID                   string
	destinationID             string
	destinationID5            string
	subscriptionID            string
	fcmServerKey              string
	fcmSenderId               string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`EventNotificationsV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(eventnotificationsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			instanceID = config["GUID"]
			if instanceID == "" {
				Skip("Unable to load service InstanceID configuration property, skipping tests")
			}
			fmt.Printf("Service GUID: %s\n", instanceID)

			fcmServerKey = config["FCM_KEY"]
			if fcmServerKey == "" {
				Skip("Unable to load service FCM_KEY configuration property, skipping tests")
			}
			fmt.Printf("Service fcmServerKey: %s\n", fcmServerKey)

			fcmSenderId = config["FCM_ID"]
			if fcmSenderId == "" {
				Skip("Unable to load service fcmSenderId configuration property, skipping tests")
			}
			fmt.Printf("Service fcmSenderId: %s\n", fcmSenderId)

			safariCertificatePath = config["SAFARI_CERTIFICATE"]
			if safariCertificatePath == "" {
				Skip("Unable to load service safariCertificatePath configuration property, skipping tests")
			}
			fmt.Printf("Service safariCertificatePath: %s\n", safariCertificatePath)

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			eventNotificationsServiceOptions := &eventnotificationsv1.EventNotificationsV1Options{}

			eventNotificationsService, err = eventnotificationsv1.NewEventNotificationsV1UsingExternalConfig(eventNotificationsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(eventNotificationsService).ToNot(BeNil())
		})
	})

	Describe(`EventNotificationsV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`CreateSources request example`, func() {
			fmt.Println("\nCreateSources() result:")
			// begin-create_sources

			createSourcesOptions := eventNotificationsService.NewCreateSourcesOptions(
				instanceID,
				"Event Notification Create Source Acme",
				"This source is used for Acme Bank",
			)
			createSourcesOptions.SetEnabled(false)

			sourceResponse, response, err := eventNotificationsService.CreateSources(createSourcesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(sourceResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_sources

			sourceID = *sourceResponse.ID

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(sourceResponse).ToNot(BeNil())

		})

		It(`ListSources request example`, func() {
			fmt.Println("\nListSources() result:")
			// begin-list_sources

			listSourcesOptions := eventNotificationsService.NewListSourcesOptions(
				instanceID,
			)

			sourceList, response, err := eventNotificationsService.ListSources(listSourcesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(sourceList, "", "  ")
			fmt.Println(string(b))

			// end-list_sources

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sourceList).ToNot(BeNil())

		})
		It(`GetSource request example`, func() {
			fmt.Println("\nGetSource() result:")
			// begin-get_source

			getSourceOptions := eventNotificationsService.NewGetSourceOptions(
				instanceID,
				sourceID,
			)

			source, response, err := eventNotificationsService.GetSource(getSourceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(source, "", "  ")
			fmt.Println(string(b))

			// end-get_source

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(source).ToNot(BeNil())

		})

		It(`UpdateSource request example`, func() {
			fmt.Println("\nUpdateSource() result:")
			// begin-update_source

			updateSourceOptions := eventNotificationsService.NewUpdateSourceOptions(
				instanceID,
				sourceID,
			)
			updateSourceOptions.SetName(*core.StringPtr("Event Notification update Source Acme"))
			updateSourceOptions.SetDescription(*core.StringPtr("This source is used for updated Acme Bank"))
			updateSourceOptions.SetEnabled(true)

			source, response, err := eventNotificationsService.UpdateSource(updateSourceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(source, "", "  ")
			fmt.Println(string(b))

			// end-update_source

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(source).ToNot(BeNil())

		})

		It(`CreateTopic request example`, func() {
			fmt.Println("\nCreateTopic() result:")
			// begin-create_topic

			rulesModel := &eventnotificationsv1.Rules{
				Enabled:            core.BoolPtr(false),
				EventTypeFilter:    core.StringPtr("$.notification_event_info.event_type == 'cert_manager'"),
				NotificationFilter: core.StringPtr("$.notification.findings[0].severity == 'MODERATE'"),
			}

			topicUpdateSourcesItemModel := &eventnotificationsv1.TopicUpdateSourcesItem{
				ID:    core.StringPtr(sourceID),
				Rules: []eventnotificationsv1.Rules{*rulesModel},
			}

			createTopicOptions := &eventnotificationsv1.CreateTopicOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(topicName),
				Description: core.StringPtr("This topic is used for routing all compliance related notifications to the appropriate destinations"),
				Sources:     []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel},
			}

			topicResponse, response, err := eventNotificationsService.CreateTopic(createTopicOptions)
			if err != nil {
				panic(err)
			}
			topicID = string(*topicResponse.ID)

			b, _ := json.MarshalIndent(topicResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_topic

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(topicResponse).ToNot(BeNil())

		})
		It(`ListTopics request example`, func() {
			fmt.Println("\nListTopics() result:")
			// begin-list_topics

			listTopicsOptions := eventNotificationsService.NewListTopicsOptions(
				instanceID,
			)

			topicList, response, err := eventNotificationsService.ListTopics(listTopicsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(topicList, "", "  ")
			fmt.Println(string(b))

			// end-list_topics

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(topicList).ToNot(BeNil())

		})
		It(`GetTopic request example`, func() {
			fmt.Println("\nGetTopic() result:")
			// begin-get_topic

			getTopicOptions := eventNotificationsService.NewGetTopicOptions(
				instanceID,
				topicID,
			)

			topic, response, err := eventNotificationsService.GetTopic(getTopicOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(topic, "", "  ")
			fmt.Println(string(b))

			// end-get_topic

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(topic).ToNot(BeNil())

		})
		It(`ReplaceTopic request example`, func() {
			fmt.Println("\nReplaceTopic() result:")
			// begin-replace_topic

			rulesModel := &eventnotificationsv1.Rules{
				Enabled:         core.BoolPtr(true),
				EventTypeFilter: core.StringPtr("$.*"),
			}

			topicUpdateSourcesItemModel := &eventnotificationsv1.TopicUpdateSourcesItem{
				ID:    core.StringPtr(sourceID),
				Rules: []eventnotificationsv1.Rules{*rulesModel},
			}

			replaceTopicOptions := eventNotificationsService.NewReplaceTopicOptions(
				instanceID,
				topicID,
			)
			replaceTopicOptions.SetSources([]eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel})
			replaceTopicOptions.SetName("Updated Admin Topic Compliance")

			topic, response, err := eventNotificationsService.ReplaceTopic(replaceTopicOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(topic, "", "  ")
			fmt.Println(string(b))

			// end-replace_topic

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(topic).ToNot(BeNil())

		})
		It(`CreateDestination request example`, func() {
			fmt.Println("\nCreateDestination() result:")
			// begin-create_destination

			createDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"FCM_destination",
				eventnotificationsv1.CreateDestinationOptionsTypePushAndroidConst,
			)

			destinationConfigParamsModel := &eventnotificationsv1.DestinationConfigParamsFcmDestinationConfig{
				ServerKey: core.StringPtr(fcmServerKey),
				SenderID:  core.StringPtr(fcmSenderId),
			}

			destinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsModel,
			}

			createDestinationOptions.SetConfig(destinationConfigModel)

			destinationResponse, response, err := eventNotificationsService.CreateDestination(createDestinationOptions)
			if err != nil {
				panic(err)
			}
			destinationID = string(*destinationResponse.ID)

			b, _ := json.MarshalIndent(destinationResponse, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			createDestinationOptions = eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"Safari_destination",
				eventnotificationsv1.CreateDestinationOptionsTypePushSafariConst,
			)

			certificatefile, err := os.Open(safariCertificatePath)
			if err != nil {
				panic(err)
			}
			createDestinationOptions.Certificate = certificatefile

			destinationConfigParamsSafariModel := &eventnotificationsv1.DestinationConfigParamsSafariDestinationConfig{
				CertType:        core.StringPtr("p12"),
				Password:        core.StringPtr("safari"),
				WebsiteURL:      core.StringPtr("https://ensafaripush.mybluemix.net"),
				WebsiteName:     core.StringPtr("NodeJS Starter Application"),
				URLFormatString: core.StringPtr("https://ensafaripush.mybluemix.net/%@/?flight=%@"),
				WebsitePushID:   core.StringPtr("web.net.mybluemix.ensafaripush"),
			}

			destinationConfigModel = &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsSafariModel,
			}

			createDestinationOptions.SetConfig(destinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(createDestinationOptions)
			if err != nil {
				panic(err)
			}

			b, _ = json.MarshalIndent(destinationResponse, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID5 = *destinationResponse.ID

			// end-create_destination

		})
		It(`ListDestinations request example`, func() {
			fmt.Println("\nListDestinations() result:")
			// begin-list_destinations

			listDestinationsOptions := eventNotificationsService.NewListDestinationsOptions(
				instanceID,
			)

			destinationList, response, err := eventNotificationsService.ListDestinations(listDestinationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(destinationList, "", "  ")
			fmt.Println(string(b))

			// end-list_destinations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationList).ToNot(BeNil())

		})
		It(`GetDestination request example`, func() {
			fmt.Println("\nGetDestination() result:")
			// begin-get_destination

			getDestinationOptions := eventNotificationsService.NewGetDestinationOptions(
				instanceID,
				destinationID,
			)

			destination, response, err := eventNotificationsService.GetDestination(getDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			// end-get_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

		})
		It(`UpdateDestination request example`, func() {
			fmt.Println("\nUpdateDestination() result:")
			// begin-update_destination

			destinationConfigParamsModel := &eventnotificationsv1.DestinationConfigParamsFcmDestinationConfig{
				ServerKey: core.StringPtr(fcmServerKey),
				SenderID:  core.StringPtr(fcmSenderId),
			}
			destinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsModel,
			}

			updateDestinationOptions := eventNotificationsService.NewUpdateDestinationOptions(
				instanceID,
				destinationID,
			)

			updateDestinationOptions.SetName("Admin FCM Compliance")
			updateDestinationOptions.SetDescription("This destination is for creating admin FCM to receive compliance notifications")
			updateDestinationOptions.SetConfig(destinationConfigModel)

			destination, response, err := eventNotificationsService.UpdateDestination(updateDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			// end-update_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

			safaridestinationConfigParamsModel := &eventnotificationsv1.DestinationConfigParamsSafariDestinationConfig{
				CertType:        core.StringPtr("p12"),
				Password:        core.StringPtr("safari"),
				URLFormatString: core.StringPtr("https://ensafaripush.mybluemix.net/%@/?flight=%@"),
				WebsiteName:     core.StringPtr("NodeJS Starter Application"),
				WebsiteURL:      core.StringPtr("https://ensafaripush.mybluemix.net"),
				WebsitePushID:   core.StringPtr("web.net.mybluemix.ensafaripush"),
			}

			safaridestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: safaridestinationConfigParamsModel,
			}

			name := "Safari_dest"
			description := "This destination is for Safari"
			safariupdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID5),
				Name:        core.StringPtr(name),
				Description: core.StringPtr(description),
				Config:      safaridestinationConfigModel,
			}

			certificatefile, err := os.Open(safariCertificatePath)
			if err != nil {
				panic(err)
			}

			safariupdateDestinationOptions.Certificate = certificatefile

			safaridestination, safariresponse, err := eventNotificationsService.UpdateDestination(safariupdateDestinationOptions)

			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(safaridestination, "", "  ")
			fmt.Println(string(b))

			// end-update_destination

			Expect(err).To(BeNil())
			Expect(safariresponse.StatusCode).To(Equal(200))
			Expect(safaridestination).ToNot(BeNil())

		})

		It(`CreateSubscription request example`, func() {
			fmt.Println("\nCreateSubscription() result:")

			subscriptionName := "FCM subscription"
			// begin-create_subscription

			createSubscriptionOptions := eventNotificationsService.NewCreateSubscriptionOptions(
				instanceID,
				subscriptionName,
				destinationID,
				topicID,
			)
			createSubscriptionOptions.SetDescription("Subscription for the FCM")

			subscription, response, err := eventNotificationsService.CreateSubscription(createSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			subscriptionID = string(*subscription.ID)
			b, _ := json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))

			// end-create_subscription

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())

		})
		It(`ListSubscriptions request example`, func() {
			fmt.Println("\nListSubscriptions() result:")
			// begin-list_subscriptions

			listSubscriptionsOptions := eventNotificationsService.NewListSubscriptionsOptions(
				instanceID,
			)

			subscriptionList, response, err := eventNotificationsService.ListSubscriptions(listSubscriptionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(subscriptionList, "", "  ")
			fmt.Println(string(b))

			// end-list_subscriptions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscriptionList).ToNot(BeNil())

		})
		It(`GetSubscription request example`, func() {
			fmt.Println("\nGetSubscription() result:")
			// begin-get_subscription

			getSubscriptionOptions := eventNotificationsService.NewGetSubscriptionOptions(
				instanceID,
				subscriptionID,
			)

			subscription, response, err := eventNotificationsService.GetSubscription(getSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))

			// end-get_subscription

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())

		})
		It(`UpdateSubscription request example`, func() {
			fmt.Println("\nUpdateSubscription() result:")
			// begin-update_subscription

			updateSubscriptionOptions := eventNotificationsService.NewUpdateSubscriptionOptions(
				instanceID,
				subscriptionID,
			)

			updateSubscriptionOptions.SetDescription("Update FCM subscription")
			updateSubscriptionOptions.SetName("Update_FCM_subscription")

			subscription, response, err := eventNotificationsService.UpdateSubscription(updateSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))

			// end-update_subscription

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())

		})
		It(`SendNotifications request example`, func() {
			fmt.Println("\nSendNotifications() result:")

			notificationID := "1234-1234-sdfs-234"
			notificationSeverity := "MEDIUM"
			typeValue := "com.acme.offer:new"
			//userId := "userId"
			notificationsSouce := "1234-1234-sdfs-234:test"
			specVersion := "1.0"

			// begin-send_notifications

			notificationCreateModel := &eventnotificationsv1.NotificationCreate{}

			notificationCreateModel.Ibmenseverity = &notificationSeverity
			notificationCreateModel.ID = &notificationID
			notificationCreateModel.Source = &notificationsSouce
			notificationCreateModel.Ibmensourceid = &sourceID
			notificationCreateModel.Type = &typeValue
			notificationCreateModel.Time = &strfmt.DateTime{}
			notificationCreateModel.Specversion = &specVersion

			notificationDevicesModel := "{\"user_ids\": [\"userId\"]}"
			notificationSafariBodyModel := "{\"en_data\": {\"alert\": \"Alert message\"}}"

			notificationCreateModel.Ibmenpushto = &notificationDevicesModel

			apnsOptions := map[string]interface{}{
				"aps": map[string]interface{}{
					"alert": "alert message",
					"badge": 5,
				},
			}

			ibmenapnsbody, _ := json.Marshal(apnsOptions)
			ibmenapnsbodyString := string(ibmenapnsbody)

			fcmOptions := map[string]interface{}{
				"notification": map[string]interface{}{
					"title": "alert title",
					"body":  "alert message",
				},
			}
			ibmenfcmbody, _ := json.Marshal(fcmOptions)
			ibmenfcmbodyString := string(ibmenfcmbody)

			apnsHeaders := map[string]interface{}{
				"apns-collapse-id": "collapse-id",
			}
			ibmenapnsheaderbody, _ := json.Marshal(apnsHeaders)
			ibmenapnsheaderstring := string(ibmenapnsheaderbody)

			notificationCreateModel.Ibmenfcmbody = &ibmenfcmbodyString
			notificationCreateModel.Ibmenapnsbody = &ibmenapnsbodyString
			notificationCreateModel.Ibmenapnsheaders = &ibmenapnsheaderstring
			notificationCreateModel.Ibmensafaribody = &notificationSafariBodyModel
			notificationCreateModel.Ibmendefaultshort = core.StringPtr("Offer Alert")
			notificationCreateModel.Ibmendefaultlong = core.StringPtr("Alert on expiring offers")

			sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
			sendNotificationsOptionsModel.InstanceID = &instanceID
			sendNotificationsOptionsModel.Body = notificationCreateModel

			notificationResponse, response, err := eventNotificationsService.SendNotifications(sendNotificationsOptionsModel)

			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(notificationResponse, "", "  ")
			fmt.Println(string(b))

			// end-send_notifications

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(notificationResponse).ToNot(BeNil())

		})

		It(`SendNotifications 2 request example`, func() {
			fmt.Println("\nSendNotifications() result:")

			notificationID := "1234-1234-sdfs-234"
			notificationSeverity := "MEDIUM"
			typeValue := "com.acme.offer:new"
			//userId := "userId"
			notificationsSouce := "1234-1234-sdfs-234:test"
			specVersion := "1.0"
			// begin-send_notifications

			notificationCreateModel := &eventnotificationsv1.NotificationCreate{}

			notificationCreateModel.Ibmenseverity = &notificationSeverity
			notificationCreateModel.ID = &notificationID
			notificationCreateModel.Source = &notificationsSouce
			notificationCreateModel.Ibmensourceid = &sourceID
			notificationCreateModel.Type = &typeValue
			notificationCreateModel.Time = &strfmt.DateTime{}
			notificationCreateModel.Specversion = &specVersion
			notificationCreateModel.Ibmendefaultshort = core.StringPtr("Offer Alert")
			notificationCreateModel.Ibmendefaultlong = core.StringPtr("Alert on expiring offers")

			notificationDevicesModel := "{\"user_ids\": [\"userId\"]}"

			notificationCreateModel.Ibmenpushto = &notificationDevicesModel

			apnsOptions := map[string]interface{}{
				"aps": map[string]interface{}{
					"alert": "alert message",
					"badge": 5,
				},
			}

			ibmenapnsbody, _ := json.Marshal(apnsOptions)
			ibmenapnsbodyString := string(ibmenapnsbody)

			fcmOptions := map[string]interface{}{
				"notification": map[string]interface{}{
					"title": "alert title",
					"body":  "alert message",
				},
			}
			ibmenfcmbody, _ := json.Marshal(fcmOptions)
			ibmenfcmbodyString := string(ibmenfcmbody)

			apnsHeaders := map[string]interface{}{
				"apns-collapse-id": "collapse-id",
			}
			ibmenapnsheaderbody, _ := json.Marshal(apnsHeaders)
			ibmenapnsheaderstring := string(ibmenapnsheaderbody)

			notificationCreateModel.Ibmenfcmbody = &ibmenfcmbodyString
			notificationCreateModel.Ibmenapnsbody = &ibmenapnsbodyString
			notificationCreateModel.Ibmenapnsheaders = &ibmenapnsheaderstring

			sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
			sendNotificationsOptionsModel.InstanceID = &instanceID
			sendNotificationsOptionsModel.Body = notificationCreateModel

			notificationResponse, response, err := eventNotificationsService.SendNotifications(sendNotificationsOptionsModel)

			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(notificationResponse, "", "  ")
			fmt.Println(string(b))

			// end-send_notifications

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(notificationResponse).ToNot(BeNil())

		})

		It(`DeleteSubscription request example`, func() {
			// begin-delete_subscription

			deleteSubscriptionOptions := eventNotificationsService.NewDeleteSubscriptionOptions(
				instanceID,
				subscriptionID,
			)

			response, err := eventNotificationsService.DeleteSubscription(deleteSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteSubscription(): %d\n", response.StatusCode)
			}

			// end-delete_subscription

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTopic request example`, func() {
			// begin-delete_topic

			deleteTopicOptions := eventNotificationsService.NewDeleteTopicOptions(
				instanceID,
				topicID,
			)

			response, err := eventNotificationsService.DeleteTopic(deleteTopicOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_topic

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteDestination request example`, func() {
			// begin-delete_destination

			deleteDestinationOptions := eventNotificationsService.NewDeleteDestinationOptions(
				instanceID,
				destinationID,
			)

			response, err := eventNotificationsService.DeleteDestination(deleteDestinationOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			deleteDestinationOptions = eventNotificationsService.NewDeleteDestinationOptions(
				instanceID,
				destinationID5,
			)

			response, err = eventNotificationsService.DeleteDestination(deleteDestinationOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteSource request example`, func() {
			// begin-delete_source

			deleteSourceOptions := eventNotificationsService.NewDeleteSourceOptions(
				instanceID,
				sourceID,
			)

			response, err := eventNotificationsService.DeleteSource(deleteSourceOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteSource(): %d\n", response.StatusCode)
			}

			// end-delete_source

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
