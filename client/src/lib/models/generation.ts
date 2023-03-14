import type InstanceModel from "./instance"
import type { Scores as ScoresModel } from "./statistic"

export default interface Generation {
	instances: Array<InstanceModel>;
}

export function get_last_generation_scores(generation: Generation): ScoresModel {
	const scores = <ScoresModel>{
		worst: 0,
		best: 0,
	}

	const current_scores = generation.instances.map(instance => instance.score)

	if (generation.instances.length) {
		scores.best = Math.max(...current_scores)
		scores.worst = Math.min(...current_scores)
	}

	return scores
}