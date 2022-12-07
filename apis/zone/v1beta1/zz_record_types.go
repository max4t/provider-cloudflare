/*
Copyright 2022 Max T.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DataObservation struct {
}

type DataParameters struct {

	// +kubebuilder:validation:Optional
	Algorithm *float64 `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// +kubebuilder:validation:Optional
	Altitude *float64 `json:"altitude,omitempty" tf:"altitude,omitempty"`

	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// +kubebuilder:validation:Optional
	Digest *string `json:"digest,omitempty" tf:"digest,omitempty"`

	// +kubebuilder:validation:Optional
	DigestType *float64 `json:"digestType,omitempty" tf:"digest_type,omitempty"`

	// +kubebuilder:validation:Optional
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// +kubebuilder:validation:Optional
	Flags *string `json:"flags,omitempty" tf:"flags,omitempty"`

	// +kubebuilder:validation:Optional
	KeyTag *float64 `json:"keyTag,omitempty" tf:"key_tag,omitempty"`

	// +kubebuilder:validation:Optional
	LatDegrees *float64 `json:"latDegrees,omitempty" tf:"lat_degrees,omitempty"`

	// +kubebuilder:validation:Optional
	LatDirection *string `json:"latDirection,omitempty" tf:"lat_direction,omitempty"`

	// +kubebuilder:validation:Optional
	LatMinutes *float64 `json:"latMinutes,omitempty" tf:"lat_minutes,omitempty"`

	// +kubebuilder:validation:Optional
	LatSeconds *float64 `json:"latSeconds,omitempty" tf:"lat_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	LongDegrees *float64 `json:"longDegrees,omitempty" tf:"long_degrees,omitempty"`

	// +kubebuilder:validation:Optional
	LongDirection *string `json:"longDirection,omitempty" tf:"long_direction,omitempty"`

	// +kubebuilder:validation:Optional
	LongMinutes *float64 `json:"longMinutes,omitempty" tf:"long_minutes,omitempty"`

	// +kubebuilder:validation:Optional
	LongSeconds *float64 `json:"longSeconds,omitempty" tf:"long_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	MatchingType *float64 `json:"matchingType,omitempty" tf:"matching_type,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PrecisionHorz *float64 `json:"precisionHorz,omitempty" tf:"precision_horz,omitempty"`

	// +kubebuilder:validation:Optional
	PrecisionVert *float64 `json:"precisionVert,omitempty" tf:"precision_vert,omitempty"`

	// +kubebuilder:validation:Optional
	Preference *float64 `json:"preference,omitempty" tf:"preference,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	Proto *string `json:"proto,omitempty" tf:"proto,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *float64 `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// +kubebuilder:validation:Optional
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// +kubebuilder:validation:Optional
	Replacement *string `json:"replacement,omitempty" tf:"replacement,omitempty"`

	// +kubebuilder:validation:Optional
	Selector *float64 `json:"selector,omitempty" tf:"selector,omitempty"`

	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// +kubebuilder:validation:Optional
	Type *float64 `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Usage *float64 `json:"usage,omitempty" tf:"usage,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type RecordObservation struct {
	CreatedOn *string `json:"createdOn,omitempty" tf:"created_on,omitempty"`

	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	ModifiedOn *string `json:"modifiedOn,omitempty" tf:"modified_on,omitempty"`

	Proxiable *bool `json:"proxiable,omitempty" tf:"proxiable,omitempty"`
}

type RecordParameters struct {

	// +kubebuilder:validation:Optional
	AllowOverwrite *bool `json:"allowOverwrite,omitempty" tf:"allow_overwrite,omitempty"`

	// +kubebuilder:validation:Optional
	Data []DataParameters `json:"data,omitempty" tf:"data,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	Proxied *bool `json:"proxied,omitempty" tf:"proxied,omitempty"`

	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// +crossplane:generate:reference:type=Zone
	// +crossplane:generate:reference:refFieldName=ZoneRef
	// +crossplane:generate:reference:selectorFieldName=ZoneSelector
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneRef *v1.Reference `json:"zoneRef,omitempty" tf:"-"`

	// Selector for a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneSelector *v1.Selector `json:"zoneSelector,omitempty" tf:"-"`
}

// RecordSpec defines the desired state of Record
type RecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordParameters `json:"forProvider"`
}

// RecordStatus defines the observed state of Record.
type RecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Record is the Schema for the Records API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Record struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecordSpec   `json:"spec"`
	Status            RecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordList contains a list of Records
type RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Record `json:"items"`
}

// Repository type metadata.
var (
	Record_Kind             = "Record"
	Record_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Record_Kind}.String()
	Record_KindAPIVersion   = Record_Kind + "." + CRDGroupVersion.String()
	Record_GroupVersionKind = CRDGroupVersion.WithKind(Record_Kind)
)

func init() {
	SchemeBuilder.Register(&Record{}, &RecordList{})
}
