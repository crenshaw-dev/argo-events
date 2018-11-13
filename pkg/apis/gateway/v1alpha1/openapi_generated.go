// +build !ignore_autogenerated

/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by main. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.Gateway":                    schema_pkg_apis_gateway_v1alpha1_Gateway(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayList":                schema_pkg_apis_gateway_v1alpha1_GatewayList(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayNotificationWatcher": schema_pkg_apis_gateway_v1alpha1_GatewayNotificationWatcher(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewaySpec":                schema_pkg_apis_gateway_v1alpha1_GatewaySpec(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayStatus":              schema_pkg_apis_gateway_v1alpha1_GatewayStatus(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.NodeStatus":                 schema_pkg_apis_gateway_v1alpha1_NodeStatus(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.NotificationWatchers":       schema_pkg_apis_gateway_v1alpha1_NotificationWatchers(ref),
		"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.SensorNotificationWatcher":  schema_pkg_apis_gateway_v1alpha1_SensorNotificationWatcher(ref),
	}
}

func schema_pkg_apis_gateway_v1alpha1_Gateway(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Gateway is the definition of a gateway resource",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayStatus"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewaySpec"),
						},
					},
				},
				Required: []string{"metadata", "status", "spec"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewaySpec", "github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_gateway_v1alpha1_GatewayList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GatewayList is the list of Gateway resources",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.Gateway"),
									},
								},
							},
						},
					},
				},
				Required: []string{"metadata", "items"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.Gateway", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_gateway_v1alpha1_GatewayNotificationWatcher(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GatewayNotificationWatcher is the gateway interested in listening to notifications from this gateway",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the gateway name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"port": {
						SchemaProps: spec.SchemaProps{
							Description: "Port is http server port on which gateway is running",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Description: "Endpoint is REST API endpoint to post event to. Events are sent using HTTP POST method to this endpoint.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name", "port", "endpoint"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_gateway_v1alpha1_GatewaySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GatewaySpec represents gateway specifications",
				Properties: map[string]spec.Schema{
					"deploySpec": {
						SchemaProps: spec.SchemaProps{
							Description: "DeploySpec is description of gateway",
							Ref:         ref("k8s.io/api/core/v1.PodSpec"),
						},
					},
					"configMap": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigMap is name of the configmap for gateway-processor",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type is the type of gateway",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "Version is used for marking event version",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"serviceSpec": {
						SchemaProps: spec.SchemaProps{
							Description: "ServiceSpec is the specifications of the service to expose the gateway",
							Ref:         ref("k8s.io/api/core/v1.ServiceSpec"),
						},
					},
					"watchers": {
						SchemaProps: spec.SchemaProps{
							Description: "Watchers are components which are interested listening to notifications from this gateway These only need to be specified when gateway dispatch mechanism is through HTTP POST notifications. In future, support for NATS, KAFKA will be added as a means to dispatch notifications in which case specifying watchers would be unnecessary.",
							Ref:         ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.NotificationWatchers"),
						},
					},
					"rpcPort": {
						SchemaProps: spec.SchemaProps{
							Description: "RPCPort if provided is used to communicate between gRPC gateway client and gRPC gateway server",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"httpServerPort": {
						SchemaProps: spec.SchemaProps{
							Description: "HTTPServerPort if provided is used to communicate between gateway client and server over http",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dispatchMechanism": {
						SchemaProps: spec.SchemaProps{
							Description: "DispatchMechanism is the underlying mechanism used to send events from gateway to watchers(components interested in listening to event from this gateway)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imageVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "ImageVersion is the version for gateway components images to run",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"deploySpec", "type", "version", "dispatchMechanism"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.NotificationWatchers", "k8s.io/api/core/v1.PodSpec", "k8s.io/api/core/v1.ServiceSpec"},
	}
}

func schema_pkg_apis_gateway_v1alpha1_GatewayStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GatewayStatus contains information about the status of a gateway.",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is the high-level summary of the gateway",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"startedAt": {
						SchemaProps: spec.SchemaProps{
							Description: "StartedAt is the time at which this gateway was initiated",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message is a human readable string indicating details about a gateway in its phase",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nodes": {
						SchemaProps: spec.SchemaProps{
							Description: "Nodes is a mapping between a node ID and the node's status it records the states for the configurations of gateway.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.NodeStatus"),
									},
								},
							},
						},
					},
				},
				Required: []string{"phase"},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.NodeStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_gateway_v1alpha1_NodeStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodeStatus describes the status for an individual node in the gateway configurations. A single node can represent one configuration.",
				Properties: map[string]spec.Schema{
					"id": {
						SchemaProps: spec.SchemaProps{
							Description: "ID is a unique identifier of a node within a sensor It is a hash of the node name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"timeID": {
						SchemaProps: spec.SchemaProps{
							Description: "TimeID is used to resolve events arriving out of order for same node",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is a unique name in the node tree used to generate the node ID",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"displayName": {
						SchemaProps: spec.SchemaProps{
							Description: "DisplayName is the human readable representation of the node",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase of the node",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"startedAt": {
						SchemaProps: spec.SchemaProps{
							Description: "StartedAt is the time at which this node started",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.MicroTime"),
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message store data or something to save for configuration",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"updateTime": {
						SchemaProps: spec.SchemaProps{
							Description: "UpdateTime is the time when node(gateway configuration) was updated",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.MicroTime"),
						},
					},
				},
				Required: []string{"id", "timeID", "name", "displayName", "phase"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.MicroTime"},
	}
}

func schema_pkg_apis_gateway_v1alpha1_NotificationWatchers(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NotificationWatchers are components which are interested listening to notifications from this gateway",
				Properties: map[string]spec.Schema{
					"gateways": {
						SchemaProps: spec.SchemaProps{
							Description: "Gateways is the list of gateways interested in listening to notifications from this gateway",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayNotificationWatcher"),
									},
								},
							},
						},
					},
					"sensors": {
						SchemaProps: spec.SchemaProps{
							Description: "Sensors is the list of sensors interested in listening to notifications from this gateway",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.SensorNotificationWatcher"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.GatewayNotificationWatcher", "github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1.SensorNotificationWatcher"},
	}
}

func schema_pkg_apis_gateway_v1alpha1_SensorNotificationWatcher(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SensorNotificationWatcher is the sensor interested in listening to notifications from this gateway",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is name of the sensor",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
		Dependencies: []string{},
	}
}