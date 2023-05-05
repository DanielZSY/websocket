package conngateway

import (
	"github.com/zeromicro/go-zero/core/logx"
	"imservice/app/conn/internal/types"
	"imservice/common/utils"
)

func ReqLog[REQ IReq, RESP IResp](c *types.UserConn, method string, body IBody, req REQ, resp RESP, err error) {
	reqStr := utils.AnyToString(req)
	respStr := utils.AnyToString(resp)
	reqId := body.GetReqId()
	if err != nil {
		logx.WithContext(c.Ctx).Errorf("reqId: %s, method: %s, req: %s, resp: %s, error: %v", reqId, method, reqStr, respStr, err)
	} else {
		logx.WithContext(c.Ctx).Debugf("reqId: %s, method: %s, req: %s, resp: %s", reqId, method, reqStr, respStr)
	}
}