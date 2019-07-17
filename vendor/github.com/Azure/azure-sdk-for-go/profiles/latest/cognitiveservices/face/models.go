// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package face

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/face"

type AccessoryType = original.AccessoryType

const (
	Glasses  AccessoryType = original.Glasses
	HeadWear AccessoryType = original.HeadWear
	Mask     AccessoryType = original.Mask
)

type AttributeType = original.AttributeType

const (
	AttributeTypeAccessories AttributeType = original.AttributeTypeAccessories
	AttributeTypeAge         AttributeType = original.AttributeTypeAge
	AttributeTypeBlur        AttributeType = original.AttributeTypeBlur
	AttributeTypeEmotion     AttributeType = original.AttributeTypeEmotion
	AttributeTypeExposure    AttributeType = original.AttributeTypeExposure
	AttributeTypeFacialHair  AttributeType = original.AttributeTypeFacialHair
	AttributeTypeGender      AttributeType = original.AttributeTypeGender
	AttributeTypeGlasses     AttributeType = original.AttributeTypeGlasses
	AttributeTypeHair        AttributeType = original.AttributeTypeHair
	AttributeTypeHeadPose    AttributeType = original.AttributeTypeHeadPose
	AttributeTypeMakeup      AttributeType = original.AttributeTypeMakeup
	AttributeTypeNoise       AttributeType = original.AttributeTypeNoise
	AttributeTypeOcclusion   AttributeType = original.AttributeTypeOcclusion
	AttributeTypeSmile       AttributeType = original.AttributeTypeSmile
)

type BlurLevel = original.BlurLevel

const (
	High   BlurLevel = original.High
	Low    BlurLevel = original.Low
	Medium BlurLevel = original.Medium
)

type ExposureLevel = original.ExposureLevel

const (
	GoodExposure  ExposureLevel = original.GoodExposure
	OverExposure  ExposureLevel = original.OverExposure
	UnderExposure ExposureLevel = original.UnderExposure
)

type FindSimilarMatchMode = original.FindSimilarMatchMode

const (
	MatchFace   FindSimilarMatchMode = original.MatchFace
	MatchPerson FindSimilarMatchMode = original.MatchPerson
)

type Gender = original.Gender

const (
	Female Gender = original.Female
	Male   Gender = original.Male
)

type GlassesType = original.GlassesType

const (
	NoGlasses       GlassesType = original.NoGlasses
	ReadingGlasses  GlassesType = original.ReadingGlasses
	Sunglasses      GlassesType = original.Sunglasses
	SwimmingGoggles GlassesType = original.SwimmingGoggles
)

type HairColorType = original.HairColorType

const (
	Black   HairColorType = original.Black
	Blond   HairColorType = original.Blond
	Brown   HairColorType = original.Brown
	Gray    HairColorType = original.Gray
	Other   HairColorType = original.Other
	Red     HairColorType = original.Red
	Unknown HairColorType = original.Unknown
	White   HairColorType = original.White
)

type NoiseLevel = original.NoiseLevel

const (
	NoiseLevelHigh   NoiseLevel = original.NoiseLevelHigh
	NoiseLevelLow    NoiseLevel = original.NoiseLevelLow
	NoiseLevelMedium NoiseLevel = original.NoiseLevelMedium
)

type OperationStatusType = original.OperationStatusType

const (
	Failed     OperationStatusType = original.Failed
	Notstarted OperationStatusType = original.Notstarted
	Running    OperationStatusType = original.Running
	Succeeded  OperationStatusType = original.Succeeded
)

type RecognitionModel = original.RecognitionModel

const (
	Recognition01 RecognitionModel = original.Recognition01
	Recognition02 RecognitionModel = original.Recognition02
)

type SnapshotApplyMode = original.SnapshotApplyMode

const (
	CreateNew SnapshotApplyMode = original.CreateNew
)

type SnapshotObjectType = original.SnapshotObjectType

const (
	SnapshotObjectTypeFaceList         SnapshotObjectType = original.SnapshotObjectTypeFaceList
	SnapshotObjectTypeLargeFaceList    SnapshotObjectType = original.SnapshotObjectTypeLargeFaceList
	SnapshotObjectTypeLargePersonGroup SnapshotObjectType = original.SnapshotObjectTypeLargePersonGroup
	SnapshotObjectTypePersonGroup      SnapshotObjectType = original.SnapshotObjectTypePersonGroup
)

type TrainingStatusType = original.TrainingStatusType

const (
	TrainingStatusTypeFailed     TrainingStatusType = original.TrainingStatusTypeFailed
	TrainingStatusTypeNonstarted TrainingStatusType = original.TrainingStatusTypeNonstarted
	TrainingStatusTypeRunning    TrainingStatusType = original.TrainingStatusTypeRunning
	TrainingStatusTypeSucceeded  TrainingStatusType = original.TrainingStatusTypeSucceeded
)

