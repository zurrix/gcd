// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Runtime functionality.
// API Version: 1.1

package gcdapi

import (
	"encoding/json"
	"github.com/zurrix/gcd/gcdmessage"
)

// Mirror object referencing original JavaScript object.
type RuntimeRemoteObject struct {
	Type          string                `json:"type"`                    // Object type.
	Subtype       string                `json:"subtype,omitempty"`       // Object subtype hint. Specified for <code>object</code> type values only.
	ClassName     string                `json:"className,omitempty"`     // Object class (constructor) name. Specified for <code>object</code> type values only.
	Value         string                `json:"value,omitempty"`         // Remote object value in case of primitive values or JSON values (if it was requested), or description string if the value can not be JSON-stringified (like NaN, Infinity, -Infinity, -0).
	Description   string                `json:"description,omitempty"`   // String representation of the object.
	ObjectId      string                `json:"objectId,omitempty"`      // Unique object identifier (for non-primitive values).
	Preview       *RuntimeObjectPreview `json:"preview,omitempty"`       // Preview containing abbreviated property values. Specified for <code>object</code> type values only.
	CustomPreview *RuntimeCustomPreview `json:"customPreview,omitempty"` //
}

// No Description.
type RuntimeCustomPreview struct {
	Header                     string `json:"header"`                     //
	HasBody                    bool   `json:"hasBody"`                    //
	FormatterObjectId          string `json:"formatterObjectId"`          //
	BindRemoteObjectFunctionId string `json:"bindRemoteObjectFunctionId"` //
	ConfigObjectId             string `json:"configObjectId,omitempty"`   //
}

// Object containing abbreviated remote object value.
type RuntimeObjectPreview struct {
	Type        string                    `json:"type"`                  // Object type.
	Subtype     string                    `json:"subtype,omitempty"`     // Object subtype hint. Specified for <code>object</code> type values only.
	Description string                    `json:"description,omitempty"` // String representation of the object.
	Overflow    bool                      `json:"overflow"`              // True iff some of the properties or entries of the original object did not fit.
	Properties  []*RuntimePropertyPreview `json:"properties"`            // List of the properties.
	Entries     []*RuntimeEntryPreview    `json:"entries,omitempty"`     // List of the entries. Specified for <code>map</code> and <code>set</code> subtype values only.
}

// No Description.
type RuntimePropertyPreview struct {
	Name         string                `json:"name"`                   // Property name.
	Type         string                `json:"type"`                   // Object type. Accessor means that the property itself is an accessor property.
	Value        string                `json:"value,omitempty"`        // User-friendly property value string.
	ValuePreview *RuntimeObjectPreview `json:"valuePreview,omitempty"` // Nested value preview.
	Subtype      string                `json:"subtype,omitempty"`      // Object subtype hint. Specified for <code>object</code> type values only.
}

// No Description.
type RuntimeEntryPreview struct {
	Key   *RuntimeObjectPreview `json:"key,omitempty"` // Preview of the key. Specified for map-like collection entries.
	Value *RuntimeObjectPreview `json:"value"`         // Preview of the value.
}

// Object property descriptor.
type RuntimePropertyDescriptor struct {
	Name         string               `json:"name"`                // Property name or symbol description.
	Value        *RuntimeRemoteObject `json:"value,omitempty"`     // The value associated with the property.
	Writable     bool                 `json:"writable,omitempty"`  // True if the value associated with the property may be changed (data descriptors only).
	Get          *RuntimeRemoteObject `json:"get,omitempty"`       // A function which serves as a getter for the property, or <code>undefined</code> if there is no getter (accessor descriptors only).
	Set          *RuntimeRemoteObject `json:"set,omitempty"`       // A function which serves as a setter for the property, or <code>undefined</code> if there is no setter (accessor descriptors only).
	Configurable bool                 `json:"configurable"`        // True if the type of this property descriptor may be changed and if the property may be deleted from the corresponding object.
	Enumerable   bool                 `json:"enumerable"`          // True if this property shows up during enumeration of the properties on the corresponding object.
	WasThrown    bool                 `json:"wasThrown,omitempty"` // True if the result was thrown during the evaluation.
	IsOwn        bool                 `json:"isOwn,omitempty"`     // True if the property is owned for the object.
	Symbol       *RuntimeRemoteObject `json:"symbol,omitempty"`    // Property symbol object, if the property is of the <code>symbol</code> type.
}

