package main

//For the Cuckoo Bloom Filter we have only a few primitives:

type Database interface {
	// Lookup checks if the given key exists in the Database.
	// If the key exists, the return value should be true with a possibility of afalse positives.
	//  If the key does not exist, the return value has to be false.
	Lookup(key uint32) bool

	// Put inserts a key into the Filter.
	Put(key uint32) error

	// Delete removes the given key from the Filter.
	Delete(key uint32)
}

//The control functions from the last lab can aslo be used here:
//expect the Cuckoo Bloom filter to be created in memory, possiby written to a file and deleted
type ControlDatabase interface {

	// Open opens a KV database with the given name.
	Open(name string) error

	// Close closes the KV database.
	Close(name string) error

	// Create creates a new KV database with the given name.
	// It returns an error if the database already exists.
	Create(name string) error

	// DeleteDB deletes the KV database with the given name.
	DeleteDB(name string) error
}
