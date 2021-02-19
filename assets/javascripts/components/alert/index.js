'use strict';

export const DEFAULT_SELECTOR = 'div.alert';

const DEFAULT_OPTIONS = {
    hideAfter: 5000,
    hideAlertClass: 'alert--hide'
};

class Alert {
    /**
     * Initializer
     *
     * @param {Element} elementRef - Alert element
     */
    constructor(elementRef) {
        // Initialise attributes
        this.options = DEFAULT_OPTIONS;

        // Root alert container
        this.elementRef = elementRef;

        // Bind functions
        this.onHideAlert = this.onHideAlert.bind(this);

        this._setup();
    }

    /**
     * Hide alert element in screen after setTimeout is executed
     */
    onHideAlert() {
        this.elementRef.classList.add(this.options.hideAlertClass)
    }

    // Private

    _setup() {
        setTimeout(this.onHideAlert, this.options.hideAfter)
    }
}

export default Alert;
