package eventnotificationsv1

import (
	"context"
	"encoding/json"

	common "github.com/IBM/event-notifications-go-admin-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
)

// SendNotifications : Send a notification
// Send a notification.
func (eventNotifications *EventNotificationsV1) SendNotifications(sendNotificationsOptions *SendNotificationsOptions) (result *NotificationResponse, response *core.DetailedResponse, err error) {
	return eventNotifications.SendNotificationsWithContext(context.Background(), sendNotificationsOptions)
}

// SendNotificationsWithContext is an alternate form of the SendNotifications method which supports a Context parameter
func (eventNotifications *EventNotificationsV1) SendNotificationsWithContext(ctx context.Context, sendNotificationsOptions *SendNotificationsOptions) (result *NotificationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(sendNotificationsOptions, "sendNotificationsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(sendNotificationsOptions, "sendNotificationsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *sendNotificationsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = eventNotifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(eventNotifications.Service.Options.URL, `/v1/instances/{instance_id}/notifications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range sendNotificationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("event_notifications", "V1", "SendNotifications")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if sendNotificationsOptions.Ibmenseverity != nil {
		body["ibmenseverity"] = sendNotificationsOptions.Ibmenseverity
	}
	if sendNotificationsOptions.Ibmensourceid != nil {
		body["ibmensourceid"] = sendNotificationsOptions.Ibmensourceid
	}
	if sendNotificationsOptions.Subject != nil {
		body["subject"] = sendNotificationsOptions.Subject
	}
	if sendNotificationsOptions.ID != nil {
		body["id"] = sendNotificationsOptions.ID
	}
	if sendNotificationsOptions.Source != nil {
		body["source"] = sendNotificationsOptions.Source
	}
	if sendNotificationsOptions.Type != nil {
		body["type"] = sendNotificationsOptions.Type
	}
	if sendNotificationsOptions.Time != nil {
		body["time"] = sendNotificationsOptions.Time
	}
	if sendNotificationsOptions.Data != nil {
		body["data"] = sendNotificationsOptions.Data
	}
	if sendNotificationsOptions.Ibmenfcmbody != nil {
		ibmenfcmbody, _ := json.Marshal(sendNotificationsOptions.Ibmenfcmbody)
		body["ibmenfcmbody"] = string(ibmenfcmbody)
	}
	if sendNotificationsOptions.Ibmenapnsbody != nil {
		ibmenapnsbody, _ := json.Marshal(sendNotificationsOptions.Ibmenapnsbody)
		body["ibmenapnsbody"] = string(ibmenapnsbody)
	}
	if sendNotificationsOptions.Ibmenpushto != nil {
		ibmenpushto, _ := json.Marshal(sendNotificationsOptions.Ibmenpushto)
		body["ibmenpushto"] = string(ibmenpushto)
	}
	if sendNotificationsOptions.Ibmenapnsheaders != nil {
		ibmenapnsheaders, _ := json.Marshal(sendNotificationsOptions.Ibmenapnsheaders)
		body["ibmenapnsheaders"] = string(ibmenapnsheaders)
	}
	body["datacontenttype"] = "application/json"
	if sendNotificationsOptions.Datacontenttype != nil {
		body["datacontenttype"] = sendNotificationsOptions.Datacontenttype
	}
	body["specversion"] = "1.0"
	if sendNotificationsOptions.Specversion != nil {
		body["specversion"] = sendNotificationsOptions.Specversion
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = eventNotifications.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNotificationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SendNotificationsOptions : The SendNotifications options.
type SendNotificationsOptions struct {
	// Unique identifier for IBM Cloud Event Notifications instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The Notifications id.
	Ibmenseverity *string `json:"ibmenseverity" validate:"required"`

	// The Event Notifications source id.
	Ibmensourceid *string `json:"ibmensourceid" validate:"required"`

	// The Notifications subject.
	Subject *string `json:"subject" validate:"required"`

	// The Notifications id.
	ID *string `json:"id" validate:"required"`

	// The source of Notifications.
	Source *string `json:"source" validate:"required"`

	// The Notifications type.
	Type *string `json:"type" validate:"required"`

	// The Notifications time.
	Time *strfmt.DateTime `json:"time" validate:"required"`

	// The Notifications data for webhook.
	Data map[string]interface{} `json:"data,omitempty"`

	Ibmenfcmbody NotificationFcmBodyIntf `json:"ibmenfcmbody,omitempty"`

	// Payload describing a APNs Notifications body.
	Ibmenapnsbody NotificationApnsBodyIntf `json:"ibmenapnsbody,omitempty"`

	// Payload describing a FCM Notifications targets.
	Ibmenpushto *NotificationDevices `json:"ibmenpushto,omitempty"`

	// The attributes for an FCM/APNs notification.
	Ibmenapnsheaders map[string]interface{} `json:"ibmenapnsheaders,omitempty"`

	// The Notifications content type.
	Datacontenttype *string `json:"datacontenttype,omitempty"`

	// The Notifications specversion.
	Specversion *string `json:"specversion,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSendNotificationsOptions : Instantiate SendNotificationsOptions
func (*EventNotificationsV1) NewSendNotificationsOptions(instanceID string, ibmenseverity string, ibmensourceid string, subject string, id string, source string, typeVar string, time *strfmt.DateTime) *SendNotificationsOptions {
	return &SendNotificationsOptions{
		InstanceID:    core.StringPtr(instanceID),
		Ibmenseverity: core.StringPtr(ibmenseverity),
		Ibmensourceid: core.StringPtr(ibmensourceid),
		Subject:       core.StringPtr(subject),
		ID:            core.StringPtr(id),
		Source:        core.StringPtr(source),
		Type:          core.StringPtr(typeVar),
		Time:          time,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *SendNotificationsOptions) SetInstanceID(instanceID string) *SendNotificationsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetIbmenseverity : Allow user to set Ibmenseverity
func (_options *SendNotificationsOptions) SetIbmenseverity(ibmenseverity string) *SendNotificationsOptions {
	_options.Ibmenseverity = core.StringPtr(ibmenseverity)
	return _options
}

// SetIbmensourceid : Allow user to set Ibmensourceid
func (_options *SendNotificationsOptions) SetIbmensourceid(ibmensourceid string) *SendNotificationsOptions {
	_options.Ibmensourceid = core.StringPtr(ibmensourceid)
	return _options
}

// SetSubject : Allow user to set Subject
func (_options *SendNotificationsOptions) SetSubject(subject string) *SendNotificationsOptions {
	_options.Subject = core.StringPtr(subject)
	return _options
}

// SetID : Allow user to set ID
func (_options *SendNotificationsOptions) SetID(id string) *SendNotificationsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetSource : Allow user to set Source
func (_options *SendNotificationsOptions) SetSource(source string) *SendNotificationsOptions {
	_options.Source = core.StringPtr(source)
	return _options
}

// SetType : Allow user to set Type
func (_options *SendNotificationsOptions) SetType(typeVar string) *SendNotificationsOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetTime : Allow user to set Time
func (_options *SendNotificationsOptions) SetTime(time *strfmt.DateTime) *SendNotificationsOptions {
	_options.Time = time
	return _options
}

// SetData : Allow user to set Data
func (_options *SendNotificationsOptions) SetData(data map[string]interface{}) *SendNotificationsOptions {
	_options.Data = data
	return _options
}

// SetIbmenfcmbody : Allow user to set Ibmenfcmbody
func (_options *SendNotificationsOptions) SetIbmenfcmbody(ibmenfcmbody NotificationFcmBodyIntf) *SendNotificationsOptions {
	_options.Ibmenfcmbody = ibmenfcmbody
	return _options
}

// SetIbmenapnsbody : Allow user to set Ibmenapnsbody
func (_options *SendNotificationsOptions) SetIbmenapnsbody(ibmenapnsbody NotificationApnsBodyIntf) *SendNotificationsOptions {
	_options.Ibmenapnsbody = ibmenapnsbody
	return _options
}

// SetIbmenpushto : Allow user to set Ibmenpushto
func (_options *SendNotificationsOptions) SetIbmenpushto(ibmenpushto *NotificationDevices) *SendNotificationsOptions {
	_options.Ibmenpushto = ibmenpushto
	return _options
}

// SetIbmenapnsheaders : Allow user to set Ibmenapnsheaders
func (_options *SendNotificationsOptions) SetIbmenapnsheaders(ibmenapnsheaders map[string]interface{}) *SendNotificationsOptions {
	_options.Ibmenapnsheaders = ibmenapnsheaders
	return _options
}

// SetDatacontenttype : Allow user to set Datacontenttype
func (_options *SendNotificationsOptions) SetDatacontenttype(datacontenttype string) *SendNotificationsOptions {
	_options.Datacontenttype = core.StringPtr(datacontenttype)
	return _options
}

// SetSpecversion : Allow user to set Specversion
func (_options *SendNotificationsOptions) SetSpecversion(specversion string) *SendNotificationsOptions {
	_options.Specversion = core.StringPtr(specversion)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SendNotificationsOptions) SetHeaders(param map[string]string) *SendNotificationsOptions {
	options.Headers = param
	return options
}
