/**
 * @file app.js
 *
 * @copyright © 2024 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @brief
 *
 * @author s3mat3
 */
'use strict';

export default class App {
    constructor() {}

    entryLifeCycleHooks() {
        // 1
        document.addEventListener("DOMContentLoaded", (e) => {
           this.onDOMContentLoaded(e);
        });
        // 2
        document.addEventListener('readystatechange', (e) => {
            this.onReadyStateChange(e);
        });
        // 3
        window.addEventListener("load", (e) => {
            this.onLoad(e);
        });
        // 4 cancelable in case when leave alert
        // window.addEventListener("beforeunload", (e) => {
        //     console.log('before unload', cnt++, e);
        //     this.onBeforeunload(e);
        // });
        // 5 no cancel
        window.addEventListener("pagehide", (e) => {
            this.onPageHide(e);
        });
        // 6
        window.addEventListener("unload", (e) => {
            this.onUnLoad(e);
        });
    }
    /////////////////////////
    // life cycle handlers //
    /////////////////////////
    onReadyStateChange() {}
    onLoad() {}
    onDOMContentLoaded() {}
    onBeforeunload() {}
    onPageHide() {}
    onUnLoad() {}
}
/* //<-- app.js ends here*/
