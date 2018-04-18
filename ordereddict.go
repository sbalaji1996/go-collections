package main

import "fmt"

type ordereddict struct {
    ordering []interface{}
    items map[interface{}]interface{}
}

func create() *ordereddict {
    return &ordereddict{make([]interface{}, 0), make(map[interface{}]interface{})}
}

func (o *ordereddict) add(key, value interface{}) {
    o.items[key] = value
    o.ordering = append(o.ordering, key)
}

func (o *ordereddict) get(key interface{}) interface{} {
    val, present := o.items[key]
    if present == true {
        return val
    } else {
        return nil
    }
}

func (o *ordereddict) get_ordering() []interface{} {
    return o.ordering
}

func (o *ordereddict) return_in_order() []interface{} {
    res := make([]interface{}, len(o.ordering))
    for index, key := range o.ordering {
        res[index] = o.items[key]
    }
    return res
}

func main() {
    o := create()
    o.add('a', 1)
    o.add('b', 2)
    o.add(97, 3)
    fmt.Println(o.get_ordering()[0])
    fmt.Println(o.get(o.get_ordering()[0]))
    fmt.Println(o.get(o.get_ordering()[2]))

    fmt.Println(o.return_in_order())
}