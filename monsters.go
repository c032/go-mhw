package mhw

var _ Creature = (*LargeMonster)(nil)

type LargeMonster struct {
	name   string
	rawURL string
}

func (lm *LargeMonster) Name() string {
	return lm.name
}

func (lm *LargeMonster) RawURL() string {
	return lm.rawURL
}
