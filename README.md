# IBM Cloud Event Notifications Go Admin SDK 0.17.0

Go client library to interact with the various [IBM Cloud Event Notifications APIs](https://cloud.ibm.com/apidocs?category=event-notifications).

## Table of Contents

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [Go modules](#go-modules)
  - [`go get` command](#go-get-command)
- [Initialize SDK](#initialize-sdk)
- [Using the SDK](#using-the-sdk)
- [Set Environment](#set-environment)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Event Notifications Go SDK allows developers to programmatically interact with Event Notifications service in IBM cloud.

| Service Name                                                                     | Package name         |
| -------------------------------------------------------------------------------- | -------------------- |
| <!-- [Example Service](https://cloud.ibm.com/apidocs/example-service)            | exampleservicev1 --> |
| [Event Notifications Service](https://cloud.ibm.com/apidocs/event-notifications) | eventnotificationsv1 |

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

- An [IBM Cloud][ibm-cloud-onboarding] account.
- An Event Notifications Instance
- Go version 1.18 or above.

## Installation

Install using the command.

```
go get -u github.com/IBM/event-notifications-go-admin-sdk
```

### Go modules

If your application uses Go modules for dependency management (recommended), just add an import for each service
that you will use in your application.  
Here is an example:

```go
import (
	"github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
)
```

Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's
`go.mod` file.

In the example above, the `eventnotificationsv1` part of the import path is the package name
associated with the Example Service service.
See the service table above to find the approprate package name for the services used by your application.

## Initialize SDK

Initialize the sdk to connect with your Event Notifications service instance.

```go
func initInstance() *eventnotificationsv1.EventNotificationsV1 {

    // IAM API key based authentication
	authenticator := &core.IamAuthenticator{
		ApiKey: <apikey>, // Event notifications service instance APIKey
	}

	// Set the options for the Event notification instance.
	options := &eventnotificationsv1.EventNotificationsV1Options{
		Authenticator: authenticator,
		URL:           "https://" + region + ".event-notifications.cloud.ibm.com/event-notifications",
	}
	eventNotificationsService, err := eventnotificationsv1.NewEventNotificationsV1(options)
	if err != nil {
		panic(err)
	}
	return eventNotificationsService

}
```

To configure service URL for Private Endpoint

If you enabled service endpoints in your account, you can send API requests over the IBM Cloud private network. In the initialisation, the base endpoint URLs of the IAM(authenticator) & Event Notification(service) should be modified to point to private endpoints.

1. Setting client options programmatically

```go
	authenticator := &core.IamAuthenticator{
		ApiKey: "<iam-api-key>",
		URL: "https://private.iam.cloud.ibm.com",
	}

	options := &eventnotificationsv1.EventNotificationsV1Options{
		Authenticator: authenticator,
		URL:           "https://private." + region + ".event-notifications.cloud.ibm.com/event-notifications",
	}
```

2. Using external configuration properties

```go
   EVENT_NOTIFICATIONS_AUTH_URL = https://private.iam.cloud.ibm.com/identity/token
```

- region : Region of the Event Notifications Instance

## Using the SDK

SDK Methods to consume

- [Source](#source)
  - [Create Source](#create-source)
  - [List Sources](#list-sources)
  - [Get Source](#get-sources)
  - [Update Source](#update-source)
  - [Delete Source](#delete-source)
- [Topics](#topics)
  - [Create Topics](#create-topic)
  - [List Topics](#list-topic)
  - [Get Topic](#get-topic)
  - [Update Topics](#update-topic)
  - [Delete Topics](#delete-topic)
- [Destinations](#destinations)
  - [Create Destination](#create-destination)
  - [List Destinations](#list-destinations)
  - [Get Destination](#get-destination)
  - [Update Destination](#update-destination)
  - [Delete Destination](#delete-destination)
  - [Custom Domain_Name_verification](#custom-domain-name-verification)
  - [Test Destination](#test-destination)
- [Templates](#templates)
  - [Create Template](#create-template)
  - [List Templates](#list-templates)
  - [Get Template](#get-template)
  - [Update Template](#update-template)
  - [Delete Template](#delete-template)
  - [List Predefined Templates](#list-predefined-templates)
  - [Get Predefined Template](#get-predefined-template)
- [Push Destination APIs](#push-destination-apis)
  - [Create Destination tag subscription](#create-destination-tag-subscription)
  - [List Destination tag subscription](#list-destination-tag-subscription)
  - [Delete Destination device tag subscription](#delete-destination-device-tag-subscription)
- [Subscriptions](#subscriptions)
  - [Create Subscription](#create-subscription)
  - [List Subscriptions](#list-subscriptions)
  - [Get Subscription](#get-subscription)
  - [Update Subscription](#update-subscription)
  - [Delete Subscription](#delete-subscription)
- [Integration](#integration)
  - [Create Integration](#create-integration)
  - [Get Integration](#get-integration)
  - [List Integrations](#list-integrations)
  - [Update Integration](#update-integration)
- [SMTP Configurations](#SMTPConfigurations)
  - [Create SMTP Configuration](#create-smtp-configuration)
  - [Create SMTP User](#create-smtp-user)
  - [Get SMTP Configuration](#get-smtp-configuration)
  - [Get SMTP User](#get-smtp-user)
  - [Get SMTP Allowed Ips](#get-smtp-allowed-ips)
  - [List SMTP Configurations](#list-smtp-configurations)
  - [List SMTP Users](#list-smtp-users)
  - [Update SMTP Configuration](#update-smtp-configuration)
  - [Update SMTP User](#update-smtp-user)
  - [Delete SMTP User](#delete-smtp-user)
  - [Delete SMTP Configuration](#delete-smtp-user)
  - [Verify SMTP](#verify-smtp)
- [Metrics](#Metrics)
  - [Get Metrics](#get-metrics)
- [Send Notifications](#send-notifications)

## Source

### Create Source

```go
createSourcesOptions := eventNotificationsService.NewCreateSourcesOptions(
	    <instance-id>, // Event notifications service instance GUID
		<source-name>,
		<source-description>,
	)
createSourcesOptions.SetEnabled(false)

sourceResponse, response, err := eventNotificationsService.CreateSources(createSourcesOptions)
```

### List Sources

```go
listSourcesOptions := eventNotificationsService.NewListSourcesOptions(
	<instance-id>, // Event notifications service instance GUID
)

sourceList, response, err := eventNotificationsService.ListSource(listSourcesOptions)

if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(sourceList, "", "  ")
fmt.Println(string(b))
```

### Get Sources

```go
getSourceOptions := eventNotificationsService.NewGetSourceOptions(
	<instance-id>, // Event notifications service instance GUID
	<source-id>,   // Event notifications service instance Source ID
)

source, response, err := eventNotificationsService.GetSource(getSourceOptions)

if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(source, "", "  ")
fmt.Println(string(b))
```

### Update Source

```go
updateSourceOptions := eventNotificationsService.NewUpdateSourceOptions(
		<instance-id>, // Event notifications service instance GUID
	    <source-id>,   // Event notifications service instance Source ID
	)
updateSourceOptions.SetName(*core.StringPtr(<source-updated-name>))
updateSourceOptions.SetDescription(*core.StringPtr(<source-updated-description>))
updateSourceOptions.SetEnabled(true)

source, response, err := eventNotificationsService.UpdateSource(updateSourceOptions)
```

### Delete Source

```go
deleteSourceOptions := eventNotificationsService.NewDeleteSourceOptions(
	<instance-id>, // Event notifications service instance GUID
	<source-id>,   // Event notifications service instance Source ID
)

response, err := eventNotificationsService.DeleteSource(deleteSourceOptions)
```

## Topics

### Create Topic

```go

// Filters applied in case of periodic-timer as source. EventTypeFilter, NotificationFilter are mutually exclusive with EventScheduleFilter

eventScheduleFilterAttributesModel := new(eventnotificationsv1.EventScheduleFilterAttributes)
eventScheduleFilterAttributesModel.StartsAt = CreateMockDateTime("2024-12-20T05:15:00.000Z")
eventScheduleFilterAttributesModel.EndsAt = CreateMockDateTime("2024-12-20T20:30:00.000Z")
eventScheduleFilterAttributesModel.Expression = core.StringPtr("* * * * *")

rulesModel = &eventnotificationsv1.Rules{
	Enabled:             core.BoolPtr(true),
	EventScheduleFilter: eventScheduleFilterAttributesModel,
}

rulesModel := &eventnotificationsv1.Rules{
	Enabled:            core.BoolPtr(false),
	EventTypeFilter:    core.StringPtr("$.notification_event_info.event_type == 'cert_manager'"), // Add your event type filter here.
	NotificationFilter: core.StringPtr("$.notification.findings[0].severity == 'MODERATE'"), // Add your notification filter here.
}

topicUpdateSourcesItemModel := &eventnotificationsv1.TopicUpdateSourcesItem{
	ID:    core.StringPtr(<source-id>),
	Rules: []eventnotificationsv1.Rules{*rulesModel},
}

createTopicOptions := &eventnotificationsv1.CreateTopicOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	Name:        core.StringPtr(<topic-name>]),
	Description: core.StringPtr(<topic-description>),
	Sources:     []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel},
}

topic, response, err := eventNotificationsService.CreateTopic(createTopicOptions)

if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(topic, "", "  ")
fmt.Println(string(b))
```

### List Topics

```go
listTopicsOptions := eventNotificationsService.NewListTopicsOptions(
	<instance-id>,
)

topicList, response, err := eventNotificationsService.ListTopic(listTopicsOptions)

if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(topicList, "", "  ")
fmt.Println(string(b))
```

### Get Topic

```go
getTopicOptions := eventNotificationsService.NewGetTopicOptions(
	<instance-id>, // Event notifications service instance GUID
	<topic-id>, // Event notifications service instance Topic ID
)

topic, response, err := eventNotificationsService.GetTopic(getTopicOptions)

if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(topic, "", "  ")
fmt.Println(string(b))
```

### Update Topic

```go

rulesModel := &eventnotificationsv1.Rules{
	Enabled:            core.BoolPtr(true),
	EventTypeFilter:    core.StringPtr("$.notification_event_info.event_type == 'core_cert_manager'"), // Add your event type filter here.
	NotificationFilter: core.StringPtr("$.notification.findings[0].severity == 'SEVERE'"), // Add your notification filter here.
}

topicUpdateSourcesItemModel := &eventnotificationsv1.TopicUpdateSourcesItem{
	ID:    core.StringPtr(<source-id>),  // Event notifications service instance Source ID
	Rules: []eventnotificationsv1.Rules{*rulesModel},
}

replaceTopicOptions := &eventnotificationsv1.ReplaceTopicOptions{
	InstanceID:  core.StringPtr(<instance-id>), // Event notifications service instance GUID
	ID:          core.StringPtr(<topic-id>),    // Event notifications service instance Topic ID
	Name:        core.StringPtr(<topic-update-name>),  // Event notifications service instance Topic Name
	Description: core.StringPtr(<topic-update-description>), // Event notifications service instance Topic description
	Sources:     []eventnotificationsv1.TopicUpdateSourcesItem{*topicUpdateSourcesItemModel},
}

topic, response, err := eventNotificationsInstance.ReplaceTopic(replaceTopicOptions)

if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(topic, "", "  ")
fmt.Println(string(b))

```

### Delete Topic

```go
deleteTopicOptions := eventNotificationsService.NewDeleteTopicOptions(
	<instance-id>,
	<topic-id>,
)

response, err := eventNotificationsService.DeleteTopic(deleteTopicOptions)

if err != nil {
	panic(err)
}

```

## Destinations

### Create Destination

```go
createDestinationOptions := eventNotificationsService.NewCreateDestinationOptions(
	<instance-id>,
	<destination-name>,
	<destination-type>,
)
destinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfWebhookDestinationConfig{
	URL:              core.StringPtr(<destination-config-url>),
	Verb:             core.StringPtr(<destination-config-verb>),
	CustomHeaders:    make(map[string]string),
	SensitiveHeaders: []string{<header-key>},
}
destinationConfigModel := &eventnotificationsv1.DestinationConfig{
	Params: destinationConfigParamsModel,
}
createDestinationOptions.SetConfig(destinationConfigModel)

destination, response, err := eventNotificationsService.CreateDestination(createDestinationOptions)

if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(destination, "", "  ")
fmt.Println(string(b))
```

Among the supported destinations, if you need to create Push Notification destinations, you have the additional option of choosing a destination of production type or pre-production type.
Set `pre_prod` boolean parameter to _true_ to configure destination as pre-production destination else set the value as _false_.
Supported destinations are Android, iOS, Chrome, Firefox and Safari.

### List Destinations

```go
listDestinationsOptions := eventNotificationsService.NewListDestinationsOptions(
	<instance-id>,
)

destinationList, response, err := eventNotificationsService.ListDestinations(listDestinationsOptions)
if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(destinationList, "", "  ")
fmt.Println(string(b))
```

### Get Destination

```go
getDestinationOptions := eventNotificationsService.NewGetDestinationOptions(
	<instance-id>,       // Event notifications service instance GUID
	<destination-id>,    // Event notifications service instance Destination ID
)

destination, response, err := eventNotificationsService.GetDestination(getDestinationOptions)

if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(destination, "", "  ")
fmt.Println(string(b))
```

### Update Destination

```go
destinationConfigParamsModel := &eventnotificationsv1.DestinationConfigOneOfWebhookDestinationConfig{
	URL:              core.StringPtr(<destination-config-update-url>),
	Verb:             core.StringPtr(<destination-config-update-verb>),
	CustomHeaders:    make(map[string]string),
	SensitiveHeaders: []string{<header-key>},
}

destinationConfigModel := &eventnotificationsv1.DestinationConfig{
	Params: destinationConfigParamsModel,
}

updateDestinationOptions := eventNotificationsService.NewUpdateDestinationOptions(
	<instance-id>,      // Event notifications service instance GUID
	<destination-id>,   // Event notifications service instance Destination ID
)

updateDestinationOptions.SetName(<destination-update-name>)
updateDestinationOptions.SetDescription(<destination-update-description>)
updateDestinationOptions.SetConfig(destinationConfigModel)

destination, response, err := eventNotificationsService.UpdateDestination(updateDestinationOptions)

if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(destination, "", "  ")
fmt.Println(string(b))

```

### Delete Destination

```go
deleteDestinationOptions := eventNotificationsService.NewDeleteDestinationOptions(
	<instance-id>,		// Event notifications service instance GUID
	<destination-id>,	// Event notifications service instance Destination ID
)

response, err := eventNotificationsService.DeleteDestination(deleteDestinationOptions)

if err != nil {
	panic(err)
}
```

### Test Destination

This functionality allows you to test a destination. The feature simplifies the process of verifying whether a destination is functioning correctly.
Currently, this functionality supports following destinations:

1. Slack
2. PagerDuty
3. ServiceNow
4. Microsoft&reg; Teams
5. IBM Cloud Code Engine
6. IBM Cloud Object Storage

```go
testDestinationOptions := &eventnotificationsv1.TestDestinationOptions{
	<instance-id>,		// Event notifications service instance GUID
	<destination-id>,	// Event notifications service instance Destination ID
}

result, response, err := eventNotificationsService.TestDestination(testDestinationOptions)
```

Once the test is completed, you will be presented with the results. These results will typically include:

- **Status**: Whether the test is successful or failed
- **Response Code**: If test fails, then the response code sent from the end destination client is returned
- **Response Message**: If test fails, then the response message sent from the end destination client is returned

### Custom Domain Name Verification

After creation of the custom email destination with your domain name, make sure its validated for the right ownership. This can be done with SPF and DKIM verification.

- Sender Policy Framework (SPF), which is used to authenticate the sender of an email. SPF specifies the mail servers that are allowed to send email for your domain.
- DomainKeys Identified Mail (DKIM), which allows an organization to take responsibility for transmitting a message by signing it. DKIM allows the receiver to check the email that claimed to have come from a specific domain, is authorized by the owner of that domain.

```go
customSpfDkimUpdateDestinationOptions := &eventnotificationsv1.UpdateVerifyDestinationOptions{
	InstanceID: core.StringPtr(<instance-id>),       // Event notifications service instance GUID
	ID:         core.StringPtr(<destination-id>),	 // Event notifications service instance Destination ID
	Type:       core.StringPtr(<verification-type>), // verification type spf/dkim
}

result, response, err := eventNotificationsService.UpdateVerifyDestination(customSpfUpdateDestinationOptions)

if err != nil {
	panic(err)
}
```

## Templates

Template is a pre-defined layout, that may include content like images, text and dynamic content based on event. Rather than creating a new content from scratch each time, you can use a template as a base and configure them in subscription.
supports the following templates:

- Custom Email notification
- Custom Email invitation

### Create Template

#### Custom Email Template

```go
templConfig := &eventnotificationsv1.TemplateConfigOneOfEmailTemplateConfig{
	Body:    core.StringPtr(<base 64 encoded html content>),
	Subject: core.StringPtr(<email-subject>),
}

createTemplateOptions := &eventnotificationsv1.CreateTemplateOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	Name:        core.StringPtr(<name>),
	Type:        core.StringPtr(<template-type>),
	Description: core.StringPtr(<description>),
	Params:      templConfig,
}

templateResponse, response, err := eventNotificationsService.CreateTemplate(createTemplateOptions)
```

For custom email supported template type values: smtp_custom.invitation, smtp_custom.notification

#### Slack Template

```go
slackTemplConfig := &eventnotificationsv1.TemplateConfigOneOfSlackTemplateConfig{
	Body: core.StringPtr(<json body encoded in to base 64 format>),
}

createTemplateOptions = &eventnotificationsv1.CreateTemplateOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	Name:        core.StringPtr(<name>),
	Type:        core.StringPtr(<template-type>),
	Description: core.StringPtr(<description>),
	Params:      slackTemplConfig,
}

templateResponse, response, err = eventNotificationsService.CreateTemplate(createTemplateOptions)
```

For slack template supported template type value: slack.notification

#### Webhook Template

```go
webhookTemplConfig := &eventnotificationsv1.TemplateConfigOneOfWebhookTemplateConfig{
	Body: core.StringPtr(<json body encoded in to base 64 format>),
}

createTemplateOptions = &eventnotificationsv1.CreateTemplateOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	Name:        core.StringPtr(<name>),
	Type:        core.StringPtr(<template-type>),
	Description: core.StringPtr(<description>),
	Params:      webhookTemplConfig,
}

templateResponse, response, err = eventNotificationsService.CreateTemplate(createTemplateOptions)
```

For webhook template supported template type value: webhook.notification

#### Pagerduty Template

```go
templateConfig := &eventnotificationsv1.TemplateConfigOneOfPagerdutyTemplateConfig{
	Body:    core.StringPtr(<base 64 encoded json body>),
}

createTemplateOptions := &eventnotificationsv1.CreateTemplateOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<template-id>),
	Name:        core.StringPtr(<name>),
	Type:        core.StringPtr(<template-type>),
	Description: core.StringPtr(<description>),
	Params:      templateConfig,
}

templateResponse, response, err := eventNotificationsService.CreateTemplate(createTemplateOptions)
```

For pagerduty template supported template type value: pagerduty.notification

#### Event Streams Template

```go
templateConfig := &eventnotificationsv1.TemplateConfigOneOfEventStreamsTemplateConfig{
	Body:    core.StringPtr(<base 64 encoded json body>),
}

createTemplateOptions := &eventnotificationsv1.CreateTemplateOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<template-id>),
	Name:        core.StringPtr(<name>),
	Type:        core.StringPtr(<template-type>),
	Description: core.StringPtr(<description>),
	Params:      templateConfig,
}

templateResponse, response, err := eventNotificationsService.CreateTemplate(createTemplateOptions)
```

For event streams template supported template type value: event_streams.notification

### List Templates

```go
listTemplatesOptions := eventNotificationsService.NewListTemplatesOptions(
	InstanceID: core.StringPtr(<instance-id>),
)

templatesList, response, err := eventNotificationsService.ListTemplates(listTemplatesOptions)
```

### Get Template

```go
getTemplateOptions := &eventnotificationsv1.GetTemplateOptions{
	InstanceID: core.StringPtr(<instance-id>),
	ID:         core.StringPtr(<template-id>),
}

template, response, err := eventNotificationsService.GetTemplate(getTemplateOptions)
```

### Update Template

#### Update Email Template

```go
templateConfig := &eventnotificationsv1.TemplateConfigOneOfEmailTemplateConfig{
	Body:    core.StringPtr(<base 64 encoded html content>),
	Subject: core.StringPtr(<email-subject>),
}

replaceTemplateOptions := &eventnotificationsv1.ReplaceTemplateOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<template-id>),
	Name:        core.StringPtr(<name>),
	Type:        core.StringPtr(<template-type>),
	Description: core.StringPtr(<description>),
	Params:      templateConfig,
}

templateResponse, response, err := eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
```

For custom email supported template type values: smtp_custom.invitation, smtp_custom.notification

#### Update Slack Template

```go
templateConfig := &eventnotificationsv1.TemplateConfigOneOfSlackTemplateConfig{
	Body:    core.StringPtr(<base 64 encoded json body>),
}

replaceTemplateOptions := &eventnotificationsv1.ReplaceTemplateOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<template-id>),
	Name:        core.StringPtr(<name>),
	Type:        core.StringPtr(<template-type>),
	Description: core.StringPtr(<description>),
	Params:      templateConfig,
}

templateResponse, response, err := eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
```

For slack template supported template type value: slack.notification

#### Update Webhook Template

```go
templateConfig := &eventnotificationsv1.TemplateConfigOneOfWebhookTemplateConfig{
	Body:    core.StringPtr(<base 64 encoded json body>),
}

replaceTemplateOptions := &eventnotificationsv1.ReplaceTemplateOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<template-id>),
	Name:        core.StringPtr(<name>),
	Type:        core.StringPtr(<template-type>),
	Description: core.StringPtr(<description>),
	Params:      templateConfig,
}

templateResponse, response, err := eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
```

For webhook template supported template type value: webhook.notification

#### Update Pagerduty Template

```go
templateConfig := &eventnotificationsv1.TemplateConfigOneOfPagerdutyTemplateConfig{
	Body:    core.StringPtr(<base 64 encoded json body>),
}

replaceTemplateOptions := &eventnotificationsv1.ReplaceTemplateOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<template-id>),
	Name:        core.StringPtr(<name>),
	Type:        core.StringPtr(<template-type>),
	Description: core.StringPtr(<description>),
	Params:      templateConfig,
}

templateResponse, response, err := eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
```

For pagerduty template supported template type value: pagerduty.notification

#### Update Event Streams Template

```go
templateConfig := &eventnotificationsv1.TemplateConfigOneOfEventStreamsTemplateConfig{
	Body:    core.StringPtr(<base 64 encoded json body>),
}

replaceTemplateOptions := &eventnotificationsv1.ReplaceTemplateOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<template-id>),
	Name:        core.StringPtr(<name>),
	Type:        core.StringPtr(<template-type>),
	Description: core.StringPtr(<description>),
	Params:      templateConfig,
}

templateResponse, response, err := eventNotificationsService.ReplaceTemplate(replaceTemplateOptions)
```

For event streams template supported template type value: event_streams.notification

### Delete Template

```go
deleteTemplateOptions := &eventnotificationsv1.DeleteTemplateOptions{
	InstanceID: core.StringPtr(<instance-id>),
	ID:         core.StringPtr(<template-id>),
}

response, err := eventNotificationsService.DeleteTemplate(deleteTemplateOptions)
```

### List Predefined Templates

```go
listpredefinedtemplatesOptions := &eventnotificationsv1.ListPreDefinedTemplatesOptions{
				InstanceID: core.StringPtr(instanceID),
				Source:     core.StringPtr(<source-type>),
				Type:       core.StringPtr(<destination-template-type>),
				Offset:     core.Int64Ptr(int64(0)),
				Limit:      core.Int64Ptr(int64(1)),
				Search:     core.StringPtr(search),
			}

```

### Get Predefined Template
```go
getPredefinedTemplateOptions := &eventnotificationsv1.GetPreDefinedTemplateOptions{
				InstanceID: core.StringPtr(instanceID),
				ID:         core.StringPtr(<pre-defined-template-id>),
			}
```

## Push Destination APIs

### Create Destination tag subscription

```go
createTagsSubscriptionOptions := eventNotificationsService.NewCreateTagsSubscriptionOptions(
	<instance-id>,		// Event notifications service instance GUID
	<destination-id>,	// Event notifications service instance Destination ID
	<device-id>,		// Event notifications service device ID
	<tag-name>,			// Event notifications service tag name
)

destinationTagsSubscriptionResponse, response, err := eventNotificationsService.CreateTagsSubscription(createTagsSubscriptionOptions)

if err != nil {
	panic(err)
}
```

### List Destination tag subscription

```go
listTagsSubscriptionOptions := eventNotificationsService.NewListTagsSubscriptionOptions(
	<instance-id>,		// Event notifications service instance GUID
	<destination-id>,	// Event notifications service instance Destination ID
)

tagsSubscriptionList, response, err := eventNotificationsService.ListTagsSubscription(listTagsSubscriptionOptions)

if err != nil {
	panic(err)
}
```

### Delete Destination device tag subscription

```go
deleteTagsSubscriptionOptions := eventNotificationsService.NewDeleteTagsSubscriptionOptions(
	<instance-id>,		// Event notifications service instance GUID
	<destination-id>,	// Event notifications service instance Destination ID
)

deleteTagsSubscriptionOptions.SetDeviceID(<device-id>)
deleteTagsSubscriptionOptions.SetTagName(<tag-name>)
response, err := eventNotificationsService.DeleteTagsSubscription(deleteTagsSubscriptionOptions)
if err != nil {
	panic(err)
}
```

## Subscriptions

### Create Subscription

```go
`While Creating Subscription use any of one option from webhook or email`

subscriptionCreateAttributesModel := &eventnotificationsv1.SubscriptionCreateAttributes{
	SigningEnabled: core.BoolPtr(false),
	TemplateIDNotification: core.StringPtr(<webhook-template-id>),
}

createSubscriptionOptions := eventNotificationsService.NewCreateSubscriptionOptions(
	<instance-id>,	// Event notifications service instance GUID
	<subscription-name>,
	<destination-id>, // Event notifications service instance Destination ID
	<topic-id>,  // Event notifications service instance Topic ID
	subscriptionCreateAttributesModel,
)

createSubscriptionOptions.SetDescription(<subscription-description>)

subscription, response, err := eventNotificationsService.CreateSubscription(createSubscriptionOptions)

if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(subscription, "", "  ")
fmt.Println(string(b))
```

### List Subscriptions

```go
listSubscriptionsOptions := eventNotificationsService.NewListSubscriptionsOptions(
	<instance-id>,	// Event notifications service instance GUID
)

subscriptionList, response, err := eventNotificationsService.ListSubscriptions(listSubscriptionsOptions)

if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(subscriptionList, "", "  ")
fmt.Println(string(b))
```

### Get Subscription

```go
getSubscriptionOptions := eventNotificationsService.NewGetSubscriptionOptions(
	<instance-id>,	// Event notifications service instance GUID
	<subscription-id>,	// Event notifications service instance Subscription ID
)

subscription, response, err := eventNotificationsService.GetSubscription(getSubscriptionOptions)

if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(subscription, "", "  ")
fmt.Println(string(b))
```

### Update Subscription

```go

updateSubscriptionOptions := eventNotificationsService.NewUpdateSubscriptionOptions(
	<instance-id>,	// Event notifications service instance GUID
	<subscription-id>,	// Event notifications service instance Subscription ID
)

subscriptionUpdateAttributesModel := &eventnotificationsv1.SubscriptionUpdateAttributesWebhookAttributes{
	SigningEnabled: core.BoolPtr(true),
	TemplateIDNotification: core.StringPtr(<webhook-template-id>),
}

updateSubscriptionOptions.SetAttributes(subscriptionUpdateAttributesModel)
updateSubscriptionOptions.SetDescription(<subscription-update-description>)
updateSubscriptionOptions.SetName(<subscription-update-name>)

subscription, response, err := eventNotificationsService.UpdateSubscription(updateSubscriptionOptions)

if err != nil {
	panic(err)
}

b, _ := json.MarshalIndent(subscription, "", "  ")
fmt.Println(string(b))

```

### Delete Subscription

```go
deleteSubscriptionOptions := eventNotificationsService.NewDeleteSubscriptionOptions(
	<instance-id>,	// Event notifications service instance GUID
	<subscription-id>,	// Event notifications service instance Subscriptions ID
)

response, err := eventNotificationsService.DeleteSubscription(deleteSubscriptionOptions)

if err != nil {
	panic(err)
}
```

## Integration

### Create Integration

```go
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

```

### Get Integration

```go
getIntegrationOptions := &eventnotificationsv1.GetIntegrationOptions{
	InstanceID: core.StringPtr(<instance-id>),
	ID:         core.StringPtr(<integration-id>),
}

integrationResponse, response, err := eventNotificationsService.GetIntegration(getIntegrationOptions)
```

### List Integrations

```go

listIntegrationsOptions := &eventnotificationsv1.ListIntegrationsOptions{
	InstanceID: core.StringPtr(<instance-id>),
	Limit:      core.Int64Ptr(<limit>),
	Offset:     core.Int64Ptr(<Offset>),
	Search:     core.StringPtr(<search>),
}

integrationResponse, response, err := eventNotificationsService.ListIntegrations(listIntegrationsOptions)
```

### Update Integration

For kms/hs-crypto-

```go
integrationMetadata := &eventnotificationsv1.IntegrationMetadata{
	Endpoint:  core.StringPtr(<end-point-url>),
	CRN:       core.StringPtr(<crn>),
	RootKeyID: core.StringPtr(<root-key-id>),
}

replaceIntegrationsOptions := &eventnotificationsv1.ReplaceIntegrationOptions{
	InstanceID: core.StringPtr(instanceID),
	ID:         core.StringPtr(integrationId),
	Type:       core.StringPtr("kms/hs-crypto"),
	Metadata:   integrationMetadata,
}

integrationResponse, response, err := eventNotificationsService.ReplaceIntegration(replaceIntegrationsOptions)
```

For Cloud Object Storage-

```go
integrationMetadata := &eventnotificationsv1.IntegrationMetadata{
	Endpoint:  core.StringPtr(<COS-end-point-url>),
	CRN:       core.StringPtr(<COS-instance-crn>),
	BucketName: core.StringPtr(<COS-bucket-name>),
}

replaceIntegrationsOptions := &eventnotificationsv1.ReplaceIntegrationOptions{
	InstanceID: core.StringPtr(instanceID),
	ID:         core.StringPtr(integrationId),
	Type:       core.StringPtr("collect-failed-events"),
	Metadata:   integrationMetadata,
}

integrationResponse, response, err := eventNotificationsService.ReplaceIntegration(replaceIntegrationsOptions)
```

## SMTPConfigurations

### Create SMTP Configuration

```go
createSMTPConfigurationOptions := &eventnotificationsv1.CreateSMTPConfigurationOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	Domain:      core.StringPtr(<domain-name>),
	Description: core.StringPtr(<description>),
	Name:        core.StringPtr(<name>),
}

smtpConfig, response, err := eventNotificationsService.CreateSMTPConfiguration(createSMTPConfigurationOptions)

```

### Create SMTP User

```go
createSMTPUserOptions := &eventnotificationsv1.CreateSMTPUserOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<smtp-Config-id)>,
	Description: core.StringPtr(<description),
}

user, response, err := eventNotificationsService.CreateSMTPUser(createSMTPUserOptions)

```

### Get SMTP Configuration

```go
getSMTPconfigurationOptions := &eventnotificationsv1.GetSMTPConfigurationOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<smtp-Config-id)>,
}

smtpConfiguration, response, err := eventNotificationsService.GetSMTPConfiguration(getSMTPconfigurationOptions)

```

### Get SMTP User

```go
getSMTPUserOptions := &eventnotificationsv1.GetSMTPUserOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<smtp-Config-id)>,
	UserID:     core.StringPtr(<user-id>),
}

SMTPUser, response, err := eventNotificationsService.GetSMTPUser(getSMTPUserOptions)
```

### Get SMTP Allowed Ips

```go
getSMTPAllowedIPsOptions := &eventnotificationsv1.GetSMTPAllowedIpsOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<smtp-Config-id)>,
}

smtpAllowedIPs, response, err := eventNotificationsService.GetSMTPAllowedIps(getSMTPAllowedIPsOptions)
```

### List SMTP Configurations

```go
listSMTPConfigurationsOptions := &eventnotificationsv1.ListSMTPConfigurationsOptions{
	InstanceID: core.StringPtr(<instance-id>),
	Limit:      core.Int64Ptr(<limit>),
	Offset:     core.Int64Ptr(<offset>),
	Search:     core.StringPtr(<search>),
}

smtpConfigurations, response, err := eventNotificationsService.ListSMTPConfigurations(listSMTPConfigurationsOptions)
```

### List SMTP Users

```go
listSMTPUsersOptions := &eventnotificationsv1.ListSMTPUsersOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<smtp-Config-id)>,
	Limit:      core.Int64Ptr(<limit>),
	Offset:     core.Int64Ptr(<offset>),
	Search:     core.StringPtr(<search>),
}

smtpUsers, response, err := eventNotificationsService.ListSMTPUsers(listSMTPUsersOptions)
```

### Update SMTP Configuration

```go
updateSMTPConfigurationOptions := &eventnotificationsv1.UpdateSMTPConfigurationOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<smtp-Config-id)>,
	Name:        core.StringPtr(<name>),
	Description: core.StringPtr(<description>),
}

updateSMTPConfiguration, response, err := eventNotificationsService.UpdateSMTPConfiguration(updateSMTPConfigurationOptions)
```

### Update SMTP User

```go
updateSMTPUserOptions := &eventnotificationsv1.UpdateSMTPUserOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<smtp-Config-id)>,
	Description: core.StringPtr(<description>),
	UserID:      core.StringPtr(<user-id>),
}

updateSMTPUser, response, err := eventNotificationsService.UpdateSMTPUser(updateSMTPUserOptions)
```

### Delete SMTP User

```go
deleteSMTPUserOptions := &eventnotificationsv1.DeleteSMTPUserOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<smtp-Config-id)>,
	UserID:     core.StringPtr(<user-id>),
}

response, err := eventNotificationsService.DeleteSMTPUser(deleteSMTPUserOptions)
```

### Delete SMTP Configuration

```go
deleteSMTPConfigurationOptions := &eventnotificationsv1.DeleteSMTPConfigurationOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<smtp-Config-id)>,
}

response, err := eventNotificationsService.DeleteSMTPConfiguration(deleteSMTPConfigurationOptions)
```

### Verify SMTP

```go
updateVerifySMTPOptions := &eventnotificationsv1.UpdateVerifySMTPOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<smtp-Config-id)>,
	Type:       core.StringPtr(<verification-type>),
}

verifySMTP, response, err := eventNotificationsService.UpdateVerifySMTP(updateVerifySMTPOptions)
```

supported verification types are dkim,spf and en_authorization.

## Metrics

### Get Metrics

```go
getMetricsOptions := &eventnotificationsv1.GetMetricsOptions{
	InstanceID:      core.StringPtr(<instance-id>),
	DestinationType: core.StringPtr("smtp_custom"),
	Gte:             core.StringPtr(<gte-timestamp>),
	Lte:             core.StringPtr(<lte-timestamp>),
	EmailTo:         core.StringPtr(<email-to>),
	DestinationID:              core.StringPtr(<destination-id>),
	NotificationID:  core.StringPtr(<notification-id>),
	Subject:         core.StringPtr(<subject>),
}

metrics, response, err := eventNotificationsService.GetMetrics(getMetricsOptions)
```

## Send Notifications

```go
notificationevicesModel := map[string]interface{}{
	UserIds: []string{"<user-ids>"},
	FcmDevices: []string{"<fcm-device-ids>"},
	ApnsDevices: []string{"<apns-device-ids>"},
	Tags: []string{"<tag-names>"},
	Platforms: []string{"<device-platforms>"},
}
devicesbody, _ := json.Marshal(notificationevicesModel)
devicesbodyString := string(devicesbody)


notificationID := "<notification-id>"
notificationSeverity := "<notification-severity>"
typeValue := "<notification-type>"
notificationsSouce := "<notification-source>"
specVersion := "1.0"

notificationDevicesModel := "{\"user_ids\": [\"userId\"]}"
notificationFcmBodyModel := "{\"message\": {\"android\": {\"notification\": {\"title\": \"Alert message\",\"body\": \"Bob wants to play Poker\"},\"data\": {\"name\": \"Willie Greenholt\",\"description\": \"notification for the Poker\"}}}}"
notificationAPNsBodyModel := "{\"alert\": \"Game Request\", \"badge\": 5 }"
notificationSafariBodyModel := "{\"aps\":{\"alert\":{\"title\":\"FlightA998NowBoarding\",\"body\":\"BoardinghasbegunforFlightA998.\",\"action\":\"View\"},\"url-args\":[\"boarding\",\"A998\"]}}}"
mailTo := "[\"abc@ibm.com\", \"def@us.ibm.com\"]"
smsTo := "[\"+911234567890\", \"+911224567890\"]"
slackTo := "[\"C07FALXBH4G\",\"C07FALXBU4G\"]";
templates := "[\"149b0e11-8a7c-4fda-a847-5d79e01b71dc\"]"
htmlBody := "\"Hi  ,<br/>Certificate expiring in 90 days.<br/><br/>Please login to <a href=\"https: //cloud.ibm.com/security-compliance/dashboard\">Security and Complaince dashboard</a> to find more information<br/>\""
mms := "{\"content\": \"mms content\", \"content_type\": \"image/png\"}"
markdown := "**Event Summary** \n\n**Toolchain ID:** `4414af34-a5c7-47d3-8f05-add4af6d78a6`  \n**Content Type:** `application/json`\n\n---\n\n *Pipeline Run Details*\n\n- **Namespace:** `PR`\n- **Trigger Name:** `manual`\n- **Triggered By:** `nitish.kulkarni3@ibm.com`\n- **Build Number:** `343`\n- **Pipeline Link:** [View Pipeline Run](https://cloud.ibm.com/devops/pipelines/tekton/e9cd5aa3-a3f2-4776-8acc-26a35922386e/runs/f29ac6f5-bd2f-4a26-abb8-4249be8dbab7?env_id=ibm:yp:us-south)"


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
notificationCreateModel.Ibmenpushto = &devicesbodyString
notificationCreateModel.Ibmenmailto = &mailTo
notificationCreateModel.Ibmensmsto = &smsTo
notificationCreateModel.Ibmenslackto = &slackTo
notificationCreateModel.Ibmentemplates = &templates
notificationCreateModel.Ibmensubject = core.StringPtr("Notification subject")
notificationCreateModel.Ibmenhtmlbody = core.StringPtr(htmlBody)
notificationCreateModel.Ibmendefaultshort = core.StringPtr("Alert message")
notificationCreateModel.Ibmendefaultlong = core.StringPtr("Alert message on expiring offer")
notificationCreateModel.Ibmenmarkdown = &markdown

sendNotificationsOptionsModel := new(eventnotificationsv1.SendNotificationsOptions)
sendNotificationsOptionsModel.InstanceID = &instanceID
sendNotificationsOptionsModel.Body = notificationCreateModel

notificationResponse, response, err := eventNotificationsService.SendNotifications(sendNotificationsOptionsModel)

if err != nil {
	panic(err)
}
```

<details open>
<summary>Send Notifications Variables</summary>
<br>

- **ibmenpushto** - Set up the push notifications targets.
  - **user_ids** (_Array of String_) - Send notification to the specified userIds.
  - **fcm_devices** (_Array of String_) - Send notification to the list of specified Android devices.
  - **apns_devices** (_Array of String_) - Send notification to the list of specified iOS devices.
  - **chrome_devices** (_Array of String_) - Send notification to the list of specified Chrome devices.
  - **firefox_devices** (_Array of string_) - Send notification to the list of specified Firefox devices.
  - **tags** (_Array of string_) - Send notification to the devices that have subscribed to any of these tags.
  - **platforms** (_Array of string_) - Send notification to the devices of the specified platforms.
    - Pass 'G' for google (Android) devices.
    - Pass 'A' for iOS devices.
    - Pass 'WEB_FIREFOX' for Firefox browser.
    - Pass 'WEB_CHROME' for Chrome browser.
- **Event Notifications SendNotificationsOptions** - Event Notifications Send Notifications method.
  - **instance_id\*** (_string_) - Unique identifier for IBM Cloud Event Notifications instance.
  - **ibmenseverity** (_string_) - Severity for the notifications. Some sources can have the concept of an Event severity. Hence a handy way is provided to specify a severity of the event. example: LOW, HIGH, MEDIUM
  - **id\*** (_string_) - A unique identifier that identifies each event. source+id must be unique. The backend should be able to uniquely track this id in logs and other records. Send unique ID for each send notification. Same ID can be sent in case of failure of send notification. source+id will be logged in IBM Cloud Logging service. Using this combination we will be able to trace the event movement from one system to another and will aid in debugging and tracing.
  - **source\*** (_string_) - Source of the notifications. This is the identifier of the event producer. A way to uniquely identify the source of the event. For IBM Cloud services this is the crn of the service instance producing the events. For API sources this can be something the event producer backend can uniquely identify itself with.
  - **ibmensourceid\*** (_string_) - This is the ID of the source created in EN. This is available in the EN UI in the "Sources" section.
  - **type** (_string_) - This describes the type of event. It is of the form <event-type-name>:<sub-type> This type is defined by the producer. The event type name has to be prefixed with the reverse DNS names so the event type is uniquely identified. The same event type can be produced by 2 different sources. It is highly recommended to use hyphen - as a separator instead of \_.
  - **data** (_string_) - The payload for webhook notification. If data is added as part of payload then its mandatory to add **datacontenttype**.
  - **datacontenttype** - The notification content type. example: application/json
  - **time** (_string_) - Time of the notifications. UTC time stamp when the event occurred. Must be in the RFC 3339 format.
  - **ibmenpushto** (_string_) - Targets for the FCM notifications. This contains details about the destination where you want to send push notification. This attribute is mandatory for successful delivery from an Android FCM or APNS destination.
  - **ibmenfcmbody** (_string_) - Set payload string specific to Android platform [Refer this FCM official [link](https://firebase.google.com/docs/cloud-messaging/http-server-ref#notification-payload-support)].
  - **ibmenhuaweibody** (_string_) - Set payload string specific to Android platform [Refer this FCM official [link](https://firebase.google.com/docs/cloud-messaging/http-server-ref#notification-payload-support)].
  - **ibmenapnsbody** (_string_) - Set payload string specific to iOS platform [Refer this APNs official doc [link](https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CreatingtheNotificationPayload.html)].
  - **ibmensafaribody** (_string_) - Set payload string specific to safari platform [Refer this Safari official doc [link](https://developer.huawei.com/consumer/en/hms/huawei-pushkit)].
  - **ibmenapnsheaders** (_string_) - Set headers required for the APNs message [Refer this APNs official [link](https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/sending_notification_requests_to_apns)(Table 1 Header fields for a POST request)]
  - **ibmenchromebody** (_string_) - Message body for the Chrome notifications. Refer [this official documentation](https://developer.mozilla.org/en-US/docs/Web/API/Notification/Notification) for more.
  - **ibmenfirefoxbody** (_string_) - Message body for the Firefox notifications. Refer [this official documentation](https://developer.mozilla.org/en-US/docs/Web/API/Notification/Notification) for more.
  - **ibmenchromeheaders** (_string_) - Headers for the Chrome notifications. Refer [this official documentation](https://developer.mozilla.org/en-US/docs/Web/API/Notification/Notification) for more.
  - **ibmenfirefoxheaders** (_string_) - Headers for the Firefox notifications. Refer [this official documentation](https://developer.mozilla.org/en-US/docs/Web/API/Notification/Notification) for more.
  - **ibmendefaultshort\*** (_string_) - Default short text for the message.
  - **ibmendefaultlong\*** (_string_) - Default long text for the message.
  - **specversion\*** (_string_) - Spec version of the Event Notifications. Default value is `1.0`.
  - **ibmenhtmlbody** (_string_) - The html body of notification for email.
  - **ibmenmailto** (_Array of string_) - Array of email ids to which the notification to be sent.
  - **ibmensmsto** (_Array of string_) - Array of SMS numbers to which the notification to be sent.
  - **ibmenslackto** (_Array of string_) - Array of Slack channel/member ids to which the notification to be sent.
  - **ibmentemplates** (_Array of string_) - Array of template IDs that needs to be applied while sending notificatin for custom domain email and slack destination.
  - **ibmenmarkdown** (_string_) - The markdown content of pretty formatting.

Note: variable with \* represents the mandatory attribute.

</details>

## Set Environment

Find [event_notifications_v1.env.hide](https://github.com/IBM/event-notifications-go-admin-sdk/blob/main/event_notifications_v1.env.hide) in the repo and rename it to `event_notifications_v1.env`. After that add the values for,

- `EVENT_NOTIFICATIONS_URL` - Add the Event Notifications service instance Url.
- `EVENT_NOTIFICATIONS_APIKEY` - Add the Event Notifications service instance apikey.
- `EVENT_NOTIFICATIONS_GUID` - Add the Event Notifications service instance GUID.

**Optional**

- `EVENT_NOTIFICATIONS_AUTH_URL` - Add the IAM url if you are using IBM test cloud.
- `EVENT_NOTIFICATIONS_FCM_KEY` - Add firebase server key for Android FCM destination.
- `EVENT_NOTIFICATIONS_FCM_ID` - Add firebase sender Id for Android FCM destination.
- `EVENT_NOTIFICATIONS_FCM_PROJECT_ID` - fcm project id
- `EVENT_NOTIFICATIONS_FCM_CLIENT_EMAIL` - fcm client email
- `EVENT_NOTIFICATIONS_FCM_PRIVATE_KEY` - fcm private key
- `EVENT_NOTIFICATIONS_SAFARI_CERTIFICATE` - safari certificate path

- `EVENT_NOTIFICATIONS_SNOW_CLIENT_ID` - service now client id
- `EVENT_NOTIFICATIONS_SNOW_CLIENT_SECRET` - service now client secret
- `EVENT_NOTIFICATIONS_SNOW_USER_NAME` - service now user name
- `EVENT_NOTIFICATIONS_SNOW_PASSWORD` - service now password
- `EVENT_NOTIFICATIONS_SNOW_INSTANCE_NAME` - service now instance name

- `EVENT_NOTIFICATIONS_COS_BUCKET_NAME` - cloud object storage bucket name
- `EVENT_NOTIFICATIONS_COS_INSTANCE` - cloud object storage instance id
- `EVENT_NOTIFICATIONS_COS_INSTANCE_CRN` - cloud object storage instance crn
- `EVENT_NOTIFICATIONS_COS_ENDPOINT` - cloud object storage end point

- `EVENT_NOTIFICATIONS_CODE_ENGINE_URL` - code engine app url
- `EVENT_NOTIFICATIONS_CODE_ENGINE_PROJECT_CRN` - code engine project crn
- `EVENT_NOTIFICATIONS_HUAWEI_CLIENT_SECRET` - huawei client secret
- `EVENT_NOTIFICATIONS_HUAWEI_CLIENT_ID` - huawei client id

- `EVENT_NOTIFICATIONS_SLACK_URL` - slack webhook url
- `EVENT_NOTIFICATIONS_MS_TEAMS_URL` - msteams webhook url
- `EVENT_NOTIFICATIONS_PD_ROUTING_KEY` - pagerduty routing key
- `EVENT_NOTIFICATIONS_PD_API_KEY` - pagerduty api key
- `EVENT_NOTIFICATIONS_TEMPLATE_BODY` - base 64 encoded html content
- `EVENT_NOTIFICATIONS_SLACK_TEMPLATE_BODY` - base 64 encoded json body
- `EVENT_NOTIFICATIONS_WEBHOOK_TEMPLATE_BODY` - base 64 encoded json body
- `EVENT_NOTIFICATIONS_PAGERDUTY_TEMPLATE_BODY` - base 64 encoded json body
- `EVENT_NOTIFICATIONS_EVENT_STREAMS_CRN` - Event Streams Instance CRN
- `EVENT_NOTIFICATIONS_EVENT_STREAMS_ENDPOINT` - Event Streams instance endpoint
- `EVENT_NOTIFICATIONS_EVENT_STREAMS_TOPIC` - Event streams instance topic
- `EVENT_NOTIFICATIONS_EVENT_STREAMS_TEMPLATE_BODY` - base 64 encoded json body

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## ⚠️ Deprecation Notice (Attributes)

### Pagerduty Destination Configuration

> The following attribute from DestinationConfigOneOfPagerDutyDestinationConfig
is **deprecated** and will be removed in a future release:

- `APIKey`

This attribute no longer recommended for use and may not be supported in upcoming versions of the SDK. Only `RoutingKey` is expected to be passed.

## Open source @ IBM

Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
