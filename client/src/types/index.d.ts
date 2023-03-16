import type { InitialContext, MutationMethod, Options, SelectionMethod } from "./core";
import type { ScoredPopulation } from "./population";

// source: main.go

export interface Arguments {
  text: string
  max_generations: number /* int */
  max_cycle: number /* int */
  max_depth: number /* int */
  max_cut: number /* int */
  time_limit: number /* int */
  population_size: number /* int */
  elitism_amount: number /* int */
  tournament_size: number /* int */
  tournament_probability: number /* float64 */
  crossover_new_instances: number
  selection_method: SelectionMethod
  mutation_method: MutationMethod
}
export interface RunningSolver {
  population: Population /* population.Population */;
  context: InitialContext /* core.InitialContext */;
  options: Options /* core.Options */;
  generation: number /* int */;
  start: unknown /* time.Time */;
  time_limit_ms: number /* int */;
}
export interface WASMGenerationReturn {
  scored_population: ScoredPopulation /* population.ScoredPopulation */;
  running_solver: RunningSolver;
}
