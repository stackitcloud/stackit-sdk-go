package v3alpha1api

import (
	"context"
	"net/http"
)

// assert the implementation matches the interface
var _ DefaultAPI = &DefaultAPIServiceMock{}

type DefaultAPIServiceMock struct {
	CloneRequestExecuteMock                   *func(r ApiCloneRequestRequest) (*CloneResponse, *http.Response, error)
	CreateDatabaseRequestExecuteMock          *func(r ApiCreateDatabaseRequestRequest) (*CreateDatabaseResponse, *http.Response, error)
	CreateInstanceRequestExecuteMock          *func(r ApiCreateInstanceRequestRequest) (*CreateInstanceResponse, *http.Response, error)
	CreateUserRequestExecuteMock              *func(r ApiCreateUserRequestRequest) (*CreateUserResponse, *http.Response, error)
	DeleteDatabaseRequestExecuteMock          *func(r ApiDeleteDatabaseRequestRequest) (*http.Response, error)
	DeleteInstanceRequestExecuteMock          *func(r ApiDeleteInstanceRequestRequest) (*http.Response, error)
	DeleteUserRequestExecuteMock              *func(r ApiDeleteUserRequestRequest) (*http.Response, error)
	GetBackupRequestExecuteMock               *func(r ApiGetBackupRequestRequest) (*GetBackupResponse, *http.Response, error)
	GetCollationsRequestExecuteMock           *func(r ApiGetCollationsRequestRequest) (*GetCollationsResponse, *http.Response, error)
	GetDatabaseRequestExecuteMock             *func(r ApiGetDatabaseRequestRequest) (*GetDatabaseResponse, *http.Response, error)
	GetFlavorsRequestExecuteMock              *func(r ApiGetFlavorsRequestRequest) (*GetFlavorsResponse, *http.Response, error)
	GetInstanceRequestExecuteMock             *func(r ApiGetInstanceRequestRequest) (*GetInstanceResponse, *http.Response, error)
	GetUserRequestExecuteMock                 *func(r ApiGetUserRequestRequest) (*GetUserResponse, *http.Response, error)
	GetVersionsRequestExecuteMock             *func(r ApiGetVersionsRequestRequest) (*GetVersionsResponse, *http.Response, error)
	ListBackupsRequestExecuteMock             *func(r ApiListBackupsRequestRequest) (*ListBackupResponse, *http.Response, error)
	ListDatabasesRequestExecuteMock           *func(r ApiListDatabasesRequestRequest) (*ListDatabasesResponse, *http.Response, error)
	ListInstancesRequestExecuteMock           *func(r ApiListInstancesRequestRequest) (*ListInstancesResponse, *http.Response, error)
	ListRolesRequestExecuteMock               *func(r ApiListRolesRequestRequest) (*ListRolesResponse, *http.Response, error)
	ListUsersRequestExecuteMock               *func(r ApiListUsersRequestRequest) (*ListUserResponse, *http.Response, error)
	ProtectInstanceRequestExecuteMock         *func(r ApiProtectInstanceRequestRequest) (*ProtectInstanceResponse, *http.Response, error)
	ResetUserRequestExecuteMock               *func(r ApiResetUserRequestRequest) (*ResetUserResponse, *http.Response, error)
	UpdateDatabasePartiallyRequestExecuteMock *func(r ApiUpdateDatabasePartiallyRequestRequest) (*http.Response, error)
	UpdateDatabaseRequestExecuteMock          *func(r ApiUpdateDatabaseRequestRequest) (*http.Response, error)
	UpdateInstancePartiallyRequestExecuteMock *func(r ApiUpdateInstancePartiallyRequestRequest) (*http.Response, error)
	UpdateInstanceRequestExecuteMock          *func(r ApiUpdateInstanceRequestRequest) (*http.Response, error)
	UpdateUserPartiallyRequestExecuteMock     *func(r ApiUpdateUserPartiallyRequestRequest) (*http.Response, error)
	UpdateUserRequestExecuteMock              *func(r ApiUpdateUserRequestRequest) (*http.Response, error)
}

func (a *DefaultAPIServiceMock) CloneRequest(ctx context.Context, projectId string, region string, instanceId string) ApiCloneRequestRequest {
	return ApiCloneRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) CloneRequestExecute(r ApiCloneRequestRequest) (*CloneResponse, *http.Response, error) {
	if a.CloneRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.CloneRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) CreateDatabaseRequest(ctx context.Context, projectId string, region string, instanceId string) ApiCreateDatabaseRequestRequest {
	return ApiCreateDatabaseRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) CreateDatabaseRequestExecute(r ApiCreateDatabaseRequestRequest) (*CreateDatabaseResponse, *http.Response, error) {
	if a.CreateDatabaseRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.CreateDatabaseRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) CreateInstanceRequest(ctx context.Context, projectId string, region string) ApiCreateInstanceRequestRequest {
	return ApiCreateInstanceRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
	}
}

