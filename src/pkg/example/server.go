package example

import (
    context "context"
    errors "errors"
    fmt "fmt"
    io "io"
    log "log"
    net "net"
    time "time"

    hex "encoding/hex"
    rand "crypto/rand"

    grpc "google.golang.org/grpc"
    jwt "github.com/pascaldekloe/jwt"
    proto "github.com/golang/protobuf/proto"
    reflection "google.golang.org/grpc/reflection"
    uuid "github.com/google/uuid"

    authentication "github.com/jeroenvm/archetype-go-axon/src/pkg/authentication"
    axonserver "github.com/jeroenvm/archetype-go-axon/src/pkg/grpc/axonserver"
    grpcExample "github.com/jeroenvm/archetype-go-axon/src/pkg/grpc/example"
    trusted "github.com/jeroenvm/archetype-go-axon/src/pkg/trusted"
)

type GreeterServer struct {
    conn *grpc.ClientConn;
    clientInfo *axonserver.ClientIdentification;
}

func (s *GreeterServer) Greet(c context.Context, greeting *grpcExample.Greeting) (*grpcExample.Acknowledgement, error) {
    message := (*greeting).Message
    log.Printf("Server: Received greeting: %v", message)
    ack := grpcExample.Acknowledgement{
        Message: "Good day to you too!",
    }
    command := grpcExample.GreetCommand {
        AggregateIdentifier: "single_aggregate",
        Message: greeting,
    }
    data, err := proto.Marshal(&command)
    if err != nil {
        log.Printf("Server: Error while marshalling command")
        return nil, errors.New("Marshalling error")
    }
    serializedCommand := axonserver.SerializedObject{
        Type: "GreetCommand",
        Data: data,
    }
    err = SubmitCommand(&serializedCommand, s.conn, s.clientInfo)
    if err != nil {
        return nil, err
    }
    return &ack, nil
}

func (s *GreeterServer) Greetings(empty *grpcExample.Empty, greetingsServer grpcExample.GreeterService_GreetingsServer) error {
    greeting := grpcExample.Greeting {
        Message: "Hello, World!",
    }
    log.Printf("Server: Greetings streamed reply: %v", greeting)
    greetingsServer.Send(&greeting)
    log.Printf("Server: Greetings streamed reply sent")

    eventStoreClient := axonserver.NewEventStoreClient(s.conn)
    requestEvents := axonserver.GetAggregateEventsRequest {
        AggregateId: "single_aggregate",
        InitialSequence: 0,
        AllowSnapshots: false,
    }
    log.Printf("Server: Request events: %v", requestEvents)
    client, e := eventStoreClient.ListAggregateEvents(context.Background(), &requestEvents)
    if e != nil {
        log.Printf("Server: Error while requesting aggregate events: %v", e)
        return nil
    }
    for {
        eventMessage, e := client.Recv()
        if e == io.EOF {
            log.Printf("Server: End of stream")
            break
        } else if e != nil {
            log.Printf("Server: Error while receiving next event: %v", e)
            break
        }
        log.Printf("Server: Received event: %v", eventMessage)
        if eventMessage.Payload != nil {
            log.Printf("Server: Payload type: %v", eventMessage.Payload.Type)
            if (eventMessage.Payload.Type != "GreetedEvent") {
                continue
            }
            log.Printf("Server: Payload: %v", eventMessage.Payload)
            event := grpcExample.GreetedEvent{}
            e = proto.Unmarshal(eventMessage.Payload.Data, &event)
            if e != nil {
                log.Printf("Server: Error while unmarshalling GreetedEvent")
                continue
            }
            log.Printf("Server: GreetedEvent: %v", event)
            log.Printf("Server: Greetings streamed reply: %v", event.Message)
            greetingsServer.Send(event.Message)
            log.Printf("Server: Greetings streamed reply sent")
        }
    }

    return nil
}

var empty = grpcExample.Empty{}

func (s *GreeterServer) Record(c context.Context, greeting *grpcExample.Empty) (*grpcExample.Empty, error) {
    command := grpcExample.RecordCommand {
        AggregateIdentifier: "single_aggregate",
    }
    data, err := proto.Marshal(&command)
    if err != nil {
        log.Printf("Server: Error while marshalling command")
        return nil, errors.New("Marshalling error")
    }
    serializedCommand := axonserver.SerializedObject{
        Type: "RecordCommand",
        Data: data,
    }
    err = SubmitCommand(&serializedCommand, s.conn, s.clientInfo)
    if err != nil {
        return nil, err
    }
    return &empty, nil
}

func (s *GreeterServer) Stop(c context.Context, greeting *grpcExample.Empty) (*grpcExample.Empty, error) {
    command := grpcExample.StopCommand {
        AggregateIdentifier: "single_aggregate",
    }
    data, err := proto.Marshal(&command)
    if err != nil {
        log.Printf("Server: Error while marshalling command")
        return nil, errors.New("Marshalling error")
    }
    serializedCommand := axonserver.SerializedObject{
        Type: "StopCommand",
        Data: data,
    }
    err = SubmitCommand(&serializedCommand, s.conn, s.clientInfo)
    if err != nil {
        return nil, err
    }
    return &empty, nil
}

