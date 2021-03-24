package cvm

import (
	"encoding/hex"
	"evm-demo/common"
	"evm-demo/vm"

	"math/big"
	"testing"
)

// 进行测试

var normalAddress, _ = hex.DecodeString("123456abc")
var contractAddress, _ = hex.DecodeString("987654321")
var normalAccount = common.BytesToAddress(normalAddress)
var contactAccont = common.BytesToAddress(contractAddress)

var byteCodeStr = "608060405234801561001057600080fd5b506101c6806100206000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632b225f29146100515780638da5cb5b146100e1575b600080fd5b34801561005d57600080fd5b50610066610138565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100a657808201518184015260208101905061008b565b50505050905090810190601f1680156100d35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156100ed57600080fd5b506100f6610175565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60606040805190810160405280601081526020017f42617365436f6e747261637456302e3100000000000000000000000000000000815250905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820c61f3af4e4533c265ddfd2cc0900fef0674b5ee9fb9bda1cc516b8be90271a740029"
var byteCode, _ = hex.DecodeString(byteCodeStr)

var input, _ = hex.DecodeString("4d9b3d5d")

func TestRunVM(t *testing.T) {

	// 创建账户State
	stateDb := NewAccountStateDb()
	// 创建一个普通账户
	stateDb.CreateAccount(normalAccount)
	stateDb.CreateAccount(contactAccont)

	stateDb.AddBalance(normalAccount, big.NewInt(0x878999988776612))
	stateDb.SetCode(contactAccont, byteCode)

	evmCtx := NewEVMContext(normalAccount, 100, 1200000, 1)
	vmenv := vm.NewEVM(evmCtx, stateDb, vm.Config{})

	ret, leftgas, err := vmenv.Call(vm.AccountRef(normalAccount), contactAccont, input, 1000000, big.NewInt(0))
	t.Logf("ret: %v, leftGas: %v, err: %v, len(ret): %v, hexret: %v", ret, leftgas, err, len(ret), hex.EncodeToString(ret))
}
