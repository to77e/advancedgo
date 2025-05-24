package hw06

const (
	maxNameLength = 42 // Maximum length of the name field in symbols.
)

// The mr field contains mana and respect, where mana is shifted by 4 bits.
const (
	manaShift   = 4      // Shift for a mana in the mr field.
	manaMask    = 0xFFF0 // Mask for a mana in the mr field.
	respectMask = 0x000F // Mask for respect in the mr field.
)

// The hs field contains health and strength, where health is shifted by 4 bits.
const (
	healthShift  = 4      // Shift for health in the hs field.
	healthMask   = 0xFFF0 // Mask for health in the hs field.
	strengthMask = 0x000F // Mask for strength in the hs field.
)

// The el field contains experience and level, where experience is shifted by 4 bits.
const (
	experienceShift = 4    // Shift for experience in the el field.
	experienceMask  = 0xF0 // Mask for experience in the el field.
	levelMask       = 0x0F // Mask for level in the el field.
)

// The hgfpt field contains flags for hasHouse, hasGun, hasFamily, and personType.
const (
	hasHouseMask    = 0x01 // Mask for the hasHouse bit in the hgfpt field.
	hasGunMask      = 0x02 // Mask for the hasGun bit in the hgfpt field.
	hasFamilyMask   = 0x04 // Mask for the hasFamily bit in the hgfpt field.
	personTypeShift = 4    // Shift for a person type in the hgfpt field.
	personTypeMask  = 0xF0 // Mask for a person type in the hgfpt field.
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		person.name = [maxNameLength]byte{}
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
		person.mr = person.mr | ((uint16(mana) << manaShift) & manaMask)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.hs = person.hs | ((uint16(health) << healthShift) & healthMask)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.mr = person.mr | (uint16(respect) & respectMask)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.hs = person.hs | (uint16(strength) & strengthMask)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.el = person.el | ((uint8(experience) << experienceShift) & experienceMask)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.el = person.el | (uint8(level) & levelMask)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hgfpt = person.hgfpt | hasHouseMask
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hgfpt = person.hgfpt | hasGunMask
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.hgfpt = person.hgfpt | hasFamilyMask
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.hgfpt = person.hgfpt | (uint8(personType<<personTypeShift) & personTypeMask)
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	x     int32               // 4 bytes
	y     int32               // 4 bytes
	z     int32               // 4 bytes
	gold  uint32              // 4 bytes
	name  [maxNameLength]byte // 42 symbols (42 bytes)
	mr    uint16              // mana - 1 byte; respect - 1 byte
	hs    uint16              // health - 1 byte; strength - 1 byte
	el    uint8               // experience - 4 bits; level - 4 bits
	hgfpt uint8               // hasHouse - 1 bit; hasGun - 1 bit; hasFamily - 1 bit; personType - 4 bits
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
	return int(p.mr >> manaShift)
}

func (p *GamePerson) Health() int {
	return int(p.hs >> healthShift)
}

func (p *GamePerson) Respect() int {
	return int(p.mr & respectMask)
}

func (p *GamePerson) Strength() int {
	return int(p.hs & strengthMask)
}

func (p *GamePerson) Experience() int {
	return int(p.el >> experienceShift)
}

func (p *GamePerson) Level() int {
	return int(p.el & levelMask)
}

func (p *GamePerson) HasHouse() bool {
	return (p.hgfpt & hasHouseMask) != 0
}

func (p *GamePerson) HasGun() bool {
	return (p.hgfpt & hasGunMask) != 0
}

func (p *GamePerson) HasFamily() bool {
	return (p.hgfpt & hasFamilyMask) != 0
}

func (p *GamePerson) Type() int {
	return int((p.hgfpt & personTypeMask) >> personTypeShift)
}
