import type { Instance } from "./instance";
import type { Simulation } from "./simulation";

// source: population.go

export interface ScoredInstance {
  instance: Instance /* instance.Instance */;
  simulation: Simulation /* simulation.Simulation */;
  score: number /* int */;
  cycle: number /* int */;
}
export interface ScoredPopulation {
  instances: ScoredInstance[];
}
export interface Population {
  instances: Instance /* instance.Instance */[];
}
