mod env;
mod span;
mod vm;

use env::{Env, RunOutput};
use owasm::core;
use span::Span;

use failure::{bail, Error};
use std::panic::{catch_unwind};

use owasm::core::cache::{Cache, CacheOptions};

// Cache initializing section
#[repr(C)]
pub struct cache_t {}

#[no_mangle]
pub extern "C" fn init_cache(size: u32) -> *mut cache_t {
    let r = catch_unwind(|| do_init_cache(size)).unwrap_or_else(|_| bail!("Caught panic"));
    match r {
        Ok(t) => t as *mut cache_t,
        Err(_) => std::ptr::null_mut()
    }
}

fn do_init_cache(size: u32) -> Result<*mut Cache, Error> {
    let cache = Cache::new( CacheOptions { cache_size: size });
    let out = Box::new(cache);
    let res = Ok(Box::into_raw(out));
    res
}

#[no_mangle]
pub unsafe extern "C" fn release_cache(cache: *mut cache_t) {
    if !cache.is_null() {
        // this will free cache when it goes out of scope
        let _ = Box::from_raw(cache as *mut Cache);
    }
}

// Compile and execute section
#[no_mangle]
pub extern "C" fn do_compile(input: Span, output: &mut Span) -> owasm::core::error::Error {
    match core::compile(input.read()) {
        Ok(out) => {
            output.write(&out);
            owasm::core::error::Error::NoError
        }
        Err(e) => e,
    }
}

#[no_mangle]
pub extern "C" fn do_run(
    cache: &mut Cache,
    code: Span,
    gas_limit: u32,
    span_size: i64,
    is_prepare: bool,
    env: Env,
    output: &mut RunOutput,
) -> owasm::core::error::Error {
    let vm_env = vm::VMEnv::new(env, span_size);
    match core::run(cache, code.read(), gas_limit, is_prepare, vm_env) {
        Ok(gas_used) => {
            output.gas_used = gas_used;
            owasm::core::error::Error::NoError
        }
        Err(e) => e,
    }
}
