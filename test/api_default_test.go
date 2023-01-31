/*
Proxmox

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_openapi_DefaultApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DefaultApiService AddCorosyncNode", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.AddCorosyncNode(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService AddCustomNodeCertificate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.AddCustomNodeCertificate(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService AddRepository", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.AddRepository(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ApplyNetworkInterfaceConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ApplyNetworkInterfaceConfiguration(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ApplyVirtualMachineConfigurationAsync", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string
        var vmId string

        resp, httpRes, err := apiClient.DefaultApi.ApplyVirtualMachineConfigurationAsync(context.Background(), node, vmId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ApplyVirtualMachineConfigurationSync", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string
        var vmId string

        resp, httpRes, err := apiClient.DefaultApi.ApplyVirtualMachineConfigurationSync(context.Background(), node, vmId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ChangeRepositoryProperties", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ChangeRepositoryProperties(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService CreateClusterConfig", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.CreateClusterConfig(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService CreateDirectory", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.CreateDirectory(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService CreateLVM", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.CreateLVM(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService CreateLVMThin", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.CreateLVMThin(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService CreateNetworkInterface", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.CreateNetworkInterface(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService CreatePool", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.CreatePool(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService CreateStorage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.CreateStorage(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService CreateTicket", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.CreateTicket(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService CreateZFSPool", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.CreateZFSPool(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService DeleteDirectory", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string
        var name string

        resp, httpRes, err := apiClient.DefaultApi.DeleteDirectory(context.Background(), node, name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService DeleteLVM", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string
        var name string

        resp, httpRes, err := apiClient.DefaultApi.DeleteLVM(context.Background(), node, name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService DeleteLVMThin", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string
        var name string

        resp, httpRes, err := apiClient.DefaultApi.DeleteLVMThin(context.Background(), node, name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService DeleteNetworkInterface", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string
        var interface_ string

        resp, httpRes, err := apiClient.DefaultApi.DeleteNetworkInterface(context.Background(), node, interface_).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService DeleteNodeCertificate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.DeleteNodeCertificate(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService DeletePool", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var poolId string

        resp, httpRes, err := apiClient.DefaultApi.DeletePool(context.Background(), poolId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService DeleteStorage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var storage string

        resp, httpRes, err := apiClient.DefaultApi.DeleteStorage(context.Background(), storage).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService DeleteZFSPool", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string
        var name string

        resp, httpRes, err := apiClient.DefaultApi.DeleteZFSPool(context.Background(), node, name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetAccessControlList", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.GetAccessControlList(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetClusterApiVersion", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.GetClusterApiVersion(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetClusterJoinInformation", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.GetClusterJoinInformation(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetClusterTotemSettings", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.GetClusterTotemSettings(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetNetworkInterface", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string
        var interface_ string

        resp, httpRes, err := apiClient.DefaultApi.GetNetworkInterface(context.Background(), node, interface_).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetPackageChangelog", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.GetPackageChangelog(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetPool", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var poolId string

        resp, httpRes, err := apiClient.DefaultApi.GetPool(context.Background(), poolId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetSmartHealth", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.GetSmartHealth(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetStorage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var storage string

        resp, httpRes, err := apiClient.DefaultApi.GetStorage(context.Background(), storage).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetVersion", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.GetVersion(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetVirtualMachineConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string
        var vmId string

        resp, httpRes, err := apiClient.DefaultApi.GetVirtualMachineConfiguration(context.Background(), node, vmId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetVirtualMachineStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string
        var vmId string

        resp, httpRes, err := apiClient.DefaultApi.GetVirtualMachineStatus(context.Background(), node, vmId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService GetZFSPoolStatus", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string
        var name string

        resp, httpRes, err := apiClient.DefaultApi.GetZFSPoolStatus(context.Background(), node, name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService InitializeGPT", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.InitializeGPT(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService JoinCluster", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.JoinCluster(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListCorosyncNodes", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.ListCorosyncNodes(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListCpuCapabilities", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListCpuCapabilities(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListDirectories", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListDirectories(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListDisks", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListDisks(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListLVMThins", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListLVMThins(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListLVMs", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListLVMs(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListMachineCapabilities", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListMachineCapabilities(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListNetworkInterfaces", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListNetworkInterfaces(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListNodeCertificates", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListNodeCertificates(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListNodes", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.ListNodes(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListPackages", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListPackages(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListPciDeviceMediatedDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string
        var deviceId string

        resp, httpRes, err := apiClient.DefaultApi.ListPciDeviceMediatedDevices(context.Background(), node, deviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListPciDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListPciDevices(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListPools", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.ListPools(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListRepositoriesInformation", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListRepositoriesInformation(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListStorage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.ListStorage(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListUpdates", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListUpdates(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListUsbDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListUsbDevices(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListVirtualMachines", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListVirtualMachines(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ListZFSPools", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.ListZFSPools(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ModifyPool", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.ModifyPool(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService ModifyStorage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var storage string

        resp, httpRes, err := apiClient.DefaultApi.ModifyStorage(context.Background(), storage).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService OrderNodeCertificate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.OrderNodeCertificate(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService RemoveCorosyncNode", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.RemoveCorosyncNode(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService RenewNodeCertificate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.RenewNodeCertificate(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService RevertNetworkInterfaceConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.RevertNetworkInterfaceConfiguration(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService UpdateAccessControlList", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DefaultApi.UpdateAccessControlList(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService UpdateNetworkInterface", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string
        var interface_ string

        resp, httpRes, err := apiClient.DefaultApi.UpdateNetworkInterface(context.Background(), node, interface_).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService WipeDisk", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var node string

        resp, httpRes, err := apiClient.DefaultApi.WipeDisk(context.Background(), node).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
