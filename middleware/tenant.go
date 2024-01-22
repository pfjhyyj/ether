package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/tenant/domain"
	"github.com/pfjhyyj/ether/common"
	"github.com/sirupsen/logrus"
	"net/http"
)

func TenantMiddleware(strict bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		logs := logrus.WithContext(c)
		host := c.Request.Host

		tenantRepo := domain.GetTenantRepository()
		tenant, err := tenantRepo.GetTenantByDomain(c, c.Request.Host)
		if err != nil || tenant == nil {
			logs.WithError(err).Warnf("fail to get tenant by domain %s", host)
			if strict {
				c.AbortWithStatusJSON(http.StatusOK, &common.Response{
					Code: common.RequestError,
					Msg:  "Unknown domain",
				})
				return
			}
		} else {
			c.Set(common.CtxTenantIDKey, tenant.TenantId)
		}

		c.Next()
	}
}
