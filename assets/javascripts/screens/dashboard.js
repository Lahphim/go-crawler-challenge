'use strict';

import UploadFileForm, { DEFAULT_SELECTOR as UPLOAD_FILE_FORM_SELECTOR } from 'Components/UploadFileForm';

const SELECTORS = {
    screen: 'body.dashboard.index'
};

class DashboardScreen {
    constructor() {
        this.uploadFileForm = document.querySelector(UPLOAD_FILE_FORM_SELECTOR);

        this._setup();
    }

    // Private

    _setup() {
        this._setupUploadFileForm();
    }

    _setupUploadFileForm() {
        if (this.uploadFileForm !== null) {
            new UploadFileForm(this.uploadFileForm);
        }
    }
}

let isDashboardScreen = document.querySelector(SELECTORS.screen) !== null;

if (isDashboardScreen) {
    new DashboardScreen();
}