// Object internal property descriptor. This property isn't normally visible in JavaScript code.
type RuntimeInternalPropertyDescriptor struct {
	Name  string               `json:"name"`            // Conventional property name.
	Value *RuntimeRemoteObject `json:"value,omitempty"` // The value associated with the property.
}

// Represents function call argument. Either remote object id <code>objectId</code> or primitive <code>value</code> or neither of (for undefined) them should be specified.
type RuntimeCallArgument struct {
	Value    string `json:"value,omitempty"`    // Primitive value, or description string if the value can not be JSON-stringified (like NaN, Infinity, -Infinity, -0).
	ObjectId string `json:"objectId,omitempty"` // Remote object handle.
	Type     string `json:"type,omitempty"`     // Object type.
}

// Description of an isolated world.
type RuntimeExecutionContextDescription struct {
	Id        int    `json:"id"`        // Unique id of the execution context. It can be used to specify in which execution context script evaluation should be performed.
	IsDefault bool   `json:"isDefault"` // Whether context is the default page context (as opposite to e.g. context of content script).
	Origin    string `json:"origin"`    // Execution context origin.
	Name      string `json:"name"`      // Human readable name describing given context.
	FrameId   string `json:"frameId"`   // Id of the owning frame. May be an empty string if the context is not associated with a frame.
}

// Detailed information on exception (or error) that was thrown during script compilation or execution.
type RuntimeExceptionDetails struct {
	Text     string             `json:"text"`               // Exception text.
	Url      string             `json:"url,omitempty"`      // URL of the message origin.
	ScriptId string             `json:"scriptId,omitempty"` // Script ID of the message origin.
	Line     int                `json:"line,omitempty"`     // Line number in the resource that generated this message.
	Column   int                `json:"column,omitempty"`   // Column number in the resource that generated this message.
	Stack    *RuntimeStackTrace `json:"stack,omitempty"`    // JavaScript stack trace for assertions and error messages.
}

// Stack entry for runtime errors and assertions.
type RuntimeCallFrame struct {
	FunctionName string `json:"functionName"` // JavaScript function name.
	ScriptId     string `json:"scriptId"`     // JavaScript script id.
	Url          string `json:"url"`          // JavaScript script name or url.
	LineNumber   int    `json:"lineNumber"`   // JavaScript script line number.
	ColumnNumber int    `json:"columnNumber"` // JavaScript script column number.
}

// Call frames for assertions or error messages.
type RuntimeStackTrace struct {
	Description string              `json:"description,omitempty"` // String label of this stack trace. For async traces this may be a name of the function that initiated the async call.
	CallFrames  []*RuntimeCallFrame `json:"callFrames"`            // JavaScript function name.
	Parent      *RuntimeStackTrace  `json:"parent,omitempty"`      // Asynchronous JavaScript stack trace that preceded this stack, if available.
}

// Issued when new execution context is created.
type RuntimeExecutionContextCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Context *RuntimeExecutionContextDescription `json:"context"` // A newly created execution contex.
	} `json:"Params,omitempty"`
}

// Issued when execution context is destroyed.
type RuntimeExecutionContextDestroyedEvent struct {
	Method string `json:"method"`
	Params struct {
		ExecutionContextId int `json:"executionContextId"` // Id of the destroyed context
	} `json:"Params,omitempty"`
}

//
type RuntimeInspectRequestedEvent struct {
	Method string `json:"method"`
	Params struct {
		Object *RuntimeRemoteObject   `json:"object"` //
		Hints  map[string]interface{} `json:"hints"`  //
	} `json:"Params,omitempty"`
}

type Runtime struct {
	target gcdmessage.ChromeTargeter
}

func NewRuntime(target gcdmessage.ChromeTargeter) *Runtime {
	c := &Runtime{target: target}
	return c
}