func (a *DefaultAPIServiceMock) CreateInstanceRequestExecute(r ApiCreateInstanceRequestRequest) (*CreateInstanceResponse, *http.Response, error) {
	if a.CreateInstanceRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.CreateInstanceRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) CreateUserRequest(ctx context.Context, projectId string, region string, instanceId string) ApiCreateUserRequestRequest {
	return ApiCreateUserRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) CreateUserRequestExecute(r ApiCreateUserRequestRequest) (*CreateUserResponse, *http.Response, error) {
	if a.CreateUserRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.CreateUserRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) DeleteDatabaseRequest(ctx context.Context, projectId string, region string, instanceId string, databaseId int32) ApiDeleteDatabaseRequestRequest {
	return ApiDeleteDatabaseRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
		databaseId: databaseId,
	}
}

func (a *DefaultAPIServiceMock) DeleteDatabaseRequestExecute(r ApiDeleteDatabaseRequestRequest) (*http.Response, error) {
	if a.DeleteDatabaseRequestExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.DeleteDatabaseRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) DeleteInstanceRequest(ctx context.Context, projectId string, region string, instanceId string) ApiDeleteInstanceRequestRequest {
	return ApiDeleteInstanceRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) DeleteInstanceRequestExecute(r ApiDeleteInstanceRequestRequest) (*http.Response, error) {
	if a.DeleteInstanceRequestExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.DeleteInstanceRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) DeleteUserRequest(ctx context.Context, projectId string, region string, instanceId string, userId int32) ApiDeleteUserRequestRequest {
	return ApiDeleteUserRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
		userId:     userId,
	}
}

func (a *DefaultAPIServiceMock) DeleteUserRequestExecute(r ApiDeleteUserRequestRequest) (*http.Response, error) {
	if a.DeleteUserRequestExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.DeleteUserRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) GetBackupRequest(ctx context.Context, projectId string, region string, instanceId string, backupId int32) ApiGetBackupRequestRequest {
	return ApiGetBackupRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
		backupId:   backupId,
	}
}

func (a *DefaultAPIServiceMock) GetBackupRequestExecute(r ApiGetBackupRequestRequest) (*GetBackupResponse, *http.Response, error) {
	if a.GetBackupRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.GetBackupRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) GetCollationsRequest(ctx context.Context, projectId string, region string, instanceId string) ApiGetCollationsRequestRequest {
	return ApiGetCollationsRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) GetCollationsRequestExecute(r ApiGetCollationsRequestRequest) (*GetCollationsResponse, *http.Response, error) {
	if a.GetCollationsRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.GetCollationsRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) GetDatabaseRequest(ctx context.Context, projectId string, region string, instanceId string, databaseId int32) ApiGetDatabaseRequestRequest {
	return ApiGetDatabaseRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
		databaseId: databaseId,
	}
}

func (a *DefaultAPIServiceMock) GetDatabaseRequestExecute(r ApiGetDatabaseRequestRequest) (*GetDatabaseResponse, *http.Response, error) {
	if a.GetDatabaseRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.GetDatabaseRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) GetFlavorsRequest(ctx context.Context, projectId string, region string) ApiGetFlavorsRequestRequest {
	return ApiGetFlavorsRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
	}
}

func (a *DefaultAPIServiceMock) GetFlavorsRequestExecute(r ApiGetFlavorsRequestRequest) (*GetFlavorsResponse, *http.Response, error) {
	if a.GetFlavorsRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.GetFlavorsRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) GetInstanceRequest(ctx context.Context, projectId string, region string, instanceId string) ApiGetInstanceRequestRequest {
	return ApiGetInstanceRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) GetInstanceRequestExecute(r ApiGetInstanceRequestRequest) (*GetInstanceResponse, *http.Response, error) {
	if a.GetInstanceRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.GetInstanceRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) GetUserRequest(ctx context.Context, projectId string, region string, instanceId string, userId int32) ApiGetUserRequestRequest {
	return ApiGetUserRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
		userId:     userId,
	}
}

func (a *DefaultAPIServiceMock) GetUserRequestExecute(r ApiGetUserRequestRequest) (*GetUserResponse, *http.Response, error) {
	if a.GetUserRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.GetUserRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) GetVersionsRequest(ctx context.Context, projectId string, region string) ApiGetVersionsRequestRequest {
	return ApiGetVersionsRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
	}
}

