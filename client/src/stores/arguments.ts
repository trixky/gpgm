import { writable } from 'svelte/store';
import Config from '../config'
import type ArgumentsModel from '../models/arguments';

function generate_default_arguments(): ArgumentsModel {
    return <ArgumentsModel>{
        generations: Config.io.generations.default,
        population: Config.io.population.default,
        deep: Config.io.deep.default,
        delay: Config.io.delay.default,
    }
}

function create_argument_store() {
	const { subscribe, update, set } = writable(generate_default_arguments());

	return {
		subscribe,
        default: () => {
            set(generate_default_arguments())
        },
		update_generations: (generations: number) => {
            update(args => {
                args.generations = generations
                return args
            })
        },
        update_population: (population: number) => {
            update(args => {
                args.population = population
                return args
            })
        },
        update_deep: (deep: number) => {
            update(args => {
                args.deep = deep
                return args
            })
        },
        update_delay: (delay: number) => {
            update(args => {
                args.delay = delay
                return args
            })
        },
	};
}

const argument_store = create_argument_store();

export default argument_store;