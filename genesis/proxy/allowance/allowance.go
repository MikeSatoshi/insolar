package allowance

import (
        "github.com/insolar/insolar/core"
        "github.com/insolar/insolar/logicrunner/goplugin/proxyctx"
)



// Reference to class of this contract
var ClassReference = core.NewRefFromBase58("")

// Contract proxy type
type Allowance struct {
    Reference core.RecordRef
}

type ContractHolder struct {
	data []byte
}

func (r *ContractHolder) AsChild(objRef core.RecordRef) *Allowance {
    ref, err := proxyctx.Current.SaveAsChild(objRef, ClassReference, r.data)
    if err != nil {
        panic(err)
    }
    return &Allowance{Reference: ref}
}

func (r *ContractHolder) AsDelegate(objRef core.RecordRef) *Allowance {
    ref, err := proxyctx.Current.SaveAsDelegate(objRef, ClassReference, r.data)
    if err != nil {
        panic(err)
    }
    return &Allowance{Reference: ref}
}

// GetObject
func GetObject(ref core.RecordRef) (r *Allowance) {
    return &Allowance{Reference: ref}
}

func GetClass() core.RecordRef {
    return ClassReference
}

func GetImplementationFrom(object core.RecordRef) *Allowance {
    ref, err := proxyctx.Current.GetDelegate(object, ClassReference)
    if err != nil {
        panic(err)
    }
    return GetObject(ref)
}


func New( to *core.RecordRef, amount uint, expire int64 ) *ContractHolder {
    var args [3]interface{}
	args[0] = to
	args[1] = amount
	args[2] = expire


    var argsSerialized []byte
    err := proxyctx.Current.Serialize(args, &argsSerialized)
    if err != nil {
        panic(err)
    }

    data, err := proxyctx.Current.RouteConstructorCall(ClassReference, "New", argsSerialized)
    if err != nil {
		panic(err)
    }

    return &ContractHolder{data: data}
}


// GetReference
func (r *Allowance) GetReference() core.RecordRef {
    return r.Reference
}

// GetClass
func (r *Allowance) GetClass() core.RecordRef {
    return ClassReference
}


func (r *Allowance) IsExpired(  ) ( bool ) {
    var args [0]interface{}

    var argsSerialized []byte

    err := proxyctx.Current.Serialize(args, &argsSerialized)
    if err != nil {
        panic(err)
    }

    res, err := proxyctx.Current.RouteCall(r.Reference, "IsExpired", argsSerialized)
    if err != nil {
   		panic(err)
    }

    resList := [1]interface{}{}
	var a0 bool
	resList[0] = a0

    err = proxyctx.Current.Deserialize(res, &resList)
    if err != nil {
        panic(err)
    }

    return resList[0].(bool)
}

func (r *Allowance) TakeAmount(  ) ( uint ) {
    var args [0]interface{}

    var argsSerialized []byte

    err := proxyctx.Current.Serialize(args, &argsSerialized)
    if err != nil {
        panic(err)
    }

    res, err := proxyctx.Current.RouteCall(r.Reference, "TakeAmount", argsSerialized)
    if err != nil {
   		panic(err)
    }

    resList := [1]interface{}{}
	var a0 uint
	resList[0] = a0

    err = proxyctx.Current.Deserialize(res, &resList)
    if err != nil {
        panic(err)
    }

    return resList[0].(uint)
}

func (r *Allowance) GetBalanceForOwner(  ) ( uint ) {
    var args [0]interface{}

    var argsSerialized []byte

    err := proxyctx.Current.Serialize(args, &argsSerialized)
    if err != nil {
        panic(err)
    }

    res, err := proxyctx.Current.RouteCall(r.Reference, "GetBalanceForOwner", argsSerialized)
    if err != nil {
   		panic(err)
    }

    resList := [1]interface{}{}
	var a0 uint
	resList[0] = a0

    err = proxyctx.Current.Deserialize(res, &resList)
    if err != nil {
        panic(err)
    }

    return resList[0].(uint)
}

func (r *Allowance) DeleteExpiredAllowance(  ) ( uint ) {
    var args [0]interface{}

    var argsSerialized []byte

    err := proxyctx.Current.Serialize(args, &argsSerialized)
    if err != nil {
        panic(err)
    }

    res, err := proxyctx.Current.RouteCall(r.Reference, "DeleteExpiredAllowance", argsSerialized)
    if err != nil {
   		panic(err)
    }

    resList := [1]interface{}{}
	var a0 uint
	resList[0] = a0

    err = proxyctx.Current.Deserialize(res, &resList)
    if err != nil {
        panic(err)
    }

    return resList[0].(uint)
}

