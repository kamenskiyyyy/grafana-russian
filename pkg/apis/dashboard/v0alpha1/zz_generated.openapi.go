//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v0alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationActions":       schema_pkg_apis_dashboard_v0alpha1_AnnotationActions(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationPermission":    schema_pkg_apis_dashboard_v0alpha1_AnnotationPermission(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.Dashboard":               schema_pkg_apis_dashboard_v0alpha1_Dashboard(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardAccess":         schema_pkg_apis_dashboard_v0alpha1_DashboardAccess(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardList":           schema_pkg_apis_dashboard_v0alpha1_DashboardList(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardSummary":        schema_pkg_apis_dashboard_v0alpha1_DashboardSummary(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardSummaryList":    schema_pkg_apis_dashboard_v0alpha1_DashboardSummaryList(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardSummarySpec":    schema_pkg_apis_dashboard_v0alpha1_DashboardSummarySpec(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardVersionInfo":    schema_pkg_apis_dashboard_v0alpha1_DashboardVersionInfo(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardVersionList":    schema_pkg_apis_dashboard_v0alpha1_DashboardVersionList(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardWithAccessInfo": schema_pkg_apis_dashboard_v0alpha1_DashboardWithAccessInfo(ref),
		"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.VersionsQueryOptions":    schema_pkg_apis_dashboard_v0alpha1_VersionsQueryOptions(ref),
	}
}

func schema_pkg_apis_dashboard_v0alpha1_AnnotationActions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"canAdd": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canEdit": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canDelete": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
				},
				Required: []string{"canAdd", "canEdit", "canDelete"},
			},
		},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_AnnotationPermission(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"dashboard": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationActions"),
						},
					},
					"organization": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationActions"),
						},
					},
				},
				Required: []string{"dashboard", "organization"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationActions"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_Dashboard(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard object's metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "The dashboard body (unstructured for now)",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
				},
				Required: []string{"spec"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardAccess(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Information about how the requesting user can use a given dashboard",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"slug": {
						SchemaProps: spec.SchemaProps{
							Description: "Metadata fields",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"url": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"canSave": {
						SchemaProps: spec.SchemaProps{
							Description: "The permissions part",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"canEdit": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canAdmin": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canStar": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"canDelete": {
						SchemaProps: spec.SchemaProps{
							Default: false,
							Type:    []string{"boolean"},
							Format:  "",
						},
					},
					"annotationsPermissions": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationPermission"),
						},
					},
				},
				Required: []string{"canSave", "canEdit", "canAdmin", "canStar", "canDelete", "annotationsPermissions"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.AnnotationPermission"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.Dashboard"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.Dashboard", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardSummary(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "The dashboard body",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardSummarySpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardSummarySpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardSummaryList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardSummary"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardSummary", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardSummarySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"title": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"tags": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"title"},
			},
		},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardVersionInfo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "The internal ID for this version (will be replaced with resourceVersion)",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"parentVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "If the dashboard came from a previous version, it is set here",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"created": {
						SchemaProps: spec.SchemaProps{
							Description: "The creation timestamp for this version",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"createdBy": {
						SchemaProps: spec.SchemaProps{
							Description: "The user who created this version",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message passed while saving the version",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"version", "created"},
			},
		},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardVersionList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardVersionInfo"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardVersionInfo", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_DashboardWithAccessInfo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "This is like the legacy DTO where access and metadata are all returned in a single call",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard object's metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "The dashboard body (unstructured for now)",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
					"access": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardAccess"),
						},
					},
				},
				Required: []string{"spec", "access"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured", "github.com/grafana/grafana/pkg/apis/dashboard/v0alpha1.DashboardAccess", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dashboard_v0alpha1_VersionsQueryOptions(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"path": {
						SchemaProps: spec.SchemaProps{
							Description: "Path is the URL path",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
				},
			},
		},
	}
}
