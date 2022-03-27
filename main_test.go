package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)
//"github.com/stretchr/testify/assert"
func TestTerraformHelloWorldExample(t *testing.T) {
	// retryable errors in terraform testing.
	/*terraformOptionsSecretStore := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "D:/AA-HTCWORKS/Terraform/ec2/modules/secretmanager",
		
	})

	defer terraform.Destroy(t, terraformOptionsSecretStore)

	terraform.InitAndApply(t, terraformOptionsSecretStore)

	secret_id := terraform.Output(t, terraformOptionsSecretStore, "secret_id")
	fmt.Println(secret_id)*/

//map_2 := map[string]string{
//		"RedshiftQueryOwner": "**********",
//	}

//arr:= [4]string{"geek", "gfg", "Geeks1231", "GeeksforGeeks"}

	terraformOptionsSecretValues := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "D:/AA-HTCWORKS/Terraform/ec2/",
		// tags type should be map in terraform file
		Vars: map[string]interface{}{
			"tags" : map[string]string{
				"RedshiftQueryOwner": "*******",
			},
		},
	})

	defer terraform.Destroy(t, terraformOptionsSecretValues)

	terraform.InitAndApply(t, terraformOptionsSecretValues)

	output := terraform.Output(t, terraformOptionsSecretValues, "test")
	
	assert.Equal(t, "redshift49", output)
}
