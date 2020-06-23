/*
 * Copyright (c) 2019. Baidu Inc. All Rights Reserved.
 */

package pb

import (
	"encoding/hex"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"math/big"
)

type HexID []byte

func (h HexID) MarshalJSON() ([]byte, error) {
	hex := hex.EncodeToString(h)
	return json.Marshal(hex)
}
func (m *HexID) Reset()         { *m = HexID{} }
func (m *HexID) String() string { return proto.CompactTextString(m) }
func (*HexID) ProtoMessage()    {}

type JsonTxInput struct {
	RefTxid   string `json:"refTxid"`
	RefOffset int32  `json:"refOffset"`
	FromAddr  string `json:"fromAddr"`
	Amount    int64  `json:"amount"`
}

func (m *JsonTxInput) Reset()         { *m = JsonTxInput{} }
func (m *JsonTxInput) String() string { return proto.CompactTextString(m) }
func (*JsonTxInput) ProtoMessage()    {}

type JsonTxOutput struct {
	Amount int64  `json:"amount"`
	ToAddr string `json:"toAddr"`
}

func (m *JsonTxOutput) Reset()         { *m = JsonTxOutput{} }
func (m *JsonTxOutput) String() string { return proto.CompactTextString(m) }
func (*JsonTxOutput) ProtoMessage()    {}

type JsonTxInputExt struct {
	Bucket    string `json:"bucket"`
	Key       string `json:"key"`
	RefTxid   string `json:"refTxid"`
	RefOffset int32  `json:"refOffset"`
}

func (m *JsonTxInputExt) Reset()         { *m = JsonTxInputExt{} }
func (m *JsonTxInputExt) String() string { return proto.CompactTextString(m) }
func (*JsonTxInputExt) ProtoMessage()    {}

type JsonTxOutputExt struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

func (m *JsonTxOutputExt) Reset()         { *m = JsonTxOutputExt{} }
func (m *JsonTxOutputExt) String() string { return proto.CompactTextString(m) }
func (*JsonTxOutputExt) ProtoMessage()    {}

type JsonResourceLimit struct {
	Type  string `json:"type"`
	Limit int64  `json:"limit"`
}

func (m *JsonResourceLimit) Reset()         { *m = JsonResourceLimit{} }
func (m *JsonResourceLimit) String() string { return proto.CompactTextString(m) }
func (*JsonResourceLimit) ProtoMessage()    {}

type JsonInvokeRequest struct {
	ModuleName    string              `json:"moduleName"`
	ContractName  string              `json:"contractName"`
	MethodName    string              `json:"methodName"`
	Args          map[string]string   `json:"args"`
	ResouceLimits []JsonResourceLimit `json:"resource_limits"`
}

func (m *JsonInvokeRequest) Reset()         { *m = JsonInvokeRequest{} }
func (m *JsonInvokeRequest) String() string { return proto.CompactTextString(m) }
func (*JsonInvokeRequest) ProtoMessage()    {}

type JsonSignatureInfo struct {
	PublicKey string `json:"publickey"`
	Sign      string `json:"sign"`
}

func (m *JsonSignatureInfo) Reset()         { *m = JsonSignatureInfo{} }
func (m *JsonSignatureInfo) String() string { return proto.CompactTextString(m) }
func (*JsonSignatureInfo) ProtoMessage()    {}

type JsonSignInfo struct {
	Address   string `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	PublicKey string `protobuf:"bytes,2,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	Sign      []byte `protobuf:"bytes,3,opt,name=Sign,proto3" json:"Sign,omitempty"`
}

func (m *JsonSignInfo) Reset()         { *m = JsonSignInfo{} }
func (m *JsonSignInfo) String() string { return proto.CompactTextString(m) }
func (*JsonSignInfo) ProtoMessage()    {}

type JsonQCSignInfos struct {
	// QCSignInfos
	QCSignInfos []*JsonSignInfo `protobuf:"bytes,1,rep,name=QCSignInfos,proto3" json:"QCSignInfos,omitempty"`
}

