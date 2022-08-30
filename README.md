
# IBM Cloud Event Notifications Go Admin SDK 0.1.5
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
To configure service URL for Private Endpoint

```go
options := &eventnotificationsv1.EventNotificationsV1Options{
		Authenticator: authenticator,
		URL:           "https://private." + region + ".event-notifications.cloud.ibm.com/event-notifications",
	}
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
- [Push Destination APIs](#push-destination-apis)
	- [Create Destination tag subscription](#create-destination-tag-subscription)
	- [List Destination tag subscription](#list-destination-tag-subscription)
	- [List Destination device tag subscriptions](#list-destination-device-tag-subscriptions)
	- [Get Device Count](#get-device-count)
	- [Delete Destination device tag subscription](#delete-destination-device-tag-subscription)
- [Subscriptions](#subscriptions)
	- [Create Subscription](#create-subscription)
	- [List Subscriptions](#list-subscriptions)
	- [Get Subscription](#get-subscription)
	- [Update Subscription](#update-subscription)
	- [Delete Subscription](#delete-subscription)
- [Send Notifications](#send-notifications)
- [Get Device Count](#get-device-count)


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
### Get Device Count

```go
getDeviceCountOptions := &eventnotificationsv1.GetDeviceCountOptions{
	<instance-id>,		// Event notifications service instance GUID
	<destination-id>,	// Event notifications service instance Destination ID
}

deviceCount, response, err := eventNotificationsService.GetDeviceCount(getDeviceCountOptions)
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
notificationFcmBodyModel := "{\"title\" : \"Portugal vs. Denmark\", \"badge\": \"great match\"}"
notificationAPNsBodyModel := "{\"alert\": \"Game Request\", \"badge\": 5 }"
notificationSafariBodyModel := "{\"aps\":{\"alert\":{\"title\":\"FlightA998NowBoarding\",\"body\":\"BoardinghasbegunforFlightA998.\",\"action\":\"View\"},\"url-args\":[\"boarding\",\"A998\"]}}}"

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

- **Ibmenpushto** - Set up the the push notifications tragets.
  - *user_ids* (Array of **String**) - Send notification to the specified userIds.
  - *fcm_devices* (Array of **String**) - Send notification to the list of specified Android devices.
  - *fcm_devices* (Array of **String**) - Send notification to the list of specified iOS devices.
  - *_devices* (Array of **String**) - Send notification to the list of specified Chrome devices.
  - *firefox_devices* (Array of **String**) - Send notification to the list of specified Firefox devices.
  - *tags* (Array of **String**) - Send notification to the devices that have subscribed to any of these tags.
  - *platforms* (Array of **String**) - Send notification to the devices of the specified platforms. 
  	- Pass 'G' for google (Android) devices.
	- Pass 'A' for iOS devices.
	- Pass 'WEB_FIREFOX' for Firefox browser.
	- Pass 'WEB_CHROME' for Chrome browser.
- **ibmenfcmbody** - Set payload specific to Android platform [Refer this FCM official [link](https://firebase.google.com/docs/cloud-messaging/http-server-ref#notification-payload-support)].
- **ibmenfcmbody** - Set payload specific to iOS platform [Refer this APNs official doc [link](https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CreatingtheNotificationPayload.html)].
- **ibmenapnsheaders** - Set headers required for the APNs message [Refer this APNs official [link](https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/sending_notification_requests_to_apns)(Table 1 Header fields for a POST request)].
- **Event Notificaitons SendNotificationsOptions** - Event Notificaitons Send notificaitons method. 
  - *InstanceID* (**String**) - Event Notificaitons instance AppGUID. 
  - *ibmenseverity* (**String**) - Severity for the notifications. Some sources can have the concept of an Event severity. Hence a handy way is provided to specify a severity of the event.
  - *ID* (**String**) - A unique identifier that identifies each event. source+id must be unique. The backend should be able to uniquely track this id in logs and other records. Send unique ID for each send notification. Same ID can be sent in case of failure of send notification. source+id will be logged in IBM Cloud Logging service. Using this combination we will be able to trace the event movement from one system to another and will aid in debugging and tracing.
  - *Source* (**String**) - This is the identifier of the event producer. A way to uniquely identify the source of the event. For IBM Cloud services this is the crn of the service instance producing the events. For API sources this can be something the event producer backend can uniquely identify itself with.
  - *Ibmensourceid* (**String**) - This is the ID of the source created in EN. This is available in the EN UI in the "Sources" section.
  - *Type* (**String**) - This describes the type of event. It is of the form <event-type-name>:<sub-type> This type is defined by the producer. The event type name has to be prefixed with the reverse DNS names so the event type is uniquely identified. The same event type can be produced by 2 different sources. It is highly recommended to use hyphen - as a separator instead of _.
  - *Time* (**String**) - Time of the notifications. UTC time stamp when the event occurred. Must be in the RFC 3339 format.
  - *Ibmenpushto* (**string**) - Targets for the FCM notifications. This contains details about the destination where you want to send push notification. This attribute is mandatory for successful delivery from an Android FCM or APNS destination
  - *Ibmenfcmbody* (**string**) - Message body for the FCM notifications. [Refer this FCM official [link](https://firebase.google.com/docs/cloud-messaging/http-server-ref#notification-payload-support)].
  - *Ibmenapnsbody* (**string**) - Message body for the APNs notifications. [Refer this APNs official doc [link](https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CreatingtheNotificationPayload.html)].
  - *Ibmensafaribody* (**string**) - Message body for the Safari notifications. [Refer this Safari official doc [link](https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CreatingtheNotificationPayload.html)].
  - *Ibmenapnsheaders* (**string**) - Headers for the APNs notifications. [Refer this APNs official [link](https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/sending_notification_requests_to_apns)(Table 1 Header fields for a POST request)]
  - *Ibmenchromebody* (**string**) - Message body for the Chrome notifications. [Refer this APNs official [link](https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/sending_notification_requests_to_apns)(Table 1 Header fields for a POST request)]
  - *Ibmenfirefoxbody* (**string**) - Message body for the Firefox notifications. [Refer [this official documentation](https://developer.mozilla.org/en-US/docs/Web/API/Notification/Notification) for more].
  - *Ibmenchromeheaders* (**string**) - Headers for the Chrome notifications. Refer [this official documentation](https://developer.mozilla.org/en-US/docs/Web/API/Notification/Notification) for more.
  - *Ibmenfirefoxheaders* (**string**) - Headers for the Firefox notifications. Refer [this official documentation](https://developer.mozilla.org/en-US/docs/Web/API/Notification/Notification) for more.
  - *Ibmendefaultshort* (**string**) - Default short text for the message.
  - *Ibmendefaultlong* (**string**) - Default long text for the message.
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
