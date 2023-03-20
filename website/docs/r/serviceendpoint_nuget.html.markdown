---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_serviceendpoint_nuget"
description: |-
  Manages a NuGet server endpoint within Azure DevOps organization.
---

# azuredevops_serviceendpoint_nuget

Manages a NuGet service endpoint within Azure DevOps.

## Example Usage

```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
  description        = "Managed by Terraform"
}

resource "azuredevops_serviceendpoint_nuget" "example" {
  project_id            = azuredevops_project.example.id
  service_endpoint_name = "Example NuGet"
  url                   = "https://api.nuget.org/v3/index.json"
  access_token          = "AbcDEf123_0x"
  description           = "Managed by Terraform"
}
```

## Argument Reference

The following arguments are supported:

- `project_id` - (Required) The ID of the project.
- `service_endpoint_name` - (Required) The Service Endpoint name.
- `url` - (Required) URL of the NuGet feed to connect with.
- `access_token` - (Required) The access-token/ApiKey for NuGet feed.
- `description` - (Optional) The Service Endpoint description.

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service endpoint.
- `project_id` - The ID of the project.
- `service_endpoint_name` - The Service Endpoint name.

## Relevant Links

- [Azure DevOps Service REST API 6.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0)
- [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml)
- [NuGet ApiKey](https://learn.microsoft.com/en-in/nuget/nuget-org/scoped-api-keys)
- [NuGet packages in Azure Artifacts](https://learn.microsoft.com/en-us/azure/devops/artifacts/get-started-nuget?view=azure-devops&tabs=windows)

## Import

Azure DevOps Service Endpoint NuGet can be imported using the **projectID/serviceEndpointID**, e.g.

```sh
terraform import azuredevops_serviceendpoint_nuget.example 00000000-0000-0000-0000-000000000000/00000000-0000-0000-0000-000000000000
```
