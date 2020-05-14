package test

import (
    "fmt"
    "os"
    "strings"
    "testing"

    "github.com/gruntwork-io/terratest/modules/random"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/gruntwork-io/terratest/modules/test-structure"
    "github.com/stretchr/testify/require"

    test_polkadot "github.com/vladimir-babichev/terratest-polkadot-deployer"
)

func TestTerraformClusterCreation(t *testing.T) {
    t.Parallel()

    terraformDir := "../"

    // Test configuration
    test_structure.RunTestStage(t, "create_terratest_options", func() {
        nodeCount := 2
        nodePort := 30100
        uniqueID := strings.ToLower(random.UniqueId())
        clusterName := fmt.Sprintf("test-polkadot-%s", uniqueID)
        doToken := os.Getenv("DIGITALOCEAN_TOKEN")
        require.NotEmpty(t, doToken, "DIGITALOCEAN_TOKEN variable is not set")

        terraformOptions := &terraform.Options{
            TerraformDir: terraformDir,
            Vars: map[string]interface{} {
                "cluster_name": clusterName,
                "do_token":     doToken,
                "location":     "lon1",
                "machine_type": "s-1vcpu-2gb",
                "node_count":   nodeCount,
            },
            NoColor: true,
        }

        test_structure.SaveInt(t, terraformDir, "nodeCount", nodeCount)
        test_structure.SaveInt(t, terraformDir, "nodePort", nodePort)
        test_structure.SaveString(t, terraformDir, "clusterName", clusterName)
        test_structure.SaveString(t, terraformDir, "uniqueID", uniqueID)
        test_structure.SaveTerraformOptions(t, terraformDir, terraformOptions)
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
        test_polkadot.ValidateNodeCount(t, terraformDir)
    })

    // Validate external connectivity to the service
    test_structure.RunTestStage(t, "validate_service", func() {
        test_polkadot.ValidateServiceAvailability(t, terraformDir)
    })
}
