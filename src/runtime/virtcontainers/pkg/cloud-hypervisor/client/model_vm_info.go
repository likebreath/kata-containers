/*
 * Cloud Hypervisor API
 *
 * Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// VmInfo Virtual Machine information
type VmInfo struct {
	Config VmConfig `json:"config"`
	State string `json:"state"`
	MemoryActualSize int64 `json:"memory_actual_size,omitempty"`
	DeviceTree map[string]DeviceNode `json:"device_tree,omitempty"`
}
