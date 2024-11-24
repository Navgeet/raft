package raft

import (
	"encoding/binary"
	"io"

	pb "github.com/navgeet/raft/internal/protobuf"
	"google.golang.org/protobuf/proto"
)

func encodeSnapshot(w io.Writer, snapshot *Snapshot) error {
	pbSnapshot := &pb.Snapshot{LastIncludedIndex: snapshot.LastIncludedIndex, LastIncludedTerm: snapshot.LastIncludedTerm, Data: snapshot.Data}
	buf, err := proto.Marshal(pbSnapshot)
	if err != nil {
		return err
	}
	size := int32(len(buf))
	if err := binary.Write(w, binary.BigEndian, size); err != nil {
		return err
	}
	if _, err := w.Write(buf); err != nil {
		return err
	}
	return nil
}

func decodeSnapshot(r io.Reader) (Snapshot, error) {
	var size int32
	if err := binary.Read(r, binary.BigEndian, &size); err != nil {
		return Snapshot{}, err
	}

	buf := make([]byte, size)
	if _, err := io.ReadFull(r, buf); err != nil {
		return Snapshot{}, err
	}

	pbSnapshot := &pb.Snapshot{}
	if err := proto.Unmarshal(buf, pbSnapshot); err != nil {
		return Snapshot{}, err
	}

	snapshot := Snapshot{
		LastIncludedIndex: pbSnapshot.GetLastIncludedIndex(),
		LastIncludedTerm:  pbSnapshot.GetLastIncludedTerm(),
		Data:              pbSnapshot.GetData(),
	}

	return snapshot, nil
}

func encodePersistentState(w io.Writer, persistentState *PersistentState) error {
	pbState := &pb.StorageState{Term: persistentState.Term, VotedFor: persistentState.VotedFor}
	buf, err := proto.Marshal(pbState)
	if err != nil {
		return err
	}
	size := int32(len(buf))
	if err := binary.Write(w, binary.BigEndian, size); err != nil {
		return err
	}
	if _, err := w.Write(buf); err != nil {
		return err
	}
	return nil
}

func decodePersistentState(r io.Reader) (PersistentState, error) {
	var size int32
	if err := binary.Read(r, binary.BigEndian, &size); err != nil {
		return PersistentState{}, err
	}

	buf := make([]byte, size)
	if _, err := io.ReadFull(r, buf); err != nil {
		return PersistentState{}, err
	}

	pbPersistentState := &pb.StorageState{}
	if err := proto.Unmarshal(buf, pbPersistentState); err != nil {
		return PersistentState{}, err
	}

	persistentState := PersistentState{
		Term:     pbPersistentState.GetTerm(),
		VotedFor: pbPersistentState.GetVotedFor(),
	}

	return persistentState, nil
}

func encodeLogEntry(w io.Writer, entry *LogEntry) error {
	pbEntry := &pb.LogEntry{
		Index:  entry.Index,
		Term:   entry.Term,
		Data:   entry.Data,
		Offset: entry.Offset,
	}

	buf, err := proto.Marshal(pbEntry)
	if err != nil {
		return err
	}

	size := int32(len(buf))
	if err := binary.Write(w, binary.BigEndian, size); err != nil {
		return err
	}

	if _, err := w.Write(buf); err != nil {
		return err
	}

	return nil
}

func decodeLogEntry(r io.Reader) (LogEntry, error) {
	var size int32
	if err := binary.Read(r, binary.BigEndian, &size); err != nil {
		return LogEntry{}, err
	}

	buf := make([]byte, size)
	if _, err := io.ReadFull(r, buf); err != nil {
		return LogEntry{}, err
	}

	pbEntry := &pb.LogEntry{}
	if err := proto.Unmarshal(buf, pbEntry); err != nil {
		return LogEntry{}, err
	}

	entry := LogEntry{
		Index:  pbEntry.GetIndex(),
		Term:   pbEntry.GetTerm(),
		Data:   pbEntry.GetData(),
		Offset: pbEntry.GetOffset(),
	}

	return entry, nil
}
