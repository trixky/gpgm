import { writable } from 'svelte/store';
import { generate_empty_generations, generate_random_generation, random_instance_number } from './temp_generations';
import type GenerationModel from '../models/generation';
import { get_last_generation_scores } from '../models/generation';
import type {Scores as ScoresModel} from '../models/statistic'

import StatisticStore from './statistic';

function sort_generation(generation: GenerationModel): GenerationModel {
	return <GenerationModel>{
		instances: generation.instances.filter(_ => true).sort((a, b) => b.score - a.score)
	}
}

function create_generation_store() {
	const { subscribe, update, set } = writable(generate_empty_generations());

	return {
		subscribe,
		push: (generation: GenerationModel) => {

			const current_score = get_last_generation_scores(generation)
			StatisticStore.set_insert_score(current_score) // can be optimized using the sorted array

			update(generations => {
				generations.push(sort_generation(generation))
				return generations
			})
		},
		push_random: (width: number | undefined = undefined) => {
			update(generations => {
				if (!width) {
					width = generations.length ? generations[0].instances.length : random_instance_number()
				}

				const random_generation = generate_random_generation(width)
				const current_score = get_last_generation_scores(random_generation)

				StatisticStore.set_insert_score(current_score) // can be optimized using the sorted array

				generations.push(sort_generation(random_generation))
				return generations
			})
		},
		reset: () => {
			set(generate_empty_generations())
			StatisticStore.reset()
		}
	};
}

const generation_store = create_generation_store();

export default generation_store;