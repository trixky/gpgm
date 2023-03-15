// source: context.go

export interface Product {
  name: string;
  quantity: number /* int */;
}
export interface Process {
  name: string;
  inputs: { [key: string]: number /* int */ };
  outputs: { [key: string]: number /* int */ };
  delay: number /* int */;
  parent: number /* int */[];
}
export interface InitialContext {
  stock: Stock;
  processes: Process[];
  optimize: { [key: string]: boolean };
  score_ratio: { [key: string]: number /* int */ };
}

//////////
// source: options.go

export type SelectionMethod = number /* int64 */;
export const RandomSelection: SelectionMethod = 0;
export const TournamentSelection: SelectionMethod = 1;
export type MutationMethod = number /* int64 */;
export const LinearMutation: MutationMethod = 0;
export const LogarithmicMutation: MutationMethod = 1;

export interface Options {
  max_generation: number /* int */;
  time_limit_ms: number /* int */;
  max_cycle: number /* int */;
  max_depth: number /* int */;
  /**
   * Population
   */
  population_size: number /* int */;
  elitism_amount: number /* int */;
  selection_method: SelectionMethod;
  tournament_size: number /* int */;
  tournament_probability: number /* float64 */;
  crossover_new_instances: number /* int */;
  /**
   * Mutation
   */
  mutation_chance: number /* float64 */;
  mutation_method: MutationMethod;
  /**
   * Genetic
   */
  n_entry: number /* int */;
  history_part_max_length: number /* int */;
  history_key_max_length: number /* int */;
  random_cut: boolean;
  max_cut: number /* int */;
}

//////////
// source: stock.go

export type Stock = { [key: string]: number /* int */ };
