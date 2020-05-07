package authentication

import (
    log "log"

    proto "github.com/golang/protobuf/proto"

    axon_utils "github.com/jeroenvm/archetype-go-axon/src/pkg/axon_utils"
    grpc_example "github.com/jeroenvm/archetype-go-axon/src/pkg/grpc/example"
)

type Projection struct {
    Credentials map[string]string
}

type CredentialsAddedSourceEvent struct {
    event *grpc_example.CredentialsAddedEvent
}

type CredentialsRemovedSourceEvent struct {
    event *grpc_example.CredentialsRemovedEvent
}

func RestoreProjection(aggregateIdentifier string, clientConnection *axon_utils.ClientConnection) *Projection {
    projection := Projection{
        Credentials: make(map[string]string),
    }
    axon_utils.RestoreProjection("Authentication", aggregateIdentifier, projection, clientConnection, prepareUnmarshal)
    return &projection
}

func Apply(event axon_utils.SourceEvent, aggregateIdentifier string, clientConnection *axon_utils.ClientConnection) {
    projection := RestoreProjection(aggregateIdentifier, clientConnection)
    event.ApplyTo(projection)
}

func prepareUnmarshal(payloadType string) (sourceEvent axon_utils.SourceEvent) {
    log.Printf("Credentials Projection: Payload type: %v", payloadType)
    if (payloadType == "CredentialsAddedEvent") {
        sourceEvent = &CredentialsAddedSourceEvent{
            event: &grpc_example.CredentialsAddedEvent{},
        }
    } else if (payloadType == "CredentialsRemovedEvent") {
        sourceEvent = &CredentialsRemovedSourceEvent{
            event: &grpc_example.CredentialsRemovedEvent{},
        }
    }
    return sourceEvent
}

func (sourceEvent *CredentialsAddedSourceEvent) GetEvent() proto.Message {
    return sourceEvent.event
}
func (sourceEvent *CredentialsAddedSourceEvent) ApplyTo(projectionWrapper interface{}) {
    projection := projectionWrapper.(Projection)
    projection.Credentials[sourceEvent.event.Credentials.Identifier] = sourceEvent.event.Credentials.Secret
}

func (sourceEvent *CredentialsRemovedSourceEvent) GetEvent() proto.Message {
    return sourceEvent.event
}
func (sourceEvent *CredentialsRemovedSourceEvent) ApplyTo(projectionWrapper interface{}) {
    projection := projectionWrapper.(Projection)
    projection.Credentials[sourceEvent.event.Identifier] = ""
}
