# automatically generated by the FlatBuffers compiler, do not modify

# namespace: MyGame

import flatbuffers
from flatbuffers.compat import import_numpy
np = import_numpy()

class Monster(object):
    __slots__ = ['_tab']

    @classmethod
    def GetRootAs(cls, buf, offset=0):
        n = flatbuffers.encode.Get(flatbuffers.packer.uoffset, buf, offset)
        x = Monster()
        x.Init(buf, n + offset)
        return x

    @classmethod
    def GetRootAsMonster(cls, buf, offset=0):
        """This method is deprecated. Please switch to GetRootAs."""
        return cls.GetRootAs(buf, offset)
    # Monster
    def Init(self, buf, pos):
        self._tab = flatbuffers.table.Table(buf, pos)

    # Monster
    def Pos(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(4))
        if o != 0:
            x = o + self._tab.Pos
            from MyGame.Vec3 import Vec3
            obj = Vec3()
            obj.Init(self._tab.Bytes, x)
            return obj
        return None

    # Monster
    def Mana(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(6))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Int16Flags, o + self._tab.Pos)
        return 150

    # Monster
    def Hp(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(8))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Int16Flags, o + self._tab.Pos)
        return 100

    # Monster
    def Name(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(10))
        if o != 0:
            return self._tab.String(o + self._tab.Pos)
        return None

    # Monster
    def Inventory(self, j):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(14))
        if o != 0:
            a = self._tab.Vector(o)
            return self._tab.Get(flatbuffers.number_types.Uint8Flags, a + flatbuffers.number_types.UOffsetTFlags.py_type(j * 1))
        return 0

    # Monster
    def InventoryAsNumpy(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(14))
        if o != 0:
            return self._tab.GetVectorAsNumpy(flatbuffers.number_types.Uint8Flags, o)
        return 0

    # Monster
    def InventoryLength(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(14))
        if o != 0:
            return self._tab.VectorLen(o)
        return 0

    # Monster
    def InventoryIsNone(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(14))
        return o == 0

    # Monster
    def Color(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(16))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Int8Flags, o + self._tab.Pos)
        return 3

    # Monster
    def TestType(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(18))
        if o != 0:
            return self._tab.Get(flatbuffers.number_types.Uint8Flags, o + self._tab.Pos)
        return 0

    # Monster
    def Test(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(20))
        if o != 0:
            from flatbuffers.table import Table
            obj = Table(bytearray(), 0)
            self._tab.Union(obj, o)
            return obj
        return None

def Start(builder): builder.StartObject(9)
def MonsterStart(builder):
    """This method is deprecated. Please switch to Start."""
    return Start(builder)
def AddPos(builder, pos): builder.PrependStructSlot(0, flatbuffers.number_types.UOffsetTFlags.py_type(pos), 0)
def MonsterAddPos(builder, pos):
    """This method is deprecated. Please switch to AddPos."""
    return AddPos(builder, pos)
def AddMana(builder, mana): builder.PrependInt16Slot(1, mana, 150)
def MonsterAddMana(builder, mana):
    """This method is deprecated. Please switch to AddMana."""
    return AddMana(builder, mana)
def AddHp(builder, hp): builder.PrependInt16Slot(2, hp, 100)
def MonsterAddHp(builder, hp):
    """This method is deprecated. Please switch to AddHp."""
    return AddHp(builder, hp)
def AddName(builder, name): builder.PrependUOffsetTRelativeSlot(3, flatbuffers.number_types.UOffsetTFlags.py_type(name), 0)
def MonsterAddName(builder, name):
    """This method is deprecated. Please switch to AddName."""
    return AddName(builder, name)
def AddInventory(builder, inventory): builder.PrependUOffsetTRelativeSlot(5, flatbuffers.number_types.UOffsetTFlags.py_type(inventory), 0)
def MonsterAddInventory(builder, inventory):
    """This method is deprecated. Please switch to AddInventory."""
    return AddInventory(builder, inventory)
def StartInventoryVector(builder, numElems): return builder.StartVector(1, numElems, 1)
def MonsterStartInventoryVector(builder, numElems):
    """This method is deprecated. Please switch to Start."""
    return StartInventoryVector(builder, numElems)
def AddColor(builder, color): builder.PrependInt8Slot(6, color, 3)
def MonsterAddColor(builder, color):
    """This method is deprecated. Please switch to AddColor."""
    return AddColor(builder, color)
def AddTestType(builder, testType): builder.PrependUint8Slot(7, testType, 0)
def MonsterAddTestType(builder, testType):
    """This method is deprecated. Please switch to AddTestType."""
    return AddTestType(builder, testType)
def AddTest(builder, test): builder.PrependUOffsetTRelativeSlot(8, flatbuffers.number_types.UOffsetTFlags.py_type(test), 0)
def MonsterAddTest(builder, test):
    """This method is deprecated. Please switch to AddTest."""
    return AddTest(builder, test)
def End(builder): return builder.EndObject()
def MonsterEnd(builder):
    """This method is deprecated. Please switch to End."""
    return End(builder)