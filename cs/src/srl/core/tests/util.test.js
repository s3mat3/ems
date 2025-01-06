/**
 * @file util.test.js
 *
 * @copyright © 2024 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @brief
 *
 * @author s3mat3
 */

import { test, expect, describe, beforeAll } from 'vitest';
import {isElement, isMatchType, toStringFromArgs, attribute} from '../util';

describe('test typename matcher', () => {
    test('match number', () =>{
        let i = 0;
        expect(isMatchType(['Number'], i)).toBeTruthy();
        i = 1.0e10;
        expect(isMatchType(['Number'], i)).toBeTruthy();
        i = NaN;
        expect(isMatchType(['Number'], i)).toBeTruthy(); // caution NaN is Number when critical number use first check isNaN
        i = BigInt(0);
        expect(isMatchType(['Number'], i)).toBeFalsy();
        i = {};
        expect(isMatchType(['Number'], i)).toBeFalsy();
        i = [];
        expect(isMatchType(['Number'], i)).toBeFalsy();
        i = new Set();
        expect(isMatchType(['Number'], i)).toBeFalsy();
        i = new Map();
        expect(isMatchType(['Number'], i)).toBeFalsy();
    });
    test('match Object', () =>{
        let i = {};
        expect(isMatchType(['Object'], i)).toBeTruthy();
        i = new Object();
        expect(isMatchType(['Object'], i)).toBeTruthy();
        i = 1.0e10;
        expect(isMatchType(['Object'], i)).toBeFalsy();
        i = BigInt(0);
        expect(isMatchType(['Object'], i)).toBeFalsy();
        i = [];
        expect(isMatchType(['Object'], i)).toBeFalsy();
        i = new Set();
        expect(isMatchType(['Object'], i)).toBeFalsy();
        i = new Map();
        expect(isMatchType(['Object'], i)).toBeFalsy();
    });
    test('match array', () => {
        let i = [];
        expect(isMatchType(['Array'], i)).toBeTruthy();
        i = new Array();
        expect(isMatchType(['Array'], i)).toBeTruthy();
        i = 1.0e10;
        expect(isMatchType(['Array'], i)).toBeFalsy();
        i = BigInt(0);
        expect(isMatchType(['Array'], i)).toBeFalsy();
        i = {};
        expect(isMatchType(['Array'], i)).toBeFalsy();
        i = new Set();
        expect(isMatchType(['Array'], i)).toBeFalsy();
        i = new Map();
        expect(isMatchType(['Array'], i)).toBeFalsy();
    });
    test('match Map', () => {
        let i = new Map();
        expect(isMatchType(['Map'], i)).toBeTruthy();
        i = 1.0e10;
        expect(isMatchType(['Map'], i)).toBeFalsy();
        i = BigInt(0);
        expect(isMatchType(['Map'], i)).toBeFalsy();
        i = {};
        expect(isMatchType(['Map'], i)).toBeFalsy();
        i = new Set();
        expect(isMatchType(['Map'], i)).toBeFalsy();
        i = [];
        expect(isMatchType(['Map'], i)).toBeFalsy();
    });
    test('match Set', () => {
        let i = new Set();
        expect(isMatchType(['Set'], i)).toBeTruthy();
        i = 1.0e10;
        expect(isMatchType(['Set'], i)).toBeFalsy();
        i = BigInt(0);
        expect(isMatchType(['Set'], i)).toBeFalsy();
        i = {};
        expect(isMatchType(['Set'], i)).toBeFalsy();
        i = new Map();
        expect(isMatchType(['Set'], i)).toBeFalsy();
        i = [];
        expect(isMatchType(['Set'], i)).toBeFalsy();
    });
});

describe('test element checker', () => {
    test('Fake test', () => {
        const o = {};
        const a = [];
        const s = new Set();
        const m = new Map();
        expect(isElement(o)).toBeFalsy();
        expect(isElement(a)).toBeFalsy();
        expect(isElement(s)).toBeFalsy();
        expect(isElement(m)).toBeFalsy();
        let obj = {};
        obj.nodeType = 1;
        expect(isElement(obj)).toBeTruthy(); // this test must be under DOM environment
    });
});
/* //<-- util.test.js ends here*/
