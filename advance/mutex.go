package advance

import (
	"bytes"
	"sync"
)

var (
	mu sync.RWMutex
	cache = make(map[string]string)
)

func get(k string) string {
	mu.RLock()
	defer mu.RUnlock()
	return cache[k]
}

func set(k, v string) {
	mu.Lock()
	defer mu.Unlock()
	cache[k] = v 
}

// sync.Once = guaranteed single execution:
var (
	once sync.Once
	db *Database
)

func getDB() *Database {
	once.Do(func() {
		db = openDB()
	})
	return db
}

var wg sync.WaitGroup

var bufPool = sync.Pool {
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func process(n int) int {
	return n
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			process(n)
		}(i)
	}
	wg.Wait()

	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)
}