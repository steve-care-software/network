package layers

import (
	"reflect"
	"testing"
)

func TestAssignableResource_withCompile_Success(t *testing.T) {
	compile := "myVariable"
	ins := NewAssignableResourceWithCompileForTests(compile)

	if !ins.IsCompile() {
		t.Errorf("the assignableResource was expected to contain a compile")
		return
	}

	if ins.IsDecompile() {
		t.Errorf("the assignableResource was expected to NOT contain a decompile")
		return
	}

	if ins.IsAmountByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an amountByQuery")
		return
	}

	if ins.IsListByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an amountByQuery")
		return
	}

	if ins.IsRetrieveByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an retrieveByQuery")
		return
	}

	if ins.IsRetrieveByHash() {
		t.Errorf("the assignableResource was expected to NOT contain an retrieveByHash")
		return
	}

	if ins.IsAmount() {
		t.Errorf("the assignableResource was expected to NOT contain an amount")
		return
	}

	retCompile := ins.Compile()
	if !reflect.DeepEqual(compile, retCompile) {
		t.Errorf("the compile is invalid")
		return
	}
}

func TestAssignableResource_withDecompile_Success(t *testing.T) {
	decompile := "myVariable"
	ins := NewAssignableResourceWithDecompileForTests(decompile)

	if ins.IsCompile() {
		t.Errorf("the assignableResource was expected to NOT contain a compile")
		return
	}

	if !ins.IsDecompile() {
		t.Errorf("the assignableResource was expected to contain a decompile")
		return
	}

	if ins.IsAmountByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an amountByQuery")
		return
	}

	if ins.IsListByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an amountByQuery")
		return
	}

	if ins.IsRetrieveByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an retrieveByQuery")
		return
	}

	if ins.IsRetrieveByHash() {
		t.Errorf("the assignableResource was expected to NOT contain an retrieveByHash")
		return
	}

	if ins.IsAmount() {
		t.Errorf("the assignableResource was expected to NOT contain an amount")
		return
	}

	retDecompile := ins.Decompile()
	if decompile != retDecompile {
		t.Errorf("the decompile was expected to be '%s', '%s' returned", decompile, retDecompile)
		return
	}
}

func TestAssignableResource_withAmountByQuery_Success(t *testing.T) {
	amountByQuery := "myVariable"
	ins := NewAssignableResourceWithAmountByQueryForTests(amountByQuery)

	if ins.IsCompile() {
		t.Errorf("the assignableResource was expected to NOT contain a compile")
		return
	}

	if ins.IsDecompile() {
		t.Errorf("the assignableResource was expected to NOT contain a decompile")
		return
	}

	if !ins.IsAmountByQuery() {
		t.Errorf("the assignableResource was expected to contain an amountByQuery")
		return
	}

	if ins.IsListByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an amountByQuery")
		return
	}

	if ins.IsRetrieveByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an retrieveByQuery")
		return
	}

	if ins.IsRetrieveByHash() {
		t.Errorf("the assignableResource was expected to NOT contain an retrieveByHash")
		return
	}

	if ins.IsAmount() {
		t.Errorf("the assignableResource was expected to NOT contain an amount")
		return
	}

	retAmountByQuery := ins.AmountByQuery()
	if !reflect.DeepEqual(amountByQuery, retAmountByQuery) {
		t.Errorf("the amountByQuery is invalid")
		return
	}
}

func TestAssignableResource_withListByQuery_Success(t *testing.T) {
	listByQuery := "myVariable"
	ins := NewAssignableResourceWithListByQueryForTests(listByQuery)

	if ins.IsCompile() {
		t.Errorf("the assignableResource was expected to NOT contain a compile")
		return
	}

	if ins.IsDecompile() {
		t.Errorf("the assignableResource was expected to NOT contain a decompile")
		return
	}

	if ins.IsAmountByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an amountByQuery")
		return
	}

	if !ins.IsListByQuery() {
		t.Errorf("the assignableResource was expected to contain an amountByQuery")
		return
	}

	if ins.IsRetrieveByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an retrieveByQuery")
		return
	}

	if ins.IsRetrieveByHash() {
		t.Errorf("the assignableResource was expected to NOT contain an retrieveByHash")
		return
	}

	if ins.IsAmount() {
		t.Errorf("the assignableResource was expected to NOT contain an amount")
		return
	}

	retListByQuery := ins.ListByQuery()
	if !reflect.DeepEqual(listByQuery, retListByQuery) {
		t.Errorf("the listByQuery is invalid")
		return
	}
}

