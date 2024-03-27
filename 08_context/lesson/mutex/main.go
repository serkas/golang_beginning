package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"proj/lessons/08_context/lesson/mutex/memorycache"
	"proj/lessons/08_context/lesson/mutex/weather"
	"sync"
	"time"
)

func main() {
	defer stopwatchPrinter(time.Now())
	cities := []string{
		"Kyiv",
		"Berlin",
		"Vienna",
		"Sofia",
		"Zagreb",
		"Prague",
		"Copenhagen",
		"Tallinn",
		"Helsinki",
		"Paris",
		"Athens",
	}

	//single loop weather
	//for _, city := range cities {
	//	currentWeather, err := weather.GetCurrentWeather(city)
	//	if err != nil {
	//		log.Printf("got error: %s", err)
	//		continue
	//	}
	//
	//	fmt.Printf("Current weather in %s is %s", city, currentWeather)
	//}

	// repeats
	//moreCities := enlargeAndShuffleSlice(cities, 10)
	//for _, city := range moreCities {
	//	currentWeather, err := weather.GetCurrentWeather(city)
	//	if err != nil {
	//		log.Printf("got error: %s", err)
	//		continue
	//	}
	//
	//	fmt.Printf("Current weather in %s is %s", city, currentWeather)
	//}

	// repeats with caching
	//moreCities := enlargeAndShuffleSlice(cities, 10)
	//cache := memorycache.New()
	//for _, city := range moreCities {
	//	// try cache
	//	weatherValue, err := cache.Get(city)
	//	if err != nil {
	//		currentWeather, err := weather.GetCurrentWeather(city)
	//		if err != nil {
	//			log.Printf("got error: %s", err)
	//			continue
	//		}
	//
	//		cache.Set(city, currentWeather) // store value for future
	//		weatherValue = currentWeather
	//	}
	//
	//	fmt.Printf("Current weather in %s is %s", city, weatherValue)
	//}

	//parallel execution
	moreCities := enlargeAndShuffleSlice(cities, 10)
	cityTasks := make(chan string)
	go func() {
		for _, city := range moreCities {
			cityTasks <- city
		}
		close(cityTasks) // no more tasks
	}()

	numWorkers := 10
	//start pool of workers
	//
	//for w := 0; w <= numWorkers; w++ {
	//	go func(workerID int) {
	//		for city := range cityTasks {
	//			currentWeather, err := weather.GetCurrentWeather(city)
	//			if err != nil {
	//				log.Printf("got error: %s", err)
	//				continue
	//			}
	//			fmt.Printf("w%d: Current weather in %s is %s", workerID, city, currentWeather)
	//		}
	//	}(w)
	//}

	//time.Sleep(5 * time.Second)

	// start a better pool of workers (with waitGroup)
	//wg := sync.WaitGroup{}
	//wg.Add(numWorkers)
	//for w := 0; w <= numWorkers; w++ {
	//	go func(workerID int) {
	//		for city := range cityTasks {
	//			currentWeather, err := weather.GetCurrentWeather(city)
	//			if err != nil {
	//				log.Printf("got error: %s", err)
	//				continue
	//			}
	//			fmt.Printf("w%d: Current weather in %s is %s", workerID, city, currentWeather)
	//		}
	//
	//		wg.Done()
	//	}(w)
	//}
	//wg.Wait() // waiting exact moment all workers are done

	//start a better pool of workers with cache
	cache := memorycache.New()
	waitingGroup := &sync.WaitGroup{}
	for workerID := 0; workerID <= numWorkers; workerID++ {
		waitingGroup.Add(1)
		w := worker{
			id:    workerID,
			cache: cache,
		}
		go func(wg *sync.WaitGroup) {
			for city := range cityTasks {
				w.handleSingleCity(city)
			}
			wg.Done()
		}(waitingGroup)
	}

	waitingGroup.Wait()
}

type worker struct {
	id    int
	cache *memorycache.Cache
}

func (w *worker) handleSingleCity(city string) {
	weatherValue, err := w.cache.Get(city)
	if err != nil {
		currentWeather, err := weather.GetCurrentWeather(city)
		if err != nil && !errors.Is(err, memorycache.ErrNotFound) {
			log.Printf("got error: %s", err)
			return
		}

		w.cache.Set(city, currentWeather) // store value for future
		weatherValue = currentWeather
	}

	fmt.Printf("w%d: Current weather in %s is %s", w.id, city, weatherValue)
}

func stopwatchPrinter(start time.Time) {
	fmt.Printf("Execution took %v\n", time.Since(start))
}

func enlargeAndShuffleSlice(in []string, n int) []string {
	enlarged := make([]string, 0, len(in)*n)
	for i := 0; i < n; i++ {
		enlarged = append(enlarged, in...)
	}

	// shuffle
	for i := range enlarged {
		j := rand.Intn(i + 1)
		enlarged[i], enlarged[j] = enlarged[j], enlarged[i]
	}

	return enlarged
}
