/*
Object Pool Pattern:
This is a creational design pattern in which a pool of objects is initialized and created beforehand and keep it in pool.
As and when needed a client can request an object from the pool use it and return it to the pool.
Basicaclly, pool is a container which contains some amount of objects.
So, when an object is taken from the pool, it is not avilable in the pool until it is put back.
3 major characteristics of this pattern are:
1. Client
2. Reusable Object
3. Object Pool

Common Usage of object pool pattern:
Threads -: creating a thread is an extremely expensive operation object pool pattern can be used to reuse threads.
Database conenctions -: creating a single connection is time consuming object pool pattern can be used to reuse connections.
Socket connections -: creating a single connection is time consuming object pool pattern can be used to reuse connections.
Large Graphics object -: as bitmaps, fonts etc. are also faster to reload that create from scratch.

WHY:
Object pool design pattern is used when the cost to create the object of the class is high
and the number of such objects that will be needed at a paticular time is limited.
*/

package objectpool

import (
	"fmt"
	"strconv"
)

type iDBConnection interface {
	getID() string
}

type connection struct {
	id string
}

func (c *connection) getID() string {
	return c.id
}

type poolConnections struct {
	idle     []iDBConnection
	active   []iDBConnection
	capacity int
}

func (p *poolConnections) getConnection() iDBConnection {
	if len(p.idle) == 0 {
		panic("no connection present in the pool")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("get pool connection with id: %s\n", obj.getID())
	return obj
}

func (p *poolConnections) returnConnection(target iDBConnection) error {
	err := p.remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Printf("return pool object with ID: %s\n", target.getID())
	return nil
}

func (p *poolConnections) remove(target iDBConnection) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getID() == target.getID() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("target pool object doesn't belong to the pool")
}

func initPool(poolObjects []iDBConnection) *poolConnections {
	if len(poolObjects) == 0 {
		return nil
	}
	active := make([]iDBConnection, 0)
	pool := &poolConnections{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
	}
	return pool
}

func DemoObjectPoolPattern() {
	connections := make([]iDBConnection, 0)
	for i := 0; i < 3; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool := initPool(connections)
	fmt.Println("Initial connection idle count", len(pool.idle))
	fmt.Println("Initial connection active count", len(pool.active))

	// Decrease connection counter
	conn1 := pool.getConnection()
	fmt.Println("connection count after taking first connection out", len(pool.idle))
	conn2 := pool.getConnection()
	fmt.Println("connection count after taking second connection out", len(pool.idle))
	conn3 := pool.getConnection()
	fmt.Println("connection count after taking third connection out", len(pool.idle))

	// Increase connection counter
	pool.returnConnection(conn1)
	fmt.Println("connection count after returning first connection", len(pool.idle))
	pool.returnConnection(conn2)
	fmt.Println("connection count after returning second connection", len(pool.idle))
	pool.returnConnection(conn3)
	fmt.Println("connection count after returning third connection", len(pool.idle))

	fmt.Println("DemoObjectPoolPattern finished")
}