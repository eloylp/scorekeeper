package scorekeeper

type Scorer struct {
	scores map[string]int
}

func NewScorer() Scorer {
	return Scorer{
		scores: make(map[string]int),
	}
}

func (s Scorer) Get(player string) int {
	return s.scores[player]
}

func (s Scorer) Add(player string, point int) {
	currentScore := s.scores[player]
	s.scores[player] = point + currentScore
}

func (s Scorer) Subs(user string, points int) {
	newScore := s.scores[user] - points
	if newScore < 0 {
		newScore = 0
	}
	s.scores[user] = newScore
}
