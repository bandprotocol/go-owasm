package api

// #include "bindings.h"
//
// Error cGetCalldata(env_t *e, Span *calldata);
// Error cGetCalldata_cgo(env_t *e, Span *calldata) { return cGetCalldata(e, calldata); }
// Error cSetReturnData(env_t *e, Span data);
// Error cSetReturnData_cgo(env_t *e, Span data) { return cSetReturnData(e, data); }
// uint64_t cGetAskCount(env_t *e);
// uint64_t cGetAskCount_cgo(env_t *e) { return cGetAskCount(e); }
// uint64_t cGetMinCount(env_t *e);
// uint64_t cGetMinCount_cgo(env_t *e) { return cGetMinCount(e); }
// int64_t cGetPrepareTime(env_t *e);
// int64_t cGetPrepareTime_cgo(env_t *e) { return cGetPrepareTime(e); }
// Error cGetExecuteTime(env_t *e, int64_t *val);
// Error cGetExecuteTime_cgo(env_t *e, int64_t *val) { return cGetExecuteTime(e, val); }
// Error cGetAnsCount(env_t *e, uint64_t *val);
// Error cGetAnsCount_cgo(env_t *e, uint64_t *val) { return cGetAnsCount(e, val); }
// Error cAskExternalData(env_t *e, uint64_t eid, uint64_t did, Span data);
// Error cAskExternalData_cgo(env_t *e, uint64_t eid, uint64_t did, Span data) { return cAskExternalData(e, eid, did, data); }
// Error cGetExternalDataStatus(env_t *e, uint64_t eid, uint64_t vid, int64_t *status);
// Error cGetExternalDataStatus_cgo(env_t *e, uint64_t eid, uint64_t vid, int64_t *status) { return cGetExternalDataStatus(e, eid, vid, status); }
// Error cGetExternalData(env_t *e, uint64_t eid, uint64_t vid, Span *data);
// Error cGetExternalData_cgo(env_t *e, uint64_t eid, uint64_t vid, Span *data) { return cGetExternalData(e, eid, vid, data); }
import "C"
