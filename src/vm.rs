use crate::env::Env;
use std::ffi::c_void;

pub struct VMLogic {
    env: Env,
}

impl VMLogic {
    pub fn new(env: Env) -> VMLogic {
        VMLogic { env: env }
    }

    pub fn get_ask_count(&self) -> Result<i64, ()> {
        Ok((self.env.dis.get_ask_count)(self.env.env))
    }
}

// pub fn create_vm_state(env: Env) -> *mut c_void {
//     let state = Box::new(VMLogic::new(env));
//     Box::into_raw(state) as *mut c_void
// }

// pub fn destroy_vm_state(ptr: *mut c_void) {
//     let b = unsafe { Box::from_raw(ptr as *mut VMLogic) };
//     std::mem::drop(b);
// }
