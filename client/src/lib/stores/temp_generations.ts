import type GenerationModel from '$lib/models/generation';
import type InstanceModel from '$lib/models/instance';

const min_instances = 30
const max_instances = 500
const diff_instances = max_instances - min_instances

export function random_instance_number() {
    return Math.ceil(Math.random() * diff_instances + min_instances);
}

export function generate_random_generation(nb_instance: number | undefined): GenerationModel {
    const instances = <Array<InstanceModel>>[]

    if (nb_instance == undefined) {
        nb_instance = random_instance_number()
    }

    for (let i = 0; i < nb_instance; i++) {
        instances.push(<InstanceModel>{
            score: Math.ceil(Math.random() * 100)
        })
    }

    return <GenerationModel>{
        instances
    }
}

export function generate_empty_generations(): Array<GenerationModel> {
    return []
}

export function generate_random_generations(): Array<GenerationModel> {


    const nb_instance = random_instance_number()
    const nb_generation = 200

    const generations: Array<GenerationModel> = []

    for (let g = 0; g < nb_generation; g++) {
        generations.push(generate_random_generation(nb_instance))
    }

    return generations
}

export const random_generated: Array<GenerationModel> = []
