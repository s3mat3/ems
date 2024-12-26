/**
 * @file state.js
 *
 * @copyright © 2024 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @brief
 *
 * @author s3mat3
 */
'use strict';

import Notification from "./notification";
export class State {
    /**
     * @type {Object | Array}
     */
    $ = {};
    /**
     * @type {class Notification}
     */
    events = null;
    /**
     * 
     */
    constructor(initial_value = null, type_name = 'state') {
        const self = this;
        const obj = (initial_value) ? initial_value : {};
        this.$ = {};
        this.events = new Notification();
        this.$ = new Proxy(obj, {
            set(target, key, newValue) {
                if (newValue === target[key]) return true; /// nothig cause same data
                target[key] = newValue;
                // console.log({'target': target, 'key': key, 'new': newValue});
                self.events.notify(type_name, self);
                return true;
            },
            get(target, key) {
                return target[key];
            },
            deleteProperty(target, key) {
                delete target[key];
                self.events.notify(`${type_name}_delete`, self);
                return true;
            },
        }); /* //<-- Proxy construct ends here */
    } /* //<-- constructor ends here */
    /**
     * 
     */
    addUpdateAction(type = 'state', callback, context = this) {
        console.log(callback);
        this.events.connect(type, callback);
    }
    addDeleteAction(type= 'state', callback) {
        this.events.connect(`${type}_delete`, callback, context = this);
    }
} /* //<-- class State ends here */
/**
 * Helper class State constructor
 *  @param {Object | Array} initialValue is state valiable object
 *  @param {Function} updater is callback when state update  Default is null
 *  @param {Function} deleter is callback when state delete  Default is null
 *  @param {Object} context of execute default this
 *  @returns {Object proxy} State object
 */
export function useState(initialValue, type_name = 'state',updateAction = null, deleteAction = null, context = this) {
    const s = new State(initialValue, type_name);
    (updateAction && typeof updateAction == 'function') ? s.addUpdateAction(type_name, updateAction, context) : '';
    (deleteAction && typeof deleteAction == 'function') ? s.addDeleteAction(type_name, deleteAction, context) : '';
    return s;
} /* //<-- function useState ends here */


/* //<-- state.js ends here*/
