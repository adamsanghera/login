# Login µService

## Interface-Based Design

In `login.go`, we see the basic interfaces, that are implemented in this package or others.

We have 3 core interfaces, whose functionality is then exposed by a centralized `login` interface.

The core interfaces are `Category`, `Authenticator`, and `Badge`.

## Useful Analogy for Understanding the Interfaces

A real-world phenomenon that maps well to these interfaces is the concept of Citizenship, Passports and Identification.

In this model, `Category` is the Nation, `Badge` is the Passport, and `Authenticator` is a method of confirming existence and association with some identity (like a name, which is associated with a password, photograph, etc.).

You can declare yourself as existing, apply for citizenship, then apply for a passport – exclusively in that order.

Existing gives you the right to apply for citizeship; citizenship gives you the right to apply for a passport.

Passports are temporary documents, and can be revoked at any time for misbehavior.
Passports can also be renewed before they are expired.
Passports can be forged!! It is important to guard against this, with the use of an Authenticator.

Nations can be joined or abandoned.
One can be a citizen of numerous nations.
Every Nation issues exactly 1 kind of Passport (this might change later).

# Roadmap

## Done
1. Get a basic registration/provisioning/validation system set up.  No fancy hashing.
2. Hash at the server level.  No plaintext passwords are persisted.

## Up and coming
3. Replace the current mumbo jumbo implementations with well-codified interfaces and implementations
  1. Create the interfaces
  2. Create the implementations
4. Allow Categories to issue multiple forms of Badges (R, W, X, RX, WX, RW, RWX).
5. 