func (a *DefaultAPIServiceMock) GetVersionsRequestExecute(r ApiGetVersionsRequestRequest) (*GetVersionsResponse, *http.Response, error) {
	if a.GetVersionsRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.GetVersionsRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListBackupsRequest(ctx context.Context, projectId string, region string, instanceId string) ApiListBackupsRequestRequest {
	return ApiListBackupsRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) ListBackupsRequestExecute(r ApiListBackupsRequestRequest) (*ListBackupResponse, *http.Response, error) {
	if a.ListBackupsRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListBackupsRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListDatabasesRequest(ctx context.Context, projectId string, region string, instanceId string) ApiListDatabasesRequestRequest {
	return ApiListDatabasesRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) ListDatabasesRequestExecute(r ApiListDatabasesRequestRequest) (*ListDatabasesResponse, *http.Response, error) {
	if a.ListDatabasesRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListDatabasesRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListInstancesRequest(ctx context.Context, projectId string, region string) ApiListInstancesRequestRequest {
	return ApiListInstancesRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
	}
}

func (a *DefaultAPIServiceMock) ListInstancesRequestExecute(r ApiListInstancesRequestRequest) (*ListInstancesResponse, *http.Response, error) {
	if a.ListInstancesRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListInstancesRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListRolesRequest(ctx context.Context, projectId string, region string, instanceId string) ApiListRolesRequestRequest {
	return ApiListRolesRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) ListRolesRequestExecute(r ApiListRolesRequestRequest) (*ListRolesResponse, *http.Response, error) {
	if a.ListRolesRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListRolesRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListUsersRequest(ctx context.Context, projectId string, region string, instanceId string) ApiListUsersRequestRequest {
	return ApiListUsersRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) ListUsersRequestExecute(r ApiListUsersRequestRequest) (*ListUserResponse, *http.Response, error) {
	if a.ListUsersRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListUsersRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ProtectInstanceRequest(ctx context.Context, projectId string, region string, instanceId string) ApiProtectInstanceRequestRequest {
	return ApiProtectInstanceRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) ProtectInstanceRequestExecute(r ApiProtectInstanceRequestRequest) (*ProtectInstanceResponse, *http.Response, error) {
	if a.ProtectInstanceRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ProtectInstanceRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ResetUserRequest(ctx context.Context, projectId string, region string, instanceId string, userId int32) ApiResetUserRequestRequest {
	return ApiResetUserRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
		userId:     userId,
	}
}

func (a *DefaultAPIServiceMock) ResetUserRequestExecute(r ApiResetUserRequestRequest) (*ResetUserResponse, *http.Response, error) {
	if a.ResetUserRequestExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ResetUserRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) UpdateDatabasePartiallyRequest(ctx context.Context, projectId string, region string, instanceId string, databaseId int32) ApiUpdateDatabasePartiallyRequestRequest {
	return ApiUpdateDatabasePartiallyRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
		databaseId: databaseId,
	}
}

func (a *DefaultAPIServiceMock) UpdateDatabasePartiallyRequestExecute(r ApiUpdateDatabasePartiallyRequestRequest) (*http.Response, error) {
	if a.UpdateDatabasePartiallyRequestExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.UpdateDatabasePartiallyRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) UpdateDatabaseRequest(ctx context.Context, projectId string, region string, instanceId string, databaseId int32) ApiUpdateDatabaseRequestRequest {
	return ApiUpdateDatabaseRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
		databaseId: databaseId,
	}
}

func (a *DefaultAPIServiceMock) UpdateDatabaseRequestExecute(r ApiUpdateDatabaseRequestRequest) (*http.Response, error) {
	if a.UpdateDatabaseRequestExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.UpdateDatabaseRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) UpdateInstancePartiallyRequest(ctx context.Context, projectId string, region string, instanceId string) ApiUpdateInstancePartiallyRequestRequest {
	return ApiUpdateInstancePartiallyRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) UpdateInstancePartiallyRequestExecute(r ApiUpdateInstancePartiallyRequestRequest) (*http.Response, error) {
	if a.UpdateInstancePartiallyRequestExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.UpdateInstancePartiallyRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) UpdateInstanceRequest(ctx context.Context, projectId string, region string, instanceId string) ApiUpdateInstanceRequestRequest {
	return ApiUpdateInstanceRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) UpdateInstanceRequestExecute(r ApiUpdateInstanceRequestRequest) (*http.Response, error) {
	if a.UpdateInstanceRequestExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.UpdateInstanceRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) UpdateUserPartiallyRequest(ctx context.Context, projectId string, region string, instanceId string, userId int32) ApiUpdateUserPartiallyRequestRequest {
	return ApiUpdateUserPartiallyRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
		userId:     userId,
	}
}

func (a *DefaultAPIServiceMock) UpdateUserPartiallyRequestExecute(r ApiUpdateUserPartiallyRequestRequest) (*http.Response, error) {
	if a.UpdateUserPartiallyRequestExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.UpdateUserPartiallyRequestExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) UpdateUserRequest(ctx context.Context, projectId string, region string, instanceId string, userId int32) ApiUpdateUserRequestRequest {
	return ApiUpdateUserRequestRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		region:     region,
		instanceId: instanceId,
		userId:     userId,
	}
}

func (a *DefaultAPIServiceMock) UpdateUserRequestExecute(r ApiUpdateUserRequestRequest) (*http.Response, error) {
	if a.UpdateUserRequestExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.UpdateUserRequestExecuteMock)(r)
}
