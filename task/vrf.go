package task

import (
	"bytes"
	"errors"
	"github.com/axieinfinity/bridge-core/utils"
	"github.com/axieinfinity/bridge-v2/contract"
	"github.com/axieinfinity/bridge-v2/models"
	"github.com/axieinfinity/bridge-v2/oracle"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"time"
)

func (r *task) sendFullFillRandomSeedTransactions(task *models.Task) (doneTasks, processingTasks, failedTasks []*models.Task, tx *ethtypes.Transaction) {
	log.Info("[RoninTask][sendFullFillRandomSeedTransactions] Processing task")
	// check whether current operator is assigned or not
	var (
		err         error
		event       = new(contract.RoninVRFCoordinatorRandomSeedRequested)
		caller      *contract.RoninVRFCoordinatorCaller
		contractAbi *abi.ABI
		assignee    common.Address
		finalized   bool
	)
	caller, err = contract.NewRoninVRFCoordinatorCaller(common.HexToAddress(r.contracts[VRF_CONTRACT]), r.client)
	if err != nil {
		goto ERROR
	}
	contractAbi, err = contract.RoninVRFCoordinatorMetaData.GetAbi()
	if err != nil {
		goto ERROR
	}
	err = r.util.UnpackLog(*contractAbi, event, "RandomSeedRequested", common.Hex2Bytes(task.Data))
	if err != nil {
		goto ERROR
	}
	// req is finalized, GOTO update done
	finalized, err = isFinalized(caller, event.ReqHash)
	if err != nil {
		goto ERROR
	}
	if finalized {
		goto DONE
	}
	// TODO: add assignee to request event
	// currently, we set the assignee address to operator itself to bypass the condition
	assignee = r.listener.GetBridgeOperatorSign().GetAddress()
	if isAssigned(r.listener.GetBridgeOperatorSign(), r.listener.GetBridgeOperatorSign().GetAddress()) {
		tx, err = tryFulfilRandomSeed(r, event.ReqHash, event.Request, assignee)
		if err != nil {
			goto ERROR
		}
		goto PROCESSING
	}
	// if not, check if the task has been waiting too long
	if waitTooLong(r, task) {
		// try to get the next assignee
		assignee, err = getNextAssignee(caller)
		if err != nil {
			goto ERROR
		}
		// FIXME: update createdTime of current task to now
		task.CreatedAt = time.Now().Unix()
		if !isAssigned(r.listener.GetBridgeOperatorSign(), assignee) {
			goto PENDING
		}
		tx, err = tryFulfilRandomSeed(r, event.ReqHash, event.Request, assignee)
		if err != nil {
			goto ERROR
		}
		goto PROCESSING
	}
PENDING:
	// otherwise task is still pending but retry count is increased by 1
	task.Retries++
	go updateTasks(r.store, []*models.Task{task}, STATUS_PENDING, "", 0, r.releaseTasksCh)
	return nil, nil, nil, nil
PROCESSING:
	// append task to prcessing tasks
	processingTasks = append(processingTasks, task)
	return
DONE:
	// append task to done tasks
	doneTasks = append(doneTasks, task)
	return
ERROR:
	r.appendFailedTask(err, task, failedTasks)
	return nil, nil, failedTasks, nil
}

func waitTooLong(r *task, task *models.Task) bool {
	deadline := int64(VRFConfig.WaitForBlock) * int64(r.listener.Config().LoadInterval.Seconds())
	return task.CreatedAt+deadline < time.Now().Unix()
}

func tryFulfilRandomSeed(r *task, requestHash [32]byte, request contract.SLARandomRequest, assignee common.Address) (*ethtypes.Transaction, error) {
	transactor, err := contract.NewRoninVRFCoordinatorTransactor(common.HexToAddress(r.contracts[VRF_CONTRACT]), r.client)
	if err != nil {
		return nil, err
	}
	proof, err := generateVRFProof(assignee, requestHash, request)
	if err != nil {
		return nil, err
	}
	// otherwise create new fulfill request transaction and add task to processing
	return r.util.SendContractTransaction(r.listener.GetBridgeOperatorSign(), r.chainId, r.gasLimitBumpRatio, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return transactor.FulfillRandomSeed(opts, proof, request)
	})
}

func isAssigned(signer utils.ISign, assignee common.Address) bool {
	if signer == nil || bytes.Equal(signer.GetAddress().Bytes(), assignee.Bytes()) {
		return false
	}
	return true
}

func isFinalized(caller *contract.RoninVRFCoordinatorCaller, req [32]byte) (bool, error) {
	result, err := caller.RequestFinalized(nil, req)
	if err != nil {
		return false, err
	}
	return result, nil
}

func generateVRFProof(assignee common.Address, requestHash [32]byte, request contract.SLARandomRequest) (contract.VRFProof, error) {
	res, err := oracle.GenerateProofResponse(VRFConfig.Key, request, VRFConfig.KeyHash, assignee, requestHash)
	if err != nil {
		return contract.VRFProof{}, err
	}
	return res.Proof, nil
}

// TODO: call to contract to get next assignee in case current assignee cannot fulfill request
func getNextAssignee(caller *contract.RoninVRFCoordinatorCaller) (common.Address, error) {
	return common.Address{}, errors.New("function getNextAssignee is not implemented yet")
}
