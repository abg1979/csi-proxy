// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	v1 "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1"
	impl "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy/impl"
)

func autoConvert_v1_ComputeDoubleRequest_To_impl_ComputeDoubleRequest(in *v1.ComputeDoubleRequest, out *impl.ComputeDoubleRequest) error {
	out.Input64 = in.Input64
	return nil
}

// Convert_v1_ComputeDoubleRequest_To_impl_ComputeDoubleRequest is an autogenerated conversion function.
func Convert_v1_ComputeDoubleRequest_To_impl_ComputeDoubleRequest(in *v1.ComputeDoubleRequest, out *impl.ComputeDoubleRequest) error {
	return autoConvert_v1_ComputeDoubleRequest_To_impl_ComputeDoubleRequest(in, out)
}

func autoConvert_impl_ComputeDoubleRequest_To_v1_ComputeDoubleRequest(in *impl.ComputeDoubleRequest, out *v1.ComputeDoubleRequest) error {
	out.Input64 = in.Input64
	return nil
}

// Convert_impl_ComputeDoubleRequest_To_v1_ComputeDoubleRequest is an autogenerated conversion function.
func Convert_impl_ComputeDoubleRequest_To_v1_ComputeDoubleRequest(in *impl.ComputeDoubleRequest, out *v1.ComputeDoubleRequest) error {
	return autoConvert_impl_ComputeDoubleRequest_To_v1_ComputeDoubleRequest(in, out)
}

func autoConvert_v1_ComputeDoubleResponse_To_impl_ComputeDoubleResponse(in *v1.ComputeDoubleResponse, out *impl.ComputeDoubleResponse) error {
	out.Response = in.Response
	return nil
}

// Convert_v1_ComputeDoubleResponse_To_impl_ComputeDoubleResponse is an autogenerated conversion function.
func Convert_v1_ComputeDoubleResponse_To_impl_ComputeDoubleResponse(in *v1.ComputeDoubleResponse, out *impl.ComputeDoubleResponse) error {
	return autoConvert_v1_ComputeDoubleResponse_To_impl_ComputeDoubleResponse(in, out)
}

func autoConvert_impl_ComputeDoubleResponse_To_v1_ComputeDoubleResponse(in *impl.ComputeDoubleResponse, out *v1.ComputeDoubleResponse) error {
	out.Response = in.Response
	return nil
}

// Convert_impl_ComputeDoubleResponse_To_v1_ComputeDoubleResponse is an autogenerated conversion function.
func Convert_impl_ComputeDoubleResponse_To_v1_ComputeDoubleResponse(in *impl.ComputeDoubleResponse, out *v1.ComputeDoubleResponse) error {
	return autoConvert_impl_ComputeDoubleResponse_To_v1_ComputeDoubleResponse(in, out)
}

func autoConvert_v1_TellMeAPoemRequest_To_impl_TellMeAPoemRequest(in *v1.TellMeAPoemRequest, out *impl.TellMeAPoemRequest) error {
	out.IWantATitle = in.IWantATitle
	return nil
}

// Convert_v1_TellMeAPoemRequest_To_impl_TellMeAPoemRequest is an autogenerated conversion function.
func Convert_v1_TellMeAPoemRequest_To_impl_TellMeAPoemRequest(in *v1.TellMeAPoemRequest, out *impl.TellMeAPoemRequest) error {
	return autoConvert_v1_TellMeAPoemRequest_To_impl_TellMeAPoemRequest(in, out)
}

func autoConvert_impl_TellMeAPoemRequest_To_v1_TellMeAPoemRequest(in *impl.TellMeAPoemRequest, out *v1.TellMeAPoemRequest) error {
	out.IWantATitle = in.IWantATitle
	return nil
}

// Convert_impl_TellMeAPoemRequest_To_v1_TellMeAPoemRequest is an autogenerated conversion function.
func Convert_impl_TellMeAPoemRequest_To_v1_TellMeAPoemRequest(in *impl.TellMeAPoemRequest, out *v1.TellMeAPoemRequest) error {
	return autoConvert_impl_TellMeAPoemRequest_To_v1_TellMeAPoemRequest(in, out)
}

func autoConvert_v1_TellMeAPoemResponse_To_impl_TellMeAPoemResponse(in *v1.TellMeAPoemResponse, out *impl.TellMeAPoemResponse) error {
	out.Title = in.Title
	out.Lines = *(*[]string)(unsafe.Pointer(&in.Lines))
	return nil
}

// Convert_v1_TellMeAPoemResponse_To_impl_TellMeAPoemResponse is an autogenerated conversion function.
func Convert_v1_TellMeAPoemResponse_To_impl_TellMeAPoemResponse(in *v1.TellMeAPoemResponse, out *impl.TellMeAPoemResponse) error {
	return autoConvert_v1_TellMeAPoemResponse_To_impl_TellMeAPoemResponse(in, out)
}

func autoConvert_impl_TellMeAPoemResponse_To_v1_TellMeAPoemResponse(in *impl.TellMeAPoemResponse, out *v1.TellMeAPoemResponse) error {
	out.Title = in.Title
	out.Lines = *(*[]string)(unsafe.Pointer(&in.Lines))
	return nil
}

// Convert_impl_TellMeAPoemResponse_To_v1_TellMeAPoemResponse is an autogenerated conversion function.
func Convert_impl_TellMeAPoemResponse_To_v1_TellMeAPoemResponse(in *impl.TellMeAPoemResponse, out *v1.TellMeAPoemResponse) error {
	return autoConvert_impl_TellMeAPoemResponse_To_v1_TellMeAPoemResponse(in, out)
}
