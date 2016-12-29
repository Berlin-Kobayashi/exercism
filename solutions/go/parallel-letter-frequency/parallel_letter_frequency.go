// Package letter implements a library for determining the frequency of letters in strings.
package letter

import "sync"

//ConcurrentFrequency determines the frequency of letters in the given strings.
func ConcurrentFrequency(strings []string) FreqMap {
	var wg sync.WaitGroup
	channel := make(chan FreqMap, len(strings))
	for _, s := range strings {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			channel <- Frequency(s)
		}(s)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	result := <-channel
	for subResult := range channel {
		for letter, frequency := range subResult {
			result[letter] += frequency
		}
	}

	return result
}
