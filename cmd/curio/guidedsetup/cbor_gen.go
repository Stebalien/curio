// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package guidedsetup

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"

	abi "github.com/filecoin-project/go-state-types/abi"

	sector "github.com/filecoin-project/curio/lib/types/sector"

	storiface "github.com/filecoin-project/lotus/storage/sealer/storiface"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *SectorInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{184, 36}); err != nil {
		return err
	}

	// t.CommD (cid.Cid) (struct)
	if len("CommD") > 8192 {
		return xerrors.Errorf("Value in field \"CommD\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("CommD"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("CommD")); err != nil {
		return err
	}

	if t.CommD == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.CommD); err != nil {
			return xerrors.Errorf("failed to write cid field t.CommD: %w", err)
		}
	}

	// t.CommR (cid.Cid) (struct)
	if len("CommR") > 8192 {
		return xerrors.Errorf("Value in field \"CommR\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("CommR"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("CommR")); err != nil {
		return err
	}

	if t.CommR == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.CommR); err != nil {
			return xerrors.Errorf("failed to write cid field t.CommR: %w", err)
		}
	}

	// t.Proof ([]uint8) (slice)
	if len("Proof") > 8192 {
		return xerrors.Errorf("Value in field \"Proof\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Proof"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Proof")); err != nil {
		return err
	}

	if len(t.Proof) > 2097152 {
		return xerrors.Errorf("Byte array in field t.Proof was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Proof))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Proof); err != nil {
		return err
	}

	// t.State (sector.SectorState) (string)
	if len("State") > 8192 {
		return xerrors.Errorf("Value in field \"State\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("State"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("State")); err != nil {
		return err
	}

	if len(t.State) > 8192 {
		return xerrors.Errorf("Value in field t.State was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.State))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.State)); err != nil {
		return err
	}

	// t.Pieces ([]sector.SafeSectorPiece) (slice)
	if len("Pieces") > 8192 {
		return xerrors.Errorf("Value in field \"Pieces\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("Pieces"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("Pieces")); err != nil {
		return err
	}

	if len(t.Pieces) > 8192 {
		return xerrors.Errorf("Slice value in field t.Pieces was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Pieces))); err != nil {
		return err
	}
	for _, v := range t.Pieces {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}

	}

	// t.LastErr (string) (string)
	if len("LastErr") > 8192 {
		return xerrors.Errorf("Value in field \"LastErr\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("LastErr"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("LastErr")); err != nil {
		return err
	}

	if len(t.LastErr) > 8192 {
		return xerrors.Errorf("Value in field t.LastErr was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.LastErr))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.LastErr)); err != nil {
		return err
	}

	// t.CCUpdate (bool) (bool)
	if len("CCUpdate") > 8192 {
		return xerrors.Errorf("Value in field \"CCUpdate\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("CCUpdate"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("CCUpdate")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.CCUpdate); err != nil {
		return err
	}

	// t.SeedEpoch (abi.ChainEpoch) (int64)
	if len("SeedEpoch") > 8192 {
		return xerrors.Errorf("Value in field \"SeedEpoch\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("SeedEpoch"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("SeedEpoch")); err != nil {
		return err
	}

	if t.SeedEpoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.SeedEpoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.SeedEpoch-1)); err != nil {
			return err
		}
	}

	// t.SeedValue (abi.InteractiveSealRandomness) (slice)
	if len("SeedValue") > 8192 {
		return xerrors.Errorf("Value in field \"SeedValue\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("SeedValue"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("SeedValue")); err != nil {
		return err
	}

	if len(t.SeedValue) > 2097152 {
		return xerrors.Errorf("Byte array in field t.SeedValue was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.SeedValue))); err != nil {
		return err
	}

	if _, err := cw.Write(t.SeedValue); err != nil {
		return err
	}

	// t.SectorType (abi.RegisteredSealProof) (int64)
	if len("SectorType") > 8192 {
		return xerrors.Errorf("Value in field \"SectorType\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("SectorType"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("SectorType")); err != nil {
		return err
	}

	if t.SectorType >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.SectorType)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.SectorType-1)); err != nil {
			return err
		}
	}

	// t.TicketEpoch (abi.ChainEpoch) (int64)
	if len("TicketEpoch") > 8192 {
		return xerrors.Errorf("Value in field \"TicketEpoch\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("TicketEpoch"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("TicketEpoch")); err != nil {
		return err
	}

	if t.TicketEpoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.TicketEpoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.TicketEpoch-1)); err != nil {
			return err
		}
	}

	// t.TicketValue (abi.SealRandomness) (slice)
	if len("TicketValue") > 8192 {
		return xerrors.Errorf("Value in field \"TicketValue\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("TicketValue"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("TicketValue")); err != nil {
		return err
	}

	if len(t.TicketValue) > 2097152 {
		return xerrors.Errorf("Byte array in field t.TicketValue was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.TicketValue))); err != nil {
		return err
	}

	if _, err := cw.Write(t.TicketValue); err != nil {
		return err
	}

	// t.CreationTime (int64) (int64)
	if len("CreationTime") > 8192 {
		return xerrors.Errorf("Value in field \"CreationTime\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("CreationTime"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("CreationTime")); err != nil {
		return err
	}

	if t.CreationTime >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.CreationTime)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.CreationTime-1)); err != nil {
			return err
		}
	}

	// t.SectorNumber (abi.SectorNumber) (uint64)
	if len("SectorNumber") > 8192 {
		return xerrors.Errorf("Value in field \"SectorNumber\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("SectorNumber"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("SectorNumber")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.SectorNumber)); err != nil {
		return err
	}

	// t.TerminatedAt (abi.ChainEpoch) (int64)
	if len("TerminatedAt") > 8192 {
		return xerrors.Errorf("Value in field \"TerminatedAt\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("TerminatedAt"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("TerminatedAt")); err != nil {
		return err
	}

	if t.TerminatedAt >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.TerminatedAt)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.TerminatedAt-1)); err != nil {
			return err
		}
	}

	// t.UpdateSealed (cid.Cid) (struct)
	if len("UpdateSealed") > 8192 {
		return xerrors.Errorf("Value in field \"UpdateSealed\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("UpdateSealed"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("UpdateSealed")); err != nil {
		return err
	}

	if t.UpdateSealed == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.UpdateSealed); err != nil {
			return xerrors.Errorf("failed to write cid field t.UpdateSealed: %w", err)
		}
	}

	// t.CommitMessage (cid.Cid) (struct)
	if len("CommitMessage") > 8192 {
		return xerrors.Errorf("Value in field \"CommitMessage\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("CommitMessage"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("CommitMessage")); err != nil {
		return err
	}

	if t.CommitMessage == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.CommitMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.CommitMessage: %w", err)
		}
	}

	// t.InvalidProofs (uint64) (uint64)
	if len("InvalidProofs") > 8192 {
		return xerrors.Errorf("Value in field \"InvalidProofs\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("InvalidProofs"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("InvalidProofs")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.InvalidProofs)); err != nil {
		return err
	}

	// t.PreCommit1Out (storiface.PreCommit1Out) (slice)
	if len("PreCommit1Out") > 8192 {
		return xerrors.Errorf("Value in field \"PreCommit1Out\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PreCommit1Out"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("PreCommit1Out")); err != nil {
		return err
	}

	if len(t.PreCommit1Out) > 2097152 {
		return xerrors.Errorf("Byte array in field t.PreCommit1Out was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.PreCommit1Out))); err != nil {
		return err
	}

	if _, err := cw.Write(t.PreCommit1Out); err != nil {
		return err
	}

	// t.FaultReportMsg (cid.Cid) (struct)
	if len("FaultReportMsg") > 8192 {
		return xerrors.Errorf("Value in field \"FaultReportMsg\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("FaultReportMsg"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("FaultReportMsg")); err != nil {
		return err
	}

	if t.FaultReportMsg == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.FaultReportMsg); err != nil {
			return xerrors.Errorf("failed to write cid field t.FaultReportMsg: %w", err)
		}
	}

	// t.UpdateUnsealed (cid.Cid) (struct)
	if len("UpdateUnsealed") > 8192 {
		return xerrors.Errorf("Value in field \"UpdateUnsealed\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("UpdateUnsealed"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("UpdateUnsealed")); err != nil {
		return err
	}

	if t.UpdateUnsealed == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.UpdateUnsealed); err != nil {
			return xerrors.Errorf("failed to write cid field t.UpdateUnsealed: %w", err)
		}
	}

	// t.PreCommit1Fails (uint64) (uint64)
	if len("PreCommit1Fails") > 8192 {
		return xerrors.Errorf("Value in field \"PreCommit1Fails\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PreCommit1Fails"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("PreCommit1Fails")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.PreCommit1Fails)); err != nil {
		return err
	}

	// t.PreCommit2Fails (uint64) (uint64)
	if len("PreCommit2Fails") > 8192 {
		return xerrors.Errorf("Value in field \"PreCommit2Fails\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PreCommit2Fails"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("PreCommit2Fails")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.PreCommit2Fails)); err != nil {
		return err
	}

	// t.PreCommitTipSet (types.TipSetKey) (struct)
	if len("PreCommitTipSet") > 8192 {
		return xerrors.Errorf("Value in field \"PreCommitTipSet\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PreCommitTipSet"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("PreCommitTipSet")); err != nil {
		return err
	}

	if err := t.PreCommitTipSet.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.RemoteDataCache (storiface.SectorLocation) (struct)
	if len("RemoteDataCache") > 8192 {
		return xerrors.Errorf("Value in field \"RemoteDataCache\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("RemoteDataCache"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("RemoteDataCache")); err != nil {
		return err
	}

	if err := t.RemoteDataCache.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.PreCommitDeposit (big.Int) (struct)
	if len("PreCommitDeposit") > 8192 {
		return xerrors.Errorf("Value in field \"PreCommitDeposit\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PreCommitDeposit"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("PreCommitDeposit")); err != nil {
		return err
	}

	if err := t.PreCommitDeposit.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.PreCommitMessage (cid.Cid) (struct)
	if len("PreCommitMessage") > 8192 {
		return xerrors.Errorf("Value in field \"PreCommitMessage\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("PreCommitMessage"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("PreCommitMessage")); err != nil {
		return err
	}

	if t.PreCommitMessage == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.PreCommitMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.PreCommitMessage: %w", err)
		}
	}

	// t.RemoteDataSealed (storiface.SectorLocation) (struct)
	if len("RemoteDataSealed") > 8192 {
		return xerrors.Errorf("Value in field \"RemoteDataSealed\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("RemoteDataSealed"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("RemoteDataSealed")); err != nil {
		return err
	}

	if err := t.RemoteDataSealed.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.TerminateMessage (cid.Cid) (struct)
	if len("TerminateMessage") > 8192 {
		return xerrors.Errorf("Value in field \"TerminateMessage\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("TerminateMessage"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("TerminateMessage")); err != nil {
		return err
	}

	if t.TerminateMessage == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.TerminateMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.TerminateMessage: %w", err)
		}
	}

	// t.RemoteDataUnsealed (storiface.SectorLocation) (struct)
	if len("RemoteDataUnsealed") > 8192 {
		return xerrors.Errorf("Value in field \"RemoteDataUnsealed\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("RemoteDataUnsealed"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("RemoteDataUnsealed")); err != nil {
		return err
	}

	if err := t.RemoteDataUnsealed.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ReplicaUpdateProof (storiface.ReplicaUpdateProof) (slice)
	if len("ReplicaUpdateProof") > 8192 {
		return xerrors.Errorf("Value in field \"ReplicaUpdateProof\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ReplicaUpdateProof"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("ReplicaUpdateProof")); err != nil {
		return err
	}

	if len(t.ReplicaUpdateProof) > 2097152 {
		return xerrors.Errorf("Byte array in field t.ReplicaUpdateProof was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.ReplicaUpdateProof))); err != nil {
		return err
	}

	if _, err := cw.Write(t.ReplicaUpdateProof); err != nil {
		return err
	}

	// t.RemoteDataFinalized (bool) (bool)
	if len("RemoteDataFinalized") > 8192 {
		return xerrors.Errorf("Value in field \"RemoteDataFinalized\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("RemoteDataFinalized"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("RemoteDataFinalized")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.RemoteDataFinalized); err != nil {
		return err
	}

	// t.ReplicaUpdateMessage (cid.Cid) (struct)
	if len("ReplicaUpdateMessage") > 8192 {
		return xerrors.Errorf("Value in field \"ReplicaUpdateMessage\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ReplicaUpdateMessage"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("ReplicaUpdateMessage")); err != nil {
		return err
	}

	if t.ReplicaUpdateMessage == nil {
		if _, err := cw.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(cw, *t.ReplicaUpdateMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.ReplicaUpdateMessage: %w", err)
		}
	}

	// t.RemoteCommit1Endpoint (string) (string)
	if len("RemoteCommit1Endpoint") > 8192 {
		return xerrors.Errorf("Value in field \"RemoteCommit1Endpoint\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("RemoteCommit1Endpoint"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("RemoteCommit1Endpoint")); err != nil {
		return err
	}

	if len(t.RemoteCommit1Endpoint) > 8192 {
		return xerrors.Errorf("Value in field t.RemoteCommit1Endpoint was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.RemoteCommit1Endpoint))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.RemoteCommit1Endpoint)); err != nil {
		return err
	}

	// t.RemoteCommit2Endpoint (string) (string)
	if len("RemoteCommit2Endpoint") > 8192 {
		return xerrors.Errorf("Value in field \"RemoteCommit2Endpoint\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("RemoteCommit2Endpoint"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("RemoteCommit2Endpoint")); err != nil {
		return err
	}

	if len(t.RemoteCommit2Endpoint) > 8192 {
		return xerrors.Errorf("Value in field t.RemoteCommit2Endpoint was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.RemoteCommit2Endpoint))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.RemoteCommit2Endpoint)); err != nil {
		return err
	}

	// t.RemoteSealingDoneEndpoint (string) (string)
	if len("RemoteSealingDoneEndpoint") > 8192 {
		return xerrors.Errorf("Value in field \"RemoteSealingDoneEndpoint\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("RemoteSealingDoneEndpoint"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("RemoteSealingDoneEndpoint")); err != nil {
		return err
	}

	if len(t.RemoteSealingDoneEndpoint) > 8192 {
		return xerrors.Errorf("Value in field t.RemoteSealingDoneEndpoint was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.RemoteSealingDoneEndpoint))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.RemoteSealingDoneEndpoint)); err != nil {
		return err
	}
	return nil
}

func (t *SectorInfo) UnmarshalCBOR(r io.Reader) (err error) {
	*t = SectorInfo{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("SectorInfo: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringWithMax(cr, 8192)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.CommD (cid.Cid) (struct)
		case "CommD":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.CommD: %w", err)
					}

					t.CommD = &c
				}

			}
			// t.CommR (cid.Cid) (struct)
		case "CommR":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.CommR: %w", err)
					}

					t.CommR = &c
				}

			}
			// t.Proof ([]uint8) (slice)
		case "Proof":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > 2097152 {
				return fmt.Errorf("t.Proof: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Proof = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.Proof); err != nil {
				return err
			}

			// t.State (sector.SectorState) (string)
		case "State":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.State = sector.SectorState(sval)
			}
			// t.Pieces ([]sector.SafeSectorPiece) (slice)
		case "Pieces":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > 8192 {
				return fmt.Errorf("t.Pieces: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}

			if extra > 0 {
				t.Pieces = make([]sector.SafeSectorPiece, extra)
			}

			for i := 0; i < int(extra); i++ {
				{
					var maj byte
					var extra uint64
					var err error
					_ = maj
					_ = extra
					_ = err

					{

						if err := t.Pieces[i].UnmarshalCBOR(cr); err != nil {
							return xerrors.Errorf("unmarshaling t.Pieces[i]: %w", err)
						}

					}

				}
			}
			// t.LastErr (string) (string)
		case "LastErr":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.LastErr = string(sval)
			}
			// t.CCUpdate (bool) (bool)
		case "CCUpdate":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.CCUpdate = false
			case 21:
				t.CCUpdate = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.SeedEpoch (abi.ChainEpoch) (int64)
		case "SeedEpoch":
			{
				maj, extra, err := cr.ReadHeader()
				if err != nil {
					return err
				}
				var extraI int64
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.SeedEpoch = abi.ChainEpoch(extraI)
			}
			// t.SeedValue (abi.InteractiveSealRandomness) (slice)
		case "SeedValue":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > 2097152 {
				return fmt.Errorf("t.SeedValue: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.SeedValue = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.SeedValue); err != nil {
				return err
			}

			// t.SectorType (abi.RegisteredSealProof) (int64)
		case "SectorType":
			{
				maj, extra, err := cr.ReadHeader()
				if err != nil {
					return err
				}
				var extraI int64
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.SectorType = abi.RegisteredSealProof(extraI)
			}
			// t.TicketEpoch (abi.ChainEpoch) (int64)
		case "TicketEpoch":
			{
				maj, extra, err := cr.ReadHeader()
				if err != nil {
					return err
				}
				var extraI int64
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.TicketEpoch = abi.ChainEpoch(extraI)
			}
			// t.TicketValue (abi.SealRandomness) (slice)
		case "TicketValue":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > 2097152 {
				return fmt.Errorf("t.TicketValue: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.TicketValue = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.TicketValue); err != nil {
				return err
			}

			// t.CreationTime (int64) (int64)
		case "CreationTime":
			{
				maj, extra, err := cr.ReadHeader()
				if err != nil {
					return err
				}
				var extraI int64
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.CreationTime = int64(extraI)
			}
			// t.SectorNumber (abi.SectorNumber) (uint64)
		case "SectorNumber":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.SectorNumber = abi.SectorNumber(extra)

			}
			// t.TerminatedAt (abi.ChainEpoch) (int64)
		case "TerminatedAt":
			{
				maj, extra, err := cr.ReadHeader()
				if err != nil {
					return err
				}
				var extraI int64
				switch maj {
				case cbg.MajUnsignedInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 positive overflow")
					}
				case cbg.MajNegativeInt:
					extraI = int64(extra)
					if extraI < 0 {
						return fmt.Errorf("int64 negative overflow")
					}
					extraI = -1 - extraI
				default:
					return fmt.Errorf("wrong type for int64 field: %d", maj)
				}

				t.TerminatedAt = abi.ChainEpoch(extraI)
			}
			// t.UpdateSealed (cid.Cid) (struct)
		case "UpdateSealed":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.UpdateSealed: %w", err)
					}

					t.UpdateSealed = &c
				}

			}
			// t.CommitMessage (cid.Cid) (struct)
		case "CommitMessage":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.CommitMessage: %w", err)
					}

					t.CommitMessage = &c
				}

			}
			// t.InvalidProofs (uint64) (uint64)
		case "InvalidProofs":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.InvalidProofs = uint64(extra)

			}
			// t.PreCommit1Out (storiface.PreCommit1Out) (slice)
		case "PreCommit1Out":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > 2097152 {
				return fmt.Errorf("t.PreCommit1Out: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.PreCommit1Out = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.PreCommit1Out); err != nil {
				return err
			}

			// t.FaultReportMsg (cid.Cid) (struct)
		case "FaultReportMsg":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.FaultReportMsg: %w", err)
					}

					t.FaultReportMsg = &c
				}

			}
			// t.UpdateUnsealed (cid.Cid) (struct)
		case "UpdateUnsealed":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.UpdateUnsealed: %w", err)
					}

					t.UpdateUnsealed = &c
				}

			}
			// t.PreCommit1Fails (uint64) (uint64)
		case "PreCommit1Fails":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.PreCommit1Fails = uint64(extra)

			}
			// t.PreCommit2Fails (uint64) (uint64)
		case "PreCommit2Fails":

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.PreCommit2Fails = uint64(extra)

			}
			// t.PreCommitTipSet (types.TipSetKey) (struct)
		case "PreCommitTipSet":

			{

				if err := t.PreCommitTipSet.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.PreCommitTipSet: %w", err)
				}

			}
			// t.RemoteDataCache (storiface.SectorLocation) (struct)
		case "RemoteDataCache":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.RemoteDataCache = new(storiface.SectorLocation)
					if err := t.RemoteDataCache.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.RemoteDataCache pointer: %w", err)
					}
				}

			}
			// t.PreCommitDeposit (big.Int) (struct)
		case "PreCommitDeposit":

			{

				if err := t.PreCommitDeposit.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.PreCommitDeposit: %w", err)
				}

			}
			// t.PreCommitMessage (cid.Cid) (struct)
		case "PreCommitMessage":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.PreCommitMessage: %w", err)
					}

					t.PreCommitMessage = &c
				}

			}
			// t.RemoteDataSealed (storiface.SectorLocation) (struct)
		case "RemoteDataSealed":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.RemoteDataSealed = new(storiface.SectorLocation)
					if err := t.RemoteDataSealed.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.RemoteDataSealed pointer: %w", err)
					}
				}

			}
			// t.TerminateMessage (cid.Cid) (struct)
		case "TerminateMessage":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.TerminateMessage: %w", err)
					}

					t.TerminateMessage = &c
				}

			}
			// t.RemoteDataUnsealed (storiface.SectorLocation) (struct)
		case "RemoteDataUnsealed":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.RemoteDataUnsealed = new(storiface.SectorLocation)
					if err := t.RemoteDataUnsealed.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.RemoteDataUnsealed pointer: %w", err)
					}
				}

			}
			// t.ReplicaUpdateProof (storiface.ReplicaUpdateProof) (slice)
		case "ReplicaUpdateProof":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > 2097152 {
				return fmt.Errorf("t.ReplicaUpdateProof: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.ReplicaUpdateProof = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.ReplicaUpdateProof); err != nil {
				return err
			}

			// t.RemoteDataFinalized (bool) (bool)
		case "RemoteDataFinalized":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.RemoteDataFinalized = false
			case 21:
				t.RemoteDataFinalized = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.ReplicaUpdateMessage (cid.Cid) (struct)
		case "ReplicaUpdateMessage":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(cr)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.ReplicaUpdateMessage: %w", err)
					}

					t.ReplicaUpdateMessage = &c
				}

			}
			// t.RemoteCommit1Endpoint (string) (string)
		case "RemoteCommit1Endpoint":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.RemoteCommit1Endpoint = string(sval)
			}
			// t.RemoteCommit2Endpoint (string) (string)
		case "RemoteCommit2Endpoint":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.RemoteCommit2Endpoint = string(sval)
			}
			// t.RemoteSealingDoneEndpoint (string) (string)
		case "RemoteSealingDoneEndpoint":

			{
				sval, err := cbg.ReadStringWithMax(cr, 8192)
				if err != nil {
					return err
				}

				t.RemoteSealingDoneEndpoint = string(sval)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}