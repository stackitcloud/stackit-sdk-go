package v1api

import (
	"context"
	"net/http"
)

// assert the implementation matches the interface
var _ DefaultAPI = &DefaultAPIServiceMock{}

type DefaultAPIServiceMock struct {
	CloneInstanceExecuteMock          *func(r ApiCloneInstanceRequest) (*CloneInstanceResponse, *http.Response, error)
	CreateDatabaseExecuteMock         *func(r ApiCreateDatabaseRequest) (*InstanceCreateDatabaseResponse, *http.Response, error)
	CreateInstanceExecuteMock         *func(r ApiCreateInstanceRequest) (*CreateInstanceResponse, *http.Response, error)
	CreateUserExecuteMock             *func(r ApiCreateUserRequest) (*CreateUserResponse, *http.Response, error)
	DeleteDatabaseExecuteMock         *func(r ApiDeleteDatabaseRequest) (*http.Response, error)
	DeleteInstanceExecuteMock         *func(r ApiDeleteInstanceRequest) (*http.Response, error)
	DeleteUserExecuteMock             *func(r ApiDeleteUserRequest) (*http.Response, error)
	ForceDeleteInstanceExecuteMock    *func(r ApiForceDeleteInstanceRequest) (*http.Response, error)
	GetBackupExecuteMock              *func(r ApiGetBackupRequest) (*GetBackupResponse, *http.Response, error)
	GetInstanceExecuteMock            *func(r ApiGetInstanceRequest) (*InstanceResponse, *http.Response, error)
	GetUserExecuteMock                *func(r ApiGetUserRequest) (*GetUserResponse, *http.Response, error)
	ListBackupsExecuteMock            *func(r ApiListBackupsRequest) (*ListBackupsResponse, *http.Response, error)
	ListDatabaseParametersExecuteMock *func(r ApiListDatabaseParametersRequest) (*PostgresDatabaseParameterResponse, *http.Response, error)
	ListDatabasesExecuteMock          *func(r ApiListDatabasesRequest) (*InstanceListDatabasesResponse, *http.Response, error)
	ListFlavorsExecuteMock            *func(r ApiListFlavorsRequest) (*ListFlavorsResponse, *http.Response, error)
	ListInstancesExecuteMock          *func(r ApiListInstancesRequest) (*ListInstancesResponse, *http.Response, error)
	ListMetricsExecuteMock            *func(r ApiListMetricsRequest) (*InstanceMetricsResponse, *http.Response, error)
	ListStoragesExecuteMock           *func(r ApiListStoragesRequest) (*ListStoragesResponse, *http.Response, error)
	ListUsersExecuteMock              *func(r ApiListUsersRequest) (*ListUsersResponse, *http.Response, error)
	ListVersionsExecuteMock           *func(r ApiListVersionsRequest) (*ListVersionsResponse, *http.Response, error)
	PartialUpdateInstanceExecuteMock  *func(r ApiPartialUpdateInstanceRequest) (*PartialUpdateInstanceResponse, *http.Response, error)
	PartialUpdateUserExecuteMock      *func(r ApiPartialUpdateUserRequest) (*http.Response, error)
	ResetUserExecuteMock              *func(r ApiResetUserRequest) (*ResetUserResponse, *http.Response, error)
	UpdateBackupScheduleExecuteMock   *func(r ApiUpdateBackupScheduleRequest) (*http.Response, error)
	UpdateInstanceExecuteMock         *func(r ApiUpdateInstanceRequest) (*PartialUpdateInstanceResponse, *http.Response, error)
	UpdateUserExecuteMock             *func(r ApiUpdateUserRequest) (*http.Response, error)
}

func (a *DefaultAPIServiceMock) CloneInstance(ctx context.Context, projectId string, instanceId string) ApiCloneInstanceRequest {
	return ApiCloneInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) CloneInstanceExecute(r ApiCloneInstanceRequest) (*CloneInstanceResponse, *http.Response, error) {
	if a.CloneInstanceExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.CloneInstanceExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) CreateDatabase(ctx context.Context, projectId string, instanceId string) ApiCreateDatabaseRequest {
	return ApiCreateDatabaseRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) CreateDatabaseExecute(r ApiCreateDatabaseRequest) (*InstanceCreateDatabaseResponse, *http.Response, error) {
	if a.CreateDatabaseExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.CreateDatabaseExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) CreateInstance(ctx context.Context, projectId string) ApiCreateInstanceRequest {
	return ApiCreateInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
	}
}

