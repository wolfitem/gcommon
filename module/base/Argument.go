package base


//堆栈信息中过滤掉前N个
const SKIP=3

type CheckItem struct {
	CheckFunc CheckFunc
	Value interface{}
	Param string
}

func  CheckList(items ...CheckItem) (err error){
	for _, v := range items {
		if err=v.CheckFunc(v.Value,v.Param);err!=nil{
			return
		}
	}
	return
}

func  Assert(f CheckFunc,v interface{},param string){
	err:=f(v,param)
	panic_stack(err)
}

func  AssertList(items ...CheckItem){
	for _, v := range items {
		if err:=v.CheckFunc(v.Value,v.Param);err!=nil{
			panic_stack(err)
			return
		}
	}
}


func CheckError(err error){
	panic_stack(err)
}

func panic_stack(err error){
	if err!=nil{
		panic(New_Skip(err.Error(),SKIP).Stack())
	}
}