// Evaluate - Evaluates expression on global object.
// expression - Expression to evaluate.
// objectGroup - Symbolic group name that can be used to release multiple objects.
// includeCommandLineAPI - Determines whether Command Line API should be available during the evaluation.
// doNotPauseOnExceptionsAndMuteConsole - Specifies whether evaluation should stop on exceptions and mute console. Overrides setPauseOnException state.
// contextId - Specifies in which isolated context to perform evaluation. Each content script lives in an isolated context and this parameter may be used to specify one of those contexts. If the parameter is omitted or 0 the evaluation will be performed in the context of the inspected page.
// returnByValue - Whether the result is expected to be a JSON object that should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// userGesture - Whether execution should be treated as initiated by user in the UI.
// Returns -  result - Evaluation result. wasThrown - True if the result was thrown during the evaluation. exceptionDetails - Exception details.
func (c *Runtime) Evaluate(expression string, objectGroup string, includeCommandLineAPI bool, doNotPauseOnExceptionsAndMuteConsole bool, contextId int, returnByValue bool, generatePreview bool, userGesture bool) (*RuntimeRemoteObject, bool, *RuntimeExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 8)
	paramRequest["expression"] = expression
	paramRequest["objectGroup"] = objectGroup
	paramRequest["includeCommandLineAPI"] = includeCommandLineAPI
	paramRequest["doNotPauseOnExceptionsAndMuteConsole"] = doNotPauseOnExceptionsAndMuteConsole
	paramRequest["contextId"] = contextId
	paramRequest["returnByValue"] = returnByValue
	paramRequest["generatePreview"] = generatePreview
	paramRequest["userGesture"] = userGesture
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.evaluate", Params: paramRequest})
	if err != nil {
		return nil, false, nil, err
	}

	var chromeData struct {
		Result struct {
			Result           *RuntimeRemoteObject
			WasThrown        bool
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, false, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, false, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, false, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.WasThrown, chromeData.Result.ExceptionDetails, nil
}

// CallFunctionOn - Calls function with given declaration on the given object. Object group of the result is inherited from the target object.
// objectId - Identifier of the object to call function on.
// functionDeclaration - Declaration of the function to call.
// arguments - Call arguments. All call arguments must belong to the same JavaScript world as the target object.
// doNotPauseOnExceptionsAndMuteConsole - Specifies whether function call should stop on exceptions and mute console. Overrides setPauseOnException state.
// returnByValue - Whether the result is expected to be a JSON object which should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// userGesture - Whether execution should be treated as initiated by user in the UI.
// Returns -  result - Call result. wasThrown - True if the result was thrown during the evaluation.
func (c *Runtime) CallFunctionOn(objectId string, functionDeclaration string, arguments *RuntimeCallArgument, doNotPauseOnExceptionsAndMuteConsole bool, returnByValue bool, generatePreview bool, userGesture bool) (*RuntimeRemoteObject, bool, error) {
	paramRequest := make(map[string]interface{}, 7)
	paramRequest["objectId"] = objectId
	paramRequest["functionDeclaration"] = functionDeclaration
	paramRequest["arguments"] = arguments
	paramRequest["doNotPauseOnExceptionsAndMuteConsole"] = doNotPauseOnExceptionsAndMuteConsole
	paramRequest["returnByValue"] = returnByValue
	paramRequest["generatePreview"] = generatePreview
	paramRequest["userGesture"] = userGesture
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.callFunctionOn", Params: paramRequest})
	if err != nil {
		return nil, false, err
	}

	var chromeData struct {
		Result struct {
			Result    *RuntimeRemoteObject
			WasThrown bool
		}
	}

	if resp == nil {
		return nil, false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, false, err
	}

	return chromeData.Result.Result, chromeData.Result.WasThrown, nil
}

// GetProperties - Returns properties of a given object. Object group of the result is inherited from the target object.
// objectId - Identifier of the object to return properties for.
// ownProperties - If true, returns properties belonging only to the element itself, not to its prototype chain.
// accessorPropertiesOnly - If true, returns accessor properties (with getter/setter) only; internal properties are not returned either.
// generatePreview - Whether preview should be generated for the results.
// Returns -  result - Object properties. internalProperties - Internal object properties (only of the element itself). exceptionDetails - Exception details.
func (c *Runtime) GetProperties(objectId string, ownProperties bool, accessorPropertiesOnly bool, generatePreview bool) ([]*RuntimePropertyDescriptor, []*RuntimeInternalPropertyDescriptor, *RuntimeExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["objectId"] = objectId
	paramRequest["ownProperties"] = ownProperties
	paramRequest["accessorPropertiesOnly"] = accessorPropertiesOnly
	paramRequest["generatePreview"] = generatePreview
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.getProperties", Params: paramRequest})
	if err != nil {
		return nil, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			Result             []*RuntimePropertyDescriptor
			InternalProperties []*RuntimeInternalPropertyDescriptor
			ExceptionDetails   *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.InternalProperties, chromeData.Result.ExceptionDetails, nil
}

// ReleaseObject - Releases remote object with given id.
// objectId - Identifier of the object to release.
func (c *Runtime) ReleaseObject(objectId string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectId"] = objectId
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.releaseObject", Params: paramRequest})
}

// ReleaseObjectGroup - Releases all remote objects that belong to a given group.
// objectGroup - Symbolic object group name.
func (c *Runtime) ReleaseObjectGroup(objectGroup string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["objectGroup"] = objectGroup
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.releaseObjectGroup", Params: paramRequest})
}

// Tells inspected instance(worker or page) that it can run in case it was started paused.
func (c *Runtime) Run() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.run"})
}

// Enables reporting of execution contexts creation by means of <code>executionContextCreated</code> event. When the reporting gets enabled the event will be sent immediately for each existing execution context.
func (c *Runtime) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.enable"})
}

