import type { Arguments } from "../types"
import { LinearMutation, LogarithmicMutation, RandomSelection, TournamentSelection } from "../types/core.d"

interface NumericConfig {
    min: number,
    max: number,
    default: number
}

interface ChoiceConfig {
    default: number
    choices: { value: number, label: string }[]
}

export type NumericConfigKeys = keyof Omit<Arguments, "text" | "selection_method" | "mutation_method">;

interface ArgumentsConfig {
    io: {
        selection_method: ChoiceConfig
        mutation_method: ChoiceConfig
    } &
    { [key in NumericConfigKeys]: NumericConfig }
}

export const config: ArgumentsConfig = {
    io: {
        max_generations: {
            min: 1,
            max: 100,
            default: 6
        },
        population_size: {
            min: 1,
            max: 100,
            default: 6
        },
        max_cycle: {
            min: 10,
            max: 100000,
            default: 1000
        },
        max_depth: {
            min: 1,
            max: 10,
            default: 6
        },
        time_limit: {
            min: 500,
            max: 600000,
            default: 5000
        },
        max_cut: {
            min: 1,
            max: 5,
            default: 1
        },
        crossover_new_instances: {
            min: 1,
            max: 10,
            default: 1
        },
        elitism_amount: {
            min: 0,
            max: 10,
            default: 1
        },
        tournament_size: {
            min: 1,
            max: 50,
            default: 10
        },
        tournament_probability: {
            min: 0.0,
            max: 1.0,
            default: 0.77
        },
        selection_method: {
            default: TournamentSelection,
            choices: [
                { value: RandomSelection, label: 'Random' },
                { value: TournamentSelection, label: 'Tournament' }
            ]
        },
        mutation_method: {
            default: LogarithmicMutation,
            choices: [
                { value: LinearMutation, label: 'Linear' },
                { value: LogarithmicMutation, label: 'Logarithmic' }
            ]
        },
    }
}
export default config