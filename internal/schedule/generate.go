package schedule

type GenerateSpec struct {
	Schedule  Schedule
	Employees []Employee
	MergeStat MergeStat
}

type Generator struct {
	Spec GenerateSpec
}

func (g *Generator) Generate() Schedule {
	return g.Spec.Schedule
}