// Disables reporting of execution contexts creation.
func (c *Runtime) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.disable"})
}

// SetCustomObjectFormatterEnabled -
// enabled -
func (c *Runtime) SetCustomObjectFormatterEnabled(enabled bool) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["enabled"] = enabled
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.setCustomObjectFormatterEnabled", Params: paramRequest})
}

// CompileScript - Compiles expression.
// expression - Expression to compile.
// sourceURL - Source url to be set for the script.
// persistScript - Specifies whether the compiled script should be persisted.
// executionContextId - Specifies in which isolated context to perform script run. Each content script lives in an isolated context and this parameter is used to specify one of those contexts.
// Returns -  scriptId - Id of the script. exceptionDetails - Exception details.
func (c *Runtime) CompileScript(expression string, sourceURL string, persistScript bool, executionContextId int) (string, *RuntimeExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["expression"] = expression
	paramRequest["sourceURL"] = sourceURL
	paramRequest["persistScript"] = persistScript
	paramRequest["executionContextId"] = executionContextId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.compileScript", Params: paramRequest})
	if err != nil {
		return "", nil, err
	}

	var chromeData struct {
		Result struct {
			ScriptId         string
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return "", nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", nil, err
	}

	return chromeData.Result.ScriptId, chromeData.Result.ExceptionDetails, nil
}

// RunScript - Runs script with given id in a given context.
// scriptId - Id of the script to run.
// executionContextId - Specifies in which isolated context to perform script run. Each content script lives in an isolated context and this parameter is used to specify one of those contexts.
// objectGroup - Symbolic group name that can be used to release multiple objects.
// doNotPauseOnExceptionsAndMuteConsole - Specifies whether script run should stop on exceptions and mute console. Overrides setPauseOnException state.
// includeCommandLineAPI - Determines whether Command Line API should be available during the evaluation.
// Returns -  result - Run result. exceptionDetails - Exception details.
func (c *Runtime) RunScript(scriptId string, executionContextId int, objectGroup string, doNotPauseOnExceptionsAndMuteConsole bool, includeCommandLineAPI bool) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["scriptId"] = scriptId
	paramRequest["executionContextId"] = executionContextId
	paramRequest["objectGroup"] = objectGroup
	paramRequest["doNotPauseOnExceptionsAndMuteConsole"] = doNotPauseOnExceptionsAndMuteConsole
	paramRequest["includeCommandLineAPI"] = includeCommandLineAPI
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.runScript", Params: paramRequest})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			Result           *RuntimeRemoteObject
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.ExceptionDetails, nil
}
