"use strict";
(function() {

Error.stackTraceLimit = Infinity;

var $global, $module;
if (typeof window !== "undefined") { /* web page */
  $global = window;
} else if (typeof self !== "undefined") { /* web worker */
  $global = self;
} else if (typeof global !== "undefined") { /* Node.js */
  $global = global;
  $global.require = require;
} else { /* others (e.g. Nashorn) */
  $global = this;
}

if ($global === undefined || $global.Array === undefined) {
  throw new Error("no global object found");
}
if (typeof module !== "undefined") {
  $module = module;
}

var $packages = {}, $idCounter = 0;
var $keys = function(m) { return m ? Object.keys(m) : []; };
var $flushConsole = function() {};
var $throwRuntimeError; /* set by package "runtime" */
var $throwNilPointerError = function() { $throwRuntimeError("invalid memory address or nil pointer dereference"); };
var $call = function(fn, rcvr, args) { return fn.apply(rcvr, args); };
var $makeFunc = function(fn) { return function() { return $externalize(fn(this, new ($sliceType($jsObjectPtr))($global.Array.prototype.slice.call(arguments, []))), $emptyInterface); }; };
var $unused = function(v) {};

var $mapArray = function(array, f) {
  var newArray = new array.constructor(array.length);
  for (var i = 0; i < array.length; i++) {
    newArray[i] = f(array[i]);
  }
  return newArray;
};

var $methodVal = function(recv, name) {
  var vals = recv.$methodVals || {};
  recv.$methodVals = vals; /* noop for primitives */
  var f = vals[name];
  if (f !== undefined) {
    return f;
  }
  var method = recv[name];
  f = function() {
    $stackDepthOffset--;
    try {
      return method.apply(recv, arguments);
    } finally {
      $stackDepthOffset++;
    }
  };
  vals[name] = f;
  return f;
};

var $methodExpr = function(typ, name) {
  var method = typ.prototype[name];
  if (method.$expr === undefined) {
    method.$expr = function() {
      $stackDepthOffset--;
      try {
        if (typ.wrapped) {
          arguments[0] = new typ(arguments[0]);
        }
        return Function.call.apply(method, arguments);
      } finally {
        $stackDepthOffset++;
      }
    };
  }
  return method.$expr;
};

var $ifaceMethodExprs = {};
var $ifaceMethodExpr = function(name) {
  var expr = $ifaceMethodExprs["$" + name];
  if (expr === undefined) {
    expr = $ifaceMethodExprs["$" + name] = function() {
      $stackDepthOffset--;
      try {
        return Function.call.apply(arguments[0][name], arguments);
      } finally {
        $stackDepthOffset++;
      }
    };
  }
  return expr;
};

var $subslice = function(slice, low, high, max) {
  if (low < 0 || high < low || max < high || high > slice.$capacity || max > slice.$capacity) {
    $throwRuntimeError("slice bounds out of range");
  }
  var s = new slice.constructor(slice.$array);
  s.$offset = slice.$offset + low;
  s.$length = slice.$length - low;
  s.$capacity = slice.$capacity - low;
  if (high !== undefined) {
    s.$length = high - low;
  }
  if (max !== undefined) {
    s.$capacity = max - low;
  }
  return s;
};

var $substring = function(str, low, high) {
  if (low < 0 || high < low || high > str.length) {
    $throwRuntimeError("slice bounds out of range");
  }
  return str.substring(low, high);
};

var $sliceToArray = function(slice) {
  if (slice.$array.constructor !== Array) {
    return slice.$array.subarray(slice.$offset, slice.$offset + slice.$length);
  }
  return slice.$array.slice(slice.$offset, slice.$offset + slice.$length);
};

var $decodeRune = function(str, pos) {
  var c0 = str.charCodeAt(pos);

  if (c0 < 0x80) {
    return [c0, 1];
  }

  if (c0 !== c0 || c0 < 0xC0) {
    return [0xFFFD, 1];
  }

  var c1 = str.charCodeAt(pos + 1);
  if (c1 !== c1 || c1 < 0x80 || 0xC0 <= c1) {
    return [0xFFFD, 1];
  }

  if (c0 < 0xE0) {
    var r = (c0 & 0x1F) << 6 | (c1 & 0x3F);
    if (r <= 0x7F) {
      return [0xFFFD, 1];
    }
    return [r, 2];
  }

  var c2 = str.charCodeAt(pos + 2);
  if (c2 !== c2 || c2 < 0x80 || 0xC0 <= c2) {
    return [0xFFFD, 1];
  }

  if (c0 < 0xF0) {
    var r = (c0 & 0x0F) << 12 | (c1 & 0x3F) << 6 | (c2 & 0x3F);
    if (r <= 0x7FF) {
      return [0xFFFD, 1];
    }
    if (0xD800 <= r && r <= 0xDFFF) {
      return [0xFFFD, 1];
    }
    return [r, 3];
  }

  var c3 = str.charCodeAt(pos + 3);
  if (c3 !== c3 || c3 < 0x80 || 0xC0 <= c3) {
    return [0xFFFD, 1];
  }

  if (c0 < 0xF8) {
    var r = (c0 & 0x07) << 18 | (c1 & 0x3F) << 12 | (c2 & 0x3F) << 6 | (c3 & 0x3F);
    if (r <= 0xFFFF || 0x10FFFF < r) {
      return [0xFFFD, 1];
    }
    return [r, 4];
  }

  return [0xFFFD, 1];
};

var $encodeRune = function(r) {
  if (r < 0 || r > 0x10FFFF || (0xD800 <= r && r <= 0xDFFF)) {
    r = 0xFFFD;
  }
  if (r <= 0x7F) {
    return String.fromCharCode(r);
  }
  if (r <= 0x7FF) {
    return String.fromCharCode(0xC0 | r >> 6, 0x80 | (r & 0x3F));
  }
  if (r <= 0xFFFF) {
    return String.fromCharCode(0xE0 | r >> 12, 0x80 | (r >> 6 & 0x3F), 0x80 | (r & 0x3F));
  }
  return String.fromCharCode(0xF0 | r >> 18, 0x80 | (r >> 12 & 0x3F), 0x80 | (r >> 6 & 0x3F), 0x80 | (r & 0x3F));
};

var $stringToBytes = function(str) {
  var array = new Uint8Array(str.length);
  for (var i = 0; i < str.length; i++) {
    array[i] = str.charCodeAt(i);
  }
  return array;
};

var $bytesToString = function(slice) {
  if (slice.$length === 0) {
    return "";
  }
  var str = "";
  for (var i = 0; i < slice.$length; i += 10000) {
    str += String.fromCharCode.apply(undefined, slice.$array.subarray(slice.$offset + i, slice.$offset + Math.min(slice.$length, i + 10000)));
  }
  return str;
};

var $stringToRunes = function(str) {
  var array = new Int32Array(str.length);
  var rune, j = 0;
  for (var i = 0; i < str.length; i += rune[1], j++) {
    rune = $decodeRune(str, i);
    array[j] = rune[0];
  }
  return array.subarray(0, j);
};

var $runesToString = function(slice) {
  if (slice.$length === 0) {
    return "";
  }
  var str = "";
  for (var i = 0; i < slice.$length; i++) {
    str += $encodeRune(slice.$array[slice.$offset + i]);
  }
  return str;
};

var $copyString = function(dst, src) {
  var n = Math.min(src.length, dst.$length);
  for (var i = 0; i < n; i++) {
    dst.$array[dst.$offset + i] = src.charCodeAt(i);
  }
  return n;
};

var $copySlice = function(dst, src) {
  var n = Math.min(src.$length, dst.$length);
  $copyArray(dst.$array, src.$array, dst.$offset, src.$offset, n, dst.constructor.elem);
  return n;
};

var $copyArray = function(dst, src, dstOffset, srcOffset, n, elem) {
  if (n === 0 || (dst === src && dstOffset === srcOffset)) {
    return;
  }

  if (src.subarray) {
    dst.set(src.subarray(srcOffset, srcOffset + n), dstOffset);
    return;
  }

  switch (elem.kind) {
  case $kindArray:
  case $kindStruct:
    if (dst === src && dstOffset > srcOffset) {
      for (var i = n - 1; i >= 0; i--) {
        elem.copy(dst[dstOffset + i], src[srcOffset + i]);
      }
      return;
    }
    for (var i = 0; i < n; i++) {
      elem.copy(dst[dstOffset + i], src[srcOffset + i]);
    }
    return;
  }

  if (dst === src && dstOffset > srcOffset) {
    for (var i = n - 1; i >= 0; i--) {
      dst[dstOffset + i] = src[srcOffset + i];
    }
    return;
  }
  for (var i = 0; i < n; i++) {
    dst[dstOffset + i] = src[srcOffset + i];
  }
};

var $clone = function(src, type) {
  var clone = type.zero();
  type.copy(clone, src);
  return clone;
};

var $pointerOfStructConversion = function(obj, type) {
  if(obj.$proxies === undefined) {
    obj.$proxies = {};
    obj.$proxies[obj.constructor.string] = obj;
  }
  var proxy = obj.$proxies[type.string];
  if (proxy === undefined) {
    var properties = {};
    for (var i = 0; i < type.elem.fields.length; i++) {
      (function(fieldProp) {
        properties[fieldProp] = {
          get: function() { return obj[fieldProp]; },
          set: function(value) { obj[fieldProp] = value; }
        };
      })(type.elem.fields[i].prop);
    }
    proxy = Object.create(type.prototype, properties);
    proxy.$val = proxy;
    obj.$proxies[type.string] = proxy;
    proxy.$proxies = obj.$proxies;
  }
  return proxy;
};

var $append = function(slice) {
  return $internalAppend(slice, arguments, 1, arguments.length - 1);
};

var $appendSlice = function(slice, toAppend) {
  if (toAppend.constructor === String) {
    var bytes = $stringToBytes(toAppend);
    return $internalAppend(slice, bytes, 0, bytes.length);
  }
  return $internalAppend(slice, toAppend.$array, toAppend.$offset, toAppend.$length);
};

var $internalAppend = function(slice, array, offset, length) {
  if (length === 0) {
    return slice;
  }

  var newArray = slice.$array;
  var newOffset = slice.$offset;
  var newLength = slice.$length + length;
  var newCapacity = slice.$capacity;

  if (newLength > newCapacity) {
    newOffset = 0;
    newCapacity = Math.max(newLength, slice.$capacity < 1024 ? slice.$capacity * 2 : Math.floor(slice.$capacity * 5 / 4));

    if (slice.$array.constructor === Array) {
      newArray = slice.$array.slice(slice.$offset, slice.$offset + slice.$length);
      newArray.length = newCapacity;
      var zero = slice.constructor.elem.zero;
      for (var i = slice.$length; i < newCapacity; i++) {
        newArray[i] = zero();
      }
    } else {
      newArray = new slice.$array.constructor(newCapacity);
      newArray.set(slice.$array.subarray(slice.$offset, slice.$offset + slice.$length));
    }
  }

  $copyArray(newArray, array, newOffset + slice.$length, offset, length, slice.constructor.elem);

  var newSlice = new slice.constructor(newArray);
  newSlice.$offset = newOffset;
  newSlice.$length = newLength;
  newSlice.$capacity = newCapacity;
  return newSlice;
};

var $equal = function(a, b, type) {
  if (type === $jsObjectPtr) {
    return a === b;
  }
  switch (type.kind) {
  case $kindComplex64:
  case $kindComplex128:
    return a.$real === b.$real && a.$imag === b.$imag;
  case $kindInt64:
  case $kindUint64:
    return a.$high === b.$high && a.$low === b.$low;
  case $kindArray:
    if (a.length !== b.length) {
      return false;
    }
    for (var i = 0; i < a.length; i++) {
      if (!$equal(a[i], b[i], type.elem)) {
        return false;
      }
    }
    return true;
  case $kindStruct:
    for (var i = 0; i < type.fields.length; i++) {
      var f = type.fields[i];
      if (!$equal(a[f.prop], b[f.prop], f.typ)) {
        return false;
      }
    }
    return true;
  case $kindInterface:
    return $interfaceIsEqual(a, b);
  default:
    return a === b;
  }
};

var $interfaceIsEqual = function(a, b) {
  if (a === $ifaceNil || b === $ifaceNil) {
    return a === b;
  }
  if (a.constructor !== b.constructor) {
    return false;
  }
  if (a.constructor === $jsObjectPtr) {
    return a.object === b.object;
  }
  if (!a.constructor.comparable) {
    $throwRuntimeError("comparing uncomparable type " + a.constructor.string);
  }
  return $equal(a.$val, b.$val, a.constructor);
};

var $min = Math.min;
var $mod = function(x, y) { return x % y; };
var $parseInt = parseInt;
var $parseFloat = function(f) {
  if (f !== undefined && f !== null && f.constructor === Number) {
    return f;
  }
  return parseFloat(f);
};

var $froundBuf = new Float32Array(1);
var $fround = Math.fround || function(f) {
  $froundBuf[0] = f;
  return $froundBuf[0];
};

var $imul = Math.imul || function(a, b) {
  var ah = (a >>> 16) & 0xffff;
  var al = a & 0xffff;
  var bh = (b >>> 16) & 0xffff;
  var bl = b & 0xffff;
  return ((al * bl) + (((ah * bl + al * bh) << 16) >>> 0) >> 0);
};

var $floatKey = function(f) {
  if (f !== f) {
    $idCounter++;
    return "NaN$" + $idCounter;
  }
  return String(f);
};

var $flatten64 = function(x) {
  return x.$high * 4294967296 + x.$low;
};

var $shiftLeft64 = function(x, y) {
  if (y === 0) {
    return x;
  }
  if (y < 32) {
    return new x.constructor(x.$high << y | x.$low >>> (32 - y), (x.$low << y) >>> 0);
  }
  if (y < 64) {
    return new x.constructor(x.$low << (y - 32), 0);
  }
  return new x.constructor(0, 0);
};

var $shiftRightInt64 = function(x, y) {
  if (y === 0) {
    return x;
  }
  if (y < 32) {
    return new x.constructor(x.$high >> y, (x.$low >>> y | x.$high << (32 - y)) >>> 0);
  }
  if (y < 64) {
    return new x.constructor(x.$high >> 31, (x.$high >> (y - 32)) >>> 0);
  }
  if (x.$high < 0) {
    return new x.constructor(-1, 4294967295);
  }
  return new x.constructor(0, 0);
};

var $shiftRightUint64 = function(x, y) {
  if (y === 0) {
    return x;
  }
  if (y < 32) {
    return new x.constructor(x.$high >>> y, (x.$low >>> y | x.$high << (32 - y)) >>> 0);
  }
  if (y < 64) {
    return new x.constructor(0, x.$high >>> (y - 32));
  }
  return new x.constructor(0, 0);
};

var $mul64 = function(x, y) {
  var high = 0, low = 0;
  if ((y.$low & 1) !== 0) {
    high = x.$high;
    low = x.$low;
  }
  for (var i = 1; i < 32; i++) {
    if ((y.$low & 1<<i) !== 0) {
      high += x.$high << i | x.$low >>> (32 - i);
      low += (x.$low << i) >>> 0;
    }
  }
  for (var i = 0; i < 32; i++) {
    if ((y.$high & 1<<i) !== 0) {
      high += x.$low << i;
    }
  }
  return new x.constructor(high, low);
};

var $div64 = function(x, y, returnRemainder) {
  if (y.$high === 0 && y.$low === 0) {
    $throwRuntimeError("integer divide by zero");
  }

  var s = 1;
  var rs = 1;

  var xHigh = x.$high;
  var xLow = x.$low;
  if (xHigh < 0) {
    s = -1;
    rs = -1;
    xHigh = -xHigh;
    if (xLow !== 0) {
      xHigh--;
      xLow = 4294967296 - xLow;
    }
  }

  var yHigh = y.$high;
  var yLow = y.$low;
  if (y.$high < 0) {
    s *= -1;
    yHigh = -yHigh;
    if (yLow !== 0) {
      yHigh--;
      yLow = 4294967296 - yLow;
    }
  }

  var high = 0, low = 0, n = 0;
  while (yHigh < 2147483648 && ((xHigh > yHigh) || (xHigh === yHigh && xLow > yLow))) {
    yHigh = (yHigh << 1 | yLow >>> 31) >>> 0;
    yLow = (yLow << 1) >>> 0;
    n++;
  }
  for (var i = 0; i <= n; i++) {
    high = high << 1 | low >>> 31;
    low = (low << 1) >>> 0;
    if ((xHigh > yHigh) || (xHigh === yHigh && xLow >= yLow)) {
      xHigh = xHigh - yHigh;
      xLow = xLow - yLow;
      if (xLow < 0) {
        xHigh--;
        xLow += 4294967296;
      }
      low++;
      if (low === 4294967296) {
        high++;
        low = 0;
      }
    }
    yLow = (yLow >>> 1 | yHigh << (32 - 1)) >>> 0;
    yHigh = yHigh >>> 1;
  }

  if (returnRemainder) {
    return new x.constructor(xHigh * rs, xLow * rs);
  }
  return new x.constructor(high * s, low * s);
};

var $divComplex = function(n, d) {
  var ninf = n.$real === Infinity || n.$real === -Infinity || n.$imag === Infinity || n.$imag === -Infinity;
  var dinf = d.$real === Infinity || d.$real === -Infinity || d.$imag === Infinity || d.$imag === -Infinity;
  var nnan = !ninf && (n.$real !== n.$real || n.$imag !== n.$imag);
  var dnan = !dinf && (d.$real !== d.$real || d.$imag !== d.$imag);
  if(nnan || dnan) {
    return new n.constructor(NaN, NaN);
  }
  if (ninf && !dinf) {
    return new n.constructor(Infinity, Infinity);
  }
  if (!ninf && dinf) {
    return new n.constructor(0, 0);
  }
  if (d.$real === 0 && d.$imag === 0) {
    if (n.$real === 0 && n.$imag === 0) {
      return new n.constructor(NaN, NaN);
    }
    return new n.constructor(Infinity, Infinity);
  }
  var a = Math.abs(d.$real);
  var b = Math.abs(d.$imag);
  if (a <= b) {
    var ratio = d.$real / d.$imag;
    var denom = d.$real * ratio + d.$imag;
    return new n.constructor((n.$real * ratio + n.$imag) / denom, (n.$imag * ratio - n.$real) / denom);
  }
  var ratio = d.$imag / d.$real;
  var denom = d.$imag * ratio + d.$real;
  return new n.constructor((n.$imag * ratio + n.$real) / denom, (n.$imag - n.$real * ratio) / denom);
};

var $kindBool = 1;
var $kindInt = 2;
var $kindInt8 = 3;
var $kindInt16 = 4;
var $kindInt32 = 5;
var $kindInt64 = 6;
var $kindUint = 7;
var $kindUint8 = 8;
var $kindUint16 = 9;
var $kindUint32 = 10;
var $kindUint64 = 11;
var $kindUintptr = 12;
var $kindFloat32 = 13;
var $kindFloat64 = 14;
var $kindComplex64 = 15;
var $kindComplex128 = 16;
var $kindArray = 17;
var $kindChan = 18;
var $kindFunc = 19;
var $kindInterface = 20;
var $kindMap = 21;
var $kindPtr = 22;
var $kindSlice = 23;
var $kindString = 24;
var $kindStruct = 25;
var $kindUnsafePointer = 26;

var $methodSynthesizers = [];
var $addMethodSynthesizer = function(f) {
  if ($methodSynthesizers === null) {
    f();
    return;
  }
  $methodSynthesizers.push(f);
};
var $synthesizeMethods = function() {
  $methodSynthesizers.forEach(function(f) { f(); });
  $methodSynthesizers = null;
};

var $ifaceKeyFor = function(x) {
  if (x === $ifaceNil) {
    return 'nil';
  }
  var c = x.constructor;
  return c.string + '$' + c.keyFor(x.$val);
};

var $identity = function(x) { return x; };

var $typeIDCounter = 0;

var $idKey = function(x) {
  if (x.$id === undefined) {
    $idCounter++;
    x.$id = $idCounter;
  }
  return String(x.$id);
};

var $newType = function(size, kind, string, named, pkg, exported, constructor) {
  var typ;
  switch(kind) {
  case $kindBool:
  case $kindInt:
  case $kindInt8:
  case $kindInt16:
  case $kindInt32:
  case $kindUint:
  case $kindUint8:
  case $kindUint16:
  case $kindUint32:
  case $kindUintptr:
  case $kindUnsafePointer:
    typ = function(v) { this.$val = v; };
    typ.wrapped = true;
    typ.keyFor = $identity;
    break;

  case $kindString:
    typ = function(v) { this.$val = v; };
    typ.wrapped = true;
    typ.keyFor = function(x) { return "$" + x; };
    break;

  case $kindFloat32:
  case $kindFloat64:
    typ = function(v) { this.$val = v; };
    typ.wrapped = true;
    typ.keyFor = function(x) { return $floatKey(x); };
    break;

  case $kindInt64:
    typ = function(high, low) {
      this.$high = (high + Math.floor(Math.ceil(low) / 4294967296)) >> 0;
      this.$low = low >>> 0;
      this.$val = this;
    };
    typ.keyFor = function(x) { return x.$high + "$" + x.$low; };
    break;

  case $kindUint64:
    typ = function(high, low) {
      this.$high = (high + Math.floor(Math.ceil(low) / 4294967296)) >>> 0;
      this.$low = low >>> 0;
      this.$val = this;
    };
    typ.keyFor = function(x) { return x.$high + "$" + x.$low; };
    break;

  case $kindComplex64:
    typ = function(real, imag) {
      this.$real = $fround(real);
      this.$imag = $fround(imag);
      this.$val = this;
    };
    typ.keyFor = function(x) { return x.$real + "$" + x.$imag; };
    break;

  case $kindComplex128:
    typ = function(real, imag) {
      this.$real = real;
      this.$imag = imag;
      this.$val = this;
    };
    typ.keyFor = function(x) { return x.$real + "$" + x.$imag; };
    break;

  case $kindArray:
    typ = function(v) { this.$val = v; };
    typ.wrapped = true;
    typ.ptr = $newType(4, $kindPtr, "*" + string, false, "", false, function(array) {
      this.$get = function() { return array; };
      this.$set = function(v) { typ.copy(this, v); };
      this.$val = array;
    });
    typ.init = function(elem, len) {
      typ.elem = elem;
      typ.len = len;
      typ.comparable = elem.comparable;
      typ.keyFor = function(x) {
        return Array.prototype.join.call($mapArray(x, function(e) {
          return String(elem.keyFor(e)).replace(/\\/g, "\\\\").replace(/\$/g, "\\$");
        }), "$");
      };
      typ.copy = function(dst, src) {
        $copyArray(dst, src, 0, 0, src.length, elem);
      };
      typ.ptr.init(typ);
      Object.defineProperty(typ.ptr.nil, "nilCheck", { get: $throwNilPointerError });
    };
    break;

  case $kindChan:
    typ = function(v) { this.$val = v; };
    typ.wrapped = true;
    typ.keyFor = $idKey;
    typ.init = function(elem, sendOnly, recvOnly) {
      typ.elem = elem;
      typ.sendOnly = sendOnly;
      typ.recvOnly = recvOnly;
    };
    break;

  case $kindFunc:
    typ = function(v) { this.$val = v; };
    typ.wrapped = true;
    typ.init = function(params, results, variadic) {
      typ.params = params;
      typ.results = results;
      typ.variadic = variadic;
      typ.comparable = false;
    };
    break;

  case $kindInterface:
    typ = { implementedBy: {}, missingMethodFor: {} };
    typ.keyFor = $ifaceKeyFor;
    typ.init = function(methods) {
      typ.methods = methods;
      methods.forEach(function(m) {
        $ifaceNil[m.prop] = $throwNilPointerError;
      });
    };
    break;

  case $kindMap:
    typ = function(v) { this.$val = v; };
    typ.wrapped = true;
    typ.init = function(key, elem) {
      typ.key = key;
      typ.elem = elem;
      typ.comparable = false;
    };
    break;

  case $kindPtr:
    typ = constructor || function(getter, setter, target) {
      this.$get = getter;
      this.$set = setter;
      this.$target = target;
      this.$val = this;
    };
    typ.keyFor = $idKey;
    typ.init = function(elem) {
      typ.elem = elem;
      typ.wrapped = (elem.kind === $kindArray);
      typ.nil = new typ($throwNilPointerError, $throwNilPointerError);
    };
    break;

  case $kindSlice:
    typ = function(array) {
      if (array.constructor !== typ.nativeArray) {
        array = new typ.nativeArray(array);
      }
      this.$array = array;
      this.$offset = 0;
      this.$length = array.length;
      this.$capacity = array.length;
      this.$val = this;
    };
    typ.init = function(elem) {
      typ.elem = elem;
      typ.comparable = false;
      typ.nativeArray = $nativeArray(elem.kind);
      typ.nil = new typ([]);
    };
    break;

  case $kindStruct:
    typ = function(v) { this.$val = v; };
    typ.wrapped = true;
    typ.ptr = $newType(4, $kindPtr, "*" + string, false, pkg, exported, constructor);
    typ.ptr.elem = typ;
    typ.ptr.prototype.$get = function() { return this; };
    typ.ptr.prototype.$set = function(v) { typ.copy(this, v); };
    typ.init = function(pkgPath, fields) {
      typ.pkgPath = pkgPath;
      typ.fields = fields;
      fields.forEach(function(f) {
        if (!f.typ.comparable) {
          typ.comparable = false;
        }
      });
      typ.keyFor = function(x) {
        var val = x.$val;
        return $mapArray(fields, function(f) {
          return String(f.typ.keyFor(val[f.prop])).replace(/\\/g, "\\\\").replace(/\$/g, "\\$");
        }).join("$");
      };
      typ.copy = function(dst, src) {
        for (var i = 0; i < fields.length; i++) {
          var f = fields[i];
          switch (f.typ.kind) {
          case $kindArray:
          case $kindStruct:
            f.typ.copy(dst[f.prop], src[f.prop]);
            continue;
          default:
            dst[f.prop] = src[f.prop];
            continue;
          }
        }
      };
      /* nil value */
      var properties = {};
      fields.forEach(function(f) {
        properties[f.prop] = { get: $throwNilPointerError, set: $throwNilPointerError };
      });
      typ.ptr.nil = Object.create(constructor.prototype, properties);
      typ.ptr.nil.$val = typ.ptr.nil;
      /* methods for embedded fields */
      $addMethodSynthesizer(function() {
        var synthesizeMethod = function(target, m, f) {
          if (target.prototype[m.prop] !== undefined) { return; }
          target.prototype[m.prop] = function() {
            var v = this.$val[f.prop];
            if (f.typ === $jsObjectPtr) {
              v = new $jsObjectPtr(v);
            }
            if (v.$val === undefined) {
              v = new f.typ(v);
            }
            return v[m.prop].apply(v, arguments);
          };
        };
        fields.forEach(function(f) {
          if (f.anonymous) {
            $methodSet(f.typ).forEach(function(m) {
              synthesizeMethod(typ, m, f);
              synthesizeMethod(typ.ptr, m, f);
            });
            $methodSet($ptrType(f.typ)).forEach(function(m) {
              synthesizeMethod(typ.ptr, m, f);
            });
          }
        });
      });
    };
    break;

  default:
    $panic(new $String("invalid kind: " + kind));
  }

  switch (kind) {
  case $kindBool:
  case $kindMap:
    typ.zero = function() { return false; };
    break;

  case $kindInt:
  case $kindInt8:
  case $kindInt16:
  case $kindInt32:
  case $kindUint:
  case $kindUint8 :
  case $kindUint16:
  case $kindUint32:
  case $kindUintptr:
  case $kindUnsafePointer:
  case $kindFloat32:
  case $kindFloat64:
    typ.zero = function() { return 0; };
    break;

  case $kindString:
    typ.zero = function() { return ""; };
    break;

  case $kindInt64:
  case $kindUint64:
  case $kindComplex64:
  case $kindComplex128:
    var zero = new typ(0, 0);
    typ.zero = function() { return zero; };
    break;

  case $kindPtr:
  case $kindSlice:
    typ.zero = function() { return typ.nil; };
    break;

  case $kindChan:
    typ.zero = function() { return $chanNil; };
    break;

  case $kindFunc:
    typ.zero = function() { return $throwNilPointerError; };
    break;

  case $kindInterface:
    typ.zero = function() { return $ifaceNil; };
    break;

  case $kindArray:
    typ.zero = function() {
      var arrayClass = $nativeArray(typ.elem.kind);
      if (arrayClass !== Array) {
        return new arrayClass(typ.len);
      }
      var array = new Array(typ.len);
      for (var i = 0; i < typ.len; i++) {
        array[i] = typ.elem.zero();
      }
      return array;
    };
    break;

  case $kindStruct:
    typ.zero = function() { return new typ.ptr(); };
    break;

  default:
    $panic(new $String("invalid kind: " + kind));
  }

  typ.id = $typeIDCounter;
  $typeIDCounter++;
  typ.size = size;
  typ.kind = kind;
  typ.string = string;
  typ.named = named;
  typ.pkg = pkg;
  typ.exported = exported;
  typ.methods = [];
  typ.methodSetCache = null;
  typ.comparable = true;
  return typ;
};

var $methodSet = function(typ) {
  if (typ.methodSetCache !== null) {
    return typ.methodSetCache;
  }
  var base = {};

  var isPtr = (typ.kind === $kindPtr);
  if (isPtr && typ.elem.kind === $kindInterface) {
    typ.methodSetCache = [];
    return [];
  }

  var current = [{typ: isPtr ? typ.elem : typ, indirect: isPtr}];

  var seen = {};

  while (current.length > 0) {
    var next = [];
    var mset = [];

    current.forEach(function(e) {
      if (seen[e.typ.string]) {
        return;
      }
      seen[e.typ.string] = true;

      if (e.typ.named) {
        mset = mset.concat(e.typ.methods);
        if (e.indirect) {
          mset = mset.concat($ptrType(e.typ).methods);
        }
      }

      switch (e.typ.kind) {
      case $kindStruct:
        e.typ.fields.forEach(function(f) {
          if (f.anonymous) {
            var fTyp = f.typ;
            var fIsPtr = (fTyp.kind === $kindPtr);
            next.push({typ: fIsPtr ? fTyp.elem : fTyp, indirect: e.indirect || fIsPtr});
          }
        });
        break;

      case $kindInterface:
        mset = mset.concat(e.typ.methods);
        break;
      }
    });

    mset.forEach(function(m) {
      if (base[m.name] === undefined) {
        base[m.name] = m;
      }
    });

    current = next;
  }

  typ.methodSetCache = [];
  Object.keys(base).sort().forEach(function(name) {
    typ.methodSetCache.push(base[name]);
  });
  return typ.methodSetCache;
};

var $Bool          = $newType( 1, $kindBool,          "bool",           true, "", false, null);
var $Int           = $newType( 4, $kindInt,           "int",            true, "", false, null);
var $Int8          = $newType( 1, $kindInt8,          "int8",           true, "", false, null);
var $Int16         = $newType( 2, $kindInt16,         "int16",          true, "", false, null);
var $Int32         = $newType( 4, $kindInt32,         "int32",          true, "", false, null);
var $Int64         = $newType( 8, $kindInt64,         "int64",          true, "", false, null);
var $Uint          = $newType( 4, $kindUint,          "uint",           true, "", false, null);
var $Uint8         = $newType( 1, $kindUint8,         "uint8",          true, "", false, null);
var $Uint16        = $newType( 2, $kindUint16,        "uint16",         true, "", false, null);
var $Uint32        = $newType( 4, $kindUint32,        "uint32",         true, "", false, null);
var $Uint64        = $newType( 8, $kindUint64,        "uint64",         true, "", false, null);
var $Uintptr       = $newType( 4, $kindUintptr,       "uintptr",        true, "", false, null);
var $Float32       = $newType( 4, $kindFloat32,       "float32",        true, "", false, null);
var $Float64       = $newType( 8, $kindFloat64,       "float64",        true, "", false, null);
var $Complex64     = $newType( 8, $kindComplex64,     "complex64",      true, "", false, null);
var $Complex128    = $newType(16, $kindComplex128,    "complex128",     true, "", false, null);
var $String        = $newType( 8, $kindString,        "string",         true, "", false, null);
var $UnsafePointer = $newType( 4, $kindUnsafePointer, "unsafe.Pointer", true, "", false, null);

var $nativeArray = function(elemKind) {
  switch (elemKind) {
  case $kindInt:
    return Int32Array;
  case $kindInt8:
    return Int8Array;
  case $kindInt16:
    return Int16Array;
  case $kindInt32:
    return Int32Array;
  case $kindUint:
    return Uint32Array;
  case $kindUint8:
    return Uint8Array;
  case $kindUint16:
    return Uint16Array;
  case $kindUint32:
    return Uint32Array;
  case $kindUintptr:
    return Uint32Array;
  case $kindFloat32:
    return Float32Array;
  case $kindFloat64:
    return Float64Array;
  default:
    return Array;
  }
};
var $toNativeArray = function(elemKind, array) {
  var nativeArray = $nativeArray(elemKind);
  if (nativeArray === Array) {
    return array;
  }
  return new nativeArray(array);
};
var $arrayTypes = {};
var $arrayType = function(elem, len) {
  var typeKey = elem.id + "$" + len;
  var typ = $arrayTypes[typeKey];
  if (typ === undefined) {
    typ = $newType(12, $kindArray, "[" + len + "]" + elem.string, false, "", false, null);
    $arrayTypes[typeKey] = typ;
    typ.init(elem, len);
  }
  return typ;
};

var $chanType = function(elem, sendOnly, recvOnly) {
  var string = (recvOnly ? "<-" : "") + "chan" + (sendOnly ? "<- " : " ") + elem.string;
  var field = sendOnly ? "SendChan" : (recvOnly ? "RecvChan" : "Chan");
  var typ = elem[field];
  if (typ === undefined) {
    typ = $newType(4, $kindChan, string, false, "", false, null);
    elem[field] = typ;
    typ.init(elem, sendOnly, recvOnly);
  }
  return typ;
};
var $Chan = function(elem, capacity) {
  if (capacity < 0 || capacity > 2147483647) {
    $throwRuntimeError("makechan: size out of range");
  }
  this.$elem = elem;
  this.$capacity = capacity;
  this.$buffer = [];
  this.$sendQueue = [];
  this.$recvQueue = [];
  this.$closed = false;
};
var $chanNil = new $Chan(null, 0);
$chanNil.$sendQueue = $chanNil.$recvQueue = { length: 0, push: function() {}, shift: function() { return undefined; }, indexOf: function() { return -1; } };

var $funcTypes = {};
var $funcType = function(params, results, variadic) {
  var typeKey = $mapArray(params, function(p) { return p.id; }).join(",") + "$" + $mapArray(results, function(r) { return r.id; }).join(",") + "$" + variadic;
  var typ = $funcTypes[typeKey];
  if (typ === undefined) {
    var paramTypes = $mapArray(params, function(p) { return p.string; });
    if (variadic) {
      paramTypes[paramTypes.length - 1] = "..." + paramTypes[paramTypes.length - 1].substr(2);
    }
    var string = "func(" + paramTypes.join(", ") + ")";
    if (results.length === 1) {
      string += " " + results[0].string;
    } else if (results.length > 1) {
      string += " (" + $mapArray(results, function(r) { return r.string; }).join(", ") + ")";
    }
    typ = $newType(4, $kindFunc, string, false, "", false, null);
    $funcTypes[typeKey] = typ;
    typ.init(params, results, variadic);
  }
  return typ;
};

var $interfaceTypes = {};
var $interfaceType = function(methods) {
  var typeKey = $mapArray(methods, function(m) { return m.pkg + "," + m.name + "," + m.typ.id; }).join("$");
  var typ = $interfaceTypes[typeKey];
  if (typ === undefined) {
    var string = "interface {}";
    if (methods.length !== 0) {
      string = "interface { " + $mapArray(methods, function(m) {
        return (m.pkg !== "" ? m.pkg + "." : "") + m.name + m.typ.string.substr(4);
      }).join("; ") + " }";
    }
    typ = $newType(8, $kindInterface, string, false, "", false, null);
    $interfaceTypes[typeKey] = typ;
    typ.init(methods);
  }
  return typ;
};
var $emptyInterface = $interfaceType([]);
var $ifaceNil = {};
var $error = $newType(8, $kindInterface, "error", true, "", false, null);
$error.init([{prop: "Error", name: "Error", pkg: "", typ: $funcType([], [$String], false)}]);

var $mapTypes = {};
var $mapType = function(key, elem) {
  var typeKey = key.id + "$" + elem.id;
  var typ = $mapTypes[typeKey];
  if (typ === undefined) {
    typ = $newType(4, $kindMap, "map[" + key.string + "]" + elem.string, false, "", false, null);
    $mapTypes[typeKey] = typ;
    typ.init(key, elem);
  }
  return typ;
};
var $makeMap = function(keyForFunc, entries) {
  var m = {};
  for (var i = 0; i < entries.length; i++) {
    var e = entries[i];
    m[keyForFunc(e.k)] = e;
  }
  return m;
};

var $ptrType = function(elem) {
  var typ = elem.ptr;
  if (typ === undefined) {
    typ = $newType(4, $kindPtr, "*" + elem.string, false, "", elem.exported, null);
    elem.ptr = typ;
    typ.init(elem);
  }
  return typ;
};

var $newDataPointer = function(data, constructor) {
  if (constructor.elem.kind === $kindStruct) {
    return data;
  }
  return new constructor(function() { return data; }, function(v) { data = v; });
};

var $indexPtr = function(array, index, constructor) {
  array.$ptr = array.$ptr || {};
  return array.$ptr[index] || (array.$ptr[index] = new constructor(function() { return array[index]; }, function(v) { array[index] = v; }));
};

var $sliceType = function(elem) {
  var typ = elem.slice;
  if (typ === undefined) {
    typ = $newType(12, $kindSlice, "[]" + elem.string, false, "", false, null);
    elem.slice = typ;
    typ.init(elem);
  }
  return typ;
};
var $makeSlice = function(typ, length, capacity) {
  capacity = capacity || length;
  if (length < 0 || length > 2147483647) {
    $throwRuntimeError("makeslice: len out of range");
  }
  if (capacity < 0 || capacity < length || capacity > 2147483647) {
    $throwRuntimeError("makeslice: cap out of range");
  }
  var array = new typ.nativeArray(capacity);
  if (typ.nativeArray === Array) {
    for (var i = 0; i < capacity; i++) {
      array[i] = typ.elem.zero();
    }
  }
  var slice = new typ(array);
  slice.$length = length;
  return slice;
};

var $structTypes = {};
var $structType = function(pkgPath, fields) {
  var typeKey = $mapArray(fields, function(f) { return f.name + "," + f.typ.id + "," + f.tag; }).join("$");
  var typ = $structTypes[typeKey];
  if (typ === undefined) {
    var string = "struct { " + $mapArray(fields, function(f) {
      return f.name + " " + f.typ.string + (f.tag !== "" ? (" \"" + f.tag.replace(/\\/g, "\\\\").replace(/"/g, "\\\"") + "\"") : "");
    }).join("; ") + " }";
    if (fields.length === 0) {
      string = "struct {}";
    }
    typ = $newType(0, $kindStruct, string, false, "", false, function() {
      this.$val = this;
      for (var i = 0; i < fields.length; i++) {
        var f = fields[i];
        var arg = arguments[i];
        this[f.prop] = arg !== undefined ? arg : f.typ.zero();
      }
    });
    $structTypes[typeKey] = typ;
    typ.init(pkgPath, fields);
  }
  return typ;
};

var $assertType = function(value, type, returnTuple) {
  var isInterface = (type.kind === $kindInterface), ok, missingMethod = "";
  if (value === $ifaceNil) {
    ok = false;
  } else if (!isInterface) {
    ok = value.constructor === type;
  } else {
    var valueTypeString = value.constructor.string;
    ok = type.implementedBy[valueTypeString];
    if (ok === undefined) {
      ok = true;
      var valueMethodSet = $methodSet(value.constructor);
      var interfaceMethods = type.methods;
      for (var i = 0; i < interfaceMethods.length; i++) {
        var tm = interfaceMethods[i];
        var found = false;
        for (var j = 0; j < valueMethodSet.length; j++) {
          var vm = valueMethodSet[j];
          if (vm.name === tm.name && vm.pkg === tm.pkg && vm.typ === tm.typ) {
            found = true;
            break;
          }
        }
        if (!found) {
          ok = false;
          type.missingMethodFor[valueTypeString] = tm.name;
          break;
        }
      }
      type.implementedBy[valueTypeString] = ok;
    }
    if (!ok) {
      missingMethod = type.missingMethodFor[valueTypeString];
    }
  }

  if (!ok) {
    if (returnTuple) {
      return [type.zero(), false];
    }
    $panic(new $packages["runtime"].TypeAssertionError.ptr("", (value === $ifaceNil ? "" : value.constructor.string), type.string, missingMethod));
  }

  if (!isInterface) {
    value = value.$val;
  }
  if (type === $jsObjectPtr) {
    value = value.object;
  }
  return returnTuple ? [value, true] : value;
};

var $stackDepthOffset = 0;
var $getStackDepth = function() {
  var err = new Error();
  if (err.stack === undefined) {
    return undefined;
  }
  return $stackDepthOffset + err.stack.split("\n").length;
};

var $panicStackDepth = null, $panicValue;
var $callDeferred = function(deferred, jsErr, fromPanic) {
  if (!fromPanic && deferred !== null && deferred.index >= $curGoroutine.deferStack.length) {
    throw jsErr;
  }
  if (jsErr !== null) {
    var newErr = null;
    try {
      $curGoroutine.deferStack.push(deferred);
      $panic(new $jsErrorPtr(jsErr));
    } catch (err) {
      newErr = err;
    }
    $curGoroutine.deferStack.pop();
    $callDeferred(deferred, newErr);
    return;
  }
  if ($curGoroutine.asleep) {
    return;
  }

  $stackDepthOffset--;
  var outerPanicStackDepth = $panicStackDepth;
  var outerPanicValue = $panicValue;

  var localPanicValue = $curGoroutine.panicStack.pop();
  if (localPanicValue !== undefined) {
    $panicStackDepth = $getStackDepth();
    $panicValue = localPanicValue;
  }

  try {
    while (true) {
      if (deferred === null) {
        deferred = $curGoroutine.deferStack[$curGoroutine.deferStack.length - 1];
        if (deferred === undefined) {
          /* The panic reached the top of the stack. Clear it and throw it as a JavaScript error. */
          $panicStackDepth = null;
          if (localPanicValue.Object instanceof Error) {
            throw localPanicValue.Object;
          }
          var msg;
          if (localPanicValue.constructor === $String) {
            msg = localPanicValue.$val;
          } else if (localPanicValue.Error !== undefined) {
            msg = localPanicValue.Error();
          } else if (localPanicValue.String !== undefined) {
            msg = localPanicValue.String();
          } else {
            msg = localPanicValue;
          }
          throw new Error(msg);
        }
      }
      var call = deferred.pop();
      if (call === undefined) {
        $curGoroutine.deferStack.pop();
        if (localPanicValue !== undefined) {
          deferred = null;
          continue;
        }
        return;
      }
      var r = call[0].apply(call[2], call[1]);
      if (r && r.$blk !== undefined) {
        deferred.push([r.$blk, [], r]);
        if (fromPanic) {
          throw null;
        }
        return;
      }

      if (localPanicValue !== undefined && $panicStackDepth === null) {
        throw null; /* error was recovered */
      }
    }
  } finally {
    if (localPanicValue !== undefined) {
      if ($panicStackDepth !== null) {
        $curGoroutine.panicStack.push(localPanicValue);
      }
      $panicStackDepth = outerPanicStackDepth;
      $panicValue = outerPanicValue;
    }
    $stackDepthOffset++;
  }
};

var $panic = function(value) {
  $curGoroutine.panicStack.push(value);
  $callDeferred(null, null, true);
};
var $recover = function() {
  if ($panicStackDepth === null || ($panicStackDepth !== undefined && $panicStackDepth !== $getStackDepth() - 2)) {
    return $ifaceNil;
  }
  $panicStackDepth = null;
  return $panicValue;
};
var $throw = function(err) { throw err; };

var $noGoroutine = { asleep: false, exit: false, deferStack: [], panicStack: [] };
var $curGoroutine = $noGoroutine, $totalGoroutines = 0, $awakeGoroutines = 0, $checkForDeadlock = true;
var $mainFinished = false;
var $go = function(fun, args, direct) {
  $totalGoroutines++;
  $awakeGoroutines++;
  var $goroutine = function() {
    try {
      $curGoroutine = $goroutine;
      var r = fun.apply(undefined, args);
      if (r && r.$blk !== undefined) {
        fun = function() { return r.$blk(); };
        args = [];
        return;
      }
      $goroutine.exit = true;
    } catch (err) {
      if (!$goroutine.exit) {
        throw err;
      }
    } finally {
      $curGoroutine = $noGoroutine;
      if ($goroutine.exit) { /* also set by runtime.Goexit() */
        $totalGoroutines--;
        $goroutine.asleep = true;
      }
      if ($goroutine.asleep) {
        $awakeGoroutines--;
        if (!$mainFinished && $awakeGoroutines === 0 && $checkForDeadlock) {
          console.error("fatal error: all goroutines are asleep - deadlock!");
          if ($global.process !== undefined) {
            $global.process.exit(2);
          }
        }
      }
    }
  };
  $goroutine.asleep = false;
  $goroutine.exit = false;
  $goroutine.deferStack = [];
  $goroutine.panicStack = [];
  $schedule($goroutine);
};

var $scheduled = [];
var $runScheduled = function() {
  try {
    var r;
    while ((r = $scheduled.shift()) !== undefined) {
      r();
    }
  } finally {
    if ($scheduled.length > 0) {
      setTimeout($runScheduled, 0);
    }
  }
};

var $schedule = function(goroutine) {
  if (goroutine.asleep) {
    goroutine.asleep = false;
    $awakeGoroutines++;
  }
  $scheduled.push(goroutine);
  if ($curGoroutine === $noGoroutine) {
    $runScheduled();
  }
};

var $setTimeout = function(f, t) {
  $awakeGoroutines++;
  return setTimeout(function() {
    $awakeGoroutines--;
    f();
  }, t);
};

var $block = function() {
  if ($curGoroutine === $noGoroutine) {
    $throwRuntimeError("cannot block in JavaScript callback, fix by wrapping code in goroutine");
  }
  $curGoroutine.asleep = true;
};

var $send = function(chan, value) {
  if (chan.$closed) {
    $throwRuntimeError("send on closed channel");
  }
  var queuedRecv = chan.$recvQueue.shift();
  if (queuedRecv !== undefined) {
    queuedRecv([value, true]);
    return;
  }
  if (chan.$buffer.length < chan.$capacity) {
    chan.$buffer.push(value);
    return;
  }

  var thisGoroutine = $curGoroutine;
  var closedDuringSend;
  chan.$sendQueue.push(function(closed) {
    closedDuringSend = closed;
    $schedule(thisGoroutine);
    return value;
  });
  $block();
  return {
    $blk: function() {
      if (closedDuringSend) {
        $throwRuntimeError("send on closed channel");
      }
    }
  };
};
var $recv = function(chan) {
  var queuedSend = chan.$sendQueue.shift();
  if (queuedSend !== undefined) {
    chan.$buffer.push(queuedSend(false));
  }
  var bufferedValue = chan.$buffer.shift();
  if (bufferedValue !== undefined) {
    return [bufferedValue, true];
  }
  if (chan.$closed) {
    return [chan.$elem.zero(), false];
  }

  var thisGoroutine = $curGoroutine;
  var f = { $blk: function() { return this.value; } };
  var queueEntry = function(v) {
    f.value = v;
    $schedule(thisGoroutine);
  };
  chan.$recvQueue.push(queueEntry);
  $block();
  return f;
};
var $close = function(chan) {
  if (chan.$closed) {
    $throwRuntimeError("close of closed channel");
  }
  chan.$closed = true;
  while (true) {
    var queuedSend = chan.$sendQueue.shift();
    if (queuedSend === undefined) {
      break;
    }
    queuedSend(true); /* will panic */
  }
  while (true) {
    var queuedRecv = chan.$recvQueue.shift();
    if (queuedRecv === undefined) {
      break;
    }
    queuedRecv([chan.$elem.zero(), false]);
  }
};
var $select = function(comms) {
  var ready = [];
  var selection = -1;
  for (var i = 0; i < comms.length; i++) {
    var comm = comms[i];
    var chan = comm[0];
    switch (comm.length) {
    case 0: /* default */
      selection = i;
      break;
    case 1: /* recv */
      if (chan.$sendQueue.length !== 0 || chan.$buffer.length !== 0 || chan.$closed) {
        ready.push(i);
      }
      break;
    case 2: /* send */
      if (chan.$closed) {
        $throwRuntimeError("send on closed channel");
      }
      if (chan.$recvQueue.length !== 0 || chan.$buffer.length < chan.$capacity) {
        ready.push(i);
      }
      break;
    }
  }

  if (ready.length !== 0) {
    selection = ready[Math.floor(Math.random() * ready.length)];
  }
  if (selection !== -1) {
    var comm = comms[selection];
    switch (comm.length) {
    case 0: /* default */
      return [selection];
    case 1: /* recv */
      return [selection, $recv(comm[0])];
    case 2: /* send */
      $send(comm[0], comm[1]);
      return [selection];
    }
  }

  var entries = [];
  var thisGoroutine = $curGoroutine;
  var f = { $blk: function() { return this.selection; } };
  var removeFromQueues = function() {
    for (var i = 0; i < entries.length; i++) {
      var entry = entries[i];
      var queue = entry[0];
      var index = queue.indexOf(entry[1]);
      if (index !== -1) {
        queue.splice(index, 1);
      }
    }
  };
  for (var i = 0; i < comms.length; i++) {
    (function(i) {
      var comm = comms[i];
      switch (comm.length) {
      case 1: /* recv */
        var queueEntry = function(value) {
          f.selection = [i, value];
          removeFromQueues();
          $schedule(thisGoroutine);
        };
        entries.push([comm[0].$recvQueue, queueEntry]);
        comm[0].$recvQueue.push(queueEntry);
        break;
      case 2: /* send */
        var queueEntry = function() {
          if (comm[0].$closed) {
            $throwRuntimeError("send on closed channel");
          }
          f.selection = [i];
          removeFromQueues();
          $schedule(thisGoroutine);
          return comm[1];
        };
        entries.push([comm[0].$sendQueue, queueEntry]);
        comm[0].$sendQueue.push(queueEntry);
        break;
      }
    })(i);
  }
  $block();
  return f;
};

var $jsObjectPtr, $jsErrorPtr;

var $needsExternalization = function(t) {
  switch (t.kind) {
    case $kindBool:
    case $kindInt:
    case $kindInt8:
    case $kindInt16:
    case $kindInt32:
    case $kindUint:
    case $kindUint8:
    case $kindUint16:
    case $kindUint32:
    case $kindUintptr:
    case $kindFloat32:
    case $kindFloat64:
      return false;
    default:
      return t !== $jsObjectPtr;
  }
};

var $externalize = function(v, t) {
  if (t === $jsObjectPtr) {
    return v;
  }
  switch (t.kind) {
  case $kindBool:
  case $kindInt:
  case $kindInt8:
  case $kindInt16:
  case $kindInt32:
  case $kindUint:
  case $kindUint8:
  case $kindUint16:
  case $kindUint32:
  case $kindUintptr:
  case $kindFloat32:
  case $kindFloat64:
    return v;
  case $kindInt64:
  case $kindUint64:
    return $flatten64(v);
  case $kindArray:
    if ($needsExternalization(t.elem)) {
      return $mapArray(v, function(e) { return $externalize(e, t.elem); });
    }
    return v;
  case $kindFunc:
    return $externalizeFunction(v, t, false);
  case $kindInterface:
    if (v === $ifaceNil) {
      return null;
    }
    if (v.constructor === $jsObjectPtr) {
      return v.$val.object;
    }
    return $externalize(v.$val, v.constructor);
  case $kindMap:
    var m = {};
    var keys = $keys(v);
    for (var i = 0; i < keys.length; i++) {
      var entry = v[keys[i]];
      m[$externalize(entry.k, t.key)] = $externalize(entry.v, t.elem);
    }
    return m;
  case $kindPtr:
    if (v === t.nil) {
      return null;
    }
    return $externalize(v.$get(), t.elem);
  case $kindSlice:
    if ($needsExternalization(t.elem)) {
      return $mapArray($sliceToArray(v), function(e) { return $externalize(e, t.elem); });
    }
    return $sliceToArray(v);
  case $kindString:
    if ($isASCII(v)) {
      return v;
    }
    var s = "", r;
    for (var i = 0; i < v.length; i += r[1]) {
      r = $decodeRune(v, i);
      var c = r[0];
      if (c > 0xFFFF) {
        var h = Math.floor((c - 0x10000) / 0x400) + 0xD800;
        var l = (c - 0x10000) % 0x400 + 0xDC00;
        s += String.fromCharCode(h, l);
        continue;
      }
      s += String.fromCharCode(c);
    }
    return s;
  case $kindStruct:
    var timePkg = $packages["time"];
    if (timePkg !== undefined && v.constructor === timePkg.Time.ptr) {
      var milli = $div64(v.UnixNano(), new $Int64(0, 1000000));
      return new Date($flatten64(milli));
    }

    var noJsObject = {};
    var searchJsObject = function(v, t) {
      if (t === $jsObjectPtr) {
        return v;
      }
      switch (t.kind) {
      case $kindPtr:
        if (v === t.nil) {
          return noJsObject;
        }
        return searchJsObject(v.$get(), t.elem);
      case $kindStruct:
        var f = t.fields[0];
        return searchJsObject(v[f.prop], f.typ);
      case $kindInterface:
        return searchJsObject(v.$val, v.constructor);
      default:
        return noJsObject;
      }
    };
    var o = searchJsObject(v, t);
    if (o !== noJsObject) {
      return o;
    }

    o = {};
    for (var i = 0; i < t.fields.length; i++) {
      var f = t.fields[i];
      if (!f.exported) {
        continue;
      }
      o[f.name] = $externalize(v[f.prop], f.typ);
    }
    return o;
  }
  $throwRuntimeError("cannot externalize " + t.string);
};

var $externalizeFunction = function(v, t, passThis) {
  if (v === $throwNilPointerError) {
    return null;
  }
  if (v.$externalizeWrapper === undefined) {
    $checkForDeadlock = false;
    v.$externalizeWrapper = function() {
      var args = [];
      for (var i = 0; i < t.params.length; i++) {
        if (t.variadic && i === t.params.length - 1) {
          var vt = t.params[i].elem, varargs = [];
          for (var j = i; j < arguments.length; j++) {
            varargs.push($internalize(arguments[j], vt));
          }
          args.push(new (t.params[i])(varargs));
          break;
        }
        args.push($internalize(arguments[i], t.params[i]));
      }
      var canBlock = $curGoroutine.canBlock;
      $curGoroutine.canBlock = false;
      try {
        var result = v.apply(passThis ? this : undefined, args);
      } finally {
        $curGoroutine.canBlock = canBlock;
      }
      switch (t.results.length) {
      case 0:
        return;
      case 1:
        return $externalize(result, t.results[0]);
      default:
        for (var i = 0; i < t.results.length; i++) {
          result[i] = $externalize(result[i], t.results[i]);
        }
        return result;
      }
    };
  }
  return v.$externalizeWrapper;
};

var $internalize = function(v, t, recv) {
  if (t === $jsObjectPtr) {
    return v;
  }
  if (t === $jsObjectPtr.elem) {
    $throwRuntimeError("cannot internalize js.Object, use *js.Object instead");
  }
  if (v && v.__internal_object__ !== undefined) {
    return $assertType(v.__internal_object__, t, false);
  }
  var timePkg = $packages["time"];
  if (timePkg !== undefined && t === timePkg.Time) {
    if (!(v !== null && v !== undefined && v.constructor === Date)) {
      $throwRuntimeError("cannot internalize time.Time from " + typeof v + ", must be Date");
    }
    return timePkg.Unix(new $Int64(0, 0), new $Int64(0, v.getTime() * 1000000));
  }
  switch (t.kind) {
  case $kindBool:
    return !!v;
  case $kindInt:
    return parseInt(v);
  case $kindInt8:
    return parseInt(v) << 24 >> 24;
  case $kindInt16:
    return parseInt(v) << 16 >> 16;
  case $kindInt32:
    return parseInt(v) >> 0;
  case $kindUint:
    return parseInt(v);
  case $kindUint8:
    return parseInt(v) << 24 >>> 24;
  case $kindUint16:
    return parseInt(v) << 16 >>> 16;
  case $kindUint32:
  case $kindUintptr:
    return parseInt(v) >>> 0;
  case $kindInt64:
  case $kindUint64:
    return new t(0, v);
  case $kindFloat32:
  case $kindFloat64:
    return parseFloat(v);
  case $kindArray:
    if (v.length !== t.len) {
      $throwRuntimeError("got array with wrong size from JavaScript native");
    }
    return $mapArray(v, function(e) { return $internalize(e, t.elem); });
  case $kindFunc:
    return function() {
      var args = [];
      for (var i = 0; i < t.params.length; i++) {
        if (t.variadic && i === t.params.length - 1) {
          var vt = t.params[i].elem, varargs = arguments[i];
          for (var j = 0; j < varargs.$length; j++) {
            args.push($externalize(varargs.$array[varargs.$offset + j], vt));
          }
          break;
        }
        args.push($externalize(arguments[i], t.params[i]));
      }
      var result = v.apply(recv, args);
      switch (t.results.length) {
      case 0:
        return;
      case 1:
        return $internalize(result, t.results[0]);
      default:
        for (var i = 0; i < t.results.length; i++) {
          result[i] = $internalize(result[i], t.results[i]);
        }
        return result;
      }
    };
  case $kindInterface:
    if (t.methods.length !== 0) {
      $throwRuntimeError("cannot internalize " + t.string);
    }
    if (v === null) {
      return $ifaceNil;
    }
    if (v === undefined) {
      return new $jsObjectPtr(undefined);
    }
    switch (v.constructor) {
    case Int8Array:
      return new ($sliceType($Int8))(v);
    case Int16Array:
      return new ($sliceType($Int16))(v);
    case Int32Array:
      return new ($sliceType($Int))(v);
    case Uint8Array:
      return new ($sliceType($Uint8))(v);
    case Uint16Array:
      return new ($sliceType($Uint16))(v);
    case Uint32Array:
      return new ($sliceType($Uint))(v);
    case Float32Array:
      return new ($sliceType($Float32))(v);
    case Float64Array:
      return new ($sliceType($Float64))(v);
    case Array:
      return $internalize(v, $sliceType($emptyInterface));
    case Boolean:
      return new $Bool(!!v);
    case Date:
      if (timePkg === undefined) {
        /* time package is not present, internalize as &js.Object{Date} so it can be externalized into original Date. */
        return new $jsObjectPtr(v);
      }
      return new timePkg.Time($internalize(v, timePkg.Time));
    case Function:
      var funcType = $funcType([$sliceType($emptyInterface)], [$jsObjectPtr], true);
      return new funcType($internalize(v, funcType));
    case Number:
      return new $Float64(parseFloat(v));
    case String:
      return new $String($internalize(v, $String));
    default:
      if ($global.Node && v instanceof $global.Node) {
        return new $jsObjectPtr(v);
      }
      var mapType = $mapType($String, $emptyInterface);
      return new mapType($internalize(v, mapType));
    }
  case $kindMap:
    var m = {};
    var keys = $keys(v);
    for (var i = 0; i < keys.length; i++) {
      var k = $internalize(keys[i], t.key);
      m[t.key.keyFor(k)] = { k: k, v: $internalize(v[keys[i]], t.elem) };
    }
    return m;
  case $kindPtr:
    if (t.elem.kind === $kindStruct) {
      return $internalize(v, t.elem);
    }
  case $kindSlice:
    return new t($mapArray(v, function(e) { return $internalize(e, t.elem); }));
  case $kindString:
    v = String(v);
    if ($isASCII(v)) {
      return v;
    }
    var s = "";
    var i = 0;
    while (i < v.length) {
      var h = v.charCodeAt(i);
      if (0xD800 <= h && h <= 0xDBFF) {
        var l = v.charCodeAt(i + 1);
        var c = (h - 0xD800) * 0x400 + l - 0xDC00 + 0x10000;
        s += $encodeRune(c);
        i += 2;
        continue;
      }
      s += $encodeRune(h);
      i++;
    }
    return s;
  case $kindStruct:
    var noJsObject = {};
    var searchJsObject = function(t) {
      if (t === $jsObjectPtr) {
        return v;
      }
      if (t === $jsObjectPtr.elem) {
        $throwRuntimeError("cannot internalize js.Object, use *js.Object instead");
      }
      switch (t.kind) {
      case $kindPtr:
        return searchJsObject(t.elem);
      case $kindStruct:
        var f = t.fields[0];
        var o = searchJsObject(f.typ);
        if (o !== noJsObject) {
          var n = new t.ptr();
          n[f.prop] = o;
          return n;
        }
        return noJsObject;
      default:
        return noJsObject;
      }
    };
    var o = searchJsObject(t);
    if (o !== noJsObject) {
      return o;
    }
  }
  $throwRuntimeError("cannot internalize " + t.string);
};

/* $isASCII reports whether string s contains only ASCII characters. */
var $isASCII = function(s) {
  for (var i = 0; i < s.length; i++) {
    if (s.charCodeAt(i) >= 128) {
      return false;
    }
  }
  return true;
};

$packages["github.com/gopherjs/gopherjs/js"] = (function() {
	var $pkg = {}, $init, Object, Error, M, sliceType, sliceType$1, ptrType, sliceType$2, funcType, ptrType$1, MakeFunc, Keys, MakeWrapper, init;
	Object = $pkg.Object = $newType(0, $kindStruct, "js.Object", true, "github.com/gopherjs/gopherjs/js", true, function(object_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.object = null;
			return;
		}
		this.object = object_;
	});
	Error = $pkg.Error = $newType(0, $kindStruct, "js.Error", true, "github.com/gopherjs/gopherjs/js", true, function(Object_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.Object = null;
			return;
		}
		this.Object = Object_;
	});
	M = $pkg.M = $newType(4, $kindMap, "js.M", true, "github.com/gopherjs/gopherjs/js", true, null);
	sliceType = $sliceType($emptyInterface);
	sliceType$1 = $sliceType($String);
	ptrType = $ptrType(Object);
	sliceType$2 = $sliceType(ptrType);
	funcType = $funcType([sliceType$2], [ptrType], true);
	ptrType$1 = $ptrType(Error);
	Object.ptr.prototype.Get = function(key) {
		var $ptr, key, o;
		o = this;
		return o.object[$externalize(key, $String)];
	};
	Object.prototype.Get = function(key) { return this.$val.Get(key); };
	Object.ptr.prototype.Set = function(key, value) {
		var $ptr, key, o, value;
		o = this;
		o.object[$externalize(key, $String)] = $externalize(value, $emptyInterface);
	};
	Object.prototype.Set = function(key, value) { return this.$val.Set(key, value); };
	Object.ptr.prototype.Delete = function(key) {
		var $ptr, key, o;
		o = this;
		delete o.object[$externalize(key, $String)];
	};
	Object.prototype.Delete = function(key) { return this.$val.Delete(key); };
	Object.ptr.prototype.Length = function() {
		var $ptr, o;
		o = this;
		return $parseInt(o.object.length);
	};
	Object.prototype.Length = function() { return this.$val.Length(); };
	Object.ptr.prototype.Index = function(i) {
		var $ptr, i, o;
		o = this;
		return o.object[i];
	};
	Object.prototype.Index = function(i) { return this.$val.Index(i); };
	Object.ptr.prototype.SetIndex = function(i, value) {
		var $ptr, i, o, value;
		o = this;
		o.object[i] = $externalize(value, $emptyInterface);
	};
	Object.prototype.SetIndex = function(i, value) { return this.$val.SetIndex(i, value); };
	Object.ptr.prototype.Call = function(name, args) {
		var $ptr, args, name, o, obj;
		o = this;
		return (obj = o.object, obj[$externalize(name, $String)].apply(obj, $externalize(args, sliceType)));
	};
	Object.prototype.Call = function(name, args) { return this.$val.Call(name, args); };
	Object.ptr.prototype.Invoke = function(args) {
		var $ptr, args, o;
		o = this;
		return o.object.apply(undefined, $externalize(args, sliceType));
	};
	Object.prototype.Invoke = function(args) { return this.$val.Invoke(args); };
	Object.ptr.prototype.New = function(args) {
		var $ptr, args, o;
		o = this;
		return new ($global.Function.prototype.bind.apply(o.object, [undefined].concat($externalize(args, sliceType))));
	};
	Object.prototype.New = function(args) { return this.$val.New(args); };
	Object.ptr.prototype.Bool = function() {
		var $ptr, o;
		o = this;
		return !!(o.object);
	};
	Object.prototype.Bool = function() { return this.$val.Bool(); };
	Object.ptr.prototype.String = function() {
		var $ptr, o;
		o = this;
		return $internalize(o.object, $String);
	};
	Object.prototype.String = function() { return this.$val.String(); };
	Object.ptr.prototype.Int = function() {
		var $ptr, o;
		o = this;
		return $parseInt(o.object) >> 0;
	};
	Object.prototype.Int = function() { return this.$val.Int(); };
	Object.ptr.prototype.Int64 = function() {
		var $ptr, o;
		o = this;
		return $internalize(o.object, $Int64);
	};
	Object.prototype.Int64 = function() { return this.$val.Int64(); };
	Object.ptr.prototype.Uint64 = function() {
		var $ptr, o;
		o = this;
		return $internalize(o.object, $Uint64);
	};
	Object.prototype.Uint64 = function() { return this.$val.Uint64(); };
	Object.ptr.prototype.Float = function() {
		var $ptr, o;
		o = this;
		return $parseFloat(o.object);
	};
	Object.prototype.Float = function() { return this.$val.Float(); };
	Object.ptr.prototype.Interface = function() {
		var $ptr, o;
		o = this;
		return $internalize(o.object, $emptyInterface);
	};
	Object.prototype.Interface = function() { return this.$val.Interface(); };
	Object.ptr.prototype.Unsafe = function() {
		var $ptr, o;
		o = this;
		return o.object;
	};
	Object.prototype.Unsafe = function() { return this.$val.Unsafe(); };
	Error.ptr.prototype.Error = function() {
		var $ptr, err;
		err = this;
		return "JavaScript error: " + $internalize(err.Object.message, $String);
	};
	Error.prototype.Error = function() { return this.$val.Error(); };
	Error.ptr.prototype.Stack = function() {
		var $ptr, err;
		err = this;
		return $internalize(err.Object.stack, $String);
	};
	Error.prototype.Stack = function() { return this.$val.Stack(); };
	MakeFunc = function(fn) {
		var $ptr, fn;
		return $makeFunc(fn);
	};
	$pkg.MakeFunc = MakeFunc;
	Keys = function(o) {
		var $ptr, a, i, o, s;
		if (o === null || o === undefined) {
			return sliceType$1.nil;
		}
		a = $global.Object.keys(o);
		s = $makeSlice(sliceType$1, $parseInt(a.length));
		i = 0;
		while (true) {
			if (!(i < $parseInt(a.length))) { break; }
			((i < 0 || i >= s.$length) ? ($throwRuntimeError("index out of range"), undefined) : s.$array[s.$offset + i] = $internalize(a[i], $String));
			i = i + (1) >> 0;
		}
		return s;
	};
	$pkg.Keys = Keys;
	MakeWrapper = function(i) {
		var $ptr, i, i$1, m, methods, o, v;
		v = i;
		o = new ($global.Object)();
		o.__internal_object__ = v;
		methods = v.constructor.methods;
		i$1 = 0;
		while (true) {
			if (!(i$1 < $parseInt(methods.length))) { break; }
			m = [m];
			m[0] = methods[i$1];
			if (!($internalize(m[0].pkg, $String) === "")) {
				i$1 = i$1 + (1) >> 0;
				continue;
			}
			o[$externalize($internalize(m[0].name, $String), $String)] = $externalize((function(m) { return function(args) {
				var $ptr, args;
				return $externalizeFunction(v[$externalize($internalize(m[0].prop, $String), $String)], m[0].typ, $externalize(true, $Bool)).apply(v, $externalize(args, sliceType$2));
			}; })(m), funcType);
			i$1 = i$1 + (1) >> 0;
		}
		return o;
	};
	$pkg.MakeWrapper = MakeWrapper;
	init = function() {
		var $ptr, e;
		e = new Error.ptr(null);
		$unused(e);
	};
	ptrType.methods = [{prop: "Get", name: "Get", pkg: "", typ: $funcType([$String], [ptrType], false)}, {prop: "Set", name: "Set", pkg: "", typ: $funcType([$String, $emptyInterface], [], false)}, {prop: "Delete", name: "Delete", pkg: "", typ: $funcType([$String], [], false)}, {prop: "Length", name: "Length", pkg: "", typ: $funcType([], [$Int], false)}, {prop: "Index", name: "Index", pkg: "", typ: $funcType([$Int], [ptrType], false)}, {prop: "SetIndex", name: "SetIndex", pkg: "", typ: $funcType([$Int, $emptyInterface], [], false)}, {prop: "Call", name: "Call", pkg: "", typ: $funcType([$String, sliceType], [ptrType], true)}, {prop: "Invoke", name: "Invoke", pkg: "", typ: $funcType([sliceType], [ptrType], true)}, {prop: "New", name: "New", pkg: "", typ: $funcType([sliceType], [ptrType], true)}, {prop: "Bool", name: "Bool", pkg: "", typ: $funcType([], [$Bool], false)}, {prop: "String", name: "String", pkg: "", typ: $funcType([], [$String], false)}, {prop: "Int", name: "Int", pkg: "", typ: $funcType([], [$Int], false)}, {prop: "Int64", name: "Int64", pkg: "", typ: $funcType([], [$Int64], false)}, {prop: "Uint64", name: "Uint64", pkg: "", typ: $funcType([], [$Uint64], false)}, {prop: "Float", name: "Float", pkg: "", typ: $funcType([], [$Float64], false)}, {prop: "Interface", name: "Interface", pkg: "", typ: $funcType([], [$emptyInterface], false)}, {prop: "Unsafe", name: "Unsafe", pkg: "", typ: $funcType([], [$Uintptr], false)}];
	ptrType$1.methods = [{prop: "Error", name: "Error", pkg: "", typ: $funcType([], [$String], false)}, {prop: "Stack", name: "Stack", pkg: "", typ: $funcType([], [$String], false)}];
	Object.init("github.com/gopherjs/gopherjs/js", [{prop: "object", name: "object", anonymous: false, exported: false, typ: ptrType, tag: ""}]);
	Error.init("", [{prop: "Object", name: "Object", anonymous: true, exported: true, typ: ptrType, tag: ""}]);
	M.init($String, $emptyInterface);
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		init();
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["runtime/internal/sys"] = (function() {
	var $pkg = {}, $init;
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["runtime"] = (function() {
	var $pkg = {}, $init, js, sys, TypeAssertionError, errorString, ptrType$4, init, throw$1;
	js = $packages["github.com/gopherjs/gopherjs/js"];
	sys = $packages["runtime/internal/sys"];
	TypeAssertionError = $pkg.TypeAssertionError = $newType(0, $kindStruct, "runtime.TypeAssertionError", true, "runtime", true, function(interfaceString_, concreteString_, assertedString_, missingMethod_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.interfaceString = "";
			this.concreteString = "";
			this.assertedString = "";
			this.missingMethod = "";
			return;
		}
		this.interfaceString = interfaceString_;
		this.concreteString = concreteString_;
		this.assertedString = assertedString_;
		this.missingMethod = missingMethod_;
	});
	errorString = $pkg.errorString = $newType(8, $kindString, "runtime.errorString", true, "runtime", false, null);
	ptrType$4 = $ptrType(TypeAssertionError);
	init = function() {
		var $ptr, e, jsPkg;
		jsPkg = $packages[$externalize("github.com/gopherjs/gopherjs/js", $String)];
		$jsObjectPtr = jsPkg.Object.ptr;
		$jsErrorPtr = jsPkg.Error.ptr;
		$throwRuntimeError = throw$1;
		e = $ifaceNil;
		e = new TypeAssertionError.ptr("", "", "", "");
		$unused(e);
	};
	throw$1 = function(s) {
		var $ptr, s;
		$panic(new errorString((s)));
	};
	TypeAssertionError.ptr.prototype.RuntimeError = function() {
		var $ptr;
	};
	TypeAssertionError.prototype.RuntimeError = function() { return this.$val.RuntimeError(); };
	TypeAssertionError.ptr.prototype.Error = function() {
		var $ptr, e, inter;
		e = this;
		inter = e.interfaceString;
		if (inter === "") {
			inter = "interface";
		}
		if (e.concreteString === "") {
			return "interface conversion: " + inter + " is nil, not " + e.assertedString;
		}
		if (e.missingMethod === "") {
			return "interface conversion: " + inter + " is " + e.concreteString + ", not " + e.assertedString;
		}
		return "interface conversion: " + e.concreteString + " is not " + e.assertedString + ": missing method " + e.missingMethod;
	};
	TypeAssertionError.prototype.Error = function() { return this.$val.Error(); };
	errorString.prototype.RuntimeError = function() {
		var $ptr, e;
		e = this.$val;
	};
	$ptrType(errorString).prototype.RuntimeError = function() { return new errorString(this.$get()).RuntimeError(); };
	errorString.prototype.Error = function() {
		var $ptr, e;
		e = this.$val;
		return "runtime error: " + (e);
	};
	$ptrType(errorString).prototype.Error = function() { return new errorString(this.$get()).Error(); };
	ptrType$4.methods = [{prop: "RuntimeError", name: "RuntimeError", pkg: "", typ: $funcType([], [], false)}, {prop: "Error", name: "Error", pkg: "", typ: $funcType([], [$String], false)}];
	errorString.methods = [{prop: "RuntimeError", name: "RuntimeError", pkg: "", typ: $funcType([], [], false)}, {prop: "Error", name: "Error", pkg: "", typ: $funcType([], [$String], false)}];
	TypeAssertionError.init("runtime", [{prop: "interfaceString", name: "interfaceString", anonymous: false, exported: false, typ: $String, tag: ""}, {prop: "concreteString", name: "concreteString", anonymous: false, exported: false, typ: $String, tag: ""}, {prop: "assertedString", name: "assertedString", anonymous: false, exported: false, typ: $String, tag: ""}, {prop: "missingMethod", name: "missingMethod", anonymous: false, exported: false, typ: $String, tag: ""}]);
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		$r = js.$init(); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = sys.$init(); /* */ $s = 2; case 2: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		init();
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["github.com/oskca/gopherjs-dom"] = (function() {
	var $pkg = {}, $init, js;
	js = $packages["github.com/gopherjs/gopherjs/js"];
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		$r = js.$init(); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["github.com/oskca/gopherjs-json"] = (function() {
	var $pkg = {}, $init, js, json, Stringify, Parse;
	js = $packages["github.com/gopherjs/gopherjs/js"];
	Stringify = function(obj) {
		var $ptr, obj;
		return $internalize(json.stringify($externalize(obj, $emptyInterface)), $String);
	};
	$pkg.Stringify = Stringify;
	Parse = function(jsonStr) {
		var $ptr, jsonStr;
		return json.parse($externalize(jsonStr, $String));
	};
	$pkg.Parse = Parse;
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		$r = js.$init(); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		json = $global.JSON;
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
	(function() {
/*!
 * Vue.js v2.1.10
 * (c) 2014-2017 Evan You
 * Released under the MIT License.
 */
!function(e,t){"object"==typeof exports&&"undefined"!=typeof module?module.exports=t():"function"==typeof define&&define.amd?define(t):e.Vue=t()}(this,function(){"use strict";function e(e){return null==e?"":"object"==typeof e?JSON.stringify(e,null,2):String(e)}function t(e){var t=parseFloat(e);return isNaN(t)?e:t}function n(e,t){for(var n=Object.create(null),r=e.split(","),i=0;i<r.length;i++)n[r[i]]=!0;return t?function(e){return n[e.toLowerCase()]}:function(e){return n[e]}}function r(e,t){if(e.length){var n=e.indexOf(t);if(n>-1)return e.splice(n,1)}}function i(e,t){return ii.call(e,t)}function o(e){return"string"==typeof e||"number"==typeof e}function a(e){var t=Object.create(null);return function(n){var r=t[n];return r||(t[n]=e(n))}}function s(e,t){function n(n){var r=arguments.length;return r?r>1?e.apply(t,arguments):e.call(t,n):e.call(t)}return n._length=e.length,n}function c(e,t){t=t||0;for(var n=e.length-t,r=new Array(n);n--;)r[n]=e[n+t];return r}function u(e,t){for(var n in t)e[n]=t[n];return e}function l(e){return null!==e&&"object"==typeof e}function f(e){return li.call(e)===fi}function p(e){for(var t={},n=0;n<e.length;n++)e[n]&&u(t,e[n]);return t}function d(){}function v(e){return e.reduce(function(e,t){return e.concat(t.staticKeys||[])},[]).join(",")}function h(e,t){var n=l(e),r=l(t);return n&&r?JSON.stringify(e)===JSON.stringify(t):!n&&!r&&String(e)===String(t)}function m(e,t){for(var n=0;n<e.length;n++)if(h(e[n],t))return n;return-1}function g(e){var t=(e+"").charCodeAt(0);return 36===t||95===t}function y(e,t,n,r){Object.defineProperty(e,t,{value:n,enumerable:!!r,writable:!0,configurable:!0})}function _(e){if(!hi.test(e)){var t=e.split(".");return function(e){for(var n=0;n<t.length;n++){if(!e)return;e=e[t[n]]}return e}}}function b(e){return/native code/.test(e.toString())}function $(e){Ei.target&&Ii.push(Ei.target),Ei.target=e}function w(){Ei.target=Ii.pop()}function C(e,t){e.__proto__=t}function x(e,t,n){for(var r=0,i=n.length;r<i;r++){var o=n[r];y(e,o,t[o])}}function k(e,t){if(l(e)){var n;return i(e,"__ob__")&&e.__ob__ instanceof Di?n=e.__ob__:Mi.shouldConvert&&!xi()&&(Array.isArray(e)||f(e))&&Object.isExtensible(e)&&!e._isVue&&(n=new Di(e)),t&&n&&n.vmCount++,n}}function A(e,t,n,r){var i=new Ei,o=Object.getOwnPropertyDescriptor(e,t);if(!o||o.configurable!==!1){var a=o&&o.get,s=o&&o.set,c=k(n);Object.defineProperty(e,t,{enumerable:!0,configurable:!0,get:function(){var t=a?a.call(e):n;return Ei.target&&(i.depend(),c&&c.dep.depend(),Array.isArray(t)&&T(t)),t},set:function(t){var r=a?a.call(e):n;t===r||t!==t&&r!==r||(s?s.call(e,t):n=t,c=k(t),i.notify())}})}}function O(e,t,n){if(Array.isArray(e))return e.length=Math.max(e.length,t),e.splice(t,1,n),n;if(i(e,t))return void(e[t]=n);var r=e.__ob__;if(!(e._isVue||r&&r.vmCount))return r?(A(r.value,t,n),r.dep.notify(),n):void(e[t]=n)}function S(e,t){var n=e.__ob__;e._isVue||n&&n.vmCount||i(e,t)&&(delete e[t],n&&n.dep.notify())}function T(e){for(var t=void 0,n=0,r=e.length;n<r;n++)t=e[n],t&&t.__ob__&&t.__ob__.dep.depend(),Array.isArray(t)&&T(t)}function E(e,t){if(!t)return e;for(var n,r,o,a=Object.keys(t),s=0;s<a.length;s++)n=a[s],r=e[n],o=t[n],i(e,n)?f(r)&&f(o)&&E(r,o):O(e,n,o);return e}function I(e,t){return t?e?e.concat(t):Array.isArray(t)?t:[t]:e}function j(e,t){var n=Object.create(e||null);return t?u(n,t):n}function N(e){var t=e.props;if(t){var n,r,i,o={};if(Array.isArray(t))for(n=t.length;n--;)r=t[n],"string"==typeof r&&(i=ai(r),o[i]={type:null});else if(f(t))for(var a in t)r=t[a],i=ai(a),o[i]=f(r)?r:{type:r};e.props=o}}function L(e){var t=e.directives;if(t)for(var n in t){var r=t[n];"function"==typeof r&&(t[n]={bind:r,update:r})}}function M(e,t,n){function r(r){var i=Pi[r]||Ri;l[r]=i(e[r],t[r],n,r)}N(t),L(t);var o=t.extends;if(o&&(e="function"==typeof o?M(e,o.options,n):M(e,o,n)),t.mixins)for(var a=0,s=t.mixins.length;a<s;a++){var c=t.mixins[a];c.prototype instanceof Ue&&(c=c.options),e=M(e,c,n)}var u,l={};for(u in e)r(u);for(u in t)i(e,u)||r(u);return l}function D(e,t,n,r){if("string"==typeof n){var o=e[t];if(i(o,n))return o[n];var a=ai(n);if(i(o,a))return o[a];var s=si(a);if(i(o,s))return o[s];var c=o[n]||o[a]||o[s];return c}}function P(e,t,n,r){var o=t[e],a=!i(n,e),s=n[e];if(H(Boolean,o.type)&&(a&&!i(o,"default")?s=!1:H(String,o.type)||""!==s&&s!==ui(e)||(s=!0)),void 0===s){s=R(r,o,e);var c=Mi.shouldConvert;Mi.shouldConvert=!0,k(s),Mi.shouldConvert=c}return s}function R(e,t,n){if(i(t,"default")){var r=t.default;return l(r),e&&e.$options.propsData&&void 0===e.$options.propsData[n]&&void 0!==e[n]?e[n]:"function"==typeof r&&t.type!==Function?r.call(e):r}}function F(e){var t=e&&e.toString().match(/^\s*function (\w+)/);return t&&t[1]}function H(e,t){if(!Array.isArray(t))return F(t)===F(e);for(var n=0,r=t.length;n<r;n++)if(F(t[n])===F(e))return!0;return!1}function U(e){return new Hi(void 0,void 0,void 0,String(e))}function B(e){var t=new Hi(e.tag,e.data,e.children,e.text,e.elm,e.context,e.componentOptions);return t.ns=e.ns,t.isStatic=e.isStatic,t.key=e.key,t.isCloned=!0,t}function z(e){for(var t=new Array(e.length),n=0;n<e.length;n++)t[n]=B(e[n]);return t}function V(e,t,n,r,i){if(e){var o=n.$options._base;if(l(e)&&(e=o.extend(e)),"function"==typeof e){if(!e.cid)if(e.resolved)e=e.resolved;else if(e=Y(e,o,function(){n.$forceUpdate()}),!e)return;He(e),t=t||{};var a=Q(t,e);if(e.options.functional)return J(e,a,t,n,r);var s=t.on;t.on=t.nativeOn,e.options.abstract&&(t={}),ee(t);var c=e.options.name||i,u=new Hi("vue-component-"+e.cid+(c?"-"+c:""),t,void 0,void 0,void 0,n,{Ctor:e,propsData:a,listeners:s,tag:i,children:r});return u}}}function J(e,t,n,r,i){var o={},a=e.options.props;if(a)for(var s in a)o[s]=P(s,a,t);var c=Object.create(r),u=function(e,t,n,r){return ue(c,e,t,n,r,!0)},l=e.options.render.call(null,u,{props:o,data:n,parent:r,children:i,slots:function(){return ve(i,r)}});return l instanceof Hi&&(l.functionalContext=r,n.slot&&((l.data||(l.data={})).slot=n.slot)),l}function K(e,t,n,r){var i=e.componentOptions,o={_isComponent:!0,parent:t,propsData:i.propsData,_componentTag:i.tag,_parentVnode:e,_parentListeners:i.listeners,_renderChildren:i.children,_parentElm:n||null,_refElm:r||null},a=e.data.inlineTemplate;return a&&(o.render=a.render,o.staticRenderFns=a.staticRenderFns),new i.Ctor(o)}function q(e,t,n,r){if(!e.componentInstance||e.componentInstance._isDestroyed){var i=e.componentInstance=K(e,Zi,n,r);i.$mount(t?e.elm:void 0,t)}else if(e.data.keepAlive){var o=e;W(o,o)}}function W(e,t){var n=t.componentOptions,r=t.componentInstance=e.componentInstance;r._updateFromParent(n.propsData,n.listeners,t,n.children)}function Z(e){e.componentInstance._isMounted||(e.componentInstance._isMounted=!0,we(e.componentInstance,"mounted")),e.data.keepAlive&&(e.componentInstance._inactive=!1,we(e.componentInstance,"activated"))}function G(e){e.componentInstance._isDestroyed||(e.data.keepAlive?(e.componentInstance._inactive=!0,we(e.componentInstance,"deactivated")):e.componentInstance.$destroy())}function Y(e,t,n){if(!e.requested){e.requested=!0;var r=e.pendingCallbacks=[n],i=!0,o=function(n){if(l(n)&&(n=t.extend(n)),e.resolved=n,!i)for(var o=0,a=r.length;o<a;o++)r[o](n)},a=function(e){},s=e(o,a);return s&&"function"==typeof s.then&&!e.resolved&&s.then(o,a),i=!1,e.resolved}e.pendingCallbacks.push(n)}function Q(e,t){var n=t.options.props;if(n){var r={},i=e.attrs,o=e.props,a=e.domProps;if(i||o||a)for(var s in n){var c=ui(s);X(r,o,s,c,!0)||X(r,i,s,c)||X(r,a,s,c)}return r}}function X(e,t,n,r,o){if(t){if(i(t,n))return e[n]=t[n],o||delete t[n],!0;if(i(t,r))return e[n]=t[r],o||delete t[r],!0}return!1}function ee(e){e.hook||(e.hook={});for(var t=0;t<Ji.length;t++){var n=Ji[t],r=e.hook[n],i=Vi[n];e.hook[n]=r?te(i,r):i}}function te(e,t){return function(n,r,i,o){e(n,r,i,o),t(n,r,i,o)}}function ne(e,t,n,r){r+=t;var i=e.__injected||(e.__injected={});if(!i[r]){i[r]=!0;var o=e[t];o?e[t]=function(){o.apply(this,arguments),n.apply(this,arguments)}:e[t]=n}}function re(e){var t={fn:e,invoker:function(){var e=arguments,n=t.fn;if(Array.isArray(n))for(var r=0;r<n.length;r++)n[r].apply(null,e);else n.apply(null,arguments)}};return t}function ie(e,t,n,r,i){var o,a,s,c;for(o in e)a=e[o],s=t[o],c=Ki(o),a&&(s?a!==s&&(s.fn=a,e[o]=s):(a.invoker||(a=e[o]=re(a)),n(c.name,a.invoker,c.once,c.capture)));for(o in t)e[o]||(c=Ki(o),r(c.name,t[o].invoker,c.capture))}function oe(e){for(var t=0;t<e.length;t++)if(Array.isArray(e[t]))return Array.prototype.concat.apply([],e);return e}function ae(e){return o(e)?[U(e)]:Array.isArray(e)?se(e):void 0}function se(e,t){var n,r,i,a=[];for(n=0;n<e.length;n++)r=e[n],null!=r&&"boolean"!=typeof r&&(i=a[a.length-1],Array.isArray(r)?a.push.apply(a,se(r,(t||"")+"_"+n)):o(r)?i&&i.text?i.text+=String(r):""!==r&&a.push(U(r)):r.text&&i&&i.text?a[a.length-1]=U(i.text+r.text):(r.tag&&null==r.key&&null!=t&&(r.key="__vlist"+t+"_"+n+"__"),a.push(r)));return a}function ce(e){return e&&e.filter(function(e){return e&&e.componentOptions})[0]}function ue(e,t,n,r,i,a){return(Array.isArray(n)||o(n))&&(i=r,r=n,n=void 0),a&&(i=Wi),le(e,t,n,r,i)}function le(e,t,n,r,i){if(n&&n.__ob__)return zi();if(!t)return zi();Array.isArray(r)&&"function"==typeof r[0]&&(n=n||{},n.scopedSlots={default:r[0]},r.length=0),i===Wi?r=ae(r):i===qi&&(r=oe(r));var o,a;if("string"==typeof t){var s;a=vi.getTagNamespace(t),o=vi.isReservedTag(t)?new Hi(vi.parsePlatformTagName(t),n,r,void 0,void 0,e):(s=D(e.$options,"components",t))?V(s,n,e,r,t):new Hi(t,n,r,void 0,void 0,e)}else o=V(t,n,e,r);return o?(a&&fe(o,a),o):zi()}function fe(e,t){if(e.ns=t,"foreignObject"!==e.tag&&e.children)for(var n=0,r=e.children.length;n<r;n++){var i=e.children[n];i.tag&&!i.ns&&fe(i,t)}}function pe(e){e.$vnode=null,e._vnode=null,e._staticTrees=null;var t=e.$options._parentVnode,n=t&&t.context;e.$slots=ve(e.$options._renderChildren,n),e.$scopedSlots={},e._c=function(t,n,r,i){return ue(e,t,n,r,i,!1)},e.$createElement=function(t,n,r,i){return ue(e,t,n,r,i,!0)}}function de(n){function r(e,t,n){if(Array.isArray(e))for(var r=0;r<e.length;r++)e[r]&&"string"!=typeof e[r]&&i(e[r],t+"_"+r,n);else i(e,t,n)}function i(e,t,n){e.isStatic=!0,e.key=t,e.isOnce=n}n.prototype.$nextTick=function(e){return Ai(e,this)},n.prototype._render=function(){var e=this,t=e.$options,n=t.render,r=t.staticRenderFns,i=t._parentVnode;if(e._isMounted)for(var o in e.$slots)e.$slots[o]=z(e.$slots[o]);i&&i.data.scopedSlots&&(e.$scopedSlots=i.data.scopedSlots),r&&!e._staticTrees&&(e._staticTrees=[]),e.$vnode=i;var a;try{a=n.call(e._renderProxy,e.$createElement)}catch(t){if(!vi.errorHandler)throw t;vi.errorHandler.call(null,t,e),a=e._vnode}return a instanceof Hi||(a=zi()),a.parent=i,a},n.prototype._s=e,n.prototype._v=U,n.prototype._n=t,n.prototype._e=zi,n.prototype._q=h,n.prototype._i=m,n.prototype._m=function(e,t){var n=this._staticTrees[e];return n&&!t?Array.isArray(n)?z(n):B(n):(n=this._staticTrees[e]=this.$options.staticRenderFns[e].call(this._renderProxy),r(n,"__static__"+e,!1),n)},n.prototype._o=function(e,t,n){return r(e,"__once__"+t+(n?"_"+n:""),!0),e},n.prototype._f=function(e){return D(this.$options,"filters",e,!0)||di},n.prototype._l=function(e,t){var n,r,i,o,a;if(Array.isArray(e)||"string"==typeof e)for(n=new Array(e.length),r=0,i=e.length;r<i;r++)n[r]=t(e[r],r);else if("number"==typeof e)for(n=new Array(e),r=0;r<e;r++)n[r]=t(r+1,r);else if(l(e))for(o=Object.keys(e),n=new Array(o.length),r=0,i=o.length;r<i;r++)a=o[r],n[r]=t(e[a],a,r);return n},n.prototype._t=function(e,t,n,r){var i=this.$scopedSlots[e];if(i)return n=n||{},r&&u(n,r),i(n)||t;var o=this.$slots[e];return o||t},n.prototype._b=function(e,t,n,r){if(n)if(l(n)){Array.isArray(n)&&(n=p(n));for(var i in n)if("class"===i||"style"===i)e[i]=n[i];else{var o=e.attrs&&e.attrs.type,a=r||vi.mustUseProp(t,o,i)?e.domProps||(e.domProps={}):e.attrs||(e.attrs={});a[i]=n[i]}}else;return e},n.prototype._k=function(e,t,n){var r=vi.keyCodes[t]||n;return Array.isArray(r)?r.indexOf(e)===-1:r!==e}}function ve(e,t){var n={};if(!e)return n;for(var r,i,o=[],a=0,s=e.length;a<s;a++)if(i=e[a],(i.context===t||i.functionalContext===t)&&i.data&&(r=i.data.slot)){var c=n[r]||(n[r]=[]);"template"===i.tag?c.push.apply(c,i.children):c.push(i)}else o.push(i);return o.length&&(1!==o.length||" "!==o[0].text&&!o[0].isComment)&&(n.default=o),n}function he(e){e._events=Object.create(null),e._hasHookEvent=!1;var t=e.$options._parentListeners;t&&ye(e,t)}function me(e,t,n){n?Bi.$once(e,t):Bi.$on(e,t)}function ge(e,t){Bi.$off(e,t)}function ye(e,t,n){Bi=e,ie(t,n||{},me,ge,e)}function _e(e){var t=/^hook:/;e.prototype.$on=function(e,n){var r=this;return(r._events[e]||(r._events[e]=[])).push(n),t.test(e)&&(r._hasHookEvent=!0),r},e.prototype.$once=function(e,t){function n(){r.$off(e,n),t.apply(r,arguments)}var r=this;return n.fn=t,r.$on(e,n),r},e.prototype.$off=function(e,t){var n=this;if(!arguments.length)return n._events=Object.create(null),n;var r=n._events[e];if(!r)return n;if(1===arguments.length)return n._events[e]=null,n;for(var i,o=r.length;o--;)if(i=r[o],i===t||i.fn===t){r.splice(o,1);break}return n},e.prototype.$emit=function(e){var t=this,n=t._events[e];if(n){n=n.length>1?c(n):n;for(var r=c(arguments,1),i=0,o=n.length;i<o;i++)n[i].apply(t,r)}return t}}function be(e){var t=e.$options,n=t.parent;if(n&&!t.abstract){for(;n.$options.abstract&&n.$parent;)n=n.$parent;n.$children.push(e)}e.$parent=n,e.$root=n?n.$root:e,e.$children=[],e.$refs={},e._watcher=null,e._inactive=!1,e._isMounted=!1,e._isDestroyed=!1,e._isBeingDestroyed=!1}function $e(e){e.prototype._mount=function(e,t){var n=this;return n.$el=e,n.$options.render||(n.$options.render=zi),we(n,"beforeMount"),n._watcher=new no(n,function(){n._update(n._render(),t)},d),t=!1,null==n.$vnode&&(n._isMounted=!0,we(n,"mounted")),n},e.prototype._update=function(e,t){var n=this;n._isMounted&&we(n,"beforeUpdate");var r=n.$el,i=n._vnode,o=Zi;Zi=n,n._vnode=e,i?n.$el=n.__patch__(i,e):n.$el=n.__patch__(n.$el,e,t,!1,n.$options._parentElm,n.$options._refElm),Zi=o,r&&(r.__vue__=null),n.$el&&(n.$el.__vue__=n),n.$vnode&&n.$parent&&n.$vnode===n.$parent._vnode&&(n.$parent.$el=n.$el)},e.prototype._updateFromParent=function(e,t,n,r){var i=this,o=!(!i.$options._renderChildren&&!r);if(i.$options._parentVnode=n,i.$vnode=n,i._vnode&&(i._vnode.parent=n),i.$options._renderChildren=r,e&&i.$options.props){Mi.shouldConvert=!1;for(var a=i.$options._propKeys||[],s=0;s<a.length;s++){var c=a[s];i[c]=P(c,i.$options.props,e,i)}Mi.shouldConvert=!0,i.$options.propsData=e}if(t){var u=i.$options._parentListeners;i.$options._parentListeners=t,ye(i,t,u)}o&&(i.$slots=ve(r,n.context),i.$forceUpdate())},e.prototype.$forceUpdate=function(){var e=this;e._watcher&&e._watcher.update()},e.prototype.$destroy=function(){var e=this;if(!e._isBeingDestroyed){we(e,"beforeDestroy"),e._isBeingDestroyed=!0;var t=e.$parent;!t||t._isBeingDestroyed||e.$options.abstract||r(t.$children,e),e._watcher&&e._watcher.teardown();for(var n=e._watchers.length;n--;)e._watchers[n].teardown();e._data.__ob__&&e._data.__ob__.vmCount--,e._isDestroyed=!0,we(e,"destroyed"),e.$off(),e.$el&&(e.$el.__vue__=null),e.__patch__(e._vnode,null)}}}function we(e,t){var n=e.$options[t];if(n)for(var r=0,i=n.length;r<i;r++)n[r].call(e);e._hasHookEvent&&e.$emit("hook:"+t)}function Ce(){Gi.length=0,Yi={},Qi=Xi=!1}function xe(){Xi=!0;var e,t,n;for(Gi.sort(function(e,t){return e.id-t.id}),eo=0;eo<Gi.length;eo++)e=Gi[eo],t=e.id,Yi[t]=null,e.run();for(eo=Gi.length;eo--;)e=Gi[eo],n=e.vm,n._watcher===e&&n._isMounted&&we(n,"updated");ki&&vi.devtools&&ki.emit("flush"),Ce()}function ke(e){var t=e.id;if(null==Yi[t]){if(Yi[t]=!0,Xi){for(var n=Gi.length-1;n>=0&&Gi[n].id>e.id;)n--;Gi.splice(Math.max(n,eo)+1,0,e)}else Gi.push(e);Qi||(Qi=!0,Ai(xe))}}function Ae(e){ro.clear(),Oe(e,ro)}function Oe(e,t){var n,r,i=Array.isArray(e);if((i||l(e))&&Object.isExtensible(e)){if(e.__ob__){var o=e.__ob__.dep.id;if(t.has(o))return;t.add(o)}if(i)for(n=e.length;n--;)Oe(e[n],t);else for(r=Object.keys(e),n=r.length;n--;)Oe(e[r[n]],t)}}function Se(e){e._watchers=[];var t=e.$options;t.props&&Te(e,t.props),t.methods&&Ne(e,t.methods),t.data?Ee(e):k(e._data={},!0),t.computed&&Ie(e,t.computed),t.watch&&Le(e,t.watch)}function Te(e,t){var n=e.$options.propsData||{},r=e.$options._propKeys=Object.keys(t),i=!e.$parent;Mi.shouldConvert=i;for(var o=function(i){var o=r[i];A(e,o,P(o,t,n,e))},a=0;a<r.length;a++)o(a);Mi.shouldConvert=!0}function Ee(e){var t=e.$options.data;t=e._data="function"==typeof t?t.call(e):t||{},f(t)||(t={});for(var n=Object.keys(t),r=e.$options.props,o=n.length;o--;)r&&i(r,n[o])||Pe(e,n[o]);k(t,!0)}function Ie(e,t){for(var n in t){var r=t[n];"function"==typeof r?(io.get=je(r,e),io.set=d):(io.get=r.get?r.cache!==!1?je(r.get,e):s(r.get,e):d,io.set=r.set?s(r.set,e):d),Object.defineProperty(e,n,io)}}function je(e,t){var n=new no(t,e,d,{lazy:!0});return function(){return n.dirty&&n.evaluate(),Ei.target&&n.depend(),n.value}}function Ne(e,t){for(var n in t)e[n]=null==t[n]?d:s(t[n],e)}function Le(e,t){for(var n in t){var r=t[n];if(Array.isArray(r))for(var i=0;i<r.length;i++)Me(e,n,r[i]);else Me(e,n,r)}}function Me(e,t,n){var r;f(n)&&(r=n,n=n.handler),"string"==typeof n&&(n=e[n]),e.$watch(t,n,r)}function De(e){var t={};t.get=function(){return this._data},Object.defineProperty(e.prototype,"$data",t),e.prototype.$set=O,e.prototype.$delete=S,e.prototype.$watch=function(e,t,n){var r=this;n=n||{},n.user=!0;var i=new no(r,e,t,n);return n.immediate&&t.call(r,i.value),function(){i.teardown()}}}function Pe(e,t){g(t)||Object.defineProperty(e,t,{configurable:!0,enumerable:!0,get:function(){return e._data[t]},set:function(n){e._data[t]=n}})}function Re(e){e.prototype._init=function(e){var t=this;t._uid=oo++,t._isVue=!0,e&&e._isComponent?Fe(t,e):t.$options=M(He(t.constructor),e||{},t),t._renderProxy=t,t._self=t,be(t),he(t),pe(t),we(t,"beforeCreate"),Se(t),we(t,"created"),t.$options.el&&t.$mount(t.$options.el)}}function Fe(e,t){var n=e.$options=Object.create(e.constructor.options);n.parent=t.parent,n.propsData=t.propsData,n._parentVnode=t._parentVnode,n._parentListeners=t._parentListeners,n._renderChildren=t._renderChildren,n._componentTag=t._componentTag,n._parentElm=t._parentElm,n._refElm=t._refElm,t.render&&(n.render=t.render,n.staticRenderFns=t.staticRenderFns)}function He(e){var t=e.options;if(e.super){var n=e.super.options,r=e.superOptions,i=e.extendOptions;n!==r&&(e.superOptions=n,i.render=t.render,i.staticRenderFns=t.staticRenderFns,i._scopeId=t._scopeId,t=e.options=M(n,i),t.name&&(t.components[t.name]=e))}return t}function Ue(e){this._init(e)}function Be(e){e.use=function(e){if(!e.installed){var t=c(arguments,1);return t.unshift(this),"function"==typeof e.install?e.install.apply(e,t):e.apply(null,t),e.installed=!0,this}}}function ze(e){e.mixin=function(e){this.options=M(this.options,e)}}function Ve(e){e.cid=0;var t=1;e.extend=function(e){e=e||{};var n=this,r=n.cid,i=e._Ctor||(e._Ctor={});if(i[r])return i[r];var o=e.name||n.options.name,a=function(e){this._init(e)};return a.prototype=Object.create(n.prototype),a.prototype.constructor=a,a.cid=t++,a.options=M(n.options,e),a.super=n,a.extend=n.extend,a.mixin=n.mixin,a.use=n.use,vi._assetTypes.forEach(function(e){a[e]=n[e]}),o&&(a.options.components[o]=a),a.superOptions=n.options,a.extendOptions=e,i[r]=a,a}}function Je(e){vi._assetTypes.forEach(function(t){e[t]=function(e,n){return n?("component"===t&&f(n)&&(n.name=n.name||e,n=this.options._base.extend(n)),"directive"===t&&"function"==typeof n&&(n={bind:n,update:n}),this.options[t+"s"][e]=n,n):this.options[t+"s"][e]}})}function Ke(e){return e&&(e.Ctor.options.name||e.tag)}function qe(e,t){return"string"==typeof e?e.split(",").indexOf(t)>-1:e.test(t)}function We(e,t){for(var n in e){var r=e[n];if(r){var i=Ke(r.componentOptions);i&&!t(i)&&(Ze(r),e[n]=null)}}}function Ze(e){e&&(e.componentInstance._inactive||we(e.componentInstance,"deactivated"),e.componentInstance.$destroy())}function Ge(e){var t={};t.get=function(){return vi},Object.defineProperty(e,"config",t),e.util=Fi,e.set=O,e.delete=S,e.nextTick=Ai,e.options=Object.create(null),vi._assetTypes.forEach(function(t){e.options[t+"s"]=Object.create(null)}),e.options._base=e,u(e.options.components,co),Be(e),ze(e),Ve(e),Je(e)}function Ye(e){for(var t=e.data,n=e,r=e;r.componentInstance;)r=r.componentInstance._vnode,r.data&&(t=Qe(r.data,t));for(;n=n.parent;)n.data&&(t=Qe(t,n.data));return Xe(t)}function Qe(e,t){return{staticClass:et(e.staticClass,t.staticClass),class:e.class?[e.class,t.class]:t.class}}function Xe(e){var t=e.class,n=e.staticClass;return n||t?et(n,tt(t)):""}function et(e,t){return e?t?e+" "+t:e:t||""}function tt(e){var t="";if(!e)return t;if("string"==typeof e)return e;if(Array.isArray(e)){for(var n,r=0,i=e.length;r<i;r++)e[r]&&(n=tt(e[r]))&&(t+=n+" ");return t.slice(0,-1)}if(l(e)){for(var o in e)e[o]&&(t+=o+" ");return t.slice(0,-1)}return t}function nt(e){return wo(e)?"svg":"math"===e?"math":void 0}function rt(e){if(!gi)return!0;if(xo(e))return!1;if(e=e.toLowerCase(),null!=ko[e])return ko[e];var t=document.createElement(e);return e.indexOf("-")>-1?ko[e]=t.constructor===window.HTMLUnknownElement||t.constructor===window.HTMLElement:ko[e]=/HTMLUnknownElement/.test(t.toString())}function it(e){if("string"==typeof e){if(e=document.querySelector(e),!e)return document.createElement("div")}return e}function ot(e,t){var n=document.createElement(e);return"select"!==e?n:(t.data&&t.data.attrs&&"multiple"in t.data.attrs&&n.setAttribute("multiple","multiple"),n)}function at(e,t){return document.createElementNS(bo[e],t)}function st(e){return document.createTextNode(e)}function ct(e){return document.createComment(e)}function ut(e,t,n){e.insertBefore(t,n)}function lt(e,t){e.removeChild(t)}function ft(e,t){e.appendChild(t)}function pt(e){return e.parentNode}function dt(e){return e.nextSibling}function vt(e){return e.tagName}function ht(e,t){e.textContent=t}function mt(e,t,n){e.setAttribute(t,n)}function gt(e,t){var n=e.data.ref;if(n){var i=e.context,o=e.componentInstance||e.elm,a=i.$refs;t?Array.isArray(a[n])?r(a[n],o):a[n]===o&&(a[n]=void 0):e.data.refInFor?Array.isArray(a[n])&&a[n].indexOf(o)<0?a[n].push(o):a[n]=[o]:a[n]=o}}function yt(e){return null==e}function _t(e){return null!=e}function bt(e,t){return e.key===t.key&&e.tag===t.tag&&e.isComment===t.isComment&&!e.data==!t.data}function $t(e,t,n){var r,i,o={};for(r=t;r<=n;++r)i=e[r].key,_t(i)&&(o[i]=r);return o}function wt(e){function t(e){return new Hi(O.tagName(e).toLowerCase(),{},[],void 0,e)}function r(e,t){function n(){0===--n.listeners&&i(e)}return n.listeners=t,n}function i(e){var t=O.parentNode(e);t&&O.removeChild(t,e)}function a(e,t,n,r,i){if(e.isRootInsert=!i,!s(e,t,n,r)){var o=e.data,a=e.children,c=e.tag;_t(c)?(e.elm=e.ns?O.createElementNS(e.ns,c):O.createElement(c,e),v(e),f(e,a,t),_t(o)&&d(e,t),l(n,e.elm,r)):e.isComment?(e.elm=O.createComment(e.text),l(n,e.elm,r)):(e.elm=O.createTextNode(e.text),l(n,e.elm,r))}}function s(e,t,n,r){var i=e.data;if(_t(i)){var o=_t(e.componentInstance)&&i.keepAlive;if(_t(i=i.hook)&&_t(i=i.init)&&i(e,!1,n,r),_t(e.componentInstance))return c(e,t),o&&u(e,t,n,r),!0}}function c(e,t){e.data.pendingInsert&&t.push.apply(t,e.data.pendingInsert),e.elm=e.componentInstance.$el,p(e)?(d(e,t),v(e)):(gt(e),t.push(e))}function u(e,t,n,r){for(var i,o=e;o.componentInstance;)if(o=o.componentInstance._vnode,_t(i=o.data)&&_t(i=i.transition)){for(i=0;i<k.activate.length;++i)k.activate[i](So,o);t.push(o);break}l(n,e.elm,r)}function l(e,t,n){e&&(n?O.insertBefore(e,t,n):O.appendChild(e,t))}function f(e,t,n){if(Array.isArray(t))for(var r=0;r<t.length;++r)a(t[r],n,e.elm,null,!0);else o(e.text)&&O.appendChild(e.elm,O.createTextNode(e.text))}function p(e){for(;e.componentInstance;)e=e.componentInstance._vnode;return _t(e.tag)}function d(e,t){for(var n=0;n<k.create.length;++n)k.create[n](So,e);C=e.data.hook,_t(C)&&(C.create&&C.create(So,e),C.insert&&t.push(e))}function v(e){var t;_t(t=e.context)&&_t(t=t.$options._scopeId)&&O.setAttribute(e.elm,t,""),_t(t=Zi)&&t!==e.context&&_t(t=t.$options._scopeId)&&O.setAttribute(e.elm,t,"")}function h(e,t,n,r,i,o){for(;r<=i;++r)a(n[r],o,e,t)}function m(e){var t,n,r=e.data;if(_t(r))for(_t(t=r.hook)&&_t(t=t.destroy)&&t(e),t=0;t<k.destroy.length;++t)k.destroy[t](e);if(_t(t=e.children))for(n=0;n<e.children.length;++n)m(e.children[n])}function g(e,t,n,r){for(;n<=r;++n){var o=t[n];_t(o)&&(_t(o.tag)?(y(o),m(o)):i(o.elm))}}function y(e,t){if(t||_t(e.data)){var n=k.remove.length+1;for(t?t.listeners+=n:t=r(e.elm,n),_t(C=e.componentInstance)&&_t(C=C._vnode)&&_t(C.data)&&y(C,t),C=0;C<k.remove.length;++C)k.remove[C](e,t);_t(C=e.data.hook)&&_t(C=C.remove)?C(e,t):t()}else i(e.elm)}function _(e,t,n,r,i){for(var o,s,c,u,l=0,f=0,p=t.length-1,d=t[0],v=t[p],m=n.length-1,y=n[0],_=n[m],$=!i;l<=p&&f<=m;)yt(d)?d=t[++l]:yt(v)?v=t[--p]:bt(d,y)?(b(d,y,r),d=t[++l],y=n[++f]):bt(v,_)?(b(v,_,r),v=t[--p],_=n[--m]):bt(d,_)?(b(d,_,r),$&&O.insertBefore(e,d.elm,O.nextSibling(v.elm)),d=t[++l],_=n[--m]):bt(v,y)?(b(v,y,r),$&&O.insertBefore(e,v.elm,d.elm),v=t[--p],y=n[++f]):(yt(o)&&(o=$t(t,l,p)),s=_t(y.key)?o[y.key]:null,yt(s)?(a(y,r,e,d.elm),y=n[++f]):(c=t[s],bt(c,y)?(b(c,y,r),t[s]=void 0,$&&O.insertBefore(e,y.elm,d.elm),y=n[++f]):(a(y,r,e,d.elm),y=n[++f])));l>p?(u=yt(n[m+1])?null:n[m+1].elm,h(e,u,n,f,m,r)):f>m&&g(e,t,l,p)}function b(e,t,n,r){if(e!==t){if(t.isStatic&&e.isStatic&&t.key===e.key&&(t.isCloned||t.isOnce))return t.elm=e.elm,void(t.componentInstance=e.componentInstance);var i,o=t.data,a=_t(o);a&&_t(i=o.hook)&&_t(i=i.prepatch)&&i(e,t);var s=t.elm=e.elm,c=e.children,u=t.children;if(a&&p(t)){for(i=0;i<k.update.length;++i)k.update[i](e,t);_t(i=o.hook)&&_t(i=i.update)&&i(e,t)}yt(t.text)?_t(c)&&_t(u)?c!==u&&_(s,c,u,n,r):_t(u)?(_t(e.text)&&O.setTextContent(s,""),h(s,null,u,0,u.length-1,n)):_t(c)?g(s,c,0,c.length-1):_t(e.text)&&O.setTextContent(s,""):e.text!==t.text&&O.setTextContent(s,t.text),a&&_t(i=o.hook)&&_t(i=i.postpatch)&&i(e,t)}}function $(e,t,n){if(n&&e.parent)e.parent.data.pendingInsert=t;else for(var r=0;r<t.length;++r)t[r].data.hook.insert(t[r])}function w(e,t,n){t.elm=e;var r=t.tag,i=t.data,o=t.children;if(_t(i)&&(_t(C=i.hook)&&_t(C=C.init)&&C(t,!0),_t(C=t.componentInstance)))return c(t,n),!0;if(_t(r)){if(_t(o))if(e.hasChildNodes()){for(var a=!0,s=e.firstChild,u=0;u<o.length;u++){if(!s||!w(s,o[u],n)){a=!1;break}s=s.nextSibling}if(!a||s)return!1}else f(t,o,n);if(_t(i))for(var l in i)if(!S(l)){d(t,n);break}}else e.data!==t.text&&(e.data=t.text);return!0}var C,x,k={},A=e.modules,O=e.nodeOps;for(C=0;C<To.length;++C)for(k[To[C]]=[],x=0;x<A.length;++x)void 0!==A[x][To[C]]&&k[To[C]].push(A[x][To[C]]);var S=n("attrs,style,class,staticClass,staticStyle,key");return function(e,n,r,i,o,s){if(!n)return void(e&&m(e));var c=!1,u=[];if(e){var l=_t(e.nodeType);if(!l&&bt(e,n))b(e,n,u,i);else{if(l){if(1===e.nodeType&&e.hasAttribute("server-rendered")&&(e.removeAttribute("server-rendered"),r=!0),r&&w(e,n,u))return $(n,u,!0),e;e=t(e)}var f=e.elm,d=O.parentNode(f);if(a(n,u,f._leaveCb?null:d,O.nextSibling(f)),n.parent){for(var v=n.parent;v;)v.elm=n.elm,v=v.parent;if(p(n))for(var h=0;h<k.create.length;++h)k.create[h](So,n.parent)}null!==d?g(d,[e],0,0):_t(e.tag)&&m(e)}}else c=!0,a(n,u,o,s);return $(n,u,c),n.elm}}function Ct(e,t){(e.data.directives||t.data.directives)&&xt(e,t)}function xt(e,t){var n,r,i,o=e===So,a=t===So,s=kt(e.data.directives,e.context),c=kt(t.data.directives,t.context),u=[],l=[];for(n in c)r=s[n],i=c[n],r?(i.oldValue=r.value,Ot(i,"update",t,e),i.def&&i.def.componentUpdated&&l.push(i)):(Ot(i,"bind",t,e),i.def&&i.def.inserted&&u.push(i));if(u.length){var f=function(){for(var n=0;n<u.length;n++)Ot(u[n],"inserted",t,e)};o?ne(t.data.hook||(t.data.hook={}),"insert",f,"dir-insert"):f()}if(l.length&&ne(t.data.hook||(t.data.hook={}),"postpatch",function(){for(var n=0;n<l.length;n++)Ot(l[n],"componentUpdated",t,e)},"dir-postpatch"),!o)for(n in s)c[n]||Ot(s[n],"unbind",e,e,a)}function kt(e,t){var n=Object.create(null);if(!e)return n;var r,i;for(r=0;r<e.length;r++)i=e[r],i.modifiers||(i.modifiers=Io),n[At(i)]=i,i.def=D(t.$options,"directives",i.name,!0);return n}function At(e){return e.rawName||e.name+"."+Object.keys(e.modifiers||{}).join(".")}function Ot(e,t,n,r,i){var o=e.def&&e.def[t];o&&o(n.elm,e,n,r,i)}function St(e,t){if(e.data.attrs||t.data.attrs){var n,r,i,o=t.elm,a=e.data.attrs||{},s=t.data.attrs||{};s.__ob__&&(s=t.data.attrs=u({},s));for(n in s)r=s[n],i=a[n],i!==r&&Tt(o,n,r);bi&&s.value!==a.value&&Tt(o,"value",s.value);for(n in a)null==s[n]&&(go(n)?o.removeAttributeNS(mo,yo(n)):vo(n)||o.removeAttribute(n))}}function Tt(e,t,n){ho(t)?_o(n)?e.removeAttribute(t):e.setAttribute(t,t):vo(t)?e.setAttribute(t,_o(n)||"false"===n?"false":"true"):go(t)?_o(n)?e.removeAttributeNS(mo,yo(t)):e.setAttributeNS(mo,t,n):_o(n)?e.removeAttribute(t):e.setAttribute(t,n)}function Et(e,t){var n=t.elm,r=t.data,i=e.data;if(r.staticClass||r.class||i&&(i.staticClass||i.class)){var o=Ye(t),a=n._transitionClasses;a&&(o=et(o,tt(a))),o!==n._prevClass&&(n.setAttribute("class",o),n._prevClass=o)}}function It(e,t,n,r){if(n){var i=t,o=uo;t=function(n){jt(e,t,r,o),1===arguments.length?i(n):i.apply(null,arguments)}}uo.addEventListener(e,t,r)}function jt(e,t,n,r){(r||uo).removeEventListener(e,t,n)}function Nt(e,t){if(e.data.on||t.data.on){var n=t.data.on||{},r=e.data.on||{};uo=t.elm,ie(n,r,It,jt,t.context)}}function Lt(e,t){if(e.data.domProps||t.data.domProps){var n,r,i=t.elm,o=e.data.domProps||{},a=t.data.domProps||{};a.__ob__&&(a=t.data.domProps=u({},a));for(n in o)null==a[n]&&(i[n]="");for(n in a)if(r=a[n],"textContent"!==n&&"innerHTML"!==n||(t.children&&(t.children.length=0),r!==o[n]))if("value"===n){i._value=r;var s=null==r?"":String(r);Mt(i,t,s)&&(i.value=s)}else i[n]=r}}function Mt(e,t,n){return!e.composing&&("option"===t.tag||Dt(e,n)||Pt(t,n))}function Dt(e,t){return document.activeElement!==e&&e.value!==t}function Pt(e,n){var r=e.elm.value,i=e.elm._vModifiers;return i&&i.number||"number"===e.elm.type?t(r)!==t(n):i&&i.trim?r.trim()!==n.trim():r!==n}function Rt(e){var t=Ft(e.style);return e.staticStyle?u(e.staticStyle,t):t}function Ft(e){return Array.isArray(e)?p(e):"string"==typeof e?Po(e):e}function Ht(e,t){var n,r={};if(t)for(var i=e;i.componentInstance;)i=i.componentInstance._vnode,i.data&&(n=Rt(i.data))&&u(r,n);(n=Rt(e.data))&&u(r,n);for(var o=e;o=o.parent;)o.data&&(n=Rt(o.data))&&u(r,n);return r}function Ut(e,t){var n=t.data,r=e.data;if(n.staticStyle||n.style||r.staticStyle||r.style){var i,o,a=t.elm,s=e.data.staticStyle,c=e.data.style||{},l=s||c,f=Ft(t.data.style)||{};t.data.style=f.__ob__?u({},f):f;var p=Ht(t,!0);for(o in l)null==p[o]&&Ho(a,o,"");for(o in p)i=p[o],i!==l[o]&&Ho(a,o,null==i?"":i)}}function Bt(e,t){if(t&&t.trim())if(e.classList)t.indexOf(" ")>-1?t.split(/\s+/).forEach(function(t){return e.classList.add(t)}):e.classList.add(t);else{var n=" "+e.getAttribute("class")+" ";n.indexOf(" "+t+" ")<0&&e.setAttribute("class",(n+t).trim())}}function zt(e,t){if(t&&t.trim())if(e.classList)t.indexOf(" ")>-1?t.split(/\s+/).forEach(function(t){return e.classList.remove(t)}):e.classList.remove(t);else{for(var n=" "+e.getAttribute("class")+" ",r=" "+t+" ";n.indexOf(r)>=0;)n=n.replace(r," ");e.setAttribute("class",n.trim())}}function Vt(e){Yo(function(){Yo(e)})}function Jt(e,t){(e._transitionClasses||(e._transitionClasses=[])).push(t),Bt(e,t)}function Kt(e,t){e._transitionClasses&&r(e._transitionClasses,t),zt(e,t)}function qt(e,t,n){var r=Wt(e,t),i=r.type,o=r.timeout,a=r.propCount;if(!i)return n();var s=i===Jo?Wo:Go,c=0,u=function(){e.removeEventListener(s,l),n()},l=function(t){t.target===e&&++c>=a&&u()};setTimeout(function(){c<a&&u()},o+1),e.addEventListener(s,l)}function Wt(e,t){var n,r=window.getComputedStyle(e),i=r[qo+"Delay"].split(", "),o=r[qo+"Duration"].split(", "),a=Zt(i,o),s=r[Zo+"Delay"].split(", "),c=r[Zo+"Duration"].split(", "),u=Zt(s,c),l=0,f=0;t===Jo?a>0&&(n=Jo,l=a,f=o.length):t===Ko?u>0&&(n=Ko,l=u,f=c.length):(l=Math.max(a,u),n=l>0?a>u?Jo:Ko:null,f=n?n===Jo?o.length:c.length:0);var p=n===Jo&&Qo.test(r[qo+"Property"]);return{type:n,timeout:l,propCount:f,hasTransform:p}}function Zt(e,t){for(;e.length<t.length;)e=e.concat(e);return Math.max.apply(null,t.map(function(t,n){
return Gt(t)+Gt(e[n])}))}function Gt(e){return 1e3*Number(e.slice(0,-1))}function Yt(e,t){var n=e.elm;n._leaveCb&&(n._leaveCb.cancelled=!0,n._leaveCb());var r=Xt(e.data.transition);if(r&&!n._enterCb&&1===n.nodeType){for(var i=r.css,o=r.type,a=r.enterClass,s=r.enterToClass,c=r.enterActiveClass,u=r.appearClass,l=r.appearToClass,f=r.appearActiveClass,p=r.beforeEnter,d=r.enter,v=r.afterEnter,h=r.enterCancelled,m=r.beforeAppear,g=r.appear,y=r.afterAppear,_=r.appearCancelled,b=Zi,$=Zi.$vnode;$&&$.parent;)$=$.parent,b=$.context;var w=!b._isMounted||!e.isRootInsert;if(!w||g||""===g){var C=w?u:a,x=w?f:c,k=w?l:s,A=w?m||p:p,O=w&&"function"==typeof g?g:d,S=w?y||v:v,T=w?_||h:h,E=i!==!1&&!bi,I=O&&(O._length||O.length)>1,j=n._enterCb=en(function(){E&&(Kt(n,k),Kt(n,x)),j.cancelled?(E&&Kt(n,C),T&&T(n)):S&&S(n),n._enterCb=null});e.data.show||ne(e.data.hook||(e.data.hook={}),"insert",function(){var t=n.parentNode,r=t&&t._pending&&t._pending[e.key];r&&r.tag===e.tag&&r.elm._leaveCb&&r.elm._leaveCb(),O&&O(n,j)},"transition-insert"),A&&A(n),E&&(Jt(n,C),Jt(n,x),Vt(function(){Jt(n,k),Kt(n,C),j.cancelled||I||qt(n,o,j)})),e.data.show&&(t&&t(),O&&O(n,j)),E||I||j()}}}function Qt(e,t){function n(){g.cancelled||(e.data.show||((r.parentNode._pending||(r.parentNode._pending={}))[e.key]=e),l&&l(r),h&&(Jt(r,s),Jt(r,u),Vt(function(){Jt(r,c),Kt(r,s),g.cancelled||m||qt(r,a,g)})),f&&f(r,g),h||m||g())}var r=e.elm;r._enterCb&&(r._enterCb.cancelled=!0,r._enterCb());var i=Xt(e.data.transition);if(!i)return t();if(!r._leaveCb&&1===r.nodeType){var o=i.css,a=i.type,s=i.leaveClass,c=i.leaveToClass,u=i.leaveActiveClass,l=i.beforeLeave,f=i.leave,p=i.afterLeave,d=i.leaveCancelled,v=i.delayLeave,h=o!==!1&&!bi,m=f&&(f._length||f.length)>1,g=r._leaveCb=en(function(){r.parentNode&&r.parentNode._pending&&(r.parentNode._pending[e.key]=null),h&&(Kt(r,c),Kt(r,u)),g.cancelled?(h&&Kt(r,s),d&&d(r)):(t(),p&&p(r)),r._leaveCb=null});v?v(n):n()}}function Xt(e){if(e){if("object"==typeof e){var t={};return e.css!==!1&&u(t,Xo(e.name||"v")),u(t,e),t}return"string"==typeof e?Xo(e):void 0}}function en(e){var t=!1;return function(){t||(t=!0,e())}}function tn(e,t){t.data.show||Yt(t)}function nn(e,t,n){var r=t.value,i=e.multiple;if(!i||Array.isArray(r)){for(var o,a,s=0,c=e.options.length;s<c;s++)if(a=e.options[s],i)o=m(r,on(a))>-1,a.selected!==o&&(a.selected=o);else if(h(on(a),r))return void(e.selectedIndex!==s&&(e.selectedIndex=s));i||(e.selectedIndex=-1)}}function rn(e,t){for(var n=0,r=t.length;n<r;n++)if(h(on(t[n]),e))return!1;return!0}function on(e){return"_value"in e?e._value:e.value}function an(e){e.target.composing=!0}function sn(e){e.target.composing=!1,cn(e.target,"input")}function cn(e,t){var n=document.createEvent("HTMLEvents");n.initEvent(t,!0,!0),e.dispatchEvent(n)}function un(e){return!e.componentInstance||e.data&&e.data.transition?e:un(e.componentInstance._vnode)}function ln(e){var t=e&&e.componentOptions;return t&&t.Ctor.options.abstract?ln(ce(t.children)):e}function fn(e){var t={},n=e.$options;for(var r in n.propsData)t[r]=e[r];var i=n._parentListeners;for(var o in i)t[ai(o)]=i[o].fn;return t}function pn(e,t){return/\d-keep-alive$/.test(t.tag)?e("keep-alive"):null}function dn(e){for(;e=e.parent;)if(e.data.transition)return!0}function vn(e,t){return t.key===e.key&&t.tag===e.tag}function hn(e){e.elm._moveCb&&e.elm._moveCb(),e.elm._enterCb&&e.elm._enterCb()}function mn(e){e.data.newPos=e.elm.getBoundingClientRect()}function gn(e){var t=e.data.pos,n=e.data.newPos,r=t.left-n.left,i=t.top-n.top;if(r||i){e.data.moved=!0;var o=e.elm.style;o.transform=o.WebkitTransform="translate("+r+"px,"+i+"px)",o.transitionDuration="0s"}}function yn(e,t){var n=document.createElement("div");return n.innerHTML='<div a="'+e+'">',n.innerHTML.indexOf(t)>0}function _n(e){return pa=pa||document.createElement("div"),pa.innerHTML=e,pa.textContent}function bn(e,t){return t&&(e=e.replace(os,"\n")),e.replace(rs,"<").replace(is,">").replace(as,"&").replace(ss,'"')}function $n(e,t){function n(t){f+=t,e=e.substring(t)}function r(){var t=e.match(Ca);if(t){var r={tagName:t[1],attrs:[],start:f};n(t[0].length);for(var i,o;!(i=e.match(xa))&&(o=e.match(ba));)n(o[0].length),r.attrs.push(o);if(i)return r.unarySlash=i[1],n(i[0].length),r.end=f,r}}function i(e){var n=e.tagName,r=e.unarySlash;u&&("p"===s&&ma(n)&&o(s),ha(n)&&s===n&&o(n));for(var i=l(n)||"html"===n&&"head"===s||!!r,a=e.attrs.length,f=new Array(a),p=0;p<a;p++){var d=e.attrs[p];Ta&&d[0].indexOf('""')===-1&&(""===d[3]&&delete d[3],""===d[4]&&delete d[4],""===d[5]&&delete d[5]);var v=d[3]||d[4]||d[5]||"";f[p]={name:d[1],value:bn(v,t.shouldDecodeNewlines)}}i||(c.push({tag:n,lowerCasedTag:n.toLowerCase(),attrs:f}),s=n,r=""),t.start&&t.start(n,f,i,e.start,e.end)}function o(e,n,r){var i,o;if(null==n&&(n=f),null==r&&(r=f),e&&(o=e.toLowerCase()),e)for(i=c.length-1;i>=0&&c[i].lowerCasedTag!==o;i--);else i=0;if(i>=0){for(var a=c.length-1;a>=i;a--)t.end&&t.end(c[a].tag,n,r);c.length=i,s=i&&c[i-1].tag}else"br"===o?t.start&&t.start(e,[],!0,n,r):"p"===o&&(t.start&&t.start(e,[],!1,n,r),t.end&&t.end(e,n,r))}for(var a,s,c=[],u=t.expectHTML,l=t.isUnaryTag||pi,f=0;e;){if(a=e,s&&ts(s)){var p=s.toLowerCase(),d=ns[p]||(ns[p]=new RegExp("([\\s\\S]*?)(</"+p+"[^>]*>)","i")),v=0,h=e.replace(d,function(e,n,r){return v=r.length,"script"!==p&&"style"!==p&&"noscript"!==p&&(n=n.replace(/<!--([\s\S]*?)-->/g,"$1").replace(/<!\[CDATA\[([\s\S]*?)]]>/g,"$1")),t.chars&&t.chars(n),""});f+=e.length-h.length,e=h,o(p,f-v,f)}else{var m=e.indexOf("<");if(0===m){if(Oa.test(e)){var g=e.indexOf("-->");if(g>=0){n(g+3);continue}}if(Sa.test(e)){var y=e.indexOf("]>");if(y>=0){n(y+2);continue}}var _=e.match(Aa);if(_){n(_[0].length);continue}var b=e.match(ka);if(b){var $=f;n(b[0].length),o(b[1],$,f);continue}var w=r();if(w){i(w);continue}}var C=void 0,x=void 0,k=void 0;if(m>0){for(x=e.slice(m);!(ka.test(x)||Ca.test(x)||Oa.test(x)||Sa.test(x)||(k=x.indexOf("<",1),k<0));)m+=k,x=e.slice(m);C=e.substring(0,m),n(m)}m<0&&(C=e,e=""),t.chars&&C&&t.chars(C)}if(e===a&&t.chars){t.chars(e);break}}o()}function wn(e){function t(){(a||(a=[])).push(e.slice(v,i).trim()),v=i+1}var n,r,i,o,a,s=!1,c=!1,u=!1,l=!1,f=0,p=0,d=0,v=0;for(i=0;i<e.length;i++)if(r=n,n=e.charCodeAt(i),s)39===n&&92!==r&&(s=!1);else if(c)34===n&&92!==r&&(c=!1);else if(u)96===n&&92!==r&&(u=!1);else if(l)47===n&&92!==r&&(l=!1);else if(124!==n||124===e.charCodeAt(i+1)||124===e.charCodeAt(i-1)||f||p||d){switch(n){case 34:c=!0;break;case 39:s=!0;break;case 96:u=!0;break;case 40:d++;break;case 41:d--;break;case 91:p++;break;case 93:p--;break;case 123:f++;break;case 125:f--}if(47===n){for(var h=i-1,m=void 0;h>=0&&(m=e.charAt(h)," "===m);h--);m&&/[\w$]/.test(m)||(l=!0)}}else void 0===o?(v=i+1,o=e.slice(0,i).trim()):t();if(void 0===o?o=e.slice(0,i).trim():0!==v&&t(),a)for(i=0;i<a.length;i++)o=Cn(o,a[i]);return o}function Cn(e,t){var n=t.indexOf("(");if(n<0)return'_f("'+t+'")('+e+")";var r=t.slice(0,n),i=t.slice(n+1);return'_f("'+r+'")('+e+","+i}function xn(e,t){var n=t?ls(t):cs;if(n.test(e)){for(var r,i,o=[],a=n.lastIndex=0;r=n.exec(e);){i=r.index,i>a&&o.push(JSON.stringify(e.slice(a,i)));var s=wn(r[1].trim());o.push("_s("+s+")"),a=i+r[0].length}return a<e.length&&o.push(JSON.stringify(e.slice(a))),o.join("+")}}function kn(e){console.error("[Vue parser]: "+e)}function An(e,t){return e?e.map(function(e){return e[t]}).filter(function(e){return e}):[]}function On(e,t,n){(e.props||(e.props=[])).push({name:t,value:n})}function Sn(e,t,n){(e.attrs||(e.attrs=[])).push({name:t,value:n})}function Tn(e,t,n,r,i,o){(e.directives||(e.directives=[])).push({name:t,rawName:n,value:r,arg:i,modifiers:o})}function En(e,t,n,r,i){r&&r.capture&&(delete r.capture,t="!"+t),r&&r.once&&(delete r.once,t="~"+t);var o;r&&r.native?(delete r.native,o=e.nativeEvents||(e.nativeEvents={})):o=e.events||(e.events={});var a={value:n,modifiers:r},s=o[t];Array.isArray(s)?i?s.unshift(a):s.push(a):s?o[t]=i?[a,s]:[s,a]:o[t]=a}function In(e,t,n){var r=jn(e,":"+t)||jn(e,"v-bind:"+t);if(null!=r)return wn(r);if(n!==!1){var i=jn(e,t);if(null!=i)return JSON.stringify(i)}}function jn(e,t){var n;if(null!=(n=e.attrsMap[t]))for(var r=e.attrsList,i=0,o=r.length;i<o;i++)if(r[i].name===t){r.splice(i,1);break}return n}function Nn(e){if(Ia=e,Ea=Ia.length,Na=La=Ma=0,e.indexOf("[")<0||e.lastIndexOf("]")<Ea-1)return{exp:e,idx:null};for(;!Mn();)ja=Ln(),Dn(ja)?Rn(ja):91===ja&&Pn(ja);return{exp:e.substring(0,La),idx:e.substring(La+1,Ma)}}function Ln(){return Ia.charCodeAt(++Na)}function Mn(){return Na>=Ea}function Dn(e){return 34===e||39===e}function Pn(e){var t=1;for(La=Na;!Mn();)if(e=Ln(),Dn(e))Rn(e);else if(91===e&&t++,93===e&&t--,0===t){Ma=Na;break}}function Rn(e){for(var t=e;!Mn()&&(e=Ln(),e!==t););}function Fn(e,t){Da=t.warn||kn,Pa=t.getTagNamespace||pi,Ra=t.mustUseProp||pi,Fa=t.isPreTag||pi,Ha=An(t.modules,"preTransformNode"),Ua=An(t.modules,"transformNode"),Ba=An(t.modules,"postTransformNode"),za=t.delimiters;var n,r,i=[],o=t.preserveWhitespace!==!1,a=!1,s=!1;return $n(e,{expectHTML:t.expectHTML,isUnaryTag:t.isUnaryTag,shouldDecodeNewlines:t.shouldDecodeNewlines,start:function(e,o,c){function u(e){}var l=r&&r.ns||Pa(e);_i&&"svg"===l&&(o=rr(o));var f={type:1,tag:e,attrsList:o,attrsMap:tr(o),parent:r,children:[]};l&&(f.ns=l),nr(f)&&!xi()&&(f.forbidden=!0);for(var p=0;p<Ha.length;p++)Ha[p](f,t);if(a||(Hn(f),f.pre&&(a=!0)),Fa(f.tag)&&(s=!0),a)Un(f);else{Vn(f),Jn(f),Zn(f),Bn(f),f.plain=!f.key&&!o.length,zn(f),Gn(f),Yn(f);for(var d=0;d<Ua.length;d++)Ua[d](f,t);Qn(f)}if(n?i.length||n.if&&(f.elseif||f.else)&&(u(f),Wn(n,{exp:f.elseif,block:f})):(n=f,u(n)),r&&!f.forbidden)if(f.elseif||f.else)Kn(f,r);else if(f.slotScope){r.plain=!1;var v=f.slotTarget||"default";(r.scopedSlots||(r.scopedSlots={}))[v]=f}else r.children.push(f),f.parent=r;c||(r=f,i.push(f));for(var h=0;h<Ba.length;h++)Ba[h](f,t)},end:function(){var e=i[i.length-1],t=e.children[e.children.length-1];t&&3===t.type&&" "===t.text&&e.children.pop(),i.length-=1,r=i[i.length-1],e.pre&&(a=!1),Fa(e.tag)&&(s=!1)},chars:function(e){if(r&&(!_i||"textarea"!==r.tag||r.attrsMap.placeholder!==e)){var t=r.children;if(e=s||e.trim()?ys(e):o&&t.length?" ":""){var n;!a&&" "!==e&&(n=xn(e,za))?t.push({type:2,expression:n,text:e}):" "===e&&" "===t[t.length-1].text||r.children.push({type:3,text:e})}}}}),n}function Hn(e){null!=jn(e,"v-pre")&&(e.pre=!0)}function Un(e){var t=e.attrsList.length;if(t)for(var n=e.attrs=new Array(t),r=0;r<t;r++)n[r]={name:e.attrsList[r].name,value:JSON.stringify(e.attrsList[r].value)};else e.pre||(e.plain=!0)}function Bn(e){var t=In(e,"key");t&&(e.key=t)}function zn(e){var t=In(e,"ref");t&&(e.ref=t,e.refInFor=Xn(e))}function Vn(e){var t;if(t=jn(e,"v-for")){var n=t.match(ps);if(!n)return;e.for=n[2].trim();var r=n[1].trim(),i=r.match(ds);i?(e.alias=i[1].trim(),e.iterator1=i[2].trim(),i[3]&&(e.iterator2=i[3].trim())):e.alias=r}}function Jn(e){var t=jn(e,"v-if");if(t)e.if=t,Wn(e,{exp:t,block:e});else{null!=jn(e,"v-else")&&(e.else=!0);var n=jn(e,"v-else-if");n&&(e.elseif=n)}}function Kn(e,t){var n=qn(t.children);n&&n.if&&Wn(n,{exp:e.elseif,block:e})}function qn(e){for(var t=e.length;t--;){if(1===e[t].type)return e[t];e.pop()}}function Wn(e,t){e.ifConditions||(e.ifConditions=[]),e.ifConditions.push(t)}function Zn(e){var t=jn(e,"v-once");null!=t&&(e.once=!0)}function Gn(e){if("slot"===e.tag)e.slotName=In(e,"name");else{var t=In(e,"slot");t&&(e.slotTarget='""'===t?'"default"':t),"template"===e.tag&&(e.slotScope=jn(e,"scope"))}}function Yn(e){var t;(t=In(e,"is"))&&(e.component=t),null!=jn(e,"inline-template")&&(e.inlineTemplate=!0)}function Qn(e){var t,n,r,i,o,a,s,c,u=e.attrsList;for(t=0,n=u.length;t<n;t++)if(r=i=u[t].name,o=u[t].value,fs.test(r))if(e.hasBindings=!0,s=er(r),s&&(r=r.replace(gs,"")),vs.test(r))r=r.replace(vs,""),o=wn(o),c=!1,s&&(s.prop&&(c=!0,r=ai(r),"innerHtml"===r&&(r="innerHTML")),s.camel&&(r=ai(r))),c||Ra(e.tag,e.attrsMap.type,r)?On(e,r,o):Sn(e,r,o);else if(hs.test(r))r=r.replace(hs,""),En(e,r,o,s);else{r=r.replace(fs,"");var l=r.match(ms);l&&(a=l[1])&&(r=r.slice(0,-(a.length+1))),Tn(e,r,i,o,a,s)}else Sn(e,r,JSON.stringify(o))}function Xn(e){for(var t=e;t;){if(void 0!==t.for)return!0;t=t.parent}return!1}function er(e){var t=e.match(gs);if(t){var n={};return t.forEach(function(e){n[e.slice(1)]=!0}),n}}function tr(e){for(var t={},n=0,r=e.length;n<r;n++)t[e[n].name]=e[n].value;return t}function nr(e){return"style"===e.tag||"script"===e.tag&&(!e.attrsMap.type||"text/javascript"===e.attrsMap.type)}function rr(e){for(var t=[],n=0;n<e.length;n++){var r=e[n];_s.test(r.name)||(r.name=r.name.replace(bs,""),t.push(r))}return t}function ir(e,t){e&&(Va=$s(t.staticKeys||""),Ja=t.isReservedTag||pi,ar(e),sr(e,!1))}function or(e){return n("type,tag,attrsList,attrsMap,plain,parent,children,attrs"+(e?","+e:""))}function ar(e){if(e.static=ur(e),1===e.type){if(!Ja(e.tag)&&"slot"!==e.tag&&null==e.attrsMap["inline-template"])return;for(var t=0,n=e.children.length;t<n;t++){var r=e.children[t];ar(r),r.static||(e.static=!1)}}}function sr(e,t){if(1===e.type){if((e.static||e.once)&&(e.staticInFor=t),e.static&&e.children.length&&(1!==e.children.length||3!==e.children[0].type))return void(e.staticRoot=!0);if(e.staticRoot=!1,e.children)for(var n=0,r=e.children.length;n<r;n++)sr(e.children[n],t||!!e.for);e.ifConditions&&cr(e.ifConditions,t)}}function cr(e,t){for(var n=1,r=e.length;n<r;n++)sr(e[n].block,t)}function ur(e){return 2!==e.type&&(3===e.type||!(!e.pre&&(e.hasBindings||e.if||e.for||ri(e.tag)||!Ja(e.tag)||lr(e)||!Object.keys(e).every(Va))))}function lr(e){for(;e.parent;){if(e=e.parent,"template"!==e.tag)return!1;if(e.for)return!0}return!1}function fr(e,t){var n=t?"nativeOn:{":"on:{";for(var r in e)n+='"'+r+'":'+pr(r,e[r])+",";return n.slice(0,-1)+"}"}function pr(e,t){if(t){if(Array.isArray(t))return"["+t.map(function(t){return pr(e,t)}).join(",")+"]";if(t.modifiers){var n="",r=[];for(var i in t.modifiers)ks[i]?n+=ks[i]:r.push(i);r.length&&(n=dr(r)+n);var o=Cs.test(t.value)?t.value+"($event)":t.value;return"function($event){"+n+o+"}"}return ws.test(t.value)||Cs.test(t.value)?t.value:"function($event){"+t.value+"}"}return"function(){}"}function dr(e){return"if("+e.map(vr).join("&&")+")return;"}function vr(e){var t=parseInt(e,10);if(t)return"$event.keyCode!=="+t;var n=xs[e];return"_k($event.keyCode,"+JSON.stringify(e)+(n?","+JSON.stringify(n):"")+")"}function hr(e,t){e.wrapData=function(n){return"_b("+n+",'"+e.tag+"',"+t.value+(t.modifiers&&t.modifiers.prop?",true":"")+")"}}function mr(e,t){var n=Ya,r=Ya=[],i=Qa;Qa=0,Xa=t,Ka=t.warn||kn,qa=An(t.modules,"transformCode"),Wa=An(t.modules,"genData"),Za=t.directives||{},Ga=t.isReservedTag||pi;var o=e?gr(e):'_c("div")';return Ya=n,Qa=i,{render:"with(this){return "+o+"}",staticRenderFns:r}}function gr(e){if(e.staticRoot&&!e.staticProcessed)return yr(e);if(e.once&&!e.onceProcessed)return _r(e);if(e.for&&!e.forProcessed)return wr(e);if(e.if&&!e.ifProcessed)return br(e);if("template"!==e.tag||e.slotTarget){if("slot"===e.tag)return Lr(e);var t;if(e.component)t=Mr(e.component,e);else{var n=e.plain?void 0:Cr(e),r=e.inlineTemplate?null:Sr(e,!0);t="_c('"+e.tag+"'"+(n?","+n:"")+(r?","+r:"")+")"}for(var i=0;i<qa.length;i++)t=qa[i](e,t);return t}return Sr(e)||"void 0"}function yr(e){return e.staticProcessed=!0,Ya.push("with(this){return "+gr(e)+"}"),"_m("+(Ya.length-1)+(e.staticInFor?",true":"")+")"}function _r(e){if(e.onceProcessed=!0,e.if&&!e.ifProcessed)return br(e);if(e.staticInFor){for(var t="",n=e.parent;n;){if(n.for){t=n.key;break}n=n.parent}return t?"_o("+gr(e)+","+Qa++ +(t?","+t:"")+")":gr(e)}return yr(e)}function br(e){return e.ifProcessed=!0,$r(e.ifConditions.slice())}function $r(e){function t(e){return e.once?_r(e):gr(e)}if(!e.length)return"_e()";var n=e.shift();return n.exp?"("+n.exp+")?"+t(n.block)+":"+$r(e):""+t(n.block)}function wr(e){var t=e.for,n=e.alias,r=e.iterator1?","+e.iterator1:"",i=e.iterator2?","+e.iterator2:"";return e.forProcessed=!0,"_l(("+t+"),function("+n+r+i+"){return "+gr(e)+"})"}function Cr(e){var t="{",n=xr(e);n&&(t+=n+","),e.key&&(t+="key:"+e.key+","),e.ref&&(t+="ref:"+e.ref+","),e.refInFor&&(t+="refInFor:true,"),e.pre&&(t+="pre:true,"),e.component&&(t+='tag:"'+e.tag+'",');for(var r=0;r<Wa.length;r++)t+=Wa[r](e);if(e.attrs&&(t+="attrs:{"+Dr(e.attrs)+"},"),e.props&&(t+="domProps:{"+Dr(e.props)+"},"),e.events&&(t+=fr(e.events)+","),e.nativeEvents&&(t+=fr(e.nativeEvents,!0)+","),e.slotTarget&&(t+="slot:"+e.slotTarget+","),e.scopedSlots&&(t+=Ar(e.scopedSlots)+","),e.inlineTemplate){var i=kr(e);i&&(t+=i+",")}return t=t.replace(/,$/,"")+"}",e.wrapData&&(t=e.wrapData(t)),t}function xr(e){var t=e.directives;if(t){var n,r,i,o,a="directives:[",s=!1;for(n=0,r=t.length;n<r;n++){i=t[n],o=!0;var c=Za[i.name]||As[i.name];c&&(o=!!c(e,i,Ka)),o&&(s=!0,a+='{name:"'+i.name+'",rawName:"'+i.rawName+'"'+(i.value?",value:("+i.value+"),expression:"+JSON.stringify(i.value):"")+(i.arg?',arg:"'+i.arg+'"':"")+(i.modifiers?",modifiers:"+JSON.stringify(i.modifiers):"")+"},")}return s?a.slice(0,-1)+"]":void 0}}function kr(e){var t=e.children[0];if(1===t.type){var n=mr(t,Xa);return"inlineTemplate:{render:function(){"+n.render+"},staticRenderFns:["+n.staticRenderFns.map(function(e){return"function(){"+e+"}"}).join(",")+"]}"}}function Ar(e){return"scopedSlots:{"+Object.keys(e).map(function(t){return Or(t,e[t])}).join(",")+"}"}function Or(e,t){return e+":function("+String(t.attrsMap.scope)+"){return "+("template"===t.tag?Sr(t)||"void 0":gr(t))+"}"}function Sr(e,t){var n=e.children;if(n.length){var r=n[0];if(1===n.length&&r.for&&"template"!==r.tag&&"slot"!==r.tag)return gr(r);var i=Tr(n);return"["+n.map(jr).join(",")+"]"+(t&&i?","+i:"")}}function Tr(e){for(var t=0,n=0;n<e.length;n++){var r=e[n];if(1===r.type){if(Er(r)||r.ifConditions&&r.ifConditions.some(function(e){return Er(e.block)})){t=2;break}(Ir(r)||r.ifConditions&&r.ifConditions.some(function(e){return Ir(e.block)}))&&(t=1)}}return t}function Er(e){return void 0!==e.for||"template"===e.tag||"slot"===e.tag}function Ir(e){return!Ga(e.tag)}function jr(e){return 1===e.type?gr(e):Nr(e)}function Nr(e){return"_v("+(2===e.type?e.expression:Pr(JSON.stringify(e.text)))+")"}function Lr(e){var t=e.slotName||'"default"',n=Sr(e),r="_t("+t+(n?","+n:""),i=e.attrs&&"{"+e.attrs.map(function(e){return ai(e.name)+":"+e.value}).join(",")+"}",o=e.attrsMap["v-bind"];return!i&&!o||n||(r+=",null"),i&&(r+=","+i),o&&(r+=(i?"":",null")+","+o),r+")"}function Mr(e,t){var n=t.inlineTemplate?null:Sr(t,!0);return"_c("+e+","+Cr(t)+(n?","+n:"")+")"}function Dr(e){for(var t="",n=0;n<e.length;n++){var r=e[n];t+='"'+r.name+'":'+Pr(r.value)+","}return t.slice(0,-1)}function Pr(e){return e.replace(/\u2028/g,"\\u2028").replace(/\u2029/g,"\\u2029")}function Rr(e,t){var n=Fn(e.trim(),t);ir(n,t);var r=mr(n,t);return{ast:n,render:r.render,staticRenderFns:r.staticRenderFns}}function Fr(e,t){var n=(t.warn||kn,jn(e,"class"));n&&(e.staticClass=JSON.stringify(n));var r=In(e,"class",!1);r&&(e.classBinding=r)}function Hr(e){var t="";return e.staticClass&&(t+="staticClass:"+e.staticClass+","),e.classBinding&&(t+="class:"+e.classBinding+","),t}function Ur(e,t){var n=(t.warn||kn,jn(e,"style"));n&&(e.staticStyle=JSON.stringify(Po(n)));var r=In(e,"style",!1);r&&(e.styleBinding=r)}function Br(e){var t="";return e.staticStyle&&(t+="staticStyle:"+e.staticStyle+","),e.styleBinding&&(t+="style:("+e.styleBinding+"),"),t}function zr(e,t,n){es=n;var r=t.value,i=t.modifiers,o=e.tag,a=e.attrsMap.type;return"select"===o?qr(e,r,i):"input"===o&&"checkbox"===a?Vr(e,r,i):"input"===o&&"radio"===a?Jr(e,r,i):Kr(e,r,i),!0}function Vr(e,t,n){var r=n&&n.number,i=In(e,"value")||"null",o=In(e,"true-value")||"true",a=In(e,"false-value")||"false";On(e,"checked","Array.isArray("+t+")?_i("+t+","+i+")>-1"+("true"===o?":("+t+")":":_q("+t+","+o+")")),En(e,"click","var $$a="+t+",$$el=$event.target,$$c=$$el.checked?("+o+"):("+a+");if(Array.isArray($$a)){var $$v="+(r?"_n("+i+")":i)+",$$i=_i($$a,$$v);if($$c){$$i<0&&("+t+"=$$a.concat($$v))}else{$$i>-1&&("+t+"=$$a.slice(0,$$i).concat($$a.slice($$i+1)))}}else{"+t+"=$$c}",null,!0)}function Jr(e,t,n){var r=n&&n.number,i=In(e,"value")||"null";i=r?"_n("+i+")":i,On(e,"checked","_q("+t+","+i+")"),En(e,"click",Wr(t,i),null,!0)}function Kr(e,t,n){var r=e.attrsMap.type,i=n||{},o=i.lazy,a=i.number,s=i.trim,c=o||_i&&"range"===r?"change":"input",u=!o&&"range"!==r,l="input"===e.tag||"textarea"===e.tag,f=l?"$event.target.value"+(s?".trim()":""):s?"(typeof $event === 'string' ? $event.trim() : $event)":"$event";f=a||"number"===r?"_n("+f+")":f;var p=Wr(t,f);l&&u&&(p="if($event.target.composing)return;"+p),On(e,"value",l?"_s("+t+")":"("+t+")"),En(e,c,p,null,!0),(s||a||"number"===r)&&En(e,"blur","$forceUpdate()")}function qr(e,t,n){var r=n&&n.number,i='Array.prototype.filter.call($event.target.options,function(o){return o.selected}).map(function(o){var val = "_value" in o ? o._value : o.value;return '+(r?"_n(val)":"val")+"})"+(null==e.attrsMap.multiple?"[0]":""),o=Wr(t,i);En(e,"change",o,null,!0)}function Wr(e,t){var n=Nn(e);return null===n.idx?e+"="+t:"var $$exp = "+n.exp+", $$idx = "+n.idx+";if (!Array.isArray($$exp)){"+e+"="+t+"}else{$$exp.splice($$idx, 1, "+t+")}"}function Zr(e,t){t.value&&On(e,"textContent","_s("+t.value+")")}function Gr(e,t){t.value&&On(e,"innerHTML","_s("+t.value+")")}function Yr(e,t){return t=t?u(u({},js),t):js,Rr(e,t)}function Qr(e,t,n){var r=(t&&t.warn||Si,t&&t.delimiters?String(t.delimiters)+e:e);if(Is[r])return Is[r];var i={},o=Yr(e,t);i.render=Xr(o.render);var a=o.staticRenderFns.length;i.staticRenderFns=new Array(a);for(var s=0;s<a;s++)i.staticRenderFns[s]=Xr(o.staticRenderFns[s]);return Is[r]=i}function Xr(e){try{return new Function(e)}catch(e){return d}}function ei(e){if(e.outerHTML)return e.outerHTML;var t=document.createElement("div");return t.appendChild(e.cloneNode(!0)),t.innerHTML}var ti,ni,ri=n("slot,component",!0),ii=Object.prototype.hasOwnProperty,oi=/-(\w)/g,ai=a(function(e){return e.replace(oi,function(e,t){return t?t.toUpperCase():""})}),si=a(function(e){return e.charAt(0).toUpperCase()+e.slice(1)}),ci=/([^-])([A-Z])/g,ui=a(function(e){return e.replace(ci,"$1-$2").replace(ci,"$1-$2").toLowerCase()}),li=Object.prototype.toString,fi="[object Object]",pi=function(){return!1},di=function(e){return e},vi={optionMergeStrategies:Object.create(null),silent:!1,devtools:!1,errorHandler:null,ignoredElements:[],keyCodes:Object.create(null),isReservedTag:pi,isUnknownElement:pi,getTagNamespace:d,parsePlatformTagName:di,mustUseProp:pi,_assetTypes:["component","directive","filter"],_lifecycleHooks:["beforeCreate","created","beforeMount","mounted","beforeUpdate","updated","beforeDestroy","destroyed","activated","deactivated"],_maxUpdateCount:100},hi=/[^\w.$]/,mi="__proto__"in{},gi="undefined"!=typeof window,yi=gi&&window.navigator.userAgent.toLowerCase(),_i=yi&&/msie|trident/.test(yi),bi=yi&&yi.indexOf("msie 9.0")>0,$i=yi&&yi.indexOf("edge/")>0,wi=yi&&yi.indexOf("android")>0,Ci=yi&&/iphone|ipad|ipod|ios/.test(yi),xi=function(){return void 0===ti&&(ti=!gi&&"undefined"!=typeof global&&"server"===global.process.env.VUE_ENV),ti},ki=gi&&window.__VUE_DEVTOOLS_GLOBAL_HOOK__,Ai=function(){function e(){r=!1;var e=n.slice(0);n.length=0;for(var t=0;t<e.length;t++)e[t]()}var t,n=[],r=!1;if("undefined"!=typeof Promise&&b(Promise)){var i=Promise.resolve(),o=function(e){console.error(e)};t=function(){i.then(e).catch(o),Ci&&setTimeout(d)}}else if("undefined"==typeof MutationObserver||!b(MutationObserver)&&"[object MutationObserverConstructor]"!==MutationObserver.toString())t=function(){setTimeout(e,0)};else{var a=1,s=new MutationObserver(e),c=document.createTextNode(String(a));s.observe(c,{characterData:!0}),t=function(){a=(a+1)%2,c.data=String(a)}}return function(e,i){var o;if(n.push(function(){e&&e.call(i),o&&o(i)}),r||(r=!0,t()),!e&&"undefined"!=typeof Promise)return new Promise(function(e){o=e})}}();ni="undefined"!=typeof Set&&b(Set)?Set:function(){function e(){this.set=Object.create(null)}return e.prototype.has=function(e){return this.set[e]===!0},e.prototype.add=function(e){this.set[e]=!0},e.prototype.clear=function(){this.set=Object.create(null)},e}();var Oi,Si=d,Ti=0,Ei=function(){this.id=Ti++,this.subs=[]};Ei.prototype.addSub=function(e){this.subs.push(e)},Ei.prototype.removeSub=function(e){r(this.subs,e)},Ei.prototype.depend=function(){Ei.target&&Ei.target.addDep(this)},Ei.prototype.notify=function(){for(var e=this.subs.slice(),t=0,n=e.length;t<n;t++)e[t].update()},Ei.target=null;var Ii=[],ji=Array.prototype,Ni=Object.create(ji);["push","pop","shift","unshift","splice","sort","reverse"].forEach(function(e){var t=ji[e];y(Ni,e,function(){for(var n=arguments,r=arguments.length,i=new Array(r);r--;)i[r]=n[r];var o,a=t.apply(this,i),s=this.__ob__;switch(e){case"push":o=i;break;case"unshift":o=i;break;case"splice":o=i.slice(2)}return o&&s.observeArray(o),s.dep.notify(),a})});var Li=Object.getOwnPropertyNames(Ni),Mi={shouldConvert:!0,isSettingProps:!1},Di=function(e){if(this.value=e,this.dep=new Ei,this.vmCount=0,y(e,"__ob__",this),Array.isArray(e)){var t=mi?C:x;t(e,Ni,Li),this.observeArray(e)}else this.walk(e)};Di.prototype.walk=function(e){for(var t=Object.keys(e),n=0;n<t.length;n++)A(e,t[n],e[t[n]])},Di.prototype.observeArray=function(e){for(var t=0,n=e.length;t<n;t++)k(e[t])};var Pi=vi.optionMergeStrategies;Pi.data=function(e,t,n){return n?e||t?function(){var r="function"==typeof t?t.call(n):t,i="function"==typeof e?e.call(n):void 0;return r?E(r,i):i}:void 0:t?"function"!=typeof t?e:e?function(){return E(t.call(this),e.call(this))}:t:e},vi._lifecycleHooks.forEach(function(e){Pi[e]=I}),vi._assetTypes.forEach(function(e){Pi[e+"s"]=j}),Pi.watch=function(e,t){if(!t)return e;if(!e)return t;var n={};u(n,e);for(var r in t){var i=n[r],o=t[r];i&&!Array.isArray(i)&&(i=[i]),n[r]=i?i.concat(o):[o]}return n},Pi.props=Pi.methods=Pi.computed=function(e,t){if(!t)return e;if(!e)return t;var n=Object.create(null);return u(n,e),u(n,t),n};var Ri=function(e,t){return void 0===t?e:t},Fi=Object.freeze({defineReactive:A,_toString:e,toNumber:t,makeMap:n,isBuiltInTag:ri,remove:r,hasOwn:i,isPrimitive:o,cached:a,camelize:ai,capitalize:si,hyphenate:ui,bind:s,toArray:c,extend:u,isObject:l,isPlainObject:f,toObject:p,noop:d,no:pi,identity:di,genStaticKeys:v,looseEqual:h,looseIndexOf:m,isReserved:g,def:y,parsePath:_,hasProto:mi,inBrowser:gi,UA:yi,isIE:_i,isIE9:bi,isEdge:$i,isAndroid:wi,isIOS:Ci,isServerRendering:xi,devtools:ki,nextTick:Ai,get _Set(){return ni},mergeOptions:M,resolveAsset:D,warn:Si,formatComponentName:Oi,validateProp:P}),Hi=function(e,t,n,r,i,o,a){this.tag=e,this.data=t,this.children=n,this.text=r,this.elm=i,this.ns=void 0,this.context=o,this.functionalContext=void 0,this.key=t&&t.key,this.componentOptions=a,this.componentInstance=void 0,this.parent=void 0,this.raw=!1,this.isStatic=!1,this.isRootInsert=!0,this.isComment=!1,this.isCloned=!1,this.isOnce=!1},Ui={child:{}};Ui.child.get=function(){return this.componentInstance},Object.defineProperties(Hi.prototype,Ui);var Bi,zi=function(){var e=new Hi;return e.text="",e.isComment=!0,e},Vi={init:q,prepatch:W,insert:Z,destroy:G},Ji=Object.keys(Vi),Ki=a(function(e){var t="~"===e.charAt(0);e=t?e.slice(1):e;var n="!"===e.charAt(0);return e=n?e.slice(1):e,{name:e,once:t,capture:n}}),qi=1,Wi=2,Zi=null,Gi=[],Yi={},Qi=!1,Xi=!1,eo=0,to=0,no=function(e,t,n,r){this.vm=e,e._watchers.push(this),r?(this.deep=!!r.deep,this.user=!!r.user,this.lazy=!!r.lazy,this.sync=!!r.sync):this.deep=this.user=this.lazy=this.sync=!1,this.cb=n,this.id=++to,this.active=!0,this.dirty=this.lazy,this.deps=[],this.newDeps=[],this.depIds=new ni,this.newDepIds=new ni,this.expression="","function"==typeof t?this.getter=t:(this.getter=_(t),this.getter||(this.getter=function(){})),this.value=this.lazy?void 0:this.get()};no.prototype.get=function(){$(this);var e=this.getter.call(this.vm,this.vm);return this.deep&&Ae(e),w(),this.cleanupDeps(),e},no.prototype.addDep=function(e){var t=e.id;this.newDepIds.has(t)||(this.newDepIds.add(t),this.newDeps.push(e),this.depIds.has(t)||e.addSub(this))},no.prototype.cleanupDeps=function(){for(var e=this,t=this.deps.length;t--;){var n=e.deps[t];e.newDepIds.has(n.id)||n.removeSub(e)}var r=this.depIds;this.depIds=this.newDepIds,this.newDepIds=r,this.newDepIds.clear(),r=this.deps,this.deps=this.newDeps,this.newDeps=r,this.newDeps.length=0},no.prototype.update=function(){this.lazy?this.dirty=!0:this.sync?this.run():ke(this)},no.prototype.run=function(){if(this.active){var e=this.get();if(e!==this.value||l(e)||this.deep){var t=this.value;if(this.value=e,this.user)try{this.cb.call(this.vm,e,t)}catch(e){if(!vi.errorHandler)throw e;vi.errorHandler.call(null,e,this.vm)}else this.cb.call(this.vm,e,t)}}},no.prototype.evaluate=function(){this.value=this.get(),this.dirty=!1},no.prototype.depend=function(){for(var e=this,t=this.deps.length;t--;)e.deps[t].depend()},no.prototype.teardown=function(){var e=this;if(this.active){this.vm._isBeingDestroyed||r(this.vm._watchers,this);for(var t=this.deps.length;t--;)e.deps[t].removeSub(e);this.active=!1}};var ro=new ni,io={enumerable:!0,configurable:!0,get:d,set:d},oo=0;Re(Ue),De(Ue),_e(Ue),$e(Ue),de(Ue);var ao=[String,RegExp],so={name:"keep-alive",abstract:!0,props:{include:ao,exclude:ao},created:function(){this.cache=Object.create(null)},destroyed:function(){var e=this;for(var t in this.cache)Ze(e.cache[t])},watch:{include:function(e){We(this.cache,function(t){return qe(e,t)})},exclude:function(e){We(this.cache,function(t){return!qe(e,t)})}},render:function(){var e=ce(this.$slots.default),t=e&&e.componentOptions;if(t){var n=Ke(t);if(n&&(this.include&&!qe(this.include,n)||this.exclude&&qe(this.exclude,n)))return e;var r=null==e.key?t.Ctor.cid+(t.tag?"::"+t.tag:""):e.key;this.cache[r]?e.componentInstance=this.cache[r].componentInstance:this.cache[r]=e,e.data.keepAlive=!0}return e}},co={KeepAlive:so};Ge(Ue),Object.defineProperty(Ue.prototype,"$isServer",{get:xi}),Ue.version="2.1.10";var uo,lo,fo=n("input,textarea,option,select"),po=function(e,t,n){return"value"===n&&fo(e)&&"button"!==t||"selected"===n&&"option"===e||"checked"===n&&"input"===e||"muted"===n&&"video"===e},vo=n("contenteditable,draggable,spellcheck"),ho=n("allowfullscreen,async,autofocus,autoplay,checked,compact,controls,declare,default,defaultchecked,defaultmuted,defaultselected,defer,disabled,enabled,formnovalidate,hidden,indeterminate,inert,ismap,itemscope,loop,multiple,muted,nohref,noresize,noshade,novalidate,nowrap,open,pauseonexit,readonly,required,reversed,scoped,seamless,selected,sortable,translate,truespeed,typemustmatch,visible"),mo="http://www.w3.org/1999/xlink",go=function(e){return":"===e.charAt(5)&&"xlink"===e.slice(0,5)},yo=function(e){return go(e)?e.slice(6,e.length):""},_o=function(e){return null==e||e===!1},bo={svg:"http://www.w3.org/2000/svg",math:"http://www.w3.org/1998/Math/MathML"},$o=n("html,body,base,head,link,meta,style,title,address,article,aside,footer,header,h1,h2,h3,h4,h5,h6,hgroup,nav,section,div,dd,dl,dt,figcaption,figure,hr,img,li,main,ol,p,pre,ul,a,b,abbr,bdi,bdo,br,cite,code,data,dfn,em,i,kbd,mark,q,rp,rt,rtc,ruby,s,samp,small,span,strong,sub,sup,time,u,var,wbr,area,audio,map,track,video,embed,object,param,source,canvas,script,noscript,del,ins,caption,col,colgroup,table,thead,tbody,td,th,tr,button,datalist,fieldset,form,input,label,legend,meter,optgroup,option,output,progress,select,textarea,details,dialog,menu,menuitem,summary,content,element,shadow,template"),wo=n("svg,animate,circle,clippath,cursor,defs,desc,ellipse,filter,font-face,g,glyph,image,line,marker,mask,missing-glyph,path,pattern,polygon,polyline,rect,switch,symbol,text,textpath,tspan,use,view",!0),Co=function(e){return"pre"===e},xo=function(e){return $o(e)||wo(e)},ko=Object.create(null),Ao=Object.freeze({createElement:ot,createElementNS:at,createTextNode:st,createComment:ct,insertBefore:ut,removeChild:lt,appendChild:ft,parentNode:pt,nextSibling:dt,tagName:vt,setTextContent:ht,setAttribute:mt}),Oo={create:function(e,t){gt(t)},update:function(e,t){e.data.ref!==t.data.ref&&(gt(e,!0),gt(t))},destroy:function(e){gt(e,!0)}},So=new Hi("",{},[]),To=["create","activate","update","remove","destroy"],Eo={create:Ct,
update:Ct,destroy:function(e){Ct(e,So)}},Io=Object.create(null),jo=[Oo,Eo],No={create:St,update:St},Lo={create:Et,update:Et},Mo={create:Nt,update:Nt},Do={create:Lt,update:Lt},Po=a(function(e){var t={},n=/;(?![^(]*\))/g,r=/:(.+)/;return e.split(n).forEach(function(e){if(e){var n=e.split(r);n.length>1&&(t[n[0].trim()]=n[1].trim())}}),t}),Ro=/^--/,Fo=/\s*!important$/,Ho=function(e,t,n){Ro.test(t)?e.style.setProperty(t,n):Fo.test(n)?e.style.setProperty(t,n.replace(Fo,""),"important"):e.style[Bo(t)]=n},Uo=["Webkit","Moz","ms"],Bo=a(function(e){if(lo=lo||document.createElement("div"),e=ai(e),"filter"!==e&&e in lo.style)return e;for(var t=e.charAt(0).toUpperCase()+e.slice(1),n=0;n<Uo.length;n++){var r=Uo[n]+t;if(r in lo.style)return r}}),zo={create:Ut,update:Ut},Vo=gi&&!bi,Jo="transition",Ko="animation",qo="transition",Wo="transitionend",Zo="animation",Go="animationend";Vo&&(void 0===window.ontransitionend&&void 0!==window.onwebkittransitionend&&(qo="WebkitTransition",Wo="webkitTransitionEnd"),void 0===window.onanimationend&&void 0!==window.onwebkitanimationend&&(Zo="WebkitAnimation",Go="webkitAnimationEnd"));var Yo=gi&&window.requestAnimationFrame?window.requestAnimationFrame.bind(window):setTimeout,Qo=/\b(transform|all)(,|$)/,Xo=a(function(e){return{enterClass:e+"-enter",leaveClass:e+"-leave",appearClass:e+"-enter",enterToClass:e+"-enter-to",leaveToClass:e+"-leave-to",appearToClass:e+"-enter-to",enterActiveClass:e+"-enter-active",leaveActiveClass:e+"-leave-active",appearActiveClass:e+"-enter-active"}}),ea=gi?{create:tn,activate:tn,remove:function(e,t){e.data.show?t():Qt(e,t)}}:{},ta=[No,Lo,Mo,Do,zo,ea],na=ta.concat(jo),ra=wt({nodeOps:Ao,modules:na});bi&&document.addEventListener("selectionchange",function(){var e=document.activeElement;e&&e.vmodel&&cn(e,"input")});var ia={inserted:function(e,t,n){if("select"===n.tag){var r=function(){nn(e,t,n.context)};r(),(_i||$i)&&setTimeout(r,0)}else"textarea"!==n.tag&&"text"!==e.type||(e._vModifiers=t.modifiers,t.modifiers.lazy||(wi||(e.addEventListener("compositionstart",an),e.addEventListener("compositionend",sn)),bi&&(e.vmodel=!0)))},componentUpdated:function(e,t,n){if("select"===n.tag){nn(e,t,n.context);var r=e.multiple?t.value.some(function(t){return rn(t,e.options)}):t.value!==t.oldValue&&rn(t.value,e.options);r&&cn(e,"change")}}},oa={bind:function(e,t,n){var r=t.value;n=un(n);var i=n.data&&n.data.transition,o=e.__vOriginalDisplay="none"===e.style.display?"":e.style.display;r&&i&&!bi?(n.data.show=!0,Yt(n,function(){e.style.display=o})):e.style.display=r?o:"none"},update:function(e,t,n){var r=t.value,i=t.oldValue;if(r!==i){n=un(n);var o=n.data&&n.data.transition;o&&!bi?(n.data.show=!0,r?Yt(n,function(){e.style.display=e.__vOriginalDisplay}):Qt(n,function(){e.style.display="none"})):e.style.display=r?e.__vOriginalDisplay:"none"}},unbind:function(e,t,n,r,i){i||(e.style.display=e.__vOriginalDisplay)}},aa={model:ia,show:oa},sa={name:String,appear:Boolean,css:Boolean,mode:String,type:String,enterClass:String,leaveClass:String,enterToClass:String,leaveToClass:String,enterActiveClass:String,leaveActiveClass:String,appearClass:String,appearActiveClass:String,appearToClass:String},ca={name:"transition",props:sa,abstract:!0,render:function(e){var t=this,n=this.$slots.default;if(n&&(n=n.filter(function(e){return e.tag}),n.length)){var r=this.mode,i=n[0];if(dn(this.$vnode))return i;var a=ln(i);if(!a)return i;if(this._leaving)return pn(e,i);var s="__transition-"+this._uid+"-",c=a.key=null==a.key?s+a.tag:o(a.key)?0===String(a.key).indexOf(s)?a.key:s+a.key:a.key,l=(a.data||(a.data={})).transition=fn(this),f=this._vnode,p=ln(f);if(a.data.directives&&a.data.directives.some(function(e){return"show"===e.name})&&(a.data.show=!0),p&&p.data&&!vn(a,p)){var d=p&&(p.data.transition=u({},l));if("out-in"===r)return this._leaving=!0,ne(d,"afterLeave",function(){t._leaving=!1,t.$forceUpdate()},c),pn(e,i);if("in-out"===r){var v,h=function(){v()};ne(l,"afterEnter",h,c),ne(l,"enterCancelled",h,c),ne(d,"delayLeave",function(e){v=e},c)}}return i}}},ua=u({tag:String,moveClass:String},sa);delete ua.mode;var la={props:ua,render:function(e){for(var t=this.tag||this.$vnode.data.tag||"span",n=Object.create(null),r=this.prevChildren=this.children,i=this.$slots.default||[],o=this.children=[],a=fn(this),s=0;s<i.length;s++){var c=i[s];c.tag&&null!=c.key&&0!==String(c.key).indexOf("__vlist")&&(o.push(c),n[c.key]=c,(c.data||(c.data={})).transition=a)}if(r){for(var u=[],l=[],f=0;f<r.length;f++){var p=r[f];p.data.transition=a,p.data.pos=p.elm.getBoundingClientRect(),n[p.key]?u.push(p):l.push(p)}this.kept=e(t,null,u),this.removed=l}return e(t,null,o)},beforeUpdate:function(){this.__patch__(this._vnode,this.kept,!1,!0),this._vnode=this.kept},updated:function(){var e=this.prevChildren,t=this.moveClass||(this.name||"v")+"-move";if(e.length&&this.hasMove(e[0].elm,t)){e.forEach(hn),e.forEach(mn),e.forEach(gn);document.body.offsetHeight;e.forEach(function(e){if(e.data.moved){var n=e.elm,r=n.style;Jt(n,t),r.transform=r.WebkitTransform=r.transitionDuration="",n.addEventListener(Wo,n._moveCb=function e(r){r&&!/transform$/.test(r.propertyName)||(n.removeEventListener(Wo,e),n._moveCb=null,Kt(n,t))})}})}},methods:{hasMove:function(e,t){if(!Vo)return!1;if(null!=this._hasMove)return this._hasMove;Jt(e,t);var n=Wt(e);return Kt(e,t),this._hasMove=n.hasTransform}}},fa={Transition:ca,TransitionGroup:la};Ue.config.isUnknownElement=rt,Ue.config.isReservedTag=xo,Ue.config.getTagNamespace=nt,Ue.config.mustUseProp=po,u(Ue.options.directives,aa),u(Ue.options.components,fa),Ue.prototype.__patch__=gi?ra:d,Ue.prototype.$mount=function(e,t){return e=e&&gi?it(e):void 0,this._mount(e,t)},setTimeout(function(){vi.devtools&&ki&&ki.emit("init",Ue)},0);var pa,da=!!gi&&yn("\n","&#10;"),va=n("area,base,br,col,embed,frame,hr,img,input,isindex,keygen,link,meta,param,source,track,wbr",!0),ha=n("colgroup,dd,dt,li,options,p,td,tfoot,th,thead,tr,source",!0),ma=n("address,article,aside,base,blockquote,body,caption,col,colgroup,dd,details,dialog,div,dl,dt,fieldset,figcaption,figure,footer,form,h1,h2,h3,h4,h5,h6,head,header,hgroup,hr,html,legend,li,menuitem,meta,optgroup,option,param,rp,rt,source,style,summary,tbody,td,tfoot,th,thead,title,tr,track",!0),ga=/([^\s"'<>\/=]+)/,ya=/(?:=)/,_a=[/"([^"]*)"+/.source,/'([^']*)'+/.source,/([^\s"'=<>`]+)/.source],ba=new RegExp("^\\s*"+ga.source+"(?:\\s*("+ya.source+")\\s*(?:"+_a.join("|")+"))?"),$a="[a-zA-Z_][\\w\\-\\.]*",wa="((?:"+$a+"\\:)?"+$a+")",Ca=new RegExp("^<"+wa),xa=/^\s*(\/?)>/,ka=new RegExp("^<\\/"+wa+"[^>]*>"),Aa=/^<!DOCTYPE [^>]+>/i,Oa=/^<!--/,Sa=/^<!\[/,Ta=!1;"x".replace(/x(.)?/g,function(e,t){Ta=""===t});var Ea,Ia,ja,Na,La,Ma,Da,Pa,Ra,Fa,Ha,Ua,Ba,za,Va,Ja,Ka,qa,Wa,Za,Ga,Ya,Qa,Xa,es,ts=n("script,style",!0),ns={},rs=/&lt;/g,is=/&gt;/g,os=/&#10;/g,as=/&amp;/g,ss=/&quot;/g,cs=/\{\{((?:.|\n)+?)\}\}/g,us=/[-.*+?^${}()|[\]\/\\]/g,ls=a(function(e){var t=e[0].replace(us,"\\$&"),n=e[1].replace(us,"\\$&");return new RegExp(t+"((?:.|\\n)+?)"+n,"g")}),fs=/^v-|^@|^:/,ps=/(.*?)\s+(?:in|of)\s+(.*)/,ds=/\((\{[^}]*\}|[^,]*),([^,]*)(?:,([^,]*))?\)/,vs=/^:|^v-bind:/,hs=/^@|^v-on:/,ms=/:(.*)$/,gs=/\.[^.]+/g,ys=a(_n),_s=/^xmlns:NS\d+/,bs=/^NS\d+:/,$s=a(or),ws=/^\s*([\w$_]+|\([^)]*?\))\s*=>|^function\s*\(/,Cs=/^\s*[A-Za-z_$][\w$]*(?:\.[A-Za-z_$][\w$]*|\['.*?']|\[".*?"]|\[\d+]|\[[A-Za-z_$][\w$]*])*\s*$/,xs={esc:27,tab:9,enter:13,space:32,up:38,left:37,right:39,down:40,delete:[8,46]},ks={stop:"$event.stopPropagation();",prevent:"$event.preventDefault();",self:"if($event.target !== $event.currentTarget)return;",ctrl:"if(!$event.ctrlKey)return;",shift:"if(!$event.shiftKey)return;",alt:"if(!$event.altKey)return;",meta:"if(!$event.metaKey)return;"},As={bind:hr,cloak:d},Os={staticKeys:["staticClass"],transformNode:Fr,genData:Hr},Ss={staticKeys:["staticStyle"],transformNode:Ur,genData:Br},Ts=[Os,Ss],Es={model:zr,text:Zr,html:Gr},Is=Object.create(null),js={expectHTML:!0,modules:Ts,staticKeys:v(Ts),directives:Es,isReservedTag:xo,isUnaryTag:va,mustUseProp:po,getTagNamespace:nt,isPreTag:Co},Ns=a(function(e){var t=it(e);return t&&t.innerHTML}),Ls=Ue.prototype.$mount;return Ue.prototype.$mount=function(e,t){if(e=e&&it(e),e===document.body||e===document.documentElement)return this;var n=this.$options;if(!n.render){var r=n.template;if(r)if("string"==typeof r)"#"===r.charAt(0)&&(r=Ns(r));else{if(!r.nodeType)return this;r=r.innerHTML}else e&&(r=ei(e));if(r){var i=Qr(r,{warn:Si,shouldDecodeNewlines:da,delimiters:n.delimiters},this),o=i.render,a=i.staticRenderFns;n.render=o,n.staticRenderFns=a}}return Ls.call(this,e,t)},Ue.compile=Qr,Ue});
	}).call($global);
$packages["github.com/oskca/gopherjs-vue/jscode/minified"] = (function() {
	var $pkg = {}, $init;
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["errors"] = (function() {
	var $pkg = {}, $init, errorString, ptrType, New;
	errorString = $pkg.errorString = $newType(0, $kindStruct, "errors.errorString", true, "errors", false, function(s_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.s = "";
			return;
		}
		this.s = s_;
	});
	ptrType = $ptrType(errorString);
	New = function(text) {
		var $ptr, text;
		return new errorString.ptr(text);
	};
	$pkg.New = New;
	errorString.ptr.prototype.Error = function() {
		var $ptr, e;
		e = this;
		return e.s;
	};
	errorString.prototype.Error = function() { return this.$val.Error(); };
	ptrType.methods = [{prop: "Error", name: "Error", pkg: "", typ: $funcType([], [$String], false)}];
	errorString.init("errors", [{prop: "s", name: "s", anonymous: false, exported: false, typ: $String, tag: ""}]);
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["internal/race"] = (function() {
	var $pkg = {}, $init, Acquire, Release;
	Acquire = function(addr) {
		var $ptr, addr;
	};
	$pkg.Acquire = Acquire;
	Release = function(addr) {
		var $ptr, addr;
	};
	$pkg.Release = Release;
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["sync/atomic"] = (function() {
	var $pkg = {}, $init, js, CompareAndSwapInt32, AddInt32;
	js = $packages["github.com/gopherjs/gopherjs/js"];
	CompareAndSwapInt32 = function(addr, old, new$1) {
		var $ptr, addr, new$1, old;
		if (addr.$get() === old) {
			addr.$set(new$1);
			return true;
		}
		return false;
	};
	$pkg.CompareAndSwapInt32 = CompareAndSwapInt32;
	AddInt32 = function(addr, delta) {
		var $ptr, addr, delta, new$1;
		new$1 = addr.$get() + delta >> 0;
		addr.$set(new$1);
		return new$1;
	};
	$pkg.AddInt32 = AddInt32;
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		$r = js.$init(); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["sync"] = (function() {
	var $pkg = {}, $init, js, race, runtime, atomic, Pool, Mutex, poolLocalInternal, poolLocal, notifyList, ptrType, sliceType, ptrType$1, chanType, sliceType$1, ptrType$6, ptrType$7, sliceType$4, funcType, ptrType$16, arrayType$2, semWaiters, expunged, allPools, runtime_registerPoolCleanup, runtime_Semacquire, runtime_SemacquireMutex, runtime_Semrelease, runtime_notifyListCheck, runtime_canSpin, runtime_nanotime, poolCleanup, init, indexLocal, init$1, runtime_doSpin;
	js = $packages["github.com/gopherjs/gopherjs/js"];
	race = $packages["internal/race"];
	runtime = $packages["runtime"];
	atomic = $packages["sync/atomic"];
	Pool = $pkg.Pool = $newType(0, $kindStruct, "sync.Pool", true, "sync", true, function(local_, localSize_, store_, New_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.local = 0;
			this.localSize = 0;
			this.store = sliceType$4.nil;
			this.New = $throwNilPointerError;
			return;
		}
		this.local = local_;
		this.localSize = localSize_;
		this.store = store_;
		this.New = New_;
	});
	Mutex = $pkg.Mutex = $newType(0, $kindStruct, "sync.Mutex", true, "sync", true, function(state_, sema_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.state = 0;
			this.sema = 0;
			return;
		}
		this.state = state_;
		this.sema = sema_;
	});
	poolLocalInternal = $pkg.poolLocalInternal = $newType(0, $kindStruct, "sync.poolLocalInternal", true, "sync", false, function(private$0_, shared_, Mutex_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.private$0 = $ifaceNil;
			this.shared = sliceType$4.nil;
			this.Mutex = new Mutex.ptr(0, 0);
			return;
		}
		this.private$0 = private$0_;
		this.shared = shared_;
		this.Mutex = Mutex_;
	});
	poolLocal = $pkg.poolLocal = $newType(0, $kindStruct, "sync.poolLocal", true, "sync", false, function(poolLocalInternal_, pad_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.poolLocalInternal = new poolLocalInternal.ptr($ifaceNil, sliceType$4.nil, new Mutex.ptr(0, 0));
			this.pad = arrayType$2.zero();
			return;
		}
		this.poolLocalInternal = poolLocalInternal_;
		this.pad = pad_;
	});
	notifyList = $pkg.notifyList = $newType(0, $kindStruct, "sync.notifyList", true, "sync", false, function(wait_, notify_, lock_, head_, tail_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.wait = 0;
			this.notify = 0;
			this.lock = 0;
			this.head = 0;
			this.tail = 0;
			return;
		}
		this.wait = wait_;
		this.notify = notify_;
		this.lock = lock_;
		this.head = head_;
		this.tail = tail_;
	});
	ptrType = $ptrType(Pool);
	sliceType = $sliceType(ptrType);
	ptrType$1 = $ptrType($Uint32);
	chanType = $chanType($Bool, false, false);
	sliceType$1 = $sliceType(chanType);
	ptrType$6 = $ptrType($Int32);
	ptrType$7 = $ptrType(poolLocal);
	sliceType$4 = $sliceType($emptyInterface);
	funcType = $funcType([], [$emptyInterface], false);
	ptrType$16 = $ptrType(Mutex);
	arrayType$2 = $arrayType($Uint8, 100);
	Pool.ptr.prototype.Get = function() {
		var $ptr, _r, p, x, x$1, x$2, $s, $r;
		/* */ $s = 0; var $f, $c = false; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $ptr = $f.$ptr; _r = $f._r; p = $f.p; x = $f.x; x$1 = $f.x$1; x$2 = $f.x$2; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		p = this;
		/* */ if (p.store.$length === 0) { $s = 1; continue; }
		/* */ $s = 2; continue;
		/* if (p.store.$length === 0) { */ case 1:
			/* */ if (!(p.New === $throwNilPointerError)) { $s = 3; continue; }
			/* */ $s = 4; continue;
			/* if (!(p.New === $throwNilPointerError)) { */ case 3:
				_r = p.New(); /* */ $s = 5; case 5: if($c) { $c = false; _r = _r.$blk(); } if (_r && _r.$blk !== undefined) { break s; }
				$s = -1; return _r;
			/* } */ case 4:
			$s = -1; return $ifaceNil;
		/* } */ case 2:
		x$2 = (x = p.store, x$1 = p.store.$length - 1 >> 0, ((x$1 < 0 || x$1 >= x.$length) ? ($throwRuntimeError("index out of range"), undefined) : x.$array[x.$offset + x$1]));
		p.store = $subslice(p.store, 0, (p.store.$length - 1 >> 0));
		$s = -1; return x$2;
		/* */ } return; } if ($f === undefined) { $f = { $blk: Pool.ptr.prototype.Get }; } $f.$ptr = $ptr; $f._r = _r; $f.p = p; $f.x = x; $f.x$1 = x$1; $f.x$2 = x$2; $f.$s = $s; $f.$r = $r; return $f;
	};
	Pool.prototype.Get = function() { return this.$val.Get(); };
	Pool.ptr.prototype.Put = function(x) {
		var $ptr, p, x;
		p = this;
		if ($interfaceIsEqual(x, $ifaceNil)) {
			return;
		}
		p.store = $append(p.store, x);
	};
	Pool.prototype.Put = function(x) { return this.$val.Put(x); };
	runtime_registerPoolCleanup = function(cleanup) {
		var $ptr, cleanup;
	};
	runtime_Semacquire = function(s) {
		var $ptr, _entry, _key, _r, ch, s, $s, $r;
		/* */ $s = 0; var $f, $c = false; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $ptr = $f.$ptr; _entry = $f._entry; _key = $f._key; _r = $f._r; ch = $f.ch; s = $f.s; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		/* */ if (s.$get() === 0) { $s = 1; continue; }
		/* */ $s = 2; continue;
		/* if (s.$get() === 0) { */ case 1:
			ch = new $Chan($Bool, 0);
			_key = s; (semWaiters || $throwRuntimeError("assignment to entry in nil map"))[ptrType$1.keyFor(_key)] = { k: _key, v: $append((_entry = semWaiters[ptrType$1.keyFor(s)], _entry !== undefined ? _entry.v : sliceType$1.nil), ch) };
			_r = $recv(ch); /* */ $s = 3; case 3: if($c) { $c = false; _r = _r.$blk(); } if (_r && _r.$blk !== undefined) { break s; }
			_r[0];
		/* } */ case 2:
		s.$set(s.$get() - (1) >>> 0);
		$s = -1; return;
		/* */ } return; } if ($f === undefined) { $f = { $blk: runtime_Semacquire }; } $f.$ptr = $ptr; $f._entry = _entry; $f._key = _key; $f._r = _r; $f.ch = ch; $f.s = s; $f.$s = $s; $f.$r = $r; return $f;
	};
	runtime_SemacquireMutex = function(s, lifo) {
		var $ptr, lifo, s, $s, $r;
		/* */ $s = 0; var $f, $c = false; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $ptr = $f.$ptr; lifo = $f.lifo; s = $f.s; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		$r = runtime_Semacquire(s); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$s = -1; return;
		/* */ } return; } if ($f === undefined) { $f = { $blk: runtime_SemacquireMutex }; } $f.$ptr = $ptr; $f.lifo = lifo; $f.s = s; $f.$s = $s; $f.$r = $r; return $f;
	};
	runtime_Semrelease = function(s, handoff) {
		var $ptr, _entry, _key, ch, handoff, s, w, $s, $r;
		/* */ $s = 0; var $f, $c = false; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $ptr = $f.$ptr; _entry = $f._entry; _key = $f._key; ch = $f.ch; handoff = $f.handoff; s = $f.s; w = $f.w; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		s.$set(s.$get() + (1) >>> 0);
		w = (_entry = semWaiters[ptrType$1.keyFor(s)], _entry !== undefined ? _entry.v : sliceType$1.nil);
		if (w.$length === 0) {
			$s = -1; return;
		}
		ch = (0 >= w.$length ? ($throwRuntimeError("index out of range"), undefined) : w.$array[w.$offset + 0]);
		w = $subslice(w, 1);
		_key = s; (semWaiters || $throwRuntimeError("assignment to entry in nil map"))[ptrType$1.keyFor(_key)] = { k: _key, v: w };
		if (w.$length === 0) {
			delete semWaiters[ptrType$1.keyFor(s)];
		}
		$r = $send(ch, true); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$s = -1; return;
		/* */ } return; } if ($f === undefined) { $f = { $blk: runtime_Semrelease }; } $f.$ptr = $ptr; $f._entry = _entry; $f._key = _key; $f.ch = ch; $f.handoff = handoff; $f.s = s; $f.w = w; $f.$s = $s; $f.$r = $r; return $f;
	};
	runtime_notifyListCheck = function(size) {
		var $ptr, size;
	};
	runtime_canSpin = function(i) {
		var $ptr, i;
		return false;
	};
	runtime_nanotime = function() {
		var $ptr;
		return $mul64($internalize(new ($global.Date)().getTime(), $Int64), new $Int64(0, 1000000));
	};
	Mutex.ptr.prototype.Lock = function() {
		var $ptr, awoke, delta, iter, m, new$1, old, queueLifo, starving, waitStartTime, x, x$1, $s, $r;
		/* */ $s = 0; var $f, $c = false; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $ptr = $f.$ptr; awoke = $f.awoke; delta = $f.delta; iter = $f.iter; m = $f.m; new$1 = $f.new$1; old = $f.old; queueLifo = $f.queueLifo; starving = $f.starving; waitStartTime = $f.waitStartTime; x = $f.x; x$1 = $f.x$1; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		m = this;
		if (atomic.CompareAndSwapInt32((m.$ptr_state || (m.$ptr_state = new ptrType$6(function() { return this.$target.state; }, function($v) { this.$target.state = $v; }, m))), 0, 1)) {
			if (false) {
				race.Acquire((m));
			}
			$s = -1; return;
		}
		waitStartTime = new $Int64(0, 0);
		starving = false;
		awoke = false;
		iter = 0;
		old = m.state;
		/* while (true) { */ case 1:
			/* */ if (((old & 5) === 1) && runtime_canSpin(iter)) { $s = 3; continue; }
			/* */ $s = 4; continue;
			/* if (((old & 5) === 1) && runtime_canSpin(iter)) { */ case 3:
				if (!awoke && ((old & 2) === 0) && !(((old >> 3 >> 0) === 0)) && atomic.CompareAndSwapInt32((m.$ptr_state || (m.$ptr_state = new ptrType$6(function() { return this.$target.state; }, function($v) { this.$target.state = $v; }, m))), old, old | 2)) {
					awoke = true;
				}
				runtime_doSpin();
				iter = iter + (1) >> 0;
				old = m.state;
				/* continue; */ $s = 1; continue;
			/* } */ case 4:
			new$1 = old;
			if ((old & 4) === 0) {
				new$1 = new$1 | (1);
			}
			if (!(((old & 5) === 0))) {
				new$1 = new$1 + (8) >> 0;
			}
			if (starving && !(((old & 1) === 0))) {
				new$1 = new$1 | (4);
			}
			if (awoke) {
				if ((new$1 & 2) === 0) {
					$panic(new $String("sync: inconsistent mutex state"));
				}
				new$1 = (new$1 & ~(2)) >> 0;
			}
			/* */ if (atomic.CompareAndSwapInt32((m.$ptr_state || (m.$ptr_state = new ptrType$6(function() { return this.$target.state; }, function($v) { this.$target.state = $v; }, m))), old, new$1)) { $s = 5; continue; }
			/* */ $s = 6; continue;
			/* if (atomic.CompareAndSwapInt32((m.$ptr_state || (m.$ptr_state = new ptrType$6(function() { return this.$target.state; }, function($v) { this.$target.state = $v; }, m))), old, new$1)) { */ case 5:
				if ((old & 5) === 0) {
					/* break; */ $s = 2; continue;
				}
				queueLifo = !((waitStartTime.$high === 0 && waitStartTime.$low === 0));
				if ((waitStartTime.$high === 0 && waitStartTime.$low === 0)) {
					waitStartTime = runtime_nanotime();
				}
				$r = runtime_SemacquireMutex((m.$ptr_sema || (m.$ptr_sema = new ptrType$1(function() { return this.$target.sema; }, function($v) { this.$target.sema = $v; }, m))), queueLifo); /* */ $s = 8; case 8: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
				starving = starving || (x = (x$1 = runtime_nanotime(), new $Int64(x$1.$high - waitStartTime.$high, x$1.$low - waitStartTime.$low)), (x.$high > 0 || (x.$high === 0 && x.$low > 1000000)));
				old = m.state;
				if (!(((old & 4) === 0))) {
					if (!(((old & 3) === 0)) || ((old >> 3 >> 0) === 0)) {
						$panic(new $String("sync: inconsistent mutex state"));
					}
					delta = -7;
					if (!starving || ((old >> 3 >> 0) === 1)) {
						delta = delta - (4) >> 0;
					}
					atomic.AddInt32((m.$ptr_state || (m.$ptr_state = new ptrType$6(function() { return this.$target.state; }, function($v) { this.$target.state = $v; }, m))), delta);
					/* break; */ $s = 2; continue;
				}
				awoke = true;
				iter = 0;
				$s = 7; continue;
			/* } else { */ case 6:
				old = m.state;
			/* } */ case 7:
		/* } */ $s = 1; continue; case 2:
		if (false) {
			race.Acquire((m));
		}
		$s = -1; return;
		/* */ } return; } if ($f === undefined) { $f = { $blk: Mutex.ptr.prototype.Lock }; } $f.$ptr = $ptr; $f.awoke = awoke; $f.delta = delta; $f.iter = iter; $f.m = m; $f.new$1 = new$1; $f.old = old; $f.queueLifo = queueLifo; $f.starving = starving; $f.waitStartTime = waitStartTime; $f.x = x; $f.x$1 = x$1; $f.$s = $s; $f.$r = $r; return $f;
	};
	Mutex.prototype.Lock = function() { return this.$val.Lock(); };
	Mutex.ptr.prototype.Unlock = function() {
		var $ptr, m, new$1, old, $s, $r;
		/* */ $s = 0; var $f, $c = false; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $ptr = $f.$ptr; m = $f.m; new$1 = $f.new$1; old = $f.old; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		m = this;
		if (false) {
			$unused(m.state);
			race.Release((m));
		}
		new$1 = atomic.AddInt32((m.$ptr_state || (m.$ptr_state = new ptrType$6(function() { return this.$target.state; }, function($v) { this.$target.state = $v; }, m))), -1);
		if ((((new$1 + 1 >> 0)) & 1) === 0) {
			$panic(new $String("sync: unlock of unlocked mutex"));
		}
		/* */ if ((new$1 & 4) === 0) { $s = 1; continue; }
		/* */ $s = 2; continue;
		/* if ((new$1 & 4) === 0) { */ case 1:
			old = new$1;
			/* while (true) { */ case 4:
				if (((old >> 3 >> 0) === 0) || !(((old & 7) === 0))) {
					$s = -1; return;
				}
				new$1 = ((old - 8 >> 0)) | 2;
				/* */ if (atomic.CompareAndSwapInt32((m.$ptr_state || (m.$ptr_state = new ptrType$6(function() { return this.$target.state; }, function($v) { this.$target.state = $v; }, m))), old, new$1)) { $s = 6; continue; }
				/* */ $s = 7; continue;
				/* if (atomic.CompareAndSwapInt32((m.$ptr_state || (m.$ptr_state = new ptrType$6(function() { return this.$target.state; }, function($v) { this.$target.state = $v; }, m))), old, new$1)) { */ case 6:
					$r = runtime_Semrelease((m.$ptr_sema || (m.$ptr_sema = new ptrType$1(function() { return this.$target.sema; }, function($v) { this.$target.sema = $v; }, m))), false); /* */ $s = 8; case 8: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
					$s = -1; return;
				/* } */ case 7:
				old = m.state;
			/* } */ $s = 4; continue; case 5:
			$s = 3; continue;
		/* } else { */ case 2:
			$r = runtime_Semrelease((m.$ptr_sema || (m.$ptr_sema = new ptrType$1(function() { return this.$target.sema; }, function($v) { this.$target.sema = $v; }, m))), true); /* */ $s = 9; case 9: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		/* } */ case 3:
		$s = -1; return;
		/* */ } return; } if ($f === undefined) { $f = { $blk: Mutex.ptr.prototype.Unlock }; } $f.$ptr = $ptr; $f.m = m; $f.new$1 = new$1; $f.old = old; $f.$s = $s; $f.$r = $r; return $f;
	};
	Mutex.prototype.Unlock = function() { return this.$val.Unlock(); };
	poolCleanup = function() {
		var $ptr, _i, _i$1, _ref, _ref$1, i, i$1, j, l, p, x;
		_ref = allPools;
		_i = 0;
		while (true) {
			if (!(_i < _ref.$length)) { break; }
			i = _i;
			p = ((_i < 0 || _i >= _ref.$length) ? ($throwRuntimeError("index out of range"), undefined) : _ref.$array[_ref.$offset + _i]);
			((i < 0 || i >= allPools.$length) ? ($throwRuntimeError("index out of range"), undefined) : allPools.$array[allPools.$offset + i] = ptrType.nil);
			i$1 = 0;
			while (true) {
				if (!(i$1 < ((p.localSize >> 0)))) { break; }
				l = indexLocal(p.local, i$1);
				l.poolLocalInternal.private$0 = $ifaceNil;
				_ref$1 = l.poolLocalInternal.shared;
				_i$1 = 0;
				while (true) {
					if (!(_i$1 < _ref$1.$length)) { break; }
					j = _i$1;
					(x = l.poolLocalInternal.shared, ((j < 0 || j >= x.$length) ? ($throwRuntimeError("index out of range"), undefined) : x.$array[x.$offset + j] = $ifaceNil));
					_i$1++;
				}
				l.poolLocalInternal.shared = sliceType$4.nil;
				i$1 = i$1 + (1) >> 0;
			}
			p.local = 0;
			p.localSize = 0;
			_i++;
		}
		allPools = new sliceType([]);
	};
	init = function() {
		var $ptr;
		runtime_registerPoolCleanup(poolCleanup);
	};
	indexLocal = function(l, i) {
		var $ptr, i, l, lp;
		lp = (((l) + ($imul(((i >>> 0)), 128) >>> 0) >>> 0));
		return ($pointerOfStructConversion(lp, ptrType$7));
	};
	init$1 = function() {
		var $ptr, n;
		n = new notifyList.ptr(0, 0, 0, 0, 0);
		runtime_notifyListCheck(20);
	};
	runtime_doSpin = function() {
		$throwRuntimeError("native function not implemented: sync.runtime_doSpin");
	};
	ptrType.methods = [{prop: "Get", name: "Get", pkg: "", typ: $funcType([], [$emptyInterface], false)}, {prop: "Put", name: "Put", pkg: "", typ: $funcType([$emptyInterface], [], false)}, {prop: "getSlow", name: "getSlow", pkg: "sync", typ: $funcType([], [$emptyInterface], false)}, {prop: "pin", name: "pin", pkg: "sync", typ: $funcType([], [ptrType$7], false)}, {prop: "pinSlow", name: "pinSlow", pkg: "sync", typ: $funcType([], [ptrType$7], false)}];
	ptrType$16.methods = [{prop: "Lock", name: "Lock", pkg: "", typ: $funcType([], [], false)}, {prop: "Unlock", name: "Unlock", pkg: "", typ: $funcType([], [], false)}];
	Pool.init("sync", [{prop: "local", name: "local", anonymous: false, exported: false, typ: $UnsafePointer, tag: ""}, {prop: "localSize", name: "localSize", anonymous: false, exported: false, typ: $Uintptr, tag: ""}, {prop: "store", name: "store", anonymous: false, exported: false, typ: sliceType$4, tag: ""}, {prop: "New", name: "New", anonymous: false, exported: true, typ: funcType, tag: ""}]);
	Mutex.init("sync", [{prop: "state", name: "state", anonymous: false, exported: false, typ: $Int32, tag: ""}, {prop: "sema", name: "sema", anonymous: false, exported: false, typ: $Uint32, tag: ""}]);
	poolLocalInternal.init("sync", [{prop: "private$0", name: "private", anonymous: false, exported: false, typ: $emptyInterface, tag: ""}, {prop: "shared", name: "shared", anonymous: false, exported: false, typ: sliceType$4, tag: ""}, {prop: "Mutex", name: "Mutex", anonymous: true, exported: true, typ: Mutex, tag: ""}]);
	poolLocal.init("sync", [{prop: "poolLocalInternal", name: "poolLocalInternal", anonymous: true, exported: false, typ: poolLocalInternal, tag: ""}, {prop: "pad", name: "pad", anonymous: false, exported: false, typ: arrayType$2, tag: ""}]);
	notifyList.init("sync", [{prop: "wait", name: "wait", anonymous: false, exported: false, typ: $Uint32, tag: ""}, {prop: "notify", name: "notify", anonymous: false, exported: false, typ: $Uint32, tag: ""}, {prop: "lock", name: "lock", anonymous: false, exported: false, typ: $Uintptr, tag: ""}, {prop: "head", name: "head", anonymous: false, exported: false, typ: $UnsafePointer, tag: ""}, {prop: "tail", name: "tail", anonymous: false, exported: false, typ: $UnsafePointer, tag: ""}]);
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		$r = js.$init(); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = race.$init(); /* */ $s = 2; case 2: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = runtime.$init(); /* */ $s = 3; case 3: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = atomic.$init(); /* */ $s = 4; case 4: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		allPools = sliceType.nil;
		semWaiters = {};
		expunged = (new Uint8Array(8));
		init();
		init$1();
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["io"] = (function() {
	var $pkg = {}, $init, errors, sync, errWhence, errOffset;
	errors = $packages["errors"];
	sync = $packages["sync"];
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		$r = errors.$init(); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = sync.$init(); /* */ $s = 2; case 2: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$pkg.ErrShortWrite = errors.New("short write");
		$pkg.ErrShortBuffer = errors.New("short buffer");
		$pkg.EOF = errors.New("EOF");
		$pkg.ErrUnexpectedEOF = errors.New("unexpected EOF");
		$pkg.ErrNoProgress = errors.New("multiple Read calls return no data or error");
		errWhence = errors.New("Seek: invalid whence");
		errOffset = errors.New("Seek: invalid offset");
		$pkg.ErrClosedPipe = errors.New("io: read/write on closed pipe");
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["unicode"] = (function() {
	var $pkg = {}, $init;
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["unicode/utf8"] = (function() {
	var $pkg = {}, $init;
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["strings"] = (function() {
	var $pkg = {}, $init, errors, js, io, unicode, utf8, HasPrefix;
	errors = $packages["errors"];
	js = $packages["github.com/gopherjs/gopherjs/js"];
	io = $packages["io"];
	unicode = $packages["unicode"];
	utf8 = $packages["unicode/utf8"];
	HasPrefix = function(s, prefix) {
		var $ptr, prefix, s;
		return s.length >= prefix.length && $substring(s, 0, prefix.length) === prefix;
	};
	$pkg.HasPrefix = HasPrefix;
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		$r = errors.$init(); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = js.$init(); /* */ $s = 2; case 2: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = io.$init(); /* */ $s = 3; case 3: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = unicode.$init(); /* */ $s = 4; case 4: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = utf8.$init(); /* */ $s = 5; case 5: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["github.com/oskca/gopherjs-vue"] = (function() {
	var $pkg = {}, $init, js, dom, json, minifiled, strings, pool, Component, TConfig, LifeCycleEvent, Option, CreateElement, Render, ViewModel, ptrType, sliceType, sliceType$1, ptrType$1, ptrType$2, funcType, funcType$1, sliceType$2, mapType, sliceType$3, mapType$1, ptrType$3, funcType$3, funcType$4, mapType$2, ptrType$7, sliceType$4, funcType$5, funcType$6, funcType$7, sliceType$5, funcType$8, funcType$9, funcType$10, funcType$11, funcType$12, funcType$13, funcType$14, funcType$15, funcType$16, creatorPool, vue, vMap, newComponent, NewOption, New, newViewModel;
	js = $packages["github.com/gopherjs/gopherjs/js"];
	dom = $packages["github.com/oskca/gopherjs-dom"];
	json = $packages["github.com/oskca/gopherjs-json"];
	minifiled = $packages["github.com/oskca/gopherjs-vue/jscode/minified"];
	strings = $packages["strings"];
	pool = $pkg.pool = $newType(0, $kindStruct, "vue.pool", true, "github.com/oskca/gopherjs-vue", false, function(creator_, structPtr_, counter_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.creator = $throwNilPointerError;
			this.structPtr = $ifaceNil;
			this.counter = 0;
			return;
		}
		this.creator = creator_;
		this.structPtr = structPtr_;
		this.counter = counter_;
	});
	Component = $pkg.Component = $newType(0, $kindStruct, "vue.Component", true, "github.com/oskca/gopherjs-vue", true, function(ViewModel_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.ViewModel = ptrType$1.nil;
			return;
		}
		this.ViewModel = ViewModel_;
	});
	TConfig = $pkg.TConfig = $newType(0, $kindStruct, "vue.TConfig", true, "github.com/oskca/gopherjs-vue", true, function(Object_, Silent_, OptionMergeStrategies_, Devtools_, ErrorHandler_, IgnoredElements_, KeyCodes_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.Object = null;
			this.Silent = false;
			this.OptionMergeStrategies = $ifaceNil;
			this.Devtools = false;
			this.ErrorHandler = $throwNilPointerError;
			this.IgnoredElements = sliceType$1.nil;
			this.KeyCodes = false;
			return;
		}
		this.Object = Object_;
		this.Silent = Silent_;
		this.OptionMergeStrategies = OptionMergeStrategies_;
		this.Devtools = Devtools_;
		this.ErrorHandler = ErrorHandler_;
		this.IgnoredElements = IgnoredElements_;
		this.KeyCodes = KeyCodes_;
	});
	LifeCycleEvent = $pkg.LifeCycleEvent = $newType(8, $kindString, "vue.LifeCycleEvent", true, "github.com/oskca/gopherjs-vue", true, null);
	Option = $pkg.Option = $newType(0, $kindStruct, "vue.Option", true, "github.com/oskca/gopherjs-vue", true, function(Object_, Name_, Data_, El_, Template_, Parent_, Delimiters_, Functional_, coms_, props_, mixins_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.Object = null;
			this.Name = "";
			this.Data = $ifaceNil;
			this.El = $ifaceNil;
			this.Template = "";
			this.Parent = null;
			this.Delimiters = sliceType$1.nil;
			this.Functional = sliceType$1.nil;
			this.coms = false;
			this.props = sliceType$1.nil;
			this.mixins = sliceType$2.nil;
			return;
		}
		this.Object = Object_;
		this.Name = Name_;
		this.Data = Data_;
		this.El = El_;
		this.Template = Template_;
		this.Parent = Parent_;
		this.Delimiters = Delimiters_;
		this.Functional = Functional_;
		this.coms = coms_;
		this.props = props_;
		this.mixins = mixins_;
	});
	CreateElement = $pkg.CreateElement = $newType(4, $kindFunc, "vue.CreateElement", true, "github.com/oskca/gopherjs-vue", true, null);
	Render = $pkg.Render = $newType(4, $kindFunc, "vue.Render", true, "github.com/oskca/gopherjs-vue", true, null);
	ViewModel = $pkg.ViewModel = $newType(0, $kindStruct, "vue.ViewModel", true, "github.com/oskca/gopherjs-vue", true, function(Object_, Data_, El_, Options_, Parent_, Root_, Children_, Slots_, Refs_, IsServer_, WatchEx_, Set_, Delete_, On_, Once_, Off_, Emit_, Mount_, ForceUpdate_, NextTick_, Destroy_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.Object = null;
			this.Data = null;
			this.El = null;
			this.Options = null;
			this.Parent = null;
			this.Root = null;
			this.Children = null;
			this.Slots = null;
			this.Refs = null;
			this.IsServer = false;
			this.WatchEx = $throwNilPointerError;
			this.Set = $throwNilPointerError;
			this.Delete = $throwNilPointerError;
			this.On = $throwNilPointerError;
			this.Once = $throwNilPointerError;
			this.Off = $throwNilPointerError;
			this.Emit = $throwNilPointerError;
			this.Mount = $throwNilPointerError;
			this.ForceUpdate = $throwNilPointerError;
			this.NextTick = $throwNilPointerError;
			this.Destroy = $throwNilPointerError;
			return;
		}
		this.Object = Object_;
		this.Data = Data_;
		this.El = El_;
		this.Options = Options_;
		this.Parent = Parent_;
		this.Root = Root_;
		this.Children = Children_;
		this.Slots = Slots_;
		this.Refs = Refs_;
		this.IsServer = IsServer_;
		this.WatchEx = WatchEx_;
		this.Set = Set_;
		this.Delete = Delete_;
		this.On = On_;
		this.Once = Once_;
		this.Off = Off_;
		this.Emit = Emit_;
		this.Mount = Mount_;
		this.ForceUpdate = ForceUpdate_;
		this.NextTick = NextTick_;
		this.Destroy = Destroy_;
	});
	ptrType = $ptrType(pool);
	sliceType = $sliceType(ptrType);
	sliceType$1 = $sliceType($String);
	ptrType$1 = $ptrType(ViewModel);
	ptrType$2 = $ptrType(Component);
	funcType = $funcType([], [$emptyInterface], false);
	funcType$1 = $funcType([], [], false);
	sliceType$2 = $sliceType(js.M);
	mapType = $mapType($String, ptrType$2);
	sliceType$3 = $sliceType($emptyInterface);
	mapType$1 = $mapType($String, js.M);
	ptrType$3 = $ptrType(js.Object);
	funcType$3 = $funcType([ptrType$3], [], false);
	funcType$4 = $funcType([ptrType$3, ptrType$3], [], false);
	mapType$2 = $mapType($String, $Int);
	ptrType$7 = $ptrType(Option);
	sliceType$4 = $sliceType(ptrType$3);
	funcType$5 = $funcType([ptrType$1, sliceType$4], [], false);
	funcType$6 = $funcType([ptrType$1], [$emptyInterface], false);
	funcType$7 = $funcType([ptrType$1, ptrType$3], [], false);
	sliceType$5 = $sliceType(funcType$7);
	funcType$8 = $funcType([ptrType$1], [], false);
	funcType$9 = $funcType([$String, funcType$4, $Bool], [funcType$1], false);
	funcType$10 = $funcType([$String, $emptyInterface], [], false);
	funcType$11 = $funcType([$String], [], false);
	funcType$12 = $funcType([sliceType$1], [], true);
	funcType$13 = $funcType([$String, sliceType$3], [], true);
	funcType$14 = $funcType([sliceType$3], [ptrType$3], true);
	funcType$15 = $funcType([funcType$1], [], false);
	funcType$16 = $funcType([$Bool], [], false);
	Component.ptr.prototype.New = function() {
		var $ptr, c;
		c = this;
		return newViewModel(new (c.ViewModel.Object)());
	};
	Component.prototype.New = function() { return this.$val.New(); };
	newComponent = function(o) {
		var $ptr, o;
		return new Component.ptr(newViewModel(o));
	};
	Component.ptr.prototype.Register = function(name) {
		var $ptr, c, name;
		c = this;
		vue.component($externalize(name, $String), $externalize(c, ptrType$2));
		return c;
	};
	Component.prototype.Register = function(name) { return this.$val.Register(name); };
	ViewModel.ptr.prototype.FromJS = function(obj) {
		var $ptr, _i, _ref, key, obj, v;
		v = this;
		_ref = js.Keys(obj);
		_i = 0;
		while (true) {
			if (!(_i < _ref.$length)) { break; }
			key = ((_i < 0 || _i >= _ref.$length) ? ($throwRuntimeError("index out of range"), undefined) : _ref.$array[_ref.$offset + _i]);
			if (strings.HasPrefix(key, "$") || strings.HasPrefix(key, "_")) {
				_i++;
				continue;
			}
			v.Object[$externalize(key, $String)] = obj[$externalize(key, $String)];
			_i++;
		}
		return v;
	};
	ViewModel.prototype.FromJS = function(obj) { return this.$val.FromJS(obj); };
	ViewModel.ptr.prototype.FromJSON = function(jsonStr) {
		var $ptr, jsonStr, v;
		v = this;
		return v.FromJS(json.Parse(jsonStr));
	};
	ViewModel.prototype.FromJSON = function(jsonStr) { return this.$val.FromJSON(jsonStr); };
	ViewModel.ptr.prototype.ToJS = function() {
		var $ptr, _i, _ref, key, obj, v;
		v = this;
		obj = new ($global.Object)();
		_ref = js.Keys(v.Object);
		_i = 0;
		while (true) {
			if (!(_i < _ref.$length)) { break; }
			key = ((_i < 0 || _i >= _ref.$length) ? ($throwRuntimeError("index out of range"), undefined) : _ref.$array[_ref.$offset + _i]);
			if (strings.HasPrefix(key, "$") || strings.HasPrefix(key, "_")) {
				_i++;
				continue;
			}
			obj[$externalize(key, $String)] = v.Object[$externalize(key, $String)];
			_i++;
		}
		return obj;
	};
	ViewModel.prototype.ToJS = function() { return this.$val.ToJS(); };
	ViewModel.ptr.prototype.ToJSON = function() {
		var $ptr, v;
		v = this;
		return json.Stringify(new $String(v.ToJSON()));
	};
	ViewModel.prototype.ToJSON = function() { return this.$val.ToJSON(); };
	NewOption = function() {
		var $ptr, c;
		c = new Option.ptr(new ($global.Object)(), "", $ifaceNil, $ifaceNil, "", null, sliceType$1.nil, sliceType$1.nil, false, sliceType$1.nil, sliceType$2.nil);
		c.coms = {};
		c.props = new sliceType$1([]);
		c.mixins = new sliceType$2([]);
		return c;
	};
	$pkg.NewOption = NewOption;
	Option.ptr.prototype.NewViewModel = function() {
		var $ptr, o;
		o = this;
		return newViewModel(new (vue)(o.prepare()));
	};
	Option.prototype.NewViewModel = function() { return this.$val.NewViewModel(); };
	Option.ptr.prototype.NewComponent = function() {
		var $ptr, _tuple, o, ok;
		o = this;
		_tuple = $assertType($internalize(o.Object.el, $emptyInterface), $String, true);
		ok = _tuple[1];
		if (ok) {
			$panic(new $String("Option.El in component must be a function"));
		}
		return newComponent(vue.extend(o.prepare()));
	};
	Option.prototype.NewComponent = function() { return this.$val.NewComponent(); };
	Option.ptr.prototype.prepare = function() {
		var $ptr, c, opts;
		opts = null;
		c = this;
		if ($keys(c.coms).length > 0) {
			c.Object.components = $externalize(c.coms, mapType);
		}
		if (c.props.$length > 0) {
			c.Object.props = $externalize(c.props, sliceType$1);
		}
		if (c.mixins.$length > 0) {
			c.Object.mixins = $externalize(c.mixins, sliceType$2);
		}
		opts = c.Object;
		return opts;
	};
	Option.prototype.prepare = function() { return this.$val.prepare(); };
	Option.ptr.prototype.SetDataWithMethods = function(structPtr) {
		var $ptr, c, structPtr;
		c = this;
		if ($interfaceIsEqual(structPtr, $ifaceNil)) {
			return c;
		}
		c.Object.data = $externalize(structPtr, $emptyInterface);
		c.Object.methods = js.MakeWrapper(structPtr);
		return c;
	};
	Option.prototype.SetDataWithMethods = function(structPtr) { return this.$val.SetDataWithMethods(structPtr); };
	Option.ptr.prototype.AddMethod = function(name, fn) {
		var $ptr, fn, name, o;
		o = this;
		return o.addMixin("methods", new js.M($makeMap($String.keyFor, [{ k: name, v: new $jsObjectPtr(js.MakeFunc((function $b(this$1, arguments$1) {
			var $ptr, arguments$1, this$1, vm, $s, $r;
			/* */ $s = 0; var $f, $c = false; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $ptr = $f.$ptr; arguments$1 = $f.arguments$1; this$1 = $f.this$1; vm = $f.vm; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
			vm = newViewModel(this$1);
			$r = fn(vm, arguments$1); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
			$s = -1; return $ifaceNil;
			/* */ } return; } if ($f === undefined) { $f = { $blk: $b }; } $f.$ptr = $ptr; $f.arguments$1 = arguments$1; $f.this$1 = this$1; $f.vm = vm; $f.$s = $s; $f.$r = $r; return $f;
		}))) }])));
	};
	Option.prototype.AddMethod = function(name, fn) { return this.$val.AddMethod(name, fn); };
	Option.ptr.prototype.SetRender = function(r) {
		var $ptr, fn, o, r;
		o = this;
		fn = js.MakeFunc((function $b(this$1, arguments$1) {
			var $ptr, arguments$1, createElement, jsCreateElement, this$1, vm, $s, $r;
			/* */ $s = 0; var $f, $c = false; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $ptr = $f.$ptr; arguments$1 = $f.arguments$1; createElement = $f.createElement; jsCreateElement = $f.jsCreateElement; this$1 = $f.this$1; vm = $f.vm; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
			jsCreateElement = [jsCreateElement];
			vm = newViewModel(this$1);
			jsCreateElement[0] = (0 >= arguments$1.$length ? ($throwRuntimeError("index out of range"), undefined) : arguments$1.$array[arguments$1.$offset + 0]);
			createElement = (function(jsCreateElement) { return function(tagName, data, children) {
				var $ptr, children, data, tagName, vnode;
				vnode = null;
				vnode = jsCreateElement[0][$externalize(tagName, $String)]($externalize(data, $emptyInterface), $externalize(children, sliceType$3));
				return vnode;
			}; })(jsCreateElement);
			$r = r(vm, createElement); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
			$s = -1; return $ifaceNil;
			/* */ } return; } if ($f === undefined) { $f = { $blk: $b }; } $f.$ptr = $ptr; $f.arguments$1 = arguments$1; $f.createElement = createElement; $f.jsCreateElement = jsCreateElement; $f.this$1 = this$1; $f.vm = vm; $f.$s = $s; $f.$r = $r; return $f;
		}));
		o.Object.render = fn;
	};
	Option.prototype.SetRender = function(r) { return this.$val.SetRender(r); };
	Option.ptr.prototype.AddComputed = function(name, getter, setter) {
		var $ptr, _entry, _entry$1, _key, _key$1, _key$2, conf, fnGetter, fnSetter, getter, name, o, setter;
		o = this;
		conf = {};
		_key = name; (conf || $throwRuntimeError("assignment to entry in nil map"))[$String.keyFor(_key)] = { k: _key, v: {} };
		fnGetter = js.MakeFunc((function $b(this$1, arguments$1) {
			var $ptr, _r, arguments$1, this$1, vm, $s, $r;
			/* */ $s = 0; var $f, $c = false; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $ptr = $f.$ptr; _r = $f._r; arguments$1 = $f.arguments$1; this$1 = $f.this$1; vm = $f.vm; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
			vm = newViewModel(this$1);
			_r = getter(vm); /* */ $s = 1; case 1: if($c) { $c = false; _r = _r.$blk(); } if (_r && _r.$blk !== undefined) { break s; }
			$s = -1; return _r;
			/* */ } return; } if ($f === undefined) { $f = { $blk: $b }; } $f.$ptr = $ptr; $f._r = _r; $f.arguments$1 = arguments$1; $f.this$1 = this$1; $f.vm = vm; $f.$s = $s; $f.$r = $r; return $f;
		}));
		_key$1 = "get"; ((_entry = conf[$String.keyFor(name)], _entry !== undefined ? _entry.v : false) || $throwRuntimeError("assignment to entry in nil map"))[$String.keyFor(_key$1)] = { k: _key$1, v: new $jsObjectPtr(fnGetter) };
		if (setter.$length > 0) {
			fnSetter = js.MakeFunc((function $b(this$1, arguments$1) {
				var $ptr, arguments$1, this$1, vm, $s, $r;
				/* */ $s = 0; var $f, $c = false; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $ptr = $f.$ptr; arguments$1 = $f.arguments$1; this$1 = $f.this$1; vm = $f.vm; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
				vm = newViewModel(this$1);
				$r = (0 >= setter.$length ? ($throwRuntimeError("index out of range"), undefined) : setter.$array[setter.$offset + 0])(vm, (0 >= arguments$1.$length ? ($throwRuntimeError("index out of range"), undefined) : arguments$1.$array[arguments$1.$offset + 0])); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
				$s = -1; return $ifaceNil;
				/* */ } return; } if ($f === undefined) { $f = { $blk: $b }; } $f.$ptr = $ptr; $f.arguments$1 = arguments$1; $f.this$1 = this$1; $f.vm = vm; $f.$s = $s; $f.$r = $r; return $f;
			}));
			_key$2 = "set"; ((_entry$1 = conf[$String.keyFor(name)], _entry$1 !== undefined ? _entry$1.v : false) || $throwRuntimeError("assignment to entry in nil map"))[$String.keyFor(_key$2)] = { k: _key$2, v: new $jsObjectPtr(fnSetter) };
		}
		o.addMixin("computed", new mapType$1(conf));
	};
	Option.prototype.AddComputed = function(name, getter, setter) { return this.$val.AddComputed(name, getter, setter); };
	Option.ptr.prototype.OnLifeCycleEvent = function(evt, fn) {
		var $ptr, evt, fn, o;
		o = this;
		return o.addMixin((evt), new $jsObjectPtr(js.MakeFunc((function $b(this$1, arguments$1) {
			var $ptr, arguments$1, this$1, vm, $s, $r;
			/* */ $s = 0; var $f, $c = false; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $ptr = $f.$ptr; arguments$1 = $f.arguments$1; this$1 = $f.this$1; vm = $f.vm; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
			vm = newViewModel(this$1);
			$r = fn(vm); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
			$s = -1; return $ifaceNil;
			/* */ } return; } if ($f === undefined) { $f = { $blk: $b }; } $f.$ptr = $ptr; $f.arguments$1 = arguments$1; $f.this$1 = this$1; $f.vm = vm; $f.$s = $s; $f.$r = $r; return $f;
		}))));
	};
	Option.prototype.OnLifeCycleEvent = function(evt, fn) { return this.$val.OnLifeCycleEvent(evt, fn); };
	Option.ptr.prototype.Mixin = function(val) {
		var $ptr, c, val;
		c = this;
		c.mixins = $append(c.mixins, val);
		return c;
	};
	Option.prototype.Mixin = function(val) { return this.$val.Mixin(val); };
	Option.ptr.prototype.addMixin = function(name, val) {
		var $ptr, c, name, val;
		c = this;
		return c.Mixin($makeMap($String.keyFor, [{ k: name, v: val }]));
	};
	Option.prototype.addMixin = function(name, val) { return this.$val.addMixin(name, val); };
	Option.ptr.prototype.AddSubComponent = function(name, sub) {
		var $ptr, _key, c, name, sub;
		c = this;
		_key = name; (c.coms || $throwRuntimeError("assignment to entry in nil map"))[$String.keyFor(_key)] = { k: _key, v: sub };
		return c;
	};
	Option.prototype.AddSubComponent = function(name, sub) { return this.$val.AddSubComponent(name, sub); };
	Option.ptr.prototype.AddProp = function(name) {
		var $ptr, c, name;
		c = this;
		c.props = $appendSlice(c.props, name);
		return c;
	};
	Option.prototype.AddProp = function(name) { return this.$val.AddProp(name); };
	New = function(selectorOrHTMLElement, structPtr) {
		var $ptr, _key, opt, selectorOrHTMLElement, structPtr, vm;
		opt = NewOption();
		opt.Object.el = $externalize(selectorOrHTMLElement, $emptyInterface);
		opt.SetDataWithMethods(structPtr);
		vm = opt.NewViewModel();
		_key = structPtr; (vMap || $throwRuntimeError("assignment to entry in nil map"))[$emptyInterface.keyFor(_key)] = { k: _key, v: vm };
		return vm;
	};
	$pkg.New = New;
	newViewModel = function(o) {
		var $ptr, o;
		return new ViewModel.ptr(o, null, null, null, null, null, null, null, null, false, $throwNilPointerError, $throwNilPointerError, $throwNilPointerError, $throwNilPointerError, $throwNilPointerError, $throwNilPointerError, $throwNilPointerError, $throwNilPointerError, $throwNilPointerError, $throwNilPointerError, $throwNilPointerError);
	};
	ViewModel.ptr.prototype.Watch = function(expression, callback) {
		var $ptr, callback, expression, obj, unwatcher, v;
		unwatcher = $throwNilPointerError;
		v = this;
		obj = v.Object.$watch($externalize(expression, $String), $externalize(callback, funcType$3));
		unwatcher = (function() {
			var $ptr;
			obj();
		});
		return unwatcher;
	};
	ViewModel.prototype.Watch = function(expression, callback) { return this.$val.Watch(expression, callback); };
	ptrType$2.methods = [{prop: "New", name: "New", pkg: "", typ: $funcType([], [ptrType$1], false)}, {prop: "Register", name: "Register", pkg: "", typ: $funcType([$String], [ptrType$2], false)}];
	ptrType$7.methods = [{prop: "NewViewModel", name: "NewViewModel", pkg: "", typ: $funcType([], [ptrType$1], false)}, {prop: "NewComponent", name: "NewComponent", pkg: "", typ: $funcType([], [ptrType$2], false)}, {prop: "prepare", name: "prepare", pkg: "github.com/oskca/gopherjs-vue", typ: $funcType([], [ptrType$3], false)}, {prop: "SetDataWithMethods", name: "SetDataWithMethods", pkg: "", typ: $funcType([$emptyInterface], [ptrType$7], false)}, {prop: "AddMethod", name: "AddMethod", pkg: "", typ: $funcType([$String, funcType$5], [ptrType$7], false)}, {prop: "SetRender", name: "SetRender", pkg: "", typ: $funcType([Render], [], false)}, {prop: "AddComputed", name: "AddComputed", pkg: "", typ: $funcType([$String, funcType$6, sliceType$5], [], true)}, {prop: "OnLifeCycleEvent", name: "OnLifeCycleEvent", pkg: "", typ: $funcType([LifeCycleEvent, funcType$8], [ptrType$7], false)}, {prop: "Mixin", name: "Mixin", pkg: "", typ: $funcType([js.M], [ptrType$7], false)}, {prop: "addMixin", name: "addMixin", pkg: "github.com/oskca/gopherjs-vue", typ: $funcType([$String, $emptyInterface], [ptrType$7], false)}, {prop: "AddSubComponent", name: "AddSubComponent", pkg: "", typ: $funcType([$String, ptrType$2], [ptrType$7], false)}, {prop: "AddProp", name: "AddProp", pkg: "", typ: $funcType([sliceType$1], [ptrType$7], true)}];
	ptrType$1.methods = [{prop: "FromJS", name: "FromJS", pkg: "", typ: $funcType([ptrType$3], [ptrType$1], false)}, {prop: "FromJSON", name: "FromJSON", pkg: "", typ: $funcType([$String], [ptrType$1], false)}, {prop: "ToJS", name: "ToJS", pkg: "", typ: $funcType([], [ptrType$3], false)}, {prop: "ToJSON", name: "ToJSON", pkg: "", typ: $funcType([], [$String], false)}, {prop: "Watch", name: "Watch", pkg: "", typ: $funcType([$String, funcType$3], [funcType$1], false)}];
	pool.init("github.com/oskca/gopherjs-vue", [{prop: "creator", name: "creator", anonymous: false, exported: false, typ: funcType, tag: ""}, {prop: "structPtr", name: "structPtr", anonymous: false, exported: false, typ: $emptyInterface, tag: ""}, {prop: "counter", name: "counter", anonymous: false, exported: false, typ: $Int, tag: ""}]);
	Component.init("", [{prop: "ViewModel", name: "ViewModel", anonymous: true, exported: true, typ: ptrType$1, tag: ""}]);
	TConfig.init("", [{prop: "Object", name: "Object", anonymous: true, exported: true, typ: ptrType$3, tag: ""}, {prop: "Silent", name: "Silent", anonymous: false, exported: true, typ: $Bool, tag: "js:\"silent\""}, {prop: "OptionMergeStrategies", name: "OptionMergeStrategies", anonymous: false, exported: true, typ: $emptyInterface, tag: "js:\"optionMergeStrategies\""}, {prop: "Devtools", name: "Devtools", anonymous: false, exported: true, typ: $Bool, tag: "js:\"devtools\""}, {prop: "ErrorHandler", name: "ErrorHandler", anonymous: false, exported: true, typ: funcType$4, tag: "js:\"errorHandler\""}, {prop: "IgnoredElements", name: "IgnoredElements", anonymous: false, exported: true, typ: sliceType$1, tag: "js:\"ignoredElements\""}, {prop: "KeyCodes", name: "KeyCodes", anonymous: false, exported: true, typ: mapType$2, tag: "js:\"keyCodes\""}]);
	Option.init("github.com/oskca/gopherjs-vue", [{prop: "Object", name: "Object", anonymous: true, exported: true, typ: ptrType$3, tag: ""}, {prop: "Name", name: "Name", anonymous: false, exported: true, typ: $String, tag: "js:\"name\""}, {prop: "Data", name: "Data", anonymous: false, exported: true, typ: $emptyInterface, tag: "js:\"data\""}, {prop: "El", name: "El", anonymous: false, exported: true, typ: $emptyInterface, tag: "js:\"el\""}, {prop: "Template", name: "Template", anonymous: false, exported: true, typ: $String, tag: "js:\"template\""}, {prop: "Parent", name: "Parent", anonymous: false, exported: true, typ: ptrType$3, tag: "js:\"parent\""}, {prop: "Delimiters", name: "Delimiters", anonymous: false, exported: true, typ: sliceType$1, tag: "js:\"delimiters\""}, {prop: "Functional", name: "Functional", anonymous: false, exported: true, typ: sliceType$1, tag: "js:\"functional\""}, {prop: "coms", name: "coms", anonymous: false, exported: false, typ: mapType, tag: ""}, {prop: "props", name: "props", anonymous: false, exported: false, typ: sliceType$1, tag: ""}, {prop: "mixins", name: "mixins", anonymous: false, exported: false, typ: sliceType$2, tag: ""}]);
	CreateElement.init([$String, $emptyInterface, sliceType$3], [ptrType$3], false);
	Render.init([ptrType$1, CreateElement], [], false);
	ViewModel.init("", [{prop: "Object", name: "Object", anonymous: true, exported: true, typ: ptrType$3, tag: ""}, {prop: "Data", name: "Data", anonymous: false, exported: true, typ: ptrType$3, tag: "js:\"$data\""}, {prop: "El", name: "El", anonymous: false, exported: true, typ: ptrType$3, tag: "js:\"$el\""}, {prop: "Options", name: "Options", anonymous: false, exported: true, typ: ptrType$3, tag: "js:\"$options\""}, {prop: "Parent", name: "Parent", anonymous: false, exported: true, typ: ptrType$3, tag: "js:\"$parent\""}, {prop: "Root", name: "Root", anonymous: false, exported: true, typ: ptrType$3, tag: "js:\"$root\""}, {prop: "Children", name: "Children", anonymous: false, exported: true, typ: ptrType$3, tag: "js:\"$children\""}, {prop: "Slots", name: "Slots", anonymous: false, exported: true, typ: ptrType$3, tag: "js:\"$slots\""}, {prop: "Refs", name: "Refs", anonymous: false, exported: true, typ: ptrType$3, tag: "js:\"$refs\""}, {prop: "IsServer", name: "IsServer", anonymous: false, exported: true, typ: $Bool, tag: "js:\"$isServer\""}, {prop: "WatchEx", name: "WatchEx", anonymous: false, exported: true, typ: funcType$9, tag: "js:\"$watch\""}, {prop: "Set", name: "Set", anonymous: false, exported: true, typ: funcType$10, tag: "js:\"$set\""}, {prop: "Delete", name: "Delete", anonymous: false, exported: true, typ: funcType$11, tag: "js:\"$delete\""}, {prop: "On", name: "On", anonymous: false, exported: true, typ: funcType$10, tag: "js:\"$on\""}, {prop: "Once", name: "Once", anonymous: false, exported: true, typ: funcType$10, tag: "js:\"$once\""}, {prop: "Off", name: "Off", anonymous: false, exported: true, typ: funcType$12, tag: "js:\"$off\""}, {prop: "Emit", name: "Emit", anonymous: false, exported: true, typ: funcType$13, tag: "js:\"$emit\""}, {prop: "Mount", name: "Mount", anonymous: false, exported: true, typ: funcType$14, tag: "js:\"$mount\""}, {prop: "ForceUpdate", name: "ForceUpdate", anonymous: false, exported: true, typ: funcType$1, tag: "js:\"$forceUpdate\""}, {prop: "NextTick", name: "NextTick", anonymous: false, exported: true, typ: funcType$15, tag: "js:\"$nextTick\""}, {prop: "Destroy", name: "Destroy", anonymous: false, exported: true, typ: funcType$16, tag: "js:\"$destroy\""}]);
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		$r = js.$init(); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = dom.$init(); /* */ $s = 2; case 2: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = json.$init(); /* */ $s = 3; case 3: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = minifiled.$init(); /* */ $s = 4; case 4: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = strings.$init(); /* */ $s = 5; case 5: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		creatorPool = $makeSlice(sliceType, 0);
		$pkg.Config = new TConfig.ptr(new ($global.Object)(), false, $ifaceNil, false, $throwNilPointerError, sliceType$1.nil, false);
		vue = $global.Vue;
		vMap = {};
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$packages["main"] = (function() {
	var $pkg = {}, $init, js, vue, Model, ptrType, ptrType$1, main;
	js = $packages["github.com/gopherjs/gopherjs/js"];
	vue = $packages["github.com/oskca/gopherjs-vue"];
	Model = $pkg.Model = $newType(0, $kindStruct, "main.Model", true, "main", true, function(Object_, IntValue_, Str_) {
		this.$val = this;
		if (arguments.length === 0) {
			this.Object = null;
			this.IntValue = 0;
			this.Str = "";
			return;
		}
		this.Object = Object_;
		this.IntValue = IntValue_;
		this.Str = Str_;
	});
	ptrType = $ptrType(Model);
	ptrType$1 = $ptrType(js.Object);
	Model.ptr.prototype.Inc = function() {
		var $ptr, m;
		m = this;
		m.Object.integer = ($parseInt(m.Object.integer) >> 0) + (1) >> 0;
		console.log("inc called");
	};
	Model.prototype.Inc = function() { return this.$val.Inc(); };
	Model.ptr.prototype.Repeat = function() {
		var $ptr, m;
		m = this;
		m.Object.str = $externalize($internalize(m.Object.str, $String) + $internalize(m.Object.str, $String), $String);
	};
	Model.prototype.Repeat = function() { return this.$val.Repeat(); };
	Model.ptr.prototype.Reset = function() {
		var $ptr, m;
		m = this;
		m.Object.str = $externalize("a string ", $String);
	};
	Model.prototype.Reset = function() { return this.$val.Reset(); };
	main = function() {
		var $ptr, m;
		m = new Model.ptr(new ($global.Object)(), 0, "");
		m.Object.integer = 100;
		m.Object.str = $externalize("a string", $String);
		vue.New(new $String("#app"), m);
	};
	ptrType.methods = [{prop: "Inc", name: "Inc", pkg: "", typ: $funcType([], [], false)}, {prop: "Repeat", name: "Repeat", pkg: "", typ: $funcType([], [], false)}, {prop: "Reset", name: "Reset", pkg: "", typ: $funcType([], [], false)}];
	Model.init("", [{prop: "Object", name: "Object", anonymous: true, exported: true, typ: ptrType$1, tag: ""}, {prop: "IntValue", name: "IntValue", anonymous: false, exported: true, typ: $Int, tag: "js:\"integer\""}, {prop: "Str", name: "Str", anonymous: false, exported: true, typ: $String, tag: "js:\"str\""}]);
	$init = function() {
		$pkg.$init = function() {};
		/* */ var $f, $c = false, $s = 0, $r; if (this !== undefined && this.$blk !== undefined) { $f = this; $c = true; $s = $f.$s; $r = $f.$r; } s: while (true) { switch ($s) { case 0:
		$r = js.$init(); /* */ $s = 1; case 1: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		$r = vue.$init(); /* */ $s = 2; case 2: if($c) { $c = false; $r = $r.$blk(); } if ($r && $r.$blk !== undefined) { break s; }
		if ($pkg === $mainPkg) {
			main();
			$mainFinished = true;
		}
		/* */ } return; } if ($f === undefined) { $f = { $blk: $init }; } $f.$s = $s; $f.$r = $r; return $f;
	};
	$pkg.$init = $init;
	return $pkg;
})();
$synthesizeMethods();
var $mainPkg = $packages["main"];
$packages["runtime"].$init();
$go($mainPkg.$init, []);
$flushConsole();

}).call(this);
//# sourceMappingURL=main.js.map