func (m *JsonQCSignInfos) Reset()         { *m = JsonQCSignInfos{} }
func (m *JsonQCSignInfos) String() string { return proto.CompactTextString(m) }
func (*JsonQCSignInfos) ProtoMessage()    {}

type JsonQCState int32

func (m *JsonQCState) Reset()         { *m = 0 }
func (m *JsonQCState) String() string { return proto.CompactTextString(m) }
func (*JsonQCState) ProtoMessage()    {}

type JsonQuorumCert struct {
	// The id of Proposal this QC certified.
	ProposalId string `protobuf:"bytes,1,opt,name=ProposalId,proto3" json:"ProposalId,omitempty"`
	// The msg of Proposal this QC certified.
	ProposalMsg []byte `protobuf:"bytes,2,opt,name=ProposalMsg,proto3" json:"ProposalMsg,omitempty"`
	// The current type of this QC certified.
	// the type contains `NEW_VIEW`, `PREPARE`
	Type JsonQCState `protobuf:"varint,3,opt,name=Type,proto3,enum=pb.QCState" json:"Type,omitempty"`
	// The view number of this QC certified.
	ViewNumber int64 `protobuf:"varint,4,opt,name=ViewNumber,proto3" json:"ViewNumber,omitempty"`
	// SignInfos is the signs of the leader gathered from replicas
	// of a specifically certType.
	SignInfos *JsonQCSignInfos `protobuf:"bytes,5,opt,name=SignInfos,proto3" json:"SignInfos,omitempty"`
}

func (m *JsonQuorumCert) Reset()         { *m = JsonQuorumCert{} }
func (m *JsonQuorumCert) String() string { return proto.CompactTextString(m) }
func (*JsonQuorumCert) ProtoMessage()    {}

type JsonModifyBlock struct {
	Marked          bool   `json:"marked,omitempty"`
	EffectiveHeight int64  `json:"effectiveHeight,omitempty"`
	EffectiveTxid   string `json:"effectiveTxid,omitempty"`
}

func (m *JsonModifyBlock) Reset()         { *m = JsonModifyBlock{} }
func (m *JsonModifyBlock) String() string { return proto.CompactTextString(m) }
func (*JsonModifyBlock) ProtoMessage()    {}

type JsonTransaction struct {
	Txid              string               `json:"txid,omitempty"`
	Blockid           string               `json:"blockid,omitempty"`
	TxInputs          []JsonTxInput        `json:"txInputs,omitempty"`
	TxOutputs         []JsonTxOutput       `json:"txOutputs,omitempty"`
	Desc              string               `json:"desc,omitempty"`
	Nonce             string               `json:"nonce,omitempty"`
	Timestamp         int64                `json:"timestamp,omitempty"`
	Version           int32                `json:"version,omitempty"`
	Autogen           bool                 `json:"autogen,omitempty"`
	Coinbase          bool                 `json:"coinbase,omitempty"`
	VoteCoinbase      bool                 `json:"voteCoinbase,omitempty"`
	TxInputsExt       []JsonTxInputExt     `json:"txInputsExt,omitempty"`
	TxOutputsExt      []JsonTxOutputExt    `json:"txOutputsExt,omitempty"`
	ContractRequests  []*JsonInvokeRequest `json:"contractRequests,omitempty"`
	Initiator         string               `json:"initiator"`
	AuthRequire       []string             `json:"authRequire,omitempty"`
	InitiatorSigns    []JsonSignatureInfo  `json:"initiatorSigns,omitempty"`
	AuthRequireSigns  []JsonSignatureInfo  `json:"authRequireSigns,omitempty"`
	ReceivedTimestamp int64                `json:"receivedTimestamp,omitempty"`
	//ModifyBlock       JsonModifyBlock      `json:"modifyBlock,omitempty"`
}

func (m *JsonTransaction) Reset()         { *m = JsonTransaction{} }
func (m *JsonTransaction) String() string { return proto.CompactTextString(m) }
func (*JsonTransaction) ProtoMessage()    {}

type BigInt big.Int

