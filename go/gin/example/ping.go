package main

import (
	"hash/fnv"
	"math/rand"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

func init() {

}

func getLogger(c *gin.Context) *logrus.Entry {
	if log, exist := c.Get("log"); exist {
		if logger, ok := log.(*logrus.Entry); ok {
			return logger
		}
	}
	return nil
}

func getTraceId() string {
	h := fnv.New32a()
	return string(h.Sum([]byte(strconv.FormatInt(rand.Int63(), 10))))
}

func main() {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		uni := ut.New(zh.New(), en.New(), zh_Hant_TW.New())
		trans, _ := uni.GetTranslator(c.GetHeader("locale"))
		if v, ok := binding.Validator.Engine().(*val.Validate); ok {
			zh_trans.RegisterDefaultTranslations(v, trans)
			v.RegisterValidation("prefix", func(fl val.FieldLevel) bool {
				v := fl.Field().String()
				if strings.Index(v, fl.Param()) != 0 {
					return false
				}
				return true
			})
			v.RegisterTranslation("prefix", trans, func(ut ut.Translator) error {
				ut.Add("prefix", "{0} 前綴必須是 {1}", false)
				return nil
			}, func(ut ut.Translator, fe val.FieldError) string {
				prefix := fe.Param()

				if !ok {
					return "prefix 類型必須是string"
				}
				t, err := ut.T("prefix", fe.Field(), prefix)
				if err != nil {
					return err.Error()
				}
				return t
			})

			// override
			trans.Add("min-string", "最少{1}個字", true)
		}
		c.Set("translator", trans)
	})
	router.Use(func(c *gin.Context) {
		c.Set("log", logrus.New().WithFields(logrus.Fields{
			"path":     c.Request.URL.Path,
			"trace-id": getTraceId(),
		}))
		logrus.Info("add logger")
		c.Next()
	})
	router.GET("ping", func(c *gin.Context) {
		count, err := strconv.Atoi(c.Query("count"))
		log := getLogger(c)
		if log != nil {
			log.Info("count", count)
		}
		if err != nil {
			log.Error(err)
		}
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	V1Route(router)
	ValidationRoute(router)

	router.Run(":80")
}

func V1Route(router *gin.Engine) {
	group := router.Group("v1")
	{
		group.Use(func(c *gin.Context) {
			getLogger(c).Info("the middleware only in group v1")
			c.Next()
		})
		group.GET("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}

func ValidationRoute(router *gin.Engine) {
	group := router.Group("valid")
	{
		group.POST("tag", func(c *gin.Context) {
			err := c.ShouldBindJSON(&TagParams{})
			if err != nil {
				resp := map[string]string{}
				if e, ok := err.(val.ValidationErrors); ok {
					for _, v := range e {
						transObj, _ := c.Get("translator")
						trans := transObj.(ut.Translator)
						resp[v.Field()] = v.Translate(trans)
					}
				}
				c.JSON(400, resp)
			}
		})
	}
}

type TagParams struct {
	Name  string `form:"name" json:"name" binding:"required,prefix=aaa"`
	Owner string `form:"owner" json:"owner" binding:"required,prefix=yyy,max=5,min=3"`
}
