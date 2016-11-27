// #include "myheader.h"
import "C"

func myInit() {
    var myContext C.contextType
    var stashedThing C.stashType

    // Stashes a pointer to stashedThing in some private 
    // internal struct, for later use.
    C.init(&myContext, &stashedThing)

    // Stuff happens here

    // Probably runs fine in 1.4, but has a non-deterministic
    // stack-bounds panic in 1.5.
    C.process(&myContext)

    // ...
}
