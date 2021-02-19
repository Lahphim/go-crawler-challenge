'use strict';

export const DEFAULT_SELECTOR = 'div.alert';

const TIMERS = {
    showAfter: 2000,
    hideAfter: 6000
};

const DEFAULT_OPTIONS = {
    timers: TIMERS,
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
        this.timers = DEFAULT_OPTIONS.timers;

        // Root alert container
        this.elementRef = elementRef;

        // Bind functions
        this.onShowAlert = this.onShowAlert.bind(this);
        this.onHideAlert = this.onHideAlert.bind(this);

        this._setup();
    }

    /**
     * Show alert element in screen after setTimeout is executed
     */
    onShowAlert() {
        this.elementRef.classList.remove(DEFAULT_OPTIONS.hideAlertClass)
    }

    /**
     * Hide alert element in screen after setTimeout is executed
     */
    onHideAlert() {
        this.elementRef.classList.add(DEFAULT_OPTIONS.hideAlertClass)
    }

    // Private

    _setup() {
        setTimeout(this.onShowAlert, this.timers.showAfter)
        setTimeout(this.onHideAlert, this.timers.hideAfter)
    }
}

export default Alert;
