// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package qqq provides easy printing, logging and panicing to any package.
//
// It is intentionally minimalistic - part of a "poor mans survival kit".
//
// `qqq.log.go` provides easy-to-use and easy-to-remember functions:
//  qqq - print to log iff verbose
//  see - print to log
//  die - panic to log
//
// Each function comes in one more flavour:
//  xxf - formated print
//
// `qqq.set.go` exports functions to control the package-wide verbosity.
// This allows other packages to control the verbosity of imported packages.
//
// `qqq.main.init`.go has a super-simple init() function overriding defaults
// of the standard package. Use in /cmd main() packages only!
//
// Note: Intentionally, log.Fatal is not included.
// IMO its use in packages is bad practice.
//
// The idea and the use of qqq as name were inspired by a talk I once listened to.
// TODO: Find and supply reference to the original author.
//
// Last, but not least: the ever-useful constants
//  `tab` = "\t" for TABulation and
//  `eol` = "\n" for End-Of-Line
//
// Be a happy gopher :-)
package qqq
