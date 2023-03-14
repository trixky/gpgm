import type { InitialContext, Options } from "./core";
import type { ScoredPopulation } from "./population";

// source: main.go

export interface Arguments {
  text: string;
  generations: number /* int */;
  deep: number /* int */;
  population: number /* int */;
}
export interface RunningSolver {
  population: Population /* population.Population */;
  context: InitialContext /* core.InitialContext */;
  options: Options /* core.Options */;
  generation: number /* int */;
  start: unknown /* time.Time */;
}
export interface WASMGenerationReturn {
  scored_population: ScoredPopulation /* population.ScoredPopulation */;
  running_solver: RunningSolver;
}
