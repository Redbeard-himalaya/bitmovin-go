package models

// import "github.com/bitmovin/bitmovin-go/bitmovintypes"
import "../bitmovintypes"

type CreatePrewarmedEncoderPoolRequest struct {
	Name             *string                      `json:"name,omitempty"`
	Description      *string                      `json:"description,omitempty"`
	InfrastructureId *string                      `json:"InfrastructureId,omitempty"`
	EncoderVersion   bitmovintypes.EncoderVersion `json:"encoderVersion"`
	CloudRegion      bitmovintypes.CloudRegion    `json:"cloudRegion"`
	DiskSize         bitmovintypes.DiskSize       `json:"diskSize"`
	TargetPoolSize   int64                        `json:"targetPoolSize"`
}

type PrewarmedEncoderPoolDetail struct {
	ID               string                      `json:"id,omitempty"`
	Name             *string                      `json:"name,omitempty"`
	Description      *string                      `json:"description,omitempty"`
	EncoderVersion   bitmovintypes.EncoderVersion `json:"encoderVersion,omitempty"`
	CloudRegion      bitmovintypes.CloudRegion    `json:"cloudRegion,omitempty"`
	DiskSize         bitmovintypes.DiskSize       `json:"diskSize,omitempty"`
	TargetPoolSize   int64                        `json:"targetPoolSize,omitempty"`
	Status           string                       `json:"status,omitempty"`
}

type PrewarmedEncoderPoolResponseData struct {
	Result PrewarmedEncoderPoolDetail `json:"result,omitempty"`
}

type PrewarmedEncoderPoolResponse struct {
	RequestID *string                          `json:"requestId,omitempty"`
	Status    bitmovintypes.ResponseStatus     `json:"status,omitempty"`
	Data      PrewarmedEncoderPoolResponseData `json:"data,omitempty"`
}

type PrewarmedEncoderPoolListResult struct {
	TotalCount *int64                       `json:"totalCount,omitempty"`
	Previous   *string                      `json:"previous,omitempty"`
	Next       *string                      `json:"next,omitempty"`
	Items      []PrewarmedEncoderPoolDetail `json:"items,omitempty"`
}

type PrewarmedEncoderPoolListResponseData struct {
	Result PrewarmedEncoderPoolListResult `json:"result,omitempty"`
}

type PrewarmedEncoderPoolListResponse struct {
	RequestID *string                              `json:"requestId,omitempty"`
	Status    bitmovintypes.ResponseStatus         `json:"status,omitempty"`
	Data      PrewarmedEncoderPoolListResponseData `json:"data,omitempty"`
}
