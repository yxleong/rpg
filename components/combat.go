package components

type Combat interface {
	Health() int
	AttackPower() int
	Attacking() bool
	Attack() bool
	Udapte()
	Damage(amount int)
}

type BasicCombat struct {
	health      int
	attackPower int
	atacking    bool
}

func NewBasicCombat(health, attackPower int) *BasicCombat {
	return &BasicCombat{
		health:      health,
		attackPower: attackPower,
		atacking:    false,
	}
}

func (c *BasicCombat) Health() int {
	return c.health
}

func (c *BasicCombat) AttackPower() int {
	return c.attackPower
}

func (c *BasicCombat) Damage(amount int) {
	c.health -= amount
	// if c.health < 0 {
	// 	c.health = 0
	// }
}

func (c *BasicCombat) Attacking() bool {
	return c.atacking
}

func (c *BasicCombat) Attack() bool {
	c.atacking = true
	return true
}

func (c *BasicCombat) Update() {
}

var _ Combat = (*BasicCombat)(nil)

type EnemyCombat struct {
	*BasicCombat
	attackCooldown      int
	timeSinceLastAttack int
}

func NewEnemyCombat(health, attackPower, attackCooldown int) *EnemyCombat {
	return &EnemyCombat{
		BasicCombat:         NewBasicCombat(health, attackPower),
		attackCooldown:      attackCooldown,
		timeSinceLastAttack: 0,
	}
}

func (e *EnemyCombat) Attack() bool {
	if e.timeSinceLastAttack >= e.attackCooldown {
		e.atacking = true
		e.timeSinceLastAttack = 0
		return true
	}
	return false
}

func (e *EnemyCombat) Update() {
	e.timeSinceLastAttack++
	// if e.atacking {
	// 	e.timeSinceLastAttack = 0
	// 	e.atacking = false
	// }
}

var _ Combat = (*EnemyCombat)(nil)
