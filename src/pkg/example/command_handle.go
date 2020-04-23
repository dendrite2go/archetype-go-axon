package example

import (
    context "context"
    log "log"

    grpc "google.golang.org/grpc"
    proto "github.com/golang/protobuf/proto"

    authentication "github.com/jeroenvm/archetype-go-axon/src/pkg/authentication"
    axonserver "github.com/jeroenvm/archetype-go-axon/src/pkg/grpc/axonserver"
    axon_utils "github.com/jeroenvm/archetype-go-axon/src/pkg/axon_utils"
    grpcExample "github.com/jeroenvm/archetype-go-axon/src/pkg/grpc/example"
    trusted "github.com/jeroenvm/archetype-go-axon/src/pkg/trusted"
)

func HandleCommands(host string, port int) (conn *grpc.ClientConn) {
    conn, clientInfo, _ := axon_utils.WaitForServer(host, port, "Command Handler")

    log.Printf("Command handler: Connection: %v", conn)
    client := axonserver.NewCommandServiceClient(conn)
    log.Printf("Command handler: Client: %v", client)

    stream, e := client.OpenStream(context.Background())
    log.Printf("Command handler: Stream: %v: %v", stream, e)

    axon_utils.SubscribeCommand("GreetCommand", stream, clientInfo)
    axon_utils.SubscribeCommand("RecordCommand", stream, clientInfo)
    axon_utils.SubscribeCommand("StopCommand", stream, clientInfo)
    axon_utils.SubscribeCommand("RegisterTrustedKeyCommand", stream, clientInfo)
    axon_utils.SubscribeCommand("RegisterKeyManagerCommand", stream, clientInfo)
    axon_utils.SubscribeCommand("RegisterCredentialsCommand", stream, clientInfo)

    go commandWorker(stream, conn, clientInfo.ClientId)

    return conn;
}

func commandWorker(stream axonserver.CommandService_OpenStreamClient, conn *grpc.ClientConn, clientId string) {
    for true {
        axon_utils.CommandAddPermits(1, stream, clientId)

        log.Printf("Command handler: Waiting for command")
        inbound, e := stream.Recv()
        if (e != nil) {
          log.Printf("Command handler: Error on receive: %v", e)
          break
        }
        log.Printf("Command handler: Inbound: %v", inbound)
        command := inbound.GetCommand()
        if (command != nil) {
            commandName := command.Name
            log.Printf("Received %v", commandName)
            if (commandName == "GreetCommand") {
                handleGreetCommand(command, stream, conn)
            } else if (commandName == "RecordCommand") {
                handleRecordCommand(command, stream, conn)
            } else if (commandName == "StopCommand") {
                handleStopCommand(command, stream, conn)
            } else if (commandName == "RegisterTrustedKeyCommand") {
                trusted.HandleRegisterTrustedKeyCommand(command, stream, conn)
            } else if (commandName == "RegisterKeyManagerCommand") {
                trusted.HandleRegisterKeyManagerCommand(command, stream, conn)
            } else if (commandName == "RegisterCredentialsCommand") {
                authentication.HandleRegisterCredentialsCommand(command, stream, conn)
            } else {
                log.Printf("Received unknown command: %v", commandName)
            }
        }
    }
}

func handleGreetCommand(command *axonserver.Command, stream axonserver.CommandService_OpenStreamClient, conn *grpc.ClientConn) {
    deserializedCommand := grpcExample.GreetCommand{}
    e := proto.Unmarshal(command.Payload.Data, &deserializedCommand)
    if (e != nil) {
        log.Printf("Could not unmarshal GreetCommand")
    }

    projection := RestoreProjection(deserializedCommand.AggregateIdentifier, conn)
    if !projection.Recording {
        axon_utils.ReportError(stream, command.MessageIdentifier, "EX001", "Not recording: " + deserializedCommand.AggregateIdentifier)
        return
    }

    event := grpcExample.GreetedEvent {
        Message: deserializedCommand.Message,
    }
    data, err := proto.Marshal(&event)
    if err != nil {
        log.Printf("Server: Error while marshalling event")
        return
    }
    serializedEvent := axonserver.SerializedObject{
        Type: "GreetedEvent",
        Data: data,
    }

    axon_utils.AppendEvent(&serializedEvent, deserializedCommand.AggregateIdentifier, conn)
    axon_utils.CommandRespond(stream, command.MessageIdentifier)
}

func handleRecordCommand(command *axonserver.Command, stream axonserver.CommandService_OpenStreamClient, conn *grpc.ClientConn) {
    deserializedCommand := grpcExample.RecordCommand{}
    e := proto.Unmarshal(command.Payload.Data, &deserializedCommand)
    if (e != nil) {
        log.Printf("Could not unmarshal RecordCommand")
    }
    event := grpcExample.StartedRecordingEvent {}
    data, err := proto.Marshal(&event)
    if err != nil {
        log.Printf("Server: Error while marshalling event")
        return
    }
    serializedEvent := axonserver.SerializedObject{
        Type: "StartedRecordingEvent",
        Data: data,
    }

    axon_utils.AppendEvent(&serializedEvent, deserializedCommand.AggregateIdentifier, conn)
    axon_utils.CommandRespond(stream, command.MessageIdentifier)
}

func handleStopCommand(command *axonserver.Command, stream axonserver.CommandService_OpenStreamClient, conn *grpc.ClientConn) {
    deserializedCommand := grpcExample.StopCommand{}
    e := proto.Unmarshal(command.Payload.Data, &deserializedCommand)
    if (e != nil) {
        log.Printf("Could not unmarshal StopCommand")
    }
    event := grpcExample.StoppedRecordingEvent {}
    data, err := proto.Marshal(&event)
    if err != nil {
        log.Printf("Server: Error while marshalling event")
        return
    }
    serializedEvent := axonserver.SerializedObject{
        Type: "StoppedRecordingEvent",
        Data: data,
    }

    axon_utils.AppendEvent(&serializedEvent, deserializedCommand.AggregateIdentifier, conn)
    axon_utils.CommandRespond(stream, command.MessageIdentifier)
}
