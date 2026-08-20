package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	ufw "github.com/stackitcloud/stackit-sdk-go/services/ufw/v1api"
)

type RuleStatus string

const (
	RuleStatusActive   RuleStatus = "Active"
	RuleStatusCreating RuleStatus = "Creating"
	RuleStatusUpdating RuleStatus = "Updating"
	RuleStatusPending  RuleStatus = "Pending"
	RuleStatusDeleting RuleStatus = "Deleting"
	RuleStatusError    RuleStatus = "Error"
)

func CreateRuleWaitHandler(ctx context.Context, a ufw.DefaultAPI, projectId, region, ruleId string) *wait.AsyncActionHandler[ufw.RuleResponse] {
	waitConfig := wait.WaiterHelper[ufw.RuleResponse, RuleStatus]{
		FetchInstance: a.GetRule(ctx, projectId, region, ruleId).Execute,
		GetState: func(ruleResp *ufw.RuleResponse) (RuleStatus, error) {
			if ruleResp == nil {
				return "", errors.New("empty response")
			}
			if ruleResp.Status == nil {
				return "", errors.New("status is missing")
			}
			return RuleStatus(*ruleResp.Status), nil
		},
		ActiveState: []RuleStatus{RuleStatusActive},
		ErrorState:  []RuleStatus{RuleStatusError},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}

func UpdateRuleWaitHandler(ctx context.Context, a ufw.DefaultAPI, projectId, region, ruleId string) *wait.AsyncActionHandler[ufw.RuleResponse] {
	waitConfig := wait.WaiterHelper[ufw.RuleResponse, RuleStatus]{
		FetchInstance: a.GetRule(ctx, projectId, region, ruleId).Execute,
		GetState: func(ruleResp *ufw.RuleResponse) (RuleStatus, error) {
			if ruleResp == nil {
				return "", errors.New("empty response")
			}
			if ruleResp.Status == nil {
				return "", errors.New("status is missing")
			}
			return RuleStatus(*ruleResp.Status), nil
		},
		ActiveState: []RuleStatus{RuleStatusActive},
		ErrorState:  []RuleStatus{RuleStatusError},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}

func DeleteRuleWaitHandler(ctx context.Context, a ufw.DefaultAPI, projectId, region, ruleId string) *wait.AsyncActionHandler[ufw.RuleResponse] {
	waitConfig := wait.WaiterHelper[ufw.RuleResponse, RuleStatus]{
		FetchInstance: a.GetRule(ctx, projectId, region, ruleId).Execute,
		GetState: func(ruleResp *ufw.RuleResponse) (RuleStatus, error) {
			if ruleResp == nil {
				return "", errors.New("empty response")
			}
			if ruleResp.Status == nil {
				return "", errors.New("status is missing")
			}
			return RuleStatus(*ruleResp.Status), nil
		},
		ErrorState:                 []RuleStatus{RuleStatusError},
		DeleteHttpErrorStatusCodes: []int{http.StatusNotFound},
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}
