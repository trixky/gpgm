import type { Process } from "./core";

// source: chromosome.go

export interface Chromosome {
  entry_gene: EntryGene;
  genes: PriorityGene[];
}

//////////
// source: dependencies.go

export interface InputDependencies {
  Input: string;
  ProcessDependencies: number /* int */[];
}
export interface ProcessDependencies {
  InputDependencies: InputDependencies[];
}

//////////
// source: entry_gene.go

export interface EntryGene {
  Process_ids: number /* int */[];
}

//////////
// source: genome.go

export interface Genome {
  chromosome: Chromosome[];
}

//////////
// source: instance.go

export interface Instance {
  chromosome: Chromosome;
}

//////////
// source: priority_gene.go

export interface PriorityGene {
  HistoryProcessDependencies: { [key: string]: ProcessDependencies };
  Process?: Process /* core.Process */;
}

//////////
// source: xeon.go

export interface Exon {
  value: number /* uint16 */;
}
