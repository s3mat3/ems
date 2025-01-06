/**
 * @file notification.test.js
 *
 * @copyright © 2024 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @brief
 *
 * @author s3mat3
 */
import { test, expect, describe, beforeAll } from 'vitest';
import Notification from '../notification.js';

let x_result;
let y_result;

function x(data = {}) {
    data.name = 'FUNCTION X';
    x_result =  data;
}

function y(data = {}) {
    data.name = 'FUNCTION Y';
    y_result = data;
}
// let y = (data = {}) => {data.func = 'FUNCTION Y'; y_result = data;}

describe('Notification test', () => {
    let notification;
    beforeAll (() => {
        notification = new Notification();
    });
    test('connectable', () =>{
        expect(notification.connect('a', x)).toBe(1);
        expect(notification.connect('a', y)).toBe(2);
        // console.log(notification.lists.a.includes(y));
    }) ;
    test('connect reject same', () =>{
        expect(notification.connect('a', x)).toBe(2);
        expect(notification.connect('a', y)).toBe(2);
        expect(notification.connect('aa', x)).toBe(1);
        expect(notification.connect('aa', y)).toBe(2);
        // console.log(notification.lists.a.includes(y));
    }) ;
    test('disconnectable', () =>{
        expect(notification.size('a')).toBe(2);
        notification.disconnect('a', x);
        expect(notification.size('a')).toBe(1);
        notification.disconnect('a', y);
        expect(notification.size('a')).toBe(0);
        notification.disconnect('x', y);
        expect(notification.size('x')).toBe(0);
        expect(notification.size('aa')).toBe(2);
    }) ;
    test('callback not a function in connect', () => {
        expect(() => notification.connect('aaa', 1).toThrow);
        expect(() => notification.connect('aaa', []).toThrow);
        expect(() => notification.connect('aaa', {}).toThrow);
    });
    test('callback not a function in disconnect', () => {
        expect(() => notification.disconnect('a', 1).toThrow);
        expect(() => notification.disconnect('a', []).toThrow);
        expect(() => notification.disconnect('a', {}).toThrow);
    });
    test('event name not available in disconnect', () => {
        expect(notification.size('aa')).toBe(2);
        notification.disconnect('aa', x);
        expect(notification.size('aa')).toBe(1);
        notification.disconnect('aa', y);
        expect(notification.size('aa')).toBe(0);
    });
    test('notifyable', () => {
        notification.connect('aa', x);
        notification.connect('aa', y);
        expect(x_result).toBeUndefined();
        expect(y_result).toBeUndefined();
        notification.notify('aa');
        expect(x_result.name).toBe('FUNCTION X');
        expect(y_result.name).toBe('FUNCTION Y');
        // expect(result).not.toBe([]);
        // expect(result).toMatchObject(detail);
        // for (let i in result) {
        //     expect(result[i].id).toBe(1);
        //     expect(result[i].name).toBe('hoge');
        //     // console.log(result[i]);
        // }
    });
})
/* //<-- notification.test.js ends here*/
