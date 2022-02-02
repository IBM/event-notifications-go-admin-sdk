
# IBM Cloud Event Notifications Go Admin SDK 0.0.6
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
* Go version 1.16 or above.

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
- region : Region of the Event Notifications Instance


## Using the SDK

SDK Methods to consume

- [Source](#source)
	- [List Sources](#list-sources)
	- [Get Source](#get-sources)
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
- [Destination Devices](#destination-devices)
	- [List Destination device](#list-destination-devices)
	- [Get Destination device report](#get-destination-device-report)
	- [Create Destination tag subscription](#create-destination-tag-subscription)
	- [List Destination tag subscription](#list-destination-tag-subscription)
	- [List Destination device tag subscriptions](#list-destination-device-tag-subscriptions)
	- [Delete Destination device tag subscription](#delete-destination-device-tag-subscription)
- [Subscriptions](#subscriptions)
	- [Create Subscription](#create-subscription)
	- [List Subscriptions](#list-subscriptions)
	- [Get Subscription](#get-subscription)
	- [Update Subscription](#update-subscription)
	- [Delete Subscription](#delete-subscription)
- [Send Notifications](#send-notifications)



## Source 

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
destinationConfigParamsModel := &eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig{
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
destinationConfigParamsModel := &eventnotificationsv1.DestinationConfigParamsWebhookDestinationConfig{
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

## Destination Devices

### List Destination devices

```go
listDestinationDevicesOptions := eventNotificationsService.NewListDestinationDevicesOptions(
	<instance-id>,		// Event notifications service instance GUID
	<destination-id>,	// Event notifications service instance Destination ID
)

destinationDevicesList, response, err := eventNotificationsService.ListDestinationDevices(listDestinationDevicesOptions)

if err != nil {
	panic(err)
}
```

### Get Destination device report

```go
getDestinationDevicesReportOptions := eventNotificationsService.NewGetDestinationDevicesReportOptions(
	<instance-id>,		// Event notifications service instance GUID
	<destination-id>,	// Event notifications service instance Destination ID
)

destinationDevicesReport, response, err := eventNotificationsService.GetDestinationDevicesReport(getDestinationDevicesReportOptions)

if err != nil {
	panic(err)
}
```

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

### List Destination device tag subscriptions

```go
listTagsSubscriptionsDeviceOptions := eventNotificationsService.NewListTagsSubscriptionsDeviceOptions(
	<instance-id>,		// Event notifications service instance GUID
	<destination-id>,	// Event notifications service instance Destination ID
	<device-id>,		// Event notifications service device ID
)

tagsSubscriptionList, response, err := eventNotificationsService.ListTagsSubscriptionsDevice(listTagsSubscriptionsDeviceOptions)

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

## Send Notifications

```go
notificationFcmDevicesModel := &eventnotificationsv1.NotificationFcmDevices{
	UserIds: []string{"<user-ids>"},
	FcmDevices: []string{"<device-ids>"},
	Tags: []string{"<tag-names>"},
	Platforms: []string{"<device-platforms>"},
}

lightsModel := &eventnotificationsv1.Lights{
	LedArgb:  core.StringPtr("<color-name>"),
	LedOnMs:  core.Int64Ptr(int64(0)),
	LedOffMs: core.StringPtr(""),
}

styleModel := &eventnotificationsv1.Style{
	Type:  core.StringPtr("<notification-style>"),
	Title: core.StringPtr("<notification-title>"),
	URL:   core.StringPtr("<notification-url>"),
}

notificationFcmBodyMessageDataModel := &eventnotificationsv1.NotificationFcmBodyMessageData{
	Alert:               core.StringPtr("<notification-alert>"),
	CollapseKey:         core.StringPtr("<notification-collapse_key>"),
	InteractiveCategory: core.StringPtr("<notification-category>"),
	Icon:                core.StringPtr("<notification-icon>"),
	DelayWhileIdle:      core.BoolPtr(true),
	Sync:                core.BoolPtr(true),
	Visibility:          core.StringPtr("<notification-visibility>"),
	Redact:              core.StringPtr("<notification-redact>"),
	Payload:             make(map[string]interface{}),
	Priority:            core.StringPtr("<notification-priority>"),
	Sound:               core.StringPtr("<notification-sound>"),
	TimeToLive:          core.Int64Ptr(int64(0)),
	Lights:              lightsModel,
	AndroidTitle:        core.StringPtr("<notification-title>"),
	GroupID:             core.StringPtr("<notification-group-id>"),
	Style:               styleModel,
	Type:                core.StringPtr("<notification-type>"),
}

notificationFcmBodyMessageModel := &eventnotificationsv1.NotificationFcmBodyMessage{
	Data: notificationFcmBodyMessageDataModel,
}

notificationFcmBodyModel := &eventnotificationsv1.NotificationFcmBody{
	Message: notificationFcmBodyMessageModel,
}

notificationID := "<notification-id>"
notificationSubject := "<notification-subject>"
notificationSeverity := "<notification-severity>"
typeValue := "<notification-type>"
notificationsSouce := "<notification-source>"

sendNotificationsOptions := &eventnotificationsv1.SendNotificationsOptions{
	InstanceID:      core.StringPtr(instanceID),
	Subject:         core.StringPtr(notificationSubject),
	Severity:        core.StringPtr(notificationSeverity),
	ID:              core.StringPtr(notificationID),
	Source:          core.StringPtr(notificationsSouce),
	EnSourceID:      core.StringPtr(sourceID),
	Type:            core.StringPtr(typeValue),
	Time:            "<notification-time>",
	Data:            make(map[string]interface{}),
	PushTo:          notificationFcmDevicesModel,
	MessageFcmBody:  notificationFcmBodyModel,
	Datacontenttype: core.StringPtr("application/json"),
	Specversion:     core.StringPtr("1.0"),
}

notificationResponse, response, err := eventNotificationsService.SendNotifications(sendNotificationsOptions)
if err != nil {
	panic(err)
}
```

<details open>
<summary>Send Notifications Variables</summary>
<br>

- **FCM Target NotificationFcmDevices** - Set up the the push notifications tragets.
  - *UserIds* (Array of **String**) - Send notification to the specified userIds.
  - *FcmDevices* (Array of **String**) - Send notification to the list of specified devices.
  - *Tags* (Array of **String**) - Send notification to the devices that have subscribed to any of
these tags.
  - *Platforms* (Array of **String**) - Send notification to the devices of the specified platforms. Pass 'G' for google (Android) devices.
- **Android Lights** - Allows setting the notification LED color on receiving push notification.
  - *LedArgb* (**String**) - The color of the led. The hardware will do its best approximation. Ex: `Red`
  - *LedOnMs* (**Integer**) - The number of milliseconds for the LED to be on while it's flashing. The hardware will do its best approximation.
  - *LedOffMs* (**String**) - The number of milliseconds for the LED to be off while it's flashing. The hardware will do its best approximation.
- **Android Style** - Options to specify for Android expandable notifications. The types of expandable notifications are *picture_notification*, *bigtext_notification*, and *inbox_notification*.
  - *Type* (**String**) - Specifies the type of expandable notifications. The possible values are *picture_notification*, *bigtext_notification*, and *inbox_notification*
  - *Title* (**String**) - Specifies the title of the notification. The title is displayed when the notification is expanded. Title must be specified for all three expandable notification.
  - *URL* (**String**) - An URL from which the picture has to be obtained for the notification. Must be specified for *picture_notification*.
  - *Text* (**String**) - The big text that needs to be displayed on expanding a *bigtext_notification*. Must be specified for *bigtext_notification*.
  - *Lines* (**String**) - An array of strings that is to be displayed in inbox style for inbox_notification. Must be specified for inbox_notification.
- **Android NotificationFcmBodyMessageData** - Settings specific to Android platform payload.
  - *Alert* (**String**) - The notification message to be shown to the user.
  - *CollapseKey* (**String**) - Dozed devices to display only the latest notification and discard old low priority notifications.
  - *InteractiveCategory* (**String**) - The category identifier to be used for the interactive push notifications.
  - *Icon* (**String**) - Specify the name of the icon to be displayed for the notification. Make sure the icon is already packaged with the client application.
  - *DelayWhileIdle* (**Bool**) - When this parameter is set to true, it indicates that the
message should not be sent until the device becomes active.
  - *Sync* (**Bool**) - Device group messaging makes it possible for every app instance in a group to reflect the latest messaging state
  - *Visibility* (**String**) - private/public - Visibility of this notification, which affects how and when the notifications are revealed on a secure locked screen.
  - *Redact* (**String**) - Content specified will show up on a secure locked screen on the device when visibility is set to Private
  - *Payload* (**map[string]interface{}**) - 	
Custom JSON payload that will be sent as part of the notification message.
  - *Priority* (**String**) - A string value that indicates the priority of this notification. Allowed values are 'max', 'high', 'default', 'low' and 'min'. High/Max priority notifications along with 'sound' field may be used for Heads up notification in Android 5.0 or higher.sampleval='low'.
  - *Sound* (**String**) - The sound file (on device) that will be attempted to play when the notification arrives on the device.
  - *TimeToLive* (**Integer**) - This parameter specifies how long (in seconds) the message
should be kept in GCM storage if the device is offline.
  - *Lights* (**Lights**) - Allows setting the notification LED color on receiving push notification.
  - *AndroidTitle* (**String**) - The title of Rich Push notifications.
  - *GroupID* (**String**) - Set this notification to be part of a group of notifications sharing the same key. Grouped notifications may display in a cluster or stack on devices which support such rendering.
  - *Style* (**Style**) - Options to specify for Android expandable notifications. The types of expandable notifications are *picture_notification*, *bigtext_notification*, and *inbox_notification*..
  - *Type* (**String**) - The notification Type value. The types are *DEFAULT* and *SILENT* .

- **FCM NotificationFcmBodyMessage** - Settings specific to Android platform data field.
  - *Data* (**NotificationFcmBodyMessageData**) - The `data` field for FCM notifications [Refer this FCM official [link](https://firebase.google.com/docs/cloud-messaging/concept-options)].

- **FCM NotificationFcmBody** - Settings specific to Android platform data field.
  - *Message* (**NotificationFcmBodyMessage**) - The `message` field for FCM notifications [Refer this FCM official [link](https://firebase.google.com/docs/cloud-messaging/concept-options)] 

- **Event Notificaitons SendNotificationsOptions** - Event Notificaitons Send notificaitons method. 
  - *InstanceID* (**String**) - Event Notificaitons instance AppGUID. 
  - *Subject* (**String**) - Subject for the notifications. 
  - *Severity* (**String**) - Severity for the notifications. 
  - *ID* (**ID**) - ID for the notifications. 
  - *Source* (**String**) - Source of the notifications. 
  - *EnSourceID* (**String**) - Event Notificaitons instance Source ID. 
  - *Type* (**String**) - Type for the notifications. 
  - *Time* (**String**) - Time of the notifications. 
  - *Data* (**map[string]interface{}**) - Data for the notifications. Supported only for `Webhook` destination. 
  - *PushTo* (**NotificationFcmDevices**) - Targets for the FCM notifications. 
  - *MessageFcmBody* (**NotificationFcmBodyMessage**) - Message body for the FCM notifications. 
  - *Datacontenttype* (**String**) - Data content type of the notifications. 
  - *Specversion* (**String**) - Spec version of the Event Notificaitons. Default value is `1.0`. 

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
