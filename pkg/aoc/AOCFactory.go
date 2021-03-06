package aoc

type AOCFactory struct {
	factory map[string]AdventOfCodeChallenge
}

func NewAOCFactory() *AOCFactory {
	return &AOCFactory{
		factory: make(map[string]AdventOfCodeChallenge),
	}
}

func (f *AOCFactory) Register(year, day, part string, challenge AdventOfCodeChallenge) {
	f.factory[year+day+part] = challenge
}

func (f *AOCFactory) Get(year, day, part string) AdventOfCodeChallenge {
	return f.factory[year+day+part]
}

func (f *AOCFactory) List() map[string]AdventOfCodeChallenge {
	return f.factory
}

var aocFactory *AOCFactory

func GetAOCFactory() *AOCFactory {
	if aocFactory == nil {
		aocFactory = NewAOCFactory()
	}

	return aocFactory
}
