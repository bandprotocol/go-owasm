#[derive(Copy, Clone)]
#[repr(C)]
pub struct Span {
    pub ptr: *mut u8,
    pub len: usize,
    pub cap: usize,
}

impl Span {
    // TODO
    pub fn create_readonly(data: &[u8]) -> Span {
        Span {
            ptr: data.as_ptr() as *mut u8,
            len: data.len(),
            cap: data.len(),
        }
    }

    /// TODO
    pub fn read(&self) -> &[u8] {
        unsafe { std::slice::from_raw_parts(self.ptr, self.len) }
    }

    /// TODO
    pub fn write(&mut self, data: &[u8]) {
        // TODO: Do not allow write if data.len() exceeds cap.
        self.len += data.len();
        unsafe { std::ptr::copy(data.as_ptr(), self.ptr, data.len()) }
    }
}

#[cfg(test)]
mod test {
    // use super::*;
    // TODO
}
