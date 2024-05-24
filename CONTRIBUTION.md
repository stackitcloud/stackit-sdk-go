# Contribute to the STACKIT Go SDK

Your contribution is welcome! Thank you for your interest in contributing to the STACKIT Go SDK. We greatly value your feedback, feature requests, additions to the code, bug reports or documentation extensions.

## Table of contents

- [Developer Guide](#developer-guide)
  - [Repository structure](#repository-structure)
  - [Implementing a module waiter](#implementing-a-module-waiter)
    - [Waiter structure](#waiter-structure)
    - [Notes](#notes)
- [Code Contributions](#code-contributions)
- [Bug Reports](#bug-reports)

## Developer Guide

#### Useful Make commands

These commands can be executed from the project root:

- `make project-tools`: get the required dependencies
- `make lint`: lint the code and the examples and sync dependencies. To only lint automatically generated files, run `make lint skip-non-generated-files=true`
- `make test`: run unit tests. To only test automatically generated files, run `make lint skip-non-generated-files=true`

### Repository structure

The SDK STACKIT service modules are located under `services`. The files located in `services/[service]` are automatically generated from the [REST API specs](https://github.com/stackitcloud/stackit-api-specifications), whereas the ones located in subfolders (like `wait`) are manually maintained. Therefore, changes to files located in `services/[service]` will not be accepted. Instead, consider proposing changes to the generation process in the [Generator repository](https://github.com/stackitcloud/stackit-sdk-generator).

Inside `core` you can find several packages that are used by all service modules, such as `auth`, `config` and `wait`. Examples of usage of the SDK are located under the `examples` folder.

### Implementing a module waiter

For integration with other tools such as the [STACKIT Terraform Provider](https://github.com/stackitcloud/terraform-provider-stackit) and the [STACKIT CLI](https://github.com/stackitcloud/stackit-cli), there is the need to implement `waiters` for some SDK modules. Waiters are routines for resources that are created/updated/deleted asynchronously, which wait until the creation/update/deletion of the specified resource is completed. The waiters are located in a folder named `wait` inside each service folder.

Let's suppose you want to implement the waiters for the `Create`, `Update` and `Delete` operations of a resource `bar` of service `foo`:

1. You would start by creating a new folder `wait/` inside `services/foo/`, if it not exists yet
2. Following with the creation of a file `wait.go` inside your new folder `services/foo/wait`, if it not exists yet. The Go package should be named `wait`.
3. Refer to the [Waiter structure](./CONTRIBUTION.md/#waiter-structure) section for details on the structure of the file and the methods
4. Add unit tests to the wait functions

#### Waiter structure

Below is a typical structure of a waiter for an SDK module:

```go
package wait

import (
	"context"
	"fmt"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	"github.com/stackitcloud/stackit-sdk-go/services/foo"
)

const (
	CreateSuccess = "CREATE_SUCCEEDED"
	CreateFail    = "CREATE_FAILED"
	UpdateSuccess = "UPDATE_SUCCEEDED"
	UpdateFail    = "UPDATE_FAILED"
	DeleteSuccess = "DELETE_SUCCEEDED"
	DeleteFail    = "DELETE_FAILED"
)

// APIClientInterface Interfaces needed for tests
type APIClientInterface interface {
	GetBarExecute(ctx context.Context, BarId, projectId string) (*foo.GetBarResponse, error)
}

// CreateBarWaitHandler will wait for Bar creation
func CreateBarWaitHandler(ctx context.Context, a APIClientInterface, BarId, projectId string) *wait.AsyncActionHandler[foo.GetBarResponse] {
	handler := wait.New(func() (waitFinished bool, response *foo.GetBarResponse, err error) {
		s, err := a.GetBarExecute(ctx, BarId, projectId)
		if err != nil {
			return false, nil, err
		}
		if s.Id == nil || s.Status == nil {
			return false, nil, fmt.Errorf("could not get Bar id or status from response for project %s and Bar %s", projectId, BarId)
		}
		if *s.Id == BarId && *s.Status == CreateSuccess {
			return true, s, nil
		}
		if *s.Id == BarId && *s.Status == CreateFail {
			return true, s, fmt.Errorf("create failed for Bar with id %s", BarId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(45 * time.Minute)
	return handler
}

// UpdateBarWaitHandler will wait for Bar update
func UpdateBarWaitHandler(ctx context.Context, a APIClientInterface, BarId, projectId string) *wait.AsyncActionHandler[foo.GetBarResponse] {
	handler := wait.New(func() (waitFinished bool, response *foo.GetBarResponse, err error) {
		s, err := a.GetBarExecute(ctx, BarId, projectId)
		if err != nil {
			return false, nil, err
		}
		if s.Id == nil || s.Status == nil {
			return false, nil, fmt.Errorf("could not get Bar id or status from response for project %s and Bar %s", projectId, BarId)
		}
		if *s.Id == BarId && (*s.Status == UpdateSuccess) {
			return true, s, nil
		}
		if *s.Id == BarId && (*s.Status == UpdateFail) {
			return true, s, fmt.Errorf("update failed for Bar with id %s", BarId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(30 * time.Minute)
	return handler
}

// DeleteBarWaitHandler will wait for Bar deletion
func DeleteBarWaitHandler(ctx context.Context, a APIClientInterface, BarId, projectId string) *wait.AsyncActionHandler[foo.GetBarResponse] {
	handler := wait.New(func() (waitFinished bool, response *foo.GetBarResponse, err error) {
		s, err := a.GetBarExecute(ctx, BarId, projectId)
		if err != nil {
			return false, nil, err
		}
		if s.Id == nil || s.Status == nil {
			return false, nil, fmt.Errorf("could not get Bar id or status from response for project %s and Bar %s", projectId, BarId)
		}
		if *s.Id == BarId && *s.Status == DeleteSuccess {
			return true, s, nil
		}
		if *s.Id == BarId && *s.Status == DeleteFail {
			return true, s, fmt.Errorf("delete failed for Bar with id %s", BarId)
		}
		return false, nil, nil
	})
	handler.SetTimeout(20 * time.Minute)
	return handler
}
```

#### Notes

- The states may vary from service to service
- The `id` and the `state` might not be present on the root level of the API response, this also varies from service to service. You must always match the resource `id` and the resource `state` to what is expected
- The timeout values included above are just for reference, each resource takes different amounts of time to finish the create, update or delete operations
- For some resources, after a successful delete operation the resource can't be found anymore, so a call to the `Get` method would result in an error. In those cases, the waiter can be implemented by calling the `List` method and check that the resource is not present, like in this example:

  ```go
  // DeleteBarWaitHandler will wait for Bar deletion
  func DeleteBarWaitHandler(ctx context.Context, a APIClientInterface, barId, projectId string) *wait.AsyncActionHandler[foo.ListBarsResponse] {
      handler := wait.New(func() (waitFinished bool, response *foo.ListBarsResponse, err error) {
          s, err := a.ListBarsExecute(ctx, barId, projectId)
          if err != nil {
              return false, nil, err
          }
          for i := range s {
              if *s[i].Id == barId {
                  return false, nil, nil
              }
          }
          return true, s, nil
      })
      handler.SetTimeout(10 * time.Minute)
      return handler
  }
  ```

- The main objective of the `wait` functions is to make sure that the operation was successful, which means any other special cases such as intermediate error states should also be handled

## Code Contributions

To make your contribution, follow these steps:

1. Check open or recently closed [Pull Requests](https://github.com/stackitcloud/stackit-sdk-go/pulls) and [Issues](https://github.com/stackitcloud/stackit-sdk-go/issues)to make sure the contribution you are making has not been already tackled by someone else.
2. Fork the repo.
3. Make your changes in a branch that is up-to-date with the original repo's `main` branch.
4. Commit your changes including a descriptive message.
5. Create a pull request with your changes.
6. The pull request will be reviewed by the repo maintainers. If you need to make further changes, make additional commits to keep commit history. When the PR is merged, commits will be squashed.

## Bug Reports

If you would like to report a bug, please open a [GitHub issue](https://github.com/stackitcloud/stackit-sdk-go/issues/new).

To ensure we can provide the best support to your issue, follow these guidelines:

1. Go through the existing issues to check if your issue has already been reported.
2. Make sure you are using the latest version of the provider, we will not provide bug fixes for older versions. Also, latest versions may have the fix for your bug.
3. Please provide as much information as you can about your environment, e.g. your version of Go, your version of the provider, which operating system you are using and the corresponding version.
4. Include in your issue the steps to reproduce it, along with code snippets and/or information about your specific use case. This will make the support process much easier and efficient.
