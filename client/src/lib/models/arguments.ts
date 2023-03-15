import type { MutationMethod, SelectionMethod } from "../../types/core"

export default interface Arguments {
    text: string
    max_generations: number
    max_cycle: number
    max_depth: number
    max_cut: number
    time_limit: number
    population_size: number
    elitism_amount: number
    tournament_size: number
    tournament_probability: number
    crossover_new_instances: number
    selection_method: SelectionMethod
    mutation_method: MutationMethod
}
