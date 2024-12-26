/**
 * @file util.js
 *
 * @copyright © 2024 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @brief
 *
 * @author s3mat3
 */
'use strict';

/**
 * create space separate string from comma separeted arguments
 *  @param {any} 
 *  @returns {String} space separated  
 */
export function toStringFromArgs(...args) {return args.filter((arg) => arg != undefined).join(" ");}
/** 
 * Extract type name came from string made by object.tostring.call()
 * let arr = [];
 * Object.call(arr) result => [object Array] => slice(8, -1) => 'Array' 
 *                            12345678     -1
 * @param {any} target is Type name extraction target
 * @return {String} cf Array Object ... etc
 */
export function extractTypeName(target) {return Object.prototype.toString.call(target).slice(8, -1);}
/**
 * Test if the target type is in the given type list.
 * @param {array} types is list(array) of types to match. cf.['Array', 'String', ...] or ['Object'] etc.
 * @param {any} target is inspect target
 * @return {bool} true match types and target, false unmatch types and target.
 */
export function isMatchType(types, target) {return types.includes(extractTypeName(target));}
/** 
 * Object type cheker.
 * @param {HTMLElement} obj
 * @return {boolean} true: HTMLElement, false: Not HTMLElement 
 */
export function isElement(obj) {
    return (obj && typeof obj === 'object') ? (obj.nodeType === 1) : false;
}
/* //<-- util.js ends here*/
