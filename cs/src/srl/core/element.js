/**
 * @file element.js
 *
 * @copyright © 2024 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @brief HTML tag element utilities
 *
 * @author s3mat3
 */
'use strict';

import {isElement, isMatchType} from './util'; 
/**
 * Escape HTML reserved special characters
 * in case 5 characters " ' & < >
 * Should use the entity name or entity number when you want to output any of these reserved characters.
 *  @see https://www.html.am/reference/html-special-characters.cfm
 *
 *  @param {String} target is string for output web browser
 */
function escapeHtmlSpecialChars(target) {
    return target.replace(/[\'\"&<>]/g, (match) => {
        const escapes = {
            '\'': '&#x27;',
            '\"': '&quot;',
            '&': '&amp;',
            '<': '&lt;',
            '>': '&gt;',
        };
        return escapes[match];
  });
}
/**
 * Update attribute(s) for target element
 *  usage
 *  const attrs = {className: 'button class-button', id: 'button_001'};
 *  updateAttributes(buttonElement, attrs);
 *  or updateAttributes(buttonElement, {className: 'button class-button', id: 'button_001'});
 *
 *  @param {HTMLElement} element is update target
 *  @param {Map | Object <String, (Value | String)>} HTML attribute(s)
 */
function updateAttributes(element, attrs) {
    if (element && isElement(element) && isMatchType(['Map', 'Object'], attrs)) {
        for (const [k, value] of Object.entries(attrs)) {
            const name = (k && k === 'className') ? 'class' : k;
            element.setAttribute(name, value);
        }
    }
}
/**
 * Build element from HTML markup string
 *  @param {String} html markup string
 *  @returns {HTMLElement} 
 */
function buildElement(html) {
    const elm = document.createElement('template');
    elm.insertAdjacentHTML('afterbegin', html);
    return elm.firstElementChild;
}
/**
 * Building node elements
 *  @returns {HTMLElement}  
 */
function buildNode(element, ...children) {
    // console.log(children);
    if (children.length > 0) {
        const fragment = document.createDocumentFragment();
        for (const child of children) {
            if (isElement(child)) {
                fragment.appendChild(child);
            }
        }
        element.appendChild(fragment);
    }
    return element;
} /* //<-- function buildNode ends here */
/**
 * Build html element from template literal. (Using taged template literal)
 * @caution All placeholder variables included html special chars are escaped
 * @usage element`<html-tag>${some_variable}</html-tag>`;
 */
function element(strings, ...values) {
    const html = strings.reduce((acc, current, index) => {
        const value = values[index - 1];
        return (typeof value === 'string') ? acc + escapeHtmlSpecialChars(value) + current
            : acc + String(value) + current;
    });
    return buildElement(html);
} /* //<-- function element ends here */
/**
 * Target element insert into destination element
 *  @param {HTMLElement} destination to inserted
 *  @param {HTMLElement} target of insert
 *  @returns {HTMLElement} inserted
 */
function insertElement(destination, target) {
    if (! isElement(destination)) throw new TypeError('destination neccesary instance of HTMLElement');
    if (isElement(target)) {
        destination.insertAdjacentElement('afterbegin', target);
    } else if (typeof target === 'string') {
        destination.insertAdjacentHTML('afterbegin', target);
    } // else {} no action (no effect)
    return destination;
} /* //<-- function insertElement ends here */
/**
 * Target element insert into destination element
 *  @param {HTMLElement} destination to inserted
 *  @param {HTMLElement} target of insert
 *  @returns {HTMLElement} inserted
 */
function appendElement(destination, target) {
    if (! isElement(destination)) throw new TypeError('destination neccesary instance of HTMLElement');
    if (isElement(target)) {
        destination.insertAdjacentElement('beforeend', target);
    } else if (typeof target === 'string') destination.insertAdjacentHTML('beforeend', target);
    // else {} no action (no effect)
    return destination;
} /* //<-- function insertElement ends here */
/**
 * Attach element to parent element
 *  @param { HTMLElement } parent to attached
 *  @param { HTMLElement } element is taget of attach
 *  @returns { HTMLElement } attached
 */
function attach(parent, element) {
    if (parent && isElement(parent)) {
        (element) ? parent.replaceChildren(element) : '';
        return parent;
    } else {
        const temp = document.createElement('template');
        temp.insertAdjacentElement('afterbegin', element);
        return insertElement(document.querySelector('body'), element);
    }
} /* //<-- function attach ends here */

export {
    attach,
    buildNode,
    appendElement,
    insertElement,
    element,
    buildElement,
    updateAttributes,
    escapeHtmlSpecialChars,
    isElement,
    isMatchType,
};
/* //<-- element.js ends here*/
