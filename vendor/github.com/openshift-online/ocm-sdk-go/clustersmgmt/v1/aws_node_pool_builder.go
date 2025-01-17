/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// AWSNodePoolBuilder contains the data and logic needed to build 'AWS_node_pool' objects.
//
// Representation of aws node pool specific parameters.
type AWSNodePoolBuilder struct {
	bitmap_               uint32
	id                    string
	href                  string
	availabilityZoneTypes map[string]string
	instanceProfile       string
	instanceType          string
	subnetOutposts        map[string]string
	tags                  map[string]string
}

// NewAWSNodePool creates a new builder of 'AWS_node_pool' objects.
func NewAWSNodePool() *AWSNodePoolBuilder {
	return &AWSNodePoolBuilder{}
}

// Link sets the flag that indicates if this is a link.
func (b *AWSNodePoolBuilder) Link(value bool) *AWSNodePoolBuilder {
	b.bitmap_ |= 1
	return b
}

// ID sets the identifier of the object.
func (b *AWSNodePoolBuilder) ID(value string) *AWSNodePoolBuilder {
	b.id = value
	b.bitmap_ |= 2
	return b
}

// HREF sets the link to the object.
func (b *AWSNodePoolBuilder) HREF(value string) *AWSNodePoolBuilder {
	b.href = value
	b.bitmap_ |= 4
	return b
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *AWSNodePoolBuilder) Empty() bool {
	return b == nil || b.bitmap_&^1 == 0
}

// AvailabilityZoneTypes sets the value of the 'availability_zone_types' attribute to the given value.
func (b *AWSNodePoolBuilder) AvailabilityZoneTypes(value map[string]string) *AWSNodePoolBuilder {
	b.availabilityZoneTypes = value
	if value != nil {
		b.bitmap_ |= 8
	} else {
		b.bitmap_ &^= 8
	}
	return b
}

// InstanceProfile sets the value of the 'instance_profile' attribute to the given value.
func (b *AWSNodePoolBuilder) InstanceProfile(value string) *AWSNodePoolBuilder {
	b.instanceProfile = value
	b.bitmap_ |= 16
	return b
}

// InstanceType sets the value of the 'instance_type' attribute to the given value.
func (b *AWSNodePoolBuilder) InstanceType(value string) *AWSNodePoolBuilder {
	b.instanceType = value
	b.bitmap_ |= 32
	return b
}

// SubnetOutposts sets the value of the 'subnet_outposts' attribute to the given value.
func (b *AWSNodePoolBuilder) SubnetOutposts(value map[string]string) *AWSNodePoolBuilder {
	b.subnetOutposts = value
	if value != nil {
		b.bitmap_ |= 64
	} else {
		b.bitmap_ &^= 64
	}
	return b
}

// Tags sets the value of the 'tags' attribute to the given value.
func (b *AWSNodePoolBuilder) Tags(value map[string]string) *AWSNodePoolBuilder {
	b.tags = value
	if value != nil {
		b.bitmap_ |= 128
	} else {
		b.bitmap_ &^= 128
	}
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AWSNodePoolBuilder) Copy(object *AWSNodePool) *AWSNodePoolBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.id = object.id
	b.href = object.href
	if len(object.availabilityZoneTypes) > 0 {
		b.availabilityZoneTypes = map[string]string{}
		for k, v := range object.availabilityZoneTypes {
			b.availabilityZoneTypes[k] = v
		}
	} else {
		b.availabilityZoneTypes = nil
	}
	b.instanceProfile = object.instanceProfile
	b.instanceType = object.instanceType
	if len(object.subnetOutposts) > 0 {
		b.subnetOutposts = map[string]string{}
		for k, v := range object.subnetOutposts {
			b.subnetOutposts[k] = v
		}
	} else {
		b.subnetOutposts = nil
	}
	if len(object.tags) > 0 {
		b.tags = map[string]string{}
		for k, v := range object.tags {
			b.tags[k] = v
		}
	} else {
		b.tags = nil
	}
	return b
}

// Build creates a 'AWS_node_pool' object using the configuration stored in the builder.
func (b *AWSNodePoolBuilder) Build() (object *AWSNodePool, err error) {
	object = new(AWSNodePool)
	object.id = b.id
	object.href = b.href
	object.bitmap_ = b.bitmap_
	if b.availabilityZoneTypes != nil {
		object.availabilityZoneTypes = make(map[string]string)
		for k, v := range b.availabilityZoneTypes {
			object.availabilityZoneTypes[k] = v
		}
	}
	object.instanceProfile = b.instanceProfile
	object.instanceType = b.instanceType
	if b.subnetOutposts != nil {
		object.subnetOutposts = make(map[string]string)
		for k, v := range b.subnetOutposts {
			object.subnetOutposts[k] = v
		}
	}
	if b.tags != nil {
		object.tags = make(map[string]string)
		for k, v := range b.tags {
			object.tags[k] = v
		}
	}
	return
}
