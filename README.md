# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2023-01-01
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/awlsring/proxmox-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**AddCorosyncNode**](docs/DefaultApi.md#addcorosyncnode) | **Post** /cluster/config/nodes/{node} | 
*DefaultApi* | [**AddCustomNodeCertificate**](docs/DefaultApi.md#addcustomnodecertificate) | **Post** /nodes/{node}/certificates/custom | 
*DefaultApi* | [**AddRepository**](docs/DefaultApi.md#addrepository) | **Put** /nodes/{node}/apt/repositories | 
*DefaultApi* | [**ApplyNetworkInterfaceConfiguration**](docs/DefaultApi.md#applynetworkinterfaceconfiguration) | **Put** /nodes/{node}/network | 
*DefaultApi* | [**ApplyVirtualMachineConfigurationAsync**](docs/DefaultApi.md#applyvirtualmachineconfigurationasync) | **Post** /nodes/{node}/qemu/{vmId}/config | 
*DefaultApi* | [**ApplyVirtualMachineConfigurationSync**](docs/DefaultApi.md#applyvirtualmachineconfigurationsync) | **Put** /nodes/{node}/qemu/{vmId}/config | 
*DefaultApi* | [**ChangeRepositoryProperties**](docs/DefaultApi.md#changerepositoryproperties) | **Post** /nodes/{node}/apt/repositories | 
*DefaultApi* | [**CloneVirtualMachine**](docs/DefaultApi.md#clonevirtualmachine) | **Post** /nodes/{node}/qemu/{vmId}/clone | 
*DefaultApi* | [**CreateClusterConfig**](docs/DefaultApi.md#createclusterconfig) | **Post** /cluster/config | 
*DefaultApi* | [**CreateDirectory**](docs/DefaultApi.md#createdirectory) | **Post** /nodes/{node}/disks/directory | 
*DefaultApi* | [**CreateLVM**](docs/DefaultApi.md#createlvm) | **Post** /nodes/{node}/disks/lvm | 
*DefaultApi* | [**CreateLVMThin**](docs/DefaultApi.md#createlvmthin) | **Post** /nodes/{node}/disks/lvmthin | 
*DefaultApi* | [**CreateNetworkInterface**](docs/DefaultApi.md#createnetworkinterface) | **Post** /nodes/{node}/network | 
*DefaultApi* | [**CreatePool**](docs/DefaultApi.md#createpool) | **Post** /pools | 
*DefaultApi* | [**CreateStorage**](docs/DefaultApi.md#createstorage) | **Post** /storage | 
*DefaultApi* | [**CreateTicket**](docs/DefaultApi.md#createticket) | **Post** /access/ticket | 
*DefaultApi* | [**CreateVirtualMachine**](docs/DefaultApi.md#createvirtualmachine) | **Post** /nodes/{node}/qemu | 
*DefaultApi* | [**CreateVirtualMachineTemplate**](docs/DefaultApi.md#createvirtualmachinetemplate) | **Post** /nodes/{node}/qemu/{vmId}/template | 
*DefaultApi* | [**CreateZFSPool**](docs/DefaultApi.md#createzfspool) | **Post** /nodes/{node}/disks/zfs | 
*DefaultApi* | [**DeleteDirectory**](docs/DefaultApi.md#deletedirectory) | **Delete** /nodes/{node}/disks/directory/{name} | 
*DefaultApi* | [**DeleteLVM**](docs/DefaultApi.md#deletelvm) | **Delete** /nodes/{node}/disks/lvm/{name} | 
*DefaultApi* | [**DeleteLVMThin**](docs/DefaultApi.md#deletelvmthin) | **Delete** /nodes/{node}/disks/lvmthin/{name} | 
*DefaultApi* | [**DeleteNetworkInterface**](docs/DefaultApi.md#deletenetworkinterface) | **Delete** /nodes/{node}/network/{interface} | 
*DefaultApi* | [**DeleteNodeCertificate**](docs/DefaultApi.md#deletenodecertificate) | **Delete** /nodes/{node}/certificates/acme/certificate | 
*DefaultApi* | [**DeletePool**](docs/DefaultApi.md#deletepool) | **Delete** /pools/{poolId} | 
*DefaultApi* | [**DeleteStorage**](docs/DefaultApi.md#deletestorage) | **Delete** /storage/{storage} | 
*DefaultApi* | [**DeleteVirtualMachine**](docs/DefaultApi.md#deletevirtualmachine) | **Delete** /nodes/{node}/qemu/{vmId} | 
*DefaultApi* | [**DeleteZFSPool**](docs/DefaultApi.md#deletezfspool) | **Delete** /nodes/{node}/disks/zfs/{name} | 
*DefaultApi* | [**GetAccessControlList**](docs/DefaultApi.md#getaccesscontrollist) | **Get** /access/acl | 
*DefaultApi* | [**GetClusterApiVersion**](docs/DefaultApi.md#getclusterapiversion) | **Get** /cluster/config/apiversion | 
*DefaultApi* | [**GetClusterJoinInformation**](docs/DefaultApi.md#getclusterjoininformation) | **Get** /cluster/config/join | 
*DefaultApi* | [**GetClusterTotemSettings**](docs/DefaultApi.md#getclustertotemsettings) | **Get** /cluster/config/totem | 
*DefaultApi* | [**GetNetworkInterface**](docs/DefaultApi.md#getnetworkinterface) | **Get** /nodes/{node}/network/{interface} | 
*DefaultApi* | [**GetPackageChangelog**](docs/DefaultApi.md#getpackagechangelog) | **Get** /nodes/{node}/apt/changelog | 
*DefaultApi* | [**GetPendingVirtualMachineCloudInitChanges**](docs/DefaultApi.md#getpendingvirtualmachinecloudinitchanges) | **Get** /nodes/{node}/qemu/{vmId}/cloudinit | 
*DefaultApi* | [**GetPool**](docs/DefaultApi.md#getpool) | **Get** /pools/{poolId} | 
*DefaultApi* | [**GetSmartHealth**](docs/DefaultApi.md#getsmarthealth) | **Get** /nodes/{node}/disks/smart | 
*DefaultApi* | [**GetStorage**](docs/DefaultApi.md#getstorage) | **Get** /storage/{storage} | 
*DefaultApi* | [**GetVersion**](docs/DefaultApi.md#getversion) | **Get** /version | 
*DefaultApi* | [**GetVirtualMachineCloudInit**](docs/DefaultApi.md#getvirtualmachinecloudinit) | **Get** /nodes/{node}/qemu/{vmId}/cloudinit/dump | 
*DefaultApi* | [**GetVirtualMachineConfiguration**](docs/DefaultApi.md#getvirtualmachineconfiguration) | **Get** /nodes/{node}/qemu/{vmId}/config | 
*DefaultApi* | [**GetVirtualMachineFeatureSupport**](docs/DefaultApi.md#getvirtualmachinefeaturesupport) | **Get** /nodes/{node}/qemu/{vmId}/feature | 
*DefaultApi* | [**GetVirtualMachineStatus**](docs/DefaultApi.md#getvirtualmachinestatus) | **Get** /nodes/{node}/qemu/{vmId}/status/current | 
*DefaultApi* | [**GetZFSPoolStatus**](docs/DefaultApi.md#getzfspoolstatus) | **Get** /nodes/{node}/disks/zfs/{name} | 
*DefaultApi* | [**InitializeGPT**](docs/DefaultApi.md#initializegpt) | **Post** /nodes/{node}/disks/smart | 
*DefaultApi* | [**JoinCluster**](docs/DefaultApi.md#joincluster) | **Post** /cluster/config/join | 
*DefaultApi* | [**ListCorosyncNodes**](docs/DefaultApi.md#listcorosyncnodes) | **Get** /cluster/config/nodes | 
*DefaultApi* | [**ListCpuCapabilities**](docs/DefaultApi.md#listcpucapabilities) | **Get** /nodes/{node}/capabilities/qemu/cpu | 
*DefaultApi* | [**ListDirectories**](docs/DefaultApi.md#listdirectories) | **Get** /nodes/{node}/disks/directory | 
*DefaultApi* | [**ListDisks**](docs/DefaultApi.md#listdisks) | **Get** /nodes/{node}/disks/list | 
*DefaultApi* | [**ListLVMThins**](docs/DefaultApi.md#listlvmthins) | **Get** /nodes/{node}/disks/lvmthin | 
*DefaultApi* | [**ListLVMs**](docs/DefaultApi.md#listlvms) | **Get** /nodes/{node}/disks/lvm | 
*DefaultApi* | [**ListMachineCapabilities**](docs/DefaultApi.md#listmachinecapabilities) | **Get** /nodes/{node}/capabilities/qemu/machines | 
*DefaultApi* | [**ListNetworkInterfaces**](docs/DefaultApi.md#listnetworkinterfaces) | **Get** /nodes/{node}/network | 
*DefaultApi* | [**ListNodeCertificates**](docs/DefaultApi.md#listnodecertificates) | **Get** /nodes/{node}/certificates/info | 
*DefaultApi* | [**ListNodes**](docs/DefaultApi.md#listnodes) | **Get** /nodes | 
*DefaultApi* | [**ListPackages**](docs/DefaultApi.md#listpackages) | **Get** /nodes/{node}/apt/versions | 
*DefaultApi* | [**ListPciDeviceMediatedDevices**](docs/DefaultApi.md#listpcidevicemediateddevices) | **Get** /nodes/{node}/hardware/pci/{deviceId} | 
*DefaultApi* | [**ListPciDevices**](docs/DefaultApi.md#listpcidevices) | **Get** /nodes/{node}/hardware/pci | 
*DefaultApi* | [**ListPendingVirtualMachineConfigurationChanges**](docs/DefaultApi.md#listpendingvirtualmachineconfigurationchanges) | **Get** /nodes/{node}/qemu/{vmId}/pending | 
*DefaultApi* | [**ListPools**](docs/DefaultApi.md#listpools) | **Get** /pools | 
*DefaultApi* | [**ListRepositoriesInformation**](docs/DefaultApi.md#listrepositoriesinformation) | **Get** /nodes/{node}/apt/repository | 
*DefaultApi* | [**ListStorage**](docs/DefaultApi.md#liststorage) | **Get** /storage | 
*DefaultApi* | [**ListUpdates**](docs/DefaultApi.md#listupdates) | **Get** /nodes/{node}/apt/update | 
*DefaultApi* | [**ListUsbDevices**](docs/DefaultApi.md#listusbdevices) | **Get** /nodes/{node}/hardware/usb | 
*DefaultApi* | [**ListVirtualMachines**](docs/DefaultApi.md#listvirtualmachines) | **Get** /nodes/{node}/qemu | 
*DefaultApi* | [**ListZFSPools**](docs/DefaultApi.md#listzfspools) | **Get** /nodes/{node}/disks/zfs | 
*DefaultApi* | [**ModifyPool**](docs/DefaultApi.md#modifypool) | **Put** /pools | 
*DefaultApi* | [**ModifyStorage**](docs/DefaultApi.md#modifystorage) | **Put** /storage/{storage} | 
*DefaultApi* | [**OrderNodeCertificate**](docs/DefaultApi.md#ordernodecertificate) | **Post** /nodes/{node}/certificates/acme/certificate | 
*DefaultApi* | [**RegenerateVirtualMachineCloudInit**](docs/DefaultApi.md#regeneratevirtualmachinecloudinit) | **Put** /nodes/{node}/qemu/{vmId}/cloudinit | 
*DefaultApi* | [**RemoveCorosyncNode**](docs/DefaultApi.md#removecorosyncnode) | **Delete** /cluster/config/nodes/{node} | 
*DefaultApi* | [**RenewNodeCertificate**](docs/DefaultApi.md#renewnodecertificate) | **Put** /nodes/{node}/certificates/acme/certificate | 
*DefaultApi* | [**ResizeVirtualMachineDisk**](docs/DefaultApi.md#resizevirtualmachinedisk) | **Put** /nodes/{node}/qemu/{vmId}/resize | 
*DefaultApi* | [**RevertNetworkInterfaceConfiguration**](docs/DefaultApi.md#revertnetworkinterfaceconfiguration) | **Delete** /nodes/{node}/network | 
*DefaultApi* | [**UnlinkVirtualMachineDisks**](docs/DefaultApi.md#unlinkvirtualmachinedisks) | **Put** /nodes/{node}/qemu/{vmId}/unlink | 
*DefaultApi* | [**UpdateAccessControlList**](docs/DefaultApi.md#updateaccesscontrollist) | **Put** /access/acl | 
*DefaultApi* | [**UpdateNetworkInterface**](docs/DefaultApi.md#updatenetworkinterface) | **Put** /nodes/{node}/network/{interface} | 
*DefaultApi* | [**WipeDisk**](docs/DefaultApi.md#wipedisk) | **Put** /nodes/{node}/disks/smart | 


## Documentation For Models

 - [AccessControlSummary](docs/AccessControlSummary.md)
 - [AccessControlType](docs/AccessControlType.md)
 - [AddCorosyncNodeRequestContent](docs/AddCorosyncNodeRequestContent.md)
 - [AddCorosyncNodeResponseContent](docs/AddCorosyncNodeResponseContent.md)
 - [AddCustomNodeCertificateRequestContent](docs/AddCustomNodeCertificateRequestContent.md)
 - [AddCustomNodeCertificateResponseContent](docs/AddCustomNodeCertificateResponseContent.md)
 - [ApplyNetworkInterfaceConfigurationResponseContent](docs/ApplyNetworkInterfaceConfigurationResponseContent.md)
 - [ApplyVirtualMachineConfigurationAsyncRequestContent](docs/ApplyVirtualMachineConfigurationAsyncRequestContent.md)
 - [ApplyVirtualMachineConfigurationAsyncResponseContent](docs/ApplyVirtualMachineConfigurationAsyncResponseContent.md)
 - [ApplyVirtualMachineConfigurationSyncRequestContent](docs/ApplyVirtualMachineConfigurationSyncRequestContent.md)
 - [CloneVirtualMachineDiskFormat](docs/CloneVirtualMachineDiskFormat.md)
 - [CloneVirtualMachineRequestContent](docs/CloneVirtualMachineRequestContent.md)
 - [CloneVirtualMachineResponseContent](docs/CloneVirtualMachineResponseContent.md)
 - [CloudInitType](docs/CloudInitType.md)
 - [CorosyncNodeSummary](docs/CorosyncNodeSummary.md)
 - [CorosyncSettings](docs/CorosyncSettings.md)
 - [CpuCapabilitySummary](docs/CpuCapabilitySummary.md)
 - [CreateClusterConfigRequestContent](docs/CreateClusterConfigRequestContent.md)
 - [CreateClusterConfigResponseContent](docs/CreateClusterConfigResponseContent.md)
 - [CreateDirectoryRequestContent](docs/CreateDirectoryRequestContent.md)
 - [CreateDirectoryResponseContent](docs/CreateDirectoryResponseContent.md)
 - [CreateLVMRequestContent](docs/CreateLVMRequestContent.md)
 - [CreateLVMResponseContent](docs/CreateLVMResponseContent.md)
 - [CreateLVMThinRequestContent](docs/CreateLVMThinRequestContent.md)
 - [CreateLVMThinResponseContent](docs/CreateLVMThinResponseContent.md)
 - [CreateNetworkInterfaceRequestContent](docs/CreateNetworkInterfaceRequestContent.md)
 - [CreatePoolRequestContent](docs/CreatePoolRequestContent.md)
 - [CreateStorageRequestContent](docs/CreateStorageRequestContent.md)
 - [CreateStorageResponseContent](docs/CreateStorageResponseContent.md)
 - [CreateTicketRequestContent](docs/CreateTicketRequestContent.md)
 - [CreateTicketResponseContent](docs/CreateTicketResponseContent.md)
 - [CreateVirtualMachineRequestContent](docs/CreateVirtualMachineRequestContent.md)
 - [CreateVirtualMachineResponseContent](docs/CreateVirtualMachineResponseContent.md)
 - [CreateVirtualMachineTemplateRequestContent](docs/CreateVirtualMachineTemplateRequestContent.md)
 - [CreateVirtualMachineTemplateResponseContent](docs/CreateVirtualMachineTemplateResponseContent.md)
 - [CreateZFSPoolRequestContent](docs/CreateZFSPoolRequestContent.md)
 - [CreateZFSPoolResponseContent](docs/CreateZFSPoolResponseContent.md)
 - [DeleteDirectoryResponseContent](docs/DeleteDirectoryResponseContent.md)
 - [DeleteLVMResponseContent](docs/DeleteLVMResponseContent.md)
 - [DeleteLVMThinResponseContent](docs/DeleteLVMThinResponseContent.md)
 - [DeleteNodeCertificateResponseContent](docs/DeleteNodeCertificateResponseContent.md)
 - [DeleteVirtualMachineResponseContent](docs/DeleteVirtualMachineResponseContent.md)
 - [DeleteZFSPoolResponseContent](docs/DeleteZFSPoolResponseContent.md)
 - [DirectoryFileSystem](docs/DirectoryFileSystem.md)
 - [DirectorySummary](docs/DirectorySummary.md)
 - [DiskSummary](docs/DiskSummary.md)
 - [DiskType](docs/DiskType.md)
 - [DiskTypeFilter](docs/DiskTypeFilter.md)
 - [FileInfoSummary](docs/FileInfoSummary.md)
 - [FileRepositorySummary](docs/FileRepositorySummary.md)
 - [FileSummary](docs/FileSummary.md)
 - [GetAccessControlListResponseContent](docs/GetAccessControlListResponseContent.md)
 - [GetClusterApiVersionResponseContent](docs/GetClusterApiVersionResponseContent.md)
 - [GetClusterJoinInformationResponseContent](docs/GetClusterJoinInformationResponseContent.md)
 - [GetClusterTotemSettingsResponseContent](docs/GetClusterTotemSettingsResponseContent.md)
 - [GetNetworkInterfaceResponseContent](docs/GetNetworkInterfaceResponseContent.md)
 - [GetPackageChangelogResponseContent](docs/GetPackageChangelogResponseContent.md)
 - [GetPendingVirtualMachineCloudInitChangesResponseContent](docs/GetPendingVirtualMachineCloudInitChangesResponseContent.md)
 - [GetPoolResponseContent](docs/GetPoolResponseContent.md)
 - [GetSmartHealthResponseContent](docs/GetSmartHealthResponseContent.md)
 - [GetStorageResponseContent](docs/GetStorageResponseContent.md)
 - [GetVersionResponseContent](docs/GetVersionResponseContent.md)
 - [GetVirtualMachineCloudInitResponseContent](docs/GetVirtualMachineCloudInitResponseContent.md)
 - [GetVirtualMachineConfigurationResponseContent](docs/GetVirtualMachineConfigurationResponseContent.md)
 - [GetVirtualMachineFeatureSupportResponseContent](docs/GetVirtualMachineFeatureSupportResponseContent.md)
 - [GetVirtualMachineStatusResponseContent](docs/GetVirtualMachineStatusResponseContent.md)
 - [GetZFSPoolStatusResponseContent](docs/GetZFSPoolStatusResponseContent.md)
 - [InitializeGPTRequestContent](docs/InitializeGPTRequestContent.md)
 - [InitializeGPTResponseContent](docs/InitializeGPTResponseContent.md)
 - [InternalServerErrorResponseContent](docs/InternalServerErrorResponseContent.md)
 - [InvalidInputErrorResponseContent](docs/InvalidInputErrorResponseContent.md)
 - [JoinClusterRequestContent](docs/JoinClusterRequestContent.md)
 - [JoinClusterResponseContent](docs/JoinClusterResponseContent.md)
 - [JoinInformation](docs/JoinInformation.md)
 - [LVMSummary](docs/LVMSummary.md)
 - [LVMThinSummary](docs/LVMThinSummary.md)
 - [LinkSummary](docs/LinkSummary.md)
 - [ListCorosyncNodesResponseContent](docs/ListCorosyncNodesResponseContent.md)
 - [ListCpuCapabilitiesResponseContent](docs/ListCpuCapabilitiesResponseContent.md)
 - [ListDirectoriesResponseContent](docs/ListDirectoriesResponseContent.md)
 - [ListDisksResponseContent](docs/ListDisksResponseContent.md)
 - [ListLVMThinsResponseContent](docs/ListLVMThinsResponseContent.md)
 - [ListLVMsResponseContent](docs/ListLVMsResponseContent.md)
 - [ListMachineCapabilitiesResponseContent](docs/ListMachineCapabilitiesResponseContent.md)
 - [ListNetworkInterfacesResponseContent](docs/ListNetworkInterfacesResponseContent.md)
 - [ListNodeCertificatesResponseContent](docs/ListNodeCertificatesResponseContent.md)
 - [ListNodesResponseContent](docs/ListNodesResponseContent.md)
 - [ListPackagesResponseContent](docs/ListPackagesResponseContent.md)
 - [ListPciDeviceMediatedDevicesResponseContent](docs/ListPciDeviceMediatedDevicesResponseContent.md)
 - [ListPciDevicesResponseContent](docs/ListPciDevicesResponseContent.md)
 - [ListPendingVirtualMachineConfigurationChangesResponseContent](docs/ListPendingVirtualMachineConfigurationChangesResponseContent.md)
 - [ListPoolsResponseContent](docs/ListPoolsResponseContent.md)
 - [ListRepositoriesInformationResponseContent](docs/ListRepositoriesInformationResponseContent.md)
 - [ListStorageResponseContent](docs/ListStorageResponseContent.md)
 - [ListUpdatesResponseContent](docs/ListUpdatesResponseContent.md)
 - [ListUsbDevicesResponseContent](docs/ListUsbDevicesResponseContent.md)
 - [ListVirtualMachinesResponseContent](docs/ListVirtualMachinesResponseContent.md)
 - [ListZFSPoolsResponseContent](docs/ListZFSPoolsResponseContent.md)
 - [MachineCapabilitySummary](docs/MachineCapabilitySummary.md)
 - [MachineType](docs/MachineType.md)
 - [ModifyPoolRequestContent](docs/ModifyPoolRequestContent.md)
 - [ModifyStorageRequestContent](docs/ModifyStorageRequestContent.md)
 - [ModifyStorageResponseContent](docs/ModifyStorageResponseContent.md)
 - [NetworkInterfaceBondHashPolicy](docs/NetworkInterfaceBondHashPolicy.md)
 - [NetworkInterfaceBondMode](docs/NetworkInterfaceBondMode.md)
 - [NetworkInterfaceMethod](docs/NetworkInterfaceMethod.md)
 - [NetworkInterfaceSummary](docs/NetworkInterfaceSummary.md)
 - [NetworkInterfaceType](docs/NetworkInterfaceType.md)
 - [NodeCertificate](docs/NodeCertificate.md)
 - [NodeStatus](docs/NodeStatus.md)
 - [NodeSummary](docs/NodeSummary.md)
 - [OrderNodeCertificateRequestContent](docs/OrderNodeCertificateRequestContent.md)
 - [OrderNodeCertificateResponseContent](docs/OrderNodeCertificateResponseContent.md)
 - [PackageSummary](docs/PackageSummary.md)
 - [PciMediatedDeviceSummary](docs/PciMediatedDeviceSummary.md)
 - [PendingVirtualMachineCloudInitField](docs/PendingVirtualMachineCloudInitField.md)
 - [PoolInfo](docs/PoolInfo.md)
 - [PoolMemberType](docs/PoolMemberType.md)
 - [ProxmoxSupportSummary](docs/ProxmoxSupportSummary.md)
 - [RenewNodeCertificateRequestContent](docs/RenewNodeCertificateRequestContent.md)
 - [RenewNodeCertificateResponseContent](docs/RenewNodeCertificateResponseContent.md)
 - [RepositoriesReport](docs/RepositoriesReport.md)
 - [RepositorySummary](docs/RepositorySummary.md)
 - [StorageConfiguration](docs/StorageConfiguration.md)
 - [StoragePreallocation](docs/StoragePreallocation.md)
 - [StorageSMBVersion](docs/StorageSMBVersion.md)
 - [StorageSummary](docs/StorageSummary.md)
 - [StorageTransport](docs/StorageTransport.md)
 - [StorageType](docs/StorageType.md)
 - [Ticket](docs/Ticket.md)
 - [TotemSummary](docs/TotemSummary.md)
 - [UpdateAccessControlListRequestContent](docs/UpdateAccessControlListRequestContent.md)
 - [UpdateNetworkInterfaceRequestContent](docs/UpdateNetworkInterfaceRequestContent.md)
 - [UpdateSummary](docs/UpdateSummary.md)
 - [UsbDeviceSummary](docs/UsbDeviceSummary.md)
 - [VersionSummary](docs/VersionSummary.md)
 - [VirtualMachineArchitecture](docs/VirtualMachineArchitecture.md)
 - [VirtualMachineBalloonSummary](docs/VirtualMachineBalloonSummary.md)
 - [VirtualMachineBios](docs/VirtualMachineBios.md)
 - [VirtualMachineCloudInitType](docs/VirtualMachineCloudInitType.md)
 - [VirtualMachineConfigLock](docs/VirtualMachineConfigLock.md)
 - [VirtualMachineConfigurationSummary](docs/VirtualMachineConfigurationSummary.md)
 - [VirtualMachineDiskTarget](docs/VirtualMachineDiskTarget.md)
 - [VirtualMachineFeatureSupportSummary](docs/VirtualMachineFeatureSupportSummary.md)
 - [VirtualMachineHighAvailabilityStatus](docs/VirtualMachineHighAvailabilityStatus.md)
 - [VirtualMachineHugePages](docs/VirtualMachineHugePages.md)
 - [VirtualMachineKeyboard](docs/VirtualMachineKeyboard.md)
 - [VirtualMachineNicBlockStatistics](docs/VirtualMachineNicBlockStatistics.md)
 - [VirtualMachineNicStatus](docs/VirtualMachineNicStatus.md)
 - [VirtualMachineOperatingSystem](docs/VirtualMachineOperatingSystem.md)
 - [VirtualMachineScsiControllerType](docs/VirtualMachineScsiControllerType.md)
 - [VirtualMachineStatus](docs/VirtualMachineStatus.md)
 - [VirtualMachineStatusSummary](docs/VirtualMachineStatusSummary.md)
 - [VirtualMachineSummary](docs/VirtualMachineSummary.md)
 - [ZFSCompression](docs/ZFSCompression.md)
 - [ZFSPoolStatusChild](docs/ZFSPoolStatusChild.md)
 - [ZFSPoolStatusSummary](docs/ZFSPoolStatusSummary.md)
 - [ZFSPoolSummary](docs/ZFSPoolSummary.md)
 - [ZFSRaidLevel](docs/ZFSRaidLevel.md)


## Documentation For Authorization



### smithy.api.httpApiKeyAuth


### smithy.api.httpBasicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



