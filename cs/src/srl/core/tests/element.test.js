/**
 * @file util-dom.test.js
 *
 * @copyright © 2024 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @brief
 *
 * @author s3mat3
 */
import { test, expect, describe, beforeAll } from 'vitest';
import { JSDOM } from "jsdom";
import '@testing-library/jest-dom';

import {isElement, escapeHtmlSpecialChars, element, buildElement} from '../element'

describe('util test in with', () => {
    let browser = null;
    let window;
    let document;
    const content_simple = 'Hello world';
    const content_escape = '"&<>"';
    beforeAll(() => {
        browser = new JSDOM();
        window = browser.window;
        document = window.document;
    });
    test('isElement', () => {
        const t1 = document.createElement('div');
        const t2 = document.createElement('template');
        const o = {};
        const a = [];
        const s = 'div';
        expect(isElement(t1)).toBeTruthy();
        expect(isElement(t2)).toBeTruthy();
        expect(isElement(o)).toBeFalsy();
        expect(isElement(a)).toBeFalsy();
        expect(isElement(s)).toBeFalsy();
    });
    // test('build element simple', () => {
    //     expect(element`<div>${content_simple}</div>`).toBe('<div>Hello world</div>');
    // });
    // test('build element escaped', () => {
    //     const elm = element`<div>${content_escape}</div>`;
    //     console.log(typeof elm.innerHTML);
    //     expect(elm.innerHtml).toBe('&amp;&lt;&gt;');
    // });
    test('build element', () => {
        let parent = browser.window.document.createElement('div');
        let elm = buildElement(`<p>${content_simple}</p>`);
        expect(isElement(elm)).toBeTruthy();
        expect(elm.innerHTML).toBe('Hello world');
        elm = buildElement(`<p>${content_simple}</p>`);
        expect(isElement(elm)).toBeTruthy();
        expect(elm.innerHTML).toBe('Hello world');
    });
});

describe('escapeHtmlSpecialChars', () => {
    test('simple string', () => {
        const mesg = 'Hello world';
        expect(escapeHtmlSpecialChars(mesg)).toBe('Hello world');
    });
    test('escape quote', () => {
        const mesg = '"Hello world"';
        const escaped = escapeHtmlSpecialChars(mesg);
        expect(escaped).toContain('&quot;');
        expect(escaped).toBe('&quot;Hello world&quot;');
    })
    test('escape single quote', () => {
        const mesg = "'Hello world'";
        const escaped = escapeHtmlSpecialChars(mesg);
        expect(escaped).toContain('&#x27;');
        expect(escaped).toBe('&#x27;Hello world&#x27;');
    })
    test('escape &', () => {
        const mesg = "Hello & world";
        const escaped = escapeHtmlSpecialChars(mesg);
        expect(escaped).toContain('&amp;');
        expect(escaped).toBe('Hello &amp; world');
    })
    test('escape <>', () => {
        const mesg = '<Hello world>';
        const escaped = escapeHtmlSpecialChars(mesg);
        expect(escaped).toContain('&lt;');
        expect(escaped).toContain('&gt;');
        expect(escaped).toBe('&lt;Hello world&gt;');
    })
});

/* //<-- util-dom.test.js ends here*/
