// Code generated by 'yaegi extract github.com/noxworld-dev/opennox-lib/object'. DO NOT EDIT.

package imports

import (
	"reflect"

	"github.com/noxworld-dev/opennox-lib/object"
)

func init() {
	Symbols["github.com/noxworld-dev/opennox-lib/object/object"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ArmorArmArmor":                reflect.ValueOf(object.ArmorArmArmor),
		"ArmorBack":                    reflect.ValueOf(object.ArmorBack),
		"ArmorBoots":                   reflect.ValueOf(object.ArmorBoots),
		"ArmorBreastplate":             reflect.ValueOf(object.ArmorBreastplate),
		"ArmorClassNames":              reflect.ValueOf(&object.ArmorClassNames).Elem(),
		"ArmorHelmet":                  reflect.ValueOf(object.ArmorHelmet),
		"ArmorLegArmor":                reflect.ValueOf(object.ArmorLegArmor),
		"ArmorPants":                   reflect.ValueOf(object.ArmorPants),
		"ArmorShield":                  reflect.ValueOf(object.ArmorShield),
		"ArmorShirt":                   reflect.ValueOf(object.ArmorShirt),
		"BookAbility":                  reflect.ValueOf(object.BookAbility),
		"BookClassNames":               reflect.ValueOf(&object.BookClassNames).Elem(),
		"BookFieldGuide":               reflect.ValueOf(object.BookFieldGuide),
		"BookSpell":                    reflect.ValueOf(object.BookSpell),
		"ClassArmor":                   reflect.ValueOf(object.ClassArmor),
		"ClassClientPersist":           reflect.ValueOf(object.ClassClientPersist),
		"ClassClientPredict":           reflect.ValueOf(object.ClassClientPredict),
		"ClassComplex":                 reflect.ValueOf(object.ClassComplex),
		"ClassDangerous":               reflect.ValueOf(object.ClassDangerous),
		"ClassDoor":                    reflect.ValueOf(object.ClassDoor),
		"ClassElevator":                reflect.ValueOf(object.ClassElevator),
		"ClassElevatorShaft":           reflect.ValueOf(object.ClassElevatorShaft),
		"ClassExit":                    reflect.ValueOf(object.ClassExit),
		"ClassFire":                    reflect.ValueOf(object.ClassFire),
		"ClassFlag":                    reflect.ValueOf(object.ClassFlag),
		"ClassFood":                    reflect.ValueOf(object.ClassFood),
		"ClassHole":                    reflect.ValueOf(object.ClassHole),
		"ClassImmobile":                reflect.ValueOf(object.ClassImmobile),
		"ClassInfoBook":                reflect.ValueOf(object.ClassInfoBook),
		"ClassKey":                     reflect.ValueOf(object.ClassKey),
		"ClassLight":                   reflect.ValueOf(object.ClassLight),
		"ClassMissile":                 reflect.ValueOf(object.ClassMissile),
		"ClassMonster":                 reflect.ValueOf(object.ClassMonster),
		"ClassMonsterGenerator":        reflect.ValueOf(object.ClassMonsterGenerator),
		"ClassNames":                   reflect.ValueOf(&object.ClassNames).Elem(),
		"ClassNotStackable":            reflect.ValueOf(object.ClassNotStackable),
		"ClassObstacle":                reflect.ValueOf(object.ClassObstacle),
		"ClassPickup":                  reflect.ValueOf(object.ClassPickup),
		"ClassPlayer":                  reflect.ValueOf(object.ClassPlayer),
		"ClassReadable":                reflect.ValueOf(object.ClassReadable),
		"ClassSimple":                  reflect.ValueOf(object.ClassSimple),
		"ClassTransporter":             reflect.ValueOf(object.ClassTransporter),
		"ClassTreasure":                reflect.ValueOf(object.ClassTreasure),
		"ClassTrigger":                 reflect.ValueOf(object.ClassTrigger),
		"ClassVisibleEnable":           reflect.ValueOf(object.ClassVisibleEnable),
		"ClassWand":                    reflect.ValueOf(object.ClassWand),
		"ClassWeapon":                  reflect.ValueOf(object.ClassWeapon),
		"ExitClassNames":               reflect.ValueOf(&object.ExitClassNames).Elem(),
		"ExitQuest":                    reflect.ValueOf(object.ExitQuest),
		"ExitQuestWarp":                reflect.ValueOf(object.ExitQuestWarp),
		"FlagActive":                   reflect.ValueOf(object.FlagActive),
		"FlagAirborne":                 reflect.ValueOf(object.FlagAirborne),
		"FlagAllowOverlap":             reflect.ValueOf(object.FlagAllowOverlap),
		"FlagBelow":                    reflect.ValueOf(object.FlagBelow),
		"FlagBouncy":                   reflect.ValueOf(object.FlagBouncy),
		"FlagDead":                     reflect.ValueOf(object.FlagDead),
		"FlagDestroyed":                reflect.ValueOf(object.FlagDestroyed),
		"FlagEditVisible":              reflect.ValueOf(object.FlagEditVisible),
		"FlagEnabled":                  reflect.ValueOf(object.FlagEnabled),
		"FlagEquipped":                 reflect.ValueOf(object.FlagEquipped),
		"FlagFalling":                  reflect.ValueOf(object.FlagFalling),
		"FlagFlicker":                  reflect.ValueOf(object.FlagFlicker),
		"FlagInHole":                   reflect.ValueOf(object.FlagInHole),
		"FlagMarked":                   reflect.ValueOf(object.FlagMarked),
		"FlagMissileHit":               reflect.ValueOf(object.FlagMissileHit),
		"FlagNoAutoDrop":               reflect.ValueOf(object.FlagNoAutoDrop),
		"FlagNoCollide":                reflect.ValueOf(object.FlagNoCollide),
		"FlagNoCollideOwner":           reflect.ValueOf(object.FlagNoCollideOwner),
		"FlagNoPushCharacters":         reflect.ValueOf(object.FlagNoPushCharacters),
		"FlagNoUpdate":                 reflect.ValueOf(object.FlagNoUpdate),
		"FlagNoUpdateMask":             reflect.ValueOf(object.FlagNoUpdateMask),
		"FlagNotMovableMask":           reflect.ValueOf(object.FlagNotMovableMask),
		"FlagOnObject":                 reflect.ValueOf(object.FlagOnObject),
		"FlagOwnerVisible":             reflect.ValueOf(object.FlagOwnerVisible),
		"FlagPartitioned":              reflect.ValueOf(object.FlagPartitioned),
		"FlagPending":                  reflect.ValueOf(object.FlagPending),
		"FlagRespawn":                  reflect.ValueOf(object.FlagRespawn),
		"FlagSelected":                 reflect.ValueOf(object.FlagSelected),
		"FlagShadow":                   reflect.ValueOf(object.FlagShadow),
		"FlagShort":                    reflect.ValueOf(object.FlagShort),
		"FlagSightDestroy":             reflect.ValueOf(object.FlagSightDestroy),
		"FlagStill":                    reflect.ValueOf(object.FlagStill),
		"FlagTransient":                reflect.ValueOf(object.FlagTransient),
		"FlagTranslucent":              reflect.ValueOf(object.FlagTranslucent),
		"FlagsNames":                   reflect.ValueOf(&object.FlagsNames).Elem(),
		"FoodApple":                    reflect.ValueOf(object.FoodApple),
		"FoodClassNames":               reflect.ValueOf(&object.FoodClassNames).Elem(),
		"FoodCurePoisonPotion":         reflect.ValueOf(object.FoodCurePoisonPotion),
		"FoodFireProtectPotion":        reflect.ValueOf(object.FoodFireProtectPotion),
		"FoodHastePotion":              reflect.ValueOf(object.FoodHastePotion),
		"FoodHealthPotion":             reflect.ValueOf(object.FoodHealthPotion),
		"FoodInfravisionPotion":        reflect.ValueOf(object.FoodInfravisionPotion),
		"FoodInvisibilityPotion":       reflect.ValueOf(object.FoodInvisibilityPotion),
		"FoodInvulnerabilityPotion":    reflect.ValueOf(object.FoodInvulnerabilityPotion),
		"FoodJug":                      reflect.ValueOf(object.FoodJug),
		"FoodManaPotion":               reflect.ValueOf(object.FoodManaPotion),
		"FoodMushroom":                 reflect.ValueOf(object.FoodMushroom),
		"FoodPoisonProtectPotion":      reflect.ValueOf(object.FoodPoisonProtectPotion),
		"FoodPotion":                   reflect.ValueOf(object.FoodPotion),
		"FoodShieldPotion":             reflect.ValueOf(object.FoodShieldPotion),
		"FoodShockProtectPotion":       reflect.ValueOf(object.FoodShockProtectPotion),
		"FoodSimple":                   reflect.ValueOf(object.FoodSimple),
		"FoodVampirismPotion":          reflect.ValueOf(object.FoodVampirismPotion),
		"GeneratorClassNames":          reflect.ValueOf(&object.GeneratorClassNames).Elem(),
		"GeneratorNE":                  reflect.ValueOf(object.GeneratorNE),
		"GeneratorNW":                  reflect.ValueOf(object.GeneratorNW),
		"GeneratorSE":                  reflect.ValueOf(object.GeneratorSE),
		"GeneratorSW":                  reflect.ValueOf(object.GeneratorSW),
		"MaskTargets":                  reflect.ValueOf(object.MaskTargets),
		"MaskUnits":                    reflect.ValueOf(object.MaskUnits),
		"MaterialAnimalHide":           reflect.ValueOf(object.MaterialAnimalHide),
		"MaterialCloth":                reflect.ValueOf(object.MaterialCloth),
		"MaterialDiamond":              reflect.ValueOf(object.MaterialDiamond),
		"MaterialEarth":                reflect.ValueOf(object.MaterialEarth),
		"MaterialFlesh":                reflect.ValueOf(object.MaterialFlesh),
		"MaterialGlass":                reflect.ValueOf(object.MaterialGlass),
		"MaterialLiquid":               reflect.ValueOf(object.MaterialLiquid),
		"MaterialMagic":                reflect.ValueOf(object.MaterialMagic),
		"MaterialMetal":                reflect.ValueOf(object.MaterialMetal),
		"MaterialMud":                  reflect.ValueOf(object.MaterialMud),
		"MaterialNames":                reflect.ValueOf(&object.MaterialNames).Elem(),
		"MaterialNone":                 reflect.ValueOf(object.MaterialNone),
		"MaterialPaper":                reflect.ValueOf(object.MaterialPaper),
		"MaterialSnow":                 reflect.ValueOf(object.MaterialSnow),
		"MaterialStone":                reflect.ValueOf(object.MaterialStone),
		"MaterialWood":                 reflect.ValueOf(object.MaterialWood),
		"MissileClassNames":            reflect.ValueOf(&object.MissileClassNames).Elem(),
		"MissileMagic":                 reflect.ValueOf(object.MissileMagic),
		"MissileMissileCounterSpell":   reflect.ValueOf(object.MissileMissileCounterSpell),
		"MonsterBomber":                reflect.ValueOf(object.MonsterBomber),
		"MonsterClassNames":            reflect.ValueOf(&object.MonsterClassNames).Elem(),
		"MonsterFemaleNPC":             reflect.ValueOf(object.MonsterFemaleNPC),
		"MonsterHasSoul":               reflect.ValueOf(object.MonsterHasSoul),
		"MonsterImmuneElectricity":     reflect.ValueOf(object.MonsterImmuneElectricity),
		"MonsterImmuneFear":            reflect.ValueOf(object.MonsterImmuneFear),
		"MonsterImmuneFire":            reflect.ValueOf(object.MonsterImmuneFire),
		"MonsterImmunePoison":          reflect.ValueOf(object.MonsterImmunePoison),
		"MonsterLarge":                 reflect.ValueOf(object.MonsterLarge),
		"MonsterLookAround":            reflect.ValueOf(object.MonsterLookAround),
		"MonsterMedium":                reflect.ValueOf(object.MonsterMedium),
		"MonsterMigrate":               reflect.ValueOf(object.MonsterMigrate),
		"MonsterMonitor":               reflect.ValueOf(object.MonsterMonitor),
		"MonsterNPC":                   reflect.ValueOf(object.MonsterNPC),
		"MonsterNoSpellTarget":         reflect.ValueOf(object.MonsterNoSpellTarget),
		"MonsterNoTarget":              reflect.ValueOf(object.MonsterNoTarget),
		"MonsterShopkeeper":            reflect.ValueOf(object.MonsterShopkeeper),
		"MonsterSmall":                 reflect.ValueOf(object.MonsterSmall),
		"MonsterUndead":                reflect.ValueOf(object.MonsterUndead),
		"MonsterWarcryStun":            reflect.ValueOf(object.MonsterWarcryStun),
		"MonsterWoundedNPC":            reflect.ValueOf(object.MonsterWoundedNPC),
		"OtherChestNE":                 reflect.ValueOf(object.OtherChestNE),
		"OtherChestNW":                 reflect.ValueOf(object.OtherChestNW),
		"OtherChestSE":                 reflect.ValueOf(object.OtherChestSE),
		"OtherChestSW":                 reflect.ValueOf(object.OtherChestSW),
		"OtherClassNames":              reflect.ValueOf(&object.OtherClassNames).Elem(),
		"OtherGate":                    reflect.ValueOf(object.OtherGate),
		"OtherHeavy":                   reflect.ValueOf(object.OtherHeavy),
		"OtherInvisibleObelisk":        reflect.ValueOf(object.OtherInvisibleObelisk),
		"OtherLOTD":                    reflect.ValueOf(object.OtherLOTD),
		"OtherLava":                    reflect.ValueOf(object.OtherLava),
		"OtherStoneDoor":               reflect.ValueOf(object.OtherStoneDoor),
		"OtherTech":                    reflect.ValueOf(object.OtherTech),
		"OtherUseable":                 reflect.ValueOf(object.OtherUseable),
		"OtherVisibleObelisk":          reflect.ValueOf(object.OtherVisibleObelisk),
		"ParseArmorClass":              reflect.ValueOf(object.ParseArmorClass),
		"ParseArmorClassSet":           reflect.ValueOf(object.ParseArmorClassSet),
		"ParseBookClass":               reflect.ValueOf(object.ParseBookClass),
		"ParseBookClassSet":            reflect.ValueOf(object.ParseBookClassSet),
		"ParseClass":                   reflect.ValueOf(object.ParseClass),
		"ParseClassSet":                reflect.ValueOf(object.ParseClassSet),
		"ParseExitClass":               reflect.ValueOf(object.ParseExitClass),
		"ParseExitClassSet":            reflect.ValueOf(object.ParseExitClassSet),
		"ParseFlag":                    reflect.ValueOf(object.ParseFlag),
		"ParseFlagSet":                 reflect.ValueOf(object.ParseFlagSet),
		"ParseFoodClass":               reflect.ValueOf(object.ParseFoodClass),
		"ParseFoodClassSet":            reflect.ValueOf(object.ParseFoodClassSet),
		"ParseGeneratorClass":          reflect.ValueOf(object.ParseGeneratorClass),
		"ParseGeneratorClassSet":       reflect.ValueOf(object.ParseGeneratorClassSet),
		"ParseMaterial":                reflect.ValueOf(object.ParseMaterial),
		"ParseMaterialSet":             reflect.ValueOf(object.ParseMaterialSet),
		"ParseMissileClass":            reflect.ValueOf(object.ParseMissileClass),
		"ParseMissileClassSet":         reflect.ValueOf(object.ParseMissileClassSet),
		"ParseMonsterClass":            reflect.ValueOf(object.ParseMonsterClass),
		"ParseMonsterClassSet":         reflect.ValueOf(object.ParseMonsterClassSet),
		"ParseOtherClass":              reflect.ValueOf(object.ParseOtherClass),
		"ParseOtherClassSet":           reflect.ValueOf(object.ParseOtherClassSet),
		"ParseSubClass":                reflect.ValueOf(object.ParseSubClass),
		"ParseSubClassSet":             reflect.ValueOf(object.ParseSubClassSet),
		"ParseWeaponClass":             reflect.ValueOf(object.ParseWeaponClass),
		"ParseWeaponClassSet":          reflect.ValueOf(object.ParseWeaponClassSet),
		"SubClassNames":                reflect.ValueOf(&object.SubClassNames).Elem(),
		"WeaponArrow":                  reflect.ValueOf(object.WeaponArrow),
		"WeaponAxe":                    reflect.ValueOf(object.WeaponAxe),
		"WeaponBolt":                   reflect.ValueOf(object.WeaponBolt),
		"WeaponBow":                    reflect.ValueOf(object.WeaponBow),
		"WeaponChakram":                reflect.ValueOf(object.WeaponChakram),
		"WeaponClassNames":             reflect.ValueOf(&object.WeaponClassNames).Elem(),
		"WeaponCrossbow":               reflect.ValueOf(object.WeaponCrossbow),
		"WeaponFlag":                   reflect.ValueOf(object.WeaponFlag),
		"WeaponGreatSword":             reflect.ValueOf(object.WeaponGreatSword),
		"WeaponHammer":                 reflect.ValueOf(object.WeaponHammer),
		"WeaponLongSword":              reflect.ValueOf(object.WeaponLongSword),
		"WeaponMace":                   reflect.ValueOf(object.WeaponMace),
		"WeaponOgreAxe":                reflect.ValueOf(object.WeaponOgreAxe),
		"WeaponQuiver":                 reflect.ValueOf(object.WeaponQuiver),
		"WeaponShuriken":               reflect.ValueOf(object.WeaponShuriken),
		"WeaponStaff":                  reflect.ValueOf(object.WeaponStaff),
		"WeaponStaffDeathRay":          reflect.ValueOf(object.WeaponStaffDeathRay),
		"WeaponStaffFireball":          reflect.ValueOf(object.WeaponStaffFireball),
		"WeaponStaffForceOfNature":     reflect.ValueOf(object.WeaponStaffForceOfNature),
		"WeaponStaffLightning":         reflect.ValueOf(object.WeaponStaffLightning),
		"WeaponStaffOblivionHalberd":   reflect.ValueOf(object.WeaponStaffOblivionHalberd),
		"WeaponStaffOblivionHeart":     reflect.ValueOf(object.WeaponStaffOblivionHeart),
		"WeaponStaffOblivionOrb":       reflect.ValueOf(object.WeaponStaffOblivionOrb),
		"WeaponStaffOblivionWierdling": reflect.ValueOf(object.WeaponStaffOblivionWierdling),
		"WeaponStaffSulphorousFlare":   reflect.ValueOf(object.WeaponStaffSulphorousFlare),
		"WeaponStaffSulphorousShower":  reflect.ValueOf(object.WeaponStaffSulphorousShower),
		"WeaponStaffTripleFireball":    reflect.ValueOf(object.WeaponStaffTripleFireball),
		"WeaponSword":                  reflect.ValueOf(object.WeaponSword),

		// type definitions
		"ArmorClass":     reflect.ValueOf((*object.ArmorClass)(nil)),
		"BookClass":      reflect.ValueOf((*object.BookClass)(nil)),
		"Class":          reflect.ValueOf((*object.Class)(nil)),
		"ExitClass":      reflect.ValueOf((*object.ExitClass)(nil)),
		"Flags":          reflect.ValueOf((*object.Flags)(nil)),
		"FoodClass":      reflect.ValueOf((*object.FoodClass)(nil)),
		"GeneratorClass": reflect.ValueOf((*object.GeneratorClass)(nil)),
		"Material":       reflect.ValueOf((*object.Material)(nil)),
		"MissileClass":   reflect.ValueOf((*object.MissileClass)(nil)),
		"MonsterClass":   reflect.ValueOf((*object.MonsterClass)(nil)),
		"OtherClass":     reflect.ValueOf((*object.OtherClass)(nil)),
		"SubClass":       reflect.ValueOf((*object.SubClass)(nil)),
		"WeaponClass":    reflect.ValueOf((*object.WeaponClass)(nil)),
	}
}
