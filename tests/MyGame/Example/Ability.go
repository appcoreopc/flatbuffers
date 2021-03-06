// automatically generated by the FlatBuffers compiler, do not modify

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Ability struct {
	_tab flatbuffers.Struct
}

func (rcv *Ability) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Ability) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Ability) Id() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Ability) MutateId(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Ability) Distance() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Ability) MutateDistance(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func CreateAbility(builder *flatbuffers.Builder, id uint32, distance uint32) flatbuffers.UOffsetT {
	builder.Prep(4, 8)
	builder.PrependUint32(distance)
	builder.PrependUint32(id)
	return builder.Offset()
}