type APIError = original.APIError
type Accessory = original.Accessory
type ApplySnapshotRequest = original.ApplySnapshotRequest
type Attributes = original.Attributes
type BaseClient = original.BaseClient
type Blur = original.Blur
type Client = original.Client
type Coordinate = original.Coordinate
type DetectedFace = original.DetectedFace
type Emotion = original.Emotion
type Error = original.Error
type Exposure = original.Exposure
type FacialHair = original.FacialHair
type FindSimilarRequest = original.FindSimilarRequest
type GroupRequest = original.GroupRequest
type GroupResult = original.GroupResult
type Hair = original.Hair
type HairColor = original.HairColor
type HeadPose = original.HeadPose
type IdentifyCandidate = original.IdentifyCandidate
type IdentifyRequest = original.IdentifyRequest
type IdentifyResult = original.IdentifyResult
type ImageURL = original.ImageURL
type Landmarks = original.Landmarks
type LargeFaceList = original.LargeFaceList
type LargeFaceListClient = original.LargeFaceListClient
type LargePersonGroup = original.LargePersonGroup
type LargePersonGroupClient = original.LargePersonGroupClient
type LargePersonGroupPersonClient = original.LargePersonGroupPersonClient
type List = original.List
type ListClient = original.ListClient
type ListDetectedFace = original.ListDetectedFace
type ListIdentifyResult = original.ListIdentifyResult
type ListLargeFaceList = original.ListLargeFaceList
type ListLargePersonGroup = original.ListLargePersonGroup
type ListList = original.ListList
type ListPersistedFace = original.ListPersistedFace
type ListPerson = original.ListPerson
type ListPersonGroup = original.ListPersonGroup
type ListSimilarFace = original.ListSimilarFace
type ListSnapshot = original.ListSnapshot
type Makeup = original.Makeup
type MetaDataContract = original.MetaDataContract
type NameAndUserDataContract = original.NameAndUserDataContract
type Noise = original.Noise
type Occlusion = original.Occlusion
type OperationStatus = original.OperationStatus
type PersistedFace = original.PersistedFace
type Person = original.Person
type PersonGroup = original.PersonGroup
type PersonGroupClient = original.PersonGroupClient
type PersonGroupPersonClient = original.PersonGroupPersonClient
type Rectangle = original.Rectangle
type SimilarFace = original.SimilarFace
type Snapshot = original.Snapshot
type SnapshotClient = original.SnapshotClient
type TakeSnapshotRequest = original.TakeSnapshotRequest
type TrainingStatus = original.TrainingStatus
type UpdateFaceRequest = original.UpdateFaceRequest
type UpdateSnapshotRequest = original.UpdateSnapshotRequest
type VerifyFaceToFaceRequest = original.VerifyFaceToFaceRequest
type VerifyFaceToPersonRequest = original.VerifyFaceToPersonRequest
type VerifyResult = original.VerifyResult

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewClient(endpoint string) Client {
	return original.NewClient(endpoint)
}
func NewLargeFaceListClient(endpoint string) LargeFaceListClient {
	return original.NewLargeFaceListClient(endpoint)
}
func NewLargePersonGroupClient(endpoint string) LargePersonGroupClient {
	return original.NewLargePersonGroupClient(endpoint)
}
func NewLargePersonGroupPersonClient(endpoint string) LargePersonGroupPersonClient {
	return original.NewLargePersonGroupPersonClient(endpoint)
}
func NewListClient(endpoint string) ListClient {
	return original.NewListClient(endpoint)
}
func NewPersonGroupClient(endpoint string) PersonGroupClient {
	return original.NewPersonGroupClient(endpoint)
}
func NewPersonGroupPersonClient(endpoint string) PersonGroupPersonClient {
	return original.NewPersonGroupPersonClient(endpoint)
}
func NewSnapshotClient(endpoint string) SnapshotClient {
	return original.NewSnapshotClient(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleAccessoryTypeValues() []AccessoryType {
	return original.PossibleAccessoryTypeValues()
}
func PossibleAttributeTypeValues() []AttributeType {
	return original.PossibleAttributeTypeValues()
}
func PossibleBlurLevelValues() []BlurLevel {
	return original.PossibleBlurLevelValues()
}
func PossibleExposureLevelValues() []ExposureLevel {
	return original.PossibleExposureLevelValues()
}
func PossibleFindSimilarMatchModeValues() []FindSimilarMatchMode {
	return original.PossibleFindSimilarMatchModeValues()
}
func PossibleGenderValues() []Gender {
	return original.PossibleGenderValues()
}
func PossibleGlassesTypeValues() []GlassesType {
	return original.PossibleGlassesTypeValues()
}
func PossibleHairColorTypeValues() []HairColorType {
	return original.PossibleHairColorTypeValues()
}
func PossibleNoiseLevelValues() []NoiseLevel {
	return original.PossibleNoiseLevelValues()
}
func PossibleOperationStatusTypeValues() []OperationStatusType {
	return original.PossibleOperationStatusTypeValues()
}
func PossibleRecognitionModelValues() []RecognitionModel {
	return original.PossibleRecognitionModelValues()
}
func PossibleSnapshotApplyModeValues() []SnapshotApplyMode {
	return original.PossibleSnapshotApplyModeValues()
}
func PossibleSnapshotObjectTypeValues() []SnapshotObjectType {
	return original.PossibleSnapshotObjectTypeValues()
}
func PossibleTrainingStatusTypeValues() []TrainingStatusType {
	return original.PossibleTrainingStatusTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}