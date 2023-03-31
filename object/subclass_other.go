package object

import "encoding/json"

var OtherClassNames = []string{
	"HEAVY", "LAVA", "GATE", "VISIBLE_OBELISK", "INVISIBLE_OBELISK", "TECH", "LOTD",
	"USEABLE", "CHEST_NW", "CHEST_NE", "CHEST_SE", "CHEST_SW", "STONE_DOOR",
}

func (c SubClass) AsOther() OtherClass {
	return OtherClass(c)
}

func ParseOtherClass(s string) (OtherClass, error) {
	v, err := parseEnum("other class", s, OtherClassNames)
	return OtherClass(v), err
}

func ParseOtherClassSet(s string) (OtherClass, error) {
	v, err := parseEnumSet("other class", s, OtherClassNames)
	return OtherClass(v), err
}

var _ enum[OtherClass] = OtherClass(0)

type OtherClass uint32

const (
	OtherHeavy            = OtherClass(1 << iota) // 0x1
	OtherLava                                     // 0x2
	OtherGate                                     // 0x4
	OtherVisibleObelisk                           // 0x8
	OtherInvisibleObelisk                         // 0x10
	OtherTech                                     // 0x20
	OtherLOTD                                     // 0x40
	OtherUseable                                  // 0x80
	OtherChestNW                                  // 0x100
	OtherChestNE                                  // 0x200
	OtherChestSE                                  // 0x400
	OtherChestSW                                  // 0x800
	OtherStoneDoor                                // 0x1000
)

func (c OtherClass) Has(c2 OtherClass) bool {
	return c&c2 != 0
}

func (c OtherClass) HasAny(c2 OtherClass) bool {
	return c&c2 != 0
}

func (c OtherClass) Split() []OtherClass {
	return splitBits(c)
}

func (c OtherClass) String() string {
	return stringBits(uint32(c), OtherClassNames)
}

func (c OtherClass) MarshalJSON() ([]byte, error) {
	var arr []string
	for _, s := range c.Split() {
		arr = append(arr, s.String())
	}
	return json.Marshal(arr)
}
