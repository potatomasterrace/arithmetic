#include "./memory.h"

void* malloc_proxy(size_t size){
    ptr = malloc(size)
    printf('\r\n%d: pointer freed\r\n',ptr)
    return ptr;
}

void free_proxy(void* ptr){
    printf('\r\n%d: pointer freed\r\n',ptr)
    free(ptr);
}
