package wait

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	ufw "github.com/stackitcloud/stackit-sdk-go/services/ufw/v1api"
)

func CreateRuleWaitHandler(ctx context.Context, a ufw.DefaultAPI, projectId, region, ruleId string) *wait.AsyncActionHandler[ufw.RuleResponse] {
	return createOrUpdateRuleWaitHandler(ctx, a, projectId, region, ruleId, []ufw.RuleResponseStatus{ufw.RULERESPONSESTATUS_ACTIVE}, nil)
}

func UpdateRuleWaitHandler(ctx context.Context, a ufw.DefaultAPI, projectId, region, ruleId string) *wait.AsyncActionHandler[ufw.RuleResponse] {
	return createOrUpdateInstanceWaitHandler(ctx, a, projectId, region, ruleId, []ufw.RuleResponseStatus{ufw.RULERESPONSESTATUS_ACTIVE}, nil)
}

func DeleteRuleWaitHandler(ctx context.Context, a ufw.DefaultAPI, projectId, region, ruleId string) *wait.AsyncActionHandler[ufw.RuleResponse] {
	return createOrUpdateInstanceWaitHandler(ctx, a, projectId, region, ruleId, nil, []int{http.StatusNotFound})
}

func createOrUpdateInstanceWaitHandler(ctx context.Context, a ufw.DefaultAPI, projectId, region, ruleId string, activeStates []ufw.RuleResponseStatus, deleteHttpErrorStatusCodes []int) *wait.AsyncActionHandler[ufw.RuleResponse] {
	waitConfig := wait.WaiterHelper[ufw.RuleResponse, ufw.RuleResponseStatus]{
		FetchInstance: a.GetRule(ctx, projectId, region, ruleId).Execute,
		GetState: func(ruleResp *ufw.RuleResponse) (ufw.RuleResponseStatus, error) {
			if ruleResp == nil {
				return "", errors.New("empty response")
			}
			return ruleResp.Status, nil
		},
		ActiveState:                activeStates,
		ErrorState:                 []ufw.RuleResponseStatus{ufw.RULERESPONSESTATUS_ERROR},
		DeleteHttpErrorStatusCodes: deleteHttpErrorStatusCodes,
	}

	handler := wait.New(waitConfig.Wait())
	handler.SetTimeout(5 * time.Minute)
	return handler
}