func TestAssignableResource_withRetrieveByQuery_Success(t *testing.T) {
	retrieveByQuery := "myVariable"
	ins := NewAssignableResourceWithRetrieveByQueryForTests(retrieveByQuery)

	if ins.IsCompile() {
		t.Errorf("the assignableResource was expected to NOT contain a compile")
		return
	}

	if ins.IsDecompile() {
		t.Errorf("the assignableResource was expected to NOT contain a decompile")
		return
	}

	if ins.IsAmountByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an amountByQuery")
		return
	}

	if ins.IsListByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an amountByQuery")
		return
	}

	if !ins.IsRetrieveByQuery() {
		t.Errorf("the assignableResource was expected to contain an retrieveByQuery")
		return
	}

	if ins.IsRetrieveByHash() {
		t.Errorf("the assignableResource was expected to NOT contain an retrieveByHash")
		return
	}

	if ins.IsAmount() {
		t.Errorf("the assignableResource was expected to NOT contain an amount")
		return
	}

	retRetrieveByQuery := ins.RetrieveByQuery()
	if !reflect.DeepEqual(retrieveByQuery, retRetrieveByQuery) {
		t.Errorf("the retrieveByQuery is invalid")
		return
	}
}

func TestAssignableResource_withRetrieveByHash_Success(t *testing.T) {
	retrieveByHash := "myVariable"
	ins := NewAssignableResourceWithRetrieveByHashForTests(retrieveByHash)

	if ins.IsCompile() {
		t.Errorf("the assignableResource was expected to NOT contain a compile")
		return
	}

	if ins.IsDecompile() {
		t.Errorf("the assignableResource was expected to NOT contain a decompile")
		return
	}

	if ins.IsAmountByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an amountByQuery")
		return
	}

	if ins.IsListByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an amountByQuery")
		return
	}

	if ins.IsRetrieveByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an retrieveByQuery")
		return
	}

	if !ins.IsRetrieveByHash() {
		t.Errorf("the assignableResource was expected to contain an retrieveByHash")
		return
	}

	if ins.IsAmount() {
		t.Errorf("the assignableResource was expected to NOT contain an amount")
		return
	}

	retRetrieveByHash := ins.RetrieveByHash()
	if !reflect.DeepEqual(retrieveByHash, retRetrieveByHash) {
		t.Errorf("the retrieveByHash is invalid")
		return
	}
}

func TestAssignableResource_withAmount_Success(t *testing.T) {
	ins := NewAssignableResourceWithAmountForTests()

	if ins.IsCompile() {
		t.Errorf("the assignableResource was expected to NOT contain a compile")
		return
	}

	if ins.IsDecompile() {
		t.Errorf("the assignableResource was expected to NOT contain a decompile")
		return
	}

	if ins.IsAmountByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an amountByQuery")
		return
	}

	if ins.IsListByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an amountByQuery")
		return
	}

	if ins.IsRetrieveByQuery() {
		t.Errorf("the assignableResource was expected to NOT contain an retrieveByQuery")
		return
	}

	if ins.IsRetrieveByHash() {
		t.Errorf("the assignableResource was expected to NOT contain an retrieveByHash")
		return
	}

	if !ins.IsAmount() {
		t.Errorf("the assignableResource was expected to contain an amount")
		return
	}
}

func TestAssignableResource_withoutParam_returnsError(t *testing.T) {
	_, err := NewAssignableResourceBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
