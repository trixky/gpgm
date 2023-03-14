import type { InitialContext, Process, Stock } from "./core";
import type { Instance } from "./instance";

// source: expected.go

export interface ExpectedStock {
  name: string;
  quantity: number /* int */;
  remaining_cycles: number /* int */;
}

// source: simulation.go

export interface ExecutedProcess {
  cycle: number /* int */;
  process: Process /* core.Process */;
  quantity: number /* int */;
}
export interface Simulation {
  initial_context: InitialContext /* core.InitialContext */;
  instance: Instance /* instance.Instance */;
  stock: Stock /* core.Stock */;
  expected_stock: ExpectedStock[];
  history: ExecutedProcess[];
  cycle: number /* int */;
}
export interface HistoryAction {
  Name: string;
  Amount: number /* int */;
}
export interface HistoryEntry {
  Cycle: number /* int */;
  Actions: HistoryAction[];
}
