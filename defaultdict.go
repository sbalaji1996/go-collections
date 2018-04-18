package main

import "fmt"

type defaultdict struct {
    default_val interface{}
    items map[interface{}]interface{}
}

func create(default_val interface{}) *defaultdict {
    return &defaultdict{default_val, make(map[interface{}]interface{})}
}

func (d *defaultdict) set(key, value interface{}) {
    d.items[key] = value
}

func (d *defaultdict) get(key interface{}) interface{} {
    val, present := d.items[key]
    if present == true {
        return val
    } else {
        return d.default_val
    }
}

func (d *defaultdict) delete(key interface{}) {
    d.items[key] = d.default_val
}

func main() {
    d := create(0)
    fmt.Println(d.get(1))
    d.set(1, 2)
    fmt.Println(d.get(1))
    d.set('a', 2)
    fmt.Println(d.get('a'))
    d.delete('a')
    fmt.Println(d.get('a'))
}