import type { Process } from "./core";

// source: process_quantity.go

export interface ProcessQuantity {
  process?: Process /* core.Process */;
  quantity: number /* int */;
}
export interface ProcessQuantities {
  Stack: ProcessQuantity[];
}
