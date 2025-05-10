package hw06

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		person.name = [42]byte{}
		for i := 0; i < len(name); i++ {
			if i >= len(person.name) {
				break
			}
			person.name[i] = name[i]
		}
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.x = int32(x)
		person.y = int32(y)
		person.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.mr = person.mr | ((uint16(mana) << 4) & 0xFFF0)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.hs = person.hs | ((uint16(health) << 4) & 0xFFF0)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.mr = person.mr | (uint16(respect) & 0x000F)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.hs = person.hs | (uint16(strength) & 0x000F)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.el = person.el | ((uint8(experience) << 4) & 0xF0)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.el = person.el | (uint8(level) & 0x0F)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hgfpt = person.hgfpt | 0x01
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hgfpt = person.hgfpt | 0x02
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hgfpt = person.hgfpt | 0x04
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.hgfpt = person.hgfpt | (uint8(personType<<4) & 0xF0)
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	x     int32    // 4
	y     int32    // 4
	z     int32    // 4
	gold  uint32   // 4
	name  [42]byte // 42
	mr    uint16   // mana, respect 2
	hs    uint16   // health, strength 2
	el    uint8    // experience,level 1
	hgfpt uint8    // hasHouse,hasGun,hasFamily,personType 1
}

func NewGamePerson(options ...Option) GamePerson {
	person := GamePerson{}
	for _, option := range options {
		option(&person)
	}
	return person
}

func (p *GamePerson) Name() string {
	return string(p.name[:])
}

func (p *GamePerson) X() int {
	return int(p.x)
}

func (p *GamePerson) Y() int {
	return int(p.y)
}

func (p *GamePerson) Z() int {
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return int(p.mr >> 4)
}

func (p *GamePerson) Health() int {
	return int(p.hs >> 4)
}

func (p *GamePerson) Respect() int {
	return int(p.mr & 0x000F)
}

func (p *GamePerson) Strength() int {
	return int(p.hs & 0x000F)
}

func (p *GamePerson) Experience() int {
	return int(p.el >> 4)
}

func (p *GamePerson) Level() int {
	return int(p.el & 0x0F)
}

func (p *GamePerson) HasHouse() bool {
	return (p.hgfpt & 0x01) != 0
}

func (p *GamePerson) HasGun() bool {
	return (p.hgfpt & 0x02) != 0
}

func (p *GamePerson) HasFamily() bool {
	return (p.hgfpt & 0x04) != 0
}

func (p *GamePerson) Type() int {
	return int((p.hgfpt & 0xF0) >> 4)
}
