import TransactionHistory from "./transaction-history";
import Transaction from "./transaction";

export default class Account {
  protected _balance: number
  protected _history: TransactionHistory

  constructor(initialBalance: number = 0) {
    this._balance = initialBalance
    this._history = new TransactionHistory()
  }

  balance(): number {
    return this._balance
  }

  history(): TransactionHistory {
    return this._history
  }

  apply(transaction: Transaction): void {
    transaction.applyTo(this)
    this._history.add(transaction)
  }
  
  deposit(amount: number): void {
    this._balance += amount
  }
  
  withdraw(amount: number): void {
    if (amount > this._balance)
      throw new Error('Not enough funds')
    
    this._balance -= amount
  }
}