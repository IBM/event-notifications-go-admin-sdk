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
	"log"
	"os"
	"time"

	"github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
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
		safariCertificatePath     string
		search                    string = ""
		topicName                 string = "WebhookTopic"
		sourceID                  string
		topicID                   string
		topicID2                  string
		topicID3                  string
		destinationID             string
		destinationID1            string
		destinationID2            string
		destinationID3            string
		destinationID4            string
		destinationID5            string
		destinationID6            string
		destinationID7            string
		destinationID8            string
		destinationID9            string
		destinationID10           string
		destinationID11           string
		destinationID12           string
		destinationID13           string
		destinationID14           string
		destinationID15           string
		destinationID16           string
		destinationID17           string
		destinationID18           string
		subscriptionID            string
		subscriptionID1           string
		subscriptionID2           string
		subscriptionID3           string
		subscriptionID4           string
		subscriptionID5           string
		subscriptionID6           string
		subscriptionID7           string
		subscriptionID8           string
		subscriptionID9           string
		subscriptionID10          string
		subscriptionID11          string
		subscriptionID12          string
		subscriptionID13          string
		subscriptionID14          string
		subscriptionID15          string
		subscriptionID16          string
		subscriptionID17          string
		subscriptionID18          string
		fcmServerKey              string
		fcmSenderId               string
		integrationId             string
		sNowClientID              string
		sNowClientSecret          string
		sNowUserName              string
		sNowPassword              string
		sNowInstanceName          string
		fcmPrivateKey             string
		fcmProjectID              string
		fcmClientEmail            string
		codeEngineURL             string
		codeEngineProjectCRN      string
		huaweiClientSecret        string
		huaweiClientID            string
		cosBucketName             string
		cosInstanceID             string
		cosEndPoint               string
		templateInvitationID      string
		templateNotificationID    string
		slackURL                  string
		teamsURL                  string
		pagerDutyApiKey           string
		pagerDutyRoutingKey       string
		templateBody              string
		cosInstanceCRN            string
		cosIntegrationID          string
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

			fcmProjectID = config["FCM_PROJECT_ID"]
			if fcmProjectID == "" {
				Skip("Unable to load service fcmProjectID configuration property, skipping tests")
			}
			fmt.Printf("Service fcmProjectID: %s\n", fcmProjectID)

			fcmClientEmail = config["FCM_CLIENT_EMAIL"]
			if fcmClientEmail == "" {
				Skip("Unable to load service fcmclientEmail configuration property, skipping tests")
			}
			fmt.Printf("Service fcmClientEmail: %s\n", fcmClientEmail)

			fcmPrivateKey = config["FCM_PRIVATE_KEY"]
			if fcmPrivateKey == "" {
				Skip("Unable to load service fcmPrivateKey configuration property, skipping tests")
			}
			fmt.Printf("Service fcmPrivateKey: %s\n", fcmPrivateKey)

			safariCertificatePath = config["SAFARI_CERTIFICATE"]
			if safariCertificatePath == "" {
				Skip("Unable to load service safariCertificatePath configuration property, skipping tests")
			}
			fmt.Printf("Service safariCertificatePath: %s\n", safariCertificatePath)

			sNowClientID = config["SNOW_CLIENT_ID"]
			if sNowClientID == "" {
				Skip("Unable to load service sNowClientID configuration property, skipping tests")
			}
			fmt.Printf("Service sNowClientID: %s\n", sNowClientID)

			sNowClientSecret = config["SNOW_CLIENT_SECRET"]
			if sNowClientSecret == "" {
				Skip("Unable to load service sNowClientSecret configuration property, skipping tests")
			}
			fmt.Printf("Service sNowClientSecret: %s\n", sNowClientSecret)

			sNowUserName = config["SNOW_USER_NAME"]
			if sNowUserName == "" {
				Skip("Unable to load service sNowUserName configuration property, skipping tests")
			}
			fmt.Printf("Service sNowUserName: %s\n", sNowUserName)

			sNowPassword = config["SNOW_PASSWORD"]
			if sNowPassword == "" {
				Skip("Unable to load service sNowPassword configuration property, skipping tests")
			}
			fmt.Printf("Service sNowPassword: %s\n", sNowPassword)

			sNowInstanceName = config["SNOW_INSTANCE_NAME"]
			if sNowInstanceName == "" {
				Skip("Unable to load service sNowInstanceName configuration property, skipping tests")
			}
			fmt.Printf("Service sNowInstanceName: %s\n", sNowInstanceName)

			codeEngineURL = config["CODE_ENGINE_URL"]
			if codeEngineURL == "" {
				Skip("Unable to load code engine url configuration property, skipping tests")
			}
			fmt.Printf("code engine url: %s\n", codeEngineURL)

			huaweiClientID = config["HUAWEI_CLIENT_ID"]
			if huaweiClientID == "" {
				Skip("Unable to load huawei client ID configuration property, skipping tests")
			}
			fmt.Printf("huawei client ID: %s\n", huaweiClientID)

			huaweiClientSecret = config["HUAWEI_CLIENT_SECRET"]
			if huaweiClientSecret == "" {
				Skip("Unable to load huawei client secret configuration property, skipping tests")
			}
			fmt.Printf("huawei client secret: %s\n", huaweiClientSecret)

			cosBucketName = config["COS_BUCKET_NAME"]
			if cosBucketName == "" {
				Skip("Unable to load cos bucket name configuration property, skipping tests")
			}
			fmt.Printf("cos bucket name: %s\n", cosBucketName)

			cosInstanceID = config["COS_INSTANCE"]
			if cosInstanceID == "" {
				Skip("Unable to load cos instance ID configuration property, skipping tests")
			}
			fmt.Printf("cos Instance ID: %s\n", cosInstanceID)

			cosEndPoint = config["COS_ENDPOINT"]
			if cosEndPoint == "" {
				Skip("Unable to load cos end point configuration property, skipping tests")
			}
			fmt.Printf("cos end point: %s\n", cosEndPoint)

			slackURL = config["SLACK_URL"]
			if slackURL == "" {
				Skip("Unable to load slack configuration property, skipping tests")
			}
			fmt.Printf("slack url: %s\n", slackURL)

			teamsURL = config["MS_TEAMS_URL"]
			if teamsURL == "" {
				Skip("Unable to load msteams end point configuration property, skipping tests")
			}
			fmt.Printf("msteams url: %s\n", teamsURL)

			pagerDutyApiKey = config["PD_API_KEY"]
			if pagerDutyApiKey == "" {
				Skip("Unable to load  pagerDutyApiKey configuration property, skipping tests")
			}
			fmt.Printf(" pagerDutyApiKey: %s\n", pagerDutyApiKey)

			pagerDutyRoutingKey = config["PD_ROUTING_KEY"]
			if pagerDutyRoutingKey == "" {
				Skip("Unable to load pagerDutyRoutingKey configuration property, skipping tests")
			}
			fmt.Printf("pagerDutyRoutingKey: %s\n", pagerDutyRoutingKey)

			templateBody = config["TEMPLATE_BODY"]
			if templateBody == "" {
				Skip("Unable to load templateBody configuration property, skipping tests")
			}
			fmt.Printf("TemplateBody: %s\n", templateBody)

			cosInstanceCRN = config["COS_INSTANCE_CRN"]
			if cosInstanceCRN == "" {
				Skip("Unable to load cosInstanceCRN configuration property, skipping tests")
			}
			fmt.Printf("cosInstanceCRN: %s\n", cosInstanceCRN)

			codeEngineProjectCRN = config["CODE_ENGINE_PROJECT_CRN"]
			if codeEngineProjectCRN == "" {
				Skip("Unable to load CODE_ENGINE_PROJECT_CRN configuration property, skipping tests")
			}
			fmt.Printf("CODE_ENGINE_PROJECT_CRN: %s\n", codeEngineProjectCRN)

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

	Describe(`CreateIntegration - Create integration of an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateIntegration(createIntegrationOptions *CreateIntegrationOptions)`, func() {

			integrationMetadata := &eventnotificationsv1.IntegrationCreateMetadata{
				Endpoint:   core.StringPtr(cosEndPoint),
				CRN:        core.StringPtr(cosInstanceCRN),
				BucketName: core.StringPtr(cosBucketName),
			}

			createIntegrationsOptions := &eventnotificationsv1.CreateIntegrationOptions{
				InstanceID: core.StringPtr(instanceID),
				Type:       core.StringPtr("collect_failed_events"),
				Metadata:   integrationMetadata,
			}

			integrationCreateResponse, response, err := eventNotificationsService.CreateIntegration(createIntegrationsOptions)

			cosIntegrationID = string(*integrationCreateResponse.ID)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(integrationCreateResponse).ToNot(BeNil())
		})
	})

	Describe(`ListIntegrations - Lists all integrations for an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListIntegrations(listIntegrationsOptions *ListIntegrationsOptions)`, func() {

			listIntegrationsOptions := &eventnotificationsv1.ListIntegrationsOptions{
				InstanceID: core.StringPtr(instanceID),
				Limit:      core.Int64Ptr(int64(1)),
				Offset:     core.Int64Ptr(int64(0)),
				Search:     core.StringPtr(search),
			}

			integrationResponse, response, err := eventNotificationsService.ListIntegrations(listIntegrationsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(integrationResponse).ToNot(BeNil())

			integrationId = string(*integrationResponse.Integrations[0].ID)
		})
	})

	Describe(`GetIntegration - Get integration of an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetIntegration(getIntegrationOptions *GetIntegrationOptions)`, func() {

			listIntegrationsOptions := &eventnotificationsv1.GetIntegrationOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(integrationId),
			}

			integrationResponse, response, err := eventNotificationsService.GetIntegration(listIntegrationsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(integrationResponse).ToNot(BeNil())
		})
	})

	Describe(`UpdateIntegration - Update integration of an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateIntegration(updateIntegrationOptions *UpdateIntegrationOptions)`, func() {

			integrationMetadata := &eventnotificationsv1.IntegrationMetadata{
				Endpoint:   core.StringPtr(cosEndPoint),
				CRN:        core.StringPtr(cosInstanceCRN),
				BucketName: core.StringPtr(cosBucketName),
			}

			replaceIntegrationsOptions := &eventnotificationsv1.ReplaceIntegrationOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(cosIntegrationID),
				Type:       core.StringPtr("collect_failed_events"),
				Metadata:   integrationMetadata,
			}

			integrationResponse, response, err := eventNotificationsService.ReplaceIntegration(replaceIntegrationsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(integrationResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateSources - Create a new API Source`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSources(createSourcesOptions *CreateSourcesOptions)`, func() {

			createSourcesOptions := &eventnotificationsv1.CreateSourcesOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr("Event Notification Create Source Acme"),
				Description: core.StringPtr("This source is used for Acme Bank"),
				Enabled:     core.BoolPtr(false),
			}

			sourceResponse, response, err := eventNotificationsService.CreateSources(createSourcesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(sourceResponse).ToNot(BeNil())

			sourceID = *sourceResponse.ID

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

	Describe(`UpdateSource - Update details of a Source`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSource(updateSourceOptions *UpdateSourceOptions)`, func() {

			updateSourceOptions := &eventnotificationsv1.UpdateSourceOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(sourceID),
				Name:        core.StringPtr("Event Notification update Source Acme"),
				Description: core.StringPtr("This source is used for updated Acme Bank"),
				Enabled:     core.BoolPtr(true),
			}

			source, response, err := eventNotificationsService.UpdateSource(updateSourceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(source).ToNot(BeNil())

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

			topicUpdateSourcesItemModel := &eventnotificationsv1.SourcesItems{
				ID:    core.StringPtr(sourceID),
				Rules: []eventnotificationsv1.Rules{*rulesModel},
			}

			description := core.StringPtr("Topic for Webhook notifications")
			name := core.StringPtr(topicName)
			createTopicOptions := &eventnotificationsv1.CreateTopicOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        name,
				Description: description,
				Sources:     []eventnotificationsv1.SourcesItems{*topicUpdateSourcesItemModel},
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
				Sources:     []eventnotificationsv1.SourcesItems{*topicUpdateSourcesItemModel},
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

			topicUpdateSourcesItemModel = &eventnotificationsv1.SourcesItems{
				ID:    core.StringPtr(sourceID),
				Rules: []eventnotificationsv1.Rules{*rulesModel},
			}

			createTopicOptions = &eventnotificationsv1.CreateTopicOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr("FCM_topic"),
				Description: core.StringPtr("This topic is used for routing all compliance related notifications to the appropriate destinations"),
				Sources:     []eventnotificationsv1.SourcesItems{*topicUpdateSourcesItemModel},
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

			topicUpdateSourcesItemModel := &eventnotificationsv1.SourcesItems{
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
				Sources:     []eventnotificationsv1.SourcesItems{*topicUpdateSourcesItemModel},
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

			destinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfWebhookDestinationConfig{
				URL:  core.StringPtr("https://gcm.com"),
				Verb: core.StringPtr("get"),
				CustomHeaders: map[string]string{
					"gcm_apikey": "api_key_value",
				},
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

			createFCMDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"FCM_destination",
				eventnotificationsv1.CreateDestinationOptionsTypePushAndroidConst,
			)

			destinationConfigParamsFCMModel := &eventnotificationsv1.DestinationConfigOneOfFcmDestinationConfig{
				ServerKey: core.StringPtr(fcmServerKey),
				SenderID:  core.StringPtr(fcmSenderId),
			}

			fcmDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsFCMModel,
			}

			createFCMDestinationOptions.SetConfig(fcmDestinationConfigModel)

			destinationResponse, response, err = eventNotificationsService.CreateDestination(createFCMDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID3 = *destinationResponse.ID

			createSlackDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"Slack_destination",
				eventnotificationsv1.CreateDestinationOptionsTypeSlackConst,
			)

			destinationConfigParamsSlackModel := &eventnotificationsv1.DestinationConfigOneOfSlackDestinationConfig{
				URL: core.StringPtr(slackURL),
			}

			slackDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsSlackModel,
			}

			createSlackDestinationOptions.SetConfig(slackDestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(createSlackDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID4 = *destinationResponse.ID

			createSafariDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"Safari_destination",
				eventnotificationsv1.CreateDestinationOptionsTypePushSafariConst)

			certificatefile, err := os.Open(safariCertificatePath)
			if err != nil {
				panic(err)
			}
			createSafariDestinationOptions.Certificate = certificatefile

			destinationConfigParamsSafariModel := &eventnotificationsv1.DestinationConfigOneOfSafariDestinationConfig{
				CertType:        core.StringPtr("p12"),
				Password:        core.StringPtr("safari"),
				WebsiteURL:      core.StringPtr("https://ensafaripush.mybluemix.net"),
				WebsiteName:     core.StringPtr("NodeJS Starter Application"),
				URLFormatString: core.StringPtr("https://ensafaripush.mybluemix.net/%@/?flight=%@"),
				WebsitePushID:   core.StringPtr("web.net.mybluemix.ensafaripush"),
			}

			safariDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsSafariModel,
			}

			createSafariDestinationOptions.SetConfig(safariDestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(createSafariDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID5 = *destinationResponse.ID

			createMSTeamsDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"MSTeams_destination",
				eventnotificationsv1.CreateDestinationOptionsTypeMsteamsConst,
			)

			destinationConfigParamsMSTeaMSModel := &eventnotificationsv1.DestinationConfigOneOfMsTeamsDestinationConfig{
				URL: core.StringPtr(teamsURL),
			}

			msTeamsDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsMSTeaMSModel,
			}

			createMSTeamsDestinationOptions.SetConfig(msTeamsDestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(createMSTeamsDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID6 = *destinationResponse.ID

			cfCreateDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"Cloud_Functions_destination",
				eventnotificationsv1.CreateDestinationOptionsTypeIbmcfConst,
			)

			destinationConfigParamsCloudFunctionsModel := &eventnotificationsv1.DestinationConfigOneOfIBMCloudFunctionsDestinationConfig{
				URL:    core.StringPtr("https://www.ibmcfendpoint.com/"),
				APIKey: core.StringPtr("sdslknsdlfnlsejifw900"),
			}

			cfdestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsCloudFunctionsModel,
			}

			cfCreateDestinationOptions.SetConfig(cfdestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(cfCreateDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID7 = *destinationResponse.ID

			chromeCreateDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"Chrome_destination",
				eventnotificationsv1.CreateDestinationOptionsTypePushChromeConst,
			)

			destinationConfigParamsChromeModel := &eventnotificationsv1.DestinationConfigOneOfChromeDestinationConfig{
				APIKey:     core.StringPtr("sdslknsdlfnlsejifw900"),
				WebsiteURL: core.StringPtr("https://cloud.ibm.com"),
			}

			chromeDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsChromeModel,
			}

			chromeCreateDestinationOptions.SetConfig(chromeDestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(chromeCreateDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID8 = *destinationResponse.ID

			fireCreateDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"Firefox_destination",
				eventnotificationsv1.CreateDestinationOptionsTypePushFirefoxConst,
			)

			destinationConfigParamsfireModel := &eventnotificationsv1.DestinationConfigOneOfFirefoxDestinationConfig{
				WebsiteURL: core.StringPtr("https://cloud.ibm.com"),
			}

			fireDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsfireModel,
			}

			fireCreateDestinationOptions.SetConfig(fireDestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(fireCreateDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID9 = *destinationResponse.ID

			pagerDutyCreateDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"PagerDuty_destination",
				eventnotificationsv1.CreateDestinationOptionsTypePagerdutyConst,
			)

			destinationConfigParamsPDModel := &eventnotificationsv1.DestinationConfigOneOfPagerDutyDestinationConfig{
				APIKey:     core.StringPtr(pagerDutyApiKey),
				RoutingKey: core.StringPtr(pagerDutyRoutingKey),
			}

			pagerDutyDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsPDModel,
			}

			pagerDutyCreateDestinationOptions.SetConfig(pagerDutyDestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(pagerDutyCreateDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID10 = *destinationResponse.ID

			serviceNowCreateDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"servicenow_destination",
				eventnotificationsv1.CreateDestinationOptionsTypeServicenowConst,
			)

			destinationConfigParamsServiceNowModel := &eventnotificationsv1.DestinationConfigOneOfServiceNowDestinationConfig{
				ClientID:     core.StringPtr(sNowClientID),
				ClientSecret: core.StringPtr(sNowClientSecret),
				Username:     core.StringPtr(sNowUserName),
				Password:     core.StringPtr(sNowPassword),
				InstanceName: core.StringPtr(sNowInstanceName),
			}

			serviceNowDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsServiceNowModel,
			}

			serviceNowCreateDestinationOptions.SetConfig(serviceNowDestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(serviceNowCreateDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID11 = *destinationResponse.ID

			createFCMDestinationOptions = eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"FCM_destination_V1",
				eventnotificationsv1.CreateDestinationOptionsTypePushAndroidConst,
			)

			destinationConfigParamsFCMModel = &eventnotificationsv1.DestinationConfigOneOfFcmDestinationConfig{
				ProjectID:   core.StringPtr(fcmProjectID),
				PrivateKey:  core.StringPtr(fcmPrivateKey),
				ClientEmail: core.StringPtr(fcmClientEmail),
			}

			fcmDestinationConfigModel = &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsFCMModel,
			}

			createFCMDestinationOptions.SetConfig(fcmDestinationConfigModel)

			destinationResponse, response, err = eventNotificationsService.CreateDestination(createFCMDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID12 = *destinationResponse.ID

			destinationConfigCEParamsModel := &eventnotificationsv1.DestinationConfigOneOfCodeEngineDestinationConfig{
				URL:  core.StringPtr(codeEngineURL),
				Verb: core.StringPtr("get"),
				Type: core.StringPtr("application"),
				CustomHeaders: map[string]string{
					"authorization": "api_key_value",
				},
				SensitiveHeaders: []string{"authorization"},
			}

			destinationConfigCEModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigCEParamsModel,
			}

			ceName := "codeengine_destination"
			ceTypeVal := "ibmce"
			ceDescription := "codeengine Destination"
			createCEDestinationOptions := &eventnotificationsv1.CreateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(ceName),
				Type:        core.StringPtr(ceTypeVal),
				Description: core.StringPtr(ceDescription),
				Config:      destinationConfigCEModel,
			}

			destinationCEResponse, response, err := eventNotificationsService.CreateDestination(createCEDestinationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationCEResponse).ToNot(BeNil())
			Expect(destinationCEResponse.Name).To(Equal(core.StringPtr(ceName)))
			Expect(destinationCEResponse.Type).To(Equal(core.StringPtr(ceTypeVal)))
			Expect(destinationCEResponse.Description).To(Equal(core.StringPtr(ceDescription)))

			destinationID13 = *destinationCEResponse.ID

			cosDestinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfIBMCloudObjectStorageDestinationConfig{
				BucketName: core.StringPtr(cosBucketName),
				InstanceID: core.StringPtr(cosInstanceID),
				Endpoint:   core.StringPtr(cosEndPoint),
			}

			cosDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: cosDestinationConfigParamsModel,
			}

			cosName := "cos_destination"
			costypeVal := eventnotificationsv1.CreateDestinationOptionsTypeIbmcosConst
			cosDescription := "cos Destination"
			cosCreateDestinationOptions := &eventnotificationsv1.CreateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(cosName),
				Type:        core.StringPtr(costypeVal),
				Description: core.StringPtr(cosDescription),
				Config:      cosDestinationConfigModel,
			}

			destinationResponse, response, err = eventNotificationsService.CreateDestination(cosCreateDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID14 = *destinationResponse.ID

			huaweiDestinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfHuaweiDestinationConfig{
				ClientID:     core.StringPtr(huaweiClientID),
				ClientSecret: core.StringPtr(huaweiClientSecret),
				PreProd:      core.BoolPtr(false),
			}

			huaweiDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: huaweiDestinationConfigParamsModel,
			}

			huaweiName := "huawei_destination"
			huaweitypeVal := eventnotificationsv1.CreateDestinationOptionsTypePushHuaweiConst
			huaweiDescription := "huawei Destination"
			huaweiCreateDestinationOptions := &eventnotificationsv1.CreateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(huaweiName),
				Type:        core.StringPtr(huaweitypeVal),
				Description: core.StringPtr(huaweiDescription),
				Config:      huaweiDestinationConfigModel,
			}

			destinationResponse, response, err = eventNotificationsService.CreateDestination(huaweiCreateDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID15 = *destinationResponse.ID

			customDestinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig{
				Domain: core.StringPtr("test.event-notifications.test.cloud.ibm.com"),
			}

			customDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: customDestinationConfigParamsModel,
			}

			customName := "custom_email_destination"
			customtypeVal := eventnotificationsv1.CreateDestinationOptionsTypeSMTPCustomConst
			customDescription := "custom Destination"
			customCreateDestinationOptions := &eventnotificationsv1.CreateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(customName),
				Type:        core.StringPtr(customtypeVal),
				Description: core.StringPtr(customDescription),
				Config:      customDestinationConfigModel,
			}

			destinationResponse, response, err = eventNotificationsService.CreateDestination(customCreateDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID16 = *destinationResponse.ID

			customSMSName := "custom_sms_destination"
			customSMSTypeVal := eventnotificationsv1.CreateDestinationOptionsTypeSmsCustomConst
			customSMSDescription := "custom sms Destination"
			customSMSCreateDestinationOptions := &eventnotificationsv1.CreateDestinationOptions{
				InstanceID:          core.StringPtr(instanceID),
				Name:                core.StringPtr(customSMSName),
				Type:                core.StringPtr(customSMSTypeVal),
				Description:         core.StringPtr(customSMSDescription),
				CollectFailedEvents: core.BoolPtr(false),
			}

			destinationResponse, response, err = eventNotificationsService.CreateDestination(customSMSCreateDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID17 = *destinationResponse.ID

			ceName = "codeengine_job_destination"
			ceDescription = "codeengine job Destination"
			destinationConfigCEJobParamsModel := &eventnotificationsv1.DestinationConfigOneOfCodeEngineDestinationConfig{
				ProjectCRN: core.StringPtr(codeEngineProjectCRN),
				JobName:    core.StringPtr("custom-job"),
				Type:       core.StringPtr("job"),
			}

			destinationConfigCEJobsModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigCEJobParamsModel,
			}

			createCEJobDestinationOptions := &eventnotificationsv1.CreateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(ceName),
				Type:        core.StringPtr(ceTypeVal),
				Description: core.StringPtr(ceDescription),
				Config:      destinationConfigCEJobsModel,
			}

			destinationCEJobResponse, response, err := eventNotificationsService.CreateDestination(createCEJobDestinationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationCEJobResponse).ToNot(BeNil())
			Expect(destinationCEJobResponse.Name).To(Equal(core.StringPtr(ceName)))
			Expect(destinationCEJobResponse.Type).To(Equal(core.StringPtr(ceTypeVal)))
			Expect(destinationCEJobResponse.Description).To(Equal(core.StringPtr(ceDescription)))

			destinationID18 = *destinationCEJobResponse.ID

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

	Describe(`TestDestination - Test Destination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`TestDestination(TestDestinationOptions *testDestinationOptions)`, func() {

			for _, ID := range []string{destinationID4, destinationID6, destinationID10, destinationID14} {
				testDestinationOptions := &eventnotificationsv1.TestDestinationOptions{
					InstanceID: core.StringPtr(instanceID),
					ID:         core.StringPtr(ID),
				}

				_, response, err := eventNotificationsService.TestDestination(testDestinationOptions)
				if err != nil {
					panic(err)
				}
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
			}
		})
	})

	Describe(`CreateTemplate - Create a new template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTemplate(CreateTemplateOptions *createTemplateOptions)`, func() {

			name := "template invitation"
			description := "template invitation description"
			templateTypeInvitation := "smtp_custom.invitation"
			templateTypeNotification := "smtp_custom.notification"

			templConfig := &eventnotificationsv1.TemplateConfig{
				Body:    core.StringPtr(templateBody),
				Subject: core.StringPtr("Hi this is invitation for invitation message"),
			}

			createTemplateOptions := &eventnotificationsv1.CreateTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeInvitation),
				Description: core.StringPtr(description),
				Params:      templConfig,
			}

			templateResponse, response, err := eventNotificationsService.CreateTemplate(createTemplateOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(templateResponse).ToNot(BeNil())
			Expect(templateResponse.Name).To(Equal(core.StringPtr(name)))
			Expect(templateResponse.Description).To(Equal(core.StringPtr(description)))

			templateInvitationID = *templateResponse.ID

			name = "template notification"
			description = "template notification description"

			templConfig = &eventnotificationsv1.TemplateConfig{
				Body:    core.StringPtr(templateBody),
				Subject: core.StringPtr("Hi this is template for notification"),
			}

			createTemplateOptions = &eventnotificationsv1.CreateTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeNotification),
				Description: core.StringPtr(description),
				Params:      templConfig,
			}

			templateResponse, response, err = eventNotificationsService.CreateTemplate(createTemplateOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(templateResponse).ToNot(BeNil())
			Expect(templateResponse.Name).To(Equal(core.StringPtr(name)))
			Expect(templateResponse.Description).To(Equal(core.StringPtr(description)))

			templateNotificationID = *templateResponse.ID
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
				Limit:      core.Int64Ptr(int64(20)),
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
					if destinationID1 != "" {
						break
					}
				}
				if *ID.Type == "sms_ibm" {
					destinationID1 = *ID.ID
					if destinationID2 != "" {
						break
					}
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

	Describe(`GetTemplate - Get details of a Template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDestination(getTemplateOptions *GetTemplateOptions)`, func() {

			getTemplateOptions := &eventnotificationsv1.GetTemplateOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(templateInvitationID),
			}

			template, response, err := eventNotificationsService.GetTemplate(getTemplateOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(template).ToNot(BeNil())
		})
	})

	Describe(`UpdateDestination - Update details of a Destination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDestination(updateDestinationOptions *UpdateDestinationOptions)`, func() {

			//webhook
			destinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfWebhookDestinationConfig{
				URL:  core.StringPtr("https://cloud.ibm.com/nhwebhook/sendwebhook"),
				Verb: core.StringPtr("post"),
				CustomHeaders: map[string]string{
					"authorization": "authorization key",
				},
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

			//FCM
			destinationConfigParamsFCMModel := &eventnotificationsv1.DestinationConfigOneOfFcmDestinationConfig{
				ServerKey: core.StringPtr(fcmServerKey),
				SenderID:  core.StringPtr(fcmSenderId),
			}

			fcmDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsFCMModel,
			}

			fcmName := "fcm_destination_update"
			fcmDescription := "This destination is for FCM"
			fcmUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID3),
				Name:        core.StringPtr(fcmName),
				Description: core.StringPtr(fcmDescription),
				Config:      fcmDestinationConfigModel,
			}

			destination, response, err = eventNotificationsService.UpdateDestination(fcmUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())
			Expect(destination.ID).To(Equal(core.StringPtr(destinationID3)))
			Expect(destination.Name).To(Equal(core.StringPtr(fcmName)))
			Expect(destination.Description).To(Equal(core.StringPtr(fcmDescription)))

			//slack
			destinationConfigParamsSlackModel := &eventnotificationsv1.DestinationConfigOneOfSlackDestinationConfig{
				URL: core.StringPtr(slackURL),
			}

			slackDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsSlackModel,
			}

			slackName := "slack_destination_update"
			slackDescription := "This destination is for slack"
			slackUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID4),
				Name:        core.StringPtr(slackName),
				Description: core.StringPtr(slackDescription),
				Config:      slackDestinationConfigModel,
			}

			destination, response, err = eventNotificationsService.UpdateDestination(slackUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())
			Expect(destination.ID).To(Equal(core.StringPtr(destinationID4)))
			Expect(destination.Name).To(Equal(core.StringPtr(slackName)))
			Expect(destination.Description).To(Equal(core.StringPtr(slackDescription)))

			//safari
			safaridestinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfSafariDestinationConfig{
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

			safariName := "Safari_dest"
			safaridescription := "This destination is for Safari"
			safariupdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID5),
				Name:        core.StringPtr(safariName),
				Description: core.StringPtr(safaridescription),
				Config:      safaridestinationConfigModel,
			}

			certificatefile, err := os.Open(safariCertificatePath)
			if err != nil {
				panic(err)
			}

			safariupdateDestinationOptions.Certificate = certificatefile

			safaridestination, safariresponse, err := eventNotificationsService.UpdateDestination(safariupdateDestinationOptions)

			Expect(err).To(BeNil())
			Expect(safariresponse.StatusCode).To(Equal(200))
			Expect(safaridestination).ToNot(BeNil())
			Expect(safaridestination.ID).To(Equal(core.StringPtr(destinationID5)))
			Expect(safaridestination.Name).To(Equal(core.StringPtr(safariName)))
			Expect(safaridestination.Description).To(Equal(core.StringPtr(safaridescription)))

			//MSTeams

			destinationConfigParamsMSTeaMSModel := &eventnotificationsv1.DestinationConfigOneOfMsTeamsDestinationConfig{
				URL: core.StringPtr(teamsURL),
			}

			msTeamsDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsMSTeaMSModel,
			}

			teamsName := "Msteams_dest"
			teamsDescription := "This destination is for MSTeams"
			msTeamsupdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID6),
				Name:        core.StringPtr(teamsName),
				Description: core.StringPtr(teamsDescription),
				Config:      msTeamsDestinationConfigModel,
			}

			destination, response, err = eventNotificationsService.UpdateDestination(msTeamsupdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())
			Expect(destination.ID).To(Equal(core.StringPtr(destinationID6)))
			Expect(destination.Name).To(Equal(core.StringPtr(teamsName)))
			Expect(destination.Description).To(Equal(core.StringPtr(teamsDescription)))

			//cloud functins
			destinationConfigParamsCloudFunctionskModel := &eventnotificationsv1.DestinationConfigOneOfIBMCloudFunctionsDestinationConfig{
				URL:    core.StringPtr("https://www.ibmcfendpoint.com/"),
				APIKey: core.StringPtr("sdslknsdlfnlsejifw900"),
			}

			cfdestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsCloudFunctionskModel,
			}

			cfName := "cf_dest"
			cfDescription := "This destination is for cloud functions"
			cfupdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID7),
				Name:        core.StringPtr(cfName),
				Description: core.StringPtr(cfDescription),
				Config:      cfdestinationConfigModel,
			}

			cfdestination, cfresponse, err := eventNotificationsService.UpdateDestination(cfupdateDestinationOptions)

			Expect(err).To(BeNil())
			Expect(cfresponse.StatusCode).To(Equal(200))
			Expect(cfdestination).ToNot(BeNil())
			Expect(cfdestination.ID).To(Equal(core.StringPtr(destinationID7)))
			Expect(cfdestination.Name).To(Equal(core.StringPtr(cfName)))
			Expect(cfdestination.Description).To(Equal(core.StringPtr(cfDescription)))

			//Chrome
			destinationConfigParamsChromeModel := &eventnotificationsv1.DestinationConfigOneOfChromeDestinationConfig{
				APIKey:     core.StringPtr("sdslknsdlfnlsejifw900"),
				WebsiteURL: core.StringPtr("https://cloud.ibm.com"),
			}

			chromeDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsChromeModel,
			}

			chromeName := "chrome_dest"
			chromeDescription := "This destination is for chrome"
			chromeupdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID8),
				Name:        core.StringPtr(chromeName),
				Description: core.StringPtr(chromeDescription),
				Config:      chromeDestinationConfigModel,
			}

			destination, response, err = eventNotificationsService.UpdateDestination(chromeupdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())
			Expect(destination.ID).To(Equal(core.StringPtr(destinationID8)))
			Expect(destination.Name).To(Equal(core.StringPtr(chromeName)))
			Expect(destination.Description).To(Equal(core.StringPtr(chromeDescription)))

			//Firefox
			destinationConfigParamsfireModel := &eventnotificationsv1.DestinationConfigOneOfFirefoxDestinationConfig{
				WebsiteURL: core.StringPtr("https://cloud.ibm.com"),
			}

			fireDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsfireModel,
			}

			fireName := "fire_dest"
			fireDescription := "This destination is for firefox"
			fireUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID9),
				Name:        core.StringPtr(fireName),
				Description: core.StringPtr(fireDescription),
				Config:      fireDestinationConfigModel,
			}

			destination, response, err = eventNotificationsService.UpdateDestination(fireUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())
			Expect(destination.ID).To(Equal(core.StringPtr(destinationID9)))
			Expect(destination.Name).To(Equal(core.StringPtr(fireName)))
			Expect(destination.Description).To(Equal(core.StringPtr(fireDescription)))

			destinationConfigParamsPDModel := &eventnotificationsv1.DestinationConfigOneOfPagerDutyDestinationConfig{
				APIKey:     core.StringPtr(pagerDutyApiKey),
				RoutingKey: core.StringPtr(pagerDutyRoutingKey),
			}

			pagerDutyDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsPDModel,
			}

			pdName := "Pagerduty_dest_update"
			pdDescription := "This destination update is for Pagerduty"
			pagerDutyUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID10),
				Name:        core.StringPtr(pdName),
				Description: core.StringPtr(pdDescription),
				Config:      pagerDutyDestinationConfigModel,
			}

			destination, response, err = eventNotificationsService.UpdateDestination(pagerDutyUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())
			Expect(destination.ID).To(Equal(core.StringPtr(destinationID10)))
			Expect(destination.Name).To(Equal(core.StringPtr(pdName)))
			Expect(destination.Description).To(Equal(core.StringPtr(pdDescription)))

			destinationConfigParamsServiceNowModel := &eventnotificationsv1.DestinationConfigOneOfServiceNowDestinationConfig{
				ClientID:     core.StringPtr(sNowClientID),
				ClientSecret: core.StringPtr(sNowClientSecret),
				Username:     core.StringPtr(sNowUserName),
				Password:     core.StringPtr(sNowPassword),
				InstanceName: core.StringPtr(sNowInstanceName),
			}

			serviceNowDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsServiceNowModel,
			}

			serviceNowName := "ServiceNow_dest_update"
			serviceNowDescription := "This destination update is for ServiceNow"
			serviceNowUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID11),
				Name:        core.StringPtr(serviceNowName),
				Description: core.StringPtr(serviceNowDescription),
				Config:      serviceNowDestinationConfigModel,
			}

			destination, response, err = eventNotificationsService.UpdateDestination(serviceNowUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())
			Expect(destination.ID).To(Equal(core.StringPtr(destinationID11)))
			Expect(destination.Name).To(Equal(core.StringPtr(serviceNowName)))
			Expect(destination.Description).To(Equal(core.StringPtr(serviceNowDescription)))

			destinationConfigParamsFCMModel = &eventnotificationsv1.DestinationConfigOneOfFcmDestinationConfig{
				ProjectID:   core.StringPtr(fcmProjectID),
				PrivateKey:  core.StringPtr(fcmPrivateKey),
				ClientEmail: core.StringPtr(fcmClientEmail),
			}

			fcmDestinationConfigModel = &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsFCMModel,
			}

			fcmName = "fcm_V1_destination_update"
			fcmDescription = "This destination is for FCM V1"
			fcmUpdateDestinationOptions = &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID12),
				Name:        core.StringPtr(fcmName),
				Description: core.StringPtr(fcmDescription),
				Config:      fcmDestinationConfigModel,
			}

			destination, response, err = eventNotificationsService.UpdateDestination(fcmUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())
			Expect(destination.ID).To(Equal(core.StringPtr(destinationID12)))
			Expect(destination.Name).To(Equal(core.StringPtr(fcmName)))
			Expect(destination.Description).To(Equal(core.StringPtr(fcmDescription)))

			destinationConfigCEParamsModel := &eventnotificationsv1.DestinationConfigOneOfCodeEngineDestinationConfig{
				URL:  core.StringPtr(codeEngineURL),
				Verb: core.StringPtr("get"),
				Type: core.StringPtr("application"),
				CustomHeaders: map[string]string{
					"authorization": "authorization key",
				},
				SensitiveHeaders: []string{"authorization"},
			}

			destinationConfigCEModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigCEParamsModel,
			}

			ceName := "code engine updated"
			ceDescription := "This destination is updated for creating code engine notifications"
			updateCEDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID13),
				Name:        core.StringPtr(ceName),
				Description: core.StringPtr(ceDescription),
				Config:      destinationConfigCEModel,
			}

			ceDestination, response, err := eventNotificationsService.UpdateDestination(updateCEDestinationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ceDestination).ToNot(BeNil())
			Expect(ceDestination.ID).To(Equal(core.StringPtr(destinationID13)))
			Expect(ceDestination.Name).To(Equal(core.StringPtr(ceName)))
			Expect(ceDestination.Description).To(Equal(core.StringPtr(ceDescription)))

			cosDestinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfIBMCloudObjectStorageDestinationConfig{
				BucketName: core.StringPtr(cosBucketName),
				InstanceID: core.StringPtr(cosInstanceID),
				Endpoint:   core.StringPtr(cosEndPoint),
			}

			cosDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: cosDestinationConfigParamsModel,
			}

			cosName := "cos_destination update"
			cosDescription := "cos Destination updated"
			cosUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(cosName),
				ID:          core.StringPtr(destinationID14),
				Description: core.StringPtr(cosDescription),
				Config:      cosDestinationConfigModel,
			}

			cosDestination, response, err := eventNotificationsService.UpdateDestination(cosUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cosDestination).ToNot(BeNil())
			Expect(cosDestination.ID).To(Equal(core.StringPtr(destinationID14)))
			Expect(cosDestination.Name).To(Equal(core.StringPtr(cosName)))
			Expect(cosDestination.Description).To(Equal(core.StringPtr(cosDescription)))

			huaweiDestinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfHuaweiDestinationConfig{
				ClientID:     core.StringPtr(huaweiClientID),
				ClientSecret: core.StringPtr(huaweiClientSecret),
				PreProd:      core.BoolPtr(false),
			}

			huaweiDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: huaweiDestinationConfigParamsModel,
			}

			huaweiName := "huawei_destination update"
			huaweiDescription := "huawei Destination updated"
			huaweiUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(huaweiName),
				ID:          core.StringPtr(destinationID15),
				Description: core.StringPtr(huaweiDescription),
				Config:      huaweiDestinationConfigModel,
			}

			huaweiDestination, response, err := eventNotificationsService.UpdateDestination(huaweiUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(huaweiDestination).ToNot(BeNil())
			Expect(huaweiDestination.ID).To(Equal(core.StringPtr(destinationID15)))
			Expect(huaweiDestination.Name).To(Equal(core.StringPtr(huaweiName)))
			Expect(huaweiDestination.Description).To(Equal(core.StringPtr(huaweiDescription)))

			customDestinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig{
				Domain: core.StringPtr("test.event-notifications.test.cloud.ibm.com"),
			}

			customDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: customDestinationConfigParamsModel,
			}

			customName := "custom_email_destination update"
			customDescription := "custom email Destination updated"
			customUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(customName),
				ID:          core.StringPtr(destinationID16),
				Description: core.StringPtr(customDescription),
				Config:      customDestinationConfigModel,
			}

			customDestination, response, err := eventNotificationsService.UpdateDestination(customUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customDestination).ToNot(BeNil())
			Expect(customDestination.ID).To(Equal(core.StringPtr(destinationID16)))
			Expect(customDestination.Name).To(Equal(core.StringPtr(customName)))
			Expect(customDestination.Description).To(Equal(core.StringPtr(customDescription)))

			customSpfUpdateDestinationOptions := &eventnotificationsv1.UpdateVerifyDestinationOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(destinationID16),
				Type:       core.StringPtr("spf"),
			}

			spfResult, response, err := eventNotificationsService.UpdateVerifyDestination(customSpfUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(spfResult).ToNot(BeNil())

			customDkimUpdateDestinationOptions := &eventnotificationsv1.UpdateVerifyDestinationOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(destinationID16),
				Type:       core.StringPtr("dkim"),
			}

			dkimResult, response, err := eventNotificationsService.UpdateVerifyDestination(customDkimUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dkimResult).ToNot(BeNil())

			customSMSName := "custom_sms_destination update"
			customSMSDescription := "custom sms Destination updated"
			customSMSUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:          core.StringPtr(instanceID),
				Name:                core.StringPtr(customSMSName),
				ID:                  core.StringPtr(destinationID17),
				Description:         core.StringPtr(customSMSDescription),
				CollectFailedEvents: core.BoolPtr(false),
			}

			customSMSDestination, response, err := eventNotificationsService.UpdateDestination(customSMSUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customSMSDestination).ToNot(BeNil())
			Expect(customSMSDestination.ID).To(Equal(core.StringPtr(destinationID17)))
			Expect(customSMSDestination.Name).To(Equal(core.StringPtr(customSMSName)))
			Expect(customSMSDestination.Description).To(Equal(core.StringPtr(customSMSDescription)))

			destinationConfigCEJobParamsModel := &eventnotificationsv1.DestinationConfigOneOfCodeEngineDestinationConfig{
				ProjectCRN: core.StringPtr(codeEngineProjectCRN),
				JobName:    core.StringPtr("custom-job"),
				Type:       core.StringPtr("job"),
			}

			destinationConfigCEJobModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigCEJobParamsModel,
			}

			ceName = "code engine job updated"
			ceDescription = "This destination is updated for creating code engine job"
			updateCEJobDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID18),
				Name:        core.StringPtr(ceName),
				Description: core.StringPtr(ceDescription),
				Config:      destinationConfigCEJobModel,
			}

			ceJobDestination, response, err := eventNotificationsService.UpdateDestination(updateCEJobDestinationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ceJobDestination).ToNot(BeNil())
			Expect(ceJobDestination.ID).To(Equal(core.StringPtr(destinationID18)))
			Expect(ceJobDestination.Name).To(Equal(core.StringPtr(ceName)))
			Expect(ceJobDestination.Description).To(Equal(core.StringPtr(ceDescription)))

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

	Describe(`UpdateTemplate - Update a template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTemplate(UpdateTemplateOptions *updateTemplateOptions)`, func() {

			name := "template invitation"
			description := "template invitation description"
			templateTypeInvitation := "smtp_custom.invitation"
			templateTypeNotification := "smtp_custom.notification"

			templateConfig := &eventnotificationsv1.TemplateConfig{
				Body:    core.StringPtr(templateBody),
				Subject: core.StringPtr("Hi this is invitation for invitation message"),
			}

			replaceTemplateOptions := &eventnotificationsv1.ReplaceTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(templateInvitationID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeInvitation),
				Description: core.StringPtr(description),
				Params:      templateConfig,
			}

			templateResponse, response, err := eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResponse).ToNot(BeNil())
			Expect(templateResponse.ID).To(Equal(core.StringPtr(templateInvitationID)))
			Expect(templateResponse.Name).To(Equal(core.StringPtr(name)))
			Expect(templateResponse.Description).To(Equal(core.StringPtr(description)))

			name = "template notification"
			description = "template notification description"

			templateConfig = &eventnotificationsv1.TemplateConfig{
				Body:    core.StringPtr(templateBody),
				Subject: core.StringPtr("Hi this is template for notification"),
			}

			replaceTemplateOptions = &eventnotificationsv1.ReplaceTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(templateNotificationID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeNotification),
				Description: core.StringPtr(description),
				Params:      templateConfig,
			}

			templateResponse, response, err = eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResponse).ToNot(BeNil())
			Expect(templateResponse.ID).To(Equal(core.StringPtr(templateNotificationID)))
			Expect(templateResponse.Name).To(Equal(core.StringPtr(name)))
			Expect(templateResponse.Description).To(Equal(core.StringPtr(description)))
		})
	})

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

			subscriptionCreateAttributesSMSModel := &eventnotificationsv1.SubscriptionCreateAttributesSmsAttributes{
				Invited: []string{"+12064563059", "+12267054625"},
			}
			smsName := core.StringPtr("subscription_sms")
			smsDescription := core.StringPtr("Subscription for sms")
			createSMSSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          smsName,
				Description:   smsDescription,
				DestinationID: core.StringPtr(destinationID1),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreateAttributesSMSModel,
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createSMSSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.Attributes).ToNot(BeNil())
			Expect(subscription.Description).To(Equal(smsDescription))
			Expect(subscription.Name).To(Equal(smsName))
			subscriptionID1 = *subscription.ID

			subscriptionCreateAttributesEmailModel := &eventnotificationsv1.SubscriptionCreateAttributesEmailAttributes{
				Invited:                []string{"tester1@gmail.com", "tester3@ibm.com"},
				AddNotificationPayload: core.BoolPtr(true),
				ReplyToMail:            core.StringPtr("testerreply@gmail.com"),
				ReplyToName:            core.StringPtr("rester_reply"),
				FromName:               core.StringPtr("Test IBM email"),
			}
			emailName := core.StringPtr("subscription_email")
			emailDescription := core.StringPtr("Subscription for email")
			createEmailSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          emailName,
				Description:   emailDescription,
				DestinationID: core.StringPtr(destinationID2),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreateAttributesEmailModel,
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createEmailSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.Attributes).ToNot(BeNil())
			Expect(subscription.Description).To(Equal(emailDescription))
			Expect(subscription.Name).To(Equal(emailName))
			subscriptionID2 = *subscription.ID

			Expect(subscriptionID2).ToNot(Equal(subscriptionID))

			createFCMSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("FCM subscription"),
				Description:   core.StringPtr("Subscription for the FCM"),
				DestinationID: core.StringPtr(destinationID3),
				TopicID:       core.StringPtr(topicID3),
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createFCMSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID3 = string(*subscription.ID)

			subscriptionCreateSlackAttributesModel := &eventnotificationsv1.SubscriptionCreateAttributesSlackAttributes{
				AttachmentColor: core.StringPtr("#0000FF"),
			}

			createSlackSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("Slack subscription"),
				Description:   core.StringPtr("Subscription for the Slack"),
				DestinationID: core.StringPtr(destinationID4),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreateSlackAttributesModel,
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createSlackSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID4 = string(*subscription.ID)

			createSafariSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("Safari subscription"),
				Description:   core.StringPtr("Subscription for the Safari"),
				DestinationID: core.StringPtr(destinationID5),
				TopicID:       core.StringPtr(topicID),
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createSafariSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID5 = string(*subscription.ID)

			createMSTeamsSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("MSTeams subscription"),
				Description:   core.StringPtr("Subscription for MSTeams"),
				DestinationID: core.StringPtr(destinationID6),
				TopicID:       core.StringPtr(topicID),
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createMSTeamsSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID6 = string(*subscription.ID)

			createCFSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("cloud functions subscription"),
				Description:   core.StringPtr("Subscription for cloud functions"),
				DestinationID: core.StringPtr(destinationID7),
				TopicID:       core.StringPtr(topicID),
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createCFSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID7 = string(*subscription.ID)

			createChromeSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("chrome subscription"),
				Description:   core.StringPtr("Subscription for chrome"),
				DestinationID: core.StringPtr(destinationID8),
				TopicID:       core.StringPtr(topicID),
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createChromeSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID8 = string(*subscription.ID)

			createFireSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("Firefox subscription"),
				Description:   core.StringPtr("Subscription for Firefox"),
				DestinationID: core.StringPtr(destinationID9),
				TopicID:       core.StringPtr(topicID),
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createFireSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID9 = string(*subscription.ID)

			createPDSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("Pager Duty subscription"),
				Description:   core.StringPtr("Subscription for Pager Duty"),
				DestinationID: core.StringPtr(destinationID10),
				TopicID:       core.StringPtr(topicID),
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createPDSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID10 = string(*subscription.ID)

			createServiceNowSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("Service Now subscription"),
				Description:   core.StringPtr("Subscription for Service Now"),
				DestinationID: core.StringPtr(destinationID11),
				TopicID:       core.StringPtr(topicID),
				Attributes: &eventnotificationsv1.SubscriptionCreateAttributesServiceNowAttributes{
					AssignedTo:      core.StringPtr("user"),
					AssignmentGroup: core.StringPtr("test"),
				},
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createServiceNowSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID11 = string(*subscription.ID)

			createFCMSubscriptionOptions = &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("FCM subscription V1"),
				Description:   core.StringPtr("Subscription for the FCM V1"),
				DestinationID: core.StringPtr(destinationID12),
				TopicID:       core.StringPtr(topicID3),
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createFCMSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID12 = string(*subscription.ID)

			subscriptionCreateCEAttributesModel := &eventnotificationsv1.SubscriptionCreateAttributes{
				SigningEnabled: core.BoolPtr(false),
			}

			ceName := core.StringPtr("subscription_code_engine")
			ceDescription := core.StringPtr("Subscription for code engine")
			createCESubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          ceName,
				Description:   ceDescription,
				DestinationID: core.StringPtr(destinationID13),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreateCEAttributesModel,
			}

			ceSubscription, response, err := eventNotificationsService.CreateSubscription(createCESubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(ceSubscription).ToNot(BeNil())
			Expect(ceSubscription.Attributes).ToNot(BeNil())
			Expect(ceSubscription.Description).To(Equal(ceDescription))
			Expect(ceSubscription.Name).To(Equal(ceName))
			subscriptionID13 = *ceSubscription.ID

			createCOSSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("COS subscription"),
				Description:   core.StringPtr("Subscription for the COS"),
				DestinationID: core.StringPtr(destinationID14),
				TopicID:       core.StringPtr(topicID),
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createCOSSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID14 = string(*subscription.ID)

			createHuaweiSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("Huawei subscription"),
				Description:   core.StringPtr("Subscription for the Huawei"),
				DestinationID: core.StringPtr(destinationID15),
				TopicID:       core.StringPtr(topicID),
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createHuaweiSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID15 = string(*subscription.ID)

			ceName = core.StringPtr("subscription_code_engine_job")
			ceDescription = core.StringPtr("Subscription for code engine job")
			createCEJobSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          ceName,
				Description:   ceDescription,
				DestinationID: core.StringPtr(destinationID18),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreateCEAttributesModel,
			}

			ceJobSubscription, response, err := eventNotificationsService.CreateSubscription(createCEJobSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(ceJobSubscription).ToNot(BeNil())
			Expect(ceJobSubscription.Attributes).ToNot(BeNil())
			Expect(ceJobSubscription.Description).To(Equal(ceDescription))
			Expect(ceJobSubscription.Name).To(Equal(ceName))
			subscriptionID18 = *ceJobSubscription.ID

			subscriptionCreateAttributesCustomEmailModel := &eventnotificationsv1.SubscriptionCreateAttributesCustomEmailAttributes{
				Invited:                []string{"abc@gmail.com", "tester3@ibm.com"},
				AddNotificationPayload: core.BoolPtr(true),
				ReplyToMail:            core.StringPtr("testerreply@gmail.com"),
				ReplyToName:            core.StringPtr("rester_reply"),
				FromName:               core.StringPtr("Test IBM email"),
				FromEmail:              core.StringPtr("test@test.event-notifications.test.cloud.ibm.com"),
				TemplateIDInvitation:   core.StringPtr(templateInvitationID),
				TemplateIDNotification: core.StringPtr(templateNotificationID),
			}
			customEmailName := core.StringPtr("subscription_custom_email")
			customEmailDescription := core.StringPtr("Subscription for custom email")
			createCustomEmailSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          customEmailName,
				Description:   customEmailDescription,
				DestinationID: core.StringPtr(destinationID16),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreateAttributesCustomEmailModel,
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createCustomEmailSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.Attributes).ToNot(BeNil())
			Expect(subscription.Description).To(Equal(customEmailDescription))
			Expect(subscription.Name).To(Equal(customEmailName))
			subscriptionID16 = *subscription.ID

			subscriptionCreateAttributesCustomSMSModel := &eventnotificationsv1.SubscriptionCreateAttributesCustomSmsAttributes{
				Invited: []string{"+12064563059", "+12267054625"},
			}
			customSMSName := core.StringPtr("subscription_custom_sms")
			customSMSDescription := core.StringPtr("Subscription for custom sms")
			createCustomSMSSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          customSMSName,
				Description:   customSMSDescription,
				DestinationID: core.StringPtr(destinationID17),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreateAttributesCustomSMSModel,
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createCustomSMSSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.Attributes).ToNot(BeNil())
			Expect(subscription.Description).To(Equal(customSMSDescription))
			Expect(subscription.Name).To(Equal(customSMSName))
			subscriptionID17 = *subscription.ID
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

	Describe(`ListTemplates - List all templates`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTemplates(listtemplatesOptions *ListTemplatesOptions)`, func() {

			listTemplatesOptions := &eventnotificationsv1.ListTemplatesOptions{
				InstanceID: core.StringPtr(instanceID),
				Offset:     core.Int64Ptr(int64(0)),
				Limit:      core.Int64Ptr(int64(1)),
				Search:     core.StringPtr(search),
			}

			templatesList, response, err := eventNotificationsService.ListTemplates(listTemplatesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templatesList).ToNot(BeNil())
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

			UpdateAttributesSMSInvitedModel := new(eventnotificationsv1.UpdateAttributesInvited)
			UpdateAttributesSMSInvitedModel.Add = []string{"+12064512559"}

			UpdateAttributesSMSSubscribedModel := new(eventnotificationsv1.UpdateAttributesSubscribed)
			UpdateAttributesSMSSubscribedModel.Remove = []string{"+12064512559"}

			UpdateAttributesSMSUnSubscribedModel := new(eventnotificationsv1.UpdateAttributesUnsubscribed)
			UpdateAttributesSMSUnSubscribedModel.Remove = []string{"+12064512559"}

			subscriptionUpdateSMSAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesSmsUpdateAttributes{
				Invited:      UpdateAttributesSMSInvitedModel,
				Subscribed:   UpdateAttributesSMSSubscribedModel,
				Unsubscribed: UpdateAttributesSMSUnSubscribedModel,
			}
			smsName := core.StringPtr("subscription_sms_update")
			smsDescription := core.StringPtr("Subscription update for sms")
			updateSubscriptionOptions = &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        smsName,
				Description: smsDescription,
				ID:          core.StringPtr(subscriptionID1),
				Attributes:  subscriptionUpdateSMSAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID1)))
			Expect(subscription.Name).To(Equal(smsName))
			Expect(subscription.Description).To(Equal(smsDescription))

			UpdateAttributesInvitedModel := new(eventnotificationsv1.UpdateAttributesInvited)
			UpdateAttributesInvitedModel.Add = []string{"tester4@ibm.com"}

			UpdateAttributessubscribedModel := new(eventnotificationsv1.UpdateAttributesSubscribed)
			UpdateAttributessubscribedModel.Remove = []string{"tester3@ibm.com"}

			UpdateAttributesUnSubscribedModel := new(eventnotificationsv1.UpdateAttributesUnsubscribed)
			UpdateAttributesUnSubscribedModel.Remove = []string{"tester3@ibm.com"}

			subscriptionUpdateEmailAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesEmailUpdateAttributes{
				Invited:                UpdateAttributesInvitedModel,
				AddNotificationPayload: core.BoolPtr(true),
				ReplyToMail:            core.StringPtr("testerreply@gmail.com"),
				ReplyToName:            core.StringPtr("rester_reply"),
				FromName:               core.StringPtr("Test IBM email"),
				Subscribed:             UpdateAttributessubscribedModel,
				Unsubscribed:           UpdateAttributesUnSubscribedModel,
			}
			emailName := core.StringPtr("subscription_email_update")
			emailDescription := core.StringPtr("Subscription update for email")
			updateSubscriptionOptions = &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        emailName,
				Description: emailDescription,
				ID:          core.StringPtr(subscriptionID2),
				Attributes:  subscriptionUpdateEmailAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID2)))
			Expect(subscription.Name).To(Equal(emailName))
			Expect(subscription.Description).To(Equal(emailDescription))

			fcmName := core.StringPtr("subscription_FCM_update")
			fcmDescription := core.StringPtr("Subscription update for FCM")
			updateFCMSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        fcmName,
				Description: fcmDescription,
				ID:          core.StringPtr(subscriptionID3),
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateFCMSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID3)))
			Expect(subscription.Name).To(Equal(fcmName))
			Expect(subscription.Description).To(Equal(fcmDescription))

			subscriptionUpdateSlackAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesSlackAttributes{
				AttachmentColor: core.StringPtr("#0000FF"),
			}

			slackName := core.StringPtr("subscription_slack_update")
			slackDescription := core.StringPtr("Subscription update for slack")
			updateSlackSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        slackName,
				Description: slackDescription,
				ID:          core.StringPtr(subscriptionID4),
				Attributes:  subscriptionUpdateSlackAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateSlackSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID4)))
			Expect(subscription.Name).To(Equal(slackName))
			Expect(subscription.Description).To(Equal(slackDescription))

			safariName := core.StringPtr("subscription_FCM")
			safariDescription := core.StringPtr("Subscription for FCM")
			updateSafariSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        safariName,
				Description: safariDescription,
				ID:          core.StringPtr(subscriptionID5),
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateSafariSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID5)))
			Expect(subscription.Name).To(Equal(safariName))
			Expect(subscription.Description).To(Equal(safariDescription))

			teamsName := core.StringPtr("subscription_MSTeams")
			teamsDescription := core.StringPtr("Subscription for MSTeams")
			updateTeamsSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        teamsName,
				Description: teamsDescription,
				ID:          core.StringPtr(subscriptionID6),
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateTeamsSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID6)))
			Expect(subscription.Name).To(Equal(teamsName))
			Expect(subscription.Description).To(Equal(teamsDescription))

			cfName := core.StringPtr("subscription_cloudfunctions")
			cfDescription := core.StringPtr("Subscription for cloud functions")
			updateCFSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        cfName,
				Description: cfDescription,
				ID:          core.StringPtr(subscriptionID7),
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateCFSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID7)))
			Expect(subscription.Name).To(Equal(cfName))
			Expect(subscription.Description).To(Equal(cfDescription))

			chromeName := core.StringPtr("subscription_Chrome")
			chromeDescription := core.StringPtr("Subscription for Chrome")
			updateChromeSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        chromeName,
				Description: chromeDescription,
				ID:          core.StringPtr(subscriptionID8),
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateChromeSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID8)))
			Expect(subscription.Name).To(Equal(chromeName))
			Expect(subscription.Description).To(Equal(chromeDescription))

			fireName := core.StringPtr("subscription_Firefox_update")
			fireDescription := core.StringPtr("Subscription for Firefox")
			updateFirefoxSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        fireName,
				Description: fireDescription,
				ID:          core.StringPtr(subscriptionID9),
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateFirefoxSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID9)))
			Expect(subscription.Name).To(Equal(fireName))
			Expect(subscription.Description).To(Equal(fireDescription))

			pdName := core.StringPtr("subscription_Pager_Duty_update")
			pdDescription := core.StringPtr("Subscription update for Pager Duty")
			updatePDSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        pdName,
				Description: pdDescription,
				ID:          core.StringPtr(subscriptionID10),
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updatePDSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID10)))
			Expect(subscription.Name).To(Equal(pdName))
			Expect(subscription.Description).To(Equal(pdDescription))

			serviceNowName := core.StringPtr("subscription_Service_Now_update")
			serviceNowDescription := core.StringPtr("Subscription update for Service_Now")
			updateServiceNowSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        serviceNowName,
				Description: serviceNowDescription,
				ID:          core.StringPtr(subscriptionID11),
				Attributes: &eventnotificationsv1.SubscriptionUpdateAttributesServiceNowAttributes{
					AssignedTo:      core.StringPtr("user"),
					AssignmentGroup: core.StringPtr("test"),
				},
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateServiceNowSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID11)))
			Expect(subscription.Name).To(Equal(serviceNowName))
			Expect(subscription.Description).To(Equal(serviceNowDescription))

			fcmName = core.StringPtr("subscription_FCM_V1_update")
			fcmDescription = core.StringPtr("Subscription update for FCM V1")
			updateFCMSubscriptionOptions = &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        fcmName,
				Description: fcmDescription,
				ID:          core.StringPtr(subscriptionID12),
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateFCMSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID12)))
			Expect(subscription.Name).To(Equal(fcmName))
			Expect(subscription.Description).To(Equal(fcmDescription))

			subscriptionCEUpdateAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesWebhookAttributes{
				SigningEnabled: core.BoolPtr(true),
			}

			ceName := core.StringPtr("code_engine_sub_updated")
			ceDescription := core.StringPtr("Update code engine subscription")
			updateCESubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(subscriptionID13),
				Name:        ceName,
				Description: ceDescription,
				Attributes:  subscriptionCEUpdateAttributesModel,
			}

			ceSubscription, response, err := eventNotificationsService.UpdateSubscription(updateCESubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ceSubscription).ToNot(BeNil())
			Expect(ceSubscription.ID).To(Equal(core.StringPtr(subscriptionID13)))
			Expect(ceSubscription.Name).To(Equal(ceName))
			Expect(ceSubscription.Description).To(Equal(ceDescription))

			cosName := core.StringPtr("subscription_COS_update")
			cosDescription := core.StringPtr("Subscription update COS")
			updatecCOSSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        cosName,
				Description: cosDescription,
				ID:          core.StringPtr(subscriptionID14),
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updatecCOSSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID14)))
			Expect(subscription.Name).To(Equal(cosName))
			Expect(subscription.Description).To(Equal(cosDescription))

			huaweiName := core.StringPtr("subscription_Huawei_update")
			huaweiDescription := core.StringPtr("Subscription update Huawei")
			updatecHuaweiSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        huaweiName,
				Description: huaweiDescription,
				ID:          core.StringPtr(subscriptionID15),
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updatecHuaweiSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID15)))
			Expect(subscription.Name).To(Equal(huaweiName))
			Expect(subscription.Description).To(Equal(huaweiDescription))

			UpdateAttributesCustomInvitedModel := new(eventnotificationsv1.UpdateAttributesInvited)
			UpdateAttributesCustomInvitedModel.Add = []string{"abc@gmail.com", "tester3@ibm.com"}

			UpdateAttributesCustomSubscribedModel := new(eventnotificationsv1.UpdateAttributesSubscribed)
			UpdateAttributesCustomSubscribedModel.Remove = []string{"tester3@ibm.com"}

			UpdateAttributesCustomUnSubscribedModel := new(eventnotificationsv1.UpdateAttributesUnsubscribed)
			UpdateAttributesCustomUnSubscribedModel.Remove = []string{"tester3@ibm.com"}

			subscriptionUpdateCustomEmailAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesCustomEmailUpdateAttributes{
				Invited:                UpdateAttributesCustomInvitedModel,
				AddNotificationPayload: core.BoolPtr(true),
				ReplyToMail:            core.StringPtr("testerreply@gmail.com"),
				ReplyToName:            core.StringPtr("rester_reply"),
				FromName:               core.StringPtr("Test IBM email"),
				FromEmail:              core.StringPtr("test@test.event-notifications.test.cloud.ibm.com"),
				Subscribed:             UpdateAttributesCustomSubscribedModel,
				Unsubscribed:           UpdateAttributesCustomUnSubscribedModel,
				TemplateIDInvitation:   core.StringPtr(templateInvitationID),
				TemplateIDNotification: core.StringPtr(templateNotificationID),
			}
			customEmailName := core.StringPtr("subscription_custom_email_update")
			CustomEmailDescription := core.StringPtr("Subscription update for custom email")
			updateSubscriptionOptions = &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        customEmailName,
				Description: CustomEmailDescription,
				ID:          core.StringPtr(subscriptionID16),
				Attributes:  subscriptionUpdateCustomEmailAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID16)))
			Expect(subscription.Name).To(Equal(customEmailName))
			Expect(subscription.Description).To(Equal(CustomEmailDescription))

			UpdateAttributesCustomSMSInvitedModel := new(eventnotificationsv1.UpdateAttributesInvited)
			UpdateAttributesCustomSMSInvitedModel.Add = []string{"+12064512559"}

			UpdateAttributesCustomSMSSubscribedModel := new(eventnotificationsv1.UpdateAttributesSubscribed)
			UpdateAttributesCustomSMSSubscribedModel.Remove = []string{"+12064512559"}

			UpdateAttributesCustomSMSUnSubscribedModel := new(eventnotificationsv1.UpdateAttributesUnsubscribed)
			UpdateAttributesCustomSMSUnSubscribedModel.Remove = []string{"+12064512559"}

			subscriptionUpdateCustomSMSAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesCustomSmsUpdateAttributes{
				Invited:      UpdateAttributesSMSInvitedModel,
				Subscribed:   UpdateAttributesSMSSubscribedModel,
				Unsubscribed: UpdateAttributesSMSUnSubscribedModel,
			}
			customSMSName := core.StringPtr("subscription_custom_sms_update")
			customSMSDescription := core.StringPtr("Subscription update for custom sms")
			updateSubscriptionOptions = &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        customSMSName,
				Description: customSMSDescription,
				ID:          core.StringPtr(subscriptionID17),
				Attributes:  subscriptionUpdateCustomSMSAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID17)))
			Expect(subscription.Name).To(Equal(customSMSName))
			Expect(subscription.Description).To(Equal(customSMSDescription))

			subscriptionCEJobUpdateAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesWebhookAttributes{
				SigningEnabled: core.BoolPtr(true),
			}

			ceName = core.StringPtr("code_engine_job_sub_updated")
			ceDescription = core.StringPtr("Update code engine job subscription")
			updateCEJobSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(subscriptionID18),
				Name:        ceName,
				Description: ceDescription,
				Attributes:  subscriptionCEJobUpdateAttributesModel,
			}

			ceJobSubscription, response, err := eventNotificationsService.UpdateSubscription(updateCEJobSubscriptionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ceJobSubscription).ToNot(BeNil())
			Expect(ceJobSubscription.ID).To(Equal(core.StringPtr(subscriptionID18)))
			Expect(ceJobSubscription.Name).To(Equal(ceName))
			Expect(ceJobSubscription.Description).To(Equal(ceDescription))
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

	Describe(`GetEnabledCountries - Get enabled countries`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEnabledCountries(getEnabledCountriesOptions *GetEnabledCountriesOptions)`, func() {

			getEnabledCountriesOptions := &eventnotificationsv1.GetEnabledCountriesOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(destinationID17),
			}

			enabledCountries, response, err := eventNotificationsService.GetEnabledCountries(getEnabledCountriesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(enabledCountries).ToNot(BeNil())
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

	Describe(`SendNotifications - Send a notification`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SendNotifications(sendNotificationsOptions *SendNotificationsOptions)`, func() {

			notificationDevicesModel := "{\"platforms\":[\"push_ios\",\"push_android\",\"push_chrome\",\"push_firefox\",\"push_huawei\"]}"
			notificationFcmBodyModel := "{\"message\": {\"android\": {\"notification\": {\"title\": \"Breaking News\",\"body\": \"New news story available.\"},\"data\": {\"name\": \"Willie Greenholt\",\"description\": \"description\"}}}}"
			notificationHuaweiBodyModel := "{\"message\": {\"android\": {\"notification\": {\"title\": \"Breaking News\",\"body\": \"New news story available.\"},\"data\": {\"name\": \"Willie Greenholt\",\"description\": \"description\"}}}}"
			notificationAPNsBodyModel := "{\"aps\":{\"alert\":{\"title\":\"Hello!! GameRequest\",\"body\":\"Bob wants to play poker\",\"action-loc-key\":\"PLAY\"},\"badge\":5}}"
			notificationSafariBodyModel := "{\"en_data\": {\"alert\": \"Alert message\"}}"
			mailTo := "[\"abc@ibm.com\", \"def@us.ibm.com\"]"
			smsTo := "[\"+911234567890\", \"+911224567890\"]"
			htmlBody := "\"Hi  ,<br/>Certificate expiring in 90 days.<br/><br/>Please login to <a href=\"https: //cloud.ibm.com/security-compliance/dashboard\">Security and Complaince dashboard</a> to find more information<br/>\""

			notificationSeverity := "MEDIUM"
			typeValue := "com.acme.offer:new"
			notificationsSouce := "1234-1234-sdfs-234:test"
			specVersion := "1.0"

			notificationCreateModel := &eventnotificationsv1.NotificationCreate{}
			notificationCreateModel.Ibmenseverity = &notificationSeverity
			notificationCreateModel.ID = &instanceID
			notificationCreateModel.Source = &notificationsSouce
			notificationCreateModel.Ibmensourceid = &sourceID
			notificationCreateModel.Type = &typeValue
			notificationCreateModel.Time = &strfmt.DateTime{}
			notificationCreateModel.Specversion = &specVersion
			notificationCreateModel.Ibmenfcmbody = &notificationFcmBodyModel
			notificationCreateModel.Ibmenapnsbody = &notificationAPNsBodyModel
			notificationCreateModel.Ibmensafaribody = &notificationSafariBodyModel
			notificationCreateModel.Ibmenhuaweibody = &notificationHuaweiBodyModel
			notificationCreateModel.Ibmenpushto = &notificationDevicesModel
			notificationCreateModel.Ibmenmailto = &mailTo
			notificationCreateModel.Ibmensmsto = &smsTo
			notificationCreateModel.Ibmensubject = core.StringPtr("Notification subject")
			notificationCreateModel.Ibmenhtmlbody = core.StringPtr(htmlBody)
			notificationCreateModel.Ibmendefaultshort = core.StringPtr("Alert message")
			notificationCreateModel.Ibmendefaultlong = core.StringPtr("Alert message on expiring offer")

			sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
			sendNotificationsOptionsModel.InstanceID = &instanceID
			sendNotificationsOptionsModel.Body = notificationCreateModel

			notificationResponse, response, err := eventNotificationsService.SendNotifications(sendNotificationsOptionsModel)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(notificationResponse).ToNot(BeNil())

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

			notificationResponse, response, err = eventNotificationsService.SendNotifications(sendNotificationsOptionsModel)

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

			for _, ID := range []string{subscriptionID, subscriptionID1, subscriptionID2, subscriptionID3, subscriptionID4, subscriptionID5, subscriptionID6, subscriptionID7, subscriptionID8, subscriptionID9, subscriptionID10, subscriptionID11, subscriptionID12, subscriptionID13, subscriptionID14, subscriptionID15, subscriptionID16, subscriptionID17, subscriptionID18} {

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

	Describe(`DeleteTemplate - Delete a Template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTemplate(deleteTemplateOptions *DeleteTemplateOptions)`, func() {

			for _, ID := range []string{templateInvitationID, templateNotificationID} {

				deleteTemplateOptions := &eventnotificationsv1.DeleteTemplateOptions{
					InstanceID: core.StringPtr(instanceID),
					ID:         core.StringPtr(ID),
				}

				response, err := eventNotificationsService.DeleteTemplate(deleteTemplateOptions)

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

	Describe(`DeleteDestination - Delete a Destination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDestination(deleteDestinationOptions *DeleteDestinationOptions)`, func() {

			for _, ID := range []string{destinationID, destinationID3, destinationID4, destinationID5, destinationID6, destinationID7, destinationID8, destinationID9, destinationID10, destinationID11, destinationID12, destinationID13, destinationID14, destinationID15, destinationID16, destinationID17, destinationID18} {
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

	Describe(`DeleteSource - Delete a Source`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSource(deleteSourceOptions *DeleteSourceOptions)`, func() {

			deleteSourceOptions := &eventnotificationsv1.DeleteSourceOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(sourceID),
			}

			response, err := eventNotificationsService.DeleteSource(deleteSourceOptions)

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
})

//
// Utility functions are declared in the unit test file
//