func (b *BigInt) MarshalJSON() ([]byte, error) {
	str := (*big.Int)(b).String()
	return json.Marshal(str)
}
func (m *BigInt) Reset()         { *m = BigInt{} }
func (m *BigInt) String() string { return proto.CompactTextString(m) }
func (*BigInt) ProtoMessage()    {}

type JsonInternalBlock struct {
	Version      int32              `json:"version"`
	Blockid      HexID              `json:"blockid"`
	PreHash      HexID              `json:"preHash"`
	Proposer     string             `json:"proposer"`
	Sign         HexID              `json:"sign"`
	Pubkey       string             `json:"pubkey"`
	MerkleRoot   HexID              `json:"merkleRoot"`
	Height       int64              `json:"height"`
	Timestamp    int64              `json:"timestamp"`
	Transactions []*JsonTransaction `json:"transactions"`
	TxCount      int32              `json:"txCount"`
	MerkleTree   []HexID            `json:"merkleTree"`
	InTrunk      bool               `json:"inTrunk"`
	NextHash     HexID              `json:"nextHash"`
	FailedTxs    map[string]string  `json:"failedTxs"`
	CurTerm      int64              `json:"curTerm"`
	CurBlockNum  int64              `json:"curBlockNum"`
	Justify      *JsonQuorumCert    `json:"justify"`
}

func (m *JsonInternalBlock) Reset()         { *m = JsonInternalBlock{} }
func (m *JsonInternalBlock) String() string { return proto.CompactTextString(m) }
func (*JsonInternalBlock) ProtoMessage()    {}

type JsonLedgerMeta struct {
	// RootBlockid root block id
	RootBlockid HexID `json:"rootBlockid"`
	// TipBlockid TipBlockid
	TipBlockid HexID `json:"tipBlockid"`
	// TrunkHeight TrunkHeight
	TrunkHeight int64 `json:"trunkHeight"`
}

func (m *JsonLedgerMeta) Reset()         { *m = JsonLedgerMeta{} }
func (m *JsonLedgerMeta) String() string { return proto.CompactTextString(m) }
func (*JsonLedgerMeta) ProtoMessage()    {}

type JsonUtxoMeta struct {
	// LatestBlockid LatestBlockid
	LatestBlockid HexID `json:"latestBlockid"`
	// LockKeyList LockKeyList
	LockKeyList []string `json:"lockKeyList"`
	// UtxoTotal UtxoTotal
	UtxoTotal string `json:"utxoTotal"`
	// Average confirmed dealy (ms)
	AvgDelay int64 `json:"avgDelay"`
	// Current unconfirmed tx amount
	UnconfirmTxAmount int64 `json:"unconfirmed"`
	// MaxBlockSize MaxBlockSize
	MaxBlockSize int64 `json:"maxBlockSize"`
	// ReservedContracts ReservedContracts
	ReservedContracts []JsonInvokeRequest `json:"reservedContracts"`
	// ForbiddenContract forbidden contract
	ForbiddenContract JsonInvokeRequest `json:"forbiddenContract"`
	// NewAccountResourceAmount resource amount of creating an account
	NewAccountResourceAmount int64 `json:"newAccountResourceAmount"`
	TransferFeeAmount        int64 `json:"transfer_fee_amount,omitempty"`
	// IrreversibleBlockHeight irreversible block height
	IrreversibleBlockHeight int64 `json:"irreversibleBlockHeight"`
	// IrreversibleSlideWindow irreversible slide window
	IrreversibleSlideWindow int64 `json:"irreversibleSlideWindow"`
	// GasPrice gas rate to utxo for different type resources
	GasPrice GasPrice `json:"gasPrice"`
}

func (m *JsonUtxoMeta) Reset()         { *m = JsonUtxoMeta{} }
func (m *JsonUtxoMeta) String() string { return proto.CompactTextString(m) }
func (*JsonUtxoMeta) ProtoMessage()    {}

type JsonContractStatData struct {
	AccountCount  int64 `json:"accountCount"`
	ContractCount int64 `json:"contractCount"`
}

