// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package entry

import (
	context "context"
	types "github.com/spiffe/spire/proto/spire/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EntryClient is the client API for Entry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntryClient interface {
	// Lists entries.
	//
	// The caller must be local or present an admin X509-SVID.
	ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error)
	// Gets an entry. If the entry does not exist, NOT_FOUND is returned.
	//
	// The caller must be local or present an admin X509-SVID.
	GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*types.Entry, error)
	// Batch creates one or more entries.
	//
	// The caller must be local or present an admin X509-SVID.
	BatchCreateEntry(ctx context.Context, in *BatchCreateEntryRequest, opts ...grpc.CallOption) (*BatchCreateEntryResponse, error)
	// Batch updates one or more entries.
	//
	// The caller must be local or present an admin X509-SVID.
	BatchUpdateEntry(ctx context.Context, in *BatchUpdateEntryRequest, opts ...grpc.CallOption) (*BatchUpdateEntryResponse, error)
	// Batch deletes one or more entries.
	//
	// The caller must be local or present an admin X509-SVID.
	BatchDeleteEntry(ctx context.Context, in *BatchDeleteEntryRequest, opts ...grpc.CallOption) (*BatchDeleteEntryResponse, error)
	// Gets the entries the caller is authorized for.
	//
	// The caller must present an active agent X509-SVID. See the Agent
	// AttestAgent/RenewAgent RPCs.
	GetAuthorizedEntries(ctx context.Context, in *GetAuthorizedEntriesRequest, opts ...grpc.CallOption) (*GetAuthorizedEntriesResponse, error)
}

type entryClient struct {
	cc grpc.ClientConnInterface
}

func NewEntryClient(cc grpc.ClientConnInterface) EntryClient {
	return &entryClient{cc}
}

