'use strict';

export const DEFAULT_SELECTOR = '#upload-file-form';

const UPLOAD_FILE_FORM_OPEN_EVENT = 'UPLOAD_FILE_FORM_OPEN';

class UploadFileForm {
    constructor(elementRef) {
        // Root alert container
        this.elementRef = elementRef;
        this.elementActionControl = document.querySelector(this.elementRef.dataset.actionControl);

        // Bind functions
        this.onClickActionControl = this.onClickActionControl.bind(this);
        this.onOpenUploadFileForm = this.onOpenUploadFileForm.bind(this);

        this._setup();
    }

    // Event Handlers

    onClickActionControl() {
        console.log("clicking!!");
    }

    onOpenUploadFileForm() {
        console.log("open it!!");
    }

    // Private

    _setup() {
        this.elementActionControl.addEventListener('click', this.onClickActionControl);
        this.elementRef.addEventListener(UPLOAD_FILE_FORM_OPEN_EVENT, this.onOpenUploadFileForm);
    }
}

export default UploadFileForm;
