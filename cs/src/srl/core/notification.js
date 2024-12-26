/**
 * @file notification.js
 *
 * @copyright © 2024 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @brief
 *
 * @author s3mat3
 */

export default class Notification {
/**
 * listenrs
 * @type {Map}
 */
#listeners;
    /**
     *  initialize
     */
    constructor() {
        this.#listeners = new Map();
    }
    // for debug
    get lists() {return this.#listeners;}
    /**
     * Connect listener's callback
     * ** CAUTION ** callback function can not entry anonymous function
     *
     *  @param {String} key is event name(or type) for waiting listener
     *  @param {Function} listener is callback function as to recive event notice
     *  @retern {Number} listener's size on key
     *  @param {Object} context of execute default this
     *  @throw TypeError when listener is not a function
     */
    connect(key, listener, context = this) {
        if (typeof listener !== 'function') throw new TypeError('Not a function, nesseary listener is function', 'notification.js');
        const obj = {ctx: context, cb: listener};
        if (! this.#listeners.has(key)) this.#listeners.set(key, new Set());
        let f = false;
        this.#listeners.get(key).forEach((o) => {
            if (o.cb === listener) f = true;
        });
        if (! f) this.#listeners.get(key).add(obj);
        return this.#listeners.get(key).size
    }
    /**
     * Disconnect listener's callback
     * ** CAUTION ** callback function can not entry anonymous function
     *
     *  @param {String} key is event name(type) for waiting listener
     *  @param {Function} callback for listenr
     */
    disconnect(key, listener) {
        if (typeof listener !== 'function') throw new TypeError('Not a function, nesseary listener is function', 'notification.js');
        const listeners = this.#listeners.get(key);
        if (! listeners) return;
        listeners.forEach((obj) => {
            if (obj.cb === listener) {
                listeners.delete(obj);
            }
        })
    }
    /**
     * Execute notify for listener all waiting
     *
     *  @param {String} key is event name(type) for waiting listener
     *  @param {Object} detail of Notification
     */
    notify(key, obj) {
        const listeners = this.#listeners.get(key);
        if (! listeners) {
            return;
        }
        listeners.forEach((listener) => {
            listener.cb.call(listener.ctx,obj);
        });
    }

    size(key) {
        const listeners = this.#listeners.get(key);
        if (! listeners) return 0;
        return listeners.size;
    }
}

/* //<-- notification.js ends here*/
