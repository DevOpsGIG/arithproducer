# arithproducer

The Arithmetic program generates a never-ending (infinite loop) of random basic arithmetic tasks of type publisher.Task.

An example of task struct created is:

```
arith := &publisher.Arith{
  Meta:  publisher.Meta{
    TaskType: "arithmetic",
  },
  A:        10,
  B:        20,
  Operator: "+",
}
```

Every task is created on a new goroutine by the `Assembly` function.
Newly generated tasks then concurrently calls the `Send` function from the `publish package` passing the job thought the task channel
and the expecting the results over the response channel.
