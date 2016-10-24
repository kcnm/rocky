package engine

type Weapon interface {
	Card() WeaponCard
	Attack() int
	Durability() int

	LoseDurability()
}
