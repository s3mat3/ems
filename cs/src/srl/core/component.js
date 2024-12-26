/**
 * @file component.js
 *
 * @copyright © 2024 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @brief Component abstract class
 *
 * @author s3mat3
 */
'use strict';
import * as tag from './element';
import Notification from './notification';
/**
 * Component abstract class
 */
export default class Component {
    /**
     * Event notify variable for occured in ui parts
     * @type {Notification}
     */
    uiEvent = null;
    /**
     * Holder for this component element
     * @type {HTMLElement}
     */
    element = null;
    /**
     * Container for this component element
     * @type {HTMLElement}
     */
    container = null;
    constructor(container = null) {
        this.uiEvent = new Notification();
        this.container = container;
    }
    // getter
    /**
     * Get this element
     */
    get component() {return this.element = this.buildElement();}
    /**
     * Append child element to this element node (push front)
     */
    appendElement(child) {
        return (child instanceof Component) ? tag.appendElement(this.element, child.component)
            : tag.appendElement(this.element, child);
    }
    /**
     * Insert child element to this element node (push back)
     */
    insertElement(child) {
        return (child instanceof Component) ? tag.insertElement(this.element, child.component)
            : tag.insertElement(this.element, child);
    }
    /**
     * Build this element
     * @caution This is virtual method must be imprimentation in derived
     */
    buildElement() {console.error('This is virtual method must be imprimentation in derived');}
    /**
     * Hook for before mount
     */
    beforeMount() {}
    /**
     * Hook for after mount
     */
    mounted() {}
    /**
     * Mount(attach) this element node to parent node
     */
    mount() {
        this.element = this.buildElement();
        this.beforeMount();
        (this.container) ? tag.attach(this.container, this.element) : document.querySelector('body').insertAdjacentElement('afterbegin', this.element);
        this.mounted();
        return this.element;
    }
    /**
     * Hook for before unmount
     */
    beforeUnmount() {}
    /**
     * Hook for after unmount
     */
    unmounted() {}
    /**
     * Unmount(detach) from parent container
     */
    unmount() {
        this.beforeUnmount();
        (this.container) ? this.container.replaceChildren() : this.element.innerText = '';
        this.unmounted();
    }
}

export {
    tag,
    Component
}
/* //<-- component.js ends here*/
