//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Elasticsearch) DeepCopyInto(out *Elasticsearch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Elasticsearch.
func (in *Elasticsearch) DeepCopy() *Elasticsearch {
	if in == nil {
		return nil
	}
	out := new(Elasticsearch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Elasticsearch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchAlert) DeepCopyInto(out *ElasticsearchAlert) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AdditionalRuleLabels != nil {
		in, out := &in.AdditionalRuleLabels, &out.AdditionalRuleLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Groups = in.Groups
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchAlert.
func (in *ElasticsearchAlert) DeepCopy() *ElasticsearchAlert {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchAlertGroups) DeepCopyInto(out *ElasticsearchAlertGroups) {
	*out = *in
	out.Database = in.Database
	out.Provisioner = in.Provisioner
	out.OpsManager = in.OpsManager
	out.Stash = in.Stash
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchAlertGroups.
func (in *ElasticsearchAlertGroups) DeepCopy() *ElasticsearchAlertGroups {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchAlertGroups)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDatabaseAlert) DeepCopyInto(out *ElasticsearchDatabaseAlert) {
	*out = *in
	out.Rules = in.Rules
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDatabaseAlert.
func (in *ElasticsearchDatabaseAlert) DeepCopy() *ElasticsearchDatabaseAlert {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDatabaseAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchDatabaseAlertRules) DeepCopyInto(out *ElasticsearchDatabaseAlertRules) {
	*out = *in
	out.ElasticsearchHeapUsageTooHigh = in.ElasticsearchHeapUsageTooHigh
	out.ElasticsearchHeapUsageWarning = in.ElasticsearchHeapUsageWarning
	out.ElasticsearchDiskOutOfSpace = in.ElasticsearchDiskOutOfSpace
	out.ElasticsearchDiskSpaceLow = in.ElasticsearchDiskSpaceLow
	out.ElasticsearchClusterRed = in.ElasticsearchClusterRed
	out.ElasticsearchClusterYellow = in.ElasticsearchClusterYellow
	out.ElasticsearchHealthyNodes = in.ElasticsearchHealthyNodes
	out.ElasticsearchHealthyDataNodes = in.ElasticsearchHealthyDataNodes
	out.ElasticsearchRelocatingShards = in.ElasticsearchRelocatingShards
	out.ElasticsearchInitializingShards = in.ElasticsearchInitializingShards
	out.ElasticsearchUnassignedShards = in.ElasticsearchUnassignedShards
	out.ElasticsearchPendingTasks = in.ElasticsearchPendingTasks
	out.ElasticsearchNoNewDocuments10M = in.ElasticsearchNoNewDocuments10M
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchDatabaseAlertRules.
func (in *ElasticsearchDatabaseAlertRules) DeepCopy() *ElasticsearchDatabaseAlertRules {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchDatabaseAlertRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchList) DeepCopyInto(out *ElasticsearchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Elasticsearch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchList.
func (in *ElasticsearchList) DeepCopy() *ElasticsearchList {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchSpec) DeepCopyInto(out *ElasticsearchSpec) {
	*out = *in
	out.Metadata = in.Metadata
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchSpec.
func (in *ElasticsearchSpec) DeepCopy() *ElasticsearchSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchSpecSpec) DeepCopyInto(out *ElasticsearchSpecSpec) {
	*out = *in
	in.Alert.DeepCopyInto(&out.Alert)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchSpecSpec.
func (in *ElasticsearchSpecSpec) DeepCopy() *ElasticsearchSpecSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchSpecSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixedAlert) DeepCopyInto(out *FixedAlert) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixedAlert.
func (in *FixedAlert) DeepCopy() *FixedAlert {
	if in == nil {
		return nil
	}
	out := new(FixedAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FloatValAlertConfig) DeepCopyInto(out *FloatValAlertConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FloatValAlertConfig.
func (in *FloatValAlertConfig) DeepCopy() *FloatValAlertConfig {
	if in == nil {
		return nil
	}
	out := new(FloatValAlertConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntValAlert) DeepCopyInto(out *IntValAlert) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntValAlert.
func (in *IntValAlert) DeepCopy() *IntValAlert {
	if in == nil {
		return nil
	}
	out := new(IntValAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariaDBAlert) DeepCopyInto(out *MariaDBAlert) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AdditionalRuleLabels != nil {
		in, out := &in.AdditionalRuleLabels, &out.AdditionalRuleLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Groups = in.Groups
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariaDBAlert.
func (in *MariaDBAlert) DeepCopy() *MariaDBAlert {
	if in == nil {
		return nil
	}
	out := new(MariaDBAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariaDBAlertGroups) DeepCopyInto(out *MariaDBAlertGroups) {
	*out = *in
	out.Database = in.Database
	out.Cluster = in.Cluster
	out.Provisioner = in.Provisioner
	out.OpsManager = in.OpsManager
	out.Stash = in.Stash
	out.SchemaManager = in.SchemaManager
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariaDBAlertGroups.
func (in *MariaDBAlertGroups) DeepCopy() *MariaDBAlertGroups {
	if in == nil {
		return nil
	}
	out := new(MariaDBAlertGroups)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariaDBClusterAlert) DeepCopyInto(out *MariaDBClusterAlert) {
	*out = *in
	out.Rules = in.Rules
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariaDBClusterAlert.
func (in *MariaDBClusterAlert) DeepCopy() *MariaDBClusterAlert {
	if in == nil {
		return nil
	}
	out := new(MariaDBClusterAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariaDBClusterAlertRules) DeepCopyInto(out *MariaDBClusterAlertRules) {
	*out = *in
	out.GaleraReplicationLatencyTooLong = in.GaleraReplicationLatencyTooLong
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariaDBClusterAlertRules.
func (in *MariaDBClusterAlertRules) DeepCopy() *MariaDBClusterAlertRules {
	if in == nil {
		return nil
	}
	out := new(MariaDBClusterAlertRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariaDBDatabaseAlert) DeepCopyInto(out *MariaDBDatabaseAlert) {
	*out = *in
	out.Rules = in.Rules
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariaDBDatabaseAlert.
func (in *MariaDBDatabaseAlert) DeepCopy() *MariaDBDatabaseAlert {
	if in == nil {
		return nil
	}
	out := new(MariaDBDatabaseAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariaDBDatabaseAlertRules) DeepCopyInto(out *MariaDBDatabaseAlertRules) {
	*out = *in
	out.MySQLInstanceDown = in.MySQLInstanceDown
	out.MySQLServiceDown = in.MySQLServiceDown
	out.MySQLTooManyConnections = in.MySQLTooManyConnections
	out.MySQLHighThreadsRunning = in.MySQLHighThreadsRunning
	out.MySQLSlowQueries = in.MySQLSlowQueries
	out.MySQLInnoDBLogWaits = in.MySQLInnoDBLogWaits
	out.MySQLRestarted = in.MySQLRestarted
	out.MySQLHighQPS = in.MySQLHighQPS
	out.MySQLHighIncomingBytes = in.MySQLHighIncomingBytes
	out.MySQLHighOutgoingBytes = in.MySQLHighOutgoingBytes
	out.MySQLTooManyOpenFiles = in.MySQLTooManyOpenFiles
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariaDBDatabaseAlertRules.
func (in *MariaDBDatabaseAlertRules) DeepCopy() *MariaDBDatabaseAlertRules {
	if in == nil {
		return nil
	}
	out := new(MariaDBDatabaseAlertRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mariadb) DeepCopyInto(out *Mariadb) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mariadb.
func (in *Mariadb) DeepCopy() *Mariadb {
	if in == nil {
		return nil
	}
	out := new(Mariadb)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Mariadb) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariadbList) DeepCopyInto(out *MariadbList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Mariadb, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariadbList.
func (in *MariadbList) DeepCopy() *MariadbList {
	if in == nil {
		return nil
	}
	out := new(MariadbList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MariadbList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariadbSpec) DeepCopyInto(out *MariadbSpec) {
	*out = *in
	out.Metadata = in.Metadata
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariadbSpec.
func (in *MariadbSpec) DeepCopy() *MariadbSpec {
	if in == nil {
		return nil
	}
	out := new(MariadbSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariadbSpecSpec) DeepCopyInto(out *MariadbSpecSpec) {
	*out = *in
	in.Alert.DeepCopyInto(&out.Alert)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariadbSpecSpec.
func (in *MariadbSpecSpec) DeepCopy() *MariadbSpecSpec {
	if in == nil {
		return nil
	}
	out := new(MariadbSpecSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAlert) DeepCopyInto(out *MongoDBAlert) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AdditionalRuleLabels != nil {
		in, out := &in.AdditionalRuleLabels, &out.AdditionalRuleLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Rules = in.Rules
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAlert.
func (in *MongoDBAlert) DeepCopy() *MongoDBAlert {
	if in == nil {
		return nil
	}
	out := new(MongoDBAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAlertRules) DeepCopyInto(out *MongoDBAlertRules) {
	*out = *in
	out.MongodbVirtualMemoryUsage = in.MongodbVirtualMemoryUsage
	out.MongodbReplicationLag = in.MongodbReplicationLag
	out.MongodbNumberCursorsOpen = in.MongodbNumberCursorsOpen
	out.MongodbCursorsTimeouts = in.MongodbCursorsTimeouts
	out.MongodbTooManyConnections = in.MongodbTooManyConnections
	out.MongoDBPhaseCritical = in.MongoDBPhaseCritical
	out.MongoDBDown = in.MongoDBDown
	out.MongodbHighLatency = in.MongodbHighLatency
	out.MongodbHighTicketUtilization = in.MongodbHighTicketUtilization
	out.MongodbRecurrentCursorTimeout = in.MongodbRecurrentCursorTimeout
	out.MongodbRecurrentMemoryPageFaults = in.MongodbRecurrentMemoryPageFaults
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAlertRules.
func (in *MongoDBAlertRules) DeepCopy() *MongoDBAlertRules {
	if in == nil {
		return nil
	}
	out := new(MongoDBAlertRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mongodb) DeepCopyInto(out *Mongodb) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mongodb.
func (in *Mongodb) DeepCopy() *Mongodb {
	if in == nil {
		return nil
	}
	out := new(Mongodb)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Mongodb) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbList) DeepCopyInto(out *MongodbList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Mongodb, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbList.
func (in *MongodbList) DeepCopy() *MongodbList {
	if in == nil {
		return nil
	}
	out := new(MongodbList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongodbList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbSpec) DeepCopyInto(out *MongodbSpec) {
	*out = *in
	out.Metadata = in.Metadata
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbSpec.
func (in *MongodbSpec) DeepCopy() *MongodbSpec {
	if in == nil {
		return nil
	}
	out := new(MongodbSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbSpecSpec) DeepCopyInto(out *MongodbSpecSpec) {
	*out = *in
	in.Alert.DeepCopyInto(&out.Alert)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbSpecSpec.
func (in *MongodbSpecSpec) DeepCopy() *MongodbSpecSpec {
	if in == nil {
		return nil
	}
	out := new(MongodbSpecSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLAlert) DeepCopyInto(out *MySQLAlert) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AdditionalRuleLabels != nil {
		in, out := &in.AdditionalRuleLabels, &out.AdditionalRuleLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Groups = in.Groups
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLAlert.
func (in *MySQLAlert) DeepCopy() *MySQLAlert {
	if in == nil {
		return nil
	}
	out := new(MySQLAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLAlertGroups) DeepCopyInto(out *MySQLAlertGroups) {
	*out = *in
	out.Database = in.Database
	out.Group = in.Group
	out.Provisioner = in.Provisioner
	out.OpsManager = in.OpsManager
	out.Stash = in.Stash
	out.SchemaManager = in.SchemaManager
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLAlertGroups.
func (in *MySQLAlertGroups) DeepCopy() *MySQLAlertGroups {
	if in == nil {
		return nil
	}
	out := new(MySQLAlertGroups)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLDatabaseAlert) DeepCopyInto(out *MySQLDatabaseAlert) {
	*out = *in
	out.Rules = in.Rules
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLDatabaseAlert.
func (in *MySQLDatabaseAlert) DeepCopy() *MySQLDatabaseAlert {
	if in == nil {
		return nil
	}
	out := new(MySQLDatabaseAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLDatabaseAlertRules) DeepCopyInto(out *MySQLDatabaseAlertRules) {
	*out = *in
	out.MySQLInstanceDown = in.MySQLInstanceDown
	out.MySQLServiceDown = in.MySQLServiceDown
	out.MySQLTooManyConnections = in.MySQLTooManyConnections
	out.MySQLHighThreadsRunning = in.MySQLHighThreadsRunning
	out.MySQLSlowQueries = in.MySQLSlowQueries
	out.MySQLInnoDBLogWaits = in.MySQLInnoDBLogWaits
	out.MySQLRestarted = in.MySQLRestarted
	out.MySQLHighQPS = in.MySQLHighQPS
	out.MySQLHighIncomingBytes = in.MySQLHighIncomingBytes
	out.MySQLHighOutgoingBytes = in.MySQLHighOutgoingBytes
	out.MySQLTooManyOpenFiles = in.MySQLTooManyOpenFiles
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLDatabaseAlertRules.
func (in *MySQLDatabaseAlertRules) DeepCopy() *MySQLDatabaseAlertRules {
	if in == nil {
		return nil
	}
	out := new(MySQLDatabaseAlertRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLGroupAlert) DeepCopyInto(out *MySQLGroupAlert) {
	*out = *in
	out.Rules = in.Rules
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLGroupAlert.
func (in *MySQLGroupAlert) DeepCopy() *MySQLGroupAlert {
	if in == nil {
		return nil
	}
	out := new(MySQLGroupAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLGroupAlertRules) DeepCopyInto(out *MySQLGroupAlertRules) {
	*out = *in
	out.MySQLHighReplicationDelay = in.MySQLHighReplicationDelay
	out.MySQLHighReplicationTransportTime = in.MySQLHighReplicationTransportTime
	out.MySQLHighReplicationApplyTime = in.MySQLHighReplicationApplyTime
	out.MySQLReplicationHighTransactionTime = in.MySQLReplicationHighTransactionTime
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLGroupAlertRules.
func (in *MySQLGroupAlertRules) DeepCopy() *MySQLGroupAlertRules {
	if in == nil {
		return nil
	}
	out := new(MySQLGroupAlertRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mysql) DeepCopyInto(out *Mysql) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mysql.
func (in *Mysql) DeepCopy() *Mysql {
	if in == nil {
		return nil
	}
	out := new(Mysql)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Mysql) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlList) DeepCopyInto(out *MysqlList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Mysql, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlList.
func (in *MysqlList) DeepCopy() *MysqlList {
	if in == nil {
		return nil
	}
	out := new(MysqlList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MysqlList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlSpec) DeepCopyInto(out *MysqlSpec) {
	*out = *in
	out.Metadata = in.Metadata
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlSpec.
func (in *MysqlSpec) DeepCopy() *MysqlSpec {
	if in == nil {
		return nil
	}
	out := new(MysqlSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MysqlSpecSpec) DeepCopyInto(out *MysqlSpecSpec) {
	*out = *in
	in.Alert.DeepCopyInto(&out.Alert)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MysqlSpecSpec.
func (in *MysqlSpecSpec) DeepCopy() *MysqlSpecSpec {
	if in == nil {
		return nil
	}
	out := new(MysqlSpecSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsManagerAlert) DeepCopyInto(out *OpsManagerAlert) {
	*out = *in
	out.Rules = in.Rules
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsManagerAlert.
func (in *OpsManagerAlert) DeepCopy() *OpsManagerAlert {
	if in == nil {
		return nil
	}
	out := new(OpsManagerAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsManagerAlertRules) DeepCopyInto(out *OpsManagerAlertRules) {
	*out = *in
	out.OpsRequestOnProgress = in.OpsRequestOnProgress
	out.OpsRequestStatusProgressingToLong = in.OpsRequestStatusProgressingToLong
	out.OpsRequestFailed = in.OpsRequestFailed
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsManagerAlertRules.
func (in *OpsManagerAlertRules) DeepCopy() *OpsManagerAlertRules {
	if in == nil {
		return nil
	}
	out := new(OpsManagerAlertRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Postgres) DeepCopyInto(out *Postgres) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Postgres.
func (in *Postgres) DeepCopy() *Postgres {
	if in == nil {
		return nil
	}
	out := new(Postgres)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Postgres) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresAlert) DeepCopyInto(out *PostgresAlert) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AdditionalRuleLabels != nil {
		in, out := &in.AdditionalRuleLabels, &out.AdditionalRuleLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Groups = in.Groups
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresAlert.
func (in *PostgresAlert) DeepCopy() *PostgresAlert {
	if in == nil {
		return nil
	}
	out := new(PostgresAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresAlertGroups) DeepCopyInto(out *PostgresAlertGroups) {
	*out = *in
	out.Database = in.Database
	out.Provisioner = in.Provisioner
	out.OpsManager = in.OpsManager
	out.Stash = in.Stash
	out.SchemaManager = in.SchemaManager
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresAlertGroups.
func (in *PostgresAlertGroups) DeepCopy() *PostgresAlertGroups {
	if in == nil {
		return nil
	}
	out := new(PostgresAlertGroups)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresDatabaseAlert) DeepCopyInto(out *PostgresDatabaseAlert) {
	*out = *in
	out.Rules = in.Rules
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresDatabaseAlert.
func (in *PostgresDatabaseAlert) DeepCopy() *PostgresDatabaseAlert {
	if in == nil {
		return nil
	}
	out := new(PostgresDatabaseAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresDatabaseAlertRules) DeepCopyInto(out *PostgresDatabaseAlertRules) {
	*out = *in
	out.PostgresInstanceDown = in.PostgresInstanceDown
	out.PostgresRestarted = in.PostgresRestarted
	out.PostgresTooManyConnections = in.PostgresTooManyConnections
	out.PostgresqlNotEnoughConnections = in.PostgresqlNotEnoughConnections
	out.PostgresSlowQueries = in.PostgresSlowQueries
	out.PostgresqlReplicationLag = in.PostgresqlReplicationLag
	out.PostgresqlHighRollbackRate = in.PostgresqlHighRollbackRate
	out.PostgresqlSplitBrain = in.PostgresqlSplitBrain
	out.PostgresqlTooManyLocksAcquired = in.PostgresqlTooManyLocksAcquired
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresDatabaseAlertRules.
func (in *PostgresDatabaseAlertRules) DeepCopy() *PostgresDatabaseAlertRules {
	if in == nil {
		return nil
	}
	out := new(PostgresDatabaseAlertRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresList) DeepCopyInto(out *PostgresList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Postgres, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresList.
func (in *PostgresList) DeepCopy() *PostgresList {
	if in == nil {
		return nil
	}
	out := new(PostgresList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresSpec) DeepCopyInto(out *PostgresSpec) {
	*out = *in
	out.Metadata = in.Metadata
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresSpec.
func (in *PostgresSpec) DeepCopy() *PostgresSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresSpecSpec) DeepCopyInto(out *PostgresSpecSpec) {
	*out = *in
	in.Alert.DeepCopyInto(&out.Alert)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresSpecSpec.
func (in *PostgresSpecSpec) DeepCopy() *PostgresSpecSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresSpecSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionerAlert) DeepCopyInto(out *ProvisionerAlert) {
	*out = *in
	out.Rules = in.Rules
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionerAlert.
func (in *ProvisionerAlert) DeepCopy() *ProvisionerAlert {
	if in == nil {
		return nil
	}
	out := new(ProvisionerAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionerAlertRules) DeepCopyInto(out *ProvisionerAlertRules) {
	*out = *in
	out.AppPhaseNotReady = in.AppPhaseNotReady
	out.AppPhaseCritical = in.AppPhaseCritical
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionerAlertRules.
func (in *ProvisionerAlertRules) DeepCopy() *ProvisionerAlertRules {
	if in == nil {
		return nil
	}
	out := new(ProvisionerAlertRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaManagerAlert) DeepCopyInto(out *SchemaManagerAlert) {
	*out = *in
	out.Rules = in.Rules
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaManagerAlert.
func (in *SchemaManagerAlert) DeepCopy() *SchemaManagerAlert {
	if in == nil {
		return nil
	}
	out := new(SchemaManagerAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaManagerAlertRules) DeepCopyInto(out *SchemaManagerAlertRules) {
	*out = *in
	out.SchemaPendingForTooLong = in.SchemaPendingForTooLong
	out.SchemaInProgressForTooLong = in.SchemaInProgressForTooLong
	out.SchemaTerminatingForTooLong = in.SchemaTerminatingForTooLong
	out.SchemaFailed = in.SchemaFailed
	out.SchemaExpired = in.SchemaExpired
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaManagerAlertRules.
func (in *SchemaManagerAlertRules) DeepCopy() *SchemaManagerAlertRules {
	if in == nil {
		return nil
	}
	out := new(SchemaManagerAlertRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashAlert) DeepCopyInto(out *StashAlert) {
	*out = *in
	out.Rules = in.Rules
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashAlert.
func (in *StashAlert) DeepCopy() *StashAlert {
	if in == nil {
		return nil
	}
	out := new(StashAlert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StashAlertRules) DeepCopyInto(out *StashAlertRules) {
	*out = *in
	out.BackupSessionFailed = in.BackupSessionFailed
	out.RestoreSessionFailed = in.RestoreSessionFailed
	out.NoBackupSessionForTooLong = in.NoBackupSessionForTooLong
	out.RepositoryCorrupted = in.RepositoryCorrupted
	out.RepositoryStorageRunningLow = in.RepositoryStorageRunningLow
	out.BackupSessionPeriodTooLong = in.BackupSessionPeriodTooLong
	out.RestoreSessionPeriodTooLong = in.RestoreSessionPeriodTooLong
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StashAlertRules.
func (in *StashAlertRules) DeepCopy() *StashAlertRules {
	if in == nil {
		return nil
	}
	out := new(StashAlertRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringValAlert) DeepCopyInto(out *StringValAlert) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringValAlert.
func (in *StringValAlert) DeepCopy() *StringValAlert {
	if in == nil {
		return nil
	}
	out := new(StringValAlert)
	in.DeepCopyInto(out)
	return out
}