func (a *DefaultAPIServiceMock) CreateInstanceExecute(r ApiCreateInstanceRequest) (*CreateInstanceResponse, *http.Response, error) {
	if a.CreateInstanceExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.CreateInstanceExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) CreateUser(ctx context.Context, projectId string, instanceId string) ApiCreateUserRequest {
	return ApiCreateUserRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) CreateUserExecute(r ApiCreateUserRequest) (*CreateUserResponse, *http.Response, error) {
	if a.CreateUserExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.CreateUserExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) DeleteDatabase(ctx context.Context, projectId string, instanceId string, databaseId string) ApiDeleteDatabaseRequest {
	return ApiDeleteDatabaseRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
		databaseId: databaseId,
	}
}

func (a *DefaultAPIServiceMock) DeleteDatabaseExecute(r ApiDeleteDatabaseRequest) (*http.Response, error) {
	if a.DeleteDatabaseExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.DeleteDatabaseExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) DeleteInstance(ctx context.Context, projectId string, instanceId string) ApiDeleteInstanceRequest {
	return ApiDeleteInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) DeleteInstanceExecute(r ApiDeleteInstanceRequest) (*http.Response, error) {
	if a.DeleteInstanceExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.DeleteInstanceExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) DeleteUser(ctx context.Context, projectId string, instanceId string, userId string) ApiDeleteUserRequest {
	return ApiDeleteUserRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
		userId:     userId,
	}
}

func (a *DefaultAPIServiceMock) DeleteUserExecute(r ApiDeleteUserRequest) (*http.Response, error) {
	if a.DeleteUserExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.DeleteUserExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ForceDeleteInstance(ctx context.Context, projectId string, instanceId string) ApiForceDeleteInstanceRequest {
	return ApiForceDeleteInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) ForceDeleteInstanceExecute(r ApiForceDeleteInstanceRequest) (*http.Response, error) {
	if a.ForceDeleteInstanceExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.ForceDeleteInstanceExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) GetBackup(ctx context.Context, projectId string, instanceId string, backupId string) ApiGetBackupRequest {
	return ApiGetBackupRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
		backupId:   backupId,
	}
}

func (a *DefaultAPIServiceMock) GetBackupExecute(r ApiGetBackupRequest) (*GetBackupResponse, *http.Response, error) {
	if a.GetBackupExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.GetBackupExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) GetInstance(ctx context.Context, projectId string, instanceId string) ApiGetInstanceRequest {
	return ApiGetInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) GetInstanceExecute(r ApiGetInstanceRequest) (*InstanceResponse, *http.Response, error) {
	if a.GetInstanceExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.GetInstanceExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) GetUser(ctx context.Context, projectId string, instanceId string, userId string) ApiGetUserRequest {
	return ApiGetUserRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
		userId:     userId,
	}
}

