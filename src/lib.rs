mod env;
mod span;
mod vm;

use env::Env;
use parity_wasm::elements::{self};
use pwasm_utils::{self, rules};
use span::Span;
use std::ffi::c_void;
use wasmer_runtime::{instantiate, Ctx};
use wasmer_runtime_core::{func, imports, wasmparser, Func};

#[no_mangle]
pub extern "C" fn do_compile(input: Span, output: &mut Span) {
    output.write(&compile(input.read()));
}

#[no_mangle]
pub extern "C" fn do_run(code: Span, env: Env) -> i32 {
    run(code.read(), env)
}

fn compile(code: &[u8]) -> Vec<u8> {
    // Check that the given Wasm code is indeed a valid Wasm.
    wasmparser::validate(code, None).unwrap();
    // Simple gas rule. Every opcode and memory growth costs 1 gas.
    let gas_rules = rules::Set::new(1, Default::default()).with_grow_cost(1);
    // Start the compiling chains. TODO: Add more safeguards.
    let module = elements::deserialize_buffer(code).unwrap();
    let module = pwasm_utils::inject_gas_counter(module, &gas_rules).unwrap();
    // Serialize the final Wasm code back to bytes.
    elements::serialize(module).unwrap()
}

struct ImportReference(*mut c_void);
unsafe impl Send for ImportReference {}
unsafe impl Sync for ImportReference {}

fn run(code: &[u8], env: Env) -> i32 {
    let vm = &mut vm::VMLogic::new(env);
    let raw_ptr = vm as *mut _ as *mut c_void;
    let import_reference = ImportReference(raw_ptr);
    let import_object = imports! {
        move || {
            let dtor = (|_: *mut c_void| {}) as fn(*mut c_void);
            (import_reference.0, dtor)
        },
        "env" => {
            "gas" => func!(|_: &mut Ctx, n: u32| {
                println!("HEY {:?}", n)
            }),
            "get_ask_count" => func!(|ctx: &mut Ctx| {
                let vm: &mut vm::VMLogic = unsafe { &mut *(ctx.data as *mut vm::VMLogic) };
                vm.get_ask_count()
            }),
        },
    };
    let instance = instantiate(code, &import_object).unwrap();
    let add_one: Func<(), i64> = instance.exports.get("prepare").unwrap();
    add_one.call().unwrap() as i32
}

// #[no_mangle]
// pub extern "C" fn compile(code: Buffer) -> Buffer {
//     let compiler = Box::new(StreamingCompiler::<SinglePassMCG, _, _, _, _>::new(
//         move || {
//             let mut chain = MiddlewareChain::new();
//             chain.push(metering::Metering::new(1000));
//             chain
//         },
//     ));
//     let module = compile_with(code.read(), compiler.as_ref()).unwrap();
//     let cache = module.cache().unwrap();
//     Buffer::from_vec(cache.serialize().unwrap())
// }

// pub fn run(code: &[u8]) -> Result<i32, wasmer_runtime_core::error::RuntimeError> {
//     let import_object = imports! {
//         "env" => {
//             "gas" => func!(|_: &mut Ctx, n: u32| {
//                 println!("HEY {:?}", n)
//             }),
//         },
//     };

//     let better_code = do_compile(code);
//     let instance = instantiate(&better_code, &import_object).unwrap();
//     let add_one: Func<(), i32> = instance.exports.get("prepare").unwrap();
//     add_one.call()
// }

// #[cfg(test)]
// mod test {
//     use super::*;

//     static EXAMPLE_WASM: &'static [u8] = include_bytes!("ez.wasm");
//     static EXAMPLE_WASM2: &'static [u8] = include_bytes!("ez2.wasm");

//     #[test]
//     fn test_run_ez() {
//         println!("{:?}", run(EXAMPLE_WASM))
//     }

//     #[test]
//     fn test_run_ez2() {
//         println!("{:?}", run(EXAMPLE_WASM2))
//     }
// }
