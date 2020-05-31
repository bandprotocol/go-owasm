use crate::env::Env;
use crate::span::Span;

pub struct VMLogic {
    env: Env,
}

impl VMLogic {
    pub fn new(env: Env) -> VMLogic {
        VMLogic { env: env }
    }

    pub fn get_calldata(&self) -> Span {
        (self.env.dis.get_calldata)(self.env.env)
    }

    pub fn set_return_data(&self, data: &[u8]) {
        (self.env.dis.set_return_data)(self.env.env, Span::create(data))
    }

    pub fn get_ask_count(&self) -> i64 {
        (self.env.dis.get_ask_count)(self.env.env)
    }

    pub fn get_min_count(&self) -> i64 {
        (self.env.dis.get_min_count)(self.env.env)
    }

    pub fn get_ans_count(&self) -> i64 {
        (self.env.dis.get_ans_count)(self.env.env)
    }
}
