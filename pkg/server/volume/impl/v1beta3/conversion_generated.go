// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1beta3

import (
	unsafe "unsafe"

	v1beta3 "github.com/kubernetes-csi/csi-proxy/client/api/volume/v1beta3"
	impl "github.com/kubernetes-csi/csi-proxy/pkg/server/volume/impl"
)

func autoConvert_v1beta3_FormatVolumeRequest_To_impl_FormatVolumeRequest(in *v1beta3.FormatVolumeRequest, out *impl.FormatVolumeRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta3_FormatVolumeRequest_To_impl_FormatVolumeRequest is an autogenerated conversion function.
func Convert_v1beta3_FormatVolumeRequest_To_impl_FormatVolumeRequest(in *v1beta3.FormatVolumeRequest, out *impl.FormatVolumeRequest) error {
	return autoConvert_v1beta3_FormatVolumeRequest_To_impl_FormatVolumeRequest(in, out)
}

func autoConvert_impl_FormatVolumeRequest_To_v1beta3_FormatVolumeRequest(in *impl.FormatVolumeRequest, out *v1beta3.FormatVolumeRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_impl_FormatVolumeRequest_To_v1beta3_FormatVolumeRequest is an autogenerated conversion function.
func Convert_impl_FormatVolumeRequest_To_v1beta3_FormatVolumeRequest(in *impl.FormatVolumeRequest, out *v1beta3.FormatVolumeRequest) error {
	return autoConvert_impl_FormatVolumeRequest_To_v1beta3_FormatVolumeRequest(in, out)
}

func autoConvert_v1beta3_FormatVolumeResponse_To_impl_FormatVolumeResponse(in *v1beta3.FormatVolumeResponse, out *impl.FormatVolumeResponse) error {
	return nil
}

// Convert_v1beta3_FormatVolumeResponse_To_impl_FormatVolumeResponse is an autogenerated conversion function.
func Convert_v1beta3_FormatVolumeResponse_To_impl_FormatVolumeResponse(in *v1beta3.FormatVolumeResponse, out *impl.FormatVolumeResponse) error {
	return autoConvert_v1beta3_FormatVolumeResponse_To_impl_FormatVolumeResponse(in, out)
}

func autoConvert_impl_FormatVolumeResponse_To_v1beta3_FormatVolumeResponse(in *impl.FormatVolumeResponse, out *v1beta3.FormatVolumeResponse) error {
	return nil
}

// Convert_impl_FormatVolumeResponse_To_v1beta3_FormatVolumeResponse is an autogenerated conversion function.
func Convert_impl_FormatVolumeResponse_To_v1beta3_FormatVolumeResponse(in *impl.FormatVolumeResponse, out *v1beta3.FormatVolumeResponse) error {
	return autoConvert_impl_FormatVolumeResponse_To_v1beta3_FormatVolumeResponse(in, out)
}

func autoConvert_v1beta3_GetDiskNumberFromVolumeIDRequest_To_impl_GetDiskNumberFromVolumeIDRequest(in *v1beta3.GetDiskNumberFromVolumeIDRequest, out *impl.GetDiskNumberFromVolumeIDRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta3_GetDiskNumberFromVolumeIDRequest_To_impl_GetDiskNumberFromVolumeIDRequest is an autogenerated conversion function.
func Convert_v1beta3_GetDiskNumberFromVolumeIDRequest_To_impl_GetDiskNumberFromVolumeIDRequest(in *v1beta3.GetDiskNumberFromVolumeIDRequest, out *impl.GetDiskNumberFromVolumeIDRequest) error {
	return autoConvert_v1beta3_GetDiskNumberFromVolumeIDRequest_To_impl_GetDiskNumberFromVolumeIDRequest(in, out)
}

func autoConvert_impl_GetDiskNumberFromVolumeIDRequest_To_v1beta3_GetDiskNumberFromVolumeIDRequest(in *impl.GetDiskNumberFromVolumeIDRequest, out *v1beta3.GetDiskNumberFromVolumeIDRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_impl_GetDiskNumberFromVolumeIDRequest_To_v1beta3_GetDiskNumberFromVolumeIDRequest is an autogenerated conversion function.
func Convert_impl_GetDiskNumberFromVolumeIDRequest_To_v1beta3_GetDiskNumberFromVolumeIDRequest(in *impl.GetDiskNumberFromVolumeIDRequest, out *v1beta3.GetDiskNumberFromVolumeIDRequest) error {
	return autoConvert_impl_GetDiskNumberFromVolumeIDRequest_To_v1beta3_GetDiskNumberFromVolumeIDRequest(in, out)
}

func autoConvert_v1beta3_GetDiskNumberFromVolumeIDResponse_To_impl_GetDiskNumberFromVolumeIDResponse(in *v1beta3.GetDiskNumberFromVolumeIDResponse, out *impl.GetDiskNumberFromVolumeIDResponse) error {
	out.DiskNumber = in.DiskNumber
	return nil
}

// Convert_v1beta3_GetDiskNumberFromVolumeIDResponse_To_impl_GetDiskNumberFromVolumeIDResponse is an autogenerated conversion function.
func Convert_v1beta3_GetDiskNumberFromVolumeIDResponse_To_impl_GetDiskNumberFromVolumeIDResponse(in *v1beta3.GetDiskNumberFromVolumeIDResponse, out *impl.GetDiskNumberFromVolumeIDResponse) error {
	return autoConvert_v1beta3_GetDiskNumberFromVolumeIDResponse_To_impl_GetDiskNumberFromVolumeIDResponse(in, out)
}

func autoConvert_impl_GetDiskNumberFromVolumeIDResponse_To_v1beta3_GetDiskNumberFromVolumeIDResponse(in *impl.GetDiskNumberFromVolumeIDResponse, out *v1beta3.GetDiskNumberFromVolumeIDResponse) error {
	out.DiskNumber = in.DiskNumber
	return nil
}

// Convert_impl_GetDiskNumberFromVolumeIDResponse_To_v1beta3_GetDiskNumberFromVolumeIDResponse is an autogenerated conversion function.
func Convert_impl_GetDiskNumberFromVolumeIDResponse_To_v1beta3_GetDiskNumberFromVolumeIDResponse(in *impl.GetDiskNumberFromVolumeIDResponse, out *v1beta3.GetDiskNumberFromVolumeIDResponse) error {
	return autoConvert_impl_GetDiskNumberFromVolumeIDResponse_To_v1beta3_GetDiskNumberFromVolumeIDResponse(in, out)
}

func autoConvert_v1beta3_GetVolumeIDFromTargetPathRequest_To_impl_GetVolumeIDFromTargetPathRequest(in *v1beta3.GetVolumeIDFromTargetPathRequest, out *impl.GetVolumeIDFromTargetPathRequest) error {
	out.TargetPath = in.TargetPath
	return nil
}

// Convert_v1beta3_GetVolumeIDFromTargetPathRequest_To_impl_GetVolumeIDFromTargetPathRequest is an autogenerated conversion function.
func Convert_v1beta3_GetVolumeIDFromTargetPathRequest_To_impl_GetVolumeIDFromTargetPathRequest(in *v1beta3.GetVolumeIDFromTargetPathRequest, out *impl.GetVolumeIDFromTargetPathRequest) error {
	return autoConvert_v1beta3_GetVolumeIDFromTargetPathRequest_To_impl_GetVolumeIDFromTargetPathRequest(in, out)
}

func autoConvert_impl_GetVolumeIDFromTargetPathRequest_To_v1beta3_GetVolumeIDFromTargetPathRequest(in *impl.GetVolumeIDFromTargetPathRequest, out *v1beta3.GetVolumeIDFromTargetPathRequest) error {
	out.TargetPath = in.TargetPath
	return nil
}

// Convert_impl_GetVolumeIDFromTargetPathRequest_To_v1beta3_GetVolumeIDFromTargetPathRequest is an autogenerated conversion function.
func Convert_impl_GetVolumeIDFromTargetPathRequest_To_v1beta3_GetVolumeIDFromTargetPathRequest(in *impl.GetVolumeIDFromTargetPathRequest, out *v1beta3.GetVolumeIDFromTargetPathRequest) error {
	return autoConvert_impl_GetVolumeIDFromTargetPathRequest_To_v1beta3_GetVolumeIDFromTargetPathRequest(in, out)
}

func autoConvert_v1beta3_GetVolumeIDFromTargetPathResponse_To_impl_GetVolumeIDFromTargetPathResponse(in *v1beta3.GetVolumeIDFromTargetPathResponse, out *impl.GetVolumeIDFromTargetPathResponse) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta3_GetVolumeIDFromTargetPathResponse_To_impl_GetVolumeIDFromTargetPathResponse is an autogenerated conversion function.
func Convert_v1beta3_GetVolumeIDFromTargetPathResponse_To_impl_GetVolumeIDFromTargetPathResponse(in *v1beta3.GetVolumeIDFromTargetPathResponse, out *impl.GetVolumeIDFromTargetPathResponse) error {
	return autoConvert_v1beta3_GetVolumeIDFromTargetPathResponse_To_impl_GetVolumeIDFromTargetPathResponse(in, out)
}

func autoConvert_impl_GetVolumeIDFromTargetPathResponse_To_v1beta3_GetVolumeIDFromTargetPathResponse(in *impl.GetVolumeIDFromTargetPathResponse, out *v1beta3.GetVolumeIDFromTargetPathResponse) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_impl_GetVolumeIDFromTargetPathResponse_To_v1beta3_GetVolumeIDFromTargetPathResponse is an autogenerated conversion function.
func Convert_impl_GetVolumeIDFromTargetPathResponse_To_v1beta3_GetVolumeIDFromTargetPathResponse(in *impl.GetVolumeIDFromTargetPathResponse, out *v1beta3.GetVolumeIDFromTargetPathResponse) error {
	return autoConvert_impl_GetVolumeIDFromTargetPathResponse_To_v1beta3_GetVolumeIDFromTargetPathResponse(in, out)
}

func autoConvert_v1beta3_GetVolumeStatsRequest_To_impl_GetVolumeStatsRequest(in *v1beta3.GetVolumeStatsRequest, out *impl.GetVolumeStatsRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta3_GetVolumeStatsRequest_To_impl_GetVolumeStatsRequest is an autogenerated conversion function.
func Convert_v1beta3_GetVolumeStatsRequest_To_impl_GetVolumeStatsRequest(in *v1beta3.GetVolumeStatsRequest, out *impl.GetVolumeStatsRequest) error {
	return autoConvert_v1beta3_GetVolumeStatsRequest_To_impl_GetVolumeStatsRequest(in, out)
}

func autoConvert_impl_GetVolumeStatsRequest_To_v1beta3_GetVolumeStatsRequest(in *impl.GetVolumeStatsRequest, out *v1beta3.GetVolumeStatsRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_impl_GetVolumeStatsRequest_To_v1beta3_GetVolumeStatsRequest is an autogenerated conversion function.
func Convert_impl_GetVolumeStatsRequest_To_v1beta3_GetVolumeStatsRequest(in *impl.GetVolumeStatsRequest, out *v1beta3.GetVolumeStatsRequest) error {
	return autoConvert_impl_GetVolumeStatsRequest_To_v1beta3_GetVolumeStatsRequest(in, out)
}

func autoConvert_v1beta3_GetVolumeStatsResponse_To_impl_GetVolumeStatsResponse(in *v1beta3.GetVolumeStatsResponse, out *impl.GetVolumeStatsResponse) error {
	out.TotalBytes = in.TotalBytes
	out.UsedBytes = in.UsedBytes
	return nil
}

// Convert_v1beta3_GetVolumeStatsResponse_To_impl_GetVolumeStatsResponse is an autogenerated conversion function.
func Convert_v1beta3_GetVolumeStatsResponse_To_impl_GetVolumeStatsResponse(in *v1beta3.GetVolumeStatsResponse, out *impl.GetVolumeStatsResponse) error {
	return autoConvert_v1beta3_GetVolumeStatsResponse_To_impl_GetVolumeStatsResponse(in, out)
}

func autoConvert_impl_GetVolumeStatsResponse_To_v1beta3_GetVolumeStatsResponse(in *impl.GetVolumeStatsResponse, out *v1beta3.GetVolumeStatsResponse) error {
	out.TotalBytes = in.TotalBytes
	out.UsedBytes = in.UsedBytes
	return nil
}

// Convert_impl_GetVolumeStatsResponse_To_v1beta3_GetVolumeStatsResponse is an autogenerated conversion function.
func Convert_impl_GetVolumeStatsResponse_To_v1beta3_GetVolumeStatsResponse(in *impl.GetVolumeStatsResponse, out *v1beta3.GetVolumeStatsResponse) error {
	return autoConvert_impl_GetVolumeStatsResponse_To_v1beta3_GetVolumeStatsResponse(in, out)
}

func autoConvert_v1beta3_IsVolumeFormattedRequest_To_impl_IsVolumeFormattedRequest(in *v1beta3.IsVolumeFormattedRequest, out *impl.IsVolumeFormattedRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta3_IsVolumeFormattedRequest_To_impl_IsVolumeFormattedRequest is an autogenerated conversion function.
func Convert_v1beta3_IsVolumeFormattedRequest_To_impl_IsVolumeFormattedRequest(in *v1beta3.IsVolumeFormattedRequest, out *impl.IsVolumeFormattedRequest) error {
	return autoConvert_v1beta3_IsVolumeFormattedRequest_To_impl_IsVolumeFormattedRequest(in, out)
}

func autoConvert_impl_IsVolumeFormattedRequest_To_v1beta3_IsVolumeFormattedRequest(in *impl.IsVolumeFormattedRequest, out *v1beta3.IsVolumeFormattedRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_impl_IsVolumeFormattedRequest_To_v1beta3_IsVolumeFormattedRequest is an autogenerated conversion function.
func Convert_impl_IsVolumeFormattedRequest_To_v1beta3_IsVolumeFormattedRequest(in *impl.IsVolumeFormattedRequest, out *v1beta3.IsVolumeFormattedRequest) error {
	return autoConvert_impl_IsVolumeFormattedRequest_To_v1beta3_IsVolumeFormattedRequest(in, out)
}

func autoConvert_v1beta3_IsVolumeFormattedResponse_To_impl_IsVolumeFormattedResponse(in *v1beta3.IsVolumeFormattedResponse, out *impl.IsVolumeFormattedResponse) error {
	out.Formatted = in.Formatted
	return nil
}

// Convert_v1beta3_IsVolumeFormattedResponse_To_impl_IsVolumeFormattedResponse is an autogenerated conversion function.
func Convert_v1beta3_IsVolumeFormattedResponse_To_impl_IsVolumeFormattedResponse(in *v1beta3.IsVolumeFormattedResponse, out *impl.IsVolumeFormattedResponse) error {
	return autoConvert_v1beta3_IsVolumeFormattedResponse_To_impl_IsVolumeFormattedResponse(in, out)
}

func autoConvert_impl_IsVolumeFormattedResponse_To_v1beta3_IsVolumeFormattedResponse(in *impl.IsVolumeFormattedResponse, out *v1beta3.IsVolumeFormattedResponse) error {
	out.Formatted = in.Formatted
	return nil
}

// Convert_impl_IsVolumeFormattedResponse_To_v1beta3_IsVolumeFormattedResponse is an autogenerated conversion function.
func Convert_impl_IsVolumeFormattedResponse_To_v1beta3_IsVolumeFormattedResponse(in *impl.IsVolumeFormattedResponse, out *v1beta3.IsVolumeFormattedResponse) error {
	return autoConvert_impl_IsVolumeFormattedResponse_To_v1beta3_IsVolumeFormattedResponse(in, out)
}

func autoConvert_v1beta3_ListVolumesOnDiskRequest_To_impl_ListVolumesOnDiskRequest(in *v1beta3.ListVolumesOnDiskRequest, out *impl.ListVolumesOnDiskRequest) error {
	out.DiskNumber = in.DiskNumber
	out.PartitionNumber = in.PartitionNumber
	return nil
}

// Convert_v1beta3_ListVolumesOnDiskRequest_To_impl_ListVolumesOnDiskRequest is an autogenerated conversion function.
func Convert_v1beta3_ListVolumesOnDiskRequest_To_impl_ListVolumesOnDiskRequest(in *v1beta3.ListVolumesOnDiskRequest, out *impl.ListVolumesOnDiskRequest) error {
	return autoConvert_v1beta3_ListVolumesOnDiskRequest_To_impl_ListVolumesOnDiskRequest(in, out)
}

func autoConvert_impl_ListVolumesOnDiskRequest_To_v1beta3_ListVolumesOnDiskRequest(in *impl.ListVolumesOnDiskRequest, out *v1beta3.ListVolumesOnDiskRequest) error {
	out.DiskNumber = in.DiskNumber
	out.PartitionNumber = in.PartitionNumber
	return nil
}

// Convert_impl_ListVolumesOnDiskRequest_To_v1beta3_ListVolumesOnDiskRequest is an autogenerated conversion function.
func Convert_impl_ListVolumesOnDiskRequest_To_v1beta3_ListVolumesOnDiskRequest(in *impl.ListVolumesOnDiskRequest, out *v1beta3.ListVolumesOnDiskRequest) error {
	return autoConvert_impl_ListVolumesOnDiskRequest_To_v1beta3_ListVolumesOnDiskRequest(in, out)
}

func autoConvert_v1beta3_ListVolumesOnDiskResponse_To_impl_ListVolumesOnDiskResponse(in *v1beta3.ListVolumesOnDiskResponse, out *impl.ListVolumesOnDiskResponse) error {
	out.VolumeIds = *(*[]string)(unsafe.Pointer(&in.VolumeIds))
	return nil
}

// Convert_v1beta3_ListVolumesOnDiskResponse_To_impl_ListVolumesOnDiskResponse is an autogenerated conversion function.
func Convert_v1beta3_ListVolumesOnDiskResponse_To_impl_ListVolumesOnDiskResponse(in *v1beta3.ListVolumesOnDiskResponse, out *impl.ListVolumesOnDiskResponse) error {
	return autoConvert_v1beta3_ListVolumesOnDiskResponse_To_impl_ListVolumesOnDiskResponse(in, out)
}

func autoConvert_impl_ListVolumesOnDiskResponse_To_v1beta3_ListVolumesOnDiskResponse(in *impl.ListVolumesOnDiskResponse, out *v1beta3.ListVolumesOnDiskResponse) error {
	out.VolumeIds = *(*[]string)(unsafe.Pointer(&in.VolumeIds))
	return nil
}

// Convert_impl_ListVolumesOnDiskResponse_To_v1beta3_ListVolumesOnDiskResponse is an autogenerated conversion function.
func Convert_impl_ListVolumesOnDiskResponse_To_v1beta3_ListVolumesOnDiskResponse(in *impl.ListVolumesOnDiskResponse, out *v1beta3.ListVolumesOnDiskResponse) error {
	return autoConvert_impl_ListVolumesOnDiskResponse_To_v1beta3_ListVolumesOnDiskResponse(in, out)
}

func autoConvert_v1beta3_MountVolumeRequest_To_impl_MountVolumeRequest(in *v1beta3.MountVolumeRequest, out *impl.MountVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.TargetPath = in.TargetPath
	return nil
}

// Convert_v1beta3_MountVolumeRequest_To_impl_MountVolumeRequest is an autogenerated conversion function.
func Convert_v1beta3_MountVolumeRequest_To_impl_MountVolumeRequest(in *v1beta3.MountVolumeRequest, out *impl.MountVolumeRequest) error {
	return autoConvert_v1beta3_MountVolumeRequest_To_impl_MountVolumeRequest(in, out)
}

func autoConvert_impl_MountVolumeRequest_To_v1beta3_MountVolumeRequest(in *impl.MountVolumeRequest, out *v1beta3.MountVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.TargetPath = in.TargetPath
	return nil
}

// Convert_impl_MountVolumeRequest_To_v1beta3_MountVolumeRequest is an autogenerated conversion function.
func Convert_impl_MountVolumeRequest_To_v1beta3_MountVolumeRequest(in *impl.MountVolumeRequest, out *v1beta3.MountVolumeRequest) error {
	return autoConvert_impl_MountVolumeRequest_To_v1beta3_MountVolumeRequest(in, out)
}

func autoConvert_v1beta3_MountVolumeResponse_To_impl_MountVolumeResponse(in *v1beta3.MountVolumeResponse, out *impl.MountVolumeResponse) error {
	return nil
}

// Convert_v1beta3_MountVolumeResponse_To_impl_MountVolumeResponse is an autogenerated conversion function.
func Convert_v1beta3_MountVolumeResponse_To_impl_MountVolumeResponse(in *v1beta3.MountVolumeResponse, out *impl.MountVolumeResponse) error {
	return autoConvert_v1beta3_MountVolumeResponse_To_impl_MountVolumeResponse(in, out)
}

func autoConvert_impl_MountVolumeResponse_To_v1beta3_MountVolumeResponse(in *impl.MountVolumeResponse, out *v1beta3.MountVolumeResponse) error {
	return nil
}

// Convert_impl_MountVolumeResponse_To_v1beta3_MountVolumeResponse is an autogenerated conversion function.
func Convert_impl_MountVolumeResponse_To_v1beta3_MountVolumeResponse(in *impl.MountVolumeResponse, out *v1beta3.MountVolumeResponse) error {
	return autoConvert_impl_MountVolumeResponse_To_v1beta3_MountVolumeResponse(in, out)
}

func autoConvert_v1beta3_ResizeVolumeRequest_To_impl_ResizeVolumeRequest(in *v1beta3.ResizeVolumeRequest, out *impl.ResizeVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.SizeBytes = in.SizeBytes
	return nil
}

// Convert_v1beta3_ResizeVolumeRequest_To_impl_ResizeVolumeRequest is an autogenerated conversion function.
func Convert_v1beta3_ResizeVolumeRequest_To_impl_ResizeVolumeRequest(in *v1beta3.ResizeVolumeRequest, out *impl.ResizeVolumeRequest) error {
	return autoConvert_v1beta3_ResizeVolumeRequest_To_impl_ResizeVolumeRequest(in, out)
}

func autoConvert_impl_ResizeVolumeRequest_To_v1beta3_ResizeVolumeRequest(in *impl.ResizeVolumeRequest, out *v1beta3.ResizeVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.SizeBytes = in.SizeBytes
	return nil
}

// Convert_impl_ResizeVolumeRequest_To_v1beta3_ResizeVolumeRequest is an autogenerated conversion function.
func Convert_impl_ResizeVolumeRequest_To_v1beta3_ResizeVolumeRequest(in *impl.ResizeVolumeRequest, out *v1beta3.ResizeVolumeRequest) error {
	return autoConvert_impl_ResizeVolumeRequest_To_v1beta3_ResizeVolumeRequest(in, out)
}

func autoConvert_v1beta3_ResizeVolumeResponse_To_impl_ResizeVolumeResponse(in *v1beta3.ResizeVolumeResponse, out *impl.ResizeVolumeResponse) error {
	return nil
}

// Convert_v1beta3_ResizeVolumeResponse_To_impl_ResizeVolumeResponse is an autogenerated conversion function.
func Convert_v1beta3_ResizeVolumeResponse_To_impl_ResizeVolumeResponse(in *v1beta3.ResizeVolumeResponse, out *impl.ResizeVolumeResponse) error {
	return autoConvert_v1beta3_ResizeVolumeResponse_To_impl_ResizeVolumeResponse(in, out)
}

func autoConvert_impl_ResizeVolumeResponse_To_v1beta3_ResizeVolumeResponse(in *impl.ResizeVolumeResponse, out *v1beta3.ResizeVolumeResponse) error {
	return nil
}

// Convert_impl_ResizeVolumeResponse_To_v1beta3_ResizeVolumeResponse is an autogenerated conversion function.
func Convert_impl_ResizeVolumeResponse_To_v1beta3_ResizeVolumeResponse(in *impl.ResizeVolumeResponse, out *v1beta3.ResizeVolumeResponse) error {
	return autoConvert_impl_ResizeVolumeResponse_To_v1beta3_ResizeVolumeResponse(in, out)
}

func autoConvert_v1beta3_UnmountVolumeRequest_To_impl_UnmountVolumeRequest(in *v1beta3.UnmountVolumeRequest, out *impl.UnmountVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.TargetPath = in.TargetPath
	return nil
}

// Convert_v1beta3_UnmountVolumeRequest_To_impl_UnmountVolumeRequest is an autogenerated conversion function.
func Convert_v1beta3_UnmountVolumeRequest_To_impl_UnmountVolumeRequest(in *v1beta3.UnmountVolumeRequest, out *impl.UnmountVolumeRequest) error {
	return autoConvert_v1beta3_UnmountVolumeRequest_To_impl_UnmountVolumeRequest(in, out)
}

func autoConvert_impl_UnmountVolumeRequest_To_v1beta3_UnmountVolumeRequest(in *impl.UnmountVolumeRequest, out *v1beta3.UnmountVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.TargetPath = in.TargetPath
	return nil
}

// Convert_impl_UnmountVolumeRequest_To_v1beta3_UnmountVolumeRequest is an autogenerated conversion function.
func Convert_impl_UnmountVolumeRequest_To_v1beta3_UnmountVolumeRequest(in *impl.UnmountVolumeRequest, out *v1beta3.UnmountVolumeRequest) error {
	return autoConvert_impl_UnmountVolumeRequest_To_v1beta3_UnmountVolumeRequest(in, out)
}

func autoConvert_v1beta3_UnmountVolumeResponse_To_impl_UnmountVolumeResponse(in *v1beta3.UnmountVolumeResponse, out *impl.UnmountVolumeResponse) error {
	return nil
}

// Convert_v1beta3_UnmountVolumeResponse_To_impl_UnmountVolumeResponse is an autogenerated conversion function.
func Convert_v1beta3_UnmountVolumeResponse_To_impl_UnmountVolumeResponse(in *v1beta3.UnmountVolumeResponse, out *impl.UnmountVolumeResponse) error {
	return autoConvert_v1beta3_UnmountVolumeResponse_To_impl_UnmountVolumeResponse(in, out)
}

func autoConvert_impl_UnmountVolumeResponse_To_v1beta3_UnmountVolumeResponse(in *impl.UnmountVolumeResponse, out *v1beta3.UnmountVolumeResponse) error {
	return nil
}

// Convert_impl_UnmountVolumeResponse_To_v1beta3_UnmountVolumeResponse is an autogenerated conversion function.
func Convert_impl_UnmountVolumeResponse_To_v1beta3_UnmountVolumeResponse(in *impl.UnmountVolumeResponse, out *v1beta3.UnmountVolumeResponse) error {
	return autoConvert_impl_UnmountVolumeResponse_To_v1beta3_UnmountVolumeResponse(in, out)
}

func autoConvert_v1beta3_WriteVolumeCacheRequest_To_impl_WriteVolumeCacheRequest(in *v1beta3.WriteVolumeCacheRequest, out *impl.WriteVolumeCacheRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta3_WriteVolumeCacheRequest_To_impl_WriteVolumeCacheRequest is an autogenerated conversion function.
func Convert_v1beta3_WriteVolumeCacheRequest_To_impl_WriteVolumeCacheRequest(in *v1beta3.WriteVolumeCacheRequest, out *impl.WriteVolumeCacheRequest) error {
	return autoConvert_v1beta3_WriteVolumeCacheRequest_To_impl_WriteVolumeCacheRequest(in, out)
}

func autoConvert_impl_WriteVolumeCacheRequest_To_v1beta3_WriteVolumeCacheRequest(in *impl.WriteVolumeCacheRequest, out *v1beta3.WriteVolumeCacheRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_impl_WriteVolumeCacheRequest_To_v1beta3_WriteVolumeCacheRequest is an autogenerated conversion function.
func Convert_impl_WriteVolumeCacheRequest_To_v1beta3_WriteVolumeCacheRequest(in *impl.WriteVolumeCacheRequest, out *v1beta3.WriteVolumeCacheRequest) error {
	return autoConvert_impl_WriteVolumeCacheRequest_To_v1beta3_WriteVolumeCacheRequest(in, out)
}

func autoConvert_v1beta3_WriteVolumeCacheResponse_To_impl_WriteVolumeCacheResponse(in *v1beta3.WriteVolumeCacheResponse, out *impl.WriteVolumeCacheResponse) error {
	return nil
}

// Convert_v1beta3_WriteVolumeCacheResponse_To_impl_WriteVolumeCacheResponse is an autogenerated conversion function.
func Convert_v1beta3_WriteVolumeCacheResponse_To_impl_WriteVolumeCacheResponse(in *v1beta3.WriteVolumeCacheResponse, out *impl.WriteVolumeCacheResponse) error {
	return autoConvert_v1beta3_WriteVolumeCacheResponse_To_impl_WriteVolumeCacheResponse(in, out)
}

func autoConvert_impl_WriteVolumeCacheResponse_To_v1beta3_WriteVolumeCacheResponse(in *impl.WriteVolumeCacheResponse, out *v1beta3.WriteVolumeCacheResponse) error {
	return nil
}

// Convert_impl_WriteVolumeCacheResponse_To_v1beta3_WriteVolumeCacheResponse is an autogenerated conversion function.
func Convert_impl_WriteVolumeCacheResponse_To_v1beta3_WriteVolumeCacheResponse(in *impl.WriteVolumeCacheResponse, out *v1beta3.WriteVolumeCacheResponse) error {
	return autoConvert_impl_WriteVolumeCacheResponse_To_v1beta3_WriteVolumeCacheResponse(in, out)
}
