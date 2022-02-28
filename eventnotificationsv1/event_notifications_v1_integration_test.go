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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the eventnotificationsv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`EventNotificationsV1 Integration Tests`, func() {

	const externalConfigFile = "../event_notifications_v1.env"

	var (
		err                       error
		eventNotificationsService *eventnotificationsv1.EventNotificationsV1
		serviceURL                string
		config                    map[string]string
		instanceID                string
		search                    string = ""
		topicName                 string = "WebhookTopic"
		sourceID                  string
		topicID                   string
		topicID2                  string
		topicID3                  string
		destinationID             string
		destinationID2            string
		destinationID3            string
		subscriptionID            string
		subscriptionID2           string
		subscriptionID3           string
		fcmServerKey              string
		fcmSenderId               string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(eventnotificationsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)

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

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			eventNotificationsServiceOptions := &eventnotificationsv1.EventNotificationsV1Options{}

			eventNotificationsService, err = eventnotificationsv1.NewEventNotificationsV1UsingExternalConfig(eventNotificationsServiceOptions)

			Expect(err).To(BeNil())
			Expect(eventNotificationsService).ToNot(BeNil())
			Expect(eventNotificationsService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			eventNotificationsService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`ListSources - List all Sources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSources(listSourcesOptions *ListSourcesOptions)`, func() {

			listSourcesOptions := &eventnotificationsv1.ListSourcesOptions{
				InstanceID: core.StringPtr(instanceID),
				Limit:      core.Int64Ptr(int64(1)),
				Offset:     core.Int64Ptr(int64(0)),
				Search:     core.StringPtr(search),
			}

			sourceList, response, err := eventNotificationsService.ListSources(listSourcesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sourceList).ToNot(BeNil())
			//	Expect(sourceList.TotalCount).To(Equal(core.Int64Ptr(3)))
			sourceID = *sourceList.Sources[len(sourceList.Sources)-1].ID

			listSourcesOptions = &eventnotificationsv1.ListSourcesOptions{
				InstanceID: core.StringPtr(instanceID),
				Limit:      core.Int64Ptr(int64(1)),
				Offset:     core.Int64Ptr(int64(1)),
				Search:     core.StringPtr(search),
			}

			sourceList, response, err = eventNotificationsService.ListSources(listSourcesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sourceList).ToNot(BeNil())
			sourceID1 := *sourceList.Sources[len(sourceList.Sources)-1].ID

			Expect(sourceID).ToNot(Equal(sourceID1))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`GetSource - Get a Source`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSource(getSourceOptions *GetSourceOptions)`, func() {

			getSourceOptions := &eventnotificationsv1.GetSourceOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(sourceID),
			}

			source, response, err := eventNotificationsService.GetSource(getSourceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(source).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})

	Describe(`CreateTopic - Create a new Topic`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTopic(createTopicOptions *CreateTopicOptions)`, func() {

			rulesModel := &eventnotificationsv1.Rules{
				Enabled:            core.BoolPtr(false),
				EventTypeFilter:    core.StringPtr("$.notification_event_info.event_type == 'cert_manager'"),
				NotificationFilter: core.StringPtr("$.notification.findings[0].severity == 'MODERATE'"),
			}

			topicUpdateSourcesItemModel := &eventnotificationsv1.TopicUpdateSourcesItem{
				ID:    core.StringPtr(sourceID),
				Rules: []eventnotificationsv1.Rules{*rulesModel},
			}

			description := core.StringPtr("Topic for Webhook notifications")
			name := core.StringPtr(topicName)
			createTopicOptions := &eventnotificationsv1.CreateTopicOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        name,
				Description: description,
				Sources:     []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel},
			}

			topicResponse, response, err := eventNotificationsService.CreateTopic(createTopicOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(topicResponse).ToNot(BeNil())
			Expect(topicResponse.Name).To(Equal(name))
			Expect(topicResponse.Description).To(Equal(description))

			topicID = *topicResponse.ID

			description = core.StringPtr("Topic 2 for Webhook notifications")
			name = core.StringPtr("topic2")
			createTopicOptions = &eventnotificationsv1.CreateTopicOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        name,
				Description: description,
				Sources:     []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel},
			}

			topicResponse, response, err = eventNotificationsService.CreateTopic(createTopicOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(topicResponse).ToNot(BeNil())
			Expect(topicResponse.Name).To(Equal(name))
			Expect(topicResponse.Description).To(Equal(description))

			topicID2 = *topicResponse.ID

			Expect(topicID).ToNot(Equal(topicID2))

			rulesModel = &eventnotificationsv1.Rules{
				Enabled:            core.BoolPtr(false),
				EventTypeFilter:    core.StringPtr("$.notification_event_info.event_type == 'cert_manager'"),
				NotificationFilter: core.StringPtr("$.notification.findings[0].severity == 'MODERATE'"),
			}

			topicUpdateSourcesItemModel = &eventnotificationsv1.TopicUpdateSourcesItem{
				ID:    core.StringPtr(sourceID),
				Rules: []eventnotificationsv1.Rules{*rulesModel},
			}

			createTopicOptions = &eventnotificationsv1.CreateTopicOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr("FCM_topic"),
				Description: core.StringPtr("This topic is used for routing all compliance related notifications to the appropriate destinations"),
				Sources:     []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel},
			}

			topicResponse, response, err = eventNotificationsService.CreateTopic(createTopicOptions)
			if err != nil {
				panic(err)
			}

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(topicResponse).ToNot(BeNil())

			topicID3 = string(*topicResponse.ID)

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			// 409
			// 415
			// 500
			//
		})
	})

	Describe(`ListTopics - List all Topics`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTopics(listTopicsOptions *ListTopicsOptions)`, func() {

			listTopicsOptions := &eventnotificationsv1.ListTopicsOptions{
				InstanceID: core.StringPtr(instanceID),
				Limit:      core.Int64Ptr(int64(1)),
				Offset:     core.Int64Ptr(int64(0)),
				Search:     core.StringPtr(search),
			}

			topicList, response, err := eventNotificationsService.ListTopics(listTopicsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(topicList).ToNot(BeNil())

			topicid1 := topicList.Topics[0].ID

			listTopicsOptions = &eventnotificationsv1.ListTopicsOptions{
				InstanceID: core.StringPtr(instanceID),
				Limit:      core.Int64Ptr(int64(1)),
				Offset:     core.Int64Ptr(int64(1)),
				Search:     core.StringPtr(search),
			}

			topicList, response, err = eventNotificationsService.ListTopics(listTopicsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(topicList).ToNot(BeNil())
			topicid2 := topicList.Topics[0].ID
			Expect(topicid1).ToNot(Equal(topicid2))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`GetTopic - Get details of a Topic`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTopic(getTopicOptions *GetTopicOptions)`, func() {

			getTopicOptions := &eventnotificationsv1.GetTopicOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(topicID),
				Include:    core.StringPtr(""),
			}

			topic, response, err := eventNotificationsService.GetTopic(getTopicOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(topic).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})

	Describe(`ReplaceTopic - Update details of a Topic`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceTopic(replaceTopicOptions *ReplaceTopicOptions)`, func() {

			rulesModel := &eventnotificationsv1.Rules{
				Enabled:            core.BoolPtr(true),
				EventTypeFilter:    core.StringPtr("$.notification_event_info.event_type == 'core_cert_manager'"),
				NotificationFilter: core.StringPtr("$.notification.findings[0].severity == 'SEVERE'"),
			}

			topicUpdateSourcesItemModel := &eventnotificationsv1.TopicUpdateSourcesItem{
				ID:    core.StringPtr(sourceID),
				Rules: []eventnotificationsv1.Rules{*rulesModel},
			}

			description := core.StringPtr("Updated Topic for Webhook notifications")
			name := core.StringPtr(topicName)

			replaceTopicOptions := &eventnotificationsv1.ReplaceTopicOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(topicID),
				Name:        name,
				Description: description,
				Sources:     []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel},
			}

			topic, response, err := eventNotificationsService.ReplaceTopic(replaceTopicOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(topic).ToNot(BeNil())
			Expect(topic.Name).To(Equal(name))
			Expect(topic.ID).To(Equal(core.StringPtr(topicID)))
			Expect(topic.Description).To(Equal(description))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			// 409
			// 415
			// 500
			//
		})
	})

	Describe(`CreateDestination - Create a new Destination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDestination(createDestinationOptions *CreateDestinationOptions)`, func() {

			destinationConfigParamsModel := &eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig{
				URL:              core.StringPtr("https://gcm.com"),
				Verb:             core.StringPtr("get"),
				CustomHeaders:    make(map[string]string),
				SensitiveHeaders: []string{"gcm_apikey"},
			}

			destinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsModel,
			}

			name := "Webhook_destination"
			typeVal := "webhook"
			description := "Webhook Destination"
			createDestinationOptions := &eventnotificationsv1.CreateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(typeVal),
				Description: core.StringPtr(description),
				Config:      destinationConfigModel,
			}

			destinationResponse, response, err := eventNotificationsService.CreateDestination(createDestinationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())
			Expect(destinationResponse.Name).To(Equal(core.StringPtr(name)))
			Expect(destinationResponse.Type).To(Equal(core.StringPtr(typeVal)))
			Expect(destinationResponse.Description).To(Equal(core.StringPtr(description)))

			destinationID = *destinationResponse.ID

			createDestinationOptions = eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"FCM_destination",
				eventnotificationsv1.CreateDestinationOptionsTypePushAndroidConst,
			)

			destinationConfigParamsFCMModel := &eventnotificationsv1.DestinationConfigParamsFcmDestinationConfig{
				ServerKey: core.StringPtr(fcmServerKey),
				SenderID:  core.StringPtr(fcmSenderId),
			}

			destinationConfigModel = &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsFCMModel,
			}

			createDestinationOptions.SetConfig(destinationConfigModel)

			destinationResponse, response, err = eventNotificationsService.CreateDestination(createDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID3 = *destinationResponse.ID

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 409
			// 415
			// 500
			//
		})
	})

	Describe(`ListDestinations - List all Destinations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDestinations(listDestinationsOptions *ListDestinationsOptions)`, func() {

			listDestinationsOptions := &eventnotificationsv1.ListDestinationsOptions{
				InstanceID: core.StringPtr(instanceID),
				Limit:      core.Int64Ptr(int64(1)),
				Offset:     core.Int64Ptr(int64(0)),
				Search:     core.StringPtr(search),
			}

			destinationList, response, err := eventNotificationsService.ListDestinations(listDestinationsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationList).ToNot(BeNil())

			destinationId1 := destinationList.Destinations[0].ID

			listDestinationsOptions = &eventnotificationsv1.ListDestinationsOptions{
				InstanceID: core.StringPtr(instanceID),
				Limit:      core.Int64Ptr(int64(1)),
				Offset:     core.Int64Ptr(int64(1)),
				Search:     core.StringPtr(search),
			}

			destinationList, response, err = eventNotificationsService.ListDestinations(listDestinationsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationList).ToNot(BeNil())

			destinationId2 := destinationList.Destinations[0].ID
			Expect(destinationId2).ToNot(Equal(destinationId1))

			listDestinationsOptions = &eventnotificationsv1.ListDestinationsOptions{
				InstanceID: core.StringPtr(instanceID),
				Limit:      core.Int64Ptr(int64(10)),
				Offset:     core.Int64Ptr(int64(0)),
				Search:     core.StringPtr(search),
			}

			destinationList, response, err = eventNotificationsService.ListDestinations(listDestinationsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationList).ToNot(BeNil())
			//Expect(len(destinationList.Destinations)).To(Equal(4))

			for _, ID := range destinationList.Destinations {
				if destinationID != *ID.ID && *ID.Type == "smtp_ibm" {
					destinationID2 = *ID.ID
					break
				}
			}

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`GetDestination - Get details of a Destination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDestination(getDestinationOptions *GetDestinationOptions)`, func() {

			getDestinationOptions := &eventnotificationsv1.GetDestinationOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(destinationID),
			}

			destination, response, err := eventNotificationsService.GetDestination(getDestinationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())
			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})

	Describe(`UpdateDestination - Update details of a Destination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDestination(updateDestinationOptions *UpdateDestinationOptions)`, func() {

			destinationConfigParamsModel := &eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig{
				URL:              core.StringPtr("https://cloud.ibm.com/nhwebhook/sendwebhook"),
				Verb:             core.StringPtr("post"),
				CustomHeaders:    make(map[string]string),
				SensitiveHeaders: []string{"authorization"},
			}

			destinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsModel,
			}

			name := "Admin Webhook Compliance"
			description := "This destination is for creating admin Webhook to receive compliance notifications"
			updateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID),
				Name:        core.StringPtr(name),
				Description: core.StringPtr(description),
				Config:      destinationConfigModel,
			}

			destination, response, err := eventNotificationsService.UpdateDestination(updateDestinationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())
			Expect(destination.ID).To(Equal(core.StringPtr(destinationID)))
			Expect(destination.Name).To(Equal(core.StringPtr(name)))
			Expect(destination.Description).To(Equal(core.StringPtr(description)))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			// 409
			// 415
			// 500
			//
		})
	})

	Describe(`ListDestinationDevices - Get list of Destination devices`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDestinationDevices(listDestinationDevicesOptions *ListDestinationDevicesOptions)`, func() {

			listDestinationDevicesOptions := &eventnotificationsv1.ListDestinationDevicesOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(destinationID3),
				Limit:      core.Int64Ptr(int64(1)),
				Offset:     core.Int64Ptr(int64(0)),
				Search:     core.StringPtr(""),
			}

			destinationDevicesList, response, err := eventNotificationsService.ListDestinationDevices(listDestinationDevicesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationDevicesList).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})

	Describe(`GetDestinationDevicesReport - Retrieves report of destination devices registered`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDestinationDevicesReport(getDestinationDevicesReportOptions *GetDestinationDevicesReportOptions)`, func() {

			getDestinationDevicesReportOptions := &eventnotificationsv1.GetDestinationDevicesReportOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(destinationID3),
				Days:       core.Int64Ptr(int64(1)),
			}

			destinationDevicesReport, response, err := eventNotificationsService.GetDestinationDevicesReport(getDestinationDevicesReportOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destinationDevicesReport).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})

	/*
		Describe(`CreateTagsSubscription - Create a new Tag subscription`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`CreateTagsSubscription(createTagsSubscriptionOptions *CreateTagsSubscriptionOptions)`, func() {

				tagName := "IBM_test"
				createTagsSubscriptionOptions := &eventnotificationsv1.CreateTagsSubscriptionOptions{
					InstanceID: core.StringPtr(instanceID),
					ID:         core.StringPtr(destinationID3),
					DeviceID:   core.StringPtr(destinationDeviceID),
					TagName:    core.StringPtr(tagName),
				}

				destinationTagsSubscriptionResponse, response, err := eventNotificationsService.CreateTagsSubscription(createTagsSubscriptionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(201))
				Expect(destinationTagsSubscriptionResponse).ToNot(BeNil())

				//
				// The following status codes aren't covered by tests.
				// Please provide integration tests for these too.
				//
				// 400
				// 401
				// 409
				// 415
				// 500
				//
			})
		})

		Describe(`ListTagsSubscription - List all Tag Subscriptions`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ListTagsSubscription(listTagsSubscriptionOptions *ListTagsSubscriptionOptions)`, func() {

				listTagsSubscriptionOptions := &eventnotificationsv1.ListTagsSubscriptionOptions{
					InstanceID: core.StringPtr(instanceID),
					ID:         core.StringPtr(destinationID3),
				}

				tagsSubscriptionList, response, err := eventNotificationsService.ListTagsSubscription(listTagsSubscriptionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(tagsSubscriptionList).ToNot(BeNil())

				//
				// The following status codes aren't covered by tests.
				// Please provide integration tests for these too.
				//
				// 401
				// 500
				//
			})
		})

		Describe(`ListTagsSubscriptionsDevice - List all Tag Subscriptions for a device`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`ListTagsSubscriptionsDevice(listTagsSubscriptionsDeviceOptions *ListTagsSubscriptionsDeviceOptions)`, func() {

				listTagsSubscriptionsDeviceOptions := &eventnotificationsv1.ListTagsSubscriptionsDeviceOptions{
					InstanceID: core.StringPtr(instanceID),
					ID:         core.StringPtr(destinationID3),
					DeviceID:   core.StringPtr(destinationDeviceID),
				}

				tagsSubscriptionList, response, err := eventNotificationsService.ListTagsSubscriptionsDevice(listTagsSubscriptionsDeviceOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(tagsSubscriptionList).ToNot(BeNil())

				//
				// The following status codes aren't covered by tests.
				// Please provide integration tests for these too.
				//
				// 401
				// 500
				//
			})
		})
	*/

	Describe(`CreateSubscription - Create a new Subscription`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSubscription(createSubscriptionOptions *CreateSubscriptionOptions)`, func() {

			subscriptionCreateAttributesModel := &eventnotificationsv1.SubscriptionCreateAttributes{
				SigningEnabled: core.BoolPtr(false),
			}

			name := core.StringPtr("subscription_web")
			description := core.StringPtr("Subscription for web")
			createSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          name,
				Description:   description,
				DestinationID: core.StringPtr(destinationID),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreateAttributesModel,
			}

			subscription, response, err := eventNotificationsService.CreateSubscription(createSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.Attributes).ToNot(BeNil())
			Expect(subscription.Description).To(Equal(description))
			Expect(subscription.Name).To(Equal(name))
			subscriptionID = *subscription.ID

			subscriptionCreateAttributesEmailModel := &eventnotificationsv1.SubscriptionCreateAttributesEmailAttributes{
				To:                     []string{"tester1@gmail.com", "tester3@ibm.com"},
				AddNotificationPayload: core.BoolPtr(true),
				ReplyToMail:            core.StringPtr("testerreply@gmail.com"),
				ReplyToName:            core.StringPtr("rester_reply"),
				FromName:               core.StringPtr("Test IBM email"),
			}
			name = core.StringPtr("subscription_email")
			description = core.StringPtr("Subscription for email")
			createSubscriptionOptions = &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          name,
				Description:   description,
				DestinationID: core.StringPtr(destinationID2),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreateAttributesEmailModel,
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.Attributes).ToNot(BeNil())
			Expect(subscription.Description).To(Equal(description))
			Expect(subscription.Name).To(Equal(name))
			subscriptionID2 = *subscription.ID

			Expect(subscriptionID2).ToNot(Equal(subscriptionID))

			createSubscriptionOptions = &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("FCM subscription"),
				Description:   core.StringPtr("Subscription for the FCM"),
				DestinationID: core.StringPtr(destinationID3),
				TopicID:       core.StringPtr(topicID3),
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID3 = string(*subscription.ID)
			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			// 409
			// 415
			// 500
			//
		})
	})

	Describe(`ListSubscriptions - List all Subscriptions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSubscriptions(listSubscriptionsOptions *ListSubscriptionsOptions)`, func() {

			listSubscriptionsOptions := &eventnotificationsv1.ListSubscriptionsOptions{
				InstanceID: core.StringPtr(instanceID),
				Offset:     core.Int64Ptr(int64(0)),
				Limit:      core.Int64Ptr(int64(1)),
				Search:     core.StringPtr(search),
			}

			subscriptionList, response, err := eventNotificationsService.ListSubscriptions(listSubscriptionsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscriptionList).ToNot(BeNil())

			subscriptionId1 := subscriptionList.Subscriptions[0].ID

			listSubscriptionsOptions = &eventnotificationsv1.ListSubscriptionsOptions{
				InstanceID: core.StringPtr(instanceID),
				Offset:     core.Int64Ptr(int64(1)),
				Limit:      core.Int64Ptr(int64(1)),
				Search:     core.StringPtr(search),
			}

			subscriptionList, response, err = eventNotificationsService.ListSubscriptions(listSubscriptionsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscriptionList).ToNot(BeNil())

			subscriptionId2 := subscriptionList.Subscriptions[0].ID

			Expect(subscriptionId2).ToNot(Equal(subscriptionId1))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 500
			//
		})
	})

	Describe(`GetSubscription - Get details of a Subscription`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSubscription(getSubscriptionOptions *GetSubscriptionOptions)`, func() {

			getSubscriptionOptions := &eventnotificationsv1.GetSubscriptionOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(subscriptionID),
			}

			subscription, response, err := eventNotificationsService.GetSubscription(getSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})

	Describe(`UpdateSubscription - Update details of a Subscription`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSubscription(updateSubscriptionOptions *UpdateSubscriptionOptions)`, func() {

			subscriptionUpdateAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesWebhookAttributes{
				SigningEnabled: core.BoolPtr(true),
			}

			name := core.StringPtr("Webhook_sub_updated")
			description := core.StringPtr("Update Webhook subscription")
			updateSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(subscriptionID),
				Name:        name,
				Description: description,
				Attributes:  subscriptionUpdateAttributesModel,
			}

			subscription, response, err := eventNotificationsService.UpdateSubscription(updateSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID)))
			Expect(subscription.Name).To(Equal(name))
			Expect(subscription.Description).To(Equal(description))

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			// 409
			// 415
			// 500
			//
		})
	})

	Describe(`SendNotifications - Send a notification`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SendNotifications(sendNotificationsOptions *SendNotificationsOptions)`, func() {

			notificationDevicesModel := &eventnotificationsv1.NotificationDevices{
				UserIds: []string{"userId"},
			}

			lightsModel := &eventnotificationsv1.Lights{
				LedArgb:  core.StringPtr("RED"),
				LedOnMs:  core.Int64Ptr(int64(0)),
				LedOffMs: core.StringPtr("20"),
			}

			styleModel := &eventnotificationsv1.Style{
				Type:  core.StringPtr("picture_notification"),
				Title: core.StringPtr("hello"),
				URL:   core.StringPtr("url.ibm.com"),
			}

			// FCM EN Data
			notificationFcmBodyMessageDataModel := &eventnotificationsv1.NotificationFcmBodyMessageData{
				Alert:               core.StringPtr("Alert message"),
				CollapseKey:         core.StringPtr("collapse_key"),
				InteractiveCategory: core.StringPtr("category_test"),
				Icon:                core.StringPtr("test.png"),
				DelayWhileIdle:      core.BoolPtr(true),
				Sync:                core.BoolPtr(true),
				Visibility:          core.StringPtr("0"),
				Redact:              core.StringPtr("redact test alert"),
				Payload:             make(map[string]interface{}),
				Priority:            core.StringPtr("MIN"),
				Sound:               core.StringPtr("newSound"),
				TimeToLive:          core.Int64Ptr(int64(0)),
				Lights:              lightsModel,
				AndroidTitle:        core.StringPtr("IBM test title"),
				GroupID:             core.StringPtr("Group_ID_1"),
				Style:               styleModel,
				Type:                core.StringPtr("DEFAULT"),
			}

			notificationFcmBodyModel := &eventnotificationsv1.NotificationFcmBodyMessageEnData{
				EnData: notificationFcmBodyMessageDataModel,
			}

			notificationApnsBodyMessageDataModel := &eventnotificationsv1.NotificationApnsBodyMessageData{
				Alert:                    core.StringPtr("Alert message"),
				Badge:                    core.Int64Ptr(int64(38)),
				InteractiveCategory:      core.StringPtr("InteractiveCategory"),
				IosActionKey:             core.StringPtr("IosActionKey"),
				Payload:                  map[string]interface{}{"testKey": "testValue"},
				Sound:                    core.StringPtr("sound.wav"),
				TitleLocKey:              core.StringPtr("TitleLocKey"),
				LocKey:                   core.StringPtr("LocKey"),
				LaunchImage:              core.StringPtr("image.png"),
				TitleLocArgs:             []string{"TitleLocArgs1", "TitleLocArgs2"},
				LocArgs:                  []string{"LocArgs1", "LocArgs2"},
				Title:                    core.StringPtr("Message Title"),
				Subtitle:                 core.StringPtr("Message SubTitle"),
				AttachmentURL:            core.StringPtr("https://testimage.sub.png"),
				Type:                     core.StringPtr("DEFAULT"),
				ApnsCollapseID:           core.StringPtr("ApnsCollapseID"),
				ApnsThreadID:             core.StringPtr("ApnsThreadID"),
				ApnsGroupSummaryArg:      core.StringPtr("ApnsGroupSummaryArg"),
				ApnsGroupSummaryArgCount: core.Int64Ptr(int64(38)),
			}

			notificationCreateMessageApnsBodyModel := &eventnotificationsv1.NotificationApnsBodyMessageEnData{
				EnData: notificationApnsBodyMessageDataModel,
			}

			notificationID := "1234-1234-sdfs-234"
			notificationSubject := "FCM_Subject"
			notificationSeverity := "MEDIUM"
			typeValue := "com.acme.offer:new"
			notificationsSouce := "1234-1234-sdfs-234:test"

			sendNotificationsOptions := &eventnotificationsv1.SendNotificationsOptions{
				InstanceID:      core.StringPtr(instanceID),
				Subject:         core.StringPtr(notificationSubject),
				Severity:        core.StringPtr(notificationSeverity),
				ID:              core.StringPtr(notificationID),
				Source:          core.StringPtr(notificationsSouce),
				EnSourceID:      core.StringPtr(sourceID),
				Type:            core.StringPtr(typeValue),
				Time:            CreateMockDateTime("2019-01-01T12:00:00.000Z"),
				Data:            make(map[string]interface{}),
				PushTo:          notificationDevicesModel,
				MessageFcmBody:  notificationFcmBodyModel,
				MessageApnsBody: notificationCreateMessageApnsBodyModel,
				Datacontenttype: core.StringPtr("application/json"),
				Specversion:     core.StringPtr("1.0"),
			}

			notificationResponse, response, err := eventNotificationsService.SendNotifications(sendNotificationsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(notificationResponse).ToNot(BeNil())

			apnsOptions := &eventnotificationsv1.NotificationApnsBody{}
			fcmOptions := &eventnotificationsv1.NotificationFcmBody{}

			apnsOptions.SetProperties(map[string]interface{}{
				"aps": map[string]interface{}{
					"alert": "Game Request",
					"badge": 5,
				},
			})
			fcmOptions.SetProperties(map[string]interface{}{
				"notification": map[string]interface{}{
					"title": "Portugal vs. Denmark",
					"body":  "great match!",
				},
			})

			apnsHeaders := map[string]interface{}{
				"apns-collapse-id": "123",
			}

			sendNotificationsOptions.MessageFcmBody = fcmOptions
			sendNotificationsOptions.MessageApnsBody = apnsOptions
			sendNotificationsOptions.MessageApnsHeaders = apnsHeaders
			notificationResponse, response, err = eventNotificationsService.SendNotifications(sendNotificationsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(notificationResponse).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 415
			// 500
			//
		})
	})

	Describe(`DeleteSubscription - Delete a Subscription`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSubscription(deleteSubscriptionOptions *DeleteSubscriptionOptions)`, func() {

			for _, ID := range []string{subscriptionID, subscriptionID2, subscriptionID3} {

				deleteSubscriptionOptions := &eventnotificationsv1.DeleteSubscriptionOptions{
					InstanceID: core.StringPtr(instanceID),
					ID:         core.StringPtr(ID),
				}

				response, err := eventNotificationsService.DeleteSubscription(deleteSubscriptionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(204))
			}

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})

	Describe(`DeleteTopic - Delete a Topic`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTopic(deleteTopicOptions *DeleteTopicOptions)`, func() {
			for _, ID := range []string{topicID, topicID2, topicID3} {
				deleteTopicOptions := &eventnotificationsv1.DeleteTopicOptions{
					InstanceID: core.StringPtr(instanceID),
					ID:         core.StringPtr(ID),
				}

				response, err := eventNotificationsService.DeleteTopic(deleteTopicOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(204))
			}

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})

	/*
		Describe(`DeleteTagsSubscription - Delete a Tag subcription`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`DeleteTagsSubscription(deleteTagsSubscriptionOptions *DeleteTagsSubscriptionOptions)`, func() {

				deleteTagsSubscriptionOptions := &eventnotificationsv1.DeleteTagsSubscriptionOptions{
					InstanceID: core.StringPtr("testString"),
					ID:         core.StringPtr("testString"),
					DeviceID:   core.StringPtr("testString"),
					TagName:    core.StringPtr("testString"),
				}

				response, err := eventNotificationsService.DeleteTagsSubscription(deleteTagsSubscriptionOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(204))

				//
				// The following status codes aren't covered by tests.
				// Please provide integration tests for these too.
				//
				// 401
				// 404
				// 500
				//
			})
		})
	*/

	Describe(`DeleteDestination - Delete a Destination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDestination(deleteDestinationOptions *DeleteDestinationOptions)`, func() {

			for _, ID := range []string{destinationID, destinationID3} {
				deleteDestinationOptions := &eventnotificationsv1.DeleteDestinationOptions{
					InstanceID: core.StringPtr(instanceID),
					ID:         core.StringPtr(ID),
				}

				response, err := eventNotificationsService.DeleteDestination(deleteDestinationOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(204))
			}

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})
})

//
// Utility functions are declared in the unit test file
//
