# nil-pointer

Tool to automatically generate Unit-Tests to find nil pointer errors.

Example for nil pointer errors:

* Provide nil pointer to func, whiche accepts pointer argument (e.g. to struct)
* Function, which accepts variadic arguments, where one is nil

To catch the nil pointer panics, every function call has to be wraped in a special crafted function, which recovers the panics and is able to call arbitary functions (e.g with differen arguments and types)