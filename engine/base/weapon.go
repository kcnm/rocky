package base

type Weapon interface {
	Card() WeaponCard
	Attack() int
	Durability() int

	LoseDurability()
}
