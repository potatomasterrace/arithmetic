
#ifndef ARTITHMETIC_MEMORY_MANAGEMENT_H
#define ARTITHMETIC_MEMORY_MANAGEMENT_H
#include <malloc.h>
#include <stdio.h>
#include <stdint.h>

#if UINTPTR_MAX == 0xffffffff
#define PTR_ALLOCATION_MESSAGE "Pointer malloc(ed) (%s): 0x%08lx\r\n"
#define PTR_FREE_MESSAGE "Pointer free(ed) (%s):   0x%08lx\r\n"
#elif UINTPTR_MAX == 0xffffffffffffffff
#define PTR_ALLOCATION_MESSAGE "Pointer malloc(ed) (%s): 0x%016lx\r\n"
#define PTR_FREE_MESSAGE "Pointer free(ed) (%s):   0x%016lx\r\n"
#endif

#define malloc_ptr(type,ptr)\

// Macros for debugging memory management 
#define MALLOC_PROXY(type,typename,var, size)             \
    type var = (type) malloc(size);                       \
    printf(PTR_ALLOCATION_MESSAGE, typename,(void*) var); \


#define FREE_PROXY(type, ptr)                 \
  ({                                          \
  printf(PTR_FREE_MESSAGE, type,(void*) ptr); \
  free(ptr);                                  \
  })                                          \


#endif /* ARTITHMETIC_MEMORY_MANAGEMENT_H */