func (m *JsonContractStatData) Reset()         { *m = JsonContractStatData{} }
func (m *JsonContractStatData) String() string { return proto.CompactTextString(m) }
func (*JsonContractStatData) ProtoMessage()    {}

type ChainStatus struct {
	Name       string         `json:"name"`
	LedgerMeta JsonLedgerMeta `json:"ledger"`
	UtxoMeta   JsonUtxoMeta   `json:"utxo"`
	// add BranchBlockid
	BranchBlockid []string `json:"branchBlockid"`
}

func (m *ChainStatus) Reset()         { *m = ChainStatus{} }
func (m *ChainStatus) String() string { return proto.CompactTextString(m) }
func (*ChainStatus) ProtoMessage()    {}

type SystemStatus struct {
	ChainStatus []ChainStatus `json:"blockchains"`
	Peers       []string      `json:"peers"`
	Speeds      *Speeds       `json:"speeds"`
}

func (m *SystemStatus) Reset()         { *m = SystemStatus{} }
func (m *SystemStatus) String() string { return proto.CompactTextString(m) }
func (*SystemStatus) ProtoMessage()    {}

type TriggerDesc struct {
	Module string      `json:"module"`
	Method string      `json:"method"`
	Args   interface{} `json:"args"`
	Height int64       `json:"height"`
}

func (m *TriggerDesc) Reset()         { *m = TriggerDesc{} }
func (m *TriggerDesc) String() string { return proto.CompactTextString(m) }
func (*TriggerDesc) ProtoMessage()    {}

type ContractDesc struct {
	Module  string      `json:"module"`
	Method  string      `json:"method"`
	Args    interface{} `json:"args"`
	Trigger TriggerDesc `json:"trigger"`
}

func (m *ContractDesc) Reset()         { *m = ContractDesc{} }
func (m *ContractDesc) String() string { return proto.CompactTextString(m) }
func (*ContractDesc) ProtoMessage()    {}

func FromSystemStatusPB(statuspb *SystemsStatus) *SystemStatus {
	status := &SystemStatus{}
	for _, chain := range statuspb.GetBcsStatus() {
		ledgerMeta := chain.GetMeta()
		utxoMeta := chain.GetUtxoMeta()
		ReservedContracts := utxoMeta.GetReservedContracts()
		rcs := []JsonInvokeRequest{}
		for _, rcpb := range ReservedContracts {
			args := map[string]string{}
			for k, v := range rcpb.GetArgs() {
				args[k] = string(v)
			}
			rc := JsonInvokeRequest{
				ModuleName:   rcpb.GetModuleName(),
				ContractName: rcpb.GetContractName(),
				MethodName:   rcpb.GetMethodName(),
				Args:         args,
			}
			rcs = append(rcs, rc)
		}
		forbiddenContract := utxoMeta.GetForbiddenContract()
		args := forbiddenContract.GetArgs()
		originalArgs := map[string]string{}
		for key, value := range args {
			originalArgs[key] = string(value)
		}
		forbiddenContractMap := JsonInvokeRequest{
			ModuleName:   forbiddenContract.GetModuleName(),
			ContractName: forbiddenContract.GetContractName(),
			MethodName:   forbiddenContract.GetMethodName(),
			Args:         originalArgs,
		}
		gasPricePB := utxoMeta.GetGasPrice()
		gasPrice := GasPrice{
			CpuRate:  gasPricePB.GetCpuRate(),
			MemRate:  gasPricePB.GetMemRate(),
			DiskRate: gasPricePB.GetDiskRate(),
			XfeeRate: gasPricePB.GetXfeeRate(),
		}
		status.ChainStatus = append(status.ChainStatus, ChainStatus{
			Name: chain.GetBcname(),
			LedgerMeta: JsonLedgerMeta{
				RootBlockid: ledgerMeta.GetRootBlockid(),
				TipBlockid:  ledgerMeta.GetTipBlockid(),
				TrunkHeight: ledgerMeta.GetTrunkHeight(),
			},
			UtxoMeta: JsonUtxoMeta{
				LatestBlockid:            utxoMeta.GetLatestBlockid(),
				LockKeyList:              utxoMeta.GetLockKeyList(),
				UtxoTotal:                utxoMeta.GetUtxoTotal(),
				AvgDelay:                 utxoMeta.GetAvgDelay(),
				UnconfirmTxAmount:        utxoMeta.GetUnconfirmTxAmount(),
				MaxBlockSize:             utxoMeta.GetMaxBlockSize(),
				NewAccountResourceAmount: utxoMeta.GetNewAccountResourceAmount(),
				TransferFeeAmount:        utxoMeta.GetTransferFeeAmount(),
				ReservedContracts:        rcs,
				ForbiddenContract:        forbiddenContractMap,
				// Irreversible block height & slide window
				IrreversibleBlockHeight: utxoMeta.GetIrreversibleBlockHeight(),
				IrreversibleSlideWindow: utxoMeta.GetIrreversibleSlideWindow(),
				// add GasPrice value
				GasPrice: gasPrice,
			},
			BranchBlockid: chain.GetBranchBlockid(),
		})
	}
	status.Peers = statuspb.GetPeerUrls()
	status.Speeds = statuspb.GetSpeeds()
	return status
}

