#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct {
  uint8_t *ptr;
  uintptr_t len;
  uintptr_t cap;
} Span;

typedef struct {
  uint8_t _private[0];
} env_t;

typedef struct {
  int64_t (*get_ask_count)(const env_t*);
} EnvDispatcher;

typedef struct {
  env_t *env;
  EnvDispatcher dis;
} Env;

void do_compile(Span input, Span *output);

int32_t do_run(Span code, Env env);
