package instance

type Gene struct {
	FirstPriorityExon Exon   `json:"first_priority_exon"`
	LastPriorityExon  Exon   `json:"last_priority_exon"`
	RatioExons        []Exon `json:"ratio_exons"`
}

// # EXPERIMENTAL #
// Cross generates a child by cross overing another one
func (g *Gene) Cross(gg *Gene) (child Gene) {
	child.RatioExons = make([]Exon, len(g.RatioExons))

	child.FirstPriorityExon = g.FirstPriorityExon.Cross(&gg.FirstPriorityExon)
	child.LastPriorityExon = g.LastPriorityExon.Cross(&gg.LastPriorityExon)

	for index, exon := range g.RatioExons {
		child.RatioExons[index] = exon.Cross(&gg.RatioExons[index])
	}

	return
}

// Mutate generates a child by mutation
func (g *Gene) Mutate(max int) (child Gene) {
	child.RatioExons = make([]Exon, len(g.RatioExons))

	child.LastPriorityExon = g.LastPriorityExon.Mutate(max)
	child.LastPriorityExon = g.LastPriorityExon.Mutate(max)

	for index, exon := range g.RatioExons {
		child.RatioExons[index] = exon.Mutate(max)
	}

	return
}
