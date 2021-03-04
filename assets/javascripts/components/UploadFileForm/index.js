'use strict';

import { dispatchEventFromElement } from 'Helpers/event';

export const DEFAULT_SELECTOR = '#upload-file-form';

const UPLOAD_FILE_FORM_OPEN_EVENT = 'UPLOAD_FILE_FORM_OPEN';

class UploadFileForm {
    /**
     * Initializer
     *
     * @param {Element} elementRef - Upload file form element
     */
    constructor(elementRef) {
        // Root alert container
        this.elementRef = elementRef;
        this.elementFilePicker = this.elementRef.querySelector('input[type="file"]');
        this.elementActionControl = document.querySelector(this.elementRef.dataset.actionControl);

        // Bind functions
        this.onClickActionControl = this.onClickActionControl.bind(this);
        this.onFileKeywordSelected = this.onFileKeywordSelected.bind(this);
        this.onOpenUploadFileForm = this.onOpenUploadFileForm.bind(this);

        this._setup();
    }

    // Event Handlers

    /**
     * Dispatch `UPLOAD_FILE_FORM_OPEN` event when click the action control button
     */
    onClickActionControl() {
        dispatchEventFromElement(this.elementRef, UPLOAD_FILE_FORM_OPEN_EVENT);
    }

    /**
     * Submit the upload file form when select a CSV file from the browser popup
     */
    onFileKeywordSelected() {
        if (this._validFilePicker()) {
            this.elementRef.submit();
        }
    }

    /**
     * Open the browser popup when receive `UPLOAD_FILE_FORM_OPEN` event
     */
    onOpenUploadFileForm() {
        this.elementFilePicker.click();
    }

    // Private

    _setup() {
        this.elementRef.addEventListener(UPLOAD_FILE_FORM_OPEN_EVENT, this.onOpenUploadFileForm);
        this.elementFilePicker.addEventListener('change', this.onFileKeywordSelected, false);
        this.elementActionControl.addEventListener('click', this.onClickActionControl);
    }

    /**
     * Validate file picker with these conditions.
     * 1. The file size can not exceed 3MB.
     * 2. User can only pick a file.
     */
    _validFilePicker() {
        let file = this.elementFilePicker.files[0];

        return (file !== undefined && this.elementFilePicker.files.length === 1 && file.size <= (1024 * 1024 * 3))
    }
}

export default UploadFileForm;
