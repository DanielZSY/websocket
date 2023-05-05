package api

import (
	"encoding/json"
	zh1 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"imservice/define/retcode"
	"io"
	"net/http"
)

type RetData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ConnRender(conn *websocket.Conn, data interface{}) (err error) {
	err = conn.WriteJSON(RetData{
		Code: retcode.SUCCESS,
		Msg:  "success",
		Data: data,
	})

	return
}

func Render(w http.ResponseWriter, code int, msg string, data interface{}) (str string) {
	var retData RetData

	retData.Code = code
	retData.Msg = msg
	retData.Data = data

	retJson, _ := json.Marshal(retData)
	str = string(retJson)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = io.WriteString(w, str)
	return
}

func Validate(inputData interface{}) error {

	validate := v10.New()
	zh := zh1.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")

	_ = zh2.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(inputData)
	if err != nil {
		for _, err := range err.(v10.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}

	return nil
}