func FromPBJustify(qc *QuorumCert) *JsonQuorumCert {
	justify := &JsonQuorumCert{}
	if qc != nil {
		justify.ProposalId = hex.EncodeToString(qc.ProposalId)
		justify.ProposalMsg = qc.ProposalMsg
		justify.Type = JsonQCState(int(qc.Type))
		justify.ViewNumber = qc.ViewNumber
		justify.SignInfos = &JsonQCSignInfos{
			QCSignInfos: make([]*JsonSignInfo, 0),
		}
		for _, sign := range qc.SignInfos.QCSignInfos {
			tmpSign := &JsonSignInfo{
				Address:   sign.Address,
				PublicKey: sign.PublicKey,
				Sign:      sign.Sign,
			}
			justify.SignInfos.QCSignInfos = append(justify.SignInfos.QCSignInfos, tmpSign)
		}
	}
	return justify
}

func FromInternalBlockPB(block *InternalBlock) *JsonInternalBlock {
	iblock := &JsonInternalBlock{
		Version:     block.Version,
		Blockid:     block.Blockid,
		PreHash:     block.PreHash,
		Proposer:    string(block.Proposer),
		Sign:        block.Sign,
		Pubkey:      string(block.Pubkey),
		MerkleRoot:  block.MerkleRoot,
		Height:      block.Height,
		Timestamp:   block.Timestamp,
		TxCount:     block.TxCount,
		InTrunk:     block.InTrunk,
		NextHash:    block.NextHash,
		FailedTxs:   block.FailedTxs,
		CurTerm:     block.CurTerm,
		CurBlockNum: block.CurBlockNum,
	}
	iblock.MerkleTree = make([]HexID, len(block.MerkleTree))
	for i := range block.MerkleTree {
		iblock.MerkleTree[i] = block.MerkleTree[i]
	}
	iblock.Transactions = make([]*JsonTransaction, len(block.Transactions))
	for i := range block.Transactions {
		iblock.Transactions[i] = GetTxFromTransaction(block.Transactions[i])
	}
	iblock.Justify = FromPBJustify(block.Justify)
	return iblock
}

