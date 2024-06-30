
# IBM Cloud Event Notifications Go Admin SDK 0.6.1
Go client library to interact with the various [IBM Cloud Event Notifications APIs](https://cloud.ibm.com/apidocs?category=event-notifications).


## Table of Contents


<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  * [Go modules](#go-modules)
  * [`go get` command](#go-get-command)
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

Service Name | Package name 
--- | --- 
<!-- [Example Service](https://cloud.ibm.com/apidocs/example-service) | exampleservicev1 -->
[Event Notifications Service](https://cloud.ibm.com/apidocs/event-notifications) | eventnotificationsv1

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An Event Notifications Instance
* Go version 1.18 or above.

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
1) Setting client options programmatically
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
2) Using external configuration properties
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
	- [Update SMTP Allowed Ips](#update-smtp-allowed-ips)
	- [Delete SMTP User](#delete-smtp-user)
	- [Delete SMTP Configuration](#delete-smtp-user)
	- [Verify SMTP](#verify-smtp)

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
Set `pre_prod` boolean parameter to *true* to configure destination as pre-production destination else set the value as *false*.
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
6. IBM Cloud Functions
7. IBM Cloud Object Storage

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

* Sender Policy Framework (SPF), which is used to authenticate the sender of an email. SPF specifies the mail servers that are allowed to send email for your domain.
* DomainKeys Identified Mail (DKIM), which allows an organization to take responsibility for transmitting a message by signing it. DKIM allows the receiver to check the email that claimed to have come from a specific domain, is authorized by the owner of that domain.

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

### Delete Template
```go
deleteTemplateOptions := &eventnotificationsv1.DeleteTemplateOptions{
	InstanceID: core.StringPtr(<instance-id>),
	ID:         core.StringPtr(<template-id>),
}

response, err := eventNotificationsService.DeleteTemplate(deleteTemplateOptions)
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

### Update SMTP Allowed IPs

```go
updateSMTPAllowedOptions := &eventnotificationsv1.UpdateSMTPAllowedIpsOptions{
	InstanceID:  core.StringPtr(<instance-id>),
	ID:          core.StringPtr(<smtp-Config-id)>,
	Subnets:    []string{"<subnet-ip>"},
}

subnets, response, err := eventNotificationsService.UpdateSMTPAllowedIps(updateSMTPAllowedOptions)
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
templates := "[\"149b0e11-8a7c-4fda-a847-5d79e01b71dc\"]"
htmlBody := "\"Hi  ,<br/>Certificate expiring in 90 days.<br/><br/>Please login to <a href=\"https: //cloud.ibm.com/security-compliance/dashboard\">Security and Complaince dashboard</a> to find more information<br/>\""
mms := "{\"content\": \"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\", \"content_type\": \"image/png\"}"

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
notificationCreateModel.Ibmentemplates = &templates
notificationCreateModel.Ibmensubject = core.StringPtr("Notification subject")
notificationCreateModel.Ibmenhtmlbody = core.StringPtr(htmlBody)
notificationCreateModel.Ibmendefaultshort = core.StringPtr("Alert message")
notificationCreateModel.Ibmendefaultlong = core.StringPtr("Alert message on expiring offer")

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
  - **instance_id** (_string_) - Unique identifier for IBM Cloud Event Notifications instance.
  - **ibmenseverity** (_string_) - Severity for the notifications. Some sources can have the concept of an Event severity. Hence a handy way is provided to specify a severity of the event. example: LOW, HIGH, MEDIUM
  - **id*** (_string_) - A unique identifier that identifies each event. source+id must be unique. The backend should be able to uniquely track this id in logs and other records. Send unique ID for each send notification. Same ID can be sent in case of failure of send notification. source+id will be logged in IBM Cloud Logging service. Using this combination we will be able to trace the event movement from one system to another and will aid in debugging and tracing.
  - **source*** (_string_) - Source of the notifications. This is the identifier of the event producer. A way to uniquely identify the source of the event. For IBM Cloud services this is the crn of the service instance producing the events. For API sources this can be something the event producer backend can uniquely identify itself with.
  - **ibmensourceid*** (_string_) - This is the ID of the source created in EN. This is available in the EN UI in the "Sources" section.
  - **type** (_string_) - This describes the type of event. It is of the form <event-type-name>:<sub-type> This type is defined by the producer. The event type name has to be prefixed with the reverse DNS names so the event type is uniquely identified. The same event type can be produced by 2 different sources. It is highly recommended to use hyphen - as a separator instead of _.
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
  - **ibmendefaultshort*** (_string_) - Default short text for the message.
  - **ibmendefaultlong*** (_string_) - Default long text for the message.
  - **specversion*** (_string_) - Spec version of the Event Notifications. Default value is `1.0`.
  - **ibmenhtmlbody*** (_string_) - The html body of notification for email.
  - **ibmenmailto*** (_Array of string_) - Array of email ids to which the notification to be sent.
  - **ibmensmsto*** (_Array of string_) - Array of SMS numbers to which the notification to be sent.
  - **ibmentemplates*** (_Array of string_) - Array of template IDs that needs to be applied while sending notificatin for custom domain email and slack destination.

Note: variable with * represents the mandatory attribute.
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

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at 
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).


## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
