package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsTrue(t *testing.T) {
	assert.Nil(t,IsTrue(true,"is true"))
	assert.NotNil(t,IsTrue(false,"is false"))
	assert.Nil(t,IsTrue("true","is true"))
	assert.Nil(t,IsTrue("TRUE","is true"))
	assert.NotNil(t,IsTrue("ABC","is FALSE"))
}

func Test_IsFasle(t *testing.T) {
	assert.NotNil(t,IsFalse(true,"is true"))
	assert.Nil(t,IsFalse(false,"is false"))
	assert.NotNil(t,IsFalse("true","is true"))
	assert.NotNil(t,IsFalse("TRUE","is true"))
	assert.NotNil(t,IsFalse("ABC","is FALSE"))
}

func Test_IsPositive(t *testing.T){
	assert.NotNil(t,IsPositive(true,"value"))
	assert.NotNil(t,IsPositive("sss","value"))
	assert.NotNil(t,IsPositive("0","value"))
	assert.NotNil(t,IsPositive("-1","value"))
	assert.Nil(t,IsPositive("1","value"))
}

func Test_IsNotEmpty(t *testing.T){
	assert.Nil(t,IsNotEmpty(true,"value"))
	assert.Nil(t,IsNotEmpty("sss","value"))
	assert.Nil(t,IsNotEmpty("0","value"))
	assert.NotNil(t,IsNotEmpty("","value"))
	assert.NotNil(t,IsNotEmpty(nil,"value"))
}

func Test_IsEmpty(t *testing.T){
	assert.NotNil(t,IsEmpty(true,"value"))
	assert.NotNil(t,IsEmpty("sss","value"))
	assert.NotNil(t,IsEmpty("0","value"))
	assert.Nil(t,IsEmpty("","value"))
	assert.Nil(t,IsEmpty(nil,"value"))
}

func Test_IsNil(t *testing.T){
	assert.Nil(t,IsNil(nil,"value"))
	assert.NotNil(t,IsNil("","value"))
}

func Test_IsNotNil(t *testing.T){
	assert.NotNil(t,IsNotNil(nil,"value"))
	assert.Nil(t,IsNotNil("","value"))
}

func Test_HasAny(t *testing.T){

	assert.NotNil(t,HasAny(nil,"value"))
	arr:=[]interface{}{"a","b"}
	assert.Nil(t,HasAny(arr,"value"))
	//todo
	//arr2:=[]CheckItem{{},{}}
	//assert.Nil(t,HasAny(arr2,"value"))
}