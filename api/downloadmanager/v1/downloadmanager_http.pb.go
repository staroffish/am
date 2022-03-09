// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	common "github.com/staroffish/am/api/common"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type DownloadmanagerHTTPServer interface {
	AddTask(context.Context, *AddTaskRequest) (*common.Empty, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*common.Empty, error)
	ListTask(context.Context, *ListTaskRequest) (*ListTaskReply, error)
	ScanTask(context.Context, *common.Empty) (*ScanTaskReply, error)
	ScanTaskAndDownload(context.Context, *common.Empty) (*ScanTaskAndDownloadReply, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*common.Empty, error)
}

func RegisterDownloadmanagerHTTPServer(s *http.Server, srv DownloadmanagerHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/download_manager/scan_and_download", _Downloadmanager_ScanTaskAndDownload0_HTTP_Handler(srv))
	r.POST("/v1/download_manager/scan", _Downloadmanager_ScanTask0_HTTP_Handler(srv))
	r.POST("/v1/download_manager/download_task/save", _Downloadmanager_AddTask0_HTTP_Handler(srv))
	r.POST("/v1/download_manager/download_task/update", _Downloadmanager_UpdateTask0_HTTP_Handler(srv))
	r.DELETE("/v1/download_manager/download_task/{id}", _Downloadmanager_DeleteTask0_HTTP_Handler(srv))
	r.GET("/v1/download_manager/download_task", _Downloadmanager_ListTask0_HTTP_Handler(srv))
}

func _Downloadmanager_ScanTaskAndDownload0_HTTP_Handler(srv DownloadmanagerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in common.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.downloadmanager.v1.Downloadmanager/ScanTaskAndDownload")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ScanTaskAndDownload(ctx, req.(*common.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ScanTaskAndDownloadReply)
		return ctx.Result(200, reply)
	}
}

func _Downloadmanager_ScanTask0_HTTP_Handler(srv DownloadmanagerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in common.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.downloadmanager.v1.Downloadmanager/ScanTask")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ScanTask(ctx, req.(*common.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ScanTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Downloadmanager_AddTask0_HTTP_Handler(srv DownloadmanagerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.downloadmanager.v1.Downloadmanager/AddTask")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddTask(ctx, req.(*AddTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Empty)
		return ctx.Result(200, reply)
	}
}

func _Downloadmanager_UpdateTask0_HTTP_Handler(srv DownloadmanagerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.downloadmanager.v1.Downloadmanager/UpdateTask")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTask(ctx, req.(*UpdateTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Empty)
		return ctx.Result(200, reply)
	}
}

func _Downloadmanager_DeleteTask0_HTTP_Handler(srv DownloadmanagerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.downloadmanager.v1.Downloadmanager/DeleteTask")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTask(ctx, req.(*DeleteTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*common.Empty)
		return ctx.Result(200, reply)
	}
}

func _Downloadmanager_ListTask0_HTTP_Handler(srv DownloadmanagerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.downloadmanager.v1.Downloadmanager/ListTask")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTask(ctx, req.(*ListTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTaskReply)
		return ctx.Result(200, reply)
	}
}

type DownloadmanagerHTTPClient interface {
	AddTask(ctx context.Context, req *AddTaskRequest, opts ...http.CallOption) (rsp *common.Empty, err error)
	DeleteTask(ctx context.Context, req *DeleteTaskRequest, opts ...http.CallOption) (rsp *common.Empty, err error)
	ListTask(ctx context.Context, req *ListTaskRequest, opts ...http.CallOption) (rsp *ListTaskReply, err error)
	ScanTask(ctx context.Context, req *common.Empty, opts ...http.CallOption) (rsp *ScanTaskReply, err error)
	ScanTaskAndDownload(ctx context.Context, req *common.Empty, opts ...http.CallOption) (rsp *ScanTaskAndDownloadReply, err error)
	UpdateTask(ctx context.Context, req *UpdateTaskRequest, opts ...http.CallOption) (rsp *common.Empty, err error)
}

type DownloadmanagerHTTPClientImpl struct {
	cc *http.Client
}

func NewDownloadmanagerHTTPClient(client *http.Client) DownloadmanagerHTTPClient {
	return &DownloadmanagerHTTPClientImpl{client}
}

func (c *DownloadmanagerHTTPClientImpl) AddTask(ctx context.Context, in *AddTaskRequest, opts ...http.CallOption) (*common.Empty, error) {
	var out common.Empty
	pattern := "/v1/download_manager/download_task/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.downloadmanager.v1.Downloadmanager/AddTask"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DownloadmanagerHTTPClientImpl) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...http.CallOption) (*common.Empty, error) {
	var out common.Empty
	pattern := "/v1/download_manager/download_task/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.downloadmanager.v1.Downloadmanager/DeleteTask"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DownloadmanagerHTTPClientImpl) ListTask(ctx context.Context, in *ListTaskRequest, opts ...http.CallOption) (*ListTaskReply, error) {
	var out ListTaskReply
	pattern := "/v1/download_manager/download_task"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.downloadmanager.v1.Downloadmanager/ListTask"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DownloadmanagerHTTPClientImpl) ScanTask(ctx context.Context, in *common.Empty, opts ...http.CallOption) (*ScanTaskReply, error) {
	var out ScanTaskReply
	pattern := "/v1/download_manager/scan"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.downloadmanager.v1.Downloadmanager/ScanTask"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DownloadmanagerHTTPClientImpl) ScanTaskAndDownload(ctx context.Context, in *common.Empty, opts ...http.CallOption) (*ScanTaskAndDownloadReply, error) {
	var out ScanTaskAndDownloadReply
	pattern := "/v1/download_manager/scan_and_download"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.downloadmanager.v1.Downloadmanager/ScanTaskAndDownload"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DownloadmanagerHTTPClientImpl) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...http.CallOption) (*common.Empty, error) {
	var out common.Empty
	pattern := "/v1/download_manager/download_task/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.downloadmanager.v1.Downloadmanager/UpdateTask"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
