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

export interface Options {
  max_generation: number /* int */;
  time_limit_seconds: number /* int */;
  max_cycle: number /* int */;
  max_depth: number /* int */;
  n_entry: number /* int */;
  history_part_max_length: number /* int */;
  history_key_max_length: number /* int */;
  population_size: number /* int */;
  use_elitism: boolean;
  elitism_amount: number /* int */;
  random_cut: boolean;
  max_cut: number /* int */;
}

//////////
// source: stock.go

export type Stock = { [key: string]: number /* int */ };
