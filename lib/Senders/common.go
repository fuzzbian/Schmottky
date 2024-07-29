package common

var PointChannel = make(chan complex128, 1e7)
var LevChannel = make(chan int, 1e7)
var LevTracking bool = false