func (a *DefaultAPIServiceMock) GetUserExecute(r ApiGetUserRequest) (*GetUserResponse, *http.Response, error) {
	if a.GetUserExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.GetUserExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListBackups(ctx context.Context, projectId string, instanceId string) ApiListBackupsRequest {
	return ApiListBackupsRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) ListBackupsExecute(r ApiListBackupsRequest) (*ListBackupsResponse, *http.Response, error) {
	if a.ListBackupsExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListBackupsExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListDatabaseParameters(ctx context.Context, projectId string, instanceId string) ApiListDatabaseParametersRequest {
	return ApiListDatabaseParametersRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) ListDatabaseParametersExecute(r ApiListDatabaseParametersRequest) (*PostgresDatabaseParameterResponse, *http.Response, error) {
	if a.ListDatabaseParametersExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListDatabaseParametersExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListDatabases(ctx context.Context, projectId string, instanceId string) ApiListDatabasesRequest {
	return ApiListDatabasesRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) ListDatabasesExecute(r ApiListDatabasesRequest) (*InstanceListDatabasesResponse, *http.Response, error) {
	if a.ListDatabasesExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListDatabasesExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListFlavors(ctx context.Context, projectId string) ApiListFlavorsRequest {
	return ApiListFlavorsRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
	}
}

func (a *DefaultAPIServiceMock) ListFlavorsExecute(r ApiListFlavorsRequest) (*ListFlavorsResponse, *http.Response, error) {
	if a.ListFlavorsExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListFlavorsExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListInstances(ctx context.Context, projectId string) ApiListInstancesRequest {
	return ApiListInstancesRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
	}
}

func (a *DefaultAPIServiceMock) ListInstancesExecute(r ApiListInstancesRequest) (*ListInstancesResponse, *http.Response, error) {
	if a.ListInstancesExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListInstancesExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListMetrics(ctx context.Context, projectId string, instanceId string, metric string) ApiListMetricsRequest {
	return ApiListMetricsRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
		metric:     metric,
	}
}

func (a *DefaultAPIServiceMock) ListMetricsExecute(r ApiListMetricsRequest) (*InstanceMetricsResponse, *http.Response, error) {
	if a.ListMetricsExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListMetricsExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListStorages(ctx context.Context, projectId string, flavorId string) ApiListStoragesRequest {
	return ApiListStoragesRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		flavorId:   flavorId,
	}
}

func (a *DefaultAPIServiceMock) ListStoragesExecute(r ApiListStoragesRequest) (*ListStoragesResponse, *http.Response, error) {
	if a.ListStoragesExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListStoragesExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListUsers(ctx context.Context, projectId string, instanceId string) ApiListUsersRequest {
	return ApiListUsersRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) ListUsersExecute(r ApiListUsersRequest) (*ListUsersResponse, *http.Response, error) {
	if a.ListUsersExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListUsersExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ListVersions(ctx context.Context, projectId string) ApiListVersionsRequest {
	return ApiListVersionsRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
	}
}

func (a *DefaultAPIServiceMock) ListVersionsExecute(r ApiListVersionsRequest) (*ListVersionsResponse, *http.Response, error) {
	if a.ListVersionsExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ListVersionsExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) PartialUpdateInstance(ctx context.Context, projectId string, instanceId string) ApiPartialUpdateInstanceRequest {
	return ApiPartialUpdateInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) PartialUpdateInstanceExecute(r ApiPartialUpdateInstanceRequest) (*PartialUpdateInstanceResponse, *http.Response, error) {
	if a.PartialUpdateInstanceExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.PartialUpdateInstanceExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) PartialUpdateUser(ctx context.Context, projectId string, instanceId string, userId string) ApiPartialUpdateUserRequest {
	return ApiPartialUpdateUserRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
		userId:     userId,
	}
}

func (a *DefaultAPIServiceMock) PartialUpdateUserExecute(r ApiPartialUpdateUserRequest) (*http.Response, error) {
	if a.PartialUpdateUserExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.PartialUpdateUserExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) ResetUser(ctx context.Context, projectId string, instanceId string, userId string) ApiResetUserRequest {
	return ApiResetUserRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
		userId:     userId,
	}
}

func (a *DefaultAPIServiceMock) ResetUserExecute(r ApiResetUserRequest) (*ResetUserResponse, *http.Response, error) {
	if a.ResetUserExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.ResetUserExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) UpdateBackupSchedule(ctx context.Context, projectId string, instanceId string) ApiUpdateBackupScheduleRequest {
	return ApiUpdateBackupScheduleRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) UpdateBackupScheduleExecute(r ApiUpdateBackupScheduleRequest) (*http.Response, error) {
	if a.UpdateBackupScheduleExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.UpdateBackupScheduleExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) UpdateInstance(ctx context.Context, projectId string, instanceId string) ApiUpdateInstanceRequest {
	return ApiUpdateInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
	}
}

func (a *DefaultAPIServiceMock) UpdateInstanceExecute(r ApiUpdateInstanceRequest) (*PartialUpdateInstanceResponse, *http.Response, error) {
	if a.UpdateInstanceExecuteMock == nil {
		return nil, &http.Response{}, nil
	}

	return (*a.UpdateInstanceExecuteMock)(r)
}

func (a *DefaultAPIServiceMock) UpdateUser(ctx context.Context, projectId string, instanceId string, userId string) ApiUpdateUserRequest {
	return ApiUpdateUserRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		instanceId: instanceId,
		userId:     userId,
	}
}

func (a *DefaultAPIServiceMock) UpdateUserExecute(r ApiUpdateUserRequest) (*http.Response, error) {
	if a.UpdateUserExecuteMock == nil {
		return &http.Response{}, nil
	}

	return (*a.UpdateUserExecuteMock)(r)
}
