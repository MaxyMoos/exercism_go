package gigasecond

import "time"

const TestVersion = 3

// Using the standard library Time.Add function and Second constant
// the exercise is straightforward
func AddGigasecond(input time.Time) time.Time {
    return input.Add( time.Second * 1e9 )
}

