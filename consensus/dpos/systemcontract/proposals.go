package systemcontract

import (
	"errors"
	"github.com/hypnosisfoundation/go-hypnosis/accounts/abi"
	"github.com/hypnosisfoundation/go-hypnosis/common"
	"github.com/hypnosisfoundation/go-hypnosis/common/hexutil"
	"github.com/hypnosisfoundation/go-hypnosis/consensus/dpos/vmcaller"
	"github.com/hypnosisfoundation/go-hypnosis/core"
	"github.com/hypnosisfoundation/go-hypnosis/core/state"
	"github.com/hypnosisfoundation/go-hypnosis/core/types"
	"github.com/hypnosisfoundation/go-hypnosis/log"
	"github.com/hypnosisfoundation/go-hypnosis/params"
	"math"
	"math/big"
)

type Proposals struct {
	abi          abi.ABI
	contractAddr common.Address
}

type ProposalInfo struct {
	Id          [4]byte
	Proposer    common.Address
	PType       uint8
	Deposit     *big.Int
	Rate        uint8
	Name        string
	Details     string
	InitBlock   *big.Int
	Guarantee   common.Address
	UpdateBlock *big.Int
	Status      uint8
}

// NewProposals return Proposals contract instance
func NewProposals() *Proposals {
	return &Proposals{
		abi:          abiMap[ValidatorProposalsContractName],
		contractAddr: ValidatorProposalsContractAddr,
	}
}

// AddressProposalSets function AddressProposalSets
func (p *Proposals) AddressProposalSets(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig, addr common.Address, page *big.Int, size *big.Int) ([][4]byte, error) {
	method := "addressProposalSets"
	data, err := p.abi.Pack(method, addr, page, size)

	if err != nil {
		log.Error("can't pack Proposals contract method", "method", method)
		return [][4]byte{}, err
	}

	msg := vmcaller.NewLegacyMessage(header.Coinbase, &p.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)
	result, err := vmcaller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		log.Error("AddressProposalSets result", "error", err)
		return [][4]byte{}, err
	}

	ret, err := p.abi.Unpack(method, result)
	if err != nil {
		return [][4]byte{}, err
	}
	if proposalIds, ok := ret[0].([][4]byte); !ok {
		return [][4]byte{}, errors.New("invalid AddressProposalSets result format")
	} else {
		return proposalIds, nil
	}
}

// AllProposalSets function AllProposalSets
func (p *Proposals) AllProposalSets(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig, page *big.Int, size *big.Int) ([][4]byte, error) {
	method := "allProposalSets"
	data, err := p.abi.Pack(method, page, size)

	if err != nil {
		log.Error("can't pack Proposals contract method", "method", method)
		return [][4]byte{}, err
	}

	msg := vmcaller.NewLegacyMessage(header.Coinbase, &p.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)
	result, err := vmcaller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		log.Error("AllProposalSets result", "error", err)
		return [][4]byte{}, err
	}

	ret, err := p.abi.Unpack(method, result)
	if err != nil {
		return [][4]byte{}, err
	}
	if proposalIds, ok := ret[0].([][4]byte); !ok {
		return [][4]byte{}, errors.New("invalid AddressProposalSets result format")
	} else {
		return proposalIds, nil
	}
}

// AddressProposals function AddressProposals
func (p *Proposals) AddressProposals(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig, addr common.Address, page *big.Int, size *big.Int) ([]ProposalInfo, error) {
	method := "addressProposals"
	data, err := p.abi.Pack(method, addr, page, size)

	if err != nil {
		log.Error("can't pack Proposals contract method", "method", method)
		return []ProposalInfo{}, err
	}

	msg := vmcaller.NewLegacyMessage(header.Coinbase, &p.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)
	result, err := vmcaller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		log.Error("AddressProposals result", "error", err)
		return []ProposalInfo{}, err
	}
	var proposalInfo []ProposalInfo
	err = p.abi.UnpackIntoInterface(&proposalInfo, method, result)
	if err != nil {
		log.Error("AddressProposals Unpack", "error", err)
		return []ProposalInfo{}, err
	}
	return proposalInfo, nil
}

// AllProposals function AllProposals
func (p *Proposals) AllProposals(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig, page *big.Int, size *big.Int) ([]ProposalInfo, error) {
	method := "allProposals"
	data, err := p.abi.Pack(method, page, size)

	if err != nil {
		log.Error("can't pack Proposals contract method", "method", method)
		return []ProposalInfo{}, err
	}

	msg := vmcaller.NewLegacyMessage(header.Coinbase, &p.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)
	result, err := vmcaller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		log.Error("AllProposals result", "error", err)
		return []ProposalInfo{}, err
	}
	var proposalInfo []ProposalInfo
	err = p.abi.UnpackIntoInterface(&proposalInfo, method, result)
	if err != nil {
		log.Error("AllProposals Unpack", "error", err)
		return []ProposalInfo{}, err
	}
	return proposalInfo, nil
}

// AddressProposalCount function AddressProposalCount
func (p *Proposals) AddressProposalCount(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig, addr common.Address) (*big.Int, error) {
	method := "addressProposalCount"
	data, err := p.abi.Pack(method, addr)

	if err != nil {
		log.Error("can't pack Proposals contract method", "method", method)
		return big.NewInt(0), err
	}

	msg := vmcaller.NewLegacyMessage(header.Coinbase, &p.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)
	result, err := vmcaller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		log.Error("AddressProposalCount result", "error", err)
		return big.NewInt(0), err
	}

	ret, err := p.abi.Unpack(method, result)
	if err != nil {
		return big.NewInt(0), err
	}
	if count, ok := ret[0].(*big.Int); !ok {
		return big.NewInt(0), errors.New("invalid AddressProposalCount result format")
	} else {
		return count, nil
	}
}

// ProposalCount function ProposalCount
func (p *Proposals) ProposalCount(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (*big.Int, error) {
	method := "proposalCount"
	data, err := p.abi.Pack(method)

	if err != nil {
		log.Error("can't pack Proposals contract method", "method", method)
		return big.NewInt(0), err
	}

	msg := vmcaller.NewLegacyMessage(header.Coinbase, &p.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)
	result, err := vmcaller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		log.Error("AllProposalSets result", "error", err)
		return big.NewInt(0), err
	}

	ret, err := p.abi.Unpack(method, result)
	if err != nil {
		return big.NewInt(0), err
	}
	if count, ok := ret[0].(*big.Int); !ok {
		return big.NewInt(0), errors.New("invalid AddressProposalSets result format")
	} else {
		return count, nil
	}
}

// GetProposal function GetProposal
func (p *Proposals) GetProposal(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig, id string) (*ProposalInfo, error) {
	method := "proposalInfos"
	idBytes, err := hexutil.Decode(id)
	if err != nil {
		return &ProposalInfo{}, err
	}
	var idByte4 [4]byte
	copy(idByte4[:], idBytes[:4])
	data, err := p.abi.Pack(method, idByte4)

	if err != nil {
		log.Error("can't pack Proposals contract method", "method", method)
		return &ProposalInfo{}, err
	}

	msg := vmcaller.NewLegacyMessage(header.Coinbase, &p.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)
	result, err := vmcaller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		log.Error("GetProposal result", "error", err)
		return &ProposalInfo{}, err
	}
	proposalInfo := &ProposalInfo{}
	err = p.abi.UnpackIntoInterface(proposalInfo, method, result)
	if err != nil {
		log.Error("GetProposal Unpack", "error", err)
		return &ProposalInfo{}, err
	}
	return proposalInfo, nil
}
