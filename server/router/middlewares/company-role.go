package middlewares

import (
	"contract/common"
	"contract/dao/company"
	"contract/model"
	"contract/router/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetCompanyRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := common.GetSession(c)

		companyIdInHeader := c.GetHeader("companyId")
		if companyIdInHeader == "" {
			response.FailWithCode(c, "无法获取到公司信息", response.AuthErrorCode)
			c.Abort()
			return
		}

		companyId, err := strconv.ParseUint(companyIdInHeader, 10, 64)
		if err != nil {
			response.FailWithCode(c, "公司信息不正确", response.AuthErrorCode)
			c.Abort()
			return
		}

		role, err := companyDao.FindCompanyMember(&model.CompanyMember{
			UserId:    session.UserId,
			CompanyId: uint(companyId),
		}, "company_member.id, company_member.role")
		if err != nil || role.ID == 0 {
			response.FailWithCode(c, "非该公司成员", response.AuthErrorCode)
			c.Abort()
			return
		}
		if len(roles) > 0 {
			if !common.Contains(role.Role, roles) {
				response.FailWithCode(c, "没有操作权限", response.AuthErrorCode)
				c.Abort()
				return
			}
		}
		c.Set("companyRole", role.Role)
	}
}
