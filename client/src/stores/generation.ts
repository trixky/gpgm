import { writable } from 'svelte/store';
import { generate_empty_generations, generate_random_generation, random_instance_number } from './temp_generations';
import type GenerationModel from '../models/generation';

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
				generations.push(sort_generation(generate_random_generation(width)))
				return generations
			})
		},
		reset: () => {
			set(generate_empty_generations())
		}
	};
}

const generation_store = create_generation_store();

export default generation_store;