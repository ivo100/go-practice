package main

import (
	"fmt"
	"sync"
)

// To execute Go code, please declare a func main() in a package "main"
/*

in memory kv store

.set(key, value string)
.get(key string) string
.unset(key string)

//transactions
.begin()
.rollback()
.commit()

//
get("a") -> nil
set("a", "a value")
get("a") -> "a value"
unset("a")
get("a") -> nil

//
get("a") -> nil
set("a", "a value")
get("a") -> "a value"
begin()
get("a") -> "a value"
set("a", "a different value")
get("a") -> "a different value"
rollback()
get("a") -> "a value"

//
get("a") -> nil
set("a", "a value")
get("a") -> "a value"
begin()
get("a") -> "a value"
set("a", "a different value")
get("a") -> "a different value"
commit()
get("a") -> "a different value"
rollback()
get("a") -> "a different value"

//
get("a") -> nil
set("a", "a value")
get("a") -> "a value"
begin()
get("a") -> "a value"
set("a", "a different value")
begin()
set("a", "a third value")
get("a") -> "a third value"
rollback()
get("a") -> "a different value"
rollback()
get("a") -> "a value"


//
get("a") -> nil
set("a", "a value")
get("a") -> "a value"
begin()
get("a") -> "a value"
set("a", "a different value")
begin()
set("a", "a third value")
get("a") -> "a third value"
commit()
get("a") -> "a different value"
rollback()
get("a") -> "a different value"

set a "1"
begin
   set a 5
	 begin
	    unset a
			commit
//
get a
A ""



*/

package main

import (
"fmt"
"sync"
)

type KvStore interface {
	set(key, value string)
	get(key string) (string, bool)
	unset(key string)

	//transactions
	begin()
	rollback()
	commit()
}

type store struct {
	//st sync.Map
	txLevel int
	levels []sync.Map
	deleted []string

}

func NewStore() KvStore {
	s := &store{}
	s.levels = make([]sync.Map, 0)
	s.levels = append(s.levels, sync.Map{})
	s.deleted = make([]string, 0)
	return s
}


func (s *store) set(key, value string) {
	s.levels[s.txLevel].Store(key, value)
}

func (s *store) get(key string) (string, bool) {
	v, ok := s.levels[s.txLevel].Load(key)
	if !ok {
		return "", false
	}
	return v.(string), ok
}

func (s *store) unset(key string) {
	s.deleted = append(s.deleted, key)
	s.levels[s.txLevel].Delete(key)
}

/*
      m.Range(func(key, value interface{}) bool {
               mm, ok := value.(*sync.Map)
   if !ok {
     // ...
   }
               mm.Store("xxxxxxx", c)
   c++
               return true
       })
*/

func (s *store) begin() {
	s.levels = append(s.levels, sync.Map{})
	s.deleted = make([]string, 0)
	m := s.levels[s.txLevel+1]
	// copy the level
	s.levels[s.txLevel].Range(
		func (k, v any) bool {
			m.Store(k, v)
			return true
		},
	)
	s.txLevel++
}

func (s *store) commit() {
	if s.txLevel == 0 {
		panic("wrong level")
	}
	//todo: check for level error
	m := s.levels[s.txLevel-1]
	// merge the level to previous
	s.levels[s.txLevel].Range(
		func (k, v any) bool {
			m.Store(k, v)
			return true
		},
	)
	// handle deleted = thombstones etc....

	s.txLevel--
}

func (s *store) rollback() {
	if s.txLevel == 0 {
		panic("wrong level")
	}
	s.txLevel--
}

func main() {

	kvStore := NewStore()

	v, ok := kvStore.get("a")
	fmt.Printf("value: %s, found: %v\n", v, ok)

	kvStore.set("a", "1")
	v, ok = kvStore.get("a")
	fmt.Printf("value: %s, found: %v\n", v, ok)

	kvStore.unset("a")
	v, ok = kvStore.get("a")
	fmt.Printf("value: %s, found: %v\n", v, ok)

	kvStore.begin()

	kvStore.set("a", "22222")
	v, ok = kvStore.get("a")
	fmt.Printf("value: %s, found: %v\n", v, ok)

	kvStore.commit()

	v, ok = kvStore.get("a")
	fmt.Printf("value: %s, found: %v\n", v, ok)

	kvStore.rollback()

	v, ok = kvStore.get("a")
	fmt.Printf("value: %s, found: %v\n", v, ok)

}

