/**
 * @file message-modal.js
 *
 * @copyright © 2024 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @brief
 *
 * @author s3mat3
 */
'use strict';
import {Component, tag} from '../../core/component'
// import { isElement } from '../../core/util';

export default class MessageDialog extends Component {
    #base = null;
    props = {
        dialog_class: '',
        title_class: '',
        content_class: '',
        title: '',
        message: '',
    };

    constructor(container = null, props = {}) {
        super(container);
        this.props.title = (props.title) ? props.title : 'No title';
        this.props.message = (props.message) ? props.message : '';
        this.props.dialog_class = (props.dialog_class) ? props.dialog_class : '';
        this.props.title_class = (props.title_class) ? props.title_class : '';
        this.props.content_class = (props.content_class) ? props.content_class : '';
        this.buildBase();
    }

    open() {
        this.mount();
        this.#base.querySelector('.message-dialog').showModal();
    }

    close() {
        const d = this.#base.querySelector('.message-dialog');
        if (d) {
            d.close();
        }
    }

    buildBase() {
        let dialog = 'message-dialog';
        dialog += (this.props.dialog_class) ? ` ${this.props.dialog_class}` : '';
        let title = 'message-dialog-title is-center';
        title += (this.props.title_class) ? ` ${this.props.title_class}` : '';
        let content = 'message-dialog-content';
        content += (this.props.content_class) ? ` ${this.props.content_class}` : '';
        this.#base = tag.element`
<div>
  <dialog class="${dialog}">
    <div class="${title}"></div>
    <div class="${content}">
      <div class="message-dialog-body row">
        <div class="icon col-3 is-center"></div>
        <div class="message col-9"></div>
      </div>
      <div class="message-dialog-command row">
        <div class="col-4 is-center pos-left">
          <button class="message-dialog-ok is-normal">
            <span>OK</span>
          </button>
        </div>
        <div class="col-4 is-center pos-center">
        </div>
        <div class="col-4 is-center pos-right">
          <button class="message-dialog-ok is-normal">
            <span class="icon">close</span>
             <span>close</span>
          </button>
        </div>
      </div>
    </div>
  </dialog>
</div>`;
        return this.#base;
    }

    mountTitle() {
        const container = this.#base.querySelector(".message-dialog-title");
        container.innerText = this.props.title;
        return this.#base;
    }
    mountBody() {
        let container = this.#base.querySelector(".message-dialog-body > .message");
        container.innerText = this.props.message;
        // container = this.#base.querySelector(".message-dialog-body > .icon");
        // container.innerText = 'hoge';
        return this.#base;
    }
    mountButton() {
        const elm = this.#base.querySelector(".message-dialog-ok");
        elm.addEventListener('click', () => {
            this.close();
        }, false);
        // const container = this.#base.querySelector(".message-dialog-command > .pos-center");
        // if (this.props.command) {
        // } else {
        //     container.insertAdjacentHtml('<button>');
        // }
        // return this.#base;
    }

    buildElement() {
        this.mountTitle();
        this.mountBody();
        this.mountButton();
        return this.element = this.#base;
    }
}
/* //<-- message-modal.js ends here*/
