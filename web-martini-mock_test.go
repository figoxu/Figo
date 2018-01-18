package Figo

import (
	"log"
	"testing"
)

func TestHttpHelperMockBuilder_FormHelper(t *testing.T) {
	mockBuilder := NewHttpHelperMockBuilder()
	mockBuilder.MockVal("fv", "123.42")
	mockBuilder.MockVal("iv", "123")
	formHelper := mockBuilder.FormHelper()
	log.Println(formHelper.String("fv"))
	log.Println(formHelper.Float32("fv"))
	log.Println(formHelper.Int("iv"))
}

func TestHttpHelperMockBuilder_ParamHelper(t *testing.T) {

	mockBuilder := NewHttpHelperMockBuilder()
	mockBuilder.MockVal("fv", "123.42")
	mockBuilder.MockVal("iv", "123")
	paramHelper := mockBuilder.ParamHelper()
	log.Println(paramHelper.String("fv"))
	log.Println(paramHelper.Float64("fv"))
	log.Println(paramHelper.Int("iv"))
}
