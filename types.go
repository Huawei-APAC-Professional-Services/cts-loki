package main

/*
{
    "Records":[
        {
            "eventVersion":"", //Version number. The current version is 3.0.
            "eventSource":"", //Message source. The value is fixed to OBS.
            "eventRegion":"", //Region where the event occurs
            "eventTime":"", //Time when an event occurs, in the ISO-8601 format, for example, 2020-07-10T09:24:11.418Z
            "eventName":"", //Name of the event that triggers the notification
            "userIdentity":{
                "ID":"" //Billing ID of the user who triggers the event
            },
            "requestParameters":{
                "sourceIPAddress":"" //Source IP address of the request
            },
            "responseElements":{
                "x-obs-request-id":"", //ID of the request
                "x-obs-id-2":"" ///Special characters for locating problems
            },
            "obs":{
                "Version":"1.0",
                "configurationId":"", //Name of the event notification rule in OBS that matches the event
                "bucket":{
                    "name":"examplebucket",
                    "ownerIdentity":{
                        "ID":"" //Account ID of the bucket owner
                    },
                    "bucket":"" //Bucket name
                },
               "object":{
                    "key":"", //Object name
                    "eTag":"", //ETag of the object
                    "size": , //Object size
                    "versionId":"null", //Version ID of the object
                    "sequencer":""  //Identifier that defines the event sequence of a specific object
                }
            }
        }
    ]
}
*/

type OBSNotification struct {
	EventVersion string `json:"eventVersion"`
	EventSource  string `json:"eventSource"`
	EventRegion  string `json:"eventRegion"`
	EventTime    string `json:"eventTime"`
	EventName    string `json:"eventName"`
	UserIdentity struct {
		ID string `json:"ID"`
	} `json:"userIdentity"`
	RequestParameters struct {
		SourceIPAddress string `json:"sourceIPAddress"`
	} `json:"requestParameters"`
	ResponseElements struct {
		RequestId string `json:"x-obs-request-id"`
		OBSId     string `json:"x-obs-id-2"`
	} `json:"responseElements"`
	OBS struct {
		Version         string `json:"Version"`
		ConfigurationId string `json:"configurationId"`
		Bucket          struct {
			Name          string `json:"name"`
			OwnerIdentity struct {
				ID string `json:"ID"`
			} `json:"ownerIdentity"`
			Bucket string `json:"bucket"`
		} `json:"bucket"`
		Object struct {
			Key       string  `json:"key"`
			Etag      string  `json:"eTag"`
			Size      float64 `json:"size"`
			VersionId string  `json:"versionId"`
			Sequencer string  `json:"sequencer"`
		} `json:"object"`
	} `json:"obs"`
}

type SMNEvent struct {
	Records []OBSNotification `json:"Records"`
}
