// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package HTTPReqCall

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Resp struct {
	_tab flatbuffers.Table
}

func GetRootAsResp(buf []byte, offset flatbuffers.UOffsetT) *Resp {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Resp{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsResp(buf []byte, offset flatbuffers.UOffsetT) *Resp {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Resp{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Resp) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Resp) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Resp) Id() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Resp) MutateId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Resp) ActionType() Action {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return Action(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Resp) MutateActionType(n Action) bool {
	return rcv._tab.MutateByteSlot(6, byte(n))
}

func (rcv *Resp) Action(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func RespStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func RespAddId(builder *flatbuffers.Builder, id uint32) {
	builder.PrependUint32Slot(0, id, 0)
}
func RespAddActionType(builder *flatbuffers.Builder, actionType Action) {
	builder.PrependByteSlot(1, byte(actionType), 0)
}
func RespAddAction(builder *flatbuffers.Builder, action flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(action), 0)
}
func RespEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}