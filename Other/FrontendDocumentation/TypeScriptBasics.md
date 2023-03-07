# TypeScript Documentationn

# Basics
- Just JavaScript but with types(if you know JavaScript that helps I guess)
- We want to avoid bugs in our code so by having a way to check as we code helps us stop them before they happen
- Types in TypeScript are more for structure rather than exclusive objects

```typescript
class Car {
  drive() {
    // hit the gas
  }
}
class Golfer {
  drive() {
    // hit the ball far
  }
}
// No error?
let w: Car = new Golfer();
```
- tell me this doesn't look super weird, but due to the classes being the same TypeScript doesn't see a problem with assigning them

## Types
### Primitive Types
- **string** => strings
- **number** => not int or float, everything is number
- **boolean** => true and false(duh!)

### Arrays
```typescript
number[] = [1, 2, 3]
```
- can also be written as Array< number >
### any
- TypeScript exclusive(i love gatekeeping) type that you can use whenever you wanna avoid errors
- you can do anything "syntactically legal"
```typescript
let obj: any = { x: 0 };
// None of the following lines of code will throw compiler errors.
// Using `any` disables all further type checking, and it is assumed 
// you know the environment better than TypeScript.
obj.foo();
obj();
obj.bar = 100;
obj = "hello";
const n: number = obj;
```

### Type Annotations
- you can specify the type of a variable when using let, var, const
```typescript
let myName: string = "Alice";
```
- you can also do and it is able to infer the type
```typescript
let myName = "Alice";
```

### Functions
- you can specify the types of input and output variables
- when you declare a function you can add type annotations like when using let, var, const
```typescript
function greet(name: string) {
  console.log("Hello, " + name.toUpperCase() + "!!");
}
```
- you can also add the return type annotations
```typescript
function getFavoriteNumber(): number {
  return 26;
}
```
- i dont understand anonymous functions but this is a placeholder
### Object Types
- to define an object type you lists its properities and types
- this defines an object of type `pt` that has `x` and `y` properties with types `number`
- it gives you control over the typing of each step
```typescript
function printCoord(pt: { x: number; y: number }) {
  console.log("The coordinate's x value is " + pt.x);
  console.log("The coordinate's y value is " + pt.y);
}
```
### Optional Properties
- you don't have to pass in the number of properties a function is asking for if you tell it to be optional
```
export function twoFer(name?:string): string {
  let result:string = "One for you, one for me.";
  if(name)
  {
    result = "One for " + name + ", one for me.";
  }
return result;
}
```
### Union Types
- these look absolutely wild
- type that is formed from two or more other types, representing values that may be any of tose types
- the types in the unions types are called `members`
```typescript
function printId(id: number | string) {
  console.log("Your ID is: " + id);
}
// OK
printId(101);
// OK
printId("202");
// Error
printId({ myID: 22342 });
```
- to work with these you cannot call something that would only work on one type
- you can create `branches` in if statements where you can ensure the type and then call functions appropriately
```typescript
function printId(id: number | string) {
  if (typeof id === "string") {
    // In this branch, id is of type 'string'
    console.log(id.toUpperCase());
  } else {
    // Here, id is of type 'number'
    console.log(id);
  }
}
```
### Type Aliases
- a name for any type, basically typedef in c and just defining your own type
```typescript
type Point = {
  x: number;
  y: number;
};
 
// Exactly the same as the earlier example
function printCoord(pt: Point) {
  console.log("The coordinate's x value is " + pt.x);
  console.log("The coordinate's y value is " + pt.y);
}
```
- a type alias can name a union type
```typescript
type ID = number | string;
```
### Interfaces
- interface is another way to name object types
- seems basically the same as type alias

### Type Assertions
continue here