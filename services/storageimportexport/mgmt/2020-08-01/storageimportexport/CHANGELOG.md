Generated from https://github.com/Azure/azure-rest-api-specs/tree/3c764635e7d442b3e74caf593029fcd440b3ef82

Code generator @microsoft.azure/autorest.go@~2.1.161

## Breaking Changes

- Function `NewListJobsResponsePage` parameter(s) have been changed from `(func(context.Context, ListJobsResponse) (ListJobsResponse, error))` to `(ListJobsResponse, func(context.Context, ListJobsResponse) (ListJobsResponse, error))`
- Type of `JobDetails.DeliveryPackage` has been changed from `*PackageInformation` to `*DeliveryPackageInformation`
- Type of `UpdateJobParametersProperties.DeliveryPackage` has been changed from `*PackageInformation` to `*DeliveryPackageInformation`

## New Content

- New struct `DeliveryPackageInformation`
