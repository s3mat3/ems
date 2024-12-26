//import '@assets/styles/style.css'
import 'material-symbols';
//import 'chota';
import '../../assets/styles/srl.scss';

import App from '../../srl/core/app';
import BasicTable from '../../srl/components/parts/basic-table';
// import EditableDialog from '../../srl/components/parts/editable-dialog';
import MessageDialog from '../../srl/components/parts/message-dialog';

import {tableData, tableHeader, rowInfo} from './table-data'

class ShowCase extends App {
    constructor() {
        super();
        this.tableView = document.querySelector('.basic-table-wrapper');
        this.table = new BasicTable(this.tableView, {name: 'Basic Table'});
        this.table.header = tableHeader;
        this.table.rows = tableData;
        // this.modalView = document.querySelector('.editable-modal-wrapper');
        // console.log(this.modalView);
        // this.modal = new EditableModal(this.modalView, {title:'ADD', mode: 'add'});
        // this.modal.infos = rowInfo;
    }

    onLoad(e) {
        this.table.mount();
        // this.modal.mount();
        const btn = document.querySelector('#button_editable_modal_open');
        btn.addEventListener('click', (e) => {
            this.modal.open();
        })
        // danger message
        const btn_danger = document.querySelector('#message_open_danger');
        btn_danger.addEventListener('click', (e) => {
            let messageView = document.querySelector('.message-dialog-wrapper');
            let props = {
                dialog_class: 'is-danger',
                content_class: 'is-error',
                title: 'Message danger',
                message: 'This action with danger!',
            };
            let message = new MessageDialog(messageView, props);
            message.open();
        });
        // error message
        const btn_error = document.querySelector('#message_open_error');
        btn_error.addEventListener('click', (e) => {
            let messageView = document.querySelector('.message-dialog-wrapper');
            let props = {
                dialog_class: 'is-error',
                content_class: 'is-error',
                title: 'Message error',
                message: '1st error \n piyo hoge fuga',
            };
            let message = new MessageDialog(messageView, props);
            message.open();
        });
        // warning message
        const btn_warn = document.querySelector('#message_open_warning');
        btn_warn.addEventListener('click', (e) => {
            let messageView = document.querySelector('.message-dialog-wrapper');
            let props = {
                dialog_class: 'is-warning',
                content_class: 'is-error',
                title: 'Message warning',
                message: 'Warning!! \n piyo hoge fuga',
            };
            let message = new MessageDialog(messageView, props);
            message.open();
        });
        // success message
        const btn_succ = document.querySelector('#message_open_success');
        btn_succ.addEventListener('click', (e) => {
            let messageView = document.querySelector('.message-dialog-wrapper');
            let props = {
                dialog_class: 'is-success',
                content_class: 'is-error',
                title: 'Message success',
                message: 'Wao!! \n piyo hoge fuga',
            };
            let message = new MessageDialog(messageView, props);
            message.open();
        });
        // Info message
        const btn_info = document.querySelector('#message_open_info');
        btn_info.addEventListener('click', (e) => {
            let messageView = document.querySelector('.message-dialog-wrapper');
            let props = {
                dialog_class: 'is-info',
                content_class: 'is-error',
                title: 'Message info',
                message: 'Information \n piyo hoge fuga',
            };
            let message = new MessageDialog(messageView, props);
            message.open();
        });
        // praimary message
        const btn_prim = document.querySelector('#message_open_primary');
        btn_prim.addEventListener('click', (e) => {
            let messageView = document.querySelector('.message-dialog-wrapper');
            let props = {
                dialog_class: 'is-primary',
                content_class: 'is-error',
                title: 'Message primary',
                message: 'Primary??? \n piyo hoge fuga',
            };
            let message = new MessageDialog(messageView, props);
            message.open();
        });
    }
}


const app = new ShowCase();
app.entryLifeCycleHooks();
