// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package ipfix_export

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	memclnt "github.com/brmcdoug/go-vpp-sr/vppbinapi/memclnt"
)

// RPCService defines RPC service ipfix_export.
type RPCService interface {
	IpfixAllExporterGet(ctx context.Context, in *IpfixAllExporterGet) (RPCService_IpfixAllExporterGetClient, error)
	IpfixClassifyStreamDump(ctx context.Context, in *IpfixClassifyStreamDump) (RPCService_IpfixClassifyStreamDumpClient, error)
	IpfixClassifyTableAddDel(ctx context.Context, in *IpfixClassifyTableAddDel) (*IpfixClassifyTableAddDelReply, error)
	IpfixClassifyTableDump(ctx context.Context, in *IpfixClassifyTableDump) (RPCService_IpfixClassifyTableDumpClient, error)
	IpfixExporterCreateDelete(ctx context.Context, in *IpfixExporterCreateDelete) (*IpfixExporterCreateDeleteReply, error)
	IpfixExporterDump(ctx context.Context, in *IpfixExporterDump) (RPCService_IpfixExporterDumpClient, error)
	IpfixFlush(ctx context.Context, in *IpfixFlush) (*IpfixFlushReply, error)
	SetIpfixClassifyStream(ctx context.Context, in *SetIpfixClassifyStream) (*SetIpfixClassifyStreamReply, error)
	SetIpfixExporter(ctx context.Context, in *SetIpfixExporter) (*SetIpfixExporterReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) IpfixAllExporterGet(ctx context.Context, in *IpfixAllExporterGet) (RPCService_IpfixAllExporterGetClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpfixAllExporterGetClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpfixAllExporterGetClient interface {
	Recv() (*IpfixAllExporterDetails, error)
	api.Stream
}

type serviceClient_IpfixAllExporterGetClient struct {
	api.Stream
}

func (c *serviceClient_IpfixAllExporterGetClient) Recv() (*IpfixAllExporterDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpfixAllExporterDetails:
		return m, nil
	case *IpfixAllExporterGetReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpfixClassifyStreamDump(ctx context.Context, in *IpfixClassifyStreamDump) (RPCService_IpfixClassifyStreamDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpfixClassifyStreamDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpfixClassifyStreamDumpClient interface {
	Recv() (*IpfixClassifyStreamDetails, error)
	api.Stream
}

type serviceClient_IpfixClassifyStreamDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpfixClassifyStreamDumpClient) Recv() (*IpfixClassifyStreamDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpfixClassifyStreamDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpfixClassifyTableAddDel(ctx context.Context, in *IpfixClassifyTableAddDel) (*IpfixClassifyTableAddDelReply, error) {
	out := new(IpfixClassifyTableAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpfixClassifyTableDump(ctx context.Context, in *IpfixClassifyTableDump) (RPCService_IpfixClassifyTableDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpfixClassifyTableDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpfixClassifyTableDumpClient interface {
	Recv() (*IpfixClassifyTableDetails, error)
	api.Stream
}

type serviceClient_IpfixClassifyTableDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpfixClassifyTableDumpClient) Recv() (*IpfixClassifyTableDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpfixClassifyTableDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpfixExporterCreateDelete(ctx context.Context, in *IpfixExporterCreateDelete) (*IpfixExporterCreateDeleteReply, error) {
	out := new(IpfixExporterCreateDeleteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpfixExporterDump(ctx context.Context, in *IpfixExporterDump) (RPCService_IpfixExporterDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpfixExporterDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpfixExporterDumpClient interface {
	Recv() (*IpfixExporterDetails, error)
	api.Stream
}

type serviceClient_IpfixExporterDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpfixExporterDumpClient) Recv() (*IpfixExporterDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpfixExporterDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpfixFlush(ctx context.Context, in *IpfixFlush) (*IpfixFlushReply, error) {
	out := new(IpfixFlushReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SetIpfixClassifyStream(ctx context.Context, in *SetIpfixClassifyStream) (*SetIpfixClassifyStreamReply, error) {
	out := new(SetIpfixClassifyStreamReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SetIpfixExporter(ctx context.Context, in *SetIpfixExporter) (*SetIpfixExporterReply, error) {
	out := new(SetIpfixExporterReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
