// source: checker.go

export interface LocalExpectedStock {
  Product: string;
  Quantity: number /* int */;
  AvailableAtCycle: number /* int */;
}
export interface TokenWithQuantity {
  Name: string;
  Quantity: number /* int */;
}
