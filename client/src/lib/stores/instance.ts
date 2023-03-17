import { writable } from 'svelte/store';
import type GenerationModel from '$lib/models/generation';

function generate_default_instance_store(): Array<Array<number>> {
    return <Array<Array<number>>>[]
}

function create_instance_store() {
    const { subscribe, update, set } = writable(generate_default_instance_store());

    return {
        subscribe,
        reset: () => {
            set(generate_default_instance_store())
        },
        insert_population: (population: GenerationModel) => {
            update(instances => {
                if (instances.length) {
                    population.scores.sort((a, b) => a + b).forEach((score, index) => {
                        instances[index].push(score)
                    })
                } else {
                    instances = population.scores.sort((a, b) => a + b).map(score => [score])
                }

                return instances
            })
        },
    };
}

const statistic_store = create_instance_store();

export default statistic_store;
