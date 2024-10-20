package models

// The current usage of the virtual resources assigned to the NF instances belonging to a  particular network slice instance.
type ResourceUsage struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	CpuUsage int32 `json:"cpuUsage,omitempty" yaml:"cpuUsage" bson:"cpuUsage,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MemoryUsage int32 `json:"memoryUsage,omitempty" yaml:"memoryUsage" bson:"memoryUsage,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	StorageUsage int32 `json:"storageUsage,omitempty" yaml:"storageUsage" bson:"storageUsage,omitempty"`
}
