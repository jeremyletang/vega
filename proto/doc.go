// Package proto contains protocol buffers definitions and autogenerated files.
//
// In order to add a new API endpoint, add a new rpc entry as well as request
// and response message types. Example:
//
//     service trading {
//         // ...
//         rpc SomeNewEndpoint(SomeNewEndpointRequest) returns (SomeNewEndpointResponse);
//         // ...
//     }
//
//     message SomeNewEndpointRequest {
//         string somestr = 1;
//         int64 someint = 2;
//     }
//
//     message SomeNewEndpointResponse {
//         string someanswer = 1;
//         repeated string somestringlist = 2;
//     }
package proto