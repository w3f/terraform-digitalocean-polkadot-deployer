package test

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/test-structure"
    test_polkadot "github.com/vladimir-babichev/terratest-polkadot-deployer"
)

func TestTerraformClusterCreation(t *testing.T) {
    t.Parallel()

    terraformDir := "../"

    // Test configuration
    test_structure.RunTestStage(t, "create_terratest_options", func() {
        createTerraformOptions(t, terraformDir)
        createHelmOptions(t, terraformDir)
    })

    // At the end of the test, run `terraform destroy` to clean up any resources that were created
    defer test_structure.RunTestStage(t, "cleanup", func() {
        test_polkadot.DestroyTerraformStack(t, terraformDir)
    })

    // Deploy infrastructure
    test_structure.RunTestStage(t, "setup_infrastructure", func() {
        test_polkadot.CreateTerraformStack(t, terraformDir)
    })

    // Configure kubectl
    test_structure.RunTestStage(t, "setup_kubectl", func() {
        test_polkadot.SetupKubeconfig(t, terraformDir)
    })

    // Validate cluster size
    test_structure.RunTestStage(t, "validate_node_count", func() {
        test_polkadot.TestNodeCount(t, terraformDir)
    })

    // Validate external connectivity to the service
    test_structure.RunTestStage(t, "validate_service", func() {
        test_polkadot.TestServiceAvailability(t, terraformDir)
    })

    // Validate no resources will change on subsequent terraform executions
    test_structure.RunTestStage(t, "validate_plan", func() {
        test_polkadot.TestResourceChanges(t, terraformDir)
    })
}
