## Ticker.Stop()
- Stop turns off a ticker. After Stop, no more ticks will be sent.
- Stop does not close the channel, to prevent a concurrent goroutine reading from the channel from seeing an erroneous "tick".


