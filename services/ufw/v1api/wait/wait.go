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
	RuleStatusActive RuleStatus = "Active"
	RuleStatusError  RuleStatus = "Error"
)

func CreateRuleWaitHandler(ctx context.Context, a ufw.DefaultAPI, projectId, region, ruleId string) *wait.AsyncActionHandler[ufw.RuleResponse] {
	return ruleWaitHandler(ctx, a, projectId, region, ruleId, []RuleStatus{RuleStatusActive}, nil)
}

func UpdateRuleWaitHandler(ctx context.Context, a ufw.DefaultAPI, projectId, region, ruleId string) *wait.AsyncActionHandler[ufw.RuleResponse] {
	return ruleWaitHandler(ctx, a, projectId, region, ruleId, []RuleStatus{RuleStatusActive}, nil)
}

func DeleteRuleWaitHandler(ctx context.Context, a ufw.DefaultAPI, projectId, region, ruleId string) *wait.AsyncActionHandler[ufw.RuleResponse] {
	return ruleWaitHandler(ctx, a, projectId, region, ruleId, nil, []int{http.StatusNotFound})
}

func ruleWaitHandler(ctx context.Context, a ufw.DefaultAPI, projectId, region, ruleId string, activeStates []RuleStatus, deleteHttpErrorStatusCodes []int) *wait.AsyncActionHandler[ufw.RuleResponse] {
	waitConfig := wait.WaiterHelper[ufw.RuleResponse, RuleStatus]{
		FetchInstance: a.GetRule(ctx, projectId, region, ruleId).Execute,
		GetState: func(ruleResp *ufw.RuleResponse) (RuleStatus, error) {
			if ruleResp == nil {
				return "", errors.New("empty response")
			}
			return RuleStatus(*ruleResp.Status), nil
		},
		ActiveState:                activeStates,
		ErrorState:                 []RuleStatus{RuleStatusError},
		DeleteHttpErrorStatusCodes: deleteHttpErrorStatusCodes,
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}
