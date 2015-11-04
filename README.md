The code in this repository is an implementation example of a Liquid Handler grpc interface for the Antha OS.

The implementation is both based on stdin/stdout interaction via simple input and using the [gocui](https://github.com/jroimartin/gocui).

[Server](server/) Usage:

```
go run main.go manual.go
```

A [sample Client](sampleClient/) is provided, usage:

```
go run main.go
```

Implementation
==============

[Protocol Buffers](https://developers.google.com/protocol-buffers) are used for communication from AnthaOS to each device.
[openAntha] (github.com/antha-lang/openAntha) repository defines two interfaces for a liquidHandler; 
[LiquidhandlingDriver and ExtendedLiquidhandlingDriver](github.com/antha-lang/openAntha/microArch/driver/liquidhandling/driver.go)

The given .proto file fulfills ExtendedLiquidHandling interface exposing the following rpc methods.

```
rpc AddPlateTo (AddPlateToRequest) returns (AddPlateToReply) {}
rpc Aspirate (AspirateRequest) returns (AspirateReply) {}
rpc Close (CloseRequest) returns (CloseReply) {}
rpc Dispense (DispenseRequest) returns (DispenseReply) {}
rpc Finalize (FinalizeRequest) returns (FinalizeReply) {}
rpc GetCapabilities (GetCapabilitiesRequest) returns (GetCapabilitiesReply) {}
rpc GetCurrentPosition (GetCurrentPositionRequest) returns (GetCurrentPositionReply) {}
rpc GetHeadState (GetHeadStateRequest) returns (GetHeadStateReply) {}
rpc GetPositionState (GetPositionStateRequest) returns (GetPositionStateReply) {}
rpc GetStatus (GetStatusRequest) returns (GetStatusReply) {}
rpc Go (GoRequest) returns (GoReply) {}
rpc Initialize (InitializeRequest) returns (InitializeReply) {}
rpc LightsOff (LightsOffRequest) returns (LightsOffReply) {}
rpc LightsOn (LightsOnRequest) returns (LightsOnReply) {}
rpc LoadAdaptor (LoadAdaptorRequest) returns (LoadAdaptorReply) {}
rpc LoadHead (LoadHeadRequest) returns (LoadHeadReply) {}
rpc LoadTips (LoadTipsRequest) returns (LoadTipsReply) {}
rpc Message (MessageRequest) returns (MessageReply) {}
rpc Mix (MixRequest) returns (MixReply) {}
rpc Move (MoveRequest) returns (MoveReply) {}
rpc MoveRaw (MoveRawRequest) returns (MoveRawReply) {}
rpc Open (OpenRequest) returns (OpenReply) {}
rpc RemoveAllPlates (RemoveAllPlatesRequest) returns (RemoveAllPlatesReply) {}
rpc RemovePlateAt (RemovePlateAtRequest) returns (RemovePlateAtReply) {}
rpc ResetPistons (ResetPistonsRequest) returns (ResetPistonsReply) {}
rpc SetDriveSpeed (SetDriveSpeedRequest) returns (SetDriveSpeedReply) {}
rpc SetPipetteSpeed (SetPipetteSpeedRequest) returns (SetPipetteSpeedReply) {}
rpc SetPositionState (SetPositionStateRequest) returns (SetPositionStateReply) {}
rpc Stop (StopRequest) returns (StopReply) {}
rpc UnloadAdaptor (UnloadAdaptorRequest) returns (UnloadAdaptorReply) {}
rpc UnloadHead (UnloadHeadRequest) returns (UnloadHeadReply) {}
rpc UnloadTips (UnloadTipsRequest) returns (UnloadTipsReply) {}
rpc UpdateMetaData (UpdateMetaDataRequest) returns (UpdateMetaDataReply) {}
rpc Wait (WaitRequest) returns (WaitReply) {}
```

Proto Code Generation
---------------------

In this example we have chosen [Go](https://golang.org) as the implementation language. Go is part of the extended set 
of languages supported by Protocol Buffers via [Third Party Add ons](https://github.com/google/protobuf/wiki/Third-Party-Add-ons).

After [installing](https://github.com/golang/protobuf) the required code we can proceed to compile the .proto file into
a .pb.go go source code file. Using the following code. [lhdriver.proto](ExtendedLiquidhandlingDriver/lhdriver.proto) is 
a [version 3](https://developers.google.com/protocol-buffers/docs/reference/proto3-spec) proto file.

```
protoc lhdriver.proto --go_out=plugins=grpc:./
```

Will generate a lhdriver.pb.go file in the same directory where it is being called from. The generated file contains an 
implementation of a client interface that will communicate with a server instance. The server implementation, however,
must be implemented by us:

```
type ExtendedLiquidhandlingDriverServer interface {
    AddPlateTo(context.Context, *AddPlateToRequest) (*AddPlateToReply, error)
    Aspirate(context.Context, *AspirateRequest) (*AspirateReply, error)
    Close(context.Context, *CloseRequest) (*CloseReply, error)
    Dispense(context.Context, *DispenseRequest) (*DispenseReply, error)
    Finalize(context.Context, *FinalizeRequest) (*FinalizeReply, error)
    GetCapabilities(context.Context, *GetCapabilitiesRequest) (*GetCapabilitiesReply, error)
    GetCurrentPosition(context.Context, *GetCurrentPositionRequest) (*GetCurrentPositionReply, error)
    GetHeadState(context.Context, *GetHeadStateRequest) (*GetHeadStateReply, error)
    GetPositionState(context.Context, *GetPositionStateRequest) (*GetPositionStateReply, error)
    GetStatus(context.Context, *GetStatusRequest) (*GetStatusReply, error)
    Go(context.Context, *GoRequest) (*GoReply, error)
    Initialize(context.Context, *InitializeRequest) (*InitializeReply, error)
    LightsOff(context.Context, *LightsOffRequest) (*LightsOffReply, error)
    LightsOn(context.Context, *LightsOnRequest) (*LightsOnReply, error)
    LoadAdaptor(context.Context, *LoadAdaptorRequest) (*LoadAdaptorReply, error)
    LoadHead(context.Context, *LoadHeadRequest) (*LoadHeadReply, error)
    LoadTips(context.Context, *LoadTipsRequest) (*LoadTipsReply, error)
    Message(context.Context, *MessageRequest) (*MessageReply, error)
    Mix(context.Context, *MixRequest) (*MixReply, error)
    Move(context.Context, *MoveRequest) (*MoveReply, error)
    MoveRaw(context.Context, *MoveRawRequest) (*MoveRawReply, error)
    Open(context.Context, *OpenRequest) (*OpenReply, error)
    RemoveAllPlates(context.Context, *RemoveAllPlatesRequest) (*RemoveAllPlatesReply, error)
    RemovePlateAt(context.Context, *RemovePlateAtRequest) (*RemovePlateAtReply, error)
    ResetPistons(context.Context, *ResetPistonsRequest) (*ResetPistonsReply, error)
    SetDriveSpeed(context.Context, *SetDriveSpeedRequest) (*SetDriveSpeedReply, error)
    SetPipetteSpeed(context.Context, *SetPipetteSpeedRequest) (*SetPipetteSpeedReply, error)
    SetPositionState(context.Context, *SetPositionStateRequest) (*SetPositionStateReply, error)
    Stop(context.Context, *StopRequest) (*StopReply, error)
    UnloadAdaptor(context.Context, *UnloadAdaptorRequest) (*UnloadAdaptorReply, error)
    UnloadHead(context.Context, *UnloadHeadRequest) (*UnloadHeadReply, error)
    UnloadTips(context.Context, *UnloadTipsRequest) (*UnloadTipsReply, error)
    UpdateMetaData(context.Context, *UpdateMetaDataRequest) (*UpdateMetaDataReply, error)
    Wait(context.Context, *WaitRequest) (*WaitReply, error)
}
```

Server implementation
---------------------

[manual.go](./manual.go) implements that interface. A [server](./server/main.go) implementation that will let us define
certain variables can be executed using

```
go run server/main.go 
  -port=50051: Sepcify the port at which the server will be listening
  -view="cli": Specify the wished view to display the messages: cli | cui
```

Both cli and cui can be found in [./cli](./cli) and are attached as interfaces of [ManualExecuter](executor.go) that
exposes methods to display messages to the user and get a feedback from him.

Client example
--------------
[sampleClient](./sampleClient/main.go) instantiates a client driver and performs two different rpc methods on the same
server. To test it, we first need to launch our lhdriver server in the expected port

```
go run ./server/main.go -port=50051 -view=cli
```

Then execute the client

```
go run ./sampleClient/main.go
```

We should see the following Ouput in the driver side

```
$ go run server/main.go  -port=50051 -view=cli
2015/10/08 11:15:55 Listening at : 50051
[{Remove all plates from the deck. []}]
fb0270df-812d-4237-bef6-b3dd33add283  >  {Remove All Plates [{Remove all plates from the deck. []}]}
>> 'y' or Write error:
```

And the following in the client side

```
$ go run main.go 
2015/10/08 11:19:39 Done 1: Ret_1:<Arg_1:true Arg_3:"GO ACK" > 
2015/10/08 11:19:39 Done 2: Ret_1:<Arg_1:true Arg_3:"Stop ACK" >
```

The client will wait for an acknowledge on the server side that no error has occurred while executing RemoveAllPlates. As
the implementation of the Go and Stop methods implement an immediate response, The response is sent back from the server
as soon as the request is attended.
 
If we acknowledge that no error has happened on the server side manually by entering "y" and hitting the return key, the
following output will be appended on the client side

```
2015/10/08 11:19:44 Done 3: Ret_1:<Arg_1:true Arg_3:"REMOVEALLPLATES ACK" > 
```

AnthaRun as Client
==================
Using [antharun](http://github.com/antha-lang/antha) as a client to the manualLiquidHandler allows the debugging of protocols and study of the different actions that play a role in the execution of such protocol. The following command will run a compiled protocol instantiating a grpc driver on the antharun side that will connect on the specified port to the manualLiquidHandler execution server.

```
antharun --workflow wf.json --parameters params.json --driver localhost:50051
```
