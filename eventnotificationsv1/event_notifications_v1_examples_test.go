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
	"encoding/json"
	"fmt"
	"os"
	"time"

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
	topicName                 string = "Admin Topic Compliance"
	sourceID                  string = ""
	topicID                   string
	destinationID             string
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
			sourceID = string(*sourceList.Sources[len(sourceList.Sources)-1].ID)

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
				"GCM_destination",
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

			// end-create_destination

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

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

			// destinationConfigParamsModel1 := &eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig{
			// 	URL:              core.StringPtr("https://cloud.ibm.com/nhwebhook/sendwebhook"),
			// 	Verb:             core.StringPtr("post"),
			// 	SensitiveHeaders: []string{"authorization"},
			// }

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

		})
		It(`ListDestinationDevices request example`, func() {
			fmt.Println("\nListDestinationDevices() result:")
			// begin-list_destination_devices

			listDestinationDevicesOptions := eventNotificationsService.NewListDestinationDevicesOptions(
				instanceID,
				destinationID,
			)

			destinationDevicesList, response, err := eventNotificationsService.ListDestinationDevices(listDestinationDevicesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(destinationDevicesList, "", "  ")
			fmt.Println(string(b))

			// end-list_destination_devices

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationDevicesList).ToNot(BeNil())

		})
		It(`GetDestinationDevicesReport request example`, func() {
			fmt.Println("\nGetDestinationDevicesReport() result:")
			// begin-get_destination_devices_report

			getDestinationDevicesReportOptions := eventNotificationsService.NewGetDestinationDevicesReportOptions(
				instanceID,
				destinationID,
			)

			destinationDevicesReport, response, err := eventNotificationsService.GetDestinationDevicesReport(getDestinationDevicesReportOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(destinationDevicesReport, "", "  ")
			fmt.Println(string(b))

			// end-get_destination_devices_report

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationDevicesReport).ToNot(BeNil())

		})
		/*
			It(`CreateTagsSubscription request example`, func() {
				fmt.Println("\nCreateTagsSubscription() result:")

				tagName := "IBM_test"
				// begin-create_tags_subscription

				createTagsSubscriptionOptions := eventNotificationsService.NewCreateTagsSubscriptionOptions(
					instanceID,
					destinationID,
					destinationDeviceID,
					tagName,
				)

				destinationTagsSubscriptionResponse, response, err := eventNotificationsService.CreateTagsSubscription(createTagsSubscriptionOptions)
				if err != nil {
					panic(err)
				}
				b, _ := json.MarshalIndent(destinationTagsSubscriptionResponse, "", "  ")
				fmt.Println(string(b))

				// end-create_tags_subscription

				tagSubscriptionID = string(*destinationTagsSubscriptionResponse.ID)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(201))
				Expect(destinationTagsSubscriptionResponse).ToNot(BeNil())

			})
			It(`ListTagsSubscription request example`, func() {
				fmt.Println("\nListTagsSubscription() result:")
				// begin-list_tags_subscription

				listTagsSubscriptionOptions := eventNotificationsService.NewListTagsSubscriptionOptions(
					instanceID,
					destinationID,
				)

				tagsSubscriptionList, response, err := eventNotificationsService.ListTagsSubscription(listTagsSubscriptionOptions)
				if err != nil {
					panic(err)
				}
				b, _ := json.MarshalIndent(tagsSubscriptionList, "", "  ")
				fmt.Println(string(b))

				// end-list_tags_subscription

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(tagsSubscriptionList).ToNot(BeNil())

			})

			It(`ListTagsSubscriptionsDevice request example`, func() {
				fmt.Println("\nListTagsSubscriptionsDevice() result:")
				// begin-list_tags_subscriptions_device

				listTagsSubscriptionsDeviceOptions := eventNotificationsService.NewListTagsSubscriptionsDeviceOptions(
					instanceID,
					destinationID,
					destinationDeviceID,
				)

				tagsSubscriptionList, response, err := eventNotificationsService.ListTagsSubscriptionsDevice(listTagsSubscriptionsDeviceOptions)
				if err != nil {
					panic(err)
				}
				b, _ := json.MarshalIndent(tagsSubscriptionList, "", "  ")
				fmt.Println(string(b))

				// end-list_tags_subscriptions_device

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(tagsSubscriptionList).ToNot(BeNil())

			})
		*/
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
			notificationSubject := "FCM_Subject"
			notificationSeverity := "MEDIUM"
			typeValue := "com.acme.offer:new"
			now := time.Now()
			date := strfmt.DateTime(now)
			userId := "userId"
			alertMessage := "Message"
			notificationsSouce := "1234-1234-sdfs-234:test"

			// begin-send_notifications

			sendNotificationsOptions := eventNotificationsService.NewSendNotificationsOptions(
				instanceID,
				notificationSubject,
				notificationSeverity,
				notificationID,
				notificationsSouce,
				sourceID,
				typeValue,
				&date,
			)

			sendNotificationsOptions.PushTo = &eventnotificationsv1.NotificationFcmDevices{
				UserIds: []string{userId},
			}

			sendNotificationsOptions.MessageFcmBody = &eventnotificationsv1.NotificationFcmBody{
				&eventnotificationsv1.NotificationFcmBodyMessage{
					Data: &eventnotificationsv1.NotificationFcmBodyMessageData{
						Alert:          core.StringPtr(alertMessage),
						DelayWhileIdle: core.BoolPtr(true),
						TimeToLive:     core.Int64Ptr(100),
					},
				},
			}

			notificationResponse, response, err := eventNotificationsService.SendNotifications(sendNotificationsOptions)
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

			// end-delete_subscription
			fmt.Printf("\nDeleteSubscription() response status code: %d\n", response.StatusCode)

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
			fmt.Printf("\nDeleteTopic() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		/*
			It(`DeleteTagsSubscription request example`, func() {

				tagName := "IBM_test"
				// begin-delete_tags_subscription

				deleteTagsSubscriptionOptions := eventNotificationsService.NewDeleteTagsSubscriptionOptions(
					instanceID,
					destinationID,
				)

				deleteTagsSubscriptionOptions.SetDeviceID(destinationDeviceID)
				deleteTagsSubscriptionOptions.SetTagName(tagName)
				response, err := eventNotificationsService.DeleteTagsSubscription(deleteTagsSubscriptionOptions)
				if err != nil {
					panic(err)
				}

				// end-delete_tags_subscription
				fmt.Printf("\nDeleteTagsSubscription() response status code: %d\n", response.StatusCode)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(204))

			})*/
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
			fmt.Printf("\nDeleteDestination() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
