import { writable } from 'svelte/store';
import {generate_empty_generations, generate_random_generation, random_instance_number} from './temp_generations';
import type GenerationModel from '../models/generation';

function createGenerationStore() {
	const { subscribe, update, set } = writable(generate_empty_generations());

	return {
		subscribe,
		push: (generation: GenerationModel) => {
			update(generations => {
				generations.push(generation)
				return generations
			})
		},
		push_random: () => {
			update(generations => {
				const width = generations.length ? generations[0].instances.length : random_instance_number()
				generations.push(generate_random_generation(width))
				return generations
			})
		},
		reset: () => {
			set(generate_empty_generations())
		}
	};
}

const generationStore = createGenerationStore();


export default generationStore;