func GetTxFromTransaction(tx *Transaction) *JsonTransaction {
	t := &JsonTransaction{
		Txid:              hex.EncodeToString(tx.Txid),
		Blockid:           hex.EncodeToString(tx.Blockid),
		Nonce:             tx.Nonce,
		Timestamp:         tx.Timestamp,
		Version:           tx.Version,
		Desc:              string(tx.Desc),
		Autogen:           tx.Autogen,
		Coinbase:          tx.Coinbase,
		VoteCoinbase:      tx.VoteCoinbase,
		Initiator:         tx.Initiator,
		ReceivedTimestamp: tx.ReceivedTimestamp,
	}
	for _, input := range tx.TxInputs {
		t.TxInputs = append(t.TxInputs, JsonTxInput{
			RefTxid:   hex.EncodeToString(input.RefTxid),
			RefOffset: input.RefOffset,
			FromAddr:  string(input.FromAddr),
			Amount:    big.NewInt(0).SetBytes(input.Amount).Int64(),
		})
	}
	for _, output := range tx.TxOutputs {
		t.TxOutputs = append(t.TxOutputs, JsonTxOutput{
			Amount: big.NewInt(0).SetBytes(output.Amount).Int64(),
			ToAddr: string(output.ToAddr),
		})
	}
	for _, inputExt := range tx.TxInputsExt {
		t.TxInputsExt = append(t.TxInputsExt, JsonTxInputExt{
			Bucket:    inputExt.Bucket,
			Key:       string(inputExt.Key),
			RefTxid:   hex.EncodeToString(inputExt.RefTxid),
			RefOffset: inputExt.RefOffset,
		})
	}
	for _, outputExt := range tx.TxOutputsExt {
		v := string(outputExt.Value)
		if len(v) > 30 {
			v = "value too long"
		}
		t.TxOutputsExt = append(t.TxOutputsExt, JsonTxOutputExt{
			Bucket: outputExt.Bucket,
			Key:    string(outputExt.Key),
			Value:  v,
		})
	}
	if tx.ContractRequests != nil {
		for i := 0; i < len(tx.ContractRequests); i++ {
			req := tx.ContractRequests[i]
			tmpReq := &JsonInvokeRequest{
				ModuleName:   req.ModuleName,
				ContractName: req.ContractName,
				MethodName:   req.MethodName,
				Args:         map[string]string{},
			}
			for argKey, argV := range req.Args {
				v := string(argV)
				if len(argV) > 30 {
					v = "value too long"
				}
				tmpReq.Args[argKey] = v
			}
			for _, rlimit := range req.ResourceLimits {
				resource := JsonResourceLimit{
					Type:  rlimit.Type.String(),
					Limit: rlimit.Limit,
				}
				tmpReq.ResouceLimits = append(tmpReq.ResouceLimits, resource)
			}
			t.ContractRequests = append(t.ContractRequests, tmpReq)
		}
	}

	t.AuthRequire = append(t.AuthRequire, tx.AuthRequire...)

	for _, initsign := range tx.InitiatorSigns {
		t.InitiatorSigns = append(t.InitiatorSigns, JsonSignatureInfo{
			PublicKey: initsign.PublicKey,
			Sign:      hex.EncodeToString(initsign.Sign),
		})
	}

	for _, authSign := range tx.AuthRequireSigns {
		t.AuthRequireSigns = append(t.AuthRequireSigns, JsonSignatureInfo{
			PublicKey: authSign.PublicKey,
			Sign:      hex.EncodeToString(authSign.Sign),
		})
	}

	//if tx.ModifyBlock != nil {
	//	t.ModifyBlock = JsonModifyBlock{
	//		EffectiveHeight: tx.ModifyBlock.EffectiveHeight,
	//		Marked:          tx.ModifyBlock.Marked,
	//		EffectiveTxid:   tx.ModifyBlock.EffectiveTxid,
	//	}
	//}
	return t
}

func GetTxsFromTransactions(txs []*Transaction) []*JsonTransaction {
	tempTxs := []*JsonTransaction{}
	for _, v := range txs {
		tx := GetTxFromTransaction(v)
		tempTxs = append(tempTxs, tx)
	}
	return tempTxs
}

type Json struct {
	Tx  *JsonTransaction
	Txs []*JsonTransaction
}

func (m *Json) Reset()         { *m = Json{} }
func (m *Json) String() string { return proto.CompactTextString(m) }
func (*Json) ProtoMessage()    {}

type JsonTxStatus struct {
	Bcname string `protobuf:"bytes,2,opt,name=bcname,proto3" json:"bcname,omitempty"`
	Txid   string `protobuf:"bytes,3,opt,name=txid,proto3" json:"txid,omitempty"`
}

func (m *JsonTxStatus) Reset()         { *m = JsonTxStatus{} }
func (m *JsonTxStatus) String() string { return proto.CompactTextString(m) }
func (*JsonTxStatus) ProtoMessage()    {}
