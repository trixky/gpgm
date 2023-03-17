import { writable } from 'svelte/store';
import type GenerationModel from '$lib/models/generation';
import type InstanceModel from '$lib/models/instance';
import type Instances from '$lib/models/instance';

function generate_default_instance_store(): Array<Instances> {
    return <Array<Instances>>[]
}

function create_instance_store() {
    const { subscribe, update, set } = writable(generate_default_instance_store());

    return {
        subscribe,
        reset: () => {
            set(generate_default_instance_store())
        },
        // --------------------- insert
        insert_population: (population: GenerationModel) => {
            update(instances => {
                if (instances.length) {
                    population.scores.sort((a, b) => a + b).forEach((score, index) => {
                        instances[index].scores.push(score)
                    })
                } else {
                    instances = population.scores.map(score =>
                        <InstanceModel>{
                            scores: [score]
                        }).sort((a, b) => a.scores[0] + b.scores[0])
                }

                return instances
            })
        },
    };
}

const statistic_store = create_instance_store();

export default statistic_store;
