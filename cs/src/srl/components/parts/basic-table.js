/**
 * @file basic-table.js
 *
 * @copyright © 2024 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @brief Simple table
 *
 * This table has sort invoke event, when row single click to change rows color and when double click to fire event.
 * @author s3mat3
 */
// Slopy reactive system
'use strict';
import {Component, tag} from '../../core/component';

export default class BasicTable extends Component {
    folding = false;
    /**
     * Table header data [Object1, Object2, Object3,...,ObjectM];
     * Oject = {name: 'xxxx', sortable: boolean};
     * @type {Array[Object]}
     */
    #header = [];
    /**
     * Table row data [rowData1,rowData2,...rowDataM];
     * rowData = [col1, col2, col3,...,colN];
     * @type {Array[Array]}
     */
    #rows = [];
    dataInfo = {total: 0, page_offset: 25, page: 0};
    props = {
        name: '',
    }
    /**
     * Table row element when clicked row
     * @type {HTMLElement} (table.tbody.tr)
     */
    #selected = null;
    /**
     * Table row position when clicked row
     * @type {Number} (table.tbody.tr.rowIndex)
     */
    #index = 0;
    /**
     * Table frame element
     * @type {HTMLElement}
     */
    #base = null;
    constructor(container = null, props = {}) {
        super(container);
        this.props.name = (props.name) ? props.name : 'No name given';
        this.buildBase(); // create base element when construct this object
    }
    // setters
    set rows(rowData) {this.#rows = rowData;}
    set header(header) {this.#header = header;}
    // geter
    get rows() {return this.#rows;}
    get header() {return this.#header;}
    get selected() {return this.#selected;}
    /**
     * Building table base element
     *  @returns {HTMLEleement}
     */
    buildBase() {
        this.#base = tag.element`
      <table class="basic-table noselect">
        <caption class="basic-table-caption noselect sticky"></caption>
        <thead class="basic-table-header">
        </thead>
        <tbody class="basic-table-body">
        </tbody>
      </table>`;
        return this.#base;
    }
    /**
     * Table caption mount on base element
     */
    mountCaption() {
        const container = this.#base.querySelector('.basic-table-caption');
        container.innerText = this.props.name;
        container.addEventListener('dblclick', () => {this.hideBody();});
        return this.#base;
    }
    /**
     * Table header mount on base element
     */
    mountHeader() {
        const container = this.#base.querySelector('.basic-table-header');
        container.replaceChildren();
        container.insertAdjacentHTML('afterbegin', '<tr class="basic-table-header">');
        const tr = container.querySelector('.basic-table-header');
        this.#header.forEach((v) => {
            let cls = (v.sortable) ? ' class="sort-desc"': '';
            let el = tag.buildElement(`<th${cls}>${tag.escapeHtmlSpecialChars(v.name)}</th>`);
            el.addEventListener('click', (e) => {
                let dir = '';
                if (el.classList.contains('sort-desc')) {
                    el.classList.remove('sort-desc');
                    el.classList.add('sort-asc');
                    dir = 'asc';
                } else if (el.classList.contains('sort-asc')){
                    el.classList.remove('sort-asc');
                    el.classList.add('sort-desc');
                    dir = 'desc';
                } else {
                    el.classList.remove('sort-desc');
                    el.classList.remove('sort-asc');
                }
                if (dir)
                    this.uiEvent.notify('sort', {'dir': dir});
            });
            tr.insertAdjacentElement('beforeend', el);
        });
        return this.#base;
    }
    /**
     * Table body mount on base element
     */
    mountBody() {
        const container = this.#base.querySelector('.basic-table-body');
        container.replaceChildren();
        for (const d of this.#rows) {
            let row = `<tr class="basic-table-row">\n`;
            d.forEach((v) => {
                row += `\t<td>${v}</td>\n`;
            });
            row += '</tr>\n';
            let el = tag.buildElement(row);
            // entry
            el.addEventListener('click', (e) => {
                const now = e.currentTarget;
                if (this.#selected) {
                    this.#selected.classList.remove('selected');
                }
                this.#selected = now;
                this.#selected.classList.add('selected');
                this.index = this.#selected.rowIndex;
                this.uiEvent.notify('click_row', {});
            });
            // entry
            el.addEventListener('dblclick', (e) => {
                let idx = this.index - 1;
                this.uiEvent.notify('dblclick_row', {index: idx, row: this.#rows[idx]});
            });
            tag.appendElement(container, el);
        }
        return this.#base;
    }

    hideBody() {
        this.folding = ! this.folding;
        let target = document.querySelector('.basic-table-body');
        if (this.folding) {
            (target) ? target.hidden = true : '';
        } else {
            (target) ? target.hidden = false : '';
            (this.#selected) ? this.#selected.classList.remove('selected') : '';
        }
    }

    buildElement() {
        this.mountCaption();
        this.mountHeader();
        this.mountBody();
        return this.element = this.#base;
    }
}

/* //<-- basic-table.js ends here*/
