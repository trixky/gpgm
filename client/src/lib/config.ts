import type { Arguments } from "../types"
import { LinearMutation, LogarithmicMutation, RandomSelection, TournamentSelection } from "../types/core.d"

interface TextareaConfig {
    cols: number,
    row: number,
}

interface NumericConfig {
    min: number,
    max: number,
    default: number
}

interface ChoiceConfig {
    default: number
    choices: number[]
}

interface ArgumentsConfig {
    ui: {
        input: TextareaConfig
        output: TextareaConfig
    }
    io: {
        selection_method: ChoiceConfig
        mutation_method: ChoiceConfig
    } &
    { [key in keyof Omit<Arguments, "text" | "selection_method" | "mutation_method">]: NumericConfig }
}

export const config: ArgumentsConfig = {
    ui: {
        input: {
            cols: 84,
            row: 30,
        },
        output: {
            cols: 42,
            row: 10,
        }
    },
    io: {
        max_generations: {
            min: 1,
            max: 10000,
            default: 6
        },
        population_size: {
            min: 1,
            max: 1000,
            default: 6
        },
        max_cycle: {
            min: 10,
            max: 1000,
            default: 1000
        },
        max_depth: {
            min: 1,
            max: 10,
            default: 1
        },
        time_limit: {
            min: 500,
            max: 60000,
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
            min: 1,
            max: 10,
            default: 1
        },
        tournament_size: {
            min: 1,
            max: 50,
            default: 10
        },
        tournament_probability: {
            min: 0,
            max: 1,
            default: 0.77
        },
        selection_method: {
            default: TournamentSelection,
            choices: [RandomSelection, TournamentSelection]
        },
        mutation_method: {
            default: LogarithmicMutation,
            choices: [LinearMutation, LogarithmicMutation]
        },
    }
}
export default config