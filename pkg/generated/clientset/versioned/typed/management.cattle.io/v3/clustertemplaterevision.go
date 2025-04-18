/*
Copyright 2025 Rancher Labs, Inc.

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

package v3

import (
	"context"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ClusterTemplateRevisionsGetter has a method to return a ClusterTemplateRevisionInterface.
// A group's client should implement this interface.
type ClusterTemplateRevisionsGetter interface {
	ClusterTemplateRevisions(namespace string) ClusterTemplateRevisionInterface
}

// ClusterTemplateRevisionInterface has methods to work with ClusterTemplateRevision resources.
type ClusterTemplateRevisionInterface interface {
	Create(ctx context.Context, clusterTemplateRevision *v3.ClusterTemplateRevision, opts v1.CreateOptions) (*v3.ClusterTemplateRevision, error)
	Update(ctx context.Context, clusterTemplateRevision *v3.ClusterTemplateRevision, opts v1.UpdateOptions) (*v3.ClusterTemplateRevision, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, clusterTemplateRevision *v3.ClusterTemplateRevision, opts v1.UpdateOptions) (*v3.ClusterTemplateRevision, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.ClusterTemplateRevision, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.ClusterTemplateRevisionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterTemplateRevision, err error)
	ClusterTemplateRevisionExpansion
}

// clusterTemplateRevisions implements ClusterTemplateRevisionInterface
type clusterTemplateRevisions struct {
	*gentype.ClientWithList[*v3.ClusterTemplateRevision, *v3.ClusterTemplateRevisionList]
}

// newClusterTemplateRevisions returns a ClusterTemplateRevisions
func newClusterTemplateRevisions(c *ManagementV3Client, namespace string) *clusterTemplateRevisions {
	return &clusterTemplateRevisions{
		gentype.NewClientWithList[*v3.ClusterTemplateRevision, *v3.ClusterTemplateRevisionList](
			"clustertemplaterevisions",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v3.ClusterTemplateRevision { return &v3.ClusterTemplateRevision{} },
			func() *v3.ClusterTemplateRevisionList { return &v3.ClusterTemplateRevisionList{} }),
	}
}
