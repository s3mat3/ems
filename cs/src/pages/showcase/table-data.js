/**
 * @file table-data.js
 *
 * @copyright © 2024 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @brief
 *
 * @author s3mat3
 */
'use strict';
export const tableHeader = [
    {name: 'ID', sortable: true},
    {name: 'Name', sortable: true},
    {name: 'Description', sortable: false},
];

export const tableData = [
    [1,  'hoge', 'fuga'],
    [2,  'hoge', 'fuga'],
    [3,  'hoge', 'fuga'],
    [4,  'hoge', 'fuga'],
    [5,  'hoge', 'fuga'],
    [6,  'hoge', 'fuga'],
    [7,  'hoge', 'fuga'],
    [8,  'hoge', 'fuga'],
    [9,  'hoge', 'fuga'],
    [10, 'hoge', 'fuga'],
    [11, 'hoge', 'fuga'],
    [12, 'hoge', 'fuga'],
    [13, 'hoge', 'fuga'],
    [14, 'hoge', 'fuga'],
    [15, 'hoge', 'fuga'],
    [16, 'hoge', 'fuga'],
    [17, 'hoge', 'fuga'],
    [18, 'hoge', 'fuga'],
    [19, 'hoge', 'fuga'],
    [20, 'hoge', 'fuga'],
];

export const rowInfo = [
    {
        feild: 'id',
        name: 'ID',
        key: true,
        required: true,
        editable: true,
        hidden: false,
        sortable: false,
        input: 'line', // line | lines | list | time | date | date-time
        list_spec: {
            name: 'datas',
            default: 1,
        }
    },
    {
        feild: 'name',
        name: 'Name',
        key: false,
        required: true,
        editable: true,
        hidden: false,
        sortable: false,
        input: 'line', // line | lines | list | time | date | date-time
        list_spec: {
            name: 'datas',
            default: 1,
        }
    },
    {
        feild: 'status',
        name: 'Status',
        key: false,
        required: true,
        editable: true,
        hidden: false,
        sortable: false,
        input: 'list', // line | lines | list | time | date | date-time
        list_spec: {
            name: 'datas',
            default: 1,
        }
    }
];
/* //<-- table-data.js ends here*/
