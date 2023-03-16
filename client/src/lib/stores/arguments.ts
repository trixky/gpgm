import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import config from '$lib/config'
import type ArgumentsModel from '$lib/models/arguments';

const defaultValue = () => JSON.stringify(<ArgumentsModel>{
    max_generations: config.io.max_generations.default,
    population_size: config.io.population_size.default,
    max_cycle: config.io.max_cycle.default,
    max_depth: config.io.max_depth.default,
    time_limit: config.io.time_limit.default,
    max_cut: config.io.max_cut.default,
    crossover_new_instances: config.io.crossover_new_instances.default,
    elitism_amount: config.io.elitism_amount.default,
    tournament_size: config.io.tournament_size.default,
    tournament_probability: config.io.tournament_probability.default,
    selection_method: config.io.selection_method.default,
    mutation_method: config.io.mutation_method.default,
})

export const store = writable<ArgumentsModel>(JSON.parse(browser ? localStorage.getItem('arguments') ?? defaultValue() : defaultValue()));

store.subscribe((value) => browser ? localStorage.setItem('arguments', JSON.stringify(value)) : undefined)

export default store;
