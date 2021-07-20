# automatically generated by the FlatBuffers compiler, do not modify

# namespace: A6

import flatbuffers
from flatbuffers.compat import import_numpy
np = import_numpy()

class TextEntry(object):
    __slots__ = ['_tab']

    @classmethod
    def GetRootAs(cls, buf, offset=0):
        n = flatbuffers.encode.Get(flatbuffers.packer.uoffset, buf, offset)
        x = TextEntry()
        x.Init(buf, n + offset)
        return x

    @classmethod
    def GetRootAsTextEntry(cls, buf, offset=0):
        """This method is deprecated. Please switch to GetRootAs."""
        return cls.GetRootAs(buf, offset)
    # TextEntry
    def Init(self, buf, pos):
        self._tab = flatbuffers.table.Table(buf, pos)

    # TextEntry
    def Name(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(4))
        if o != 0:
            return self._tab.String(o + self._tab.Pos)
        return None

    # TextEntry
    def Value(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(6))
        if o != 0:
            return self._tab.String(o + self._tab.Pos)
        return None

def Start(builder): builder.StartObject(2)
def TextEntryStart(builder):
    """This method is deprecated. Please switch to Start."""
    return Start(builder)
def AddName(builder, name): builder.PrependUOffsetTRelativeSlot(0, flatbuffers.number_types.UOffsetTFlags.py_type(name), 0)
def TextEntryAddName(builder, name):
    """This method is deprecated. Please switch to AddName."""
    return AddName(builder, name)
def AddValue(builder, value): builder.PrependUOffsetTRelativeSlot(1, flatbuffers.number_types.UOffsetTFlags.py_type(value), 0)
def TextEntryAddValue(builder, value):
    """This method is deprecated. Please switch to AddValue."""
    return AddValue(builder, value)
def End(builder): return builder.EndObject()
def TextEntryEnd(builder):
    """This method is deprecated. Please switch to End."""
    return End(builder)