
#ifndef ARTITHMETIC_MEMORY_H
#define ARTITHMETIC_MEMORY_H
#include <malloc.h>
#include <stdio.h>


void* malloc_proxy(size_t size){
    return malloc(size);
}


void free_proxy(void* ptr){
    free(ptr);
}
#endif /* ARTITHMETIC_MEMORY_H */