func (s *GreeterServer) Search(query *grpcExample.SearchQuery, greetingsServer grpcExample.GreeterService_SearchServer) error {
    greeting := grpcExample.Greeting {
        Message: "Hello, World!",
    }
    log.Printf("Server: Search streamed reply: %v", greeting)
    greetingsServer.Send(&greeting)
    log.Printf("Server: Search streamed reply sent")

    queryData, e := proto.Marshal(query)
    if e != nil {
        log.Printf("Server: Error while marshalling query object: %v", e)
        return e
    }

    serializedQuery := axonserver.SerializedObject{
        Type: "SearchQuery",
        Data: queryData,
    }

    eventStoreClient := axonserver.NewQueryServiceClient(s.conn)
    id := uuid.New()
    queryRequest := axonserver.QueryRequest {
        MessageIdentifier: id.String(),
        Query: "SearchQuery",
        Payload: &serializedQuery,
    }
    log.Printf("Server: Query request: %v", queryRequest)
    client, e := eventStoreClient.Query(context.Background(), &queryRequest)
    if e != nil {
        log.Printf("Server: Error while submitting query: %v", e)
        return nil
    }
    for {
        queryResponse, e := client.Recv()
        if e == io.EOF {
            log.Printf("Server: End of stream")
            break
        } else if e != nil {
            log.Printf("Server: Error while receiving next event: %v", e)
            break
        }
        log.Printf("Server: Received query response: %v", queryResponse)
        if queryResponse.Payload != nil {
            log.Printf("Server: Payload type: %v", queryResponse.Payload.Type)
            if (queryResponse.Payload.Type != "Greeting") {
                continue
            }
            log.Printf("Server: Payload: %v", queryResponse.Payload)
            greeting := grpcExample.Greeting{}
            e = proto.Unmarshal(queryResponse.Payload.Data, &greeting)
            if e != nil {
                log.Printf("Server: Error while unmarshalling Greeting")
                continue
            }
            log.Printf("Server: Greeting: %v", greeting)
            greetingsServer.Send(&greeting)
            log.Printf("Server: Search streamed reply sent")
        }
    }

    return nil
}

func (s *GreeterServer) Authorize(ctx context.Context, credentials *grpcExample.Credentials) (*grpcExample.AccessToken, error) {
    accessToken := grpcExample.AccessToken{
        Jwt: "",
    }
    if authentication.Authenticate(credentials.Identifier, credentials.Secret) {
        var claims jwt.Claims
        claims.Subject = credentials.Identifier
        claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
        token, e := trusted.CreateJWT(claims)
        if e != nil {
            return nil, e
        }
        accessToken.Jwt = token
    }
    return &accessToken, nil
}

func (s *GreeterServer) ListTrustedKeys(req *grpcExample.Empty, streamServer grpcExample.GreeterService_ListTrustedKeysServer) error {
    trustedKey := grpcExample.PublicKey {}
    for name, key := range trusted.TrustedKeys {
        trustedKey.Name = name
        trustedKey.PublicKey = key
        log.Printf("Server: Trusted keys streamed reply: %v", trustedKey)
        streamServer.Send(&trustedKey)
        log.Printf("Server: Trusted keys streamed reply sent")
    }
    return nil
}

func (s *GreeterServer) SetPrivateKey(ctx context.Context, request *grpcExample.PrivateKey) (*grpcExample.Empty, error) {
    trusted.SetPrivateKey(request.Name, request.PrivateKey)

    empty := grpcExample.Empty{}
    return &empty, nil
}

func (s *GreeterServer) ChangeTrustedKeys(stream grpcExample.GreeterService_ChangeTrustedKeysServer) error {
    var status = grpcExample.Status{}
    response := grpcExample.TrustedKeyResponse{}
    nonce := make([]byte, 64)
    first := true
    for true {
        request, e := stream.Recv();
        if e != nil {
            log.Printf("Server: Change trusted keys: error receiving request: %v", e)
            return e
        }

        status.Code = 500
        status.Message = "Internal Server Error"

        if first {
            first = false
            status.Code = 200
            status.Message = "OK"
        } else {
            if request.Signature == nil {
                status.Code = 200
                status.Message = "End of stream"
                response.Status = &status
                response.Nonce = nil
                _ = stream.Send(&response)
                return nil
            }
            e = trusted.AddTrustedKey(request, nonce)
            if e == nil {
                status.Code = 200
                status.Message = "OK"
            } else {
                status.Code = 400
                status.Message = e.Error()
            }
        }

        rand.Reader.Read(nonce)
        hexNonce := hex.EncodeToString(nonce)
        log.Printf("Next nonce: %v", hexNonce)

        response.Status = &status
        response.Nonce = nonce
        e = stream.Send(&response)
        if e != nil {
            log.Printf("Server: Change trusted keys: error sending response: %v", e)
            return e
        }
    }
    return errors.New("Server: Change trusted keys: unexpected end of stream")
}

func (s *GreeterServer) ChangeCredentials(stream grpcExample.GreeterService_ChangeCredentialsServer) error {
    for true {
        credentials, e := stream.Recv()
        if e != nil {
            log.Printf("Error while receiving credentials: %v", e)
            return e
        }
        if credentials.Signature == nil {
            break
        }
        authentication.SetCredentials(credentials)
    }
    empty = grpcExample.Empty{}
    return stream.SendAndClose(&empty)
}

func Serve(conn *grpc.ClientConn, clientInfo *axonserver.ClientIdentification) {
    port := 8181
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        log.Fatalf("Server: Failed to listen: %v", err)
        panic("Server: Panic!")
    }
    log.Printf("Server: Listening on port: %d", port)
    grpcServer := grpc.NewServer()
    grpcExample.RegisterGreeterServiceServer(grpcServer, &GreeterServer{conn,clientInfo})
    reflection.Register(grpcServer)
    // ... // determine whether to use TLS
    grpcServer.Serve(lis)
}