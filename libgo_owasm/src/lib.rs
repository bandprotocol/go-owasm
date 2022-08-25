mod env;
mod span;
mod vm;

use env::{Env, RunOutput};

use span::Span;

use failure::{bail, Error as FailureError};
use std::panic::catch_unwind;

use owasm_vm;
use owasm_vm::cache::{Cache, CacheOptions};
use owasm_vm::error::Error;

// Cache initializing section
#[repr(C)]
pub struct cache_t {}

pub fn to_cache(ptr: *mut cache_t) -> Option<&'static mut Cache> {
    if ptr.is_null() {
        None
    } else {
        let c = unsafe { &mut *(ptr as *mut Cache) };
        Some(c)
    }
}

#[no_mangle]
pub extern "C" fn init_cache(size: u32) -> *mut cache_t {
    let r = catch_unwind(|| do_init_cache(size)).unwrap_or_else(|_| bail!("Caught panic"));
    match r {
        Ok(t) => t as *mut cache_t,
        Err(_) => std::ptr::null_mut(),
    }
}

fn do_init_cache(size: u32) -> Result<*mut Cache, FailureError> {
    let cache = Cache::new(CacheOptions { cache_size: size });
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
pub extern "C" fn do_compile(input: Span, output: &mut Span) -> Error {
    match owasm_vm::compile(input.read()) {
        Ok(out) => {
            output.write(&out);
            Error::NoError
        }
        Err(e) => e,
    }
}

#[no_mangle]
pub extern "C" fn do_run(
    cache: *mut cache_t,
    code: Span,
    gas_limit: u64,
    span_size: i64,
    is_prepare: bool,
    env: Env,
    output: &mut RunOutput,
) -> Error {
    if !cache.is_null() {
        let vm_env = vm::VMEnv::new(env, span_size);
        match owasm_vm::run(to_cache(cache).unwrap(), code.read(), gas_limit, is_prepare, vm_env) {
            Ok(gas_used) => {
                output.gas_used = gas_used;
                Error::NoError
            }
            Err(e) => e,
        }
    } else {
        Error::UnknownError
    }
}
