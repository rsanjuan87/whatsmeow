package waMsgTransport

import (
	"github.com/rsanjuan87/whatsmeow/proto/armadilloutil"
	"github.com/rsanjuan87/whatsmeow/proto/instamadilloTransportPayload"
	"go.mau.fi/whatsmeow/proto/waMsgApplication"
	"github.com/rsanjuan87/whatsmeow/proto/armadilloutil"
	"github.com/rsanjuan87/whatsmeow/proto/waMsgApplication"
)

const (
	FBMessageApplicationVersion = 2
	IGMessageApplicationVersion = 3
)

func (msg *MessageTransport_Payload) DecodeFB() (*waMsgApplication.MessageApplication, error) {
	return armadilloutil.Unmarshal(&waMsgApplication.MessageApplication{}, msg.GetApplicationPayload(), FBMessageApplicationVersion)
}

func (msg *MessageTransport_Payload) DecodeIG() (*instamadilloTransportPayload.TransportPayload, error) {
	return armadilloutil.Unmarshal(&instamadilloTransportPayload.TransportPayload{}, msg.GetApplicationPayload(), IGMessageApplicationVersion)
}
