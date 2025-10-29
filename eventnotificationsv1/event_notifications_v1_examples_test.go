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
	"strings"

	"github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
const externalConfigFile = "../event_notifications_v1.env"

var (
	eventNotificationsService *eventnotificationsv1.EventNotificationsV1
	config                    map[string]string
	configLoaded              bool = false
	instanceID                string
	safariCertificatePath     string
	topicName                 string = "Admin Topic Compliance"
	sourceID                  string = ""
	search                    string = ""
	topicID                   string
	destinationID             string
	destinationID1            string
	destinationID2            string
	destinationID3            string
	destinationID4            string
	destinationID5            string
	destinationID6            string
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
	destinationID19           string
	destinationID20           string
	destinationID21           string
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
	huaweiClientSecret        string
	huaweiClientID            string
	cosBucketName             string
	cosInstanceID             string
	cosEndPoint               string
	templateInvitationID      string
	templateNotificationID    string
	slackTemplateID           string
	templateBody              string
	cosInstanceCRN            string
	slackTemplateBody         string
	cosIntegrationID          string
	codeEngineProjectCRN      string
	smtpConfigID              string
	smtpUserID                string
	notificationID            string
	slackDMToken              string
	slackChannelID            string
	webhookTemplateID         string
	webhookTemplateBody       string
	pdTemplateID              string
	pdTemplateBody            string
	eventstreamscrn           string
	eventstreamsendpoint      string
	estopicName               string
	esTemplateID              string
	esTemplateBody            string
	cejobTemplateID           string
	cejobTemplateBody         string
	ceappTemplateID           string
	ceappTemplateBody         string
	testNotificationID        string
	appConfigCRN              string
	appConfigEnvID            string
	appConfigFeatureID        string
	appConfigTemplateBody     string
	acTemplateID              string
	smtpcloneconfigID         string
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
			fcmPrivateKey = strings.ReplaceAll(fcmPrivateKey, "\\n", "\n")
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

			templateBody = config["TEMPLATE_BODY"]
			if templateBody == "" {
				Skip("Unable to load templateBody configuration property, skipping tests")
			}
			fmt.Printf("TemplateBody: %s\n", templateBody)

			slackTemplateBody = config["SLACK_TEMPLATE_BODY"]
			if slackTemplateBody == "" {
				Skip("Unable to load slackTemplateBody configuration property, skipping tests")
			}
			fmt.Printf("slackTemplateBody: %s\n", slackTemplateBody)

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

			slackDMToken = config["SLACK_DM_TOKEN"]
			if slackDMToken == "" {
				Skip("Unable to load SLACK_DM_TOKEN configuration property, skipping tests")
			}
			fmt.Printf("SLACK_DM_TOKEN: %s\n", slackDMToken)

			slackChannelID = config["SLACK_CHANNEL_ID"]
			if slackChannelID == "" {
				Skip("Unable to load SLACK_CHANNEL_ID configuration property, skipping tests")
			}
			fmt.Printf("SLACK_CHANNEL_ID: %s\n", slackChannelID)

			webhookTemplateBody = config["WEBHOOK_TEMPLATE_BODY"]
			if webhookTemplateBody == "" {
				Skip("Unable to load webhookTemplateBody configuration property, skipping tests")
			}
			fmt.Printf("webhookTemplateBody: %s\n", webhookTemplateBody)

			pdTemplateBody = config["PAGERDUTY_TEMPLATE_BODY"]
			if pdTemplateBody == "" {
				Skip("Unable to load pdTemplateBody configuration property, skipping tests")
			}
			fmt.Printf("PAGERDUTY_TEMPLATE_BODY: %s\n", pdTemplateBody)

			eventstreamscrn = config["EVENT_STREAMS_CRN"]
			if eventstreamscrn == "" {
				Skip("Unable to load eventstreamscrn configuration property, skipping tests")
			}
			fmt.Printf("EVENT_STREAMS_CRN: %s\n", eventstreamscrn)

			eventstreamsendpoint = config["EVENT_STREAMS_ENDPOINT"]
			if eventstreamsendpoint == "" {
				Skip("Unable to load eventstreamsendpoint configuration property, skipping tests")
			}
			fmt.Printf("EVENT_STREAMS_ENDPOINT: %s\n", eventstreamsendpoint)

			estopicName = config["EVENT_STREAMS_TOPIC"]
			if estopicName == "" {
				Skip("Unable to load estopicName configuration property, skipping tests")
			}
			fmt.Printf("EVENT_STREAMS_TOPIC: %s\n", estopicName)

			esTemplateBody = config["EVENT_STREAMS_TEMPLATE_BODY"]
			if esTemplateBody == "" {
				Skip("Unable to load esTemplateBody configuration property, skipping tests")
			}
			fmt.Printf("EVENT_STREAMS_TEMPLATE_BODY: %s\n", esTemplateBody)

			cejobTemplateBody = config["CODE_ENGINE_JOB_TEMPLATE_BODY"]
			if cejobTemplateBody == "" {
				Skip("Unable to load cejobTemplateBody configuration property, skipping tests")
			}
			fmt.Printf("CODE_ENGINE_JOB_TEMPLATE_BODY: %s\n", cejobTemplateBody)

			ceappTemplateBody = config["CODE_ENGINE_APP_TEMPLATE_BODY"]
			if ceappTemplateBody == "" {
				Skip("Unable to load ceappTemplateBody configuration property, skipping tests")
			}
			fmt.Printf("CODE_ENGINE_APP_TEMPLATE_BODY: %s\n", ceappTemplateBody)

			appConfigCRN = config["APP_CONFIGURATION_INSTANCE_CRN"]
			if appConfigCRN == "" {
				Skip("Unable to load appConfigCRN  configuration property, skipping tests")
			}
			fmt.Printf("APP_CONFIGURATION_INSTANCE_CRN: %s\n", appConfigCRN)

			appConfigEnvID = config["APP_CONFIGURATION_ENVIRONMENT_ID"]
			if appConfigCRN == "" {
				Skip("Unable to load appConfigEnvID  configuration property, skipping tests")
			}
			fmt.Printf("APP_CONFIGURATION_ENVIRONMENT_ID: %s\n", appConfigEnvID)

			appConfigFeatureID = config["APP_CONFIGURATION_FEATURE_ID"]
			if appConfigCRN == "" {
				Skip("Unable to load appConfigFeatureID  configuration property, skipping tests")
			}
			fmt.Printf("APP_CONFIGURATION_FEATURE_ID: %s\n", appConfigFeatureID)

			appConfigTemplateBody = config["APP_CONFIGURATION_TEMPLATE_BODY"]
			if appConfigCRN == "" {
				Skip("Unable to load appConfigTemplateBody configuration property, skipping tests")
			}
			fmt.Printf("APP_CONFIGURATION_TEMPLATE_BODY: %s\n", appConfigTemplateBody)

			smtpcloneconfigID = config["SMTP_CLONE_CONFIG_ID"]
			if smtpcloneconfigID == "" {
				Skip("Unable to load smtpcloneconfigID configuration property, skipping tests")
			}
			fmt.Printf("SMTP_CLONE_CONFIG_ID: %s\n", smtpcloneconfigID)

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

	Describe(`CreateIntegration - Create integration of an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateIntegration(createIntegrationOptions *CreateIntegrationOptions)`, func() {
			// begin-create_integration
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
			// end-create_integration
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(integrationCreateResponse).ToNot(BeNil())
		})
	})

	Describe(`EventNotificationsV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`listIntegrations request example`, func() {
			// begin-list_integrations

			listIntegrationsOptions := &eventnotificationsv1.ListIntegrationsOptions{
				InstanceID: core.StringPtr(instanceID),
				Limit:      core.Int64Ptr(int64(1)),
				Offset:     core.Int64Ptr(int64(0)),
				Search:     core.StringPtr(search),
			}

			integrationResponse, response, err := eventNotificationsService.ListIntegrations(listIntegrationsOptions)

			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from listIntegrations(): %d\n", response.StatusCode)
			}
			integrationId = string(*integrationResponse.Integrations[0].ID)
			// end-list_integrations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})

		It(`getIntegration request example`, func() {
			// begin-get_integration

			listIntegrationsOptions := &eventnotificationsv1.GetIntegrationOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(integrationId),
			}

			_, response, err := eventNotificationsService.GetIntegration(listIntegrationsOptions)

			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from getIntegration(): %d\n", response.StatusCode)
			}

			// end-get_integration
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})

		It(`updateIntegration request example`, func() {
			// begin-replace_integration

			integrationMetadata := &eventnotificationsv1.IntegrationMetadata{
				Endpoint:  core.StringPtr("https://private.us-south.kms.cloud.ibm.com"),
				CRN:       core.StringPtr("insert CRN"),
				RootKeyID: core.StringPtr("insert Root Key Id"),
			}

			replaceIntegrationsOptions := &eventnotificationsv1.ReplaceIntegrationOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(integrationId),
				Type:       core.StringPtr("kms/hs-crypto"),
				Metadata:   integrationMetadata,
			}

			_, response, err := eventNotificationsService.ReplaceIntegration(replaceIntegrationsOptions)

			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from updateIntegration(): %d\n", response.StatusCode)
			}

			//COS Integration
			integrationCOSMetadata := &eventnotificationsv1.IntegrationMetadata{
				Endpoint:   core.StringPtr(cosEndPoint),
				CRN:        core.StringPtr(cosInstanceCRN),
				BucketName: core.StringPtr(cosBucketName),
			}

			replaceCOSIntegrationsOptions := &eventnotificationsv1.ReplaceIntegrationOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(cosIntegrationID),
				Type:       core.StringPtr("collect_failed_events"),
				Metadata:   integrationCOSMetadata,
			}

			_, response, err = eventNotificationsService.ReplaceIntegration(replaceCOSIntegrationsOptions)

			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from updateIntegration(): %d\n", response.StatusCode)
			}

			// end-replace_integration
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

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

			topicUpdateSourcesItemModel := &eventnotificationsv1.SourcesItems{
				ID:    core.StringPtr(sourceID),
				Rules: []eventnotificationsv1.Rules{*rulesModel},
			}

			createTopicOptions := &eventnotificationsv1.CreateTopicOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(topicName),
				Description: core.StringPtr("This topic is used for routing all compliance related notifications to the appropriate destinations"),
				Sources:     []eventnotificationsv1.SourcesItems{*topicUpdateSourcesItemModel},
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

			topicUpdateSourcesItemModel := &eventnotificationsv1.SourcesItems{
				ID:    core.StringPtr(sourceID),
				Rules: []eventnotificationsv1.Rules{*rulesModel},
			}

			replaceTopicOptions := eventNotificationsService.NewReplaceTopicOptions(
				instanceID,
				topicID,
			)
			replaceTopicOptions.SetSources([]eventnotificationsv1.SourcesItems{*topicUpdateSourcesItemModel})
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

			destinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfFcmDestinationConfig{
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

			//webhook
			webHookDestinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfWebhookDestinationConfig{
				URL:  core.StringPtr("https://gcm.com"),
				Verb: core.StringPtr("get"),
				CustomHeaders: map[string]string{
					"gcm_apikey": "api_key_value",
				},
				SensitiveHeaders: []string{"gcm_apikey"},
			}

			webHookDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: webHookDestinationConfigParamsModel,
			}

			name := "Webhook_destination"
			typeVal := "webhook"
			description := "Webhook Destination"
			createWebHookDestinationOptions := &eventnotificationsv1.CreateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(typeVal),
				Description: core.StringPtr(description),
				Config:      webHookDestinationConfigModel,
			}

			destinationResponse, response, err = eventNotificationsService.CreateDestination(createWebHookDestinationOptions)
			if err != nil {
				panic(err)
			}

			b, _ = json.MarshalIndent(destinationResponse, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID3 = *destinationResponse.ID

			//slack
			createSlackDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"Slack_destination",
				eventnotificationsv1.CreateDestinationOptionsTypeSlackConst,
			)

			destinationConfigParamsSlackModel := &eventnotificationsv1.DestinationConfigOneOfSlackDestinationConfig{
				URL:  core.StringPtr("https://api.slack.com/myslack"),
				Type: core.StringPtr("incoming_webhook"),
			}

			slackDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsSlackModel,
			}

			createSlackDestinationOptions.SetConfig(slackDestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(createSlackDestinationOptions)
			if err != nil {
				panic(err)
			}

			b, _ = json.MarshalIndent(destinationResponse, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID4 = *destinationResponse.ID

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

			destinationConfigParamsSafariModel := &eventnotificationsv1.DestinationConfigOneOfSafariDestinationConfig{
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

			//MSTeams
			createMSTeamsDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"MSTeams_destination",
				eventnotificationsv1.CreateDestinationOptionsTypeMsteamsConst,
			)

			destinationConfigParamsMSTeaMSModel := &eventnotificationsv1.DestinationConfigOneOfMsTeamsDestinationConfig{
				URL: core.StringPtr("https://teams.microsoft.com"),
			}

			msTeamsDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsMSTeaMSModel,
			}

			createMSTeamsDestinationOptions.SetConfig(msTeamsDestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(createMSTeamsDestinationOptions)
			if err != nil {
				panic(err)
			}

			b, _ = json.MarshalIndent(destinationResponse, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID6 = *destinationResponse.ID

			//Chrome
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

			b, _ = json.MarshalIndent(destinationResponse, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID8 = *destinationResponse.ID

			//Firefox
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

			b, _ = json.MarshalIndent(destinationResponse, "", "  ")
			fmt.Println(string(b))

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
				APIKey:     core.StringPtr("insert API key here"),
				RoutingKey: core.StringPtr("insert Routing Key here"),
			}

			pagerDutyDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsPDModel,
			}

			pagerDutyCreateDestinationOptions.SetConfig(pagerDutyDestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(pagerDutyCreateDestinationOptions)
			if err != nil {
				panic(err)
			}

			b, _ = json.MarshalIndent(destinationResponse, "", "  ")
			fmt.Println(string(b))

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
			b, _ = json.MarshalIndent(destinationResponse, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID11 = *destinationResponse.ID

			createFCMV1DestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"FCM_destination_V1",
				eventnotificationsv1.CreateDestinationOptionsTypePushAndroidConst,
			)

			destinationFCMV1ConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfFcmDestinationConfig{
				ProjectID:   core.StringPtr(fcmProjectID),
				PrivateKey:  core.StringPtr(fcmPrivateKey),
				ClientEmail: core.StringPtr(fcmClientEmail),
			}

			destinationFCMV1ConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationFCMV1ConfigParamsModel,
			}

			createDestinationOptions.SetConfig(destinationFCMV1ConfigModel)

			destinationResponse, response, err = eventNotificationsService.CreateDestination(createFCMV1DestinationOptions)
			if err != nil {
				panic(err)
			}

			b, _ = json.MarshalIndent(destinationResponse, "", "  ")
			fmt.Println(string(b))

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

			destinationResponse, response, err = eventNotificationsService.CreateDestination(createCEDestinationOptions)
			if err != nil {
				panic(err)
			}

			b, _ = json.MarshalIndent(destinationResponse, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID13 = *destinationResponse.ID

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

			b, _ = json.MarshalIndent(destinationResponse, "", "  ")
			fmt.Println(string(b))

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

			b, _ = json.MarshalIndent(destinationResponse, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID15 = *destinationResponse.ID

			customDestinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig{
				Domain: core.StringPtr("abc.event-notifications.test.cloud.ibm.com"),
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

			destinationID18 = *destinationCEJobResponse.ID

			createSlackDMDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"Slack_DM_destination",
				eventnotificationsv1.CreateDestinationOptionsTypeSlackConst,
			)

			destinationConfigParamsSlackDMModel := &eventnotificationsv1.DestinationConfigOneOfSlackDirectMessageDestinationConfig{
				Token: core.StringPtr(slackDMToken),
				Type:  core.StringPtr("direct_message"),
			}

			slackDMDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsSlackDMModel,
			}

			createSlackDMDestinationOptions.SetConfig(slackDMDestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(createSlackDMDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID19 = *destinationResponse.ID

			createEventStreamsDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"Event_Streams_destination",
				eventnotificationsv1.CreateDestinationOptionsTypeEventStreamsConst,
			)

			destinationConfigParamsEventStreamsModel := &eventnotificationsv1.DestinationConfigOneOfEventStreamsDestinationConfig{
				CRN:      core.StringPtr(eventstreamscrn),
				Endpoint: core.StringPtr(eventstreamsendpoint),
				Topic:    core.StringPtr(estopicName),
			}

			eventstreamsDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsEventStreamsModel,
			}

			createEventStreamsDestinationOptions.SetConfig(eventstreamsDestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(createEventStreamsDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID20 = *destinationResponse.ID

			createAppConfigurationsDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
				instanceID,
				"App_Configuration_destination",
				eventnotificationsv1.CreateDestinationOptionsTypeAppConfigurationConst,
			)

			destinationConfigParamsAppConfigurationModel := &eventnotificationsv1.DestinationConfigOneOfAppConfigurationDestinationConfig{
				Type:          core.StringPtr("features"),
				CRN:           core.StringPtr(appConfigCRN),
				EnvironmentID: core.StringPtr(appConfigEnvID),
				FeatureID:     core.StringPtr(appConfigFeatureID),
			}

			appconfigurationDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsAppConfigurationModel,
			}

			createAppConfigurationsDestinationOptions.SetConfig(appconfigurationDestinationConfigModel)
			destinationResponse, response, err = eventNotificationsService.CreateDestination(createAppConfigurationsDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(destinationResponse).ToNot(BeNil())

			destinationID21 = *destinationResponse.ID
			// end-create_destination

		})

		It(`TestDestination(TestDestinationOptions *testDestinationOptions)`, func() {

			// begin-test_destination
			testDestinationOptions := &eventnotificationsv1.TestDestinationOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(destinationID14),
			}

			_, response, err := eventNotificationsService.TestDestination(testDestinationOptions)
			// end-test_destination
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})

		It(`TestWebhookDestination(TestDestinationOptions *testDestinationOptions)`, func() {

			// begin-test_destination
			testDestinationOptions := &eventnotificationsv1.TestDestinationOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(destinationID),
			}

			testWebhookResponse, response, err := eventNotificationsService.TestDestination(testDestinationOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(testWebhookResponse).ToNot(BeNil())

			if testResponse, ok := testWebhookResponse.(*eventnotificationsv1.TestDestinationResponse); ok {
				if testResponse.NotificationID != nil {
					testNotificationID = *testResponse.NotificationID
				}
			}

			getNoticationOptions := &eventnotificationsv1.GetNotificationsStatusOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(testNotificationID),
			}

			getNotificationResponse, response, err := eventNotificationsService.GetNotificationsStatus(getNoticationOptions)
			// end-test_destination
			if err != nil {
				panic(err)
			}

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getNotificationResponse).ToNot(BeNil())
			Expect(getNotificationResponse.Status).To(Equal(core.StringPtr("success")))

		})

		It(`CreateTemplate request example`, func() {
			fmt.Println("\nCreateTemplate() result:")
			// begin-create_template
			name := "template invitation"
			description := "template invitation description"
			templateTypeInvitation := "smtp_custom.invitation"
			templateTypeNotification := "smtp_custom.notification"
			templateTypeSlack := "slack.notification"
			templateTypeWebhook := "webhook.notification"
			templateTypePD := "pagerduty.notification"
			templateTypeES := "event_streams.notification"
			templateTypeCEJOB := "ibmcejob.notification"
			templateTypeCEAPP := "ibmceapp.notification"
			templateTypeAC := "app_configuration.notification"

			templConfig := &eventnotificationsv1.TemplateConfigOneOfEmailTemplateConfig{
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

			templateInvitationID = *templateResponse.ID

			name = "template notification"
			description = "template notification description"

			templConfig = &eventnotificationsv1.TemplateConfigOneOfEmailTemplateConfig{
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

			templateNotificationID = *templateResponse.ID

			name = "slack template"
			description = "slack template description"

			slackTemplConfig := &eventnotificationsv1.TemplateConfigOneOfSlackTemplateConfig{
				Body: core.StringPtr(slackTemplateBody),
			}

			createTemplateOptions = &eventnotificationsv1.CreateTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeSlack),
				Description: core.StringPtr(description),
				Params:      slackTemplConfig,
			}

			templateResponse, response, err = eventNotificationsService.CreateTemplate(createTemplateOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(templateResponse).ToNot(BeNil())

			slackTemplateID = *templateResponse.ID

			name = "webhook template"
			description = "webhook template description"

			webhookTemplConfig := &eventnotificationsv1.TemplateConfigOneOfWebhookTemplateConfig{
				Body: core.StringPtr(webhookTemplateBody),
			}

			createTemplateOptions = &eventnotificationsv1.CreateTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeWebhook),
				Description: core.StringPtr(description),
				Params:      webhookTemplConfig,
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

			webhookTemplateID = *templateResponse.ID

			name = "Pagerduty template"
			description = "pagerduty template description"

			pdTemplConfig := &eventnotificationsv1.TemplateConfigOneOfPagerdutyTemplateConfig{
				Body: core.StringPtr(pdTemplateBody),
			}

			createTemplateOptions = &eventnotificationsv1.CreateTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypePD),
				Description: core.StringPtr(description),
				Params:      pdTemplConfig,
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

			pdTemplateID = *templateResponse.ID

			name = "Event Streams template"
			description = "Event Streams template description"

			eventstreamsTemplConfig := &eventnotificationsv1.TemplateConfigOneOfEventStreamsTemplateConfig{
				Body: core.StringPtr(esTemplateBody),
			}

			createTemplateOptions = &eventnotificationsv1.CreateTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeES),
				Description: core.StringPtr(description),
				Params:      eventstreamsTemplConfig,
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

			esTemplateID = *templateResponse.ID

			name = "Code Engine Job template"
			description = "Code Engine Job template description"

			codeenginejobTemplConfig := &eventnotificationsv1.TemplateConfigOneOfCodeEngineJobTemplateConfig{
				Body: core.StringPtr(cejobTemplateBody),
			}

			createTemplateOptions = &eventnotificationsv1.CreateTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeCEJOB),
				Description: core.StringPtr(description),
				Params:      codeenginejobTemplConfig,
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

			cejobTemplateID = *templateResponse.ID

			name = "Code Engine app template"
			description = "Code Engine app template description"

			codeengineappTemplConfig := &eventnotificationsv1.TemplateConfigOneOfCodeEngineApplicationTemplateConfig{
				Body: core.StringPtr(ceappTemplateBody),
			}

			createTemplateOptions = &eventnotificationsv1.CreateTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeCEAPP),
				Description: core.StringPtr(description),
				Params:      codeengineappTemplConfig,
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

			ceappTemplateID = *templateResponse.ID

			name = "App Configuration template"
			description = "App Configuration template description"

			appconfigTemplConfig := &eventnotificationsv1.TemplateConfigOneOfAppConfigurationTemplateConfig{
				Body: core.StringPtr(appConfigTemplateBody),
			}

			createTemplateOptions = &eventnotificationsv1.CreateTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeAC),
				Description: core.StringPtr(description),
				Params:      appconfigTemplConfig,
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

			acTemplateID = *templateResponse.ID
			// end-create_template
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

		It(`GetTemplate request example`, func() {
			fmt.Println("\nGetTemplate() result:")
			// begin-get_template

			getTemplateOptions := &eventnotificationsv1.GetTemplateOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(templateInvitationID),
			}

			template, response, err := eventNotificationsService.GetTemplate(getTemplateOptions)
			// end-get_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(template).ToNot(BeNil())

		})

		It(`GetPreDefinedTemplate request example`, func() {
			fmt.Println("\nGetTemplate() result:")
			// begin-get_pre_defined_template

			predefinedtemplateID := "0cacb9a0-d43a-4042-920d-d4a3f7d4cbd5"
			getPredefinedTemplateOptions := &eventnotificationsv1.GetPreDefinedTemplateOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(predefinedtemplateID),
			}

			predefinedtemplate, response, err := eventNotificationsService.GetPreDefinedTemplate(getPredefinedTemplateOptions)
			// end-get_pre_defined_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(predefinedtemplate).ToNot(BeNil())

		})

		It(`UpdateDestination request example`, func() {
			fmt.Println("\nUpdateDestination() result:")
			// begin-update_destination

			//webhook
			webHookDestinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfWebhookDestinationConfig{
				URL:  core.StringPtr("https://cloud.ibm.com/nhwebhook/sendwebhook"),
				Verb: core.StringPtr("post"),
				CustomHeaders: map[string]string{
					"authorization": "authorization key",
				},
				SensitiveHeaders: []string{"authorization"},
			}

			webHookDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: webHookDestinationConfigParamsModel,
			}

			webName := "Admin Webhook Compliance"
			webDescription := "This destination is for creating admin Webhook to receive compliance notifications"
			webUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID3),
				Name:        core.StringPtr(webName),
				Description: core.StringPtr(webDescription),
				Config:      webHookDestinationConfigModel,
			}

			destination, response, err := eventNotificationsService.UpdateDestination(webUpdateDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

			//slack
			destinationConfigParamsSlackModel := &eventnotificationsv1.DestinationConfigOneOfSlackDestinationConfig{
				URL:  core.StringPtr("https://api.slack.com/myslack"),
				Type: core.StringPtr("incoming_webhook"),
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

			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

			//FCM
			destinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfFcmDestinationConfig{
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

			destination, response, err = eventNotificationsService.UpdateDestination(updateDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

			//Safari
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

			Expect(err).To(BeNil())
			Expect(safariresponse.StatusCode).To(Equal(200))
			Expect(safaridestination).ToNot(BeNil())

			//MSTeams

			destinationConfigParamsMSTeaMSModel := &eventnotificationsv1.DestinationConfigOneOfMsTeamsDestinationConfig{
				URL: core.StringPtr("https://teams.microsoft.com"),
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

			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

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

			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

			//Firefox
			destinationConfigParamsfireModel := &eventnotificationsv1.DestinationConfigOneOfFirefoxDestinationConfig{
				WebsiteURL: core.StringPtr("https://cloud.ibm.com"),
			}

			fireDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsfireModel,
			}

			fireName := "Firefox_destination"
			fireDescription := "This destination is for Firefox"
			fireUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID9),
				Name:        core.StringPtr(fireName),
				Description: core.StringPtr(fireDescription),
				Config:      fireDestinationConfigModel,
			}

			destination, response, err = eventNotificationsService.UpdateDestination(fireUpdateDestinationOptions)

			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

			destinationConfigParamsPDModel := &eventnotificationsv1.DestinationConfigOneOfPagerDutyDestinationConfig{
				APIKey:     core.StringPtr("insert API Key here"),
				RoutingKey: core.StringPtr("insert Routing Key here"),
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
			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

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
			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

			//FCM V1

			destinationConfigFCMV1ParamsModel := &eventnotificationsv1.DestinationConfigOneOfFcmDestinationConfig{
				ProjectID:   core.StringPtr(fcmProjectID),
				PrivateKey:  core.StringPtr(fcmPrivateKey),
				ClientEmail: core.StringPtr(fcmClientEmail),
			}
			destinationConfigFCMV1Model := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigFCMV1ParamsModel,
			}

			updateFCMV1DestinationOptions := eventNotificationsService.NewUpdateDestinationOptions(
				instanceID,
				destinationID12,
			)

			updateFCMV1DestinationOptions.SetName("Admin FCM V1 Compliance")
			updateFCMV1DestinationOptions.SetDescription("This destination is for creating admin FCM V1 to receive compliance notifications")
			updateFCMV1DestinationOptions.SetConfig(destinationConfigFCMV1Model)

			destination, response, err = eventNotificationsService.UpdateDestination(updateFCMV1DestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

			//Code Engine

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

			destination, response, err = eventNotificationsService.UpdateDestination(updateCEDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

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

			destination, response, err = eventNotificationsService.UpdateDestination(cosUpdateDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

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

			destination, response, err = eventNotificationsService.UpdateDestination(huaweiUpdateDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

			customDestinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfCustomDomainEmailDestinationConfig{
				Domain: core.StringPtr("abc.event-notifications.test.cloud.ibm.com"),
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

			destination, response, err = eventNotificationsService.UpdateDestination(customUpdateDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(destination, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())

			customEmailUpdateDestinationOptions := &eventnotificationsv1.UpdateVerifyDestinationOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(destinationID16),
				Type:       core.StringPtr("spf/dkim"),
			}

			spfDkimResult, response, err := eventNotificationsService.UpdateVerifyDestination(customEmailUpdateDestinationOptions)
			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(spfDkimResult, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(spfDkimResult).ToNot(BeNil())

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

			destinationConfigParamsSlackDMModel := &eventnotificationsv1.DestinationConfigOneOfSlackDirectMessageDestinationConfig{
				Token: core.StringPtr(slackDMToken),
				Type:  core.StringPtr("direct_message"),
			}

			slackDMDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsSlackDMModel,
			}

			slackDMName := "slack_DM_destination_update"
			slackDMDescription := "This destination is for slack DM"
			slackDMUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID19),
				Name:        core.StringPtr(slackDMName),
				Description: core.StringPtr(slackDMDescription),
				Config:      slackDMDestinationConfigModel,
			}

			destination, response, err = eventNotificationsService.UpdateDestination(slackDMUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())
			Expect(destination.ID).To(Equal(core.StringPtr(destinationID19)))
			Expect(destination.Name).To(Equal(core.StringPtr(slackDMName)))
			Expect(destination.Description).To(Equal(core.StringPtr(slackDMDescription)))

			destinationConfigParamsEventStreamsModel := &eventnotificationsv1.DestinationConfigOneOfEventStreamsDestinationConfig{
				CRN:      core.StringPtr(eventstreamscrn),
				Endpoint: core.StringPtr(eventstreamsendpoint),
				Topic:    core.StringPtr(estopicName),
			}

			eventstreamsDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsEventStreamsModel,
			}

			eventStreamsName := "Event_Streams_destination_update"
			eventStreamsDescription := "This destination is for Event Streams Destination Update"
			eventStreamsUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID20),
				Name:        core.StringPtr(eventStreamsName),
				Description: core.StringPtr(eventStreamsDescription),
				Config:      eventstreamsDestinationConfigModel,
			}

			destination, response, err = eventNotificationsService.UpdateDestination(eventStreamsUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())
			Expect(destination.ID).To(Equal(core.StringPtr(destinationID20)))
			Expect(destination.Name).To(Equal(core.StringPtr(eventStreamsName)))
			Expect(destination.Description).To(Equal(core.StringPtr(eventStreamsDescription)))

			destinationConfigParamsAppConfigModel := &eventnotificationsv1.DestinationConfigOneOfAppConfigurationDestinationConfig{
				Type:          core.StringPtr("features"),
				CRN:           core.StringPtr(appConfigCRN),
				EnvironmentID: core.StringPtr(appConfigEnvID),
				FeatureID:     core.StringPtr(appConfigFeatureID),
			}

			appconfigDestinationConfigModel := &eventnotificationsv1.DestinationConfig{
				Params: destinationConfigParamsAppConfigModel,
			}

			appConfigName := "App_Configuration_destination_update"
			appConfigDescription := "This destination is for AppConfiguration Destination Update"
			appConfigUpdateDestinationOptions := &eventnotificationsv1.UpdateDestinationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(destinationID21),
				Name:        core.StringPtr(appConfigName),
				Description: core.StringPtr(appConfigDescription),
				Config:      appconfigDestinationConfigModel,
			}

			destination, response, err = eventNotificationsService.UpdateDestination(appConfigUpdateDestinationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(destination).ToNot(BeNil())
			Expect(destination.ID).To(Equal(core.StringPtr(destinationID21)))
			Expect(destination.Name).To(Equal(core.StringPtr(appConfigName)))
			Expect(destination.Description).To(Equal(core.StringPtr(appConfigDescription)))

			// end-update_destination
		})

		It(`UpdateTemplate request example`, func() {
			fmt.Println("\nUpdateTemplate() result:")
			// begin-replace_template
			name := "template invitation"
			description := "template invitation description"
			templateTypeInvitation := "smtp_custom.invitation"
			templateTypeNotification := "smtp_custom.notification"
			templateTypeSlack := "slack.notification"
			templateTypeWebhook := "webhook.notification"
			templateTypePD := "pagerduty.notification"
			templateTypeES := "event_streams.notification"
			templateTypeCEJOB := "ibmcejob.notification"
			templateTypeCEAPP := "ibmceapp.notification"
			templateTypeAC := "app_configuration.notification"

			templateConfig := &eventnotificationsv1.TemplateConfigOneOfEmailTemplateConfig{
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

			templateConfig = &eventnotificationsv1.TemplateConfigOneOfEmailTemplateConfig{
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

			name = "slack template"
			description = "slack template description"

			slackTemplateConfig := &eventnotificationsv1.TemplateConfigOneOfSlackTemplateConfig{
				Body: core.StringPtr(slackTemplateBody),
			}

			replaceTemplateOptions = &eventnotificationsv1.ReplaceTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(slackTemplateID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeSlack),
				Description: core.StringPtr(description),
				Params:      slackTemplateConfig,
			}

			templateResponse, response, err = eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResponse).ToNot(BeNil())
			Expect(templateResponse.ID).To(Equal(core.StringPtr(slackTemplateID)))
			Expect(templateResponse.Name).To(Equal(core.StringPtr(name)))
			Expect(templateResponse.Description).To(Equal(core.StringPtr(description)))

			name = "webhook template"
			description = "webhook template description"

			webhookTemplateConfig := &eventnotificationsv1.TemplateConfigOneOfWebhookTemplateConfig{
				Body: core.StringPtr(webhookTemplateBody),
			}

			replaceTemplateOptions = &eventnotificationsv1.ReplaceTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(webhookTemplateID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeWebhook),
				Description: core.StringPtr(description),
				Params:      webhookTemplateConfig,
			}

			templateResponse, response, err = eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResponse).ToNot(BeNil())
			Expect(templateResponse.ID).To(Equal(core.StringPtr(webhookTemplateID)))
			Expect(templateResponse.Name).To(Equal(core.StringPtr(name)))
			Expect(templateResponse.Description).To(Equal(core.StringPtr(description)))

			name = "webhook template"
			description = "webhook template description"

			name = "PagerDuty template update"
			description = "pagerduty template description"

			pdTemplateConfig := &eventnotificationsv1.TemplateConfigOneOfPagerdutyTemplateConfig{
				Body: core.StringPtr(pdTemplateBody),
			}

			replaceTemplateOptions = &eventnotificationsv1.ReplaceTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(pdTemplateID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypePD),
				Description: core.StringPtr(description),
				Params:      pdTemplateConfig,
			}

			templateResponse, response, err = eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResponse).ToNot(BeNil())
			Expect(templateResponse.ID).To(Equal(core.StringPtr(pdTemplateID)))
			Expect(templateResponse.Name).To(Equal(core.StringPtr(name)))
			Expect(templateResponse.Description).To(Equal(core.StringPtr(description)))

			name = "Event Streams template update"
			description = "Event Streams template description update"

			eventstreamsTemplateConfig := &eventnotificationsv1.TemplateConfigOneOfEventStreamsTemplateConfig{
				Body: core.StringPtr(esTemplateBody),
			}

			replaceTemplateOptions = &eventnotificationsv1.ReplaceTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(esTemplateID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeES),
				Description: core.StringPtr(description),
				Params:      eventstreamsTemplateConfig,
			}

			templateResponse, response, err = eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResponse).ToNot(BeNil())
			Expect(templateResponse.ID).To(Equal(core.StringPtr(esTemplateID)))
			Expect(templateResponse.Name).To(Equal(core.StringPtr(name)))
			Expect(templateResponse.Description).To(Equal(core.StringPtr(description)))

			name = "Code Engine Job template update"
			description = "Code Engine Job description update"

			codeenginejobTemplateConfig := &eventnotificationsv1.TemplateConfigOneOfCodeEngineJobTemplateConfig{
				Body: core.StringPtr(cejobTemplateBody),
			}

			replaceTemplateOptions = &eventnotificationsv1.ReplaceTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(cejobTemplateID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeCEJOB),
				Description: core.StringPtr(description),
				Params:      codeenginejobTemplateConfig,
			}

			templateResponse, response, err = eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResponse).ToNot(BeNil())
			Expect(templateResponse.ID).To(Equal(core.StringPtr(cejobTemplateID)))
			Expect(templateResponse.Name).To(Equal(core.StringPtr(name)))
			Expect(templateResponse.Description).To(Equal(core.StringPtr(description)))

			name = "Code Engine APP template update"
			description = "Code Engine APP description update"

			codeengineappTemplateConfig := &eventnotificationsv1.TemplateConfigOneOfCodeEngineApplicationTemplateConfig{
				Body: core.StringPtr(ceappTemplateBody),
			}

			replaceTemplateOptions = &eventnotificationsv1.ReplaceTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(ceappTemplateID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeCEAPP),
				Description: core.StringPtr(description),
				Params:      codeengineappTemplateConfig,
			}

			templateResponse, response, err = eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResponse).ToNot(BeNil())
			Expect(templateResponse.ID).To(Equal(core.StringPtr(ceappTemplateID)))
			Expect(templateResponse.Name).To(Equal(core.StringPtr(name)))
			Expect(templateResponse.Description).To(Equal(core.StringPtr(description)))

			name = "App Configuration template update"
			description = "App Configuration description update"

			appconfigTemplateConfig := &eventnotificationsv1.TemplateConfigOneOfAppConfigurationTemplateConfig{
				Body: core.StringPtr(ceappTemplateBody),
			}

			replaceTemplateOptions = &eventnotificationsv1.ReplaceTemplateOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(acTemplateID),
				Name:        core.StringPtr(name),
				Type:        core.StringPtr(templateTypeAC),
				Description: core.StringPtr(description),
				Params:      appconfigTemplateConfig,
			}

			templateResponse, response, err = eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateResponse).ToNot(BeNil())
			Expect(templateResponse.ID).To(Equal(core.StringPtr(acTemplateID)))
			Expect(templateResponse.Name).To(Equal(core.StringPtr(name)))
			Expect(templateResponse.Description).To(Equal(core.StringPtr(description)))

			// end-replace_template
		})

		It(`CreateSubscription request example`, func() {
			fmt.Println("\nCreateSubscription() result:")

			subscriptionName := "FCM subscription"
			// begin-create_subscription
			//FCM
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

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())

			//SMS
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

			//Email
			subscriptionCreateAttributesEmailModel := &eventnotificationsv1.SubscriptionCreateAttributesEmailAttributes{
				Invited:                []string{"tester1@gmail.com", "tester3@ibm.com"},
				AddNotificationPayload: core.BoolPtr(true),
				ReplyToMail:            core.StringPtr("testerreply@gmail.com"),
				ReplyToName:            core.StringPtr("rester_reply"),
				FromName:               core.StringPtr("Test IBM email"),
			}
			subscriptionName = "subscription_email"
			description := core.StringPtr("Subscription for email")
			createSubscriptionOptions = &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr(subscriptionName),
				Description:   description,
				DestinationID: core.StringPtr(destinationID2),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreateAttributesEmailModel,
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createSubscriptionOptions)
			if err != nil {
				panic(err)
			}

			subscriptionID2 = string(*subscription.ID)
			b, _ = json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())

			webSubscriptionCreateAttributesModel := &eventnotificationsv1.SubscriptionCreateAttributes{
				SigningEnabled:         core.BoolPtr(false),
				TemplateIDNotification: core.StringPtr(webhookTemplateID),
			}

			webName := core.StringPtr("subscription_web")
			webDescription := core.StringPtr("Subscription for web")
			createWebSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          webName,
				Description:   webDescription,
				DestinationID: core.StringPtr(destinationID3),
				TopicID:       core.StringPtr(topicID),
				Attributes:    webSubscriptionCreateAttributesModel,
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createWebSubscriptionOptions)

			b, _ = json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID3 = *subscription.ID

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
			b, _ = json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID4 = string(*subscription.ID)

			subscriptionCreateSlackAttributesModel := &eventnotificationsv1.SubscriptionCreateAttributesSlackAttributes{
				AttachmentColor:        core.StringPtr("#0000FF"),
				TemplateIDNotification: core.StringPtr(slackTemplateID),
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
			b, _ = json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID5 = string(*subscription.ID)

			subscriptionCreateAttributesCustomEmailModel := &eventnotificationsv1.SubscriptionCreateAttributesCustomEmailAttributes{
				Invited:                []string{"abc@gmail.com", "tester3@ibm.com"},
				AddNotificationPayload: core.BoolPtr(true),
				ReplyToMail:            core.StringPtr("testerreply@gmail.com"),
				ReplyToName:            core.StringPtr("rester_reply"),
				FromName:               core.StringPtr("Test IBM email"),
				FromEmail:              core.StringPtr("test@abc.event-notifications.test.cloud.ibm.com"),
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

			b, _ = json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID6 = string(*subscription.ID)

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
			subscriptionID7 = *subscription.ID

			slackDirectMessageChannel := &eventnotificationsv1.ChannelCreateAttributes{
				ID: core.StringPtr(slackChannelID),
			}

			subscriptionCreateSlackDMAttributesModel := &eventnotificationsv1.SubscriptionCreateAttributesSlackDirectMessageAttributes{
				Channels:               []eventnotificationsv1.ChannelCreateAttributes{*slackDirectMessageChannel},
				TemplateIDNotification: core.StringPtr(slackTemplateID),
			}

			createSlackDMSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("Slack DM subscription"),
				Description:   core.StringPtr("Subscription for the Slack DM"),
				DestinationID: core.StringPtr(destinationID19),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreateSlackDMAttributesModel,
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createSlackDMSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID8 = string(*subscription.ID)

			subscriptionCreatePDAttributesModel := &eventnotificationsv1.SubscriptionCreateAttributes{
				TemplateIDNotification: core.StringPtr(pdTemplateID),
			}

			pdName := core.StringPtr("subscription_pd")
			pdDescription := core.StringPtr("Subscription for pagerduty")
			createPDSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          pdName,
				Description:   pdDescription,
				DestinationID: core.StringPtr(destinationID10),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreatePDAttributesModel,
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createPDSubscriptionOptions)

			b, _ = json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID9 = *subscription.ID

			subscriptionCreateEventStreamsAttributesModel := &eventnotificationsv1.SubscriptionCreateAttributesEventstreamsAttributes{
				TemplateIDNotification: core.StringPtr(esTemplateID),
			}

			createEventSTreamsSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("Event Streams subscription"),
				Description:   core.StringPtr("Subscription for the Event Streams Destination"),
				DestinationID: core.StringPtr(destinationID20),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreateEventStreamsAttributesModel,
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createEventSTreamsSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID10 = string(*subscription.ID)

			subscriptionCreateCEAttributesModel := &eventnotificationsv1.SubscriptionCreateAttributes{
				TemplateIDNotification: core.StringPtr(cejobTemplateID),
			}

			ceName := core.StringPtr("subscription_code_engine_job")
			ceDescription := core.StringPtr("Subscription for code engine job")
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
			subscriptionID11 = *ceJobSubscription.ID

			subscriptionCreateCEAttributesModel = &eventnotificationsv1.SubscriptionCreateAttributes{
				TemplateIDNotification: core.StringPtr(ceappTemplateID),
			}

			ceName = core.StringPtr("subscription_code_engine_app")
			ceDescription = core.StringPtr("Subscription for code engine application")
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
			subscriptionID12 = *ceSubscription.ID

			subscriptionCreateAppConfigAttributesModel := &eventnotificationsv1.SubscriptionCreateAttributesAppConfigurationAttributes{
				TemplateIDNotification: core.StringPtr(acTemplateID),
			}

			createAppConfigSubscriptionOptions := &eventnotificationsv1.CreateSubscriptionOptions{
				InstanceID:    core.StringPtr(instanceID),
				Name:          core.StringPtr("App Config subscription"),
				Description:   core.StringPtr("Subscription for the App Config Destination"),
				DestinationID: core.StringPtr(destinationID21),
				TopicID:       core.StringPtr(topicID),
				Attributes:    subscriptionCreateAppConfigAttributesModel,
			}

			subscription, response, err = eventNotificationsService.CreateSubscription(createAppConfigSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subscription).ToNot(BeNil())
			subscriptionID13 = string(*subscription.ID)

			// end-create_subscription

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
		It(`ListTemplates request example`, func() {
			fmt.Println("\nListTemplates() result:")
			// begin-list_templates

			listTemplatesOptions := eventNotificationsService.NewListTemplatesOptions(
				instanceID,
			)

			templatesList, response, err := eventNotificationsService.ListTemplates(listTemplatesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templatesList, "", "  ")
			fmt.Println(string(b))

			// end-list_templates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templatesList).ToNot(BeNil())

		})
		It(`ListPreDefinedTemplates request example`, func() {
			fmt.Println("\nListPreDefinedTemplates() result:")
			// begin-list_pre_defined_templates

			listpredefinedtemplatesOptions := &eventnotificationsv1.ListPreDefinedTemplatesOptions{
				InstanceID: core.StringPtr(instanceID),
				Source:     core.StringPtr("logs"),
				Type:       core.StringPtr("slack.notification"),
				Offset:     core.Int64Ptr(int64(0)),
				Limit:      core.Int64Ptr(int64(1)),
				Search:     core.StringPtr(search),
			}

			predefinedtemplatesList, response, err := eventNotificationsService.ListPreDefinedTemplates(listpredefinedtemplatesOptions)

			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(predefinedtemplatesList, "", "  ")
			fmt.Println(string(b))

			// end-list_pre_defined_templates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(predefinedtemplatesList).ToNot(BeNil())
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

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())

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
			if err != nil {
				panic(err)
			}

			b, _ = json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))
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
			name := core.StringPtr("subscription_email")
			description := core.StringPtr("Subscription for email")
			updateSubscriptionOptions = &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        name,
				Description: description,
				ID:          core.StringPtr(subscriptionID2),
				Attributes:  subscriptionUpdateEmailAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateSubscriptionOptions)

			if err != nil {
				panic(err)
			}

			b, _ = json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID2)))
			Expect(subscription.Name).To(Equal(name))
			Expect(subscription.Description).To(Equal(description))

			webSubscriptionUpdateAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesWebhookAttributes{
				SigningEnabled:         core.BoolPtr(true),
				TemplateIDNotification: core.StringPtr(webhookTemplateID),
			}

			webName := core.StringPtr("Webhook_sub_updated")
			webDescription := core.StringPtr("Update Webhook subscription")
			webUpdateSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(subscriptionID3),
				Name:        webName,
				Description: webDescription,
				Attributes:  webSubscriptionUpdateAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(webUpdateSubscriptionOptions)

			if err != nil {
				panic(err)
			}

			b, _ = json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID3)))
			Expect(subscription.Name).To(Equal(webName))
			Expect(subscription.Description).To(Equal(webDescription))

			serviceNowName := core.StringPtr("subscription_Service_Now_update")
			serviceNowDescription := core.StringPtr("Subscription update for Service_Now")
			updateServiceNowSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        serviceNowName,
				Description: serviceNowDescription,
				ID:          core.StringPtr(subscriptionID4),
				Attributes: &eventnotificationsv1.SubscriptionUpdateAttributesServiceNowAttributes{
					AssignedTo:      core.StringPtr("user"),
					AssignmentGroup: core.StringPtr("test"),
				},
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateServiceNowSubscriptionOptions)
			if err != nil {
				panic(err)
			}

			b, _ = json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID4)))
			Expect(subscription.Name).To(Equal(serviceNowName))
			Expect(subscription.Description).To(Equal(serviceNowDescription))

			subscriptionUpdateSlackAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesSlackAttributes{
				AttachmentColor:        core.StringPtr("#0000FF"),
				TemplateIDNotification: core.StringPtr(slackTemplateID),
			}

			slackName := core.StringPtr("subscription_slack_update")
			slackDescription := core.StringPtr("Subscription update for slack")
			updateSlackSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        slackName,
				Description: slackDescription,
				ID:          core.StringPtr(subscriptionID5),
				Attributes:  subscriptionUpdateSlackAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateSlackSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID5)))
			Expect(subscription.Name).To(Equal(slackName))
			Expect(subscription.Description).To(Equal(slackDescription))

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
				FromEmail:              core.StringPtr("test@abc.event-notifications.test.cloud.ibm.com"),
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
				ID:          core.StringPtr(subscriptionID6),
				Attributes:  subscriptionUpdateCustomEmailAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateSubscriptionOptions)
			if err != nil {
				panic(err)
			}
			b, _ = json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID6)))
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
				ID:          core.StringPtr(subscriptionID7),
				Attributes:  subscriptionUpdateCustomSMSAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID7)))
			Expect(subscription.Name).To(Equal(customSMSName))
			Expect(subscription.Description).To(Equal(customSMSDescription))

			slackDirectMessageChannel := &eventnotificationsv1.ChannelUpdateAttributes{
				ID:        core.StringPtr(slackChannelID),
				Operation: core.StringPtr("add"),
			}

			subscriptionUpdateSlackDMAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesSlackDirectMessageUpdateAttributes{
				Channels:               []eventnotificationsv1.ChannelUpdateAttributes{*slackDirectMessageChannel},
				TemplateIDNotification: core.StringPtr(slackTemplateID),
			}

			slackDMName := core.StringPtr("subscription_slack_DM_update")
			slackDMDescription := core.StringPtr("Subscription update for slack DM")
			updateSlackDMSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        slackDMName,
				Description: slackDMDescription,
				ID:          core.StringPtr(subscriptionID8),
				Attributes:  subscriptionUpdateSlackDMAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateSlackDMSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID8)))
			Expect(subscription.Name).To(Equal(slackDMName))
			Expect(subscription.Description).To(Equal(slackDMDescription))

			subscriptionUpdatePagerDutyAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesPagerDutyAttributes{
				TemplateIDNotification: core.StringPtr(webhookTemplateID),
			}

			pdName := core.StringPtr("subscription_Pager_Duty_update")
			pdDescription := core.StringPtr("Subscription update for Pager Duty")
			updatePDSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(subscriptionID9),
				Name:        pdName,
				Description: pdDescription,
				Attributes:  subscriptionUpdatePagerDutyAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updatePDSubscriptionOptions)

			if err != nil {
				panic(err)
			}

			b, _ = json.MarshalIndent(subscription, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID9)))
			Expect(subscription.Name).To(Equal(pdName))
			Expect(subscription.Description).To(Equal(pdDescription))

			subscriptionUpdateEventStreamsAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesEventstreamsAttributes{
				TemplateIDNotification: core.StringPtr(esTemplateID),
			}

			eventstreamsName := core.StringPtr("subscription_Event_Streams_update")
			eventstreamsDescription := core.StringPtr("Subscription update for Event Streams")
			updateEventStreamsSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        eventstreamsName,
				Description: eventstreamsDescription,
				ID:          core.StringPtr(subscriptionID10),
				Attributes:  subscriptionUpdateEventStreamsAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateEventStreamsSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID10)))
			Expect(subscription.Name).To(Equal(eventstreamsName))
			Expect(subscription.Description).To(Equal(eventstreamsDescription))

			subscriptionCEJobUpdateAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesCodeEngineAttributes{
				TemplateIDNotification: core.StringPtr(cejobTemplateID),
			}

			ceName := core.StringPtr("code_engine_job_sub_updated")
			ceDescription := core.StringPtr("Update code engine job subscription")
			updateCEJobSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(subscriptionID11),
				Name:        ceName,
				Description: ceDescription,
				Attributes:  subscriptionCEJobUpdateAttributesModel,
			}

			ceJobSubscription, response, err := eventNotificationsService.UpdateSubscription(updateCEJobSubscriptionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ceJobSubscription).ToNot(BeNil())
			Expect(ceJobSubscription.ID).To(Equal(core.StringPtr(subscriptionID11)))
			Expect(ceJobSubscription.Name).To(Equal(ceName))
			Expect(ceJobSubscription.Description).To(Equal(ceDescription))

			subscriptionCEUpdateAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesCodeEngineAttributes{
				TemplateIDNotification: core.StringPtr(ceappTemplateID),
			}

			ceName = core.StringPtr("code_engine_sub_updated")
			ceDescription = core.StringPtr("Update code engine subscription")
			updateCESubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(subscriptionID12),
				Name:        ceName,
				Description: ceDescription,
				Attributes:  subscriptionCEUpdateAttributesModel,
			}

			ceSubscription, response, err := eventNotificationsService.UpdateSubscription(updateCESubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ceSubscription).ToNot(BeNil())
			Expect(ceSubscription.ID).To(Equal(core.StringPtr(subscriptionID12)))
			Expect(ceSubscription.Name).To(Equal(ceName))
			Expect(ceSubscription.Description).To(Equal(ceDescription))

			subscriptionUpdateAppConfigAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesAppConfigurationAttributes{
				TemplateIDNotification: core.StringPtr(esTemplateID),
			}

			appConfigName := core.StringPtr("subscription_App_Config_update")
			appConfigDescription := core.StringPtr("Subscription update for App Configuration")
			updateAppConfigSubscriptionOptions := &eventnotificationsv1.UpdateSubscriptionOptions{
				InstanceID:  core.StringPtr(instanceID),
				Name:        appConfigName,
				Description: appConfigDescription,
				ID:          core.StringPtr(subscriptionID13),
				Attributes:  subscriptionUpdateAppConfigAttributesModel,
			}

			subscription, response, err = eventNotificationsService.UpdateSubscription(updateAppConfigSubscriptionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subscription).ToNot(BeNil())
			Expect(subscription.ID).To(Equal(core.StringPtr(subscriptionID13)))
			Expect(subscription.Name).To(Equal(appConfigName))
			Expect(subscription.Description).To(Equal(appConfigDescription))

			// end-update_subscription

		})

		It(`GetEnabledCountries request example`, func() {
			fmt.Println("\nGetEnabledCountries() result:")
			// begin-get_enabled_countries
			getEnabledCountriesOptions := &eventnotificationsv1.GetEnabledCountriesOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(destinationID17),
			}

			enabledCountries, response, err := eventNotificationsService.GetEnabledCountries(getEnabledCountriesOptions)
			// end-get_enabled_countries

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(enabledCountries).ToNot(BeNil())
		})

		It(`SendNotifications request example`, func() {
			fmt.Println("\nSendNotifications() result:")

			notificationID := "1234-1234-sdfs-234"
			notificationSeverity := "LOW"
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

			notificationDevicesModel := "{\"platforms\":[\"push_ios\",\"push_android\",\"push_chrome\",\"push_firefox\",\"push_huawei\"]}"
			notificationSafariBodyModel := "{\"en_data\": {\"alert\": \"Alert message\"}}"
			mailTo := "[\"abc@ibm.com\", \"def@us.ibm.com\"]"
			templates := "[\"149b0e11-8a7c-4fda-a847-5d79e01b71dc\"]"
			smsTo := "[\"+911234567890\", \"+911224567890\"]"
			slackTo := "[\"C07FALXBH4G\"]"
			mms := "{\"url\": \"https://cloud.ibm.com/avatar/v1/avatar/migrationsegment/logo_ibm.png\"}"
			htmlBody := "\"Hi  ,<br/>Certificate expiring in 90 days.<br/><br/>Please login to <a href=\"https: //cloud.ibm.com/security-compliance/dashboard\">Security and Complaince dashboard</a> to find more information<br/>\""
			markdown := "**Event Summary** \n\n**Toolchain ID:** `4414af34-a5c7-47d3-8f05-add4af6d78a6`  \n**Content Type:** `application/json`\n\n---\n\n *Pipeline Run Details*\n\n- **Namespace:** `PR`\n- **Trigger Name:** `manual`\n- **Triggered By:** `nitish.kulkarni3@ibm.com`\n- **Build Number:** `343`\n- **Pipeline Link:** [View Pipeline Run](https://cloud.ibm.com/devops/pipelines/tekton/e9cd5aa3-a3f2-4776-8acc-26a35922386e/runs/f29ac6f5-bd2f-4a26-abb8-4249be8dbab7?env_id=ibm:yp:us-south)"

			notificationCreateModel.Ibmenpushto = &notificationDevicesModel

			apnsOptions := map[string]interface{}{
				"aps": map[string]interface{}{
					"alert": "APNS alert",
					"badge": 5,
				},
			}

			ibmenapnsbody, _ := json.Marshal(apnsOptions)
			ibmenapnsbodyString := string(ibmenapnsbody)

			fcmOptions := map[string]interface{}{
				"notification": map[string]interface{}{
					"title": "FCM alert",
					"body":  "alert message for FCM",
				},
			}
			ibmenfcmbody, _ := json.Marshal(fcmOptions)
			ibmenfcmbodyString := string(ibmenfcmbody)

			apnsHeaders := map[string]interface{}{
				"apns-collapse-id": "collapse-id",
			}
			ibmenapnsheaderbody, _ := json.Marshal(apnsHeaders)
			ibmenapnsheaderstring := string(ibmenapnsheaderbody)
			notificationHuaweiBodyModel := "{\"message\": {\"android\": {\"notification\": {\"title\": \"Breaking News\",\"body\": \"New news story available.\"},\"data\": {\"name\": \"Willie Greenholt\",\"description\": \"description\"}}}}"

			notificationCreateModel.Ibmenfcmbody = &ibmenfcmbodyString
			notificationCreateModel.Ibmenapnsbody = &ibmenapnsbodyString
			notificationCreateModel.Ibmenapnsheaders = &ibmenapnsheaderstring
			notificationCreateModel.Ibmensafaribody = &notificationSafariBodyModel
			notificationCreateModel.Ibmenhuaweibody = &notificationHuaweiBodyModel
			notificationCreateModel.Ibmenmailto = &mailTo
			notificationCreateModel.Ibmentemplates = &templates
			notificationCreateModel.Ibmensmsto = &smsTo
			notificationCreateModel.Ibmenslackto = &slackTo
			notificationCreateModel.Ibmenmms = &mms
			notificationCreateModel.Ibmensubject = core.StringPtr("Notification subject")
			notificationCreateModel.Ibmenhtmlbody = core.StringPtr(htmlBody)
			notificationCreateModel.Ibmenmarkdown = core.StringPtr(markdown)
			notificationCreateModel.Ibmendefaultshort = core.StringPtr("This is simple test alert from IBM Cloud Event Notifications service.")
			notificationCreateModel.Ibmendefaultlong = core.StringPtr("Hi, we are making sure from our side that the service is available for consumption.")

			sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
			sendNotificationsOptionsModel.InstanceID = &instanceID
			sendNotificationsOptionsModel.Body = notificationCreateModel

			notificationResponse, response, err := eventNotificationsService.SendNotifications(sendNotificationsOptionsModel)
			notificationID = *notificationResponse.NotificationID

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

		It(`GetMetrics(getMetricsOptions *GetMetricsOptions)`, func() {
			// begin-metrics
			getMetricsOptions := &eventnotificationsv1.GetMetricsOptions{
				InstanceID:      core.StringPtr(instanceID),
				DestinationType: core.StringPtr("smtp_custom"),
				Gte:             core.StringPtr("2024-08-01T17:18:43Z"),
				Lte:             core.StringPtr("2024-08-02T11:55:22Z"),
				EmailTo:         core.StringPtr("mobileb@us.ibm.com"),
				DestinationID:   core.StringPtr(destinationID16),
				NotificationID:  core.StringPtr(notificationID),
				Subject:         core.StringPtr("Test Metrics Subject"),
			}

			metrics, response, err := eventNotificationsService.GetMetrics(getMetricsOptions)
			// end-metrics
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(metrics).ToNot(BeNil())
		})

		It(`CreateSMTPConfiguration request example`, func() {
			// begin-create_smtp_configuration
			name := "SMTP configuration"
			description := "SMTP configuration description"
			domain := "mailx.event-notifications.test.cloud.ibm.com"

			createSMTPConfigurationOptions := &eventnotificationsv1.CreateSMTPConfigurationOptions{
				InstanceID:  core.StringPtr(instanceID),
				Domain:      core.StringPtr(domain),
				Description: core.StringPtr(description),
				Name:        core.StringPtr(name),
			}

			smtpConfig, response, err := eventNotificationsService.CreateSMTPConfiguration(createSMTPConfigurationOptions)
			// end-create_smtp_configuration
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(smtpConfig).ToNot(BeNil())
			Expect(smtpConfig.Name).To(Equal(core.StringPtr(name)))
			Expect(smtpConfig.Description).To(Equal(core.StringPtr(description)))
			Expect(smtpConfig.Domain).To(Equal(core.StringPtr(domain)))
			smtpConfigID = *smtpConfig.ID
		})

		It(`UpdateVerifySMTP request example`, func() {
			// begin-update_verify_smtp
			updateVerifySMTPOptions := &eventnotificationsv1.UpdateVerifySMTPOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(smtpConfigID),
				Type:       core.StringPtr("dkim,spf,en_authorization"),
			}

			verifySMTP, response, err := eventNotificationsService.UpdateVerifySMTP(updateVerifySMTPOptions)
			// end-update_verify_smtp
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(verifySMTP.Status).ToNot(BeNil())
		})

		It(`CreateSMTPUser request example`, func() {
			// begin-create_smtp_user
			description := "smtp user description"
			createSMTPUserOptions := &eventnotificationsv1.CreateSMTPUserOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(smtpConfigID),
				Description: core.StringPtr(description),
			}

			user, response, err := eventNotificationsService.CreateSMTPUser(createSMTPUserOptions)
			// end-create_smtp_user
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(user.Domain).ToNot(BeNil())
			Expect(user.Username).ToNot(BeNil())
			Expect(user.Password).ToNot(BeNil())
			Expect(user.SMTPConfigID).ToNot(BeNil())
			Expect(user.Description).To(Equal(core.StringPtr(description)))
			smtpUserID = *user.ID
		})

		It(`CreateCloneSMTPUser(createSMTPUserOptions *CreateSMTPUserOptions)`, func() {
			// begin-clone_smtp_user
			description := "clone smtp user"
			createSMTPUserOptions := &eventnotificationsv1.CreateSMTPUserOptions{
				InstanceID:      core.StringPtr(instanceID),
				ID:              core.StringPtr(smtpcloneconfigID),
				Description:     core.StringPtr(description),
				UsernameToClone: core.StringPtr(smtpUserID),
			}

			user, response, err := eventNotificationsService.CreateSMTPUser(createSMTPUserOptions)
			// end-clone_smtp_user
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(user.Domain).ToNot(BeNil())
			Expect(user.Username).ToNot(BeNil())
			Expect(user.Password).ToNot(BeNil())
			Expect(user.SMTPConfigID).ToNot(BeNil())
			Expect(user.Description).To(Equal(core.StringPtr(description)))
			smtpUserID = *user.ID
		})

		It(`ListSMTPConfigurations request example`, func() {
			// begin-list_smtp_configurations
			listSMTPConfigurationsOptions := &eventnotificationsv1.ListSMTPConfigurationsOptions{
				InstanceID: core.StringPtr(instanceID),
				Limit:      core.Int64Ptr(int64(1)),
				Offset:     core.Int64Ptr(int64(0)),
				Search:     core.StringPtr(search),
			}

			smtpConfigurations, response, err := eventNotificationsService.ListSMTPConfigurations(listSMTPConfigurationsOptions)
			// end-list_smtp_configurations
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(smtpConfigurations.TotalCount).To(Equal(core.Int64Ptr(int64(1))))
		})

		It(`ListSMTPUsers example request`, func() {
			// begin-list_smtp_users
			listSMTPUsersOptions := &eventnotificationsv1.ListSMTPUsersOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(smtpConfigID),
				Limit:      core.Int64Ptr(int64(1)),
				Offset:     core.Int64Ptr(int64(0)),
				Search:     core.StringPtr(search),
			}

			smtpUsers, response, err := eventNotificationsService.ListSMTPUsers(listSMTPUsersOptions)
			// end-list_smtp_users
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(smtpUsers.TotalCount).To(Equal(core.Int64Ptr(int64(1))))
		})

		It(`GetSMTPConfiguration example request`, func() {
			// begin-get_smtp_configuration
			getSMTPconfigurationOptions := &eventnotificationsv1.GetSMTPConfigurationOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(smtpConfigID),
			}

			smtpConfiguration, response, err := eventNotificationsService.GetSMTPConfiguration(getSMTPconfigurationOptions)
			// end-get_smtp_configuration
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(smtpConfiguration.Domain).ToNot(BeNil())
			Expect(smtpConfiguration.Name).ToNot(BeNil())
			Expect(smtpConfiguration.Description).ToNot(BeNil())
		})

		It(`GetSMTPAllowedIps example request`, func() {
			// begin-get_smtp_allowed_ips
			getSMTPAllowedIPsOptions := &eventnotificationsv1.GetSMTPAllowedIpsOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(smtpConfigID),
			}

			smtpAllowedIPs, response, err := eventNotificationsService.GetSMTPAllowedIps(getSMTPAllowedIPsOptions)
			// end-get_smtp_allowed_ips
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(smtpAllowedIPs.Subnets[0]).ToNot(BeNil())
		})

		It(`GetSMTPUser example request`, func() {
			// begin-get_smtp_user
			getSMTPUserOptions := &eventnotificationsv1.GetSMTPUserOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(smtpConfigID),
				UserID:     core.StringPtr(smtpUserID),
			}

			SMTPUser, response, err := eventNotificationsService.GetSMTPUser(getSMTPUserOptions)
			// end-get_smtp_user
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(SMTPUser.Domain).ToNot(BeNil())
			Expect(SMTPUser.Description).ToNot(BeNil())
			Expect(SMTPUser.Username).ToNot(BeNil())
		})

		It(`UpdateSMTPConfiguration example request`, func() {
			// begin-update_smtp_configuration
			name := "SMTP configuration name update"
			description := "SMTP configuration description update"

			updateSMTPConfigurationOptions := &eventnotificationsv1.UpdateSMTPConfigurationOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(smtpConfigID),
				Name:        core.StringPtr(name),
				Description: core.StringPtr(description),
			}

			updateSMTPConfiguration, response, err := eventNotificationsService.UpdateSMTPConfiguration(updateSMTPConfigurationOptions)
			// end-update_smtp_configuration
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateSMTPConfiguration.Name).To(Equal(core.StringPtr(name)))
			Expect(updateSMTPConfiguration.Description).To(Equal(core.StringPtr(description)))
		})

		It(`UpdateSMTPUser request example`, func() {
			// begin-update_smtp_user
			description := "SMTP user description update"

			updateSMTPUserOptions := &eventnotificationsv1.UpdateSMTPUserOptions{
				InstanceID:  core.StringPtr(instanceID),
				ID:          core.StringPtr(smtpConfigID),
				Description: core.StringPtr(description),
				UserID:      core.StringPtr(smtpUserID),
			}

			updateSMTPUser, response, err := eventNotificationsService.UpdateSMTPUser(updateSMTPUserOptions)
			// end-update_smtp_user
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateSMTPUser.Description).To(Equal(core.StringPtr(description)))
		})

		It(`DeleteTemplate request example`, func() {
			// begin-delete_template
			for _, ID := range []string{templateInvitationID, templateNotificationID, slackTemplateID} {

				deleteTemplateOptions := &eventnotificationsv1.DeleteTemplateOptions{
					InstanceID: core.StringPtr(instanceID),
					ID:         core.StringPtr(ID),
				}

				response, err := eventNotificationsService.DeleteTemplate(deleteTemplateOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(204))
			}
			// end-delete_template
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
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteSubscription(): %d\n", response.StatusCode)
			}

			for _, ID := range []string{subscriptionID1, subscriptionID2, subscriptionID3, subscriptionID4, subscriptionID5, subscriptionID6, subscriptionID7, subscriptionID8, subscriptionID9, subscriptionID10, subscriptionID11} {

				deleteSubscriptionOptions := &eventnotificationsv1.DeleteSubscriptionOptions{
					InstanceID: core.StringPtr(instanceID),
					ID:         core.StringPtr(ID),
				}

				response, err := eventNotificationsService.DeleteSubscription(deleteSubscriptionOptions)
				if err != nil {
					panic(err)
				}
				if response.StatusCode != 204 {
					fmt.Printf("\nUnexpected response status code received from DeleteSubscription(): %d\n", response.StatusCode)
				}
			}

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteSMTPUser request example`, func() {

			for _, ID := range []string{smtpUserID} {
				// begin-delete_smtp_user
				deleteSMTPUserOptions := &eventnotificationsv1.DeleteSMTPUserOptions{
					InstanceID: core.StringPtr(instanceID),
					ID:         core.StringPtr(smtpConfigID),
					UserID:     core.StringPtr(ID),
				}

				response, err := eventNotificationsService.DeleteSMTPUser(deleteSMTPUserOptions)
				// end-delete_smtp_user
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(204))
			}
		})

		It(`DeleteSMTPConfiguration request example)`, func() {

			for _, ID := range []string{smtpConfigID} {
				// begin-delete_smtp_configuration
				deleteSMTPConfigurationOptions := &eventnotificationsv1.DeleteSMTPConfigurationOptions{
					InstanceID: core.StringPtr(instanceID),
					ID:         core.StringPtr(ID),
				}

				response, err := eventNotificationsService.DeleteSMTPConfiguration(deleteSMTPConfigurationOptions)
				// end-delete_smtp_configuration

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(204))
			}
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

			for _, ID := range []string{destinationID3, destinationID4, destinationID5, destinationID6, destinationID8, destinationID9, destinationID10, destinationID11, destinationID12, destinationID13, destinationID14, destinationID15, destinationID16, destinationID17, destinationID18, destinationID19} {
				deleteDestinationOptions := &eventnotificationsv1.DeleteDestinationOptions{
					InstanceID: core.StringPtr(instanceID),
					ID:         core.StringPtr(ID),
				}

				response, err := eventNotificationsService.DeleteDestination(deleteDestinationOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(204))
			}

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