func (c *entryClient) ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error) {
	out := new(ListEntriesResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.entry.v1.Entry/ListEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryClient) GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*types.Entry, error) {
	out := new(types.Entry)
	err := c.cc.Invoke(ctx, "/spire.api.server.entry.v1.Entry/GetEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryClient) BatchCreateEntry(ctx context.Context, in *BatchCreateEntryRequest, opts ...grpc.CallOption) (*BatchCreateEntryResponse, error) {
	out := new(BatchCreateEntryResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.entry.v1.Entry/BatchCreateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryClient) BatchUpdateEntry(ctx context.Context, in *BatchUpdateEntryRequest, opts ...grpc.CallOption) (*BatchUpdateEntryResponse, error) {
	out := new(BatchUpdateEntryResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.entry.v1.Entry/BatchUpdateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryClient) BatchDeleteEntry(ctx context.Context, in *BatchDeleteEntryRequest, opts ...grpc.CallOption) (*BatchDeleteEntryResponse, error) {
	out := new(BatchDeleteEntryResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.entry.v1.Entry/BatchDeleteEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entryClient) GetAuthorizedEntries(ctx context.Context, in *GetAuthorizedEntriesRequest, opts ...grpc.CallOption) (*GetAuthorizedEntriesResponse, error) {
	out := new(GetAuthorizedEntriesResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.entry.v1.Entry/GetAuthorizedEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntryServer is the server API for Entry service.
// All implementations must embed UnimplementedEntryServer
// for forward compatibility
type EntryServer interface {
	// Lists entries.
	//
	// The caller must be local or present an admin X509-SVID.
	ListEntries(context.Context, *ListEntriesRequest) (*ListEntriesResponse, error)
	// Gets an entry. If the entry does not exist, NOT_FOUND is returned.
	//
	// The caller must be local or present an admin X509-SVID.
	GetEntry(context.Context, *GetEntryRequest) (*types.Entry, error)
	// Batch creates one or more entries.
	//
	// The caller must be local or present an admin X509-SVID.
	BatchCreateEntry(context.Context, *BatchCreateEntryRequest) (*BatchCreateEntryResponse, error)
	// Batch updates one or more entries.
	//
	// The caller must be local or present an admin X509-SVID.
	BatchUpdateEntry(context.Context, *BatchUpdateEntryRequest) (*BatchUpdateEntryResponse, error)
	// Batch deletes one or more entries.
	//
	// The caller must be local or present an admin X509-SVID.
	BatchDeleteEntry(context.Context, *BatchDeleteEntryRequest) (*BatchDeleteEntryResponse, error)
	// Gets the entries the caller is authorized for.
	//
	// The caller must present an active agent X509-SVID. See the Agent
	// AttestAgent/RenewAgent RPCs.
	GetAuthorizedEntries(context.Context, *GetAuthorizedEntriesRequest) (*GetAuthorizedEntriesResponse, error)
	mustEmbedUnimplementedEntryServer()
}

// UnimplementedEntryServer must be embedded to have forward compatible implementations.
type UnimplementedEntryServer struct {
}

func (UnimplementedEntryServer) ListEntries(context.Context, *ListEntriesRequest) (*ListEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntries not implemented")
}
func (UnimplementedEntryServer) GetEntry(context.Context, *GetEntryRequest) (*types.Entry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntry not implemented")
}
func (UnimplementedEntryServer) BatchCreateEntry(context.Context, *BatchCreateEntryRequest) (*BatchCreateEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreateEntry not implemented")
}
func (UnimplementedEntryServer) BatchUpdateEntry(context.Context, *BatchUpdateEntryRequest) (*BatchUpdateEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchUpdateEntry not implemented")
}
func (UnimplementedEntryServer) BatchDeleteEntry(context.Context, *BatchDeleteEntryRequest) (*BatchDeleteEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteEntry not implemented")
}
func (UnimplementedEntryServer) GetAuthorizedEntries(context.Context, *GetAuthorizedEntriesRequest) (*GetAuthorizedEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorizedEntries not implemented")
}
func (UnimplementedEntryServer) mustEmbedUnimplementedEntryServer() {}

// UnsafeEntryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntryServer will
// result in compilation errors.
type UnsafeEntryServer interface {
	mustEmbedUnimplementedEntryServer()
}

func RegisterEntryServer(s grpc.ServiceRegistrar, srv EntryServer) {
	s.RegisterService(&_Entry_serviceDesc, srv)
}

func _Entry_ListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServer).ListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.entry.v1.Entry/ListEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServer).ListEntries(ctx, req.(*ListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Entry_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.entry.v1.Entry/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServer).GetEntry(ctx, req.(*GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Entry_BatchCreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServer).BatchCreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.entry.v1.Entry/BatchCreateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServer).BatchCreateEntry(ctx, req.(*BatchCreateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Entry_BatchUpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchUpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServer).BatchUpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.entry.v1.Entry/BatchUpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServer).BatchUpdateEntry(ctx, req.(*BatchUpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Entry_BatchDeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchDeleteEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServer).BatchDeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.entry.v1.Entry/BatchDeleteEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServer).BatchDeleteEntry(ctx, req.(*BatchDeleteEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Entry_GetAuthorizedEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorizedEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServer).GetAuthorizedEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.entry.v1.Entry/GetAuthorizedEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServer).GetAuthorizedEntries(ctx, req.(*GetAuthorizedEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Entry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.api.server.entry.v1.Entry",
	HandlerType: (*EntryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEntries",
			Handler:    _Entry_ListEntries_Handler,
		},
		{
			MethodName: "GetEntry",
			Handler:    _Entry_GetEntry_Handler,
		},
		{
			MethodName: "BatchCreateEntry",
			Handler:    _Entry_BatchCreateEntry_Handler,
		},
		{
			MethodName: "BatchUpdateEntry",
			Handler:    _Entry_BatchUpdateEntry_Handler,
		},
		{
			MethodName: "BatchDeleteEntry",
			Handler:    _Entry_BatchDeleteEntry_Handler,
		},
		{
			MethodName: "GetAuthorizedEntries",
			Handler:    _Entry_GetAuthorizedEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spire/api/server/entry/v1/entry.proto",
}
